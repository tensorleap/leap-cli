# SetExperimentPropertiesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**ExperimentId** | **string** |  | 
**Properties** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 

## Methods

### NewSetExperimentPropertiesRequest

`func NewSetExperimentPropertiesRequest(projectId string, experimentId string, properties map[string]interface{}, ) *SetExperimentPropertiesRequest`

NewSetExperimentPropertiesRequest instantiates a new SetExperimentPropertiesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetExperimentPropertiesRequestWithDefaults

`func NewSetExperimentPropertiesRequestWithDefaults() *SetExperimentPropertiesRequest`

NewSetExperimentPropertiesRequestWithDefaults instantiates a new SetExperimentPropertiesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *SetExperimentPropertiesRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SetExperimentPropertiesRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SetExperimentPropertiesRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetExperimentId

`func (o *SetExperimentPropertiesRequest) GetExperimentId() string`

GetExperimentId returns the ExperimentId field if non-nil, zero value otherwise.

### GetExperimentIdOk

`func (o *SetExperimentPropertiesRequest) GetExperimentIdOk() (*string, bool)`

GetExperimentIdOk returns a tuple with the ExperimentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExperimentId

`func (o *SetExperimentPropertiesRequest) SetExperimentId(v string)`

SetExperimentId sets ExperimentId field to given value.


### GetProperties

`func (o *SetExperimentPropertiesRequest) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *SetExperimentPropertiesRequest) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *SetExperimentPropertiesRequest) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


