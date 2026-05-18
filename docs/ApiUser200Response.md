# ApiUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Perfs** | Pointer to [**ApiUser200ResponseAllOfPerfs**](ApiUser200ResponseAllOfPerfs.md) |  | [optional] 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
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
**Url** | **string** |  | 
**Playing** | Pointer to **string** |  | [optional] 
**Count** | Pointer to [**ApiUser200ResponseAllOfCount**](ApiUser200ResponseAllOfCount.md) |  | [optional] 
**Streaming** | Pointer to **bool** |  | [optional] 
**Streamer** | Pointer to [**ApiUser200ResponseAllOfStreamer**](ApiUser200ResponseAllOfStreamer.md) |  | [optional] 
**Followable** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**Following** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**Blocking** | Pointer to **bool** | only appears if the request is [authenticated with OAuth2](#description/authentication) | [optional] 
**FideId** | Pointer to **float32** |  | [optional] 

## Methods

### NewApiUser200Response

`func NewApiUser200Response(id string, username string, url string, ) *ApiUser200Response`

NewApiUser200Response instantiates a new ApiUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUser200ResponseWithDefaults

`func NewApiUser200ResponseWithDefaults() *ApiUser200Response`

NewApiUser200ResponseWithDefaults instantiates a new ApiUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUser200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUser200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUser200Response) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *ApiUser200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ApiUser200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ApiUser200Response) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *ApiUser200Response) GetPerfs() ApiUser200ResponseAllOfPerfs`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *ApiUser200Response) GetPerfsOk() (*ApiUser200ResponseAllOfPerfs, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *ApiUser200Response) SetPerfs(v ApiUser200ResponseAllOfPerfs)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *ApiUser200Response) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *ApiUser200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUser200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUser200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUser200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ApiUser200Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ApiUser200Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFlair

`func (o *ApiUser200Response) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiUser200Response) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiUser200Response) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiUser200Response) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ApiUser200Response) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ApiUser200Response) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ApiUser200Response) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ApiUser200Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetDisabled

`func (o *ApiUser200Response) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ApiUser200Response) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ApiUser200Response) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ApiUser200Response) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTosViolation

`func (o *ApiUser200Response) GetTosViolation() bool`

GetTosViolation returns the TosViolation field if non-nil, zero value otherwise.

### GetTosViolationOk

`func (o *ApiUser200Response) GetTosViolationOk() (*bool, bool)`

GetTosViolationOk returns a tuple with the TosViolation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTosViolation

`func (o *ApiUser200Response) SetTosViolation(v bool)`

SetTosViolation sets TosViolation field to given value.

### HasTosViolation

`func (o *ApiUser200Response) HasTosViolation() bool`

HasTosViolation returns a boolean if a field has been set.

### GetProfile

`func (o *ApiUser200Response) GetProfile() ApiUser200ResponseAllOfProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ApiUser200Response) GetProfileOk() (*ApiUser200ResponseAllOfProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ApiUser200Response) SetProfile(v ApiUser200ResponseAllOfProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ApiUser200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetSeenAt

`func (o *ApiUser200Response) GetSeenAt() int64`

GetSeenAt returns the SeenAt field if non-nil, zero value otherwise.

### GetSeenAtOk

`func (o *ApiUser200Response) GetSeenAtOk() (*int64, bool)`

GetSeenAtOk returns a tuple with the SeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeenAt

`func (o *ApiUser200Response) SetSeenAt(v int64)`

SetSeenAt sets SeenAt field to given value.

### HasSeenAt

`func (o *ApiUser200Response) HasSeenAt() bool`

HasSeenAt returns a boolean if a field has been set.

### GetPlayTime

`func (o *ApiUser200Response) GetPlayTime() ApiUser200ResponseAllOfPlayTime`

GetPlayTime returns the PlayTime field if non-nil, zero value otherwise.

### GetPlayTimeOk

`func (o *ApiUser200Response) GetPlayTimeOk() (*ApiUser200ResponseAllOfPlayTime, bool)`

GetPlayTimeOk returns a tuple with the PlayTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayTime

`func (o *ApiUser200Response) SetPlayTime(v ApiUser200ResponseAllOfPlayTime)`

SetPlayTime sets PlayTime field to given value.

### HasPlayTime

`func (o *ApiUser200Response) HasPlayTime() bool`

HasPlayTime returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUser200Response) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUser200Response) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUser200Response) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUser200Response) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiUser200Response) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiUser200Response) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiUser200Response) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiUser200Response) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetVerified

