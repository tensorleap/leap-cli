# AddDashboardParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**Name** | **string** |  | 
**Items** | [**[]DashboardItem**](DashboardItem.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewAddDashboardParams

`func NewAddDashboardParams(projectId string, name string, items []DashboardItem, ) *AddDashboardParams`

NewAddDashboardParams instantiates a new AddDashboardParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDashboardParamsWithDefaults

`func NewAddDashboardParamsWithDefaults() *AddDashboardParams`

NewAddDashboardParamsWithDefaults instantiates a new AddDashboardParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *AddDashboardParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddDashboardParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddDashboardParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetName

`func (o *AddDashboardParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddDashboardParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddDashboardParams) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *AddDashboardParams) GetItems() []DashboardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AddDashboardParams) GetItemsOk() (*[]DashboardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AddDashboardParams) SetItems(v []DashboardItem)`

SetItems sets Items field to given value.


### GetDescription

`func (o *AddDashboardParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddDashboardParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddDashboardParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddDashboardParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


