# DatasetSplittingJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**SplitAcrossMetadata** | **[]string** |  | 
**KeepTogetherMetadata** | **[]string** |  | 
**SplitsToResplit** | [**[]SplitSubset**](SplitSubset.md) |  | 
**InferenceArtifactId** | **string** |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewDatasetSplittingJobParams

`func NewDatasetSplittingJobParams(digest string, splitAcrossMetadata []string, keepTogetherMetadata []string, splitsToResplit []SplitSubset, inferenceArtifactId string, versionId string, type_ string, ) *DatasetSplittingJobParams`

NewDatasetSplittingJobParams instantiates a new DatasetSplittingJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetSplittingJobParamsWithDefaults

`func NewDatasetSplittingJobParamsWithDefaults() *DatasetSplittingJobParams`

NewDatasetSplittingJobParamsWithDefaults instantiates a new DatasetSplittingJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *DatasetSplittingJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *DatasetSplittingJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *DatasetSplittingJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetSplitAcrossMetadata

`func (o *DatasetSplittingJobParams) GetSplitAcrossMetadata() []string`

GetSplitAcrossMetadata returns the SplitAcrossMetadata field if non-nil, zero value otherwise.

### GetSplitAcrossMetadataOk

`func (o *DatasetSplittingJobParams) GetSplitAcrossMetadataOk() (*[]string, bool)`

GetSplitAcrossMetadataOk returns a tuple with the SplitAcrossMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAcrossMetadata

`func (o *DatasetSplittingJobParams) SetSplitAcrossMetadata(v []string)`

SetSplitAcrossMetadata sets SplitAcrossMetadata field to given value.


### GetKeepTogetherMetadata

`func (o *DatasetSplittingJobParams) GetKeepTogetherMetadata() []string`

GetKeepTogetherMetadata returns the KeepTogetherMetadata field if non-nil, zero value otherwise.

### GetKeepTogetherMetadataOk

`func (o *DatasetSplittingJobParams) GetKeepTogetherMetadataOk() (*[]string, bool)`

GetKeepTogetherMetadataOk returns a tuple with the KeepTogetherMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepTogetherMetadata

`func (o *DatasetSplittingJobParams) SetKeepTogetherMetadata(v []string)`

SetKeepTogetherMetadata sets KeepTogetherMetadata field to given value.


### GetSplitsToResplit

`func (o *DatasetSplittingJobParams) GetSplitsToResplit() []SplitSubset`

GetSplitsToResplit returns the SplitsToResplit field if non-nil, zero value otherwise.

### GetSplitsToResplitOk

`func (o *DatasetSplittingJobParams) GetSplitsToResplitOk() (*[]SplitSubset, bool)`

GetSplitsToResplitOk returns a tuple with the SplitsToResplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsToResplit

`func (o *DatasetSplittingJobParams) SetSplitsToResplit(v []SplitSubset)`

SetSplitsToResplit sets SplitsToResplit field to given value.


### GetInferenceArtifactId

`func (o *DatasetSplittingJobParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *DatasetSplittingJobParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *DatasetSplittingJobParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVersionId

`func (o *DatasetSplittingJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *DatasetSplittingJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *DatasetSplittingJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *DatasetSplittingJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DatasetSplittingJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DatasetSplittingJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


