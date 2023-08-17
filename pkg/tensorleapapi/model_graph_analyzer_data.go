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

// checks if the GraphAnalyzerData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphAnalyzerData{}

// GraphAnalyzerData struct for GraphAnalyzerData
type GraphAnalyzerData struct {
	// Construct a type with a set of properties K of type T
	Nodes map[string]interface{} `json:"nodes"`
	Hash *string `json:"hash,omitempty"`
	GeneralError *string `json:"general_error,omitempty"`
}

// NewGraphAnalyzerData instantiates a new GraphAnalyzerData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphAnalyzerData(nodes map[string]interface{}) *GraphAnalyzerData {
	this := GraphAnalyzerData{}
	this.Nodes = nodes
	return &this
}

// NewGraphAnalyzerDataWithDefaults instantiates a new GraphAnalyzerData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphAnalyzerDataWithDefaults() *GraphAnalyzerData {
	this := GraphAnalyzerData{}
	return &this
}

// GetNodes returns the Nodes field value
func (o *GraphAnalyzerData) GetNodes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Nodes
}

// GetNodesOk returns a tuple with the Nodes field value
// and a boolean to check if the value has been set.
func (o *GraphAnalyzerData) GetNodesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Nodes, true
}

// SetNodes sets field value
func (o *GraphAnalyzerData) SetNodes(v map[string]interface{}) {
	o.Nodes = v
}

// GetHash returns the Hash field value if set, zero value otherwise.
func (o *GraphAnalyzerData) GetHash() string {
	if o == nil || IsNil(o.Hash) {
		var ret string
		return ret
	}
	return *o.Hash
}

// GetHashOk returns a tuple with the Hash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphAnalyzerData) GetHashOk() (*string, bool) {
	if o == nil || IsNil(o.Hash) {
		return nil, false
	}
	return o.Hash, true
}

// HasHash returns a boolean if a field has been set.
func (o *GraphAnalyzerData) HasHash() bool {
	if o != nil && !IsNil(o.Hash) {
		return true
	}

	return false
}

// SetHash gets a reference to the given string and assigns it to the Hash field.
func (o *GraphAnalyzerData) SetHash(v string) {
	o.Hash = &v
}

// GetGeneralError returns the GeneralError field value if set, zero value otherwise.
func (o *GraphAnalyzerData) GetGeneralError() string {
	if o == nil || IsNil(o.GeneralError) {
		var ret string
		return ret
	}
	return *o.GeneralError
}

// GetGeneralErrorOk returns a tuple with the GeneralError field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphAnalyzerData) GetGeneralErrorOk() (*string, bool) {
	if o == nil || IsNil(o.GeneralError) {
		return nil, false
	}
	return o.GeneralError, true
}

// HasGeneralError returns a boolean if a field has been set.
func (o *GraphAnalyzerData) HasGeneralError() bool {
	if o != nil && !IsNil(o.GeneralError) {
		return true
	}

	return false
}

// SetGeneralError gets a reference to the given string and assigns it to the GeneralError field.
func (o *GraphAnalyzerData) SetGeneralError(v string) {
	o.GeneralError = &v
}

func (o GraphAnalyzerData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphAnalyzerData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodes"] = o.Nodes
	if !IsNil(o.Hash) {
		toSerialize["hash"] = o.Hash
	}
	if !IsNil(o.GeneralError) {
		toSerialize["general_error"] = o.GeneralError
	}
	return toSerialize, nil
}

type NullableGraphAnalyzerData struct {
	value *GraphAnalyzerData
	isSet bool
}

func (v NullableGraphAnalyzerData) Get() *GraphAnalyzerData {
	return v.value
}

func (v *NullableGraphAnalyzerData) Set(val *GraphAnalyzerData) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphAnalyzerData) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphAnalyzerData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphAnalyzerData(val *GraphAnalyzerData) *NullableGraphAnalyzerData {
	return &NullableGraphAnalyzerData{value: val, isSet: true}
}

func (v NullableGraphAnalyzerData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphAnalyzerData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


