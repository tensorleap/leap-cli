# MutualInformationElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureName** | **string** |  | 
**FeatureValue** | Pointer to **string** |  | [optional] 
**Direction** | [**Direction**](Direction.md) |  | 
**IsCategorical** | **bool** |  | 

## Methods

### NewMutualInformationElement

`func NewMutualInformationElement(featureName string, direction Direction, isCategorical bool, ) *MutualInformationElement`

NewMutualInformationElement instantiates a new MutualInformationElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualInformationElementWithDefaults

`func NewMutualInformationElementWithDefaults() *MutualInformationElement`

NewMutualInformationElementWithDefaults instantiates a new MutualInformationElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureName

`func (o *MutualInformationElement) GetFeatureName() string`

GetFeatureName returns the FeatureName field if non-nil, zero value otherwise.

### GetFeatureNameOk

`func (o *MutualInformationElement) GetFeatureNameOk() (*string, bool)`

GetFeatureNameOk returns a tuple with the FeatureName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureName

`func (o *MutualInformationElement) SetFeatureName(v string)`

SetFeatureName sets FeatureName field to given value.


### GetFeatureValue

`func (o *MutualInformationElement) GetFeatureValue() string`

GetFeatureValue returns the FeatureValue field if non-nil, zero value otherwise.

### GetFeatureValueOk

`func (o *MutualInformationElement) GetFeatureValueOk() (*string, bool)`

GetFeatureValueOk returns a tuple with the FeatureValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureValue

`func (o *MutualInformationElement) SetFeatureValue(v string)`

SetFeatureValue sets FeatureValue field to given value.

### HasFeatureValue

`func (o *MutualInformationElement) HasFeatureValue() bool`

HasFeatureValue returns a boolean if a field has been set.

### GetDirection

`func (o *MutualInformationElement) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *MutualInformationElement) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *MutualInformationElement) SetDirection(v Direction)`

SetDirection sets Direction field to given value.


### GetIsCategorical

`func (o *MutualInformationElement) GetIsCategorical() bool`

GetIsCategorical returns the IsCategorical field if non-nil, zero value otherwise.

### GetIsCategoricalOk

`func (o *MutualInformationElement) GetIsCategoricalOk() (*bool, bool)`

GetIsCategoricalOk returns a tuple with the IsCategorical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCategorical

`func (o *MutualInformationElement) SetIsCategorical(v bool)`

SetIsCategorical sets IsCategorical field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


