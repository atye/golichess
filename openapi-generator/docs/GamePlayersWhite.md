# GamePlayersWhite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**LightUser**](LightUser.md) |  | 
**Rating** | **int32** |  | 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**Analysis** | Pointer to [**GamePlayerAiAnalysis**](GamePlayerAiAnalysis.md) |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 
**Berserk** | Pointer to **bool** | Whether the player berserked. Only present in Arena tournament games. | [optional] 
**AiLevel** | **int32** |  | 

## Methods

### NewGamePlayersWhite

`func NewGamePlayersWhite(user LightUser, rating int32, aiLevel int32, ) *GamePlayersWhite`

NewGamePlayersWhite instantiates a new GamePlayersWhite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamePlayersWhiteWithDefaults

`func NewGamePlayersWhiteWithDefaults() *GamePlayersWhite`

NewGamePlayersWhiteWithDefaults instantiates a new GamePlayersWhite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GamePlayersWhite) GetUser() LightUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GamePlayersWhite) GetUserOk() (*LightUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GamePlayersWhite) SetUser(v LightUser)`

SetUser sets User field to given value.


### GetRating

`func (o *GamePlayersWhite) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GamePlayersWhite) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GamePlayersWhite) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetRatingDiff

`func (o *GamePlayersWhite) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *GamePlayersWhite) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *GamePlayersWhite) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *GamePlayersWhite) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetName

`func (o *GamePlayersWhite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GamePlayersWhite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GamePlayersWhite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GamePlayersWhite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisional

`func (o *GamePlayersWhite) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *GamePlayersWhite) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *GamePlayersWhite) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *GamePlayersWhite) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetAnalysis

`func (o *GamePlayersWhite) GetAnalysis() GamePlayerAiAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GamePlayersWhite) GetAnalysisOk() (*GamePlayerAiAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GamePlayersWhite) SetAnalysis(v GamePlayerAiAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GamePlayersWhite) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTeam

`func (o *GamePlayersWhite) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GamePlayersWhite) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GamePlayersWhite) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GamePlayersWhite) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetBerserk

`func (o *GamePlayersWhite) GetBerserk() bool`

GetBerserk returns the Berserk field if non-nil, zero value otherwise.

### GetBerserkOk

`func (o *GamePlayersWhite) GetBerserkOk() (*bool, bool)`

GetBerserkOk returns a tuple with the Berserk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerserk

`func (o *GamePlayersWhite) SetBerserk(v bool)`

SetBerserk sets Berserk field to given value.

### HasBerserk

`func (o *GamePlayersWhite) HasBerserk() bool`

HasBerserk returns a boolean if a field has been set.

### GetAiLevel

`func (o *GamePlayersWhite) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *GamePlayersWhite) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *GamePlayersWhite) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


