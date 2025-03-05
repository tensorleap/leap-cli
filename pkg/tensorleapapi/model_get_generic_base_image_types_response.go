/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetGenericBaseImageTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGenericBaseImageTypesResponse{}

// GetGenericBaseImageTypesResponse struct for GetGenericBaseImageTypesResponse
type GetGenericBaseImageTypesResponse struct {
	BaseImageTypes       []GenericBaseImage `json:"baseImageTypes"`
	DefaultBaseImageType string             `json:"defaultBaseImageType"`
}

// NewGetGenericBaseImageTypesResponse instantiates a new GetGenericBaseImageTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGenericBaseImageTypesResponse(baseImageTypes []GenericBaseImage, defaultBaseImageType string) *GetGenericBaseImageTypesResponse {
	this := GetGenericBaseImageTypesResponse{}
	this.BaseImageTypes = baseImageTypes
	this.DefaultBaseImageType = defaultBaseImageType
	return &this
}

// NewGetGenericBaseImageTypesResponseWithDefaults instantiates a new GetGenericBaseImageTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGenericBaseImageTypesResponseWithDefaults() *GetGenericBaseImageTypesResponse {
	this := GetGenericBaseImageTypesResponse{}
	return &this
}

// GetBaseImageTypes returns the BaseImageTypes field value
func (o *GetGenericBaseImageTypesResponse) GetBaseImageTypes() []GenericBaseImage {
	if o == nil {
		var ret []GenericBaseImage
		return ret
	}

	return o.BaseImageTypes
}

// GetBaseImageTypesOk returns a tuple with the BaseImageTypes field value
// and a boolean to check if the value has been set.
func (o *GetGenericBaseImageTypesResponse) GetBaseImageTypesOk() ([]GenericBaseImage, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseImageTypes, true
}

// SetBaseImageTypes sets field value
func (o *GetGenericBaseImageTypesResponse) SetBaseImageTypes(v []GenericBaseImage) {
	o.BaseImageTypes = v
}

// GetDefaultBaseImageType returns the DefaultBaseImageType field value
func (o *GetGenericBaseImageTypesResponse) GetDefaultBaseImageType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBaseImageType
}

// GetDefaultBaseImageTypeOk returns a tuple with the DefaultBaseImageType field value
// and a boolean to check if the value has been set.
func (o *GetGenericBaseImageTypesResponse) GetDefaultBaseImageTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBaseImageType, true
}

// SetDefaultBaseImageType sets field value
func (o *GetGenericBaseImageTypesResponse) SetDefaultBaseImageType(v string) {
	o.DefaultBaseImageType = v
}

func (o GetGenericBaseImageTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGenericBaseImageTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["baseImageTypes"] = o.BaseImageTypes
	toSerialize["defaultBaseImageType"] = o.DefaultBaseImageType
	return toSerialize, nil
}

type NullableGetGenericBaseImageTypesResponse struct {
	value *GetGenericBaseImageTypesResponse
	isSet bool
}

func (v NullableGetGenericBaseImageTypesResponse) Get() *GetGenericBaseImageTypesResponse {
	return v.value
}

func (v *NullableGetGenericBaseImageTypesResponse) Set(val *GetGenericBaseImageTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGenericBaseImageTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGenericBaseImageTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGenericBaseImageTypesResponse(val *GetGenericBaseImageTypesResponse) *NullableGetGenericBaseImageTypesResponse {
	return &NullableGetGenericBaseImageTypesResponse{value: val, isSet: true}
}

func (v NullableGetGenericBaseImageTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGenericBaseImageTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
