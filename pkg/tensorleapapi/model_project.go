/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the Project type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Project{}

// Project struct for Project
type Project struct {
	Cid              string           `json:"cid"`
	TeamId           string           `json:"teamId"`
	CreatedBy        string           `json:"createdBy"`
	CreatedAt        NullableTime     `json:"createdAt"`
	Name             string           `json:"name"`
	LastAccessed     time.Time        `json:"lastAccessed"`
	Status           ProjectStatus    `json:"status"`
	Access           ProjectAccess    `json:"access"`
	Tags             []string         `json:"tags"`
	BgImagePath      *string          `json:"bgImagePath,omitempty"`
	HubPublishPolicy HubPublishPolicy `json:"hubPublishPolicy"`
	// Construct a type with a set of properties K of type T
	Categories  map[string]interface{} `json:"categories,omitempty"`
	Description *string                `json:"description,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(cid string, teamId string, createdBy string, createdAt NullableTime, name string, lastAccessed time.Time, status ProjectStatus, access ProjectAccess, tags []string, hubPublishPolicy HubPublishPolicy) *Project {
	this := Project{}
	this.Cid = cid
	this.TeamId = teamId
	this.CreatedBy = createdBy
	this.CreatedAt = createdAt
	this.Name = name
	this.LastAccessed = lastAccessed
	this.Status = status
	this.Access = access
	this.Tags = tags
	this.HubPublishPolicy = hubPublishPolicy
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetCid returns the Cid field value
func (o *Project) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *Project) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *Project) SetCid(v string) {
	o.Cid = v
}

// GetTeamId returns the TeamId field value
func (o *Project) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *Project) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *Project) SetTeamId(v string) {
	o.TeamId = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *Project) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *Project) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *Project) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetCreatedAt returns the CreatedAt field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *Project) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.CreatedAt.Get()
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Project) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreatedAt.Get(), o.CreatedAt.IsSet()
}

// SetCreatedAt sets field value
func (o *Project) SetCreatedAt(v time.Time) {
	o.CreatedAt.Set(&v)
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetLastAccessed returns the LastAccessed field value
func (o *Project) GetLastAccessed() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastAccessed
}

// GetLastAccessedOk returns a tuple with the LastAccessed field value
// and a boolean to check if the value has been set.
func (o *Project) GetLastAccessedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastAccessed, true
}

// SetLastAccessed sets field value
func (o *Project) SetLastAccessed(v time.Time) {
	o.LastAccessed = v
}

// GetStatus returns the Status field value
func (o *Project) GetStatus() ProjectStatus {
	if o == nil {
		var ret ProjectStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Project) GetStatusOk() (*ProjectStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Project) SetStatus(v ProjectStatus) {
	o.Status = v
}

// GetAccess returns the Access field value
func (o *Project) GetAccess() ProjectAccess {
	if o == nil {
		var ret ProjectAccess
		return ret
	}

	return o.Access
}

// GetAccessOk returns a tuple with the Access field value
// and a boolean to check if the value has been set.
func (o *Project) GetAccessOk() (*ProjectAccess, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Access, true
}

// SetAccess sets field value
func (o *Project) SetAccess(v ProjectAccess) {
	o.Access = v
}

// GetTags returns the Tags field value
func (o *Project) GetTags() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value
// and a boolean to check if the value has been set.
func (o *Project) GetTagsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tags, true
}

// SetTags sets field value
func (o *Project) SetTags(v []string) {
	o.Tags = v
}

// GetBgImagePath returns the BgImagePath field value if set, zero value otherwise.
func (o *Project) GetBgImagePath() string {
	if o == nil || IsNil(o.BgImagePath) {
		var ret string
		return ret
	}
	return *o.BgImagePath
}

// GetBgImagePathOk returns a tuple with the BgImagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBgImagePathOk() (*string, bool) {
	if o == nil || IsNil(o.BgImagePath) {
		return nil, false
	}
	return o.BgImagePath, true
}

// HasBgImagePath returns a boolean if a field has been set.
func (o *Project) HasBgImagePath() bool {
	if o != nil && !IsNil(o.BgImagePath) {
		return true
	}

	return false
}

// SetBgImagePath gets a reference to the given string and assigns it to the BgImagePath field.
func (o *Project) SetBgImagePath(v string) {
	o.BgImagePath = &v
}

// GetHubPublishPolicy returns the HubPublishPolicy field value
func (o *Project) GetHubPublishPolicy() HubPublishPolicy {
	if o == nil {
		var ret HubPublishPolicy
		return ret
	}

	return o.HubPublishPolicy
}

// GetHubPublishPolicyOk returns a tuple with the HubPublishPolicy field value
// and a boolean to check if the value has been set.
func (o *Project) GetHubPublishPolicyOk() (*HubPublishPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubPublishPolicy, true
}

// SetHubPublishPolicy sets field value
func (o *Project) SetHubPublishPolicy(v HubPublishPolicy) {
	o.HubPublishPolicy = v
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *Project) GetCategories() map[string]interface{} {
	if o == nil || IsNil(o.Categories) {
		var ret map[string]interface{}
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCategoriesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Categories) {
		return map[string]interface{}{}, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *Project) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given map[string]interface{} and assigns it to the Categories field.
func (o *Project) SetCategories(v map[string]interface{}) {
	o.Categories = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Project) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Project) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Project) SetDescription(v string) {
	o.Description = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Project) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["teamId"] = o.TeamId
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["createdAt"] = o.CreatedAt.Get()
	toSerialize["name"] = o.Name
	toSerialize["lastAccessed"] = o.LastAccessed
	toSerialize["status"] = o.Status
	toSerialize["access"] = o.Access
	toSerialize["tags"] = o.Tags
	if !IsNil(o.BgImagePath) {
		toSerialize["bgImagePath"] = o.BgImagePath
	}
	toSerialize["hubPublishPolicy"] = o.HubPublishPolicy
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
