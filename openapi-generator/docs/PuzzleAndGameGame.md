# PuzzleAndGameGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clock** | **string** |  | 
**Id** | **string** |  | 
**Perf** | [**PuzzleAndGameGamePerf**](PuzzleAndGameGamePerf.md) |  | 
**Pgn** | **string** |  | 
**Players** | [**[]PuzzleAndGameGamePlayersInner**](PuzzleAndGameGamePlayersInner.md) |  | 
**Rated** | **bool** |  | 

## Methods

### NewPuzzleAndGameGame

`func NewPuzzleAndGameGame(clock string, id string, perf PuzzleAndGameGamePerf, pgn string, players []PuzzleAndGameGamePlayersInner, rated bool, ) *PuzzleAndGameGame`

NewPuzzleAndGameGame instantiates a new PuzzleAndGameGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleAndGameGameWithDefaults

`func NewPuzzleAndGameGameWithDefaults() *PuzzleAndGameGame`

NewPuzzleAndGameGameWithDefaults instantiates a new PuzzleAndGameGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClock

`func (o *PuzzleAndGameGame) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *PuzzleAndGameGame) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *PuzzleAndGameGame) SetClock(v string)`

SetClock sets Clock field to given value.


### GetId

`func (o *PuzzleAndGameGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleAndGameGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleAndGameGame) SetId(v string)`

SetId sets Id field to given value.


### GetPerf

`func (o *PuzzleAndGameGame) GetPerf() PuzzleAndGameGamePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *PuzzleAndGameGame) GetPerfOk() (*PuzzleAndGameGamePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *PuzzleAndGameGame) SetPerf(v PuzzleAndGameGamePerf)`

SetPerf sets Perf field to given value.


### GetPgn

`func (o *PuzzleAndGameGame) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *PuzzleAndGameGame) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *PuzzleAndGameGame) SetPgn(v string)`

SetPgn sets Pgn field to given value.


### GetPlayers

`func (o *PuzzleAndGameGame) GetPlayers() []PuzzleAndGameGamePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *PuzzleAndGameGame) GetPlayersOk() (*[]PuzzleAndGameGamePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *PuzzleAndGameGame) SetPlayers(v []PuzzleAndGameGamePlayersInner)`

SetPlayers sets Players field to given value.


### GetRated

`func (o *PuzzleAndGameGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *PuzzleAndGameGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *PuzzleAndGameGame) SetRated(v bool)`

SetRated sets Rated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


