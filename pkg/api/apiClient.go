package api

import (
	"context"

	. "github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func CreateAuthenticatedContext(parentCtx context.Context, apiKey string,  baseUrl string) context.Context {
	ctx := context.WithValue(parentCtx, ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
	ctx = context.WithValue(ctx, ContextAccessToken, apiKey)
	return ctx
}

func getApiClient() *DefaultApiService {
	cfg := NewConfiguration()
	cfg.Servers = ServerConfigurations{
		{
			URL: "{baseUrl}",
			Variables: map[string]ServerVariable{
				"baseUrl": {},
			},
		},
	}
	return NewAPIClient(cfg).DefaultApi
}

var ApiClient = getApiClient()
