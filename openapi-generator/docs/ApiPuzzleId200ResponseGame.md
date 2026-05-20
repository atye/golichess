# ApiPuzzleId200ResponseGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clock** | **string** |  | 
**Id** | **string** |  | 
**Perf** | [**ApiPuzzleDaily200ResponseGamePerf**](ApiPuzzleDaily200ResponseGamePerf.md) |  | 
**Pgn** | **string** |  | 
**Players** | [**[]ApiPuzzleId200ResponseGamePlayersInner**](ApiPuzzleId200ResponseGamePlayersInner.md) |  | 
**Rated** | **bool** |  | 

## Methods

### NewApiPuzzleId200ResponseGame

`func NewApiPuzzleId200ResponseGame(clock string, id string, perf ApiPuzzleDaily200ResponseGamePerf, pgn string, players []ApiPuzzleId200ResponseGamePlayersInner, rated bool, ) *ApiPuzzleId200ResponseGame`

NewApiPuzzleId200ResponseGame instantiates a new ApiPuzzleId200ResponseGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleId200ResponseGameWithDefaults

`func NewApiPuzzleId200ResponseGameWithDefaults() *ApiPuzzleId200ResponseGame`

NewApiPuzzleId200ResponseGameWithDefaults instantiates a new ApiPuzzleId200ResponseGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClock

`func (o *ApiPuzzleId200ResponseGame) GetClock() string`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiPuzzleId200ResponseGame) GetClockOk() (*string, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiPuzzleId200ResponseGame) SetClock(v string)`

SetClock sets Clock field to given value.


### GetId

`func (o *ApiPuzzleId200ResponseGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPuzzleId200ResponseGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPuzzleId200ResponseGame) SetId(v string)`

SetId sets Id field to given value.


### GetPerf

`func (o *ApiPuzzleId200ResponseGame) GetPerf() ApiPuzzleDaily200ResponseGamePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiPuzzleId200ResponseGame) GetPerfOk() (*ApiPuzzleDaily200ResponseGamePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiPuzzleId200ResponseGame) SetPerf(v ApiPuzzleDaily200ResponseGamePerf)`

SetPerf sets Perf field to given value.


### GetPgn

`func (o *ApiPuzzleId200ResponseGame) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *ApiPuzzleId200ResponseGame) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *ApiPuzzleId200ResponseGame) SetPgn(v string)`

SetPgn sets Pgn field to given value.


### GetPlayers

`func (o *ApiPuzzleId200ResponseGame) GetPlayers() []ApiPuzzleId200ResponseGamePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *ApiPuzzleId200ResponseGame) GetPlayersOk() (*[]ApiPuzzleId200ResponseGamePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *ApiPuzzleId200ResponseGame) SetPlayers(v []ApiPuzzleId200ResponseGamePlayersInner)`

SetPlayers sets Players field to given value.


### GetRated

`func (o *ApiPuzzleId200ResponseGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiPuzzleId200ResponseGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiPuzzleId200ResponseGame) SetRated(v bool)`

SetRated sets Rated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


