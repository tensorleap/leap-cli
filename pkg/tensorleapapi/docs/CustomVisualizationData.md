# CustomVisualizationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**Name** | **string** |  | 
**Type** | [**CustomVisualizationType**](CustomVisualizationType.md) |  | 

## Methods

### NewCustomVisualizationData

`func NewCustomVisualizationData(data map[string]interface{}, name string, type_ CustomVisualizationType, ) *CustomVisualizationData`

NewCustomVisualizationData instantiates a new CustomVisualizationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomVisualizationDataWithDefaults

`func NewCustomVisualizationDataWithDefaults() *CustomVisualizationData`

NewCustomVisualizationDataWithDefaults instantiates a new CustomVisualizationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomVisualizationData) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomVisualizationData) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomVisualizationData) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetName

`func (o *CustomVisualizationData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomVisualizationData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomVisualizationData) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *CustomVisualizationData) GetType() CustomVisualizationType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomVisualizationData) GetTypeOk() (*CustomVisualizationType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomVisualizationData) SetType(v CustomVisualizationType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


