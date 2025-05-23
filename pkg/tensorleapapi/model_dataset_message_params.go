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

// checks if the DatasetMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetMessageParams{}

// DatasetMessageParams struct for DatasetMessageParams
type DatasetMessageParams struct {
	Message     string  `json:"message"`
	MessageCode *string `json:"message_code,omitempty"`
	Line        float64 `json:"line"`
	FuncName    string  `json:"func_name"`
	Module      Module  `json:"module"`
}

// NewDatasetMessageParams instantiates a new DatasetMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetMessageParams(message string, line float64, funcName string, module Module) *DatasetMessageParams {
	this := DatasetMessageParams{}
	this.Message = message
	this.Line = line
	this.FuncName = funcName
	this.Module = module
	return &this
}

// NewDatasetMessageParamsWithDefaults instantiates a new DatasetMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetMessageParamsWithDefaults() *DatasetMessageParams {
	this := DatasetMessageParams{}
	return &this
}

// GetMessage returns the Message field value
func (o *DatasetMessageParams) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *DatasetMessageParams) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *DatasetMessageParams) SetMessage(v string) {
	o.Message = v
}

// GetMessageCode returns the MessageCode field value if set, zero value otherwise.
func (o *DatasetMessageParams) GetMessageCode() string {
	if o == nil || IsNil(o.MessageCode) {
		var ret string
		return ret
	}
	return *o.MessageCode
}

// GetMessageCodeOk returns a tuple with the MessageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetMessageParams) GetMessageCodeOk() (*string, bool) {
	if o == nil || IsNil(o.MessageCode) {
		return nil, false
	}
	return o.MessageCode, true
}

// HasMessageCode returns a boolean if a field has been set.
func (o *DatasetMessageParams) HasMessageCode() bool {
	if o != nil && !IsNil(o.MessageCode) {
		return true
	}

	return false
}

// SetMessageCode gets a reference to the given string and assigns it to the MessageCode field.
func (o *DatasetMessageParams) SetMessageCode(v string) {
	o.MessageCode = &v
}

// GetLine returns the Line field value
func (o *DatasetMessageParams) GetLine() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Line
}

// GetLineOk returns a tuple with the Line field value
// and a boolean to check if the value has been set.
func (o *DatasetMessageParams) GetLineOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Line, true
}

// SetLine sets field value
func (o *DatasetMessageParams) SetLine(v float64) {
	o.Line = v
}

// GetFuncName returns the FuncName field value
func (o *DatasetMessageParams) GetFuncName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FuncName
}

// GetFuncNameOk returns a tuple with the FuncName field value
// and a boolean to check if the value has been set.
func (o *DatasetMessageParams) GetFuncNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FuncName, true
}

// SetFuncName sets field value
func (o *DatasetMessageParams) SetFuncName(v string) {
	o.FuncName = v
}

// GetModule returns the Module field value
func (o *DatasetMessageParams) GetModule() Module {
	if o == nil {
		var ret Module
		return ret
	}

	return o.Module
}

// GetModuleOk returns a tuple with the Module field value
// and a boolean to check if the value has been set.
func (o *DatasetMessageParams) GetModuleOk() (*Module, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Module, true
}

// SetModule sets field value
func (o *DatasetMessageParams) SetModule(v Module) {
	o.Module = v
}

func (o DatasetMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["message"] = o.Message
	if !IsNil(o.MessageCode) {
		toSerialize["message_code"] = o.MessageCode
	}
	toSerialize["line"] = o.Line
	toSerialize["func_name"] = o.FuncName
	toSerialize["module"] = o.Module
	return toSerialize, nil
}

type NullableDatasetMessageParams struct {
	value *DatasetMessageParams
	isSet bool
}

func (v NullableDatasetMessageParams) Get() *DatasetMessageParams {
	return v.value
}

func (v *NullableDatasetMessageParams) Set(val *DatasetMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetMessageParams(val *DatasetMessageParams) *NullableDatasetMessageParams {
	return &NullableDatasetMessageParams{value: val, isSet: true}
}

func (v NullableDatasetMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
