# OpeningExplorerMastersTopGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Winner** | [**GameColor**](GameColor.md) |  | 
**White** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Black** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Year** | **int32** |  | 
**Month** | Pointer to **string** |  | [optional] 
**Uci** | **string** |  | 

## Methods

### NewOpeningExplorerMastersTopGamesInner

`func NewOpeningExplorerMastersTopGamesInner(id string, winner GameColor, white OpeningExplorerGamePlayer, black OpeningExplorerGamePlayer, year int32, uci string, ) *OpeningExplorerMastersTopGamesInner`

NewOpeningExplorerMastersTopGamesInner instantiates a new OpeningExplorerMastersTopGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMastersTopGamesInnerWithDefaults

`func NewOpeningExplorerMastersTopGamesInnerWithDefaults() *OpeningExplorerMastersTopGamesInner`

NewOpeningExplorerMastersTopGamesInnerWithDefaults instantiates a new OpeningExplorerMastersTopGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningExplorerMastersTopGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerMastersTopGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerMastersTopGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerMastersTopGamesInner) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerMastersTopGamesInner) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerMastersTopGamesInner) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.


### GetWhite

`func (o *OpeningExplorerMastersTopGamesInner) GetWhite() OpeningExplorerGamePlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMastersTopGamesInner) GetWhiteOk() (*OpeningExplorerGamePlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMastersTopGamesInner) SetWhite(v OpeningExplorerGamePlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerMastersTopGamesInner) GetBlack() OpeningExplorerGamePlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMastersTopGamesInner) GetBlackOk() (*OpeningExplorerGamePlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMastersTopGamesInner) SetBlack(v OpeningExplorerGamePlayer)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerMastersTopGamesInner) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerMastersTopGamesInner) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerMastersTopGamesInner) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerMastersTopGamesInner) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerMastersTopGamesInner) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerMastersTopGamesInner) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *OpeningExplorerMastersTopGamesInner) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetUci

`func (o *OpeningExplorerMastersTopGamesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerMastersTopGamesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerMastersTopGamesInner) SetUci(v string)`

SetUci sets Uci field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


