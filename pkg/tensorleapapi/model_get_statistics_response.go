/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.360
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetStatisticsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatisticsResponse{}

// GetStatisticsResponse struct for GetStatisticsResponse
type GetStatisticsResponse struct {
	ActiveUsers float64 `json:"activeUsers"`
	OpenIssues  float64 `json:"openIssues"`
	Tests       float64 `json:"tests"`
	Projects    float64 `json:"projects"`
	Networks    float64 `json:"networks"`
	Sessions    float64 `json:"sessions"`
}

// NewGetStatisticsResponse instantiates a new GetStatisticsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatisticsResponse(activeUsers float64, openIssues float64, tests float64, projects float64, networks float64, sessions float64) *GetStatisticsResponse {
	this := GetStatisticsResponse{}
	this.ActiveUsers = activeUsers
	this.OpenIssues = openIssues
	this.Tests = tests
	this.Projects = projects
	this.Networks = networks
	this.Sessions = sessions
	return &this
}

// NewGetStatisticsResponseWithDefaults instantiates a new GetStatisticsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatisticsResponseWithDefaults() *GetStatisticsResponse {
	this := GetStatisticsResponse{}
	return &this
}

// GetActiveUsers returns the ActiveUsers field value
func (o *GetStatisticsResponse) GetActiveUsers() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.ActiveUsers
}

// GetActiveUsersOk returns a tuple with the ActiveUsers field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetActiveUsersOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ActiveUsers, true
}

// SetActiveUsers sets field value
func (o *GetStatisticsResponse) SetActiveUsers(v float64) {
	o.ActiveUsers = v
}

// GetOpenIssues returns the OpenIssues field value
func (o *GetStatisticsResponse) GetOpenIssues() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.OpenIssues
}

// GetOpenIssuesOk returns a tuple with the OpenIssues field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetOpenIssuesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpenIssues, true
}

// SetOpenIssues sets field value
func (o *GetStatisticsResponse) SetOpenIssues(v float64) {
	o.OpenIssues = v
}

// GetTests returns the Tests field value
func (o *GetStatisticsResponse) GetTests() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Tests
}

// GetTestsOk returns a tuple with the Tests field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetTestsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tests, true
}

// SetTests sets field value
func (o *GetStatisticsResponse) SetTests(v float64) {
	o.Tests = v
}

// GetProjects returns the Projects field value
func (o *GetStatisticsResponse) GetProjects() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetProjectsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Projects, true
}

// SetProjects sets field value
func (o *GetStatisticsResponse) SetProjects(v float64) {
	o.Projects = v
}

// GetNetworks returns the Networks field value
func (o *GetStatisticsResponse) GetNetworks() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetNetworksOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Networks, true
}

// SetNetworks sets field value
func (o *GetStatisticsResponse) SetNetworks(v float64) {
	o.Networks = v
}

// GetSessions returns the Sessions field value
func (o *GetStatisticsResponse) GetSessions() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Sessions
}

// GetSessionsOk returns a tuple with the Sessions field value
// and a boolean to check if the value has been set.
func (o *GetStatisticsResponse) GetSessionsOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sessions, true
}

// SetSessions sets field value
func (o *GetStatisticsResponse) SetSessions(v float64) {
	o.Sessions = v
}

func (o GetStatisticsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatisticsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["activeUsers"] = o.ActiveUsers
	toSerialize["openIssues"] = o.OpenIssues
	toSerialize["tests"] = o.Tests
	toSerialize["projects"] = o.Projects
	toSerialize["networks"] = o.Networks
	toSerialize["sessions"] = o.Sessions
	return toSerialize, nil
}

type NullableGetStatisticsResponse struct {
	value *GetStatisticsResponse
	isSet bool
}

func (v NullableGetStatisticsResponse) Get() *GetStatisticsResponse {
	return v.value
}

func (v *NullableGetStatisticsResponse) Set(val *GetStatisticsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatisticsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatisticsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatisticsResponse(val *GetStatisticsResponse) *NullableGetStatisticsResponse {
	return &NullableGetStatisticsResponse{value: val, isSet: true}
}

func (v NullableGetStatisticsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatisticsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
