# GetTeamUsersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]SlimUserData**](SlimUserData.md) |  | 

## Methods

### NewGetTeamUsersResponse

`func NewGetTeamUsersResponse(users []SlimUserData, ) *GetTeamUsersResponse`

NewGetTeamUsersResponse instantiates a new GetTeamUsersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamUsersResponseWithDefaults

`func NewGetTeamUsersResponseWithDefaults() *GetTeamUsersResponse`

NewGetTeamUsersResponseWithDefaults instantiates a new GetTeamUsersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *GetTeamUsersResponse) GetUsers() []SlimUserData`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *GetTeamUsersResponse) GetUsersOk() (*[]SlimUserData, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *GetTeamUsersResponse) SetUsers(v []SlimUserData)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


