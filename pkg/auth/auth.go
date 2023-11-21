package auth

import (
	"context"

	"github.com/tensorleap/leap-cli/pkg/api"
)

type Env struct {
	Name   string
	ApiUrl string
	ApiKey string
}

func NewEnv(name, apiUrl, apiKey string) *Env {
	return &Env{
		Name:   name,
		ApiUrl: apiUrl,
		ApiKey: apiKey,
	}
}

func (a *Env) PrintWhoami(ctx context.Context) error {
	authCtx := a.AuthContext(ctx)
	return PrintWhoami(authCtx)
}

func (a *Env) AuthContext(ctx context.Context) context.Context {
	return api.CreateAuthenticatedContext(ctx, a.ApiKey, a.ApiUrl)
}
