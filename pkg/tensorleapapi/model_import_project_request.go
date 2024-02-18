/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ImportProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImportProjectRequest{}

// ImportProjectRequest struct for ImportProjectRequest
type ImportProjectRequest struct {
	Name             string           `json:"name"`
	Description      string           `json:"description"`
	Tags             []string         `json:"tags"`
	HubPublishPolicy HubPublishPolicy `json:"hubPublishPolicy"`
	BgImageUrl       *string          `json:"bgImageUrl,omitempty"`
	// Construct a type with a set of properties K of type T
	Categories map[string]interface{} `json:"categories,omitempty"`
	ImportUrl  string                 `json:"importUrl"`
}

// NewImportProjectRequest instantiates a new ImportProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportProjectRequest(name string, description string, tags []string, hubPublishPolicy HubPublishPolicy, importUrl string) *ImportProjectRequest {
	this := ImportProjectRequest{}
	this.Name = name
	this.Description = description
	this.Tags = tags
	this.HubPublishPolicy = hubPublishPolicy
	this.ImportUrl = importUrl
	return &this
}

// NewImportProjectRequestWithDefaults instantiates a new ImportProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportProjectRequestWithDefaults() *ImportProjectRequest {
	this := ImportProjectRequest{}
	return &this
}

// GetName returns the Name field value
func (o *ImportProjectRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ImportProjectRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ImportProjectRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ImportProjectRequest) SetDescription(v string) {
	o.Description = v
}

// GetTags returns the Tags field value
func (o *ImportProjectRequest) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ImportProjectRequest) SetTags(v []string) {
	o.Tags = v
}

// GetHubPublishPolicy returns the HubPublishPolicy field value
func (o *ImportProjectRequest) GetHubPublishPolicy() HubPublishPolicy {
	if o == nil {
		var ret HubPublishPolicy
		return ret
	}

	return o.HubPublishPolicy
}

// GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field value
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetHubPublishPolicyOk() (*HubPublishPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubPublishPolicy, true
}

// SetHubPublishPolicy sets field value
func (o *ImportProjectRequest) SetHubPublishPolicy(v HubPublishPolicy) {
	o.HubPublishPolicy = v
}

// GetBgImageUrl returns the BgImageUrl field value if set, zero value otherwise.
func (o *ImportProjectRequest) GetBgImageUrl() string {
	if o == nil || IsNil(o.BgImageUrl) {
		var ret string
		return ret
	}
	return *o.BgImageUrl
}

// GetBgImageUrlOk returns a tuple with the BgImageUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetBgImageUrlOk() (*string, bool) {
	if o == nil || IsNil(o.BgImageUrl) {
		return nil, false
	}
	return o.BgImageUrl, true
}

// HasBgImageUrl returns a boolean if a field has been set.
func (o *ImportProjectRequest) HasBgImageUrl() bool {
	if o != nil && !IsNil(o.BgImageUrl) {
		return true
	}

	return false
}

// SetBgImageUrl gets a reference to the given string and assigns it to the BgImageUrl field.
func (o *ImportProjectRequest) SetBgImageUrl(v string) {
	o.BgImageUrl = &v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *ImportProjectRequest) GetCategories() map[string]interface{} {
	if o == nil || IsNil(o.Categories) {
		var ret map[string]interface{}
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetCategoriesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Categories) {
		return map[string]interface{}{}, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *ImportProjectRequest) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given map[string]interface{} and assigns it to the Categories field.
func (o *ImportProjectRequest) SetCategories(v map[string]interface{}) {
	o.Categories = v
}

// GetImportUrl returns the ImportUrl field value
func (o *ImportProjectRequest) GetImportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ImportUrl
}

// GetImportUrlOk returns a tuple with the ImportUrl field value
// and a boolean to check if the value has been set.
func (o *ImportProjectRequest) GetImportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ImportUrl, true
}

// SetImportUrl sets field value
func (o *ImportProjectRequest) SetImportUrl(v string) {
	o.ImportUrl = v
}

func (o ImportProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImportProjectRequest) ToMap() (map[string]interface{}, error) {
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
	toSerialize["importUrl"] = o.ImportUrl
	return toSerialize, nil
}

type NullableImportProjectRequest struct {
	value *ImportProjectRequest
	isSet bool
}

func (v NullableImportProjectRequest) Get() *ImportProjectRequest {
	return v.value
}

func (v *NullableImportProjectRequest) Set(val *ImportProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableImportProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableImportProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportProjectRequest(val *ImportProjectRequest) *NullableImportProjectRequest {
	return &NullableImportProjectRequest{value: val, isSet: true}
}

func (v NullableImportProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
