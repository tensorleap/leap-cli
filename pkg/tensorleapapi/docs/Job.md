# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**Version** | **string** |  | 
**Type** | [**JobTypeEnum**](JobTypeEnum.md) |  | 
**Status** | [**JobStatus**](JobStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**Params** | Pointer to [**JobParams**](JobParams.md) |  | [optional] 
**SessionRun** | **NullableString** |  | 
**Organization** | **string** |  | 

## Methods

### NewJob

`func NewJob(id string, createdBy string, version string, type_ JobTypeEnum, status JobStatus, createdAt time.Time, updatedAt time.Time, sessionRun NullableString, organization string, ) *Job`

NewJob instantiates a new Job object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobWithDefaults

`func NewJobWithDefaults() *Job`

NewJobWithDefaults instantiates a new Job object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Job) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Job) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Job) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *Job) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Job) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Job) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetVersion

`func (o *Job) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *Job) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *Job) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetType

`func (o *Job) GetType() JobTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Job) GetTypeOk() (*JobTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Job) SetType(v JobTypeEnum)`

SetType sets Type field to given value.


### GetStatus

`func (o *Job) GetStatus() JobStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Job) GetStatusOk() (*JobStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Job) SetStatus(v JobStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *Job) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Job) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Job) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Job) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Job) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Job) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetParams

`func (o *Job) GetParams() JobParams`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *Job) GetParamsOk() (*JobParams, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *Job) SetParams(v JobParams)`

SetParams sets Params field to given value.

### HasParams

`func (o *Job) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetSessionRun

`func (o *Job) GetSessionRun() string`

GetSessionRun returns the SessionRun field if non-nil, zero value otherwise.

### GetSessionRunOk

`func (o *Job) GetSessionRunOk() (*string, bool)`

GetSessionRunOk returns a tuple with the SessionRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRun

`func (o *Job) SetSessionRun(v string)`

SetSessionRun sets SessionRun field to given value.


### SetSessionRunNil

`func (o *Job) SetSessionRunNil(b bool)`

 SetSessionRunNil sets the value for SessionRun to be an explicit nil

### UnsetSessionRun
`func (o *Job) UnsetSessionRun()`

UnsetSessionRun ensures that no value is present for SessionRun, not even an explicit nil
### GetOrganization

`func (o *Job) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Job) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Job) SetOrganization(v string)`

SetOrganization sets Organization field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


