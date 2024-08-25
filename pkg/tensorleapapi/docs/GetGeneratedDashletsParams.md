# GetGeneratedDashletsParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**SessionRunIds** | **[]string** |  | 
**Threshold** | Pointer to **float64** |  | [optional] 
**ResetPrevious** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetGeneratedDashletsParams

`func NewGetGeneratedDashletsParams(projectId string, sessionRunIds []string, ) *GetGeneratedDashletsParams`

NewGetGeneratedDashletsParams instantiates a new GetGeneratedDashletsParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGeneratedDashletsParamsWithDefaults

`func NewGetGeneratedDashletsParamsWithDefaults() *GetGeneratedDashletsParams`

NewGetGeneratedDashletsParamsWithDefaults instantiates a new GetGeneratedDashletsParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GetGeneratedDashletsParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetGeneratedDashletsParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetGeneratedDashletsParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetSessionRunIds

`func (o *GetGeneratedDashletsParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *GetGeneratedDashletsParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *GetGeneratedDashletsParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.


### GetThreshold

`func (o *GetGeneratedDashletsParams) GetThreshold() float64`

GetThreshold returns the Threshold field if non-nil, zero value otherwise.

### GetThresholdOk

`func (o *GetGeneratedDashletsParams) GetThresholdOk() (*float64, bool)`

GetThresholdOk returns a tuple with the Threshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreshold

`func (o *GetGeneratedDashletsParams) SetThreshold(v float64)`

SetThreshold sets Threshold field to given value.

### HasThreshold

`func (o *GetGeneratedDashletsParams) HasThreshold() bool`

HasThreshold returns a boolean if a field has been set.

### GetResetPrevious

`func (o *GetGeneratedDashletsParams) GetResetPrevious() bool`

GetResetPrevious returns the ResetPrevious field if non-nil, zero value otherwise.

### GetResetPreviousOk

`func (o *GetGeneratedDashletsParams) GetResetPreviousOk() (*bool, bool)`

GetResetPreviousOk returns a tuple with the ResetPrevious field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetPrevious

`func (o *GetGeneratedDashletsParams) SetResetPrevious(v bool)`

SetResetPrevious sets ResetPrevious field to given value.

### HasResetPrevious

`func (o *GetGeneratedDashletsParams) HasResetPrevious() bool`

HasResetPrevious returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


