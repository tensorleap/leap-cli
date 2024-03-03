/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.467
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the ClusterInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ClusterInsight{}

// ClusterInsight struct for ClusterInsight
type ClusterInsight struct {
	Type               ScatterInsightType         `json:"type"`
	Filter             ScatterFilter              `json:"filter"`
	MutualInfoElements []MutualInformationElement `json:"mutual_info_elements,omitempty"`
	BlobPath           *string                    `json:"blob_path,omitempty"`
	NSamples           float64                    `json:"n_samples"`
	AvgMetric          float64                    `json:"avg_metric"`
	MetricName         string                     `json:"metric_name"`
	// Construct a type with a set of properties K of type T
	Configuration map[string]interface{} `json:"configuration,omitempty"`
}

// NewClusterInsight instantiates a new ClusterInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterInsight(type_ ScatterInsightType, filter ScatterFilter, nSamples float64, avgMetric float64, metricName string) *ClusterInsight {
	this := ClusterInsight{}
	this.Type = type_
	this.Filter = filter
	this.NSamples = nSamples
	this.AvgMetric = avgMetric
	this.MetricName = metricName
	return &this
}

// NewClusterInsightWithDefaults instantiates a new ClusterInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterInsightWithDefaults() *ClusterInsight {
	this := ClusterInsight{}
	return &this
}

// GetType returns the Type field value
func (o *ClusterInsight) GetType() ScatterInsightType {
	if o == nil {
		var ret ScatterInsightType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetTypeOk() (*ScatterInsightType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ClusterInsight) SetType(v ScatterInsightType) {
	o.Type = v
}

// GetFilter returns the Filter field value
func (o *ClusterInsight) GetFilter() ScatterFilter {
	if o == nil {
		var ret ScatterFilter
		return ret
	}

	return o.Filter
}

// GetFilterOk returns a tuple with the Filter field value
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetFilterOk() (*ScatterFilter, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Filter, true
}

// SetFilter sets field value
func (o *ClusterInsight) SetFilter(v ScatterFilter) {
	o.Filter = v
}

// GetMutualInfoElements returns the MutualInfoElements field value if set, zero value otherwise.
func (o *ClusterInsight) GetMutualInfoElements() []MutualInformationElement {
	if o == nil || IsNil(o.MutualInfoElements) {
		var ret []MutualInformationElement
		return ret
	}
	return o.MutualInfoElements
}

// GetMutualInfoElementsOk returns a tuple with the MutualInfoElements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetMutualInfoElementsOk() ([]MutualInformationElement, bool) {
	if o == nil || IsNil(o.MutualInfoElements) {
		return nil, false
	}
	return o.MutualInfoElements, true
}

// HasMutualInfoElements returns a boolean if a field has been set.
func (o *ClusterInsight) HasMutualInfoElements() bool {
	if o != nil && !IsNil(o.MutualInfoElements) {
		return true
	}

	return false
}

// SetMutualInfoElements gets a reference to the given []MutualInformationElement and assigns it to the MutualInfoElements field.
func (o *ClusterInsight) SetMutualInfoElements(v []MutualInformationElement) {
	o.MutualInfoElements = v
}

// GetBlobPath returns the BlobPath field value if set, zero value otherwise.
func (o *ClusterInsight) GetBlobPath() string {
	if o == nil || IsNil(o.BlobPath) {
		var ret string
		return ret
	}
	return *o.BlobPath
}

// GetBlobPathOk returns a tuple with the BlobPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetBlobPathOk() (*string, bool) {
	if o == nil || IsNil(o.BlobPath) {
		return nil, false
	}
	return o.BlobPath, true
}

// HasBlobPath returns a boolean if a field has been set.
func (o *ClusterInsight) HasBlobPath() bool {
	if o != nil && !IsNil(o.BlobPath) {
		return true
	}

	return false
}

// SetBlobPath gets a reference to the given string and assigns it to the BlobPath field.
func (o *ClusterInsight) SetBlobPath(v string) {
	o.BlobPath = &v
}

// GetNSamples returns the NSamples field value
func (o *ClusterInsight) GetNSamples() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.NSamples
}

// GetNSamplesOk returns a tuple with the NSamples field value
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetNSamplesOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NSamples, true
}

// SetNSamples sets field value
func (o *ClusterInsight) SetNSamples(v float64) {
	o.NSamples = v
}

// GetAvgMetric returns the AvgMetric field value
func (o *ClusterInsight) GetAvgMetric() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.AvgMetric
}

// GetAvgMetricOk returns a tuple with the AvgMetric field value
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetAvgMetricOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvgMetric, true
}

// SetAvgMetric sets field value
func (o *ClusterInsight) SetAvgMetric(v float64) {
	o.AvgMetric = v
}

// GetMetricName returns the MetricName field value
func (o *ClusterInsight) GetMetricName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetricName
}

// GetMetricNameOk returns a tuple with the MetricName field value
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetMetricNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetricName, true
}

// SetMetricName sets field value
func (o *ClusterInsight) SetMetricName(v string) {
	o.MetricName = v
}

// GetConfiguration returns the Configuration field value if set, zero value otherwise.
func (o *ClusterInsight) GetConfiguration() map[string]interface{} {
	if o == nil || IsNil(o.Configuration) {
		var ret map[string]interface{}
		return ret
	}
	return o.Configuration
}

// GetConfigurationOk returns a tuple with the Configuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterInsight) GetConfigurationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Configuration) {
		return map[string]interface{}{}, false
	}
	return o.Configuration, true
}

// HasConfiguration returns a boolean if a field has been set.
func (o *ClusterInsight) HasConfiguration() bool {
	if o != nil && !IsNil(o.Configuration) {
		return true
	}

	return false
}

// SetConfiguration gets a reference to the given map[string]interface{} and assigns it to the Configuration field.
func (o *ClusterInsight) SetConfiguration(v map[string]interface{}) {
	o.Configuration = v
}

func (o ClusterInsight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ClusterInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["filter"] = o.Filter
	if !IsNil(o.MutualInfoElements) {
		toSerialize["mutual_info_elements"] = o.MutualInfoElements
	}
	if !IsNil(o.BlobPath) {
		toSerialize["blob_path"] = o.BlobPath
	}
	toSerialize["n_samples"] = o.NSamples
	toSerialize["avg_metric"] = o.AvgMetric
	toSerialize["metric_name"] = o.MetricName
	if !IsNil(o.Configuration) {
		toSerialize["configuration"] = o.Configuration
	}
	return toSerialize, nil
}

type NullableClusterInsight struct {
	value *ClusterInsight
	isSet bool
}

func (v NullableClusterInsight) Get() *ClusterInsight {
	return v.value
}

func (v *NullableClusterInsight) Set(val *ClusterInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterInsight(val *ClusterInsight) *NullableClusterInsight {
	return &NullableClusterInsight{value: val, isSet: true}
}

func (v NullableClusterInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
