# SlimUserData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Role** | **string** |  | 
**RecordSession** | **bool** |  | 
**Local** | [**SlimUserDataLocal**](SlimUserDataLocal.md) |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewSlimUserData

`func NewSlimUserData(id string, role string, recordSession bool, local SlimUserDataLocal, createdAt time.Time, ) *SlimUserData`

NewSlimUserData instantiates a new SlimUserData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimUserDataWithDefaults

`func NewSlimUserDataWithDefaults() *SlimUserData`

NewSlimUserDataWithDefaults instantiates a new SlimUserData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimUserData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimUserData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimUserData) SetId(v string)`

SetId sets Id field to given value.


### GetRole

`func (o *SlimUserData) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SlimUserData) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SlimUserData) SetRole(v string)`

SetRole sets Role field to given value.


### GetRecordSession

`func (o *SlimUserData) GetRecordSession() bool`

GetRecordSession returns the RecordSession field if non-nil, zero value otherwise.

### GetRecordSessionOk

`func (o *SlimUserData) GetRecordSessionOk() (*bool, bool)`

GetRecordSessionOk returns a tuple with the RecordSession field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSession

`func (o *SlimUserData) SetRecordSession(v bool)`

SetRecordSession sets RecordSession field to given value.


### GetLocal

`func (o *SlimUserData) GetLocal() SlimUserDataLocal`

GetLocal returns the Local field if non-nil, zero value otherwise.

### GetLocalOk

`func (o *SlimUserData) GetLocalOk() (*SlimUserDataLocal, bool)`

GetLocalOk returns a tuple with the Local field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocal

`func (o *SlimUserData) SetLocal(v SlimUserDataLocal)`

SetLocal sets Local field to given value.


### GetCreatedAt

`func (o *SlimUserData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlimUserData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlimUserData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


