# GetProjectIssuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | [**[]Issue**](Issue.md) |  | 

## Methods

### NewGetProjectIssuesResponse

`func NewGetProjectIssuesResponse(issues []Issue, ) *GetProjectIssuesResponse`

NewGetProjectIssuesResponse instantiates a new GetProjectIssuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetProjectIssuesResponseWithDefaults

`func NewGetProjectIssuesResponseWithDefaults() *GetProjectIssuesResponse`

NewGetProjectIssuesResponseWithDefaults instantiates a new GetProjectIssuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *GetProjectIssuesResponse) GetIssues() []Issue`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetProjectIssuesResponse) GetIssuesOk() (*[]Issue, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetProjectIssuesResponse) SetIssues(v []Issue)`

SetIssues sets Issues field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


