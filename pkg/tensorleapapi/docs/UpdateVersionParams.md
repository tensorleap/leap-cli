# UpdateVersionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VersionId** | **string** |  | 
**Data** | [**ModelGraph**](ModelGraph.md) |  | 
**CodeIntegrationVersionId** | Pointer to **string** |  | [optional] 
**DatasetSetup** | [**DatasetSetup**](DatasetSetup.md) |  | 
**OverrideHash** | **bool** |  | 
**Hash** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateVersionParams

`func NewUpdateVersionParams(versionId string, data ModelGraph, datasetSetup DatasetSetup, overrideHash bool, ) *UpdateVersionParams`

NewUpdateVersionParams instantiates a new UpdateVersionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVersionParamsWithDefaults

`func NewUpdateVersionParamsWithDefaults() *UpdateVersionParams`

NewUpdateVersionParamsWithDefaults instantiates a new UpdateVersionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersionId

`func (o *UpdateVersionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UpdateVersionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UpdateVersionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetData

`func (o *UpdateVersionParams) GetData() ModelGraph`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *UpdateVersionParams) GetDataOk() (*ModelGraph, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *UpdateVersionParams) SetData(v ModelGraph)`

SetData sets Data field to given value.


### GetCodeIntegrationVersionId

`func (o *UpdateVersionParams) GetCodeIntegrationVersionId() string`

GetCodeIntegrationVersionId returns the CodeIntegrationVersionId field if non-nil, zero value otherwise.

### GetCodeIntegrationVersionIdOk

`func (o *UpdateVersionParams) GetCodeIntegrationVersionIdOk() (*string, bool)`

GetCodeIntegrationVersionIdOk returns a tuple with the CodeIntegrationVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeIntegrationVersionId

`func (o *UpdateVersionParams) SetCodeIntegrationVersionId(v string)`

SetCodeIntegrationVersionId sets CodeIntegrationVersionId field to given value.

### HasCodeIntegrationVersionId

`func (o *UpdateVersionParams) HasCodeIntegrationVersionId() bool`

HasCodeIntegrationVersionId returns a boolean if a field has been set.

### GetDatasetSetup

`func (o *UpdateVersionParams) GetDatasetSetup() DatasetSetup`

GetDatasetSetup returns the DatasetSetup field if non-nil, zero value otherwise.

### GetDatasetSetupOk

`func (o *UpdateVersionParams) GetDatasetSetupOk() (*DatasetSetup, bool)`

GetDatasetSetupOk returns a tuple with the DatasetSetup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetSetup

`func (o *UpdateVersionParams) SetDatasetSetup(v DatasetSetup)`

SetDatasetSetup sets DatasetSetup field to given value.


### GetOverrideHash

`func (o *UpdateVersionParams) GetOverrideHash() bool`

GetOverrideHash returns the OverrideHash field if non-nil, zero value otherwise.

### GetOverrideHashOk

`func (o *UpdateVersionParams) GetOverrideHashOk() (*bool, bool)`

GetOverrideHashOk returns a tuple with the OverrideHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideHash

`func (o *UpdateVersionParams) SetOverrideHash(v bool)`

SetOverrideHash sets OverrideHash field to given value.


### GetHash

`func (o *UpdateVersionParams) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *UpdateVersionParams) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *UpdateVersionParams) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *UpdateVersionParams) HasHash() bool`

HasHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


