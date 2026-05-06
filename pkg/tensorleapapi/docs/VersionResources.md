# VersionResources

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InferenceArtifactId** | **string** |  | 
**VisArtifactId** | **string** |  | 
**InferenceResources** | [**InferenceArtifactResources**](InferenceArtifactResources.md) |  | 
**VisResources** | [**VisArtifactResources**](VisArtifactResources.md) |  | 
**EsMetricsIndex** | Pointer to **string** |  | [optional] 
**EsInspectionIndex** | Pointer to **string** |  | [optional] 
**EsModelId** | Pointer to **string** |  | [optional] 
**PopulationExplorationDigestSeedCount** | Pointer to **float64** |  | [optional] 
**InsightsIndexCounter** | Pointer to **float64** |  | [optional] 
**CsvBlobPath** | Pointer to **string** |  | [optional] 
**OverridedVersionId** | Pointer to **string** |  | [optional] 

## Methods

### NewVersionResources

`func NewVersionResources(inferenceArtifactId string, visArtifactId string, inferenceResources InferenceArtifactResources, visResources VisArtifactResources, ) *VersionResources`

NewVersionResources instantiates a new VersionResources object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionResourcesWithDefaults

`func NewVersionResourcesWithDefaults() *VersionResources`

NewVersionResourcesWithDefaults instantiates a new VersionResources object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInferenceArtifactId

