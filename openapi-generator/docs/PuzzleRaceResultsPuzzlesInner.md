# PuzzleRaceResultsPuzzlesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Puzzle ID | 
**Fen** | **string** | X-FEN position of the puzzle | 
**Line** | **string** | Solution moves sequence | 
**Rating** | **int32** | Puzzle Glicko2 rating | 

## Methods

### NewPuzzleRaceResultsPuzzlesInner

`func NewPuzzleRaceResultsPuzzlesInner(id string, fen string, line string, rating int32, ) *PuzzleRaceResultsPuzzlesInner`

NewPuzzleRaceResultsPuzzlesInner instantiates a new PuzzleRaceResultsPuzzlesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleRaceResultsPuzzlesInnerWithDefaults

`func NewPuzzleRaceResultsPuzzlesInnerWithDefaults() *PuzzleRaceResultsPuzzlesInner`

NewPuzzleRaceResultsPuzzlesInnerWithDefaults instantiates a new PuzzleRaceResultsPuzzlesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PuzzleRaceResultsPuzzlesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleRaceResultsPuzzlesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleRaceResultsPuzzlesInner) SetId(v string)`

SetId sets Id field to given value.


### GetFen

`func (o *PuzzleRaceResultsPuzzlesInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *PuzzleRaceResultsPuzzlesInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *PuzzleRaceResultsPuzzlesInner) SetFen(v string)`

SetFen sets Fen field to given value.


### GetLine

`func (o *PuzzleRaceResultsPuzzlesInner) GetLine() string`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *PuzzleRaceResultsPuzzlesInner) GetLineOk() (*string, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *PuzzleRaceResultsPuzzlesInner) SetLine(v string)`

SetLine sets Line field to given value.


### GetRating

`func (o *PuzzleRaceResultsPuzzlesInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PuzzleRaceResultsPuzzlesInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PuzzleRaceResultsPuzzlesInner) SetRating(v int32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


