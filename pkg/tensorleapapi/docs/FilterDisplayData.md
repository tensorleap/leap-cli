# FilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**InsightData** | [**InsightType**](InsightType.md) |  | 
**SessionRun** | [**FilterSessionRun**](FilterSessionRun.md) |  | 
**Limit** | **float64** |  | 
**SampleIds** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**Epoch** | **float64** |  | 
**FiltersUsed** | **[]string** |  | 

## Methods

### NewFilterDisplayData

`func NewFilterDisplayData(type_ string, insightData InsightType, sessionRun FilterSessionRun, limit float64, sampleIds []SampleIdentity, epoch float64, filtersUsed []string, ) *FilterDisplayData`

NewFilterDisplayData instantiates a new FilterDisplayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterDisplayDataWithDefaults

`func NewFilterDisplayDataWithDefaults() *FilterDisplayData`

NewFilterDisplayDataWithDefaults instantiates a new FilterDisplayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FilterDisplayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FilterDisplayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FilterDisplayData) SetType(v string)`

SetType sets Type field to given value.


### GetInsightData

`func (o *FilterDisplayData) GetInsightData() InsightType`

GetInsightData returns the InsightData field if non-nil, zero value otherwise.

### GetInsightDataOk

`func (o *FilterDisplayData) GetInsightDataOk() (*InsightType, bool)`

GetInsightDataOk returns a tuple with the InsightData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightData

`func (o *FilterDisplayData) SetInsightData(v InsightType)`

SetInsightData sets InsightData field to given value.


### GetSessionRun

`func (o *FilterDisplayData) GetSessionRun() FilterSessionRun`

GetSessionRun returns the SessionRun field if non-nil, zero value otherwise.

### GetSessionRunOk

`func (o *FilterDisplayData) GetSessionRunOk() (*FilterSessionRun, bool)`

GetSessionRunOk returns a tuple with the SessionRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRun

`func (o *FilterDisplayData) SetSessionRun(v FilterSessionRun)`

SetSessionRun sets SessionRun field to given value.


### GetLimit

`func (o *FilterDisplayData) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FilterDisplayData) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FilterDisplayData) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### GetSampleIds

`func (o *FilterDisplayData) GetSampleIds() []SampleIdentity`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *FilterDisplayData) GetSampleIdsOk() (*[]SampleIdentity, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *FilterDisplayData) SetSampleIds(v []SampleIdentity)`

SetSampleIds sets SampleIds field to given value.


### GetEpoch

`func (o *FilterDisplayData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *FilterDisplayData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *FilterDisplayData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetFiltersUsed

`func (o *FilterDisplayData) GetFiltersUsed() []string`

GetFiltersUsed returns the FiltersUsed field if non-nil, zero value otherwise.

### GetFiltersUsedOk

`func (o *FilterDisplayData) GetFiltersUsedOk() (*[]string, bool)`

GetFiltersUsedOk returns a tuple with the FiltersUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersUsed

`func (o *FilterDisplayData) SetFiltersUsed(v []string)`

SetFiltersUsed sets FiltersUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


