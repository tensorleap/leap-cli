# GetJobsFilterParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**SessionRunIds** | Pointer to **[]string** |  | [optional] 
**Types** | Pointer to [**[]JobType**](JobType.md) |  | [optional] 
**SubTypes** | Pointer to [**[]JobSubType**](JobSubType.md) |  | [optional] 
**Trigger** | Pointer to [**JobTrigger**](JobTrigger.md) |  | [optional] 
**Status** | Pointer to [**[]JobStatus**](JobStatus.md) |  | [optional] 
**ProjectId** | Pointer to **string** |  | [optional] 
**ShowJobsFromDate** | Pointer to **time.Time** |  | [optional] 
**Cid** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetJobsFilterParams

`func NewGetJobsFilterParams() *GetJobsFilterParams`

NewGetJobsFilterParams instantiates a new GetJobsFilterParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetJobsFilterParamsWithDefaults

`func NewGetJobsFilterParamsWithDefaults() *GetJobsFilterParams`

NewGetJobsFilterParamsWithDefaults instantiates a new GetJobsFilterParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastUpdated

`func (o *GetJobsFilterParams) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetJobsFilterParams) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetJobsFilterParams) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetJobsFilterParams) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSessionRunIds

`func (o *GetJobsFilterParams) GetSessionRunIds() []string`

GetSessionRunIds returns the SessionRunIds field if non-nil, zero value otherwise.

### GetSessionRunIdsOk

`func (o *GetJobsFilterParams) GetSessionRunIdsOk() (*[]string, bool)`

GetSessionRunIdsOk returns a tuple with the SessionRunIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunIds

`func (o *GetJobsFilterParams) SetSessionRunIds(v []string)`

SetSessionRunIds sets SessionRunIds field to given value.

### HasSessionRunIds

`func (o *GetJobsFilterParams) HasSessionRunIds() bool`

HasSessionRunIds returns a boolean if a field has been set.

### GetTypes

`func (o *GetJobsFilterParams) GetTypes() []JobType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *GetJobsFilterParams) GetTypesOk() (*[]JobType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *GetJobsFilterParams) SetTypes(v []JobType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *GetJobsFilterParams) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetSubTypes

`func (o *GetJobsFilterParams) GetSubTypes() []JobSubType`

GetSubTypes returns the SubTypes field if non-nil, zero value otherwise.

### GetSubTypesOk

`func (o *GetJobsFilterParams) GetSubTypesOk() (*[]JobSubType, bool)`

GetSubTypesOk returns a tuple with the SubTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubTypes

`func (o *GetJobsFilterParams) SetSubTypes(v []JobSubType)`

SetSubTypes sets SubTypes field to given value.

### HasSubTypes

`func (o *GetJobsFilterParams) HasSubTypes() bool`

HasSubTypes returns a boolean if a field has been set.

### GetTrigger

`func (o *GetJobsFilterParams) GetTrigger() JobTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *GetJobsFilterParams) GetTriggerOk() (*JobTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *GetJobsFilterParams) SetTrigger(v JobTrigger)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *GetJobsFilterParams) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetStatus

`func (o *GetJobsFilterParams) GetStatus() []JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetJobsFilterParams) GetStatusOk() (*[]JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetJobsFilterParams) SetStatus(v []JobStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetJobsFilterParams) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetProjectId

`func (o *GetJobsFilterParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetJobsFilterParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetJobsFilterParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *GetJobsFilterParams) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetShowJobsFromDate

`func (o *GetJobsFilterParams) GetShowJobsFromDate() time.Time`

GetShowJobsFromDate returns the ShowJobsFromDate field if non-nil, zero value otherwise.

### GetShowJobsFromDateOk

`func (o *GetJobsFilterParams) GetShowJobsFromDateOk() (*time.Time, bool)`

GetShowJobsFromDateOk returns a tuple with the ShowJobsFromDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowJobsFromDate

`func (o *GetJobsFilterParams) SetShowJobsFromDate(v time.Time)`

SetShowJobsFromDate sets ShowJobsFromDate field to given value.

### HasShowJobsFromDate

`func (o *GetJobsFilterParams) HasShowJobsFromDate() bool`

HasShowJobsFromDate returns a boolean if a field has been set.

### GetCid

`func (o *GetJobsFilterParams) GetCid() []string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *GetJobsFilterParams) GetCidOk() (*[]string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *GetJobsFilterParams) SetCid(v []string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *GetJobsFilterParams) HasCid() bool`

HasCid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


