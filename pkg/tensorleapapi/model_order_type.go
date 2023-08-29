/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.378-rc.3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// OrderType the model 'OrderType'
type OrderType string

// List of OrderType
const (
	ORDERTYPE_DESC OrderType = "desc"
	ORDERTYPE_ASC OrderType = "asc"
)

// All allowed values of OrderType enum
var AllowedOrderTypeEnumValues = []OrderType{
	"desc",
	"asc",
}

func (v *OrderType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrderType(value)
	for _, existing := range AllowedOrderTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrderType", value)
}

// NewOrderTypeFromValue returns a pointer to a valid OrderType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrderTypeFromValue(v string) (*OrderType, error) {
	ev := OrderType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrderType: valid values are %v", v, AllowedOrderTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrderType) IsValid() bool {
	for _, existing := range AllowedOrderTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrderType value
func (v OrderType) Ptr() *OrderType {
	return &v
}

type NullableOrderType struct {
	value *OrderType
	isSet bool
}

func (v NullableOrderType) Get() *OrderType {
	return v.value
}

func (v *NullableOrderType) Set(val *OrderType) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderType) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderType(val *OrderType) *NullableOrderType {
	return &NullableOrderType{value: val, isSet: true}
}

func (v NullableOrderType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

