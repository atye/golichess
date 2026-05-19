# ApiStreamEvent200ResponseOneOf1Game

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FullId** | **string** |  | 
**GameId** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Color** | Pointer to **string** |  | [optional] 
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
**Winner** | Pointer to **string** |  | [optional] 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Compat** | Pointer to [**ApiStreamEvent200ResponseOneOfGameCompat**](ApiStreamEvent200ResponseOneOfGameCompat.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOf1Game

`func NewApiStreamEvent200ResponseOneOf1Game(fullId string, gameId string, ) *ApiStreamEvent200ResponseOneOf1Game`

NewApiStreamEvent200ResponseOneOf1Game instantiates a new ApiStreamEvent200ResponseOneOf1Game object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOf1GameWithDefaults

`func NewApiStreamEvent200ResponseOneOf1GameWithDefaults() *ApiStreamEvent200ResponseOneOf1Game`

NewApiStreamEvent200ResponseOneOf1GameWithDefaults instantiates a new ApiStreamEvent200ResponseOneOf1Game object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFullId

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetFullId(v string)`

SetFullId sets FullId field to given value.


### GetGameId

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetFen

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetColor

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasColor() bool`

HasColor returns a boolean if a field has been set.

### GetLastMove

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetSource

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetStatus() StreamGame200ResponseInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetStatus(v StreamGame200ResponseInnerOneOfStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVariant

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRating

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetRated

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetHasMoved

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetHasMoved() bool`

GetHasMoved returns the HasMoved field if non-nil, zero value otherwise.

### GetHasMovedOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetHasMovedOk() (*bool, bool)`

GetHasMovedOk returns a tuple with the HasMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMoved

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetHasMoved(v bool)`

SetHasMoved sets HasMoved field to given value.

### HasHasMoved

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasHasMoved() bool`

HasHasMoved returns a boolean if a field has been set.

### GetOpponent

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetOpponent() ApiStreamEvent200ResponseOneOfGameOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetOpponentOk() (*ApiStreamEvent200ResponseOneOfGameOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetOpponent(v ApiStreamEvent200ResponseOneOfGameOpponent)`

SetOpponent sets Opponent field to given value.

### HasOpponent

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasOpponent() bool`

HasOpponent returns a boolean if a field has been set.

### GetIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetIsMyTurn() bool`

GetIsMyTurn returns the IsMyTurn field if non-nil, zero value otherwise.

### GetIsMyTurnOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetIsMyTurnOk() (*bool, bool)`

GetIsMyTurnOk returns a tuple with the IsMyTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetIsMyTurn(v bool)`

SetIsMyTurn sets IsMyTurn field to given value.

### HasIsMyTurn

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasIsMyTurn() bool`

HasIsMyTurn returns a boolean if a field has been set.

### GetSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSecondsLeft() int32`

GetSecondsLeft returns the SecondsLeft field if non-nil, zero value otherwise.

### GetSecondsLeftOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetSecondsLeftOk() (*int32, bool)`

GetSecondsLeftOk returns a tuple with the SecondsLeft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetSecondsLeft(v int32)`

SetSecondsLeft sets SecondsLeft field to given value.

### HasSecondsLeft

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasSecondsLeft() bool`

HasSecondsLeft returns a boolean if a field has been set.

### GetWinner

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetRatingDiff

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetCompat

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetCompat() ApiStreamEvent200ResponseOneOfGameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetCompatOk() (*ApiStreamEvent200ResponseOneOfGameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetCompat(v ApiStreamEvent200ResponseOneOfGameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasCompat() bool`

HasCompat returns a boolean if a field has been set.

### GetId

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTournamentId

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *ApiStreamEvent200ResponseOneOf1Game) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *ApiStreamEvent200ResponseOneOf1Game) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *ApiStreamEvent200ResponseOneOf1Game) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


