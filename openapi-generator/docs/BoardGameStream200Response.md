# BoardGameStream200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**Clock** | Pointer to [**BoardGameStream200ResponseOneOfClock**](BoardGameStream200ResponseOneOfClock.md) |  | [optional] 
**Speed** | **string** |  | 
**Perf** | [**BoardGameStream200ResponseOneOfPerf**](BoardGameStream200ResponseOneOfPerf.md) |  | 
**Rated** | **bool** |  | 
**CreatedAt** | **int64** |  | 
**White** | [**BoardGameStream200ResponseOneOfWhite**](BoardGameStream200ResponseOneOfWhite.md) |  | 
**Black** | [**BoardGameStream200ResponseOneOfWhite**](BoardGameStream200ResponseOneOfWhite.md) |  | 
**InitialFen** | **string** |  | [default to "startpos"]
**State** | [**BoardGameStream200ResponseOneOfState**](BoardGameStream200ResponseOneOfState.md) |  | 
**DaysPerTurn** | Pointer to **int32** | If the game is correspondence | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 
**Moves** | **string** | Current moves in UCI format (King to rook for Chess690-compatible castling notation)  | 
**Wtime** | **int32** | Integer of milliseconds White has left on the clock | 
**Btime** | **int32** | Integer of milliseconds Black has left on the clock | 
**Winc** | **int32** | Integer of White Fisher increment. | 
**Binc** | **int32** | Integer of Black Fisher increment. | 
**Status** | **string** |  | 
**Winner** | Pointer to **string** | Color of the winner, if any | [optional] 
**Wdraw** | Pointer to **bool** | true if white is offering draw, else omitted | [optional] 
**Bdraw** | Pointer to **bool** | true if black is offering draw, else omitted | [optional] 
**Wtakeback** | Pointer to **bool** | true if white is proposing takeback, else omitted | [optional] 
**Btakeback** | Pointer to **bool** | true if black is proposing takeback, else omitted | [optional] 
**Expiration** | Pointer to [**BoardGameStream200ResponseOneOfStateExpiration**](BoardGameStream200ResponseOneOfStateExpiration.md) |  | [optional] 
**Room** | **string** |  | 
**Username** | **string** |  | 
**Text** | **string** |  | 
**Gone** | **bool** |  | 
**ClaimWinInSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewBoardGameStream200Response

`func NewBoardGameStream200Response(type_ string, id string, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, speed string, perf BoardGameStream200ResponseOneOfPerf, rated bool, createdAt int64, white BoardGameStream200ResponseOneOfWhite, black BoardGameStream200ResponseOneOfWhite, initialFen string, state BoardGameStream200ResponseOneOfState, moves string, wtime int32, btime int32, winc int32, binc int32, status string, room string, username string, text string, gone bool, ) *BoardGameStream200Response`

NewBoardGameStream200Response instantiates a new BoardGameStream200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardGameStream200ResponseWithDefaults

`func NewBoardGameStream200ResponseWithDefaults() *BoardGameStream200Response`

NewBoardGameStream200ResponseWithDefaults instantiates a new BoardGameStream200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BoardGameStream200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardGameStream200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardGameStream200Response) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BoardGameStream200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardGameStream200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardGameStream200Response) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *BoardGameStream200Response) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *BoardGameStream200Response) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *BoardGameStream200Response) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *BoardGameStream200Response) GetClock() BoardGameStream200ResponseOneOfClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *BoardGameStream200Response) GetClockOk() (*BoardGameStream200ResponseOneOfClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *BoardGameStream200Response) SetClock(v BoardGameStream200ResponseOneOfClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *BoardGameStream200Response) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetSpeed

`func (o *BoardGameStream200Response) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *BoardGameStream200Response) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *BoardGameStream200Response) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *BoardGameStream200Response) GetPerf() BoardGameStream200ResponseOneOfPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *BoardGameStream200Response) GetPerfOk() (*BoardGameStream200ResponseOneOfPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *BoardGameStream200Response) SetPerf(v BoardGameStream200ResponseOneOfPerf)`

SetPerf sets Perf field to given value.


### GetRated

