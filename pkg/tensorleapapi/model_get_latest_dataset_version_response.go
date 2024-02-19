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

// checks if the GetLatestDatasetVersionResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetLatestDatasetVersionResponse{}

// GetLatestDatasetVersionResponse struct for GetLatestDatasetVersionResponse
type GetLatestDatasetVersionResponse struct {
	LatestVersion *DatasetVersion `json:"latestVersion,omitempty"`
}

// NewGetLatestDatasetVersionResponse instantiates a new GetLatestDatasetVersionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetLatestDatasetVersionResponse() *GetLatestDatasetVersionResponse {
	this := GetLatestDatasetVersionResponse{}
	return &this
}

// NewGetLatestDatasetVersionResponseWithDefaults instantiates a new GetLatestDatasetVersionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetLatestDatasetVersionResponseWithDefaults() *GetLatestDatasetVersionResponse {
	this := GetLatestDatasetVersionResponse{}
	return &this
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *GetLatestDatasetVersionResponse) GetLatestVersion() DatasetVersion {
	if o == nil || IsNil(o.LatestVersion) {
		var ret DatasetVersion
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetLatestDatasetVersionResponse) GetLatestVersionOk() (*DatasetVersion, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *GetLatestDatasetVersionResponse) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given DatasetVersion and assigns it to the LatestVersion field.
func (o *GetLatestDatasetVersionResponse) SetLatestVersion(v DatasetVersion) {
	o.LatestVersion = &v
}

func (o GetLatestDatasetVersionResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetLatestDatasetVersionResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	return toSerialize, nil
}

type NullableGetLatestDatasetVersionResponse struct {
	value *GetLatestDatasetVersionResponse
	isSet bool
}

func (v NullableGetLatestDatasetVersionResponse) Get() *GetLatestDatasetVersionResponse {
	return v.value
}

func (v *NullableGetLatestDatasetVersionResponse) Set(val *GetLatestDatasetVersionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetLatestDatasetVersionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetLatestDatasetVersionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetLatestDatasetVersionResponse(val *GetLatestDatasetVersionResponse) *NullableGetLatestDatasetVersionResponse {
	return &NullableGetLatestDatasetVersionResponse{value: val, isSet: true}
}

func (v NullableGetLatestDatasetVersionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetLatestDatasetVersionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
