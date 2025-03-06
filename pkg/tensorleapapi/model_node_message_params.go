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

// checks if the NodeMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NodeMessageParams{}

// NodeMessageParams struct for NodeMessageParams
type NodeMessageParams struct {
	Message     string   `json:"message"`
	MessageCode *string  `json:"message_code,omitempty"`
	LayerName   string   `json:"layer_name"`
	LayerType   NodeType `json:"layer_type"`
	Module      Module   `json:"module"`
}

// NewNodeMessageParams instantiates a new NodeMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNodeMessageParams(message string, layerName string, layerType NodeType, module Module) *NodeMessageParams {
	this := NodeMessageParams{}
	this.Message = message
	this.LayerName = layerName
	this.LayerType = layerType
	this.Module = module
	return &this
}

// NewNodeMessageParamsWithDefaults instantiates a new NodeMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNodeMessageParamsWithDefaults() *NodeMessageParams {
	this := NodeMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *NodeMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *NodeMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *NodeMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *NodeMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NodeMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *NodeMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *NodeMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetLayerName returns the LayerName field value
func (o *NodeMessageParams) GetLayerName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LayerName
}

// GetLayerNameOk returns a tuple with the LayerName field value
// and a boolean to check if the value has been set.
func (o *NodeMessageParams) GetLayerNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayerName, true
}

// SetLayerName sets field value
func (o *NodeMessageParams) SetLayerName(v string) {
	o.LayerName = v
}

// GetLayerType returns the LayerType field value
func (o *NodeMessageParams) GetLayerType() NodeType {
	if o == nil {
		var ret NodeType
		return ret
	}

	return o.LayerType
}

// GetLayerTypeOk returns a tuple with the LayerType field value
// and a boolean to check if the value has been set.
func (o *NodeMessageParams) GetLayerTypeOk() (*NodeType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LayerType, true
}

// SetLayerType sets field value
func (o *NodeMessageParams) SetLayerType(v NodeType) {
	o.LayerType = v
}

// GetModule returns the Module field value
func (o *NodeMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *NodeMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *NodeMessageParams) SetModule(v Module) {
	o.Module = v
}

func (o NodeMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NodeMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["layer_name"] = o.LayerName
	toSerialize["layer_type"] = o.LayerType
	toSerialize["module"] = o.Module
	return toSerialize, nil
}

type NullableNodeMessageParams struct {
	value *NodeMessageParams
	isSet bool
}

func (v NullableNodeMessageParams) Get() *NodeMessageParams {
	return v.value
}

func (v *NullableNodeMessageParams) Set(val *NodeMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeMessageParams(val *NodeMessageParams) *NullableNodeMessageParams {
	return &NullableNodeMessageParams{value: val, isSet: true}
}

func (v NullableNodeMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
