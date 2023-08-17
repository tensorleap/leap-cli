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

// checks if the DatasetMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetMetadata{}

// DatasetMetadata struct for DatasetMetadata
type DatasetMetadata struct {
	Setup           *DatasetSetup       `json:"setup,omitempty"`
	SetupStatus     *DatasetSetupStatus `json:"setupStatus,omitempty"`
	ModelSetup      *ModelSetup         `json:"modelSetup,omitempty"`
	SecretManagerId *string             `json:"secretManagerId,omitempty"`
}

// NewDatasetMetadata instantiates a new DatasetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetMetadata() *DatasetMetadata {
	this := DatasetMetadata{}
	return &this
}

// NewDatasetMetadataWithDefaults instantiates a new DatasetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetMetadataWithDefaults() *DatasetMetadata {
	this := DatasetMetadata{}
	return &this
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *DatasetMetadata) GetSetup() DatasetSetup {
	if o == nil || IsNil(o.Setup) {
		var ret DatasetSetup
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetMetadata) GetSetupOk() (*DatasetSetup, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *DatasetMetadata) HasSetup() bool {
	if o != nil && !IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given DatasetSetup and assigns it to the Setup field.
func (o *DatasetMetadata) SetSetup(v DatasetSetup) {
	o.Setup = &v
}

// GetSetupStatus returns the SetupStatus field value if set, zero value otherwise.
func (o *DatasetMetadata) GetSetupStatus() DatasetSetupStatus {
	if o == nil || IsNil(o.SetupStatus) {
		var ret DatasetSetupStatus
		return ret
	}
	return *o.SetupStatus
}

// GetSetupStatusOk returns a tuple with the SetupStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetMetadata) GetSetupStatusOk() (*DatasetSetupStatus, bool) {
	if o == nil || IsNil(o.SetupStatus) {
		return nil, false
	}
	return o.SetupStatus, true
}

// HasSetupStatus returns a boolean if a field has been set.
func (o *DatasetMetadata) HasSetupStatus() bool {
	if o != nil && !IsNil(o.SetupStatus) {
		return true
	}

	return false
}

// SetSetupStatus gets a reference to the given DatasetSetupStatus and assigns it to the SetupStatus field.
func (o *DatasetMetadata) SetSetupStatus(v DatasetSetupStatus) {
	o.SetupStatus = &v
}

// GetModelSetup returns the ModelSetup field value if set, zero value otherwise.
func (o *DatasetMetadata) GetModelSetup() ModelSetup {
	if o == nil || IsNil(o.ModelSetup) {
		var ret ModelSetup
		return ret
	}
	return *o.ModelSetup
}

// GetModelSetupOk returns a tuple with the ModelSetup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetMetadata) GetModelSetupOk() (*ModelSetup, bool) {
	if o == nil || IsNil(o.ModelSetup) {
		return nil, false
	}
	return o.ModelSetup, true
}

// HasModelSetup returns a boolean if a field has been set.
func (o *DatasetMetadata) HasModelSetup() bool {
	if o != nil && !IsNil(o.ModelSetup) {
		return true
	}

	return false
}

// SetModelSetup gets a reference to the given ModelSetup and assigns it to the ModelSetup field.
func (o *DatasetMetadata) SetModelSetup(v ModelSetup) {
	o.ModelSetup = &v
}

// GetSecretManagerId returns the SecretManagerId field value if set, zero value otherwise.
func (o *DatasetMetadata) GetSecretManagerId() string {
	if o == nil || IsNil(o.SecretManagerId) {
		var ret string
		return ret
	}
	return *o.SecretManagerId
}

// GetSecretManagerIdOk returns a tuple with the SecretManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DatasetMetadata) GetSecretManagerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretManagerId) {
		return nil, false
	}
	return o.SecretManagerId, true
}

// HasSecretManagerId returns a boolean if a field has been set.
func (o *DatasetMetadata) HasSecretManagerId() bool {
	if o != nil && !IsNil(o.SecretManagerId) {
		return true
	}

	return false
}

// SetSecretManagerId gets a reference to the given string and assigns it to the SecretManagerId field.
func (o *DatasetMetadata) SetSecretManagerId(v string) {
	o.SecretManagerId = &v
}

func (o DatasetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Setup) {
		toSerialize["setup"] = o.Setup
	}
	if !IsNil(o.SetupStatus) {
		toSerialize["setupStatus"] = o.SetupStatus
	}
	if !IsNil(o.ModelSetup) {
		toSerialize["modelSetup"] = o.ModelSetup
	}
	if !IsNil(o.SecretManagerId) {
		toSerialize["secretManagerId"] = o.SecretManagerId
	}
	return toSerialize, nil
}

type NullableDatasetMetadata struct {
	value *DatasetMetadata
	isSet bool
}

func (v NullableDatasetMetadata) Get() *DatasetMetadata {
	return v.value
}

func (v *NullableDatasetMetadata) Set(val *DatasetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetMetadata(val *DatasetMetadata) *NullableDatasetMetadata {
	return &NullableDatasetMetadata{value: val, isSet: true}
}

func (v NullableDatasetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
