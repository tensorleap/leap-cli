/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleAnalysisViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleAnalysisViz{}

// SampleAnalysisViz struct for SampleAnalysisViz
type SampleAnalysisViz struct {
	Type            string           `json:"type"`
	Title           string           `json:"title"`
	SubTitle        string           `json:"sub_title"`
	Guid            string           `json:"guid"`
	VisualizedItems []VisualizedItem `json:"visualized_items"`
	// Construct a type with a set of properties K of type T
	MetadataMap map[string]interface{} `json:"metadata_map"`
}

// NewSampleAnalysisViz instantiates a new SampleAnalysisViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleAnalysisViz(type_ string, title string, subTitle string, guid string, visualizedItems []VisualizedItem, metadataMap map[string]interface{}) *SampleAnalysisViz {
	this := SampleAnalysisViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.VisualizedItems = visualizedItems
	this.MetadataMap = metadataMap
	return &this
}

// NewSampleAnalysisVizWithDefaults instantiates a new SampleAnalysisViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleAnalysisVizWithDefaults() *SampleAnalysisViz {
	this := SampleAnalysisViz{}
	return &this
}

// GetType returns the Type field value
func (o *SampleAnalysisViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SampleAnalysisViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *SampleAnalysisViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *SampleAnalysisViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *SampleAnalysisViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *SampleAnalysisViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *SampleAnalysisViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *SampleAnalysisViz) SetGuid(v string) {
	o.Guid = v
}

// GetVisualizedItems returns the VisualizedItems field value
func (o *SampleAnalysisViz) GetVisualizedItems() []VisualizedItem {
	if o == nil {
		var ret []VisualizedItem
		return ret
	}

	return o.VisualizedItems
}

// GetVisualizedItemsOk returns a tuple with the VisualizedItems field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetVisualizedItemsOk() ([]VisualizedItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisualizedItems, true
}

// SetVisualizedItems sets field value
func (o *SampleAnalysisViz) SetVisualizedItems(v []VisualizedItem) {
	o.VisualizedItems = v
}

// GetMetadataMap returns the MetadataMap field value
func (o *SampleAnalysisViz) GetMetadataMap() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.MetadataMap
}

// GetMetadataMapOk returns a tuple with the MetadataMap field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisViz) GetMetadataMapOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.MetadataMap, true
}

// SetMetadataMap sets field value
func (o *SampleAnalysisViz) SetMetadataMap(v map[string]interface{}) {
	o.MetadataMap = v
}

func (o SampleAnalysisViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleAnalysisViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["visualized_items"] = o.VisualizedItems
	toSerialize["metadata_map"] = o.MetadataMap
	return toSerialize, nil
}

type NullableSampleAnalysisViz struct {
	value *SampleAnalysisViz
	isSet bool
}

func (v NullableSampleAnalysisViz) Get() *SampleAnalysisViz {
	return v.value
}

func (v *NullableSampleAnalysisViz) Set(val *SampleAnalysisViz) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleAnalysisViz) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleAnalysisViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleAnalysisViz(val *SampleAnalysisViz) *NullableSampleAnalysisViz {
	return &NullableSampleAnalysisViz{value: val, isSet: true}
}

func (v NullableSampleAnalysisViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleAnalysisViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
