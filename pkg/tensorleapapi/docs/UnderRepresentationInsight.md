# UnderRepresentationInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**UnderRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**UnderRepresentationNSamples** | **float64** |  | 
**OverRepresentationDataset** | [**DataStateType**](DataStateType.md) |  | 
**OverRepresentationNSamples** | **float64** |  | 

## Methods

### NewUnderRepresentationInsight

`func NewUnderRepresentationInsight(type_ ScatterInsightType, filter ScatterFilter, underRepresentationDataset DataStateType, underRepresentationNSamples float64, overRepresentationDataset DataStateType, overRepresentationNSamples float64, ) *UnderRepresentationInsight`

NewUnderRepresentationInsight instantiates a new UnderRepresentationInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnderRepresentationInsightWithDefaults

`func NewUnderRepresentationInsightWithDefaults() *UnderRepresentationInsight`

NewUnderRepresentationInsightWithDefaults instantiates a new UnderRepresentationInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *UnderRepresentationInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnderRepresentationInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnderRepresentationInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilter

`func (o *UnderRepresentationInsight) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *UnderRepresentationInsight) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *UnderRepresentationInsight) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


### GetMutualInfoElements

`func (o *UnderRepresentationInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *UnderRepresentationInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *UnderRepresentationInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.

### HasMutualInfoElements

`func (o *UnderRepresentationInsight) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

### GetUnderRepresentationDataset

`func (o *UnderRepresentationInsight) GetUnderRepresentationDataset() DataStateType`

GetUnderRepresentationDataset returns the UnderRepresentationDataset field if non-nil, zero value otherwise.

### GetUnderRepresentationDatasetOk

`func (o *UnderRepresentationInsight) GetUnderRepresentationDatasetOk() (*DataStateType, bool)`

GetUnderRepresentationDatasetOk returns a tuple with the UnderRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationDataset

`func (o *UnderRepresentationInsight) SetUnderRepresentationDataset(v DataStateType)`

SetUnderRepresentationDataset sets UnderRepresentationDataset field to given value.


### GetUnderRepresentationNSamples

`func (o *UnderRepresentationInsight) GetUnderRepresentationNSamples() float64`

GetUnderRepresentationNSamples returns the UnderRepresentationNSamples field if non-nil, zero value otherwise.

### GetUnderRepresentationNSamplesOk

`func (o *UnderRepresentationInsight) GetUnderRepresentationNSamplesOk() (*float64, bool)`

GetUnderRepresentationNSamplesOk returns a tuple with the UnderRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderRepresentationNSamples

`func (o *UnderRepresentationInsight) SetUnderRepresentationNSamples(v float64)`

SetUnderRepresentationNSamples sets UnderRepresentationNSamples field to given value.


### GetOverRepresentationDataset

`func (o *UnderRepresentationInsight) GetOverRepresentationDataset() DataStateType`

GetOverRepresentationDataset returns the OverRepresentationDataset field if non-nil, zero value otherwise.

### GetOverRepresentationDatasetOk

`func (o *UnderRepresentationInsight) GetOverRepresentationDatasetOk() (*DataStateType, bool)`

GetOverRepresentationDatasetOk returns a tuple with the OverRepresentationDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationDataset

`func (o *UnderRepresentationInsight) SetOverRepresentationDataset(v DataStateType)`

SetOverRepresentationDataset sets OverRepresentationDataset field to given value.


### GetOverRepresentationNSamples

`func (o *UnderRepresentationInsight) GetOverRepresentationNSamples() float64`

GetOverRepresentationNSamples returns the OverRepresentationNSamples field if non-nil, zero value otherwise.

### GetOverRepresentationNSamplesOk

`func (o *UnderRepresentationInsight) GetOverRepresentationNSamplesOk() (*float64, bool)`

GetOverRepresentationNSamplesOk returns a tuple with the OverRepresentationNSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverRepresentationNSamples

`func (o *UnderRepresentationInsight) SetOverRepresentationNSamples(v float64)`

SetOverRepresentationNSamples sets OverRepresentationNSamples field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


