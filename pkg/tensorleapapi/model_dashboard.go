/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Dashboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Dashboard{}

// Dashboard struct for Dashboard
type Dashboard struct {
	Cid string `json:"cid"`
	ProjectId string `json:"projectId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedBy string `json:"createdBy"`
	Name string `json:"name"`
	Items []DashboardItem `json:"items"`
	Description *string `json:"description,omitempty"`
}

// NewDashboard instantiates a new Dashboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDashboard(cid string, projectId string, createdAt time.Time, updatedAt time.Time, createdBy string, name string, items []DashboardItem) *Dashboard {
	this := Dashboard{}
	this.Cid = cid
	this.ProjectId = projectId
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.CreatedBy = createdBy
	this.Name = name
	this.Items = items
	return &this
}

// NewDashboardWithDefaults instantiates a new Dashboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDashboardWithDefaults() *Dashboard {
	this := Dashboard{}
	return &this
}

// GetCid returns the Cid field value
func (o *Dashboard) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Dashboard) SetCid(v string) {
	o.Cid = v
}

// GetProjectId returns the ProjectId field value
func (o *Dashboard) GetProjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectId
}

// GetProjectIdOk returns a tuple with the ProjectId field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetProjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectId, true
}

// SetProjectId sets field value
func (o *Dashboard) SetProjectId(v string) {
	o.ProjectId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Dashboard) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Dashboard) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Dashboard) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Dashboard) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Dashboard) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Dashboard) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetName returns the Name field value
func (o *Dashboard) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Dashboard) SetName(v string) {
	o.Name = v
}

// GetItems returns the Items field value
func (o *Dashboard) GetItems() []DashboardItem {
	if o == nil {
		var ret []DashboardItem
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *Dashboard) GetItemsOk() ([]DashboardItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.Items, true
}

// SetItems sets field value
func (o *Dashboard) SetItems(v []DashboardItem) {
	o.Items = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Dashboard) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Dashboard) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Dashboard) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Dashboard) SetDescription(v string) {
	o.Description = &v
}

func (o Dashboard) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Dashboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["projectId"] = o.ProjectId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["updatedAt"] = o.UpdatedAt
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["name"] = o.Name
	toSerialize["items"] = o.Items
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableDashboard struct {
	value *Dashboard
	isSet bool
}

func (v NullableDashboard) Get() *Dashboard {
	return v.value
}

func (v *NullableDashboard) Set(val *Dashboard) {
	v.value = val
	v.isSet = true
}

func (v NullableDashboard) IsSet() bool {
	return v.isSet
}

func (v *NullableDashboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDashboard(val *Dashboard) *NullableDashboard {
	return &NullableDashboard{value: val, isSet: true}
}

func (v NullableDashboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDashboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


