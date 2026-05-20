# ApiUserActivity200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Interval** | [**ApiUserActivity200ResponseInnerInterval**](ApiUserActivity200ResponseInnerInterval.md) |  | 
**Games** | Pointer to [**ApiUserActivity200ResponseInnerGames**](ApiUserActivity200ResponseInnerGames.md) |  | [optional] 
**Puzzles** | Pointer to [**ApiUserActivity200ResponseInnerPuzzles**](ApiUserActivity200ResponseInnerPuzzles.md) |  | [optional] 
**Storm** | Pointer to [**ApiUser200ResponseAllOfPerfsStorm**](ApiUser200ResponseAllOfPerfsStorm.md) |  | [optional] 
**Racer** | Pointer to [**ApiUser200ResponseAllOfPerfsStorm**](ApiUser200ResponseAllOfPerfsStorm.md) |  | [optional] 
**Streak** | Pointer to [**ApiUser200ResponseAllOfPerfsStorm**](ApiUser200ResponseAllOfPerfsStorm.md) |  | [optional] 
**Tournaments** | Pointer to [**ApiUserActivity200ResponseInnerTournaments**](ApiUserActivity200ResponseInnerTournaments.md) |  | [optional] 
**Practice** | Pointer to [**[]ApiUserActivity200ResponseInnerPracticeInner**](ApiUserActivity200ResponseInnerPracticeInner.md) |  | [optional] 
**Simuls** | Pointer to **[]string** |  | [optional] 
**CorrespondenceMoves** | Pointer to [**ApiUserActivity200ResponseInnerCorrespondenceMoves**](ApiUserActivity200ResponseInnerCorrespondenceMoves.md) |  | [optional] 
**CorrespondenceEnds** | Pointer to [**ApiUserActivity200ResponseInnerCorrespondenceEnds**](ApiUserActivity200ResponseInnerCorrespondenceEnds.md) |  | [optional] 
**Follows** | Pointer to [**ApiUserActivity200ResponseInnerFollows**](ApiUserActivity200ResponseInnerFollows.md) |  | [optional] 
**Studies** | Pointer to [**[]ApiUserActivity200ResponseInnerTournamentsBestInnerTournament**](ApiUserActivity200ResponseInnerTournamentsBestInnerTournament.md) |  | [optional] 
**Teams** | Pointer to [**[]ApiUserActivity200ResponseInnerTeamsInner**](ApiUserActivity200ResponseInnerTeamsInner.md) |  | [optional] 
**Posts** | Pointer to [**[]ApiUserActivity200ResponseInnerPostsInner**](ApiUserActivity200ResponseInnerPostsInner.md) |  | [optional] 
**Patron** | Pointer to [**ApiUserActivity200ResponseInnerPatron**](ApiUserActivity200ResponseInnerPatron.md) |  | [optional] 
**Stream** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiUserActivity200ResponseInner

`func NewApiUserActivity200ResponseInner(interval ApiUserActivity200ResponseInnerInterval, ) *ApiUserActivity200ResponseInner`

NewApiUserActivity200ResponseInner instantiates a new ApiUserActivity200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserActivity200ResponseInnerWithDefaults

`func NewApiUserActivity200ResponseInnerWithDefaults() *ApiUserActivity200ResponseInner`

NewApiUserActivity200ResponseInnerWithDefaults instantiates a new ApiUserActivity200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterval

`func (o *ApiUserActivity200ResponseInner) GetInterval() ApiUserActivity200ResponseInnerInterval`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *ApiUserActivity200ResponseInner) GetIntervalOk() (*ApiUserActivity200ResponseInnerInterval, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *ApiUserActivity200ResponseInner) SetInterval(v ApiUserActivity200ResponseInnerInterval)`

SetInterval sets Interval field to given value.


### GetGames

`func (o *ApiUserActivity200ResponseInner) GetGames() ApiUserActivity200ResponseInnerGames`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *ApiUserActivity200ResponseInner) GetGamesOk() (*ApiUserActivity200ResponseInnerGames, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *ApiUserActivity200ResponseInner) SetGames(v ApiUserActivity200ResponseInnerGames)`

SetGames sets Games field to given value.

### HasGames

`func (o *ApiUserActivity200ResponseInner) HasGames() bool`

HasGames returns a boolean if a field has been set.

### GetPuzzles

`func (o *ApiUserActivity200ResponseInner) GetPuzzles() ApiUserActivity200ResponseInnerPuzzles`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *ApiUserActivity200ResponseInner) GetPuzzlesOk() (*ApiUserActivity200ResponseInnerPuzzles, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *ApiUserActivity200ResponseInner) SetPuzzles(v ApiUserActivity200ResponseInnerPuzzles)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *ApiUserActivity200ResponseInner) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetStorm

`func (o *ApiUserActivity200ResponseInner) GetStorm() ApiUser200ResponseAllOfPerfsStorm`

