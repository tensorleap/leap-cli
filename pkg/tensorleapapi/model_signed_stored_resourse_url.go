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

// checks if the SignedStoredResourseUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignedStoredResourseUrl{}

// SignedStoredResourseUrl struct for SignedStoredResourseUrl
type SignedStoredResourseUrl struct {
	FileUrl string `json:"fileUrl"`
}

// NewSignedStoredResourseUrl instantiates a new SignedStoredResourseUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignedStoredResourseUrl(fileUrl string) *SignedStoredResourseUrl {
	this := SignedStoredResourseUrl{}
	this.FileUrl = fileUrl
	return &this
}

// NewSignedStoredResourseUrlWithDefaults instantiates a new SignedStoredResourseUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignedStoredResourseUrlWithDefaults() *SignedStoredResourseUrl {
	this := SignedStoredResourseUrl{}
	return &this
}

// GetFileUrl returns the FileUrl field value
func (o *SignedStoredResourseUrl) GetFileUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value
// and a boolean to check if the value has been set.
func (o *SignedStoredResourseUrl) GetFileUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileUrl, true
}

// SetFileUrl sets field value
func (o *SignedStoredResourseUrl) SetFileUrl(v string) {
	o.FileUrl = v
}

func (o SignedStoredResourseUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignedStoredResourseUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileUrl"] = o.FileUrl
	return toSerialize, nil
}

type NullableSignedStoredResourseUrl struct {
	value *SignedStoredResourseUrl
	isSet bool
}

func (v NullableSignedStoredResourseUrl) Get() *SignedStoredResourseUrl {
	return v.value
}

func (v *NullableSignedStoredResourseUrl) Set(val *SignedStoredResourseUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableSignedStoredResourseUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableSignedStoredResourseUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignedStoredResourseUrl(val *SignedStoredResourseUrl) *NullableSignedStoredResourseUrl {
	return &NullableSignedStoredResourseUrl{value: val, isSet: true}
}

func (v NullableSignedStoredResourseUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignedStoredResourseUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
