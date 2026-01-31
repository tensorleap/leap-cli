package root_cmd

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/code"
	"github.com/tensorleap/leap-cli/pkg/local"
	"github.com/tensorleap/leap-cli/pkg/log"
	"github.com/tensorleap/leap-cli/pkg/model"
	"github.com/tensorleap/leap-cli/pkg/project"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
	"github.com/tensorleap/leap-cli/pkg/workspace"
)

type VersionWithCode struct {
	VersionId      string
	VersionName    string
	CodeSnapshotId string
	Status         model.VersionStatus
	UpdatedAt      time.Time
}

func NewPullCmd() *cobra.Command {
	var outputDir string
	var openAfter bool
	var tempDir bool
	var archiveOnly bool

	var cmd = &cobra.Command{
		Use:   "pull",
		Short: "Pull version code",
		Long: `Pull version code and save it to a local directory.

Examples:
  # Pull code to current directory
  leap pull

  # Pull and open the directory
  leap pull -o

  # Pull to a temp directory and open it
  leap pull -to

  # Download as archive (tar.gz) without extracting
  leap pull -a
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := auth.RequireAuthSimple(cmd.Context()); err != nil {
				return err
			}
			ctx := cmd.Context()

			var projectId string
			workspaceConfig, err := workspace.GetWorkspaceConfig()
			if err == nil && workspaceConfig != nil {
				projectId = workspaceConfig.ProjectId
			}

			currentProject, _, err := project.GetProjectFromProjectId(ctx, projectId, false)
			if err != nil {
				return err
			}

			selectedVersion, err := askUserToSelectVersionWithCode(ctx, currentProject.GetCid())
			if err != nil {
				return err
			}

			if selectedVersion == nil {
				return fmt.Errorf("no version selected")
			}

			if selectedVersion.CodeSnapshotId == "" {
				return fmt.Errorf("selected version does not have code")
			}

			codeSnapshot, err := code.GetCodeSnapshot(ctx, currentProject.GetCid(), selectedVersion.CodeSnapshotId)
			if err != nil {
				return fmt.Errorf("failed to get code snapshot: %w", err)
			}

			dateStr := selectedVersion.UpdatedAt.Local().Format("2006-01-02_15-04")
			baseName := local.SanitizeDirectoryName(fmt.Sprintf("%s_%s", selectedVersion.VersionName, dateStr))

			if archiveOnly {
				var targetPath string
				if tempDir {
					hash := generateShortHash()
					targetPath = filepath.Join(os.TempDir(), fmt.Sprintf("%s_%s", baseName, hash))
				} else if outputDir != "" {
					targetPath = outputDir
				} else {
					targetPath = baseName
				}

				if _, err := os.Stat(targetPath); !os.IsNotExist(err) && !tempDir {
					overwrite := false
					prompt := &survey.Confirm{
						Message: fmt.Sprintf("File '%s' already exists. Do you want to overwrite it?", targetPath),
						Default: false,
					}
					err := survey.AskOne(prompt, &overwrite)
					if err != nil {
						return err
					}
					if !overwrite {
						log.Info("Pull cancelled")
						return nil
					}
				}

				log.Infof("Downloading archive for version '%s'...", selectedVersion.VersionName)
				err = code.DownloadCodeSnapshotArchive(ctx, codeSnapshot, targetPath)
				if err != nil {
					return fmt.Errorf("failed to download archive: %w", err)
				}

				log.Infof("Successfully downloaded archive to: %s", targetPath)
			} else {
				var targetPath string
				if tempDir {
					hash := generateShortHash()
					targetPath = filepath.Join(os.TempDir(), fmt.Sprintf("%s_%s", baseName, hash))
				} else if outputDir != "" {
					targetPath = outputDir
				} else {
					targetPath = baseName
				}

				if _, err := os.Stat(targetPath); !os.IsNotExist(err) && !tempDir {
					overwrite := false
					prompt := &survey.Confirm{
						Message: fmt.Sprintf("Directory '%s' already exists. Do you want to overwrite it?", targetPath),
						Default: false,
					}
					err := survey.AskOne(prompt, &overwrite)
					if err != nil {
						return err
					}
					if !overwrite {
						log.Info("Pull cancelled")
						return nil
					}
					if err := os.RemoveAll(targetPath); err != nil {
						return fmt.Errorf("failed to remove existing directory: %w", err)
					}
				}

				if err := os.MkdirAll(targetPath, 0755); err != nil {
					return fmt.Errorf("failed to create output directory: %w", err)
				}

				log.Infof("Pulling code for version '%s'...", selectedVersion.VersionName)
				files, err := code.CloneCodeSnapshot(ctx, codeSnapshot, targetPath, "")
				if err != nil {
					return fmt.Errorf("failed to download code: %w", err)
				}

				log.Infof("Successfully pulled %d files to: %s", len(files), targetPath)

				if openAfter {
					if err := local.OpenDirectorySmart(targetPath); err != nil {
						log.Warnf("Failed to open directory: %v", err)
					}
				}
			}

			return nil
		},
	}

	cmd.Flags().StringVarP(&outputDir, "output", "O", "", "Output path (default: version name and updated date)")
	cmd.Flags().BoolVarP(&openAfter, "open", "o", false, "Open the directory/location after downloading")
	cmd.Flags().BoolVarP(&tempDir, "temp", "t", false, "Save to temp directory")
	cmd.Flags().BoolVarP(&archiveOnly, "archive", "a", false, "Download as tar.gz archive without extracting")

	return cmd
}

func generateShortHash() string {
	bytes := make([]byte, 4)
	_, _ = rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func askUserToSelectVersionWithCode(ctx context.Context, projectId string) (*VersionWithCode, error) {
	versions, err := model.GetVersions(ctx, projectId)
	if err != nil {
		return nil, err
	}

	if len(versions) == 0 {
		return nil, fmt.Errorf("no versions found in project")
	}

	versionsWithCode := []tensorleapapi.SlimVersion{}
	for _, version := range versions {
		if version.HasCodeSnapshotId() {
			versionsWithCode = append(versionsWithCode, version)
		}
	}

	if len(versionsWithCode) == 0 {
		return nil, fmt.Errorf("no versions with code found in project")
	}

	runsStatusesPerVersionId, err := model.GetRunsStatusesPerVersionId(ctx, projectId)
	if err != nil {
		return nil, err
	}

	maxLengthOfVersionName := 0
	for _, version := range versionsWithCode {
		if len(version.GetNotes()) > maxLengthOfVersionName {
			maxLengthOfVersionName = len(version.GetNotes())
		}
	}

	versionInfos := []VersionWithCode{}
	options := []string{}

	for _, version := range versionsWithCode {
		status, _ := model.CalcVersionStatus(&version, runsStatusesPerVersionId[version.GetCid()])
		displayName := model.FormatVersionDisplayName(&version, status, maxLengthOfVersionName)
		options = append(options, displayName)

		versionInfos = append(versionInfos, VersionWithCode{
			VersionId:      version.GetCid(),
			VersionName:    version.GetNotes(),
			CodeSnapshotId: version.GetCodeSnapshotId(),
			Status:         status,
			UpdatedAt:      version.GetUpdatedAt(),
		})
	}

	selectedIndex := 0
	prompt := &survey.Select{
		Message: "Select version to pull code from",
		Options: options,
		Default: selectedIndex,
	}
	err = survey.AskOne(prompt, &selectedIndex)
	if err != nil {
		return nil, err
	}

	return &versionInfos[selectedIndex], nil
}
