/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the EarlyStopParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EarlyStopParams{}

// EarlyStopParams struct for EarlyStopParams
type EarlyStopParams struct {
	Patience float64 `json:"patience"`
}

// NewEarlyStopParams instantiates a new EarlyStopParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEarlyStopParams(patience float64) *EarlyStopParams {
	this := EarlyStopParams{}
	this.Patience = patience
	return &this
}

// NewEarlyStopParamsWithDefaults instantiates a new EarlyStopParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEarlyStopParamsWithDefaults() *EarlyStopParams {
	this := EarlyStopParams{}
	return &this
}

// GetPatience returns the Patience field value
func (o *EarlyStopParams) GetPatience() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Patience
}

// GetPatienceOk returns a tuple with the Patience field value
// and a boolean to check if the value has been set.
func (o *EarlyStopParams) GetPatienceOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Patience, true
}

// SetPatience sets field value
func (o *EarlyStopParams) SetPatience(v float64) {
	o.Patience = v
}

func (o EarlyStopParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EarlyStopParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["patience"] = o.Patience
	return toSerialize, nil
}

type NullableEarlyStopParams struct {
	value *EarlyStopParams
	isSet bool
}

func (v NullableEarlyStopParams) Get() *EarlyStopParams {
	return v.value
}

func (v *NullableEarlyStopParams) Set(val *EarlyStopParams) {
	v.value = val
	v.isSet = true
}

func (v NullableEarlyStopParams) IsSet() bool {
	return v.isSet
}

func (v *NullableEarlyStopParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEarlyStopParams(val *EarlyStopParams) *NullableEarlyStopParams {
	return &NullableEarlyStopParams{value: val, isSet: true}
}

func (v NullableEarlyStopParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEarlyStopParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
