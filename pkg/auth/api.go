package auth

import (
	"context"

	"github.com/tensorleap/leap-cli/pkg/api"
	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func GetWhoami(ctx context.Context) (*tensorleapapi.UserData, error) {
	userData, _, err := api.ApiClient.WhoAmI(ctx).Execute()
	if err != nil {
		return nil, err
	}
	return userData, nil
}
