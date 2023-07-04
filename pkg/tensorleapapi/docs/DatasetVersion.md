# DatasetVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cid** | **string** |  | 
**TeamId** | **string** |  | 
**DatasetId** | **string** |  | 
**Note** | **string** |  | 
**CreatedAt** | **string** |  | 
**CreatedBy** | **string** |  | 
**TestStatus** | [**TestStatus**](TestStatus.md) |  | 
**Metadata** | [**DatasetMetadata**](DatasetMetadata.md) |  | 
**BlobPath** | **string** |  | 
**CodeEntryFile** | **string** |  | 

## Methods

### NewDatasetVersion

`func NewDatasetVersion(cid string, teamId string, datasetId string, note string, createdAt string, createdBy string, testStatus TestStatus, metadata DatasetMetadata, blobPath string, codeEntryFile string, ) *DatasetVersion`

NewDatasetVersion instantiates a new DatasetVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDatasetVersionWithDefaults

`func NewDatasetVersionWithDefaults() *DatasetVersion`

NewDatasetVersionWithDefaults instantiates a new DatasetVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCid

`func (o *DatasetVersion) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *DatasetVersion) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *DatasetVersion) SetCid(v string)`

SetCid sets Cid field to given value.


### GetTeamId

`func (o *DatasetVersion) GetTeamId() string`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *DatasetVersion) GetTeamIdOk() (*string, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *DatasetVersion) SetTeamId(v string)`

SetTeamId sets TeamId field to given value.


### GetDatasetId

`func (o *DatasetVersion) GetDatasetId() string`

GetDatasetId returns the DatasetId field if non-nil, zero value otherwise.

### GetDatasetIdOk

`func (o *DatasetVersion) GetDatasetIdOk() (*string, bool)`

GetDatasetIdOk returns a tuple with the DatasetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatasetId

`func (o *DatasetVersion) SetDatasetId(v string)`

SetDatasetId sets DatasetId field to given value.


### GetNote

`func (o *DatasetVersion) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *DatasetVersion) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *DatasetVersion) SetNote(v string)`

SetNote sets Note field to given value.


### GetCreatedAt

`func (o *DatasetVersion) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DatasetVersion) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DatasetVersion) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedBy

`func (o *DatasetVersion) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *DatasetVersion) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *DatasetVersion) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetTestStatus

`func (o *DatasetVersion) GetTestStatus() TestStatus`

GetTestStatus returns the TestStatus field if non-nil, zero value otherwise.

### GetTestStatusOk

`func (o *DatasetVersion) GetTestStatusOk() (*TestStatus, bool)`

GetTestStatusOk returns a tuple with the TestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestStatus

`func (o *DatasetVersion) SetTestStatus(v TestStatus)`

SetTestStatus sets TestStatus field to given value.


### GetMetadata

`func (o *DatasetVersion) GetMetadata() DatasetMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DatasetVersion) GetMetadataOk() (*DatasetMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DatasetVersion) SetMetadata(v DatasetMetadata)`

SetMetadata sets Metadata field to given value.


### GetBlobPath

`func (o *DatasetVersion) GetBlobPath() string`

GetBlobPath returns the BlobPath field if non-nil, zero value otherwise.

### GetBlobPathOk

`func (o *DatasetVersion) GetBlobPathOk() (*string, bool)`

GetBlobPathOk returns a tuple with the BlobPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobPath

`func (o *DatasetVersion) SetBlobPath(v string)`

SetBlobPath sets BlobPath field to given value.


### GetCodeEntryFile

`func (o *DatasetVersion) GetCodeEntryFile() string`

GetCodeEntryFile returns the CodeEntryFile field if non-nil, zero value otherwise.

### GetCodeEntryFileOk

`func (o *DatasetVersion) GetCodeEntryFileOk() (*string, bool)`

GetCodeEntryFileOk returns a tuple with the CodeEntryFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCodeEntryFile

`func (o *DatasetVersion) SetCodeEntryFile(v string)`

SetCodeEntryFile sets CodeEntryFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


