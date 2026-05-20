# ResultsByTournament200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rank** | **int32** |  | 
**Score** | **int32** |  | 
**Rating** | **int32** |  | 
**Username** | **string** |  | 
**Performance** | **int32** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Sheet** | Pointer to [**ApiTournamentPost200ResponseStandingPlayersInnerSheet**](ApiTournamentPost200ResponseStandingPlayersInnerSheet.md) |  | [optional] 

## Methods

### NewResultsByTournament200Response

`func NewResultsByTournament200Response(rank int32, score int32, rating int32, username string, performance int32, ) *ResultsByTournament200Response`

NewResultsByTournament200Response instantiates a new ResultsByTournament200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResultsByTournament200ResponseWithDefaults

`func NewResultsByTournament200ResponseWithDefaults() *ResultsByTournament200Response`

NewResultsByTournament200ResponseWithDefaults instantiates a new ResultsByTournament200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRank

`func (o *ResultsByTournament200Response) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ResultsByTournament200Response) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ResultsByTournament200Response) SetRank(v int32)`

SetRank sets Rank field to given value.


### GetScore

`func (o *ResultsByTournament200Response) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ResultsByTournament200Response) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ResultsByTournament200Response) SetScore(v int32)`

SetScore sets Score field to given value.


### GetRating

`func (o *ResultsByTournament200Response) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ResultsByTournament200Response) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ResultsByTournament200Response) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetUsername

`func (o *ResultsByTournament200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ResultsByTournament200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ResultsByTournament200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerformance

`func (o *ResultsByTournament200Response) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *ResultsByTournament200Response) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *ResultsByTournament200Response) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.


### GetTitle

`func (o *ResultsByTournament200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ResultsByTournament200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ResultsByTournament200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ResultsByTournament200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ResultsByTournament200Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ResultsByTournament200Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTeam

`func (o *ResultsByTournament200Response) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ResultsByTournament200Response) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ResultsByTournament200Response) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ResultsByTournament200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFlair

`func (o *ResultsByTournament200Response) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ResultsByTournament200Response) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ResultsByTournament200Response) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ResultsByTournament200Response) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatronColor

`func (o *ResultsByTournament200Response) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ResultsByTournament200Response) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ResultsByTournament200Response) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ResultsByTournament200Response) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetSheet

`func (o *ResultsByTournament200Response) GetSheet() ApiTournamentPost200ResponseStandingPlayersInnerSheet`

GetSheet returns the Sheet field if non-nil, zero value otherwise.

### GetSheetOk

`func (o *ResultsByTournament200Response) GetSheetOk() (*ApiTournamentPost200ResponseStandingPlayersInnerSheet, bool)`

GetSheetOk returns a tuple with the Sheet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSheet

`func (o *ResultsByTournament200Response) SetSheet(v ApiTournamentPost200ResponseStandingPlayersInnerSheet)`

SetSheet sets Sheet field to given value.

### HasSheet

`func (o *ResultsByTournament200Response) HasSheet() bool`

HasSheet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


