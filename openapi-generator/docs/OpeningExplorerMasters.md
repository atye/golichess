# OpeningExplorerMasters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerOpening**](OpeningExplorerOpening.md) |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerMastersMovesInner**](OpeningExplorerMastersMovesInner.md) |  | 
**TopGames** | [**[]OpeningExplorerMastersTopGamesInner**](OpeningExplorerMastersTopGamesInner.md) |  | 

## Methods

### NewOpeningExplorerMasters

`func NewOpeningExplorerMasters(opening NullableOpeningExplorerOpening, white int32, draws int32, black int32, moves []OpeningExplorerMastersMovesInner, topGames []OpeningExplorerMastersTopGamesInner, ) *OpeningExplorerMasters`

NewOpeningExplorerMasters instantiates a new OpeningExplorerMasters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMastersWithDefaults

`func NewOpeningExplorerMastersWithDefaults() *OpeningExplorerMasters`

NewOpeningExplorerMastersWithDefaults instantiates a new OpeningExplorerMasters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerMasters) GetOpening() OpeningExplorerOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerMasters) GetOpeningOk() (*OpeningExplorerOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerMasters) SetOpening(v OpeningExplorerOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerMasters) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerMasters) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetWhite

`func (o *OpeningExplorerMasters) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMasters) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMasters) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerMasters) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerMasters) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerMasters) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerMasters) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMasters) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMasters) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerMasters) GetMoves() []OpeningExplorerMastersMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerMasters) GetMovesOk() (*[]OpeningExplorerMastersMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerMasters) SetMoves(v []OpeningExplorerMastersMovesInner)`

SetMoves sets Moves field to given value.


### GetTopGames

`func (o *OpeningExplorerMasters) GetTopGames() []OpeningExplorerMastersTopGamesInner`

GetTopGames returns the TopGames field if non-nil, zero value otherwise.

### GetTopGamesOk

`func (o *OpeningExplorerMasters) GetTopGamesOk() (*[]OpeningExplorerMastersTopGamesInner, bool)`

GetTopGamesOk returns a tuple with the TopGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopGames

`func (o *OpeningExplorerMasters) SetTopGames(v []OpeningExplorerMastersTopGamesInner)`

SetTopGames sets TopGames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


