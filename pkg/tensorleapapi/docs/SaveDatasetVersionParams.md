# SaveDatasetVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DatasetId** | **string** |  | 
**Setup** | Pointer to [**DatasetSetup**](DatasetSetup.md) |  | [optional] 
**SecretManagerId** | Pointer to **string** |  | [optional] 
**CodeUrl** | **string** |  | 
**CodeEntryFile** | **string** |  | 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewSaveDatasetVersionParams

`func NewSaveDatasetVersionParams(datasetId string, codeUrl string, codeEntryFile string, ) *SaveDatasetVersionParams`

NewSaveDatasetVersionParams instantiates a new SaveDatasetVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDatasetVersionParamsWithDefaults

`func NewSaveDatasetVersionParamsWithDefaults() *SaveDatasetVersionParams`

NewSaveDatasetVersionParamsWithDefaults instantiates a new SaveDatasetVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatasetId

`func (o *SaveDatasetVersionParams) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *SaveDatasetVersionParams) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *SaveDatasetVersionParams) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetSetup

`func (o *SaveDatasetVersionParams) GetSetup() DatasetSetup`

GetSetup returns the Setup field if non-nil, zero value otherwise.

### GetSetupOk

`func (o *SaveDatasetVersionParams) GetSetupOk() (*DatasetSetup, bool)`

GetSetupOk returns a tuple with the Setup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetup

`func (o *SaveDatasetVersionParams) SetSetup(v DatasetSetup)`

SetSetup sets Setup field to given value.

### HasSetup

`func (o *SaveDatasetVersionParams) HasSetup() bool`

HasSetup returns a boolean if a field has been set.

### GetSecretManagerId

`func (o *SaveDatasetVersionParams) GetSecretManagerId() string`

GetSecretManagerId returns the SecretManagerId field if non-nil, zero value otherwise.

### GetSecretManagerIdOk

`func (o *SaveDatasetVersionParams) GetSecretManagerIdOk() (*string, bool)`

GetSecretManagerIdOk returns a tuple with the SecretManagerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretManagerId

`func (o *SaveDatasetVersionParams) SetSecretManagerId(v string)`

SetSecretManagerId sets SecretManagerId field to given value.

### HasSecretManagerId

`func (o *SaveDatasetVersionParams) HasSecretManagerId() bool`

HasSecretManagerId returns a boolean if a field has been set.

### GetCodeUrl

`func (o *SaveDatasetVersionParams) GetCodeUrl() string`

GetCodeUrl returns the CodeUrl field if non-nil, zero value otherwise.

### GetCodeUrlOk

`func (o *SaveDatasetVersionParams) GetCodeUrlOk() (*string, bool)`

GetCodeUrlOk returns a tuple with the CodeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeUrl

`func (o *SaveDatasetVersionParams) SetCodeUrl(v string)`

SetCodeUrl sets CodeUrl field to given value.


### GetCodeEntryFile

`func (o *SaveDatasetVersionParams) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *SaveDatasetVersionParams) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *SaveDatasetVersionParams) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.


### GetNote

`func (o *SaveDatasetVersionParams) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *SaveDatasetVersionParams) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *SaveDatasetVersionParams) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *SaveDatasetVersionParams) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


