# NumberSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Category** | [**SettingsCategory**](SettingsCategory.md) |  | 
**Type** | **string** |  | 
**Def** | Pointer to **NullableFloat64** |  | [optional] 
**Options** | Pointer to **[]float64** |  | [optional] 
**Step** | Pointer to **float64** |  | [optional] 
**Max** | Pointer to **float64** |  | [optional] 
**Min** | Pointer to **float64** |  | [optional] 
**Required** | Pointer to **bool** |  | [optional] 
**Placeholder** | Pointer to **string** |  | [optional] 
**Suffix** | Pointer to **string** |  | [optional] 
**LimitedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewNumberSchema

`func NewNumberSchema(category SettingsCategory, type_ string, ) *NumberSchema`

NewNumberSchema instantiates a new NumberSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNumberSchemaWithDefaults

`func NewNumberSchemaWithDefaults() *NumberSchema`

NewNumberSchemaWithDefaults instantiates a new NumberSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *NumberSchema) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NumberSchema) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NumberSchema) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NumberSchema) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *NumberSchema) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NumberSchema) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NumberSchema) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NumberSchema) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCategory

`func (o *NumberSchema) GetCategory() SettingsCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NumberSchema) GetCategoryOk() (*SettingsCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NumberSchema) SetCategory(v SettingsCategory)`

SetCategory sets Category field to given value.


### GetType

`func (o *NumberSchema) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NumberSchema) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NumberSchema) SetType(v string)`

SetType sets Type field to given value.


### GetDef

`func (o *NumberSchema) GetDef() float64`

GetDef returns the Def field if non-nil, zero value otherwise.

### GetDefOk

`func (o *NumberSchema) GetDefOk() (*float64, bool)`

GetDefOk returns a tuple with the Def field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDef

`func (o *NumberSchema) SetDef(v float64)`

SetDef sets Def field to given value.

### HasDef

`func (o *NumberSchema) HasDef() bool`

HasDef returns a boolean if a field has been set.

### SetDefNil

`func (o *NumberSchema) SetDefNil(b bool)`

 SetDefNil sets the value for Def to be an explicit nil

### UnsetDef
`func (o *NumberSchema) UnsetDef()`

UnsetDef ensures that no value is present for Def, not even an explicit nil
### GetOptions

`func (o *NumberSchema) GetOptions() []float64`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *NumberSchema) GetOptionsOk() (*[]float64, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *NumberSchema) SetOptions(v []float64)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *NumberSchema) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetStep

`func (o *NumberSchema) GetStep() float64`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *NumberSchema) GetStepOk() (*float64, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *NumberSchema) SetStep(v float64)`

SetStep sets Step field to given value.

### HasStep

`func (o *NumberSchema) HasStep() bool`

HasStep returns a boolean if a field has been set.

### GetMax

`func (o *NumberSchema) GetMax() float64`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *NumberSchema) GetMaxOk() (*float64, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *NumberSchema) SetMax(v float64)`

SetMax sets Max field to given value.

### HasMax

`func (o *NumberSchema) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *NumberSchema) GetMin() float64`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *NumberSchema) GetMinOk() (*float64, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *NumberSchema) SetMin(v float64)`

SetMin sets Min field to given value.

### HasMin

`func (o *NumberSchema) HasMin() bool`

HasMin returns a boolean if a field has been set.

### GetRequired

`func (o *NumberSchema) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *NumberSchema) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *NumberSchema) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *NumberSchema) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetPlaceholder

`func (o *NumberSchema) GetPlaceholder() string`

GetPlaceholder returns the Placeholder field if non-nil, zero value otherwise.

### GetPlaceholderOk

`func (o *NumberSchema) GetPlaceholderOk() (*string, bool)`

GetPlaceholderOk returns a tuple with the Placeholder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaceholder

`func (o *NumberSchema) SetPlaceholder(v string)`

SetPlaceholder sets Placeholder field to given value.

### HasPlaceholder

`func (o *NumberSchema) HasPlaceholder() bool`

HasPlaceholder returns a boolean if a field has been set.

### GetSuffix

`func (o *NumberSchema) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *NumberSchema) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *NumberSchema) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *NumberSchema) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetLimitedBy

`func (o *NumberSchema) GetLimitedBy() string`

GetLimitedBy returns the LimitedBy field if non-nil, zero value otherwise.

### GetLimitedByOk

`func (o *NumberSchema) GetLimitedByOk() (*string, bool)`

GetLimitedByOk returns a tuple with the LimitedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitedBy

`func (o *NumberSchema) SetLimitedBy(v string)`

SetLimitedBy sets LimitedBy field to given value.

### HasLimitedBy

`func (o *NumberSchema) HasLimitedBy() bool`

HasLimitedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


