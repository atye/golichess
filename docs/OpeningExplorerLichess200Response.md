# OpeningExplorerLichess200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerMaster200ResponseOpening**](OpeningExplorerMaster200ResponseOpening.md) |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerLichess200ResponseMovesInner**](OpeningExplorerLichess200ResponseMovesInner.md) |  | 
**TopGames** | [**[]OpeningExplorerLichess200ResponseTopGamesInner**](OpeningExplorerLichess200ResponseTopGamesInner.md) |  | 
**RecentGames** | Pointer to [**[]OpeningExplorerLichess200ResponseTopGamesInner**](OpeningExplorerLichess200ResponseTopGamesInner.md) |  | [optional] 
**History** | Pointer to [**[]OpeningExplorerLichess200ResponseHistoryInner**](OpeningExplorerLichess200ResponseHistoryInner.md) |  | [optional] 

## Methods

### NewOpeningExplorerLichess200Response

`func NewOpeningExplorerLichess200Response(opening NullableOpeningExplorerMaster200ResponseOpening, white int32, draws int32, black int32, moves []OpeningExplorerLichess200ResponseMovesInner, topGames []OpeningExplorerLichess200ResponseTopGamesInner, ) *OpeningExplorerLichess200Response`

NewOpeningExplorerLichess200Response instantiates a new OpeningExplorerLichess200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerLichess200ResponseWithDefaults

`func NewOpeningExplorerLichess200ResponseWithDefaults() *OpeningExplorerLichess200Response`

NewOpeningExplorerLichess200ResponseWithDefaults instantiates a new OpeningExplorerLichess200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerLichess200Response) GetOpening() OpeningExplorerMaster200ResponseOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerLichess200Response) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerLichess200Response) SetOpening(v OpeningExplorerMaster200ResponseOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerLichess200Response) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerLichess200Response) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetWhite

`func (o *OpeningExplorerLichess200Response) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerLichess200Response) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerLichess200Response) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerLichess200Response) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerLichess200Response) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerLichess200Response) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerLichess200Response) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerLichess200Response) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerLichess200Response) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerLichess200Response) GetMoves() []OpeningExplorerLichess200ResponseMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerLichess200Response) GetMovesOk() (*[]OpeningExplorerLichess200ResponseMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerLichess200Response) SetMoves(v []OpeningExplorerLichess200ResponseMovesInner)`

SetMoves sets Moves field to given value.


### GetTopGames

`func (o *OpeningExplorerLichess200Response) GetTopGames() []OpeningExplorerLichess200ResponseTopGamesInner`

GetTopGames returns the TopGames field if non-nil, zero value otherwise.

### GetTopGamesOk

`func (o *OpeningExplorerLichess200Response) GetTopGamesOk() (*[]OpeningExplorerLichess200ResponseTopGamesInner, bool)`

GetTopGamesOk returns a tuple with the TopGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopGames

`func (o *OpeningExplorerLichess200Response) SetTopGames(v []OpeningExplorerLichess200ResponseTopGamesInner)`

SetTopGames sets TopGames field to given value.


### GetRecentGames

`func (o *OpeningExplorerLichess200Response) GetRecentGames() []OpeningExplorerLichess200ResponseTopGamesInner`

GetRecentGames returns the RecentGames field if non-nil, zero value otherwise.

### GetRecentGamesOk

`func (o *OpeningExplorerLichess200Response) GetRecentGamesOk() (*[]OpeningExplorerLichess200ResponseTopGamesInner, bool)`

GetRecentGamesOk returns a tuple with the RecentGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentGames

`func (o *OpeningExplorerLichess200Response) SetRecentGames(v []OpeningExplorerLichess200ResponseTopGamesInner)`

SetRecentGames sets RecentGames field to given value.

### HasRecentGames

`func (o *OpeningExplorerLichess200Response) HasRecentGames() bool`

HasRecentGames returns a boolean if a field has been set.

### GetHistory

`func (o *OpeningExplorerLichess200Response) GetHistory() []OpeningExplorerLichess200ResponseHistoryInner`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *OpeningExplorerLichess200Response) GetHistoryOk() (*[]OpeningExplorerLichess200ResponseHistoryInner, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *OpeningExplorerLichess200Response) SetHistory(v []OpeningExplorerLichess200ResponseHistoryInner)`

SetHistory sets History field to given value.

### HasHistory

`func (o *OpeningExplorerLichess200Response) HasHistory() bool`

HasHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


