/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the VisualizerMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizerMessageParams{}

// VisualizerMessageParams struct for VisualizerMessageParams
type VisualizerMessageParams struct {
	Message        string   `json:"message"`
	MessageCode    *string  `json:"message_code,omitempty"`
	VisualizerName string   `json:"visualizer_name"`
	NodeId         string   `json:"node_id"`
	LineNumber     *float64 `json:"line_number,omitempty"`
	Module         Module   `json:"module"`
}

// NewVisualizerMessageParams instantiates a new VisualizerMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizerMessageParams(message string, visualizerName string, nodeId string, module Module) *VisualizerMessageParams {
	this := VisualizerMessageParams{}
	this.Message = message
	this.VisualizerName = visualizerName
	this.NodeId = nodeId
	this.Module = module
	return &this
}

// NewVisualizerMessageParamsWithDefaults instantiates a new VisualizerMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizerMessageParamsWithDefaults() *VisualizerMessageParams {
	this := VisualizerMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *VisualizerMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *VisualizerMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *VisualizerMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *VisualizerMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *VisualizerMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetVisualizerName returns the VisualizerName field value
func (o *VisualizerMessageParams) GetVisualizerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizerName
}

// GetVisualizerNameOk returns a tuple with the VisualizerName field value
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetVisualizerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizerName, true
}

// SetVisualizerName sets field value
func (o *VisualizerMessageParams) SetVisualizerName(v string) {
	o.VisualizerName = v
}

// GetNodeId returns the NodeId field value
func (o *VisualizerMessageParams) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *VisualizerMessageParams) SetNodeId(v string) {
	o.NodeId = v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *VisualizerMessageParams) GetLineNumber() float64 {
	if o == nil || IsNil(o.LineNumber) {
		var ret float64
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetLineNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.LineNumber) {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *VisualizerMessageParams) HasLineNumber() bool {
	if o != nil && !IsNil(o.LineNumber) {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given float64 and assigns it to the LineNumber field.
func (o *VisualizerMessageParams) SetLineNumber(v float64) {
	o.LineNumber = &v
}

// GetModule returns the Module field value
func (o *VisualizerMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *VisualizerMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *VisualizerMessageParams) SetModule(v Module) {
	o.Module = v
}

func (o VisualizerMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizerMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["visualizer_name"] = o.VisualizerName
	toSerialize["node_id"] = o.NodeId
	if !IsNil(o.LineNumber) {
		toSerialize["line_number"] = o.LineNumber
	}
	toSerialize["module"] = o.Module
	return toSerialize, nil
}

type NullableVisualizerMessageParams struct {
	value *VisualizerMessageParams
	isSet bool
}

func (v NullableVisualizerMessageParams) Get() *VisualizerMessageParams {
	return v.value
}

func (v *NullableVisualizerMessageParams) Set(val *VisualizerMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizerMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizerMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizerMessageParams(val *VisualizerMessageParams) *NullableVisualizerMessageParams {
	return &NullableVisualizerMessageParams{value: val, isSet: true}
}

func (v NullableVisualizerMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizerMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
