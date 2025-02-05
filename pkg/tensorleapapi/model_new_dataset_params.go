/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the NewDatasetParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NewDatasetParams{}

// NewDatasetParams struct for NewDatasetParams
type NewDatasetParams struct {
	Name NullableString `json:"name"`
}

// NewNewDatasetParams instantiates a new NewDatasetParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDatasetParams(name NullableString) *NewDatasetParams {
	this := NewDatasetParams{}
	this.Name = name
	return &this
}

// NewNewDatasetParamsWithDefaults instantiates a new NewDatasetParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDatasetParamsWithDefaults() *NewDatasetParams {
	this := NewDatasetParams{}
	return &this
}

// GetName returns the Name field value
// If the value is explicit nil, the zero value for string will be returned
func (o *NewDatasetParams) GetName() string {
	if o == nil || o.Name.Get() == nil {
		var ret string
		return ret
	}

	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NewDatasetParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// SetName sets field value
func (o *NewDatasetParams) SetName(v string) {
	o.Name.Set(&v)
}

func (o NewDatasetParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NewDatasetParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name.Get()
	return toSerialize, nil
}

type NullableNewDatasetParams struct {
	value *NewDatasetParams
	isSet bool
}

func (v NullableNewDatasetParams) Get() *NewDatasetParams {
	return v.value
}

func (v *NullableNewDatasetParams) Set(val *NewDatasetParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDatasetParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDatasetParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDatasetParams(val *NewDatasetParams) *NullableNewDatasetParams {
	return &NullableNewDatasetParams{value: val, isSet: true}
}

func (v NullableNewDatasetParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDatasetParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
