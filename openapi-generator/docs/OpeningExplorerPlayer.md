# OpeningExplorerPlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerOpening**](OpeningExplorerOpening.md) |  | 
**QueuePosition** | **int32** | Waiting for other players to be indexed first | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerPlayerMovesInner**](OpeningExplorerPlayerMovesInner.md) |  | 
**RecentGames** | [**[]OpeningExplorerPlayerRecentGamesInner**](OpeningExplorerPlayerRecentGamesInner.md) |  | 

## Methods

### NewOpeningExplorerPlayer

`func NewOpeningExplorerPlayer(opening NullableOpeningExplorerOpening, queuePosition int32, white int32, draws int32, black int32, moves []OpeningExplorerPlayerMovesInner, recentGames []OpeningExplorerPlayerRecentGamesInner, ) *OpeningExplorerPlayer`

NewOpeningExplorerPlayer instantiates a new OpeningExplorerPlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayerWithDefaults

`func NewOpeningExplorerPlayerWithDefaults() *OpeningExplorerPlayer`

NewOpeningExplorerPlayerWithDefaults instantiates a new OpeningExplorerPlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerPlayer) GetOpening() OpeningExplorerOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerPlayer) GetOpeningOk() (*OpeningExplorerOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerPlayer) SetOpening(v OpeningExplorerOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerPlayer) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerPlayer) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetQueuePosition

`func (o *OpeningExplorerPlayer) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *OpeningExplorerPlayer) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *OpeningExplorerPlayer) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.


### GetWhite

`func (o *OpeningExplorerPlayer) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayer) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayer) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerPlayer) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerPlayer) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerPlayer) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerPlayer) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayer) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayer) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerPlayer) GetMoves() []OpeningExplorerPlayerMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerPlayer) GetMovesOk() (*[]OpeningExplorerPlayerMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerPlayer) SetMoves(v []OpeningExplorerPlayerMovesInner)`

SetMoves sets Moves field to given value.


### GetRecentGames

`func (o *OpeningExplorerPlayer) GetRecentGames() []OpeningExplorerPlayerRecentGamesInner`

GetRecentGames returns the RecentGames field if non-nil, zero value otherwise.

### GetRecentGamesOk

`func (o *OpeningExplorerPlayer) GetRecentGamesOk() (*[]OpeningExplorerPlayerRecentGamesInner, bool)`

GetRecentGamesOk returns a tuple with the RecentGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentGames

`func (o *OpeningExplorerPlayer) SetRecentGames(v []OpeningExplorerPlayerRecentGamesInner)`

SetRecentGames sets RecentGames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


