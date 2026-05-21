# GameStreamGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Rated** | Pointer to **bool** |  | [optional] 
**Variant** | Pointer to [**VariantKey**](VariantKey.md) |  | [optional] [default to VARIANTKEY_STANDARD]
**Speed** | Pointer to [**Speed**](Speed.md) |  | [optional] 
**Perf** | Pointer to [**PerfType**](PerfType.md) |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to [**GameStatusId**](GameStatusId.md) |  | [optional] 
**StatusName** | Pointer to [**GameStatusName**](GameStatusName.md) |  | [optional] 
**Clock** | Pointer to [**GameStreamGameClock**](GameStreamGameClock.md) |  | [optional] 
**Players** | Pointer to [**GameStreamGamePlayers**](GameStreamGamePlayers.md) |  | [optional] 
**Winner** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 

## Methods

### NewGameStreamGame

`func NewGameStreamGame(id string, ) *GameStreamGame`

NewGameStreamGame instantiates a new GameStreamGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameStreamGameWithDefaults

`func NewGameStreamGameWithDefaults() *GameStreamGame`

NewGameStreamGameWithDefaults instantiates a new GameStreamGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GameStreamGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameStreamGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameStreamGame) SetId(v string)`

SetId sets Id field to given value.


### GetRated

`func (o *GameStreamGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameStreamGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameStreamGame) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *GameStreamGame) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetVariant

`func (o *GameStreamGame) GetVariant() VariantKey`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameStreamGame) GetVariantOk() (*VariantKey, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameStreamGame) SetVariant(v VariantKey)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *GameStreamGame) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *GameStreamGame) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameStreamGame) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameStreamGame) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GameStreamGame) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *GameStreamGame) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameStreamGame) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameStreamGame) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *GameStreamGame) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GameStreamGame) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameStreamGame) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameStreamGame) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GameStreamGame) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *GameStreamGame) GetStatus() GameStatusId`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GameStreamGame) GetStatusOk() (*GameStatusId, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GameStreamGame) SetStatus(v GameStatusId)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GameStreamGame) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusName

`func (o *GameStreamGame) GetStatusName() GameStatusName`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *GameStreamGame) GetStatusNameOk() (*GameStatusName, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *GameStreamGame) SetStatusName(v GameStatusName)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *GameStreamGame) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetClock

`func (o *GameStreamGame) GetClock() GameStreamGameClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GameStreamGame) GetClockOk() (*GameStreamGameClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GameStreamGame) SetClock(v GameStreamGameClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GameStreamGame) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetPlayers

`func (o *GameStreamGame) GetPlayers() GameStreamGamePlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GameStreamGame) GetPlayersOk() (*GameStreamGamePlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GameStreamGame) SetPlayers(v GameStreamGamePlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *GameStreamGame) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetWinner

`func (o *GameStreamGame) GetWinner() GameColor`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GameStreamGame) GetWinnerOk() (*GameColor, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GameStreamGame) SetWinner(v GameColor)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GameStreamGame) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