GetStorm returns the Storm field if non-nil, zero value otherwise.

### GetStormOk

`func (o *ApiUserActivity200ResponseInner) GetStormOk() (*ApiUser200ResponseAllOfPerfsStorm, bool)`

GetStormOk returns a tuple with the Storm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorm

`func (o *ApiUserActivity200ResponseInner) SetStorm(v ApiUser200ResponseAllOfPerfsStorm)`

SetStorm sets Storm field to given value.

### HasStorm

`func (o *ApiUserActivity200ResponseInner) HasStorm() bool`

HasStorm returns a boolean if a field has been set.

### GetRacer

`func (o *ApiUserActivity200ResponseInner) GetRacer() ApiUser200ResponseAllOfPerfsStorm`

GetRacer returns the Racer field if non-nil, zero value otherwise.

### GetRacerOk

`func (o *ApiUserActivity200ResponseInner) GetRacerOk() (*ApiUser200ResponseAllOfPerfsStorm, bool)`

GetRacerOk returns a tuple with the Racer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacer

`func (o *ApiUserActivity200ResponseInner) SetRacer(v ApiUser200ResponseAllOfPerfsStorm)`

SetRacer sets Racer field to given value.

### HasRacer

`func (o *ApiUserActivity200ResponseInner) HasRacer() bool`

HasRacer returns a boolean if a field has been set.

### GetStreak

`func (o *ApiUserActivity200ResponseInner) GetStreak() ApiUser200ResponseAllOfPerfsStorm`

GetStreak returns the Streak field if non-nil, zero value otherwise.

### GetStreakOk

`func (o *ApiUserActivity200ResponseInner) GetStreakOk() (*ApiUser200ResponseAllOfPerfsStorm, bool)`

GetStreakOk returns a tuple with the Streak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreak

`func (o *ApiUserActivity200ResponseInner) SetStreak(v ApiUser200ResponseAllOfPerfsStorm)`

SetStreak sets Streak field to given value.

### HasStreak

`func (o *ApiUserActivity200ResponseInner) HasStreak() bool`

HasStreak returns a boolean if a field has been set.

### GetTournaments

`func (o *ApiUserActivity200ResponseInner) GetTournaments() ApiUserActivity200ResponseInnerTournaments`

GetTournaments returns the Tournaments field if non-nil, zero value otherwise.

### GetTournamentsOk

`func (o *ApiUserActivity200ResponseInner) GetTournamentsOk() (*ApiUserActivity200ResponseInnerTournaments, bool)`

GetTournamentsOk returns a tuple with the Tournaments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournaments

`func (o *ApiUserActivity200ResponseInner) SetTournaments(v ApiUserActivity200ResponseInnerTournaments)`

SetTournaments sets Tournaments field to given value.

### HasTournaments

`func (o *ApiUserActivity200ResponseInner) HasTournaments() bool`

HasTournaments returns a boolean if a field has been set.

### GetPractice

`func (o *ApiUserActivity200ResponseInner) GetPractice() []ApiUserActivity200ResponseInnerPracticeInner`

GetPractice returns the Practice field if non-nil, zero value otherwise.

### GetPracticeOk

`func (o *ApiUserActivity200ResponseInner) GetPracticeOk() (*[]ApiUserActivity200ResponseInnerPracticeInner, bool)`

GetPracticeOk returns a tuple with the Practice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPractice

`func (o *ApiUserActivity200ResponseInner) SetPractice(v []ApiUserActivity200ResponseInnerPracticeInner)`

SetPractice sets Practice field to given value.

### HasPractice

`func (o *ApiUserActivity200ResponseInner) HasPractice() bool`

HasPractice returns a boolean if a field has been set.

### GetSimuls

`func (o *ApiUserActivity200ResponseInner) GetSimuls() []string`

GetSimuls returns the Simuls field if non-nil, zero value otherwise.

### GetSimulsOk

`func (o *ApiUserActivity200ResponseInner) GetSimulsOk() (*[]string, bool)`

GetSimulsOk returns a tuple with the Simuls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimuls

`func (o *ApiUserActivity200ResponseInner) SetSimuls(v []string)`

SetSimuls sets Simuls field to given value.

### HasSimuls

`func (o *ApiUserActivity200ResponseInner) HasSimuls() bool`

HasSimuls returns a boolean if a field has been set.

### GetCorrespondenceMoves

`func (o *ApiUserActivity200ResponseInner) GetCorrespondenceMoves() ApiUserActivity200ResponseInnerCorrespondenceMoves`

GetCorrespondenceMoves returns the CorrespondenceMoves field if non-nil, zero value otherwise.

### GetCorrespondenceMovesOk

`func (o *ApiUserActivity200ResponseInner) GetCorrespondenceMovesOk() (*ApiUserActivity200ResponseInnerCorrespondenceMoves, bool)`

GetCorrespondenceMovesOk returns a tuple with the CorrespondenceMoves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondenceMoves

