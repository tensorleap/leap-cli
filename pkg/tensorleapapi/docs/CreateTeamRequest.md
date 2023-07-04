# CreateTeamRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PublicName** | **string** |  | 

## Methods

### NewCreateTeamRequest

`func NewCreateTeamRequest(name string, publicName string, ) *CreateTeamRequest`

NewCreateTeamRequest instantiates a new CreateTeamRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTeamRequestWithDefaults

`func NewCreateTeamRequestWithDefaults() *CreateTeamRequest`

NewCreateTeamRequestWithDefaults instantiates a new CreateTeamRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateTeamRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateTeamRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateTeamRequest) SetName(v string)`

SetName sets Name field to given value.


### GetPublicName

`func (o *CreateTeamRequest) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *CreateTeamRequest) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *CreateTeamRequest) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


