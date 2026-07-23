# AutoSyntheticDataJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**InitialSimulationFilters** | Pointer to [**[]ESFilter**](ESFilter.md) |  | [optional] 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**SimulationNames** | **[]string** |  | 
**InferenceArtifactId** | **string** |  | 
**VersionId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewAutoSyntheticDataJobParams

`func NewAutoSyntheticDataJobParams(digest string, targetFilters []ESFilter, simulationNames []string, inferenceArtifactId string, versionId string, type_ string, ) *AutoSyntheticDataJobParams`

NewAutoSyntheticDataJobParams instantiates a new AutoSyntheticDataJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoSyntheticDataJobParamsWithDefaults

`func NewAutoSyntheticDataJobParamsWithDefaults() *AutoSyntheticDataJobParams`

NewAutoSyntheticDataJobParamsWithDefaults instantiates a new AutoSyntheticDataJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *AutoSyntheticDataJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AutoSyntheticDataJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AutoSyntheticDataJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetInitialSimulationFilters

`func (o *AutoSyntheticDataJobParams) GetInitialSimulationFilters() []ESFilter`

GetInitialSimulationFilters returns the InitialSimulationFilters field if non-nil, zero value otherwise.

### GetInitialSimulationFiltersOk

`func (o *AutoSyntheticDataJobParams) GetInitialSimulationFiltersOk() (*[]ESFilter, bool)`

GetInitialSimulationFiltersOk returns a tuple with the InitialSimulationFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialSimulationFilters

`func (o *AutoSyntheticDataJobParams) SetInitialSimulationFilters(v []ESFilter)`

SetInitialSimulationFilters sets InitialSimulationFilters field to given value.

### HasInitialSimulationFilters

`func (o *AutoSyntheticDataJobParams) HasInitialSimulationFilters() bool`

HasInitialSimulationFilters returns a boolean if a field has been set.

### GetTargetFilters

`func (o *AutoSyntheticDataJobParams) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *AutoSyntheticDataJobParams) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *AutoSyntheticDataJobParams) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.


### GetSimulationNames

`func (o *AutoSyntheticDataJobParams) GetSimulationNames() []string`

GetSimulationNames returns the SimulationNames field if non-nil, zero value otherwise.

### GetSimulationNamesOk

`func (o *AutoSyntheticDataJobParams) GetSimulationNamesOk() (*[]string, bool)`

GetSimulationNamesOk returns a tuple with the SimulationNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulationNames

`func (o *AutoSyntheticDataJobParams) SetSimulationNames(v []string)`

SetSimulationNames sets SimulationNames field to given value.


### GetInferenceArtifactId

`func (o *AutoSyntheticDataJobParams) GetInferenceArtifactId() string`

GetInferenceArtifactId returns the InferenceArtifactId field if non-nil, zero value otherwise.

### GetInferenceArtifactIdOk

`func (o *AutoSyntheticDataJobParams) GetInferenceArtifactIdOk() (*string, bool)`

GetInferenceArtifactIdOk returns a tuple with the InferenceArtifactId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInferenceArtifactId

`func (o *AutoSyntheticDataJobParams) SetInferenceArtifactId(v string)`

SetInferenceArtifactId sets InferenceArtifactId field to given value.


### GetVersionId

`func (o *AutoSyntheticDataJobParams) GetVersionId() string`

GetVersionId returns the VersionId field if non-nil, zero value otherwise.

### GetVersionIdOk

`func (o *AutoSyntheticDataJobParams) GetVersionIdOk() (*string, bool)`

GetVersionIdOk returns a tuple with the VersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionId

`func (o *AutoSyntheticDataJobParams) SetVersionId(v string)`

SetVersionId sets VersionId field to given value.


### GetType

`func (o *AutoSyntheticDataJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AutoSyntheticDataJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AutoSyntheticDataJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


