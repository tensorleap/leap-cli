# AnalyzeParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**AnalyzeTypeEnum**](AnalyzeTypeEnum.md) |  | 
**FromEpoch** | **float64** |  | 
**BatchSize** | Pointer to **float64** |  | [optional] 
**SampleIdentity** | Pointer to [**SampleIdentity**](SampleIdentity.md) |  | [optional] 
**FromDatasetSlice** | Pointer to [**DataStateType**](DataStateType.md) |  | [optional] 
**ExtId** | **string** |  | 

## Methods

### NewAnalyzeParams

`func NewAnalyzeParams(type_ AnalyzeTypeEnum, fromEpoch float64, extId string, ) *AnalyzeParams`

NewAnalyzeParams instantiates a new AnalyzeParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyzeParamsWithDefaults

`func NewAnalyzeParamsWithDefaults() *AnalyzeParams`

NewAnalyzeParamsWithDefaults instantiates a new AnalyzeParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AnalyzeParams) GetType() AnalyzeTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyzeParams) GetTypeOk() (*AnalyzeTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyzeParams) SetType(v AnalyzeTypeEnum)`

SetType sets Type field to given value.


### GetFromEpoch

`func (o *AnalyzeParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *AnalyzeParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *AnalyzeParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetBatchSize

`func (o *AnalyzeParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *AnalyzeParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *AnalyzeParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *AnalyzeParams) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetSampleIdentity

`func (o *AnalyzeParams) GetSampleIdentity() SampleIdentity`

GetSampleIdentity returns the SampleIdentity field if non-nil, zero value otherwise.

### GetSampleIdentityOk

`func (o *AnalyzeParams) GetSampleIdentityOk() (*SampleIdentity, bool)`

GetSampleIdentityOk returns a tuple with the SampleIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentity

`func (o *AnalyzeParams) SetSampleIdentity(v SampleIdentity)`

SetSampleIdentity sets SampleIdentity field to given value.

### HasSampleIdentity

`func (o *AnalyzeParams) HasSampleIdentity() bool`

HasSampleIdentity returns a boolean if a field has been set.

### GetFromDatasetSlice

`func (o *AnalyzeParams) GetFromDatasetSlice() DataStateType`

GetFromDatasetSlice returns the FromDatasetSlice field if non-nil, zero value otherwise.

### GetFromDatasetSliceOk

`func (o *AnalyzeParams) GetFromDatasetSliceOk() (*DataStateType, bool)`

GetFromDatasetSliceOk returns a tuple with the FromDatasetSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatasetSlice

`func (o *AnalyzeParams) SetFromDatasetSlice(v DataStateType)`

SetFromDatasetSlice sets FromDatasetSlice field to given value.

### HasFromDatasetSlice

`func (o *AnalyzeParams) HasFromDatasetSlice() bool`

HasFromDatasetSlice returns a boolean if a field has been set.

### GetExtId

`func (o *AnalyzeParams) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *AnalyzeParams) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *AnalyzeParams) SetExtId(v string)`

SetExtId sets ExtId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


