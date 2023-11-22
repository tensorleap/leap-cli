package storage

import "fmt"

// MockFileSystem is a simple in-memory virtual file system.
type MockFileSystem map[string][]byte

// MockLockFileApi is a mocked implementation of LockFileApi using MockFileSystem.
type MockLockFileApi struct {
	FileSystem MockFileSystem
}

func (m *MockLockFileApi) Exists(file string) (bool, error) {
	_, ok := m.FileSystem[file]
	return ok, nil
}

func (m *MockLockFileApi) Create(file string, b []byte) error {
	m.FileSystem[file] = b
	return nil
}

func (m *MockLockFileApi) Read(file string) ([]byte, error) {
	content, ok := m.FileSystem[file]
	if !ok {
		return nil, fmt.Errorf("file not found")
	}
	return content, nil
}

func (m *MockLockFileApi) Update(file string, b []byte) error {
	m.FileSystem[file] = b
	return nil
}

func (m *MockLockFileApi) Delete(file string) error {
	delete(m.FileSystem, file)
	return nil
}
