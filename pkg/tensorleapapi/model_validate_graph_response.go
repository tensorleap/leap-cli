/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ValidateGraphResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidateGraphResponse{}

// ValidateGraphResponse struct for ValidateGraphResponse
type ValidateGraphResponse struct {
	Digest string              `json:"digest"`
	Data   *GraphValidatorData `json:"data,omitempty"`
	Status JobStatus           `json:"status"`
}

// NewValidateGraphResponse instantiates a new ValidateGraphResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidateGraphResponse(digest string, status JobStatus) *ValidateGraphResponse {
	this := ValidateGraphResponse{}
	this.Digest = digest
	this.Status = status
	return &this
}

// NewValidateGraphResponseWithDefaults instantiates a new ValidateGraphResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidateGraphResponseWithDefaults() *ValidateGraphResponse {
	this := ValidateGraphResponse{}
	return &this
}

// GetDigest returns the Digest field value
func (o *ValidateGraphResponse) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *ValidateGraphResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *ValidateGraphResponse) SetDigest(v string) {
	o.Digest = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ValidateGraphResponse) GetData() GraphValidatorData {
	if o == nil || IsNil(o.Data) {
		var ret GraphValidatorData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidateGraphResponse) GetDataOk() (*GraphValidatorData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ValidateGraphResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GraphValidatorData and assigns it to the Data field.
func (o *ValidateGraphResponse) SetData(v GraphValidatorData) {
	o.Data = &v
}

// GetStatus returns the Status field value
func (o *ValidateGraphResponse) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ValidateGraphResponse) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ValidateGraphResponse) SetStatus(v JobStatus) {
	o.Status = v
}

func (o ValidateGraphResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidateGraphResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digest"] = o.Digest
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableValidateGraphResponse struct {
	value *ValidateGraphResponse
	isSet bool
}

func (v NullableValidateGraphResponse) Get() *ValidateGraphResponse {
	return v.value
}

func (v *NullableValidateGraphResponse) Set(val *ValidateGraphResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableValidateGraphResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableValidateGraphResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidateGraphResponse(val *ValidateGraphResponse) *NullableValidateGraphResponse {
	return &NullableValidateGraphResponse{value: val, isSet: true}
}

func (v NullableValidateGraphResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidateGraphResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
