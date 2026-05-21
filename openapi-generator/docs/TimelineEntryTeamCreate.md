# TimelineEntryTeamCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Date** | **float32** |  | 
**Data** | [**TimelineEntryTeamJoinData**](TimelineEntryTeamJoinData.md) |  | 

## Methods

### NewTimelineEntryTeamCreate

`func NewTimelineEntryTeamCreate(type_ string, date float32, data TimelineEntryTeamJoinData, ) *TimelineEntryTeamCreate`

NewTimelineEntryTeamCreate instantiates a new TimelineEntryTeamCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEntryTeamCreateWithDefaults

`func NewTimelineEntryTeamCreateWithDefaults() *TimelineEntryTeamCreate`

NewTimelineEntryTeamCreateWithDefaults instantiates a new TimelineEntryTeamCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TimelineEntryTeamCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TimelineEntryTeamCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TimelineEntryTeamCreate) SetType(v string)`

SetType sets Type field to given value.


### GetDate

`func (o *TimelineEntryTeamCreate) GetDate() float32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TimelineEntryTeamCreate) GetDateOk() (*float32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TimelineEntryTeamCreate) SetDate(v float32)`

SetDate sets Date field to given value.


### GetData

`func (o *TimelineEntryTeamCreate) GetData() TimelineEntryTeamJoinData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TimelineEntryTeamCreate) GetDataOk() (*TimelineEntryTeamJoinData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TimelineEntryTeamCreate) SetData(v TimelineEntryTeamJoinData)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


