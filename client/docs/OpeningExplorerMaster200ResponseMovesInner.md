# OpeningExplorerMaster200ResponseMovesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**San** | **string** |  | 
**AverageRating** | **int32** |  | 
**White** | **int32** |  | 
**Draws** | **int32** |  | 
**Black** | **int32** |  | 
**Game** | [**NullableOpeningExplorerMaster200ResponseMovesInnerGame**](OpeningExplorerMaster200ResponseMovesInnerGame.md) |  | 
**Opening** | [**NullableOpeningExplorerMaster200ResponseOpening**](OpeningExplorerMaster200ResponseOpening.md) |  | 

## Methods

### NewOpeningExplorerMaster200ResponseMovesInner

`func NewOpeningExplorerMaster200ResponseMovesInner(uci string, san string, averageRating int32, white int32, draws int32, black int32, game NullableOpeningExplorerMaster200ResponseMovesInnerGame, opening NullableOpeningExplorerMaster200ResponseOpening, ) *OpeningExplorerMaster200ResponseMovesInner`

NewOpeningExplorerMaster200ResponseMovesInner instantiates a new OpeningExplorerMaster200ResponseMovesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMaster200ResponseMovesInnerWithDefaults

`func NewOpeningExplorerMaster200ResponseMovesInnerWithDefaults() *OpeningExplorerMaster200ResponseMovesInner`

NewOpeningExplorerMaster200ResponseMovesInnerWithDefaults instantiates a new OpeningExplorerMaster200ResponseMovesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetSan

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetSan() string`

GetSan returns the San field if non-nil, zero value otherwise.

### GetSanOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetSanOk() (*string, bool)`

GetSanOk returns a tuple with the San field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSan

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetSan(v string)`

SetSan sets San field to given value.


### GetAverageRating

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetAverageRating() int32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetAverageRatingOk() (*int32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetAverageRating(v int32)`

SetAverageRating sets AverageRating field to given value.


### GetWhite

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetWhite() int32`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetWhiteOk() (*int32, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetWhite(v int32)`

SetWhite sets White field to given value.


### GetDraws

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetDraws() int32`

GetDraws returns the Draws field if non-nil, zero value otherwise.

### GetDrawsOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetDrawsOk() (*int32, bool)`

GetDrawsOk returns a tuple with the Draws field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraws

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetDraws(v int32)`

SetDraws sets Draws field to given value.


### GetBlack

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetBlack() int32`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetBlackOk() (*int32, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetBlack(v int32)`

SetBlack sets Black field to given value.


### GetGame

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetGame() OpeningExplorerMaster200ResponseMovesInnerGame`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetGameOk() (*OpeningExplorerMaster200ResponseMovesInnerGame, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetGame(v OpeningExplorerMaster200ResponseMovesInnerGame)`

SetGame sets Game field to given value.


### SetGameNil

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetGameNil(b bool)`

 SetGameNil sets the value for Game to be an explicit nil

### UnsetGame
`func (o *OpeningExplorerMaster200ResponseMovesInner) UnsetGame()`

UnsetGame ensures that no value is present for Game, not even an explicit nil
### GetOpening

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetOpening() OpeningExplorerMaster200ResponseOpening`

GetOpening returns the Opening field if non-nil, zero value otherwise.

### GetOpeningOk

`func (o *OpeningExplorerMaster200ResponseMovesInner) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool)`

GetOpeningOk returns a tuple with the Opening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpening

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetOpening(v OpeningExplorerMaster200ResponseOpening)`

SetOpening sets Opening field to given value.


### SetOpeningNil

`func (o *OpeningExplorerMaster200ResponseMovesInner) SetOpeningNil(b bool)`

 SetOpeningNil sets the value for Opening to be an explicit nil

### UnsetOpening
`func (o *OpeningExplorerMaster200ResponseMovesInner) UnsetOpening()`

UnsetOpening ensures that no value is present for Opening, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


