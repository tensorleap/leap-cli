/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SaveDatasetVersionNoParseParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveDatasetVersionNoParseParams{}

// SaveDatasetVersionNoParseParams struct for SaveDatasetVersionNoParseParams
type SaveDatasetVersionNoParseParams struct {
	FromDatasetVersionId string  `json:"fromDatasetVersionId"`
	SecretManagerId      *string `json:"secretManagerId,omitempty"`
	CodeUrl              string  `json:"codeUrl"`
	CodeEntryFile        string  `json:"codeEntryFile"`
	Note                 *string `json:"note,omitempty"`
}

// NewSaveDatasetVersionNoParseParams instantiates a new SaveDatasetVersionNoParseParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveDatasetVersionNoParseParams(fromDatasetVersionId string, codeUrl string, codeEntryFile string) *SaveDatasetVersionNoParseParams {
	this := SaveDatasetVersionNoParseParams{}
	this.FromDatasetVersionId = fromDatasetVersionId
	this.CodeUrl = codeUrl
	this.CodeEntryFile = codeEntryFile
	return &this
}

// NewSaveDatasetVersionNoParseParamsWithDefaults instantiates a new SaveDatasetVersionNoParseParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveDatasetVersionNoParseParamsWithDefaults() *SaveDatasetVersionNoParseParams {
	this := SaveDatasetVersionNoParseParams{}
	return &this
}

// GetFromDatasetVersionId returns the FromDatasetVersionId field value
func (o *SaveDatasetVersionNoParseParams) GetFromDatasetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromDatasetVersionId
}

// GetFromDatasetVersionIdOk returns a tuple with the FromDatasetVersionId field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionNoParseParams) GetFromDatasetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromDatasetVersionId, true
}

// SetFromDatasetVersionId sets field value
func (o *SaveDatasetVersionNoParseParams) SetFromDatasetVersionId(v string) {
	o.FromDatasetVersionId = v
}

// GetSecretManagerId returns the SecretManagerId field value if set, zero value otherwise.
func (o *SaveDatasetVersionNoParseParams) GetSecretManagerId() string {
	if o == nil || IsNil(o.SecretManagerId) {
		var ret string
		return ret
	}
	return *o.SecretManagerId
}

// GetSecretManagerIdOk returns a tuple with the SecretManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionNoParseParams) GetSecretManagerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretManagerId) {
		return nil, false
	}
	return o.SecretManagerId, true
}

// HasSecretManagerId returns a boolean if a field has been set.
func (o *SaveDatasetVersionNoParseParams) HasSecretManagerId() bool {
	if o != nil && !IsNil(o.SecretManagerId) {
		return true
	}

	return false
}

// SetSecretManagerId gets a reference to the given string and assigns it to the SecretManagerId field.
func (o *SaveDatasetVersionNoParseParams) SetSecretManagerId(v string) {
	o.SecretManagerId = &v
}

// GetCodeUrl returns the CodeUrl field value
func (o *SaveDatasetVersionNoParseParams) GetCodeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeUrl
}

// GetCodeUrlOk returns a tuple with the CodeUrl field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionNoParseParams) GetCodeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeUrl, true
}

// SetCodeUrl sets field value
func (o *SaveDatasetVersionNoParseParams) SetCodeUrl(v string) {
	o.CodeUrl = v
}

// GetCodeEntryFile returns the CodeEntryFile field value
func (o *SaveDatasetVersionNoParseParams) GetCodeEntryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeEntryFile
}

// GetCodeEntryFileOk returns a tuple with the CodeEntryFile field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionNoParseParams) GetCodeEntryFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeEntryFile, true
}

// SetCodeEntryFile sets field value
func (o *SaveDatasetVersionNoParseParams) SetCodeEntryFile(v string) {
	o.CodeEntryFile = v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *SaveDatasetVersionNoParseParams) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionNoParseParams) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *SaveDatasetVersionNoParseParams) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *SaveDatasetVersionNoParseParams) SetNote(v string) {
	o.Note = &v
}

func (o SaveDatasetVersionNoParseParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveDatasetVersionNoParseParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fromDatasetVersionId"] = o.FromDatasetVersionId
	if !IsNil(o.SecretManagerId) {
		toSerialize["secretManagerId"] = o.SecretManagerId
	}
	toSerialize["codeUrl"] = o.CodeUrl
	toSerialize["codeEntryFile"] = o.CodeEntryFile
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullableSaveDatasetVersionNoParseParams struct {
	value *SaveDatasetVersionNoParseParams
	isSet bool
}

func (v NullableSaveDatasetVersionNoParseParams) Get() *SaveDatasetVersionNoParseParams {
	return v.value
}

func (v *NullableSaveDatasetVersionNoParseParams) Set(val *SaveDatasetVersionNoParseParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveDatasetVersionNoParseParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveDatasetVersionNoParseParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveDatasetVersionNoParseParams(val *SaveDatasetVersionNoParseParams) *NullableSaveDatasetVersionNoParseParams {
	return &NullableSaveDatasetVersionNoParseParams{value: val, isSet: true}
}

func (v NullableSaveDatasetVersionNoParseParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveDatasetVersionNoParseParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
