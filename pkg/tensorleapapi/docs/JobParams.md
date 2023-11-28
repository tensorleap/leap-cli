# JobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Epochs** | **float64** |  | 
**BatchSize** | **float64** |  | 
**EarlyStopParams** | Pointer to [**EarlyStopParams**](EarlyStopParams.md) |  | [optional] 
**Type** | [**ExportModelTypeEnum**](ExportModelTypeEnum.md) |  | 
**FromEpoch** | **float64** |  | 
**NumOfSamples** | Pointer to **float64** |  | [optional] 
**BatchSize** | Pointer to **float64** |  | [optional] 
**SampleIdentity** | Pointer to [**SampleIdentity**](SampleIdentity.md) |  | [optional] 
**FromDatasetSlice** | Pointer to [**DataStateType**](DataStateType.md) |  | [optional] 
**ExtId** | **string** |  | 
**Digest** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Title** | **string** |  | 
**Epoch** | **float64** |  | 

## Methods

### NewJobParams

`func NewJobParams(epochs float64, batchSize float64, type_ ExportModelTypeEnum, fromEpoch float64, extId string, title string, epoch float64, ) *JobParams`

NewJobParams instantiates a new JobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobParamsWithDefaults

`func NewJobParamsWithDefaults() *JobParams`

NewJobParamsWithDefaults instantiates a new JobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEpochs

`func (o *JobParams) GetEpochs() float64`

GetEpochs returns the Epochs field if non-nil, zero value otherwise.

### GetEpochsOk

`func (o *JobParams) GetEpochsOk() (*float64, bool)`

GetEpochsOk returns a tuple with the Epochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpochs

`func (o *JobParams) SetEpochs(v float64)`

SetEpochs sets Epochs field to given value.


### GetBatchSize

`func (o *JobParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *JobParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *JobParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.


### GetEarlyStopParams

`func (o *JobParams) GetEarlyStopParams() EarlyStopParams`

GetEarlyStopParams returns the EarlyStopParams field if non-nil, zero value otherwise.

### GetEarlyStopParamsOk

`func (o *JobParams) GetEarlyStopParamsOk() (*EarlyStopParams, bool)`

GetEarlyStopParamsOk returns a tuple with the EarlyStopParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarlyStopParams

`func (o *JobParams) SetEarlyStopParams(v EarlyStopParams)`

SetEarlyStopParams sets EarlyStopParams field to given value.

### HasEarlyStopParams

`func (o *JobParams) HasEarlyStopParams() bool`

HasEarlyStopParams returns a boolean if a field has been set.

### GetType

`func (o *JobParams) GetType() ExportModelTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JobParams) GetTypeOk() (*ExportModelTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JobParams) SetType(v ExportModelTypeEnum)`

SetType sets Type field to given value.


### GetFromEpoch

`func (o *JobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *JobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *JobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetNumOfSamples

`func (o *JobParams) GetNumOfSamples() float64`

GetNumOfSamples returns the NumOfSamples field if non-nil, zero value otherwise.

### GetNumOfSamplesOk

`func (o *JobParams) GetNumOfSamplesOk() (*float64, bool)`

GetNumOfSamplesOk returns a tuple with the NumOfSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamples

`func (o *JobParams) SetNumOfSamples(v float64)`

SetNumOfSamples sets NumOfSamples field to given value.

### HasNumOfSamples

`func (o *JobParams) HasNumOfSamples() bool`

HasNumOfSamples returns a boolean if a field has been set.

### GetBatchSize

`func (o *JobParams) GetBatchSize() float64`

GetBatchSize returns the BatchSize field if non-nil, zero value otherwise.

### GetBatchSizeOk

`func (o *JobParams) GetBatchSizeOk() (*float64, bool)`

GetBatchSizeOk returns a tuple with the BatchSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchSize

`func (o *JobParams) SetBatchSize(v float64)`

SetBatchSize sets BatchSize field to given value.

### HasBatchSize

`func (o *JobParams) HasBatchSize() bool`

HasBatchSize returns a boolean if a field has been set.

### GetSampleIdentity

`func (o *JobParams) GetSampleIdentity() SampleIdentity`

GetSampleIdentity returns the SampleIdentity field if non-nil, zero value otherwise.

### GetSampleIdentityOk

`func (o *JobParams) GetSampleIdentityOk() (*SampleIdentity, bool)`

GetSampleIdentityOk returns a tuple with the SampleIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIdentity

`func (o *JobParams) SetSampleIdentity(v SampleIdentity)`

SetSampleIdentity sets SampleIdentity field to given value.

### HasSampleIdentity

`func (o *JobParams) HasSampleIdentity() bool`

HasSampleIdentity returns a boolean if a field has been set.

### GetFromDatasetSlice

`func (o *JobParams) GetFromDatasetSlice() DataStateType`

GetFromDatasetSlice returns the FromDatasetSlice field if non-nil, zero value otherwise.

### GetFromDatasetSliceOk

`func (o *JobParams) GetFromDatasetSliceOk() (*DataStateType, bool)`

GetFromDatasetSliceOk returns a tuple with the FromDatasetSlice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDatasetSlice

`func (o *JobParams) SetFromDatasetSlice(v DataStateType)`

SetFromDatasetSlice sets FromDatasetSlice field to given value.

### HasFromDatasetSlice

`func (o *JobParams) HasFromDatasetSlice() bool`

HasFromDatasetSlice returns a boolean if a field has been set.

### GetExtId

`func (o *JobParams) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *JobParams) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *JobParams) SetExtId(v string)`

SetExtId sets ExtId field to given value.


### GetDigest

`func (o *JobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *JobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *JobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *JobParams) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFilters

`func (o *JobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *JobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *JobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *JobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetTitle

`func (o *JobParams) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *JobParams) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *JobParams) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetEpoch

`func (o *JobParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *JobParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *JobParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


