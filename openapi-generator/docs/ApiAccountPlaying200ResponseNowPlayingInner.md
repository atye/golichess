# ApiAccountPlaying200ResponseNowPlayingInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullId** | **string** |  | 
**GameId** | **string** |  | 
**Fen** | **string** |  | 
**Color** | [**GameColor**](GameColor.md) |  | 
**LastMove** | **string** |  | 
**Source** | [**GameSource**](GameSource.md) |  | 
**Status** | Pointer to [**GameStatusName**](GameStatusName.md) |  | [optional] 
**Variant** | [**Variant**](Variant.md) |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**Perf** | [**PerfType**](PerfType.md) |  | 
**Rated** | **bool** |  | 
**HasMoved** | **bool** |  | 
**Opponent** | [**ApiAccountPlaying200ResponseNowPlayingInnerOpponent**](ApiAccountPlaying200ResponseNowPlayingInnerOpponent.md) |  | 
**IsMyTurn** | **bool** |  | 
**SecondsLeft** | **int32** |  | 
**TournamentId** | Pointer to **string** |  | [optional] 
**SwissId** | Pointer to **string** |  | [optional] 
**Winner** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**RatingDiff** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiAccountPlaying200ResponseNowPlayingInner

`func NewApiAccountPlaying200ResponseNowPlayingInner(fullId string, gameId string, fen string, color GameColor, lastMove string, source GameSource, variant Variant, speed Speed, perf PerfType, rated bool, hasMoved bool, opponent ApiAccountPlaying200ResponseNowPlayingInnerOpponent, isMyTurn bool, secondsLeft int32, ) *ApiAccountPlaying200ResponseNowPlayingInner`

NewApiAccountPlaying200ResponseNowPlayingInner instantiates a new ApiAccountPlaying200ResponseNowPlayingInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiAccountPlaying200ResponseNowPlayingInnerWithDefaults

`func NewApiAccountPlaying200ResponseNowPlayingInnerWithDefaults() *ApiAccountPlaying200ResponseNowPlayingInner`

NewApiAccountPlaying200ResponseNowPlayingInnerWithDefaults instantiates a new ApiAccountPlaying200ResponseNowPlayingInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetFullId(v string)`

SetFullId sets FullId field to given value.


### GetGameId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetFen

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetFen(v string)`

SetFen sets Fen field to given value.


### GetColor

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetColor(v GameColor)`

SetColor sets Color field to given value.


### GetLastMove

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.


### GetSource

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSource() GameSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSourceOk() (*GameSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSource(v GameSource)`

SetSource sets Source field to given value.


### GetStatus

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetStatus() GameStatusName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetStatusOk() (*GameStatusName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetStatus(v GameStatusName)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariant

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetSpeed

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.


### GetRated

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetHasMoved

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetHasMoved() bool`

GetHasMoved returns the HasMoved field if non-nil, zero value otherwise.

### GetHasMovedOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetHasMovedOk() (*bool, bool)`

GetHasMovedOk returns a tuple with the HasMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoved

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetHasMoved(v bool)`

SetHasMoved sets HasMoved field to given value.


### GetOpponent

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetOpponent() ApiAccountPlaying200ResponseNowPlayingInnerOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetOpponentOk() (*ApiAccountPlaying200ResponseNowPlayingInnerOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetOpponent(v ApiAccountPlaying200ResponseNowPlayingInnerOpponent)`

SetOpponent sets Opponent field to given value.


### GetIsMyTurn

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetIsMyTurn() bool`

GetIsMyTurn returns the IsMyTurn field if non-nil, zero value otherwise.

### GetIsMyTurnOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetIsMyTurnOk() (*bool, bool)`

GetIsMyTurnOk returns a tuple with the IsMyTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMyTurn

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetIsMyTurn(v bool)`

SetIsMyTurn sets IsMyTurn field to given value.


### GetSecondsLeft

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSecondsLeft() int32`

GetSecondsLeft returns the SecondsLeft field if non-nil, zero value otherwise.

### GetSecondsLeftOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSecondsLeftOk() (*int32, bool)`

GetSecondsLeftOk returns a tuple with the SecondsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsLeft

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSecondsLeft(v int32)`

SetSecondsLeft sets SecondsLeft field to given value.


### GetTournamentId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.

### GetSwissId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSwissId() string`

GetSwissId returns the SwissId field if non-nil, zero value otherwise.

### GetSwissIdOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSwissIdOk() (*string, bool)`

GetSwissIdOk returns a tuple with the SwissId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwissId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSwissId(v string)`

SetSwissId sets SwissId field to given value.

### HasSwissId

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasSwissId() bool`

HasSwissId returns a boolean if a field has been set.

### GetWinner

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetRatingDiff

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


