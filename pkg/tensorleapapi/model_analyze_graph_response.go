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

// checks if the AnalyzeGraphResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnalyzeGraphResponse{}

// AnalyzeGraphResponse struct for AnalyzeGraphResponse
type AnalyzeGraphResponse struct {
	Digest string `json:"digest"`
	Response *GraphAnalyzerData `json:"response,omitempty"`
	Status JobStatus `json:"status"`
}

// NewAnalyzeGraphResponse instantiates a new AnalyzeGraphResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnalyzeGraphResponse(digest string, status JobStatus) *AnalyzeGraphResponse {
	this := AnalyzeGraphResponse{}
	this.Digest = digest
	this.Status = status
	return &this
}

// NewAnalyzeGraphResponseWithDefaults instantiates a new AnalyzeGraphResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnalyzeGraphResponseWithDefaults() *AnalyzeGraphResponse {
	this := AnalyzeGraphResponse{}
	return &this
}

// GetDigest returns the Digest field value
func (o *AnalyzeGraphResponse) GetDigest() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Digest
}

// GetDigestOk returns a tuple with the Digest field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphResponse) GetDigestOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Digest, true
}

// SetDigest sets field value
func (o *AnalyzeGraphResponse) SetDigest(v string) {
	o.Digest = v
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *AnalyzeGraphResponse) GetResponse() GraphAnalyzerData {
	if o == nil || IsNil(o.Response) {
		var ret GraphAnalyzerData
		return ret
	}
	return *o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphResponse) GetResponseOk() (*GraphAnalyzerData, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *AnalyzeGraphResponse) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given GraphAnalyzerData and assigns it to the Response field.
func (o *AnalyzeGraphResponse) SetResponse(v GraphAnalyzerData) {
	o.Response = &v
}

// GetStatus returns the Status field value
func (o *AnalyzeGraphResponse) GetStatus() JobStatus {
	if o == nil {
		var ret JobStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *AnalyzeGraphResponse) GetStatusOk() (*JobStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *AnalyzeGraphResponse) SetStatus(v JobStatus) {
	o.Status = v
}

func (o AnalyzeGraphResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnalyzeGraphResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["digest"] = o.Digest
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

type NullableAnalyzeGraphResponse struct {
	value *AnalyzeGraphResponse
	isSet bool
}

func (v NullableAnalyzeGraphResponse) Get() *AnalyzeGraphResponse {
	return v.value
}

func (v *NullableAnalyzeGraphResponse) Set(val *AnalyzeGraphResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyzeGraphResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyzeGraphResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyzeGraphResponse(val *AnalyzeGraphResponse) *NullableAnalyzeGraphResponse {
	return &NullableAnalyzeGraphResponse{value: val, isSet: true}
}

func (v NullableAnalyzeGraphResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyzeGraphResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


