# BroadcastGameEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Round** | **string** | ID of the round | 
**Id** | **string** | The game ID. Analogous to chapterId. | 
**Opponent** | [**BroadcastPlayerWithFed**](BroadcastPlayerWithFed.md) |  | 
**Color** | [**GameColor**](GameColor.md) |  | 
**Points** | Pointer to [**BroadcastPointStr**](BroadcastPointStr.md) |  | [optional] 
**CustomPoints** | Pointer to **float32** |  | [optional] 
**RatingDiff** | Pointer to **int32** | The change in rating for the player as a result of this game | [optional] 
**FideTC** | [**FideTimeControl**](FideTimeControl.md) |  | 
**Ongoing** | Pointer to **bool** |  | [optional] 

## Methods

### NewBroadcastGameEntry

`func NewBroadcastGameEntry(round string, id string, opponent BroadcastPlayerWithFed, color GameColor, fideTC FideTimeControl, ) *BroadcastGameEntry`

NewBroadcastGameEntry instantiates a new BroadcastGameEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastGameEntryWithDefaults

`func NewBroadcastGameEntryWithDefaults() *BroadcastGameEntry`

NewBroadcastGameEntryWithDefaults instantiates a new BroadcastGameEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRound

`func (o *BroadcastGameEntry) GetRound() string`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *BroadcastGameEntry) GetRoundOk() (*string, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *BroadcastGameEntry) SetRound(v string)`

SetRound sets Round field to given value.


### GetId

`func (o *BroadcastGameEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastGameEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastGameEntry) SetId(v string)`

SetId sets Id field to given value.


### GetOpponent

`func (o *BroadcastGameEntry) GetOpponent() BroadcastPlayerWithFed`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *BroadcastGameEntry) GetOpponentOk() (*BroadcastPlayerWithFed, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *BroadcastGameEntry) SetOpponent(v BroadcastPlayerWithFed)`

SetOpponent sets Opponent field to given value.


### GetColor

`func (o *BroadcastGameEntry) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *BroadcastGameEntry) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *BroadcastGameEntry) SetColor(v GameColor)`

SetColor sets Color field to given value.


### GetPoints

`func (o *BroadcastGameEntry) GetPoints() BroadcastPointStr`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BroadcastGameEntry) GetPointsOk() (*BroadcastPointStr, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BroadcastGameEntry) SetPoints(v BroadcastPointStr)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BroadcastGameEntry) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetCustomPoints

`func (o *BroadcastGameEntry) GetCustomPoints() float32`

GetCustomPoints returns the CustomPoints field if non-nil, zero value otherwise.

### GetCustomPointsOk

`func (o *BroadcastGameEntry) GetCustomPointsOk() (*float32, bool)`

GetCustomPointsOk returns a tuple with the CustomPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPoints

`func (o *BroadcastGameEntry) SetCustomPoints(v float32)`

SetCustomPoints sets CustomPoints field to given value.

### HasCustomPoints

`func (o *BroadcastGameEntry) HasCustomPoints() bool`

HasCustomPoints returns a boolean if a field has been set.

### GetRatingDiff

`func (o *BroadcastGameEntry) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *BroadcastGameEntry) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *BroadcastGameEntry) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *BroadcastGameEntry) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetFideTC

`func (o *BroadcastGameEntry) GetFideTC() FideTimeControl`

GetFideTC returns the FideTC field if non-nil, zero value otherwise.

### GetFideTCOk

`func (o *BroadcastGameEntry) GetFideTCOk() (*FideTimeControl, bool)`

GetFideTCOk returns a tuple with the FideTC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideTC

`func (o *BroadcastGameEntry) SetFideTC(v FideTimeControl)`

SetFideTC sets FideTC field to given value.


### GetOngoing

`func (o *BroadcastGameEntry) GetOngoing() bool`

GetOngoing returns the Ongoing field if non-nil, zero value otherwise.

### GetOngoingOk

`func (o *BroadcastGameEntry) GetOngoingOk() (*bool, bool)`

GetOngoingOk returns a tuple with the Ongoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOngoing

`func (o *BroadcastGameEntry) SetOngoing(v bool)`

SetOngoing sets Ongoing field to given value.

### HasOngoing

`func (o *BroadcastGameEntry) HasOngoing() bool`

HasOngoing returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


