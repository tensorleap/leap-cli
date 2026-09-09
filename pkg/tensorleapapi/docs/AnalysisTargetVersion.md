# AnalysisTargetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**Name** | **string** |  | 
**SerialNumber** | Pointer to **float64** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Evaluated** | **bool** |  | 
**HasInsights** | **bool** |  | 

## Methods

### NewAnalysisTargetVersion

`func NewAnalysisTargetVersion(cid string, name string, createdAt time.Time, evaluated bool, hasInsights bool, ) *AnalysisTargetVersion`

NewAnalysisTargetVersion instantiates a new AnalysisTargetVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnalysisTargetVersionWithDefaults

`func NewAnalysisTargetVersionWithDefaults() *AnalysisTargetVersion`

NewAnalysisTargetVersionWithDefaults instantiates a new AnalysisTargetVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *AnalysisTargetVersion) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *AnalysisTargetVersion) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *AnalysisTargetVersion) SetCid(v string)`

SetCid sets Cid field to given value.


### GetName

`func (o *AnalysisTargetVersion) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AnalysisTargetVersion) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AnalysisTargetVersion) SetName(v string)`

SetName sets Name field to given value.


### GetSerialNumber

`func (o *AnalysisTargetVersion) GetSerialNumber() float64`

GetSerialNumber returns the SerialNumber field if non-nil, zero value otherwise.

### GetSerialNumberOk

`func (o *AnalysisTargetVersion) GetSerialNumberOk() (*float64, bool)`

GetSerialNumberOk returns a tuple with the SerialNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSerialNumber

`func (o *AnalysisTargetVersion) SetSerialNumber(v float64)`

SetSerialNumber sets SerialNumber field to given value.

### HasSerialNumber

`func (o *AnalysisTargetVersion) HasSerialNumber() bool`

HasSerialNumber returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AnalysisTargetVersion) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AnalysisTargetVersion) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AnalysisTargetVersion) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEvaluated

`func (o *AnalysisTargetVersion) GetEvaluated() bool`

GetEvaluated returns the Evaluated field if non-nil, zero value otherwise.

### GetEvaluatedOk

`func (o *AnalysisTargetVersion) GetEvaluatedOk() (*bool, bool)`

GetEvaluatedOk returns a tuple with the Evaluated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvaluated

`func (o *AnalysisTargetVersion) SetEvaluated(v bool)`

SetEvaluated sets Evaluated field to given value.


### GetHasInsights

`func (o *AnalysisTargetVersion) GetHasInsights() bool`

GetHasInsights returns the HasInsights field if non-nil, zero value otherwise.

### GetHasInsightsOk

`func (o *AnalysisTargetVersion) GetHasInsightsOk() (*bool, bool)`

GetHasInsightsOk returns a tuple with the HasInsights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasInsights

`func (o *AnalysisTargetVersion) SetHasInsights(v bool)`

SetHasInsights sets HasInsights field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


