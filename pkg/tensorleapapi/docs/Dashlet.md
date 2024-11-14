# Dashlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Layout** | [**SizedLayout**](SizedLayout.md) |  | 
**PinFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Type** | **string** |  | 
**Data** | [**PopulationExplorationDashletData**](PopulationExplorationDashletData.md) |  | 
**Name** | **string** |  | 
**CollectionIds** | **[]string** |  | 
**Sample** | Pointer to [**SampleIdentity**](SampleIdentity.md) |  | [optional] 
**AssetId** | Pointer to [**SampleAssetId**](SampleAssetId.md) |  | [optional] 

## Methods

### NewDashlet

`func NewDashlet(cid string, layout SizedLayout, type_ string, data PopulationExplorationDashletData, name string, collectionIds []string, ) *Dashlet`

NewDashlet instantiates a new Dashlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDashletWithDefaults

`func NewDashletWithDefaults() *Dashlet`

NewDashletWithDefaults instantiates a new Dashlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Dashlet) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Dashlet) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Dashlet) SetCid(v string)`

SetCid sets Cid field to given value.


### GetLayout

`func (o *Dashlet) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *Dashlet) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *Dashlet) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.


### GetPinFilters

`func (o *Dashlet) GetPinFilters() []ESFilter`

GetPinFilters returns the PinFilters field if non-nil, zero value otherwise.

### GetPinFiltersOk

`func (o *Dashlet) GetPinFiltersOk() (*[]ESFilter, bool)`

GetPinFiltersOk returns a tuple with the PinFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFilters

`func (o *Dashlet) SetPinFilters(v []ESFilter)`

SetPinFilters sets PinFilters field to given value.

### HasPinFilters

`func (o *Dashlet) HasPinFilters() bool`

HasPinFilters returns a boolean if a field has been set.

### GetType

`func (o *Dashlet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Dashlet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Dashlet) SetType(v string)`

SetType sets Type field to given value.


### GetData

`func (o *Dashlet) GetData() PopulationExplorationDashletData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Dashlet) GetDataOk() (*PopulationExplorationDashletData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Dashlet) SetData(v PopulationExplorationDashletData)`

SetData sets Data field to given value.


### GetName

`func (o *Dashlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Dashlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Dashlet) SetName(v string)`

SetName sets Name field to given value.


### GetCollectionIds

`func (o *Dashlet) GetCollectionIds() []string`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *Dashlet) GetCollectionIdsOk() (*[]string, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *Dashlet) SetCollectionIds(v []string)`

SetCollectionIds sets CollectionIds field to given value.


### GetSample

`func (o *Dashlet) GetSample() SampleIdentity`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *Dashlet) GetSampleOk() (*SampleIdentity, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *Dashlet) SetSample(v SampleIdentity)`

SetSample sets Sample field to given value.

### HasSample

`func (o *Dashlet) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetAssetId

`func (o *Dashlet) GetAssetId() SampleAssetId`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *Dashlet) GetAssetIdOk() (*SampleAssetId, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *Dashlet) SetAssetId(v SampleAssetId)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *Dashlet) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


