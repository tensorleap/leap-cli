# CustomMessageData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | [**MessageLevel**](MessageLevel.md) |  | 
**Params** | [**CustomMessageDataParams**](CustomMessageDataParams.md) |  | 

## Methods

### NewCustomMessageData

`func NewCustomMessageData(level MessageLevel, params CustomMessageDataParams, ) *CustomMessageData`

NewCustomMessageData instantiates a new CustomMessageData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomMessageDataWithDefaults

`func NewCustomMessageDataWithDefaults() *CustomMessageData`

NewCustomMessageDataWithDefaults instantiates a new CustomMessageData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *CustomMessageData) GetLevel() MessageLevel`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *CustomMessageData) GetLevelOk() (*MessageLevel, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *CustomMessageData) SetLevel(v MessageLevel)`

SetLevel sets Level field to given value.


### GetParams

`func (o *CustomMessageData) GetParams() CustomMessageDataParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *CustomMessageData) GetParamsOk() (*CustomMessageDataParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *CustomMessageData) SetParams(v CustomMessageDataParams)`

SetParams sets Params field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


