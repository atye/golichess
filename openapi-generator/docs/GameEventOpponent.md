# GameEventOpponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**nil**](nil.md) |  | 
**Username** | **string** |  | 
**Rating** | **int32** |  | 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Ai** | **int32** | AI level, from 1 to 8, where 1 is the weakest and 8 is the strongest. | 

## Methods

### NewGameEventOpponent

`func NewGameEventOpponent(id nil, username string, rating int32, ai int32, ) *GameEventOpponent`

NewGameEventOpponent instantiates a new GameEventOpponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameEventOpponentWithDefaults

`func NewGameEventOpponentWithDefaults() *GameEventOpponent`

NewGameEventOpponentWithDefaults instantiates a new GameEventOpponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameEventOpponent) GetId() nil`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameEventOpponent) GetIdOk() (*nil, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameEventOpponent) SetId(v nil)`

SetId sets Id field to given value.


### GetUsername

`func (o *GameEventOpponent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GameEventOpponent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GameEventOpponent) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetRating

`func (o *GameEventOpponent) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GameEventOpponent) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GameEventOpponent) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetRatingDiff

`func (o *GameEventOpponent) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *GameEventOpponent) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *GameEventOpponent) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *GameEventOpponent) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetAi

`func (o *GameEventOpponent) GetAi() int32`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *GameEventOpponent) GetAiOk() (*int32, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *GameEventOpponent) SetAi(v int32)`

SetAi sets Ai field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


