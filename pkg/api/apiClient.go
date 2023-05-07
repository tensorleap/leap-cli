package api

import (
	"github.com/tensorleap/cli-go/pkg/tensorleapapi"
)

func getClient() *tensorleapapi.DefaultApiService {
  cfg := tensorleapapi.NewConfiguration()
  cfg.Servers = tensorleapapi.ServerConfigurations{
    {
      URL: "http://127.0.0.1:4589/api/v2",
    },
  }
  return tensorleapapi.NewAPIClient(cfg).DefaultApi
}

var Client = getClient()
