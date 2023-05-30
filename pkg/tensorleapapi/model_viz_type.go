/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// VizType struct for VizType
type VizType struct {
	BBoxImageViz      *BBoxImageViz
	GraphViz          *GraphViz
	HorizontalBarViz  *HorizontalBarViz
	ImageViz          *ImageViz
	SampleAnalysisViz *SampleAnalysisViz
	ScatterViz        *ScatterViz
	TextViz           *TextViz
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VizType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BBoxImageViz
	err = json.Unmarshal(data, &dst.BBoxImageViz)
	if err == nil {
		jsonBBoxImageViz, _ := json.Marshal(dst.BBoxImageViz)
		if string(jsonBBoxImageViz) == "{}" { // empty struct
			dst.BBoxImageViz = nil
		} else {
			return nil // data stored in dst.BBoxImageViz, return on the first match
		}
	} else {
		dst.BBoxImageViz = nil
	}

	// try to unmarshal JSON data into GraphViz
	err = json.Unmarshal(data, &dst.GraphViz)
	if err == nil {
		jsonGraphViz, _ := json.Marshal(dst.GraphViz)
		if string(jsonGraphViz) == "{}" { // empty struct
			dst.GraphViz = nil
		} else {
			return nil // data stored in dst.GraphViz, return on the first match
		}
	} else {
		dst.GraphViz = nil
	}

	// try to unmarshal JSON data into HorizontalBarViz
	err = json.Unmarshal(data, &dst.HorizontalBarViz)
	if err == nil {
		jsonHorizontalBarViz, _ := json.Marshal(dst.HorizontalBarViz)
		if string(jsonHorizontalBarViz) == "{}" { // empty struct
			dst.HorizontalBarViz = nil
		} else {
			return nil // data stored in dst.HorizontalBarViz, return on the first match
		}
	} else {
		dst.HorizontalBarViz = nil
	}

	// try to unmarshal JSON data into ImageViz
	err = json.Unmarshal(data, &dst.ImageViz)
	if err == nil {
		jsonImageViz, _ := json.Marshal(dst.ImageViz)
		if string(jsonImageViz) == "{}" { // empty struct
			dst.ImageViz = nil
		} else {
			return nil // data stored in dst.ImageViz, return on the first match
		}
	} else {
		dst.ImageViz = nil
	}

	// try to unmarshal JSON data into SampleAnalysisViz
	err = json.Unmarshal(data, &dst.SampleAnalysisViz)
	if err == nil {
		jsonSampleAnalysisViz, _ := json.Marshal(dst.SampleAnalysisViz)
		if string(jsonSampleAnalysisViz) == "{}" { // empty struct
			dst.SampleAnalysisViz = nil
		} else {
			return nil // data stored in dst.SampleAnalysisViz, return on the first match
		}
	} else {
		dst.SampleAnalysisViz = nil
	}

	// try to unmarshal JSON data into ScatterViz
	err = json.Unmarshal(data, &dst.ScatterViz)
	if err == nil {
		jsonScatterViz, _ := json.Marshal(dst.ScatterViz)
		if string(jsonScatterViz) == "{}" { // empty struct
			dst.ScatterViz = nil
		} else {
			return nil // data stored in dst.ScatterViz, return on the first match
		}
	} else {
		dst.ScatterViz = nil
	}

	// try to unmarshal JSON data into TextViz
	err = json.Unmarshal(data, &dst.TextViz)
	if err == nil {
		jsonTextViz, _ := json.Marshal(dst.TextViz)
		if string(jsonTextViz) == "{}" { // empty struct
			dst.TextViz = nil
		} else {
			return nil // data stored in dst.TextViz, return on the first match
		}
	} else {
		dst.TextViz = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(VizType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VizType) MarshalJSON() ([]byte, error) {
	if src.BBoxImageViz != nil {
		return json.Marshal(&src.BBoxImageViz)
	}

	if src.GraphViz != nil {
		return json.Marshal(&src.GraphViz)
	}

	if src.HorizontalBarViz != nil {
		return json.Marshal(&src.HorizontalBarViz)
	}

	if src.ImageViz != nil {
		return json.Marshal(&src.ImageViz)
	}

	if src.SampleAnalysisViz != nil {
		return json.Marshal(&src.SampleAnalysisViz)
	}

	if src.ScatterViz != nil {
		return json.Marshal(&src.ScatterViz)
	}

	if src.TextViz != nil {
		return json.Marshal(&src.TextViz)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVizType struct {
	value *VizType
	isSet bool
}

func (v NullableVizType) Get() *VizType {
	return v.value
}

func (v *NullableVizType) Set(val *VizType) {
	v.value = val
	v.isSet = true
}

func (v NullableVizType) IsSet() bool {
	return v.isSet
}

func (v *NullableVizType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVizType(val *VizType) *NullableVizType {
	return &NullableVizType{value: val, isSet: true}
}

func (v NullableVizType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVizType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
