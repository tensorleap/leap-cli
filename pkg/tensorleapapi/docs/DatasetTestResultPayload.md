# DatasetTestResultPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Display** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**IsPassed** | **bool** |  | 
**Shape** | Pointer to **[]float64** |  | [optional] 

## Methods

### NewDatasetTestResultPayload

`func NewDatasetTestResultPayload(name string, display map[string]interface{}, isPassed bool, ) *DatasetTestResultPayload`

NewDatasetTestResultPayload instantiates a new DatasetTestResultPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetTestResultPayloadWithDefaults

`func NewDatasetTestResultPayloadWithDefaults() *DatasetTestResultPayload`

NewDatasetTestResultPayloadWithDefaults instantiates a new DatasetTestResultPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DatasetTestResultPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DatasetTestResultPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DatasetTestResultPayload) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *DatasetTestResultPayload) GetDisplay() map[string]interface{}`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *DatasetTestResultPayload) GetDisplayOk() (*map[string]interface{}, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *DatasetTestResultPayload) SetDisplay(v map[string]interface{})`

SetDisplay sets Display field to given value.


### GetIsPassed

`func (o *DatasetTestResultPayload) GetIsPassed() bool`

GetIsPassed returns the IsPassed field if non-nil, zero value otherwise.

### GetIsPassedOk

`func (o *DatasetTestResultPayload) GetIsPassedOk() (*bool, bool)`

GetIsPassedOk returns a tuple with the IsPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassed

`func (o *DatasetTestResultPayload) SetIsPassed(v bool)`

SetIsPassed sets IsPassed field to given value.


### GetShape

`func (o *DatasetTestResultPayload) GetShape() []float64`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *DatasetTestResultPayload) GetShapeOk() (*[]float64, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *DatasetTestResultPayload) SetShape(v []float64)`

SetShape sets Shape field to given value.

### HasShape

`func (o *DatasetTestResultPayload) HasShape() bool`

HasShape returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


