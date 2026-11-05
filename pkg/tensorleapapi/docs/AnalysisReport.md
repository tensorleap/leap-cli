# AnalysisReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**Title** | **string** |  | 
**Summary** | **string** |  | 
**HtmlBlobPath** | **string** |  | 
**MarkdownBlobPath** | Pointer to **string** |  | [optional] 
**JsonBlobPath** | Pointer to **string** |  | [optional] 
**CreatedBy** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewAnalysisReport

`func NewAnalysisReport(cid string, projectId string, versionId string, title string, summary string, htmlBlobPath string, createdBy string, createdAt time.Time, updatedAt time.Time, ) *AnalysisReport`

NewAnalysisReport instantiates a new AnalysisReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisReportWithDefaults

`func NewAnalysisReportWithDefaults() *AnalysisReport`

NewAnalysisReportWithDefaults instantiates a new AnalysisReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *AnalysisReport) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *AnalysisReport) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *AnalysisReport) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProjectId

`func (o *AnalysisReport) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AnalysisReport) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AnalysisReport) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *AnalysisReport) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AnalysisReport) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AnalysisReport) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetInferenceArtifactId

`func (o *AnalysisReport) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *AnalysisReport) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *AnalysisReport) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *AnalysisReport) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetTitle

`func (o *AnalysisReport) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *AnalysisReport) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *AnalysisReport) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetSummary

`func (o *AnalysisReport) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *AnalysisReport) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *AnalysisReport) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetHtmlBlobPath

`func (o *AnalysisReport) GetHtmlBlobPath() string`

GetHtmlBlobPath returns the HtmlBlobPath field if non-nil, zero value otherwise.

### GetHtmlBlobPathOk

`func (o *AnalysisReport) GetHtmlBlobPathOk() (*string, bool)`

GetHtmlBlobPathOk returns a tuple with the HtmlBlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlBlobPath

`func (o *AnalysisReport) SetHtmlBlobPath(v string)`

SetHtmlBlobPath sets HtmlBlobPath field to given value.


### GetMarkdownBlobPath

`func (o *AnalysisReport) GetMarkdownBlobPath() string`

GetMarkdownBlobPath returns the MarkdownBlobPath field if non-nil, zero value otherwise.

### GetMarkdownBlobPathOk

`func (o *AnalysisReport) GetMarkdownBlobPathOk() (*string, bool)`

GetMarkdownBlobPathOk returns a tuple with the MarkdownBlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkdownBlobPath

`func (o *AnalysisReport) SetMarkdownBlobPath(v string)`

SetMarkdownBlobPath sets MarkdownBlobPath field to given value.

### HasMarkdownBlobPath

`func (o *AnalysisReport) HasMarkdownBlobPath() bool`

HasMarkdownBlobPath returns a boolean if a field has been set.

### GetJsonBlobPath

`func (o *AnalysisReport) GetJsonBlobPath() string`

GetJsonBlobPath returns the JsonBlobPath field if non-nil, zero value otherwise.

### GetJsonBlobPathOk

`func (o *AnalysisReport) GetJsonBlobPathOk() (*string, bool)`

GetJsonBlobPathOk returns a tuple with the JsonBlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonBlobPath

`func (o *AnalysisReport) SetJsonBlobPath(v string)`

SetJsonBlobPath sets JsonBlobPath field to given value.

### HasJsonBlobPath

`func (o *AnalysisReport) HasJsonBlobPath() bool`

HasJsonBlobPath returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AnalysisReport) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AnalysisReport) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AnalysisReport) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *AnalysisReport) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalysisReport) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalysisReport) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AnalysisReport) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AnalysisReport) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AnalysisReport) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


