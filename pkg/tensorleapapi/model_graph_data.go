/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.592
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the GraphData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GraphData{}

// GraphData struct for GraphData
type GraphData struct {
	Body    [][]float64  `json:"body"`
	Heatmap *Heatmap     `json:"heatmap,omitempty"`
	Type    DataTypeEnum `json:"type"`
	XLabel  *string      `json:"x_label,omitempty"`
	YLabel  *string      `json:"y_label,omitempty"`
	XRange  []float64    `json:"x_range,omitempty"`
}

// NewGraphData instantiates a new GraphData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGraphData(body [][]float64, type_ DataTypeEnum) *GraphData {
	this := GraphData{}
	this.Body = body
	this.Type = type_
	return &this
}

// NewGraphDataWithDefaults instantiates a new GraphData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGraphDataWithDefaults() *GraphData {
	this := GraphData{}
	return &this
}

// GetBody returns the Body field value
func (o *GraphData) GetBody() [][]float64 {
	if o == nil {
		var ret [][]float64
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *GraphData) GetBodyOk() ([][]float64, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *GraphData) SetBody(v [][]float64) {
	o.Body = v
}

// GetHeatmap returns the Heatmap field value if set, zero value otherwise.
func (o *GraphData) GetHeatmap() Heatmap {
	if o == nil || IsNil(o.Heatmap) {
		var ret Heatmap
		return ret
	}
	return *o.Heatmap
}

// GetHeatmapOk returns a tuple with the Heatmap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphData) GetHeatmapOk() (*Heatmap, bool) {
	if o == nil || IsNil(o.Heatmap) {
		return nil, false
	}
	return o.Heatmap, true
}

// HasHeatmap returns a boolean if a field has been set.
func (o *GraphData) HasHeatmap() bool {
	if o != nil && !IsNil(o.Heatmap) {
		return true
	}

	return false
}

// SetHeatmap gets a reference to the given Heatmap and assigns it to the Heatmap field.
func (o *GraphData) SetHeatmap(v Heatmap) {
	o.Heatmap = &v
}

// GetType returns the Type field value
func (o *GraphData) GetType() DataTypeEnum {
	if o == nil {
		var ret DataTypeEnum
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GraphData) GetTypeOk() (*DataTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GraphData) SetType(v DataTypeEnum) {
	o.Type = v
}

// GetXLabel returns the XLabel field value if set, zero value otherwise.
func (o *GraphData) GetXLabel() string {
	if o == nil || IsNil(o.XLabel) {
		var ret string
		return ret
	}
	return *o.XLabel
}

// GetXLabelOk returns a tuple with the XLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphData) GetXLabelOk() (*string, bool) {
	if o == nil || IsNil(o.XLabel) {
		return nil, false
	}
	return o.XLabel, true
}

// HasXLabel returns a boolean if a field has been set.
func (o *GraphData) HasXLabel() bool {
	if o != nil && !IsNil(o.XLabel) {
		return true
	}

	return false
}

// SetXLabel gets a reference to the given string and assigns it to the XLabel field.
func (o *GraphData) SetXLabel(v string) {
	o.XLabel = &v
}

// GetYLabel returns the YLabel field value if set, zero value otherwise.
func (o *GraphData) GetYLabel() string {
	if o == nil || IsNil(o.YLabel) {
		var ret string
		return ret
	}
	return *o.YLabel
}

// GetYLabelOk returns a tuple with the YLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphData) GetYLabelOk() (*string, bool) {
	if o == nil || IsNil(o.YLabel) {
		return nil, false
	}
	return o.YLabel, true
}

// HasYLabel returns a boolean if a field has been set.
func (o *GraphData) HasYLabel() bool {
	if o != nil && !IsNil(o.YLabel) {
		return true
	}

	return false
}

// SetYLabel gets a reference to the given string and assigns it to the YLabel field.
func (o *GraphData) SetYLabel(v string) {
	o.YLabel = &v
}

// GetXRange returns the XRange field value if set, zero value otherwise.
func (o *GraphData) GetXRange() []float64 {
	if o == nil || IsNil(o.XRange) {
		var ret []float64
		return ret
	}
	return o.XRange
}

// GetXRangeOk returns a tuple with the XRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GraphData) GetXRangeOk() ([]float64, bool) {
	if o == nil || IsNil(o.XRange) {
		return nil, false
	}
	return o.XRange, true
}

// HasXRange returns a boolean if a field has been set.
func (o *GraphData) HasXRange() bool {
	if o != nil && !IsNil(o.XRange) {
		return true
	}

	return false
}

// SetXRange gets a reference to the given []float64 and assigns it to the XRange field.
func (o *GraphData) SetXRange(v []float64) {
	o.XRange = v
}

func (o GraphData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GraphData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	if !IsNil(o.Heatmap) {
		toSerialize["heatmap"] = o.Heatmap
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.XLabel) {
		toSerialize["x_label"] = o.XLabel
	}
	if !IsNil(o.YLabel) {
		toSerialize["y_label"] = o.YLabel
	}
	if !IsNil(o.XRange) {
		toSerialize["x_range"] = o.XRange
	}
	return toSerialize, nil
}

type NullableGraphData struct {
	value *GraphData
	isSet bool
}

func (v NullableGraphData) Get() *GraphData {
	return v.value
}

func (v *NullableGraphData) Set(val *GraphData) {
	v.value = val
	v.isSet = true
}

func (v NullableGraphData) IsSet() bool {
	return v.isSet
}

func (v *NullableGraphData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGraphData(val *GraphData) *NullableGraphData {
	return &NullableGraphData{value: val, isSet: true}
}

func (v NullableGraphData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGraphData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
