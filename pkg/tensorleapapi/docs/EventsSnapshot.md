# EventsSnapshot

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Events** | [**[]JobEvent**](JobEvent.md) |  | 
**SnapshotIndex** | **float64** |  | 

## Methods

### NewEventsSnapshot

`func NewEventsSnapshot(events []JobEvent, snapshotIndex float64, ) *EventsSnapshot`

NewEventsSnapshot instantiates a new EventsSnapshot object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsSnapshotWithDefaults

`func NewEventsSnapshotWithDefaults() *EventsSnapshot`

NewEventsSnapshotWithDefaults instantiates a new EventsSnapshot object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEvents

`func (o *EventsSnapshot) GetEvents() []JobEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *EventsSnapshot) GetEventsOk() (*[]JobEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *EventsSnapshot) SetEvents(v []JobEvent)`

SetEvents sets Events field to given value.


### GetSnapshotIndex

`func (o *EventsSnapshot) GetSnapshotIndex() float64`

GetSnapshotIndex returns the SnapshotIndex field if non-nil, zero value otherwise.

### GetSnapshotIndexOk

`func (o *EventsSnapshot) GetSnapshotIndexOk() (*float64, bool)`

GetSnapshotIndexOk returns a tuple with the SnapshotIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnapshotIndex

`func (o *EventsSnapshot) SetSnapshotIndex(v float64)`

SetSnapshotIndex sets SnapshotIndex field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


