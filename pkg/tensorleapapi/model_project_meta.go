/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ProjectMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectMeta{}

// ProjectMeta struct for ProjectMeta
type ProjectMeta struct {
	Name string `json:"name"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
	HubPublishPolicy HubPublishPolicy `json:"hubPublishPolicy"`
	BgImageUrl *string `json:"bgImageUrl,omitempty"`
	// Construct a type with a set of properties K of type T
	Categories map[string]interface{} `json:"categories,omitempty"`
}

// NewProjectMeta instantiates a new ProjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectMeta(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy) *ProjectMeta {
	this := ProjectMeta{}
	this.Name = name
	this.Description = description
	this.Tags = tags
	this.HubPublishPolicy = hubPublishPolicy
	return &this
}

// NewProjectMetaWithDefaults instantiates a new ProjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectMetaWithDefaults() *ProjectMeta {
	this := ProjectMeta{}
	return &this
}

// GetName returns the Name field value
func (o *ProjectMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ProjectMeta) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ProjectMeta) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ProjectMeta) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value
func (o *ProjectMeta) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ProjectMeta) SetTags(v []string) {
	o.Tags = v
}

// GetHubPublishPolicy returns the HubPublishPolicy field value
func (o *ProjectMeta) GetHubPublishPolicy() HubPublishPolicy {
	if o == nil {
		var ret HubPublishPolicy
		return ret
	}

	return o.HubPublishPolicy
}

// GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field value
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetHubPublishPolicyOk() (*HubPublishPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubPublishPolicy, true
}

// SetHubPublishPolicy sets field value
func (o *ProjectMeta) SetHubPublishPolicy(v HubPublishPolicy) {
	o.HubPublishPolicy = v
}

// GetBgImageUrl returns the BgImageUrl field value if set, zero value otherwise.
func (o *ProjectMeta) GetBgImageUrl() string {
	if o == nil || IsNil(o.BgImageUrl) {
		var ret string
		return ret
	}
	return *o.BgImageUrl
}

// GetBgImageUrlOk returns a tuple with the BgImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetBgImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BgImageUrl) {
		return nil, false
	}
	return o.BgImageUrl, true
}

// HasBgImageUrl returns a boolean if a field has been set.
func (o *ProjectMeta) HasBgImageUrl() bool {
	if o != nil && !IsNil(o.BgImageUrl) {
		return true
	}

	return false
}

// SetBgImageUrl gets a reference to the given string and assigns it to the BgImageUrl field.
func (o *ProjectMeta) SetBgImageUrl(v string) {
	o.BgImageUrl = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ProjectMeta) GetCategories() map[string]interface{} {
	if o == nil || IsNil(o.Categories) {
		var ret map[string]interface{}
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectMeta) GetCategoriesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Categories) {
		return map[string]interface{}{}, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ProjectMeta) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given map[string]interface{} and assigns it to the Categories field.
func (o *ProjectMeta) SetCategories(v map[string]interface{}) {
	o.Categories = v
}

func (o ProjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["tags"] = o.Tags
	toSerialize["hubPublishPolicy"] = o.HubPublishPolicy
	if !IsNil(o.BgImageUrl) {
		toSerialize["bgImageUrl"] = o.BgImageUrl
	}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	return toSerialize, nil
}

type NullableProjectMeta struct {
	value *ProjectMeta
	isSet bool
}

func (v NullableProjectMeta) Get() *ProjectMeta {
	return v.value
}

func (v *NullableProjectMeta) Set(val *ProjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectMeta(val *ProjectMeta) *NullableProjectMeta {
	return &NullableProjectMeta{value: val, isSet: true}
}

func (v NullableProjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


