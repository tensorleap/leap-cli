# AnalysisInsight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Index** | **float64** |  | 
**Status** | [**AnalysisInsightStatus**](AnalysisInsightStatus.md) |  | 
**Description** | Pointer to **string** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**InsightType** | [**InsightType**](InsightType.md) |  | 
**CsvUrl** | Pointer to **string** |  | [optional] 
**ClusterBlobUrl** | Pointer to **string** |  | [optional] 
**TopPanelUrl** | Pointer to **string** |  | [optional] 
**FixingCsvUrl** | Pointer to **string** |  | [optional] 
**Engine** | **map[string]interface{}** | Construct a type with a set of properties K of type T | 
**Digest** | Pointer to [**AnalysisInsightDigest**](AnalysisInsightDigest.md) |  | [optional] 
**DigestError** | Pointer to **string** |  | [optional] 
**TopPanel** | Pointer to **map[string]interface{}** | Construct a type with a set of properties K of type T | [optional] 
**AnalyzeLinkPath** | Pointer to **string** |  | [optional] 

## Methods

### NewAnalysisInsight

`func NewAnalysisInsight(cid string, index float64, status AnalysisInsightStatus, createdAt time.Time, updatedAt time.Time, insightType InsightType, engine map[string]interface{}, ) *AnalysisInsight`

NewAnalysisInsight instantiates a new AnalysisInsight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisInsightWithDefaults

`func NewAnalysisInsightWithDefaults() *AnalysisInsight`

NewAnalysisInsightWithDefaults instantiates a new AnalysisInsight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *AnalysisInsight) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *AnalysisInsight) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *AnalysisInsight) SetCid(v string)`

SetCid sets Cid field to given value.


### GetIndex

`func (o *AnalysisInsight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *AnalysisInsight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *AnalysisInsight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetStatus

`func (o *AnalysisInsight) GetStatus() AnalysisInsightStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AnalysisInsight) GetStatusOk() (*AnalysisInsightStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AnalysisInsight) SetStatus(v AnalysisInsightStatus)`

SetStatus sets Status field to given value.


### GetDescription

`func (o *AnalysisInsight) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AnalysisInsight) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AnalysisInsight) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AnalysisInsight) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalysisInsight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalysisInsight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalysisInsight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *AnalysisInsight) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AnalysisInsight) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AnalysisInsight) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetInsightType

`func (o *AnalysisInsight) GetInsightType() InsightType`

GetInsightType returns the InsightType field if non-nil, zero value otherwise.

### GetInsightTypeOk

`func (o *AnalysisInsight) GetInsightTypeOk() (*InsightType, bool)`

GetInsightTypeOk returns a tuple with the InsightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightType

`func (o *AnalysisInsight) SetInsightType(v InsightType)`

SetInsightType sets InsightType field to given value.


### GetCsvUrl

`func (o *AnalysisInsight) GetCsvUrl() string`

GetCsvUrl returns the CsvUrl field if non-nil, zero value otherwise.

### GetCsvUrlOk

`func (o *AnalysisInsight) GetCsvUrlOk() (*string, bool)`

GetCsvUrlOk returns a tuple with the CsvUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvUrl

`func (o *AnalysisInsight) SetCsvUrl(v string)`

SetCsvUrl sets CsvUrl field to given value.

### HasCsvUrl

`func (o *AnalysisInsight) HasCsvUrl() bool`

HasCsvUrl returns a boolean if a field has been set.

### GetClusterBlobUrl

`func (o *AnalysisInsight) GetClusterBlobUrl() string`

GetClusterBlobUrl returns the ClusterBlobUrl field if non-nil, zero value otherwise.

### GetClusterBlobUrlOk

`func (o *AnalysisInsight) GetClusterBlobUrlOk() (*string, bool)`

GetClusterBlobUrlOk returns a tuple with the ClusterBlobUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterBlobUrl

`func (o *AnalysisInsight) SetClusterBlobUrl(v string)`

SetClusterBlobUrl sets ClusterBlobUrl field to given value.

### HasClusterBlobUrl

`func (o *AnalysisInsight) HasClusterBlobUrl() bool`

HasClusterBlobUrl returns a boolean if a field has been set.

### GetTopPanelUrl

`func (o *AnalysisInsight) GetTopPanelUrl() string`

GetTopPanelUrl returns the TopPanelUrl field if non-nil, zero value otherwise.

### GetTopPanelUrlOk

`func (o *AnalysisInsight) GetTopPanelUrlOk() (*string, bool)`

GetTopPanelUrlOk returns a tuple with the TopPanelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPanelUrl

`func (o *AnalysisInsight) SetTopPanelUrl(v string)`

SetTopPanelUrl sets TopPanelUrl field to given value.

### HasTopPanelUrl

`func (o *AnalysisInsight) HasTopPanelUrl() bool`

HasTopPanelUrl returns a boolean if a field has been set.

### GetFixingCsvUrl

`func (o *AnalysisInsight) GetFixingCsvUrl() string`

GetFixingCsvUrl returns the FixingCsvUrl field if non-nil, zero value otherwise.

### GetFixingCsvUrlOk

`func (o *AnalysisInsight) GetFixingCsvUrlOk() (*string, bool)`

GetFixingCsvUrlOk returns a tuple with the FixingCsvUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixingCsvUrl

`func (o *AnalysisInsight) SetFixingCsvUrl(v string)`

SetFixingCsvUrl sets FixingCsvUrl field to given value.

### HasFixingCsvUrl

`func (o *AnalysisInsight) HasFixingCsvUrl() bool`

HasFixingCsvUrl returns a boolean if a field has been set.

### GetEngine

`func (o *AnalysisInsight) GetEngine() map[string]interface{}`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *AnalysisInsight) GetEngineOk() (*map[string]interface{}, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *AnalysisInsight) SetEngine(v map[string]interface{})`

SetEngine sets Engine field to given value.


### GetDigest

`func (o *AnalysisInsight) GetDigest() AnalysisInsightDigest`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *AnalysisInsight) GetDigestOk() (*AnalysisInsightDigest, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *AnalysisInsight) SetDigest(v AnalysisInsightDigest)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *AnalysisInsight) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetDigestError

