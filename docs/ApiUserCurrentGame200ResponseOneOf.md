# ApiUserCurrentGame200ResponseOneOf

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
**Players** | [**ApiUserCurrentGame200ResponseOneOfPlayers**](ApiUserCurrentGame200ResponseOneOfPlayers.md) |  | 
**InitialFen** | Pointer to **string** |  | [optional] 
**Winner** | Pointer to **string** |  | [optional] 
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

### NewApiUserCurrentGame200ResponseOneOf

`func NewApiUserCurrentGame200ResponseOneOf(id string, rated bool, variant string, speed string, perf string, createdAt int64, lastMoveAt int64, status string, players ApiUserCurrentGame200ResponseOneOfPlayers, ) *ApiUserCurrentGame200ResponseOneOf`

NewApiUserCurrentGame200ResponseOneOf instantiates a new ApiUserCurrentGame200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserCurrentGame200ResponseOneOfWithDefaults

`func NewApiUserCurrentGame200ResponseOneOfWithDefaults() *ApiUserCurrentGame200ResponseOneOf`

NewApiUserCurrentGame200ResponseOneOfWithDefaults instantiates a new ApiUserCurrentGame200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUserCurrentGame200ResponseOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserCurrentGame200ResponseOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetRated

`func (o *ApiUserCurrentGame200ResponseOneOf) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiUserCurrentGame200ResponseOneOf) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetVariant

`func (o *ApiUserCurrentGame200ResponseOneOf) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiUserCurrentGame200ResponseOneOf) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetSpeed

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiUserCurrentGame200ResponseOneOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiUserCurrentGame200ResponseOneOf) SetPerf(v string)`

SetPerf sets Perf field to given value.


### GetCreatedAt

`func (o *ApiUserCurrentGame200ResponseOneOf) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiUserCurrentGame200ResponseOneOf) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastMoveAt

`func (o *ApiUserCurrentGame200ResponseOneOf) GetLastMoveAt() int64`

GetLastMoveAt returns the LastMoveAt field if non-nil, zero value otherwise.

### GetLastMoveAtOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetLastMoveAtOk() (*int64, bool)`

GetLastMoveAtOk returns a tuple with the LastMoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMoveAt

`func (o *ApiUserCurrentGame200ResponseOneOf) SetLastMoveAt(v int64)`

SetLastMoveAt sets LastMoveAt field to given value.


### GetStatus

`func (o *ApiUserCurrentGame200ResponseOneOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiUserCurrentGame200ResponseOneOf) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiUserCurrentGame200ResponseOneOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiUserCurrentGame200ResponseOneOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPlayers

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPlayers() ApiUserCurrentGame200ResponseOneOfPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPlayersOk() (*ApiUserCurrentGame200ResponseOneOfPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *ApiUserCurrentGame200ResponseOneOf) SetPlayers(v ApiUserCurrentGame200ResponseOneOfPlayers)`

SetPlayers sets Players field to given value.


### GetInitialFen

`func (o *ApiUserCurrentGame200ResponseOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiUserCurrentGame200ResponseOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ApiUserCurrentGame200ResponseOneOf) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetWinner

`func (o *ApiUserCurrentGame200ResponseOneOf) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ApiUserCurrentGame200ResponseOneOf) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ApiUserCurrentGame200ResponseOneOf) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetOpening

`func (o *ApiUserCurrentGame200ResponseOneOf) GetOpening() GamePgn200ResponseOneOfOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetOpeningOk() (*GamePgn200ResponseOneOfOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *ApiUserCurrentGame200ResponseOneOf) SetOpening(v GamePgn200ResponseOneOfOpening)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *ApiUserCurrentGame200ResponseOneOf) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

### GetMoves

`func (o *ApiUserCurrentGame200ResponseOneOf) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiUserCurrentGame200ResponseOneOf) SetMoves(v string)`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *ApiUserCurrentGame200ResponseOneOf) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### GetPgn

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *ApiUserCurrentGame200ResponseOneOf) SetPgn(v string)`

SetPgn sets Pgn field to given value.

### HasPgn

`func (o *ApiUserCurrentGame200ResponseOneOf) HasPgn() bool`

HasPgn returns a boolean if a field has been set.

### GetDaysPerTurn

`func (o *ApiUserCurrentGame200ResponseOneOf) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *ApiUserCurrentGame200ResponseOneOf) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *ApiUserCurrentGame200ResponseOneOf) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOf) GetAnalysis() []GamePgn200ResponseOneOfAnalysisInner`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetAnalysisOk() (*[]GamePgn200ResponseOneOfAnalysisInner, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOf) SetAnalysis(v []GamePgn200ResponseOneOfAnalysisInner)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOf) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTournament

`func (o *ApiUserCurrentGame200ResponseOneOf) GetTournament() string`

GetTournament returns the Tournament field if non-nil, zero value otherwise.

### GetTournamentOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetTournamentOk() (*string, bool)`

GetTournamentOk returns a tuple with the Tournament field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournament

`func (o *ApiUserCurrentGame200ResponseOneOf) SetTournament(v string)`

SetTournament sets Tournament field to given value.

### HasTournament

`func (o *ApiUserCurrentGame200ResponseOneOf) HasTournament() bool`

HasTournament returns a boolean if a field has been set.

### GetSwiss

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSwiss() string`

GetSwiss returns the Swiss field if non-nil, zero value otherwise.

### GetSwissOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetSwissOk() (*string, bool)`

GetSwissOk returns a tuple with the Swiss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwiss

`func (o *ApiUserCurrentGame200ResponseOneOf) SetSwiss(v string)`

SetSwiss sets Swiss field to given value.

### HasSwiss

`func (o *ApiUserCurrentGame200ResponseOneOf) HasSwiss() bool`

HasSwiss returns a boolean if a field has been set.

### GetClock

`func (o *ApiUserCurrentGame200ResponseOneOf) GetClock() GamePgn200ResponseOneOfClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetClockOk() (*GamePgn200ResponseOneOfClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiUserCurrentGame200ResponseOneOf) SetClock(v GamePgn200ResponseOneOfClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *ApiUserCurrentGame200ResponseOneOf) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetClocks

`func (o *ApiUserCurrentGame200ResponseOneOf) GetClocks() []int32`

GetClocks returns the Clocks field if non-nil, zero value otherwise.

### GetClocksOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetClocksOk() (*[]int32, bool)`

GetClocksOk returns a tuple with the Clocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClocks

`func (o *ApiUserCurrentGame200ResponseOneOf) SetClocks(v []int32)`

SetClocks sets Clocks field to given value.

### HasClocks

`func (o *ApiUserCurrentGame200ResponseOneOf) HasClocks() bool`

HasClocks returns a boolean if a field has been set.

### GetDivision

`func (o *ApiUserCurrentGame200ResponseOneOf) GetDivision() GamePgn200ResponseOneOfDivision`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *ApiUserCurrentGame200ResponseOneOf) GetDivisionOk() (*GamePgn200ResponseOneOfDivision, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *ApiUserCurrentGame200ResponseOneOf) SetDivision(v GamePgn200ResponseOneOfDivision)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *ApiUserCurrentGame200ResponseOneOf) HasDivision() bool`

HasDivision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


