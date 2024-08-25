# AnalyticsDashlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Layout** | [**SizedLayout**](SizedLayout.md) |  | 
**PinFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Type** | **string** |  | 
**Data** | [**AnalyticsDashletData**](AnalyticsDashletData.md) |  | 

## Methods

### NewAnalyticsDashlet

`func NewAnalyticsDashlet(cid string, layout SizedLayout, type_ string, data AnalyticsDashletData, ) *AnalyticsDashlet`

NewAnalyticsDashlet instantiates a new AnalyticsDashlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalyticsDashletWithDefaults

`func NewAnalyticsDashletWithDefaults() *AnalyticsDashlet`

NewAnalyticsDashletWithDefaults instantiates a new AnalyticsDashlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *AnalyticsDashlet) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *AnalyticsDashlet) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *AnalyticsDashlet) SetCid(v string)`

SetCid sets Cid field to given value.


### GetLayout

`func (o *AnalyticsDashlet) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AnalyticsDashlet) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AnalyticsDashlet) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.


### GetPinFilters

`func (o *AnalyticsDashlet) GetPinFilters() []ESFilter`

GetPinFilters returns the PinFilters field if non-nil, zero value otherwise.

### GetPinFiltersOk

`func (o *AnalyticsDashlet) GetPinFiltersOk() (*[]ESFilter, bool)`

GetPinFiltersOk returns a tuple with the PinFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFilters

`func (o *AnalyticsDashlet) SetPinFilters(v []ESFilter)`

SetPinFilters sets PinFilters field to given value.

### HasPinFilters

`func (o *AnalyticsDashlet) HasPinFilters() bool`

HasPinFilters returns a boolean if a field has been set.

### GetType

`func (o *AnalyticsDashlet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AnalyticsDashlet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AnalyticsDashlet) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *AnalyticsDashlet) GetData() AnalyticsDashletData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnalyticsDashlet) GetDataOk() (*AnalyticsDashletData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnalyticsDashlet) SetData(v AnalyticsDashletData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


