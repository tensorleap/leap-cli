# ConfusionMatrixLabelsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LabelsByPrediction** | [**[]PredictionLabels**](PredictionLabels.md) |  | 

## Methods

### NewConfusionMatrixLabelsResponse

`func NewConfusionMatrixLabelsResponse(labelsByPrediction []PredictionLabels, ) *ConfusionMatrixLabelsResponse`

NewConfusionMatrixLabelsResponse instantiates a new ConfusionMatrixLabelsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfusionMatrixLabelsResponseWithDefaults

`func NewConfusionMatrixLabelsResponseWithDefaults() *ConfusionMatrixLabelsResponse`

NewConfusionMatrixLabelsResponseWithDefaults instantiates a new ConfusionMatrixLabelsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabelsByPrediction

`func (o *ConfusionMatrixLabelsResponse) GetLabelsByPrediction() []PredictionLabels`

GetLabelsByPrediction returns the LabelsByPrediction field if non-nil, zero value otherwise.

### GetLabelsByPredictionOk

`func (o *ConfusionMatrixLabelsResponse) GetLabelsByPredictionOk() (*[]PredictionLabels, bool)`

GetLabelsByPredictionOk returns a tuple with the LabelsByPrediction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelsByPrediction

`func (o *ConfusionMatrixLabelsResponse) SetLabelsByPrediction(v []PredictionLabels)`

SetLabelsByPrediction sets LabelsByPrediction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


