/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the GetJobsFilterParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJobsFilterParams{}

// GetJobsFilterParams struct for GetJobsFilterParams
type GetJobsFilterParams struct {
	LastUpdated      *time.Time   `json:"lastUpdated,omitempty"`
	SessionRunIds    []string     `json:"sessionRunIds,omitempty"`
	Types            []JobType    `json:"types,omitempty"`
	SubTypes         []JobSubType `json:"subTypes,omitempty"`
	Trigger          *JobTrigger  `json:"trigger,omitempty"`
	Status           []JobStatus  `json:"status,omitempty"`
	ProjectId        *string      `json:"projectId,omitempty"`
	ShowJobsFromDate *time.Time   `json:"showJobsFromDate,omitempty"`
	Cid              []string     `json:"cid,omitempty"`
}

// NewGetJobsFilterParams instantiates a new GetJobsFilterParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJobsFilterParams() *GetJobsFilterParams {
	this := GetJobsFilterParams{}
	return &this
}

// NewGetJobsFilterParamsWithDefaults instantiates a new GetJobsFilterParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJobsFilterParamsWithDefaults() *GetJobsFilterParams {
	this := GetJobsFilterParams{}
	return &this
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GetJobsFilterParams) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetSessionRunIds returns the SessionRunIds field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetSessionRunIds() []string {
	if o == nil || IsNil(o.SessionRunIds) {
		var ret []string
		return ret
	}
	return o.SessionRunIds
}

// GetSessionRunIdsOk returns a tuple with the SessionRunIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetSessionRunIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SessionRunIds) {
		return nil, false
	}
	return o.SessionRunIds, true
}

// HasSessionRunIds returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasSessionRunIds() bool {
	if o != nil && !IsNil(o.SessionRunIds) {
		return true
	}

	return false
}

// SetSessionRunIds gets a reference to the given []string and assigns it to the SessionRunIds field.
func (o *GetJobsFilterParams) SetSessionRunIds(v []string) {
	o.SessionRunIds = v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetTypes() []JobType {
	if o == nil || IsNil(o.Types) {
		var ret []JobType
		return ret
	}
	return o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetTypesOk() ([]JobType, bool) {
	if o == nil || IsNil(o.Types) {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasTypes() bool {
	if o != nil && !IsNil(o.Types) {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []JobType and assigns it to the Types field.
func (o *GetJobsFilterParams) SetTypes(v []JobType) {
	o.Types = v
}

// GetSubTypes returns the SubTypes field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetSubTypes() []JobSubType {
	if o == nil || IsNil(o.SubTypes) {
		var ret []JobSubType
		return ret
	}
	return o.SubTypes
}

// GetSubTypesOk returns a tuple with the SubTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetSubTypesOk() ([]JobSubType, bool) {
	if o == nil || IsNil(o.SubTypes) {
		return nil, false
	}
	return o.SubTypes, true
}

// HasSubTypes returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasSubTypes() bool {
	if o != nil && !IsNil(o.SubTypes) {
		return true
	}

	return false
}

// SetSubTypes gets a reference to the given []JobSubType and assigns it to the SubTypes field.
func (o *GetJobsFilterParams) SetSubTypes(v []JobSubType) {
	o.SubTypes = v
}

// GetTrigger returns the Trigger field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetTrigger() JobTrigger {
	if o == nil || IsNil(o.Trigger) {
		var ret JobTrigger
		return ret
	}
	return *o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetTriggerOk() (*JobTrigger, bool) {
	if o == nil || IsNil(o.Trigger) {
		return nil, false
	}
	return o.Trigger, true
}

// HasTrigger returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasTrigger() bool {
	if o != nil && !IsNil(o.Trigger) {
		return true
	}

	return false
}

// SetTrigger gets a reference to the given JobTrigger and assigns it to the Trigger field.
func (o *GetJobsFilterParams) SetTrigger(v JobTrigger) {
	o.Trigger = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetStatus() []JobStatus {
	if o == nil || IsNil(o.Status) {
		var ret []JobStatus
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetStatusOk() ([]JobStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given []JobStatus and assigns it to the Status field.
func (o *GetJobsFilterParams) SetStatus(v []JobStatus) {
	o.Status = v
}

// GetProjectId returns the ProjectId field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetProjectId() string {
	if o == nil || IsNil(o.ProjectId) {
		var ret string
		return ret
	}
	return *o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetProjectIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProjectId) {
		return nil, false
	}
	return o.ProjectId, true
}

// HasProjectId returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasProjectId() bool {
	if o != nil && !IsNil(o.ProjectId) {
		return true
	}

	return false
}

// SetProjectId gets a reference to the given string and assigns it to the ProjectId field.
func (o *GetJobsFilterParams) SetProjectId(v string) {
	o.ProjectId = &v
}

// GetShowJobsFromDate returns the ShowJobsFromDate field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetShowJobsFromDate() time.Time {
	if o == nil || IsNil(o.ShowJobsFromDate) {
		var ret time.Time
		return ret
	}
	return *o.ShowJobsFromDate
}

// GetShowJobsFromDateOk returns a tuple with the ShowJobsFromDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetShowJobsFromDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ShowJobsFromDate) {
		return nil, false
	}
	return o.ShowJobsFromDate, true
}

