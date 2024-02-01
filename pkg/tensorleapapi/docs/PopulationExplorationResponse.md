# PopulationExplorationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**JobStatus**](JobStatus.md) |  | 
**JobId** | Pointer to **string** |  | [optional] 
**ReadyArtifacts** | [**PartialPopulationExplorationArtifacts**](PartialPopulationExplorationArtifacts.md) |  | 

## Methods

### NewPopulationExplorationResponse

`func NewPopulationExplorationResponse(status JobStatus, readyArtifacts PartialPopulationExplorationArtifacts, ) *PopulationExplorationResponse`

NewPopulationExplorationResponse instantiates a new PopulationExplorationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationExplorationResponseWithDefaults

`func NewPopulationExplorationResponseWithDefaults() *PopulationExplorationResponse`

NewPopulationExplorationResponseWithDefaults instantiates a new PopulationExplorationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PopulationExplorationResponse) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PopulationExplorationResponse) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PopulationExplorationResponse) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetJobId

`func (o *PopulationExplorationResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *PopulationExplorationResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *PopulationExplorationResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *PopulationExplorationResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### GetReadyArtifacts

`func (o *PopulationExplorationResponse) GetReadyArtifacts() PartialPopulationExplorationArtifacts`

GetReadyArtifacts returns the ReadyArtifacts field if non-nil, zero value otherwise.

### GetReadyArtifactsOk

`func (o *PopulationExplorationResponse) GetReadyArtifactsOk() (*PartialPopulationExplorationArtifacts, bool)`

GetReadyArtifactsOk returns a tuple with the ReadyArtifacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadyArtifacts

`func (o *PopulationExplorationResponse) SetReadyArtifacts(v PartialPopulationExplorationArtifacts)`

SetReadyArtifacts sets ReadyArtifacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


