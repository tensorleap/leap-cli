package entity

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

func CreateUniqueNameValidator(existingNames []string) survey.Validator {
	return func(val interface{}) error {
		name := strings.TrimSpace(val.(string))
		if name == "" {
			return fmt.Errorf("name is required")
		}

		// Check if the name already exists
		for _, existingName := range existingNames {
			if name == existingName {
				return fmt.Errorf("name already exists")
			}
		}

		return nil
	}
}
