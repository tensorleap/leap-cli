# SlimOrganization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**PublicName** | **string** |  | 
**Type** | **string** |  | 
**IsDefault** | **bool** |  | 
**MachineTypeId** | Pointer to **string** |  | [optional] 

## Methods

### NewSlimOrganization

`func NewSlimOrganization(id string, name string, publicName string, type_ string, isDefault bool, ) *SlimOrganization`

NewSlimOrganization instantiates a new SlimOrganization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimOrganizationWithDefaults

`func NewSlimOrganizationWithDefaults() *SlimOrganization`

NewSlimOrganizationWithDefaults instantiates a new SlimOrganization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimOrganization) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimOrganization) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimOrganization) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SlimOrganization) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SlimOrganization) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SlimOrganization) SetName(v string)`

SetName sets Name field to given value.


### GetPublicName

`func (o *SlimOrganization) GetPublicName() string`

GetPublicName returns the PublicName field if non-nil, zero value otherwise.

### GetPublicNameOk

`func (o *SlimOrganization) GetPublicNameOk() (*string, bool)`

GetPublicNameOk returns a tuple with the PublicName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicName

`func (o *SlimOrganization) SetPublicName(v string)`

SetPublicName sets PublicName field to given value.


### GetType

`func (o *SlimOrganization) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SlimOrganization) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SlimOrganization) SetType(v string)`

SetType sets Type field to given value.


### GetIsDefault

`func (o *SlimOrganization) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SlimOrganization) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SlimOrganization) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetMachineTypeId

`func (o *SlimOrganization) GetMachineTypeId() string`

GetMachineTypeId returns the MachineTypeId field if non-nil, zero value otherwise.

### GetMachineTypeIdOk

`func (o *SlimOrganization) GetMachineTypeIdOk() (*string, bool)`

GetMachineTypeIdOk returns a tuple with the MachineTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMachineTypeId

`func (o *SlimOrganization) SetMachineTypeId(v string)`

SetMachineTypeId sets MachineTypeId field to given value.

### HasMachineTypeId

`func (o *SlimOrganization) HasMachineTypeId() bool`

HasMachineTypeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


