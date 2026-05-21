# OpeningExplorerMastersGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Winner** | [**NullableGameColor**](GameColor.md) |  | 
**White** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Black** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Year** | **int32** |  | 
**Month** | Pointer to **string** |  | [optional] 

## Methods

### NewOpeningExplorerMastersGame

`func NewOpeningExplorerMastersGame(id string, winner NullableGameColor, white OpeningExplorerGamePlayer, black OpeningExplorerGamePlayer, year int32, ) *OpeningExplorerMastersGame`

NewOpeningExplorerMastersGame instantiates a new OpeningExplorerMastersGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerMastersGameWithDefaults

`func NewOpeningExplorerMastersGameWithDefaults() *OpeningExplorerMastersGame`

NewOpeningExplorerMastersGameWithDefaults instantiates a new OpeningExplorerMastersGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningExplorerMastersGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerMastersGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerMastersGame) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerMastersGame) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerMastersGame) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerMastersGame) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.


### SetWinnerNil

`func (o *OpeningExplorerMastersGame) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *OpeningExplorerMastersGame) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetWhite

`func (o *OpeningExplorerMastersGame) GetWhite() OpeningExplorerGamePlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerMastersGame) GetWhiteOk() (*OpeningExplorerGamePlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerMastersGame) SetWhite(v OpeningExplorerGamePlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerMastersGame) GetBlack() OpeningExplorerGamePlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerMastersGame) GetBlackOk() (*OpeningExplorerGamePlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerMastersGame) SetBlack(v OpeningExplorerGamePlayer)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerMastersGame) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerMastersGame) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerMastersGame) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerMastersGame) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerMastersGame) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerMastersGame) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *OpeningExplorerMastersGame) HasMonth() bool`

HasMonth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


