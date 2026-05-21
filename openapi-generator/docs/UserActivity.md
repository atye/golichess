# UserActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**UserActivityInterval**](UserActivityInterval.md) |  | 
**Games** | Pointer to [**UserActivityGames**](UserActivityGames.md) |  | [optional] 
**Puzzles** | Pointer to [**UserActivityPuzzles**](UserActivityPuzzles.md) |  | [optional] 
**Storm** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 
**Racer** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 
**Streak** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 
**Tournaments** | Pointer to [**UserActivityTournaments**](UserActivityTournaments.md) |  | [optional] 
**Practice** | Pointer to [**[]UserActivityPracticeInner**](UserActivityPracticeInner.md) |  | [optional] 
**Simuls** | Pointer to **[]string** |  | [optional] 
**CorrespondenceMoves** | Pointer to [**UserActivityCorrespondenceMoves**](UserActivityCorrespondenceMoves.md) |  | [optional] 
**CorrespondenceEnds** | Pointer to [**UserActivityCorrespondenceEnds**](UserActivityCorrespondenceEnds.md) |  | [optional] 
**Follows** | Pointer to [**UserActivityFollows**](UserActivityFollows.md) |  | [optional] 
**Studies** | Pointer to [**[]UserActivityTournamentsBestInnerTournament**](UserActivityTournamentsBestInnerTournament.md) |  | [optional] 
**Teams** | Pointer to [**[]UserActivityTeamsInner**](UserActivityTeamsInner.md) |  | [optional] 
**Posts** | Pointer to [**[]UserActivityPostsInner**](UserActivityPostsInner.md) |  | [optional] 
**Patron** | Pointer to [**UserActivityPatron**](UserActivityPatron.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 

## Methods

### NewUserActivity

`func NewUserActivity(interval UserActivityInterval, ) *UserActivity`

NewUserActivity instantiates a new UserActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityWithDefaults

`func NewUserActivityWithDefaults() *UserActivity`

NewUserActivityWithDefaults instantiates a new UserActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *UserActivity) GetInterval() UserActivityInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *UserActivity) GetIntervalOk() (*UserActivityInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *UserActivity) SetInterval(v UserActivityInterval)`

SetInterval sets Interval field to given value.


### GetGames

`func (o *UserActivity) GetGames() UserActivityGames`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *UserActivity) GetGamesOk() (*UserActivityGames, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *UserActivity) SetGames(v UserActivityGames)`

SetGames sets Games field to given value.

### HasGames

