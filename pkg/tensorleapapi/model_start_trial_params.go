/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the StartTrialParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartTrialParams{}

// StartTrialParams struct for StartTrialParams
type StartTrialParams struct {
	FirstName   string  `json:"firstName"`
	LastName    string  `json:"lastName"`
	Email       string  `json:"email"`
	PhoneNumber *string `json:"phoneNumber,omitempty"`
	Role        *string `json:"role,omitempty"`
	Company     *string `json:"company,omitempty"`
}

// NewStartTrialParams instantiates a new StartTrialParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartTrialParams(firstName string, lastName string, email string) *StartTrialParams {
	this := StartTrialParams{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	return &this
}

// NewStartTrialParamsWithDefaults instantiates a new StartTrialParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartTrialParamsWithDefaults() *StartTrialParams {
	this := StartTrialParams{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *StartTrialParams) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *StartTrialParams) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *StartTrialParams) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *StartTrialParams) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *StartTrialParams) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *StartTrialParams) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *StartTrialParams) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *StartTrialParams) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *StartTrialParams) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *StartTrialParams) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *StartTrialParams) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *StartTrialParams) SetRole(v string) {
	o.Role = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *StartTrialParams) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartTrialParams) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *StartTrialParams) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *StartTrialParams) SetCompany(v string) {
	o.Company = &v
}

func (o StartTrialParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartTrialParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	toSerialize["email"] = o.Email
	if !IsNil(o.PhoneNumber) {
		toSerialize["phoneNumber"] = o.PhoneNumber
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	return toSerialize, nil
}

type NullableStartTrialParams struct {
	value *StartTrialParams
	isSet bool
}

func (v NullableStartTrialParams) Get() *StartTrialParams {
	return v.value
}

func (v *NullableStartTrialParams) Set(val *StartTrialParams) {
	v.value = val
	v.isSet = true
}

func (v NullableStartTrialParams) IsSet() bool {
	return v.isSet
}

func (v *NullableStartTrialParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartTrialParams(val *StartTrialParams) *NullableStartTrialParams {
	return &NullableStartTrialParams{value: val, isSet: true}
}

func (v NullableStartTrialParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartTrialParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
