# GamesByUsers200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Rated** | Pointer to **bool** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] [default to "standard"]
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 
**StatusName** | Pointer to **string** |  | [optional] 
**Clock** | Pointer to [**GamesByUsers200ResponseInnerClock**](GamesByUsers200ResponseInnerClock.md) |  | [optional] 
**Players** | Pointer to [**GamesByUsers200ResponseInnerPlayers**](GamesByUsers200ResponseInnerPlayers.md) |  | [optional] 
**Winner** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGamesByUsers200ResponseInner

`func NewGamesByUsers200ResponseInner(id string, ) *GamesByUsers200ResponseInner`

NewGamesByUsers200ResponseInner instantiates a new GamesByUsers200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamesByUsers200ResponseInnerWithDefaults

`func NewGamesByUsers200ResponseInnerWithDefaults() *GamesByUsers200ResponseInner`

NewGamesByUsers200ResponseInnerWithDefaults instantiates a new GamesByUsers200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GamesByUsers200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GamesByUsers200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GamesByUsers200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetRated

`func (o *GamesByUsers200ResponseInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GamesByUsers200ResponseInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GamesByUsers200ResponseInner) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *GamesByUsers200ResponseInner) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetVariant

`func (o *GamesByUsers200ResponseInner) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GamesByUsers200ResponseInner) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GamesByUsers200ResponseInner) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *GamesByUsers200ResponseInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *GamesByUsers200ResponseInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GamesByUsers200ResponseInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GamesByUsers200ResponseInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *GamesByUsers200ResponseInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *GamesByUsers200ResponseInner) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GamesByUsers200ResponseInner) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GamesByUsers200ResponseInner) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *GamesByUsers200ResponseInner) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GamesByUsers200ResponseInner) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GamesByUsers200ResponseInner) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GamesByUsers200ResponseInner) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GamesByUsers200ResponseInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *GamesByUsers200ResponseInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GamesByUsers200ResponseInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GamesByUsers200ResponseInner) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GamesByUsers200ResponseInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusName

`func (o *GamesByUsers200ResponseInner) GetStatusName() string`

GetStatusName returns the StatusName field if non-nil, zero value otherwise.

### GetStatusNameOk

`func (o *GamesByUsers200ResponseInner) GetStatusNameOk() (*string, bool)`

GetStatusNameOk returns a tuple with the StatusName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusName

`func (o *GamesByUsers200ResponseInner) SetStatusName(v string)`

SetStatusName sets StatusName field to given value.

### HasStatusName

`func (o *GamesByUsers200ResponseInner) HasStatusName() bool`

HasStatusName returns a boolean if a field has been set.

### GetClock

`func (o *GamesByUsers200ResponseInner) GetClock() GamesByUsers200ResponseInnerClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GamesByUsers200ResponseInner) GetClockOk() (*GamesByUsers200ResponseInnerClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GamesByUsers200ResponseInner) SetClock(v GamesByUsers200ResponseInnerClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GamesByUsers200ResponseInner) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetPlayers

`func (o *GamesByUsers200ResponseInner) GetPlayers() GamesByUsers200ResponseInnerPlayers`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *GamesByUsers200ResponseInner) GetPlayersOk() (*GamesByUsers200ResponseInnerPlayers, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *GamesByUsers200ResponseInner) SetPlayers(v GamesByUsers200ResponseInnerPlayers)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *GamesByUsers200ResponseInner) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetWinner

`func (o *GamesByUsers200ResponseInner) GetWinner() string`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GamesByUsers200ResponseInner) GetWinnerOk() (*string, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GamesByUsers200ResponseInner) SetWinner(v string)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GamesByUsers200ResponseInner) HasWinner() bool`

HasWinner returns a boolean if a field has been set.

### SetWinnerNil

`func (o *GamesByUsers200ResponseInner) SetWinnerNil(b bool)`

 SetWinnerNil sets the value for Winner to be an explicit nil

### UnsetWinner
`func (o *GamesByUsers200ResponseInner) UnsetWinner()`

UnsetWinner ensures that no value is present for Winner, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


