# GetFieldsValuesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]ESFilter**](ESFilter.md) |  | 
**InferenceArtifactIds** | **[]string** |  | 
**Fields** | [**[]QueryFieldValues**](QueryFieldValues.md) |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetFieldsValuesRequest

`func NewGetFieldsValuesRequest(filters []ESFilter, inferenceArtifactIds []string, fields []QueryFieldValues, projectId string, ) *GetFieldsValuesRequest`

NewGetFieldsValuesRequest instantiates a new GetFieldsValuesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFieldsValuesRequestWithDefaults

`func NewGetFieldsValuesRequestWithDefaults() *GetFieldsValuesRequest`

NewGetFieldsValuesRequestWithDefaults instantiates a new GetFieldsValuesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilters

`func (o *GetFieldsValuesRequest) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetFieldsValuesRequest) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetFieldsValuesRequest) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.


### GetInferenceArtifactIds

`func (o *GetFieldsValuesRequest) GetInferenceArtifactIds() []string`

GetInferenceArtifactIds returns the InferenceArtifactIds field if non-nil, zero value otherwise.

### GetInferenceArtifactIdsOk

`func (o *GetFieldsValuesRequest) GetInferenceArtifactIdsOk() (*[]string, bool)`

GetInferenceArtifactIdsOk returns a tuple with the InferenceArtifactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactIds

`func (o *GetFieldsValuesRequest) SetInferenceArtifactIds(v []string)`

SetInferenceArtifactIds sets InferenceArtifactIds field to given value.


### GetFields

`func (o *GetFieldsValuesRequest) GetFields() []QueryFieldValues`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GetFieldsValuesRequest) GetFieldsOk() (*[]QueryFieldValues, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GetFieldsValuesRequest) SetFields(v []QueryFieldValues)`

SetFields sets Fields field to given value.


### GetProjectId

`func (o *GetFieldsValuesRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *GetFieldsValuesRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *GetFieldsValuesRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


