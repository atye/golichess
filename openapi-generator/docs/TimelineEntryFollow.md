# TimelineEntryFollow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Date** | **float32** |  | 
**Data** | [**TimelineEntryFollowData**](TimelineEntryFollowData.md) |  | 

## Methods

### NewTimelineEntryFollow

`func NewTimelineEntryFollow(type_ string, date float32, data TimelineEntryFollowData, ) *TimelineEntryFollow`

NewTimelineEntryFollow instantiates a new TimelineEntryFollow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEntryFollowWithDefaults

`func NewTimelineEntryFollowWithDefaults() *TimelineEntryFollow`

NewTimelineEntryFollowWithDefaults instantiates a new TimelineEntryFollow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimelineEntryFollow) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineEntryFollow) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineEntryFollow) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *TimelineEntryFollow) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimelineEntryFollow) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimelineEntryFollow) SetDate(v float32)`

SetDate sets Date field to given value.


### GetData

`func (o *TimelineEntryFollow) GetData() TimelineEntryFollowData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TimelineEntryFollow) GetDataOk() (*TimelineEntryFollowData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TimelineEntryFollow) SetData(v TimelineEntryFollowData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


