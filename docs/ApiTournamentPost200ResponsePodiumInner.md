# ApiTournamentPost200ResponsePodiumInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Nb** | Pointer to [**ApiTournamentPost200ResponsePodiumInnerNb**](ApiTournamentPost200ResponsePodiumInnerNb.md) |  | [optional] 
**Performance** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiTournamentPost200ResponsePodiumInner

`func NewApiTournamentPost200ResponsePodiumInner() *ApiTournamentPost200ResponsePodiumInner`

NewApiTournamentPost200ResponsePodiumInner instantiates a new ApiTournamentPost200ResponsePodiumInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTournamentPost200ResponsePodiumInnerWithDefaults

`func NewApiTournamentPost200ResponsePodiumInnerWithDefaults() *ApiTournamentPost200ResponsePodiumInner`

NewApiTournamentPost200ResponsePodiumInnerWithDefaults instantiates a new ApiTournamentPost200ResponsePodiumInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiTournamentPost200ResponsePodiumInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiTournamentPost200ResponsePodiumInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiTournamentPost200ResponsePodiumInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *ApiTournamentPost200ResponsePodiumInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiTournamentPost200ResponsePodiumInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiTournamentPost200ResponsePodiumInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ApiTournamentPost200ResponsePodiumInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ApiTournamentPost200ResponsePodiumInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPatron

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiTournamentPost200ResponsePodiumInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiTournamentPost200ResponsePodiumInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiTournamentPost200ResponsePodiumInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiTournamentPost200ResponsePodiumInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetFlair

`func (o *ApiTournamentPost200ResponsePodiumInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiTournamentPost200ResponsePodiumInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiTournamentPost200ResponsePodiumInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetRank

`func (o *ApiTournamentPost200ResponsePodiumInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiTournamentPost200ResponsePodiumInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *ApiTournamentPost200ResponsePodiumInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRating

`func (o *ApiTournamentPost200ResponsePodiumInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiTournamentPost200ResponsePodiumInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiTournamentPost200ResponsePodiumInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScore

`func (o *ApiTournamentPost200ResponsePodiumInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *ApiTournamentPost200ResponsePodiumInner) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *ApiTournamentPost200ResponsePodiumInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetNb

`func (o *ApiTournamentPost200ResponsePodiumInner) GetNb() ApiTournamentPost200ResponsePodiumInnerNb`

GetNb returns the Nb field if non-nil, zero value otherwise.

### GetNbOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetNbOk() (*ApiTournamentPost200ResponsePodiumInnerNb, bool)`

GetNbOk returns a tuple with the Nb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNb

`func (o *ApiTournamentPost200ResponsePodiumInner) SetNb(v ApiTournamentPost200ResponsePodiumInnerNb)`

SetNb sets Nb field to given value.

### HasNb

`func (o *ApiTournamentPost200ResponsePodiumInner) HasNb() bool`

HasNb returns a boolean if a field has been set.

### GetPerformance

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *ApiTournamentPost200ResponsePodiumInner) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *ApiTournamentPost200ResponsePodiumInner) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *ApiTournamentPost200ResponsePodiumInner) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


