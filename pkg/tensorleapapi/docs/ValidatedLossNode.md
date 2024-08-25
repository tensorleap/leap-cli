# ValidatedLossNode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**ConnectionName** | Pointer to **string** |  | [optional] 
**NodeId** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 
**IsCustomLoss** | **bool** |  | 

## Methods

### NewValidatedLossNode

`func NewValidatedLossNode(name string, isCustomLoss bool, ) *ValidatedLossNode`

NewValidatedLossNode instantiates a new ValidatedLossNode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatedLossNodeWithDefaults

`func NewValidatedLossNodeWithDefaults() *ValidatedLossNode`

NewValidatedLossNodeWithDefaults instantiates a new ValidatedLossNode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ValidatedLossNode) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ValidatedLossNode) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ValidatedLossNode) SetName(v string)`

SetName sets Name field to given value.


### GetConnectionName

`func (o *ValidatedLossNode) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *ValidatedLossNode) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *ValidatedLossNode) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *ValidatedLossNode) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetNodeId

`func (o *ValidatedLossNode) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *ValidatedLossNode) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *ValidatedLossNode) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *ValidatedLossNode) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetError

`func (o *ValidatedLossNode) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidatedLossNode) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidatedLossNode) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ValidatedLossNode) HasError() bool`

HasError returns a boolean if a field has been set.

### GetIsCustomLoss

`func (o *ValidatedLossNode) GetIsCustomLoss() bool`

GetIsCustomLoss returns the IsCustomLoss field if non-nil, zero value otherwise.

### GetIsCustomLossOk

`func (o *ValidatedLossNode) GetIsCustomLossOk() (*bool, bool)`

GetIsCustomLossOk returns a tuple with the IsCustomLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCustomLoss

`func (o *ValidatedLossNode) SetIsCustomLoss(v bool)`

SetIsCustomLoss sets IsCustomLoss field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


