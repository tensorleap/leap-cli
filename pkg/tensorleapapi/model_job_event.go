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

// checks if the JobEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JobEvent{}

// JobEvent struct for JobEvent
type JobEvent struct {
	Id       string            `json:"id"`
	Name     string            `json:"name"`
	Status   StatusEnum        `json:"status"`
	Progress *JobEventProgress `json:"progress,omitempty"`
}

// NewJobEvent instantiates a new JobEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobEvent(id string, name string, status StatusEnum) *JobEvent {
	this := JobEvent{}
	this.Id = id
	this.Name = name
	this.Status = status
	return &this
}

// NewJobEventWithDefaults instantiates a new JobEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobEventWithDefaults() *JobEvent {
	this := JobEvent{}
	return &this
}

// GetId returns the Id field value
func (o *JobEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobEvent) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *JobEvent) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobEvent) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobEvent) SetName(v string) {
	o.Name = v
}

// GetStatus returns the Status field value
func (o *JobEvent) GetStatus() StatusEnum {
	if o == nil {
		var ret StatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *JobEvent) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobEvent) SetStatus(v StatusEnum) {
	o.Status = v
}

// GetProgress returns the Progress field value if set, zero value otherwise.
func (o *JobEvent) GetProgress() JobEventProgress {
	if o == nil || IsNil(o.Progress) {
		var ret JobEventProgress
		return ret
	}
	return *o.Progress
}

// GetProgressOk returns a tuple with the Progress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobEvent) GetProgressOk() (*JobEventProgress, bool) {
	if o == nil || IsNil(o.Progress) {
		return nil, false
	}
	return o.Progress, true
}

// HasProgress returns a boolean if a field has been set.
func (o *JobEvent) HasProgress() bool {
	if o != nil && !IsNil(o.Progress) {
		return true
	}

	return false
}

// SetProgress gets a reference to the given JobEventProgress and assigns it to the Progress field.
func (o *JobEvent) SetProgress(v JobEventProgress) {
	o.Progress = &v
}

func (o JobEvent) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JobEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["status"] = o.Status
	if !IsNil(o.Progress) {
		toSerialize["progress"] = o.Progress
	}
	return toSerialize, nil
}

type NullableJobEvent struct {
	value *JobEvent
	isSet bool
}

func (v NullableJobEvent) Get() *JobEvent {
	return v.value
}

func (v *NullableJobEvent) Set(val *JobEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableJobEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableJobEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobEvent(val *JobEvent) *NullableJobEvent {
	return &NullableJobEvent{value: val, isSet: true}
}

func (v NullableJobEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
