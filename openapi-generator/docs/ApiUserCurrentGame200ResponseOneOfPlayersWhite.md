# ApiUserCurrentGame200ResponseOneOfPlayersWhite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser**](ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser.md) |  | 
**Rating** | **int32** |  | 
**RatingDiff** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**AiLevel** | Pointer to **int32** |  | [optional] 
**Analysis** | Pointer to [**GamePgn200ResponseOneOfPlayersWhiteAnalysis**](GamePgn200ResponseOneOfPlayersWhiteAnalysis.md) |  | [optional] 
**Team** | Pointer to **string** |  | [optional] 

## Methods

### NewApiUserCurrentGame200ResponseOneOfPlayersWhite

`func NewApiUserCurrentGame200ResponseOneOfPlayersWhite(user ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser, rating int32, ) *ApiUserCurrentGame200ResponseOneOfPlayersWhite`

NewApiUserCurrentGame200ResponseOneOfPlayersWhite instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserCurrentGame200ResponseOneOfPlayersWhiteWithDefaults

`func NewApiUserCurrentGame200ResponseOneOfPlayersWhiteWithDefaults() *ApiUserCurrentGame200ResponseOneOfPlayersWhite`

NewApiUserCurrentGame200ResponseOneOfPlayersWhiteWithDefaults instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetUser() ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetUserOk() (*ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetUser(v ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser)`

SetUser sets User field to given value.


### GetRating

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetRatingDiff

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingDiff() int32`

GetRatingDiff returns the RatingDiff field if non-nil, zero value otherwise.

### GetRatingDiffOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingDiffOk() (*int32, bool)`

GetRatingDiffOk returns a tuple with the RatingDiff field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatingDiff

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetRatingDiff(v int32)`

SetRatingDiff sets RatingDiff field to given value.

### HasRatingDiff

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasRatingDiff() bool`

HasRatingDiff returns a boolean if a field has been set.

### GetName

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProvisional

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetAiLevel

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.

### HasAiLevel

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasAiLevel() bool`

HasAiLevel returns a boolean if a field has been set.

### GetAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAnalysis() GamePgn200ResponseOneOfPlayersWhiteAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAnalysisOk() (*GamePgn200ResponseOneOfPlayersWhiteAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetAnalysis(v GamePgn200ResponseOneOfPlayersWhiteAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.

### GetTeam

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetTeam() string`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetTeamOk() (*string, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetTeam(v string)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasTeam() bool`

HasTeam returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


