# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**SessionRunId** | **string** |  | 
**ModelEpoch** | **float64** |  | 
**ModelExtId** | **string** |  | 
**OrganizationId** | **string** |  | 
**CreatedBy** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**VisualizationUuid** | **string** |  | 
**Data** | [**InsightType**](InsightType.md) |  | 

## Methods

### NewInsight

`func NewInsight(id string, sessionRunId string, modelEpoch float64, modelExtId string, organizationId string, createdBy string, createdAt time.Time, updatedAt time.Time, visualizationUuid string, data InsightType, ) *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Insight) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Insight) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Insight) SetId(v string)`

SetId sets Id field to given value.


### GetSessionRunId

`func (o *Insight) GetSessionRunId() string`

GetSessionRunId returns the SessionRunId field if non-nil, zero value otherwise.

### GetSessionRunIdOk

`func (o *Insight) GetSessionRunIdOk() (*string, bool)`

GetSessionRunIdOk returns a tuple with the SessionRunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionRunId

`func (o *Insight) SetSessionRunId(v string)`

SetSessionRunId sets SessionRunId field to given value.


### GetModelEpoch

`func (o *Insight) GetModelEpoch() float64`

GetModelEpoch returns the ModelEpoch field if non-nil, zero value otherwise.

### GetModelEpochOk

`func (o *Insight) GetModelEpochOk() (*float64, bool)`

GetModelEpochOk returns a tuple with the ModelEpoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelEpoch

`func (o *Insight) SetModelEpoch(v float64)`

SetModelEpoch sets ModelEpoch field to given value.


### GetModelExtId

`func (o *Insight) GetModelExtId() string`

GetModelExtId returns the ModelExtId field if non-nil, zero value otherwise.

### GetModelExtIdOk

`func (o *Insight) GetModelExtIdOk() (*string, bool)`

GetModelExtIdOk returns a tuple with the ModelExtId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelExtId

`func (o *Insight) SetModelExtId(v string)`

SetModelExtId sets ModelExtId field to given value.


### GetOrganizationId

`func (o *Insight) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *Insight) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *Insight) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetCreatedBy

`func (o *Insight) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *Insight) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *Insight) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreatedAt

`func (o *Insight) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Insight) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Insight) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Insight) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Insight) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Insight) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetVisualizationUuid

`func (o *Insight) GetVisualizationUuid() string`

GetVisualizationUuid returns the VisualizationUuid field if non-nil, zero value otherwise.

### GetVisualizationUuidOk

`func (o *Insight) GetVisualizationUuidOk() (*string, bool)`

GetVisualizationUuidOk returns a tuple with the VisualizationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualizationUuid

`func (o *Insight) SetVisualizationUuid(v string)`

SetVisualizationUuid sets VisualizationUuid field to given value.


### GetData

`func (o *Insight) GetData() InsightType`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Insight) GetDataOk() (*InsightType, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Insight) SetData(v InsightType)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


