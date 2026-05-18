# ApiStreamEvent200ResponseOneOf2ChallengeChallenger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Rating** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Lag** | Pointer to **int32** |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOf2ChallengeChallenger

`func NewApiStreamEvent200ResponseOneOf2ChallengeChallenger(id string, name string, ) *ApiStreamEvent200ResponseOneOf2ChallengeChallenger`

NewApiStreamEvent200ResponseOneOf2ChallengeChallenger instantiates a new ApiStreamEvent200ResponseOneOf2ChallengeChallenger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOf2ChallengeChallengerWithDefaults

`func NewApiStreamEvent200ResponseOneOf2ChallengeChallengerWithDefaults() *ApiStreamEvent200ResponseOneOf2ChallengeChallenger`

NewApiStreamEvent200ResponseOneOf2ChallengeChallengerWithDefaults instantiates a new ApiStreamEvent200ResponseOneOf2ChallengeChallenger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetName(v string)`

SetName sets Name field to given value.


### GetRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *ApiStreamEvent200ResponseOneOf2ChallengeChallenger) HasLag() bool`

HasLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


