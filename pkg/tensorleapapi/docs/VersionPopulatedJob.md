# VersionPopulatedJob

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

## Methods

### NewVersionPopulatedJob

`func NewVersionPopulatedJob(cid string, modelName string, createdAt time.Time, teamId string, ) *VersionPopulatedJob`

NewVersionPopulatedJob instantiates a new VersionPopulatedJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionPopulatedJobWithDefaults

`func NewVersionPopulatedJobWithDefaults() *VersionPopulatedJob`

NewVersionPopulatedJobWithDefaults instantiates a new VersionPopulatedJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *VersionPopulatedJob) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *VersionPopulatedJob) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *VersionPopulatedJob) SetCid(v string)`

SetCid sets Cid field to given value.


### GetExtId

`func (o *VersionPopulatedJob) GetExtId() string`

GetExtId returns the ExtId field if non-nil, zero value otherwise.

### GetExtIdOk

`func (o *VersionPopulatedJob) GetExtIdOk() (*string, bool)`

GetExtIdOk returns a tuple with the ExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtId

`func (o *VersionPopulatedJob) SetExtId(v string)`

SetExtId sets ExtId field to given value.

### HasExtId

`func (o *VersionPopulatedJob) HasExtId() bool`

HasExtId returns a boolean if a field has been set.

### GetModelName

`func (o *VersionPopulatedJob) GetModelName() string`

GetModelName returns the ModelName field if non-nil, zero value otherwise.

### GetModelNameOk

`func (o *VersionPopulatedJob) GetModelNameOk() (*string, bool)`

GetModelNameOk returns a tuple with the ModelName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelName

`func (o *VersionPopulatedJob) SetModelName(v string)`

SetModelName sets ModelName field to given value.


### GetCreatedAt

`func (o *VersionPopulatedJob) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VersionPopulatedJob) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VersionPopulatedJob) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *VersionPopulatedJob) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *VersionPopulatedJob) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *VersionPopulatedJob) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *VersionPopulatedJob) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetTeamId

`func (o *VersionPopulatedJob) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *VersionPopulatedJob) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *VersionPopulatedJob) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetHash

`func (o *VersionPopulatedJob) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *VersionPopulatedJob) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *VersionPopulatedJob) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *VersionPopulatedJob) HasHash() bool`

HasHash returns a boolean if a field has been set.

### SetHashNil

`func (o *VersionPopulatedJob) SetHashNil(b bool)`

 SetHashNil sets the value for Hash to be an explicit nil

### UnsetHash
`func (o *VersionPopulatedJob) UnsetHash()`

UnsetHash ensures that no value is present for Hash, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


