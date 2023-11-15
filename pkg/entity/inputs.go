package entity

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/log"
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

func SelectEntityByNameOrAsk[TEntity any](name string, entities []TEntity, desc *EntityDescriptor[TEntity]) (*TEntity, error) {
	if name != "" {
		entity, err := GetEntityByDisplayName(name, entities, desc)
		if err == nil {
			return entity, nil
		}
		log.Warnf("%s", err)
	}
	return SelectEntity(entities, desc)
}

func SelectEntitiesByNamesOrAsk[TEntity any](names []string, entities []TEntity, desc *EntityDescriptor[TEntity]) ([]TEntity, error) {
	var entitiesToSelect []TEntity
	if len(names) > 0 {
		for _, name := range names {
			entity, err := GetEntityByDisplayName(name, entities, desc)
			if err != nil {
				log.Warnf("%s", err)
				continue
			}
			entitiesToSelect = append(entitiesToSelect, *entity)
		}
	}
	if len(entitiesToSelect) > 0 {
		return entitiesToSelect, nil
	}
	entity, err := SelectEntity(entities, desc)
	if err != nil {
		return nil, err
	}
	return []TEntity{*entity}, nil
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
) (entity *TEntity, wasCreated bool, err error) {

	thereAreProjects := len(entities) > 0

	isCreatingNew := !thereAreProjects
	if askForNewFirst {
		var err error
		if thereAreProjects {
			isCreatingNew, err = AskUserIsCreatingNew(desc)
			if err != nil {
				return nil, false, err
			}
		}

	}

	if isCreatingNew {
		wasCreated = true
		entity, err = createEntity()
		return
	}
	entity, err = SelectEntity(entities, desc)
	return
}

func AskForName[TEntity any](existingNames []string, desc *EntityDescriptor[TEntity]) (name string, err error) {
	prompt := &survey.Input{
		Message: fmt.Sprintf("Enter %s name", desc.Name),
	}

	validate := CreateUniqueNameValidator(existingNames)

	err = survey.AskOne(prompt, &name, survey.WithValidator(validate))
	return
}
