# OpeningExplorerPlayer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerMaster200ResponseOpening**](OpeningExplorerMaster200ResponseOpening.md) |  | 
**QueuePosition** | **int32** | Waiting for other players to be indexed first | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerPlayer200ResponseMovesInner**](OpeningExplorerPlayer200ResponseMovesInner.md) |  | 
**RecentGames** | [**[]OpeningExplorerPlayer200ResponseRecentGamesInner**](OpeningExplorerPlayer200ResponseRecentGamesInner.md) |  | 

## Methods

### NewOpeningExplorerPlayer200Response

`func NewOpeningExplorerPlayer200Response(opening NullableOpeningExplorerMaster200ResponseOpening, queuePosition int32, white int32, draws int32, black int32, moves []OpeningExplorerPlayer200ResponseMovesInner, recentGames []OpeningExplorerPlayer200ResponseRecentGamesInner, ) *OpeningExplorerPlayer200Response`

NewOpeningExplorerPlayer200Response instantiates a new OpeningExplorerPlayer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayer200ResponseWithDefaults

`func NewOpeningExplorerPlayer200ResponseWithDefaults() *OpeningExplorerPlayer200Response`

NewOpeningExplorerPlayer200ResponseWithDefaults instantiates a new OpeningExplorerPlayer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerPlayer200Response) GetOpening() OpeningExplorerMaster200ResponseOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerPlayer200Response) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerPlayer200Response) SetOpening(v OpeningExplorerMaster200ResponseOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerPlayer200Response) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerPlayer200Response) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetQueuePosition

`func (o *OpeningExplorerPlayer200Response) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *OpeningExplorerPlayer200Response) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *OpeningExplorerPlayer200Response) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.


### GetWhite

`func (o *OpeningExplorerPlayer200Response) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayer200Response) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayer200Response) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerPlayer200Response) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerPlayer200Response) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerPlayer200Response) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerPlayer200Response) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayer200Response) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayer200Response) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerPlayer200Response) GetMoves() []OpeningExplorerPlayer200ResponseMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerPlayer200Response) GetMovesOk() (*[]OpeningExplorerPlayer200ResponseMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerPlayer200Response) SetMoves(v []OpeningExplorerPlayer200ResponseMovesInner)`

SetMoves sets Moves field to given value.


### GetRecentGames

`func (o *OpeningExplorerPlayer200Response) GetRecentGames() []OpeningExplorerPlayer200ResponseRecentGamesInner`

GetRecentGames returns the RecentGames field if non-nil, zero value otherwise.

### GetRecentGamesOk

`func (o *OpeningExplorerPlayer200Response) GetRecentGamesOk() (*[]OpeningExplorerPlayer200ResponseRecentGamesInner, bool)`

GetRecentGamesOk returns a tuple with the RecentGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentGames

`func (o *OpeningExplorerPlayer200Response) SetRecentGames(v []OpeningExplorerPlayer200ResponseRecentGamesInner)`

SetRecentGames sets RecentGames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