`func (o *ApiUser200Response) GetVerified() bool`

GetVerified returns the Verified field if non-nil, zero value otherwise.

### GetVerifiedOk

`func (o *ApiUser200Response) GetVerifiedOk() (*bool, bool)`

GetVerifiedOk returns a tuple with the Verified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerified

`func (o *ApiUser200Response) SetVerified(v bool)`

SetVerified sets Verified field to given value.

### HasVerified

`func (o *ApiUser200Response) HasVerified() bool`

HasVerified returns a boolean if a field has been set.

### GetUrl

`func (o *ApiUser200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiUser200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiUser200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPlaying

`func (o *ApiUser200Response) GetPlaying() string`

GetPlaying returns the Playing field if non-nil, zero value otherwise.

### GetPlayingOk

`func (o *ApiUser200Response) GetPlayingOk() (*string, bool)`

GetPlayingOk returns a tuple with the Playing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaying

`func (o *ApiUser200Response) SetPlaying(v string)`

SetPlaying sets Playing field to given value.

### HasPlaying

`func (o *ApiUser200Response) HasPlaying() bool`

HasPlaying returns a boolean if a field has been set.

### GetCount

`func (o *ApiUser200Response) GetCount() ApiUser200ResponseAllOfCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiUser200Response) GetCountOk() (*ApiUser200ResponseAllOfCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiUser200Response) SetCount(v ApiUser200ResponseAllOfCount)`

SetCount sets Count field to given value.

### HasCount

`func (o *ApiUser200Response) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetStreaming

`func (o *ApiUser200Response) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *ApiUser200Response) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *ApiUser200Response) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *ApiUser200Response) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetStreamer

`func (o *ApiUser200Response) GetStreamer() ApiUser200ResponseAllOfStreamer`

GetStreamer returns the Streamer field if non-nil, zero value otherwise.

### GetStreamerOk

`func (o *ApiUser200Response) GetStreamerOk() (*ApiUser200ResponseAllOfStreamer, bool)`

GetStreamerOk returns a tuple with the Streamer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamer

`func (o *ApiUser200Response) SetStreamer(v ApiUser200ResponseAllOfStreamer)`

SetStreamer sets Streamer field to given value.

### HasStreamer

`func (o *ApiUser200Response) HasStreamer() bool`

HasStreamer returns a boolean if a field has been set.

### GetFollowable

`func (o *ApiUser200Response) GetFollowable() bool`

GetFollowable returns the Followable field if non-nil, zero value otherwise.

### GetFollowableOk

`func (o *ApiUser200Response) GetFollowableOk() (*bool, bool)`

GetFollowableOk returns a tuple with the Followable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowable

`func (o *ApiUser200Response) SetFollowable(v bool)`

SetFollowable sets Followable field to given value.

### HasFollowable

`func (o *ApiUser200Response) HasFollowable() bool`

HasFollowable returns a boolean if a field has been set.

### GetFollowing

`func (o *ApiUser200Response) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *ApiUser200Response) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *ApiUser200Response) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *ApiUser200Response) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetBlocking

`func (o *ApiUser200Response) GetBlocking() bool`

GetBlocking returns the Blocking field if non-nil, zero value otherwise.

### GetBlockingOk

`func (o *ApiUser200Response) GetBlockingOk() (*bool, bool)`

GetBlockingOk returns a tuple with the Blocking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocking

`func (o *ApiUser200Response) SetBlocking(v bool)`

SetBlocking sets Blocking field to given value.

### HasBlocking

`func (o *ApiUser200Response) HasBlocking() bool`

HasBlocking returns a boolean if a field has been set.

### GetFideId

`func (o *ApiUser200Response) GetFideId() float32`

GetFideId returns the FideId field if non-nil, zero value otherwise.

### GetFideIdOk

`func (o *ApiUser200Response) GetFideIdOk() (*float32, bool)`

GetFideIdOk returns a tuple with the FideId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideId

`func (o *ApiUser200Response) SetFideId(v float32)`

SetFideId sets FideId field to given value.

### HasFideId

`func (o *ApiUser200Response) HasFideId() bool`

HasFideId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


