/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SaveDatasetVersionParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveDatasetVersionParams{}

// SaveDatasetVersionParams struct for SaveDatasetVersionParams
type SaveDatasetVersionParams struct {
	DatasetId       string        `json:"datasetId"`
	Setup           *DatasetSetup `json:"setup,omitempty"`
	SecretManagerId *string       `json:"secretManagerId,omitempty"`
	CodeUrl         string        `json:"codeUrl"`
	CodeEntryFile   string        `json:"codeEntryFile"`
	Branch          *string       `json:"branch,omitempty"`
	Note            *string       `json:"note,omitempty"`
}

// NewSaveDatasetVersionParams instantiates a new SaveDatasetVersionParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveDatasetVersionParams(datasetId string, codeUrl string, codeEntryFile string) *SaveDatasetVersionParams {
	this := SaveDatasetVersionParams{}
	this.DatasetId = datasetId
	this.CodeUrl = codeUrl
	this.CodeEntryFile = codeEntryFile
	return &this
}

// NewSaveDatasetVersionParamsWithDefaults instantiates a new SaveDatasetVersionParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveDatasetVersionParamsWithDefaults() *SaveDatasetVersionParams {
	this := SaveDatasetVersionParams{}
	return &this
}

// GetDatasetId returns the DatasetId field value
func (o *SaveDatasetVersionParams) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *SaveDatasetVersionParams) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *SaveDatasetVersionParams) GetSetup() DatasetSetup {
	if o == nil || IsNil(o.Setup) {
		var ret DatasetSetup
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetSetupOk() (*DatasetSetup, bool) {
	if o == nil || IsNil(o.Setup) {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *SaveDatasetVersionParams) HasSetup() bool {
	if o != nil && !IsNil(o.Setup) {
		return true
	}

	return false
}

// SetSetup gets a reference to the given DatasetSetup and assigns it to the Setup field.
func (o *SaveDatasetVersionParams) SetSetup(v DatasetSetup) {
	o.Setup = &v
}

// GetSecretManagerId returns the SecretManagerId field value if set, zero value otherwise.
func (o *SaveDatasetVersionParams) GetSecretManagerId() string {
	if o == nil || IsNil(o.SecretManagerId) {
		var ret string
		return ret
	}
	return *o.SecretManagerId
}

// GetSecretManagerIdOk returns a tuple with the SecretManagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetSecretManagerIdOk() (*string, bool) {
	if o == nil || IsNil(o.SecretManagerId) {
		return nil, false
	}
	return o.SecretManagerId, true
}

// HasSecretManagerId returns a boolean if a field has been set.
func (o *SaveDatasetVersionParams) HasSecretManagerId() bool {
	if o != nil && !IsNil(o.SecretManagerId) {
		return true
	}

	return false
}

// SetSecretManagerId gets a reference to the given string and assigns it to the SecretManagerId field.
func (o *SaveDatasetVersionParams) SetSecretManagerId(v string) {
	o.SecretManagerId = &v
}

// GetCodeUrl returns the CodeUrl field value
func (o *SaveDatasetVersionParams) GetCodeUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeUrl
}

// GetCodeUrlOk returns a tuple with the CodeUrl field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetCodeUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeUrl, true
}

// SetCodeUrl sets field value
func (o *SaveDatasetVersionParams) SetCodeUrl(v string) {
	o.CodeUrl = v
}

// GetCodeEntryFile returns the CodeEntryFile field value
func (o *SaveDatasetVersionParams) GetCodeEntryFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CodeEntryFile
}

// GetCodeEntryFileOk returns a tuple with the CodeEntryFile field value
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetCodeEntryFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CodeEntryFile, true
}

// SetCodeEntryFile sets field value
func (o *SaveDatasetVersionParams) SetCodeEntryFile(v string) {
	o.CodeEntryFile = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *SaveDatasetVersionParams) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *SaveDatasetVersionParams) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *SaveDatasetVersionParams) SetBranch(v string) {
	o.Branch = &v
}

// GetNote returns the Note field value if set, zero value otherwise.
func (o *SaveDatasetVersionParams) GetNote() string {
	if o == nil || IsNil(o.Note) {
		var ret string
		return ret
	}
	return *o.Note
}

// GetNoteOk returns a tuple with the Note field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveDatasetVersionParams) GetNoteOk() (*string, bool) {
	if o == nil || IsNil(o.Note) {
		return nil, false
	}
	return o.Note, true
}

// HasNote returns a boolean if a field has been set.
func (o *SaveDatasetVersionParams) HasNote() bool {
	if o != nil && !IsNil(o.Note) {
		return true
	}

	return false
}

// SetNote gets a reference to the given string and assigns it to the Note field.
func (o *SaveDatasetVersionParams) SetNote(v string) {
	o.Note = &v
}

func (o SaveDatasetVersionParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveDatasetVersionParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetId"] = o.DatasetId
	if !IsNil(o.Setup) {
		toSerialize["setup"] = o.Setup
	}
	if !IsNil(o.SecretManagerId) {
		toSerialize["secretManagerId"] = o.SecretManagerId
	}
	toSerialize["codeUrl"] = o.CodeUrl
	toSerialize["codeEntryFile"] = o.CodeEntryFile
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Note) {
		toSerialize["note"] = o.Note
	}
	return toSerialize, nil
}

type NullableSaveDatasetVersionParams struct {
	value *SaveDatasetVersionParams
	isSet bool
}

func (v NullableSaveDatasetVersionParams) Get() *SaveDatasetVersionParams {
	return v.value
}

func (v *NullableSaveDatasetVersionParams) Set(val *SaveDatasetVersionParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveDatasetVersionParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveDatasetVersionParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveDatasetVersionParams(val *SaveDatasetVersionParams) *NullableSaveDatasetVersionParams {
	return &NullableSaveDatasetVersionParams{value: val, isSet: true}
}

func (v NullableSaveDatasetVersionParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveDatasetVersionParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
