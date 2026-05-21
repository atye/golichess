# OpeningExplorerMastersMovesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**San** | **string** |  | 
**AverageRating** | **int32** |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Game** | [**NullableOpeningExplorerMastersGame**](OpeningExplorerMastersGame.md) |  | 
**Opening** | [**NullableOpeningExplorerOpening**](OpeningExplorerOpening.md) |  | 

## Methods

### NewOpeningExplorerMastersMovesInner

`func NewOpeningExplorerMastersMovesInner(uci string, san string, averageRating int32, white int32, draws int32, black int32, game NullableOpeningExplorerMastersGame, opening NullableOpeningExplorerOpening, ) *OpeningExplorerMastersMovesInner`

NewOpeningExplorerMastersMovesInner instantiates a new OpeningExplorerMastersMovesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMastersMovesInnerWithDefaults

`func NewOpeningExplorerMastersMovesInnerWithDefaults() *OpeningExplorerMastersMovesInner`

NewOpeningExplorerMastersMovesInnerWithDefaults instantiates a new OpeningExplorerMastersMovesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerMastersMovesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerMastersMovesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerMastersMovesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *OpeningExplorerMastersMovesInner) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *OpeningExplorerMastersMovesInner) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *OpeningExplorerMastersMovesInner) SetSan(v string)`

SetSan sets San field to given value.


### GetAverageRating

`func (o *OpeningExplorerMastersMovesInner) GetAverageRating() int32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *OpeningExplorerMastersMovesInner) GetAverageRatingOk() (*int32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *OpeningExplorerMastersMovesInner) SetAverageRating(v int32)`

SetAverageRating sets AverageRating field to given value.


### GetWhite

`func (o *OpeningExplorerMastersMovesInner) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMastersMovesInner) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMastersMovesInner) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerMastersMovesInner) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerMastersMovesInner) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerMastersMovesInner) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerMastersMovesInner) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMastersMovesInner) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMastersMovesInner) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetGame

`func (o *OpeningExplorerMastersMovesInner) GetGame() OpeningExplorerMastersGame`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *OpeningExplorerMastersMovesInner) GetGameOk() (*OpeningExplorerMastersGame, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *OpeningExplorerMastersMovesInner) SetGame(v OpeningExplorerMastersGame)`

SetGame sets Game field to given value.


### SetGameNil

`func (o *OpeningExplorerMastersMovesInner) SetGameNil(b bool)`

 SetGameNil sets the value for Game to be an explicit nil

### UnsetGame
`func (o *OpeningExplorerMastersMovesInner) UnsetGame()`

UnsetGame ensures that no value is present for Game, not even an explicit nil
### GetOpening

`func (o *OpeningExplorerMastersMovesInner) GetOpening() OpeningExplorerOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerMastersMovesInner) GetOpeningOk() (*OpeningExplorerOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerMastersMovesInner) SetOpening(v OpeningExplorerOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerMastersMovesInner) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerMastersMovesInner) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


