# DatasetPreprocess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TrainingLength** | **float64** |  | 
**ValidationLength** | **float64** |  | 
**TestLength** | Pointer to **float64** |  | [optional] 
**UnlabeledLength** | Pointer to **float64** |  | [optional] 

## Methods

### NewDatasetPreprocess

`func NewDatasetPreprocess(trainingLength float64, validationLength float64, ) *DatasetPreprocess`

NewDatasetPreprocess instantiates a new DatasetPreprocess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetPreprocessWithDefaults

`func NewDatasetPreprocessWithDefaults() *DatasetPreprocess`

NewDatasetPreprocessWithDefaults instantiates a new DatasetPreprocess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTrainingLength

`func (o *DatasetPreprocess) GetTrainingLength() float64`

GetTrainingLength returns the TrainingLength field if non-nil, zero value otherwise.

### GetTrainingLengthOk

`func (o *DatasetPreprocess) GetTrainingLengthOk() (*float64, bool)`

GetTrainingLengthOk returns a tuple with the TrainingLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingLength

`func (o *DatasetPreprocess) SetTrainingLength(v float64)`

SetTrainingLength sets TrainingLength field to given value.


### GetValidationLength

`func (o *DatasetPreprocess) GetValidationLength() float64`

GetValidationLength returns the ValidationLength field if non-nil, zero value otherwise.

### GetValidationLengthOk

`func (o *DatasetPreprocess) GetValidationLengthOk() (*float64, bool)`

GetValidationLengthOk returns a tuple with the ValidationLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationLength

`func (o *DatasetPreprocess) SetValidationLength(v float64)`

SetValidationLength sets ValidationLength field to given value.


### GetTestLength

`func (o *DatasetPreprocess) GetTestLength() float64`

GetTestLength returns the TestLength field if non-nil, zero value otherwise.

### GetTestLengthOk

`func (o *DatasetPreprocess) GetTestLengthOk() (*float64, bool)`

GetTestLengthOk returns a tuple with the TestLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestLength

`func (o *DatasetPreprocess) SetTestLength(v float64)`

SetTestLength sets TestLength field to given value.

### HasTestLength

`func (o *DatasetPreprocess) HasTestLength() bool`

HasTestLength returns a boolean if a field has been set.

### GetUnlabeledLength

`func (o *DatasetPreprocess) GetUnlabeledLength() float64`

GetUnlabeledLength returns the UnlabeledLength field if non-nil, zero value otherwise.

### GetUnlabeledLengthOk

`func (o *DatasetPreprocess) GetUnlabeledLengthOk() (*float64, bool)`

GetUnlabeledLengthOk returns a tuple with the UnlabeledLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlabeledLength

`func (o *DatasetPreprocess) SetUnlabeledLength(v float64)`

SetUnlabeledLength sets UnlabeledLength field to given value.

### HasUnlabeledLength

`func (o *DatasetPreprocess) HasUnlabeledLength() bool`

HasUnlabeledLength returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


