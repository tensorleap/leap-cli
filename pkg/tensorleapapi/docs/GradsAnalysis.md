# GradsAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GradsItem**](GradsItem.md) |  | 
**IsLoss** | **bool** |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewGradsAnalysis

`func NewGradsAnalysis(data []GradsItem, isLoss bool, type_ DataTypeEnum, ) *GradsAnalysis`

NewGradsAnalysis instantiates a new GradsAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGradsAnalysisWithDefaults

`func NewGradsAnalysisWithDefaults() *GradsAnalysis`

NewGradsAnalysisWithDefaults instantiates a new GradsAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GradsAnalysis) GetData() []GradsItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GradsAnalysis) GetDataOk() (*[]GradsItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GradsAnalysis) SetData(v []GradsItem)`

SetData sets Data field to given value.


### GetIsLoss

`func (o *GradsAnalysis) GetIsLoss() bool`

GetIsLoss returns the IsLoss field if non-nil, zero value otherwise.

### GetIsLossOk

`func (o *GradsAnalysis) GetIsLossOk() (*bool, bool)`

GetIsLossOk returns a tuple with the IsLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLoss

`func (o *GradsAnalysis) SetIsLoss(v bool)`

SetIsLoss sets IsLoss field to given value.


### GetType

`func (o *GradsAnalysis) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GradsAnalysis) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GradsAnalysis) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


