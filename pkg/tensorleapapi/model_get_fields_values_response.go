/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetFieldsValuesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetFieldsValuesResponse{}

// GetFieldsValuesResponse struct for GetFieldsValuesResponse
type GetFieldsValuesResponse struct {
	Results []GetFieldsValuesResponseResultsInner `json:"results"`
}

// NewGetFieldsValuesResponse instantiates a new GetFieldsValuesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetFieldsValuesResponse(results []GetFieldsValuesResponseResultsInner) *GetFieldsValuesResponse {
	this := GetFieldsValuesResponse{}
	this.Results = results
	return &this
}

// NewGetFieldsValuesResponseWithDefaults instantiates a new GetFieldsValuesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetFieldsValuesResponseWithDefaults() *GetFieldsValuesResponse {
	this := GetFieldsValuesResponse{}
	return &this
}

// GetResults returns the Results field value
func (o *GetFieldsValuesResponse) GetResults() []GetFieldsValuesResponseResultsInner {
	if o == nil {
		var ret []GetFieldsValuesResponseResultsInner
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *GetFieldsValuesResponse) GetResultsOk() ([]GetFieldsValuesResponseResultsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *GetFieldsValuesResponse) SetResults(v []GetFieldsValuesResponseResultsInner) {
	o.Results = v
}

func (o GetFieldsValuesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetFieldsValuesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

type NullableGetFieldsValuesResponse struct {
	value *GetFieldsValuesResponse
	isSet bool
}

func (v NullableGetFieldsValuesResponse) Get() *GetFieldsValuesResponse {
	return v.value
}

func (v *NullableGetFieldsValuesResponse) Set(val *GetFieldsValuesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetFieldsValuesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetFieldsValuesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetFieldsValuesResponse(val *GetFieldsValuesResponse) *NullableGetFieldsValuesResponse {
	return &NullableGetFieldsValuesResponse{value: val, isSet: true}
}

func (v NullableGetFieldsValuesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetFieldsValuesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
