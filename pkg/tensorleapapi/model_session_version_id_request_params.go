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

// checks if the SessionVersionIdRequestParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionVersionIdRequestParams{}

// SessionVersionIdRequestParams struct for SessionVersionIdRequestParams
type SessionVersionIdRequestParams struct {
	VersionId string `json:"versionId"`
	ProjectId string `json:"projectId"`
}

// NewSessionVersionIdRequestParams instantiates a new SessionVersionIdRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionVersionIdRequestParams(versionId string, projectId string) *SessionVersionIdRequestParams {
	this := SessionVersionIdRequestParams{}
	this.VersionId = versionId
	this.ProjectId = projectId
	return &this
}

// NewSessionVersionIdRequestParamsWithDefaults instantiates a new SessionVersionIdRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionVersionIdRequestParamsWithDefaults() *SessionVersionIdRequestParams {
	this := SessionVersionIdRequestParams{}
	return &this
}

// GetVersionId returns the VersionId field value
func (o *SessionVersionIdRequestParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *SessionVersionIdRequestParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *SessionVersionIdRequestParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetProjectId returns the ProjectId field value
func (o *SessionVersionIdRequestParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *SessionVersionIdRequestParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *SessionVersionIdRequestParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o SessionVersionIdRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionVersionIdRequestParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["versionId"] = o.VersionId
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableSessionVersionIdRequestParams struct {
	value *SessionVersionIdRequestParams
	isSet bool
}

func (v NullableSessionVersionIdRequestParams) Get() *SessionVersionIdRequestParams {
	return v.value
}

func (v *NullableSessionVersionIdRequestParams) Set(val *SessionVersionIdRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionVersionIdRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionVersionIdRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionVersionIdRequestParams(val *SessionVersionIdRequestParams) *NullableSessionVersionIdRequestParams {
	return &NullableSessionVersionIdRequestParams{value: val, isSet: true}
}

func (v NullableSessionVersionIdRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionVersionIdRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
