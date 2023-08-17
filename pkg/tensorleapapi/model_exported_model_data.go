/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the ExportedModelData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportedModelData{}

// ExportedModelData struct for ExportedModelData
type ExportedModelData struct {
	Cid             string    `json:"cid"`
	JobId           string    `json:"jobId"`
	SessionWeightId string    `json:"sessionWeightId"`
	Blob            string    `json:"blob"`
	Status          string    `json:"status"`
	SessionId       string    `json:"sessionId"`
	Title           string    `json:"title"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	CreatedBy       string    `json:"createdBy"`
	FileType        string    `json:"fileType"`
}

// NewExportedModelData instantiates a new ExportedModelData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportedModelData(cid string, jobId string, sessionWeightId string, blob string, status string, sessionId string, title string, createdAt time.Time, updatedAt time.Time, createdBy string, fileType string) *ExportedModelData {
	this := ExportedModelData{}
	this.Cid = cid
	this.JobId = jobId
	this.SessionWeightId = sessionWeightId
	this.Blob = blob
	this.Status = status
	this.SessionId = sessionId
	this.Title = title
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.CreatedBy = createdBy
	this.FileType = fileType
	return &this
}

// NewExportedModelDataWithDefaults instantiates a new ExportedModelData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportedModelDataWithDefaults() *ExportedModelData {
	this := ExportedModelData{}
	return &this
}

// GetCid returns the Cid field value
func (o *ExportedModelData) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *ExportedModelData) SetCid(v string) {
	o.Cid = v
}

// GetJobId returns the JobId field value
func (o *ExportedModelData) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *ExportedModelData) SetJobId(v string) {
	o.JobId = v
}

// GetSessionWeightId returns the SessionWeightId field value
func (o *ExportedModelData) GetSessionWeightId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionWeightId
}

// GetSessionWeightIdOk returns a tuple with the SessionWeightId field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetSessionWeightIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionWeightId, true
}

// SetSessionWeightId sets field value
func (o *ExportedModelData) SetSessionWeightId(v string) {
	o.SessionWeightId = v
}

// GetBlob returns the Blob field value
func (o *ExportedModelData) GetBlob() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Blob
}

// GetBlobOk returns a tuple with the Blob field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetBlobOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blob, true
}

// SetBlob sets field value
func (o *ExportedModelData) SetBlob(v string) {
	o.Blob = v
}

// GetStatus returns the Status field value
func (o *ExportedModelData) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ExportedModelData) SetStatus(v string) {
	o.Status = v
}

// GetSessionId returns the SessionId field value
func (o *ExportedModelData) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ExportedModelData) SetSessionId(v string) {
	o.SessionId = v
}

// GetTitle returns the Title field value
func (o *ExportedModelData) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ExportedModelData) SetTitle(v string) {
	o.Title = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *ExportedModelData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *ExportedModelData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *ExportedModelData) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *ExportedModelData) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ExportedModelData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ExportedModelData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetFileType returns the FileType field value
func (o *ExportedModelData) GetFileType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileType
}

// GetFileTypeOk returns a tuple with the FileType field value
// and a boolean to check if the value has been set.
func (o *ExportedModelData) GetFileTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileType, true
}

// SetFileType sets field value
func (o *ExportedModelData) SetFileType(v string) {
	o.FileType = v
}

func (o ExportedModelData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportedModelData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["jobId"] = o.JobId
	toSerialize["sessionWeightId"] = o.SessionWeightId
	toSerialize["blob"] = o.Blob
	toSerialize["status"] = o.Status
	toSerialize["sessionId"] = o.SessionId
	toSerialize["title"] = o.Title
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["fileType"] = o.FileType
	return toSerialize, nil
}

type NullableExportedModelData struct {
	value *ExportedModelData
	isSet bool
}

func (v NullableExportedModelData) Get() *ExportedModelData {
	return v.value
}

func (v *NullableExportedModelData) Set(val *ExportedModelData) {
	v.value = val
	v.isSet = true
}

func (v NullableExportedModelData) IsSet() bool {
	return v.isSet
}

func (v *NullableExportedModelData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportedModelData(val *ExportedModelData) *NullableExportedModelData {
	return &NullableExportedModelData{value: val, isSet: true}
}

func (v NullableExportedModelData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportedModelData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
