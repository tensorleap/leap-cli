/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.464
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SampleAnalysisDashlet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SampleAnalysisDashlet{}

// SampleAnalysisDashlet struct for SampleAnalysisDashlet
type SampleAnalysisDashlet struct {
	Cid           string          `json:"cid"`
	Layout        SizedLayout     `json:"layout"`
	Type          string          `json:"type"`
	Name          string          `json:"name"`
	CollectionIds []string        `json:"collectionIds"`
	Sample        *SampleIdentity `json:"sample,omitempty"`
	AssetId       *SampleAssetId  `json:"assetId,omitempty"`
}

// NewSampleAnalysisDashlet instantiates a new SampleAnalysisDashlet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSampleAnalysisDashlet(cid string, layout SizedLayout, type_ string, name string, collectionIds []string) *SampleAnalysisDashlet {
	this := SampleAnalysisDashlet{}
	this.Cid = cid
	this.Layout = layout
	this.Type = type_
	this.Name = name
	this.CollectionIds = collectionIds
	return &this
}

// NewSampleAnalysisDashletWithDefaults instantiates a new SampleAnalysisDashlet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSampleAnalysisDashletWithDefaults() *SampleAnalysisDashlet {
	this := SampleAnalysisDashlet{}
	return &this
}

// GetCid returns the Cid field value
func (o *SampleAnalysisDashlet) GetCid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cid
}

// GetCidOk returns a tuple with the Cid field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetCidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cid, true
}

// SetCid sets field value
func (o *SampleAnalysisDashlet) SetCid(v string) {
	o.Cid = v
}

// GetLayout returns the Layout field value
func (o *SampleAnalysisDashlet) GetLayout() SizedLayout {
	if o == nil {
		var ret SizedLayout
		return ret
	}

	return o.Layout
}

// GetLayoutOk returns a tuple with the Layout field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetLayoutOk() (*SizedLayout, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Layout, true
}

// SetLayout sets field value
func (o *SampleAnalysisDashlet) SetLayout(v SizedLayout) {
	o.Layout = v
}

// GetType returns the Type field value
func (o *SampleAnalysisDashlet) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SampleAnalysisDashlet) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SampleAnalysisDashlet) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SampleAnalysisDashlet) SetName(v string) {
	o.Name = v
}

// GetCollectionIds returns the CollectionIds field value
func (o *SampleAnalysisDashlet) GetCollectionIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.CollectionIds
}

// GetCollectionIdsOk returns a tuple with the CollectionIds field value
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetCollectionIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CollectionIds, true
}

// SetCollectionIds sets field value
func (o *SampleAnalysisDashlet) SetCollectionIds(v []string) {
	o.CollectionIds = v
}

// GetSample returns the Sample field value if set, zero value otherwise.
func (o *SampleAnalysisDashlet) GetSample() SampleIdentity {
	if o == nil || IsNil(o.Sample) {
		var ret SampleIdentity
		return ret
	}
	return *o.Sample
}

// GetSampleOk returns a tuple with the Sample field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetSampleOk() (*SampleIdentity, bool) {
	if o == nil || IsNil(o.Sample) {
		return nil, false
	}
	return o.Sample, true
}

// HasSample returns a boolean if a field has been set.
func (o *SampleAnalysisDashlet) HasSample() bool {
	if o != nil && !IsNil(o.Sample) {
		return true
	}

	return false
}

// SetSample gets a reference to the given SampleIdentity and assigns it to the Sample field.
func (o *SampleAnalysisDashlet) SetSample(v SampleIdentity) {
	o.Sample = &v
}

// GetAssetId returns the AssetId field value if set, zero value otherwise.
func (o *SampleAnalysisDashlet) GetAssetId() SampleAssetId {
	if o == nil || IsNil(o.AssetId) {
		var ret SampleAssetId
		return ret
	}
	return *o.AssetId
}

// GetAssetIdOk returns a tuple with the AssetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SampleAnalysisDashlet) GetAssetIdOk() (*SampleAssetId, bool) {
	if o == nil || IsNil(o.AssetId) {
		return nil, false
	}
	return o.AssetId, true
}

// HasAssetId returns a boolean if a field has been set.
func (o *SampleAnalysisDashlet) HasAssetId() bool {
	if o != nil && !IsNil(o.AssetId) {
		return true
	}

	return false
}

// SetAssetId gets a reference to the given SampleAssetId and assigns it to the AssetId field.
func (o *SampleAnalysisDashlet) SetAssetId(v SampleAssetId) {
	o.AssetId = &v
}

func (o SampleAnalysisDashlet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SampleAnalysisDashlet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cid"] = o.Cid
	toSerialize["layout"] = o.Layout
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	toSerialize["collectionIds"] = o.CollectionIds
	if !IsNil(o.Sample) {
		toSerialize["sample"] = o.Sample
	}
	if !IsNil(o.AssetId) {
		toSerialize["assetId"] = o.AssetId
	}
	return toSerialize, nil
}

type NullableSampleAnalysisDashlet struct {
	value *SampleAnalysisDashlet
	isSet bool
}

func (v NullableSampleAnalysisDashlet) Get() *SampleAnalysisDashlet {
	return v.value
}

func (v *NullableSampleAnalysisDashlet) Set(val *SampleAnalysisDashlet) {
	v.value = val
	v.isSet = true
}

func (v NullableSampleAnalysisDashlet) IsSet() bool {
	return v.isSet
}

func (v *NullableSampleAnalysisDashlet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSampleAnalysisDashlet(val *SampleAnalysisDashlet) *NullableSampleAnalysisDashlet {
	return &NullableSampleAnalysisDashlet{value: val, isSet: true}
}

func (v NullableSampleAnalysisDashlet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSampleAnalysisDashlet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