`func (o *BoardGameStream200Response) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *BoardGameStream200Response) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *BoardGameStream200Response) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetCreatedAt

`func (o *BoardGameStream200Response) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BoardGameStream200Response) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BoardGameStream200Response) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetWhite

`func (o *BoardGameStream200Response) GetWhite() BoardGameStream200ResponseOneOfWhite`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *BoardGameStream200Response) GetWhiteOk() (*BoardGameStream200ResponseOneOfWhite, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *BoardGameStream200Response) SetWhite(v BoardGameStream200ResponseOneOfWhite)`

SetWhite sets White field to given value.


### GetBlack

`func (o *BoardGameStream200Response) GetBlack() BoardGameStream200ResponseOneOfWhite`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *BoardGameStream200Response) GetBlackOk() (*BoardGameStream200ResponseOneOfWhite, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *BoardGameStream200Response) SetBlack(v BoardGameStream200ResponseOneOfWhite)`

SetBlack sets Black field to given value.


### GetInitialFen

`func (o *BoardGameStream200Response) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *BoardGameStream200Response) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *BoardGameStream200Response) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetState

`func (o *BoardGameStream200Response) GetState() BoardGameStream200ResponseOneOfState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BoardGameStream200Response) GetStateOk() (*BoardGameStream200ResponseOneOfState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BoardGameStream200Response) SetState(v BoardGameStream200ResponseOneOfState)`

SetState sets State field to given value.


### GetDaysPerTurn

`func (o *BoardGameStream200Response) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *BoardGameStream200Response) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *BoardGameStream200Response) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *BoardGameStream200Response) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetTournamentId

`func (o *BoardGameStream200Response) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *BoardGameStream200Response) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *BoardGameStream200Response) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *BoardGameStream200Response) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.

### GetMoves

`func (o *BoardGameStream200Response) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *BoardGameStream200Response) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *BoardGameStream200Response) SetMoves(v string)`

SetMoves sets Moves field to given value.


### GetWtime

`func (o *BoardGameStream200Response) GetWtime() int32`

GetWtime returns the Wtime field if non-nil, zero value otherwise.

### GetWtimeOk

`func (o *BoardGameStream200Response) GetWtimeOk() (*int32, bool)`

GetWtimeOk returns a tuple with the Wtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtime

`func (o *BoardGameStream200Response) SetWtime(v int32)`

SetWtime sets Wtime field to given value.


### GetBtime

`func (o *BoardGameStream200Response) GetBtime() int32`

GetBtime returns the Btime field if non-nil, zero value otherwise.

### GetBtimeOk

`func (o *BoardGameStream200Response) GetBtimeOk() (*int32, bool)`

GetBtimeOk returns a tuple with the Btime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtime

`func (o *BoardGameStream200Response) SetBtime(v int32)`

SetBtime sets Btime field to given value.


### GetWinc

`func (o *BoardGameStream200Response) GetWinc() int32`

GetWinc returns the Winc field if non-nil, zero value otherwise.

### GetWincOk

`func (o *BoardGameStream200Response) GetWincOk() (*int32, bool)`

GetWincOk returns a tuple with the Winc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinc

`func (o *BoardGameStream200Response) SetWinc(v int32)`

SetWinc sets Winc field to given value.


### GetBinc

`func (o *BoardGameStream200Response) GetBinc() int32`

GetBinc returns the Binc field if non-nil, zero value otherwise.

### GetBincOk

`func (o *BoardGameStream200Response) GetBincOk() (*int32, bool)`

GetBincOk returns a tuple with the Binc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinc

`func (o *BoardGameStream200Response) SetBinc(v int32)`

SetBinc sets Binc field to given value.


### GetStatus

`func (o *BoardGameStream200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BoardGameStream200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BoardGameStream200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWinner

`func (o *BoardGameStream200Response) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *BoardGameStream200Response) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *BoardGameStream200Response) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *BoardGameStream200Response) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetWdraw

`func (o *BoardGameStream200Response) GetWdraw() bool`

GetWdraw returns the Wdraw field if non-nil, zero value otherwise.

### GetWdrawOk

`func (o *BoardGameStream200Response) GetWdrawOk() (*bool, bool)`

GetWdrawOk returns a tuple with the Wdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdraw

`func (o *BoardGameStream200Response) SetWdraw(v bool)`

SetWdraw sets Wdraw field to given value.

### HasWdraw

`func (o *BoardGameStream200Response) HasWdraw() bool`

HasWdraw returns a boolean if a field has been set.

### GetBdraw

`func (o *BoardGameStream200Response) GetBdraw() bool`

GetBdraw returns the Bdraw field if non-nil, zero value otherwise.

### GetBdrawOk

`func (o *BoardGameStream200Response) GetBdrawOk() (*bool, bool)`

GetBdrawOk returns a tuple with the Bdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdraw

`func (o *BoardGameStream200Response) SetBdraw(v bool)`

SetBdraw sets Bdraw field to given value.

### HasBdraw

`func (o *BoardGameStream200Response) HasBdraw() bool`

HasBdraw returns a boolean if a field has been set.

### GetWtakeback

`func (o *BoardGameStream200Response) GetWtakeback() bool`

GetWtakeback returns the Wtakeback field if non-nil, zero value otherwise.

### GetWtakebackOk

`func (o *BoardGameStream200Response) GetWtakebackOk() (*bool, bool)`

GetWtakebackOk returns a tuple with the Wtakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtakeback

`func (o *BoardGameStream200Response) SetWtakeback(v bool)`

SetWtakeback sets Wtakeback field to given value.

### HasWtakeback

`func (o *BoardGameStream200Response) HasWtakeback() bool`

HasWtakeback returns a boolean if a field has been set.

### GetBtakeback

`func (o *BoardGameStream200Response) GetBtakeback() bool`

GetBtakeback returns the Btakeback field if non-nil, zero value otherwise.

### GetBtakebackOk

`func (o *BoardGameStream200Response) GetBtakebackOk() (*bool, bool)`

GetBtakebackOk returns a tuple with the Btakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtakeback

`func (o *BoardGameStream200Response) SetBtakeback(v bool)`

SetBtakeback sets Btakeback field to given value.

### HasBtakeback

`func (o *BoardGameStream200Response) HasBtakeback() bool`

HasBtakeback returns a boolean if a field has been set.

### GetExpiration

`func (o *BoardGameStream200Response) GetExpiration() BoardGameStream200ResponseOneOfStateExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BoardGameStream200Response) GetExpirationOk() (*BoardGameStream200ResponseOneOfStateExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BoardGameStream200Response) SetExpiration(v BoardGameStream200ResponseOneOfStateExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BoardGameStream200Response) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.

