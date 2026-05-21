# TimelineEntryGameEndData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullId** | **string** |  | 
**Opponent** | **string** |  | 
**Win** | **bool** |  | 
**Perf** | [**PerfType**](PerfType.md) |  | 

## Methods

### NewTimelineEntryGameEndData

`func NewTimelineEntryGameEndData(fullId string, opponent string, win bool, perf PerfType, ) *TimelineEntryGameEndData`

NewTimelineEntryGameEndData instantiates a new TimelineEntryGameEndData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineEntryGameEndDataWithDefaults

`func NewTimelineEntryGameEndDataWithDefaults() *TimelineEntryGameEndData`

NewTimelineEntryGameEndDataWithDefaults instantiates a new TimelineEntryGameEndData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullId

`func (o *TimelineEntryGameEndData) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *TimelineEntryGameEndData) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *TimelineEntryGameEndData) SetFullId(v string)`

SetFullId sets FullId field to given value.


### GetOpponent

`func (o *TimelineEntryGameEndData) GetOpponent() string`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *TimelineEntryGameEndData) GetOpponentOk() (*string, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *TimelineEntryGameEndData) SetOpponent(v string)`

SetOpponent sets Opponent field to given value.


### GetWin

`func (o *TimelineEntryGameEndData) GetWin() bool`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *TimelineEntryGameEndData) GetWinOk() (*bool, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *TimelineEntryGameEndData) SetWin(v bool)`

SetWin sets Win field to given value.


### GetPerf

`func (o *TimelineEntryGameEndData) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *TimelineEntryGameEndData) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *TimelineEntryGameEndData) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


