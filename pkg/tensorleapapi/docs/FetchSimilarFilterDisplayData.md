# FetchSimilarFilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Limit** | **float64** |  | 
**SampleIds** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**FiltersUsed** | **[]string** |  | 
**Version** | [**FilterVersion**](FilterVersion.md) |  | 

## Methods

### NewFetchSimilarFilterDisplayData

`func NewFetchSimilarFilterDisplayData(type_ string, limit float64, sampleIds []SampleIdentity, filtersUsed []string, version FilterVersion, ) *FetchSimilarFilterDisplayData`

NewFetchSimilarFilterDisplayData instantiates a new FetchSimilarFilterDisplayData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchSimilarFilterDisplayDataWithDefaults

`func NewFetchSimilarFilterDisplayDataWithDefaults() *FetchSimilarFilterDisplayData`

NewFetchSimilarFilterDisplayDataWithDefaults instantiates a new FetchSimilarFilterDisplayData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FetchSimilarFilterDisplayData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FetchSimilarFilterDisplayData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FetchSimilarFilterDisplayData) SetType(v string)`

SetType sets Type field to given value.


### GetLimit

`func (o *FetchSimilarFilterDisplayData) GetLimit() float64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *FetchSimilarFilterDisplayData) GetLimitOk() (*float64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *FetchSimilarFilterDisplayData) SetLimit(v float64)`

SetLimit sets Limit field to given value.


### GetSampleIds

`func (o *FetchSimilarFilterDisplayData) GetSampleIds() []SampleIdentity`

GetSampleIds returns the SampleIds field if non-nil, zero value otherwise.

### GetSampleIdsOk

`func (o *FetchSimilarFilterDisplayData) GetSampleIdsOk() (*[]SampleIdentity, bool)`

GetSampleIdsOk returns a tuple with the SampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleIds

`func (o *FetchSimilarFilterDisplayData) SetSampleIds(v []SampleIdentity)`

SetSampleIds sets SampleIds field to given value.


### GetFiltersUsed

`func (o *FetchSimilarFilterDisplayData) GetFiltersUsed() []string`

GetFiltersUsed returns the FiltersUsed field if non-nil, zero value otherwise.

### GetFiltersUsedOk

`func (o *FetchSimilarFilterDisplayData) GetFiltersUsedOk() (*[]string, bool)`

GetFiltersUsedOk returns a tuple with the FiltersUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltersUsed

`func (o *FetchSimilarFilterDisplayData) SetFiltersUsed(v []string)`

SetFiltersUsed sets FiltersUsed field to given value.


### GetVersion

`func (o *FetchSimilarFilterDisplayData) GetVersion() FilterVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FetchSimilarFilterDisplayData) GetVersionOk() (*FilterVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FetchSimilarFilterDisplayData) SetVersion(v FilterVersion)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


