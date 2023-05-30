/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.321
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateProjectMetaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProjectMetaRequest{}

// UpdateProjectMetaRequest struct for UpdateProjectMetaRequest
type UpdateProjectMetaRequest struct {
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Tags             []string         `json:"tags"`
	HubPublishPolicy HubPublishPolicy `json:"hubPublishPolicy"`
	BgImageUrl       *string          `json:"bgImageUrl,omitempty"`
	ProjectId        string           `json:"projectId"`
}

// NewUpdateProjectMetaRequest instantiates a new UpdateProjectMetaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProjectMetaRequest(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy, projectId string) *UpdateProjectMetaRequest {
	this := UpdateProjectMetaRequest{}
	this.Name = name
	this.Description = description
	this.Tags = tags
	this.HubPublishPolicy = hubPublishPolicy
	this.ProjectId = projectId
	return &this
}

// NewUpdateProjectMetaRequestWithDefaults instantiates a new UpdateProjectMetaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProjectMetaRequestWithDefaults() *UpdateProjectMetaRequest {
	this := UpdateProjectMetaRequest{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateProjectMetaRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProjectMetaRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *UpdateProjectMetaRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *UpdateProjectMetaRequest) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value
func (o *UpdateProjectMetaRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *UpdateProjectMetaRequest) SetTags(v []string) {
	o.Tags = v
}

// GetHubPublishPolicy returns the HubPublishPolicy field value
func (o *UpdateProjectMetaRequest) GetHubPublishPolicy() HubPublishPolicy {
	if o == nil {
		var ret HubPublishPolicy
		return ret
	}

	return o.HubPublishPolicy
}

// GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetHubPublishPolicyOk() (*HubPublishPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubPublishPolicy, true
}

// SetHubPublishPolicy sets field value
func (o *UpdateProjectMetaRequest) SetHubPublishPolicy(v HubPublishPolicy) {
	o.HubPublishPolicy = v
}

// GetBgImageUrl returns the BgImageUrl field value if set, zero value otherwise.
func (o *UpdateProjectMetaRequest) GetBgImageUrl() string {
	if o == nil || IsNil(o.BgImageUrl) {
		var ret string
		return ret
	}
	return *o.BgImageUrl
}

// GetBgImageUrlOk returns a tuple with the BgImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetBgImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BgImageUrl) {
		return nil, false
	}
	return o.BgImageUrl, true
}

// HasBgImageUrl returns a boolean if a field has been set.
func (o *UpdateProjectMetaRequest) HasBgImageUrl() bool {
	if o != nil && !IsNil(o.BgImageUrl) {
		return true
	}

	return false
}

// SetBgImageUrl gets a reference to the given string and assigns it to the BgImageUrl field.
func (o *UpdateProjectMetaRequest) SetBgImageUrl(v string) {
	o.BgImageUrl = &v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateProjectMetaRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateProjectMetaRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateProjectMetaRequest) SetProjectId(v string) {
	o.ProjectId = v
}

func (o UpdateProjectMetaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProjectMetaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["tags"] = o.Tags
	toSerialize["hubPublishPolicy"] = o.HubPublishPolicy
	if !IsNil(o.BgImageUrl) {
		toSerialize["bgImageUrl"] = o.BgImageUrl
	}
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableUpdateProjectMetaRequest struct {
	value *UpdateProjectMetaRequest
	isSet bool
}

func (v NullableUpdateProjectMetaRequest) Get() *UpdateProjectMetaRequest {
	return v.value
}

func (v *NullableUpdateProjectMetaRequest) Set(val *UpdateProjectMetaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProjectMetaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProjectMetaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProjectMetaRequest(val *UpdateProjectMetaRequest) *NullableUpdateProjectMetaRequest {
	return &NullableUpdateProjectMetaRequest{value: val, isSet: true}
}

func (v NullableUpdateProjectMetaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProjectMetaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
