/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.465
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
	"time"
)

// checks if the SingleUserModeSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SingleUserModeSettings{}

// SingleUserModeSettings struct for SingleUserModeSettings
type SingleUserModeSettings struct {
	TrialStarted   bool       `json:"trialStarted"`
	ExpirationDate *time.Time `json:"expirationDate,omitempty"`
	FirstName      string     `json:"firstName"`
	LastName       string     `json:"lastName"`
	Email          string     `json:"email"`
	PhoneNumber    *string    `json:"phoneNumber,omitempty"`
	Role           *string    `json:"role,omitempty"`
	Company        *string    `json:"company,omitempty"`
}

// NewSingleUserModeSettings instantiates a new SingleUserModeSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSingleUserModeSettings(trialStarted bool, firstName string, lastName string, email string) *SingleUserModeSettings {
	this := SingleUserModeSettings{}
	this.TrialStarted = trialStarted
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	return &this
}

// NewSingleUserModeSettingsWithDefaults instantiates a new SingleUserModeSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSingleUserModeSettingsWithDefaults() *SingleUserModeSettings {
	this := SingleUserModeSettings{}
	return &this
}

// GetTrialStarted returns the TrialStarted field value
func (o *SingleUserModeSettings) GetTrialStarted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TrialStarted
}

// GetTrialStartedOk returns a tuple with the TrialStarted field value
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetTrialStartedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TrialStarted, true
}

// SetTrialStarted sets field value
func (o *SingleUserModeSettings) SetTrialStarted(v bool) {
	o.TrialStarted = v
}

// GetExpirationDate returns the ExpirationDate field value if set, zero value otherwise.
func (o *SingleUserModeSettings) GetExpirationDate() time.Time {
	if o == nil || IsNil(o.ExpirationDate) {
		var ret time.Time
		return ret
	}
	return *o.ExpirationDate
}

// GetExpirationDateOk returns a tuple with the ExpirationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetExpirationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpirationDate) {
		return nil, false
	}
	return o.ExpirationDate, true
}

// HasExpirationDate returns a boolean if a field has been set.
func (o *SingleUserModeSettings) HasExpirationDate() bool {
	if o != nil && !IsNil(o.ExpirationDate) {
		return true
	}

	return false
}

// SetExpirationDate gets a reference to the given time.Time and assigns it to the ExpirationDate field.
func (o *SingleUserModeSettings) SetExpirationDate(v time.Time) {
	o.ExpirationDate = &v
}

// GetFirstName returns the FirstName field value
func (o *SingleUserModeSettings) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *SingleUserModeSettings) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *SingleUserModeSettings) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *SingleUserModeSettings) SetLastName(v string) {
	o.LastName = v
}

// GetEmail returns the Email field value
func (o *SingleUserModeSettings) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SingleUserModeSettings) SetEmail(v string) {
	o.Email = v
}

// GetPhoneNumber returns the PhoneNumber field value if set, zero value otherwise.
func (o *SingleUserModeSettings) GetPhoneNumber() string {
	if o == nil || IsNil(o.PhoneNumber) {
		var ret string
		return ret
	}
	return *o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetPhoneNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumber) {
		return nil, false
	}
	return o.PhoneNumber, true
}

// HasPhoneNumber returns a boolean if a field has been set.
func (o *SingleUserModeSettings) HasPhoneNumber() bool {
	if o != nil && !IsNil(o.PhoneNumber) {
		return true
	}

	return false
}

// SetPhoneNumber gets a reference to the given string and assigns it to the PhoneNumber field.
func (o *SingleUserModeSettings) SetPhoneNumber(v string) {
	o.PhoneNumber = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SingleUserModeSettings) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SingleUserModeSettings) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SingleUserModeSettings) SetRole(v string) {
	o.Role = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SingleUserModeSettings) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SingleUserModeSettings) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SingleUserModeSettings) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *SingleUserModeSettings) SetCompany(v string) {
	o.Company = &v
}

func (o SingleUserModeSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SingleUserModeSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["trialStarted"] = o.TrialStarted
	if !IsNil(o.ExpirationDate) {
		toSerialize["expirationDate"] = o.ExpirationDate
	}
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

type NullableSingleUserModeSettings struct {
	value *SingleUserModeSettings
	isSet bool
}

func (v NullableSingleUserModeSettings) Get() *SingleUserModeSettings {
	return v.value
}

func (v *NullableSingleUserModeSettings) Set(val *SingleUserModeSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableSingleUserModeSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableSingleUserModeSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSingleUserModeSettings(val *SingleUserModeSettings) *NullableSingleUserModeSettings {
	return &NullableSingleUserModeSettings{value: val, isSet: true}
}

func (v NullableSingleUserModeSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSingleUserModeSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
