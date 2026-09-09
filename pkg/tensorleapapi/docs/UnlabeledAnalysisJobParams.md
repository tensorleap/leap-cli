# UnlabeledAnalysisJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**ElementInstance** | Pointer to **bool** |  | [optional] 
**LatentSpaceType** | Pointer to **string** |  | [optional] 
**Filters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**VisArtifactId** | **string** |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewUnlabeledAnalysisJobParams

`func NewUnlabeledAnalysisJobParams(digest string, visArtifactId string, versionId string, type_ string, ) *UnlabeledAnalysisJobParams`

NewUnlabeledAnalysisJobParams instantiates a new UnlabeledAnalysisJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnlabeledAnalysisJobParamsWithDefaults

`func NewUnlabeledAnalysisJobParamsWithDefaults() *UnlabeledAnalysisJobParams`

NewUnlabeledAnalysisJobParamsWithDefaults instantiates a new UnlabeledAnalysisJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *UnlabeledAnalysisJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *UnlabeledAnalysisJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *UnlabeledAnalysisJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetElementInstance

`func (o *UnlabeledAnalysisJobParams) GetElementInstance() bool`

GetElementInstance returns the ElementInstance field if non-nil, zero value otherwise.

### GetElementInstanceOk

`func (o *UnlabeledAnalysisJobParams) GetElementInstanceOk() (*bool, bool)`

GetElementInstanceOk returns a tuple with the ElementInstance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementInstance

`func (o *UnlabeledAnalysisJobParams) SetElementInstance(v bool)`

SetElementInstance sets ElementInstance field to given value.

### HasElementInstance

`func (o *UnlabeledAnalysisJobParams) HasElementInstance() bool`

HasElementInstance returns a boolean if a field has been set.

### GetLatentSpaceType

`func (o *UnlabeledAnalysisJobParams) GetLatentSpaceType() string`

GetLatentSpaceType returns the LatentSpaceType field if non-nil, zero value otherwise.

### GetLatentSpaceTypeOk

`func (o *UnlabeledAnalysisJobParams) GetLatentSpaceTypeOk() (*string, bool)`

GetLatentSpaceTypeOk returns a tuple with the LatentSpaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatentSpaceType

`func (o *UnlabeledAnalysisJobParams) SetLatentSpaceType(v string)`

SetLatentSpaceType sets LatentSpaceType field to given value.

### HasLatentSpaceType

`func (o *UnlabeledAnalysisJobParams) HasLatentSpaceType() bool`

HasLatentSpaceType returns a boolean if a field has been set.

### GetFilters

`func (o *UnlabeledAnalysisJobParams) GetFilters() []ESFilter`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *UnlabeledAnalysisJobParams) GetFiltersOk() (*[]ESFilter, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *UnlabeledAnalysisJobParams) SetFilters(v []ESFilter)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *UnlabeledAnalysisJobParams) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### GetVisArtifactId

`func (o *UnlabeledAnalysisJobParams) GetVisArtifactId() string`

GetVisArtifactId returns the VisArtifactId field if non-nil, zero value otherwise.

### GetVisArtifactIdOk

`func (o *UnlabeledAnalysisJobParams) GetVisArtifactIdOk() (*string, bool)`

GetVisArtifactIdOk returns a tuple with the VisArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisArtifactId

`func (o *UnlabeledAnalysisJobParams) SetVisArtifactId(v string)`

SetVisArtifactId sets VisArtifactId field to given value.


### GetVersionId

`func (o *UnlabeledAnalysisJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *UnlabeledAnalysisJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *UnlabeledAnalysisJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *UnlabeledAnalysisJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UnlabeledAnalysisJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UnlabeledAnalysisJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


