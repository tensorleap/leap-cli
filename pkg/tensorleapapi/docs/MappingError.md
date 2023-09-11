# MappingError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**MappingErrorType**](MappingErrorType.md) |  | 
**Message** | **string** |  | 
**Title** | **string** |  | 

## Methods

### NewMappingError

`func NewMappingError(type_ MappingErrorType, message string, title string, ) *MappingError`

NewMappingError instantiates a new MappingError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMappingErrorWithDefaults

`func NewMappingErrorWithDefaults() *MappingError`

NewMappingErrorWithDefaults instantiates a new MappingError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MappingError) GetType() MappingErrorType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MappingError) GetTypeOk() (*MappingErrorType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MappingError) SetType(v MappingErrorType)`

SetType sets Type field to given value.


### GetMessage

`func (o *MappingError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *MappingError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *MappingError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetTitle

`func (o *MappingError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MappingError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MappingError) SetTitle(v string)`

SetTitle sets Title field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


