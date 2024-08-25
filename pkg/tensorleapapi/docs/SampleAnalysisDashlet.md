# SampleAnalysisDashlet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Layout** | [**SizedLayout**](SizedLayout.md) |  | 
**PinFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**Type** | **string** |  | 
**Name** | **string** |  | 
**CollectionIds** | **[]string** |  | 
**Sample** | Pointer to [**SampleIdentity**](SampleIdentity.md) |  | [optional] 
**AssetId** | Pointer to [**SampleAssetId**](SampleAssetId.md) |  | [optional] 

## Methods

### NewSampleAnalysisDashlet

`func NewSampleAnalysisDashlet(cid string, layout SizedLayout, type_ string, name string, collectionIds []string, ) *SampleAnalysisDashlet`

NewSampleAnalysisDashlet instantiates a new SampleAnalysisDashlet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSampleAnalysisDashletWithDefaults

`func NewSampleAnalysisDashletWithDefaults() *SampleAnalysisDashlet`

NewSampleAnalysisDashletWithDefaults instantiates a new SampleAnalysisDashlet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *SampleAnalysisDashlet) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *SampleAnalysisDashlet) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *SampleAnalysisDashlet) SetCid(v string)`

SetCid sets Cid field to given value.


### GetLayout

`func (o *SampleAnalysisDashlet) GetLayout() SizedLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SampleAnalysisDashlet) GetLayoutOk() (*SizedLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SampleAnalysisDashlet) SetLayout(v SizedLayout)`

SetLayout sets Layout field to given value.


### GetPinFilters

`func (o *SampleAnalysisDashlet) GetPinFilters() []ESFilter`

GetPinFilters returns the PinFilters field if non-nil, zero value otherwise.

### GetPinFiltersOk

`func (o *SampleAnalysisDashlet) GetPinFiltersOk() (*[]ESFilter, bool)`

GetPinFiltersOk returns a tuple with the PinFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinFilters

`func (o *SampleAnalysisDashlet) SetPinFilters(v []ESFilter)`

SetPinFilters sets PinFilters field to given value.

### HasPinFilters

`func (o *SampleAnalysisDashlet) HasPinFilters() bool`

HasPinFilters returns a boolean if a field has been set.

### GetType

`func (o *SampleAnalysisDashlet) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SampleAnalysisDashlet) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SampleAnalysisDashlet) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *SampleAnalysisDashlet) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SampleAnalysisDashlet) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SampleAnalysisDashlet) SetName(v string)`

SetName sets Name field to given value.


### GetCollectionIds

`func (o *SampleAnalysisDashlet) GetCollectionIds() []string`

GetCollectionIds returns the CollectionIds field if non-nil, zero value otherwise.

### GetCollectionIdsOk

`func (o *SampleAnalysisDashlet) GetCollectionIdsOk() (*[]string, bool)`

GetCollectionIdsOk returns a tuple with the CollectionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollectionIds

`func (o *SampleAnalysisDashlet) SetCollectionIds(v []string)`

SetCollectionIds sets CollectionIds field to given value.


### GetSample

`func (o *SampleAnalysisDashlet) GetSample() SampleIdentity`

GetSample returns the Sample field if non-nil, zero value otherwise.

### GetSampleOk

`func (o *SampleAnalysisDashlet) GetSampleOk() (*SampleIdentity, bool)`

GetSampleOk returns a tuple with the Sample field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSample

`func (o *SampleAnalysisDashlet) SetSample(v SampleIdentity)`

SetSample sets Sample field to given value.

### HasSample

`func (o *SampleAnalysisDashlet) HasSample() bool`

HasSample returns a boolean if a field has been set.

### GetAssetId

`func (o *SampleAnalysisDashlet) GetAssetId() SampleAssetId`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *SampleAnalysisDashlet) GetAssetIdOk() (*SampleAssetId, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *SampleAnalysisDashlet) SetAssetId(v SampleAssetId)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *SampleAnalysisDashlet) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


