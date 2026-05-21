# BroadcastPlayerEntry

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

## Methods

### NewBroadcastPlayerEntry

`func NewBroadcastPlayerEntry(name string, ) *BroadcastPlayerEntry`

NewBroadcastPlayerEntry instantiates a new BroadcastPlayerEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPlayerEntryWithDefaults

`func NewBroadcastPlayerEntryWithDefaults() *BroadcastPlayerEntry`

NewBroadcastPlayerEntryWithDefaults instantiates a new BroadcastPlayerEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastPlayerEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastPlayerEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastPlayerEntry) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BroadcastPlayerEntry) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastPlayerEntry) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastPlayerEntry) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastPlayerEntry) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRating

`func (o *BroadcastPlayerEntry) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BroadcastPlayerEntry) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BroadcastPlayerEntry) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BroadcastPlayerEntry) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFideId

`func (o *BroadcastPlayerEntry) GetFideId() int32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *BroadcastPlayerEntry) GetFideIdOk() (*int32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *BroadcastPlayerEntry) SetFideId(v int32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *BroadcastPlayerEntry) HasFideId() bool`

HasFideId returns a boolean if a field has been set.

### GetTeam

`func (o *BroadcastPlayerEntry) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BroadcastPlayerEntry) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BroadcastPlayerEntry) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BroadcastPlayerEntry) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFed

`func (o *BroadcastPlayerEntry) GetFed() string`

GetFed returns the Fed field if non-nil, zero value otherwise.

### GetFedOk

`func (o *BroadcastPlayerEntry) GetFedOk() (*string, bool)`

GetFedOk returns a tuple with the Fed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFed

`func (o *BroadcastPlayerEntry) SetFed(v string)`

SetFed sets Fed field to given value.

### HasFed

`func (o *BroadcastPlayerEntry) HasFed() bool`

HasFed returns a boolean if a field has been set.

### GetScore

`func (o *BroadcastPlayerEntry) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BroadcastPlayerEntry) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BroadcastPlayerEntry) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BroadcastPlayerEntry) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPlayed

`func (o *BroadcastPlayerEntry) GetPlayed() int32`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BroadcastPlayerEntry) GetPlayedOk() (*int32, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BroadcastPlayerEntry) SetPlayed(v int32)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BroadcastPlayerEntry) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRatingDiffs

`func (o *BroadcastPlayerEntry) GetRatingDiffs() StatByFideTC`

GetRatingDiffs returns the RatingDiffs field if non-nil, zero value otherwise.

### GetRatingDiffsOk

`func (o *BroadcastPlayerEntry) GetRatingDiffsOk() (*StatByFideTC, bool)`

GetRatingDiffsOk returns a tuple with the RatingDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiffs

`func (o *BroadcastPlayerEntry) SetRatingDiffs(v StatByFideTC)`

SetRatingDiffs sets RatingDiffs field to given value.

### HasRatingDiffs

`func (o *BroadcastPlayerEntry) HasRatingDiffs() bool`

HasRatingDiffs returns a boolean if a field has been set.

### GetRatingsMap

`func (o *BroadcastPlayerEntry) GetRatingsMap() StatByFideTC`

GetRatingsMap returns the RatingsMap field if non-nil, zero value otherwise.

### GetRatingsMapOk

`func (o *BroadcastPlayerEntry) GetRatingsMapOk() (*StatByFideTC, bool)`

GetRatingsMapOk returns a tuple with the RatingsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingsMap

`func (o *BroadcastPlayerEntry) SetRatingsMap(v StatByFideTC)`

SetRatingsMap sets RatingsMap field to given value.

### HasRatingsMap

`func (o *BroadcastPlayerEntry) HasRatingsMap() bool`

HasRatingsMap returns a boolean if a field has been set.

### GetPerformances

`func (o *BroadcastPlayerEntry) GetPerformances() StatByFideTC`

GetPerformances returns the Performances field if non-nil, zero value otherwise.

### GetPerformancesOk

`func (o *BroadcastPlayerEntry) GetPerformancesOk() (*StatByFideTC, bool)`

GetPerformancesOk returns a tuple with the Performances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformances

`func (o *BroadcastPlayerEntry) SetPerformances(v StatByFideTC)`

SetPerformances sets Performances field to given value.

### HasPerformances

`func (o *BroadcastPlayerEntry) HasPerformances() bool`

HasPerformances returns a boolean if a field has been set.

### GetTiebreaks

`func (o *BroadcastPlayerEntry) GetTiebreaks() []BroadcastPlayerTiebreak`

GetTiebreaks returns the Tiebreaks field if non-nil, zero value otherwise.

### GetTiebreaksOk

`func (o *BroadcastPlayerEntry) GetTiebreaksOk() (*[]BroadcastPlayerTiebreak, bool)`

GetTiebreaksOk returns a tuple with the Tiebreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreaks

`func (o *BroadcastPlayerEntry) SetTiebreaks(v []BroadcastPlayerTiebreak)`

SetTiebreaks sets Tiebreaks field to given value.

### HasTiebreaks

`func (o *BroadcastPlayerEntry) HasTiebreaks() bool`

HasTiebreaks returns a boolean if a field has been set.

### GetRank

`func (o *BroadcastPlayerEntry) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BroadcastPlayerEntry) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BroadcastPlayerEntry) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BroadcastPlayerEntry) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


