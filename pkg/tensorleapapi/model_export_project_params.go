/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ExportProjectParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExportProjectParams{}

// ExportProjectParams struct for ExportProjectParams
type ExportProjectParams struct {
	ExportUrl         string            `json:"exportUrl"`
	ProjectVersion    float64           `json:"projectVersion"`
	ExportOptions     ExportOptions     `json:"exportOptions"`
	AlreadyExported   bool              `json:"alreadyExported"`
	ProjectExportMeta ExportProjectMeta `json:"projectExportMeta"`
	CopyToUrl         *string           `json:"copyToUrl,omitempty"`
}

// NewExportProjectParams instantiates a new ExportProjectParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportProjectParams(exportUrl string, projectVersion float64, exportOptions ExportOptions, alreadyExported bool, projectExportMeta ExportProjectMeta) *ExportProjectParams {
	this := ExportProjectParams{}
	this.ExportUrl = exportUrl
	this.ProjectVersion = projectVersion
	this.ExportOptions = exportOptions
	this.AlreadyExported = alreadyExported
	this.ProjectExportMeta = projectExportMeta
	return &this
}

// NewExportProjectParamsWithDefaults instantiates a new ExportProjectParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportProjectParamsWithDefaults() *ExportProjectParams {
	this := ExportProjectParams{}
	return &this
}

// GetExportUrl returns the ExportUrl field value
func (o *ExportProjectParams) GetExportUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExportUrl
}

// GetExportUrlOk returns a tuple with the ExportUrl field value
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetExportUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportUrl, true
}

// SetExportUrl sets field value
func (o *ExportProjectParams) SetExportUrl(v string) {
	o.ExportUrl = v
}

// GetProjectVersion returns the ProjectVersion field value
func (o *ExportProjectParams) GetProjectVersion() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ProjectVersion
}

// GetProjectVersionOk returns a tuple with the ProjectVersion field value
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetProjectVersionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectVersion, true
}

// SetProjectVersion sets field value
func (o *ExportProjectParams) SetProjectVersion(v float64) {
	o.ProjectVersion = v
}

// GetExportOptions returns the ExportOptions field value
func (o *ExportProjectParams) GetExportOptions() ExportOptions {
	if o == nil {
		var ret ExportOptions
		return ret
	}

	return o.ExportOptions
}

// GetExportOptionsOk returns a tuple with the ExportOptions field value
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetExportOptionsOk() (*ExportOptions, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportOptions, true
}

// SetExportOptions sets field value
func (o *ExportProjectParams) SetExportOptions(v ExportOptions) {
	o.ExportOptions = v
}

// GetAlreadyExported returns the AlreadyExported field value
func (o *ExportProjectParams) GetAlreadyExported() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AlreadyExported
}

// GetAlreadyExportedOk returns a tuple with the AlreadyExported field value
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetAlreadyExportedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AlreadyExported, true
}

// SetAlreadyExported sets field value
func (o *ExportProjectParams) SetAlreadyExported(v bool) {
	o.AlreadyExported = v
}

// GetProjectExportMeta returns the ProjectExportMeta field value
func (o *ExportProjectParams) GetProjectExportMeta() ExportProjectMeta {
	if o == nil {
		var ret ExportProjectMeta
		return ret
	}

	return o.ProjectExportMeta
}

// GetProjectExportMetaOk returns a tuple with the ProjectExportMeta field value
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetProjectExportMetaOk() (*ExportProjectMeta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectExportMeta, true
}

// SetProjectExportMeta sets field value
func (o *ExportProjectParams) SetProjectExportMeta(v ExportProjectMeta) {
	o.ProjectExportMeta = v
}

// GetCopyToUrl returns the CopyToUrl field value if set, zero value otherwise.
func (o *ExportProjectParams) GetCopyToUrl() string {
	if o == nil || IsNil(o.CopyToUrl) {
		var ret string
		return ret
	}
	return *o.CopyToUrl
}

// GetCopyToUrlOk returns a tuple with the CopyToUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportProjectParams) GetCopyToUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CopyToUrl) {
		return nil, false
	}
	return o.CopyToUrl, true
}

// HasCopyToUrl returns a boolean if a field has been set.
func (o *ExportProjectParams) HasCopyToUrl() bool {
	if o != nil && !IsNil(o.CopyToUrl) {
		return true
	}

	return false
}

// SetCopyToUrl gets a reference to the given string and assigns it to the CopyToUrl field.
func (o *ExportProjectParams) SetCopyToUrl(v string) {
	o.CopyToUrl = &v
}

func (o ExportProjectParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportProjectParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["exportUrl"] = o.ExportUrl
	toSerialize["projectVersion"] = o.ProjectVersion
	toSerialize["exportOptions"] = o.ExportOptions
	toSerialize["alreadyExported"] = o.AlreadyExported
	toSerialize["projectExportMeta"] = o.ProjectExportMeta
	if !IsNil(o.CopyToUrl) {
		toSerialize["copyToUrl"] = o.CopyToUrl
	}
	return toSerialize, nil
}

type NullableExportProjectParams struct {
	value *ExportProjectParams
	isSet bool
}

func (v NullableExportProjectParams) Get() *ExportProjectParams {
	return v.value
}

func (v *NullableExportProjectParams) Set(val *ExportProjectParams) {
	v.value = val
	v.isSet = true
}

func (v NullableExportProjectParams) IsSet() bool {
	return v.isSet
}

func (v *NullableExportProjectParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportProjectParams(val *ExportProjectParams) *NullableExportProjectParams {
	return &NullableExportProjectParams{value: val, isSet: true}
}

func (v NullableExportProjectParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportProjectParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
