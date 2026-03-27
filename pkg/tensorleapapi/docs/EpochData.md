# EpochData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**Epoch** | **float64** |  | 
**Tags** | **[]string** |  | 
**UploadedModelFilePath** | Pointer to **string** |  | [optional] 
**ExternalData** | Pointer to [**EpochDataExternalData**](EpochDataExternalData.md) |  | [optional] 

## Methods

### NewEpochData

`func NewEpochData(versionId string, epoch float64, tags []string, ) *EpochData`

NewEpochData instantiates a new EpochData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEpochDataWithDefaults

`func NewEpochDataWithDefaults() *EpochData`

NewEpochDataWithDefaults instantiates a new EpochData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *EpochData) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *EpochData) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *EpochData) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetEpoch

`func (o *EpochData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *EpochData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *EpochData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetTags

`func (o *EpochData) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *EpochData) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *EpochData) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetUploadedModelFilePath

`func (o *EpochData) GetUploadedModelFilePath() string`

GetUploadedModelFilePath returns the UploadedModelFilePath field if non-nil, zero value otherwise.

### GetUploadedModelFilePathOk

`func (o *EpochData) GetUploadedModelFilePathOk() (*string, bool)`

GetUploadedModelFilePathOk returns a tuple with the UploadedModelFilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadedModelFilePath

`func (o *EpochData) SetUploadedModelFilePath(v string)`

SetUploadedModelFilePath sets UploadedModelFilePath field to given value.

### HasUploadedModelFilePath

`func (o *EpochData) HasUploadedModelFilePath() bool`

HasUploadedModelFilePath returns a boolean if a field has been set.

### GetExternalData

`func (o *EpochData) GetExternalData() EpochDataExternalData`

GetExternalData returns the ExternalData field if non-nil, zero value otherwise.

### GetExternalDataOk

`func (o *EpochData) GetExternalDataOk() (*EpochDataExternalData, bool)`

GetExternalDataOk returns a tuple with the ExternalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalData

`func (o *EpochData) SetExternalData(v EpochDataExternalData)`

SetExternalData sets ExternalData field to given value.

### HasExternalData

`func (o *EpochData) HasExternalData() bool`

HasExternalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


