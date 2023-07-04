# DashboardItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Layout** | [**SizedLayout**](SizedLayout.md) |  | 
**Type** | **string** |  | 
**Data** | [**CustomVisualizationData**](CustomVisualizationData.md) |  | 

## Methods

### NewDashboardItem

`func NewDashboardItem(cid string, layout SizedLayout, type_ string, data CustomVisualizationData, ) *DashboardItem`

NewDashboardItem instantiates a new DashboardItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashboardItemWithDefaults

`func NewDashboardItemWithDefaults() *DashboardItem`

NewDashboardItemWithDefaults instantiates a new DashboardItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *DashboardItem) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *DashboardItem) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *DashboardItem) SetCid(v string)`

SetCid sets Cid field to given value.


### GetLayout

`func (o *DashboardItem) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *DashboardItem) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *DashboardItem) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.


### GetType

`func (o *DashboardItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DashboardItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DashboardItem) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *DashboardItem) GetData() CustomVisualizationData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DashboardItem) GetDataOk() (*CustomVisualizationData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DashboardItem) SetData(v CustomVisualizationData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


