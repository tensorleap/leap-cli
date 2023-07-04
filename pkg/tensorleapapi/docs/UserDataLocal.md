# UserDataLocal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** |  | 
**Name** | **string** |  | 
**TeamId** | **string** |  | 
**TrialRequested** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserDataLocal

`func NewUserDataLocal(email string, name string, teamId string, ) *UserDataLocal`

NewUserDataLocal instantiates a new UserDataLocal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDataLocalWithDefaults

`func NewUserDataLocalWithDefaults() *UserDataLocal`

NewUserDataLocalWithDefaults instantiates a new UserDataLocal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserDataLocal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserDataLocal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserDataLocal) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetName

`func (o *UserDataLocal) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDataLocal) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDataLocal) SetName(v string)`

SetName sets Name field to given value.


### GetTeamId

`func (o *UserDataLocal) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *UserDataLocal) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *UserDataLocal) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetTrialRequested

`func (o *UserDataLocal) GetTrialRequested() bool`

GetTrialRequested returns the TrialRequested field if non-nil, zero value otherwise.

### GetTrialRequestedOk

`func (o *UserDataLocal) GetTrialRequestedOk() (*bool, bool)`

GetTrialRequestedOk returns a tuple with the TrialRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialRequested

`func (o *UserDataLocal) SetTrialRequested(v bool)`

SetTrialRequested sets TrialRequested field to given value.

### HasTrialRequested

`func (o *UserDataLocal) HasTrialRequested() bool`

HasTrialRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


