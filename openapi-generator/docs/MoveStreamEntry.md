# MoveStreamEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Variant** | Pointer to [**Variant**](Variant.md) |  | [optional] 
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**Perf** | Pointer to [**PerfType**](PerfType.md) |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**Fen** | **string** |  | 
**Player** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**Turns** | Pointer to **int32** |  | [optional] 
**StartedAtTurn** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to [**GameSource**](GameSource.md) |  | [optional] 
**Status** | Pointer to [**GameStatus**](GameStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**GamePlayers**](GamePlayers.md) |  | [optional] 
**Lm** | Pointer to **string** |  | [optional] 
**Wc** | **int32** |  | 
**Bc** | **int32** |  | 

## Methods

### NewMoveStreamEntry

`func NewMoveStreamEntry(id string, fen string, wc int32, bc int32, ) *MoveStreamEntry`

NewMoveStreamEntry instantiates a new MoveStreamEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveStreamEntryWithDefaults

`func NewMoveStreamEntryWithDefaults() *MoveStreamEntry`

NewMoveStreamEntryWithDefaults instantiates a new MoveStreamEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveStreamEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveStreamEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveStreamEntry) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *MoveStreamEntry) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MoveStreamEntry) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MoveStreamEntry) SetVariant(v Variant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MoveStreamEntry) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *MoveStreamEntry) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MoveStreamEntry) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MoveStreamEntry) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MoveStreamEntry) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *MoveStreamEntry) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *MoveStreamEntry) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *MoveStreamEntry) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *MoveStreamEntry) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *MoveStreamEntry) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *MoveStreamEntry) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *MoveStreamEntry) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *MoveStreamEntry) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetInitialFen

`func (o *MoveStreamEntry) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *MoveStreamEntry) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *MoveStreamEntry) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *MoveStreamEntry) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetFen

`func (o *MoveStreamEntry) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *MoveStreamEntry) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *MoveStreamEntry) SetFen(v string)`

SetFen sets Fen field to given value.


### GetPlayer

`func (o *MoveStreamEntry) GetPlayer() GameColor`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *MoveStreamEntry) GetPlayerOk() (*GameColor, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *MoveStreamEntry) SetPlayer(v GameColor)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *MoveStreamEntry) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetTurns

`func (o *MoveStreamEntry) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *MoveStreamEntry) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *MoveStreamEntry) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *MoveStreamEntry) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetStartedAtTurn

`func (o *MoveStreamEntry) GetStartedAtTurn() int32`

GetStartedAtTurn returns the StartedAtTurn field if non-nil, zero value otherwise.

### GetStartedAtTurnOk

`func (o *MoveStreamEntry) GetStartedAtTurnOk() (*int32, bool)`

GetStartedAtTurnOk returns a tuple with the StartedAtTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtTurn

`func (o *MoveStreamEntry) SetStartedAtTurn(v int32)`

SetStartedAtTurn sets StartedAtTurn field to given value.

### HasStartedAtTurn

`func (o *MoveStreamEntry) HasStartedAtTurn() bool`

HasStartedAtTurn returns a boolean if a field has been set.

### GetSource

`func (o *MoveStreamEntry) GetSource() GameSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MoveStreamEntry) GetSourceOk() (*GameSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MoveStreamEntry) SetSource(v GameSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MoveStreamEntry) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *MoveStreamEntry) GetStatus() GameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MoveStreamEntry) GetStatusOk() (*GameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MoveStreamEntry) SetStatus(v GameStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MoveStreamEntry) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveStreamEntry) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveStreamEntry) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveStreamEntry) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveStreamEntry) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastMove

`func (o *MoveStreamEntry) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *MoveStreamEntry) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *MoveStreamEntry) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *MoveStreamEntry) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetPlayers

`func (o *MoveStreamEntry) GetPlayers() GamePlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *MoveStreamEntry) GetPlayersOk() (*GamePlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *MoveStreamEntry) SetPlayers(v GamePlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *MoveStreamEntry) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetLm

`func (o *MoveStreamEntry) GetLm() string`

GetLm returns the Lm field if non-nil, zero value otherwise.

### GetLmOk

`func (o *MoveStreamEntry) GetLmOk() (*string, bool)`

GetLmOk returns a tuple with the Lm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLm

`func (o *MoveStreamEntry) SetLm(v string)`

SetLm sets Lm field to given value.

### HasLm

`func (o *MoveStreamEntry) HasLm() bool`

HasLm returns a boolean if a field has been set.

### GetWc

`func (o *MoveStreamEntry) GetWc() int32`

GetWc returns the Wc field if non-nil, zero value otherwise.

### GetWcOk

`func (o *MoveStreamEntry) GetWcOk() (*int32, bool)`

GetWcOk returns a tuple with the Wc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWc

`func (o *MoveStreamEntry) SetWc(v int32)`

SetWc sets Wc field to given value.


### GetBc

`func (o *MoveStreamEntry) GetBc() int32`

GetBc returns the Bc field if non-nil, zero value otherwise.

### GetBcOk

`func (o *MoveStreamEntry) GetBcOk() (*int32, bool)`

GetBcOk returns a tuple with the Bc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBc

`func (o *MoveStreamEntry) SetBc(v int32)`

SetBc sets Bc field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


