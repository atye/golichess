# OpeningExplorerPlayerMovesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**San** | **string** |  | 
**AverageOpponentRating** | **int32** |  | 
**Performance** | **int32** |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Game** | [**NullableOpeningExplorerPlayerGame**](OpeningExplorerPlayerGame.md) |  | 
**Opening** | [**NullableOpeningExplorerOpening**](OpeningExplorerOpening.md) |  | 

## Methods

### NewOpeningExplorerPlayerMovesInner

`func NewOpeningExplorerPlayerMovesInner(uci string, san string, averageOpponentRating int32, performance int32, white int32, draws int32, black int32, game NullableOpeningExplorerPlayerGame, opening NullableOpeningExplorerOpening, ) *OpeningExplorerPlayerMovesInner`

NewOpeningExplorerPlayerMovesInner instantiates a new OpeningExplorerPlayerMovesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayerMovesInnerWithDefaults

`func NewOpeningExplorerPlayerMovesInnerWithDefaults() *OpeningExplorerPlayerMovesInner`

NewOpeningExplorerPlayerMovesInnerWithDefaults instantiates a new OpeningExplorerPlayerMovesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerPlayerMovesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerPlayerMovesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerPlayerMovesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *OpeningExplorerPlayerMovesInner) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *OpeningExplorerPlayerMovesInner) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *OpeningExplorerPlayerMovesInner) SetSan(v string)`

SetSan sets San field to given value.


### GetAverageOpponentRating

`func (o *OpeningExplorerPlayerMovesInner) GetAverageOpponentRating() int32`

GetAverageOpponentRating returns the AverageOpponentRating field if non-nil, zero value otherwise.

### GetAverageOpponentRatingOk

`func (o *OpeningExplorerPlayerMovesInner) GetAverageOpponentRatingOk() (*int32, bool)`

GetAverageOpponentRatingOk returns a tuple with the AverageOpponentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOpponentRating

`func (o *OpeningExplorerPlayerMovesInner) SetAverageOpponentRating(v int32)`

SetAverageOpponentRating sets AverageOpponentRating field to given value.


### GetPerformance

`func (o *OpeningExplorerPlayerMovesInner) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *OpeningExplorerPlayerMovesInner) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *OpeningExplorerPlayerMovesInner) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.


### GetWhite

`func (o *OpeningExplorerPlayerMovesInner) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayerMovesInner) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayerMovesInner) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerPlayerMovesInner) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerPlayerMovesInner) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerPlayerMovesInner) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerPlayerMovesInner) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayerMovesInner) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayerMovesInner) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetGame

`func (o *OpeningExplorerPlayerMovesInner) GetGame() OpeningExplorerPlayerGame`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *OpeningExplorerPlayerMovesInner) GetGameOk() (*OpeningExplorerPlayerGame, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *OpeningExplorerPlayerMovesInner) SetGame(v OpeningExplorerPlayerGame)`

SetGame sets Game field to given value.


### SetGameNil

`func (o *OpeningExplorerPlayerMovesInner) SetGameNil(b bool)`

 SetGameNil sets the value for Game to be an explicit nil

### UnsetGame
`func (o *OpeningExplorerPlayerMovesInner) UnsetGame()`

UnsetGame ensures that no value is present for Game, not even an explicit nil
### GetOpening

`func (o *OpeningExplorerPlayerMovesInner) GetOpening() OpeningExplorerOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerPlayerMovesInner) GetOpeningOk() (*OpeningExplorerOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerPlayerMovesInner) SetOpening(v OpeningExplorerOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerPlayerMovesInner) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerPlayerMovesInner) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


