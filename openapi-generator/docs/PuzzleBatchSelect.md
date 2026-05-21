# PuzzleBatchSelect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puzzles** | Pointer to [**[]PuzzleAndGame**](PuzzleAndGame.md) |  | [optional] 
**Glicko** | Pointer to [**PuzzleGlicko**](PuzzleGlicko.md) |  | [optional] 

## Methods

### NewPuzzleBatchSelect

`func NewPuzzleBatchSelect() *PuzzleBatchSelect`

NewPuzzleBatchSelect instantiates a new PuzzleBatchSelect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleBatchSelectWithDefaults

`func NewPuzzleBatchSelectWithDefaults() *PuzzleBatchSelect`

NewPuzzleBatchSelectWithDefaults instantiates a new PuzzleBatchSelect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuzzles

`func (o *PuzzleBatchSelect) GetPuzzles() []PuzzleAndGame`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *PuzzleBatchSelect) GetPuzzlesOk() (*[]PuzzleAndGame, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *PuzzleBatchSelect) SetPuzzles(v []PuzzleAndGame)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *PuzzleBatchSelect) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetGlicko

`func (o *PuzzleBatchSelect) GetGlicko() PuzzleGlicko`

GetGlicko returns the Glicko field if non-nil, zero value otherwise.

### GetGlickoOk

`func (o *PuzzleBatchSelect) GetGlickoOk() (*PuzzleGlicko, bool)`

GetGlickoOk returns a tuple with the Glicko field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlicko

`func (o *PuzzleBatchSelect) SetGlicko(v PuzzleGlicko)`

SetGlicko sets Glicko field to given value.

### HasGlicko

`func (o *PuzzleBatchSelect) HasGlicko() bool`

HasGlicko returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


