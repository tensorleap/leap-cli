# LocalLoginResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | Pointer to **string** |  | [optional] 
**User** | Pointer to [**UserData**](UserData.md) |  | [optional] 

## Methods

### NewLocalLoginResponse

`func NewLocalLoginResponse() *LocalLoginResponse`

NewLocalLoginResponse instantiates a new LocalLoginResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalLoginResponseWithDefaults

`func NewLocalLoginResponseWithDefaults() *LocalLoginResponse`

NewLocalLoginResponseWithDefaults instantiates a new LocalLoginResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *LocalLoginResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *LocalLoginResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *LocalLoginResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *LocalLoginResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetUser

`func (o *LocalLoginResponse) GetUser() UserData`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *LocalLoginResponse) GetUserOk() (*UserData, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *LocalLoginResponse) SetUser(v UserData)`

SetUser sets User field to given value.

### HasUser

`func (o *LocalLoginResponse) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


