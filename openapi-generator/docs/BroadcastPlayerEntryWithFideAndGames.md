# BroadcastPlayerEntryWithFideAndGames

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**FideId** | Pointer to **int32** |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Fed** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Played** | Pointer to **int32** |  | [optional] 
**RatingDiffs** | Pointer to [**StatByFideTC**](StatByFideTC.md) | Rating differences by FIDE time control.  | [optional] 
**RatingsMap** | Pointer to [**StatByFideTC**](StatByFideTC.md) | Player&#39;s ratings at the time of the tournament.  | [optional] 
**Performances** | Pointer to [**StatByFideTC**](StatByFideTC.md) | Performance ratings by FIDE time control.  | [optional] 
**Tiebreaks** | Pointer to [**[]BroadcastPlayerTiebreak**](BroadcastPlayerTiebreak.md) |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Fide** | Pointer to [**BroadcastPlayerEntryWithFideAndGamesAllOfFide**](BroadcastPlayerEntryWithFideAndGamesAllOfFide.md) |  | [optional] 
**Games** | Pointer to [**[]BroadcastGameEntry**](BroadcastGameEntry.md) | List of games played by the player in the broadcast tournament | [optional] 

## Methods

### NewBroadcastPlayerEntryWithFideAndGames

`func NewBroadcastPlayerEntryWithFideAndGames(name string, ) *BroadcastPlayerEntryWithFideAndGames`

NewBroadcastPlayerEntryWithFideAndGames instantiates a new BroadcastPlayerEntryWithFideAndGames object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPlayerEntryWithFideAndGamesWithDefaults

`func NewBroadcastPlayerEntryWithFideAndGamesWithDefaults() *BroadcastPlayerEntryWithFideAndGames`

NewBroadcastPlayerEntryWithFideAndGamesWithDefaults instantiates a new BroadcastPlayerEntryWithFideAndGames object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastPlayerEntryWithFideAndGames) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastPlayerEntryWithFideAndGames) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastPlayerEntryWithFideAndGames) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastPlayerEntryWithFideAndGames) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRating

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BroadcastPlayerEntryWithFideAndGames) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BroadcastPlayerEntryWithFideAndGames) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFideId

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFideId() int32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFideIdOk() (*int32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *BroadcastPlayerEntryWithFideAndGames) SetFideId(v int32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *BroadcastPlayerEntryWithFideAndGames) HasFideId() bool`

HasFideId returns a boolean if a field has been set.

### GetTeam

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BroadcastPlayerEntryWithFideAndGames) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BroadcastPlayerEntryWithFideAndGames) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFed

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFed() string`

GetFed returns the Fed field if non-nil, zero value otherwise.

### GetFedOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFedOk() (*string, bool)`

GetFedOk returns a tuple with the Fed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFed

`func (o *BroadcastPlayerEntryWithFideAndGames) SetFed(v string)`

SetFed sets Fed field to given value.

### HasFed

`func (o *BroadcastPlayerEntryWithFideAndGames) HasFed() bool`

HasFed returns a boolean if a field has been set.

### GetScore

`func (o *BroadcastPlayerEntryWithFideAndGames) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BroadcastPlayerEntryWithFideAndGames) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BroadcastPlayerEntryWithFideAndGames) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPlayed

`func (o *BroadcastPlayerEntryWithFideAndGames) GetPlayed() int32`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetPlayedOk() (*int32, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BroadcastPlayerEntryWithFideAndGames) SetPlayed(v int32)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BroadcastPlayerEntryWithFideAndGames) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRatingDiffs

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRatingDiffs() StatByFideTC`

GetRatingDiffs returns the RatingDiffs field if non-nil, zero value otherwise.

### GetRatingDiffsOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRatingDiffsOk() (*StatByFideTC, bool)`

GetRatingDiffsOk returns a tuple with the RatingDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiffs

`func (o *BroadcastPlayerEntryWithFideAndGames) SetRatingDiffs(v StatByFideTC)`

SetRatingDiffs sets RatingDiffs field to given value.

### HasRatingDiffs

`func (o *BroadcastPlayerEntryWithFideAndGames) HasRatingDiffs() bool`

HasRatingDiffs returns a boolean if a field has been set.

### GetRatingsMap

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRatingsMap() StatByFideTC`

GetRatingsMap returns the RatingsMap field if non-nil, zero value otherwise.

### GetRatingsMapOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRatingsMapOk() (*StatByFideTC, bool)`

GetRatingsMapOk returns a tuple with the RatingsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingsMap

`func (o *BroadcastPlayerEntryWithFideAndGames) SetRatingsMap(v StatByFideTC)`

SetRatingsMap sets RatingsMap field to given value.

### HasRatingsMap

`func (o *BroadcastPlayerEntryWithFideAndGames) HasRatingsMap() bool`

HasRatingsMap returns a boolean if a field has been set.

### GetPerformances

`func (o *BroadcastPlayerEntryWithFideAndGames) GetPerformances() StatByFideTC`

GetPerformances returns the Performances field if non-nil, zero value otherwise.

### GetPerformancesOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetPerformancesOk() (*StatByFideTC, bool)`

GetPerformancesOk returns a tuple with the Performances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformances

`func (o *BroadcastPlayerEntryWithFideAndGames) SetPerformances(v StatByFideTC)`

SetPerformances sets Performances field to given value.

### HasPerformances

`func (o *BroadcastPlayerEntryWithFideAndGames) HasPerformances() bool`

HasPerformances returns a boolean if a field has been set.

### GetTiebreaks

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTiebreaks() []BroadcastPlayerTiebreak`

GetTiebreaks returns the Tiebreaks field if non-nil, zero value otherwise.

### GetTiebreaksOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetTiebreaksOk() (*[]BroadcastPlayerTiebreak, bool)`

GetTiebreaksOk returns a tuple with the Tiebreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreaks

`func (o *BroadcastPlayerEntryWithFideAndGames) SetTiebreaks(v []BroadcastPlayerTiebreak)`

SetTiebreaks sets Tiebreaks field to given value.

### HasTiebreaks

`func (o *BroadcastPlayerEntryWithFideAndGames) HasTiebreaks() bool`

HasTiebreaks returns a boolean if a field has been set.

### GetRank

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BroadcastPlayerEntryWithFideAndGames) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BroadcastPlayerEntryWithFideAndGames) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFide

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFide() BroadcastPlayerEntryWithFideAndGamesAllOfFide`

GetFide returns the Fide field if non-nil, zero value otherwise.

### GetFideOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetFideOk() (*BroadcastPlayerEntryWithFideAndGamesAllOfFide, bool)`

GetFideOk returns a tuple with the Fide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFide

`func (o *BroadcastPlayerEntryWithFideAndGames) SetFide(v BroadcastPlayerEntryWithFideAndGamesAllOfFide)`

SetFide sets Fide field to given value.

### HasFide

`func (o *BroadcastPlayerEntryWithFideAndGames) HasFide() bool`

HasFide returns a boolean if a field has been set.

### GetGames

`func (o *BroadcastPlayerEntryWithFideAndGames) GetGames() []BroadcastGameEntry`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BroadcastPlayerEntryWithFideAndGames) GetGamesOk() (*[]BroadcastGameEntry, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BroadcastPlayerEntryWithFideAndGames) SetGames(v []BroadcastGameEntry)`

SetGames sets Games field to given value.

### HasGames

`func (o *BroadcastPlayerEntryWithFideAndGames) HasGames() bool`

HasGames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


