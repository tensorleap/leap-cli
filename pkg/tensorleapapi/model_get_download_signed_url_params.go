/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetDownloadSignedUrlParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetDownloadSignedUrlParams{}

// GetDownloadSignedUrlParams struct for GetDownloadSignedUrlParams
type GetDownloadSignedUrlParams struct {
	Origin   *string `json:"origin,omitempty"`
	FileName string  `json:"fileName"`
}

// NewGetDownloadSignedUrlParams instantiates a new GetDownloadSignedUrlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetDownloadSignedUrlParams(fileName string) *GetDownloadSignedUrlParams {
	this := GetDownloadSignedUrlParams{}
	this.FileName = fileName
	return &this
}

// NewGetDownloadSignedUrlParamsWithDefaults instantiates a new GetDownloadSignedUrlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetDownloadSignedUrlParamsWithDefaults() *GetDownloadSignedUrlParams {
	this := GetDownloadSignedUrlParams{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetDownloadSignedUrlParams) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetDownloadSignedUrlParams) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetDownloadSignedUrlParams) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *GetDownloadSignedUrlParams) SetOrigin(v string) {
	o.Origin = &v
}

// GetFileName returns the FileName field value
func (o *GetDownloadSignedUrlParams) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *GetDownloadSignedUrlParams) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *GetDownloadSignedUrlParams) SetFileName(v string) {
	o.FileName = v
}

func (o GetDownloadSignedUrlParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetDownloadSignedUrlParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableGetDownloadSignedUrlParams struct {
	value *GetDownloadSignedUrlParams
	isSet bool
}

func (v NullableGetDownloadSignedUrlParams) Get() *GetDownloadSignedUrlParams {
	return v.value
}

func (v *NullableGetDownloadSignedUrlParams) Set(val *GetDownloadSignedUrlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetDownloadSignedUrlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetDownloadSignedUrlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetDownloadSignedUrlParams(val *GetDownloadSignedUrlParams) *NullableGetDownloadSignedUrlParams {
	return &NullableGetDownloadSignedUrlParams{value: val, isSet: true}
}

func (v NullableGetDownloadSignedUrlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetDownloadSignedUrlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
