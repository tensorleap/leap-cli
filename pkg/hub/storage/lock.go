package storage

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/tensorleap/leap-cli/pkg/log"
)

// Lock represents a lock on a file.
type Lock struct {
	lockFilePath string // API URL for file operations
	fileApi      LockFileApi
	owner        string // Owner of the lock
}

func (l *Lock) String() string {
	return fmt.Sprintf("owner: %s, file: %s", l.lockFilePath, l.owner)
}

// LockInfo represents the information associated with a lock.
type LockInfo struct {
	IsLocked      bool      `json:"is_locked"`
	LockOwner     string    `json:"lock_owner"`
	ExpirationUTC time.Time `json:"expiration_utc"`
}

func (l *LockInfo) StillLocked() bool {
	return l.IsLocked && time.Now().UTC().Before(l.ExpirationUTC)
}

func (l *LockInfo) StillLockedBy(owner string) bool {
	return l.StillLocked() && l.LockOwner == owner
}

type LockFileApi interface {
	Exists(file string) (bool, error)
	Create(file string, b []byte) error
	Read(file string) ([]byte, error)
	Update(file string, b []byte) error
	Delete(file string) error
}

// NewLock creates a new Lock instance with the given API URL.
func NewLock(lockFilePath string, api LockFileApi, owner string) *Lock {
	return &Lock{
		lockFilePath: lockFilePath,
		fileApi:      api,
		owner:        owner,
	}
}

// CheckLock checks if a lock exists and returns the lock information.
func (l *Lock) CheckLock() (*LockInfo, error) {

	isExist, err := l.fileApi.Exists(l.lockFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to check lock (%s): %s", l, err)
	}

	if !isExist {
		// File doesn't exist, so there's no lock
		return &LockInfo{IsLocked: false}, nil
	}
	content, err := l.fileApi.Read(l.lockFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to check lock (%s): %s", l, err)
	}
	var lockInfo LockInfo
	if err := json.Unmarshal(content, &lockInfo); err != nil {
		return nil, fmt.Errorf("failed to check lock (%s): %s", l, err)
	}

	return &lockInfo, nil
}

// Lock acquires a lock on the file.
func (l *Lock) Lock(expiration time.Time, override bool) error {
	lockInfo, err := l.CheckLock()
	if err != nil {
		return err
	}

	if lockInfo.IsLocked && !override {
		return fmt.Errorf("file is already locked by %s until %v", lockInfo.LockOwner, lockInfo.ExpirationUTC)
	}

	// Create or update the lock
	lockInfo.IsLocked = true
	lockInfo.LockOwner = l.owner
	lockInfo.ExpirationUTC = expiration

	// Send a request to the API to create or update the lock
	lockInfoJSON, err := json.Marshal(lockInfo)
	if err != nil {
		return err
	}

	err = l.fileApi.Create(l.lockFilePath, lockInfoJSON)
	if err != nil {
		return fmt.Errorf("failed to acquire lock (%s): %s", l, err)
	}

	log.Println("Lock acquired.")
	return nil
}

// Locked checks if the file is still locked.
func (l *Lock) Locked() (bool, error) {
	lockInfo, err := l.CheckLock()
	if err != nil {
		return false, err
	}

	if lockInfo.StillLockedBy(l.owner) {
		return true, nil
	}

	return false, nil
}

func (l *Lock) WaitAndLock(lockTime time.Duration, timeout time.Duration) error {
	intervalDuration := 1 * time.Second
	expiredWithoutIntervalDuration := time.Now().UTC().Add(timeout - intervalDuration)
	for {
		lockInfo, err := l.CheckLock()
		if err != nil {
			return err
		}
		if !lockInfo.StillLocked() {
			break
		}
		if time.Now().UTC().After(expiredWithoutIntervalDuration) {
			return fmt.Errorf("lock (%s) timeout", l)
		}
		time.Sleep(intervalDuration)
	}

	return l.Lock(time.Now().UTC().Add(lockTime), true)
}

// ReleaseLock releases the lock on the file.
func (l *Lock) ReleaseLock() error {

	lockInfo, err := l.CheckLock()

	if err != nil {
		return err
	}
	if !lockInfo.IsLocked {
		return fmt.Errorf("file is not locked")
	}
	if lockInfo.IsLocked && lockInfo.LockOwner != l.owner {
		return fmt.Errorf("file is locked by %s", lockInfo.LockOwner)
	}

	err = l.fileApi.Delete(l.lockFilePath)
	if err != nil {
		return fmt.Errorf("failed to release lock: %s", err)
	}

	fmt.Println("Lock released.")
	return nil
}

// ExtendExpiration extends the expiration date of the lock.
func (l *Lock) ExtendExpiration(newExpiration time.Time) error {
	lockInfo, err := l.CheckLock()
	if err != nil {
		return err
	}

	if !lockInfo.IsLocked {
		return fmt.Errorf("file is not locked")
	}

	// Update the expiration date
	lockInfo.ExpirationUTC = newExpiration

	// Send a request to the API to update the lock expiration date
	lockInfoJSON, err := json.Marshal(lockInfo)
	if err != nil {
		return err
	}

	err = l.fileApi.Update(l.lockFilePath, lockInfoJSON)
	if err != nil {
		return fmt.Errorf("failed to extend lock expiration: %s", err)
	}

	log.Println("Lock expiration extended.")
	return nil
}

type warpStorageClientAsLockApi struct {
	StorageClient
}

func (s *warpStorageClientAsLockApi) Exists(file string) (bool, error) {
	return s.StorageClient.CheckIfFileExists(file)
}

func (s *warpStorageClientAsLockApi) Create(file string, b []byte) error {
	return s.StorageClient.UploadFileBuffer(file, b, nil)
}

func (s *warpStorageClientAsLockApi) Read(file string) ([]byte, error) {
	return s.StorageClient.GetFileBuffer(file)
}

func (s *warpStorageClientAsLockApi) Delete(file string) error {
	return s.StorageClient.DeleteFile(file)
}

func (s *warpStorageClientAsLockApi) Update(file string, b []byte) error {
	return s.StorageClient.UploadFileBuffer(file, b, nil)
}

func BuildLockApiFromStorageClient(client StorageClient) LockFileApi {
	return &warpStorageClientAsLockApi{StorageClient: client}
}
