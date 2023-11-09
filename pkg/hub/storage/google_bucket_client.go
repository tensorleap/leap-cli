package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/storage"
	"github.com/tensorleap/leap-cli/pkg/log"
	"google.golang.org/api/iterator"
)

const (
	GoogleBucketNamePrefix = "project-hub-"
	GoogleProjectID        = "tensorleap-ops3"
)

type GoogleBucketClient struct {
	Name   string
	bucket *storage.BucketHandle
}

func NewGoogleBucketClient(bucket *storage.BucketHandle, name string) *GoogleBucketClient {
	return &GoogleBucketClient{
		Name:   name,
		bucket: bucket,
	}
}

func (c *GoogleBucketClient) CheckIfFileExists(fileName string) (bool, error) {
	ctx := context.Background()
	_, err := c.bucket.Object(fileName).Attrs(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			return false, nil
		}
		return false, fmt.Errorf("failed to check if file exists: %v", err)
	}
	return true, nil
}

func (c *GoogleBucketClient) GetFileBuffer(fileName string) ([]byte, error) {
	ctx := context.Background()
	rc, err := c.bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *GoogleBucketClient) UploadFile(fileName string, content io.Reader, options *UploadOptions) error {
	ctx := context.Background()
	wc := c.bucket.Object(fileName).NewWriter(ctx)
	defer wc.Close()

	if options != nil && options.NoCache {
		wc.CacheControl = "no-store"
	}
	if _, err := io.Copy(wc, content); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	log.Printf("File (%s) uploaded successfully.\n", fileName)
	return nil
}

func (c *GoogleBucketClient) UploadFileBuffer(fileName string, content []byte, options *UploadOptions) error {
	ctx := context.Background()
	wc := c.bucket.Object(fileName).NewWriter(ctx)
	defer wc.Close()
	if options != nil && options.NoCache {
		wc.CacheControl = "no-store"
	}
	if _, err := wc.Write(content); err != nil {
		return err
	}
	log.Printf("File (%s) uploaded successfully.\n", fileName)
	return nil
}

func (c *GoogleBucketClient) DownloadFile(fileName string, destFileName string) error {
	ctx := context.Background()
	rc, err := c.bucket.Object(fileName).NewReader(ctx)
	if err != nil {
		return err
	}
	defer rc.Close()

	data, err := io.ReadAll(rc)
	if err != nil {
		return err
	}

	err = os.WriteFile(destFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (c *GoogleBucketClient) GetFileMeta(fileName string) (*FileMeta, error) {
	ctx := context.Background()
	attrs, err := c.bucket.Object(fileName).Attrs(ctx)
	if err != nil {
		return nil, err
	}

	return &FileMeta{
		Name:        attrs.Name,
		Size:        attrs.Size,
		CreatedTime: attrs.Created,
		UpdatedTime: attrs.Updated,
	}, nil
}

func (c *GoogleBucketClient) DeleteFile(fileName string) error {
	ctx := context.Background()
	if err := c.bucket.Object(fileName).Delete(ctx); err != nil {
		return err
	}

	log.Printf("File %s deleted successfully.\n", fileName)
	return nil
}

func (c *GoogleBucketClient) DeleteFilesInDirectory(directoryPath string) (bool, error) {
	ctx := context.Background()
	files, err := c.ListDirectoryObjects(directoryPath, true)
	if err != nil || len(files) == 0 {
		return false, err
	}

	for _, file := range files {
		log.Infof("Deleting file: %s\n", file)
		if err := c.bucket.Object(file).Delete(ctx); err != nil {
			return false, err
		}
	}

	log.Infof("Files in directory %s deleted successfully.\n", directoryPath)
	return true, nil
}

func (c *GoogleBucketClient) ListDirectoryObjects(directoryPath string, recursive bool) ([]string, error) {
	ctx := context.Background()
	normalizedPath := directoryPath
	if !strings.HasSuffix(directoryPath, "/") {
		normalizedPath = directoryPath + "/"
	}
	query := &storage.Query{Prefix: normalizedPath}
	it := c.bucket.Objects(ctx, query)
	var objects []string
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list directory objects: %v", err)
		}

		objects = append(objects, attrs.Name)
		if recursive && strings.HasSuffix(attrs.Name, "/") {
			subObjects, err := c.ListDirectoryObjects(attrs.Name, recursive)
			if err != nil {
				return nil, err
			}
			objects = append(objects, subObjects...)
		}
	}
	return objects, nil
}

func (c *GoogleBucketClient) CreateWriteableSignedUrl(fileName string, expire time.Time) (string, error) {
	opts := &storage.SignedURLOptions{
		Scheme:  storage.SigningSchemeV4,
		Method:  "PUT",
		Expires: expire,
	}

	u, err := c.bucket.SignedURL(fileName, opts)
	if err != nil {
		return "", err
	}
	return u, nil
}

func CreateStorageClient() (StorageClient, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create storage client: %v", err)
	}

	it := client.Buckets(ctx, GoogleProjectID)
	for {
		bucketAttrs, err := it.Next()
		if err == storage.ErrBucketNotExist {
			return nil, fmt.Errorf("bucket not found")
		}
		if err != nil {
			return nil, err
		}

		if strings.HasPrefix(bucketAttrs.Name, GoogleBucketNamePrefix) {
			return NewGoogleBucketClient(client.Bucket(bucketAttrs.Name), bucketAttrs.Name), nil
		}
	}
}
