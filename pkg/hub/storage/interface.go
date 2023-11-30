package storage

import (
	"io"
	"time"
)

type UploadOptions struct {
	NoCache bool
}

type FileMeta struct {
	Name        string
	Size        int64
	CreatedTime time.Time
	UpdatedTime time.Time
}

type StorageClient interface {
	CheckIfFileExists(fileName string) (bool, error)
	GetFileBuffer(fileName string) ([]byte, error)
	UploadFile(fileName string, content io.Reader, options *UploadOptions) error
	UploadFileBuffer(fileName string, content []byte, options *UploadOptions) error
	DownloadFile(fileName string, destFileName string) error
	GetFileMeta(fileName string) (*FileMeta, error)
	DeleteFile(fileName string) error
	DeleteFilesInDirectory(directoryPath string) (bool, error)
	ListDirectoryObjects(directoryPath string, recursive bool) ([]string, error)
	CreateSignedUrl(fileName, method string, expire time.Time) (string, error)
}
