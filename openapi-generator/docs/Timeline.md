# Timeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | [**[]TimelineEntriesInner**](TimelineEntriesInner.md) |  | 
**Users** | [**map[string]TimelineUsersValue**](TimelineUsersValue.md) |  | 

## Methods

### NewTimeline

`func NewTimeline(entries []TimelineEntriesInner, users map[string]TimelineUsersValue, ) *Timeline`

NewTimeline instantiates a new Timeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineWithDefaults

`func NewTimelineWithDefaults() *Timeline`

NewTimelineWithDefaults instantiates a new Timeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *Timeline) GetEntries() []TimelineEntriesInner`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *Timeline) GetEntriesOk() (*[]TimelineEntriesInner, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *Timeline) SetEntries(v []TimelineEntriesInner)`

SetEntries sets Entries field to given value.


### GetUsers

`func (o *Timeline) GetUsers() map[string]TimelineUsersValue`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Timeline) GetUsersOk() (*map[string]TimelineUsersValue, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Timeline) SetUsers(v map[string]TimelineUsersValue)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


