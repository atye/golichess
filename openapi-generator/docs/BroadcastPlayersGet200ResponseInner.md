# BroadcastPlayersGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
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

### NewBroadcastPlayersGet200ResponseInner

`func NewBroadcastPlayersGet200ResponseInner(name string, ) *BroadcastPlayersGet200ResponseInner`

NewBroadcastPlayersGet200ResponseInner instantiates a new BroadcastPlayersGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPlayersGet200ResponseInnerWithDefaults

`func NewBroadcastPlayersGet200ResponseInnerWithDefaults() *BroadcastPlayersGet200ResponseInner`

NewBroadcastPlayersGet200ResponseInnerWithDefaults instantiates a new BroadcastPlayersGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastPlayersGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastPlayersGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastPlayersGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BroadcastPlayersGet200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastPlayersGet200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastPlayersGet200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastPlayersGet200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *BroadcastPlayersGet200ResponseInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *BroadcastPlayersGet200ResponseInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRating

`func (o *BroadcastPlayersGet200ResponseInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BroadcastPlayersGet200ResponseInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BroadcastPlayersGet200ResponseInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BroadcastPlayersGet200ResponseInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFideId

`func (o *BroadcastPlayersGet200ResponseInner) GetFideId() int32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *BroadcastPlayersGet200ResponseInner) GetFideIdOk() (*int32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *BroadcastPlayersGet200ResponseInner) SetFideId(v int32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *BroadcastPlayersGet200ResponseInner) HasFideId() bool`

HasFideId returns a boolean if a field has been set.

### GetTeam

`func (o *BroadcastPlayersGet200ResponseInner) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BroadcastPlayersGet200ResponseInner) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BroadcastPlayersGet200ResponseInner) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BroadcastPlayersGet200ResponseInner) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFed

`func (o *BroadcastPlayersGet200ResponseInner) GetFed() string`

GetFed returns the Fed field if non-nil, zero value otherwise.

### GetFedOk

`func (o *BroadcastPlayersGet200ResponseInner) GetFedOk() (*string, bool)`

GetFedOk returns a tuple with the Fed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFed

`func (o *BroadcastPlayersGet200ResponseInner) SetFed(v string)`

SetFed sets Fed field to given value.

### HasFed

`func (o *BroadcastPlayersGet200ResponseInner) HasFed() bool`

HasFed returns a boolean if a field has been set.

### GetScore

`func (o *BroadcastPlayersGet200ResponseInner) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BroadcastPlayersGet200ResponseInner) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BroadcastPlayersGet200ResponseInner) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BroadcastPlayersGet200ResponseInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPlayed

`func (o *BroadcastPlayersGet200ResponseInner) GetPlayed() int32`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BroadcastPlayersGet200ResponseInner) GetPlayedOk() (*int32, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BroadcastPlayersGet200ResponseInner) SetPlayed(v int32)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BroadcastPlayersGet200ResponseInner) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRatingDiffs

`func (o *BroadcastPlayersGet200ResponseInner) GetRatingDiffs() BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs`

GetRatingDiffs returns the RatingDiffs field if non-nil, zero value otherwise.

### GetRatingDiffsOk

`func (o *BroadcastPlayersGet200ResponseInner) GetRatingDiffsOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs, bool)`

GetRatingDiffsOk returns a tuple with the RatingDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiffs

`func (o *BroadcastPlayersGet200ResponseInner) SetRatingDiffs(v BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs)`

SetRatingDiffs sets RatingDiffs field to given value.

### HasRatingDiffs

`func (o *BroadcastPlayersGet200ResponseInner) HasRatingDiffs() bool`

HasRatingDiffs returns a boolean if a field has been set.

### GetRatingsMap

`func (o *BroadcastPlayersGet200ResponseInner) GetRatingsMap() BroadcastPlayersGet200ResponseInnerAllOfRatingsMap`

GetRatingsMap returns the RatingsMap field if non-nil, zero value otherwise.

### GetRatingsMapOk

`func (o *BroadcastPlayersGet200ResponseInner) GetRatingsMapOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingsMap, bool)`

GetRatingsMapOk returns a tuple with the RatingsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingsMap

`func (o *BroadcastPlayersGet200ResponseInner) SetRatingsMap(v BroadcastPlayersGet200ResponseInnerAllOfRatingsMap)`

SetRatingsMap sets RatingsMap field to given value.

### HasRatingsMap

`func (o *BroadcastPlayersGet200ResponseInner) HasRatingsMap() bool`

HasRatingsMap returns a boolean if a field has been set.

### GetPerformances

`func (o *BroadcastPlayersGet200ResponseInner) GetPerformances() BroadcastPlayersGet200ResponseInnerAllOfPerformances`

GetPerformances returns the Performances field if non-nil, zero value otherwise.

### GetPerformancesOk

`func (o *BroadcastPlayersGet200ResponseInner) GetPerformancesOk() (*BroadcastPlayersGet200ResponseInnerAllOfPerformances, bool)`

GetPerformancesOk returns a tuple with the Performances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformances

`func (o *BroadcastPlayersGet200ResponseInner) SetPerformances(v BroadcastPlayersGet200ResponseInnerAllOfPerformances)`

SetPerformances sets Performances field to given value.

### HasPerformances

`func (o *BroadcastPlayersGet200ResponseInner) HasPerformances() bool`

HasPerformances returns a boolean if a field has been set.

### GetTiebreaks

`func (o *BroadcastPlayersGet200ResponseInner) GetTiebreaks() []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner`

GetTiebreaks returns the Tiebreaks field if non-nil, zero value otherwise.

### GetTiebreaksOk

`func (o *BroadcastPlayersGet200ResponseInner) GetTiebreaksOk() (*[]BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner, bool)`

GetTiebreaksOk returns a tuple with the Tiebreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreaks

`func (o *BroadcastPlayersGet200ResponseInner) SetTiebreaks(v []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner)`

SetTiebreaks sets Tiebreaks field to given value.

### HasTiebreaks

`func (o *BroadcastPlayersGet200ResponseInner) HasTiebreaks() bool`

HasTiebreaks returns a boolean if a field has been set.

### GetRank

`func (o *BroadcastPlayersGet200ResponseInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BroadcastPlayersGet200ResponseInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BroadcastPlayersGet200ResponseInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BroadcastPlayersGet200ResponseInner) HasRank() bool`

HasRank returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


