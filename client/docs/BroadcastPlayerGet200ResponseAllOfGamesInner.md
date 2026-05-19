# BroadcastPlayerGet200ResponseAllOfGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Round** | **string** | ID of the round | 
**Id** | **string** | The game ID. Analogous to chapterId. | 
**Opponent** | [**BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent**](BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent.md) |  | 
**Color** | **NullableString** |  | 
**Points** | Pointer to **string** |  | [optional] 
**CustomPoints** | Pointer to **float32** |  | [optional] 
**RatingDiff** | Pointer to **int32** | The change in rating for the player as a result of this game | [optional] 
**FideTC** | **string** | FIDE rating category  | 
**Ongoing** | Pointer to **bool** |  | [optional] 

## Methods

### NewBroadcastPlayerGet200ResponseAllOfGamesInner

`func NewBroadcastPlayerGet200ResponseAllOfGamesInner(round string, id string, opponent BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent, color NullableString, fideTC string, ) *BroadcastPlayerGet200ResponseAllOfGamesInner`

NewBroadcastPlayerGet200ResponseAllOfGamesInner instantiates a new BroadcastPlayerGet200ResponseAllOfGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPlayerGet200ResponseAllOfGamesInnerWithDefaults

`func NewBroadcastPlayerGet200ResponseAllOfGamesInnerWithDefaults() *BroadcastPlayerGet200ResponseAllOfGamesInner`

NewBroadcastPlayerGet200ResponseAllOfGamesInnerWithDefaults instantiates a new BroadcastPlayerGet200ResponseAllOfGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRound

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetRound(v string)`

SetRound sets Round field to given value.


### GetId

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetOpponent

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOpponent() BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOpponentOk() (*BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetOpponent(v BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent)`

SetOpponent sets Opponent field to given value.


### GetColor

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetPoints() string`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetPointsOk() (*string, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetPoints(v string)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetCustomPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetCustomPoints() float32`

GetCustomPoints returns the CustomPoints field if non-nil, zero value otherwise.

### GetCustomPointsOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetCustomPointsOk() (*float32, bool)`

GetCustomPointsOk returns a tuple with the CustomPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetCustomPoints(v float32)`

SetCustomPoints sets CustomPoints field to given value.

### HasCustomPoints

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasCustomPoints() bool`

HasCustomPoints returns a boolean if a field has been set.

### GetRatingDiff

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetFideTC

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetFideTC() string`

GetFideTC returns the FideTC field if non-nil, zero value otherwise.

### GetFideTCOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetFideTCOk() (*string, bool)`

GetFideTCOk returns a tuple with the FideTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideTC

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetFideTC(v string)`

SetFideTC sets FideTC field to given value.


### GetOngoing

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


