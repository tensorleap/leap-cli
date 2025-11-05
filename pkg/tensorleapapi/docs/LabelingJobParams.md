# LabelingJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UseCustomLatentSpace** | Pointer to **bool** |  | [optional] 
**Digest** | **string** |  | 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**LabelingAlgorithm** | [**LabelingAlgorithm**](LabelingAlgorithm.md) |  | 
**NumOfSamplesToLabel** | Pointer to **float64** |  | [optional] 
**FromEpoch** | **float64** |  | 
**SessionRunId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewLabelingJobParams

`func NewLabelingJobParams(digest string, labelingAlgorithm LabelingAlgorithm, fromEpoch float64, sessionRunId string, type_ string, ) *LabelingJobParams`

NewLabelingJobParams instantiates a new LabelingJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLabelingJobParamsWithDefaults

`func NewLabelingJobParamsWithDefaults() *LabelingJobParams`

NewLabelingJobParamsWithDefaults instantiates a new LabelingJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUseCustomLatentSpace

`func (o *LabelingJobParams) GetUseCustomLatentSpace() bool`

GetUseCustomLatentSpace returns the UseCustomLatentSpace field if non-nil, zero value otherwise.

### GetUseCustomLatentSpaceOk

`func (o *LabelingJobParams) GetUseCustomLatentSpaceOk() (*bool, bool)`

GetUseCustomLatentSpaceOk returns a tuple with the UseCustomLatentSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCustomLatentSpace

`func (o *LabelingJobParams) SetUseCustomLatentSpace(v bool)`

SetUseCustomLatentSpace sets UseCustomLatentSpace field to given value.

### HasUseCustomLatentSpace

`func (o *LabelingJobParams) HasUseCustomLatentSpace() bool`

HasUseCustomLatentSpace returns a boolean if a field has been set.

### GetDigest

`func (o *LabelingJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *LabelingJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *LabelingJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFilters

`func (o *LabelingJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *LabelingJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *LabelingJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *LabelingJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetLabelingAlgorithm

`func (o *LabelingJobParams) GetLabelingAlgorithm() LabelingAlgorithm`

GetLabelingAlgorithm returns the LabelingAlgorithm field if non-nil, zero value otherwise.

### GetLabelingAlgorithmOk

`func (o *LabelingJobParams) GetLabelingAlgorithmOk() (*LabelingAlgorithm, bool)`

GetLabelingAlgorithmOk returns a tuple with the LabelingAlgorithm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelingAlgorithm

`func (o *LabelingJobParams) SetLabelingAlgorithm(v LabelingAlgorithm)`

SetLabelingAlgorithm sets LabelingAlgorithm field to given value.


### GetNumOfSamplesToLabel

`func (o *LabelingJobParams) GetNumOfSamplesToLabel() float64`

GetNumOfSamplesToLabel returns the NumOfSamplesToLabel field if non-nil, zero value otherwise.

### GetNumOfSamplesToLabelOk

`func (o *LabelingJobParams) GetNumOfSamplesToLabelOk() (*float64, bool)`

GetNumOfSamplesToLabelOk returns a tuple with the NumOfSamplesToLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumOfSamplesToLabel

`func (o *LabelingJobParams) SetNumOfSamplesToLabel(v float64)`

SetNumOfSamplesToLabel sets NumOfSamplesToLabel field to given value.

### HasNumOfSamplesToLabel

`func (o *LabelingJobParams) HasNumOfSamplesToLabel() bool`

HasNumOfSamplesToLabel returns a boolean if a field has been set.

### GetFromEpoch

`func (o *LabelingJobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *LabelingJobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *LabelingJobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetSessionRunId

`func (o *LabelingJobParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *LabelingJobParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *LabelingJobParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetType

`func (o *LabelingJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LabelingJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LabelingJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


