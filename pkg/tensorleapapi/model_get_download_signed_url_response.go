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

// checks if the GetDownloadSignedUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDownloadSignedUrlResponse{}

// GetDownloadSignedUrlResponse struct for GetDownloadSignedUrlResponse
type GetDownloadSignedUrlResponse struct {
	Url string `json:"url"`
}

// NewGetDownloadSignedUrlResponse instantiates a new GetDownloadSignedUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDownloadSignedUrlResponse(url string) *GetDownloadSignedUrlResponse {
	this := GetDownloadSignedUrlResponse{}
	this.Url = url
	return &this
}

// NewGetDownloadSignedUrlResponseWithDefaults instantiates a new GetDownloadSignedUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDownloadSignedUrlResponseWithDefaults() *GetDownloadSignedUrlResponse {
	this := GetDownloadSignedUrlResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *GetDownloadSignedUrlResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *GetDownloadSignedUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *GetDownloadSignedUrlResponse) SetUrl(v string) {
	o.Url = v
}

func (o GetDownloadSignedUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDownloadSignedUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableGetDownloadSignedUrlResponse struct {
	value *GetDownloadSignedUrlResponse
	isSet bool
}

func (v NullableGetDownloadSignedUrlResponse) Get() *GetDownloadSignedUrlResponse {
	return v.value
}

func (v *NullableGetDownloadSignedUrlResponse) Set(val *GetDownloadSignedUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDownloadSignedUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDownloadSignedUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDownloadSignedUrlResponse(val *GetDownloadSignedUrlResponse) *NullableGetDownloadSignedUrlResponse {
	return &NullableGetDownloadSignedUrlResponse{value: val, isSet: true}
}

func (v NullableGetDownloadSignedUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDownloadSignedUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


