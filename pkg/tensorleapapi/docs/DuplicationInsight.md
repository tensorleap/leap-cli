# DuplicationInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 
**NDuplicateSamples** | **float64** |  | 
**NCrossSubsetDuplications** | **float64** |  | 

## Methods

### NewDuplicationInsight

`func NewDuplicationInsight(type_ ScatterInsightType, filter ScatterFilter, nDuplicateSamples float64, nCrossSubsetDuplications float64, ) *DuplicationInsight`

NewDuplicationInsight instantiates a new DuplicationInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDuplicationInsightWithDefaults

`func NewDuplicationInsightWithDefaults() *DuplicationInsight`

NewDuplicationInsightWithDefaults instantiates a new DuplicationInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DuplicationInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DuplicationInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DuplicationInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilter

`func (o *DuplicationInsight) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DuplicationInsight) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DuplicationInsight) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


### GetMutualInfoElements

`func (o *DuplicationInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *DuplicationInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *DuplicationInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.

### HasMutualInfoElements

`func (o *DuplicationInsight) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

### GetBlobPath

`func (o *DuplicationInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *DuplicationInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *DuplicationInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.

### HasBlobPath

`func (o *DuplicationInsight) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.

### GetNDuplicateSamples

`func (o *DuplicationInsight) GetNDuplicateSamples() float64`

GetNDuplicateSamples returns the NDuplicateSamples field if non-nil, zero value otherwise.

### GetNDuplicateSamplesOk

`func (o *DuplicationInsight) GetNDuplicateSamplesOk() (*float64, bool)`

GetNDuplicateSamplesOk returns a tuple with the NDuplicateSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNDuplicateSamples

`func (o *DuplicationInsight) SetNDuplicateSamples(v float64)`

SetNDuplicateSamples sets NDuplicateSamples field to given value.


### GetNCrossSubsetDuplications

`func (o *DuplicationInsight) GetNCrossSubsetDuplications() float64`

GetNCrossSubsetDuplications returns the NCrossSubsetDuplications field if non-nil, zero value otherwise.

### GetNCrossSubsetDuplicationsOk

`func (o *DuplicationInsight) GetNCrossSubsetDuplicationsOk() (*float64, bool)`

GetNCrossSubsetDuplicationsOk returns a tuple with the NCrossSubsetDuplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCrossSubsetDuplications

`func (o *DuplicationInsight) SetNCrossSubsetDuplications(v float64)`

SetNCrossSubsetDuplications sets NCrossSubsetDuplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


