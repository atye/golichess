# PuzzleAndGameGamePlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | [**GameColor**](GameColor.md) |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Rating** | **int32** |  | 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 

## Methods

### NewPuzzleAndGameGamePlayersInner

`func NewPuzzleAndGameGamePlayersInner(color GameColor, id string, name string, rating int32, ) *PuzzleAndGameGamePlayersInner`

NewPuzzleAndGameGamePlayersInner instantiates a new PuzzleAndGameGamePlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPuzzleAndGameGamePlayersInnerWithDefaults

`func NewPuzzleAndGameGamePlayersInnerWithDefaults() *PuzzleAndGameGamePlayersInner`

NewPuzzleAndGameGamePlayersInnerWithDefaults instantiates a new PuzzleAndGameGamePlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *PuzzleAndGameGamePlayersInner) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *PuzzleAndGameGamePlayersInner) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *PuzzleAndGameGamePlayersInner) SetColor(v GameColor)`

SetColor sets Color field to given value.


### GetFlair

`func (o *PuzzleAndGameGamePlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *PuzzleAndGameGamePlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *PuzzleAndGameGamePlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *PuzzleAndGameGamePlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetId

`func (o *PuzzleAndGameGamePlayersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PuzzleAndGameGamePlayersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PuzzleAndGameGamePlayersInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PuzzleAndGameGamePlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PuzzleAndGameGamePlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PuzzleAndGameGamePlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetPatron

`func (o *PuzzleAndGameGamePlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *PuzzleAndGameGamePlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *PuzzleAndGameGamePlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *PuzzleAndGameGamePlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *PuzzleAndGameGamePlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *PuzzleAndGameGamePlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *PuzzleAndGameGamePlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *PuzzleAndGameGamePlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetRating

`func (o *PuzzleAndGameGamePlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *PuzzleAndGameGamePlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *PuzzleAndGameGamePlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetTitle

`func (o *PuzzleAndGameGamePlayersInner) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PuzzleAndGameGamePlayersInner) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PuzzleAndGameGamePlayersInner) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PuzzleAndGameGamePlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


