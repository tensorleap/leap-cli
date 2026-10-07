# SetAiAccessParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | Pointer to **string** |  | [optional] 
**Classes** | [**NullableAiAccessClasses**](AiAccessClasses.md) |  | 

## Methods

### NewSetAiAccessParams

`func NewSetAiAccessParams(classes NullableAiAccessClasses, ) *SetAiAccessParams`

NewSetAiAccessParams instantiates a new SetAiAccessParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetAiAccessParamsWithDefaults

`func NewSetAiAccessParamsWithDefaults() *SetAiAccessParams`

NewSetAiAccessParamsWithDefaults instantiates a new SetAiAccessParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SetAiAccessParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetAiAccessParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetAiAccessParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *SetAiAccessParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetClasses

`func (o *SetAiAccessParams) GetClasses() AiAccessClasses`

GetClasses returns the Classes field if non-nil, zero value otherwise.

### GetClassesOk

`func (o *SetAiAccessParams) GetClassesOk() (*AiAccessClasses, bool)`

GetClassesOk returns a tuple with the Classes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClasses

`func (o *SetAiAccessParams) SetClasses(v AiAccessClasses)`

SetClasses sets Classes field to given value.


### SetClassesNil

`func (o *SetAiAccessParams) SetClassesNil(b bool)`

 SetClassesNil sets the value for Classes to be an explicit nil

### UnsetClasses
`func (o *SetAiAccessParams) UnsetClasses()`

UnsetClasses ensures that no value is present for Classes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


