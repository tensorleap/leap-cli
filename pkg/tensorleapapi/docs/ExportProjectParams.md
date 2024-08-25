# ExportProjectParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExportUrl** | **string** |  | 
**ProjectVersion** | **float64** |  | 
**ExportOptions** | [**ExportOptions**](ExportOptions.md) |  | 
**AlreadyExported** | **bool** |  | 
**ProjectExportMeta** | [**ExportProjectMeta**](ExportProjectMeta.md) |  | 
**CopyToUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewExportProjectParams

`func NewExportProjectParams(exportUrl string, projectVersion float64, exportOptions ExportOptions, alreadyExported bool, projectExportMeta ExportProjectMeta, ) *ExportProjectParams`

NewExportProjectParams instantiates a new ExportProjectParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportProjectParamsWithDefaults

`func NewExportProjectParamsWithDefaults() *ExportProjectParams`

NewExportProjectParamsWithDefaults instantiates a new ExportProjectParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExportUrl

`func (o *ExportProjectParams) GetExportUrl() string`

GetExportUrl returns the ExportUrl field if non-nil, zero value otherwise.

### GetExportUrlOk

`func (o *ExportProjectParams) GetExportUrlOk() (*string, bool)`

GetExportUrlOk returns a tuple with the ExportUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportUrl

`func (o *ExportProjectParams) SetExportUrl(v string)`

SetExportUrl sets ExportUrl field to given value.


### GetProjectVersion

`func (o *ExportProjectParams) GetProjectVersion() float64`

GetProjectVersion returns the ProjectVersion field if non-nil, zero value otherwise.

### GetProjectVersionOk

`func (o *ExportProjectParams) GetProjectVersionOk() (*float64, bool)`

GetProjectVersionOk returns a tuple with the ProjectVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectVersion

`func (o *ExportProjectParams) SetProjectVersion(v float64)`

SetProjectVersion sets ProjectVersion field to given value.


### GetExportOptions

`func (o *ExportProjectParams) GetExportOptions() ExportOptions`

GetExportOptions returns the ExportOptions field if non-nil, zero value otherwise.

### GetExportOptionsOk

`func (o *ExportProjectParams) GetExportOptionsOk() (*ExportOptions, bool)`

GetExportOptionsOk returns a tuple with the ExportOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportOptions

`func (o *ExportProjectParams) SetExportOptions(v ExportOptions)`

SetExportOptions sets ExportOptions field to given value.


### GetAlreadyExported

`func (o *ExportProjectParams) GetAlreadyExported() bool`

GetAlreadyExported returns the AlreadyExported field if non-nil, zero value otherwise.

### GetAlreadyExportedOk

`func (o *ExportProjectParams) GetAlreadyExportedOk() (*bool, bool)`

GetAlreadyExportedOk returns a tuple with the AlreadyExported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyExported

`func (o *ExportProjectParams) SetAlreadyExported(v bool)`

SetAlreadyExported sets AlreadyExported field to given value.


### GetProjectExportMeta

`func (o *ExportProjectParams) GetProjectExportMeta() ExportProjectMeta`

GetProjectExportMeta returns the ProjectExportMeta field if non-nil, zero value otherwise.

### GetProjectExportMetaOk

`func (o *ExportProjectParams) GetProjectExportMetaOk() (*ExportProjectMeta, bool)`

GetProjectExportMetaOk returns a tuple with the ProjectExportMeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectExportMeta

`func (o *ExportProjectParams) SetProjectExportMeta(v ExportProjectMeta)`

SetProjectExportMeta sets ProjectExportMeta field to given value.


### GetCopyToUrl

`func (o *ExportProjectParams) GetCopyToUrl() string`

GetCopyToUrl returns the CopyToUrl field if non-nil, zero value otherwise.

### GetCopyToUrlOk

`func (o *ExportProjectParams) GetCopyToUrlOk() (*string, bool)`

GetCopyToUrlOk returns a tuple with the CopyToUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToUrl

`func (o *ExportProjectParams) SetCopyToUrl(v string)`

SetCopyToUrl sets CopyToUrl field to given value.

### HasCopyToUrl

`func (o *ExportProjectParams) HasCopyToUrl() bool`

HasCopyToUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


