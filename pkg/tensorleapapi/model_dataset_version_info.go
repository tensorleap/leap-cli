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

// checks if the DatasetVersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetVersionInfo{}

// DatasetVersionInfo struct for DatasetVersionInfo
type DatasetVersionInfo struct {
	DatasetId                  string  `json:"datasetId"`
	DatasetName                string  `json:"datasetName"`
	DatasetVersionId           string  `json:"datasetVersionId"`
	DatasetVersionDisplayIndex float64 `json:"datasetVersionDisplayIndex"`
	DatasetBranchName          string  `json:"datasetBranchName"`
}

// NewDatasetVersionInfo instantiates a new DatasetVersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetVersionInfo(datasetId string, datasetName string, datasetVersionId string, datasetVersionDisplayIndex float64, datasetBranchName string) *DatasetVersionInfo {
	this := DatasetVersionInfo{}
	this.DatasetId = datasetId
	this.DatasetName = datasetName
	this.DatasetVersionId = datasetVersionId
	this.DatasetVersionDisplayIndex = datasetVersionDisplayIndex
	this.DatasetBranchName = datasetBranchName
	return &this
}

// NewDatasetVersionInfoWithDefaults instantiates a new DatasetVersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetVersionInfoWithDefaults() *DatasetVersionInfo {
	this := DatasetVersionInfo{}
	return &this
}

// GetDatasetId returns the DatasetId field value
func (o *DatasetVersionInfo) GetDatasetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionInfo) GetDatasetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetId, true
}

// SetDatasetId sets field value
func (o *DatasetVersionInfo) SetDatasetId(v string) {
	o.DatasetId = v
}

// GetDatasetName returns the DatasetName field value
func (o *DatasetVersionInfo) GetDatasetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetName
}

// GetDatasetNameOk returns a tuple with the DatasetName field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionInfo) GetDatasetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetName, true
}

// SetDatasetName sets field value
func (o *DatasetVersionInfo) SetDatasetName(v string) {
	o.DatasetName = v
}

// GetDatasetVersionId returns the DatasetVersionId field value
func (o *DatasetVersionInfo) GetDatasetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetVersionId
}

// GetDatasetVersionIdOk returns a tuple with the DatasetVersionId field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionInfo) GetDatasetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersionId, true
}

// SetDatasetVersionId sets field value
func (o *DatasetVersionInfo) SetDatasetVersionId(v string) {
	o.DatasetVersionId = v
}

// GetDatasetVersionDisplayIndex returns the DatasetVersionDisplayIndex field value
func (o *DatasetVersionInfo) GetDatasetVersionDisplayIndex() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.DatasetVersionDisplayIndex
}

// GetDatasetVersionDisplayIndexOk returns a tuple with the DatasetVersionDisplayIndex field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionInfo) GetDatasetVersionDisplayIndexOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetVersionDisplayIndex, true
}

// SetDatasetVersionDisplayIndex sets field value
func (o *DatasetVersionInfo) SetDatasetVersionDisplayIndex(v float64) {
	o.DatasetVersionDisplayIndex = v
}

// GetDatasetBranchName returns the DatasetBranchName field value
func (o *DatasetVersionInfo) GetDatasetBranchName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DatasetBranchName
}

// GetDatasetBranchNameOk returns a tuple with the DatasetBranchName field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionInfo) GetDatasetBranchNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DatasetBranchName, true
}

// SetDatasetBranchName sets field value
func (o *DatasetVersionInfo) SetDatasetBranchName(v string) {
	o.DatasetBranchName = v
}

func (o DatasetVersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetVersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["datasetId"] = o.DatasetId
	toSerialize["datasetName"] = o.DatasetName
	toSerialize["datasetVersionId"] = o.DatasetVersionId
	toSerialize["datasetVersionDisplayIndex"] = o.DatasetVersionDisplayIndex
	toSerialize["datasetBranchName"] = o.DatasetBranchName
	return toSerialize, nil
}

type NullableDatasetVersionInfo struct {
	value *DatasetVersionInfo
	isSet bool
}

func (v NullableDatasetVersionInfo) Get() *DatasetVersionInfo {
	return v.value
}

func (v *NullableDatasetVersionInfo) Set(val *DatasetVersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetVersionInfo(val *DatasetVersionInfo) *NullableDatasetVersionInfo {
	return &NullableDatasetVersionInfo{value: val, isSet: true}
}

func (v NullableDatasetVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
