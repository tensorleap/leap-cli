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

// checks if the ScatterViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScatterViz{}

// ScatterViz struct for ScatterViz
type ScatterViz struct {
	Type        string              `json:"type"`
	Title       string              `json:"title"`
	SubTitle    string              `json:"sub_title"`
	Guid        string              `json:"guid"`
	ScatterData ScatterVizDataState `json:"scatter_data"`
}

// NewScatterViz instantiates a new ScatterViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScatterViz(type_ string, title string, subTitle string, guid string, scatterData ScatterVizDataState) *ScatterViz {
	this := ScatterViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.ScatterData = scatterData
	return &this
}

// NewScatterVizWithDefaults instantiates a new ScatterViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScatterVizWithDefaults() *ScatterViz {
	this := ScatterViz{}
	return &this
}

// GetType returns the Type field value
func (o *ScatterViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScatterViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScatterViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ScatterViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ScatterViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ScatterViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *ScatterViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *ScatterViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *ScatterViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *ScatterViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *ScatterViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *ScatterViz) SetGuid(v string) {
	o.Guid = v
}

// GetScatterData returns the ScatterData field value
func (o *ScatterViz) GetScatterData() ScatterVizDataState {
	if o == nil {
		var ret ScatterVizDataState
		return ret
	}

	return o.ScatterData
}

// GetScatterDataOk returns a tuple with the ScatterData field value
// and a boolean to check if the value has been set.
func (o *ScatterViz) GetScatterDataOk() (*ScatterVizDataState, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScatterData, true
}

// SetScatterData sets field value
func (o *ScatterViz) SetScatterData(v ScatterVizDataState) {
	o.ScatterData = v
}

func (o ScatterViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScatterViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["scatter_data"] = o.ScatterData
	return toSerialize, nil
}

type NullableScatterViz struct {
	value *ScatterViz
	isSet bool
}

func (v NullableScatterViz) Get() *ScatterViz {
	return v.value
}

func (v *NullableScatterViz) Set(val *ScatterViz) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterViz) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterViz(val *ScatterViz) *NullableScatterViz {
	return &NullableScatterViz{value: val, isSet: true}
}

func (v NullableScatterViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
