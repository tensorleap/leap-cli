/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the MetricMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricMessageParams{}

// MetricMessageParams struct for MetricMessageParams
type MetricMessageParams struct {
	Message string `json:"message"`
	MessageCode *string `json:"message_code,omitempty"`
	MetricName string `json:"metric_name"`
	LossNodeId *string `json:"loss_node_id,omitempty"`
	LossName *string `json:"loss_name,omitempty"`
	Module Module `json:"module"`
	NodeId *string `json:"node_id,omitempty"`
	LineNumber *float64 `json:"line_number,omitempty"`
}

// NewMetricMessageParams instantiates a new MetricMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricMessageParams(message string, metricName string, module Module) *MetricMessageParams {
	this := MetricMessageParams{}
	this.Message = message
	this.MetricName = metricName
	this.Module = module
	return &this
}

// NewMetricMessageParamsWithDefaults instantiates a new MetricMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricMessageParamsWithDefaults() *MetricMessageParams {
	this := MetricMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *MetricMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MetricMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *MetricMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *MetricMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *MetricMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetMetricName returns the MetricName field value
func (o *MetricMessageParams) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *MetricMessageParams) SetMetricName(v string) {
	o.MetricName = v
}

// GetLossNodeId returns the LossNodeId field value if set, zero value otherwise.
func (o *MetricMessageParams) GetLossNodeId() string {
	if o == nil || IsNil(o.LossNodeId) {
		var ret string
		return ret
	}
	return *o.LossNodeId
}

// GetLossNodeIdOk returns a tuple with the LossNodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetLossNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.LossNodeId) {
		return nil, false
	}
	return o.LossNodeId, true
}

// HasLossNodeId returns a boolean if a field has been set.
func (o *MetricMessageParams) HasLossNodeId() bool {
	if o != nil && !IsNil(o.LossNodeId) {
		return true
	}

	return false
}

// SetLossNodeId gets a reference to the given string and assigns it to the LossNodeId field.
func (o *MetricMessageParams) SetLossNodeId(v string) {
	o.LossNodeId = &v
}

// GetLossName returns the LossName field value if set, zero value otherwise.
func (o *MetricMessageParams) GetLossName() string {
	if o == nil || IsNil(o.LossName) {
		var ret string
		return ret
	}
	return *o.LossName
}

// GetLossNameOk returns a tuple with the LossName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetLossNameOk() (*string, bool) {
	if o == nil || IsNil(o.LossName) {
		return nil, false
	}
	return o.LossName, true
}

// HasLossName returns a boolean if a field has been set.
func (o *MetricMessageParams) HasLossName() bool {
	if o != nil && !IsNil(o.LossName) {
		return true
	}

	return false
}

// SetLossName gets a reference to the given string and assigns it to the LossName field.
func (o *MetricMessageParams) SetLossName(v string) {
	o.LossName = &v
}

// GetModule returns the Module field value
func (o *MetricMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *MetricMessageParams) SetModule(v Module) {
	o.Module = v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *MetricMessageParams) GetNodeId() string {
	if o == nil || IsNil(o.NodeId) {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetNodeIdOk() (*string, bool) {
	if o == nil || IsNil(o.NodeId) {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *MetricMessageParams) HasNodeId() bool {
	if o != nil && !IsNil(o.NodeId) {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *MetricMessageParams) SetNodeId(v string) {
	o.NodeId = &v
}

// GetLineNumber returns the LineNumber field value if set, zero value otherwise.
func (o *MetricMessageParams) GetLineNumber() float64 {
	if o == nil || IsNil(o.LineNumber) {
		var ret float64
		return ret
	}
	return *o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricMessageParams) GetLineNumberOk() (*float64, bool) {
	if o == nil || IsNil(o.LineNumber) {
		return nil, false
	}
	return o.LineNumber, true
}

// HasLineNumber returns a boolean if a field has been set.
func (o *MetricMessageParams) HasLineNumber() bool {
	if o != nil && !IsNil(o.LineNumber) {
		return true
	}

	return false
}

// SetLineNumber gets a reference to the given float64 and assigns it to the LineNumber field.
func (o *MetricMessageParams) SetLineNumber(v float64) {
	o.LineNumber = &v
}

func (o MetricMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["metric_name"] = o.MetricName
	if !IsNil(o.LossNodeId) {
		toSerialize["loss_node_id"] = o.LossNodeId
	}
	if !IsNil(o.LossName) {
		toSerialize["loss_name"] = o.LossName
	}
	toSerialize["module"] = o.Module
	if !IsNil(o.NodeId) {
		toSerialize["node_id"] = o.NodeId
	}
	if !IsNil(o.LineNumber) {
		toSerialize["line_number"] = o.LineNumber
	}
	return toSerialize, nil
}

type NullableMetricMessageParams struct {
	value *MetricMessageParams
	isSet bool
}

func (v NullableMetricMessageParams) Get() *MetricMessageParams {
	return v.value
}

func (v *NullableMetricMessageParams) Set(val *MetricMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricMessageParams(val *MetricMessageParams) *NullableMetricMessageParams {
	return &NullableMetricMessageParams{value: val, isSet: true}
}

func (v NullableMetricMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


