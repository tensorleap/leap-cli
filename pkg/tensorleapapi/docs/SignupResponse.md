# SignupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserCid** | **string** |  | 
**Success** | **bool** |  | 

## Methods

### NewSignupResponse

`func NewSignupResponse(userCid string, success bool, ) *SignupResponse`

NewSignupResponse instantiates a new SignupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignupResponseWithDefaults

`func NewSignupResponseWithDefaults() *SignupResponse`

NewSignupResponseWithDefaults instantiates a new SignupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserCid

`func (o *SignupResponse) GetUserCid() string`

GetUserCid returns the UserCid field if non-nil, zero value otherwise.

### GetUserCidOk

`func (o *SignupResponse) GetUserCidOk() (*string, bool)`

GetUserCidOk returns a tuple with the UserCid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCid

`func (o *SignupResponse) SetUserCid(v string)`

SetUserCid sets UserCid field to given value.


### GetSuccess

`func (o *SignupResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SignupResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SignupResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