`func (o *UserActivity) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetPuzzles

`func (o *UserActivity) GetPuzzles() UserActivityPuzzles`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *UserActivity) GetPuzzlesOk() (*UserActivityPuzzles, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *UserActivity) SetPuzzles(v UserActivityPuzzles)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *UserActivity) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetStorm

`func (o *UserActivity) GetStorm() PuzzleModePerf`

GetStorm returns the Storm field if non-nil, zero value otherwise.

### GetStormOk

`func (o *UserActivity) GetStormOk() (*PuzzleModePerf, bool)`

GetStormOk returns a tuple with the Storm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorm

`func (o *UserActivity) SetStorm(v PuzzleModePerf)`

SetStorm sets Storm field to given value.

### HasStorm

`func (o *UserActivity) HasStorm() bool`

HasStorm returns a boolean if a field has been set.

### GetRacer

`func (o *UserActivity) GetRacer() PuzzleModePerf`

GetRacer returns the Racer field if non-nil, zero value otherwise.

### GetRacerOk

`func (o *UserActivity) GetRacerOk() (*PuzzleModePerf, bool)`

GetRacerOk returns a tuple with the Racer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacer

`func (o *UserActivity) SetRacer(v PuzzleModePerf)`

SetRacer sets Racer field to given value.

### HasRacer

`func (o *UserActivity) HasRacer() bool`

HasRacer returns a boolean if a field has been set.

### GetStreak

`func (o *UserActivity) GetStreak() PuzzleModePerf`

GetStreak returns the Streak field if non-nil, zero value otherwise.

### GetStreakOk

`func (o *UserActivity) GetStreakOk() (*PuzzleModePerf, bool)`

GetStreakOk returns a tuple with the Streak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreak

`func (o *UserActivity) SetStreak(v PuzzleModePerf)`

SetStreak sets Streak field to given value.

### HasStreak

`func (o *UserActivity) HasStreak() bool`

HasStreak returns a boolean if a field has been set.

### GetTournaments

`func (o *UserActivity) GetTournaments() UserActivityTournaments`

GetTournaments returns the Tournaments field if non-nil, zero value otherwise.

### GetTournamentsOk

`func (o *UserActivity) GetTournamentsOk() (*UserActivityTournaments, bool)`

GetTournamentsOk returns a tuple with the Tournaments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournaments

`func (o *UserActivity) SetTournaments(v UserActivityTournaments)`

SetTournaments sets Tournaments field to given value.

### HasTournaments

`func (o *UserActivity) HasTournaments() bool`

HasTournaments returns a boolean if a field has been set.

### GetPractice

`func (o *UserActivity) GetPractice() []UserActivityPracticeInner`

GetPractice returns the Practice field if non-nil, zero value otherwise.

### GetPracticeOk

`func (o *UserActivity) GetPracticeOk() (*[]UserActivityPracticeInner, bool)`

GetPracticeOk returns a tuple with the Practice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPractice

`func (o *UserActivity) SetPractice(v []UserActivityPracticeInner)`

SetPractice sets Practice field to given value.

### HasPractice

`func (o *UserActivity) HasPractice() bool`

HasPractice returns a boolean if a field has been set.

### GetSimuls

`func (o *UserActivity) GetSimuls() []string`

GetSimuls returns the Simuls field if non-nil, zero value otherwise.

### GetSimulsOk

`func (o *UserActivity) GetSimulsOk() (*[]string, bool)`

GetSimulsOk returns a tuple with the Simuls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimuls

`func (o *UserActivity) SetSimuls(v []string)`

SetSimuls sets Simuls field to given value.

### HasSimuls

`func (o *UserActivity) HasSimuls() bool`

HasSimuls returns a boolean if a field has been set.

### GetCorrespondenceMoves

`func (o *UserActivity) GetCorrespondenceMoves() UserActivityCorrespondenceMoves`

GetCorrespondenceMoves returns the CorrespondenceMoves field if non-nil, zero value otherwise.

### GetCorrespondenceMovesOk

`func (o *UserActivity) GetCorrespondenceMovesOk() (*UserActivityCorrespondenceMoves, bool)`

GetCorrespondenceMovesOk returns a tuple with the CorrespondenceMoves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondenceMoves

`func (o *UserActivity) SetCorrespondenceMoves(v UserActivityCorrespondenceMoves)`

SetCorrespondenceMoves sets CorrespondenceMoves field to given value.

### HasCorrespondenceMoves

`func (o *UserActivity) HasCorrespondenceMoves() bool`

HasCorrespondenceMoves returns a boolean if a field has been set.

### GetCorrespondenceEnds

`func (o *UserActivity) GetCorrespondenceEnds() UserActivityCorrespondenceEnds`

GetCorrespondenceEnds returns the CorrespondenceEnds field if non-nil, zero value otherwise.

### GetCorrespondenceEndsOk

`func (o *UserActivity) GetCorrespondenceEndsOk() (*UserActivityCorrespondenceEnds, bool)`

GetCorrespondenceEndsOk returns a tuple with the CorrespondenceEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondenceEnds

`func (o *UserActivity) SetCorrespondenceEnds(v UserActivityCorrespondenceEnds)`

SetCorrespondenceEnds sets CorrespondenceEnds field to given value.

### HasCorrespondenceEnds

`func (o *UserActivity) HasCorrespondenceEnds() bool`

HasCorrespondenceEnds returns a boolean if a field has been set.

### GetFollows

`func (o *UserActivity) GetFollows() UserActivityFollows`

GetFollows returns the Follows field if non-nil, zero value otherwise.

### GetFollowsOk

`func (o *UserActivity) GetFollowsOk() (*UserActivityFollows, bool)`

GetFollowsOk returns a tuple with the Follows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollows

`func (o *UserActivity) SetFollows(v UserActivityFollows)`

SetFollows sets Follows field to given value.

### HasFollows

`func (o *UserActivity) HasFollows() bool`

HasFollows returns a boolean if a field has been set.

### GetStudies

`func (o *UserActivity) GetStudies() []UserActivityTournamentsBestInnerTournament`

GetStudies returns the Studies field if non-nil, zero value otherwise.

### GetStudiesOk

`func (o *UserActivity) GetStudiesOk() (*[]UserActivityTournamentsBestInnerTournament, bool)`

GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudies

`func (o *UserActivity) SetStudies(v []UserActivityTournamentsBestInnerTournament)`

SetStudies sets Studies field to given value.

### HasStudies

`func (o *UserActivity) HasStudies() bool`

HasStudies returns a boolean if a field has been set.

### GetTeams

`func (o *UserActivity) GetTeams() []UserActivityTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *UserActivity) GetTeamsOk() (*[]UserActivityTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *UserActivity) SetTeams(v []UserActivityTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *UserActivity) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetPosts

`func (o *UserActivity) GetPosts() []UserActivityPostsInner`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *UserActivity) GetPostsOk() (*[]UserActivityPostsInner, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *UserActivity) SetPosts(v []UserActivityPostsInner)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *UserActivity) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPatron

`func (o *UserActivity) GetPatron() UserActivityPatron`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *UserActivity) GetPatronOk() (*UserActivityPatron, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *UserActivity) SetPatron(v UserActivityPatron)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *UserActivity) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetStream

`func (o *UserActivity) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *UserActivity) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *UserActivity) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *UserActivity) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


