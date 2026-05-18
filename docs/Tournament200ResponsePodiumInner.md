# Tournament200ResponsePodiumInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Rank** | Pointer to **int32** |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**Nb** | Pointer to [**ApiTournamentPost200ResponsePodiumInnerNb**](ApiTournamentPost200ResponsePodiumInnerNb.md) |  | [optional] 
**Performance** | Pointer to **int32** |  | [optional] 

## Methods

### NewTournament200ResponsePodiumInner

`func NewTournament200ResponsePodiumInner() *Tournament200ResponsePodiumInner`

NewTournament200ResponsePodiumInner instantiates a new Tournament200ResponsePodiumInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTournament200ResponsePodiumInnerWithDefaults

`func NewTournament200ResponsePodiumInnerWithDefaults() *Tournament200ResponsePodiumInner`

NewTournament200ResponsePodiumInnerWithDefaults instantiates a new Tournament200ResponsePodiumInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Tournament200ResponsePodiumInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tournament200ResponsePodiumInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tournament200ResponsePodiumInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tournament200ResponsePodiumInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTitle

`func (o *Tournament200ResponsePodiumInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Tournament200ResponsePodiumInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Tournament200ResponsePodiumInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Tournament200ResponsePodiumInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *Tournament200ResponsePodiumInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *Tournament200ResponsePodiumInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *Tournament200ResponsePodiumInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *Tournament200ResponsePodiumInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *Tournament200ResponsePodiumInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *Tournament200ResponsePodiumInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *Tournament200ResponsePodiumInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *Tournament200ResponsePodiumInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetFlair

`func (o *Tournament200ResponsePodiumInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *Tournament200ResponsePodiumInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *Tournament200ResponsePodiumInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *Tournament200ResponsePodiumInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetRank

`func (o *Tournament200ResponsePodiumInner) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Tournament200ResponsePodiumInner) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Tournament200ResponsePodiumInner) SetRank(v int32)`

SetRank sets Rank field to given value.

### HasRank

`func (o *Tournament200ResponsePodiumInner) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetRating

`func (o *Tournament200ResponsePodiumInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *Tournament200ResponsePodiumInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *Tournament200ResponsePodiumInner) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *Tournament200ResponsePodiumInner) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetScore

`func (o *Tournament200ResponsePodiumInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *Tournament200ResponsePodiumInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *Tournament200ResponsePodiumInner) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *Tournament200ResponsePodiumInner) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetNb

`func (o *Tournament200ResponsePodiumInner) GetNb() ApiTournamentPost200ResponsePodiumInnerNb`

GetNb returns the Nb field if non-nil, zero value otherwise.

### GetNbOk

`func (o *Tournament200ResponsePodiumInner) GetNbOk() (*ApiTournamentPost200ResponsePodiumInnerNb, bool)`

GetNbOk returns a tuple with the Nb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNb

`func (o *Tournament200ResponsePodiumInner) SetNb(v ApiTournamentPost200ResponsePodiumInnerNb)`

SetNb sets Nb field to given value.

### HasNb

`func (o *Tournament200ResponsePodiumInner) HasNb() bool`

HasNb returns a boolean if a field has been set.

### GetPerformance

`func (o *Tournament200ResponsePodiumInner) GetPerformance() int32`

GetPerformance returns the Performance field if non-nil, zero value otherwise.

### GetPerformanceOk

`func (o *Tournament200ResponsePodiumInner) GetPerformanceOk() (*int32, bool)`

GetPerformanceOk returns a tuple with the Performance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformance

`func (o *Tournament200ResponsePodiumInner) SetPerformance(v int32)`

SetPerformance sets Performance field to given value.

### HasPerformance

`func (o *Tournament200ResponsePodiumInner) HasPerformance() bool`

HasPerformance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


