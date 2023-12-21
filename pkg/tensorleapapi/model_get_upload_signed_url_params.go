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

// checks if the GetUploadSignedUrlParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUploadSignedUrlParams{}

// GetUploadSignedUrlParams struct for GetUploadSignedUrlParams
type GetUploadSignedUrlParams struct {
	Origin   *string `json:"origin,omitempty"`
	FileName string  `json:"fileName"`
}

// NewGetUploadSignedUrlParams instantiates a new GetUploadSignedUrlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUploadSignedUrlParams(fileName string) *GetUploadSignedUrlParams {
	this := GetUploadSignedUrlParams{}
	this.FileName = fileName
	return &this
}

// NewGetUploadSignedUrlParamsWithDefaults instantiates a new GetUploadSignedUrlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUploadSignedUrlParamsWithDefaults() *GetUploadSignedUrlParams {
	this := GetUploadSignedUrlParams{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetUploadSignedUrlParams) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUploadSignedUrlParams) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetUploadSignedUrlParams) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *GetUploadSignedUrlParams) SetOrigin(v string) {
	o.Origin = &v
}

// GetFileName returns the FileName field value
func (o *GetUploadSignedUrlParams) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *GetUploadSignedUrlParams) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *GetUploadSignedUrlParams) SetFileName(v string) {
	o.FileName = v
}

func (o GetUploadSignedUrlParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUploadSignedUrlParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	toSerialize["fileName"] = o.FileName
	return toSerialize, nil
}

type NullableGetUploadSignedUrlParams struct {
	value *GetUploadSignedUrlParams
	isSet bool
}

func (v NullableGetUploadSignedUrlParams) Get() *GetUploadSignedUrlParams {
	return v.value
}

func (v *NullableGetUploadSignedUrlParams) Set(val *GetUploadSignedUrlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUploadSignedUrlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUploadSignedUrlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUploadSignedUrlParams(val *GetUploadSignedUrlParams) *NullableGetUploadSignedUrlParams {
	return &NullableGetUploadSignedUrlParams{value: val, isSet: true}
}

func (v NullableGetUploadSignedUrlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUploadSignedUrlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
