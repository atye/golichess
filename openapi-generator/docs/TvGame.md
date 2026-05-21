# TvGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**LightUser**](LightUser.md) |  | 
**Rating** | **int32** |  | 
**GameId** | **string** |  | 
**Color** | [**GameColor**](GameColor.md) |  | 

## Methods

### NewTvGame

`func NewTvGame(user LightUser, rating int32, gameId string, color GameColor, ) *TvGame`

NewTvGame instantiates a new TvGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTvGameWithDefaults

`func NewTvGameWithDefaults() *TvGame`

NewTvGameWithDefaults instantiates a new TvGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *TvGame) GetUser() LightUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TvGame) GetUserOk() (*LightUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TvGame) SetUser(v LightUser)`

SetUser sets User field to given value.


### GetRating

`func (o *TvGame) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *TvGame) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *TvGame) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetGameId

`func (o *TvGame) GetGameId() string`

GetGameId returns the GameId field if non-nil, zero value otherwise.

### GetGameIdOk

`func (o *TvGame) GetGameIdOk() (*string, bool)`

GetGameIdOk returns a tuple with the GameId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameId

`func (o *TvGame) SetGameId(v string)`

SetGameId sets GameId field to given value.


### GetColor

`func (o *TvGame) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *TvGame) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *TvGame) SetColor(v GameColor)`

SetColor sets Color field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


