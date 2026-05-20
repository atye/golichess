# ChallengeAi201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**Fen** | Pointer to **string** |  | [optional] 
**Turns** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StreamGame200ResponseInnerOneOfStatus**](StreamGame200ResponseInnerOneOfStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **int64** |  | [optional] 
**Player** | Pointer to **NullableString** |  | [optional] 
**FullId** | Pointer to **string** |  | [optional] 

## Methods

### NewChallengeAi201Response

`func NewChallengeAi201Response() *ChallengeAi201Response`

NewChallengeAi201Response instantiates a new ChallengeAi201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeAi201ResponseWithDefaults

`func NewChallengeAi201ResponseWithDefaults() *ChallengeAi201Response`

NewChallengeAi201ResponseWithDefaults instantiates a new ChallengeAi201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeAi201Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeAi201Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeAi201Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ChallengeAi201Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetVariant

`func (o *ChallengeAi201Response) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeAi201Response) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeAi201Response) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ChallengeAi201Response) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *ChallengeAi201Response) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeAi201Response) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeAi201Response) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ChallengeAi201Response) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeAi201Response) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeAi201Response) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeAi201Response) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ChallengeAi201Response) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *ChallengeAi201Response) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeAi201Response) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeAi201Response) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ChallengeAi201Response) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetFen

`func (o *ChallengeAi201Response) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ChallengeAi201Response) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ChallengeAi201Response) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *ChallengeAi201Response) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetTurns

`func (o *ChallengeAi201Response) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *ChallengeAi201Response) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *ChallengeAi201Response) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *ChallengeAi201Response) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetSource

`func (o *ChallengeAi201Response) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ChallengeAi201Response) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ChallengeAi201Response) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *ChallengeAi201Response) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *ChallengeAi201Response) GetStatus() StreamGame200ResponseInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeAi201Response) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeAi201Response) SetStatus(v StreamGame200ResponseInnerOneOfStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ChallengeAi201Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ChallengeAi201Response) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ChallengeAi201Response) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ChallengeAi201Response) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ChallengeAi201Response) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPlayer

`func (o *ChallengeAi201Response) GetPlayer() string`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *ChallengeAi201Response) GetPlayerOk() (*string, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *ChallengeAi201Response) SetPlayer(v string)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *ChallengeAi201Response) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### SetPlayerNil

`func (o *ChallengeAi201Response) SetPlayerNil(b bool)`

 SetPlayerNil sets the value for Player to be an explicit nil

### UnsetPlayer
`func (o *ChallengeAi201Response) UnsetPlayer()`

UnsetPlayer ensures that no value is present for Player, not even an explicit nil
### GetFullId

`func (o *ChallengeAi201Response) GetFullId() string`

GetFullId returns the FullId field if non-nil, zero value otherwise.

### GetFullIdOk

`func (o *ChallengeAi201Response) GetFullIdOk() (*string, bool)`

GetFullIdOk returns a tuple with the FullId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullId

`func (o *ChallengeAi201Response) SetFullId(v string)`

SetFullId sets FullId field to given value.

### HasFullId

`func (o *ChallengeAi201Response) HasFullId() bool`

HasFullId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