// HasShowJobsFromDate returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasShowJobsFromDate() bool {
	if o != nil && !IsNil(o.ShowJobsFromDate) {
		return true
	}

	return false
}

// SetShowJobsFromDate gets a reference to the given time.Time and assigns it to the ShowJobsFromDate field.
func (o *GetJobsFilterParams) SetShowJobsFromDate(v time.Time) {
	o.ShowJobsFromDate = &v
}

// GetCid returns the Cid field value if set, zero value otherwise.
func (o *GetJobsFilterParams) GetCid() []string {
	if o == nil || IsNil(o.Cid) {
		var ret []string
		return ret
	}
	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsFilterParams) GetCidOk() ([]string, bool) {
	if o == nil || IsNil(o.Cid) {
		return nil, false
	}
	return o.Cid, true
}

// HasCid returns a boolean if a field has been set.
func (o *GetJobsFilterParams) HasCid() bool {
	if o != nil && !IsNil(o.Cid) {
		return true
	}

	return false
}

// SetCid gets a reference to the given []string and assigns it to the Cid field.
func (o *GetJobsFilterParams) SetCid(v []string) {
	o.Cid = v
}

func (o GetJobsFilterParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJobsFilterParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.SessionRunIds) {
		toSerialize["sessionRunIds"] = o.SessionRunIds
	}
	if !IsNil(o.Types) {
		toSerialize["types"] = o.Types
	}
	if !IsNil(o.SubTypes) {
		toSerialize["subTypes"] = o.SubTypes
	}
	if !IsNil(o.Trigger) {
		toSerialize["trigger"] = o.Trigger
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.ProjectId) {
		toSerialize["projectId"] = o.ProjectId
	}
	if !IsNil(o.ShowJobsFromDate) {
		toSerialize["showJobsFromDate"] = o.ShowJobsFromDate
	}
	if !IsNil(o.Cid) {
		toSerialize["cid"] = o.Cid
	}
	return toSerialize, nil
}

type NullableGetJobsFilterParams struct {
	value *GetJobsFilterParams
	isSet bool
}

func (v NullableGetJobsFilterParams) Get() *GetJobsFilterParams {
	return v.value
}

func (v *NullableGetJobsFilterParams) Set(val *GetJobsFilterParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJobsFilterParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJobsFilterParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJobsFilterParams(val *GetJobsFilterParams) *NullableGetJobsFilterParams {
	return &NullableGetJobsFilterParams{value: val, isSet: true}
}

func (v NullableGetJobsFilterParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJobsFilterParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
