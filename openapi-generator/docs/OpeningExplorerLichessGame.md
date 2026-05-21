# OpeningExplorerLichessGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Winner** | [**NullableGameColor**](GameColor.md) |  | 
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**White** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Black** | [**OpeningExplorerGamePlayer**](OpeningExplorerGamePlayer.md) |  | 
**Year** | **float32** |  | 
**Month** | **NullableString** |  | 

## Methods

### NewOpeningExplorerLichessGame

`func NewOpeningExplorerLichessGame(id string, winner NullableGameColor, white OpeningExplorerGamePlayer, black OpeningExplorerGamePlayer, year float32, month NullableString, ) *OpeningExplorerLichessGame`

NewOpeningExplorerLichessGame instantiates a new OpeningExplorerLichessGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerLichessGameWithDefaults

`func NewOpeningExplorerLichessGameWithDefaults() *OpeningExplorerLichessGame`

NewOpeningExplorerLichessGameWithDefaults instantiates a new OpeningExplorerLichessGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OpeningExplorerLichessGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerLichessGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerLichessGame) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerLichessGame) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerLichessGame) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerLichessGame) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.


### SetWinnerNil

`func (o *OpeningExplorerLichessGame) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *OpeningExplorerLichessGame) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetSpeed

`func (o *OpeningExplorerLichessGame) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OpeningExplorerLichessGame) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OpeningExplorerLichessGame) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OpeningExplorerLichessGame) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetWhite

`func (o *OpeningExplorerLichessGame) GetWhite() OpeningExplorerGamePlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerLichessGame) GetWhiteOk() (*OpeningExplorerGamePlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerLichessGame) SetWhite(v OpeningExplorerGamePlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerLichessGame) GetBlack() OpeningExplorerGamePlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerLichessGame) GetBlackOk() (*OpeningExplorerGamePlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerLichessGame) SetBlack(v OpeningExplorerGamePlayer)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerLichessGame) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerLichessGame) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerLichessGame) SetYear(v float32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerLichessGame) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerLichessGame) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerLichessGame) SetMonth(v string)`

SetMonth sets Month field to given value.


### SetMonthNil

`func (o *OpeningExplorerLichessGame) SetMonthNil(b bool)`

 SetMonthNil sets the value for Month to be an explicit nil

### UnsetMonth
`func (o *OpeningExplorerLichessGame) UnsetMonth()`

UnsetMonth ensures that no value is present for Month, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


