# ExportProjectMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Tags** | **[]string** |  | 
**Categories** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**BgImagePath** | **string** |  | 
**SchemaVersion** | **float64** |  | 
**SourceProjectId** | **string** |  | 

## Methods

### NewExportProjectMeta

`func NewExportProjectMeta(name string, tags []string, categories map[string]interface{}, bgImagePath string, schemaVersion float64, sourceProjectId string, ) *ExportProjectMeta`

NewExportProjectMeta instantiates a new ExportProjectMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportProjectMetaWithDefaults

`func NewExportProjectMetaWithDefaults() *ExportProjectMeta`

NewExportProjectMetaWithDefaults instantiates a new ExportProjectMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExportProjectMeta) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExportProjectMeta) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExportProjectMeta) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ExportProjectMeta) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ExportProjectMeta) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ExportProjectMeta) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ExportProjectMeta) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *ExportProjectMeta) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ExportProjectMeta) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ExportProjectMeta) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetCategories

`func (o *ExportProjectMeta) GetCategories() map[string]interface{}`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ExportProjectMeta) GetCategoriesOk() (*map[string]interface{}, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ExportProjectMeta) SetCategories(v map[string]interface{})`

SetCategories sets Categories field to given value.


### GetBgImagePath

`func (o *ExportProjectMeta) GetBgImagePath() string`

GetBgImagePath returns the BgImagePath field if non-nil, zero value otherwise.

### GetBgImagePathOk

`func (o *ExportProjectMeta) GetBgImagePathOk() (*string, bool)`

GetBgImagePathOk returns a tuple with the BgImagePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgImagePath

`func (o *ExportProjectMeta) SetBgImagePath(v string)`

SetBgImagePath sets BgImagePath field to given value.


### GetSchemaVersion

`func (o *ExportProjectMeta) GetSchemaVersion() float64`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *ExportProjectMeta) GetSchemaVersionOk() (*float64, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *ExportProjectMeta) SetSchemaVersion(v float64)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetSourceProjectId

`func (o *ExportProjectMeta) GetSourceProjectId() string`

GetSourceProjectId returns the SourceProjectId field if non-nil, zero value otherwise.

### GetSourceProjectIdOk

`func (o *ExportProjectMeta) GetSourceProjectIdOk() (*string, bool)`

GetSourceProjectIdOk returns a tuple with the SourceProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceProjectId

`func (o *ExportProjectMeta) SetSourceProjectId(v string)`

SetSourceProjectId sets SourceProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


