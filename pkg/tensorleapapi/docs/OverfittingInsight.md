# OverfittingInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 
**NSamplesValidation** | **float64** |  | 
**NSamplesTraining** | **float64** |  | 
**AvgMetricValidation** | **float64** |  | 
**AvgMetricTraining** | **float64** |  | 
**MetricName** | **string** |  | 
**Configuration** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewOverfittingInsight

`func NewOverfittingInsight(type_ ScatterInsightType, filter ScatterFilter, nSamplesValidation float64, nSamplesTraining float64, avgMetricValidation float64, avgMetricTraining float64, metricName string, ) *OverfittingInsight`

NewOverfittingInsight instantiates a new OverfittingInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOverfittingInsightWithDefaults

`func NewOverfittingInsightWithDefaults() *OverfittingInsight`

NewOverfittingInsightWithDefaults instantiates a new OverfittingInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OverfittingInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OverfittingInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OverfittingInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilter

`func (o *OverfittingInsight) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *OverfittingInsight) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *OverfittingInsight) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


### GetMutualInfoElements

`func (o *OverfittingInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *OverfittingInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *OverfittingInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.

### HasMutualInfoElements

`func (o *OverfittingInsight) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

### GetBlobPath

`func (o *OverfittingInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *OverfittingInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *OverfittingInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.

### HasBlobPath

`func (o *OverfittingInsight) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.

### GetNSamplesValidation

`func (o *OverfittingInsight) GetNSamplesValidation() float64`

GetNSamplesValidation returns the NSamplesValidation field if non-nil, zero value otherwise.

### GetNSamplesValidationOk

`func (o *OverfittingInsight) GetNSamplesValidationOk() (*float64, bool)`

GetNSamplesValidationOk returns a tuple with the NSamplesValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesValidation

`func (o *OverfittingInsight) SetNSamplesValidation(v float64)`

SetNSamplesValidation sets NSamplesValidation field to given value.


### GetNSamplesTraining

`func (o *OverfittingInsight) GetNSamplesTraining() float64`

GetNSamplesTraining returns the NSamplesTraining field if non-nil, zero value otherwise.

### GetNSamplesTrainingOk

`func (o *OverfittingInsight) GetNSamplesTrainingOk() (*float64, bool)`

GetNSamplesTrainingOk returns a tuple with the NSamplesTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNSamplesTraining

`func (o *OverfittingInsight) SetNSamplesTraining(v float64)`

SetNSamplesTraining sets NSamplesTraining field to given value.


### GetAvgMetricValidation

`func (o *OverfittingInsight) GetAvgMetricValidation() float64`

GetAvgMetricValidation returns the AvgMetricValidation field if non-nil, zero value otherwise.

### GetAvgMetricValidationOk

`func (o *OverfittingInsight) GetAvgMetricValidationOk() (*float64, bool)`

GetAvgMetricValidationOk returns a tuple with the AvgMetricValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetricValidation

`func (o *OverfittingInsight) SetAvgMetricValidation(v float64)`

SetAvgMetricValidation sets AvgMetricValidation field to given value.


### GetAvgMetricTraining

`func (o *OverfittingInsight) GetAvgMetricTraining() float64`

GetAvgMetricTraining returns the AvgMetricTraining field if non-nil, zero value otherwise.

### GetAvgMetricTrainingOk

`func (o *OverfittingInsight) GetAvgMetricTrainingOk() (*float64, bool)`

GetAvgMetricTrainingOk returns a tuple with the AvgMetricTraining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgMetricTraining

`func (o *OverfittingInsight) SetAvgMetricTraining(v float64)`

SetAvgMetricTraining sets AvgMetricTraining field to given value.


### GetMetricName

`func (o *OverfittingInsight) GetMetricName() string`

GetMetricName returns the MetricName field if non-nil, zero value otherwise.

### GetMetricNameOk

`func (o *OverfittingInsight) GetMetricNameOk() (*string, bool)`

GetMetricNameOk returns a tuple with the MetricName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricName

`func (o *OverfittingInsight) SetMetricName(v string)`

SetMetricName sets MetricName field to given value.


### GetConfiguration

`func (o *OverfittingInsight) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *OverfittingInsight) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *OverfittingInsight) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *OverfittingInsight) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


