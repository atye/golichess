# OpeningExplorerLichess200ResponseTopGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**Id** | **string** |  | 
**Winner** | **NullableString** |  | 
**Speed** | Pointer to **string** |  | [optional] 
**White** | [**OpeningExplorerMaster200ResponseMovesInnerGameWhite**](OpeningExplorerMaster200ResponseMovesInnerGameWhite.md) |  | 
**Black** | [**OpeningExplorerMaster200ResponseMovesInnerGameWhite**](OpeningExplorerMaster200ResponseMovesInnerGameWhite.md) |  | 
**Year** | **float32** |  | 
**Month** | **NullableString** |  | 

## Methods

### NewOpeningExplorerLichess200ResponseTopGamesInner

`func NewOpeningExplorerLichess200ResponseTopGamesInner(uci string, id string, winner NullableString, white OpeningExplorerMaster200ResponseMovesInnerGameWhite, black OpeningExplorerMaster200ResponseMovesInnerGameWhite, year float32, month NullableString, ) *OpeningExplorerLichess200ResponseTopGamesInner`

NewOpeningExplorerLichess200ResponseTopGamesInner instantiates a new OpeningExplorerLichess200ResponseTopGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerLichess200ResponseTopGamesInnerWithDefaults

`func NewOpeningExplorerLichess200ResponseTopGamesInnerWithDefaults() *OpeningExplorerLichess200ResponseTopGamesInner`

NewOpeningExplorerLichess200ResponseTopGamesInnerWithDefaults instantiates a new OpeningExplorerLichess200ResponseTopGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetId

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetWinner(v string)`

SetWinner sets Winner field to given value.


### SetWinnerNil

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *OpeningExplorerLichess200ResponseTopGamesInner) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetSpeed

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetWhite

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetWhite() OpeningExplorerMaster200ResponseMovesInnerGameWhite`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetWhiteOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetWhite(v OpeningExplorerMaster200ResponseMovesInnerGameWhite)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetBlack() OpeningExplorerMaster200ResponseMovesInnerGameWhite`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetBlackOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetBlack(v OpeningExplorerMaster200ResponseMovesInnerGameWhite)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetYear(v float32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetMonth(v string)`

SetMonth sets Month field to given value.


### SetMonthNil

`func (o *OpeningExplorerLichess200ResponseTopGamesInner) SetMonthNil(b bool)`

 SetMonthNil sets the value for Month to be an explicit nil

### UnsetMonth
`func (o *OpeningExplorerLichess200ResponseTopGamesInner) UnsetMonth()`

UnsetMonth ensures that no value is present for Month, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


