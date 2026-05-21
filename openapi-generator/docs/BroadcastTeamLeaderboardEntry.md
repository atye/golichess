# BroadcastTeamLeaderboardEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the team | 
**Mp** | **float32** | Total match points scored | 
**Gp** | **float32** | Total game points scored | 
**AverageRating** | Pointer to **int32** | The average rating of the team&#39;s players | [optional] 
**Matches** | [**[]BroadcastTeamPOVMatchEntry**](BroadcastTeamPOVMatchEntry.md) |  | 
**Players** | [**[]BroadcastPlayerEntry**](BroadcastPlayerEntry.md) | Players who have played for the team and their overall score | 

## Methods

### NewBroadcastTeamLeaderboardEntry

`func NewBroadcastTeamLeaderboardEntry(name string, mp float32, gp float32, matches []BroadcastTeamPOVMatchEntry, players []BroadcastPlayerEntry, ) *BroadcastTeamLeaderboardEntry`

NewBroadcastTeamLeaderboardEntry instantiates a new BroadcastTeamLeaderboardEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTeamLeaderboardEntryWithDefaults

`func NewBroadcastTeamLeaderboardEntryWithDefaults() *BroadcastTeamLeaderboardEntry`

NewBroadcastTeamLeaderboardEntryWithDefaults instantiates a new BroadcastTeamLeaderboardEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastTeamLeaderboardEntry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastTeamLeaderboardEntry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastTeamLeaderboardEntry) SetName(v string)`

SetName sets Name field to given value.


### GetMp

`func (o *BroadcastTeamLeaderboardEntry) GetMp() float32`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *BroadcastTeamLeaderboardEntry) GetMpOk() (*float32, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *BroadcastTeamLeaderboardEntry) SetMp(v float32)`

SetMp sets Mp field to given value.


### GetGp

`func (o *BroadcastTeamLeaderboardEntry) GetGp() float32`

GetGp returns the Gp field if non-nil, zero value otherwise.

### GetGpOk

`func (o *BroadcastTeamLeaderboardEntry) GetGpOk() (*float32, bool)`

GetGpOk returns a tuple with the Gp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp

`func (o *BroadcastTeamLeaderboardEntry) SetGp(v float32)`

SetGp sets Gp field to given value.


### GetAverageRating

`func (o *BroadcastTeamLeaderboardEntry) GetAverageRating() int32`

GetAverageRating returns the AverageRating field if non-nil, zero value otherwise.

### GetAverageRatingOk

`func (o *BroadcastTeamLeaderboardEntry) GetAverageRatingOk() (*int32, bool)`

GetAverageRatingOk returns a tuple with the AverageRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageRating

`func (o *BroadcastTeamLeaderboardEntry) SetAverageRating(v int32)`

SetAverageRating sets AverageRating field to given value.

### HasAverageRating

`func (o *BroadcastTeamLeaderboardEntry) HasAverageRating() bool`

HasAverageRating returns a boolean if a field has been set.

### GetMatches

`func (o *BroadcastTeamLeaderboardEntry) GetMatches() []BroadcastTeamPOVMatchEntry`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *BroadcastTeamLeaderboardEntry) GetMatchesOk() (*[]BroadcastTeamPOVMatchEntry, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *BroadcastTeamLeaderboardEntry) SetMatches(v []BroadcastTeamPOVMatchEntry)`

SetMatches sets Matches field to given value.


### GetPlayers

`func (o *BroadcastTeamLeaderboardEntry) GetPlayers() []BroadcastPlayerEntry`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastTeamLeaderboardEntry) GetPlayersOk() (*[]BroadcastPlayerEntry, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastTeamLeaderboardEntry) SetPlayers(v []BroadcastPlayerEntry)`

SetPlayers sets Players field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


