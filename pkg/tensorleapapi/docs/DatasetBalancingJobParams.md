# DatasetBalancingJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**PercentageOfSamplesToPrune** | Pointer to **float64** |  | [optional] 
**PrioritizedMetadataTags** | Pointer to **[]string** |  | [optional] 
**MetadataTags** | **[]string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**InferenceArtifactId** | **string** |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewDatasetBalancingJobParams

`func NewDatasetBalancingJobParams(digest string, metadataTags []string, inferenceArtifactId string, versionId string, type_ string, ) *DatasetBalancingJobParams`

NewDatasetBalancingJobParams instantiates a new DatasetBalancingJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetBalancingJobParamsWithDefaults

`func NewDatasetBalancingJobParamsWithDefaults() *DatasetBalancingJobParams`

NewDatasetBalancingJobParamsWithDefaults instantiates a new DatasetBalancingJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *DatasetBalancingJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DatasetBalancingJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DatasetBalancingJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetPercentageOfSamplesToPrune

`func (o *DatasetBalancingJobParams) GetPercentageOfSamplesToPrune() float64`

GetPercentageOfSamplesToPrune returns the PercentageOfSamplesToPrune field if non-nil, zero value otherwise.

### GetPercentageOfSamplesToPruneOk

`func (o *DatasetBalancingJobParams) GetPercentageOfSamplesToPruneOk() (*float64, bool)`

GetPercentageOfSamplesToPruneOk returns a tuple with the PercentageOfSamplesToPrune field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentageOfSamplesToPrune

`func (o *DatasetBalancingJobParams) SetPercentageOfSamplesToPrune(v float64)`

SetPercentageOfSamplesToPrune sets PercentageOfSamplesToPrune field to given value.

### HasPercentageOfSamplesToPrune

`func (o *DatasetBalancingJobParams) HasPercentageOfSamplesToPrune() bool`

HasPercentageOfSamplesToPrune returns a boolean if a field has been set.

### GetPrioritizedMetadataTags

`func (o *DatasetBalancingJobParams) GetPrioritizedMetadataTags() []string`

GetPrioritizedMetadataTags returns the PrioritizedMetadataTags field if non-nil, zero value otherwise.

### GetPrioritizedMetadataTagsOk

`func (o *DatasetBalancingJobParams) GetPrioritizedMetadataTagsOk() (*[]string, bool)`

GetPrioritizedMetadataTagsOk returns a tuple with the PrioritizedMetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrioritizedMetadataTags

`func (o *DatasetBalancingJobParams) SetPrioritizedMetadataTags(v []string)`

SetPrioritizedMetadataTags sets PrioritizedMetadataTags field to given value.

### HasPrioritizedMetadataTags

`func (o *DatasetBalancingJobParams) HasPrioritizedMetadataTags() bool`

HasPrioritizedMetadataTags returns a boolean if a field has been set.

### GetMetadataTags

`func (o *DatasetBalancingJobParams) GetMetadataTags() []string`

GetMetadataTags returns the MetadataTags field if non-nil, zero value otherwise.

### GetMetadataTagsOk

`func (o *DatasetBalancingJobParams) GetMetadataTagsOk() (*[]string, bool)`

GetMetadataTagsOk returns a tuple with the MetadataTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadataTags

`func (o *DatasetBalancingJobParams) SetMetadataTags(v []string)`

SetMetadataTags sets MetadataTags field to given value.


### GetFilters

`func (o *DatasetBalancingJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *DatasetBalancingJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *DatasetBalancingJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *DatasetBalancingJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *DatasetBalancingJobParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *DatasetBalancingJobParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *DatasetBalancingJobParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVersionId

`func (o *DatasetBalancingJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DatasetBalancingJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DatasetBalancingJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *DatasetBalancingJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetBalancingJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetBalancingJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


