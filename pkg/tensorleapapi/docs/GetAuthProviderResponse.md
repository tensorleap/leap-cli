# GetAuthProviderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthProvider** | [**AuthProviderEnum**](AuthProviderEnum.md) |  | 

## Methods

### NewGetAuthProviderResponse

`func NewGetAuthProviderResponse(authProvider AuthProviderEnum, ) *GetAuthProviderResponse`

NewGetAuthProviderResponse instantiates a new GetAuthProviderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthProviderResponseWithDefaults

`func NewGetAuthProviderResponseWithDefaults() *GetAuthProviderResponse`

NewGetAuthProviderResponseWithDefaults instantiates a new GetAuthProviderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthProvider

`func (o *GetAuthProviderResponse) GetAuthProvider() AuthProviderEnum`

GetAuthProvider returns the AuthProvider field if non-nil, zero value otherwise.

### GetAuthProviderOk

`func (o *GetAuthProviderResponse) GetAuthProviderOk() (*AuthProviderEnum, bool)`

GetAuthProviderOk returns a tuple with the AuthProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthProvider

`func (o *GetAuthProviderResponse) SetAuthProvider(v AuthProviderEnum)`

SetAuthProvider sets AuthProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


