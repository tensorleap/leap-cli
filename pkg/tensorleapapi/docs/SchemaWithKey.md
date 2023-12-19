# SchemaWithKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyName** | **string** |  | 
**Schema** | [**SettingSchema**](SettingSchema.md) |  | 

## Methods

### NewSchemaWithKey

`func NewSchemaWithKey(keyName string, schema SettingSchema, ) *SchemaWithKey`

NewSchemaWithKey instantiates a new SchemaWithKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSchemaWithKeyWithDefaults

`func NewSchemaWithKeyWithDefaults() *SchemaWithKey`

NewSchemaWithKeyWithDefaults instantiates a new SchemaWithKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyName

`func (o *SchemaWithKey) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *SchemaWithKey) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *SchemaWithKey) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### GetSchema

`func (o *SchemaWithKey) GetSchema() SettingSchema`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *SchemaWithKey) GetSchemaOk() (*SettingSchema, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *SchemaWithKey) SetSchema(v SettingSchema)`

SetSchema sets Schema field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


