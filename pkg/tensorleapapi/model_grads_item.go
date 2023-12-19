/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.422
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GradsItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GradsItem{}

// GradsItem struct for GradsItem
type GradsItem struct {
	NodeName          string             `json:"node_name"`
	Label             string             `json:"label"`
	Depth             float64            `json:"depth"`
	Grads             VisData            `json:"grads"`
	VisualizerName    string             `json:"visualizer_name"`
	EncoderNames      []string           `json:"encoder_names"`
	ConnectionName    string             `json:"connection_name"`
	ChannelImportance *HorizontalBarData `json:"channel_importance,omitempty"`
}

// NewGradsItem instantiates a new GradsItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGradsItem(nodeName string, label string, depth float64, grads VisData, visualizerName string, encoderNames []string, connectionName string) *GradsItem {
	this := GradsItem{}
	this.NodeName = nodeName
	this.Label = label
	this.Depth = depth
	this.Grads = grads
	this.VisualizerName = visualizerName
	this.EncoderNames = encoderNames
	this.ConnectionName = connectionName
	return &this
}

// NewGradsItemWithDefaults instantiates a new GradsItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGradsItemWithDefaults() *GradsItem {
	this := GradsItem{}
	return &this
}

// GetNodeName returns the NodeName field value
func (o *GradsItem) GetNodeName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetNodeNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeName, true
}

// SetNodeName sets field value
func (o *GradsItem) SetNodeName(v string) {
	o.NodeName = v
}

// GetLabel returns the Label field value
func (o *GradsItem) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *GradsItem) SetLabel(v string) {
	o.Label = v
}

// GetDepth returns the Depth field value
func (o *GradsItem) GetDepth() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Depth
}

// GetDepthOk returns a tuple with the Depth field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetDepthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Depth, true
}

// SetDepth sets field value
func (o *GradsItem) SetDepth(v float64) {
	o.Depth = v
}

// GetGrads returns the Grads field value
func (o *GradsItem) GetGrads() VisData {
	if o == nil {
		var ret VisData
		return ret
	}

	return o.Grads
}

// GetGradsOk returns a tuple with the Grads field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetGradsOk() (*VisData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Grads, true
}

// SetGrads sets field value
func (o *GradsItem) SetGrads(v VisData) {
	o.Grads = v
}

// GetVisualizerName returns the VisualizerName field value
func (o *GradsItem) GetVisualizerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizerName
}

// GetVisualizerNameOk returns a tuple with the VisualizerName field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetVisualizerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizerName, true
}

// SetVisualizerName sets field value
func (o *GradsItem) SetVisualizerName(v string) {
	o.VisualizerName = v
}

// GetEncoderNames returns the EncoderNames field value
func (o *GradsItem) GetEncoderNames() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EncoderNames
}

// GetEncoderNamesOk returns a tuple with the EncoderNames field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetEncoderNamesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncoderNames, true
}

// SetEncoderNames sets field value
func (o *GradsItem) SetEncoderNames(v []string) {
	o.EncoderNames = v
}

// GetConnectionName returns the ConnectionName field value
func (o *GradsItem) GetConnectionName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value
// and a boolean to check if the value has been set.
func (o *GradsItem) GetConnectionNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConnectionName, true
}

// SetConnectionName sets field value
func (o *GradsItem) SetConnectionName(v string) {
	o.ConnectionName = v
}

// GetChannelImportance returns the ChannelImportance field value if set, zero value otherwise.
func (o *GradsItem) GetChannelImportance() HorizontalBarData {
	if o == nil || IsNil(o.ChannelImportance) {
		var ret HorizontalBarData
		return ret
	}
	return *o.ChannelImportance
}

// GetChannelImportanceOk returns a tuple with the ChannelImportance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GradsItem) GetChannelImportanceOk() (*HorizontalBarData, bool) {
	if o == nil || IsNil(o.ChannelImportance) {
		return nil, false
	}
	return o.ChannelImportance, true
}

// HasChannelImportance returns a boolean if a field has been set.
func (o *GradsItem) HasChannelImportance() bool {
	if o != nil && !IsNil(o.ChannelImportance) {
		return true
	}

	return false
}

// SetChannelImportance gets a reference to the given HorizontalBarData and assigns it to the ChannelImportance field.
func (o *GradsItem) SetChannelImportance(v HorizontalBarData) {
	o.ChannelImportance = &v
}

func (o GradsItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GradsItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["node_name"] = o.NodeName
	toSerialize["label"] = o.Label
	toSerialize["depth"] = o.Depth
	toSerialize["grads"] = o.Grads
	toSerialize["visualizer_name"] = o.VisualizerName
	toSerialize["encoder_names"] = o.EncoderNames
	toSerialize["connection_name"] = o.ConnectionName
	if !IsNil(o.ChannelImportance) {
		toSerialize["channel_importance"] = o.ChannelImportance
	}
	return toSerialize, nil
}

type NullableGradsItem struct {
	value *GradsItem
	isSet bool
}

func (v NullableGradsItem) Get() *GradsItem {
	return v.value
}

func (v *NullableGradsItem) Set(val *GradsItem) {
	v.value = val
	v.isSet = true
}

func (v NullableGradsItem) IsSet() bool {
	return v.isSet
}

func (v *NullableGradsItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGradsItem(val *GradsItem) *NullableGradsItem {
	return &NullableGradsItem{value: val, isSet: true}
}

func (v NullableGradsItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGradsItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
