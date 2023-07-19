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

// checks if the FeatureImportanceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FeatureImportanceItem{}

// FeatureImportanceItem struct for FeatureImportanceItem
type FeatureImportanceItem struct {
	Label string `json:"label"`
	Fi VisData `json:"fi"`
	GroundTruthName *string `json:"ground_truth_name,omitempty"`
}

// NewFeatureImportanceItem instantiates a new FeatureImportanceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFeatureImportanceItem(label string, fi VisData) *FeatureImportanceItem {
	this := FeatureImportanceItem{}
	this.Label = label
	this.Fi = fi
	return &this
}

// NewFeatureImportanceItemWithDefaults instantiates a new FeatureImportanceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFeatureImportanceItemWithDefaults() *FeatureImportanceItem {
	this := FeatureImportanceItem{}
	return &this
}

// GetLabel returns the Label field value
func (o *FeatureImportanceItem) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *FeatureImportanceItem) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *FeatureImportanceItem) SetLabel(v string) {
	o.Label = v
}

// GetFi returns the Fi field value
func (o *FeatureImportanceItem) GetFi() VisData {
	if o == nil {
		var ret VisData
		return ret
	}

	return o.Fi
}

// GetFiOk returns a tuple with the Fi field value
// and a boolean to check if the value has been set.
func (o *FeatureImportanceItem) GetFiOk() (*VisData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fi, true
}

// SetFi sets field value
func (o *FeatureImportanceItem) SetFi(v VisData) {
	o.Fi = v
}

// GetGroundTruthName returns the GroundTruthName field value if set, zero value otherwise.
func (o *FeatureImportanceItem) GetGroundTruthName() string {
	if o == nil || IsNil(o.GroundTruthName) {
		var ret string
		return ret
	}
	return *o.GroundTruthName
}

// GetGroundTruthNameOk returns a tuple with the GroundTruthName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FeatureImportanceItem) GetGroundTruthNameOk() (*string, bool) {
	if o == nil || IsNil(o.GroundTruthName) {
		return nil, false
	}
	return o.GroundTruthName, true
}

// HasGroundTruthName returns a boolean if a field has been set.
func (o *FeatureImportanceItem) HasGroundTruthName() bool {
	if o != nil && !IsNil(o.GroundTruthName) {
		return true
	}

	return false
}

// SetGroundTruthName gets a reference to the given string and assigns it to the GroundTruthName field.
func (o *FeatureImportanceItem) SetGroundTruthName(v string) {
	o.GroundTruthName = &v
}

func (o FeatureImportanceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FeatureImportanceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["label"] = o.Label
	toSerialize["fi"] = o.Fi
	if !IsNil(o.GroundTruthName) {
		toSerialize["ground_truth_name"] = o.GroundTruthName
	}
	return toSerialize, nil
}

type NullableFeatureImportanceItem struct {
	value *FeatureImportanceItem
	isSet bool
}

func (v NullableFeatureImportanceItem) Get() *FeatureImportanceItem {
	return v.value
}

func (v *NullableFeatureImportanceItem) Set(val *FeatureImportanceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableFeatureImportanceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableFeatureImportanceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFeatureImportanceItem(val *FeatureImportanceItem) *NullableFeatureImportanceItem {
	return &NullableFeatureImportanceItem{value: val, isSet: true}
}

func (v NullableFeatureImportanceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFeatureImportanceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


