package run

import (
	"fmt"
	"slices"
	"strings"

	"github.com/samber/lo"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"

	runPkg "github.com/tensorleap/leap-cli/pkg/run"
)

func NewListCmd() *cobra.Command {

	var types []string
	var statuses []string

	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"ls"},
		Short:   "List the runs",
		RunE: func(cmd *cobra.Command, args []string) error {
			types, err := convertAndValidateStringToEnum(types, tensorleapapi.AllowedJobSubTypeEnumValues)
			if err != nil {
				return err
			}
			statuses, err := convertAndValidateStringToEnum(statuses, tensorleapapi.AllowedJobStatusEnumValues)
			if err != nil {
				return err
			}
			runs, err := runPkg.GetRuns(cmd.Context(), types, statuses)
			if err != nil {
				return err
			}
			entity.PrintList(runs, runPkg.RunEntityDesc)
			return nil
		},
	}

	typesAsStrings := convertArrayToStrings(tensorleapapi.AllowedJobSubTypeEnumValues)
	statusesAsStrings := convertArrayToStrings(tensorleapapi.AllowedJobStatusEnumValues)
	allowedTypes := (strings.Join(typesAsStrings, ", "))
	allowedStatuses := (strings.Join(statusesAsStrings, ", "))
	cmd.Flags().StringSliceVarP(&types, "types", "t", typesAsStrings, fmt.Sprintf("Filter by types (%s)", allowedTypes))
	cmd.Flags().StringSliceVarP(&statuses, "statuses", "s", statusesAsStrings, fmt.Sprintf("Filter by statuses (%s)", allowedStatuses))

	return cmd
}

func init() {
	RootCommand.AddCommand(NewListCmd())
}

func convertArrayToStrings[T ~string](list []T) []string {
	return lo.Map(list, func(item T, _ int) string {
		return string(item)
	})
}

func convertAndValidateStringToEnum[T ~string](list []string, enum []T) ([]T, error) {
	enumAsStrings := convertArrayToStrings(enum)
	for _, item := range list {
		if !slices.Contains(enumAsStrings, item) {
			return nil, fmt.Errorf("invalid value '%s' for %T: valid values are %v", item, enum, strings.Join(enumAsStrings, ", "))
		}
	}
	return lo.Map(list, func(item string, _ int) T {
		return T(item)
	}), nil
}
