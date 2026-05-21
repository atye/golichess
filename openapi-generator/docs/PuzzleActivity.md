# PuzzleActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int32** |  | 
**Puzzle** | [**PuzzleActivityPuzzle**](PuzzleActivityPuzzle.md) |  | 
**Win** | **bool** |  | 

## Methods

### NewPuzzleActivity

`func NewPuzzleActivity(date int32, puzzle PuzzleActivityPuzzle, win bool, ) *PuzzleActivity`

NewPuzzleActivity instantiates a new PuzzleActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleActivityWithDefaults

`func NewPuzzleActivityWithDefaults() *PuzzleActivity`

NewPuzzleActivityWithDefaults instantiates a new PuzzleActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *PuzzleActivity) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *PuzzleActivity) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *PuzzleActivity) SetDate(v int32)`

SetDate sets Date field to given value.


### GetPuzzle

`func (o *PuzzleActivity) GetPuzzle() PuzzleActivityPuzzle`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *PuzzleActivity) GetPuzzleOk() (*PuzzleActivityPuzzle, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *PuzzleActivity) SetPuzzle(v PuzzleActivityPuzzle)`

SetPuzzle sets Puzzle field to given value.


### GetWin

`func (o *PuzzleActivity) GetWin() bool`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *PuzzleActivity) GetWinOk() (*bool, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *PuzzleActivity) SetWin(v bool)`

SetWin sets Win field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


