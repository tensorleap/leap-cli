/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.526
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"fmt"
)

// AuthProviderEnum the model 'AuthProviderEnum'
type AuthProviderEnum string

// List of AuthProviderEnum
const (
	AUTHPROVIDERENUM_LOCAL    AuthProviderEnum = "local"
	AUTHPROVIDERENUM_KEYCLOAK AuthProviderEnum = "keycloak"
)

// All allowed values of AuthProviderEnum enum
var AllowedAuthProviderEnumEnumValues = []AuthProviderEnum{
	"local",
	"keycloak",
}

func (v *AuthProviderEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthProviderEnum(value)
	for _, existing := range AllowedAuthProviderEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthProviderEnum", value)
}

// NewAuthProviderEnumFromValue returns a pointer to a valid AuthProviderEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthProviderEnumFromValue(v string) (*AuthProviderEnum, error) {
	ev := AuthProviderEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthProviderEnum: valid values are %v", v, AllowedAuthProviderEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthProviderEnum) IsValid() bool {
	for _, existing := range AllowedAuthProviderEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthProviderEnum value
func (v AuthProviderEnum) Ptr() *AuthProviderEnum {
	return &v
}

type NullableAuthProviderEnum struct {
	value *AuthProviderEnum
	isSet bool
}

func (v NullableAuthProviderEnum) Get() *AuthProviderEnum {
	return v.value
}

func (v *NullableAuthProviderEnum) Set(val *AuthProviderEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthProviderEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthProviderEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthProviderEnum(val *AuthProviderEnum) *NullableAuthProviderEnum {
	return &NullableAuthProviderEnum{value: val, isSet: true}
}

func (v NullableAuthProviderEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthProviderEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}