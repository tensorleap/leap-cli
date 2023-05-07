# GetOrganizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Organizations** | [**[]SlimOrganization**](SlimOrganization.md) |  | 

## Methods

### NewGetOrganizationsResponse

`func NewGetOrganizationsResponse(organizations []SlimOrganization, ) *GetOrganizationsResponse`

NewGetOrganizationsResponse instantiates a new GetOrganizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrganizationsResponseWithDefaults

`func NewGetOrganizationsResponseWithDefaults() *GetOrganizationsResponse`

NewGetOrganizationsResponseWithDefaults instantiates a new GetOrganizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrganizations

`func (o *GetOrganizationsResponse) GetOrganizations() []SlimOrganization`

GetOrganizations returns the Organizations field if non-nil, zero value otherwise.

### GetOrganizationsOk

`func (o *GetOrganizationsResponse) GetOrganizationsOk() (*[]SlimOrganization, bool)`

GetOrganizationsOk returns a tuple with the Organizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizations

`func (o *GetOrganizationsResponse) SetOrganizations(v []SlimOrganization)`

SetOrganizations sets Organizations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


