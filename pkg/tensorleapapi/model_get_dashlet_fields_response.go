/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDashletFieldsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDashletFieldsResponse{}

// GetDashletFieldsResponse struct for GetDashletFieldsResponse
type GetDashletFieldsResponse struct {
	AggregatableFields []string `json:"aggregatableFields"`
	NumericFields      []string `json:"numericFields"`
	StringFields       []string `json:"stringFields"`
}

// NewGetDashletFieldsResponse instantiates a new GetDashletFieldsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDashletFieldsResponse(aggregatableFields []string, numericFields []string, stringFields []string) *GetDashletFieldsResponse {
	this := GetDashletFieldsResponse{}
	this.AggregatableFields = aggregatableFields
	this.NumericFields = numericFields
	this.StringFields = stringFields
	return &this
}

// NewGetDashletFieldsResponseWithDefaults instantiates a new GetDashletFieldsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDashletFieldsResponseWithDefaults() *GetDashletFieldsResponse {
	this := GetDashletFieldsResponse{}
	return &this
}

// GetAggregatableFields returns the AggregatableFields field value
func (o *GetDashletFieldsResponse) GetAggregatableFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AggregatableFields
}

// GetAggregatableFieldsOk returns a tuple with the AggregatableFields field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsResponse) GetAggregatableFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AggregatableFields, true
}

// SetAggregatableFields sets field value
func (o *GetDashletFieldsResponse) SetAggregatableFields(v []string) {
	o.AggregatableFields = v
}

// GetNumericFields returns the NumericFields field value
func (o *GetDashletFieldsResponse) GetNumericFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.NumericFields
}

// GetNumericFieldsOk returns a tuple with the NumericFields field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsResponse) GetNumericFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.NumericFields, true
}

// SetNumericFields sets field value
func (o *GetDashletFieldsResponse) SetNumericFields(v []string) {
	o.NumericFields = v
}

// GetStringFields returns the StringFields field value
func (o *GetDashletFieldsResponse) GetStringFields() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.StringFields
}

// GetStringFieldsOk returns a tuple with the StringFields field value
// and a boolean to check if the value has been set.
func (o *GetDashletFieldsResponse) GetStringFieldsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.StringFields, true
}

// SetStringFields sets field value
func (o *GetDashletFieldsResponse) SetStringFields(v []string) {
	o.StringFields = v
}

func (o GetDashletFieldsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDashletFieldsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["aggregatableFields"] = o.AggregatableFields
	toSerialize["numericFields"] = o.NumericFields
	toSerialize["stringFields"] = o.StringFields
	return toSerialize, nil
}

type NullableGetDashletFieldsResponse struct {
	value *GetDashletFieldsResponse
	isSet bool
}

func (v NullableGetDashletFieldsResponse) Get() *GetDashletFieldsResponse {
	return v.value
}

func (v *NullableGetDashletFieldsResponse) Set(val *GetDashletFieldsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDashletFieldsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDashletFieldsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDashletFieldsResponse(val *GetDashletFieldsResponse) *NullableGetDashletFieldsResponse {
	return &NullableGetDashletFieldsResponse{value: val, isSet: true}
}

func (v NullableGetDashletFieldsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDashletFieldsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
