/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SessionWeightData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionWeightData{}

// SessionWeightData struct for SessionWeightData
type SessionWeightData struct {
	Cid string `json:"cid"`
	SessionId string `json:"sessionId"`
	Epoch float64 `json:"epoch"`
	GlobalStep float64 `json:"globalStep"`
	Status StatusEnum `json:"status"`
	TeamId string `json:"teamId"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}

// NewSessionWeightData instantiates a new SessionWeightData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionWeightData(cid string, sessionId string, epoch float64, globalStep float64, status StatusEnum, teamId string, createdAt time.Time, createdBy string) *SessionWeightData {
	this := SessionWeightData{}
	this.Cid = cid
	this.SessionId = sessionId
	this.Epoch = epoch
	this.GlobalStep = globalStep
	this.Status = status
	this.TeamId = teamId
	this.CreatedAt = createdAt
	this.CreatedBy = createdBy
	return &this
}

// NewSessionWeightDataWithDefaults instantiates a new SessionWeightData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionWeightDataWithDefaults() *SessionWeightData {
	this := SessionWeightData{}
	return &this
}

// GetCid returns the Cid field value
func (o *SessionWeightData) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SessionWeightData) SetCid(v string) {
	o.Cid = v
}

// GetSessionId returns the SessionId field value
func (o *SessionWeightData) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *SessionWeightData) SetSessionId(v string) {
	o.SessionId = v
}

// GetEpoch returns the Epoch field value
func (o *SessionWeightData) GetEpoch() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Epoch
}

// GetEpochOk returns a tuple with the Epoch field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetEpochOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Epoch, true
}

// SetEpoch sets field value
func (o *SessionWeightData) SetEpoch(v float64) {
	o.Epoch = v
}

// GetGlobalStep returns the GlobalStep field value
func (o *SessionWeightData) GetGlobalStep() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.GlobalStep
}

// GetGlobalStepOk returns a tuple with the GlobalStep field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetGlobalStepOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalStep, true
}

// SetGlobalStep sets field value
func (o *SessionWeightData) SetGlobalStep(v float64) {
	o.GlobalStep = v
}

// GetStatus returns the Status field value
func (o *SessionWeightData) GetStatus() StatusEnum {
	if o == nil {
		var ret StatusEnum
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetStatusOk() (*StatusEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SessionWeightData) SetStatus(v StatusEnum) {
	o.Status = v
}

// GetTeamId returns the TeamId field value
func (o *SessionWeightData) GetTeamId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TeamId
}

// GetTeamIdOk returns a tuple with the TeamId field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetTeamIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TeamId, true
}

// SetTeamId sets field value
func (o *SessionWeightData) SetTeamId(v string) {
	o.TeamId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SessionWeightData) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SessionWeightData) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SessionWeightData) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SessionWeightData) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SessionWeightData) SetCreatedBy(v string) {
	o.CreatedBy = v
}

func (o SessionWeightData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionWeightData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["sessionId"] = o.SessionId
	toSerialize["epoch"] = o.Epoch
	toSerialize["globalStep"] = o.GlobalStep
	toSerialize["status"] = o.Status
	toSerialize["teamId"] = o.TeamId
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["createdBy"] = o.CreatedBy
	return toSerialize, nil
}

type NullableSessionWeightData struct {
	value *SessionWeightData
	isSet bool
}

func (v NullableSessionWeightData) Get() *SessionWeightData {
	return v.value
}

func (v *NullableSessionWeightData) Set(val *SessionWeightData) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionWeightData) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionWeightData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionWeightData(val *SessionWeightData) *NullableSessionWeightData {
	return &NullableSessionWeightData{value: val, isSet: true}
}

func (v NullableSessionWeightData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionWeightData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


