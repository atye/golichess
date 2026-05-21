# OpeningExplorerLichess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerOpening**](OpeningExplorerOpening.md) |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerLichessMovesInner**](OpeningExplorerLichessMovesInner.md) |  | 
**TopGames** | [**[]OpeningExplorerLichessTopGamesInner**](OpeningExplorerLichessTopGamesInner.md) |  | 
**RecentGames** | Pointer to [**[]OpeningExplorerLichessTopGamesInner**](OpeningExplorerLichessTopGamesInner.md) |  | [optional] 
**History** | Pointer to [**[]OpeningExplorerLichessHistoryInner**](OpeningExplorerLichessHistoryInner.md) |  | [optional] 

## Methods

### NewOpeningExplorerLichess

`func NewOpeningExplorerLichess(opening NullableOpeningExplorerOpening, white int32, draws int32, black int32, moves []OpeningExplorerLichessMovesInner, topGames []OpeningExplorerLichessTopGamesInner, ) *OpeningExplorerLichess`

NewOpeningExplorerLichess instantiates a new OpeningExplorerLichess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerLichessWithDefaults

`func NewOpeningExplorerLichessWithDefaults() *OpeningExplorerLichess`

NewOpeningExplorerLichessWithDefaults instantiates a new OpeningExplorerLichess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerLichess) GetOpening() OpeningExplorerOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerLichess) GetOpeningOk() (*OpeningExplorerOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerLichess) SetOpening(v OpeningExplorerOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerLichess) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerLichess) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetWhite

`func (o *OpeningExplorerLichess) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerLichess) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerLichess) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerLichess) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerLichess) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerLichess) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerLichess) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerLichess) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerLichess) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerLichess) GetMoves() []OpeningExplorerLichessMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerLichess) GetMovesOk() (*[]OpeningExplorerLichessMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerLichess) SetMoves(v []OpeningExplorerLichessMovesInner)`

SetMoves sets Moves field to given value.


### GetTopGames

`func (o *OpeningExplorerLichess) GetTopGames() []OpeningExplorerLichessTopGamesInner`

GetTopGames returns the TopGames field if non-nil, zero value otherwise.

### GetTopGamesOk

`func (o *OpeningExplorerLichess) GetTopGamesOk() (*[]OpeningExplorerLichessTopGamesInner, bool)`

GetTopGamesOk returns a tuple with the TopGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopGames

`func (o *OpeningExplorerLichess) SetTopGames(v []OpeningExplorerLichessTopGamesInner)`

SetTopGames sets TopGames field to given value.


### GetRecentGames

`func (o *OpeningExplorerLichess) GetRecentGames() []OpeningExplorerLichessTopGamesInner`

GetRecentGames returns the RecentGames field if non-nil, zero value otherwise.

### GetRecentGamesOk

`func (o *OpeningExplorerLichess) GetRecentGamesOk() (*[]OpeningExplorerLichessTopGamesInner, bool)`

GetRecentGamesOk returns a tuple with the RecentGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentGames

`func (o *OpeningExplorerLichess) SetRecentGames(v []OpeningExplorerLichessTopGamesInner)`

SetRecentGames sets RecentGames field to given value.

### HasRecentGames

`func (o *OpeningExplorerLichess) HasRecentGames() bool`

HasRecentGames returns a boolean if a field has been set.

### GetHistory

`func (o *OpeningExplorerLichess) GetHistory() []OpeningExplorerLichessHistoryInner`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *OpeningExplorerLichess) GetHistoryOk() (*[]OpeningExplorerLichessHistoryInner, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *OpeningExplorerLichess) SetHistory(v []OpeningExplorerLichessHistoryInner)`

SetHistory sets History field to given value.

### HasHistory

`func (o *OpeningExplorerLichess) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


