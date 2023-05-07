# FeatureImportance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GradCam** | [**[]FeatureImportanceItem**](FeatureImportanceItem.md) |  | 
**GradByLoss** | [**[]FeatureImportanceItem**](FeatureImportanceItem.md) |  | 

## Methods

### NewFeatureImportance

`func NewFeatureImportance(gradCam []FeatureImportanceItem, gradByLoss []FeatureImportanceItem, ) *FeatureImportance`

NewFeatureImportance instantiates a new FeatureImportance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureImportanceWithDefaults

`func NewFeatureImportanceWithDefaults() *FeatureImportance`

NewFeatureImportanceWithDefaults instantiates a new FeatureImportance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGradCam

`func (o *FeatureImportance) GetGradCam() []FeatureImportanceItem`

GetGradCam returns the GradCam field if non-nil, zero value otherwise.

### GetGradCamOk

`func (o *FeatureImportance) GetGradCamOk() (*[]FeatureImportanceItem, bool)`

GetGradCamOk returns a tuple with the GradCam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradCam

`func (o *FeatureImportance) SetGradCam(v []FeatureImportanceItem)`

SetGradCam sets GradCam field to given value.


### GetGradByLoss

`func (o *FeatureImportance) GetGradByLoss() []FeatureImportanceItem`

GetGradByLoss returns the GradByLoss field if non-nil, zero value otherwise.

### GetGradByLossOk

`func (o *FeatureImportance) GetGradByLossOk() (*[]FeatureImportanceItem, bool)`

GetGradByLossOk returns a tuple with the GradByLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradByLoss

`func (o *FeatureImportance) SetGradByLoss(v []FeatureImportanceItem)`

SetGradByLoss sets GradByLoss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


