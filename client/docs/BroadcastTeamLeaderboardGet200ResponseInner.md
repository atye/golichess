# BroadcastTeamLeaderboardGet200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team | 
**Mp** | **float32** | Total match points scored | 
**Gp** | **float32** | Total game points scored | 
**AverageRating** | Pointer to **int32** | The average rating of the team&#39;s players | [optional] 
**Matches** | [**[]BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner**](BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner.md) |  | 
**Players** | [**[]BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner**](BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner.md) | Players who have played for the team and their overall score | 

## Methods

### NewBroadcastTeamLeaderboardGet200ResponseInner

`func NewBroadcastTeamLeaderboardGet200ResponseInner(name string, mp float32, gp float32, matches []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner, players []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner, ) *BroadcastTeamLeaderboardGet200ResponseInner`

NewBroadcastTeamLeaderboardGet200ResponseInner instantiates a new BroadcastTeamLeaderboardGet200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTeamLeaderboardGet200ResponseInnerWithDefaults

`func NewBroadcastTeamLeaderboardGet200ResponseInnerWithDefaults() *BroadcastTeamLeaderboardGet200ResponseInner`

NewBroadcastTeamLeaderboardGet200ResponseInnerWithDefaults instantiates a new BroadcastTeamLeaderboardGet200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetMp

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMp() float32`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMpOk() (*float32, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetMp(v float32)`

SetMp sets Mp field to given value.


### GetGp

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetGp() float32`

GetGp returns the Gp field if non-nil, zero value otherwise.

### GetGpOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetGpOk() (*float32, bool)`

GetGpOk returns a tuple with the Gp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetGp(v float32)`

SetGp sets Gp field to given value.


### GetAverageRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetAverageRating() int32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetAverageRatingOk() (*int32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetAverageRating(v int32)`

SetAverageRating sets AverageRating field to given value.

### HasAverageRating

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) HasAverageRating() bool`

HasAverageRating returns a boolean if a field has been set.

### GetMatches

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMatches() []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMatchesOk() (*[]BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetMatches(v []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner)`

SetMatches sets Matches field to given value.


### GetPlayers

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetPlayers() []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetPlayersOk() (*[]BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetPlayers(v []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner)`

SetPlayers sets Players field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


