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

// checks if the GetGeneratedDashletsBySessionRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetGeneratedDashletsBySessionRun{}

// GetGeneratedDashletsBySessionRun struct for GetGeneratedDashletsBySessionRun
type GetGeneratedDashletsBySessionRun struct {
	Dashlets     []GeneratedDashlet `json:"dashlets"`
	SessionRunId string             `json:"sessionRunId"`
	Status       JobStatus          `json:"status"`
}

// NewGetGeneratedDashletsBySessionRun instantiates a new GetGeneratedDashletsBySessionRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetGeneratedDashletsBySessionRun(dashlets []GeneratedDashlet, sessionRunId string, status JobStatus) *GetGeneratedDashletsBySessionRun {
	this := GetGeneratedDashletsBySessionRun{}
	this.Dashlets = dashlets
	this.SessionRunId = sessionRunId
	this.Status = status
	return &this
}

// NewGetGeneratedDashletsBySessionRunWithDefaults instantiates a new GetGeneratedDashletsBySessionRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetGeneratedDashletsBySessionRunWithDefaults() *GetGeneratedDashletsBySessionRun {
	this := GetGeneratedDashletsBySessionRun{}
	return &this
}

// GetDashlets returns the Dashlets field value
func (o *GetGeneratedDashletsBySessionRun) GetDashlets() []GeneratedDashlet {
	if o == nil {
		var ret []GeneratedDashlet
		return ret
	}

	return o.Dashlets
}

// GetDashletsOk returns a tuple with the Dashlets field value
// and a boolean to check if the value has been set.
func (o *GetGeneratedDashletsBySessionRun) GetDashletsOk() ([]GeneratedDashlet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dashlets, true
}

// SetDashlets sets field value
func (o *GetGeneratedDashletsBySessionRun) SetDashlets(v []GeneratedDashlet) {
	o.Dashlets = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *GetGeneratedDashletsBySessionRun) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *GetGeneratedDashletsBySessionRun) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *GetGeneratedDashletsBySessionRun) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetStatus returns the Status field value
func (o *GetGeneratedDashletsBySessionRun) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetGeneratedDashletsBySessionRun) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetGeneratedDashletsBySessionRun) SetStatus(v JobStatus) {
	o.Status = v
}

func (o GetGeneratedDashletsBySessionRun) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetGeneratedDashletsBySessionRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashlets"] = o.Dashlets
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableGetGeneratedDashletsBySessionRun struct {
	value *GetGeneratedDashletsBySessionRun
	isSet bool
}

func (v NullableGetGeneratedDashletsBySessionRun) Get() *GetGeneratedDashletsBySessionRun {
	return v.value
}

func (v *NullableGetGeneratedDashletsBySessionRun) Set(val *GetGeneratedDashletsBySessionRun) {
	v.value = val
	v.isSet = true
}

func (v NullableGetGeneratedDashletsBySessionRun) IsSet() bool {
	return v.isSet
}

func (v *NullableGetGeneratedDashletsBySessionRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetGeneratedDashletsBySessionRun(val *GetGeneratedDashletsBySessionRun) *NullableGetGeneratedDashletsBySessionRun {
	return &NullableGetGeneratedDashletsBySessionRun{value: val, isSet: true}
}

func (v NullableGetGeneratedDashletsBySessionRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetGeneratedDashletsBySessionRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
