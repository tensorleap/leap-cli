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

// checks if the GraphViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphViz{}

// GraphViz struct for GraphViz
type GraphViz struct {
	Type     string      `json:"type"`
	Title    string      `json:"title"`
	SubTitle string      `json:"sub_title"`
	Guid     string      `json:"guid"`
	Body     [][]float64 `json:"body"`
	Heatmap  *Heatmap    `json:"heatmap,omitempty"`
}

// NewGraphViz instantiates a new GraphViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphViz(type_ string, title string, subTitle string, guid string, body [][]float64) *GraphViz {
	this := GraphViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.Body = body
	return &this
}

// NewGraphVizWithDefaults instantiates a new GraphViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphVizWithDefaults() *GraphViz {
	this := GraphViz{}
	return &this
}

// GetType returns the Type field value
func (o *GraphViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GraphViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *GraphViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *GraphViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *GraphViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *GraphViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *GraphViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *GraphViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *GraphViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *GraphViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *GraphViz) SetGuid(v string) {
	o.Guid = v
}

// GetBody returns the Body field value
func (o *GraphViz) GetBody() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GraphViz) GetBodyOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *GraphViz) SetBody(v [][]float64) {
	o.Body = v
}

// GetHeatmap returns the Heatmap field value if set, zero value otherwise.
func (o *GraphViz) GetHeatmap() Heatmap {
	if o == nil || IsNil(o.Heatmap) {
		var ret Heatmap
		return ret
	}
	return *o.Heatmap
}

// GetHeatmapOk returns a tuple with the Heatmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphViz) GetHeatmapOk() (*Heatmap, bool) {
	if o == nil || IsNil(o.Heatmap) {
		return nil, false
	}
	return o.Heatmap, true
}

// HasHeatmap returns a boolean if a field has been set.
func (o *GraphViz) HasHeatmap() bool {
	if o != nil && !IsNil(o.Heatmap) {
		return true
	}

	return false
}

// SetHeatmap gets a reference to the given Heatmap and assigns it to the Heatmap field.
func (o *GraphViz) SetHeatmap(v Heatmap) {
	o.Heatmap = &v
}

func (o GraphViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["body"] = o.Body
	if !IsNil(o.Heatmap) {
		toSerialize["heatmap"] = o.Heatmap
	}
	return toSerialize, nil
}

type NullableGraphViz struct {
	value *GraphViz
	isSet bool
}

func (v NullableGraphViz) Get() *GraphViz {
	return v.value
}

func (v *NullableGraphViz) Set(val *GraphViz) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphViz) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphViz(val *GraphViz) *NullableGraphViz {
	return &NullableGraphViz{value: val, isSet: true}
}

func (v NullableGraphViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
