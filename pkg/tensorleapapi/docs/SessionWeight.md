# SessionWeight

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** |  | 
**BlobName** | Pointer to **string** |  | [optional] 
**Epoch** | **float64** |  | 
**Status** | [**StatusEnum**](StatusEnum.md) |  | 
**GlobalStep** | **float64** |  | 

## Methods

### NewSessionWeight

`func NewSessionWeight(identifier string, epoch float64, status StatusEnum, globalStep float64, ) *SessionWeight`

NewSessionWeight instantiates a new SessionWeight object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionWeightWithDefaults

`func NewSessionWeightWithDefaults() *SessionWeight`

NewSessionWeightWithDefaults instantiates a new SessionWeight object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *SessionWeight) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *SessionWeight) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *SessionWeight) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetBlobName

`func (o *SessionWeight) GetBlobName() string`

GetBlobName returns the BlobName field if non-nil, zero value otherwise.

### GetBlobNameOk

`func (o *SessionWeight) GetBlobNameOk() (*string, bool)`

GetBlobNameOk returns a tuple with the BlobName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobName

`func (o *SessionWeight) SetBlobName(v string)`

SetBlobName sets BlobName field to given value.

### HasBlobName

`func (o *SessionWeight) HasBlobName() bool`

HasBlobName returns a boolean if a field has been set.

### GetEpoch

`func (o *SessionWeight) GetEpoch() float64`

GetEpoch returns the Epoch field if non-nil, zero value otherwise.

### GetEpochOk

`func (o *SessionWeight) GetEpochOk() (*float64, bool)`

GetEpochOk returns a tuple with the Epoch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEpoch

`func (o *SessionWeight) SetEpoch(v float64)`

SetEpoch sets Epoch field to given value.


### GetStatus

`func (o *SessionWeight) GetStatus() StatusEnum`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SessionWeight) GetStatusOk() (*StatusEnum, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SessionWeight) SetStatus(v StatusEnum)`

SetStatus sets Status field to given value.


### GetGlobalStep

`func (o *SessionWeight) GetGlobalStep() float64`

GetGlobalStep returns the GlobalStep field if non-nil, zero value otherwise.

### GetGlobalStepOk

`func (o *SessionWeight) GetGlobalStepOk() (*float64, bool)`

GetGlobalStepOk returns a tuple with the GlobalStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalStep

`func (o *SessionWeight) SetGlobalStep(v float64)`

SetGlobalStep sets GlobalStep field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


