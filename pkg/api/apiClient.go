package api

import (
	"context"
	"net/http"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func CreateAuthenticatedContext(parentCtx context.Context, apiKey string, baseUrl string) context.Context {
	ctx := context.WithValue(parentCtx, tensorleapapi.ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
	ctx = context.WithValue(ctx, tensorleapapi.ContextAccessToken, apiKey)
	return ctx
}

func CreateContextWithBaseURL(parentCtx context.Context, baseUrl string) context.Context {
	return context.WithValue(parentCtx, tensorleapapi.ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
}

func getApiClient() *tensorleapapi.DefaultAPIService {
	cfg := tensorleapapi.NewConfiguration()

	cfg.HTTPClient = NewDefaultClient()

	cfg.Servers = tensorleapapi.ServerConfigurations{
		{
			URL: "{baseUrl}",
			Variables: map[string]tensorleapapi.ServerVariable{
				"baseUrl": {},
			},
		},
	}
	return tensorleapapi.NewAPIClient(cfg).DefaultAPI
}

func GetAuthFromContext(ctx context.Context) (baseUrl string, apiKey string) {
	apiKey = ctx.Value(tensorleapapi.ContextAccessToken).(string)
	baseUrl = ctx.Value(tensorleapapi.ContextServerVariables).(map[string]string)["baseUrl"]
	return
}

func AddAuthToRequestHeader(h *http.Header, apiKey string) {
	h.Add("Authorization", "Bearer "+apiKey)
}

var ApiClient = getApiClient()

// Export type aliases for backward compatibility
type DefaultAPIService = tensorleapapi.DefaultAPIService
