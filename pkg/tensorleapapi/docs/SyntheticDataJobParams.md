# SyntheticDataJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**Sources** | [**[]SyntheticDataJobParamsSourcesInner**](SyntheticDataJobParamsSourcesInner.md) |  | 
**InferenceArtifactId** | **string** |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewSyntheticDataJobParams

`func NewSyntheticDataJobParams(digest string, targetFilters []ESFilter, sources []SyntheticDataJobParamsSourcesInner, inferenceArtifactId string, versionId string, type_ string, ) *SyntheticDataJobParams`

NewSyntheticDataJobParams instantiates a new SyntheticDataJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticDataJobParamsWithDefaults

`func NewSyntheticDataJobParamsWithDefaults() *SyntheticDataJobParams`

NewSyntheticDataJobParamsWithDefaults instantiates a new SyntheticDataJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *SyntheticDataJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SyntheticDataJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SyntheticDataJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTargetFilters

`func (o *SyntheticDataJobParams) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *SyntheticDataJobParams) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *SyntheticDataJobParams) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.


### GetSources

`func (o *SyntheticDataJobParams) GetSources() []SyntheticDataJobParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SyntheticDataJobParams) GetSourcesOk() (*[]SyntheticDataJobParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SyntheticDataJobParams) SetSources(v []SyntheticDataJobParamsSourcesInner)`

SetSources sets Sources field to given value.


### GetInferenceArtifactId

`func (o *SyntheticDataJobParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *SyntheticDataJobParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *SyntheticDataJobParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVersionId

`func (o *SyntheticDataJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *SyntheticDataJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *SyntheticDataJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *SyntheticDataJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticDataJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticDataJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


