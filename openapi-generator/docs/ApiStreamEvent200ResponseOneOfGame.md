# ApiStreamEvent200ResponseOneOfGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullId** | **string** |  | 
**GameId** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StreamGame200ResponseInnerOneOfStatus**](StreamGame200ResponseInnerOneOfStatus.md) |  | [optional] 
**Variant** | Pointer to [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**HasMoved** | Pointer to **bool** |  | [optional] 
**Opponent** | Pointer to [**ApiStreamEvent200ResponseOneOfGameOpponent**](ApiStreamEvent200ResponseOneOfGameOpponent.md) |  | [optional] 
**IsMyTurn** | Pointer to **bool** |  | [optional] 
**SecondsLeft** | Pointer to **int32** |  | [optional] 
**Winner** | Pointer to **NullableString** |  | [optional] 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Compat** | Pointer to [**ApiStreamEvent200ResponseOneOfGameCompat**](ApiStreamEvent200ResponseOneOfGameCompat.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOfGame

`func NewApiStreamEvent200ResponseOneOfGame(fullId string, gameId string, ) *ApiStreamEvent200ResponseOneOfGame`

NewApiStreamEvent200ResponseOneOfGame instantiates a new ApiStreamEvent200ResponseOneOfGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOfGameWithDefaults

`func NewApiStreamEvent200ResponseOneOfGameWithDefaults() *ApiStreamEvent200ResponseOneOfGame`

NewApiStreamEvent200ResponseOneOfGameWithDefaults instantiates a new ApiStreamEvent200ResponseOneOfGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullId

`func (o *ApiStreamEvent200ResponseOneOfGame) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *ApiStreamEvent200ResponseOneOfGame) SetFullId(v string)`

SetFullId sets FullId field to given value.


### GetGameId

`func (o *ApiStreamEvent200ResponseOneOfGame) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *ApiStreamEvent200ResponseOneOfGame) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetFen

`func (o *ApiStreamEvent200ResponseOneOfGame) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ApiStreamEvent200ResponseOneOfGame) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *ApiStreamEvent200ResponseOneOfGame) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetColor

`func (o *ApiStreamEvent200ResponseOneOfGame) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiStreamEvent200ResponseOneOfGame) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ApiStreamEvent200ResponseOneOfGame) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ApiStreamEvent200ResponseOneOfGame) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ApiStreamEvent200ResponseOneOfGame) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetLastMove

`func (o *ApiStreamEvent200ResponseOneOfGame) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *ApiStreamEvent200ResponseOneOfGame) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *ApiStreamEvent200ResponseOneOfGame) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetSource

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiStreamEvent200ResponseOneOfGame) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiStreamEvent200ResponseOneOfGame) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ApiStreamEvent200ResponseOneOfGame) GetStatus() StreamGame200ResponseInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStreamEvent200ResponseOneOfGame) SetStatus(v StreamGame200ResponseInnerOneOfStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiStreamEvent200ResponseOneOfGame) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariant

`func (o *ApiStreamEvent200ResponseOneOfGame) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiStreamEvent200ResponseOneOfGame) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ApiStreamEvent200ResponseOneOfGame) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiStreamEvent200ResponseOneOfGame) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ApiStreamEvent200ResponseOneOfGame) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *ApiStreamEvent200ResponseOneOfGame) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiStreamEvent200ResponseOneOfGame) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ApiStreamEvent200ResponseOneOfGame) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRating

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiStreamEvent200ResponseOneOfGame) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiStreamEvent200ResponseOneOfGame) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetRated

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiStreamEvent200ResponseOneOfGame) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ApiStreamEvent200ResponseOneOfGame) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetHasMoved

`func (o *ApiStreamEvent200ResponseOneOfGame) GetHasMoved() bool`

GetHasMoved returns the HasMoved field if non-nil, zero value otherwise.

### GetHasMovedOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetHasMovedOk() (*bool, bool)`

GetHasMovedOk returns a tuple with the HasMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoved

`func (o *ApiStreamEvent200ResponseOneOfGame) SetHasMoved(v bool)`

SetHasMoved sets HasMoved field to given value.

### HasHasMoved

`func (o *ApiStreamEvent200ResponseOneOfGame) HasHasMoved() bool`

HasHasMoved returns a boolean if a field has been set.

### GetOpponent

`func (o *ApiStreamEvent200ResponseOneOfGame) GetOpponent() ApiStreamEvent200ResponseOneOfGameOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetOpponentOk() (*ApiStreamEvent200ResponseOneOfGameOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *ApiStreamEvent200ResponseOneOfGame) SetOpponent(v ApiStreamEvent200ResponseOneOfGameOpponent)`

SetOpponent sets Opponent field to given value.

### HasOpponent

`func (o *ApiStreamEvent200ResponseOneOfGame) HasOpponent() bool`

HasOpponent returns a boolean if a field has been set.

### GetIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOfGame) GetIsMyTurn() bool`

GetIsMyTurn returns the IsMyTurn field if non-nil, zero value otherwise.

### GetIsMyTurnOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetIsMyTurnOk() (*bool, bool)`

GetIsMyTurnOk returns a tuple with the IsMyTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOfGame) SetIsMyTurn(v bool)`

SetIsMyTurn sets IsMyTurn field to given value.

### HasIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOfGame) HasIsMyTurn() bool`

HasIsMyTurn returns a boolean if a field has been set.

### GetSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSecondsLeft() int32`

GetSecondsLeft returns the SecondsLeft field if non-nil, zero value otherwise.

### GetSecondsLeftOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetSecondsLeftOk() (*int32, bool)`

GetSecondsLeftOk returns a tuple with the SecondsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOfGame) SetSecondsLeft(v int32)`

SetSecondsLeft sets SecondsLeft field to given value.

### HasSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOfGame) HasSecondsLeft() bool`

HasSecondsLeft returns a boolean if a field has been set.

### GetWinner

`func (o *ApiStreamEvent200ResponseOneOfGame) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ApiStreamEvent200ResponseOneOfGame) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ApiStreamEvent200ResponseOneOfGame) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *ApiStreamEvent200ResponseOneOfGame) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *ApiStreamEvent200ResponseOneOfGame) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetRatingDiff

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *ApiStreamEvent200ResponseOneOfGame) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *ApiStreamEvent200ResponseOneOfGame) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetCompat

`func (o *ApiStreamEvent200ResponseOneOfGame) GetCompat() ApiStreamEvent200ResponseOneOfGameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetCompatOk() (*ApiStreamEvent200ResponseOneOfGameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiStreamEvent200ResponseOneOfGame) SetCompat(v ApiStreamEvent200ResponseOneOfGameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ApiStreamEvent200ResponseOneOfGame) HasCompat() bool`

HasCompat returns a boolean if a field has been set.

### GetId

`func (o *ApiStreamEvent200ResponseOneOfGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStreamEvent200ResponseOneOfGame) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiStreamEvent200ResponseOneOfGame) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTournamentId

`func (o *ApiStreamEvent200ResponseOneOfGame) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *ApiStreamEvent200ResponseOneOfGame) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *ApiStreamEvent200ResponseOneOfGame) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *ApiStreamEvent200ResponseOneOfGame) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


