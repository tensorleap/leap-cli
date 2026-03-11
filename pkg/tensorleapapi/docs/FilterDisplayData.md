# FilterDisplayData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Insights** | [**[]InsightFilterDisplayDataInsightsInner**](InsightFilterDisplayDataInsightsInner.md) |  | 
**Limit** | **float64** |  | 
**SampleIds** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**FiltersUsed** | **[]string** |  | 
**Version** | [**FilterVersion**](FilterVersion.md) |  | 
**CollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewFilterDisplayData

`func NewFilterDisplayData(type_ string, insights []InsightFilterDisplayDataInsightsInner, limit float64, sampleIds []SampleIdentity, filtersUsed []string, version FilterVersion, collectionId string, projectId string, ) *FilterDisplayData`

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


### GetInsights

`func (o *FilterDisplayData) GetInsights() []InsightFilterDisplayDataInsightsInner`

GetInsights returns the Insights field if non-nil, zero value otherwise.

### GetInsightsOk

`func (o *FilterDisplayData) GetInsightsOk() (*[]InsightFilterDisplayDataInsightsInner, bool)`

GetInsightsOk returns a tuple with the Insights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsights

`func (o *FilterDisplayData) SetInsights(v []InsightFilterDisplayDataInsightsInner)`

SetInsights sets Insights field to given value.


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


### GetVersion

`func (o *FilterDisplayData) GetVersion() FilterVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *FilterDisplayData) GetVersionOk() (*FilterVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *FilterDisplayData) SetVersion(v FilterVersion)`

SetVersion sets Version field to given value.


### GetCollectionId

`func (o *FilterDisplayData) GetCollectionId() string`

GetCollectionId returns the CollectionId field if non-nil, zero value otherwise.

### GetCollectionIdOk

`func (o *FilterDisplayData) GetCollectionIdOk() (*string, bool)`

GetCollectionIdOk returns a tuple with the CollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionId

`func (o *FilterDisplayData) SetCollectionId(v string)`

SetCollectionId sets CollectionId field to given value.


### GetProjectId

`func (o *FilterDisplayData) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *FilterDisplayData) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *FilterDisplayData) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