`func (o *ApiUserActivity200ResponseInner) SetCorrespondenceMoves(v ApiUserActivity200ResponseInnerCorrespondenceMoves)`

SetCorrespondenceMoves sets CorrespondenceMoves field to given value.

### HasCorrespondenceMoves

`func (o *ApiUserActivity200ResponseInner) HasCorrespondenceMoves() bool`

HasCorrespondenceMoves returns a boolean if a field has been set.

### GetCorrespondenceEnds

`func (o *ApiUserActivity200ResponseInner) GetCorrespondenceEnds() ApiUserActivity200ResponseInnerCorrespondenceEnds`

GetCorrespondenceEnds returns the CorrespondenceEnds field if non-nil, zero value otherwise.

### GetCorrespondenceEndsOk

`func (o *ApiUserActivity200ResponseInner) GetCorrespondenceEndsOk() (*ApiUserActivity200ResponseInnerCorrespondenceEnds, bool)`

GetCorrespondenceEndsOk returns a tuple with the CorrespondenceEnds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondenceEnds

`func (o *ApiUserActivity200ResponseInner) SetCorrespondenceEnds(v ApiUserActivity200ResponseInnerCorrespondenceEnds)`

SetCorrespondenceEnds sets CorrespondenceEnds field to given value.

### HasCorrespondenceEnds

`func (o *ApiUserActivity200ResponseInner) HasCorrespondenceEnds() bool`

HasCorrespondenceEnds returns a boolean if a field has been set.

### GetFollows

`func (o *ApiUserActivity200ResponseInner) GetFollows() ApiUserActivity200ResponseInnerFollows`

GetFollows returns the Follows field if non-nil, zero value otherwise.

### GetFollowsOk

`func (o *ApiUserActivity200ResponseInner) GetFollowsOk() (*ApiUserActivity200ResponseInnerFollows, bool)`

GetFollowsOk returns a tuple with the Follows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollows

`func (o *ApiUserActivity200ResponseInner) SetFollows(v ApiUserActivity200ResponseInnerFollows)`

SetFollows sets Follows field to given value.

### HasFollows

`func (o *ApiUserActivity200ResponseInner) HasFollows() bool`

HasFollows returns a boolean if a field has been set.

### GetStudies

`func (o *ApiUserActivity200ResponseInner) GetStudies() []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament`

GetStudies returns the Studies field if non-nil, zero value otherwise.

### GetStudiesOk

`func (o *ApiUserActivity200ResponseInner) GetStudiesOk() (*[]ApiUserActivity200ResponseInnerTournamentsBestInnerTournament, bool)`

GetStudiesOk returns a tuple with the Studies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStudies

`func (o *ApiUserActivity200ResponseInner) SetStudies(v []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament)`

SetStudies sets Studies field to given value.

### HasStudies

`func (o *ApiUserActivity200ResponseInner) HasStudies() bool`

HasStudies returns a boolean if a field has been set.

### GetTeams

`func (o *ApiUserActivity200ResponseInner) GetTeams() []ApiUserActivity200ResponseInnerTeamsInner`

GetTeams returns the Teams field if non-nil, zero value otherwise.

### GetTeamsOk

`func (o *ApiUserActivity200ResponseInner) GetTeamsOk() (*[]ApiUserActivity200ResponseInnerTeamsInner, bool)`

GetTeamsOk returns a tuple with the Teams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeams

`func (o *ApiUserActivity200ResponseInner) SetTeams(v []ApiUserActivity200ResponseInnerTeamsInner)`

SetTeams sets Teams field to given value.

### HasTeams

`func (o *ApiUserActivity200ResponseInner) HasTeams() bool`

HasTeams returns a boolean if a field has been set.

### GetPosts

`func (o *ApiUserActivity200ResponseInner) GetPosts() []ApiUserActivity200ResponseInnerPostsInner`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *ApiUserActivity200ResponseInner) GetPostsOk() (*[]ApiUserActivity200ResponseInnerPostsInner, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *ApiUserActivity200ResponseInner) SetPosts(v []ApiUserActivity200ResponseInnerPostsInner)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *ApiUserActivity200ResponseInner) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUserActivity200ResponseInner) GetPatron() ApiUserActivity200ResponseInnerPatron`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUserActivity200ResponseInner) GetPatronOk() (*ApiUserActivity200ResponseInnerPatron, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUserActivity200ResponseInner) SetPatron(v ApiUserActivity200ResponseInnerPatron)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUserActivity200ResponseInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetStream

`func (o *ApiUserActivity200ResponseInner) GetStream() bool`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *ApiUserActivity200ResponseInner) GetStreamOk() (*bool, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *ApiUserActivity200ResponseInner) SetStream(v bool)`

SetStream sets Stream field to given value.

### HasStream

`func (o *ApiUserActivity200ResponseInner) HasStream() bool`

HasStream returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


