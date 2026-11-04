# PublishAnalysisReportParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**Html** | **string** |  | 
**Markdown** | Pointer to **string** |  | [optional] 
**Json** | Pointer to **string** |  | [optional] 

## Methods

### NewPublishAnalysisReportParams

`func NewPublishAnalysisReportParams(projectId string, versionId string, title string, summary string, html string, ) *PublishAnalysisReportParams`

NewPublishAnalysisReportParams instantiates a new PublishAnalysisReportParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishAnalysisReportParamsWithDefaults

`func NewPublishAnalysisReportParamsWithDefaults() *PublishAnalysisReportParams`

NewPublishAnalysisReportParamsWithDefaults instantiates a new PublishAnalysisReportParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *PublishAnalysisReportParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *PublishAnalysisReportParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *PublishAnalysisReportParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *PublishAnalysisReportParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *PublishAnalysisReportParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *PublishAnalysisReportParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetTitle

`func (o *PublishAnalysisReportParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PublishAnalysisReportParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PublishAnalysisReportParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *PublishAnalysisReportParams) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *PublishAnalysisReportParams) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *PublishAnalysisReportParams) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetHtml

`func (o *PublishAnalysisReportParams) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *PublishAnalysisReportParams) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *PublishAnalysisReportParams) SetHtml(v string)`

SetHtml sets Html field to given value.


### GetMarkdown

`func (o *PublishAnalysisReportParams) GetMarkdown() string`

GetMarkdown returns the Markdown field if non-nil, zero value otherwise.

### GetMarkdownOk

`func (o *PublishAnalysisReportParams) GetMarkdownOk() (*string, bool)`

GetMarkdownOk returns a tuple with the Markdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdown

`func (o *PublishAnalysisReportParams) SetMarkdown(v string)`

SetMarkdown sets Markdown field to given value.

### HasMarkdown

`func (o *PublishAnalysisReportParams) HasMarkdown() bool`

HasMarkdown returns a boolean if a field has been set.

### GetJson

`func (o *PublishAnalysisReportParams) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *PublishAnalysisReportParams) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *PublishAnalysisReportParams) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *PublishAnalysisReportParams) HasJson() bool`

HasJson returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


