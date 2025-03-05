/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.607
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the CodeIntegrationMappingErrorsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CodeIntegrationMappingErrorsParams{}

// CodeIntegrationMappingErrorsParams struct for CodeIntegrationMappingErrorsParams
type CodeIntegrationMappingErrorsParams struct {
	ProjectId string `json:"projectId"`
	VersionId string `json:"versionId"`
	Mapping   string `json:"mapping"`
}

// NewCodeIntegrationMappingErrorsParams instantiates a new CodeIntegrationMappingErrorsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCodeIntegrationMappingErrorsParams(projectId string, versionId string, mapping string) *CodeIntegrationMappingErrorsParams {
	this := CodeIntegrationMappingErrorsParams{}
	this.ProjectId = projectId
	this.VersionId = versionId
	this.Mapping = mapping
	return &this
}

// NewCodeIntegrationMappingErrorsParamsWithDefaults instantiates a new CodeIntegrationMappingErrorsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCodeIntegrationMappingErrorsParamsWithDefaults() *CodeIntegrationMappingErrorsParams {
	this := CodeIntegrationMappingErrorsParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *CodeIntegrationMappingErrorsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *CodeIntegrationMappingErrorsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *CodeIntegrationMappingErrorsParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetVersionId returns the VersionId field value
func (o *CodeIntegrationMappingErrorsParams) GetVersionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VersionId
}

// GetVersionIdOk returns a tuple with the VersionId field value
// and a boolean to check if the value has been set.
func (o *CodeIntegrationMappingErrorsParams) GetVersionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VersionId, true
}

// SetVersionId sets field value
func (o *CodeIntegrationMappingErrorsParams) SetVersionId(v string) {
	o.VersionId = v
}

// GetMapping returns the Mapping field value
func (o *CodeIntegrationMappingErrorsParams) GetMapping() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mapping
}

// GetMappingOk returns a tuple with the Mapping field value
// and a boolean to check if the value has been set.
func (o *CodeIntegrationMappingErrorsParams) GetMappingOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mapping, true
}

// SetMapping sets field value
func (o *CodeIntegrationMappingErrorsParams) SetMapping(v string) {
	o.Mapping = v
}

func (o CodeIntegrationMappingErrorsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CodeIntegrationMappingErrorsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["versionId"] = o.VersionId
	toSerialize["mapping"] = o.Mapping
	return toSerialize, nil
}

type NullableCodeIntegrationMappingErrorsParams struct {
	value *CodeIntegrationMappingErrorsParams
	isSet bool
}

func (v NullableCodeIntegrationMappingErrorsParams) Get() *CodeIntegrationMappingErrorsParams {
	return v.value
}

func (v *NullableCodeIntegrationMappingErrorsParams) Set(val *CodeIntegrationMappingErrorsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeIntegrationMappingErrorsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeIntegrationMappingErrorsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeIntegrationMappingErrorsParams(val *CodeIntegrationMappingErrorsParams) *NullableCodeIntegrationMappingErrorsParams {
	return &NullableCodeIntegrationMappingErrorsParams{value: val, isSet: true}
}

func (v NullableCodeIntegrationMappingErrorsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeIntegrationMappingErrorsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
