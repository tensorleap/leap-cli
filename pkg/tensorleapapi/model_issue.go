/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Issue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Issue{}

// Issue struct for Issue
type Issue struct {
	Cid string `json:"cid"`
	Index float64 `json:"index"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time `json:"updatedAt"`
	TeamId string `json:"teamId"`
	ProjectId string `json:"projectId"`
	Title string `json:"title"`
	Status IssueStatus `json:"status"`
	Branch string `json:"branch"`
	Tags []string `json:"tags"`
	Assignment *string `json:"assignment,omitempty"`
	Activities []IssueActivity `json:"activities"`
}

// NewIssue instantiates a new Issue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssue(cid string, index float64, createdAt time.Time, createdBy string, updatedAt time.Time, teamId string, projectId string, title string, status IssueStatus, branch string, tags []string, activities []IssueActivity) *Issue {
	this := Issue{}
	this.Cid = cid
	this.Index = index
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	this.UpdatedAt = updatedAt
	this.TeamId = teamId
	this.ProjectId = projectId
	this.Title = title
	this.Status = status
	this.Branch = branch
	this.Tags = tags
	this.Activities = activities
	return &this
}

// NewIssueWithDefaults instantiates a new Issue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueWithDefaults() *Issue {
	this := Issue{}
	return &this
}

// GetCid returns the Cid field value
func (o *Issue) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Issue) SetCid(v string) {
	o.Cid = v
}

// GetIndex returns the Index field value
func (o *Issue) GetIndex() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Index
}

// GetIndexOk returns a tuple with the Index field value
// and a boolean to check if the value has been set.
func (o *Issue) GetIndexOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Index, true
}

// SetIndex sets field value
func (o *Issue) SetIndex(v float64) {
	o.Index = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Issue) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Issue) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Issue) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Issue) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Issue) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Issue) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Issue) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Issue) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetTeamId returns the TeamId field value
func (o *Issue) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Issue) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Issue) SetTeamId(v string) {
	o.TeamId = v
}

// GetProjectId returns the ProjectId field value
func (o *Issue) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Issue) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Issue) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTitle returns the Title field value
func (o *Issue) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *Issue) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *Issue) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *Issue) GetStatus() IssueStatus {
	if o == nil {
		var ret IssueStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Issue) GetStatusOk() (*IssueStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Issue) SetStatus(v IssueStatus) {
	o.Status = v
}

// GetBranch returns the Branch field value
func (o *Issue) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *Issue) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Branch, true
}

// SetBranch sets field value
func (o *Issue) SetBranch(v string) {
	o.Branch = v
}

// GetTags returns the Tags field value
func (o *Issue) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Issue) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Issue) SetTags(v []string) {
	o.Tags = v
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *Issue) GetAssignment() string {
	if o == nil || IsNil(o.Assignment) {
		var ret string
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Issue) GetAssignmentOk() (*string, bool) {
	if o == nil || IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *Issue) HasAssignment() bool {
	if o != nil && !IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given string and assigns it to the Assignment field.
func (o *Issue) SetAssignment(v string) {
	o.Assignment = &v
}

// GetActivities returns the Activities field value
func (o *Issue) GetActivities() []IssueActivity {
	if o == nil {
		var ret []IssueActivity
		return ret
	}

	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value
// and a boolean to check if the value has been set.
func (o *Issue) GetActivitiesOk() ([]IssueActivity, bool) {
	if o == nil {
		return nil, false
	}
	return o.Activities, true
}

// SetActivities sets field value
func (o *Issue) SetActivities(v []IssueActivity) {
	o.Activities = v
}

func (o Issue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Issue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["index"] = o.Index
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["teamId"] = o.TeamId
	toSerialize["projectId"] = o.ProjectId
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
	toSerialize["branch"] = o.Branch
	toSerialize["tags"] = o.Tags
	if !IsNil(o.Assignment) {
		toSerialize["assignment"] = o.Assignment
	}
	toSerialize["activities"] = o.Activities
	return toSerialize, nil
}

type NullableIssue struct {
	value *Issue
	isSet bool
}

func (v NullableIssue) Get() *Issue {
	return v.value
}

func (v *NullableIssue) Set(val *Issue) {
	v.value = val
	v.isSet = true
}

func (v NullableIssue) IsSet() bool {
	return v.isSet
}

func (v *NullableIssue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssue(val *Issue) *NullableIssue {
	return &NullableIssue{value: val, isSet: true}
}

func (v NullableIssue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


