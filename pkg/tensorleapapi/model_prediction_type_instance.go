/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the PredictionTypeInstance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PredictionTypeInstance{}

// PredictionTypeInstance struct for PredictionTypeInstance
type PredictionTypeInstance struct {
	Name       string   `json:"name"`
	Labels     []string `json:"labels"`
	ChannelDim *float64 `json:"channel_dim,omitempty"`
}

// NewPredictionTypeInstance instantiates a new PredictionTypeInstance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPredictionTypeInstance(name string, labels []string) *PredictionTypeInstance {
	this := PredictionTypeInstance{}
	this.Name = name
	this.Labels = labels
	return &this
}

// NewPredictionTypeInstanceWithDefaults instantiates a new PredictionTypeInstance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPredictionTypeInstanceWithDefaults() *PredictionTypeInstance {
	this := PredictionTypeInstance{}
	return &this
}

// GetName returns the Name field value
func (o *PredictionTypeInstance) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PredictionTypeInstance) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PredictionTypeInstance) SetName(v string) {
	o.Name = v
}

// GetLabels returns the Labels field value
func (o *PredictionTypeInstance) GetLabels() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *PredictionTypeInstance) GetLabelsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Labels, true
}

// SetLabels sets field value
func (o *PredictionTypeInstance) SetLabels(v []string) {
	o.Labels = v
}

// GetChannelDim returns the ChannelDim field value if set, zero value otherwise.
func (o *PredictionTypeInstance) GetChannelDim() float64 {
	if o == nil || IsNil(o.ChannelDim) {
		var ret float64
		return ret
	}
	return *o.ChannelDim
}

// GetChannelDimOk returns a tuple with the ChannelDim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PredictionTypeInstance) GetChannelDimOk() (*float64, bool) {
	if o == nil || IsNil(o.ChannelDim) {
		return nil, false
	}
	return o.ChannelDim, true
}

// HasChannelDim returns a boolean if a field has been set.
func (o *PredictionTypeInstance) HasChannelDim() bool {
	if o != nil && !IsNil(o.ChannelDim) {
		return true
	}

	return false
}

// SetChannelDim gets a reference to the given float64 and assigns it to the ChannelDim field.
func (o *PredictionTypeInstance) SetChannelDim(v float64) {
	o.ChannelDim = &v
}

func (o PredictionTypeInstance) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PredictionTypeInstance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["labels"] = o.Labels
	if !IsNil(o.ChannelDim) {
		toSerialize["channel_dim"] = o.ChannelDim
	}
	return toSerialize, nil
}

type NullablePredictionTypeInstance struct {
	value *PredictionTypeInstance
	isSet bool
}

func (v NullablePredictionTypeInstance) Get() *PredictionTypeInstance {
	return v.value
}

func (v *NullablePredictionTypeInstance) Set(val *PredictionTypeInstance) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionTypeInstance) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionTypeInstance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionTypeInstance(val *PredictionTypeInstance) *NullablePredictionTypeInstance {
	return &NullablePredictionTypeInstance{value: val, isSet: true}
}

func (v NullablePredictionTypeInstance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionTypeInstance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
