/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.404
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// ScatterLabel struct for ScatterLabel
type ScatterLabel struct {
	BBoxImageScatterLabel     *BBoxImageScatterLabel
	GraphScatterLabel         *GraphScatterLabel
	HorizontalBarScatterLabel *HorizontalBarScatterLabel
	ImageScatterLabel         *ImageScatterLabel
	MaskImageScatterLabel     *MaskImageScatterLabel
	TextMaskScatterLabel      *TextMaskScatterLabel
	TextScatterLabel          *TextScatterLabel
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScatterLabel) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BBoxImageScatterLabel
	err = json.Unmarshal(data, &dst.BBoxImageScatterLabel)
	if err == nil {
		jsonBBoxImageScatterLabel, _ := json.Marshal(dst.BBoxImageScatterLabel)
		if string(jsonBBoxImageScatterLabel) == "{}" { // empty struct
			dst.BBoxImageScatterLabel = nil
		} else {
			return nil // data stored in dst.BBoxImageScatterLabel, return on the first match
		}
	} else {
		dst.BBoxImageScatterLabel = nil
	}

	// try to unmarshal JSON data into GraphScatterLabel
	err = json.Unmarshal(data, &dst.GraphScatterLabel)
	if err == nil {
		jsonGraphScatterLabel, _ := json.Marshal(dst.GraphScatterLabel)
		if string(jsonGraphScatterLabel) == "{}" { // empty struct
			dst.GraphScatterLabel = nil
		} else {
			return nil // data stored in dst.GraphScatterLabel, return on the first match
		}
	} else {
		dst.GraphScatterLabel = nil
	}

	// try to unmarshal JSON data into HorizontalBarScatterLabel
	err = json.Unmarshal(data, &dst.HorizontalBarScatterLabel)
	if err == nil {
		jsonHorizontalBarScatterLabel, _ := json.Marshal(dst.HorizontalBarScatterLabel)
		if string(jsonHorizontalBarScatterLabel) == "{}" { // empty struct
			dst.HorizontalBarScatterLabel = nil
		} else {
			return nil // data stored in dst.HorizontalBarScatterLabel, return on the first match
		}
	} else {
		dst.HorizontalBarScatterLabel = nil
	}

	// try to unmarshal JSON data into ImageScatterLabel
	err = json.Unmarshal(data, &dst.ImageScatterLabel)
	if err == nil {
		jsonImageScatterLabel, _ := json.Marshal(dst.ImageScatterLabel)
		if string(jsonImageScatterLabel) == "{}" { // empty struct
			dst.ImageScatterLabel = nil
		} else {
			return nil // data stored in dst.ImageScatterLabel, return on the first match
		}
	} else {
		dst.ImageScatterLabel = nil
	}

	// try to unmarshal JSON data into MaskImageScatterLabel
	err = json.Unmarshal(data, &dst.MaskImageScatterLabel)
	if err == nil {
		jsonMaskImageScatterLabel, _ := json.Marshal(dst.MaskImageScatterLabel)
		if string(jsonMaskImageScatterLabel) == "{}" { // empty struct
			dst.MaskImageScatterLabel = nil
		} else {
			return nil // data stored in dst.MaskImageScatterLabel, return on the first match
		}
	} else {
		dst.MaskImageScatterLabel = nil
	}

	// try to unmarshal JSON data into TextMaskScatterLabel
	err = json.Unmarshal(data, &dst.TextMaskScatterLabel)
	if err == nil {
		jsonTextMaskScatterLabel, _ := json.Marshal(dst.TextMaskScatterLabel)
		if string(jsonTextMaskScatterLabel) == "{}" { // empty struct
			dst.TextMaskScatterLabel = nil
		} else {
			return nil // data stored in dst.TextMaskScatterLabel, return on the first match
		}
	} else {
		dst.TextMaskScatterLabel = nil
	}

	// try to unmarshal JSON data into TextScatterLabel
	err = json.Unmarshal(data, &dst.TextScatterLabel)
	if err == nil {
		jsonTextScatterLabel, _ := json.Marshal(dst.TextScatterLabel)
		if string(jsonTextScatterLabel) == "{}" { // empty struct
			dst.TextScatterLabel = nil
		} else {
			return nil // data stored in dst.TextScatterLabel, return on the first match
		}
	} else {
		dst.TextScatterLabel = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ScatterLabel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScatterLabel) MarshalJSON() ([]byte, error) {
	if src.BBoxImageScatterLabel != nil {
		return json.Marshal(&src.BBoxImageScatterLabel)
	}

	if src.GraphScatterLabel != nil {
		return json.Marshal(&src.GraphScatterLabel)
	}

	if src.HorizontalBarScatterLabel != nil {
		return json.Marshal(&src.HorizontalBarScatterLabel)
	}

	if src.ImageScatterLabel != nil {
		return json.Marshal(&src.ImageScatterLabel)
	}

	if src.MaskImageScatterLabel != nil {
		return json.Marshal(&src.MaskImageScatterLabel)
	}

	if src.TextMaskScatterLabel != nil {
		return json.Marshal(&src.TextMaskScatterLabel)
	}

	if src.TextScatterLabel != nil {
		return json.Marshal(&src.TextScatterLabel)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScatterLabel struct {
	value *ScatterLabel
	isSet bool
}

func (v NullableScatterLabel) Get() *ScatterLabel {
	return v.value
}

func (v *NullableScatterLabel) Set(val *ScatterLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableScatterLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableScatterLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScatterLabel(val *ScatterLabel) *NullableScatterLabel {
	return &NullableScatterLabel{value: val, isSet: true}
}

func (v NullableScatterLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScatterLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
