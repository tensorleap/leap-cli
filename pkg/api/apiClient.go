package api

import (
	"context"
	"crypto/tls"
	"net/http"

	. "github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func CreateAuthenticatedContext(parentCtx context.Context, apiKey string, baseUrl string) context.Context {
	ctx := context.WithValue(parentCtx, ContextServerVariables, map[string]string{
		"baseUrl": baseUrl,
	})
	ctx = context.WithValue(ctx, ContextAccessToken, apiKey)
	return ctx
}

func getApiClient() *DefaultApiService {
	cfg := NewConfiguration()

	// // Create a custom HTTP client with disabled SSL verification
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Disable SSL verification
	}
	customHttpClient := &http.Client{
		Transport: customTransport,
	}
	cfg.HTTPClient = customHttpClient // Use the custom HTTP client

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

func GetAuthFromContext(ctx context.Context) (baseUrl string, apiKey string) {
	apiKey = ctx.Value(ContextAccessToken).(string)
	baseUrl = ctx.Value(ContextServerVariables).(map[string]string)["baseUrl"]
	return
}

func AddAuthToRequestHeader(h *http.Header, apiKey string) {
	h.Add("Authorization", "Bearer "+apiKey)
}

var ApiClient = getApiClient()
