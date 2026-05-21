# GamePlayers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**White** | [**GamePlayerUser**](GamePlayerUser.md) |  | 
**Black** | [**GamePlayerUser**](GamePlayerUser.md) |  | 

## Methods

### NewGamePlayers

`func NewGamePlayers(white GamePlayerUser, black GamePlayerUser, ) *GamePlayers`

NewGamePlayers instantiates a new GamePlayers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamePlayersWithDefaults

`func NewGamePlayersWithDefaults() *GamePlayers`

NewGamePlayersWithDefaults instantiates a new GamePlayers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWhite

`func (o *GamePlayers) GetWhite() GamePlayerUser`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *GamePlayers) GetWhiteOk() (*GamePlayerUser, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *GamePlayers) SetWhite(v GamePlayerUser)`

SetWhite sets White field to given value.


### GetBlack

`func (o *GamePlayers) GetBlack() GamePlayerUser`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *GamePlayers) GetBlackOk() (*GamePlayerUser, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *GamePlayers) SetBlack(v GamePlayerUser)`

SetBlack sets Black field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


