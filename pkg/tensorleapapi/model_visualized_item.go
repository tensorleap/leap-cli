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

// checks if the VisualizedItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizedItem{}

// VisualizedItem struct for VisualizedItem
type VisualizedItem struct {
	VisualizerName string   `json:"visualizer_name"`
	Data           VisData  `json:"data"`
	EncoderNames   []string `json:"encoder_names"`
	ConnectionName string   `json:"connection_name"`
	NodeId         *string  `json:"node_id,omitempty"`
	HasError       *bool    `json:"has_error,omitempty"`
}

// NewVisualizedItem instantiates a new VisualizedItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizedItem(visualizerName string, data VisData, encoderNames []string, connectionName string) *VisualizedItem {
	this := VisualizedItem{}
	this.VisualizerName = visualizerName
	this.Data = data
	this.EncoderNames = encoderNames
	this.ConnectionName = connectionName
	return &this
}

// NewVisualizedItemWithDefaults instantiates a new VisualizedItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizedItemWithDefaults() *VisualizedItem {
	this := VisualizedItem{}
	return &this
}

// GetVisualizerName returns the VisualizerName field value
func (o *VisualizedItem) GetVisualizerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizerName
}

// GetVisualizerNameOk returns a tuple with the VisualizerName field value
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetVisualizerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizerName, true
}

// SetVisualizerName sets field value
func (o *VisualizedItem) SetVisualizerName(v string) {
	o.VisualizerName = v
}

// GetData returns the Data field value
func (o *VisualizedItem) GetData() VisData {
	if o == nil {
		var ret VisData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetDataOk() (*VisData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *VisualizedItem) SetData(v VisData) {
	o.Data = v
}

// GetEncoderNames returns the EncoderNames field value
func (o *VisualizedItem) GetEncoderNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EncoderNames
}

// GetEncoderNamesOk returns a tuple with the EncoderNames field value
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetEncoderNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncoderNames, true
}

// SetEncoderNames sets field value
func (o *VisualizedItem) SetEncoderNames(v []string) {
	o.EncoderNames = v
}

// GetConnectionName returns the ConnectionName field value
func (o *VisualizedItem) GetConnectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetConnectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *VisualizedItem) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *VisualizedItem) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *VisualizedItem) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *VisualizedItem) SetNodeId(v string) {
	o.NodeId = &v
}

// GetHasError returns the HasError field value if set, zero value otherwise.
func (o *VisualizedItem) GetHasError() bool {
	if o == nil || IsNil(o.HasError) {
		var ret bool
		return ret
	}
	return *o.HasError
}

// GetHasErrorOk returns a tuple with the HasError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizedItem) GetHasErrorOk() (*bool, bool) {
	if o == nil || IsNil(o.HasError) {
		return nil, false
	}
	return o.HasError, true
}

// HasHasError returns a boolean if a field has been set.
func (o *VisualizedItem) HasHasError() bool {
	if o != nil && !IsNil(o.HasError) {
		return true
	}

	return false
}

// SetHasError gets a reference to the given bool and assigns it to the HasError field.
func (o *VisualizedItem) SetHasError(v bool) {
	o.HasError = &v
}

func (o VisualizedItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizedItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visualizer_name"] = o.VisualizerName
	toSerialize["data"] = o.Data
	toSerialize["encoder_names"] = o.EncoderNames
	toSerialize["connection_name"] = o.ConnectionName
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.HasError) {
		toSerialize["has_error"] = o.HasError
	}
	return toSerialize, nil
}

type NullableVisualizedItem struct {
	value *VisualizedItem
	isSet bool
}

func (v NullableVisualizedItem) Get() *VisualizedItem {
	return v.value
}

func (v *NullableVisualizedItem) Set(val *VisualizedItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizedItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizedItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizedItem(val *VisualizedItem) *NullableVisualizedItem {
	return &NullableVisualizedItem{value: val, isSet: true}
}

func (v NullableVisualizedItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizedItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
