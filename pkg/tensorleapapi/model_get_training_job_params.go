/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetTrainingJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTrainingJobParams{}

// GetTrainingJobParams struct for GetTrainingJobParams
type GetTrainingJobParams struct {
	VersionId string `json:"versionId"`
}

// NewGetTrainingJobParams instantiates a new GetTrainingJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTrainingJobParams(versionId string) *GetTrainingJobParams {
	this := GetTrainingJobParams{}
	this.VersionId = versionId
	return &this
}

// NewGetTrainingJobParamsWithDefaults instantiates a new GetTrainingJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTrainingJobParamsWithDefaults() *GetTrainingJobParams {
	this := GetTrainingJobParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *GetTrainingJobParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *GetTrainingJobParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *GetTrainingJobParams) SetVersionId(v string) {
	o.VersionId = v
}

func (o GetTrainingJobParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTrainingJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	return toSerialize, nil
}

type NullableGetTrainingJobParams struct {
	value *GetTrainingJobParams
	isSet bool
}

func (v NullableGetTrainingJobParams) Get() *GetTrainingJobParams {
	return v.value
}

func (v *NullableGetTrainingJobParams) Set(val *GetTrainingJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTrainingJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTrainingJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTrainingJobParams(val *GetTrainingJobParams) *NullableGetTrainingJobParams {
	return &NullableGetTrainingJobParams{value: val, isSet: true}
}

func (v NullableGetTrainingJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTrainingJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
