package hub

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/tensorleap/leap-cli/pkg/log"
)

type Image struct {
	Name   string
	Buffer []byte
}

func ExtractProjectContextFromTar(tarStream io.Reader) (*ProjectContext, error) {
	gz, err := gzip.NewReader(tarStream)
	if err != nil {
		return nil, err
	}
	defer func() { _ = gz.Close() }()

	tarReader := tar.NewReader(gz)

	var projectMeta ProjectMeta
	var projectBgImage Image
	foundFilesCount := 0

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		if header.Typeflag == tar.TypeReg {
			if strings.HasSuffix(header.Name, "project.json") {
				foundFilesCount++
				err := json.NewDecoder(tarReader).Decode(&projectMeta)
				if err != nil {
					return nil, err
				}
			} else if strings.Contains(header.Name, "bgImage") {
				foundFilesCount++
				buf := new(bytes.Buffer)
				_, err = buf.ReadFrom(tarReader)
				if err != nil {
					return nil, err
				}
				projectBgImage = Image{
					Name:   header.Name,
					Buffer: buf.Bytes(),
				}
			}

			if foundFilesCount == 2 {
				break
			}
		}
	}

	if projectMeta.Name == "" {
		return nil, errors.New("invalid tar file, can't found project.json")
	}

	if len(projectBgImage.Buffer) == 0 {
		return nil, errors.New("invalid tar file, can't found bgImage*.[jpg|png]. make sure you add background image on the project")
	}

	log.Infof("Exported project context from tar file")

	return &ProjectContext{
		Meta:    projectMeta,
		BgImage: projectBgImage,
	}, nil
}
