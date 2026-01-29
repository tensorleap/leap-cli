# SyntheticDataJobParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Digest** | **string** |  | 
**TargetFilters** | [**[]ESFilter**](ESFilter.md) |  | 
**Sources** | [**[]GenerateSyntheticDataParamsSourcesInner**](GenerateSyntheticDataParamsSourcesInner.md) |  | 
**FromEpoch** | **float64** |  | 
**SessionRunId** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewSyntheticDataJobParams

`func NewSyntheticDataJobParams(digest string, targetFilters []ESFilter, sources []GenerateSyntheticDataParamsSourcesInner, fromEpoch float64, sessionRunId string, type_ string, ) *SyntheticDataJobParams`

NewSyntheticDataJobParams instantiates a new SyntheticDataJobParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSyntheticDataJobParamsWithDefaults

`func NewSyntheticDataJobParamsWithDefaults() *SyntheticDataJobParams`

NewSyntheticDataJobParamsWithDefaults instantiates a new SyntheticDataJobParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDigest

`func (o *SyntheticDataJobParams) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *SyntheticDataJobParams) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *SyntheticDataJobParams) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetTargetFilters

`func (o *SyntheticDataJobParams) GetTargetFilters() []ESFilter`

GetTargetFilters returns the TargetFilters field if non-nil, zero value otherwise.

### GetTargetFiltersOk

`func (o *SyntheticDataJobParams) GetTargetFiltersOk() (*[]ESFilter, bool)`

GetTargetFiltersOk returns a tuple with the TargetFilters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetFilters

`func (o *SyntheticDataJobParams) SetTargetFilters(v []ESFilter)`

SetTargetFilters sets TargetFilters field to given value.


### GetSources

`func (o *SyntheticDataJobParams) GetSources() []GenerateSyntheticDataParamsSourcesInner`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *SyntheticDataJobParams) GetSourcesOk() (*[]GenerateSyntheticDataParamsSourcesInner, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *SyntheticDataJobParams) SetSources(v []GenerateSyntheticDataParamsSourcesInner)`

SetSources sets Sources field to given value.


### GetFromEpoch

`func (o *SyntheticDataJobParams) GetFromEpoch() float64`

GetFromEpoch returns the FromEpoch field if non-nil, zero value otherwise.

### GetFromEpochOk

`func (o *SyntheticDataJobParams) GetFromEpochOk() (*float64, bool)`

GetFromEpochOk returns a tuple with the FromEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromEpoch

`func (o *SyntheticDataJobParams) SetFromEpoch(v float64)`

SetFromEpoch sets FromEpoch field to given value.


### GetSessionRunId

`func (o *SyntheticDataJobParams) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *SyntheticDataJobParams) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *SyntheticDataJobParams) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetType

`func (o *SyntheticDataJobParams) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SyntheticDataJobParams) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SyntheticDataJobParams) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


