# GameJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Rated** | **bool** |  | 
**Variant** | [**VariantKey**](VariantKey.md) |  | [default to VARIANTKEY_STANDARD]
**Speed** | [**Speed**](Speed.md) |  | 
**Perf** | **string** |  | 
**CreatedAt** | **int64** |  | 
**LastMoveAt** | **int64** |  | 
**Status** | [**GameStatusName**](GameStatusName.md) |  | 
**Source** | Pointer to **string** |  | [optional] 
**Players** | [**GamePlayers**](GamePlayers.md) |  | 
**InitialFen** | Pointer to **string** |  | [optional] 
**Winner** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**Opening** | Pointer to [**GameOpening**](GameOpening.md) |  | [optional] 
**Moves** | Pointer to **string** |  | [optional] 
**Pgn** | Pointer to **string** |  | [optional] 
**DaysPerTurn** | Pointer to **int32** |  | [optional] 
**Analysis** | Pointer to [**[]GameMoveAnalysis**](GameMoveAnalysis.md) |  | [optional] 
**ArenaTour** | Pointer to [**GameJsonArenaTour**](GameJsonArenaTour.md) |  | [optional] 
**SwissTour** | Pointer to [**ApiStudyPost200Response**](ApiStudyPost200Response.md) |  | [optional] 
**Clock** | Pointer to [**GameJsonClock**](GameJsonClock.md) |  | [optional] 
**Clocks** | Pointer to **[]int32** |  | [optional] 
**Division** | Pointer to [**GameJsonDivision**](GameJsonDivision.md) |  | [optional] 

## Methods

### NewGameJson

`func NewGameJson(id string, rated bool, variant VariantKey, speed Speed, perf string, createdAt int64, lastMoveAt int64, status GameStatusName, players GamePlayers, ) *GameJson`

NewGameJson instantiates a new GameJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameJsonWithDefaults

`func NewGameJsonWithDefaults() *GameJson`

NewGameJsonWithDefaults instantiates a new GameJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameJson) SetId(v string)`

SetId sets Id field to given value.


### GetRated

`func (o *GameJson) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameJson) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameJson) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetVariant

`func (o *GameJson) GetVariant() VariantKey`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameJson) GetVariantOk() (*VariantKey, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameJson) SetVariant(v VariantKey)`

SetVariant sets Variant field to given value.


### GetSpeed

`func (o *GameJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *GameJson) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameJson) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameJson) SetPerf(v string)`

SetPerf sets Perf field to given value.


### GetCreatedAt

`func (o *GameJson) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameJson) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameJson) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastMoveAt

`func (o *GameJson) GetLastMoveAt() int64`

GetLastMoveAt returns the LastMoveAt field if non-nil, zero value otherwise.

### GetLastMoveAtOk

`func (o *GameJson) GetLastMoveAtOk() (*int64, bool)`

GetLastMoveAtOk returns a tuple with the LastMoveAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMoveAt

`func (o *GameJson) SetLastMoveAt(v int64)`

SetLastMoveAt sets LastMoveAt field to given value.


### GetStatus

`func (o *GameJson) GetStatus() GameStatusName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameJson) GetStatusOk() (*GameStatusName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameJson) SetStatus(v GameStatusName)`

SetStatus sets Status field to given value.


### GetSource

`func (o *GameJson) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GameJson) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GameJson) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GameJson) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPlayers

`func (o *GameJson) GetPlayers() GamePlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GameJson) GetPlayersOk() (*GamePlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GameJson) SetPlayers(v GamePlayers)`

SetPlayers sets Players field to given value.


### GetInitialFen

`func (o *GameJson) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *GameJson) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *GameJson) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *GameJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetWinner

`func (o *GameJson) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameJson) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameJson) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameJson) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetOpening

`func (o *GameJson) GetOpening() GameOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *GameJson) GetOpeningOk() (*GameOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *GameJson) SetOpening(v GameOpening)`

SetOpening sets Opening field to given value.

### HasOpening

`func (o *GameJson) HasOpening() bool`

HasOpening returns a boolean if a field has been set.

### GetMoves

`func (o *GameJson) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *GameJson) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *GameJson) SetMoves(v string)`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *GameJson) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### GetPgn

`func (o *GameJson) GetPgn() string`

GetPgn returns the Pgn field if non-nil, zero value otherwise.

### GetPgnOk

`func (o *GameJson) GetPgnOk() (*string, bool)`

GetPgnOk returns a tuple with the Pgn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPgn

`func (o *GameJson) SetPgn(v string)`

SetPgn sets Pgn field to given value.

### HasPgn

`func (o *GameJson) HasPgn() bool`

HasPgn returns a boolean if a field has been set.

### GetDaysPerTurn

`func (o *GameJson) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *GameJson) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *GameJson) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *GameJson) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetAnalysis

`func (o *GameJson) GetAnalysis() []GameMoveAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GameJson) GetAnalysisOk() (*[]GameMoveAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GameJson) SetAnalysis(v []GameMoveAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GameJson) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetArenaTour

`func (o *GameJson) GetArenaTour() GameJsonArenaTour`

GetArenaTour returns the ArenaTour field if non-nil, zero value otherwise.

### GetArenaTourOk

`func (o *GameJson) GetArenaTourOk() (*GameJsonArenaTour, bool)`

GetArenaTourOk returns a tuple with the ArenaTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArenaTour

`func (o *GameJson) SetArenaTour(v GameJsonArenaTour)`

SetArenaTour sets ArenaTour field to given value.

### HasArenaTour

`func (o *GameJson) HasArenaTour() bool`

HasArenaTour returns a boolean if a field has been set.

### GetSwissTour

`func (o *GameJson) GetSwissTour() ApiStudyPost200Response`

GetSwissTour returns the SwissTour field if non-nil, zero value otherwise.

### GetSwissTourOk

`func (o *GameJson) GetSwissTourOk() (*ApiStudyPost200Response, bool)`

GetSwissTourOk returns a tuple with the SwissTour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissTour

`func (o *GameJson) SetSwissTour(v ApiStudyPost200Response)`

SetSwissTour sets SwissTour field to given value.

### HasSwissTour

`func (o *GameJson) HasSwissTour() bool`

HasSwissTour returns a boolean if a field has been set.

### GetClock

`func (o *GameJson) GetClock() GameJsonClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GameJson) GetClockOk() (*GameJsonClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GameJson) SetClock(v GameJsonClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GameJson) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetClocks

`func (o *GameJson) GetClocks() []int32`

GetClocks returns the Clocks field if non-nil, zero value otherwise.

### GetClocksOk

`func (o *GameJson) GetClocksOk() (*[]int32, bool)`

GetClocksOk returns a tuple with the Clocks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClocks

`func (o *GameJson) SetClocks(v []int32)`

SetClocks sets Clocks field to given value.

### HasClocks

`func (o *GameJson) HasClocks() bool`

HasClocks returns a boolean if a field has been set.

### GetDivision

`func (o *GameJson) GetDivision() GameJsonDivision`

GetDivision returns the Division field if non-nil, zero value otherwise.

### GetDivisionOk

`func (o *GameJson) GetDivisionOk() (*GameJsonDivision, bool)`

GetDivisionOk returns a tuple with the Division field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivision

`func (o *GameJson) SetDivision(v GameJsonDivision)`

SetDivision sets Division field to given value.

### HasDivision

`func (o *GameJson) HasDivision() bool`

HasDivision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


