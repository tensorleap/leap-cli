/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.301
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the UserDataMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataMetadata{}

// UserDataMetadata struct for UserDataMetadata
type UserDataMetadata struct {
	DnnLevel    *string `json:"dnnLevel,omitempty"`
	Company     *string `json:"company,omitempty"`
	Title       *string `json:"title,omitempty"`
	CompanySize *string `json:"companySize,omitempty"`
}

// NewUserDataMetadata instantiates a new UserDataMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataMetadata() *UserDataMetadata {
	this := UserDataMetadata{}
	return &this
}

// NewUserDataMetadataWithDefaults instantiates a new UserDataMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataMetadataWithDefaults() *UserDataMetadata {
	this := UserDataMetadata{}
	return &this
}

// GetDnnLevel returns the DnnLevel field value if set, zero value otherwise.
func (o *UserDataMetadata) GetDnnLevel() string {
	if o == nil || IsNil(o.DnnLevel) {
		var ret string
		return ret
	}
	return *o.DnnLevel
}

// GetDnnLevelOk returns a tuple with the DnnLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataMetadata) GetDnnLevelOk() (*string, bool) {
	if o == nil || IsNil(o.DnnLevel) {
		return nil, false
	}
	return o.DnnLevel, true
}

// HasDnnLevel returns a boolean if a field has been set.
func (o *UserDataMetadata) HasDnnLevel() bool {
	if o != nil && !IsNil(o.DnnLevel) {
		return true
	}

	return false
}

// SetDnnLevel gets a reference to the given string and assigns it to the DnnLevel field.
func (o *UserDataMetadata) SetDnnLevel(v string) {
	o.DnnLevel = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *UserDataMetadata) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataMetadata) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *UserDataMetadata) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *UserDataMetadata) SetCompany(v string) {
	o.Company = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserDataMetadata) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataMetadata) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserDataMetadata) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *UserDataMetadata) SetTitle(v string) {
	o.Title = &v
}

// GetCompanySize returns the CompanySize field value if set, zero value otherwise.
func (o *UserDataMetadata) GetCompanySize() string {
	if o == nil || IsNil(o.CompanySize) {
		var ret string
		return ret
	}
	return *o.CompanySize
}

// GetCompanySizeOk returns a tuple with the CompanySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataMetadata) GetCompanySizeOk() (*string, bool) {
	if o == nil || IsNil(o.CompanySize) {
		return nil, false
	}
	return o.CompanySize, true
}

// HasCompanySize returns a boolean if a field has been set.
func (o *UserDataMetadata) HasCompanySize() bool {
	if o != nil && !IsNil(o.CompanySize) {
		return true
	}

	return false
}

// SetCompanySize gets a reference to the given string and assigns it to the CompanySize field.
func (o *UserDataMetadata) SetCompanySize(v string) {
	o.CompanySize = &v
}

func (o UserDataMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DnnLevel) {
		toSerialize["dnnLevel"] = o.DnnLevel
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.CompanySize) {
		toSerialize["companySize"] = o.CompanySize
	}
	return toSerialize, nil
}

type NullableUserDataMetadata struct {
	value *UserDataMetadata
	isSet bool
}

func (v NullableUserDataMetadata) Get() *UserDataMetadata {
	return v.value
}

func (v *NullableUserDataMetadata) Set(val *UserDataMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataMetadata(val *UserDataMetadata) *NullableUserDataMetadata {
	return &NullableUserDataMetadata{value: val, isSet: true}
}

func (v NullableUserDataMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
