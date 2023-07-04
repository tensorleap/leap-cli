# GetTeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Teams** | [**[]SlimTeam**](SlimTeam.md) |  | 

## Methods

### NewGetTeamsResponse

`func NewGetTeamsResponse(teams []SlimTeam, ) *GetTeamsResponse`

NewGetTeamsResponse instantiates a new GetTeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamsResponseWithDefaults

`func NewGetTeamsResponseWithDefaults() *GetTeamsResponse`

NewGetTeamsResponseWithDefaults instantiates a new GetTeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeams

`func (o *GetTeamsResponse) GetTeams() []SlimTeam`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *GetTeamsResponse) GetTeamsOk() (*[]SlimTeam, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *GetTeamsResponse) SetTeams(v []SlimTeam)`

SetTeams sets Teams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


