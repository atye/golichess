# OpeningExplorerMaster200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Opening** | [**NullableOpeningExplorerMaster200ResponseOpening**](OpeningExplorerMaster200ResponseOpening.md) |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Moves** | [**[]OpeningExplorerMaster200ResponseMovesInner**](OpeningExplorerMaster200ResponseMovesInner.md) |  | 
**TopGames** | [**[]OpeningExplorerMaster200ResponseTopGamesInner**](OpeningExplorerMaster200ResponseTopGamesInner.md) |  | 

## Methods

### NewOpeningExplorerMaster200Response

`func NewOpeningExplorerMaster200Response(opening NullableOpeningExplorerMaster200ResponseOpening, white int32, draws int32, black int32, moves []OpeningExplorerMaster200ResponseMovesInner, topGames []OpeningExplorerMaster200ResponseTopGamesInner, ) *OpeningExplorerMaster200Response`

NewOpeningExplorerMaster200Response instantiates a new OpeningExplorerMaster200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMaster200ResponseWithDefaults

`func NewOpeningExplorerMaster200ResponseWithDefaults() *OpeningExplorerMaster200Response`

NewOpeningExplorerMaster200ResponseWithDefaults instantiates a new OpeningExplorerMaster200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOpening

`func (o *OpeningExplorerMaster200Response) GetOpening() OpeningExplorerMaster200ResponseOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerMaster200Response) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerMaster200Response) SetOpening(v OpeningExplorerMaster200ResponseOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerMaster200Response) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerMaster200Response) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil
### GetWhite

`func (o *OpeningExplorerMaster200Response) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMaster200Response) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMaster200Response) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerMaster200Response) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerMaster200Response) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerMaster200Response) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerMaster200Response) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMaster200Response) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMaster200Response) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetMoves

`func (o *OpeningExplorerMaster200Response) GetMoves() []OpeningExplorerMaster200ResponseMovesInner`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *OpeningExplorerMaster200Response) GetMovesOk() (*[]OpeningExplorerMaster200ResponseMovesInner, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *OpeningExplorerMaster200Response) SetMoves(v []OpeningExplorerMaster200ResponseMovesInner)`

SetMoves sets Moves field to given value.


### GetTopGames

`func (o *OpeningExplorerMaster200Response) GetTopGames() []OpeningExplorerMaster200ResponseTopGamesInner`

GetTopGames returns the TopGames field if non-nil, zero value otherwise.

### GetTopGamesOk

`func (o *OpeningExplorerMaster200Response) GetTopGamesOk() (*[]OpeningExplorerMaster200ResponseTopGamesInner, bool)`

GetTopGamesOk returns a tuple with the TopGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopGames

`func (o *OpeningExplorerMaster200Response) SetTopGames(v []OpeningExplorerMaster200ResponseTopGamesInner)`

SetTopGames sets TopGames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


