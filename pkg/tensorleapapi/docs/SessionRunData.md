# SessionRunData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionId** | **string** |  | 
**Name** | **string** |  | 
**Organization** | **string** |  | 
**IsEvaluate** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 
**CreatedBy** | **string** |  | 
**WeightAssets** | [**[]WeightAssetData**](WeightAssetData.md) |  | 

## Methods

### NewSessionRunData

`func NewSessionRunData(id string, sessionId string, name string, organization string, isEvaluate bool, createdAt time.Time, createdBy string, weightAssets []WeightAssetData, ) *SessionRunData`

NewSessionRunData instantiates a new SessionRunData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionRunDataWithDefaults

`func NewSessionRunDataWithDefaults() *SessionRunData`

NewSessionRunDataWithDefaults instantiates a new SessionRunData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SessionRunData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SessionRunData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SessionRunData) SetId(v string)`

SetId sets Id field to given value.


### GetSessionId

`func (o *SessionRunData) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *SessionRunData) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *SessionRunData) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.


### GetName

`func (o *SessionRunData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SessionRunData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SessionRunData) SetName(v string)`

SetName sets Name field to given value.


### GetOrganization

`func (o *SessionRunData) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SessionRunData) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SessionRunData) SetOrganization(v string)`

SetOrganization sets Organization field to given value.


### GetIsEvaluate

`func (o *SessionRunData) GetIsEvaluate() bool`

GetIsEvaluate returns the IsEvaluate field if non-nil, zero value otherwise.

### GetIsEvaluateOk

`func (o *SessionRunData) GetIsEvaluateOk() (*bool, bool)`

GetIsEvaluateOk returns a tuple with the IsEvaluate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEvaluate

`func (o *SessionRunData) SetIsEvaluate(v bool)`

SetIsEvaluate sets IsEvaluate field to given value.


### GetCreatedAt

`func (o *SessionRunData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SessionRunData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SessionRunData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *SessionRunData) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SessionRunData) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SessionRunData) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetWeightAssets

`func (o *SessionRunData) GetWeightAssets() []WeightAssetData`

GetWeightAssets returns the WeightAssets field if non-nil, zero value otherwise.

### GetWeightAssetsOk

`func (o *SessionRunData) GetWeightAssetsOk() (*[]WeightAssetData, bool)`

GetWeightAssetsOk returns a tuple with the WeightAssets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightAssets

`func (o *SessionRunData) SetWeightAssets(v []WeightAssetData)`

SetWeightAssets sets WeightAssets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


