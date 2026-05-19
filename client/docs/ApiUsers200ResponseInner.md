# ApiUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Perfs** | Pointer to [**ApiUser200ResponseAllOfPerfs**](ApiUser200ResponseAllOfPerfs.md) |  | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Disabled** | Pointer to **bool** | only appears if a user&#39;s account is closed | [optional] 
**TosViolation** | Pointer to **bool** | only appears if a user&#39;s account is marked for the violation of [Lichess TOS](https://lichess.org/terms-of-service) | [optional] 
**Profile** | Pointer to [**ApiUser200ResponseAllOfProfile**](ApiUser200ResponseAllOfProfile.md) |  | [optional] 
**SeenAt** | Pointer to **int64** |  | [optional] 
**PlayTime** | Pointer to [**ApiUser200ResponseAllOfPlayTime**](ApiUser200ResponseAllOfPlayTime.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Verified** | Pointer to **bool** |  | [optional] 

## Methods

### NewApiUsers200ResponseInner

`func NewApiUsers200ResponseInner(id string, username string, ) *ApiUsers200ResponseInner`

NewApiUsers200ResponseInner instantiates a new ApiUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsers200ResponseInnerWithDefaults

`func NewApiUsers200ResponseInnerWithDefaults() *ApiUsers200ResponseInner`

NewApiUsers200ResponseInnerWithDefaults instantiates a new ApiUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUsers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUsers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUsers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *ApiUsers200ResponseInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUsers200ResponseInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUsers200ResponseInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *ApiUsers200ResponseInner) GetPerfs() ApiUser200ResponseAllOfPerfs`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *ApiUsers200ResponseInner) GetPerfsOk() (*ApiUser200ResponseAllOfPerfs, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *ApiUsers200ResponseInner) SetPerfs(v ApiUser200ResponseAllOfPerfs)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *ApiUsers200ResponseInner) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *ApiUsers200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUsers200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUsers200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUsers200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *ApiUsers200ResponseInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiUsers200ResponseInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiUsers200ResponseInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiUsers200ResponseInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiUsers200ResponseInner) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiUsers200ResponseInner) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiUsers200ResponseInner) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiUsers200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *ApiUsers200ResponseInner) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApiUsers200ResponseInner) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApiUsers200ResponseInner) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApiUsers200ResponseInner) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTosViolation

`func (o *ApiUsers200ResponseInner) GetTosViolation() bool`

GetTosViolation returns the TosViolation field if non-nil, zero value otherwise.

### GetTosViolationOk

`func (o *ApiUsers200ResponseInner) GetTosViolationOk() (*bool, bool)`

GetTosViolationOk returns a tuple with the TosViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosViolation

`func (o *ApiUsers200ResponseInner) SetTosViolation(v bool)`

SetTosViolation sets TosViolation field to given value.

### HasTosViolation

`func (o *ApiUsers200ResponseInner) HasTosViolation() bool`

HasTosViolation returns a boolean if a field has been set.

### GetProfile

`func (o *ApiUsers200ResponseInner) GetProfile() ApiUser200ResponseAllOfProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApiUsers200ResponseInner) GetProfileOk() (*ApiUser200ResponseAllOfProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApiUsers200ResponseInner) SetProfile(v ApiUser200ResponseAllOfProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ApiUsers200ResponseInner) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeenAt

`func (o *ApiUsers200ResponseInner) GetSeenAt() int64`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *ApiUsers200ResponseInner) GetSeenAtOk() (*int64, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *ApiUsers200ResponseInner) SetSeenAt(v int64)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *ApiUsers200ResponseInner) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### GetPlayTime

`func (o *ApiUsers200ResponseInner) GetPlayTime() ApiUser200ResponseAllOfPlayTime`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *ApiUsers200ResponseInner) GetPlayTimeOk() (*ApiUser200ResponseAllOfPlayTime, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *ApiUsers200ResponseInner) SetPlayTime(v ApiUser200ResponseAllOfPlayTime)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *ApiUsers200ResponseInner) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUsers200ResponseInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUsers200ResponseInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUsers200ResponseInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUsers200ResponseInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiUsers200ResponseInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiUsers200ResponseInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiUsers200ResponseInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiUsers200ResponseInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetVerified

`func (o *ApiUsers200ResponseInner) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ApiUsers200ResponseInner) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ApiUsers200ResponseInner) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ApiUsers200ResponseInner) HasVerified() bool`

HasVerified returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


