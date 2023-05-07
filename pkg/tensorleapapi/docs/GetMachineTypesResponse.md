# GetMachineTypesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | [**[]MachineTypeOption**](MachineTypeOption.md) |  | 
**DefaultCpuType** | **string** |  | 
**DefaultGpuType** | **string** |  | 

## Methods

### NewGetMachineTypesResponse

`func NewGetMachineTypesResponse(options []MachineTypeOption, defaultCpuType string, defaultGpuType string, ) *GetMachineTypesResponse`

NewGetMachineTypesResponse instantiates a new GetMachineTypesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMachineTypesResponseWithDefaults

`func NewGetMachineTypesResponseWithDefaults() *GetMachineTypesResponse`

NewGetMachineTypesResponseWithDefaults instantiates a new GetMachineTypesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *GetMachineTypesResponse) GetOptions() []MachineTypeOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetMachineTypesResponse) GetOptionsOk() (*[]MachineTypeOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetMachineTypesResponse) SetOptions(v []MachineTypeOption)`

SetOptions sets Options field to given value.


### GetDefaultCpuType

`func (o *GetMachineTypesResponse) GetDefaultCpuType() string`

GetDefaultCpuType returns the DefaultCpuType field if non-nil, zero value otherwise.

### GetDefaultCpuTypeOk

`func (o *GetMachineTypesResponse) GetDefaultCpuTypeOk() (*string, bool)`

GetDefaultCpuTypeOk returns a tuple with the DefaultCpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCpuType

`func (o *GetMachineTypesResponse) SetDefaultCpuType(v string)`

SetDefaultCpuType sets DefaultCpuType field to given value.


### GetDefaultGpuType

`func (o *GetMachineTypesResponse) GetDefaultGpuType() string`

GetDefaultGpuType returns the DefaultGpuType field if non-nil, zero value otherwise.

### GetDefaultGpuTypeOk

`func (o *GetMachineTypesResponse) GetDefaultGpuTypeOk() (*string, bool)`

GetDefaultGpuTypeOk returns a tuple with the DefaultGpuType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGpuType

`func (o *GetMachineTypesResponse) SetDefaultGpuType(v string)`

SetDefaultGpuType sets DefaultGpuType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


