package entity

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func AskUserIsCreatingNew[TEntity any](desc *EntityDescriptor[TEntity]) (bool, error) {
	var response int

	prompt := &survey.Select{
		Message: fmt.Sprintf("Select %s to use:", desc.Name),
		Options: []string{
			"Create new",
			"Select existing",
		},
	}

	err := survey.AskOne(prompt, &response)
	if err != nil {
		return false, fmt.Errorf("no option selected: %v", err)
	}
	return response == 0, nil
}

func SelectEntity[TEntity any](entities []TEntity, desc *EntityDescriptor[TEntity]) (*TEntity, error) {
	if len(entities) == 0 {
		return nil, fmt.Errorf("no %s found, please create %s", desc.Name, desc.Name)
	}
	index := -1
	options := []string{}
	for _, entity := range entities {
		options = append(options, desc.GetDisplayName(&entity))
	}
	prompt := &survey.Select{
		Message: fmt.Sprintf("Select %s:", desc.Name),
		Options: options,
	}
	err := survey.AskOne(prompt, &index)
	if err != nil || index == -1 {
		return nil, fmt.Errorf("no %s selected", desc.Name)
	}
	return &entities[index], nil
}

func SelectEntityOrCreateOne[TEntity any](
	entities []TEntity,
	createEntity func() (*TEntity, error),
	askForNewFirst bool,
	desc *EntityDescriptor[TEntity],
) (*TEntity, error) {

	thereAreProjects := len(entities) > 0

	isCreatingNew := !thereAreProjects
	if askForNewFirst {
		var err error
		if thereAreProjects {
			isCreatingNew, err = AskUserIsCreatingNew(desc)
			if err != nil {
				return nil, err
			}
		}

	}
	if isCreatingNew {
		return createEntity()
	}
	return SelectEntity(entities, desc)
}
