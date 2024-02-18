/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.456
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the BucketAggregation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketAggregation{}

// BucketAggregation struct for BucketAggregation
type BucketAggregation struct {
	XField               string               `json:"xField"`
	DataDistributionType DataDistributionType `json:"dataDistributionType"`
	XAxisSizeInterval    float64              `json:"xAxisSizeInterval"`
	OrderParams          string               `json:"orderParams"`
}

// NewBucketAggregation instantiates a new BucketAggregation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketAggregation(xField string, dataDistributionType DataDistributionType, xAxisSizeInterval float64, orderParams string) *BucketAggregation {
	this := BucketAggregation{}
	this.XField = xField
	this.DataDistributionType = dataDistributionType
	this.XAxisSizeInterval = xAxisSizeInterval
	this.OrderParams = orderParams
	return &this
}

// NewBucketAggregationWithDefaults instantiates a new BucketAggregation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketAggregationWithDefaults() *BucketAggregation {
	this := BucketAggregation{}
	return &this
}

// GetXField returns the XField field value
func (o *BucketAggregation) GetXField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.XField
}

// GetXFieldOk returns a tuple with the XField field value
// and a boolean to check if the value has been set.
func (o *BucketAggregation) GetXFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XField, true
}

// SetXField sets field value
func (o *BucketAggregation) SetXField(v string) {
	o.XField = v
}

// GetDataDistributionType returns the DataDistributionType field value
func (o *BucketAggregation) GetDataDistributionType() DataDistributionType {
	if o == nil {
		var ret DataDistributionType
		return ret
	}

	return o.DataDistributionType
}

// GetDataDistributionTypeOk returns a tuple with the DataDistributionType field value
// and a boolean to check if the value has been set.
func (o *BucketAggregation) GetDataDistributionTypeOk() (*DataDistributionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DataDistributionType, true
}

// SetDataDistributionType sets field value
func (o *BucketAggregation) SetDataDistributionType(v DataDistributionType) {
	o.DataDistributionType = v
}

// GetXAxisSizeInterval returns the XAxisSizeInterval field value
func (o *BucketAggregation) GetXAxisSizeInterval() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.XAxisSizeInterval
}

// GetXAxisSizeIntervalOk returns a tuple with the XAxisSizeInterval field value
// and a boolean to check if the value has been set.
func (o *BucketAggregation) GetXAxisSizeIntervalOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.XAxisSizeInterval, true
}

// SetXAxisSizeInterval sets field value
func (o *BucketAggregation) SetXAxisSizeInterval(v float64) {
	o.XAxisSizeInterval = v
}

// GetOrderParams returns the OrderParams field value
func (o *BucketAggregation) GetOrderParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrderParams
}

// GetOrderParamsOk returns a tuple with the OrderParams field value
// and a boolean to check if the value has been set.
func (o *BucketAggregation) GetOrderParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderParams, true
}

// SetOrderParams sets field value
func (o *BucketAggregation) SetOrderParams(v string) {
	o.OrderParams = v
}

func (o BucketAggregation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketAggregation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["xField"] = o.XField
	toSerialize["dataDistributionType"] = o.DataDistributionType
	toSerialize["xAxisSizeInterval"] = o.XAxisSizeInterval
	toSerialize["orderParams"] = o.OrderParams
	return toSerialize, nil
}

type NullableBucketAggregation struct {
	value *BucketAggregation
	isSet bool
}

func (v NullableBucketAggregation) Get() *BucketAggregation {
	return v.value
}

func (v *NullableBucketAggregation) Set(val *BucketAggregation) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketAggregation) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketAggregation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketAggregation(val *BucketAggregation) *NullableBucketAggregation {
	return &NullableBucketAggregation{value: val, isSet: true}
}

func (v NullableBucketAggregation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketAggregation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
