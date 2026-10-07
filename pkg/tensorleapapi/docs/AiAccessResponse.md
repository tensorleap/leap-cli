# AiAccessResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Install** | [**AiAccessClasses**](AiAccessClasses.md) |  | 
**Project** | Pointer to [**NullableAiAccessClasses**](AiAccessClasses.md) |  | [optional] 
**Effective** | [**AiAccessClasses**](AiAccessClasses.md) |  | 
**ProjectOverrides** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 

## Methods

### NewAiAccessResponse

`func NewAiAccessResponse(install AiAccessClasses, effective AiAccessClasses, ) *AiAccessResponse`

NewAiAccessResponse instantiates a new AiAccessResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiAccessResponseWithDefaults

`func NewAiAccessResponseWithDefaults() *AiAccessResponse`

NewAiAccessResponseWithDefaults instantiates a new AiAccessResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstall

`func (o *AiAccessResponse) GetInstall() AiAccessClasses`

GetInstall returns the Install field if non-nil, zero value otherwise.

### GetInstallOk

`func (o *AiAccessResponse) GetInstallOk() (*AiAccessClasses, bool)`

GetInstallOk returns a tuple with the Install field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstall

`func (o *AiAccessResponse) SetInstall(v AiAccessClasses)`

SetInstall sets Install field to given value.


### GetProject

`func (o *AiAccessResponse) GetProject() AiAccessClasses`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *AiAccessResponse) GetProjectOk() (*AiAccessClasses, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *AiAccessResponse) SetProject(v AiAccessClasses)`

SetProject sets Project field to given value.

### HasProject

`func (o *AiAccessResponse) HasProject() bool`

HasProject returns a boolean if a field has been set.

### SetProjectNil

`func (o *AiAccessResponse) SetProjectNil(b bool)`

 SetProjectNil sets the value for Project to be an explicit nil

### UnsetProject
`func (o *AiAccessResponse) UnsetProject()`

UnsetProject ensures that no value is present for Project, not even an explicit nil
### GetEffective

`func (o *AiAccessResponse) GetEffective() AiAccessClasses`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *AiAccessResponse) GetEffectiveOk() (*AiAccessClasses, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *AiAccessResponse) SetEffective(v AiAccessClasses)`

SetEffective sets Effective field to given value.


### GetProjectOverrides

`func (o *AiAccessResponse) GetProjectOverrides() map[string]interface{}`

GetProjectOverrides returns the ProjectOverrides field if non-nil, zero value otherwise.

### GetProjectOverridesOk

`func (o *AiAccessResponse) GetProjectOverridesOk() (*map[string]interface{}, bool)`

GetProjectOverridesOk returns a tuple with the ProjectOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectOverrides

`func (o *AiAccessResponse) SetProjectOverrides(v map[string]interface{})`

SetProjectOverrides sets ProjectOverrides field to given value.

### HasProjectOverrides

`func (o *AiAccessResponse) HasProjectOverrides() bool`

HasProjectOverrides returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


