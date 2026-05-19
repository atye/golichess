# BoardGameStream200ResponseOneOfState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
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

## Methods

### NewBoardGameStream200ResponseOneOfState

`func NewBoardGameStream200ResponseOneOfState(type_ string, moves string, wtime int32, btime int32, winc int32, binc int32, status string, ) *BoardGameStream200ResponseOneOfState`

NewBoardGameStream200ResponseOneOfState instantiates a new BoardGameStream200ResponseOneOfState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardGameStream200ResponseOneOfStateWithDefaults

`func NewBoardGameStream200ResponseOneOfStateWithDefaults() *BoardGameStream200ResponseOneOfState`

NewBoardGameStream200ResponseOneOfStateWithDefaults instantiates a new BoardGameStream200ResponseOneOfState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BoardGameStream200ResponseOneOfState) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardGameStream200ResponseOneOfState) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardGameStream200ResponseOneOfState) SetType(v string)`

SetType sets Type field to given value.


### GetMoves

`func (o *BoardGameStream200ResponseOneOfState) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *BoardGameStream200ResponseOneOfState) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *BoardGameStream200ResponseOneOfState) SetMoves(v string)`

SetMoves sets Moves field to given value.


### GetWtime

`func (o *BoardGameStream200ResponseOneOfState) GetWtime() int32`

GetWtime returns the Wtime field if non-nil, zero value otherwise.

### GetWtimeOk

`func (o *BoardGameStream200ResponseOneOfState) GetWtimeOk() (*int32, bool)`

GetWtimeOk returns a tuple with the Wtime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtime

`func (o *BoardGameStream200ResponseOneOfState) SetWtime(v int32)`

SetWtime sets Wtime field to given value.


### GetBtime

`func (o *BoardGameStream200ResponseOneOfState) GetBtime() int32`

GetBtime returns the Btime field if non-nil, zero value otherwise.

### GetBtimeOk

`func (o *BoardGameStream200ResponseOneOfState) GetBtimeOk() (*int32, bool)`

GetBtimeOk returns a tuple with the Btime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtime

`func (o *BoardGameStream200ResponseOneOfState) SetBtime(v int32)`

SetBtime sets Btime field to given value.


### GetWinc

`func (o *BoardGameStream200ResponseOneOfState) GetWinc() int32`

GetWinc returns the Winc field if non-nil, zero value otherwise.

### GetWincOk

`func (o *BoardGameStream200ResponseOneOfState) GetWincOk() (*int32, bool)`

GetWincOk returns a tuple with the Winc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinc

`func (o *BoardGameStream200ResponseOneOfState) SetWinc(v int32)`

SetWinc sets Winc field to given value.


### GetBinc

`func (o *BoardGameStream200ResponseOneOfState) GetBinc() int32`

GetBinc returns the Binc field if non-nil, zero value otherwise.

### GetBincOk

`func (o *BoardGameStream200ResponseOneOfState) GetBincOk() (*int32, bool)`

GetBincOk returns a tuple with the Binc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBinc

`func (o *BoardGameStream200ResponseOneOfState) SetBinc(v int32)`

SetBinc sets Binc field to given value.


### GetStatus

`func (o *BoardGameStream200ResponseOneOfState) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BoardGameStream200ResponseOneOfState) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BoardGameStream200ResponseOneOfState) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetWinner

`func (o *BoardGameStream200ResponseOneOfState) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *BoardGameStream200ResponseOneOfState) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *BoardGameStream200ResponseOneOfState) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *BoardGameStream200ResponseOneOfState) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### GetWdraw

`func (o *BoardGameStream200ResponseOneOfState) GetWdraw() bool`

GetWdraw returns the Wdraw field if non-nil, zero value otherwise.

### GetWdrawOk

`func (o *BoardGameStream200ResponseOneOfState) GetWdrawOk() (*bool, bool)`

GetWdrawOk returns a tuple with the Wdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWdraw

`func (o *BoardGameStream200ResponseOneOfState) SetWdraw(v bool)`

SetWdraw sets Wdraw field to given value.

### HasWdraw

`func (o *BoardGameStream200ResponseOneOfState) HasWdraw() bool`

HasWdraw returns a boolean if a field has been set.

### GetBdraw

`func (o *BoardGameStream200ResponseOneOfState) GetBdraw() bool`

GetBdraw returns the Bdraw field if non-nil, zero value otherwise.

### GetBdrawOk

`func (o *BoardGameStream200ResponseOneOfState) GetBdrawOk() (*bool, bool)`

GetBdrawOk returns a tuple with the Bdraw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBdraw

`func (o *BoardGameStream200ResponseOneOfState) SetBdraw(v bool)`

SetBdraw sets Bdraw field to given value.

### HasBdraw

`func (o *BoardGameStream200ResponseOneOfState) HasBdraw() bool`

HasBdraw returns a boolean if a field has been set.

### GetWtakeback

`func (o *BoardGameStream200ResponseOneOfState) GetWtakeback() bool`

GetWtakeback returns the Wtakeback field if non-nil, zero value otherwise.

### GetWtakebackOk

`func (o *BoardGameStream200ResponseOneOfState) GetWtakebackOk() (*bool, bool)`

GetWtakebackOk returns a tuple with the Wtakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWtakeback

`func (o *BoardGameStream200ResponseOneOfState) SetWtakeback(v bool)`

SetWtakeback sets Wtakeback field to given value.

### HasWtakeback

`func (o *BoardGameStream200ResponseOneOfState) HasWtakeback() bool`

HasWtakeback returns a boolean if a field has been set.

### GetBtakeback

`func (o *BoardGameStream200ResponseOneOfState) GetBtakeback() bool`

GetBtakeback returns the Btakeback field if non-nil, zero value otherwise.

### GetBtakebackOk

`func (o *BoardGameStream200ResponseOneOfState) GetBtakebackOk() (*bool, bool)`

GetBtakebackOk returns a tuple with the Btakeback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBtakeback

`func (o *BoardGameStream200ResponseOneOfState) SetBtakeback(v bool)`

SetBtakeback sets Btakeback field to given value.

### HasBtakeback

`func (o *BoardGameStream200ResponseOneOfState) HasBtakeback() bool`

HasBtakeback returns a boolean if a field has been set.

### GetExpiration

`func (o *BoardGameStream200ResponseOneOfState) GetExpiration() BoardGameStream200ResponseOneOfStateExpiration`

GetExpiration returns the Expiration field if non-nil, zero value otherwise.

### GetExpirationOk

`func (o *BoardGameStream200ResponseOneOfState) GetExpirationOk() (*BoardGameStream200ResponseOneOfStateExpiration, bool)`

GetExpirationOk returns a tuple with the Expiration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiration

`func (o *BoardGameStream200ResponseOneOfState) SetExpiration(v BoardGameStream200ResponseOneOfStateExpiration)`

SetExpiration sets Expiration field to given value.

### HasExpiration

`func (o *BoardGameStream200ResponseOneOfState) HasExpiration() bool`

HasExpiration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


