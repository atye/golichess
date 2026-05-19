# GamePgn200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Rated** | **bool** |  | 
**Variant** | **string** |  | [default to "standard"]
**Speed** | **string** |  | 
**Perf** | **string** |  | 
**CreatedAt** | **int64** |  | 
**LastMoveAt** | **int64** |  | 
**Status** | **string** |  | 
**Source** | Pointer to **string** |  | [optional] 
**Players** | [**GamePgn200ResponseOneOfPlayers**](GamePgn200ResponseOneOfPlayers.md) |  | 
**InitialFen** | Pointer to **string** |  | [optional] 
**Winner** | Pointer to **NullableString** |  | [optional] 
**Opening** | Pointer to [**GamePgn200ResponseOneOfOpening**](GamePgn200ResponseOneOfOpening.md) |  | [optional] 
**Moves** | Pointer to **string** |  | [optional] 
**Pgn** | Pointer to **string** |  | [optional] 
**DaysPerTurn** | Pointer to **int32** |  | [optional] 
**Analysis** | Pointer to [**[]GamePgn200ResponseOneOfAnalysisInner**](GamePgn200ResponseOneOfAnalysisInner.md) |  | [optional] 
**Tournament** | Pointer to **string** |  | [optional] 
**Swiss** | Pointer to **string** |  | [optional] 
**Clock** | Pointer to [**GamePgn200ResponseOneOfClock**](GamePgn200ResponseOneOfClock.md) |  | [optional] 
**Clocks** | Pointer to **[]int32** |  | [optional] 
**Division** | Pointer to [**GamePgn200ResponseOneOfDivision**](GamePgn200ResponseOneOfDivision.md) |  | [optional] 

## Methods

### NewGamePgn200ResponseOneOf

`func NewGamePgn200ResponseOneOf(id string, rated bool, variant string, speed string, perf string, createdAt int64, lastMoveAt int64, status string, players GamePgn200ResponseOneOfPlayers, ) *GamePgn200ResponseOneOf`

NewGamePgn200ResponseOneOf instantiates a new GamePgn200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamePgn200ResponseOneOfWithDefaults

`func NewGamePgn200ResponseOneOfWithDefaults() *GamePgn200ResponseOneOf`

NewGamePgn200ResponseOneOfWithDefaults instantiates a new GamePgn200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GamePgn200ResponseOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GamePgn200ResponseOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GamePgn200ResponseOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetRated

`func (o *GamePgn200ResponseOneOf) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GamePgn200ResponseOneOf) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GamePgn200ResponseOneOf) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetVariant

`func (o *GamePgn200ResponseOneOf) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GamePgn200ResponseOneOf) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GamePgn200ResponseOneOf) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetSpeed

`func (o *GamePgn200ResponseOneOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GamePgn200ResponseOneOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GamePgn200ResponseOneOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *GamePgn200ResponseOneOf) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GamePgn200ResponseOneOf) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GamePgn200ResponseOneOf) SetPerf(v string)`

SetPerf sets Perf field to given value.


### GetCreatedAt

`func (o *GamePgn200ResponseOneOf) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GamePgn200ResponseOneOf) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GamePgn200ResponseOneOf) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastMoveAt

`func (o *GamePgn200ResponseOneOf) GetLastMoveAt() int64`

GetLastMoveAt returns the LastMoveAt field if non-nil, zero value otherwise.

### GetLastMoveAtOk

`func (o *GamePgn200ResponseOneOf) GetLastMoveAtOk() (*int64, bool)`

GetLastMoveAtOk returns a tuple with the LastMoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMoveAt

`func (o *GamePgn200ResponseOneOf) SetLastMoveAt(v int64)`

SetLastMoveAt sets LastMoveAt field to given value.


### GetStatus

`func (o *GamePgn200ResponseOneOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GamePgn200ResponseOneOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GamePgn200ResponseOneOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *GamePgn200ResponseOneOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GamePgn200ResponseOneOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GamePgn200ResponseOneOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GamePgn200ResponseOneOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPlayers

`func (o *GamePgn200ResponseOneOf) GetPlayers() GamePgn200ResponseOneOfPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GamePgn200ResponseOneOf) GetPlayersOk() (*GamePgn200ResponseOneOfPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GamePgn200ResponseOneOf) SetPlayers(v GamePgn200ResponseOneOfPlayers)`

SetPlayers sets Players field to given value.


### GetInitialFen

`func (o *GamePgn200ResponseOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *GamePgn200ResponseOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *GamePgn200ResponseOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *GamePgn200ResponseOneOf) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetWinner

`func (o *GamePgn200ResponseOneOf) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GamePgn200ResponseOneOf) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GamePgn200ResponseOneOf) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GamePgn200ResponseOneOf) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *GamePgn200ResponseOneOf) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *GamePgn200ResponseOneOf) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetOpening

