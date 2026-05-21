# MoveStreamEntryOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Variant** | Pointer to [**Variant**](Variant.md) |  | [optional] 
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**Perf** | Pointer to [**PerfType**](PerfType.md) |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**Fen** | Pointer to **string** |  | [optional] 
**Player** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**Turns** | Pointer to **int32** |  | [optional] 
**StartedAtTurn** | Pointer to **int32** |  | [optional] 
**Source** | Pointer to [**GameSource**](GameSource.md) |  | [optional] 
**Status** | Pointer to [**GameStatus**](GameStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**GamePlayers**](GamePlayers.md) |  | [optional] 

## Methods

### NewMoveStreamEntryOneOf

`func NewMoveStreamEntryOneOf(id string, ) *MoveStreamEntryOneOf`

NewMoveStreamEntryOneOf instantiates a new MoveStreamEntryOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveStreamEntryOneOfWithDefaults

`func NewMoveStreamEntryOneOfWithDefaults() *MoveStreamEntryOneOf`

NewMoveStreamEntryOneOfWithDefaults instantiates a new MoveStreamEntryOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveStreamEntryOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveStreamEntryOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveStreamEntryOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *MoveStreamEntryOneOf) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *MoveStreamEntryOneOf) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *MoveStreamEntryOneOf) SetVariant(v Variant)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *MoveStreamEntryOneOf) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *MoveStreamEntryOneOf) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *MoveStreamEntryOneOf) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *MoveStreamEntryOneOf) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *MoveStreamEntryOneOf) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *MoveStreamEntryOneOf) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *MoveStreamEntryOneOf) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *MoveStreamEntryOneOf) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *MoveStreamEntryOneOf) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *MoveStreamEntryOneOf) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *MoveStreamEntryOneOf) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *MoveStreamEntryOneOf) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *MoveStreamEntryOneOf) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetInitialFen

`func (o *MoveStreamEntryOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *MoveStreamEntryOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *MoveStreamEntryOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *MoveStreamEntryOneOf) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetFen

`func (o *MoveStreamEntryOneOf) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *MoveStreamEntryOneOf) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *MoveStreamEntryOneOf) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *MoveStreamEntryOneOf) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetPlayer

`func (o *MoveStreamEntryOneOf) GetPlayer() GameColor`

GetPlayer returns the Player field if non-nil, zero value otherwise.

### GetPlayerOk

`func (o *MoveStreamEntryOneOf) GetPlayerOk() (*GameColor, bool)`

GetPlayerOk returns a tuple with the Player field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayer

`func (o *MoveStreamEntryOneOf) SetPlayer(v GameColor)`

SetPlayer sets Player field to given value.

### HasPlayer

`func (o *MoveStreamEntryOneOf) HasPlayer() bool`

HasPlayer returns a boolean if a field has been set.

### GetTurns

`func (o *MoveStreamEntryOneOf) GetTurns() int32`

GetTurns returns the Turns field if non-nil, zero value otherwise.

### GetTurnsOk

`func (o *MoveStreamEntryOneOf) GetTurnsOk() (*int32, bool)`

GetTurnsOk returns a tuple with the Turns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurns

`func (o *MoveStreamEntryOneOf) SetTurns(v int32)`

SetTurns sets Turns field to given value.

### HasTurns

`func (o *MoveStreamEntryOneOf) HasTurns() bool`

HasTurns returns a boolean if a field has been set.

### GetStartedAtTurn

`func (o *MoveStreamEntryOneOf) GetStartedAtTurn() int32`

GetStartedAtTurn returns the StartedAtTurn field if non-nil, zero value otherwise.

### GetStartedAtTurnOk

`func (o *MoveStreamEntryOneOf) GetStartedAtTurnOk() (*int32, bool)`

GetStartedAtTurnOk returns a tuple with the StartedAtTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAtTurn

`func (o *MoveStreamEntryOneOf) SetStartedAtTurn(v int32)`

SetStartedAtTurn sets StartedAtTurn field to given value.

### HasStartedAtTurn

`func (o *MoveStreamEntryOneOf) HasStartedAtTurn() bool`

HasStartedAtTurn returns a boolean if a field has been set.

### GetSource

`func (o *MoveStreamEntryOneOf) GetSource() GameSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MoveStreamEntryOneOf) GetSourceOk() (*GameSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MoveStreamEntryOneOf) SetSource(v GameSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *MoveStreamEntryOneOf) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetStatus

`func (o *MoveStreamEntryOneOf) GetStatus() GameStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MoveStreamEntryOneOf) GetStatusOk() (*GameStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MoveStreamEntryOneOf) SetStatus(v GameStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MoveStreamEntryOneOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveStreamEntryOneOf) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveStreamEntryOneOf) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveStreamEntryOneOf) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveStreamEntryOneOf) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetLastMove

`func (o *MoveStreamEntryOneOf) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *MoveStreamEntryOneOf) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *MoveStreamEntryOneOf) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *MoveStreamEntryOneOf) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetPlayers

`func (o *MoveStreamEntryOneOf) GetPlayers() GamePlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *MoveStreamEntryOneOf) GetPlayersOk() (*GamePlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *MoveStreamEntryOneOf) SetPlayers(v GamePlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *MoveStreamEntryOneOf) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


