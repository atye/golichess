# GamePlayerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**LightUser**](LightUser.md) |  | 
**Rating** | **int32** |  | 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**AiLevel** | Pointer to **int32** |  | [optional] 
**Analysis** | Pointer to [**GamePlayerUserAnalysis**](GamePlayerUserAnalysis.md) |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 

## Methods

### NewGamePlayerUser

`func NewGamePlayerUser(user LightUser, rating int32, ) *GamePlayerUser`

NewGamePlayerUser instantiates a new GamePlayerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamePlayerUserWithDefaults

`func NewGamePlayerUserWithDefaults() *GamePlayerUser`

NewGamePlayerUserWithDefaults instantiates a new GamePlayerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *GamePlayerUser) GetUser() LightUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GamePlayerUser) GetUserOk() (*LightUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GamePlayerUser) SetUser(v LightUser)`

SetUser sets User field to given value.


### GetRating

`func (o *GamePlayerUser) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GamePlayerUser) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GamePlayerUser) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetRatingDiff

`func (o *GamePlayerUser) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *GamePlayerUser) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *GamePlayerUser) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *GamePlayerUser) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetName

`func (o *GamePlayerUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GamePlayerUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GamePlayerUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GamePlayerUser) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisional

`func (o *GamePlayerUser) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *GamePlayerUser) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *GamePlayerUser) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *GamePlayerUser) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetAiLevel

`func (o *GamePlayerUser) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *GamePlayerUser) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *GamePlayerUser) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.

### HasAiLevel

`func (o *GamePlayerUser) HasAiLevel() bool`

HasAiLevel returns a boolean if a field has been set.

### GetAnalysis

`func (o *GamePlayerUser) GetAnalysis() GamePlayerUserAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GamePlayerUser) GetAnalysisOk() (*GamePlayerUserAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GamePlayerUser) SetAnalysis(v GamePlayerUserAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GamePlayerUser) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTeam

`func (o *GamePlayerUser) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *GamePlayerUser) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *GamePlayerUser) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *GamePlayerUser) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


