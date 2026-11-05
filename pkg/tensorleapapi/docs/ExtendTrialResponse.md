# ExtendTrialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**RegisteredUser** | Pointer to [**ExtendTrialResponseRegisteredUser**](ExtendTrialResponseRegisteredUser.md) |  | [optional] 

## Methods

### NewExtendTrialResponse

`func NewExtendTrialResponse(success bool, ) *ExtendTrialResponse`

NewExtendTrialResponse instantiates a new ExtendTrialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExtendTrialResponseWithDefaults

`func NewExtendTrialResponseWithDefaults() *ExtendTrialResponse`

NewExtendTrialResponseWithDefaults instantiates a new ExtendTrialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ExtendTrialResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ExtendTrialResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ExtendTrialResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetRegisteredUser

`func (o *ExtendTrialResponse) GetRegisteredUser() ExtendTrialResponseRegisteredUser`

GetRegisteredUser returns the RegisteredUser field if non-nil, zero value otherwise.

### GetRegisteredUserOk

`func (o *ExtendTrialResponse) GetRegisteredUserOk() (*ExtendTrialResponseRegisteredUser, bool)`

GetRegisteredUserOk returns a tuple with the RegisteredUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegisteredUser

`func (o *ExtendTrialResponse) SetRegisteredUser(v ExtendTrialResponseRegisteredUser)`

SetRegisteredUser sets RegisteredUser field to given value.

### HasRegisteredUser

`func (o *ExtendTrialResponse) HasRegisteredUser() bool`

HasRegisteredUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


