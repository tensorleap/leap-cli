# BoolSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Def** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewBoolSchema

`func NewBoolSchema(type_ string, ) *BoolSchema`

NewBoolSchema instantiates a new BoolSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoolSchemaWithDefaults

`func NewBoolSchemaWithDefaults() *BoolSchema`

NewBoolSchemaWithDefaults instantiates a new BoolSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BoolSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoolSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoolSchema) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *BoolSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BoolSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BoolSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BoolSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDef

`func (o *BoolSchema) GetDef() bool`

GetDef returns the Def field if non-nil, zero value otherwise.

### GetDefOk

`func (o *BoolSchema) GetDefOk() (*bool, bool)`

GetDefOk returns a tuple with the Def field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDef

`func (o *BoolSchema) SetDef(v bool)`

SetDef sets Def field to given value.

### HasDef

`func (o *BoolSchema) HasDef() bool`

HasDef returns a boolean if a field has been set.

### GetDescription

`func (o *BoolSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BoolSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BoolSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BoolSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


