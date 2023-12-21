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

// checks if the GetMachineTypesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMachineTypesResponse{}

// GetMachineTypesResponse struct for GetMachineTypesResponse
type GetMachineTypesResponse struct {
	Options []MachineTypeOption `json:"options"`
	DefaultCpuType string `json:"defaultCpuType"`
	DefaultGpuType string `json:"defaultGpuType"`
}

// NewGetMachineTypesResponse instantiates a new GetMachineTypesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMachineTypesResponse(options []MachineTypeOption, defaultCpuType string, defaultGpuType string) *GetMachineTypesResponse {
	this := GetMachineTypesResponse{}
	this.Options = options
	this.DefaultCpuType = defaultCpuType
	this.DefaultGpuType = defaultGpuType
	return &this
}

// NewGetMachineTypesResponseWithDefaults instantiates a new GetMachineTypesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMachineTypesResponseWithDefaults() *GetMachineTypesResponse {
	this := GetMachineTypesResponse{}
	return &this
}

// GetOptions returns the Options field value
func (o *GetMachineTypesResponse) GetOptions() []MachineTypeOption {
	if o == nil {
		var ret []MachineTypeOption
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *GetMachineTypesResponse) GetOptionsOk() ([]MachineTypeOption, bool) {
	if o == nil {
		return nil, false
	}
	return o.Options, true
}

// SetOptions sets field value
func (o *GetMachineTypesResponse) SetOptions(v []MachineTypeOption) {
	o.Options = v
}

// GetDefaultCpuType returns the DefaultCpuType field value
func (o *GetMachineTypesResponse) GetDefaultCpuType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultCpuType
}

// GetDefaultCpuTypeOk returns a tuple with the DefaultCpuType field value
// and a boolean to check if the value has been set.
func (o *GetMachineTypesResponse) GetDefaultCpuTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultCpuType, true
}

// SetDefaultCpuType sets field value
func (o *GetMachineTypesResponse) SetDefaultCpuType(v string) {
	o.DefaultCpuType = v
}

// GetDefaultGpuType returns the DefaultGpuType field value
func (o *GetMachineTypesResponse) GetDefaultGpuType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultGpuType
}

// GetDefaultGpuTypeOk returns a tuple with the DefaultGpuType field value
// and a boolean to check if the value has been set.
func (o *GetMachineTypesResponse) GetDefaultGpuTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultGpuType, true
}

// SetDefaultGpuType sets field value
func (o *GetMachineTypesResponse) SetDefaultGpuType(v string) {
	o.DefaultGpuType = v
}

func (o GetMachineTypesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMachineTypesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["options"] = o.Options
	toSerialize["defaultCpuType"] = o.DefaultCpuType
	toSerialize["defaultGpuType"] = o.DefaultGpuType
	return toSerialize, nil
}

type NullableGetMachineTypesResponse struct {
	value *GetMachineTypesResponse
	isSet bool
}

func (v NullableGetMachineTypesResponse) Get() *GetMachineTypesResponse {
	return v.value
}

func (v *NullableGetMachineTypesResponse) Set(val *GetMachineTypesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMachineTypesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMachineTypesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMachineTypesResponse(val *GetMachineTypesResponse) *NullableGetMachineTypesResponse {
	return &NullableGetMachineTypesResponse{value: val, isSet: true}
}

func (v NullableGetMachineTypesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMachineTypesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


