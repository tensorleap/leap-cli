# CompositeVizData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]CompositeVizItem**](CompositeVizItem.md) |  | 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewCompositeVizData

`func NewCompositeVizData(data []CompositeVizItem, type_ DataTypeEnum, ) *CompositeVizData`

NewCompositeVizData instantiates a new CompositeVizData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompositeVizDataWithDefaults

`func NewCompositeVizDataWithDefaults() *CompositeVizData`

NewCompositeVizDataWithDefaults instantiates a new CompositeVizData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CompositeVizData) GetData() []CompositeVizItem`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CompositeVizData) GetDataOk() (*[]CompositeVizItem, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CompositeVizData) SetData(v []CompositeVizItem)`

SetData sets Data field to given value.


### GetType

`func (o *CompositeVizData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CompositeVizData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CompositeVizData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


