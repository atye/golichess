# UserExtended

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
**Url** | **string** |  | 
**Playing** | Pointer to **string** |  | [optional] 
**Count** | Pointer to [**Count**](Count.md) |  | [optional] 
**Streaming** | Pointer to **bool** |  | [optional] 
**Streamer** | Pointer to [**UserStreamer**](UserStreamer.md) |  | [optional] 
**Followable** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**Following** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**Blocking** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**FideId** | Pointer to **float32** |  | [optional] 

## Methods

### NewUserExtended

`func NewUserExtended(id string, username string, url string, ) *UserExtended`

NewUserExtended instantiates a new UserExtended object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserExtendedWithDefaults

`func NewUserExtendedWithDefaults() *UserExtended`

NewUserExtendedWithDefaults instantiates a new UserExtended object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserExtended) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserExtended) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserExtended) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *UserExtended) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserExtended) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserExtended) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *UserExtended) GetPerfs() Perfs`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *UserExtended) GetPerfsOk() (*Perfs, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *UserExtended) SetPerfs(v Perfs)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *UserExtended) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *UserExtended) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UserExtended) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UserExtended) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UserExtended) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *UserExtended) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *UserExtended) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *UserExtended) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *UserExtended) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UserExtended) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UserExtended) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UserExtended) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UserExtended) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *UserExtended) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *UserExtended) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *UserExtended) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *UserExtended) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTosViolation

`func (o *UserExtended) GetTosViolation() bool`

GetTosViolation returns the TosViolation field if non-nil, zero value otherwise.

### GetTosViolationOk

`func (o *UserExtended) GetTosViolationOk() (*bool, bool)`

GetTosViolationOk returns a tuple with the TosViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosViolation

`func (o *UserExtended) SetTosViolation(v bool)`

SetTosViolation sets TosViolation field to given value.

### HasTosViolation

`func (o *UserExtended) HasTosViolation() bool`

HasTosViolation returns a boolean if a field has been set.

### GetProfile

`func (o *UserExtended) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UserExtended) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UserExtended) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UserExtended) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeenAt

`func (o *UserExtended) GetSeenAt() int64`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *UserExtended) GetSeenAtOk() (*int64, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *UserExtended) SetSeenAt(v int64)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *UserExtended) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### GetPlayTime

`func (o *UserExtended) GetPlayTime() PlayTime`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *UserExtended) GetPlayTimeOk() (*PlayTime, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *UserExtended) SetPlayTime(v PlayTime)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *UserExtended) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetPatron

`func (o *UserExtended) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *UserExtended) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *UserExtended) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *UserExtended) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *UserExtended) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *UserExtended) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *UserExtended) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *UserExtended) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetVerified

`func (o *UserExtended) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *UserExtended) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *UserExtended) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *UserExtended) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUrl

`func (o *UserExtended) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserExtended) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserExtended) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPlaying

`func (o *UserExtended) GetPlaying() string`

GetPlaying returns the Playing field if non-nil, zero value otherwise.

### GetPlayingOk

`func (o *UserExtended) GetPlayingOk() (*string, bool)`

GetPlayingOk returns a tuple with the Playing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaying

`func (o *UserExtended) SetPlaying(v string)`

SetPlaying sets Playing field to given value.

### HasPlaying

`func (o *UserExtended) HasPlaying() bool`

HasPlaying returns a boolean if a field has been set.

### GetCount

`func (o *UserExtended) GetCount() Count`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *UserExtended) GetCountOk() (*Count, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *UserExtended) SetCount(v Count)`

SetCount sets Count field to given value.

### HasCount

`func (o *UserExtended) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStreaming

`func (o *UserExtended) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *UserExtended) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *UserExtended) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *UserExtended) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetStreamer

`func (o *UserExtended) GetStreamer() UserStreamer`

GetStreamer returns the Streamer field if non-nil, zero value otherwise.

### GetStreamerOk

`func (o *UserExtended) GetStreamerOk() (*UserStreamer, bool)`

GetStreamerOk returns a tuple with the Streamer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamer

`func (o *UserExtended) SetStreamer(v UserStreamer)`

SetStreamer sets Streamer field to given value.

### HasStreamer

`func (o *UserExtended) HasStreamer() bool`

HasStreamer returns a boolean if a field has been set.

### GetFollowable

`func (o *UserExtended) GetFollowable() bool`

GetFollowable returns the Followable field if non-nil, zero value otherwise.

### GetFollowableOk

`func (o *UserExtended) GetFollowableOk() (*bool, bool)`

GetFollowableOk returns a tuple with the Followable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowable

`func (o *UserExtended) SetFollowable(v bool)`

SetFollowable sets Followable field to given value.

### HasFollowable

`func (o *UserExtended) HasFollowable() bool`

HasFollowable returns a boolean if a field has been set.

### GetFollowing

`func (o *UserExtended) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *UserExtended) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *UserExtended) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *UserExtended) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetBlocking

`func (o *UserExtended) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *UserExtended) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *UserExtended) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *UserExtended) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetFideId

`func (o *UserExtended) GetFideId() float32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *UserExtended) GetFideIdOk() (*float32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *UserExtended) SetFideId(v float32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *UserExtended) HasFideId() bool`

HasFideId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


