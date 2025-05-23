/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.618
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// VisData struct for VisData
type VisData struct {
	BBoxImageData     *BBoxImageData
	GraphData         *GraphData
	HorizontalBarData *HorizontalBarData
	ImageData         *ImageData
	ImageHeatmapData  *ImageHeatmapData
	MaskImageData     *MaskImageData
	MaskTextData      *MaskTextData
	TextData          *TextData
	VideoData         *VideoData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VisData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BBoxImageData
	err = json.Unmarshal(data, &dst.BBoxImageData)
	if err == nil {
		jsonBBoxImageData, _ := json.Marshal(dst.BBoxImageData)
		if string(jsonBBoxImageData) == "{}" { // empty struct
			dst.BBoxImageData = nil
		} else {
			return nil // data stored in dst.BBoxImageData, return on the first match
		}
	} else {
		dst.BBoxImageData = nil
	}

	// try to unmarshal JSON data into GraphData
	err = json.Unmarshal(data, &dst.GraphData)
	if err == nil {
		jsonGraphData, _ := json.Marshal(dst.GraphData)
		if string(jsonGraphData) == "{}" { // empty struct
			dst.GraphData = nil
		} else {
			return nil // data stored in dst.GraphData, return on the first match
		}
	} else {
		dst.GraphData = nil
	}

	// try to unmarshal JSON data into HorizontalBarData
	err = json.Unmarshal(data, &dst.HorizontalBarData)
	if err == nil {
		jsonHorizontalBarData, _ := json.Marshal(dst.HorizontalBarData)
		if string(jsonHorizontalBarData) == "{}" { // empty struct
			dst.HorizontalBarData = nil
		} else {
			return nil // data stored in dst.HorizontalBarData, return on the first match
		}
	} else {
		dst.HorizontalBarData = nil
	}

	// try to unmarshal JSON data into ImageData
	err = json.Unmarshal(data, &dst.ImageData)
	if err == nil {
		jsonImageData, _ := json.Marshal(dst.ImageData)
		if string(jsonImageData) == "{}" { // empty struct
			dst.ImageData = nil
		} else {
			return nil // data stored in dst.ImageData, return on the first match
		}
	} else {
		dst.ImageData = nil
	}

	// try to unmarshal JSON data into ImageHeatmapData
	err = json.Unmarshal(data, &dst.ImageHeatmapData)
	if err == nil {
		jsonImageHeatmapData, _ := json.Marshal(dst.ImageHeatmapData)
		if string(jsonImageHeatmapData) == "{}" { // empty struct
			dst.ImageHeatmapData = nil
		} else {
			return nil // data stored in dst.ImageHeatmapData, return on the first match
		}
	} else {
		dst.ImageHeatmapData = nil
	}

	// try to unmarshal JSON data into MaskImageData
	err = json.Unmarshal(data, &dst.MaskImageData)
	if err == nil {
		jsonMaskImageData, _ := json.Marshal(dst.MaskImageData)
		if string(jsonMaskImageData) == "{}" { // empty struct
			dst.MaskImageData = nil
		} else {
			return nil // data stored in dst.MaskImageData, return on the first match
		}
	} else {
		dst.MaskImageData = nil
	}

	// try to unmarshal JSON data into MaskTextData
	err = json.Unmarshal(data, &dst.MaskTextData)
	if err == nil {
		jsonMaskTextData, _ := json.Marshal(dst.MaskTextData)
		if string(jsonMaskTextData) == "{}" { // empty struct
			dst.MaskTextData = nil
		} else {
			return nil // data stored in dst.MaskTextData, return on the first match
		}
	} else {
		dst.MaskTextData = nil
	}

	// try to unmarshal JSON data into TextData
	err = json.Unmarshal(data, &dst.TextData)
	if err == nil {
		jsonTextData, _ := json.Marshal(dst.TextData)
		if string(jsonTextData) == "{}" { // empty struct
			dst.TextData = nil
		} else {
			return nil // data stored in dst.TextData, return on the first match
		}
	} else {
		dst.TextData = nil
	}

	// try to unmarshal JSON data into VideoData
	err = json.Unmarshal(data, &dst.VideoData)
	if err == nil {
		jsonVideoData, _ := json.Marshal(dst.VideoData)
		if string(jsonVideoData) == "{}" { // empty struct
			dst.VideoData = nil
		} else {
			return nil // data stored in dst.VideoData, return on the first match
		}
	} else {
		dst.VideoData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(VisData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VisData) MarshalJSON() ([]byte, error) {
	if src.BBoxImageData != nil {
		return json.Marshal(&src.BBoxImageData)
	}

	if src.GraphData != nil {
		return json.Marshal(&src.GraphData)
	}

	if src.HorizontalBarData != nil {
		return json.Marshal(&src.HorizontalBarData)
	}

	if src.ImageData != nil {
		return json.Marshal(&src.ImageData)
	}

	if src.ImageHeatmapData != nil {
		return json.Marshal(&src.ImageHeatmapData)
	}

	if src.MaskImageData != nil {
		return json.Marshal(&src.MaskImageData)
	}

	if src.MaskTextData != nil {
		return json.Marshal(&src.MaskTextData)
	}

	if src.TextData != nil {
		return json.Marshal(&src.TextData)
	}

	if src.VideoData != nil {
		return json.Marshal(&src.VideoData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVisData struct {
	value *VisData
	isSet bool
}

func (v NullableVisData) Get() *VisData {
	return v.value
}

func (v *NullableVisData) Set(val *VisData) {
	v.value = val
	v.isSet = true
}

func (v NullableVisData) IsSet() bool {
	return v.isSet
}

func (v *NullableVisData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVisData(val *VisData) *NullableVisData {
	return &NullableVisData{value: val, isSet: true}
}

func (v NullableVisData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVisData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
