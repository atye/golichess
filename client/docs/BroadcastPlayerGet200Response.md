# BroadcastPlayerGet200Response

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
**Fide** | Pointer to [**BroadcastPlayerGet200ResponseAllOfFide**](BroadcastPlayerGet200ResponseAllOfFide.md) |  | [optional] 
**Games** | Pointer to [**[]BroadcastPlayerGet200ResponseAllOfGamesInner**](BroadcastPlayerGet200ResponseAllOfGamesInner.md) | List of games played by the player in the broadcast tournament | [optional] 

## Methods

### NewBroadcastPlayerGet200Response

`func NewBroadcastPlayerGet200Response(name string, ) *BroadcastPlayerGet200Response`

NewBroadcastPlayerGet200Response instantiates a new BroadcastPlayerGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPlayerGet200ResponseWithDefaults

`func NewBroadcastPlayerGet200ResponseWithDefaults() *BroadcastPlayerGet200Response`

NewBroadcastPlayerGet200ResponseWithDefaults instantiates a new BroadcastPlayerGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastPlayerGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastPlayerGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastPlayerGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *BroadcastPlayerGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BroadcastPlayerGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BroadcastPlayerGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *BroadcastPlayerGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRating

`func (o *BroadcastPlayerGet200Response) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *BroadcastPlayerGet200Response) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *BroadcastPlayerGet200Response) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *BroadcastPlayerGet200Response) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetFideId

`func (o *BroadcastPlayerGet200Response) GetFideId() int32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *BroadcastPlayerGet200Response) GetFideIdOk() (*int32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *BroadcastPlayerGet200Response) SetFideId(v int32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *BroadcastPlayerGet200Response) HasFideId() bool`

HasFideId returns a boolean if a field has been set.

### GetTeam

`func (o *BroadcastPlayerGet200Response) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *BroadcastPlayerGet200Response) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *BroadcastPlayerGet200Response) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *BroadcastPlayerGet200Response) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetFed

`func (o *BroadcastPlayerGet200Response) GetFed() string`

GetFed returns the Fed field if non-nil, zero value otherwise.

### GetFedOk

`func (o *BroadcastPlayerGet200Response) GetFedOk() (*string, bool)`

GetFedOk returns a tuple with the Fed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFed

`func (o *BroadcastPlayerGet200Response) SetFed(v string)`

SetFed sets Fed field to given value.

### HasFed

`func (o *BroadcastPlayerGet200Response) HasFed() bool`

HasFed returns a boolean if a field has been set.

### GetScore

`func (o *BroadcastPlayerGet200Response) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *BroadcastPlayerGet200Response) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *BroadcastPlayerGet200Response) SetScore(v float32)`

SetScore sets Score field to given value.

### HasScore

