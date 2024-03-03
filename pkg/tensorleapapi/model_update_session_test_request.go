/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateSessionTestRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSessionTestRequest{}

// UpdateSessionTestRequest struct for UpdateSessionTestRequest
type UpdateSessionTestRequest struct {
	Cid           string              `json:"cid"`
	ProjectId     string              `json:"projectId"`
	Name          *string             `json:"name,omitempty"`
	TestFilter    *ClientFilterParams `json:"testFilter,omitempty"`
	DatasetFilter []ESFilter          `json:"datasetFilter,omitempty"`
}

// NewUpdateSessionTestRequest instantiates a new UpdateSessionTestRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSessionTestRequest(cid string, projectId string) *UpdateSessionTestRequest {
	this := UpdateSessionTestRequest{}
	this.Cid = cid
	this.ProjectId = projectId
	return &this
}

// NewUpdateSessionTestRequestWithDefaults instantiates a new UpdateSessionTestRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSessionTestRequestWithDefaults() *UpdateSessionTestRequest {
	this := UpdateSessionTestRequest{}
	return &this
}

// GetCid returns the Cid field value
func (o *UpdateSessionTestRequest) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionTestRequest) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *UpdateSessionTestRequest) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *UpdateSessionTestRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *UpdateSessionTestRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *UpdateSessionTestRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateSessionTestRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSessionTestRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateSessionTestRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateSessionTestRequest) SetName(v string) {
	o.Name = &v
}

// GetTestFilter returns the TestFilter field value if set, zero value otherwise.
func (o *UpdateSessionTestRequest) GetTestFilter() ClientFilterParams {
	if o == nil || IsNil(o.TestFilter) {
		var ret ClientFilterParams
		return ret
	}
	return *o.TestFilter
}

// GetTestFilterOk returns a tuple with the TestFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSessionTestRequest) GetTestFilterOk() (*ClientFilterParams, bool) {
	if o == nil || IsNil(o.TestFilter) {
		return nil, false
	}
	return o.TestFilter, true
}

// HasTestFilter returns a boolean if a field has been set.
func (o *UpdateSessionTestRequest) HasTestFilter() bool {
	if o != nil && !IsNil(o.TestFilter) {
		return true
	}

	return false
}

// SetTestFilter gets a reference to the given ClientFilterParams and assigns it to the TestFilter field.
func (o *UpdateSessionTestRequest) SetTestFilter(v ClientFilterParams) {
	o.TestFilter = &v
}

// GetDatasetFilter returns the DatasetFilter field value if set, zero value otherwise.
func (o *UpdateSessionTestRequest) GetDatasetFilter() []ESFilter {
	if o == nil || IsNil(o.DatasetFilter) {
		var ret []ESFilter
		return ret
	}
	return o.DatasetFilter
}

// GetDatasetFilterOk returns a tuple with the DatasetFilter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSessionTestRequest) GetDatasetFilterOk() ([]ESFilter, bool) {
	if o == nil || IsNil(o.DatasetFilter) {
		return nil, false
	}
	return o.DatasetFilter, true
}

// HasDatasetFilter returns a boolean if a field has been set.
func (o *UpdateSessionTestRequest) HasDatasetFilter() bool {
	if o != nil && !IsNil(o.DatasetFilter) {
		return true
	}

	return false
}

// SetDatasetFilter gets a reference to the given []ESFilter and assigns it to the DatasetFilter field.
func (o *UpdateSessionTestRequest) SetDatasetFilter(v []ESFilter) {
	o.DatasetFilter = v
}

func (o UpdateSessionTestRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSessionTestRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TestFilter) {
		toSerialize["testFilter"] = o.TestFilter
	}
	if !IsNil(o.DatasetFilter) {
		toSerialize["datasetFilter"] = o.DatasetFilter
	}
	return toSerialize, nil
}

type NullableUpdateSessionTestRequest struct {
	value *UpdateSessionTestRequest
	isSet bool
}

func (v NullableUpdateSessionTestRequest) Get() *UpdateSessionTestRequest {
	return v.value
}

func (v *NullableUpdateSessionTestRequest) Set(val *UpdateSessionTestRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSessionTestRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSessionTestRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSessionTestRequest(val *UpdateSessionTestRequest) *NullableUpdateSessionTestRequest {
	return &NullableUpdateSessionTestRequest{value: val, isSet: true}
}

func (v NullableUpdateSessionTestRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSessionTestRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
