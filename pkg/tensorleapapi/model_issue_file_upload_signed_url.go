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

// checks if the IssueFileUploadSignedUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueFileUploadSignedUrl{}

// IssueFileUploadSignedUrl struct for IssueFileUploadSignedUrl
type IssueFileUploadSignedUrl struct {
	Url      string `json:"url"`
	FileName string `json:"fileName"`
}

// NewIssueFileUploadSignedUrl instantiates a new IssueFileUploadSignedUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueFileUploadSignedUrl(url string, fileName string) *IssueFileUploadSignedUrl {
	this := IssueFileUploadSignedUrl{}
	this.Url = url
	this.FileName = fileName
	return &this
}

// NewIssueFileUploadSignedUrlWithDefaults instantiates a new IssueFileUploadSignedUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueFileUploadSignedUrlWithDefaults() *IssueFileUploadSignedUrl {
	this := IssueFileUploadSignedUrl{}
	return &this
}

// GetUrl returns the Url field value
func (o *IssueFileUploadSignedUrl) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *IssueFileUploadSignedUrl) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *IssueFileUploadSignedUrl) SetUrl(v string) {
	o.Url = v
}

// GetFileName returns the FileName field value
func (o *IssueFileUploadSignedUrl) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *IssueFileUploadSignedUrl) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *IssueFileUploadSignedUrl) SetFileName(v string) {
	o.FileName = v
}

func (o IssueFileUploadSignedUrl) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueFileUploadSignedUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableIssueFileUploadSignedUrl struct {
	value *IssueFileUploadSignedUrl
	isSet bool
}

func (v NullableIssueFileUploadSignedUrl) Get() *IssueFileUploadSignedUrl {
	return v.value
}

func (v *NullableIssueFileUploadSignedUrl) Set(val *IssueFileUploadSignedUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueFileUploadSignedUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueFileUploadSignedUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueFileUploadSignedUrl(val *IssueFileUploadSignedUrl) *NullableIssueFileUploadSignedUrl {
	return &NullableIssueFileUploadSignedUrl{value: val, isSet: true}
}

func (v NullableIssueFileUploadSignedUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueFileUploadSignedUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
