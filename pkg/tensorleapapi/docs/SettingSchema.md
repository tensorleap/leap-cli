# SettingSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Def** | Pointer to **NullableFloat64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Options** | Pointer to **[]float64** |  | [optional] 
**MaxLength** | Pointer to **float64** |  | [optional] 
**MinLength** | Pointer to **float64** |  | [optional] 
**Pattern** | Pointer to **string** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Step** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 

## Methods

### NewSettingSchema

`func NewSettingSchema(type_ string, ) *SettingSchema`

NewSettingSchema instantiates a new SettingSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingSchemaWithDefaults

`func NewSettingSchemaWithDefaults() *SettingSchema`

NewSettingSchemaWithDefaults instantiates a new SettingSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SettingSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SettingSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SettingSchema) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *SettingSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SettingSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SettingSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SettingSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDef

`func (o *SettingSchema) GetDef() float64`

GetDef returns the Def field if non-nil, zero value otherwise.

### GetDefOk

`func (o *SettingSchema) GetDefOk() (*float64, bool)`

GetDefOk returns a tuple with the Def field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDef

`func (o *SettingSchema) SetDef(v float64)`

SetDef sets Def field to given value.

### HasDef

`func (o *SettingSchema) HasDef() bool`

HasDef returns a boolean if a field has been set.

### SetDefNil

`func (o *SettingSchema) SetDefNil(b bool)`

 SetDefNil sets the value for Def to be an explicit nil

### UnsetDef
`func (o *SettingSchema) UnsetDef()`

UnsetDef ensures that no value is present for Def, not even an explicit nil
### GetDescription

`func (o *SettingSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SettingSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SettingSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SettingSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOptions

`func (o *SettingSchema) GetOptions() []float64`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *SettingSchema) GetOptionsOk() (*[]float64, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *SettingSchema) SetOptions(v []float64)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *SettingSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetMaxLength

`func (o *SettingSchema) GetMaxLength() float64`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *SettingSchema) GetMaxLengthOk() (*float64, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *SettingSchema) SetMaxLength(v float64)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *SettingSchema) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *SettingSchema) GetMinLength() float64`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *SettingSchema) GetMinLengthOk() (*float64, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *SettingSchema) SetMinLength(v float64)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *SettingSchema) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetPattern

`func (o *SettingSchema) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *SettingSchema) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *SettingSchema) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *SettingSchema) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### GetRequired

`func (o *SettingSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *SettingSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *SettingSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *SettingSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *SettingSchema) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *SettingSchema) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *SettingSchema) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *SettingSchema) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetStep

`func (o *SettingSchema) GetStep() float64`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *SettingSchema) GetStepOk() (*float64, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *SettingSchema) SetStep(v float64)`

SetStep sets Step field to given value.

### HasStep

`func (o *SettingSchema) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetMax

`func (o *SettingSchema) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *SettingSchema) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *SettingSchema) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *SettingSchema) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *SettingSchema) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *SettingSchema) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *SettingSchema) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *SettingSchema) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetSuffix

`func (o *SettingSchema) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *SettingSchema) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *SettingSchema) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *SettingSchema) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


