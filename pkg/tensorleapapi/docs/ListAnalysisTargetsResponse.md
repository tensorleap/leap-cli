# ListAnalysisTargetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContractVersion** | **float64** |  | 
**Me** | [**ListAnalysisTargetsResponseMe**](ListAnalysisTargetsResponseMe.md) |  | 
**Projects** | Pointer to [**[]AnalysisTargetProject**](AnalysisTargetProject.md) |  | [optional] 
**Versions** | Pointer to [**[]AnalysisTargetVersion**](AnalysisTargetVersion.md) |  | [optional] 

## Methods

### NewListAnalysisTargetsResponse

`func NewListAnalysisTargetsResponse(contractVersion float64, me ListAnalysisTargetsResponseMe, ) *ListAnalysisTargetsResponse`

NewListAnalysisTargetsResponse instantiates a new ListAnalysisTargetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAnalysisTargetsResponseWithDefaults

`func NewListAnalysisTargetsResponseWithDefaults() *ListAnalysisTargetsResponse`

NewListAnalysisTargetsResponseWithDefaults instantiates a new ListAnalysisTargetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContractVersion

`func (o *ListAnalysisTargetsResponse) GetContractVersion() float64`

GetContractVersion returns the ContractVersion field if non-nil, zero value otherwise.

### GetContractVersionOk

`func (o *ListAnalysisTargetsResponse) GetContractVersionOk() (*float64, bool)`

GetContractVersionOk returns a tuple with the ContractVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractVersion

`func (o *ListAnalysisTargetsResponse) SetContractVersion(v float64)`

SetContractVersion sets ContractVersion field to given value.


### GetMe

`func (o *ListAnalysisTargetsResponse) GetMe() ListAnalysisTargetsResponseMe`

GetMe returns the Me field if non-nil, zero value otherwise.

### GetMeOk

`func (o *ListAnalysisTargetsResponse) GetMeOk() (*ListAnalysisTargetsResponseMe, bool)`

GetMeOk returns a tuple with the Me field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMe

`func (o *ListAnalysisTargetsResponse) SetMe(v ListAnalysisTargetsResponseMe)`

SetMe sets Me field to given value.


### GetProjects

`func (o *ListAnalysisTargetsResponse) GetProjects() []AnalysisTargetProject`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *ListAnalysisTargetsResponse) GetProjectsOk() (*[]AnalysisTargetProject, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *ListAnalysisTargetsResponse) SetProjects(v []AnalysisTargetProject)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *ListAnalysisTargetsResponse) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetVersions

`func (o *ListAnalysisTargetsResponse) GetVersions() []AnalysisTargetVersion`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListAnalysisTargetsResponse) GetVersionsOk() (*[]AnalysisTargetVersion, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListAnalysisTargetsResponse) SetVersions(v []AnalysisTargetVersion)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ListAnalysisTargetsResponse) HasVersions() bool`

HasVersions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


