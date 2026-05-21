# OpeningExplorerPlayerGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Winner** | [**NullableGameColor**](GameColor.md) |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**Mode** | **string** |  | 
**White** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Black** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Year** | **int32** |  | 
**Month** | **string** |  | 

## Methods

### NewOpeningExplorerPlayerGame

`func NewOpeningExplorerPlayerGame(id string, winner NullableGameColor, speed Speed, mode string, white OpeningExplorerGamePlayer, black OpeningExplorerGamePlayer, year int32, month string, ) *OpeningExplorerPlayerGame`

NewOpeningExplorerPlayerGame instantiates a new OpeningExplorerPlayerGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayerGameWithDefaults

`func NewOpeningExplorerPlayerGameWithDefaults() *OpeningExplorerPlayerGame`

NewOpeningExplorerPlayerGameWithDefaults instantiates a new OpeningExplorerPlayerGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningExplorerPlayerGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerPlayerGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerPlayerGame) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerPlayerGame) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerPlayerGame) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerPlayerGame) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.


### SetWinnerNil

`func (o *OpeningExplorerPlayerGame) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *OpeningExplorerPlayerGame) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetSpeed

`func (o *OpeningExplorerPlayerGame) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OpeningExplorerPlayerGame) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OpeningExplorerPlayerGame) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetMode

`func (o *OpeningExplorerPlayerGame) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OpeningExplorerPlayerGame) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OpeningExplorerPlayerGame) SetMode(v string)`

SetMode sets Mode field to given value.


### GetWhite

`func (o *OpeningExplorerPlayerGame) GetWhite() OpeningExplorerGamePlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayerGame) GetWhiteOk() (*OpeningExplorerGamePlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayerGame) SetWhite(v OpeningExplorerGamePlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerPlayerGame) GetBlack() OpeningExplorerGamePlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayerGame) GetBlackOk() (*OpeningExplorerGamePlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayerGame) SetBlack(v OpeningExplorerGamePlayer)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerPlayerGame) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerPlayerGame) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerPlayerGame) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerPlayerGame) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerPlayerGame) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerPlayerGame) SetMonth(v string)`

SetMonth sets Month field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


