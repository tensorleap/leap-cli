package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

type LicenseFlag struct {
	FilePath     string
	LicenseToken string
}

func NewLicenseFlag() *LicenseFlag {
	return &LicenseFlag{
		FilePath:     "",
		LicenseToken: "",
	}
}

const FILE_FLAG_MSG = "Path to the license file"
const TOKEN_FLAG_MSG = "License token"

func (f *LicenseFlag) AddShortFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&f.FilePath, "file", "f", "", FILE_FLAG_MSG)
	cmd.Flags().StringVarP(&f.LicenseToken, "token", "t", "", TOKEN_FLAG_MSG)
}

func (f *LicenseFlag) AddFlags(cmd *cobra.Command) {
	cmd.Flags().StringVar(&f.FilePath, "license-file", "", FILE_FLAG_MSG)
	cmd.Flags().StringVar(&f.LicenseToken, "license-token", "", TOKEN_FLAG_MSG)
}

func (f *LicenseFlag) PrepareLicenseToken() (string, error) {
	var tokenContent string
	if f.LicenseToken != "" {
		tokenContent = f.LicenseToken
	} else if f.FilePath != "" {
		fileBytes, err := os.ReadFile(filepath.Clean(f.FilePath))
		if err != nil {
			return "", err
		}
		tokenContent = string(fileBytes)
	} else {
		return "", fmt.Errorf("license token or file path is required")
	}
	return strings.TrimSpace(tokenContent), nil
}

func (f *LicenseFlag) HasLicense() bool {
	return f.LicenseToken != "" || f.FilePath != ""
}
