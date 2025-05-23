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

// checks if the LossMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LossMessageParams{}

// LossMessageParams struct for LossMessageParams
type LossMessageParams struct {
	Message     string  `json:"message"`
	MessageCode *string `json:"message_code,omitempty"`
	LossName    string  `json:"loss_name"`
	NodeId      string  `json:"node_id"`
	Module      Module  `json:"module"`
}

// NewLossMessageParams instantiates a new LossMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLossMessageParams(message string, lossName string, nodeId string, module Module) *LossMessageParams {
	this := LossMessageParams{}
	this.Message = message
	this.LossName = lossName
	this.NodeId = nodeId
	this.Module = module
	return &this
}

// NewLossMessageParamsWithDefaults instantiates a new LossMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLossMessageParamsWithDefaults() *LossMessageParams {
	this := LossMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *LossMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *LossMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *LossMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *LossMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LossMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *LossMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *LossMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetLossName returns the LossName field value
func (o *LossMessageParams) GetLossName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LossName
}

// GetLossNameOk returns a tuple with the LossName field value
// and a boolean to check if the value has been set.
func (o *LossMessageParams) GetLossNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LossName, true
}

// SetLossName sets field value
func (o *LossMessageParams) SetLossName(v string) {
	o.LossName = v
}

// GetNodeId returns the NodeId field value
func (o *LossMessageParams) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *LossMessageParams) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *LossMessageParams) SetNodeId(v string) {
	o.NodeId = v
}

// GetModule returns the Module field value
func (o *LossMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *LossMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *LossMessageParams) SetModule(v Module) {
	o.Module = v
}

func (o LossMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LossMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["loss_name"] = o.LossName
	toSerialize["node_id"] = o.NodeId
	toSerialize["module"] = o.Module
	return toSerialize, nil
}

type NullableLossMessageParams struct {
	value *LossMessageParams
	isSet bool
}

func (v NullableLossMessageParams) Get() *LossMessageParams {
	return v.value
}

func (v *NullableLossMessageParams) Set(val *LossMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableLossMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableLossMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossMessageParams(val *LossMessageParams) *NullableLossMessageParams {
	return &NullableLossMessageParams{value: val, isSet: true}
}

func (v NullableLossMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
