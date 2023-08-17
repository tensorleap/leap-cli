/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ExternalImportModelStorage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalImportModelStorage{}

// ExternalImportModelStorage struct for ExternalImportModelStorage
type ExternalImportModelStorage struct {
	Url      string `json:"url"`
	FileName string `json:"fileName"`
}

// NewExternalImportModelStorage instantiates a new ExternalImportModelStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalImportModelStorage(url string, fileName string) *ExternalImportModelStorage {
	this := ExternalImportModelStorage{}
	this.Url = url
	this.FileName = fileName
	return &this
}

// NewExternalImportModelStorageWithDefaults instantiates a new ExternalImportModelStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalImportModelStorageWithDefaults() *ExternalImportModelStorage {
	this := ExternalImportModelStorage{}
	return &this
}

// GetUrl returns the Url field value
func (o *ExternalImportModelStorage) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ExternalImportModelStorage) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ExternalImportModelStorage) SetUrl(v string) {
	o.Url = v
}

// GetFileName returns the FileName field value
func (o *ExternalImportModelStorage) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *ExternalImportModelStorage) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *ExternalImportModelStorage) SetFileName(v string) {
	o.FileName = v
}

func (o ExternalImportModelStorage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalImportModelStorage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableExternalImportModelStorage struct {
	value *ExternalImportModelStorage
	isSet bool
}

func (v NullableExternalImportModelStorage) Get() *ExternalImportModelStorage {
	return v.value
}

func (v *NullableExternalImportModelStorage) Set(val *ExternalImportModelStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalImportModelStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalImportModelStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalImportModelStorage(val *ExternalImportModelStorage) *NullableExternalImportModelStorage {
	return &NullableExternalImportModelStorage{value: val, isSet: true}
}

func (v NullableExternalImportModelStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalImportModelStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
