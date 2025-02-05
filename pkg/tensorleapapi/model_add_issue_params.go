/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddIssueParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIssueParams{}

// AddIssueParams struct for AddIssueParams
type AddIssueParams struct {
	ProjectId  string          `json:"projectId"`
	Title      string          `json:"title"`
	Status     IssueStatus     `json:"status"`
	Branch     *string         `json:"branch,omitempty"`
	Tags       []string        `json:"tags,omitempty"`
	Assignment *string         `json:"assignment,omitempty"`
	Activities []IssueActivity `json:"activities,omitempty"`
}

// NewAddIssueParams instantiates a new AddIssueParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIssueParams(projectId string, title string, status IssueStatus) *AddIssueParams {
	this := AddIssueParams{}
	this.ProjectId = projectId
	this.Title = title
	this.Status = status
	return &this
}

// NewAddIssueParamsWithDefaults instantiates a new AddIssueParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIssueParamsWithDefaults() *AddIssueParams {
	this := AddIssueParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *AddIssueParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddIssueParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetTitle returns the Title field value
func (o *AddIssueParams) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AddIssueParams) SetTitle(v string) {
	o.Title = v
}

// GetStatus returns the Status field value
func (o *AddIssueParams) GetStatus() IssueStatus {
	if o == nil {
		var ret IssueStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetStatusOk() (*IssueStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AddIssueParams) SetStatus(v IssueStatus) {
	o.Status = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *AddIssueParams) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *AddIssueParams) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *AddIssueParams) SetBranch(v string) {
	o.Branch = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AddIssueParams) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AddIssueParams) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AddIssueParams) SetTags(v []string) {
	o.Tags = v
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *AddIssueParams) GetAssignment() string {
	if o == nil || IsNil(o.Assignment) {
		var ret string
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetAssignmentOk() (*string, bool) {
	if o == nil || IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *AddIssueParams) HasAssignment() bool {
	if o != nil && !IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given string and assigns it to the Assignment field.
func (o *AddIssueParams) SetAssignment(v string) {
	o.Assignment = &v
}

// GetActivities returns the Activities field value if set, zero value otherwise.
func (o *AddIssueParams) GetActivities() []IssueActivity {
	if o == nil || IsNil(o.Activities) {
		var ret []IssueActivity
		return ret
	}
	return o.Activities
}

// GetActivitiesOk returns a tuple with the Activities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIssueParams) GetActivitiesOk() ([]IssueActivity, bool) {
	if o == nil || IsNil(o.Activities) {
		return nil, false
	}
	return o.Activities, true
}

// HasActivities returns a boolean if a field has been set.
func (o *AddIssueParams) HasActivities() bool {
	if o != nil && !IsNil(o.Activities) {
		return true
	}

	return false
}

// SetActivities gets a reference to the given []IssueActivity and assigns it to the Activities field.
func (o *AddIssueParams) SetActivities(v []IssueActivity) {
	o.Activities = v
}

func (o AddIssueParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIssueParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["title"] = o.Title
	toSerialize["status"] = o.Status
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

type NullableAddIssueParams struct {
	value *AddIssueParams
	isSet bool
}

func (v NullableAddIssueParams) Get() *AddIssueParams {
	return v.value
}

func (v *NullableAddIssueParams) Set(val *AddIssueParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddIssueParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddIssueParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddIssueParams(val *AddIssueParams) *NullableAddIssueParams {
	return &NullableAddIssueParams{value: val, isSet: true}
}

func (v NullableAddIssueParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddIssueParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
