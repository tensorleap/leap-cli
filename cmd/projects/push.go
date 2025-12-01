package projects

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func NewPushCmd() *cobra.Command {
	var secretId string
	var modelVersionName string
	var codeVersionMessage string
	var modelType string
	var modelPath string
	var branch string
	var transformInput bool
	var noWait bool
	var pythonVersion string

	var cmd = &cobra.Command{
		Use:   "push [modelPath]",
		Short: "Push new or overwrite model version (DEPRECATED)",
		Long:  `Push new or overwrite model version into a project with its model and code integration (DEPRECATED)`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Show deprecation warning
			log.Warn("⚠️  The command 'leap projects push' is deprecated and will be removed in a future version.")
			log.Warn("Please use 'leap push' instead.")
			fmt.Println()

			return fmt.Errorf("please use 'leap push' instead of 'leap projects push'")
		},
	}

	// Add the same flags as the new push command (matching the old projects push interface)
	cmd.Flags().StringVarP(&modelVersionName, "model-name", "m", "", "Model version name")
	cmd.Flags().StringVar(&modelType, "type", "", "Type is the type of the model file [JSON_TF2 / ONNX / PB_TF2 / H5_TF2]")
	cmd.Flags().StringVar(&branch, "code-branch", "", "Name of the code branch [OPTIONAL]")
	cmd.Flags().StringVar(&secretId, "secretId", "", "Secret id")
	cmd.Flags().BoolVar(&transformInput, "transform-input", false, "Transpose the input data to channel-last format")
	cmd.Flags().BoolVar(&noWait, "no-wait", false, "Do not wait for push to complete")
	cmd.Flags().StringVarP(&modelPath, "model-path", "", "", "Path to the model file")
	cmd.Flags().StringVar(&codeVersionMessage, "code-message", "", "Code version message")
	cmd.Flags().StringVar(&pythonVersion, "python-version", "", "Python version")

	// Mark command as deprecated
	cmd.Deprecated = "use 'leap push' instead"

	return cmd
}

func init() {
	RootCommand.AddCommand(NewPushCmd())
}
