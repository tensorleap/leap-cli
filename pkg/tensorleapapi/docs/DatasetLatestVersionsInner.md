# DatasetLatestVersionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latest** | Pointer to [**DatasetVersion**](DatasetVersion.md) |  | [optional] 
**LatestValid** | Pointer to [**DatasetVersion**](DatasetVersion.md) |  | [optional] 
**Branch** | **string** |  | 

## Methods

### NewDatasetLatestVersionsInner

`func NewDatasetLatestVersionsInner(branch string, ) *DatasetLatestVersionsInner`

NewDatasetLatestVersionsInner instantiates a new DatasetLatestVersionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetLatestVersionsInnerWithDefaults

`func NewDatasetLatestVersionsInnerWithDefaults() *DatasetLatestVersionsInner`

NewDatasetLatestVersionsInnerWithDefaults instantiates a new DatasetLatestVersionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatest

`func (o *DatasetLatestVersionsInner) GetLatest() DatasetVersion`

GetLatest returns the Latest field if non-nil, zero value otherwise.

### GetLatestOk

`func (o *DatasetLatestVersionsInner) GetLatestOk() (*DatasetVersion, bool)`

GetLatestOk returns a tuple with the Latest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatest

`func (o *DatasetLatestVersionsInner) SetLatest(v DatasetVersion)`

SetLatest sets Latest field to given value.

### HasLatest

`func (o *DatasetLatestVersionsInner) HasLatest() bool`

HasLatest returns a boolean if a field has been set.

### GetLatestValid

`func (o *DatasetLatestVersionsInner) GetLatestValid() DatasetVersion`

GetLatestValid returns the LatestValid field if non-nil, zero value otherwise.

### GetLatestValidOk

`func (o *DatasetLatestVersionsInner) GetLatestValidOk() (*DatasetVersion, bool)`

GetLatestValidOk returns a tuple with the LatestValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestValid

`func (o *DatasetLatestVersionsInner) SetLatestValid(v DatasetVersion)`

SetLatestValid sets LatestValid field to given value.

### HasLatestValid

`func (o *DatasetLatestVersionsInner) HasLatestValid() bool`

HasLatestValid returns a boolean if a field has been set.

### GetBranch

`func (o *DatasetLatestVersionsInner) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DatasetLatestVersionsInner) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DatasetLatestVersionsInner) SetBranch(v string)`

SetBranch sets Branch field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


