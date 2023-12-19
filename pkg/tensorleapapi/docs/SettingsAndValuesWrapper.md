# SettingsAndValuesWrapper

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | [**[]SchemaWithKey**](SchemaWithKey.md) |  | 
**Values** | [**[]ValueWithKey**](ValueWithKey.md) |  | 

## Methods

### NewSettingsAndValuesWrapper

`func NewSettingsAndValuesWrapper(schema []SchemaWithKey, values []ValueWithKey, ) *SettingsAndValuesWrapper`

NewSettingsAndValuesWrapper instantiates a new SettingsAndValuesWrapper object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsAndValuesWrapperWithDefaults

`func NewSettingsAndValuesWrapperWithDefaults() *SettingsAndValuesWrapper`

NewSettingsAndValuesWrapperWithDefaults instantiates a new SettingsAndValuesWrapper object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *SettingsAndValuesWrapper) GetSchema() []SchemaWithKey`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SettingsAndValuesWrapper) GetSchemaOk() (*[]SchemaWithKey, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SettingsAndValuesWrapper) SetSchema(v []SchemaWithKey)`

SetSchema sets Schema field to given value.


### GetValues

`func (o *SettingsAndValuesWrapper) GetValues() []ValueWithKey`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *SettingsAndValuesWrapper) GetValuesOk() (*[]ValueWithKey, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *SettingsAndValuesWrapper) SetValues(v []ValueWithKey)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


