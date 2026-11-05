# SendSupportRequestParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**PageUrl** | Pointer to **string** | The page the user was on when they wrote the message. | [optional] 

## Methods

### NewSendSupportRequestParams

`func NewSendSupportRequestParams(message string, ) *SendSupportRequestParams`

NewSendSupportRequestParams instantiates a new SendSupportRequestParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendSupportRequestParamsWithDefaults

`func NewSendSupportRequestParamsWithDefaults() *SendSupportRequestParams`

NewSendSupportRequestParamsWithDefaults instantiates a new SendSupportRequestParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SendSupportRequestParams) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendSupportRequestParams) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendSupportRequestParams) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetPageUrl

`func (o *SendSupportRequestParams) GetPageUrl() string`

GetPageUrl returns the PageUrl field if non-nil, zero value otherwise.

### GetPageUrlOk

`func (o *SendSupportRequestParams) GetPageUrlOk() (*string, bool)`

GetPageUrlOk returns a tuple with the PageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUrl

`func (o *SendSupportRequestParams) SetPageUrl(v string)`

SetPageUrl sets PageUrl field to given value.

### HasPageUrl

`func (o *SendSupportRequestParams) HasPageUrl() bool`

HasPageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


