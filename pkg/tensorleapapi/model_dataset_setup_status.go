/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetSetupStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetSetupStatus{}

// DatasetSetupStatus struct for DatasetSetupStatus
type DatasetSetupStatus struct {
	GeneralError  *string                    `json:"generalError,omitempty"`
	PrintLog      *string                    `json:"printLog,omitempty"`
	BindersStatus []DatasetTestResultPayload `json:"bindersStatus"`
}

// NewDatasetSetupStatus instantiates a new DatasetSetupStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetSetupStatus(bindersStatus []DatasetTestResultPayload) *DatasetSetupStatus {
	this := DatasetSetupStatus{}
	this.BindersStatus = bindersStatus
	return &this
}

// NewDatasetSetupStatusWithDefaults instantiates a new DatasetSetupStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetSetupStatusWithDefaults() *DatasetSetupStatus {
	this := DatasetSetupStatus{}
	return &this
}

// GetGeneralError returns the GeneralError field value if set, zero value otherwise.
func (o *DatasetSetupStatus) GetGeneralError() string {
	if o == nil || IsNil(o.GeneralError) {
		var ret string
		return ret
	}
	return *o.GeneralError
}

// GetGeneralErrorOk returns a tuple with the GeneralError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetSetupStatus) GetGeneralErrorOk() (*string, bool) {
	if o == nil || IsNil(o.GeneralError) {
		return nil, false
	}
	return o.GeneralError, true
}

// HasGeneralError returns a boolean if a field has been set.
func (o *DatasetSetupStatus) HasGeneralError() bool {
	if o != nil && !IsNil(o.GeneralError) {
		return true
	}

	return false
}

// SetGeneralError gets a reference to the given string and assigns it to the GeneralError field.
func (o *DatasetSetupStatus) SetGeneralError(v string) {
	o.GeneralError = &v
}

// GetPrintLog returns the PrintLog field value if set, zero value otherwise.
func (o *DatasetSetupStatus) GetPrintLog() string {
	if o == nil || IsNil(o.PrintLog) {
		var ret string
		return ret
	}
	return *o.PrintLog
}

// GetPrintLogOk returns a tuple with the PrintLog field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetSetupStatus) GetPrintLogOk() (*string, bool) {
	if o == nil || IsNil(o.PrintLog) {
		return nil, false
	}
	return o.PrintLog, true
}

// HasPrintLog returns a boolean if a field has been set.
func (o *DatasetSetupStatus) HasPrintLog() bool {
	if o != nil && !IsNil(o.PrintLog) {
		return true
	}

	return false
}

// SetPrintLog gets a reference to the given string and assigns it to the PrintLog field.
func (o *DatasetSetupStatus) SetPrintLog(v string) {
	o.PrintLog = &v
}

// GetBindersStatus returns the BindersStatus field value
func (o *DatasetSetupStatus) GetBindersStatus() []DatasetTestResultPayload {
	if o == nil {
		var ret []DatasetTestResultPayload
		return ret
	}

	return o.BindersStatus
}

// GetBindersStatusOk returns a tuple with the BindersStatus field value
// and a boolean to check if the value has been set.
func (o *DatasetSetupStatus) GetBindersStatusOk() ([]DatasetTestResultPayload, bool) {
	if o == nil {
		return nil, false
	}
	return o.BindersStatus, true
}

// SetBindersStatus sets field value
func (o *DatasetSetupStatus) SetBindersStatus(v []DatasetTestResultPayload) {
	o.BindersStatus = v
}

func (o DatasetSetupStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetSetupStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GeneralError) {
		toSerialize["generalError"] = o.GeneralError
	}
	if !IsNil(o.PrintLog) {
		toSerialize["printLog"] = o.PrintLog
	}
	toSerialize["bindersStatus"] = o.BindersStatus
	return toSerialize, nil
}

type NullableDatasetSetupStatus struct {
	value *DatasetSetupStatus
	isSet bool
}

func (v NullableDatasetSetupStatus) Get() *DatasetSetupStatus {
	return v.value
}

func (v *NullableDatasetSetupStatus) Set(val *DatasetSetupStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetSetupStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetSetupStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetSetupStatus(val *DatasetSetupStatus) *NullableDatasetSetupStatus {
	return &NullableDatasetSetupStatus{value: val, isSet: true}
}

func (v NullableDatasetSetupStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetSetupStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
