# GameStateEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Moves** | **string** | Current moves in UCI format (King to rook for Chess690-compatible castling notation)  | 
**Wtime** | **int32** | Integer of milliseconds White has left on the clock | 
**Btime** | **int32** | Integer of milliseconds Black has left on the clock | 
**Winc** | **int32** | Integer of White Fisher increment. | 
**Binc** | **int32** | Integer of Black Fisher increment. | 
**Status** | [**GameStatusName**](GameStatusName.md) |  | 
**Winner** | Pointer to [**GameColor**](GameColor.md) | Color of the winner, if any | [optional] 
**Wdraw** | Pointer to **bool** | true if white is offering draw, else omitted | [optional] 
**Bdraw** | Pointer to **bool** | true if black is offering draw, else omitted | [optional] 
**Wtakeback** | Pointer to **bool** | true if white is proposing takeback, else omitted | [optional] 
**Btakeback** | Pointer to **bool** | true if black is proposing takeback, else omitted | [optional] 
**Expiration** | Pointer to [**GameStateEventExpiration**](GameStateEventExpiration.md) |  | [optional] 

## Methods

### NewGameStateEvent

`func NewGameStateEvent(type_ string, moves string, wtime int32, btime int32, winc int32, binc int32, status GameStatusName, ) *GameStateEvent`

NewGameStateEvent instantiates a new GameStateEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameStateEventWithDefaults

`func NewGameStateEventWithDefaults() *GameStateEvent`

NewGameStateEventWithDefaults instantiates a new GameStateEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameStateEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameStateEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameStateEvent) SetType(v string)`

SetType sets Type field to given value.


### GetMoves

`func (o *GameStateEvent) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *GameStateEvent) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *GameStateEvent) SetMoves(v string)`

SetMoves sets Moves field to given value.


### GetWtime

`func (o *GameStateEvent) GetWtime() int32`

GetWtime returns the Wtime field if non-nil, zero value otherwise.

### GetWtimeOk

`func (o *GameStateEvent) GetWtimeOk() (*int32, bool)`

GetWtimeOk returns a tuple with the Wtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtime

`func (o *GameStateEvent) SetWtime(v int32)`

SetWtime sets Wtime field to given value.


### GetBtime

`func (o *GameStateEvent) GetBtime() int32`

GetBtime returns the Btime field if non-nil, zero value otherwise.

### GetBtimeOk

`func (o *GameStateEvent) GetBtimeOk() (*int32, bool)`

GetBtimeOk returns a tuple with the Btime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtime

`func (o *GameStateEvent) SetBtime(v int32)`

SetBtime sets Btime field to given value.


### GetWinc

`func (o *GameStateEvent) GetWinc() int32`

GetWinc returns the Winc field if non-nil, zero value otherwise.

### GetWincOk

`func (o *GameStateEvent) GetWincOk() (*int32, bool)`

GetWincOk returns a tuple with the Winc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinc

`func (o *GameStateEvent) SetWinc(v int32)`

SetWinc sets Winc field to given value.


### GetBinc

`func (o *GameStateEvent) GetBinc() int32`

GetBinc returns the Binc field if non-nil, zero value otherwise.

### GetBincOk

`func (o *GameStateEvent) GetBincOk() (*int32, bool)`

GetBincOk returns a tuple with the Binc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinc

`func (o *GameStateEvent) SetBinc(v int32)`

SetBinc sets Binc field to given value.


### GetStatus

`func (o *GameStateEvent) GetStatus() GameStatusName`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameStateEvent) GetStatusOk() (*GameStatusName, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameStateEvent) SetStatus(v GameStatusName)`

SetStatus sets Status field to given value.


### GetWinner

`func (o *GameStateEvent) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameStateEvent) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameStateEvent) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameStateEvent) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetWdraw

`func (o *GameStateEvent) GetWdraw() bool`

GetWdraw returns the Wdraw field if non-nil, zero value otherwise.

### GetWdrawOk

`func (o *GameStateEvent) GetWdrawOk() (*bool, bool)`

GetWdrawOk returns a tuple with the Wdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdraw

`func (o *GameStateEvent) SetWdraw(v bool)`

SetWdraw sets Wdraw field to given value.

### HasWdraw

`func (o *GameStateEvent) HasWdraw() bool`

HasWdraw returns a boolean if a field has been set.

### GetBdraw

`func (o *GameStateEvent) GetBdraw() bool`

GetBdraw returns the Bdraw field if non-nil, zero value otherwise.

### GetBdrawOk

`func (o *GameStateEvent) GetBdrawOk() (*bool, bool)`

GetBdrawOk returns a tuple with the Bdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdraw

`func (o *GameStateEvent) SetBdraw(v bool)`

SetBdraw sets Bdraw field to given value.

### HasBdraw

`func (o *GameStateEvent) HasBdraw() bool`

HasBdraw returns a boolean if a field has been set.

### GetWtakeback

`func (o *GameStateEvent) GetWtakeback() bool`

GetWtakeback returns the Wtakeback field if non-nil, zero value otherwise.

### GetWtakebackOk

`func (o *GameStateEvent) GetWtakebackOk() (*bool, bool)`

GetWtakebackOk returns a tuple with the Wtakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtakeback

`func (o *GameStateEvent) SetWtakeback(v bool)`

SetWtakeback sets Wtakeback field to given value.

### HasWtakeback

`func (o *GameStateEvent) HasWtakeback() bool`

HasWtakeback returns a boolean if a field has been set.

### GetBtakeback

`func (o *GameStateEvent) GetBtakeback() bool`

GetBtakeback returns the Btakeback field if non-nil, zero value otherwise.

### GetBtakebackOk

`func (o *GameStateEvent) GetBtakebackOk() (*bool, bool)`

GetBtakebackOk returns a tuple with the Btakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtakeback

`func (o *GameStateEvent) SetBtakeback(v bool)`

SetBtakeback sets Btakeback field to given value.

### HasBtakeback

`func (o *GameStateEvent) HasBtakeback() bool`

HasBtakeback returns a boolean if a field has been set.

### GetExpiration

`func (o *GameStateEvent) GetExpiration() GameStateEventExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *GameStateEvent) GetExpirationOk() (*GameStateEventExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *GameStateEvent) SetExpiration(v GameStateEventExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *GameStateEvent) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


