# AddSampleCollectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Excluded** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**ProjectId** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewAddSampleCollectionParams

`func NewAddSampleCollectionParams(projectId string, name string, ) *AddSampleCollectionParams`

NewAddSampleCollectionParams instantiates a new AddSampleCollectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSampleCollectionParamsWithDefaults

`func NewAddSampleCollectionParamsWithDefaults() *AddSampleCollectionParams`

NewAddSampleCollectionParamsWithDefaults instantiates a new AddSampleCollectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *AddSampleCollectionParams) GetIncluded() []SampleIdentity`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AddSampleCollectionParams) GetIncludedOk() (*[]SampleIdentity, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AddSampleCollectionParams) SetIncluded(v []SampleIdentity)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AddSampleCollectionParams) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetExcluded

`func (o *AddSampleCollectionParams) GetExcluded() []SampleIdentity`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *AddSampleCollectionParams) GetExcludedOk() (*[]SampleIdentity, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *AddSampleCollectionParams) SetExcluded(v []SampleIdentity)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *AddSampleCollectionParams) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetFilters

`func (o *AddSampleCollectionParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *AddSampleCollectionParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *AddSampleCollectionParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *AddSampleCollectionParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVersionId

`func (o *AddSampleCollectionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AddSampleCollectionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AddSampleCollectionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *AddSampleCollectionParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetProjectId

`func (o *AddSampleCollectionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *AddSampleCollectionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *AddSampleCollectionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetDescription

`func (o *AddSampleCollectionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddSampleCollectionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddSampleCollectionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddSampleCollectionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *AddSampleCollectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddSampleCollectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddSampleCollectionParams) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


