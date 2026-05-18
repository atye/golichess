# ApiPuzzleDaily200ResponseGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clock** | **string** |  | 
**Id** | **string** |  | 
**Perf** | [**ApiPuzzleDaily200ResponseGamePerf**](ApiPuzzleDaily200ResponseGamePerf.md) |  | 
**Pgn** | **string** |  | 
**Players** | [**[]ApiPuzzleDaily200ResponseGamePlayersInner**](ApiPuzzleDaily200ResponseGamePlayersInner.md) |  | 
**Rated** | **bool** |  | 

## Methods

### NewApiPuzzleDaily200ResponseGame

`func NewApiPuzzleDaily200ResponseGame(clock string, id string, perf ApiPuzzleDaily200ResponseGamePerf, pgn string, players []ApiPuzzleDaily200ResponseGamePlayersInner, rated bool, ) *ApiPuzzleDaily200ResponseGame`

NewApiPuzzleDaily200ResponseGame instantiates a new ApiPuzzleDaily200ResponseGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleDaily200ResponseGameWithDefaults

`func NewApiPuzzleDaily200ResponseGameWithDefaults() *ApiPuzzleDaily200ResponseGame`

NewApiPuzzleDaily200ResponseGameWithDefaults instantiates a new ApiPuzzleDaily200ResponseGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClock

`func (o *ApiPuzzleDaily200ResponseGame) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiPuzzleDaily200ResponseGame) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiPuzzleDaily200ResponseGame) SetClock(v string)`

SetClock sets Clock field to given value.


### GetId

`func (o *ApiPuzzleDaily200ResponseGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPuzzleDaily200ResponseGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPuzzleDaily200ResponseGame) SetId(v string)`

SetId sets Id field to given value.


### GetPerf

`func (o *ApiPuzzleDaily200ResponseGame) GetPerf() ApiPuzzleDaily200ResponseGamePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiPuzzleDaily200ResponseGame) GetPerfOk() (*ApiPuzzleDaily200ResponseGamePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiPuzzleDaily200ResponseGame) SetPerf(v ApiPuzzleDaily200ResponseGamePerf)`

SetPerf sets Perf field to given value.


### GetPgn

`func (o *ApiPuzzleDaily200ResponseGame) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *ApiPuzzleDaily200ResponseGame) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *ApiPuzzleDaily200ResponseGame) SetPgn(v string)`

SetPgn sets Pgn field to given value.


### GetPlayers

`func (o *ApiPuzzleDaily200ResponseGame) GetPlayers() []ApiPuzzleDaily200ResponseGamePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *ApiPuzzleDaily200ResponseGame) GetPlayersOk() (*[]ApiPuzzleDaily200ResponseGamePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *ApiPuzzleDaily200ResponseGame) SetPlayers(v []ApiPuzzleDaily200ResponseGamePlayersInner)`

SetPlayers sets Players field to given value.


### GetRated

`func (o *ApiPuzzleDaily200ResponseGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiPuzzleDaily200ResponseGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiPuzzleDaily200ResponseGame) SetRated(v bool)`

SetRated sets Rated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


