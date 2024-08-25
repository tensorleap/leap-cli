# UpdateDashboardParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DashboardId** | **string** |  | 
**Name** | **string** |  | 
**Items** | [**[]Dashlet**](Dashlet.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**IgnoreSuggestedDashletsHashes** | Pointer to **[]string** |  | [optional] 
**ProjectId** | **string** |  | 

## Methods

### NewUpdateDashboardParams

`func NewUpdateDashboardParams(dashboardId string, name string, items []Dashlet, projectId string, ) *UpdateDashboardParams`

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

`func (o *UpdateDashboardParams) GetItems() []Dashlet`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *UpdateDashboardParams) GetItemsOk() (*[]Dashlet, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *UpdateDashboardParams) SetItems(v []Dashlet)`

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

### GetIgnoreSuggestedDashletsHashes

`func (o *UpdateDashboardParams) GetIgnoreSuggestedDashletsHashes() []string`

GetIgnoreSuggestedDashletsHashes returns the IgnoreSuggestedDashletsHashes field if non-nil, zero value otherwise.

### GetIgnoreSuggestedDashletsHashesOk

`func (o *UpdateDashboardParams) GetIgnoreSuggestedDashletsHashesOk() (*[]string, bool)`

GetIgnoreSuggestedDashletsHashesOk returns a tuple with the IgnoreSuggestedDashletsHashes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSuggestedDashletsHashes

`func (o *UpdateDashboardParams) SetIgnoreSuggestedDashletsHashes(v []string)`

SetIgnoreSuggestedDashletsHashes sets IgnoreSuggestedDashletsHashes field to given value.

### HasIgnoreSuggestedDashletsHashes

`func (o *UpdateDashboardParams) HasIgnoreSuggestedDashletsHashes() bool`

HasIgnoreSuggestedDashletsHashes returns a boolean if a field has been set.

### GetProjectId

`func (o *UpdateDashboardParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateDashboardParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateDashboardParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


