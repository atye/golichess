# RacerGet200ResponsePuzzlesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Puzzle ID | 
**Fen** | **string** | X-FEN position of the puzzle | 
**Line** | **string** | Solution moves sequence | 
**Rating** | **int32** | Puzzle Glicko2 rating | 

## Methods

### NewRacerGet200ResponsePuzzlesInner

`func NewRacerGet200ResponsePuzzlesInner(id string, fen string, line string, rating int32, ) *RacerGet200ResponsePuzzlesInner`

NewRacerGet200ResponsePuzzlesInner instantiates a new RacerGet200ResponsePuzzlesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacerGet200ResponsePuzzlesInnerWithDefaults

`func NewRacerGet200ResponsePuzzlesInnerWithDefaults() *RacerGet200ResponsePuzzlesInner`

NewRacerGet200ResponsePuzzlesInnerWithDefaults instantiates a new RacerGet200ResponsePuzzlesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RacerGet200ResponsePuzzlesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RacerGet200ResponsePuzzlesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RacerGet200ResponsePuzzlesInner) SetId(v string)`

SetId sets Id field to given value.


### GetFen

`func (o *RacerGet200ResponsePuzzlesInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *RacerGet200ResponsePuzzlesInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *RacerGet200ResponsePuzzlesInner) SetFen(v string)`

SetFen sets Fen field to given value.


### GetLine

`func (o *RacerGet200ResponsePuzzlesInner) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *RacerGet200ResponsePuzzlesInner) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *RacerGet200ResponsePuzzlesInner) SetLine(v string)`

SetLine sets Line field to given value.


### GetRating

`func (o *RacerGet200ResponsePuzzlesInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *RacerGet200ResponsePuzzlesInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *RacerGet200ResponsePuzzlesInner) SetRating(v int32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


