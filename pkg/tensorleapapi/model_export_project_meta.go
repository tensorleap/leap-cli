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

// checks if the ExportProjectMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportProjectMeta{}

// ExportProjectMeta struct for ExportProjectMeta
type ExportProjectMeta struct {
	Name        string   `json:"name"`
	Description *string  `json:"description,omitempty"`
	Tags        []string `json:"tags"`
	// Construct a type with a set of properties K of type T
	Categories      map[string]interface{} `json:"categories"`
	BgImagePath     string                 `json:"bgImagePath"`
	SchemaVersion   float64                `json:"schemaVersion"`
	SourceProjectId string                 `json:"sourceProjectId"`
}

// NewExportProjectMeta instantiates a new ExportProjectMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportProjectMeta(name string, tags []string, categories map[string]interface{}, bgImagePath string, schemaVersion float64, sourceProjectId string) *ExportProjectMeta {
	this := ExportProjectMeta{}
	this.Name = name
	this.Tags = tags
	this.Categories = categories
	this.BgImagePath = bgImagePath
	this.SchemaVersion = schemaVersion
	this.SourceProjectId = sourceProjectId
	return &this
}

// NewExportProjectMetaWithDefaults instantiates a new ExportProjectMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportProjectMetaWithDefaults() *ExportProjectMeta {
	this := ExportProjectMeta{}
	return &this
}

// GetName returns the Name field value
func (o *ExportProjectMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExportProjectMeta) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExportProjectMeta) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExportProjectMeta) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExportProjectMeta) SetDescription(v string) {
	o.Description = &v
}

// GetTags returns the Tags field value
func (o *ExportProjectMeta) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *ExportProjectMeta) SetTags(v []string) {
	o.Tags = v
}

// GetCategories returns the Categories field value
func (o *ExportProjectMeta) GetCategories() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetCategoriesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Categories, true
}

// SetCategories sets field value
func (o *ExportProjectMeta) SetCategories(v map[string]interface{}) {
	o.Categories = v
}

// GetBgImagePath returns the BgImagePath field value
func (o *ExportProjectMeta) GetBgImagePath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BgImagePath
}

// GetBgImagePathOk returns a tuple with the BgImagePath field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetBgImagePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BgImagePath, true
}

// SetBgImagePath sets field value
func (o *ExportProjectMeta) SetBgImagePath(v string) {
	o.BgImagePath = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *ExportProjectMeta) GetSchemaVersion() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetSchemaVersionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *ExportProjectMeta) SetSchemaVersion(v float64) {
	o.SchemaVersion = v
}

// GetSourceProjectId returns the SourceProjectId field value
func (o *ExportProjectMeta) GetSourceProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceProjectId
}

// GetSourceProjectIdOk returns a tuple with the SourceProjectId field value
// and a boolean to check if the value has been set.
func (o *ExportProjectMeta) GetSourceProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceProjectId, true
}

// SetSourceProjectId sets field value
func (o *ExportProjectMeta) SetSourceProjectId(v string) {
	o.SourceProjectId = v
}

func (o ExportProjectMeta) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportProjectMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["tags"] = o.Tags
	toSerialize["categories"] = o.Categories
	toSerialize["bgImagePath"] = o.BgImagePath
	toSerialize["schemaVersion"] = o.SchemaVersion
	toSerialize["sourceProjectId"] = o.SourceProjectId
	return toSerialize, nil
}

type NullableExportProjectMeta struct {
	value *ExportProjectMeta
	isSet bool
}

func (v NullableExportProjectMeta) Get() *ExportProjectMeta {
	return v.value
}

func (v *NullableExportProjectMeta) Set(val *ExportProjectMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableExportProjectMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableExportProjectMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportProjectMeta(val *ExportProjectMeta) *NullableExportProjectMeta {
	return &NullableExportProjectMeta{value: val, isSet: true}
}

func (v NullableExportProjectMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportProjectMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
