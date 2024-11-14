/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetSignedUrlParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSignedUrlParams{}

// GetSignedUrlParams struct for GetSignedUrlParams
type GetSignedUrlParams struct {
	Origin   *string     `json:"origin,omitempty"`
	FileName string      `json:"fileName"`
	Expired  float64     `json:"expired"`
	Method   HttpMethods `json:"method"`
}

// NewGetSignedUrlParams instantiates a new GetSignedUrlParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSignedUrlParams(fileName string, expired float64, method HttpMethods) *GetSignedUrlParams {
	this := GetSignedUrlParams{}
	this.FileName = fileName
	this.Expired = expired
	this.Method = method
	return &this
}

// NewGetSignedUrlParamsWithDefaults instantiates a new GetSignedUrlParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSignedUrlParamsWithDefaults() *GetSignedUrlParams {
	this := GetSignedUrlParams{}
	return &this
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *GetSignedUrlParams) GetOrigin() string {
	if o == nil || IsNil(o.Origin) {
		var ret string
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSignedUrlParams) GetOriginOk() (*string, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *GetSignedUrlParams) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given string and assigns it to the Origin field.
func (o *GetSignedUrlParams) SetOrigin(v string) {
	o.Origin = &v
}

// GetFileName returns the FileName field value
func (o *GetSignedUrlParams) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *GetSignedUrlParams) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *GetSignedUrlParams) SetFileName(v string) {
	o.FileName = v
}

// GetExpired returns the Expired field value
func (o *GetSignedUrlParams) GetExpired() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Expired
}

// GetExpiredOk returns a tuple with the Expired field value
// and a boolean to check if the value has been set.
func (o *GetSignedUrlParams) GetExpiredOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expired, true
}

// SetExpired sets field value
func (o *GetSignedUrlParams) SetExpired(v float64) {
	o.Expired = v
}

// GetMethod returns the Method field value
func (o *GetSignedUrlParams) GetMethod() HttpMethods {
	if o == nil {
		var ret HttpMethods
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *GetSignedUrlParams) GetMethodOk() (*HttpMethods, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *GetSignedUrlParams) SetMethod(v HttpMethods) {
	o.Method = v
}

func (o GetSignedUrlParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSignedUrlParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	toSerialize["fileName"] = o.FileName
	toSerialize["expired"] = o.Expired
	toSerialize["method"] = o.Method
	return toSerialize, nil
}

type NullableGetSignedUrlParams struct {
	value *GetSignedUrlParams
	isSet bool
}

func (v NullableGetSignedUrlParams) Get() *GetSignedUrlParams {
	return v.value
}

func (v *NullableGetSignedUrlParams) Set(val *GetSignedUrlParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSignedUrlParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSignedUrlParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSignedUrlParams(val *GetSignedUrlParams) *NullableGetSignedUrlParams {
	return &NullableGetSignedUrlParams{value: val, isSet: true}
}

func (v NullableGetSignedUrlParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSignedUrlParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
