# CreateSessionTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Name** | **string** |  | 
**TestFilter** | [**ClientFilterParams**](ClientFilterParams.md) |  | 
**DatasetFilter** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 

## Methods

### NewCreateSessionTestRequest

`func NewCreateSessionTestRequest(projectId string, name string, testFilter ClientFilterParams, ) *CreateSessionTestRequest`

NewCreateSessionTestRequest instantiates a new CreateSessionTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionTestRequestWithDefaults

`func NewCreateSessionTestRequestWithDefaults() *CreateSessionTestRequest`

NewCreateSessionTestRequestWithDefaults instantiates a new CreateSessionTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *CreateSessionTestRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *CreateSessionTestRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *CreateSessionTestRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *CreateSessionTestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSessionTestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSessionTestRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTestFilter

`func (o *CreateSessionTestRequest) GetTestFilter() ClientFilterParams`

GetTestFilter returns the TestFilter field if non-nil, zero value otherwise.

### GetTestFilterOk

`func (o *CreateSessionTestRequest) GetTestFilterOk() (*ClientFilterParams, bool)`

GetTestFilterOk returns a tuple with the TestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFilter

`func (o *CreateSessionTestRequest) SetTestFilter(v ClientFilterParams)`

SetTestFilter sets TestFilter field to given value.


### GetDatasetFilter

`func (o *CreateSessionTestRequest) GetDatasetFilter() []ESFilter`

GetDatasetFilter returns the DatasetFilter field if non-nil, zero value otherwise.

### GetDatasetFilterOk

`func (o *CreateSessionTestRequest) GetDatasetFilterOk() (*[]ESFilter, bool)`

GetDatasetFilterOk returns a tuple with the DatasetFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetFilter

`func (o *CreateSessionTestRequest) SetDatasetFilter(v []ESFilter)`

SetDatasetFilter sets DatasetFilter field to given value.

### HasDatasetFilter

`func (o *CreateSessionTestRequest) HasDatasetFilter() bool`

HasDatasetFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


