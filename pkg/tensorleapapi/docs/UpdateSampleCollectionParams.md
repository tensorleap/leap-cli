# UpdateSampleCollectionParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Included** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Excluded** | Pointer to [**[]SampleIdentity**](SampleIdentity.md) |  | [optional] 
**Filters** | Pointer to [**[]CollectionFilterSpec**](CollectionFilterSpec.md) |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**SampleCollectionId** | **string** |  | 
**ProjectId** | **string** |  | 

## Methods

### NewUpdateSampleCollectionParams

`func NewUpdateSampleCollectionParams(sampleCollectionId string, projectId string, ) *UpdateSampleCollectionParams`

NewUpdateSampleCollectionParams instantiates a new UpdateSampleCollectionParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSampleCollectionParamsWithDefaults

`func NewUpdateSampleCollectionParamsWithDefaults() *UpdateSampleCollectionParams`

NewUpdateSampleCollectionParamsWithDefaults instantiates a new UpdateSampleCollectionParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncluded

`func (o *UpdateSampleCollectionParams) GetIncluded() []SampleIdentity`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *UpdateSampleCollectionParams) GetIncludedOk() (*[]SampleIdentity, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *UpdateSampleCollectionParams) SetIncluded(v []SampleIdentity)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *UpdateSampleCollectionParams) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetExcluded

`func (o *UpdateSampleCollectionParams) GetExcluded() []SampleIdentity`

GetExcluded returns the Excluded field if non-nil, zero value otherwise.

### GetExcludedOk

`func (o *UpdateSampleCollectionParams) GetExcludedOk() (*[]SampleIdentity, bool)`

GetExcludedOk returns a tuple with the Excluded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcluded

`func (o *UpdateSampleCollectionParams) SetExcluded(v []SampleIdentity)`

SetExcluded sets Excluded field to given value.

### HasExcluded

`func (o *UpdateSampleCollectionParams) HasExcluded() bool`

HasExcluded returns a boolean if a field has been set.

### GetFilters

`func (o *UpdateSampleCollectionParams) GetFilters() []CollectionFilterSpec`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UpdateSampleCollectionParams) GetFiltersOk() (*[]CollectionFilterSpec, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UpdateSampleCollectionParams) SetFilters(v []CollectionFilterSpec)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UpdateSampleCollectionParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVersionId

`func (o *UpdateSampleCollectionParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UpdateSampleCollectionParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UpdateSampleCollectionParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *UpdateSampleCollectionParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSampleCollectionParams) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSampleCollectionParams) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSampleCollectionParams) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSampleCollectionParams) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *UpdateSampleCollectionParams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSampleCollectionParams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSampleCollectionParams) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSampleCollectionParams) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSampleCollectionId

`func (o *UpdateSampleCollectionParams) GetSampleCollectionId() string`

GetSampleCollectionId returns the SampleCollectionId field if non-nil, zero value otherwise.

### GetSampleCollectionIdOk

`func (o *UpdateSampleCollectionParams) GetSampleCollectionIdOk() (*string, bool)`

GetSampleCollectionIdOk returns a tuple with the SampleCollectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleCollectionId

`func (o *UpdateSampleCollectionParams) SetSampleCollectionId(v string)`

SetSampleCollectionId sets SampleCollectionId field to given value.


### GetProjectId

`func (o *UpdateSampleCollectionParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateSampleCollectionParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateSampleCollectionParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


