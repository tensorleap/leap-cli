/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ImageViz type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ImageViz{}

// ImageViz struct for ImageViz
type ImageViz struct {
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	SubTitle string   `json:"sub_title"`
	Guid     string   `json:"guid"`
	Blob     string   `json:"blob"`
	Labels   []string `json:"labels,omitempty"`
}

// NewImageViz instantiates a new ImageViz object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImageViz(type_ string, title string, subTitle string, guid string, blob string) *ImageViz {
	this := ImageViz{}
	this.Type = type_
	this.Title = title
	this.SubTitle = subTitle
	this.Guid = guid
	this.Blob = blob
	return &this
}

// NewImageVizWithDefaults instantiates a new ImageViz object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImageVizWithDefaults() *ImageViz {
	this := ImageViz{}
	return &this
}

// GetType returns the Type field value
func (o *ImageViz) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ImageViz) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ImageViz) SetType(v string) {
	o.Type = v
}

// GetTitle returns the Title field value
func (o *ImageViz) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ImageViz) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ImageViz) SetTitle(v string) {
	o.Title = v
}

// GetSubTitle returns the SubTitle field value
func (o *ImageViz) GetSubTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubTitle
}

// GetSubTitleOk returns a tuple with the SubTitle field value
// and a boolean to check if the value has been set.
func (o *ImageViz) GetSubTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubTitle, true
}

// SetSubTitle sets field value
func (o *ImageViz) SetSubTitle(v string) {
	o.SubTitle = v
}

// GetGuid returns the Guid field value
func (o *ImageViz) GetGuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Guid
}

// GetGuidOk returns a tuple with the Guid field value
// and a boolean to check if the value has been set.
func (o *ImageViz) GetGuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Guid, true
}

// SetGuid sets field value
func (o *ImageViz) SetGuid(v string) {
	o.Guid = v
}

// GetBlob returns the Blob field value
func (o *ImageViz) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *ImageViz) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *ImageViz) SetBlob(v string) {
	o.Blob = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ImageViz) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImageViz) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ImageViz) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *ImageViz) SetLabels(v []string) {
	o.Labels = v
}

func (o ImageViz) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ImageViz) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["title"] = o.Title
	toSerialize["sub_title"] = o.SubTitle
	toSerialize["guid"] = o.Guid
	toSerialize["blob"] = o.Blob
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	return toSerialize, nil
}

type NullableImageViz struct {
	value *ImageViz
	isSet bool
}

func (v NullableImageViz) Get() *ImageViz {
	return v.value
}

func (v *NullableImageViz) Set(val *ImageViz) {
	v.value = val
	v.isSet = true
}

func (v NullableImageViz) IsSet() bool {
	return v.isSet
}

func (v *NullableImageViz) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImageViz(val *ImageViz) *NullableImageViz {
	return &NullableImageViz{value: val, isSet: true}
}

func (v NullableImageViz) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImageViz) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
