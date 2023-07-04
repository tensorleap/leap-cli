# UpdateSecretManagerParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Name** | **string** |  | 
**AuthData** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSecretManagerParams

`func NewUpdateSecretManagerParams(cid string, name string, ) *UpdateSecretManagerParams`

NewUpdateSecretManagerParams instantiates a new UpdateSecretManagerParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecretManagerParamsWithDefaults

`func NewUpdateSecretManagerParamsWithDefaults() *UpdateSecretManagerParams`

NewUpdateSecretManagerParamsWithDefaults instantiates a new UpdateSecretManagerParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *UpdateSecretManagerParams) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *UpdateSecretManagerParams) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *UpdateSecretManagerParams) SetCid(v string)`

SetCid sets Cid field to given value.


### GetName

`func (o *UpdateSecretManagerParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSecretManagerParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSecretManagerParams) SetName(v string)`

SetName sets Name field to given value.


### GetAuthData

`func (o *UpdateSecretManagerParams) GetAuthData() string`

GetAuthData returns the AuthData field if non-nil, zero value otherwise.

### GetAuthDataOk

`func (o *UpdateSecretManagerParams) GetAuthDataOk() (*string, bool)`

GetAuthDataOk returns a tuple with the AuthData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthData

`func (o *UpdateSecretManagerParams) SetAuthData(v string)`

SetAuthData sets AuthData field to given value.

### HasAuthData

`func (o *UpdateSecretManagerParams) HasAuthData() bool`

HasAuthData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


