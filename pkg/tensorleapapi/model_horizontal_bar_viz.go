/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the HorizontalBarViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HorizontalBarViz{}

// HorizontalBarViz struct for HorizontalBarViz
type HorizontalBarViz struct {
	Type     string    `json:"type"`
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Guid     string    `json:"guid"`
	Body     []float64 `json:"body"`
	Labels   []string  `json:"labels"`
}

// NewHorizontalBarViz instantiates a new HorizontalBarViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHorizontalBarViz(type_ string, title string, subTitle string, guid string, body []float64, labels []string) *HorizontalBarViz {
	this := HorizontalBarViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.Body = body
	this.Labels = labels
	return &this
}

// NewHorizontalBarVizWithDefaults instantiates a new HorizontalBarViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHorizontalBarVizWithDefaults() *HorizontalBarViz {
	this := HorizontalBarViz{}
	return &this
}

// GetType returns the Type field value
func (o *HorizontalBarViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HorizontalBarViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *HorizontalBarViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *HorizontalBarViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *HorizontalBarViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *HorizontalBarViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *HorizontalBarViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *HorizontalBarViz) SetGuid(v string) {
	o.Guid = v
}

// GetBody returns the Body field value
func (o *HorizontalBarViz) GetBody() []float64 {
	if o == nil {
		var ret []float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetBodyOk() ([]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *HorizontalBarViz) SetBody(v []float64) {
	o.Body = v
}

// GetLabels returns the Labels field value
func (o *HorizontalBarViz) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *HorizontalBarViz) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *HorizontalBarViz) SetLabels(v []string) {
	o.Labels = v
}

func (o HorizontalBarViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HorizontalBarViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["body"] = o.Body
	toSerialize["labels"] = o.Labels
	return toSerialize, nil
}

type NullableHorizontalBarViz struct {
	value *HorizontalBarViz
	isSet bool
}

func (v NullableHorizontalBarViz) Get() *HorizontalBarViz {
	return v.value
}

func (v *NullableHorizontalBarViz) Set(val *HorizontalBarViz) {
	v.value = val
	v.isSet = true
}

func (v NullableHorizontalBarViz) IsSet() bool {
	return v.isSet
}

func (v *NullableHorizontalBarViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHorizontalBarViz(val *HorizontalBarViz) *NullableHorizontalBarViz {
	return &NullableHorizontalBarViz{value: val, isSet: true}
}

func (v NullableHorizontalBarViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHorizontalBarViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
