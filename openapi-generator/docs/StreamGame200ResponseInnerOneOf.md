# StreamGame200ResponseInnerOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Variant** | Pointer to [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**Fen** | Pointer to **string** |  | [optional] 
**Player** | Pointer to **NullableString** |  | [optional] 
**Turns** | Pointer to **int32** |  | [optional] 
**StartedAtTurn** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StreamGame200ResponseInnerOneOfStatus**](StreamGame200ResponseInnerOneOfStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**GamePgn200ResponseOneOfPlayers**](GamePgn200ResponseOneOfPlayers.md) |  | [optional] 

## Methods

### NewStreamGame200ResponseInnerOneOf

`func NewStreamGame200ResponseInnerOneOf(id string, ) *StreamGame200ResponseInnerOneOf`

NewStreamGame200ResponseInnerOneOf instantiates a new StreamGame200ResponseInnerOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGame200ResponseInnerOneOfWithDefaults

`func NewStreamGame200ResponseInnerOneOfWithDefaults() *StreamGame200ResponseInnerOneOf`

NewStreamGame200ResponseInnerOneOfWithDefaults instantiates a new StreamGame200ResponseInnerOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamGame200ResponseInnerOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamGame200ResponseInnerOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamGame200ResponseInnerOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *StreamGame200ResponseInnerOneOf) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *StreamGame200ResponseInnerOneOf) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *StreamGame200ResponseInnerOneOf) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *StreamGame200ResponseInnerOneOf) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *StreamGame200ResponseInnerOneOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StreamGame200ResponseInnerOneOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StreamGame200ResponseInnerOneOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StreamGame200ResponseInnerOneOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *StreamGame200ResponseInnerOneOf) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *StreamGame200ResponseInnerOneOf) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *StreamGame200ResponseInnerOneOf) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *StreamGame200ResponseInnerOneOf) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *StreamGame200ResponseInnerOneOf) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *StreamGame200ResponseInnerOneOf) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *StreamGame200ResponseInnerOneOf) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *StreamGame200ResponseInnerOneOf) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetInitialFen

`func (o *StreamGame200ResponseInnerOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *StreamGame200ResponseInnerOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *StreamGame200ResponseInnerOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *StreamGame200ResponseInnerOneOf) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetFen

`func (o *StreamGame200ResponseInnerOneOf) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *StreamGame200ResponseInnerOneOf) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *StreamGame200ResponseInnerOneOf) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *StreamGame200ResponseInnerOneOf) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetPlayer

`func (o *StreamGame200ResponseInnerOneOf) GetPlayer() string`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *StreamGame200ResponseInnerOneOf) GetPlayerOk() (*string, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *StreamGame200ResponseInnerOneOf) SetPlayer(v string)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *StreamGame200ResponseInnerOneOf) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### SetPlayerNil

`func (o *StreamGame200ResponseInnerOneOf) SetPlayerNil(b bool)`

 SetPlayerNil sets the value for Player to be an explicit nil

### UnsetPlayer
`func (o *StreamGame200ResponseInnerOneOf) UnsetPlayer()`

UnsetPlayer ensures that no value is present for Player, not even an explicit nil
### GetTurns

`func (o *StreamGame200ResponseInnerOneOf) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *StreamGame200ResponseInnerOneOf) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *StreamGame200ResponseInnerOneOf) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *StreamGame200ResponseInnerOneOf) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetStartedAtTurn

`func (o *StreamGame200ResponseInnerOneOf) GetStartedAtTurn() int32`

GetStartedAtTurn returns the StartedAtTurn field if non-nil, zero value otherwise.

### GetStartedAtTurnOk

`func (o *StreamGame200ResponseInnerOneOf) GetStartedAtTurnOk() (*int32, bool)`

GetStartedAtTurnOk returns a tuple with the StartedAtTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtTurn

`func (o *StreamGame200ResponseInnerOneOf) SetStartedAtTurn(v int32)`

SetStartedAtTurn sets StartedAtTurn field to given value.

### HasStartedAtTurn

`func (o *StreamGame200ResponseInnerOneOf) HasStartedAtTurn() bool`

HasStartedAtTurn returns a boolean if a field has been set.

### GetSource

`func (o *StreamGame200ResponseInnerOneOf) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StreamGame200ResponseInnerOneOf) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StreamGame200ResponseInnerOneOf) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StreamGame200ResponseInnerOneOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *StreamGame200ResponseInnerOneOf) GetStatus() StreamGame200ResponseInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamGame200ResponseInnerOneOf) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamGame200ResponseInnerOneOf) SetStatus(v StreamGame200ResponseInnerOneOfStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StreamGame200ResponseInnerOneOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StreamGame200ResponseInnerOneOf) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StreamGame200ResponseInnerOneOf) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StreamGame200ResponseInnerOneOf) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StreamGame200ResponseInnerOneOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastMove

`func (o *StreamGame200ResponseInnerOneOf) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *StreamGame200ResponseInnerOneOf) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *StreamGame200ResponseInnerOneOf) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *StreamGame200ResponseInnerOneOf) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetPlayers

`func (o *StreamGame200ResponseInnerOneOf) GetPlayers() GamePgn200ResponseOneOfPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *StreamGame200ResponseInnerOneOf) GetPlayersOk() (*GamePgn200ResponseOneOfPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *StreamGame200ResponseInnerOneOf) SetPlayers(v GamePgn200ResponseOneOfPlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *StreamGame200ResponseInnerOneOf) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


