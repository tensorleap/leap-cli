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

// checks if the IssueAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IssueAction{}

// IssueAction struct for IssueAction
type IssueAction struct {
	Assignment *ActionAssignmentElement `json:"assignment,omitempty"`
	TagsAdded  *ActionTagElement        `json:"tagsAdded,omitempty"`
	Comment    *ActionCommentElement    `json:"comment,omitempty"`
}

// NewIssueAction instantiates a new IssueAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueAction() *IssueAction {
	this := IssueAction{}
	return &this
}

// NewIssueActionWithDefaults instantiates a new IssueAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueActionWithDefaults() *IssueAction {
	this := IssueAction{}
	return &this
}

// GetAssignment returns the Assignment field value if set, zero value otherwise.
func (o *IssueAction) GetAssignment() ActionAssignmentElement {
	if o == nil || IsNil(o.Assignment) {
		var ret ActionAssignmentElement
		return ret
	}
	return *o.Assignment
}

// GetAssignmentOk returns a tuple with the Assignment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAction) GetAssignmentOk() (*ActionAssignmentElement, bool) {
	if o == nil || IsNil(o.Assignment) {
		return nil, false
	}
	return o.Assignment, true
}

// HasAssignment returns a boolean if a field has been set.
func (o *IssueAction) HasAssignment() bool {
	if o != nil && !IsNil(o.Assignment) {
		return true
	}

	return false
}

// SetAssignment gets a reference to the given ActionAssignmentElement and assigns it to the Assignment field.
func (o *IssueAction) SetAssignment(v ActionAssignmentElement) {
	o.Assignment = &v
}

// GetTagsAdded returns the TagsAdded field value if set, zero value otherwise.
func (o *IssueAction) GetTagsAdded() ActionTagElement {
	if o == nil || IsNil(o.TagsAdded) {
		var ret ActionTagElement
		return ret
	}
	return *o.TagsAdded
}

// GetTagsAddedOk returns a tuple with the TagsAdded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAction) GetTagsAddedOk() (*ActionTagElement, bool) {
	if o == nil || IsNil(o.TagsAdded) {
		return nil, false
	}
	return o.TagsAdded, true
}

// HasTagsAdded returns a boolean if a field has been set.
func (o *IssueAction) HasTagsAdded() bool {
	if o != nil && !IsNil(o.TagsAdded) {
		return true
	}

	return false
}

// SetTagsAdded gets a reference to the given ActionTagElement and assigns it to the TagsAdded field.
func (o *IssueAction) SetTagsAdded(v ActionTagElement) {
	o.TagsAdded = &v
}

// GetComment returns the Comment field value if set, zero value otherwise.
func (o *IssueAction) GetComment() ActionCommentElement {
	if o == nil || IsNil(o.Comment) {
		var ret ActionCommentElement
		return ret
	}
	return *o.Comment
}

// GetCommentOk returns a tuple with the Comment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueAction) GetCommentOk() (*ActionCommentElement, bool) {
	if o == nil || IsNil(o.Comment) {
		return nil, false
	}
	return o.Comment, true
}

// HasComment returns a boolean if a field has been set.
func (o *IssueAction) HasComment() bool {
	if o != nil && !IsNil(o.Comment) {
		return true
	}

	return false
}

// SetComment gets a reference to the given ActionCommentElement and assigns it to the Comment field.
func (o *IssueAction) SetComment(v ActionCommentElement) {
	o.Comment = &v
}

func (o IssueAction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IssueAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assignment) {
		toSerialize["assignment"] = o.Assignment
	}
	if !IsNil(o.TagsAdded) {
		toSerialize["tagsAdded"] = o.TagsAdded
	}
	if !IsNil(o.Comment) {
		toSerialize["comment"] = o.Comment
	}
	return toSerialize, nil
}

type NullableIssueAction struct {
	value *IssueAction
	isSet bool
}

func (v NullableIssueAction) Get() *IssueAction {
	return v.value
}

func (v *NullableIssueAction) Set(val *IssueAction) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueAction) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueAction(val *IssueAction) *NullableIssueAction {
	return &NullableIssueAction{value: val, isSet: true}
}

func (v NullableIssueAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
