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

// checks if the SessionTestResultSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionTestResultSuccess{}

// SessionTestResultSuccess struct for SessionTestResultSuccess
type SessionTestResultSuccess struct {
	TestSucceeded    bool    `json:"testSucceeded"`
	SessionRunId     string  `json:"sessionRunId"`
	Aggregation      float64 `json:"aggregation"`
	SuccefullSamples float64 `json:"succefullSamples"`
	AllSamples       float64 `json:"allSamples"`
	Epoch            float64 `json:"epoch"`
	QueryStatus      string  `json:"queryStatus"`
}

// NewSessionTestResultSuccess instantiates a new SessionTestResultSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionTestResultSuccess(testSucceeded bool, sessionRunId string, aggregation float64, succefullSamples float64, allSamples float64, epoch float64, queryStatus string) *SessionTestResultSuccess {
	this := SessionTestResultSuccess{}
	this.TestSucceeded = testSucceeded
	this.SessionRunId = sessionRunId
	this.Aggregation = aggregation
	this.SuccefullSamples = succefullSamples
	this.AllSamples = allSamples
	this.Epoch = epoch
	this.QueryStatus = queryStatus
	return &this
}

// NewSessionTestResultSuccessWithDefaults instantiates a new SessionTestResultSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionTestResultSuccessWithDefaults() *SessionTestResultSuccess {
	this := SessionTestResultSuccess{}
	return &this
}

// GetTestSucceeded returns the TestSucceeded field value
func (o *SessionTestResultSuccess) GetTestSucceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TestSucceeded
}

// GetTestSucceededOk returns a tuple with the TestSucceeded field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetTestSucceededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestSucceeded, true
}

// SetTestSucceeded sets field value
func (o *SessionTestResultSuccess) SetTestSucceeded(v bool) {
	o.TestSucceeded = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *SessionTestResultSuccess) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *SessionTestResultSuccess) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetAggregation returns the Aggregation field value
func (o *SessionTestResultSuccess) GetAggregation() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Aggregation
}

// GetAggregationOk returns a tuple with the Aggregation field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetAggregationOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Aggregation, true
}

// SetAggregation sets field value
func (o *SessionTestResultSuccess) SetAggregation(v float64) {
	o.Aggregation = v
}

// GetSuccefullSamples returns the SuccefullSamples field value
func (o *SessionTestResultSuccess) GetSuccefullSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SuccefullSamples
}

// GetSuccefullSamplesOk returns a tuple with the SuccefullSamples field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetSuccefullSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuccefullSamples, true
}

// SetSuccefullSamples sets field value
func (o *SessionTestResultSuccess) SetSuccefullSamples(v float64) {
	o.SuccefullSamples = v
}

// GetAllSamples returns the AllSamples field value
func (o *SessionTestResultSuccess) GetAllSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AllSamples
}

// GetAllSamplesOk returns a tuple with the AllSamples field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetAllSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllSamples, true
}

// SetAllSamples sets field value
func (o *SessionTestResultSuccess) SetAllSamples(v float64) {
	o.AllSamples = v
}

// GetEpoch returns the Epoch field value
func (o *SessionTestResultSuccess) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *SessionTestResultSuccess) SetEpoch(v float64) {
	o.Epoch = v
}

// GetQueryStatus returns the QueryStatus field value
func (o *SessionTestResultSuccess) GetQueryStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryStatus
}

// GetQueryStatusOk returns a tuple with the QueryStatus field value
// and a boolean to check if the value has been set.
func (o *SessionTestResultSuccess) GetQueryStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryStatus, true
}

// SetQueryStatus sets field value
func (o *SessionTestResultSuccess) SetQueryStatus(v string) {
	o.QueryStatus = v
}

func (o SessionTestResultSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionTestResultSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["testSucceeded"] = o.TestSucceeded
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["aggregation"] = o.Aggregation
	toSerialize["succefullSamples"] = o.SuccefullSamples
	toSerialize["allSamples"] = o.AllSamples
	toSerialize["epoch"] = o.Epoch
	toSerialize["queryStatus"] = o.QueryStatus
	return toSerialize, nil
}

type NullableSessionTestResultSuccess struct {
	value *SessionTestResultSuccess
	isSet bool
}

func (v NullableSessionTestResultSuccess) Get() *SessionTestResultSuccess {
	return v.value
}

func (v *NullableSessionTestResultSuccess) Set(val *SessionTestResultSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionTestResultSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionTestResultSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionTestResultSuccess(val *SessionTestResultSuccess) *NullableSessionTestResultSuccess {
	return &NullableSessionTestResultSuccess{value: val, isSet: true}
}

func (v NullableSessionTestResultSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionTestResultSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
