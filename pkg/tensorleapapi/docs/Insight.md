# Insight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**SessionRunId** | **string** |  | 
**InsightType** | [**InsightType**](InsightType.md) |  | 
**Index** | **float64** |  | 
**Status** | [**InsightStatus**](InsightStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewInsight

`func NewInsight(cid string, sessionRunId string, insightType InsightType, index float64, status InsightStatus, createdAt time.Time, updatedAt time.Time, ) *Insight`

NewInsight instantiates a new Insight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInsightWithDefaults

`func NewInsightWithDefaults() *Insight`

NewInsightWithDefaults instantiates a new Insight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *Insight) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *Insight) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *Insight) SetCid(v string)`

SetCid sets Cid field to given value.


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


### GetInsightType

`func (o *Insight) GetInsightType() InsightType`

GetInsightType returns the InsightType field if non-nil, zero value otherwise.

### GetInsightTypeOk

`func (o *Insight) GetInsightTypeOk() (*InsightType, bool)`

GetInsightTypeOk returns a tuple with the InsightType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInsightType

`func (o *Insight) SetInsightType(v InsightType)`

SetInsightType sets InsightType field to given value.


### GetIndex

`func (o *Insight) GetIndex() float64`

GetIndex returns the Index field if non-nil, zero value otherwise.

### GetIndexOk

`func (o *Insight) GetIndexOk() (*float64, bool)`

GetIndexOk returns a tuple with the Index field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndex

`func (o *Insight) SetIndex(v float64)`

SetIndex sets Index field to given value.


### GetStatus

`func (o *Insight) GetStatus() InsightStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Insight) GetStatusOk() (*InsightStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Insight) SetStatus(v InsightStatus)`

SetStatus sets Status field to given value.


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



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


