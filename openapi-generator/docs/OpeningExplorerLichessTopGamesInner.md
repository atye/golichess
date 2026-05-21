# OpeningExplorerLichessTopGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Winner** | [**GameColor**](GameColor.md) |  | 
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**White** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Black** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Year** | **float32** |  | 
**Month** | **string** |  | 
**Uci** | **string** |  | 

## Methods

### NewOpeningExplorerLichessTopGamesInner

`func NewOpeningExplorerLichessTopGamesInner(id string, winner GameColor, white OpeningExplorerGamePlayer, black OpeningExplorerGamePlayer, year float32, month string, uci string, ) *OpeningExplorerLichessTopGamesInner`

NewOpeningExplorerLichessTopGamesInner instantiates a new OpeningExplorerLichessTopGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerLichessTopGamesInnerWithDefaults

`func NewOpeningExplorerLichessTopGamesInnerWithDefaults() *OpeningExplorerLichessTopGamesInner`

NewOpeningExplorerLichessTopGamesInnerWithDefaults instantiates a new OpeningExplorerLichessTopGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningExplorerLichessTopGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerLichessTopGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerLichessTopGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerLichessTopGamesInner) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerLichessTopGamesInner) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerLichessTopGamesInner) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.


### GetSpeed

`func (o *OpeningExplorerLichessTopGamesInner) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OpeningExplorerLichessTopGamesInner) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OpeningExplorerLichessTopGamesInner) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OpeningExplorerLichessTopGamesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetWhite

`func (o *OpeningExplorerLichessTopGamesInner) GetWhite() OpeningExplorerGamePlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerLichessTopGamesInner) GetWhiteOk() (*OpeningExplorerGamePlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerLichessTopGamesInner) SetWhite(v OpeningExplorerGamePlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerLichessTopGamesInner) GetBlack() OpeningExplorerGamePlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerLichessTopGamesInner) GetBlackOk() (*OpeningExplorerGamePlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerLichessTopGamesInner) SetBlack(v OpeningExplorerGamePlayer)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerLichessTopGamesInner) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerLichessTopGamesInner) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerLichessTopGamesInner) SetYear(v float32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerLichessTopGamesInner) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerLichessTopGamesInner) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerLichessTopGamesInner) SetMonth(v string)`

SetMonth sets Month field to given value.


### GetUci

`func (o *OpeningExplorerLichessTopGamesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerLichessTopGamesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerLichessTopGamesInner) SetUci(v string)`

SetUci sets Uci field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


