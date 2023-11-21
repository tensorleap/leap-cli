/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the PublishProjectRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublishProjectRequest{}

// PublishProjectRequest struct for PublishProjectRequest
type PublishProjectRequest struct {
	ProjectId  string `json:"projectId"`
	PublishUrl string `json:"publishUrl"`
}

// NewPublishProjectRequest instantiates a new PublishProjectRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublishProjectRequest(projectId string, publishUrl string) *PublishProjectRequest {
	this := PublishProjectRequest{}
	this.ProjectId = projectId
	this.PublishUrl = publishUrl
	return &this
}

// NewPublishProjectRequestWithDefaults instantiates a new PublishProjectRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublishProjectRequestWithDefaults() *PublishProjectRequest {
	this := PublishProjectRequest{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *PublishProjectRequest) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *PublishProjectRequest) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *PublishProjectRequest) SetProjectId(v string) {
	o.ProjectId = v
}

// GetPublishUrl returns the PublishUrl field value
func (o *PublishProjectRequest) GetPublishUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublishUrl
}

// GetPublishUrlOk returns a tuple with the PublishUrl field value
// and a boolean to check if the value has been set.
func (o *PublishProjectRequest) GetPublishUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublishUrl, true
}

// SetPublishUrl sets field value
func (o *PublishProjectRequest) SetPublishUrl(v string) {
	o.PublishUrl = v
}

func (o PublishProjectRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublishProjectRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["publishUrl"] = o.PublishUrl
	return toSerialize, nil
}

type NullablePublishProjectRequest struct {
	value *PublishProjectRequest
	isSet bool
}

func (v NullablePublishProjectRequest) Get() *PublishProjectRequest {
	return v.value
}

func (v *NullablePublishProjectRequest) Set(val *PublishProjectRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePublishProjectRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePublishProjectRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublishProjectRequest(val *PublishProjectRequest) *NullablePublishProjectRequest {
	return &NullablePublishProjectRequest{value: val, isSet: true}
}

func (v NullablePublishProjectRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublishProjectRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}