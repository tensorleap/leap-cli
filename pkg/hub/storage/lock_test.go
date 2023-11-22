package storage

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestLock_CheckLock(t *testing.T) {
	// Set up a mock file system
	fileSystem := make(MockFileSystem)

	// Set up the Lock with the mock file system
	lock := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner")

	// Test when the lock file doesn't exist
	lockInfo, err := lock.CheckLock()
	assert.Nil(t, err)
	assert.False(t, lockInfo.IsLocked)

	// Create a sample lock
	lockInfoJSON, _ := json.Marshal(&LockInfo{
		IsLocked:      true,
		LockOwner:     "owner",
		ExpirationUTC: time.Now().UTC().Add(1 * time.Hour),
	})
	fileSystem["test/file.lock"] = lockInfoJSON

	// Test when the lock file exists
	lockInfo, err = lock.CheckLock()
	assert.Nil(t, err)
	assert.True(t, lockInfo.IsLocked)
	assert.Equal(t, "owner", lockInfo.LockOwner)
}

func TestLock_Lock(t *testing.T) {
	// Set up a mock file system
	fileSystem := make(MockFileSystem)

	// Set up the Lock with the mock file system
	lock := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner")

	// Test acquiring a lock
	expiration := time.Now().UTC().Add(1 * time.Hour)
	err := lock.Lock(expiration, false)
	assert.Nil(t, err)

	// Check that the lock file exists and has the correct content
	lockInfoJSON, ok := fileSystem["test/file.lock"]
	assert.True(t, ok)

	var lockInfo LockInfo
	err = json.Unmarshal(lockInfoJSON, &lockInfo)
	assert.Nil(t, err)

	assert.True(t, lockInfo.IsLocked)
	assert.Equal(t, "owner", lockInfo.LockOwner)
	assert.Equal(t, expiration, lockInfo.ExpirationUTC)
}

func TestLock_WhenAlreadyLocked(t *testing.T) {
	// Set up a mock file system
	fileSystem := make(MockFileSystem)

	// Set up the Lock with the mock file system
	lock1 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner1")

	// Test acquiring a lock
	expiration := time.Now().UTC().Add(1 * time.Hour)
	err := lock1.Lock(expiration, false)
	assert.Nil(t, err)

	lock2 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner2")

	// Test acquiring a lock when the file is already locked
	err = lock2.Lock(expiration, false)
	assert.NotNil(t, err)
}

func TestLock_WhenAlreadyLockedAndOverride(t *testing.T) {
	// Set up a mock file system
	fileSystem := make(MockFileSystem)

	// Set up the Lock with the mock file system
	lock1 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner1")

	// Test acquiring a lock
	expiration := time.Now().UTC().Add(1 * time.Hour)
	err := lock1.Lock(expiration, false)
	assert.Nil(t, err)

	lock2 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner2")

	err = lock2.Lock(expiration, true)
	assert.Nil(t, err)

	// Check that the lock file exists and has the correct content
	lockInfo, err := lock1.CheckLock()
	assert.Nil(t, err)
	assert.True(t, lockInfo.IsLocked)
	assert.Equal(t, "owner2", lockInfo.LockOwner)

	isLock1Locked, err := lock1.Locked()
	assert.Nil(t, err)
	assert.False(t, isLock1Locked)
}

func TestLock_WaitAndLock(t *testing.T) {
	// Set up a mock file system
	fileSystem := make(MockFileSystem)

	// Set up the Lock with the mock file system
	lock1 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner1")

	// Test acquiring a lock
	expiration := time.Now().UTC().Add(5 * time.Second)
	err := lock1.Lock(expiration, false)
	assert.Nil(t, err)

	lock2 := NewLock("test/file.lock", &MockLockFileApi{FileSystem: fileSystem}, "owner2")

	go func() {
		time.Sleep(2 * time.Second)
		err := lock1.ReleaseLock()
		assert.Nil(t, err)
	}()

	// Test acquiring a lock when the file is already locked
	err = lock2.WaitAndLock(1*time.Hour, 3*time.Second)
	assert.Nil(t, err)

	// Check that the lock file exists and has the correct content
	isLock1Locked, err := lock1.Locked()
	assert.Nil(t, err)
	assert.False(t, isLock1Locked)

	isLock2Locked, err := lock2.Locked()
	assert.Nil(t, err)
	assert.True(t, isLock2Locked)
}
