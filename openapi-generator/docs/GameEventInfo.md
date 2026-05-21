# GameEventInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullId** | **string** |  | 
**GameId** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Color** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Source** | Pointer to [**GameSource**](GameSource.md) |  | [optional] 
**Status** | Pointer to [**GameStatus**](GameStatus.md) |  | [optional] 
**Variant** | Pointer to [**Variant**](Variant.md) |  | [optional] 
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**HasMoved** | Pointer to **bool** |  | [optional] 
**Opponent** | Pointer to [**GameEventOpponent**](GameEventOpponent.md) |  | [optional] 
**IsMyTurn** | Pointer to **bool** |  | [optional] 
**SecondsLeft** | Pointer to **int32** |  | [optional] 
**Winner** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Compat** | Pointer to [**GameCompat**](GameCompat.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 

## Methods

### NewGameEventInfo

`func NewGameEventInfo(fullId string, gameId string, ) *GameEventInfo`

NewGameEventInfo instantiates a new GameEventInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameEventInfoWithDefaults

`func NewGameEventInfoWithDefaults() *GameEventInfo`

NewGameEventInfoWithDefaults instantiates a new GameEventInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullId

`func (o *GameEventInfo) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *GameEventInfo) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *GameEventInfo) SetFullId(v string)`

SetFullId sets FullId field to given value.


### GetGameId

`func (o *GameEventInfo) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *GameEventInfo) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *GameEventInfo) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetFen

`func (o *GameEventInfo) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *GameEventInfo) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *GameEventInfo) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *GameEventInfo) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetColor

`func (o *GameEventInfo) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GameEventInfo) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GameEventInfo) SetColor(v GameColor)`

SetColor sets Color field to given value.

### HasColor

`func (o *GameEventInfo) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLastMove

`func (o *GameEventInfo) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *GameEventInfo) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *GameEventInfo) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *GameEventInfo) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetSource

`func (o *GameEventInfo) GetSource() GameSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GameEventInfo) GetSourceOk() (*GameSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GameEventInfo) SetSource(v GameSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *GameEventInfo) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *GameEventInfo) GetStatus() GameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameEventInfo) GetStatusOk() (*GameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameEventInfo) SetStatus(v GameStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GameEventInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariant

`func (o *GameEventInfo) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameEventInfo) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameEventInfo) SetVariant(v Variant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *GameEventInfo) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *GameEventInfo) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameEventInfo) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameEventInfo) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GameEventInfo) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *GameEventInfo) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameEventInfo) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameEventInfo) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *GameEventInfo) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRating

`func (o *GameEventInfo) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GameEventInfo) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GameEventInfo) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GameEventInfo) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetRated

`func (o *GameEventInfo) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameEventInfo) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameEventInfo) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *GameEventInfo) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetHasMoved

`func (o *GameEventInfo) GetHasMoved() bool`

GetHasMoved returns the HasMoved field if non-nil, zero value otherwise.

### GetHasMovedOk

`func (o *GameEventInfo) GetHasMovedOk() (*bool, bool)`

GetHasMovedOk returns a tuple with the HasMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoved

`func (o *GameEventInfo) SetHasMoved(v bool)`

SetHasMoved sets HasMoved field to given value.

### HasHasMoved

`func (o *GameEventInfo) HasHasMoved() bool`

HasHasMoved returns a boolean if a field has been set.

### GetOpponent

`func (o *GameEventInfo) GetOpponent() GameEventOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *GameEventInfo) GetOpponentOk() (*GameEventOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *GameEventInfo) SetOpponent(v GameEventOpponent)`

SetOpponent sets Opponent field to given value.

### HasOpponent

`func (o *GameEventInfo) HasOpponent() bool`

HasOpponent returns a boolean if a field has been set.

### GetIsMyTurn

`func (o *GameEventInfo) GetIsMyTurn() bool`

GetIsMyTurn returns the IsMyTurn field if non-nil, zero value otherwise.

### GetIsMyTurnOk

`func (o *GameEventInfo) GetIsMyTurnOk() (*bool, bool)`

GetIsMyTurnOk returns a tuple with the IsMyTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMyTurn

`func (o *GameEventInfo) SetIsMyTurn(v bool)`

SetIsMyTurn sets IsMyTurn field to given value.

### HasIsMyTurn

`func (o *GameEventInfo) HasIsMyTurn() bool`

HasIsMyTurn returns a boolean if a field has been set.

### GetSecondsLeft

`func (o *GameEventInfo) GetSecondsLeft() int32`

GetSecondsLeft returns the SecondsLeft field if non-nil, zero value otherwise.

### GetSecondsLeftOk

`func (o *GameEventInfo) GetSecondsLeftOk() (*int32, bool)`

GetSecondsLeftOk returns a tuple with the SecondsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsLeft

`func (o *GameEventInfo) SetSecondsLeft(v int32)`

SetSecondsLeft sets SecondsLeft field to given value.

### HasSecondsLeft

`func (o *GameEventInfo) HasSecondsLeft() bool`

HasSecondsLeft returns a boolean if a field has been set.

### GetWinner

`func (o *GameEventInfo) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameEventInfo) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameEventInfo) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameEventInfo) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetRatingDiff

`func (o *GameEventInfo) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *GameEventInfo) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *GameEventInfo) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *GameEventInfo) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetCompat

`func (o *GameEventInfo) GetCompat() GameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *GameEventInfo) GetCompatOk() (*GameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *GameEventInfo) SetCompat(v GameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *GameEventInfo) HasCompat() bool`

HasCompat returns a boolean if a field has been set.

### GetId

`func (o *GameEventInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameEventInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameEventInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GameEventInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTournamentId

`func (o *GameEventInfo) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *GameEventInfo) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *GameEventInfo) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *GameEventInfo) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


