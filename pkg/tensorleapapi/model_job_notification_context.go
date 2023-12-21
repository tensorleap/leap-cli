/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.423
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// JobNotificationContext struct for JobNotificationContext
type JobNotificationContext struct {
	JobNotificationAnalyzeContext *JobNotificationAnalyzeContext
	JobNotificationBaseContext *JobNotificationBaseContext
	JobNotificationLeepScriptContext *JobNotificationLeepScriptContext
	JobNotificationModelContext *JobNotificationModelContext
	JobNotificationSampleContext *JobNotificationSampleContext
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *JobNotificationContext) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into JobNotificationAnalyzeContext
	err = json.Unmarshal(data, &dst.JobNotificationAnalyzeContext);
	if err == nil {
		jsonJobNotificationAnalyzeContext, _ := json.Marshal(dst.JobNotificationAnalyzeContext)
		if string(jsonJobNotificationAnalyzeContext) == "{}" { // empty struct
			dst.JobNotificationAnalyzeContext = nil
		} else {
			return nil // data stored in dst.JobNotificationAnalyzeContext, return on the first match
		}
	} else {
		dst.JobNotificationAnalyzeContext = nil
	}

	// try to unmarshal JSON data into JobNotificationBaseContext
	err = json.Unmarshal(data, &dst.JobNotificationBaseContext);
	if err == nil {
		jsonJobNotificationBaseContext, _ := json.Marshal(dst.JobNotificationBaseContext)
		if string(jsonJobNotificationBaseContext) == "{}" { // empty struct
			dst.JobNotificationBaseContext = nil
		} else {
			return nil // data stored in dst.JobNotificationBaseContext, return on the first match
		}
	} else {
		dst.JobNotificationBaseContext = nil
	}

	// try to unmarshal JSON data into JobNotificationLeepScriptContext
	err = json.Unmarshal(data, &dst.JobNotificationLeepScriptContext);
	if err == nil {
		jsonJobNotificationLeepScriptContext, _ := json.Marshal(dst.JobNotificationLeepScriptContext)
		if string(jsonJobNotificationLeepScriptContext) == "{}" { // empty struct
			dst.JobNotificationLeepScriptContext = nil
		} else {
			return nil // data stored in dst.JobNotificationLeepScriptContext, return on the first match
		}
	} else {
		dst.JobNotificationLeepScriptContext = nil
	}

	// try to unmarshal JSON data into JobNotificationModelContext
	err = json.Unmarshal(data, &dst.JobNotificationModelContext);
	if err == nil {
		jsonJobNotificationModelContext, _ := json.Marshal(dst.JobNotificationModelContext)
		if string(jsonJobNotificationModelContext) == "{}" { // empty struct
			dst.JobNotificationModelContext = nil
		} else {
			return nil // data stored in dst.JobNotificationModelContext, return on the first match
		}
	} else {
		dst.JobNotificationModelContext = nil
	}

	// try to unmarshal JSON data into JobNotificationSampleContext
	err = json.Unmarshal(data, &dst.JobNotificationSampleContext);
	if err == nil {
		jsonJobNotificationSampleContext, _ := json.Marshal(dst.JobNotificationSampleContext)
		if string(jsonJobNotificationSampleContext) == "{}" { // empty struct
			dst.JobNotificationSampleContext = nil
		} else {
			return nil // data stored in dst.JobNotificationSampleContext, return on the first match
		}
	} else {
		dst.JobNotificationSampleContext = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(JobNotificationContext)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *JobNotificationContext) MarshalJSON() ([]byte, error) {
	if src.JobNotificationAnalyzeContext != nil {
		return json.Marshal(&src.JobNotificationAnalyzeContext)
	}

	if src.JobNotificationBaseContext != nil {
		return json.Marshal(&src.JobNotificationBaseContext)
	}

	if src.JobNotificationLeepScriptContext != nil {
		return json.Marshal(&src.JobNotificationLeepScriptContext)
	}

	if src.JobNotificationModelContext != nil {
		return json.Marshal(&src.JobNotificationModelContext)
	}

	if src.JobNotificationSampleContext != nil {
		return json.Marshal(&src.JobNotificationSampleContext)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableJobNotificationContext struct {
	value *JobNotificationContext
	isSet bool
}

func (v NullableJobNotificationContext) Get() *JobNotificationContext {
	return v.value
}

func (v *NullableJobNotificationContext) Set(val *JobNotificationContext) {
	v.value = val
	v.isSet = true
}

func (v NullableJobNotificationContext) IsSet() bool {
	return v.isSet
}

func (v *NullableJobNotificationContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobNotificationContext(val *JobNotificationContext) *NullableJobNotificationContext {
	return &NullableJobNotificationContext{value: val, isSet: true}
}

func (v NullableJobNotificationContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobNotificationContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


