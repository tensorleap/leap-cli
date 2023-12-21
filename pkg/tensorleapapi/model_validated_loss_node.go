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

// checks if the ValidatedLossNode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidatedLossNode{}

// ValidatedLossNode struct for ValidatedLossNode
type ValidatedLossNode struct {
	Name string `json:"name"`
	NodeId *string `json:"node_id,omitempty"`
	ConnectionName *string `json:"connection_name,omitempty"`
	Error *string `json:"error,omitempty"`
	IsCustomLoss bool `json:"is_custom_loss"`
}

// NewValidatedLossNode instantiates a new ValidatedLossNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidatedLossNode(name string, isCustomLoss bool) *ValidatedLossNode {
	this := ValidatedLossNode{}
	this.Name = name
	this.IsCustomLoss = isCustomLoss
	return &this
}

// NewValidatedLossNodeWithDefaults instantiates a new ValidatedLossNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidatedLossNodeWithDefaults() *ValidatedLossNode {
	this := ValidatedLossNode{}
	return &this
}

// GetName returns the Name field value
func (o *ValidatedLossNode) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ValidatedLossNode) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ValidatedLossNode) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *ValidatedLossNode) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedLossNode) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *ValidatedLossNode) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *ValidatedLossNode) SetNodeId(v string) {
	o.NodeId = &v
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *ValidatedLossNode) GetConnectionName() string {
	if o == nil || IsNil(o.ConnectionName) {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedLossNode) GetConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionName) {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *ValidatedLossNode) HasConnectionName() bool {
	if o != nil && !IsNil(o.ConnectionName) {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *ValidatedLossNode) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ValidatedLossNode) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidatedLossNode) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ValidatedLossNode) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ValidatedLossNode) SetError(v string) {
	o.Error = &v
}

// GetIsCustomLoss returns the IsCustomLoss field value
func (o *ValidatedLossNode) GetIsCustomLoss() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCustomLoss
}

// GetIsCustomLossOk returns a tuple with the IsCustomLoss field value
// and a boolean to check if the value has been set.
func (o *ValidatedLossNode) GetIsCustomLossOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCustomLoss, true
}

// SetIsCustomLoss sets field value
func (o *ValidatedLossNode) SetIsCustomLoss(v bool) {
	o.IsCustomLoss = v
}

func (o ValidatedLossNode) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidatedLossNode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.ConnectionName) {
		toSerialize["connection_name"] = o.ConnectionName
	}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	toSerialize["is_custom_loss"] = o.IsCustomLoss
	return toSerialize, nil
}

type NullableValidatedLossNode struct {
	value *ValidatedLossNode
	isSet bool
}

func (v NullableValidatedLossNode) Get() *ValidatedLossNode {
	return v.value
}

func (v *NullableValidatedLossNode) Set(val *ValidatedLossNode) {
	v.value = val
	v.isSet = true
}

func (v NullableValidatedLossNode) IsSet() bool {
	return v.isSet
}

func (v *NullableValidatedLossNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidatedLossNode(val *ValidatedLossNode) *NullableValidatedLossNode {
	return &NullableValidatedLossNode{value: val, isSet: true}
}

func (v NullableValidatedLossNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidatedLossNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


