# GeneratedDashlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Priority** | **float64** |  | 
**Type** | [**GeneratedDashletType**](GeneratedDashletType.md) |  | 
**Data** | [**GeneratedChartData**](GeneratedChartData.md) |  | 
**Name** | Pointer to **string** |  | [optional] 

## Methods

### NewGeneratedDashlet

`func NewGeneratedDashlet(priority float64, type_ GeneratedDashletType, data GeneratedChartData, ) *GeneratedDashlet`

NewGeneratedDashlet instantiates a new GeneratedDashlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneratedDashletWithDefaults

`func NewGeneratedDashletWithDefaults() *GeneratedDashlet`

NewGeneratedDashletWithDefaults instantiates a new GeneratedDashlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriority

`func (o *GeneratedDashlet) GetPriority() float64`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GeneratedDashlet) GetPriorityOk() (*float64, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GeneratedDashlet) SetPriority(v float64)`

SetPriority sets Priority field to given value.


### GetType

`func (o *GeneratedDashlet) GetType() GeneratedDashletType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GeneratedDashlet) GetTypeOk() (*GeneratedDashletType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GeneratedDashlet) SetType(v GeneratedDashletType)`

SetType sets Type field to given value.


### GetData

`func (o *GeneratedDashlet) GetData() GeneratedChartData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GeneratedDashlet) GetDataOk() (*GeneratedChartData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GeneratedDashlet) SetData(v GeneratedChartData)`

SetData sets Data field to given value.


### GetName

`func (o *GeneratedDashlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GeneratedDashlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GeneratedDashlet) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GeneratedDashlet) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


