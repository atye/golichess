# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Perfs** | Pointer to [**Perfs**](Perfs.md) |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Disabled** | Pointer to **bool** | only appears if a user&#39;s account is closed | [optional] 
**TosViolation** | Pointer to **bool** | only appears if a user&#39;s account is marked for the violation of [Lichess TOS](https://lichess.org/terms-of-service) | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 
**SeenAt** | Pointer to **int64** |  | [optional] 
**PlayTime** | Pointer to [**PlayTime**](PlayTime.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewUser

`func NewUser(id string, username string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *User) GetPerfs() Perfs`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *User) GetPerfsOk() (*Perfs, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *User) SetPerfs(v Perfs)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *User) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *User) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *User) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *User) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *User) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *User) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *User) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *User) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *User) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetCreatedAt

`func (o *User) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *User) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *User) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *User) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *User) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *User) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *User) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *User) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTosViolation

`func (o *User) GetTosViolation() bool`

GetTosViolation returns the TosViolation field if non-nil, zero value otherwise.

### GetTosViolationOk

`func (o *User) GetTosViolationOk() (*bool, bool)`

GetTosViolationOk returns a tuple with the TosViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosViolation

`func (o *User) SetTosViolation(v bool)`

SetTosViolation sets TosViolation field to given value.

### HasTosViolation

`func (o *User) HasTosViolation() bool`

HasTosViolation returns a boolean if a field has been set.

### GetProfile

`func (o *User) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *User) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *User) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *User) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeenAt

`func (o *User) GetSeenAt() int64`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *User) GetSeenAtOk() (*int64, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *User) SetSeenAt(v int64)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *User) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### GetPlayTime

`func (o *User) GetPlayTime() PlayTime`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *User) GetPlayTimeOk() (*PlayTime, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *User) SetPlayTime(v PlayTime)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *User) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetPatron

`func (o *User) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *User) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *User) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *User) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *User) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *User) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *User) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *User) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetVerified

`func (o *User) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *User) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *User) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *User) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