`func (o *AnalysisInsight) GetDigestError() string`

GetDigestError returns the DigestError field if non-nil, zero value otherwise.

### GetDigestErrorOk

`func (o *AnalysisInsight) GetDigestErrorOk() (*string, bool)`

GetDigestErrorOk returns a tuple with the DigestError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigestError

`func (o *AnalysisInsight) SetDigestError(v string)`

SetDigestError sets DigestError field to given value.

### HasDigestError

`func (o *AnalysisInsight) HasDigestError() bool`

HasDigestError returns a boolean if a field has been set.

### GetTopPanel

`func (o *AnalysisInsight) GetTopPanel() map[string]interface{}`

GetTopPanel returns the TopPanel field if non-nil, zero value otherwise.

### GetTopPanelOk

`func (o *AnalysisInsight) GetTopPanelOk() (*map[string]interface{}, bool)`

GetTopPanelOk returns a tuple with the TopPanel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopPanel

`func (o *AnalysisInsight) SetTopPanel(v map[string]interface{})`

SetTopPanel sets TopPanel field to given value.

### HasTopPanel

`func (o *AnalysisInsight) HasTopPanel() bool`

HasTopPanel returns a boolean if a field has been set.

### GetAnalyzeLinkPath

`func (o *AnalysisInsight) GetAnalyzeLinkPath() string`

GetAnalyzeLinkPath returns the AnalyzeLinkPath field if non-nil, zero value otherwise.

### GetAnalyzeLinkPathOk

`func (o *AnalysisInsight) GetAnalyzeLinkPathOk() (*string, bool)`

GetAnalyzeLinkPathOk returns a tuple with the AnalyzeLinkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalyzeLinkPath

`func (o *AnalysisInsight) SetAnalyzeLinkPath(v string)`

SetAnalyzeLinkPath sets AnalyzeLinkPath field to given value.

### HasAnalyzeLinkPath

`func (o *AnalysisInsight) HasAnalyzeLinkPath() bool`

HasAnalyzeLinkPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


