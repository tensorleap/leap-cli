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

// checks if the GetSampleVisualizationsPathsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSampleVisualizationsPathsParams{}

// GetSampleVisualizationsPathsParams struct for GetSampleVisualizationsPathsParams
type GetSampleVisualizationsPathsParams struct {
	ScatterSampleVisualizationsPrefix string `json:"scatterSampleVisualizationsPrefix"`
	FileNameMatch                     string `json:"fileNameMatch"`
	SampleId                          string `json:"sampleId"`
}

// NewGetSampleVisualizationsPathsParams instantiates a new GetSampleVisualizationsPathsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSampleVisualizationsPathsParams(scatterSampleVisualizationsPrefix string, fileNameMatch string, sampleId string) *GetSampleVisualizationsPathsParams {
	this := GetSampleVisualizationsPathsParams{}
	this.ScatterSampleVisualizationsPrefix = scatterSampleVisualizationsPrefix
	this.FileNameMatch = fileNameMatch
	this.SampleId = sampleId
	return &this
}

// NewGetSampleVisualizationsPathsParamsWithDefaults instantiates a new GetSampleVisualizationsPathsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSampleVisualizationsPathsParamsWithDefaults() *GetSampleVisualizationsPathsParams {
	this := GetSampleVisualizationsPathsParams{}
	return &this
}

// GetScatterSampleVisualizationsPrefix returns the ScatterSampleVisualizationsPrefix field value
func (o *GetSampleVisualizationsPathsParams) GetScatterSampleVisualizationsPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScatterSampleVisualizationsPrefix
}

// GetScatterSampleVisualizationsPrefixOk returns a tuple with the ScatterSampleVisualizationsPrefix field value
// and a boolean to check if the value has been set.
func (o *GetSampleVisualizationsPathsParams) GetScatterSampleVisualizationsPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScatterSampleVisualizationsPrefix, true
}

// SetScatterSampleVisualizationsPrefix sets field value
func (o *GetSampleVisualizationsPathsParams) SetScatterSampleVisualizationsPrefix(v string) {
	o.ScatterSampleVisualizationsPrefix = v
}

// GetFileNameMatch returns the FileNameMatch field value
func (o *GetSampleVisualizationsPathsParams) GetFileNameMatch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileNameMatch
}

// GetFileNameMatchOk returns a tuple with the FileNameMatch field value
// and a boolean to check if the value has been set.
func (o *GetSampleVisualizationsPathsParams) GetFileNameMatchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileNameMatch, true
}

// SetFileNameMatch sets field value
func (o *GetSampleVisualizationsPathsParams) SetFileNameMatch(v string) {
	o.FileNameMatch = v
}

// GetSampleId returns the SampleId field value
func (o *GetSampleVisualizationsPathsParams) GetSampleId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SampleId
}

// GetSampleIdOk returns a tuple with the SampleId field value
// and a boolean to check if the value has been set.
func (o *GetSampleVisualizationsPathsParams) GetSampleIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SampleId, true
}

// SetSampleId sets field value
func (o *GetSampleVisualizationsPathsParams) SetSampleId(v string) {
	o.SampleId = v
}

func (o GetSampleVisualizationsPathsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSampleVisualizationsPathsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scatterSampleVisualizationsPrefix"] = o.ScatterSampleVisualizationsPrefix
	toSerialize["fileNameMatch"] = o.FileNameMatch
	toSerialize["sampleId"] = o.SampleId
	return toSerialize, nil
}

type NullableGetSampleVisualizationsPathsParams struct {
	value *GetSampleVisualizationsPathsParams
	isSet bool
}

func (v NullableGetSampleVisualizationsPathsParams) Get() *GetSampleVisualizationsPathsParams {
	return v.value
}

func (v *NullableGetSampleVisualizationsPathsParams) Set(val *GetSampleVisualizationsPathsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSampleVisualizationsPathsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSampleVisualizationsPathsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSampleVisualizationsPathsParams(val *GetSampleVisualizationsPathsParams) *NullableGetSampleVisualizationsPathsParams {
	return &NullableGetSampleVisualizationsPathsParams{value: val, isSet: true}
}

func (v NullableGetSampleVisualizationsPathsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSampleVisualizationsPathsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
