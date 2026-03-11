# InsightsJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**InferenceArtifactId** | Pointer to **string** |  | [optional] 
**VersionId** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewInsightsJobParams

`func NewInsightsJobParams(type_ string, ) *InsightsJobParams`

NewInsightsJobParams instantiates a new InsightsJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightsJobParamsWithDefaults

`func NewInsightsJobParamsWithDefaults() *InsightsJobParams`

NewInsightsJobParamsWithDefaults instantiates a new InsightsJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *InsightsJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *InsightsJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *InsightsJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *InsightsJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetInferenceArtifactId

`func (o *InsightsJobParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *InsightsJobParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *InsightsJobParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.

### HasInferenceArtifactId

`func (o *InsightsJobParams) HasInferenceArtifactId() bool`

HasInferenceArtifactId returns a boolean if a field has been set.

### GetVersionId

`func (o *InsightsJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *InsightsJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *InsightsJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.

### HasVersionId

`func (o *InsightsJobParams) HasVersionId() bool`

HasVersionId returns a boolean if a field has been set.

### GetType

`func (o *InsightsJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InsightsJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InsightsJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


