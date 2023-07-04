# SessionPopulatedJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSessionPopulatedJob

`func NewSessionPopulatedJob(cid string, modelName string, createdAt time.Time, teamId string, ) *SessionPopulatedJob`

NewSessionPopulatedJob instantiates a new SessionPopulatedJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionPopulatedJobWithDefaults

`func NewSessionPopulatedJobWithDefaults() *SessionPopulatedJob`

NewSessionPopulatedJobWithDefaults instantiates a new SessionPopulatedJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *SessionPopulatedJob) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SessionPopulatedJob) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SessionPopulatedJob) SetCid(v string)`

SetCid sets Cid field to given value.


### GetExtId

`func (o *SessionPopulatedJob) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *SessionPopulatedJob) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *SessionPopulatedJob) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *SessionPopulatedJob) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModelName

`func (o *SessionPopulatedJob) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *SessionPopulatedJob) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *SessionPopulatedJob) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetCreatedAt

`func (o *SessionPopulatedJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionPopulatedJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionPopulatedJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SessionPopulatedJob) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SessionPopulatedJob) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SessionPopulatedJob) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SessionPopulatedJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTeamId

`func (o *SessionPopulatedJob) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SessionPopulatedJob) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SessionPopulatedJob) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetHash

`func (o *SessionPopulatedJob) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *SessionPopulatedJob) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *SessionPopulatedJob) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *SessionPopulatedJob) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *SessionPopulatedJob) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *SessionPopulatedJob) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil
### GetTrainingParams

`func (o *SessionPopulatedJob) GetTrainingParams() TrainingParams`

GetTrainingParams returns the TrainingParams field if non-nil, zero value otherwise.

### GetTrainingParamsOk

`func (o *SessionPopulatedJob) GetTrainingParamsOk() (*TrainingParams, bool)`

GetTrainingParamsOk returns a tuple with the TrainingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingParams

`func (o *SessionPopulatedJob) SetTrainingParams(v TrainingParams)`

SetTrainingParams sets TrainingParams field to given value.

### HasTrainingParams

`func (o *SessionPopulatedJob) HasTrainingParams() bool`

HasTrainingParams returns a boolean if a field has been set.

### GetSessionRuns

`func (o *SessionPopulatedJob) GetSessionRuns() []SessionRunData`

GetSessionRuns returns the SessionRuns field if non-nil, zero value otherwise.

### GetSessionRunsOk

`func (o *SessionPopulatedJob) GetSessionRunsOk() (*[]SessionRunData, bool)`

GetSessionRunsOk returns a tuple with the SessionRuns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRuns

`func (o *SessionPopulatedJob) SetSessionRuns(v []SessionRunData)`

SetSessionRuns sets SessionRuns field to given value.

### HasSessionRuns

`func (o *SessionPopulatedJob) HasSessionRuns() bool`

HasSessionRuns returns a boolean if a field has been set.

### GetSessionWeights

`func (o *SessionPopulatedJob) GetSessionWeights() []SessionWeightData`

GetSessionWeights returns the SessionWeights field if non-nil, zero value otherwise.

### GetSessionWeightsOk

`func (o *SessionPopulatedJob) GetSessionWeightsOk() (*[]SessionWeightData, bool)`

GetSessionWeightsOk returns a tuple with the SessionWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionWeights

`func (o *SessionPopulatedJob) SetSessionWeights(v []SessionWeightData)`

SetSessionWeights sets SessionWeights field to given value.

### HasSessionWeights

`func (o *SessionPopulatedJob) HasSessionWeights() bool`

HasSessionWeights returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


