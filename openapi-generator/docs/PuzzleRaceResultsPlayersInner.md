# PuzzleRaceResultsPlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Player username | 
**Score** | **int32** | Player&#39;s current score in the race | 
**Id** | Pointer to **string** | User ID. Missing if player is anonymous. | [optional] 
**Flair** | Pointer to **string** | User&#39;s flair icon | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewPuzzleRaceResultsPlayersInner

`func NewPuzzleRaceResultsPlayersInner(name string, score int32, ) *PuzzleRaceResultsPlayersInner`

NewPuzzleRaceResultsPlayersInner instantiates a new PuzzleRaceResultsPlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleRaceResultsPlayersInnerWithDefaults

`func NewPuzzleRaceResultsPlayersInnerWithDefaults() *PuzzleRaceResultsPlayersInner`

NewPuzzleRaceResultsPlayersInnerWithDefaults instantiates a new PuzzleRaceResultsPlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PuzzleRaceResultsPlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PuzzleRaceResultsPlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PuzzleRaceResultsPlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetScore

`func (o *PuzzleRaceResultsPlayersInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PuzzleRaceResultsPlayersInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PuzzleRaceResultsPlayersInner) SetScore(v int32)`

SetScore sets Score field to given value.


### GetId

`func (o *PuzzleRaceResultsPlayersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleRaceResultsPlayersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleRaceResultsPlayersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PuzzleRaceResultsPlayersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlair

`func (o *PuzzleRaceResultsPlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *PuzzleRaceResultsPlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *PuzzleRaceResultsPlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *PuzzleRaceResultsPlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *PuzzleRaceResultsPlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *PuzzleRaceResultsPlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *PuzzleRaceResultsPlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *PuzzleRaceResultsPlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *PuzzleRaceResultsPlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *PuzzleRaceResultsPlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *PuzzleRaceResultsPlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *PuzzleRaceResultsPlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


