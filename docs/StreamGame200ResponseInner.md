# StreamGame200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Variant** | Pointer to [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | [optional] 
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**Fen** | **string** |  | 
**Player** | Pointer to **string** |  | [optional] 
**Turns** | Pointer to **int32** |  | [optional] 
**StartedAtTurn** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**StreamGame200ResponseInnerOneOfStatus**](StreamGame200ResponseInnerOneOfStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**GamePgn200ResponseOneOfPlayers**](GamePgn200ResponseOneOfPlayers.md) |  | [optional] 
**Lm** | Pointer to **string** |  | [optional] 
**Wc** | **int32** |  | 
**Bc** | **int32** |  | 

## Methods

### NewStreamGame200ResponseInner

`func NewStreamGame200ResponseInner(id string, fen string, wc int32, bc int32, ) *StreamGame200ResponseInner`

NewStreamGame200ResponseInner instantiates a new StreamGame200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamGame200ResponseInnerWithDefaults

`func NewStreamGame200ResponseInnerWithDefaults() *StreamGame200ResponseInner`

NewStreamGame200ResponseInnerWithDefaults instantiates a new StreamGame200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamGame200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamGame200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamGame200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *StreamGame200ResponseInner) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *StreamGame200ResponseInner) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *StreamGame200ResponseInner) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *StreamGame200ResponseInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *StreamGame200ResponseInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *StreamGame200ResponseInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *StreamGame200ResponseInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *StreamGame200ResponseInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *StreamGame200ResponseInner) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *StreamGame200ResponseInner) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *StreamGame200ResponseInner) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *StreamGame200ResponseInner) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *StreamGame200ResponseInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *StreamGame200ResponseInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *StreamGame200ResponseInner) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *StreamGame200ResponseInner) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetInitialFen

`func (o *StreamGame200ResponseInner) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *StreamGame200ResponseInner) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *StreamGame200ResponseInner) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *StreamGame200ResponseInner) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetFen

`func (o *StreamGame200ResponseInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *StreamGame200ResponseInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *StreamGame200ResponseInner) SetFen(v string)`

SetFen sets Fen field to given value.


### GetPlayer

`func (o *StreamGame200ResponseInner) GetPlayer() string`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *StreamGame200ResponseInner) GetPlayerOk() (*string, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *StreamGame200ResponseInner) SetPlayer(v string)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *StreamGame200ResponseInner) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetTurns

`func (o *StreamGame200ResponseInner) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *StreamGame200ResponseInner) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *StreamGame200ResponseInner) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *StreamGame200ResponseInner) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetStartedAtTurn

`func (o *StreamGame200ResponseInner) GetStartedAtTurn() int32`

GetStartedAtTurn returns the StartedAtTurn field if non-nil, zero value otherwise.

### GetStartedAtTurnOk

`func (o *StreamGame200ResponseInner) GetStartedAtTurnOk() (*int32, bool)`

GetStartedAtTurnOk returns a tuple with the StartedAtTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtTurn

`func (o *StreamGame200ResponseInner) SetStartedAtTurn(v int32)`

SetStartedAtTurn sets StartedAtTurn field to given value.

### HasStartedAtTurn

`func (o *StreamGame200ResponseInner) HasStartedAtTurn() bool`

HasStartedAtTurn returns a boolean if a field has been set.

### GetSource

`func (o *StreamGame200ResponseInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StreamGame200ResponseInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StreamGame200ResponseInner) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StreamGame200ResponseInner) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *StreamGame200ResponseInner) GetStatus() StreamGame200ResponseInnerOneOfStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StreamGame200ResponseInner) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StreamGame200ResponseInner) SetStatus(v StreamGame200ResponseInnerOneOfStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StreamGame200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *StreamGame200ResponseInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *StreamGame200ResponseInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *StreamGame200ResponseInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *StreamGame200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastMove

`func (o *StreamGame200ResponseInner) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *StreamGame200ResponseInner) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *StreamGame200ResponseInner) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *StreamGame200ResponseInner) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetPlayers

`func (o *StreamGame200ResponseInner) GetPlayers() GamePgn200ResponseOneOfPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *StreamGame200ResponseInner) GetPlayersOk() (*GamePgn200ResponseOneOfPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *StreamGame200ResponseInner) SetPlayers(v GamePgn200ResponseOneOfPlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *StreamGame200ResponseInner) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetLm

`func (o *StreamGame200ResponseInner) GetLm() string`

GetLm returns the Lm field if non-nil, zero value otherwise.

### GetLmOk

`func (o *StreamGame200ResponseInner) GetLmOk() (*string, bool)`

GetLmOk returns a tuple with the Lm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLm

`func (o *StreamGame200ResponseInner) SetLm(v string)`

SetLm sets Lm field to given value.

### HasLm

`func (o *StreamGame200ResponseInner) HasLm() bool`

HasLm returns a boolean if a field has been set.

### GetWc

`func (o *StreamGame200ResponseInner) GetWc() int32`

GetWc returns the Wc field if non-nil, zero value otherwise.

### GetWcOk

`func (o *StreamGame200ResponseInner) GetWcOk() (*int32, bool)`

GetWcOk returns a tuple with the Wc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWc

`func (o *StreamGame200ResponseInner) SetWc(v int32)`

SetWc sets Wc field to given value.


### GetBc

`func (o *StreamGame200ResponseInner) GetBc() int32`

GetBc returns the Bc field if non-nil, zero value otherwise.

### GetBcOk

`func (o *StreamGame200ResponseInner) GetBcOk() (*int32, bool)`

GetBcOk returns a tuple with the Bc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBc

`func (o *StreamGame200ResponseInner) SetBc(v int32)`

SetBc sets Bc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


