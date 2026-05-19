# ApiTournamentPost200ResponseStandingPlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Sheet** | Pointer to [**ApiTournamentPost200ResponseStandingPlayersInnerSheet**](ApiTournamentPost200ResponseStandingPlayersInnerSheet.md) |  | [optional] 

## Methods

### NewApiTournamentPost200ResponseStandingPlayersInner

`func NewApiTournamentPost200ResponseStandingPlayersInner() *ApiTournamentPost200ResponseStandingPlayersInner`

NewApiTournamentPost200ResponseStandingPlayersInner instantiates a new ApiTournamentPost200ResponseStandingPlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTournamentPost200ResponseStandingPlayersInnerWithDefaults

`func NewApiTournamentPost200ResponseStandingPlayersInnerWithDefaults() *ApiTournamentPost200ResponseStandingPlayersInner`

NewApiTournamentPost200ResponseStandingPlayersInnerWithDefaults instantiates a new ApiTournamentPost200ResponseStandingPlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ApiTournamentPost200ResponseStandingPlayersInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPatron

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetFlair

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetRank

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRating

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScore

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetSheet

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetSheet() ApiTournamentPost200ResponseStandingPlayersInnerSheet`

GetSheet returns the Sheet field if non-nil, zero value otherwise.

### GetSheetOk

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) GetSheetOk() (*ApiTournamentPost200ResponseStandingPlayersInnerSheet, bool)`

GetSheetOk returns a tuple with the Sheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheet

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) SetSheet(v ApiTournamentPost200ResponseStandingPlayersInnerSheet)`

SetSheet sets Sheet field to given value.

### HasSheet

`func (o *ApiTournamentPost200ResponseStandingPlayersInner) HasSheet() bool`

HasSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


