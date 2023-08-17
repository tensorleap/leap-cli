package airgap

import (
	"archive/tar"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tensorleap/leap-cli/pkg/docker"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/server/manifest"
	"gopkg.in/yaml.v3"
)

func Pack(mnf *manifest.InstallationManifest, outputFile io.Writer) error {

	tarWriter := tar.NewWriter(outputFile)
	defer tarWriter.Close()

	err := AddManifest(tarWriter, mnf)
	if err != nil {
		return err
	}

	err = AddImages(tarWriter, mnf)
	if err != nil {
		return err
	}

	err = AddHelm(tarWriter, mnf.ServerHelmChart, HELM_FILE_NAME)
	if err != nil {
		return err
	}
	return nil
}

func AddHelm(tarWriter *tar.Writer, chartMeta manifest.HelmChartMeta, fileName string) error {
	log.Info("Downloading helm chart...")
	version, err := manifest.GetHelmVersion(chartMeta.RepoUrl, chartMeta.ChartName, chartMeta.Version)
	if err != nil {
		return err
	}

	tempHelmFile, clean, err := DownloadIntoTempFile(version.URLs[0], "helm.tgz")
	if err != nil {
		return err
	}
	defer clean()

	tempHelmFileStat, err := tempHelmFile.Stat()
	if err != nil {
		return err
	}
	helmHeader, err := tar.FileInfoHeader(tempHelmFileStat, tempHelmFileStat.Name())
	if err != nil {
		return err
	}
	helmHeader.Name = fileName
	err = tarWriter.WriteHeader(helmHeader)
	if err != nil {
		return err
	}
	_, err = tempHelmFile.Seek(0, 0)
	if err != nil {
		return err
	}
	_, err = io.Copy(tarWriter, tempHelmFile)
	if err != nil {
		return err
	}
	log.Info("Downloaded helm chart")
	return nil
}

func AddImages(tarWriter *tar.Writer, mnf *manifest.InstallationManifest) error {
	dockerClient, err := docker.NewClient()
	if err != nil {
		return err
	}

	images := mnf.GetAllImages()

	// create temp file to store images
	// we can't stream the images directly to the tar writer, because we need to know the size of the images file before writing it to the tar
	tempImagesFile, err := os.CreateTemp("", "images.tgz")
	defer local.CleanupTempFile(tempImagesFile)
	if err != nil {
		return err
	}

	err = docker.DownloadDockerImages(dockerClient, images, tempImagesFile)

	if err != nil {
		return err
	}

	tempImagesFileStat, err := tempImagesFile.Stat()
	if err != nil {
		return err
	}
	imagesHeader, err := tar.FileInfoHeader(tempImagesFileStat, tempImagesFileStat.Name())
	if err != nil {
		return err
	}

	imagesHeader.Name = IMAGES_FILE_NAME
	err = tarWriter.WriteHeader(imagesHeader)
	if err != nil {
		return err
	}
	_, err = tempImagesFile.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = io.Copy(tarWriter, tempImagesFile)
	if err != nil {
		return err
	}
	return nil
}

func AddManifest(tarWriter *tar.Writer, mnf *manifest.InstallationManifest) error {

	manifestBytes, err := yaml.Marshal(*mnf)
	if err != nil {
		return fmt.Errorf("failed to marshal manifest: %v", err)
	}
	manifestHeader := &tar.Header{
		Name: MANIFEST_FILE_NAME,
		Mode: 0600,
		Size: int64(len(manifestBytes)),
	}

	err = tarWriter.WriteHeader(manifestHeader)
	if err != nil {
		return err
	}
	_, err = tarWriter.Write(manifestBytes)
	if err != nil {
		return err
	}
	return nil
}

func DownloadIntoTempFile(url, tempSuffix string) (*os.File, func(), error) {
	tempFile, err := os.CreateTemp("", tempSuffix)
	if err != nil {
		return nil, nil, err
	}
	clean := func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}

	res, err := http.Get(url)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, nil, fmt.Errorf("failed downloading (%s): %v", url, res.StatusCode)
	}
	_, err = io.Copy(tempFile, res.Body)
	if err != nil {
		return nil, nil, err
	}
	_, err = tempFile.Seek(0, 0)
	if err != nil {
		return nil, nil, err
	}

	return tempFile, clean, nil
}
