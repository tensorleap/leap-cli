# FailureInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**EngineSettingKey** | Pointer to [**EngineSettingKey**](EngineSettingKey.md) |  | [optional] 

## Methods

### NewFailureInfo

`func NewFailureInfo(message string, ) *FailureInfo`

NewFailureInfo instantiates a new FailureInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFailureInfoWithDefaults

`func NewFailureInfoWithDefaults() *FailureInfo`

NewFailureInfoWithDefaults instantiates a new FailureInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FailureInfo) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FailureInfo) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FailureInfo) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetEngineSettingKey

`func (o *FailureInfo) GetEngineSettingKey() EngineSettingKey`

GetEngineSettingKey returns the EngineSettingKey field if non-nil, zero value otherwise.

### GetEngineSettingKeyOk

`func (o *FailureInfo) GetEngineSettingKeyOk() (*EngineSettingKey, bool)`

GetEngineSettingKeyOk returns a tuple with the EngineSettingKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngineSettingKey

`func (o *FailureInfo) SetEngineSettingKey(v EngineSettingKey)`

SetEngineSettingKey sets EngineSettingKey field to given value.

### HasEngineSettingKey

`func (o *FailureInfo) HasEngineSettingKey() bool`

HasEngineSettingKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


