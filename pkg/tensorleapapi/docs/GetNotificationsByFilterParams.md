# GetNotificationsByFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**JobId** | Pointer to **string** |  | [optional] 
**FromDate** | Pointer to **string** |  | [optional] 
**NotificationTypes** | Pointer to [**[]MessageLevel**](MessageLevel.md) |  | [optional] 

## Methods

### NewGetNotificationsByFilterParams

`func NewGetNotificationsByFilterParams() *GetNotificationsByFilterParams`

NewGetNotificationsByFilterParams instantiates a new GetNotificationsByFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNotificationsByFilterParamsWithDefaults

`func NewGetNotificationsByFilterParamsWithDefaults() *GetNotificationsByFilterParams`

NewGetNotificationsByFilterParamsWithDefaults instantiates a new GetNotificationsByFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobId

`func (o *GetNotificationsByFilterParams) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *GetNotificationsByFilterParams) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *GetNotificationsByFilterParams) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *GetNotificationsByFilterParams) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetFromDate

`func (o *GetNotificationsByFilterParams) GetFromDate() string`

GetFromDate returns the FromDate field if non-nil, zero value otherwise.

### GetFromDateOk

`func (o *GetNotificationsByFilterParams) GetFromDateOk() (*string, bool)`

GetFromDateOk returns a tuple with the FromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromDate

`func (o *GetNotificationsByFilterParams) SetFromDate(v string)`

SetFromDate sets FromDate field to given value.

### HasFromDate

`func (o *GetNotificationsByFilterParams) HasFromDate() bool`

HasFromDate returns a boolean if a field has been set.

### GetNotificationTypes

`func (o *GetNotificationsByFilterParams) GetNotificationTypes() []MessageLevel`

GetNotificationTypes returns the NotificationTypes field if non-nil, zero value otherwise.

### GetNotificationTypesOk

`func (o *GetNotificationsByFilterParams) GetNotificationTypesOk() (*[]MessageLevel, bool)`

GetNotificationTypesOk returns a tuple with the NotificationTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotificationTypes

`func (o *GetNotificationsByFilterParams) SetNotificationTypes(v []MessageLevel)`

SetNotificationTypes sets NotificationTypes field to given value.

### HasNotificationTypes

`func (o *GetNotificationsByFilterParams) HasNotificationTypes() bool`

HasNotificationTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