`func (o *VersionResources) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *VersionResources) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *VersionResources) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVisArtifactId

`func (o *VersionResources) GetVisArtifactId() string`

GetVisArtifactId returns the VisArtifactId field if non-nil, zero value otherwise.

### GetVisArtifactIdOk

`func (o *VersionResources) GetVisArtifactIdOk() (*string, bool)`

GetVisArtifactIdOk returns a tuple with the VisArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisArtifactId

`func (o *VersionResources) SetVisArtifactId(v string)`

SetVisArtifactId sets VisArtifactId field to given value.


### GetInferenceResources

`func (o *VersionResources) GetInferenceResources() InferenceArtifactResources`

GetInferenceResources returns the InferenceResources field if non-nil, zero value otherwise.

### GetInferenceResourcesOk

`func (o *VersionResources) GetInferenceResourcesOk() (*InferenceArtifactResources, bool)`

GetInferenceResourcesOk returns a tuple with the InferenceResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceResources

`func (o *VersionResources) SetInferenceResources(v InferenceArtifactResources)`

SetInferenceResources sets InferenceResources field to given value.


### GetVisResources

`func (o *VersionResources) GetVisResources() VisArtifactResources`

GetVisResources returns the VisResources field if non-nil, zero value otherwise.

### GetVisResourcesOk

`func (o *VersionResources) GetVisResourcesOk() (*VisArtifactResources, bool)`

GetVisResourcesOk returns a tuple with the VisResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisResources

`func (o *VersionResources) SetVisResources(v VisArtifactResources)`

SetVisResources sets VisResources field to given value.


### GetEsMetricsIndex

`func (o *VersionResources) GetEsMetricsIndex() string`

GetEsMetricsIndex returns the EsMetricsIndex field if non-nil, zero value otherwise.

### GetEsMetricsIndexOk

`func (o *VersionResources) GetEsMetricsIndexOk() (*string, bool)`

GetEsMetricsIndexOk returns a tuple with the EsMetricsIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsMetricsIndex

`func (o *VersionResources) SetEsMetricsIndex(v string)`

SetEsMetricsIndex sets EsMetricsIndex field to given value.

### HasEsMetricsIndex

`func (o *VersionResources) HasEsMetricsIndex() bool`

HasEsMetricsIndex returns a boolean if a field has been set.

### GetEsInspectionIndex

`func (o *VersionResources) GetEsInspectionIndex() string`

GetEsInspectionIndex returns the EsInspectionIndex field if non-nil, zero value otherwise.

### GetEsInspectionIndexOk

`func (o *VersionResources) GetEsInspectionIndexOk() (*string, bool)`

GetEsInspectionIndexOk returns a tuple with the EsInspectionIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsInspectionIndex

`func (o *VersionResources) SetEsInspectionIndex(v string)`

SetEsInspectionIndex sets EsInspectionIndex field to given value.

### HasEsInspectionIndex

`func (o *VersionResources) HasEsInspectionIndex() bool`

HasEsInspectionIndex returns a boolean if a field has been set.

### GetEsModelId

`func (o *VersionResources) GetEsModelId() string`

GetEsModelId returns the EsModelId field if non-nil, zero value otherwise.

### GetEsModelIdOk

`func (o *VersionResources) GetEsModelIdOk() (*string, bool)`

GetEsModelIdOk returns a tuple with the EsModelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEsModelId

`func (o *VersionResources) SetEsModelId(v string)`

SetEsModelId sets EsModelId field to given value.

### HasEsModelId

`func (o *VersionResources) HasEsModelId() bool`

HasEsModelId returns a boolean if a field has been set.

### GetPopulationExplorationDigestSeedCount

`func (o *VersionResources) GetPopulationExplorationDigestSeedCount() float64`

GetPopulationExplorationDigestSeedCount returns the PopulationExplorationDigestSeedCount field if non-nil, zero value otherwise.

### GetPopulationExplorationDigestSeedCountOk

`func (o *VersionResources) GetPopulationExplorationDigestSeedCountOk() (*float64, bool)`

GetPopulationExplorationDigestSeedCountOk returns a tuple with the PopulationExplorationDigestSeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPopulationExplorationDigestSeedCount

`func (o *VersionResources) SetPopulationExplorationDigestSeedCount(v float64)`

SetPopulationExplorationDigestSeedCount sets PopulationExplorationDigestSeedCount field to given value.

### HasPopulationExplorationDigestSeedCount

`func (o *VersionResources) HasPopulationExplorationDigestSeedCount() bool`

HasPopulationExplorationDigestSeedCount returns a boolean if a field has been set.

### GetInsightsIndexCounter

`func (o *VersionResources) GetInsightsIndexCounter() float64`

GetInsightsIndexCounter returns the InsightsIndexCounter field if non-nil, zero value otherwise.

### GetInsightsIndexCounterOk

`func (o *VersionResources) GetInsightsIndexCounterOk() (*float64, bool)`

GetInsightsIndexCounterOk returns a tuple with the InsightsIndexCounter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightsIndexCounter

`func (o *VersionResources) SetInsightsIndexCounter(v float64)`

SetInsightsIndexCounter sets InsightsIndexCounter field to given value.

### HasInsightsIndexCounter

`func (o *VersionResources) HasInsightsIndexCounter() bool`

HasInsightsIndexCounter returns a boolean if a field has been set.

### GetCsvBlobPath

`func (o *VersionResources) GetCsvBlobPath() string`

GetCsvBlobPath returns the CsvBlobPath field if non-nil, zero value otherwise.

### GetCsvBlobPathOk

`func (o *VersionResources) GetCsvBlobPathOk() (*string, bool)`

GetCsvBlobPathOk returns a tuple with the CsvBlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvBlobPath

`func (o *VersionResources) SetCsvBlobPath(v string)`

SetCsvBlobPath sets CsvBlobPath field to given value.

### HasCsvBlobPath

`func (o *VersionResources) HasCsvBlobPath() bool`

HasCsvBlobPath returns a boolean if a field has been set.

### GetOverridedVersionId

`func (o *VersionResources) GetOverridedVersionId() string`

GetOverridedVersionId returns the OverridedVersionId field if non-nil, zero value otherwise.

### GetOverridedVersionIdOk

`func (o *VersionResources) GetOverridedVersionIdOk() (*string, bool)`

GetOverridedVersionIdOk returns a tuple with the OverridedVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverridedVersionId

`func (o *VersionResources) SetOverridedVersionId(v string)`

SetOverridedVersionId sets OverridedVersionId field to given value.

### HasOverridedVersionId

`func (o *VersionResources) HasOverridedVersionId() bool`

HasOverridedVersionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


