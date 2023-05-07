# SendUserMessageParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Email** | **string** |  | 
**Telephone** | Pointer to **string** |  | [optional] 
**Content** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 

## Methods

### NewSendUserMessageParams

`func NewSendUserMessageParams(firstName string, lastName string, email string, content map[string]interface{}, ) *SendUserMessageParams`

NewSendUserMessageParams instantiates a new SendUserMessageParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendUserMessageParamsWithDefaults

`func NewSendUserMessageParamsWithDefaults() *SendUserMessageParams`

NewSendUserMessageParamsWithDefaults instantiates a new SendUserMessageParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *SendUserMessageParams) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SendUserMessageParams) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SendUserMessageParams) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *SendUserMessageParams) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SendUserMessageParams) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SendUserMessageParams) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetRole

`func (o *SendUserMessageParams) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SendUserMessageParams) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SendUserMessageParams) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SendUserMessageParams) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetCompany

`func (o *SendUserMessageParams) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SendUserMessageParams) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SendUserMessageParams) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SendUserMessageParams) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetEmail

`func (o *SendUserMessageParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendUserMessageParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendUserMessageParams) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetTelephone

`func (o *SendUserMessageParams) GetTelephone() string`

GetTelephone returns the Telephone field if non-nil, zero value otherwise.

### GetTelephoneOk

`func (o *SendUserMessageParams) GetTelephoneOk() (*string, bool)`

GetTelephoneOk returns a tuple with the Telephone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTelephone

`func (o *SendUserMessageParams) SetTelephone(v string)`

SetTelephone sets Telephone field to given value.

### HasTelephone

`func (o *SendUserMessageParams) HasTelephone() bool`

HasTelephone returns a boolean if a field has been set.

### GetContent

`func (o *SendUserMessageParams) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SendUserMessageParams) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SendUserMessageParams) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


