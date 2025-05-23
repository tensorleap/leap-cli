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

// checks if the TextViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TextViz{}

// TextViz struct for TextViz
type TextViz struct {
	Type     string    `json:"type"`
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Guid     string    `json:"guid"`
	Body     []string  `json:"body"`
	Clusters *Clusters `json:"clusters,omitempty"`
	Heatmap  *Heatmap  `json:"heatmap,omitempty"`
}

// NewTextViz instantiates a new TextViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTextViz(type_ string, title string, subTitle string, guid string, body []string) *TextViz {
	this := TextViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.Body = body
	return &this
}

// NewTextVizWithDefaults instantiates a new TextViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextVizWithDefaults() *TextViz {
	this := TextViz{}
	return &this
}

// GetType returns the Type field value
func (o *TextViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TextViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *TextViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *TextViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *TextViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *TextViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *TextViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *TextViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *TextViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *TextViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *TextViz) SetGuid(v string) {
	o.Guid = v
}

// GetBody returns the Body field value
func (o *TextViz) GetBody() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *TextViz) GetBodyOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *TextViz) SetBody(v []string) {
	o.Body = v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *TextViz) GetClusters() Clusters {
	if o == nil || IsNil(o.Clusters) {
		var ret Clusters
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextViz) GetClustersOk() (*Clusters, bool) {
	if o == nil || IsNil(o.Clusters) {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *TextViz) HasClusters() bool {
	if o != nil && !IsNil(o.Clusters) {
		return true
	}

	return false
}

// SetClusters gets a reference to the given Clusters and assigns it to the Clusters field.
func (o *TextViz) SetClusters(v Clusters) {
	o.Clusters = &v
}

// GetHeatmap returns the Heatmap field value if set, zero value otherwise.
func (o *TextViz) GetHeatmap() Heatmap {
	if o == nil || IsNil(o.Heatmap) {
		var ret Heatmap
		return ret
	}
	return *o.Heatmap
}

// GetHeatmapOk returns a tuple with the Heatmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextViz) GetHeatmapOk() (*Heatmap, bool) {
	if o == nil || IsNil(o.Heatmap) {
		return nil, false
	}
	return o.Heatmap, true
}

// HasHeatmap returns a boolean if a field has been set.
func (o *TextViz) HasHeatmap() bool {
	if o != nil && !IsNil(o.Heatmap) {
		return true
	}

	return false
}

// SetHeatmap gets a reference to the given Heatmap and assigns it to the Heatmap field.
func (o *TextViz) SetHeatmap(v Heatmap) {
	o.Heatmap = &v
}

func (o TextViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["body"] = o.Body
	if !IsNil(o.Clusters) {
		toSerialize["clusters"] = o.Clusters
	}
	if !IsNil(o.Heatmap) {
		toSerialize["heatmap"] = o.Heatmap
	}
	return toSerialize, nil
}

type NullableTextViz struct {
	value *TextViz
	isSet bool
}

func (v NullableTextViz) Get() *TextViz {
	return v.value
}

func (v *NullableTextViz) Set(val *TextViz) {
	v.value = val
	v.isSet = true
}

func (v NullableTextViz) IsSet() bool {
	return v.isSet
}

func (v *NullableTextViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextViz(val *TextViz) *NullableTextViz {
	return &NullableTextViz{value: val, isSet: true}
}

func (v NullableTextViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
