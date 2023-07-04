# UpdateSessionTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**ProjectId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**TestFilter** | Pointer to [**ClientFilterParams**](ClientFilterParams.md) |  | [optional] 
**DatasetFilter** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 

## Methods

### NewUpdateSessionTestRequest

`func NewUpdateSessionTestRequest(cid string, projectId string, ) *UpdateSessionTestRequest`

NewUpdateSessionTestRequest instantiates a new UpdateSessionTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSessionTestRequestWithDefaults

`func NewUpdateSessionTestRequestWithDefaults() *UpdateSessionTestRequest`

NewUpdateSessionTestRequestWithDefaults instantiates a new UpdateSessionTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *UpdateSessionTestRequest) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *UpdateSessionTestRequest) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *UpdateSessionTestRequest) SetCid(v string)`

SetCid sets Cid field to given value.


### GetProjectId

`func (o *UpdateSessionTestRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateSessionTestRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateSessionTestRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *UpdateSessionTestRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSessionTestRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSessionTestRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSessionTestRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTestFilter

`func (o *UpdateSessionTestRequest) GetTestFilter() ClientFilterParams`

GetTestFilter returns the TestFilter field if non-nil, zero value otherwise.

### GetTestFilterOk

`func (o *UpdateSessionTestRequest) GetTestFilterOk() (*ClientFilterParams, bool)`

GetTestFilterOk returns a tuple with the TestFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestFilter

`func (o *UpdateSessionTestRequest) SetTestFilter(v ClientFilterParams)`

SetTestFilter sets TestFilter field to given value.

### HasTestFilter

`func (o *UpdateSessionTestRequest) HasTestFilter() bool`

HasTestFilter returns a boolean if a field has been set.

### GetDatasetFilter

`func (o *UpdateSessionTestRequest) GetDatasetFilter() []ESFilter`

GetDatasetFilter returns the DatasetFilter field if non-nil, zero value otherwise.

### GetDatasetFilterOk

`func (o *UpdateSessionTestRequest) GetDatasetFilterOk() (*[]ESFilter, bool)`

GetDatasetFilterOk returns a tuple with the DatasetFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetFilter

`func (o *UpdateSessionTestRequest) SetDatasetFilter(v []ESFilter)`

SetDatasetFilter sets DatasetFilter field to given value.

### HasDatasetFilter

`func (o *UpdateSessionTestRequest) HasDatasetFilter() bool`

HasDatasetFilter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