### GetRoom

`func (o *BoardGameStream200Response) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *BoardGameStream200Response) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *BoardGameStream200Response) SetRoom(v string)`

SetRoom sets Room field to given value.


### GetUsername

`func (o *BoardGameStream200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BoardGameStream200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BoardGameStream200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetText

`func (o *BoardGameStream200Response) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *BoardGameStream200Response) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *BoardGameStream200Response) SetText(v string)`

SetText sets Text field to given value.


### GetGone

`func (o *BoardGameStream200Response) GetGone() bool`

GetGone returns the Gone field if non-nil, zero value otherwise.

### GetGoneOk

`func (o *BoardGameStream200Response) GetGoneOk() (*bool, bool)`

GetGoneOk returns a tuple with the Gone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGone

`func (o *BoardGameStream200Response) SetGone(v bool)`

SetGone sets Gone field to given value.


### GetClaimWinInSeconds

`func (o *BoardGameStream200Response) GetClaimWinInSeconds() int32`

GetClaimWinInSeconds returns the ClaimWinInSeconds field if non-nil, zero value otherwise.

### GetClaimWinInSecondsOk

`func (o *BoardGameStream200Response) GetClaimWinInSecondsOk() (*int32, bool)`

GetClaimWinInSecondsOk returns a tuple with the ClaimWinInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimWinInSeconds

`func (o *BoardGameStream200Response) SetClaimWinInSeconds(v int32)`

SetClaimWinInSeconds sets ClaimWinInSeconds field to given value.

### HasClaimWinInSeconds

`func (o *BoardGameStream200Response) HasClaimWinInSeconds() bool`

HasClaimWinInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


