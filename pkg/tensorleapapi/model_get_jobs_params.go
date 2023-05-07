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

// checks if the GetJobsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetJobsParams{}

// GetJobsParams struct for GetJobsParams
type GetJobsParams struct {
	BranchId *string `json:"branchId,omitempty"`
	SessionRunId *string `json:"sessionRunId,omitempty"`
	LastUpdated *string `json:"lastUpdated,omitempty"`
}

// NewGetJobsParams instantiates a new GetJobsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetJobsParams() *GetJobsParams {
	this := GetJobsParams{}
	return &this
}

// NewGetJobsParamsWithDefaults instantiates a new GetJobsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetJobsParamsWithDefaults() *GetJobsParams {
	this := GetJobsParams{}
	return &this
}

// GetBranchId returns the BranchId field value if set, zero value otherwise.
func (o *GetJobsParams) GetBranchId() string {
	if o == nil || IsNil(o.BranchId) {
		var ret string
		return ret
	}
	return *o.BranchId
}

// GetBranchIdOk returns a tuple with the BranchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsParams) GetBranchIdOk() (*string, bool) {
	if o == nil || IsNil(o.BranchId) {
		return nil, false
	}
	return o.BranchId, true
}

// HasBranchId returns a boolean if a field has been set.
func (o *GetJobsParams) HasBranchId() bool {
	if o != nil && !IsNil(o.BranchId) {
		return true
	}

	return false
}

// SetBranchId gets a reference to the given string and assigns it to the BranchId field.
func (o *GetJobsParams) SetBranchId(v string) {
	o.BranchId = &v
}

// GetSessionRunId returns the SessionRunId field value if set, zero value otherwise.
func (o *GetJobsParams) GetSessionRunId() string {
	if o == nil || IsNil(o.SessionRunId) {
		var ret string
		return ret
	}
	return *o.SessionRunId
}

// GetSessionRunIdOk returns a tuple with the SessionRunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsParams) GetSessionRunIdOk() (*string, bool) {
	if o == nil || IsNil(o.SessionRunId) {
		return nil, false
	}
	return o.SessionRunId, true
}

// HasSessionRunId returns a boolean if a field has been set.
func (o *GetJobsParams) HasSessionRunId() bool {
	if o != nil && !IsNil(o.SessionRunId) {
		return true
	}

	return false
}

// SetSessionRunId gets a reference to the given string and assigns it to the SessionRunId field.
func (o *GetJobsParams) SetSessionRunId(v string) {
	o.SessionRunId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GetJobsParams) GetLastUpdated() string {
	if o == nil || IsNil(o.LastUpdated) {
		var ret string
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetJobsParams) GetLastUpdatedOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *GetJobsParams) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given string and assigns it to the LastUpdated field.
func (o *GetJobsParams) SetLastUpdated(v string) {
	o.LastUpdated = &v
}

func (o GetJobsParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetJobsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BranchId) {
		toSerialize["branchId"] = o.BranchId
	}
	if !IsNil(o.SessionRunId) {
		toSerialize["sessionRunId"] = o.SessionRunId
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	return toSerialize, nil
}

type NullableGetJobsParams struct {
	value *GetJobsParams
	isSet bool
}

func (v NullableGetJobsParams) Get() *GetJobsParams {
	return v.value
}

func (v *NullableGetJobsParams) Set(val *GetJobsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableGetJobsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableGetJobsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetJobsParams(val *GetJobsParams) *NullableGetJobsParams {
	return &NullableGetJobsParams{value: val, isSet: true}
}

func (v NullableGetJobsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetJobsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


