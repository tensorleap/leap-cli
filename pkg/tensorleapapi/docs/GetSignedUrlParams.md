# GetSignedUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Origin** | Pointer to **string** |  | [optional] 
**FileName** | **string** |  | 
**Expired** | **float64** |  | 
**Method** | [**HttpMethods**](HttpMethods.md) |  | 

## Methods

### NewGetSignedUrlParams

`func NewGetSignedUrlParams(fileName string, expired float64, method HttpMethods, ) *GetSignedUrlParams`

NewGetSignedUrlParams instantiates a new GetSignedUrlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSignedUrlParamsWithDefaults

`func NewGetSignedUrlParamsWithDefaults() *GetSignedUrlParams`

NewGetSignedUrlParamsWithDefaults instantiates a new GetSignedUrlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrigin

`func (o *GetSignedUrlParams) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *GetSignedUrlParams) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *GetSignedUrlParams) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *GetSignedUrlParams) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetFileName

`func (o *GetSignedUrlParams) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *GetSignedUrlParams) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *GetSignedUrlParams) SetFileName(v string)`

SetFileName sets FileName field to given value.


### GetExpired

`func (o *GetSignedUrlParams) GetExpired() float64`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *GetSignedUrlParams) GetExpiredOk() (*float64, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *GetSignedUrlParams) SetExpired(v float64)`

SetExpired sets Expired field to given value.


### GetMethod

`func (o *GetSignedUrlParams) GetMethod() HttpMethods`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GetSignedUrlParams) GetMethodOk() (*HttpMethods, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GetSignedUrlParams) SetMethod(v HttpMethods)`

SetMethod sets Method field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


