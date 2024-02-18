/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the IssueActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueActivity{}

// IssueActivity struct for IssueActivity
type IssueActivity struct {
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`
	CreatedBy  string          `json:"createdBy"`
	ActionType IssueActionType `json:"actionType"`
	Action     IssueAction     `json:"action"`
}

// NewIssueActivity instantiates a new IssueActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueActivity(createdAt time.Time, updatedAt time.Time, createdBy string, actionType IssueActionType, action IssueAction) *IssueActivity {
	this := IssueActivity{}
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.CreatedBy = createdBy
	this.ActionType = actionType
	this.Action = action
	return &this
}

// NewIssueActivityWithDefaults instantiates a new IssueActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueActivityWithDefaults() *IssueActivity {
	this := IssueActivity{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *IssueActivity) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *IssueActivity) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *IssueActivity) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *IssueActivity) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *IssueActivity) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *IssueActivity) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *IssueActivity) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *IssueActivity) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *IssueActivity) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetActionType returns the ActionType field value
func (o *IssueActivity) GetActionType() IssueActionType {
	if o == nil {
		var ret IssueActionType
		return ret
	}

	return o.ActionType
}

// GetActionTypeOk returns a tuple with the ActionType field value
// and a boolean to check if the value has been set.
func (o *IssueActivity) GetActionTypeOk() (*IssueActionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActionType, true
}

// SetActionType sets field value
func (o *IssueActivity) SetActionType(v IssueActionType) {
	o.ActionType = v
}

// GetAction returns the Action field value
func (o *IssueActivity) GetAction() IssueAction {
	if o == nil {
		var ret IssueAction
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *IssueActivity) GetActionOk() (*IssueAction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *IssueActivity) SetAction(v IssueAction) {
	o.Action = v
}

func (o IssueActivity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["actionType"] = o.ActionType
	toSerialize["action"] = o.Action
	return toSerialize, nil
}

type NullableIssueActivity struct {
	value *IssueActivity
	isSet bool
}

func (v NullableIssueActivity) Get() *IssueActivity {
	return v.value
}

func (v *NullableIssueActivity) Set(val *IssueActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueActivity(val *IssueActivity) *NullableIssueActivity {
	return &NullableIssueActivity{value: val, isSet: true}
}

func (v NullableIssueActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
