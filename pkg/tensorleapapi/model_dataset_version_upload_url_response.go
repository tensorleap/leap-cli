/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DatasetVersionUploadUrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DatasetVersionUploadUrlResponse{}

// DatasetVersionUploadUrlResponse struct for DatasetVersionUploadUrlResponse
type DatasetVersionUploadUrlResponse struct {
	Url string `json:"url"`
}

// NewDatasetVersionUploadUrlResponse instantiates a new DatasetVersionUploadUrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDatasetVersionUploadUrlResponse(url string) *DatasetVersionUploadUrlResponse {
	this := DatasetVersionUploadUrlResponse{}
	this.Url = url
	return &this
}

// NewDatasetVersionUploadUrlResponseWithDefaults instantiates a new DatasetVersionUploadUrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDatasetVersionUploadUrlResponseWithDefaults() *DatasetVersionUploadUrlResponse {
	this := DatasetVersionUploadUrlResponse{}
	return &this
}

// GetUrl returns the Url field value
func (o *DatasetVersionUploadUrlResponse) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *DatasetVersionUploadUrlResponse) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *DatasetVersionUploadUrlResponse) SetUrl(v string) {
	o.Url = v
}

func (o DatasetVersionUploadUrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DatasetVersionUploadUrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["url"] = o.Url
	return toSerialize, nil
}

type NullableDatasetVersionUploadUrlResponse struct {
	value *DatasetVersionUploadUrlResponse
	isSet bool
}

func (v NullableDatasetVersionUploadUrlResponse) Get() *DatasetVersionUploadUrlResponse {
	return v.value
}

func (v *NullableDatasetVersionUploadUrlResponse) Set(val *DatasetVersionUploadUrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetVersionUploadUrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetVersionUploadUrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetVersionUploadUrlResponse(val *DatasetVersionUploadUrlResponse) *NullableDatasetVersionUploadUrlResponse {
	return &NullableDatasetVersionUploadUrlResponse{value: val, isSet: true}
}

func (v NullableDatasetVersionUploadUrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetVersionUploadUrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


