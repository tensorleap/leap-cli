/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SaveAnalyzerLayoutParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveAnalyzerLayoutParams{}

// SaveAnalyzerLayoutParams struct for SaveAnalyzerLayoutParams
type SaveAnalyzerLayoutParams struct {
	PanelsLayouts []PanelLayout `json:"panelsLayouts"`
	ProjectId     string        `json:"projectId"`
}

// NewSaveAnalyzerLayoutParams instantiates a new SaveAnalyzerLayoutParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveAnalyzerLayoutParams(panelsLayouts []PanelLayout, projectId string) *SaveAnalyzerLayoutParams {
	this := SaveAnalyzerLayoutParams{}
	this.PanelsLayouts = panelsLayouts
	this.ProjectId = projectId
	return &this
}

// NewSaveAnalyzerLayoutParamsWithDefaults instantiates a new SaveAnalyzerLayoutParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveAnalyzerLayoutParamsWithDefaults() *SaveAnalyzerLayoutParams {
	this := SaveAnalyzerLayoutParams{}
	return &this
}

// GetPanelsLayouts returns the PanelsLayouts field value
func (o *SaveAnalyzerLayoutParams) GetPanelsLayouts() []PanelLayout {
	if o == nil {
		var ret []PanelLayout
		return ret
	}

	return o.PanelsLayouts
}

// GetPanelsLayoutsOk returns a tuple with the PanelsLayouts field value
// and a boolean to check if the value has been set.
func (o *SaveAnalyzerLayoutParams) GetPanelsLayoutsOk() ([]PanelLayout, bool) {
	if o == nil {
		return nil, false
	}
	return o.PanelsLayouts, true
}

// SetPanelsLayouts sets field value
func (o *SaveAnalyzerLayoutParams) SetPanelsLayouts(v []PanelLayout) {
	o.PanelsLayouts = v
}

// GetProjectId returns the ProjectId field value
func (o *SaveAnalyzerLayoutParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SaveAnalyzerLayoutParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SaveAnalyzerLayoutParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o SaveAnalyzerLayoutParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveAnalyzerLayoutParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["panelsLayouts"] = o.PanelsLayouts
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableSaveAnalyzerLayoutParams struct {
	value *SaveAnalyzerLayoutParams
	isSet bool
}

func (v NullableSaveAnalyzerLayoutParams) Get() *SaveAnalyzerLayoutParams {
	return v.value
}

func (v *NullableSaveAnalyzerLayoutParams) Set(val *SaveAnalyzerLayoutParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveAnalyzerLayoutParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveAnalyzerLayoutParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveAnalyzerLayoutParams(val *SaveAnalyzerLayoutParams) *NullableSaveAnalyzerLayoutParams {
	return &NullableSaveAnalyzerLayoutParams{value: val, isSet: true}
}

func (v NullableSaveAnalyzerLayoutParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveAnalyzerLayoutParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
