# PuzzleRaceResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the puzzle race | 
**Owner** | **string** | Owner of the puzzle race | 
**Players** | [**[]PuzzleRaceResultsPlayersInner**](PuzzleRaceResultsPlayersInner.md) | List of players participating in the race | 
**Puzzles** | [**[]PuzzleRaceResultsPuzzlesInner**](PuzzleRaceResultsPuzzlesInner.md) | List of puzzles in the race | 
**FinishesAt** | **int32** | Timestamp in milliseconds when the race finishes | 
**StartsAt** | **int32** | Timestamp in milliseconds when the race started | 

## Methods

### NewPuzzleRaceResults

`func NewPuzzleRaceResults(id string, owner string, players []PuzzleRaceResultsPlayersInner, puzzles []PuzzleRaceResultsPuzzlesInner, finishesAt int32, startsAt int32, ) *PuzzleRaceResults`

NewPuzzleRaceResults instantiates a new PuzzleRaceResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleRaceResultsWithDefaults

`func NewPuzzleRaceResultsWithDefaults() *PuzzleRaceResults`

NewPuzzleRaceResultsWithDefaults instantiates a new PuzzleRaceResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PuzzleRaceResults) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleRaceResults) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleRaceResults) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *PuzzleRaceResults) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PuzzleRaceResults) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PuzzleRaceResults) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPlayers

`func (o *PuzzleRaceResults) GetPlayers() []PuzzleRaceResultsPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *PuzzleRaceResults) GetPlayersOk() (*[]PuzzleRaceResultsPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *PuzzleRaceResults) SetPlayers(v []PuzzleRaceResultsPlayersInner)`

SetPlayers sets Players field to given value.


### GetPuzzles

`func (o *PuzzleRaceResults) GetPuzzles() []PuzzleRaceResultsPuzzlesInner`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *PuzzleRaceResults) GetPuzzlesOk() (*[]PuzzleRaceResultsPuzzlesInner, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *PuzzleRaceResults) SetPuzzles(v []PuzzleRaceResultsPuzzlesInner)`

SetPuzzles sets Puzzles field to given value.


### GetFinishesAt

`func (o *PuzzleRaceResults) GetFinishesAt() int32`

GetFinishesAt returns the FinishesAt field if non-nil, zero value otherwise.

### GetFinishesAtOk

`func (o *PuzzleRaceResults) GetFinishesAtOk() (*int32, bool)`

GetFinishesAtOk returns a tuple with the FinishesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishesAt

`func (o *PuzzleRaceResults) SetFinishesAt(v int32)`

SetFinishesAt sets FinishesAt field to given value.


### GetStartsAt

`func (o *PuzzleRaceResults) GetStartsAt() int32`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *PuzzleRaceResults) GetStartsAtOk() (*int32, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *PuzzleRaceResults) SetStartsAt(v int32)`

SetStartsAt sets StartsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


