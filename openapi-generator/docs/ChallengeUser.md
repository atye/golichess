# ChallengeUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Rating** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Lag** | Pointer to **int32** |  | [optional] 

## Methods

### NewChallengeUser

`func NewChallengeUser(id string, name string, ) *ChallengeUser`

NewChallengeUser instantiates a new ChallengeUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeUserWithDefaults

`func NewChallengeUserWithDefaults() *ChallengeUser`

NewChallengeUserWithDefaults instantiates a new ChallengeUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ChallengeUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ChallengeUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ChallengeUser) SetName(v string)`

SetName sets Name field to given value.


### GetRating

`func (o *ChallengeUser) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ChallengeUser) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ChallengeUser) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *ChallengeUser) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetTitle

`func (o *ChallengeUser) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ChallengeUser) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ChallengeUser) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ChallengeUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *ChallengeUser) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ChallengeUser) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ChallengeUser) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ChallengeUser) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *ChallengeUser) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ChallengeUser) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ChallengeUser) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ChallengeUser) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ChallengeUser) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ChallengeUser) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ChallengeUser) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ChallengeUser) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetProvisional

`func (o *ChallengeUser) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *ChallengeUser) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *ChallengeUser) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *ChallengeUser) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.

### GetOnline

`func (o *ChallengeUser) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ChallengeUser) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ChallengeUser) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ChallengeUser) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetLag

`func (o *ChallengeUser) GetLag() int32`

GetLag returns the Lag field if non-nil, zero value otherwise.

### GetLagOk

`func (o *ChallengeUser) GetLagOk() (*int32, bool)`

GetLagOk returns a tuple with the Lag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLag

`func (o *ChallengeUser) SetLag(v int32)`

SetLag sets Lag field to given value.

### HasLag

`func (o *ChallengeUser) HasLag() bool`

HasLag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