`func (o *BroadcastPlayerGet200Response) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetPlayed

`func (o *BroadcastPlayerGet200Response) GetPlayed() int32`

GetPlayed returns the Played field if non-nil, zero value otherwise.

### GetPlayedOk

`func (o *BroadcastPlayerGet200Response) GetPlayedOk() (*int32, bool)`

GetPlayedOk returns a tuple with the Played field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayed

`func (o *BroadcastPlayerGet200Response) SetPlayed(v int32)`

SetPlayed sets Played field to given value.

### HasPlayed

`func (o *BroadcastPlayerGet200Response) HasPlayed() bool`

HasPlayed returns a boolean if a field has been set.

### GetRatingDiffs

`func (o *BroadcastPlayerGet200Response) GetRatingDiffs() BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs`

GetRatingDiffs returns the RatingDiffs field if non-nil, zero value otherwise.

### GetRatingDiffsOk

`func (o *BroadcastPlayerGet200Response) GetRatingDiffsOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs, bool)`

GetRatingDiffsOk returns a tuple with the RatingDiffs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiffs

`func (o *BroadcastPlayerGet200Response) SetRatingDiffs(v BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs)`

SetRatingDiffs sets RatingDiffs field to given value.

### HasRatingDiffs

`func (o *BroadcastPlayerGet200Response) HasRatingDiffs() bool`

HasRatingDiffs returns a boolean if a field has been set.

### GetRatingsMap

`func (o *BroadcastPlayerGet200Response) GetRatingsMap() BroadcastPlayersGet200ResponseInnerAllOfRatingsMap`

GetRatingsMap returns the RatingsMap field if non-nil, zero value otherwise.

### GetRatingsMapOk

`func (o *BroadcastPlayerGet200Response) GetRatingsMapOk() (*BroadcastPlayersGet200ResponseInnerAllOfRatingsMap, bool)`

GetRatingsMapOk returns a tuple with the RatingsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingsMap

`func (o *BroadcastPlayerGet200Response) SetRatingsMap(v BroadcastPlayersGet200ResponseInnerAllOfRatingsMap)`

SetRatingsMap sets RatingsMap field to given value.

### HasRatingsMap

`func (o *BroadcastPlayerGet200Response) HasRatingsMap() bool`

HasRatingsMap returns a boolean if a field has been set.

### GetPerformances

`func (o *BroadcastPlayerGet200Response) GetPerformances() BroadcastPlayersGet200ResponseInnerAllOfPerformances`

GetPerformances returns the Performances field if non-nil, zero value otherwise.

### GetPerformancesOk

`func (o *BroadcastPlayerGet200Response) GetPerformancesOk() (*BroadcastPlayersGet200ResponseInnerAllOfPerformances, bool)`

GetPerformancesOk returns a tuple with the Performances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformances

`func (o *BroadcastPlayerGet200Response) SetPerformances(v BroadcastPlayersGet200ResponseInnerAllOfPerformances)`

SetPerformances sets Performances field to given value.

### HasPerformances

`func (o *BroadcastPlayerGet200Response) HasPerformances() bool`

HasPerformances returns a boolean if a field has been set.

### GetTiebreaks

`func (o *BroadcastPlayerGet200Response) GetTiebreaks() []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner`

GetTiebreaks returns the Tiebreaks field if non-nil, zero value otherwise.

### GetTiebreaksOk

`func (o *BroadcastPlayerGet200Response) GetTiebreaksOk() (*[]BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner, bool)`

GetTiebreaksOk returns a tuple with the Tiebreaks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiebreaks

`func (o *BroadcastPlayerGet200Response) SetTiebreaks(v []BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner)`

SetTiebreaks sets Tiebreaks field to given value.

### HasTiebreaks

`func (o *BroadcastPlayerGet200Response) HasTiebreaks() bool`

HasTiebreaks returns a boolean if a field has been set.

### GetRank

`func (o *BroadcastPlayerGet200Response) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *BroadcastPlayerGet200Response) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *BroadcastPlayerGet200Response) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *BroadcastPlayerGet200Response) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetFide

`func (o *BroadcastPlayerGet200Response) GetFide() BroadcastPlayerGet200ResponseAllOfFide`

GetFide returns the Fide field if non-nil, zero value otherwise.

### GetFideOk

`func (o *BroadcastPlayerGet200Response) GetFideOk() (*BroadcastPlayerGet200ResponseAllOfFide, bool)`

GetFideOk returns a tuple with the Fide field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFide

`func (o *BroadcastPlayerGet200Response) SetFide(v BroadcastPlayerGet200ResponseAllOfFide)`

SetFide sets Fide field to given value.

### HasFide

`func (o *BroadcastPlayerGet200Response) HasFide() bool`

HasFide returns a boolean if a field has been set.

### GetGames

`func (o *BroadcastPlayerGet200Response) GetGames() []BroadcastPlayerGet200ResponseAllOfGamesInner`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BroadcastPlayerGet200Response) GetGamesOk() (*[]BroadcastPlayerGet200ResponseAllOfGamesInner, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BroadcastPlayerGet200Response) SetGames(v []BroadcastPlayerGet200ResponseAllOfGamesInner)`

SetGames sets Games field to given value.

### HasGames

`func (o *BroadcastPlayerGet200Response) HasGames() bool`

HasGames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


