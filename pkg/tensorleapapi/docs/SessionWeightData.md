# SessionWeightData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**SessionId** | **string** |  | 
**Epoch** | **float64** |  | 
**GlobalStep** | **float64** |  | 
**Status** | [**StatusEnum**](StatusEnum.md) |  | 
**TeamId** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 

## Methods

### NewSessionWeightData

`func NewSessionWeightData(cid string, sessionId string, epoch float64, globalStep float64, status StatusEnum, teamId string, createdAt time.Time, createdBy string, ) *SessionWeightData`

NewSessionWeightData instantiates a new SessionWeightData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWeightDataWithDefaults

`func NewSessionWeightDataWithDefaults() *SessionWeightData`

NewSessionWeightDataWithDefaults instantiates a new SessionWeightData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *SessionWeightData) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SessionWeightData) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SessionWeightData) SetCid(v string)`

SetCid sets Cid field to given value.


### GetSessionId

`func (o *SessionWeightData) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionWeightData) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionWeightData) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetEpoch

`func (o *SessionWeightData) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SessionWeightData) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SessionWeightData) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetGlobalStep

`func (o *SessionWeightData) GetGlobalStep() float64`

GetGlobalStep returns the GlobalStep field if non-nil, zero value otherwise.

### GetGlobalStepOk

`func (o *SessionWeightData) GetGlobalStepOk() (*float64, bool)`

GetGlobalStepOk returns a tuple with the GlobalStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStep

`func (o *SessionWeightData) SetGlobalStep(v float64)`

SetGlobalStep sets GlobalStep field to given value.


### GetStatus

`func (o *SessionWeightData) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionWeightData) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionWeightData) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetTeamId

`func (o *SessionWeightData) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *SessionWeightData) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *SessionWeightData) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetCreatedAt

`func (o *SessionWeightData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionWeightData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionWeightData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SessionWeightData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SessionWeightData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SessionWeightData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


