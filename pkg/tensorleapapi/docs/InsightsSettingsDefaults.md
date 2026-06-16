# InsightsSettingsDefaults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | [**[]MetricDefault**](MetricDefault.md) |  | 
**Domains** | **[]string** |  | 
**AvailableMetadata** | **[]string** |  | 
**LatentSpaces** | [**[]LatentSpaceDefault**](LatentSpaceDefault.md) |  | 

## Methods

### NewInsightsSettingsDefaults

`func NewInsightsSettingsDefaults(metrics []MetricDefault, domains []string, availableMetadata []string, latentSpaces []LatentSpaceDefault, ) *InsightsSettingsDefaults`

NewInsightsSettingsDefaults instantiates a new InsightsSettingsDefaults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsSettingsDefaultsWithDefaults

`func NewInsightsSettingsDefaultsWithDefaults() *InsightsSettingsDefaults`

NewInsightsSettingsDefaultsWithDefaults instantiates a new InsightsSettingsDefaults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *InsightsSettingsDefaults) GetMetrics() []MetricDefault`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InsightsSettingsDefaults) GetMetricsOk() (*[]MetricDefault, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InsightsSettingsDefaults) SetMetrics(v []MetricDefault)`

SetMetrics sets Metrics field to given value.


### GetDomains

`func (o *InsightsSettingsDefaults) GetDomains() []string`

GetDomains returns the Domains field if non-nil, zero value otherwise.

### GetDomainsOk

`func (o *InsightsSettingsDefaults) GetDomainsOk() (*[]string, bool)`

GetDomainsOk returns a tuple with the Domains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomains

`func (o *InsightsSettingsDefaults) SetDomains(v []string)`

SetDomains sets Domains field to given value.


### GetAvailableMetadata

`func (o *InsightsSettingsDefaults) GetAvailableMetadata() []string`

GetAvailableMetadata returns the AvailableMetadata field if non-nil, zero value otherwise.

### GetAvailableMetadataOk

`func (o *InsightsSettingsDefaults) GetAvailableMetadataOk() (*[]string, bool)`

GetAvailableMetadataOk returns a tuple with the AvailableMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableMetadata

`func (o *InsightsSettingsDefaults) SetAvailableMetadata(v []string)`

SetAvailableMetadata sets AvailableMetadata field to given value.


### GetLatentSpaces

`func (o *InsightsSettingsDefaults) GetLatentSpaces() []LatentSpaceDefault`

GetLatentSpaces returns the LatentSpaces field if non-nil, zero value otherwise.

### GetLatentSpacesOk

`func (o *InsightsSettingsDefaults) GetLatentSpacesOk() (*[]LatentSpaceDefault, bool)`

GetLatentSpacesOk returns a tuple with the LatentSpaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaces

`func (o *InsightsSettingsDefaults) SetLatentSpaces(v []LatentSpaceDefault)`

SetLatentSpaces sets LatentSpaces field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


