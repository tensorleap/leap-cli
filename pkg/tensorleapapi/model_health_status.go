/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the HealthStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HealthStatus{}

// HealthStatus struct for HealthStatus
type HealthStatus struct {
	Status bool `json:"status"`
	Error *string `json:"error,omitempty"`
}

// NewHealthStatus instantiates a new HealthStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHealthStatus(status bool) *HealthStatus {
	this := HealthStatus{}
	this.Status = status
	return &this
}

// NewHealthStatusWithDefaults instantiates a new HealthStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHealthStatusWithDefaults() *HealthStatus {
	this := HealthStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *HealthStatus) GetStatus() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *HealthStatus) GetStatusOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *HealthStatus) SetStatus(v bool) {
	o.Status = v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *HealthStatus) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HealthStatus) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *HealthStatus) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *HealthStatus) SetError(v string) {
	o.Error = &v
}

func (o HealthStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HealthStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableHealthStatus struct {
	value *HealthStatus
	isSet bool
}

func (v NullableHealthStatus) Get() *HealthStatus {
	return v.value
}

func (v *NullableHealthStatus) Set(val *HealthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableHealthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableHealthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHealthStatus(val *HealthStatus) *NullableHealthStatus {
	return &NullableHealthStatus{value: val, isSet: true}
}

func (v NullableHealthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHealthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


