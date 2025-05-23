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

// checks if the BBoxImageViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BBoxImageViz{}

// BBoxImageViz struct for BBoxImageViz
type BBoxImageViz struct {
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	SubTitle      string        `json:"sub_title"`
	Guid          string        `json:"guid"`
	BoundingBoxes []BoundingBox `json:"bounding_boxes"`
	Labels        []string      `json:"labels,omitempty"`
}

// NewBBoxImageViz instantiates a new BBoxImageViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBBoxImageViz(type_ string, title string, subTitle string, guid string, boundingBoxes []BoundingBox) *BBoxImageViz {
	this := BBoxImageViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.BoundingBoxes = boundingBoxes
	return &this
}

// NewBBoxImageVizWithDefaults instantiates a new BBoxImageViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBBoxImageVizWithDefaults() *BBoxImageViz {
	this := BBoxImageViz{}
	return &this
}

// GetType returns the Type field value
func (o *BBoxImageViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BBoxImageViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *BBoxImageViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *BBoxImageViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *BBoxImageViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *BBoxImageViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *BBoxImageViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *BBoxImageViz) SetGuid(v string) {
	o.Guid = v
}

// GetBoundingBoxes returns the BoundingBoxes field value
func (o *BBoxImageViz) GetBoundingBoxes() []BoundingBox {
	if o == nil {
		var ret []BoundingBox
		return ret
	}

	return o.BoundingBoxes
}

// GetBoundingBoxesOk returns a tuple with the BoundingBoxes field value
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetBoundingBoxesOk() ([]BoundingBox, bool) {
	if o == nil {
		return nil, false
	}
	return o.BoundingBoxes, true
}

// SetBoundingBoxes sets field value
func (o *BBoxImageViz) SetBoundingBoxes(v []BoundingBox) {
	o.BoundingBoxes = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *BBoxImageViz) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BBoxImageViz) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *BBoxImageViz) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *BBoxImageViz) SetLabels(v []string) {
	o.Labels = v
}

func (o BBoxImageViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BBoxImageViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["bounding_boxes"] = o.BoundingBoxes
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableBBoxImageViz struct {
	value *BBoxImageViz
	isSet bool
}

func (v NullableBBoxImageViz) Get() *BBoxImageViz {
	return v.value
}

func (v *NullableBBoxImageViz) Set(val *BBoxImageViz) {
	v.value = val
	v.isSet = true
}

func (v NullableBBoxImageViz) IsSet() bool {
	return v.isSet
}

func (v *NullableBBoxImageViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBBoxImageViz(val *BBoxImageViz) *NullableBBoxImageViz {
	return &NullableBBoxImageViz{value: val, isSet: true}
}

func (v NullableBBoxImageViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBBoxImageViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
