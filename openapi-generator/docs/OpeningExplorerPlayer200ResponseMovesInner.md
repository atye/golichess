# OpeningExplorerPlayer200ResponseMovesInner

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
**Game** | [**NullableOpeningExplorerPlayer200ResponseMovesInnerGame**](OpeningExplorerPlayer200ResponseMovesInnerGame.md) |  | 
**Opening** | [**NullableOpeningExplorerMaster200ResponseOpening**](OpeningExplorerMaster200ResponseOpening.md) |  | 

## Methods

### NewOpeningExplorerPlayer200ResponseMovesInner

`func NewOpeningExplorerPlayer200ResponseMovesInner(uci string, san string, averageOpponentRating int32, performance int32, white int32, draws int32, black int32, game NullableOpeningExplorerPlayer200ResponseMovesInnerGame, opening NullableOpeningExplorerMaster200ResponseOpening, ) *OpeningExplorerPlayer200ResponseMovesInner`

NewOpeningExplorerPlayer200ResponseMovesInner instantiates a new OpeningExplorerPlayer200ResponseMovesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayer200ResponseMovesInnerWithDefaults

`func NewOpeningExplorerPlayer200ResponseMovesInnerWithDefaults() *OpeningExplorerPlayer200ResponseMovesInner`

NewOpeningExplorerPlayer200ResponseMovesInnerWithDefaults instantiates a new OpeningExplorerPlayer200ResponseMovesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetSan(v string)`

SetSan sets San field to given value.


### GetAverageOpponentRating

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetAverageOpponentRating() int32`

GetAverageOpponentRating returns the AverageOpponentRating field if non-nil, zero value otherwise.

### GetAverageOpponentRatingOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetAverageOpponentRatingOk() (*int32, bool)`

GetAverageOpponentRatingOk returns a tuple with the AverageOpponentRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageOpponentRating

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetAverageOpponentRating(v int32)`

SetAverageOpponentRating sets AverageOpponentRating field to given value.


### GetPerformance

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.


### GetWhite

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetGame

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetGame() OpeningExplorerPlayer200ResponseMovesInnerGame`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetGameOk() (*OpeningExplorerPlayer200ResponseMovesInnerGame, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetGame(v OpeningExplorerPlayer200ResponseMovesInnerGame)`

SetGame sets Game field to given value.


### SetGameNil

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetGameNil(b bool)`

 SetGameNil sets the value for Game to be an explicit nil

### UnsetGame
`func (o *OpeningExplorerPlayer200ResponseMovesInner) UnsetGame()`

UnsetGame ensures that no value is present for Game, not even an explicit nil
### GetOpening

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetOpening() OpeningExplorerMaster200ResponseOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerPlayer200ResponseMovesInner) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetOpening(v OpeningExplorerMaster200ResponseOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerPlayer200ResponseMovesInner) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerPlayer200ResponseMovesInner) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


