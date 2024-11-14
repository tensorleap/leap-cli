# PopulationExplorationDashlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Layout** | [**SizedLayout**](SizedLayout.md) |  | 
**PinFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Type** | **string** |  | 
**Name** | **string** |  | 
**Data** | [**PopulationExplorationDashletData**](PopulationExplorationDashletData.md) |  | 

## Methods

### NewPopulationExplorationDashlet

`func NewPopulationExplorationDashlet(cid string, layout SizedLayout, type_ string, name string, data PopulationExplorationDashletData, ) *PopulationExplorationDashlet`

NewPopulationExplorationDashlet instantiates a new PopulationExplorationDashlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPopulationExplorationDashletWithDefaults

`func NewPopulationExplorationDashletWithDefaults() *PopulationExplorationDashlet`

NewPopulationExplorationDashletWithDefaults instantiates a new PopulationExplorationDashlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *PopulationExplorationDashlet) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *PopulationExplorationDashlet) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *PopulationExplorationDashlet) SetCid(v string)`

SetCid sets Cid field to given value.


### GetLayout

`func (o *PopulationExplorationDashlet) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *PopulationExplorationDashlet) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *PopulationExplorationDashlet) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.


### GetPinFilters

`func (o *PopulationExplorationDashlet) GetPinFilters() []ESFilter`

GetPinFilters returns the PinFilters field if non-nil, zero value otherwise.

### GetPinFiltersOk

`func (o *PopulationExplorationDashlet) GetPinFiltersOk() (*[]ESFilter, bool)`

GetPinFiltersOk returns a tuple with the PinFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFilters

`func (o *PopulationExplorationDashlet) SetPinFilters(v []ESFilter)`

SetPinFilters sets PinFilters field to given value.

### HasPinFilters

`func (o *PopulationExplorationDashlet) HasPinFilters() bool`

HasPinFilters returns a boolean if a field has been set.

### GetType

`func (o *PopulationExplorationDashlet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PopulationExplorationDashlet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PopulationExplorationDashlet) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PopulationExplorationDashlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PopulationExplorationDashlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PopulationExplorationDashlet) SetName(v string)`

SetName sets Name field to given value.


### GetData

`func (o *PopulationExplorationDashlet) GetData() PopulationExplorationDashletData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PopulationExplorationDashlet) GetDataOk() (*PopulationExplorationDashletData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PopulationExplorationDashlet) SetData(v PopulationExplorationDashletData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


