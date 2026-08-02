# GenerateDatasetSplittingParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** |  | 
**VersionId** | **string** |  | 
**SplitsToResplit** | [**[]SplitSubset**](SplitSubset.md) |  | 
**KeepTogetherMetadata** | **[]string** |  | 
**SplitAcrossMetadata** | **[]string** |  | 

## Methods

### NewGenerateDatasetSplittingParams

`func NewGenerateDatasetSplittingParams(projectId string, versionId string, splitsToResplit []SplitSubset, keepTogetherMetadata []string, splitAcrossMetadata []string, ) *GenerateDatasetSplittingParams`

NewGenerateDatasetSplittingParams instantiates a new GenerateDatasetSplittingParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateDatasetSplittingParamsWithDefaults

`func NewGenerateDatasetSplittingParamsWithDefaults() *GenerateDatasetSplittingParams`

NewGenerateDatasetSplittingParamsWithDefaults instantiates a new GenerateDatasetSplittingParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *GenerateDatasetSplittingParams) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GenerateDatasetSplittingParams) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GenerateDatasetSplittingParams) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetVersionId

`func (o *GenerateDatasetSplittingParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *GenerateDatasetSplittingParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *GenerateDatasetSplittingParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetSplitsToResplit

`func (o *GenerateDatasetSplittingParams) GetSplitsToResplit() []SplitSubset`

GetSplitsToResplit returns the SplitsToResplit field if non-nil, zero value otherwise.

### GetSplitsToResplitOk

`func (o *GenerateDatasetSplittingParams) GetSplitsToResplitOk() (*[]SplitSubset, bool)`

GetSplitsToResplitOk returns a tuple with the SplitsToResplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitsToResplit

`func (o *GenerateDatasetSplittingParams) SetSplitsToResplit(v []SplitSubset)`

SetSplitsToResplit sets SplitsToResplit field to given value.


### GetKeepTogetherMetadata

`func (o *GenerateDatasetSplittingParams) GetKeepTogetherMetadata() []string`

GetKeepTogetherMetadata returns the KeepTogetherMetadata field if non-nil, zero value otherwise.

### GetKeepTogetherMetadataOk

`func (o *GenerateDatasetSplittingParams) GetKeepTogetherMetadataOk() (*[]string, bool)`

GetKeepTogetherMetadataOk returns a tuple with the KeepTogetherMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepTogetherMetadata

`func (o *GenerateDatasetSplittingParams) SetKeepTogetherMetadata(v []string)`

SetKeepTogetherMetadata sets KeepTogetherMetadata field to given value.


### GetSplitAcrossMetadata

`func (o *GenerateDatasetSplittingParams) GetSplitAcrossMetadata() []string`

GetSplitAcrossMetadata returns the SplitAcrossMetadata field if non-nil, zero value otherwise.

### GetSplitAcrossMetadataOk

`func (o *GenerateDatasetSplittingParams) GetSplitAcrossMetadataOk() (*[]string, bool)`

GetSplitAcrossMetadataOk returns a tuple with the SplitAcrossMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitAcrossMetadata

`func (o *GenerateDatasetSplittingParams) SetSplitAcrossMetadata(v []string)`

SetSplitAcrossMetadata sets SplitAcrossMetadata field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


