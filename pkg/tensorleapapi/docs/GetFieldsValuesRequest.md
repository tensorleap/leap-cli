# GetFieldsValuesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filters** | [**[]ESFilter**](ESFilter.md) |  | 
**VersionIds** | Pointer to **[]string** |  | [optional] 
**Fields** | [**[]QueryFieldValues**](QueryFieldValues.md) |  | 
**ProjectId** | **string** |  | 

## Methods

### NewGetFieldsValuesRequest

`func NewGetFieldsValuesRequest(filters []ESFilter, fields []QueryFieldValues, projectId string, ) *GetFieldsValuesRequest`

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


### GetVersionIds

`func (o *GetFieldsValuesRequest) GetVersionIds() []string`

GetVersionIds returns the VersionIds field if non-nil, zero value otherwise.

### GetVersionIdsOk

`func (o *GetFieldsValuesRequest) GetVersionIdsOk() (*[]string, bool)`

GetVersionIdsOk returns a tuple with the VersionIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionIds

`func (o *GetFieldsValuesRequest) SetVersionIds(v []string)`

SetVersionIds sets VersionIds field to given value.

### HasVersionIds

`func (o *GetFieldsValuesRequest) HasVersionIds() bool`

HasVersionIds returns a boolean if a field has been set.

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


