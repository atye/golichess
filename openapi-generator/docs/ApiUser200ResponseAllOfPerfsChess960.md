# ApiUser200ResponseAllOfPerfsChess960

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Games** | **int32** |  | 
**Rating** | **int32** |  | 
**Rd** | **int32** | rating deviation | 
**Prog** | **int32** |  | 
**Prov** | Pointer to **bool** | only appears if a user&#39;s perf rating are [provisional](https://lichess.org/faq#provisional) | [optional] 
**Rank** | Pointer to **int32** | global lichess ranking, only appears for recently active players | [optional] 

## Methods

### NewApiUser200ResponseAllOfPerfsChess960

`func NewApiUser200ResponseAllOfPerfsChess960(games int32, rating int32, rd int32, prog int32, ) *ApiUser200ResponseAllOfPerfsChess960`

NewApiUser200ResponseAllOfPerfsChess960 instantiates a new ApiUser200ResponseAllOfPerfsChess960 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUser200ResponseAllOfPerfsChess960WithDefaults

`func NewApiUser200ResponseAllOfPerfsChess960WithDefaults() *ApiUser200ResponseAllOfPerfsChess960`

NewApiUser200ResponseAllOfPerfsChess960WithDefaults instantiates a new ApiUser200ResponseAllOfPerfsChess960 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGames

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetGames() int32`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetGamesOk() (*int32, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetGames(v int32)`

SetGames sets Games field to given value.


### GetRating

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetRd

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRd() int32`

GetRd returns the Rd field if non-nil, zero value otherwise.

### GetRdOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRdOk() (*int32, bool)`

GetRdOk returns a tuple with the Rd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRd

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetRd(v int32)`

SetRd sets Rd field to given value.


### GetProg

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetProg() int32`

GetProg returns the Prog field if non-nil, zero value otherwise.

### GetProgOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetProgOk() (*int32, bool)`

GetProgOk returns a tuple with the Prog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProg

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetProg(v int32)`

SetProg sets Prog field to given value.


### GetProv

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetProv() bool`

GetProv returns the Prov field if non-nil, zero value otherwise.

### GetProvOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetProvOk() (*bool, bool)`

GetProvOk returns a tuple with the Prov field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProv

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetProv(v bool)`

SetProv sets Prov field to given value.

### HasProv

`func (o *ApiUser200ResponseAllOfPerfsChess960) HasProv() bool`

HasProv returns a boolean if a field has been set.

### GetRank

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiUser200ResponseAllOfPerfsChess960) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiUser200ResponseAllOfPerfsChess960) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ApiUser200ResponseAllOfPerfsChess960) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


