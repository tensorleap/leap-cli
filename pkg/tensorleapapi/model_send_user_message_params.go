/*
node-server

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.364
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tensorleapapi

import (
	"encoding/json"
)

// checks if the SendUserMessageParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendUserMessageParams{}

// SendUserMessageParams struct for SendUserMessageParams
type SendUserMessageParams struct {
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Role      *string `json:"role,omitempty"`
	Company   *string `json:"company,omitempty"`
	Email     string  `json:"email"`
	Telephone *string `json:"telephone,omitempty"`
	// Construct a type with a set of properties K of type T
	Content map[string]interface{} `json:"content"`
}

// NewSendUserMessageParams instantiates a new SendUserMessageParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendUserMessageParams(firstName string, lastName string, email string, content map[string]interface{}) *SendUserMessageParams {
	this := SendUserMessageParams{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Email = email
	this.Content = content
	return &this
}

// NewSendUserMessageParamsWithDefaults instantiates a new SendUserMessageParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendUserMessageParamsWithDefaults() *SendUserMessageParams {
	this := SendUserMessageParams{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *SendUserMessageParams) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *SendUserMessageParams) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *SendUserMessageParams) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *SendUserMessageParams) SetLastName(v string) {
	o.LastName = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *SendUserMessageParams) GetRole() string {
	if o == nil || IsNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetRoleOk() (*string, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *SendUserMessageParams) HasRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *SendUserMessageParams) SetRole(v string) {
	o.Role = &v
}

// GetCompany returns the Company field value if set, zero value otherwise.
func (o *SendUserMessageParams) GetCompany() string {
	if o == nil || IsNil(o.Company) {
		var ret string
		return ret
	}
	return *o.Company
}

// GetCompanyOk returns a tuple with the Company field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetCompanyOk() (*string, bool) {
	if o == nil || IsNil(o.Company) {
		return nil, false
	}
	return o.Company, true
}

// HasCompany returns a boolean if a field has been set.
func (o *SendUserMessageParams) HasCompany() bool {
	if o != nil && !IsNil(o.Company) {
		return true
	}

	return false
}

// SetCompany gets a reference to the given string and assigns it to the Company field.
func (o *SendUserMessageParams) SetCompany(v string) {
	o.Company = &v
}

// GetEmail returns the Email field value
func (o *SendUserMessageParams) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *SendUserMessageParams) SetEmail(v string) {
	o.Email = v
}

// GetTelephone returns the Telephone field value if set, zero value otherwise.
func (o *SendUserMessageParams) GetTelephone() string {
	if o == nil || IsNil(o.Telephone) {
		var ret string
		return ret
	}
	return *o.Telephone
}

// GetTelephoneOk returns a tuple with the Telephone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetTelephoneOk() (*string, bool) {
	if o == nil || IsNil(o.Telephone) {
		return nil, false
	}
	return o.Telephone, true
}

// HasTelephone returns a boolean if a field has been set.
func (o *SendUserMessageParams) HasTelephone() bool {
	if o != nil && !IsNil(o.Telephone) {
		return true
	}

	return false
}

// SetTelephone gets a reference to the given string and assigns it to the Telephone field.
func (o *SendUserMessageParams) SetTelephone(v string) {
	o.Telephone = &v
}

// GetContent returns the Content field value
func (o *SendUserMessageParams) GetContent() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *SendUserMessageParams) GetContentOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Content, true
}

// SetContent sets field value
func (o *SendUserMessageParams) SetContent(v map[string]interface{}) {
	o.Content = v
}

func (o SendUserMessageParams) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendUserMessageParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstName"] = o.FirstName
	toSerialize["lastName"] = o.LastName
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Company) {
		toSerialize["company"] = o.Company
	}
	toSerialize["email"] = o.Email
	if !IsNil(o.Telephone) {
		toSerialize["telephone"] = o.Telephone
	}
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

type NullableSendUserMessageParams struct {
	value *SendUserMessageParams
	isSet bool
}

func (v NullableSendUserMessageParams) Get() *SendUserMessageParams {
	return v.value
}

func (v *NullableSendUserMessageParams) Set(val *SendUserMessageParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSendUserMessageParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSendUserMessageParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendUserMessageParams(val *SendUserMessageParams) *NullableSendUserMessageParams {
	return &NullableSendUserMessageParams{value: val, isSet: true}
}

func (v NullableSendUserMessageParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendUserMessageParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
