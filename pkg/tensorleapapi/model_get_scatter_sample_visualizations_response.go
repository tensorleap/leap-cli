/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetScatterSampleVisualizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetScatterSampleVisualizationsResponse{}

// GetScatterSampleVisualizationsResponse struct for GetScatterSampleVisualizationsResponse
type GetScatterSampleVisualizationsResponse struct {
	SamplesIds                        []string `json:"samplesIds"`
	ScatterSampleVisualizationsPrefix string   `json:"scatterSampleVisualizationsPrefix"`
}

// NewGetScatterSampleVisualizationsResponse instantiates a new GetScatterSampleVisualizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetScatterSampleVisualizationsResponse(samplesIds []string, scatterSampleVisualizationsPrefix string) *GetScatterSampleVisualizationsResponse {
	this := GetScatterSampleVisualizationsResponse{}
	this.SamplesIds = samplesIds
	this.ScatterSampleVisualizationsPrefix = scatterSampleVisualizationsPrefix
	return &this
}

// NewGetScatterSampleVisualizationsResponseWithDefaults instantiates a new GetScatterSampleVisualizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetScatterSampleVisualizationsResponseWithDefaults() *GetScatterSampleVisualizationsResponse {
	this := GetScatterSampleVisualizationsResponse{}
	return &this
}

// GetSamplesIds returns the SamplesIds field value
func (o *GetScatterSampleVisualizationsResponse) GetSamplesIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SamplesIds
}

// GetSamplesIdsOk returns a tuple with the SamplesIds field value
// and a boolean to check if the value has been set.
func (o *GetScatterSampleVisualizationsResponse) GetSamplesIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SamplesIds, true
}

// SetSamplesIds sets field value
func (o *GetScatterSampleVisualizationsResponse) SetSamplesIds(v []string) {
	o.SamplesIds = v
}

// GetScatterSampleVisualizationsPrefix returns the ScatterSampleVisualizationsPrefix field value
func (o *GetScatterSampleVisualizationsResponse) GetScatterSampleVisualizationsPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScatterSampleVisualizationsPrefix
}

// GetScatterSampleVisualizationsPrefixOk returns a tuple with the ScatterSampleVisualizationsPrefix field value
// and a boolean to check if the value has been set.
func (o *GetScatterSampleVisualizationsResponse) GetScatterSampleVisualizationsPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScatterSampleVisualizationsPrefix, true
}

// SetScatterSampleVisualizationsPrefix sets field value
func (o *GetScatterSampleVisualizationsResponse) SetScatterSampleVisualizationsPrefix(v string) {
	o.ScatterSampleVisualizationsPrefix = v
}

func (o GetScatterSampleVisualizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetScatterSampleVisualizationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["samplesIds"] = o.SamplesIds
	toSerialize["scatterSampleVisualizationsPrefix"] = o.ScatterSampleVisualizationsPrefix
	return toSerialize, nil
}

type NullableGetScatterSampleVisualizationsResponse struct {
	value *GetScatterSampleVisualizationsResponse
	isSet bool
}

func (v NullableGetScatterSampleVisualizationsResponse) Get() *GetScatterSampleVisualizationsResponse {
	return v.value
}

func (v *NullableGetScatterSampleVisualizationsResponse) Set(val *GetScatterSampleVisualizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetScatterSampleVisualizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetScatterSampleVisualizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetScatterSampleVisualizationsResponse(val *GetScatterSampleVisualizationsResponse) *NullableGetScatterSampleVisualizationsResponse {
	return &NullableGetScatterSampleVisualizationsResponse{value: val, isSet: true}
}

func (v NullableGetScatterSampleVisualizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetScatterSampleVisualizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
