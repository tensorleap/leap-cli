# LicenseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**LicenseType**](LicenseType.md) |  | 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**MaxUsers** | Pointer to **float64** |  | [optional] 
**MaxConcurrentEvaluateJobsPerUser** | Pointer to **float64** | Maximum number of concurrent Evaluate (or Continue-Evaluate) jobs a single user is allowed to have running at the same time. Undefined or 0 means unlimited. | [optional] 
**MaxProjects** | Pointer to **float64** | Projects the team may own. Undefined or 0 means unlimited. | [optional] 
**Activated** | Pointer to **bool** | False while a colored license waits for its activation key. getEnvironmentInfo always sends it to clients: true for every license not waiting for its key. | [optional] 
**ActivationUrl** | Pointer to **string** | Asked for the key before a blocked action; answers { activated, licenseToken }. | [optional] 
**Rid** | Pointer to **string** | Free Tier registration id; kept after activation to report milestones. Not for clients. | [optional] 

## Methods

### NewLicenseMetadata

`func NewLicenseMetadata(type_ LicenseType, ) *LicenseMetadata`

NewLicenseMetadata instantiates a new LicenseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseMetadataWithDefaults

`func NewLicenseMetadataWithDefaults() *LicenseMetadata`

NewLicenseMetadataWithDefaults instantiates a new LicenseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *LicenseMetadata) GetType() LicenseType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LicenseMetadata) GetTypeOk() (*LicenseType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LicenseMetadata) SetType(v LicenseType)`

SetType sets Type field to given value.


### GetExpirationDate

`func (o *LicenseMetadata) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *LicenseMetadata) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *LicenseMetadata) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *LicenseMetadata) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetMaxUsers

`func (o *LicenseMetadata) GetMaxUsers() float64`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *LicenseMetadata) GetMaxUsersOk() (*float64, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *LicenseMetadata) SetMaxUsers(v float64)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *LicenseMetadata) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMaxConcurrentEvaluateJobsPerUser

`func (o *LicenseMetadata) GetMaxConcurrentEvaluateJobsPerUser() float64`

GetMaxConcurrentEvaluateJobsPerUser returns the MaxConcurrentEvaluateJobsPerUser field if non-nil, zero value otherwise.

### GetMaxConcurrentEvaluateJobsPerUserOk

`func (o *LicenseMetadata) GetMaxConcurrentEvaluateJobsPerUserOk() (*float64, bool)`

GetMaxConcurrentEvaluateJobsPerUserOk returns a tuple with the MaxConcurrentEvaluateJobsPerUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxConcurrentEvaluateJobsPerUser

`func (o *LicenseMetadata) SetMaxConcurrentEvaluateJobsPerUser(v float64)`

SetMaxConcurrentEvaluateJobsPerUser sets MaxConcurrentEvaluateJobsPerUser field to given value.

### HasMaxConcurrentEvaluateJobsPerUser

`func (o *LicenseMetadata) HasMaxConcurrentEvaluateJobsPerUser() bool`

HasMaxConcurrentEvaluateJobsPerUser returns a boolean if a field has been set.

### GetMaxProjects

`func (o *LicenseMetadata) GetMaxProjects() float64`

GetMaxProjects returns the MaxProjects field if non-nil, zero value otherwise.

### GetMaxProjectsOk

`func (o *LicenseMetadata) GetMaxProjectsOk() (*float64, bool)`

GetMaxProjectsOk returns a tuple with the MaxProjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxProjects

`func (o *LicenseMetadata) SetMaxProjects(v float64)`

SetMaxProjects sets MaxProjects field to given value.

### HasMaxProjects

`func (o *LicenseMetadata) HasMaxProjects() bool`

HasMaxProjects returns a boolean if a field has been set.

### GetActivated

`func (o *LicenseMetadata) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *LicenseMetadata) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *LicenseMetadata) SetActivated(v bool)`

SetActivated sets Activated field to given value.

### HasActivated

`func (o *LicenseMetadata) HasActivated() bool`

HasActivated returns a boolean if a field has been set.

### GetActivationUrl

`func (o *LicenseMetadata) GetActivationUrl() string`

GetActivationUrl returns the ActivationUrl field if non-nil, zero value otherwise.

### GetActivationUrlOk

`func (o *LicenseMetadata) GetActivationUrlOk() (*string, bool)`

GetActivationUrlOk returns a tuple with the ActivationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationUrl

`func (o *LicenseMetadata) SetActivationUrl(v string)`

SetActivationUrl sets ActivationUrl field to given value.

### HasActivationUrl

`func (o *LicenseMetadata) HasActivationUrl() bool`

HasActivationUrl returns a boolean if a field has been set.

### GetRid

`func (o *LicenseMetadata) GetRid() string`

GetRid returns the Rid field if non-nil, zero value otherwise.

### GetRidOk

`func (o *LicenseMetadata) GetRidOk() (*string, bool)`

GetRidOk returns a tuple with the Rid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRid

`func (o *LicenseMetadata) SetRid(v string)`

SetRid sets Rid field to given value.

### HasRid

`func (o *LicenseMetadata) HasRid() bool`

HasRid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


