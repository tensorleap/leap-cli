# SingleUserModeSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrialStarted** | **bool** |  | 
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** |  | 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 

## Methods

### NewSingleUserModeSettings

`func NewSingleUserModeSettings(trialStarted bool, firstName string, lastName string, email string, ) *SingleUserModeSettings`

NewSingleUserModeSettings instantiates a new SingleUserModeSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSingleUserModeSettingsWithDefaults

`func NewSingleUserModeSettingsWithDefaults() *SingleUserModeSettings`

NewSingleUserModeSettingsWithDefaults instantiates a new SingleUserModeSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrialStarted

`func (o *SingleUserModeSettings) GetTrialStarted() bool`

GetTrialStarted returns the TrialStarted field if non-nil, zero value otherwise.

### GetTrialStartedOk

`func (o *SingleUserModeSettings) GetTrialStartedOk() (*bool, bool)`

GetTrialStartedOk returns a tuple with the TrialStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialStarted

`func (o *SingleUserModeSettings) SetTrialStarted(v bool)`

SetTrialStarted sets TrialStarted field to given value.


### GetFirstName

`func (o *SingleUserModeSettings) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SingleUserModeSettings) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SingleUserModeSettings) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SingleUserModeSettings) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SingleUserModeSettings) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SingleUserModeSettings) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *SingleUserModeSettings) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SingleUserModeSettings) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SingleUserModeSettings) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPhoneNumber

`func (o *SingleUserModeSettings) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SingleUserModeSettings) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SingleUserModeSettings) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SingleUserModeSettings) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetRole

`func (o *SingleUserModeSettings) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SingleUserModeSettings) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SingleUserModeSettings) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SingleUserModeSettings) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCompany

`func (o *SingleUserModeSettings) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SingleUserModeSettings) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SingleUserModeSettings) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SingleUserModeSettings) HasCompany() bool`

HasCompany returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


