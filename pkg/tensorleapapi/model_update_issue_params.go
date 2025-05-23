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

// checks if the UpdateIssueParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateIssueParams{}

// UpdateIssueParams struct for UpdateIssueParams
type UpdateIssueParams struct {
	Cid        string          `json:"cid"`
	ProjectId  string          `json:"projectId"`
	Title      *string         `json:"title,omitempty"`
	Status     *IssueStatus    `json:"status,omitempty"`
	Branch     *string         `json:"branch,omitempty"`
	Tags       []string        `json:"tags,omitempty"`
	Assignment *string         `json:"assignment,omitempty"`
	Activities []IssueActivity `json:"activities,omitempty"`
}

// NewUpdateIssueParams instantiates a new UpdateIssueParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateIssueParams(cid string, projectId string) *UpdateIssueParams {
	this := UpdateIssueParams{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewUpdateIssueParamsWithDefaults instantiates a new UpdateIssueParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateIssueParamsWithDefaults() *UpdateIssueParams {
	this := UpdateIssueParams{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateIssueParams) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateIssueParams) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateIssueParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateIssueParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UpdateIssueParams) SetTitle(v string) {
	o.Title = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetStatus() IssueStatus {
	if o == nil || IsNil(o.Status) {
		var ret IssueStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetStatusOk() (*IssueStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given IssueStatus and assigns it to the Status field.
func (o *UpdateIssueParams) SetStatus(v IssueStatus) {
	o.Status = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *UpdateIssueParams) SetBranch(v string) {
	o.Branch = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateIssueParams) SetTags(v []string) {
	o.Tags = v
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetAssignment() string {
	if o == nil || IsNil(o.Assignment) {
		var ret string
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetAssignmentOk() (*string, bool) {
	if o == nil || IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasAssignment() bool {
	if o != nil && !IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given string and assigns it to the Assignment field.
func (o *UpdateIssueParams) SetAssignment(v string) {
	o.Assignment = &v
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *UpdateIssueParams) GetActivities() []IssueActivity {
	if o == nil || IsNil(o.Activities) {
		var ret []IssueActivity
		return ret
	}
	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateIssueParams) GetActivitiesOk() ([]IssueActivity, bool) {
	if o == nil || IsNil(o.Activities) {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *UpdateIssueParams) HasActivities() bool {
	if o != nil && !IsNil(o.Activities) {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []IssueActivity and assigns it to the Activities field.
func (o *UpdateIssueParams) SetActivities(v []IssueActivity) {
	o.Activities = v
}

func (o UpdateIssueParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateIssueParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Assignment) {
		toSerialize["assignment"] = o.Assignment
	}
	if !IsNil(o.Activities) {
		toSerialize["activities"] = o.Activities
	}
	return toSerialize, nil
}

type NullableUpdateIssueParams struct {
	value *UpdateIssueParams
	isSet bool
}

func (v NullableUpdateIssueParams) Get() *UpdateIssueParams {
	return v.value
}

func (v *NullableUpdateIssueParams) Set(val *UpdateIssueParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateIssueParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateIssueParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateIssueParams(val *UpdateIssueParams) *NullableUpdateIssueParams {
	return &NullableUpdateIssueParams{value: val, isSet: true}
}

func (v NullableUpdateIssueParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateIssueParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
