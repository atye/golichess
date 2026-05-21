# PuzzleBatchSolveResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puzzles** | Pointer to [**[]PuzzleAndGame**](PuzzleAndGame.md) |  | [optional] 
**Glicko** | Pointer to [**PuzzleGlicko**](PuzzleGlicko.md) |  | [optional] 
**Rounds** | Pointer to [**[]PuzzleBatchSolveResponseRoundsInner**](PuzzleBatchSolveResponseRoundsInner.md) |  | [optional] 

## Methods

### NewPuzzleBatchSolveResponse

`func NewPuzzleBatchSolveResponse() *PuzzleBatchSolveResponse`

NewPuzzleBatchSolveResponse instantiates a new PuzzleBatchSolveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleBatchSolveResponseWithDefaults

`func NewPuzzleBatchSolveResponseWithDefaults() *PuzzleBatchSolveResponse`

NewPuzzleBatchSolveResponseWithDefaults instantiates a new PuzzleBatchSolveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuzzles

`func (o *PuzzleBatchSolveResponse) GetPuzzles() []PuzzleAndGame`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *PuzzleBatchSolveResponse) GetPuzzlesOk() (*[]PuzzleAndGame, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *PuzzleBatchSolveResponse) SetPuzzles(v []PuzzleAndGame)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *PuzzleBatchSolveResponse) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetGlicko

`func (o *PuzzleBatchSolveResponse) GetGlicko() PuzzleGlicko`

GetGlicko returns the Glicko field if non-nil, zero value otherwise.

### GetGlickoOk

`func (o *PuzzleBatchSolveResponse) GetGlickoOk() (*PuzzleGlicko, bool)`

GetGlickoOk returns a tuple with the Glicko field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlicko

`func (o *PuzzleBatchSolveResponse) SetGlicko(v PuzzleGlicko)`

SetGlicko sets Glicko field to given value.

### HasGlicko

`func (o *PuzzleBatchSolveResponse) HasGlicko() bool`

HasGlicko returns a boolean if a field has been set.

### GetRounds

`func (o *PuzzleBatchSolveResponse) GetRounds() []PuzzleBatchSolveResponseRoundsInner`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *PuzzleBatchSolveResponse) GetRoundsOk() (*[]PuzzleBatchSolveResponseRoundsInner, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *PuzzleBatchSolveResponse) SetRounds(v []PuzzleBatchSolveResponseRoundsInner)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *PuzzleBatchSolveResponse) HasRounds() bool`

HasRounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


