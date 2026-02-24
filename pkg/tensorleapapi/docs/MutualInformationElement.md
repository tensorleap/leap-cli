# MutualInformationElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Features** | [**[]MutualInformationFeature**](MutualInformationFeature.md) |  | 
**ValueInCluster** | Pointer to **float64** |  | [optional] 
**ValueOutsideCluster** | Pointer to **float64** |  | [optional] 
**Score** | Pointer to **float64** |  | [optional] 
**IsUnique** | **bool** |  | 

## Methods

### NewMutualInformationElement

`func NewMutualInformationElement(features []MutualInformationFeature, isUnique bool, ) *MutualInformationElement`

NewMutualInformationElement instantiates a new MutualInformationElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMutualInformationElementWithDefaults

`func NewMutualInformationElementWithDefaults() *MutualInformationElement`

NewMutualInformationElementWithDefaults instantiates a new MutualInformationElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatures

`func (o *MutualInformationElement) GetFeatures() []MutualInformationFeature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *MutualInformationElement) GetFeaturesOk() (*[]MutualInformationFeature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *MutualInformationElement) SetFeatures(v []MutualInformationFeature)`

SetFeatures sets Features field to given value.


### GetValueInCluster

`func (o *MutualInformationElement) GetValueInCluster() float64`

GetValueInCluster returns the ValueInCluster field if non-nil, zero value otherwise.

### GetValueInClusterOk

`func (o *MutualInformationElement) GetValueInClusterOk() (*float64, bool)`

GetValueInClusterOk returns a tuple with the ValueInCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueInCluster

`func (o *MutualInformationElement) SetValueInCluster(v float64)`

SetValueInCluster sets ValueInCluster field to given value.

### HasValueInCluster

`func (o *MutualInformationElement) HasValueInCluster() bool`

HasValueInCluster returns a boolean if a field has been set.

### GetValueOutsideCluster

`func (o *MutualInformationElement) GetValueOutsideCluster() float64`

GetValueOutsideCluster returns the ValueOutsideCluster field if non-nil, zero value otherwise.

### GetValueOutsideClusterOk

`func (o *MutualInformationElement) GetValueOutsideClusterOk() (*float64, bool)`

GetValueOutsideClusterOk returns a tuple with the ValueOutsideCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueOutsideCluster

`func (o *MutualInformationElement) SetValueOutsideCluster(v float64)`

SetValueOutsideCluster sets ValueOutsideCluster field to given value.

### HasValueOutsideCluster

`func (o *MutualInformationElement) HasValueOutsideCluster() bool`

HasValueOutsideCluster returns a boolean if a field has been set.

### GetScore

`func (o *MutualInformationElement) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *MutualInformationElement) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *MutualInformationElement) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *MutualInformationElement) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetIsUnique

`func (o *MutualInformationElement) GetIsUnique() bool`

GetIsUnique returns the IsUnique field if non-nil, zero value otherwise.

### GetIsUniqueOk

`func (o *MutualInformationElement) GetIsUniqueOk() (*bool, bool)`

GetIsUniqueOk returns a tuple with the IsUnique field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnique

`func (o *MutualInformationElement) SetIsUnique(v bool)`

SetIsUnique sets IsUnique field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


