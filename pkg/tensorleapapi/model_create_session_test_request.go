/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.285
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CreateSessionTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSessionTestRequest{}

// CreateSessionTestRequest struct for CreateSessionTestRequest
type CreateSessionTestRequest struct {
	ProjectId string `json:"projectId"`
	Name string `json:"name"`
	TestFilter ClientFilterParams `json:"testFilter"`
	DatasetFilter []ESFilter `json:"datasetFilter,omitempty"`
}

// NewCreateSessionTestRequest instantiates a new CreateSessionTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSessionTestRequest(projectId string, name string, testFilter ClientFilterParams) *CreateSessionTestRequest {
	this := CreateSessionTestRequest{}
	this.ProjectId = projectId
	this.Name = name
	this.TestFilter = testFilter
	return &this
}

// NewCreateSessionTestRequestWithDefaults instantiates a new CreateSessionTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSessionTestRequestWithDefaults() *CreateSessionTestRequest {
	this := CreateSessionTestRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *CreateSessionTestRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CreateSessionTestRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CreateSessionTestRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *CreateSessionTestRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSessionTestRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSessionTestRequest) SetName(v string) {
	o.Name = v
}

// GetTestFilter returns the TestFilter field value
func (o *CreateSessionTestRequest) GetTestFilter() ClientFilterParams {
	if o == nil {
		var ret ClientFilterParams
		return ret
	}

	return o.TestFilter
}

// GetTestFilterOk returns a tuple with the TestFilter field value
// and a boolean to check if the value has been set.
func (o *CreateSessionTestRequest) GetTestFilterOk() (*ClientFilterParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TestFilter, true
}

// SetTestFilter sets field value
func (o *CreateSessionTestRequest) SetTestFilter(v ClientFilterParams) {
	o.TestFilter = v
}

// GetDatasetFilter returns the DatasetFilter field value if set, zero value otherwise.
func (o *CreateSessionTestRequest) GetDatasetFilter() []ESFilter {
	if o == nil || IsNil(o.DatasetFilter) {
		var ret []ESFilter
		return ret
	}
	return o.DatasetFilter
}

// GetDatasetFilterOk returns a tuple with the DatasetFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSessionTestRequest) GetDatasetFilterOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.DatasetFilter) {
		return nil, false
	}
	return o.DatasetFilter, true
}

// HasDatasetFilter returns a boolean if a field has been set.
func (o *CreateSessionTestRequest) HasDatasetFilter() bool {
	if o != nil && !IsNil(o.DatasetFilter) {
		return true
	}

	return false
}

// SetDatasetFilter gets a reference to the given []ESFilter and assigns it to the DatasetFilter field.
func (o *CreateSessionTestRequest) SetDatasetFilter(v []ESFilter) {
	o.DatasetFilter = v
}

func (o CreateSessionTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSessionTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["testFilter"] = o.TestFilter
	if !IsNil(o.DatasetFilter) {
		toSerialize["datasetFilter"] = o.DatasetFilter
	}
	return toSerialize, nil
}

type NullableCreateSessionTestRequest struct {
	value *CreateSessionTestRequest
	isSet bool
}

func (v NullableCreateSessionTestRequest) Get() *CreateSessionTestRequest {
	return v.value
}

func (v *NullableCreateSessionTestRequest) Set(val *CreateSessionTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSessionTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSessionTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSessionTestRequest(val *CreateSessionTestRequest) *NullableCreateSessionTestRequest {
	return &NullableCreateSessionTestRequest{value: val, isSet: true}
}

func (v NullableCreateSessionTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSessionTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


