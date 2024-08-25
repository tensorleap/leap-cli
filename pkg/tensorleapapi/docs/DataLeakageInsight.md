# DataLeakageInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**ScatterInsightType**](ScatterInsightType.md) |  | 
**Filter** | [**ScatterFilter**](ScatterFilter.md) |  | 
**MutualInfoElements** | Pointer to [**[]MutualInformationElement**](MutualInformationElement.md) |  | [optional] 
**BlobPath** | Pointer to **string** |  | [optional] 
**InsightSubCategoryDsCuration** | Pointer to [**InsightSubCategoryDSCuration**](InsightSubCategoryDSCuration.md) |  | [optional] 
**InsightCategoryPerformance** | Pointer to [**InsightSubCategoryPerformance**](InsightSubCategoryPerformance.md) |  | [optional] 
**Index** | Pointer to **float64** |  | [optional] 
**FirstSubset** | [**DataStateType**](DataStateType.md) |  | 
**SecondSubset** | [**DataStateType**](DataStateType.md) |  | 
**NCrossSubsetsDuplications** | **float64** |  | 

## Methods

### NewDataLeakageInsight

`func NewDataLeakageInsight(type_ ScatterInsightType, filter ScatterFilter, firstSubset DataStateType, secondSubset DataStateType, nCrossSubsetsDuplications float64, ) *DataLeakageInsight`

NewDataLeakageInsight instantiates a new DataLeakageInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataLeakageInsightWithDefaults

`func NewDataLeakageInsightWithDefaults() *DataLeakageInsight`

NewDataLeakageInsightWithDefaults instantiates a new DataLeakageInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DataLeakageInsight) GetType() ScatterInsightType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DataLeakageInsight) GetTypeOk() (*ScatterInsightType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DataLeakageInsight) SetType(v ScatterInsightType)`

SetType sets Type field to given value.


### GetFilter

`func (o *DataLeakageInsight) GetFilter() ScatterFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DataLeakageInsight) GetFilterOk() (*ScatterFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DataLeakageInsight) SetFilter(v ScatterFilter)`

SetFilter sets Filter field to given value.


### GetMutualInfoElements

`func (o *DataLeakageInsight) GetMutualInfoElements() []MutualInformationElement`

GetMutualInfoElements returns the MutualInfoElements field if non-nil, zero value otherwise.

### GetMutualInfoElementsOk

`func (o *DataLeakageInsight) GetMutualInfoElementsOk() (*[]MutualInformationElement, bool)`

GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualInfoElements

`func (o *DataLeakageInsight) SetMutualInfoElements(v []MutualInformationElement)`

SetMutualInfoElements sets MutualInfoElements field to given value.

### HasMutualInfoElements

`func (o *DataLeakageInsight) HasMutualInfoElements() bool`

HasMutualInfoElements returns a boolean if a field has been set.

### GetBlobPath

`func (o *DataLeakageInsight) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *DataLeakageInsight) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *DataLeakageInsight) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.

### HasBlobPath

`func (o *DataLeakageInsight) HasBlobPath() bool`

HasBlobPath returns a boolean if a field has been set.

### GetInsightSubCategoryDsCuration

`func (o *DataLeakageInsight) GetInsightSubCategoryDsCuration() InsightSubCategoryDSCuration`

GetInsightSubCategoryDsCuration returns the InsightSubCategoryDsCuration field if non-nil, zero value otherwise.

### GetInsightSubCategoryDsCurationOk

`func (o *DataLeakageInsight) GetInsightSubCategoryDsCurationOk() (*InsightSubCategoryDSCuration, bool)`

GetInsightSubCategoryDsCurationOk returns a tuple with the InsightSubCategoryDsCuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightSubCategoryDsCuration

`func (o *DataLeakageInsight) SetInsightSubCategoryDsCuration(v InsightSubCategoryDSCuration)`

SetInsightSubCategoryDsCuration sets InsightSubCategoryDsCuration field to given value.

### HasInsightSubCategoryDsCuration

`func (o *DataLeakageInsight) HasInsightSubCategoryDsCuration() bool`

HasInsightSubCategoryDsCuration returns a boolean if a field has been set.

### GetInsightCategoryPerformance

`func (o *DataLeakageInsight) GetInsightCategoryPerformance() InsightSubCategoryPerformance`

GetInsightCategoryPerformance returns the InsightCategoryPerformance field if non-nil, zero value otherwise.

### GetInsightCategoryPerformanceOk

`func (o *DataLeakageInsight) GetInsightCategoryPerformanceOk() (*InsightSubCategoryPerformance, bool)`

GetInsightCategoryPerformanceOk returns a tuple with the InsightCategoryPerformance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightCategoryPerformance

`func (o *DataLeakageInsight) SetInsightCategoryPerformance(v InsightSubCategoryPerformance)`

SetInsightCategoryPerformance sets InsightCategoryPerformance field to given value.

### HasInsightCategoryPerformance

`func (o *DataLeakageInsight) HasInsightCategoryPerformance() bool`

HasInsightCategoryPerformance returns a boolean if a field has been set.

### GetIndex

`func (o *DataLeakageInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *DataLeakageInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *DataLeakageInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.

### HasIndex

`func (o *DataLeakageInsight) HasIndex() bool`

HasIndex returns a boolean if a field has been set.

### GetFirstSubset

`func (o *DataLeakageInsight) GetFirstSubset() DataStateType`

GetFirstSubset returns the FirstSubset field if non-nil, zero value otherwise.

### GetFirstSubsetOk

`func (o *DataLeakageInsight) GetFirstSubsetOk() (*DataStateType, bool)`

GetFirstSubsetOk returns a tuple with the FirstSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstSubset

`func (o *DataLeakageInsight) SetFirstSubset(v DataStateType)`

SetFirstSubset sets FirstSubset field to given value.


### GetSecondSubset

`func (o *DataLeakageInsight) GetSecondSubset() DataStateType`

GetSecondSubset returns the SecondSubset field if non-nil, zero value otherwise.

### GetSecondSubsetOk

`func (o *DataLeakageInsight) GetSecondSubsetOk() (*DataStateType, bool)`

GetSecondSubsetOk returns a tuple with the SecondSubset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondSubset

`func (o *DataLeakageInsight) SetSecondSubset(v DataStateType)`

SetSecondSubset sets SecondSubset field to given value.


### GetNCrossSubsetsDuplications

`func (o *DataLeakageInsight) GetNCrossSubsetsDuplications() float64`

GetNCrossSubsetsDuplications returns the NCrossSubsetsDuplications field if non-nil, zero value otherwise.

### GetNCrossSubsetsDuplicationsOk

`func (o *DataLeakageInsight) GetNCrossSubsetsDuplicationsOk() (*float64, bool)`

GetNCrossSubsetsDuplicationsOk returns a tuple with the NCrossSubsetsDuplications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNCrossSubsetsDuplications

`func (o *DataLeakageInsight) SetNCrossSubsetsDuplications(v float64)`

SetNCrossSubsetsDuplications sets NCrossSubsetsDuplications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


