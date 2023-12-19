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

// checks if the ActionTagElement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ActionTagElement{}

// ActionTagElement struct for ActionTagElement
type ActionTagElement struct {
	Tags []string `json:"tags"`
}

// NewActionTagElement instantiates a new ActionTagElement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActionTagElement(tags []string) *ActionTagElement {
	this := ActionTagElement{}
	this.Tags = tags
	return &this
}

// NewActionTagElementWithDefaults instantiates a new ActionTagElement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionTagElementWithDefaults() *ActionTagElement {
	this := ActionTagElement{}
	return &this
}

// GetTags returns the Tags field value
func (o *ActionTagElement) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ActionTagElement) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ActionTagElement) SetTags(v []string) {
	o.Tags = v
}

func (o ActionTagElement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ActionTagElement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tags"] = o.Tags
	return toSerialize, nil
}

type NullableActionTagElement struct {
	value *ActionTagElement
	isSet bool
}

func (v NullableActionTagElement) Get() *ActionTagElement {
	return v.value
}

func (v *NullableActionTagElement) Set(val *ActionTagElement) {
	v.value = val
	v.isSet = true
}

func (v NullableActionTagElement) IsSet() bool {
	return v.isSet
}

func (v *NullableActionTagElement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionTagElement(val *ActionTagElement) *NullableActionTagElement {
	return &NullableActionTagElement{value: val, isSet: true}
}

func (v NullableActionTagElement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionTagElement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
