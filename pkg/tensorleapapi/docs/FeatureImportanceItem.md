# FeatureImportanceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Fi** | [**VisData**](VisData.md) |  | 
**GroundTruthName** | Pointer to **string** |  | [optional] 

## Methods

### NewFeatureImportanceItem

`func NewFeatureImportanceItem(label string, fi VisData, ) *FeatureImportanceItem`

NewFeatureImportanceItem instantiates a new FeatureImportanceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatureImportanceItemWithDefaults

`func NewFeatureImportanceItemWithDefaults() *FeatureImportanceItem`

NewFeatureImportanceItemWithDefaults instantiates a new FeatureImportanceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *FeatureImportanceItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *FeatureImportanceItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *FeatureImportanceItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetFi

`func (o *FeatureImportanceItem) GetFi() VisData`

GetFi returns the Fi field if non-nil, zero value otherwise.

### GetFiOk

`func (o *FeatureImportanceItem) GetFiOk() (*VisData, bool)`

GetFiOk returns a tuple with the Fi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFi

`func (o *FeatureImportanceItem) SetFi(v VisData)`

SetFi sets Fi field to given value.


### GetGroundTruthName

`func (o *FeatureImportanceItem) GetGroundTruthName() string`

GetGroundTruthName returns the GroundTruthName field if non-nil, zero value otherwise.

### GetGroundTruthNameOk

`func (o *FeatureImportanceItem) GetGroundTruthNameOk() (*string, bool)`

GetGroundTruthNameOk returns a tuple with the GroundTruthName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroundTruthName

`func (o *FeatureImportanceItem) SetGroundTruthName(v string)`

SetGroundTruthName sets GroundTruthName field to given value.

### HasGroundTruthName

`func (o *FeatureImportanceItem) HasGroundTruthName() bool`

HasGroundTruthName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


