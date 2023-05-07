/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the VisualizationEntityResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VisualizationEntityResponse{}

// VisualizationEntityResponse struct for VisualizationEntityResponse
type VisualizationEntityResponse struct {
	Id string `json:"_id"`
	JobId string `json:"jobId"`
	JobParams *JobParams `json:"jobParams,omitempty"`
	Epoch float64 `json:"epoch"`
	BlobName string `json:"blobName"`
	CreatedAt time.Time `json:"createdAt"`
	Version string `json:"version"`
	SessionRunId string `json:"sessionRunId"`
	VisualizationUuid string `json:"visualizationUuid"`
	OrganizationId string `json:"organizationId"`
	Type AnalyzeTypeEnum `json:"type"`
	CreatedBy *string `json:"createdBy,omitempty"`
	Layout *SizedLayout `json:"layout,omitempty"`
}

// NewVisualizationEntityResponse instantiates a new VisualizationEntityResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVisualizationEntityResponse(id string, jobId string, epoch float64, blobName string, createdAt time.Time, version string, sessionRunId string, visualizationUuid string, organizationId string, type_ AnalyzeTypeEnum) *VisualizationEntityResponse {
	this := VisualizationEntityResponse{}
	this.Id = id
	this.JobId = jobId
	this.Epoch = epoch
	this.BlobName = blobName
	this.CreatedAt = createdAt
	this.Version = version
	this.SessionRunId = sessionRunId
	this.VisualizationUuid = visualizationUuid
	this.OrganizationId = organizationId
	this.Type = type_
	return &this
}

// NewVisualizationEntityResponseWithDefaults instantiates a new VisualizationEntityResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVisualizationEntityResponseWithDefaults() *VisualizationEntityResponse {
	this := VisualizationEntityResponse{}
	return &this
}

// GetId returns the Id field value
func (o *VisualizationEntityResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VisualizationEntityResponse) SetId(v string) {
	o.Id = v
}

// GetJobId returns the JobId field value
func (o *VisualizationEntityResponse) GetJobId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetJobIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JobId, true
}

// SetJobId sets field value
func (o *VisualizationEntityResponse) SetJobId(v string) {
	o.JobId = v
}

// GetJobParams returns the JobParams field value if set, zero value otherwise.
func (o *VisualizationEntityResponse) GetJobParams() JobParams {
	if o == nil || IsNil(o.JobParams) {
		var ret JobParams
		return ret
	}
	return *o.JobParams
}

// GetJobParamsOk returns a tuple with the JobParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetJobParamsOk() (*JobParams, bool) {
	if o == nil || IsNil(o.JobParams) {
		return nil, false
	}
	return o.JobParams, true
}

// HasJobParams returns a boolean if a field has been set.
func (o *VisualizationEntityResponse) HasJobParams() bool {
	if o != nil && !IsNil(o.JobParams) {
		return true
	}

	return false
}

// SetJobParams gets a reference to the given JobParams and assigns it to the JobParams field.
func (o *VisualizationEntityResponse) SetJobParams(v JobParams) {
	o.JobParams = &v
}

// GetEpoch returns the Epoch field value
func (o *VisualizationEntityResponse) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *VisualizationEntityResponse) SetEpoch(v float64) {
	o.Epoch = v
}

// GetBlobName returns the BlobName field value
func (o *VisualizationEntityResponse) GetBlobName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BlobName
}

// GetBlobNameOk returns a tuple with the BlobName field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetBlobNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlobName, true
}

// SetBlobName sets field value
func (o *VisualizationEntityResponse) SetBlobName(v string) {
	o.BlobName = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *VisualizationEntityResponse) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *VisualizationEntityResponse) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetVersion returns the Version field value
func (o *VisualizationEntityResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VisualizationEntityResponse) SetVersion(v string) {
	o.Version = v
}

// GetSessionRunId returns the SessionRunId field value
func (o *VisualizationEntityResponse) GetSessionRunId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetSessionRunIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionRunId, true
}

// SetSessionRunId sets field value
func (o *VisualizationEntityResponse) SetSessionRunId(v string) {
	o.SessionRunId = v
}

// GetVisualizationUuid returns the VisualizationUuid field value
func (o *VisualizationEntityResponse) GetVisualizationUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VisualizationUuid
}

// GetVisualizationUuidOk returns a tuple with the VisualizationUuid field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetVisualizationUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VisualizationUuid, true
}

// SetVisualizationUuid sets field value
func (o *VisualizationEntityResponse) SetVisualizationUuid(v string) {
	o.VisualizationUuid = v
}

// GetOrganizationId returns the OrganizationId field value
func (o *VisualizationEntityResponse) GetOrganizationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationId
}

// GetOrganizationIdOk returns a tuple with the OrganizationId field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetOrganizationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationId, true
}

// SetOrganizationId sets field value
func (o *VisualizationEntityResponse) SetOrganizationId(v string) {
	o.OrganizationId = v
}

// GetType returns the Type field value
func (o *VisualizationEntityResponse) GetType() AnalyzeTypeEnum {
	if o == nil {
		var ret AnalyzeTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetTypeOk() (*AnalyzeTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VisualizationEntityResponse) SetType(v AnalyzeTypeEnum) {
	o.Type = v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *VisualizationEntityResponse) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *VisualizationEntityResponse) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *VisualizationEntityResponse) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetLayout returns the Layout field value if set, zero value otherwise.
func (o *VisualizationEntityResponse) GetLayout() SizedLayout {
	if o == nil || IsNil(o.Layout) {
		var ret SizedLayout
		return ret
	}
	return *o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VisualizationEntityResponse) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil || IsNil(o.Layout) {
		return nil, false
	}
	return o.Layout, true
}

// HasLayout returns a boolean if a field has been set.
func (o *VisualizationEntityResponse) HasLayout() bool {
	if o != nil && !IsNil(o.Layout) {
		return true
	}

	return false
}

// SetLayout gets a reference to the given SizedLayout and assigns it to the Layout field.
func (o *VisualizationEntityResponse) SetLayout(v SizedLayout) {
	o.Layout = &v
}

func (o VisualizationEntityResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VisualizationEntityResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["_id"] = o.Id
	toSerialize["jobId"] = o.JobId
	if !IsNil(o.JobParams) {
		toSerialize["jobParams"] = o.JobParams
	}
	toSerialize["epoch"] = o.Epoch
	toSerialize["blobName"] = o.BlobName
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["version"] = o.Version
	toSerialize["sessionRunId"] = o.SessionRunId
	toSerialize["visualizationUuid"] = o.VisualizationUuid
	toSerialize["organizationId"] = o.OrganizationId
	toSerialize["type"] = o.Type
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Layout) {
		toSerialize["layout"] = o.Layout
	}
	return toSerialize, nil
}

type NullableVisualizationEntityResponse struct {
	value *VisualizationEntityResponse
	isSet bool
}

func (v NullableVisualizationEntityResponse) Get() *VisualizationEntityResponse {
	return v.value
}

func (v *NullableVisualizationEntityResponse) Set(val *VisualizationEntityResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableVisualizationEntityResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableVisualizationEntityResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisualizationEntityResponse(val *VisualizationEntityResponse) *NullableVisualizationEntityResponse {
	return &NullableVisualizationEntityResponse{value: val, isSet: true}
}

func (v NullableVisualizationEntityResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisualizationEntityResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


