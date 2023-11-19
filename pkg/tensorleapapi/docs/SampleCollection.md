# SampleCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Samples** | [**[]SampleIdentity**](SampleIdentity.md) |  | 
**CreatedBy** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Cid** | **string** |  | 
**ProjectId** | **string** |  | 
**TeamId** | **string** |  | 

## Methods

### NewSampleCollection

`func NewSampleCollection(samples []SampleIdentity, createdAt time.Time, name string, cid string, projectId string, teamId string, ) *SampleCollection`

NewSampleCollection instantiates a new SampleCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleCollectionWithDefaults

`func NewSampleCollectionWithDefaults() *SampleCollection`

NewSampleCollectionWithDefaults instantiates a new SampleCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSamples

`func (o *SampleCollection) GetSamples() []SampleIdentity`

GetSamples returns the Samples field if non-nil, zero value otherwise.

### GetSamplesOk

`func (o *SampleCollection) GetSamplesOk() (*[]SampleIdentity, bool)`

GetSamplesOk returns a tuple with the Samples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamples

`func (o *SampleCollection) SetSamples(v []SampleIdentity)`

SetSamples sets Samples field to given value.


### GetCreatedBy

`func (o *SampleCollection) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SampleCollection) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SampleCollection) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *SampleCollection) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SampleCollection) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SampleCollection) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SampleCollection) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDescription

`func (o *SampleCollection) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SampleCollection) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SampleCollection) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SampleCollection) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SampleCollection) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SampleCollection) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SampleCollection) SetName(v string)`

SetName sets Name field to given value.


### GetCid

`func (o *SampleCollection) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SampleCollection) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SampleCollection) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProjectId

`func (o *SampleCollection) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *SampleCollection) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *SampleCollection) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetTeamId

`func (o *SampleCollection) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SampleCollection) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SampleCollection) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


