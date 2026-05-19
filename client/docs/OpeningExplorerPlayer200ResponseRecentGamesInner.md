# OpeningExplorerPlayer200ResponseRecentGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uci** | **string** |  | 
**Id** | **string** |  | 
**Winner** | **NullableString** |  | 
**Speed** | **string** |  | 
**Mode** | **string** |  | 
**White** | [**OpeningExplorerMaster200ResponseMovesInnerGameWhite**](OpeningExplorerMaster200ResponseMovesInnerGameWhite.md) |  | 
**Black** | [**OpeningExplorerMaster200ResponseMovesInnerGameWhite**](OpeningExplorerMaster200ResponseMovesInnerGameWhite.md) |  | 
**Year** | **int32** |  | 
**Month** | **string** |  | 

## Methods

### NewOpeningExplorerPlayer200ResponseRecentGamesInner

`func NewOpeningExplorerPlayer200ResponseRecentGamesInner(uci string, id string, winner NullableString, speed string, mode string, white OpeningExplorerMaster200ResponseMovesInnerGameWhite, black OpeningExplorerMaster200ResponseMovesInnerGameWhite, year int32, month string, ) *OpeningExplorerPlayer200ResponseRecentGamesInner`

NewOpeningExplorerPlayer200ResponseRecentGamesInner instantiates a new OpeningExplorerPlayer200ResponseRecentGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpeningExplorerPlayer200ResponseRecentGamesInnerWithDefaults

`func NewOpeningExplorerPlayer200ResponseRecentGamesInnerWithDefaults() *OpeningExplorerPlayer200ResponseRecentGamesInner`

NewOpeningExplorerPlayer200ResponseRecentGamesInnerWithDefaults instantiates a new OpeningExplorerPlayer200ResponseRecentGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUci

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetUci() string`

GetUci returns the Uci field if non-nil, zero value otherwise.

### GetUciOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetUciOk() (*string, bool)`

GetUciOk returns a tuple with the Uci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUci

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetUci(v string)`

SetUci sets Uci field to given value.


### GetId

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetWinner

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetWinner(v string)`

SetWinner sets Winner field to given value.


### SetWinnerNil

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil
### GetSpeed

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetMode

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetMode(v string)`

SetMode sets Mode field to given value.


### GetWhite

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetWhite() OpeningExplorerMaster200ResponseMovesInnerGameWhite`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetWhiteOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetWhite(v OpeningExplorerMaster200ResponseMovesInnerGameWhite)`

SetWhite sets White field to given value.


### GetBlack

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetBlack() OpeningExplorerMaster200ResponseMovesInnerGameWhite`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetBlackOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetBlack(v OpeningExplorerMaster200ResponseMovesInnerGameWhite)`

SetBlack sets Black field to given value.


### GetYear

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetYear(v int32)`

SetYear sets Year field to given value.


### GetMonth

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *OpeningExplorerPlayer200ResponseRecentGamesInner) SetMonth(v string)`

SetMonth sets Month field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


