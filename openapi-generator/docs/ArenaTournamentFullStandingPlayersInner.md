# ArenaTournamentFullStandingPlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Sheet** | Pointer to [**ArenaSheet**](ArenaSheet.md) |  | [optional] 

## Methods

### NewArenaTournamentFullStandingPlayersInner

`func NewArenaTournamentFullStandingPlayersInner() *ArenaTournamentFullStandingPlayersInner`

NewArenaTournamentFullStandingPlayersInner instantiates a new ArenaTournamentFullStandingPlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaTournamentFullStandingPlayersInnerWithDefaults

`func NewArenaTournamentFullStandingPlayersInnerWithDefaults() *ArenaTournamentFullStandingPlayersInner`

NewArenaTournamentFullStandingPlayersInnerWithDefaults instantiates a new ArenaTournamentFullStandingPlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArenaTournamentFullStandingPlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArenaTournamentFullStandingPlayersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArenaTournamentFullStandingPlayersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ArenaTournamentFullStandingPlayersInner) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ArenaTournamentFullStandingPlayersInner) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ArenaTournamentFullStandingPlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *ArenaTournamentFullStandingPlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ArenaTournamentFullStandingPlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ArenaTournamentFullStandingPlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ArenaTournamentFullStandingPlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ArenaTournamentFullStandingPlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ArenaTournamentFullStandingPlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetFlair

`func (o *ArenaTournamentFullStandingPlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ArenaTournamentFullStandingPlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ArenaTournamentFullStandingPlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetRank

`func (o *ArenaTournamentFullStandingPlayersInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ArenaTournamentFullStandingPlayersInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ArenaTournamentFullStandingPlayersInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRating

`func (o *ArenaTournamentFullStandingPlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ArenaTournamentFullStandingPlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ArenaTournamentFullStandingPlayersInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScore

`func (o *ArenaTournamentFullStandingPlayersInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ArenaTournamentFullStandingPlayersInner) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ArenaTournamentFullStandingPlayersInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSheet

`func (o *ArenaTournamentFullStandingPlayersInner) GetSheet() ArenaSheet`

GetSheet returns the Sheet field if non-nil, zero value otherwise.

### GetSheetOk

`func (o *ArenaTournamentFullStandingPlayersInner) GetSheetOk() (*ArenaSheet, bool)`

GetSheetOk returns a tuple with the Sheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheet

`func (o *ArenaTournamentFullStandingPlayersInner) SetSheet(v ArenaSheet)`

SetSheet sets Sheet field to given value.

### HasSheet

`func (o *ArenaTournamentFullStandingPlayersInner) HasSheet() bool`

HasSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


