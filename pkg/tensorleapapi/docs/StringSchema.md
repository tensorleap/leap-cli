# StringSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | [**SettingsCategory**](SettingsCategory.md) |  | 
**Type** | **string** |  | 
**Def** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **[]string** |  | [optional] 
**MaxLength** | Pointer to **float64** |  | [optional] 
**MinLength** | Pointer to **float64** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 

## Methods

### NewStringSchema

`func NewStringSchema(category SettingsCategory, type_ string, ) *StringSchema`

NewStringSchema instantiates a new StringSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStringSchemaWithDefaults

`func NewStringSchemaWithDefaults() *StringSchema`

NewStringSchemaWithDefaults instantiates a new StringSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *StringSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StringSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StringSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StringSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *StringSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StringSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StringSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StringSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *StringSchema) GetCategory() SettingsCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StringSchema) GetCategoryOk() (*SettingsCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StringSchema) SetCategory(v SettingsCategory)`

SetCategory sets Category field to given value.


### GetType

`func (o *StringSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *StringSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *StringSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDef

`func (o *StringSchema) GetDef() string`

GetDef returns the Def field if non-nil, zero value otherwise.

### GetDefOk

`func (o *StringSchema) GetDefOk() (*string, bool)`

GetDefOk returns a tuple with the Def field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDef

`func (o *StringSchema) SetDef(v string)`

SetDef sets Def field to given value.

### HasDef

`func (o *StringSchema) HasDef() bool`

HasDef returns a boolean if a field has been set.

### SetDefNil

`func (o *StringSchema) SetDefNil(b bool)`

 SetDefNil sets the value for Def to be an explicit nil

### UnsetDef
`func (o *StringSchema) UnsetDef()`

UnsetDef ensures that no value is present for Def, not even an explicit nil
### GetOptions

`func (o *StringSchema) GetOptions() []string`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *StringSchema) GetOptionsOk() (*[]string, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *StringSchema) SetOptions(v []string)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *StringSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMaxLength

`func (o *StringSchema) GetMaxLength() float64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *StringSchema) GetMaxLengthOk() (*float64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *StringSchema) SetMaxLength(v float64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *StringSchema) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *StringSchema) GetMinLength() float64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *StringSchema) GetMinLengthOk() (*float64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *StringSchema) SetMinLength(v float64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *StringSchema) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetPattern

`func (o *StringSchema) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *StringSchema) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *StringSchema) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *StringSchema) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRequired

`func (o *StringSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *StringSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *StringSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *StringSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *StringSchema) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *StringSchema) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *StringSchema) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *StringSchema) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


