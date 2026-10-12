# ConfirmPasswordResetParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Code** | **string** |  | 
**NewPassword** | **string** |  | 

## Methods

### NewConfirmPasswordResetParams

`func NewConfirmPasswordResetParams(email string, code string, newPassword string, ) *ConfirmPasswordResetParams`

NewConfirmPasswordResetParams instantiates a new ConfirmPasswordResetParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfirmPasswordResetParamsWithDefaults

`func NewConfirmPasswordResetParamsWithDefaults() *ConfirmPasswordResetParams`

NewConfirmPasswordResetParamsWithDefaults instantiates a new ConfirmPasswordResetParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *ConfirmPasswordResetParams) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ConfirmPasswordResetParams) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ConfirmPasswordResetParams) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetCode

`func (o *ConfirmPasswordResetParams) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ConfirmPasswordResetParams) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ConfirmPasswordResetParams) SetCode(v string)`

SetCode sets Code field to given value.


### GetNewPassword

`func (o *ConfirmPasswordResetParams) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *ConfirmPasswordResetParams) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *ConfirmPasswordResetParams) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


