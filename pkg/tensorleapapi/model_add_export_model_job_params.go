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

// checks if the AddExportModelJobParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddExportModelJobParams{}

// AddExportModelJobParams struct for AddExportModelJobParams
type AddExportModelJobParams struct {
	ProjectId       string              `json:"projectId"`
	SessionWeightId string              `json:"sessionWeightId"`
	ExportModelType ExportModelTypeEnum `json:"exportModelType"`
	Title           string              `json:"title"`
	PruneModel      bool                `json:"pruneModel"`
}

// NewAddExportModelJobParams instantiates a new AddExportModelJobParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddExportModelJobParams(projectId string, sessionWeightId string, exportModelType ExportModelTypeEnum, title string, pruneModel bool) *AddExportModelJobParams {
	this := AddExportModelJobParams{}
	this.ProjectId = projectId
	this.SessionWeightId = sessionWeightId
	this.ExportModelType = exportModelType
	this.Title = title
	this.PruneModel = pruneModel
	return &this
}

// NewAddExportModelJobParamsWithDefaults instantiates a new AddExportModelJobParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddExportModelJobParamsWithDefaults() *AddExportModelJobParams {
	this := AddExportModelJobParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *AddExportModelJobParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddExportModelJobParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddExportModelJobParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetSessionWeightId returns the SessionWeightId field value
func (o *AddExportModelJobParams) GetSessionWeightId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionWeightId
}

// GetSessionWeightIdOk returns a tuple with the SessionWeightId field value
// and a boolean to check if the value has been set.
func (o *AddExportModelJobParams) GetSessionWeightIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionWeightId, true
}

// SetSessionWeightId sets field value
func (o *AddExportModelJobParams) SetSessionWeightId(v string) {
	o.SessionWeightId = v
}

// GetExportModelType returns the ExportModelType field value
func (o *AddExportModelJobParams) GetExportModelType() ExportModelTypeEnum {
	if o == nil {
		var ret ExportModelTypeEnum
		return ret
	}

	return o.ExportModelType
}

// GetExportModelTypeOk returns a tuple with the ExportModelType field value
// and a boolean to check if the value has been set.
func (o *AddExportModelJobParams) GetExportModelTypeOk() (*ExportModelTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExportModelType, true
}

// SetExportModelType sets field value
func (o *AddExportModelJobParams) SetExportModelType(v ExportModelTypeEnum) {
	o.ExportModelType = v
}

// GetTitle returns the Title field value
func (o *AddExportModelJobParams) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *AddExportModelJobParams) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *AddExportModelJobParams) SetTitle(v string) {
	o.Title = v
}

// GetPruneModel returns the PruneModel field value
func (o *AddExportModelJobParams) GetPruneModel() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.PruneModel
}

// GetPruneModelOk returns a tuple with the PruneModel field value
// and a boolean to check if the value has been set.
func (o *AddExportModelJobParams) GetPruneModelOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PruneModel, true
}

// SetPruneModel sets field value
func (o *AddExportModelJobParams) SetPruneModel(v bool) {
	o.PruneModel = v
}

func (o AddExportModelJobParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddExportModelJobParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["sessionWeightId"] = o.SessionWeightId
	toSerialize["exportModelType"] = o.ExportModelType
	toSerialize["title"] = o.Title
	toSerialize["pruneModel"] = o.PruneModel
	return toSerialize, nil
}

type NullableAddExportModelJobParams struct {
	value *AddExportModelJobParams
	isSet bool
}

func (v NullableAddExportModelJobParams) Get() *AddExportModelJobParams {
	return v.value
}

func (v *NullableAddExportModelJobParams) Set(val *AddExportModelJobParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddExportModelJobParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddExportModelJobParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddExportModelJobParams(val *AddExportModelJobParams) *NullableAddExportModelJobParams {
	return &NullableAddExportModelJobParams{value: val, isSet: true}
}

func (v NullableAddExportModelJobParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddExportModelJobParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
