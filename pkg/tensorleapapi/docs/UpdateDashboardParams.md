# UpdateDashboardParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardId** | **string** |  | 
**Name** | **string** |  | 
**Items** | [**[]DashboardItem**](DashboardItem.md) |  | 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateDashboardParams

`func NewUpdateDashboardParams(dashboardId string, name string, items []DashboardItem, ) *UpdateDashboardParams`

NewUpdateDashboardParams instantiates a new UpdateDashboardParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDashboardParamsWithDefaults

`func NewUpdateDashboardParamsWithDefaults() *UpdateDashboardParams`

NewUpdateDashboardParamsWithDefaults instantiates a new UpdateDashboardParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboardId

`func (o *UpdateDashboardParams) GetDashboardId() string`

GetDashboardId returns the DashboardId field if non-nil, zero value otherwise.

### GetDashboardIdOk

`func (o *UpdateDashboardParams) GetDashboardIdOk() (*string, bool)`

GetDashboardIdOk returns a tuple with the DashboardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboardId

`func (o *UpdateDashboardParams) SetDashboardId(v string)`

SetDashboardId sets DashboardId field to given value.


### GetName

`func (o *UpdateDashboardParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateDashboardParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateDashboardParams) SetName(v string)`

SetName sets Name field to given value.


### GetItems

`func (o *UpdateDashboardParams) GetItems() []DashboardItem`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateDashboardParams) GetItemsOk() (*[]DashboardItem, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateDashboardParams) SetItems(v []DashboardItem)`

SetItems sets Items field to given value.


### GetDescription

`func (o *UpdateDashboardParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateDashboardParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateDashboardParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateDashboardParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


