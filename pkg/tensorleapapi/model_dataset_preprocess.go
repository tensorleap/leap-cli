/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetPreprocess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetPreprocess{}

// DatasetPreprocess struct for DatasetPreprocess
type DatasetPreprocess struct {
	TrainingLength   float64  `json:"training_length"`
	ValidationLength float64  `json:"validation_length"`
	TestLength       *float64 `json:"test_length,omitempty"`
	UnlabeledLength  *float64 `json:"unlabeled_length,omitempty"`
}

// NewDatasetPreprocess instantiates a new DatasetPreprocess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetPreprocess(trainingLength float64, validationLength float64) *DatasetPreprocess {
	this := DatasetPreprocess{}
	this.TrainingLength = trainingLength
	this.ValidationLength = validationLength
	return &this
}

// NewDatasetPreprocessWithDefaults instantiates a new DatasetPreprocess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetPreprocessWithDefaults() *DatasetPreprocess {
	this := DatasetPreprocess{}
	return &this
}

// GetTrainingLength returns the TrainingLength field value
func (o *DatasetPreprocess) GetTrainingLength() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.TrainingLength
}

// GetTrainingLengthOk returns a tuple with the TrainingLength field value
// and a boolean to check if the value has been set.
func (o *DatasetPreprocess) GetTrainingLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrainingLength, true
}

// SetTrainingLength sets field value
func (o *DatasetPreprocess) SetTrainingLength(v float64) {
	o.TrainingLength = v
}

// GetValidationLength returns the ValidationLength field value
func (o *DatasetPreprocess) GetValidationLength() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ValidationLength
}

// GetValidationLengthOk returns a tuple with the ValidationLength field value
// and a boolean to check if the value has been set.
func (o *DatasetPreprocess) GetValidationLengthOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidationLength, true
}

// SetValidationLength sets field value
func (o *DatasetPreprocess) SetValidationLength(v float64) {
	o.ValidationLength = v
}

// GetTestLength returns the TestLength field value if set, zero value otherwise.
func (o *DatasetPreprocess) GetTestLength() float64 {
	if o == nil || IsNil(o.TestLength) {
		var ret float64
		return ret
	}
	return *o.TestLength
}

// GetTestLengthOk returns a tuple with the TestLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetPreprocess) GetTestLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.TestLength) {
		return nil, false
	}
	return o.TestLength, true
}

// HasTestLength returns a boolean if a field has been set.
func (o *DatasetPreprocess) HasTestLength() bool {
	if o != nil && !IsNil(o.TestLength) {
		return true
	}

	return false
}

// SetTestLength gets a reference to the given float64 and assigns it to the TestLength field.
func (o *DatasetPreprocess) SetTestLength(v float64) {
	o.TestLength = &v
}

// GetUnlabeledLength returns the UnlabeledLength field value if set, zero value otherwise.
func (o *DatasetPreprocess) GetUnlabeledLength() float64 {
	if o == nil || IsNil(o.UnlabeledLength) {
		var ret float64
		return ret
	}
	return *o.UnlabeledLength
}

// GetUnlabeledLengthOk returns a tuple with the UnlabeledLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetPreprocess) GetUnlabeledLengthOk() (*float64, bool) {
	if o == nil || IsNil(o.UnlabeledLength) {
		return nil, false
	}
	return o.UnlabeledLength, true
}

// HasUnlabeledLength returns a boolean if a field has been set.
func (o *DatasetPreprocess) HasUnlabeledLength() bool {
	if o != nil && !IsNil(o.UnlabeledLength) {
		return true
	}

	return false
}

// SetUnlabeledLength gets a reference to the given float64 and assigns it to the UnlabeledLength field.
func (o *DatasetPreprocess) SetUnlabeledLength(v float64) {
	o.UnlabeledLength = &v
}

func (o DatasetPreprocess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetPreprocess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["training_length"] = o.TrainingLength
	toSerialize["validation_length"] = o.ValidationLength
	if !IsNil(o.TestLength) {
		toSerialize["test_length"] = o.TestLength
	}
	if !IsNil(o.UnlabeledLength) {
		toSerialize["unlabeled_length"] = o.UnlabeledLength
	}
	return toSerialize, nil
}

type NullableDatasetPreprocess struct {
	value *DatasetPreprocess
	isSet bool
}

func (v NullableDatasetPreprocess) Get() *DatasetPreprocess {
	return v.value
}

func (v *NullableDatasetPreprocess) Set(val *DatasetPreprocess) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetPreprocess) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetPreprocess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetPreprocess(val *DatasetPreprocess) *NullableDatasetPreprocess {
	return &NullableDatasetPreprocess{value: val, isSet: true}
}

func (v NullableDatasetPreprocess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetPreprocess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
