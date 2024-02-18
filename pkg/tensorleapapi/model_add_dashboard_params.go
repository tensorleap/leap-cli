/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the AddDashboardParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddDashboardParams{}

// AddDashboardParams struct for AddDashboardParams
type AddDashboardParams struct {
	ProjectId   string    `json:"projectId"`
	Name        string    `json:"name"`
	Items       []Dashlet `json:"items"`
	Description *string   `json:"description,omitempty"`
}

// NewAddDashboardParams instantiates a new AddDashboardParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddDashboardParams(projectId string, name string, items []Dashlet) *AddDashboardParams {
	this := AddDashboardParams{}
	this.ProjectId = projectId
	this.Name = name
	this.Items = items
	return &this
}

// NewAddDashboardParamsWithDefaults instantiates a new AddDashboardParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddDashboardParamsWithDefaults() *AddDashboardParams {
	this := AddDashboardParams{}
	return &this
}

// GetProjectId returns the ProjectId field value
func (o *AddDashboardParams) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *AddDashboardParams) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *AddDashboardParams) SetProjectId(v string) {
	o.ProjectId = v
}

// GetName returns the Name field value
func (o *AddDashboardParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AddDashboardParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AddDashboardParams) SetName(v string) {
	o.Name = v
}

// GetItems returns the Items field value
func (o *AddDashboardParams) GetItems() []Dashlet {
	if o == nil {
		var ret []Dashlet
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AddDashboardParams) GetItemsOk() ([]Dashlet, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *AddDashboardParams) SetItems(v []Dashlet) {
	o.Items = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AddDashboardParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddDashboardParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AddDashboardParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AddDashboardParams) SetDescription(v string) {
	o.Description = &v
}

func (o AddDashboardParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddDashboardParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["projectId"] = o.ProjectId
	toSerialize["name"] = o.Name
	toSerialize["items"] = o.Items
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableAddDashboardParams struct {
	value *AddDashboardParams
	isSet bool
}

func (v NullableAddDashboardParams) Get() *AddDashboardParams {
	return v.value
}

func (v *NullableAddDashboardParams) Set(val *AddDashboardParams) {
	v.value = val
	v.isSet = true
}

func (v NullableAddDashboardParams) IsSet() bool {
	return v.isSet
}

func (v *NullableAddDashboardParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddDashboardParams(val *AddDashboardParams) *NullableAddDashboardParams {
	return &NullableAddDashboardParams{value: val, isSet: true}
}

func (v NullableAddDashboardParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddDashboardParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
