# GetJobsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]Job**](Job.md) |  | 
**MapVersionToProjectName** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**IncludeProject** | **bool** |  | 
**IsModelJobs** | **bool** |  | 

## Methods

### NewGetJobsResponse

`func NewGetJobsResponse(jobs []Job, mapVersionToProjectName map[string]interface{}, includeProject bool, isModelJobs bool, ) *GetJobsResponse`

NewGetJobsResponse instantiates a new GetJobsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobsResponseWithDefaults

`func NewGetJobsResponseWithDefaults() *GetJobsResponse`

NewGetJobsResponseWithDefaults instantiates a new GetJobsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *GetJobsResponse) GetJobs() []Job`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *GetJobsResponse) GetJobsOk() (*[]Job, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *GetJobsResponse) SetJobs(v []Job)`

SetJobs sets Jobs field to given value.


### GetMapVersionToProjectName

`func (o *GetJobsResponse) GetMapVersionToProjectName() map[string]interface{}`

GetMapVersionToProjectName returns the MapVersionToProjectName field if non-nil, zero value otherwise.

### GetMapVersionToProjectNameOk

`func (o *GetJobsResponse) GetMapVersionToProjectNameOk() (*map[string]interface{}, bool)`

GetMapVersionToProjectNameOk returns a tuple with the MapVersionToProjectName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapVersionToProjectName

`func (o *GetJobsResponse) SetMapVersionToProjectName(v map[string]interface{})`

SetMapVersionToProjectName sets MapVersionToProjectName field to given value.


### GetIncludeProject

`func (o *GetJobsResponse) GetIncludeProject() bool`

GetIncludeProject returns the IncludeProject field if non-nil, zero value otherwise.

### GetIncludeProjectOk

`func (o *GetJobsResponse) GetIncludeProjectOk() (*bool, bool)`

GetIncludeProjectOk returns a tuple with the IncludeProject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeProject

`func (o *GetJobsResponse) SetIncludeProject(v bool)`

SetIncludeProject sets IncludeProject field to given value.


### GetIsModelJobs

`func (o *GetJobsResponse) GetIsModelJobs() bool`

GetIsModelJobs returns the IsModelJobs field if non-nil, zero value otherwise.

### GetIsModelJobsOk

`func (o *GetJobsResponse) GetIsModelJobsOk() (*bool, bool)`

GetIsModelJobsOk returns a tuple with the IsModelJobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsModelJobs

`func (o *GetJobsResponse) SetIsModelJobs(v bool)`

SetIsModelJobs sets IsModelJobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


