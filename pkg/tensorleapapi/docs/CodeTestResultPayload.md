# CodeTestResultPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Display** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**IsPassed** | **bool** |  | 
**Shape** | Pointer to **[]float64** |  | [optional] 
**RawResult** | Pointer to **interface{}** |  | [optional] 
**HandlerType** | Pointer to **string** |  | [optional] 

## Methods

### NewCodeTestResultPayload

`func NewCodeTestResultPayload(name string, display map[string]interface{}, isPassed bool, ) *CodeTestResultPayload`

NewCodeTestResultPayload instantiates a new CodeTestResultPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeTestResultPayloadWithDefaults

`func NewCodeTestResultPayloadWithDefaults() *CodeTestResultPayload`

NewCodeTestResultPayloadWithDefaults instantiates a new CodeTestResultPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CodeTestResultPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CodeTestResultPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CodeTestResultPayload) SetName(v string)`

SetName sets Name field to given value.


### GetDisplay

`func (o *CodeTestResultPayload) GetDisplay() map[string]interface{}`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *CodeTestResultPayload) GetDisplayOk() (*map[string]interface{}, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *CodeTestResultPayload) SetDisplay(v map[string]interface{})`

SetDisplay sets Display field to given value.


### GetIsPassed

`func (o *CodeTestResultPayload) GetIsPassed() bool`

GetIsPassed returns the IsPassed field if non-nil, zero value otherwise.

### GetIsPassedOk

`func (o *CodeTestResultPayload) GetIsPassedOk() (*bool, bool)`

GetIsPassedOk returns a tuple with the IsPassed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPassed

`func (o *CodeTestResultPayload) SetIsPassed(v bool)`

SetIsPassed sets IsPassed field to given value.


### GetShape

`func (o *CodeTestResultPayload) GetShape() []float64`

GetShape returns the Shape field if non-nil, zero value otherwise.

### GetShapeOk

`func (o *CodeTestResultPayload) GetShapeOk() (*[]float64, bool)`

GetShapeOk returns a tuple with the Shape field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShape

`func (o *CodeTestResultPayload) SetShape(v []float64)`

SetShape sets Shape field to given value.

### HasShape

`func (o *CodeTestResultPayload) HasShape() bool`

HasShape returns a boolean if a field has been set.

### GetRawResult

`func (o *CodeTestResultPayload) GetRawResult() interface{}`

GetRawResult returns the RawResult field if non-nil, zero value otherwise.

### GetRawResultOk

`func (o *CodeTestResultPayload) GetRawResultOk() (*interface{}, bool)`

GetRawResultOk returns a tuple with the RawResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawResult

`func (o *CodeTestResultPayload) SetRawResult(v interface{})`

SetRawResult sets RawResult field to given value.

### HasRawResult

`func (o *CodeTestResultPayload) HasRawResult() bool`

HasRawResult returns a boolean if a field has been set.

### SetRawResultNil

`func (o *CodeTestResultPayload) SetRawResultNil(b bool)`

 SetRawResultNil sets the value for RawResult to be an explicit nil

### UnsetRawResult
`func (o *CodeTestResultPayload) UnsetRawResult()`

UnsetRawResult ensures that no value is present for RawResult, not even an explicit nil
### GetHandlerType

`func (o *CodeTestResultPayload) GetHandlerType() string`

GetHandlerType returns the HandlerType field if non-nil, zero value otherwise.

### GetHandlerTypeOk

`func (o *CodeTestResultPayload) GetHandlerTypeOk() (*string, bool)`

GetHandlerTypeOk returns a tuple with the HandlerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlerType

`func (o *CodeTestResultPayload) SetHandlerType(v string)`

SetHandlerType sets HandlerType field to given value.

### HasHandlerType

`func (o *CodeTestResultPayload) HasHandlerType() bool`

HasHandlerType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


