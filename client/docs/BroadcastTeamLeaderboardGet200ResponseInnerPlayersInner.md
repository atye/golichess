# BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**FideId** | Pointer to **int32** |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Fed** | Pointer to **string** |  | [optional] 
**Score** | Pointer to **float32** |  | [optional] 
**Played** | Pointer to **int32** |  | [optional] 
**RatingDiffs** | Pointer to [**BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs**](BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs.md) |  | [optional] 
**RatingsMap** | Pointer to [**BroadcastPlayersGet200ResponseInnerAllOfRatingsMap**](BroadcastPlayersGet200ResponseInnerAllOfRatingsMap.md) |  | [optional] 
**Performances** | Pointer to [**BroadcastPlayersGet200ResponseInnerAllOfPerformances**](BroadcastPlayersGet200ResponseInnerAllOfPerformances.md) |  | [optional] 
**Tiebreaks** | Pointer to [**[]BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner**](BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner.md) |  | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 

## Methods

### NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInner

`func NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInner(name string, ) *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner`

NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInner instantiates a new BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInnerWithDefaults

`func NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInnerWithDefaults() *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner`

NewBroadcastTeamLeaderboardGet200ResponseInnerPlayersInnerWithDefaults instantiates a new BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFideId

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetFideId() int32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetFideIdOk() (*int32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetFideId(v int32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasFideId() bool`

HasFideId returns a boolean if a field has been set.

### GetTeam

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetFed() string`

GetFed returns the Fed field if non-nil, zero value otherwise.

### GetFedOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetFedOk() (*string, bool)`

GetFedOk returns a tuple with the Fed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetFed(v string)`

SetFed sets Fed field to given value.

### HasFed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasFed() bool`

HasFed returns a boolean if a field has been set.

### GetScore

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPlayed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetPlayed() int32`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetPlayedOk() (*int32, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetPlayed(v int32)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRatingDiffs

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRatingDiffs() BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs`

GetRatingDiffs returns the RatingDiffs field if non-nil, zero value otherwise.

### GetRatingDiffsOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRatingDiffsOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs, bool)`

GetRatingDiffsOk returns a tuple with the RatingDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiffs

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetRatingDiffs(v BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs)`

SetRatingDiffs sets RatingDiffs field to given value.

### HasRatingDiffs

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasRatingDiffs() bool`

HasRatingDiffs returns a boolean if a field has been set.

### GetRatingsMap

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRatingsMap() BroadcastPlayersGet200ResponseInnerAllOfRatingsMap`

GetRatingsMap returns the RatingsMap field if non-nil, zero value otherwise.

### GetRatingsMapOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRatingsMapOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingsMap, bool)`

GetRatingsMapOk returns a tuple with the RatingsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingsMap

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetRatingsMap(v BroadcastPlayersGet200ResponseInnerAllOfRatingsMap)`

SetRatingsMap sets RatingsMap field to given value.

### HasRatingsMap

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasRatingsMap() bool`

HasRatingsMap returns a boolean if a field has been set.

### GetPerformances

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetPerformances() BroadcastPlayersGet200ResponseInnerAllOfPerformances`

GetPerformances returns the Performances field if non-nil, zero value otherwise.

### GetPerformancesOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetPerformancesOk() (*BroadcastPlayersGet200ResponseInnerAllOfPerformances, bool)`

GetPerformancesOk returns a tuple with the Performances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformances

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetPerformances(v BroadcastPlayersGet200ResponseInnerAllOfPerformances)`

SetPerformances sets Performances field to given value.

### HasPerformances

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasPerformances() bool`

HasPerformances returns a boolean if a field has been set.

### GetTiebreaks

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTiebreaks() []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner`

GetTiebreaks returns the Tiebreaks field if non-nil, zero value otherwise.

### GetTiebreaksOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetTiebreaksOk() (*[]BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner, bool)`

GetTiebreaksOk returns a tuple with the Tiebreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreaks

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetTiebreaks(v []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner)`

SetTiebreaks sets Tiebreaks field to given value.

### HasTiebreaks

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasTiebreaks() bool`

HasTiebreaks returns a boolean if a field has been set.

### GetRank

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


