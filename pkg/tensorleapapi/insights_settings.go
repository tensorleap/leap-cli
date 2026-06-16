/*
node-server

Hand-added client for the insightsSettings.getInsightsSettings endpoint.

This file mirrors the OpenAPI-generated style for a single POST endpoint that
the published spec does not yet expose. `scripts/update_server_api.sh` wipes
`pkg/tensorleapapi` and regenerates it, so once the endpoint lands on
node-server master this file is replaced by the generated equivalents.
*/

package tensorleapapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

// InsightsSettingsOverrides are the per-version user customizations of insight
// settings. Empty maps/slices mean "no override".
type InsightsSettingsOverrides struct {
	MetricActiveOverrides      map[string]bool   `json:"metricActiveOverrides,omitempty"`
	MetricDirectionOverrides   map[string]string `json:"metricDirectionOverrides,omitempty"`
	IgnoredDomains             []string          `json:"ignoredDomains,omitempty"`
	AddedDomains               []string          `json:"addedDomains,omitempty"`
	LatentSpaceIgnoreOverrides map[string]bool   `json:"latentSpaceIgnoreOverrides,omitempty"`
}

// GetInsightsSettingsParams is the request body for getInsightsSettings.
type GetInsightsSettingsParams struct {
	ProjectId string `json:"projectId"`
	VersionId string `json:"versionId"`
}

// NewGetInsightsSettingsParams instantiates a new GetInsightsSettingsParams.
func NewGetInsightsSettingsParams(projectId string, versionId string) *GetInsightsSettingsParams {
	return &GetInsightsSettingsParams{ProjectId: projectId, VersionId: versionId}
}

// GetInsightsSettingsResponse is the getInsightsSettings response. Only the
// overrides are modeled (defaults are ignored by the JSON decoder).
type GetInsightsSettingsResponse struct {
	Overrides InsightsSettingsOverrides `json:"overrides"`
}

type ApiGetInsightsSettingsRequest struct {
	ctx                       context.Context
	ApiService                *DefaultAPIService
	getInsightsSettingsParams *GetInsightsSettingsParams
}

func (r ApiGetInsightsSettingsRequest) GetInsightsSettingsParams(getInsightsSettingsParams GetInsightsSettingsParams) ApiGetInsightsSettingsRequest {
	r.getInsightsSettingsParams = &getInsightsSettingsParams
	return r
}

func (r ApiGetInsightsSettingsRequest) Execute() (*GetInsightsSettingsResponse, *http.Response, error) {
	return r.ApiService.GetInsightsSettingsExecute(r)
}

// GetInsightsSettings Method for GetInsightsSettings
func (a *DefaultAPIService) GetInsightsSettings(ctx context.Context) ApiGetInsightsSettingsRequest {
	return ApiGetInsightsSettingsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// GetInsightsSettingsExecute executes the request
//
//	@return GetInsightsSettingsResponse
func (a *DefaultAPIService) GetInsightsSettingsExecute(r ApiGetInsightsSettingsRequest) (*GetInsightsSettingsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *GetInsightsSettingsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "DefaultAPIService.GetInsightsSettings")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/insightsSettings/getInsightsSettings"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getInsightsSettingsParams == nil {
		return localVarReturnValue, nil, reportError("getInsightsSettingsParams is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.getInsightsSettingsParams
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
