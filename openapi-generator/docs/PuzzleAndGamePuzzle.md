# PuzzleAndGamePuzzle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**InitialPly** | **int32** |  | 
**Plays** | **int32** |  | 
**Rating** | **int32** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**LastMove** | Pointer to **string** | In UCI format, e.g. \&quot;e2e4\&quot; | [optional] 
**Solution** | **[]string** |  | 
**Themes** | **[]string** |  | 

## Methods

### NewPuzzleAndGamePuzzle

`func NewPuzzleAndGamePuzzle(id string, initialPly int32, plays int32, rating int32, solution []string, themes []string, ) *PuzzleAndGamePuzzle`

NewPuzzleAndGamePuzzle instantiates a new PuzzleAndGamePuzzle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleAndGamePuzzleWithDefaults

`func NewPuzzleAndGamePuzzleWithDefaults() *PuzzleAndGamePuzzle`

NewPuzzleAndGamePuzzleWithDefaults instantiates a new PuzzleAndGamePuzzle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PuzzleAndGamePuzzle) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleAndGamePuzzle) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleAndGamePuzzle) SetId(v string)`

SetId sets Id field to given value.


### GetInitialPly

`func (o *PuzzleAndGamePuzzle) GetInitialPly() int32`

GetInitialPly returns the InitialPly field if non-nil, zero value otherwise.

### GetInitialPlyOk

`func (o *PuzzleAndGamePuzzle) GetInitialPlyOk() (*int32, bool)`

GetInitialPlyOk returns a tuple with the InitialPly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialPly

`func (o *PuzzleAndGamePuzzle) SetInitialPly(v int32)`

SetInitialPly sets InitialPly field to given value.


### GetPlays

`func (o *PuzzleAndGamePuzzle) GetPlays() int32`

GetPlays returns the Plays field if non-nil, zero value otherwise.

### GetPlaysOk

`func (o *PuzzleAndGamePuzzle) GetPlaysOk() (*int32, bool)`

GetPlaysOk returns a tuple with the Plays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlays

`func (o *PuzzleAndGamePuzzle) SetPlays(v int32)`

SetPlays sets Plays field to given value.


### GetRating

`func (o *PuzzleAndGamePuzzle) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PuzzleAndGamePuzzle) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PuzzleAndGamePuzzle) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetFen

`func (o *PuzzleAndGamePuzzle) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *PuzzleAndGamePuzzle) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *PuzzleAndGamePuzzle) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *PuzzleAndGamePuzzle) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetLastMove

`func (o *PuzzleAndGamePuzzle) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *PuzzleAndGamePuzzle) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *PuzzleAndGamePuzzle) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *PuzzleAndGamePuzzle) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetSolution

`func (o *PuzzleAndGamePuzzle) GetSolution() []string`

GetSolution returns the Solution field if non-nil, zero value otherwise.

### GetSolutionOk

`func (o *PuzzleAndGamePuzzle) GetSolutionOk() (*[]string, bool)`

GetSolutionOk returns a tuple with the Solution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolution

`func (o *PuzzleAndGamePuzzle) SetSolution(v []string)`

SetSolution sets Solution field to given value.


### GetThemes

`func (o *PuzzleAndGamePuzzle) GetThemes() []string`

GetThemes returns the Themes field if non-nil, zero value otherwise.

### GetThemesOk

`func (o *PuzzleAndGamePuzzle) GetThemesOk() (*[]string, bool)`

GetThemesOk returns a tuple with the Themes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemes

`func (o *PuzzleAndGamePuzzle) SetThemes(v []string)`

SetThemes sets Themes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


