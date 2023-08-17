/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GetEnvironmentInfoResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEnvironmentInfoResponse{}

// GetEnvironmentInfoResponse struct for GetEnvironmentInfoResponse
type GetEnvironmentInfoResponse struct {
	ClientStoragePrefixUrl string `json:"clientStoragePrefixUrl"`
	SchemaVersion float64 `json:"schemaVersion"`
	HubDefaultNamespace string `json:"hubDefaultNamespace"`
	HubUrl string `json:"hubUrl"`
}

// NewGetEnvironmentInfoResponse instantiates a new GetEnvironmentInfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEnvironmentInfoResponse(clientStoragePrefixUrl string, schemaVersion float64, hubDefaultNamespace string, hubUrl string) *GetEnvironmentInfoResponse {
	this := GetEnvironmentInfoResponse{}
	this.ClientStoragePrefixUrl = clientStoragePrefixUrl
	this.SchemaVersion = schemaVersion
	this.HubDefaultNamespace = hubDefaultNamespace
	this.HubUrl = hubUrl
	return &this
}

// NewGetEnvironmentInfoResponseWithDefaults instantiates a new GetEnvironmentInfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEnvironmentInfoResponseWithDefaults() *GetEnvironmentInfoResponse {
	this := GetEnvironmentInfoResponse{}
	return &this
}

// GetClientStoragePrefixUrl returns the ClientStoragePrefixUrl field value
func (o *GetEnvironmentInfoResponse) GetClientStoragePrefixUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientStoragePrefixUrl
}

// GetClientStoragePrefixUrlOk returns a tuple with the ClientStoragePrefixUrl field value
// and a boolean to check if the value has been set.
func (o *GetEnvironmentInfoResponse) GetClientStoragePrefixUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientStoragePrefixUrl, true
}

// SetClientStoragePrefixUrl sets field value
func (o *GetEnvironmentInfoResponse) SetClientStoragePrefixUrl(v string) {
	o.ClientStoragePrefixUrl = v
}

// GetSchemaVersion returns the SchemaVersion field value
func (o *GetEnvironmentInfoResponse) GetSchemaVersion() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.SchemaVersion
}

// GetSchemaVersionOk returns a tuple with the SchemaVersion field value
// and a boolean to check if the value has been set.
func (o *GetEnvironmentInfoResponse) GetSchemaVersionOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaVersion, true
}

// SetSchemaVersion sets field value
func (o *GetEnvironmentInfoResponse) SetSchemaVersion(v float64) {
	o.SchemaVersion = v
}

// GetHubDefaultNamespace returns the HubDefaultNamespace field value
func (o *GetEnvironmentInfoResponse) GetHubDefaultNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubDefaultNamespace
}

// GetHubDefaultNamespaceOk returns a tuple with the HubDefaultNamespace field value
// and a boolean to check if the value has been set.
func (o *GetEnvironmentInfoResponse) GetHubDefaultNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubDefaultNamespace, true
}

// SetHubDefaultNamespace sets field value
func (o *GetEnvironmentInfoResponse) SetHubDefaultNamespace(v string) {
	o.HubDefaultNamespace = v
}

// GetHubUrl returns the HubUrl field value
func (o *GetEnvironmentInfoResponse) GetHubUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HubUrl
}

// GetHubUrlOk returns a tuple with the HubUrl field value
// and a boolean to check if the value has been set.
func (o *GetEnvironmentInfoResponse) GetHubUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HubUrl, true
}

// SetHubUrl sets field value
func (o *GetEnvironmentInfoResponse) SetHubUrl(v string) {
	o.HubUrl = v
}

func (o GetEnvironmentInfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEnvironmentInfoResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clientStoragePrefixUrl"] = o.ClientStoragePrefixUrl
	toSerialize["schemaVersion"] = o.SchemaVersion
	toSerialize["hubDefaultNamespace"] = o.HubDefaultNamespace
	toSerialize["hubUrl"] = o.HubUrl
	return toSerialize, nil
}

type NullableGetEnvironmentInfoResponse struct {
	value *GetEnvironmentInfoResponse
	isSet bool
}

func (v NullableGetEnvironmentInfoResponse) Get() *GetEnvironmentInfoResponse {
	return v.value
}

func (v *NullableGetEnvironmentInfoResponse) Set(val *GetEnvironmentInfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEnvironmentInfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEnvironmentInfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEnvironmentInfoResponse(val *GetEnvironmentInfoResponse) *NullableGetEnvironmentInfoResponse {
	return &NullableGetEnvironmentInfoResponse{value: val, isSet: true}
}

func (v NullableGetEnvironmentInfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEnvironmentInfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


