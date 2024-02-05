/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.445
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the FetchSimilarResponseReadyArtifacts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchSimilarResponseReadyArtifacts{}

// FetchSimilarResponseReadyArtifacts struct for FetchSimilarResponseReadyArtifacts
type FetchSimilarResponseReadyArtifacts struct {
	ClusterPath *string `json:"clusterPath,omitempty"`
}

// NewFetchSimilarResponseReadyArtifacts instantiates a new FetchSimilarResponseReadyArtifacts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchSimilarResponseReadyArtifacts() *FetchSimilarResponseReadyArtifacts {
	this := FetchSimilarResponseReadyArtifacts{}
	return &this
}

// NewFetchSimilarResponseReadyArtifactsWithDefaults instantiates a new FetchSimilarResponseReadyArtifacts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchSimilarResponseReadyArtifactsWithDefaults() *FetchSimilarResponseReadyArtifacts {
	this := FetchSimilarResponseReadyArtifacts{}
	return &this
}

// GetClusterPath returns the ClusterPath field value if set, zero value otherwise.
func (o *FetchSimilarResponseReadyArtifacts) GetClusterPath() string {
	if o == nil || IsNil(o.ClusterPath) {
		var ret string
		return ret
	}
	return *o.ClusterPath
}

// GetClusterPathOk returns a tuple with the ClusterPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchSimilarResponseReadyArtifacts) GetClusterPathOk() (*string, bool) {
	if o == nil || IsNil(o.ClusterPath) {
		return nil, false
	}
	return o.ClusterPath, true
}

// HasClusterPath returns a boolean if a field has been set.
func (o *FetchSimilarResponseReadyArtifacts) HasClusterPath() bool {
	if o != nil && !IsNil(o.ClusterPath) {
		return true
	}

	return false
}

// SetClusterPath gets a reference to the given string and assigns it to the ClusterPath field.
func (o *FetchSimilarResponseReadyArtifacts) SetClusterPath(v string) {
	o.ClusterPath = &v
}

func (o FetchSimilarResponseReadyArtifacts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchSimilarResponseReadyArtifacts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClusterPath) {
		toSerialize["clusterPath"] = o.ClusterPath
	}
	return toSerialize, nil
}

type NullableFetchSimilarResponseReadyArtifacts struct {
	value *FetchSimilarResponseReadyArtifacts
	isSet bool
}

func (v NullableFetchSimilarResponseReadyArtifacts) Get() *FetchSimilarResponseReadyArtifacts {
	return v.value
}

func (v *NullableFetchSimilarResponseReadyArtifacts) Set(val *FetchSimilarResponseReadyArtifacts) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchSimilarResponseReadyArtifacts) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchSimilarResponseReadyArtifacts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchSimilarResponseReadyArtifacts(val *FetchSimilarResponseReadyArtifacts) *NullableFetchSimilarResponseReadyArtifacts {
	return &NullableFetchSimilarResponseReadyArtifacts{value: val, isSet: true}
}

func (v NullableFetchSimilarResponseReadyArtifacts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchSimilarResponseReadyArtifacts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}