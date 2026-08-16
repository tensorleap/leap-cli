# AudioData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioBlob** | **string** |  | 
**Visual** | [**AudioDataVisual**](AudioDataVisual.md) |  | 
**SampleRate** | **float64** |  | 
**XRange** | Pointer to **[]float64** |  | [optional] 
**Type** | [**DataTypeEnum**](DataTypeEnum.md) |  | 

## Methods

### NewAudioData

`func NewAudioData(audioBlob string, visual AudioDataVisual, sampleRate float64, type_ DataTypeEnum, ) *AudioData`

NewAudioData instantiates a new AudioData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAudioDataWithDefaults

`func NewAudioDataWithDefaults() *AudioData`

NewAudioDataWithDefaults instantiates a new AudioData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioBlob

`func (o *AudioData) GetAudioBlob() string`

GetAudioBlob returns the AudioBlob field if non-nil, zero value otherwise.

### GetAudioBlobOk

`func (o *AudioData) GetAudioBlobOk() (*string, bool)`

GetAudioBlobOk returns a tuple with the AudioBlob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioBlob

`func (o *AudioData) SetAudioBlob(v string)`

SetAudioBlob sets AudioBlob field to given value.


### GetVisual

`func (o *AudioData) GetVisual() AudioDataVisual`

GetVisual returns the Visual field if non-nil, zero value otherwise.

### GetVisualOk

`func (o *AudioData) GetVisualOk() (*AudioDataVisual, bool)`

GetVisualOk returns a tuple with the Visual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisual

`func (o *AudioData) SetVisual(v AudioDataVisual)`

SetVisual sets Visual field to given value.


### GetSampleRate

`func (o *AudioData) GetSampleRate() float64`

GetSampleRate returns the SampleRate field if non-nil, zero value otherwise.

### GetSampleRateOk

`func (o *AudioData) GetSampleRateOk() (*float64, bool)`

GetSampleRateOk returns a tuple with the SampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSampleRate

`func (o *AudioData) SetSampleRate(v float64)`

SetSampleRate sets SampleRate field to given value.


### GetXRange

`func (o *AudioData) GetXRange() []float64`

GetXRange returns the XRange field if non-nil, zero value otherwise.

### GetXRangeOk

`func (o *AudioData) GetXRangeOk() (*[]float64, bool)`

GetXRangeOk returns a tuple with the XRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXRange

`func (o *AudioData) SetXRange(v []float64)`

SetXRange sets XRange field to given value.

### HasXRange

`func (o *AudioData) HasXRange() bool`

HasXRange returns a boolean if a field has been set.

### GetType

`func (o *AudioData) GetType() DataTypeEnum`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AudioData) GetTypeOk() (*DataTypeEnum, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AudioData) SetType(v DataTypeEnum)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


