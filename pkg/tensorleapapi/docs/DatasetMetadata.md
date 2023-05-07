# DatasetMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Setup** | Pointer to [**DatasetSetup**](DatasetSetup.md) |  | [optional] 
**SetupStatus** | Pointer to [**DatasetSetupStatus**](DatasetSetupStatus.md) |  | [optional] 
**ModelSetup** | Pointer to [**ModelSetup**](ModelSetup.md) |  | [optional] 
**SecretManagerId** | Pointer to **string** |  | [optional] 

## Methods

### NewDatasetMetadata

`func NewDatasetMetadata() *DatasetMetadata`

NewDatasetMetadata instantiates a new DatasetMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetMetadataWithDefaults

`func NewDatasetMetadataWithDefaults() *DatasetMetadata`

NewDatasetMetadataWithDefaults instantiates a new DatasetMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSetup

`func (o *DatasetMetadata) GetSetup() DatasetSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *DatasetMetadata) GetSetupOk() (*DatasetSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *DatasetMetadata) SetSetup(v DatasetSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *DatasetMetadata) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSetupStatus

`func (o *DatasetMetadata) GetSetupStatus() DatasetSetupStatus`

GetSetupStatus returns the SetupStatus field if non-nil, zero value otherwise.

### GetSetupStatusOk

`func (o *DatasetMetadata) GetSetupStatusOk() (*DatasetSetupStatus, bool)`

GetSetupStatusOk returns a tuple with the SetupStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupStatus

`func (o *DatasetMetadata) SetSetupStatus(v DatasetSetupStatus)`

SetSetupStatus sets SetupStatus field to given value.

### HasSetupStatus

`func (o *DatasetMetadata) HasSetupStatus() bool`

HasSetupStatus returns a boolean if a field has been set.

### GetModelSetup

`func (o *DatasetMetadata) GetModelSetup() ModelSetup`

GetModelSetup returns the ModelSetup field if non-nil, zero value otherwise.

### GetModelSetupOk

`func (o *DatasetMetadata) GetModelSetupOk() (*ModelSetup, bool)`

GetModelSetupOk returns a tuple with the ModelSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelSetup

`func (o *DatasetMetadata) SetModelSetup(v ModelSetup)`

SetModelSetup sets ModelSetup field to given value.

### HasModelSetup

`func (o *DatasetMetadata) HasModelSetup() bool`

HasModelSetup returns a boolean if a field has been set.

### GetSecretManagerId

`func (o *DatasetMetadata) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *DatasetMetadata) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *DatasetMetadata) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *DatasetMetadata) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


