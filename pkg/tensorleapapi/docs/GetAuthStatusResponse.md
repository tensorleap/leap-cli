# GetAuthStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**AuthStatus**](AuthStatus.md) |  | 

## Methods

### NewGetAuthStatusResponse

`func NewGetAuthStatusResponse(status AuthStatus, ) *GetAuthStatusResponse`

NewGetAuthStatusResponse instantiates a new GetAuthStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAuthStatusResponseWithDefaults

`func NewGetAuthStatusResponseWithDefaults() *GetAuthStatusResponse`

NewGetAuthStatusResponseWithDefaults instantiates a new GetAuthStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GetAuthStatusResponse) GetStatus() AuthStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAuthStatusResponse) GetStatusOk() (*AuthStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAuthStatusResponse) SetStatus(v AuthStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


