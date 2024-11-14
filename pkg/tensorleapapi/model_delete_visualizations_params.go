/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.557
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the DeleteVisualizationsParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteVisualizationsParams{}

// DeleteVisualizationsParams struct for DeleteVisualizationsParams
type DeleteVisualizationsParams struct {
	VisualizationIdsToDelete []string `json:"visualizationIdsToDelete"`
	ProjectId                string   `json:"projectId"`
}

// NewDeleteVisualizationsParams instantiates a new DeleteVisualizationsParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteVisualizationsParams(visualizationIdsToDelete []string, projectId string) *DeleteVisualizationsParams {
	this := DeleteVisualizationsParams{}
	this.VisualizationIdsToDelete = visualizationIdsToDelete
	this.ProjectId = projectId
	return &this
}

// NewDeleteVisualizationsParamsWithDefaults instantiates a new DeleteVisualizationsParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteVisualizationsParamsWithDefaults() *DeleteVisualizationsParams {
	this := DeleteVisualizationsParams{}
	return &this
}

// GetVisualizationIdsToDelete returns the VisualizationIdsToDelete field value
func (o *DeleteVisualizationsParams) GetVisualizationIdsToDelete() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.VisualizationIdsToDelete
}

// GetVisualizationIdsToDeleteOk returns a tuple with the VisualizationIdsToDelete field value
// and a boolean to check if the value has been set.
func (o *DeleteVisualizationsParams) GetVisualizationIdsToDeleteOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.VisualizationIdsToDelete, true
}

// SetVisualizationIdsToDelete sets field value
func (o *DeleteVisualizationsParams) SetVisualizationIdsToDelete(v []string) {
	o.VisualizationIdsToDelete = v
}

// GetProjectId returns the ProjectId field value
func (o *DeleteVisualizationsParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *DeleteVisualizationsParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *DeleteVisualizationsParams) SetProjectId(v string) {
	o.ProjectId = v
}

func (o DeleteVisualizationsParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteVisualizationsParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["visualizationIdsToDelete"] = o.VisualizationIdsToDelete
	toSerialize["projectId"] = o.ProjectId
	return toSerialize, nil
}

type NullableDeleteVisualizationsParams struct {
	value *DeleteVisualizationsParams
	isSet bool
}

func (v NullableDeleteVisualizationsParams) Get() *DeleteVisualizationsParams {
	return v.value
}

func (v *NullableDeleteVisualizationsParams) Set(val *DeleteVisualizationsParams) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteVisualizationsParams) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteVisualizationsParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteVisualizationsParams(val *DeleteVisualizationsParams) *NullableDeleteVisualizationsParams {
	return &NullableDeleteVisualizationsParams{value: val, isSet: true}
}

func (v NullableDeleteVisualizationsParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteVisualizationsParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
