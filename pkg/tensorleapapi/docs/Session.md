# Session

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Cid** | **string** |  | 
**ExtId** | Pointer to **string** |  | [optional] 
**ModelName** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 
**TeamId** | **string** |  | 
**Hash** | Pointer to **NullableString** |  | [optional] 
**TrainingParams** | Pointer to [**TrainingParams**](TrainingParams.md) |  | [optional] 
**SessionRuns** | Pointer to [**[]SessionRunData**](SessionRunData.md) |  | [optional] 
**SessionWeights** | Pointer to [**[]SessionWeightData**](SessionWeightData.md) |  | [optional] 

## Methods

### NewSession

`func NewSession(projectId string, cid string, modelName string, createdAt time.Time, teamId string, ) *Session`

NewSession instantiates a new Session object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWithDefaults

`func NewSessionWithDefaults() *Session`

NewSessionWithDefaults instantiates a new Session object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *Session) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *Session) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *Session) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetCid

`func (o *Session) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Session) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Session) SetCid(v string)`

SetCid sets Cid field to given value.


### GetExtId

`func (o *Session) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *Session) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *Session) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *Session) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModelName

`func (o *Session) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *Session) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *Session) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetCreatedAt

`func (o *Session) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Session) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Session) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *Session) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Session) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Session) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *Session) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTeamId

`func (o *Session) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *Session) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *Session) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetHash

`func (o *Session) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Session) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Session) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Session) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *Session) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *Session) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetTrainingParams

`func (o *Session) GetTrainingParams() TrainingParams`

GetTrainingParams returns the TrainingParams field if non-nil, zero value otherwise.

### GetTrainingParamsOk

`func (o *Session) GetTrainingParamsOk() (*TrainingParams, bool)`

GetTrainingParamsOk returns a tuple with the TrainingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingParams

`func (o *Session) SetTrainingParams(v TrainingParams)`

SetTrainingParams sets TrainingParams field to given value.

### HasTrainingParams

`func (o *Session) HasTrainingParams() bool`

HasTrainingParams returns a boolean if a field has been set.

### GetSessionRuns

`func (o *Session) GetSessionRuns() []SessionRunData`

GetSessionRuns returns the SessionRuns field if non-nil, zero value otherwise.

### GetSessionRunsOk

`func (o *Session) GetSessionRunsOk() (*[]SessionRunData, bool)`

GetSessionRunsOk returns a tuple with the SessionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRuns

`func (o *Session) SetSessionRuns(v []SessionRunData)`

SetSessionRuns sets SessionRuns field to given value.

### HasSessionRuns

`func (o *Session) HasSessionRuns() bool`

HasSessionRuns returns a boolean if a field has been set.

### GetSessionWeights

`func (o *Session) GetSessionWeights() []SessionWeightData`

GetSessionWeights returns the SessionWeights field if non-nil, zero value otherwise.

### GetSessionWeightsOk

`func (o *Session) GetSessionWeightsOk() (*[]SessionWeightData, bool)`

GetSessionWeightsOk returns a tuple with the SessionWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWeights

`func (o *Session) SetSessionWeights(v []SessionWeightData)`

SetSessionWeights sets SessionWeights field to given value.

### HasSessionWeights

`func (o *Session) HasSessionWeights() bool`

HasSessionWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


