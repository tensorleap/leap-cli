# FetchSimilarJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Epoch** | **float64** |  | 
**SampleIds** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**Limit** | **float64** |  | 
**SessionRunId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewFetchSimilarJobParams

`func NewFetchSimilarJobParams(digest string, epoch float64, sampleIds []SampleIdentity, limit float64, sessionRunId string, type_ string, ) *FetchSimilarJobParams`

NewFetchSimilarJobParams instantiates a new FetchSimilarJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSimilarJobParamsWithDefaults

`func NewFetchSimilarJobParamsWithDefaults() *FetchSimilarJobParams`

NewFetchSimilarJobParamsWithDefaults instantiates a new FetchSimilarJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *FetchSimilarJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *FetchSimilarJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *FetchSimilarJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFilters

`func (o *FetchSimilarJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *FetchSimilarJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *FetchSimilarJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *FetchSimilarJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetEpoch

`func (o *FetchSimilarJobParams) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *FetchSimilarJobParams) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *FetchSimilarJobParams) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetSampleIds

`func (o *FetchSimilarJobParams) GetSampleIds() []SampleIdentity`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *FetchSimilarJobParams) GetSampleIdsOk() (*[]SampleIdentity, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *FetchSimilarJobParams) SetSampleIds(v []SampleIdentity)`

SetSampleIds sets SampleIds field to given value.


### GetLimit

`func (o *FetchSimilarJobParams) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FetchSimilarJobParams) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FetchSimilarJobParams) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### GetSessionRunId

`func (o *FetchSimilarJobParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *FetchSimilarJobParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *FetchSimilarJobParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetType

`func (o *FetchSimilarJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchSimilarJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchSimilarJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


