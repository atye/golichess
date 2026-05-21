# GameFullEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**Clock** | Pointer to [**GameFullEventClock**](GameFullEventClock.md) |  | [optional] 
**Speed** | [**Speed**](Speed.md) |  | 
**Perf** | [**GameFullEventPerf**](GameFullEventPerf.md) |  | 
**Rated** | **bool** |  | 
**CreatedAt** | **int64** |  | 
**White** | [**GameEventPlayer**](GameEventPlayer.md) |  | 
**Black** | [**GameEventPlayer**](GameEventPlayer.md) |  | 
**InitialFen** | **string** |  | [default to "startpos"]
**State** | [**GameStateEvent**](GameStateEvent.md) |  | 
**DaysPerTurn** | Pointer to **int32** | If the game is correspondence | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 

## Methods

### NewGameFullEvent

`func NewGameFullEvent(type_ string, id string, variant Variant, speed Speed, perf GameFullEventPerf, rated bool, createdAt int64, white GameEventPlayer, black GameEventPlayer, initialFen string, state GameStateEvent, ) *GameFullEvent`

NewGameFullEvent instantiates a new GameFullEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameFullEventWithDefaults

`func NewGameFullEventWithDefaults() *GameFullEvent`

NewGameFullEventWithDefaults instantiates a new GameFullEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameFullEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameFullEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameFullEvent) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameFullEvent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameFullEvent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameFullEvent) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *GameFullEvent) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *GameFullEvent) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *GameFullEvent) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *GameFullEvent) GetClock() GameFullEventClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *GameFullEvent) GetClockOk() (*GameFullEventClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *GameFullEvent) SetClock(v GameFullEventClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *GameFullEvent) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetSpeed

`func (o *GameFullEvent) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *GameFullEvent) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *GameFullEvent) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *GameFullEvent) GetPerf() GameFullEventPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *GameFullEvent) GetPerfOk() (*GameFullEventPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *GameFullEvent) SetPerf(v GameFullEventPerf)`

SetPerf sets Perf field to given value.


### GetRated

`func (o *GameFullEvent) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *GameFullEvent) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *GameFullEvent) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetCreatedAt

`func (o *GameFullEvent) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GameFullEvent) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GameFullEvent) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetWhite

`func (o *GameFullEvent) GetWhite() GameEventPlayer`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *GameFullEvent) GetWhiteOk() (*GameEventPlayer, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *GameFullEvent) SetWhite(v GameEventPlayer)`

SetWhite sets White field to given value.


### GetBlack

`func (o *GameFullEvent) GetBlack() GameEventPlayer`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *GameFullEvent) GetBlackOk() (*GameEventPlayer, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *GameFullEvent) SetBlack(v GameEventPlayer)`

SetBlack sets Black field to given value.


### GetInitialFen

`func (o *GameFullEvent) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *GameFullEvent) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *GameFullEvent) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetState

`func (o *GameFullEvent) GetState() GameStateEvent`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GameFullEvent) GetStateOk() (*GameStateEvent, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GameFullEvent) SetState(v GameStateEvent)`

SetState sets State field to given value.


### GetDaysPerTurn

`func (o *GameFullEvent) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *GameFullEvent) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *GameFullEvent) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *GameFullEvent) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetTournamentId

`func (o *GameFullEvent) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *GameFullEvent) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *GameFullEvent) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *GameFullEvent) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


