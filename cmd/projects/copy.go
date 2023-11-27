package projects

import (
	"context"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/tensorleap/leap-cli/pkg/auth"
	"github.com/tensorleap/leap-cli/pkg/entity"
	"github.com/tensorleap/leap-cli/pkg/project"
)

func NewCopyCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "copy [env:source-name] [env:target-name]",
		Aliases: []string{"cp"},
		Short:   "Copy a project",
		RunE: func(cmd *cobra.Command, args []string) error {

			ctx := cmd.Context()
			sourceCtx, targetCtx, sourceProject, targetProjectName, err := initProjectCopyArgs(ctx, args)
			if err != nil {
				return err
			}

			err = project.CopyProject(sourceCtx, sourceProject, targetCtx, targetProjectName)

			if err != nil {
				return err
			}
			return nil
		},
	}

	return cmd
}

func initProjectCopyArgs(ctx context.Context, args []string) (
	sourceCtx, targetCtx context.Context,
	sourceProject *project.ProjectEntity,
	targetProjectName string,
	err error,
) {
	envs := auth.GetEnvs()
	if len(envs) == 0 {
		err = fmt.Errorf("no environments found. Please login an environment first")
		return
	}
	var sourceEnv, targetEnv *auth.Env
	var sourceProjectArg, targetProjectArg string
	var sourceEnvNameArg, targetEnvNameArg string

	if len(envs) == 1 {
		sourceEnv = &envs[0]
		targetEnv = &envs[0]
	}
	if len(args) > 0 {
		sourceEnvNameArg, sourceProjectArg = splitEnvAndProject(args[0])
	}
	if len(args) > 1 {
		targetEnvNameArg, targetProjectArg = splitEnvAndProject(args[1])
	}
	if sourceEnv == nil {
		sourceEnv, err = auth.SelectEnv("Select source environment", sourceEnvNameArg)
		if err != nil {
			return
		}
	}

	sourceCtx = sourceEnv.AuthContext(ctx)
	var sourceProjects, targetProjects []project.ProjectEntity
	sourceProjects, err = project.GetProjects(sourceEnv.AuthContext(ctx))
	if err != nil {
		return
	}

	sourceProject, err = entity.SelectEntityByNameOrAsk(sourceProjectArg, sourceProjects, project.ProjectEntityDesc)
	if err != nil {
		return
	}
	if targetEnv == nil {
		targetEnv, err = auth.SelectEnv("Select target environment", targetEnvNameArg)
		if err != nil {
			return
		}
	}

	targetCtx = targetEnv.AuthContext(ctx)
	targetProjects, err = project.GetProjects(targetEnv.AuthContext(ctx))
	if err != nil {
		return
	}
	targetProjectName, err = project.ValidateProjectName(targetProjectArg, sourceProject.GetName(), targetProjects)
	return
}

func splitEnvAndProject(arg string) (env, project string) {
	stringSplit := strings.Split(arg, ":")
	if len(stringSplit) == 1 {
		return "", stringSplit[0]
	}
	return stringSplit[0], stringSplit[1]
}

func init() {
	RootCommand.AddCommand(NewCopyCmd())
}
