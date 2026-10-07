# ExportBundleParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**TopK** | Pointer to **float64** |  | [optional] 
**HeatmapLabels** | Pointer to **[]string** |  | [optional] 

## Methods

### NewExportBundleParams

`func NewExportBundleParams(projectId string, versionId string, ) *ExportBundleParams`

NewExportBundleParams instantiates a new ExportBundleParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportBundleParamsWithDefaults

`func NewExportBundleParamsWithDefaults() *ExportBundleParams`

NewExportBundleParamsWithDefaults instantiates a new ExportBundleParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *ExportBundleParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ExportBundleParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ExportBundleParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *ExportBundleParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *ExportBundleParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *ExportBundleParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetTopK

`func (o *ExportBundleParams) GetTopK() float64`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *ExportBundleParams) GetTopKOk() (*float64, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *ExportBundleParams) SetTopK(v float64)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *ExportBundleParams) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetHeatmapLabels

`func (o *ExportBundleParams) GetHeatmapLabels() []string`

GetHeatmapLabels returns the HeatmapLabels field if non-nil, zero value otherwise.

### GetHeatmapLabelsOk

`func (o *ExportBundleParams) GetHeatmapLabelsOk() (*[]string, bool)`

GetHeatmapLabelsOk returns a tuple with the HeatmapLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeatmapLabels

`func (o *ExportBundleParams) SetHeatmapLabels(v []string)`

SetHeatmapLabels sets HeatmapLabels field to given value.

### HasHeatmapLabels

`func (o *ExportBundleParams) HasHeatmapLabels() bool`

HasHeatmapLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


