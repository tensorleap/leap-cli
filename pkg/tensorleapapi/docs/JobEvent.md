# JobEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Status** | [**StatusEnum**](StatusEnum.md) |  | 
**Progress** | Pointer to [**JobEventProgress**](JobEventProgress.md) |  | [optional] 

## Methods

### NewJobEvent

`func NewJobEvent(id string, name string, status StatusEnum, ) *JobEvent`

NewJobEvent instantiates a new JobEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJobEventWithDefaults

`func NewJobEventWithDefaults() *JobEvent`

NewJobEventWithDefaults instantiates a new JobEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *JobEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *JobEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *JobEvent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *JobEvent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JobEvent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JobEvent) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *JobEvent) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *JobEvent) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *JobEvent) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetProgress

`func (o *JobEvent) GetProgress() JobEventProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *JobEvent) GetProgressOk() (*JobEventProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *JobEvent) SetProgress(v JobEventProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *JobEvent) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


