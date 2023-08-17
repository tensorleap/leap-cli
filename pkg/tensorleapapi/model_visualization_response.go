/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.365
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the VisualizationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizationResponse{}

// VisualizationResponse struct for VisualizationResponse
type VisualizationResponse struct {
	Data   VisualizationData `json:"data"`
	Info   VizInfoType       `json:"info"`
	Source string            `json:"source"`
}

// NewVisualizationResponse instantiates a new VisualizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizationResponse(data VisualizationData, info VizInfoType, source string) *VisualizationResponse {
	this := VisualizationResponse{}
	this.Data = data
	this.Info = info
	this.Source = source
	return &this
}

// NewVisualizationResponseWithDefaults instantiates a new VisualizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationResponseWithDefaults() *VisualizationResponse {
	this := VisualizationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *VisualizationResponse) GetData() VisualizationData {
	if o == nil {
		var ret VisualizationData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VisualizationResponse) GetDataOk() (*VisualizationData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VisualizationResponse) SetData(v VisualizationData) {
	o.Data = v
}

// GetInfo returns the Info field value
func (o *VisualizationResponse) GetInfo() VizInfoType {
	if o == nil {
		var ret VizInfoType
		return ret
	}

	return o.Info
}

// GetInfoOk returns a tuple with the Info field value
// and a boolean to check if the value has been set.
func (o *VisualizationResponse) GetInfoOk() (*VizInfoType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Info, true
}

// SetInfo sets field value
func (o *VisualizationResponse) SetInfo(v VizInfoType) {
	o.Info = v
}

// GetSource returns the Source field value
func (o *VisualizationResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *VisualizationResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *VisualizationResponse) SetSource(v string) {
	o.Source = v
}

func (o VisualizationResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["info"] = o.Info
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableVisualizationResponse struct {
	value *VisualizationResponse
	isSet bool
}

func (v NullableVisualizationResponse) Get() *VisualizationResponse {
	return v.value
}

func (v *NullableVisualizationResponse) Set(val *VisualizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizationResponse(val *VisualizationResponse) *NullableVisualizationResponse {
	return &NullableVisualizationResponse{value: val, isSet: true}
}

func (v NullableVisualizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
