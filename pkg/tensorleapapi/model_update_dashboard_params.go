/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UpdateDashboardParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateDashboardParams{}

// UpdateDashboardParams struct for UpdateDashboardParams
type UpdateDashboardParams struct {
	DashboardId string          `json:"dashboardId"`
	Name        string          `json:"name"`
	Items       []DashboardItem `json:"items"`
	Description *string         `json:"description,omitempty"`
}

// NewUpdateDashboardParams instantiates a new UpdateDashboardParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDashboardParams(dashboardId string, name string, items []DashboardItem) *UpdateDashboardParams {
	this := UpdateDashboardParams{}
	this.DashboardId = dashboardId
	this.Name = name
	this.Items = items
	return &this
}

// NewUpdateDashboardParamsWithDefaults instantiates a new UpdateDashboardParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDashboardParamsWithDefaults() *UpdateDashboardParams {
	this := UpdateDashboardParams{}
	return &this
}

// GetDashboardId returns the DashboardId field value
func (o *UpdateDashboardParams) GetDashboardId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DashboardId
}

// GetDashboardIdOk returns a tuple with the DashboardId field value
// and a boolean to check if the value has been set.
func (o *UpdateDashboardParams) GetDashboardIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DashboardId, true
}

// SetDashboardId sets field value
func (o *UpdateDashboardParams) SetDashboardId(v string) {
	o.DashboardId = v
}

// GetName returns the Name field value
func (o *UpdateDashboardParams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateDashboardParams) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateDashboardParams) SetName(v string) {
	o.Name = v
}

// GetItems returns the Items field value
func (o *UpdateDashboardParams) GetItems() []DashboardItem {
	if o == nil {
		var ret []DashboardItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *UpdateDashboardParams) GetItemsOk() ([]DashboardItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *UpdateDashboardParams) SetItems(v []DashboardItem) {
	o.Items = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateDashboardParams) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDashboardParams) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateDashboardParams) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateDashboardParams) SetDescription(v string) {
	o.Description = &v
}

func (o UpdateDashboardParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateDashboardParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dashboardId"] = o.DashboardId
	toSerialize["name"] = o.Name
	toSerialize["items"] = o.Items
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableUpdateDashboardParams struct {
	value *UpdateDashboardParams
	isSet bool
}

func (v NullableUpdateDashboardParams) Get() *UpdateDashboardParams {
	return v.value
}

func (v *NullableUpdateDashboardParams) Set(val *UpdateDashboardParams) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDashboardParams) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDashboardParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDashboardParams(val *UpdateDashboardParams) *NullableUpdateDashboardParams {
	return &NullableUpdateDashboardParams{value: val, isSet: true}
}

func (v NullableUpdateDashboardParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDashboardParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
