# AnalysisInsightDigest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupSize** | **float64** |  | 
**GroupDefinition** | **string** |  | 
**CsvRows** | **float64** |  | 
**Split** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**Composition** | [**[]AnalysisComposition**](AnalysisComposition.md) |  | 
**Contrast** | [**[]AnalysisContrast**](AnalysisContrast.md) |  | 
**RankedBy** | Pointer to **string** |  | [optional] 
**RankedSampleIds** | **[]string** |  | 

## Methods

### NewAnalysisInsightDigest

`func NewAnalysisInsightDigest(groupSize float64, groupDefinition string, csvRows float64, split map[string]interface{}, composition []AnalysisComposition, contrast []AnalysisContrast, rankedSampleIds []string, ) *AnalysisInsightDigest`

NewAnalysisInsightDigest instantiates a new AnalysisInsightDigest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisInsightDigestWithDefaults

`func NewAnalysisInsightDigestWithDefaults() *AnalysisInsightDigest`

NewAnalysisInsightDigestWithDefaults instantiates a new AnalysisInsightDigest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupSize

`func (o *AnalysisInsightDigest) GetGroupSize() float64`

GetGroupSize returns the GroupSize field if non-nil, zero value otherwise.

### GetGroupSizeOk

`func (o *AnalysisInsightDigest) GetGroupSizeOk() (*float64, bool)`

GetGroupSizeOk returns a tuple with the GroupSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupSize

`func (o *AnalysisInsightDigest) SetGroupSize(v float64)`

SetGroupSize sets GroupSize field to given value.


### GetGroupDefinition

`func (o *AnalysisInsightDigest) GetGroupDefinition() string`

GetGroupDefinition returns the GroupDefinition field if non-nil, zero value otherwise.

### GetGroupDefinitionOk

`func (o *AnalysisInsightDigest) GetGroupDefinitionOk() (*string, bool)`

GetGroupDefinitionOk returns a tuple with the GroupDefinition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDefinition

`func (o *AnalysisInsightDigest) SetGroupDefinition(v string)`

SetGroupDefinition sets GroupDefinition field to given value.


### GetCsvRows

`func (o *AnalysisInsightDigest) GetCsvRows() float64`

GetCsvRows returns the CsvRows field if non-nil, zero value otherwise.

### GetCsvRowsOk

`func (o *AnalysisInsightDigest) GetCsvRowsOk() (*float64, bool)`

GetCsvRowsOk returns a tuple with the CsvRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvRows

`func (o *AnalysisInsightDigest) SetCsvRows(v float64)`

SetCsvRows sets CsvRows field to given value.


### GetSplit

`func (o *AnalysisInsightDigest) GetSplit() map[string]interface{}`

GetSplit returns the Split field if non-nil, zero value otherwise.

### GetSplitOk

`func (o *AnalysisInsightDigest) GetSplitOk() (*map[string]interface{}, bool)`

GetSplitOk returns a tuple with the Split field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplit

`func (o *AnalysisInsightDigest) SetSplit(v map[string]interface{})`

SetSplit sets Split field to given value.


### GetComposition

`func (o *AnalysisInsightDigest) GetComposition() []AnalysisComposition`

GetComposition returns the Composition field if non-nil, zero value otherwise.

### GetCompositionOk

`func (o *AnalysisInsightDigest) GetCompositionOk() (*[]AnalysisComposition, bool)`

GetCompositionOk returns a tuple with the Composition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposition

`func (o *AnalysisInsightDigest) SetComposition(v []AnalysisComposition)`

SetComposition sets Composition field to given value.


### GetContrast

`func (o *AnalysisInsightDigest) GetContrast() []AnalysisContrast`

GetContrast returns the Contrast field if non-nil, zero value otherwise.

### GetContrastOk

`func (o *AnalysisInsightDigest) GetContrastOk() (*[]AnalysisContrast, bool)`

GetContrastOk returns a tuple with the Contrast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContrast

`func (o *AnalysisInsightDigest) SetContrast(v []AnalysisContrast)`

SetContrast sets Contrast field to given value.


### GetRankedBy

`func (o *AnalysisInsightDigest) GetRankedBy() string`

GetRankedBy returns the RankedBy field if non-nil, zero value otherwise.

### GetRankedByOk

`func (o *AnalysisInsightDigest) GetRankedByOk() (*string, bool)`

GetRankedByOk returns a tuple with the RankedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedBy

`func (o *AnalysisInsightDigest) SetRankedBy(v string)`

SetRankedBy sets RankedBy field to given value.

### HasRankedBy

`func (o *AnalysisInsightDigest) HasRankedBy() bool`

HasRankedBy returns a boolean if a field has been set.

### GetRankedSampleIds

`func (o *AnalysisInsightDigest) GetRankedSampleIds() []string`

GetRankedSampleIds returns the RankedSampleIds field if non-nil, zero value otherwise.

### GetRankedSampleIdsOk

`func (o *AnalysisInsightDigest) GetRankedSampleIdsOk() (*[]string, bool)`

GetRankedSampleIdsOk returns a tuple with the RankedSampleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRankedSampleIds

`func (o *AnalysisInsightDigest) SetRankedSampleIds(v []string)`

SetRankedSampleIds sets RankedSampleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


