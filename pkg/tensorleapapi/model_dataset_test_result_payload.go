/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetTestResultPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetTestResultPayload{}

// DatasetTestResultPayload struct for DatasetTestResultPayload
type DatasetTestResultPayload struct {
	Name string `json:"name"`
	// Construct a type with a set of properties K of type T
	Display     map[string]interface{} `json:"display"`
	IsPassed    bool                   `json:"is_passed"`
	Shape       []float64              `json:"shape,omitempty"`
	RawResult   interface{}            `json:"raw_result,omitempty"`
	HandlerType *string                `json:"handler_type,omitempty"`
}

// NewDatasetTestResultPayload instantiates a new DatasetTestResultPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetTestResultPayload(name string, display map[string]interface{}, isPassed bool) *DatasetTestResultPayload {
	this := DatasetTestResultPayload{}
	this.Name = name
	this.Display = display
	this.IsPassed = isPassed
	return &this
}

// NewDatasetTestResultPayloadWithDefaults instantiates a new DatasetTestResultPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetTestResultPayloadWithDefaults() *DatasetTestResultPayload {
	this := DatasetTestResultPayload{}
	return &this
}

// GetName returns the Name field value
func (o *DatasetTestResultPayload) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DatasetTestResultPayload) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *DatasetTestResultPayload) SetName(v string) {
	o.Name = v
}

// GetDisplay returns the Display field value
func (o *DatasetTestResultPayload) GetDisplay() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *DatasetTestResultPayload) GetDisplayOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Display, true
}

// SetDisplay sets field value
func (o *DatasetTestResultPayload) SetDisplay(v map[string]interface{}) {
	o.Display = v
}

// GetIsPassed returns the IsPassed field value
func (o *DatasetTestResultPayload) GetIsPassed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsPassed
}

// GetIsPassedOk returns a tuple with the IsPassed field value
// and a boolean to check if the value has been set.
func (o *DatasetTestResultPayload) GetIsPassedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsPassed, true
}

// SetIsPassed sets field value
func (o *DatasetTestResultPayload) SetIsPassed(v bool) {
	o.IsPassed = v
}

// GetShape returns the Shape field value if set, zero value otherwise.
func (o *DatasetTestResultPayload) GetShape() []float64 {
	if o == nil || IsNil(o.Shape) {
		var ret []float64
		return ret
	}
	return o.Shape
}

// GetShapeOk returns a tuple with the Shape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetTestResultPayload) GetShapeOk() ([]float64, bool) {
	if o == nil || IsNil(o.Shape) {
		return nil, false
	}
	return o.Shape, true
}

// HasShape returns a boolean if a field has been set.
func (o *DatasetTestResultPayload) HasShape() bool {
	if o != nil && !IsNil(o.Shape) {
		return true
	}

	return false
}

// SetShape gets a reference to the given []float64 and assigns it to the Shape field.
func (o *DatasetTestResultPayload) SetShape(v []float64) {
	o.Shape = v
}

// GetRawResult returns the RawResult field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DatasetTestResultPayload) GetRawResult() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.RawResult
}

// GetRawResultOk returns a tuple with the RawResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DatasetTestResultPayload) GetRawResultOk() (*interface{}, bool) {
	if o == nil || IsNil(o.RawResult) {
		return nil, false
	}
	return &o.RawResult, true
}

// HasRawResult returns a boolean if a field has been set.
func (o *DatasetTestResultPayload) HasRawResult() bool {
	if o != nil && IsNil(o.RawResult) {
		return true
	}

	return false
}

// SetRawResult gets a reference to the given interface{} and assigns it to the RawResult field.
func (o *DatasetTestResultPayload) SetRawResult(v interface{}) {
	o.RawResult = v
}

// GetHandlerType returns the HandlerType field value if set, zero value otherwise.
func (o *DatasetTestResultPayload) GetHandlerType() string {
	if o == nil || IsNil(o.HandlerType) {
		var ret string
		return ret
	}
	return *o.HandlerType
}

// GetHandlerTypeOk returns a tuple with the HandlerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetTestResultPayload) GetHandlerTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HandlerType) {
		return nil, false
	}
	return o.HandlerType, true
}

// HasHandlerType returns a boolean if a field has been set.
func (o *DatasetTestResultPayload) HasHandlerType() bool {
	if o != nil && !IsNil(o.HandlerType) {
		return true
	}

	return false
}

// SetHandlerType gets a reference to the given string and assigns it to the HandlerType field.
func (o *DatasetTestResultPayload) SetHandlerType(v string) {
	o.HandlerType = &v
}

func (o DatasetTestResultPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetTestResultPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["display"] = o.Display
	toSerialize["is_passed"] = o.IsPassed
	if !IsNil(o.Shape) {
		toSerialize["shape"] = o.Shape
	}
	if o.RawResult != nil {
		toSerialize["raw_result"] = o.RawResult
	}
	if !IsNil(o.HandlerType) {
		toSerialize["handler_type"] = o.HandlerType
	}
	return toSerialize, nil
}

type NullableDatasetTestResultPayload struct {
	value *DatasetTestResultPayload
	isSet bool
}

func (v NullableDatasetTestResultPayload) Get() *DatasetTestResultPayload {
	return v.value
}

func (v *NullableDatasetTestResultPayload) Set(val *DatasetTestResultPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetTestResultPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetTestResultPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetTestResultPayload(val *DatasetTestResultPayload) *NullableDatasetTestResultPayload {
	return &NullableDatasetTestResultPayload{value: val, isSet: true}
}

func (v NullableDatasetTestResultPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetTestResultPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
