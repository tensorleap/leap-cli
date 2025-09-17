package api

import (
	"context"
	"fmt"
	"time"

	"github.com/tensorleap/leap-cli/pkg/tensorleapapi"
)

func GetDownloadSignedUrl(ctx context.Context, blobPath string) (string, error) {
	params := *tensorleapapi.NewGetDownloadSignedUrlParams(blobPath)

	base_url := GetBaseUrlFromContext(ctx)
	params.SetOrigin(base_url)

	res, res_http, err := ApiClient.GetDownloadSignedUrl(ctx).
		GetDownloadSignedUrlParams(params).
		Execute()
	if err = CheckRes(res_http, err); err != nil {
		return "", fmt.Errorf("failed to get signed url: %v", err)
	}
	return res.GetUrl(), nil
}

func GetUploadSignedUrl(ctx context.Context, blobPath string) (*tensorleapapi.ExternalImportModelStorage, error) {
	params := *tensorleapapi.NewGetUploadSignedUrlParams(blobPath)

	base_url := GetBaseUrlFromContext(ctx)
	params.SetOrigin(base_url)
	res, res_http, err := ApiClient.GetUploadSignedUrl(ctx).
		GetUploadSignedUrlParams(params).
		Execute()
	if err = CheckRes(res_http, err); err != nil {
		return nil, fmt.Errorf("failed to get signed url: %v", err)
	}
	return res, nil
}

func GetSignedUrl(ctx context.Context, url string, method tensorleapapi.HttpMethods, expire time.Duration, origin *string) (string, error) {
	expireTime := float64(expire.Seconds())
	getUrlParams := *tensorleapapi.NewGetSignedUrlParams(url, expireTime, method)
	if origin != nil {
		getUrlParams.SetOrigin(*origin)
	} else {
		base_url := GetBaseUrlFromContext(ctx)
		getUrlParams.SetOrigin(base_url)
	}
	getUrl, res, err := ApiClient.GetSignedUrl(ctx).
		GetSignedUrlParams(getUrlParams).Execute()
	if err = CheckRes(res, err); err != nil {
		return "", fmt.Errorf("failed to get signed url: %v", err)
	}
	return getUrl.Url, nil
}

func GetBaseUrlFromContext(ctx context.Context) string {
	base_url, _ := GetAuthFromContext(ctx)

	return ChangeToUIUrl(base_url)
}