`func (o *GamePgn200ResponseOneOf) GetOpening() GamePgn200ResponseOneOfOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *GamePgn200ResponseOneOf) GetOpeningOk() (*GamePgn200ResponseOneOfOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *GamePgn200ResponseOneOf) SetOpening(v GamePgn200ResponseOneOfOpening)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *GamePgn200ResponseOneOf) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

### GetMoves

`func (o *GamePgn200ResponseOneOf) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *GamePgn200ResponseOneOf) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *GamePgn200ResponseOneOf) SetMoves(v string)`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *GamePgn200ResponseOneOf) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### GetPgn

`func (o *GamePgn200ResponseOneOf) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *GamePgn200ResponseOneOf) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *GamePgn200ResponseOneOf) SetPgn(v string)`

SetPgn sets Pgn field to given value.

### HasPgn

`func (o *GamePgn200ResponseOneOf) HasPgn() bool`

HasPgn returns a boolean if a field has been set.

### GetDaysPerTurn

`func (o *GamePgn200ResponseOneOf) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *GamePgn200ResponseOneOf) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *GamePgn200ResponseOneOf) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *GamePgn200ResponseOneOf) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetAnalysis

`func (o *GamePgn200ResponseOneOf) GetAnalysis() []GamePgn200ResponseOneOfAnalysisInner`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GamePgn200ResponseOneOf) GetAnalysisOk() (*[]GamePgn200ResponseOneOfAnalysisInner, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GamePgn200ResponseOneOf) SetAnalysis(v []GamePgn200ResponseOneOfAnalysisInner)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GamePgn200ResponseOneOf) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTournament

`func (o *GamePgn200ResponseOneOf) GetTournament() string`

GetTournament returns the Tournament field if non-nil, zero value otherwise.

### GetTournamentOk

`func (o *GamePgn200ResponseOneOf) GetTournamentOk() (*string, bool)`

GetTournamentOk returns a tuple with the Tournament field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournament

`func (o *GamePgn200ResponseOneOf) SetTournament(v string)`

SetTournament sets Tournament field to given value.

### HasTournament

`func (o *GamePgn200ResponseOneOf) HasTournament() bool`

HasTournament returns a boolean if a field has been set.

### GetSwiss

`func (o *GamePgn200ResponseOneOf) GetSwiss() string`

GetSwiss returns the Swiss field if non-nil, zero value otherwise.

### GetSwissOk

`func (o *GamePgn200ResponseOneOf) GetSwissOk() (*string, bool)`

GetSwissOk returns a tuple with the Swiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiss

`func (o *GamePgn200ResponseOneOf) SetSwiss(v string)`

SetSwiss sets Swiss field to given value.

### HasSwiss

`func (o *GamePgn200ResponseOneOf) HasSwiss() bool`

HasSwiss returns a boolean if a field has been set.

### GetClock

`func (o *GamePgn200ResponseOneOf) GetClock() GamePgn200ResponseOneOfClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GamePgn200ResponseOneOf) GetClockOk() (*GamePgn200ResponseOneOfClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GamePgn200ResponseOneOf) SetClock(v GamePgn200ResponseOneOfClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GamePgn200ResponseOneOf) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetClocks

`func (o *GamePgn200ResponseOneOf) GetClocks() []int32`

GetClocks returns the Clocks field if non-nil, zero value otherwise.

### GetClocksOk

`func (o *GamePgn200ResponseOneOf) GetClocksOk() (*[]int32, bool)`

GetClocksOk returns a tuple with the Clocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClocks

`func (o *GamePgn200ResponseOneOf) SetClocks(v []int32)`

SetClocks sets Clocks field to given value.

### HasClocks

`func (o *GamePgn200ResponseOneOf) HasClocks() bool`

HasClocks returns a boolean if a field has been set.

### GetDivision

`func (o *GamePgn200ResponseOneOf) GetDivision() GamePgn200ResponseOneOfDivision`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *GamePgn200ResponseOneOf) GetDivisionOk() (*GamePgn200ResponseOneOfDivision, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *GamePgn200ResponseOneOf) SetDivision(v GamePgn200ResponseOneOfDivision)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *GamePgn200ResponseOneOf) HasDivision() bool`

HasDivision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


