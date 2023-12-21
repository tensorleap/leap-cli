/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetIssueFileUploadSignedUrl type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetIssueFileUploadSignedUrl{}

// GetIssueFileUploadSignedUrl struct for GetIssueFileUploadSignedUrl
type GetIssueFileUploadSignedUrl struct {
	FileName string `json:"fileName"`
	ProjectId string `json:"projectId"`
	IssueId string `json:"issueId"`
}

// NewGetIssueFileUploadSignedUrl instantiates a new GetIssueFileUploadSignedUrl object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetIssueFileUploadSignedUrl(fileName string, projectId string, issueId string) *GetIssueFileUploadSignedUrl {
	this := GetIssueFileUploadSignedUrl{}
	this.FileName = fileName
	this.ProjectId = projectId
	this.IssueId = issueId
	return &this
}

// NewGetIssueFileUploadSignedUrlWithDefaults instantiates a new GetIssueFileUploadSignedUrl object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetIssueFileUploadSignedUrlWithDefaults() *GetIssueFileUploadSignedUrl {
	this := GetIssueFileUploadSignedUrl{}
	return &this
}

// GetFileName returns the FileName field value
func (o *GetIssueFileUploadSignedUrl) GetFileName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FileName
}

// GetFileNameOk returns a tuple with the FileName field value
// and a boolean to check if the value has been set.
func (o *GetIssueFileUploadSignedUrl) GetFileNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileName, true
}

// SetFileName sets field value
func (o *GetIssueFileUploadSignedUrl) SetFileName(v string) {
	o.FileName = v
}

// GetProjectId returns the ProjectId field value
func (o *GetIssueFileUploadSignedUrl) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *GetIssueFileUploadSignedUrl) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *GetIssueFileUploadSignedUrl) SetProjectId(v string) {
	o.ProjectId = v
}

// GetIssueId returns the IssueId field value
func (o *GetIssueFileUploadSignedUrl) GetIssueId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IssueId
}

// GetIssueIdOk returns a tuple with the IssueId field value
// and a boolean to check if the value has been set.
func (o *GetIssueFileUploadSignedUrl) GetIssueIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IssueId, true
}

// SetIssueId sets field value
func (o *GetIssueFileUploadSignedUrl) SetIssueId(v string) {
	o.IssueId = v
}

func (o GetIssueFileUploadSignedUrl) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetIssueFileUploadSignedUrl) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fileName"] = o.FileName
	toSerialize["projectId"] = o.ProjectId
	toSerialize["issueId"] = o.IssueId
	return toSerialize, nil
}

type NullableGetIssueFileUploadSignedUrl struct {
	value *GetIssueFileUploadSignedUrl
	isSet bool
}

func (v NullableGetIssueFileUploadSignedUrl) Get() *GetIssueFileUploadSignedUrl {
	return v.value
}

func (v *NullableGetIssueFileUploadSignedUrl) Set(val *GetIssueFileUploadSignedUrl) {
	v.value = val
	v.isSet = true
}

func (v NullableGetIssueFileUploadSignedUrl) IsSet() bool {
	return v.isSet
}

func (v *NullableGetIssueFileUploadSignedUrl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetIssueFileUploadSignedUrl(val *GetIssueFileUploadSignedUrl) *NullableGetIssueFileUploadSignedUrl {
	return &NullableGetIssueFileUploadSignedUrl{value: val, isSet: true}
}

func (v NullableGetIssueFileUploadSignedUrl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetIssueFileUploadSignedUrl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


