package entity

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/AlecAivazis/survey/v2"
	"github.com/tensorleap/leap-cli/pkg/log"
)

func GetEntityById[TEntity any](entityId string, entities []TEntity, desc *EntityDescriptor[TEntity]) (*TEntity, error) {
	for _, entity := range entities {
		if desc.GetID(&entity) == entityId {
			return &entity, nil
		}
	}
	return nil, fmt.Errorf("not found %s id: %s", entityId, desc.Name)
}

func GetEntityByDisplayName[TEntity any](entityName string, entities []TEntity, desc *EntityDescriptor[TEntity]) (*TEntity, error) {
	for _, entity := range entities {
		if desc.GetDisplayName(&entity) == entityName {
			return &entity, nil
		}
	}
	return nil, fmt.Errorf("not found %s name: %s", entityName, desc.Name)
}

func GetNames[TEntity any](entities []TEntity, desc *EntityDescriptor[TEntity]) []string {
	names := []string{}
	for _, entity := range entities {
		names = append(names, desc.GetDisplayName(&entity))
	}
	return names
}

func PrintList[TEntity any](entities []TEntity, desc *EntityDescriptor[TEntity]) {
	if len(entities) == 0 {
		fmt.Printf("No %s \n", desc.PluralName)
		return
	}

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', 0)
	for _, entity := range entities {
		displayName := desc.GetDisplayName(&entity)
		id := desc.GetID(&entity)
		fmt.Fprintf(w, "%s\t%s\n", displayName, id)
	}
	w.Flush()
}

func ConfirmDeletion[TEntity any](entity *TEntity, desc *EntityDescriptor[TEntity]) (confirmed bool, err error) {

	prompt := &survey.Confirm{
		Message: fmt.Sprintf("Confirm deletion of %s: %s id: %s", desc.Name, desc.GetDisplayName(entity), desc.GetID(entity)),
	}

	err = survey.AskOne(prompt, &confirmed)

	return
}

func InfoCreation[TEntity any](entity *TEntity, desc *EntityDescriptor[TEntity]) {
	log.Infof("Successfully created %s '%s'", desc.Name, desc.GetDisplayName(entity))
}

func InfoDeletion[TEntity any](entity *TEntity, desc *EntityDescriptor[TEntity]) {
	log.Infof("Successfully deleted %s '%s'", desc.Name, desc.GetDisplayName(entity))
}
