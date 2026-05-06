# RemoveSamplesFromCollectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Excluded** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**SampleCollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewRemoveSamplesFromCollectionParams

`func NewRemoveSamplesFromCollectionParams(sampleCollectionId string, projectId string, ) *RemoveSamplesFromCollectionParams`

NewRemoveSamplesFromCollectionParams instantiates a new RemoveSamplesFromCollectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveSamplesFromCollectionParamsWithDefaults

`func NewRemoveSamplesFromCollectionParamsWithDefaults() *RemoveSamplesFromCollectionParams`

NewRemoveSamplesFromCollectionParamsWithDefaults instantiates a new RemoveSamplesFromCollectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *RemoveSamplesFromCollectionParams) GetIncluded() []SampleIdentity`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *RemoveSamplesFromCollectionParams) GetIncludedOk() (*[]SampleIdentity, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *RemoveSamplesFromCollectionParams) SetIncluded(v []SampleIdentity)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *RemoveSamplesFromCollectionParams) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetExcluded

`func (o *RemoveSamplesFromCollectionParams) GetExcluded() []SampleIdentity`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *RemoveSamplesFromCollectionParams) GetExcludedOk() (*[]SampleIdentity, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *RemoveSamplesFromCollectionParams) SetExcluded(v []SampleIdentity)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *RemoveSamplesFromCollectionParams) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetFilters

`func (o *RemoveSamplesFromCollectionParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *RemoveSamplesFromCollectionParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *RemoveSamplesFromCollectionParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *RemoveSamplesFromCollectionParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVersionId

`func (o *RemoveSamplesFromCollectionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *RemoveSamplesFromCollectionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *RemoveSamplesFromCollectionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *RemoveSamplesFromCollectionParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetSampleCollectionId

`func (o *RemoveSamplesFromCollectionParams) GetSampleCollectionId() string`

GetSampleCollectionId returns the SampleCollectionId field if non-nil, zero value otherwise.

### GetSampleCollectionIdOk

`func (o *RemoveSamplesFromCollectionParams) GetSampleCollectionIdOk() (*string, bool)`

GetSampleCollectionIdOk returns a tuple with the SampleCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCollectionId

`func (o *RemoveSamplesFromCollectionParams) SetSampleCollectionId(v string)`

SetSampleCollectionId sets SampleCollectionId field to given value.


### GetProjectId

`func (o *RemoveSamplesFromCollectionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *RemoveSamplesFromCollectionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *RemoveSamplesFromCollectionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


