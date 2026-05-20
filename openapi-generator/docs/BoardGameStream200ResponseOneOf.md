# BoardGameStream200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**Clock** | Pointer to [**BoardGameStream200ResponseOneOfClock**](BoardGameStream200ResponseOneOfClock.md) |  | [optional] 
**Speed** | **string** |  | 
**Perf** | [**BoardGameStream200ResponseOneOfPerf**](BoardGameStream200ResponseOneOfPerf.md) |  | 
**Rated** | **bool** |  | 
**CreatedAt** | **int64** |  | 
**White** | [**BoardGameStream200ResponseOneOfWhite**](BoardGameStream200ResponseOneOfWhite.md) |  | 
**Black** | [**BoardGameStream200ResponseOneOfWhite**](BoardGameStream200ResponseOneOfWhite.md) |  | 
**InitialFen** | **string** |  | [default to "startpos"]
**State** | [**BoardGameStream200ResponseOneOfState**](BoardGameStream200ResponseOneOfState.md) |  | 
**DaysPerTurn** | Pointer to **int32** | If the game is correspondence | [optional] 
**TournamentId** | Pointer to **string** |  | [optional] 

## Methods

### NewBoardGameStream200ResponseOneOf

`func NewBoardGameStream200ResponseOneOf(type_ string, id string, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, speed string, perf BoardGameStream200ResponseOneOfPerf, rated bool, createdAt int64, white BoardGameStream200ResponseOneOfWhite, black BoardGameStream200ResponseOneOfWhite, initialFen string, state BoardGameStream200ResponseOneOfState, ) *BoardGameStream200ResponseOneOf`

NewBoardGameStream200ResponseOneOf instantiates a new BoardGameStream200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardGameStream200ResponseOneOfWithDefaults

`func NewBoardGameStream200ResponseOneOfWithDefaults() *BoardGameStream200ResponseOneOf`

NewBoardGameStream200ResponseOneOfWithDefaults instantiates a new BoardGameStream200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *BoardGameStream200ResponseOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BoardGameStream200ResponseOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BoardGameStream200ResponseOneOf) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *BoardGameStream200ResponseOneOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BoardGameStream200ResponseOneOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BoardGameStream200ResponseOneOf) SetId(v string)`

SetId sets Id field to given value.


### GetVariant

`func (o *BoardGameStream200ResponseOneOf) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *BoardGameStream200ResponseOneOf) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *BoardGameStream200ResponseOneOf) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *BoardGameStream200ResponseOneOf) GetClock() BoardGameStream200ResponseOneOfClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *BoardGameStream200ResponseOneOf) GetClockOk() (*BoardGameStream200ResponseOneOfClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *BoardGameStream200ResponseOneOf) SetClock(v BoardGameStream200ResponseOneOfClock)`

SetClock sets Clock field to given value.

### HasClock

`func (o *BoardGameStream200ResponseOneOf) HasClock() bool`

HasClock returns a boolean if a field has been set.

### GetSpeed

`func (o *BoardGameStream200ResponseOneOf) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *BoardGameStream200ResponseOneOf) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *BoardGameStream200ResponseOneOf) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetPerf

`func (o *BoardGameStream200ResponseOneOf) GetPerf() BoardGameStream200ResponseOneOfPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *BoardGameStream200ResponseOneOf) GetPerfOk() (*BoardGameStream200ResponseOneOfPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *BoardGameStream200ResponseOneOf) SetPerf(v BoardGameStream200ResponseOneOfPerf)`

SetPerf sets Perf field to given value.


### GetRated

`func (o *BoardGameStream200ResponseOneOf) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *BoardGameStream200ResponseOneOf) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *BoardGameStream200ResponseOneOf) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetCreatedAt

`func (o *BoardGameStream200ResponseOneOf) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BoardGameStream200ResponseOneOf) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BoardGameStream200ResponseOneOf) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.


### GetWhite

`func (o *BoardGameStream200ResponseOneOf) GetWhite() BoardGameStream200ResponseOneOfWhite`

GetWhite returns the White field if non-nil, zero value otherwise.

### GetWhiteOk

`func (o *BoardGameStream200ResponseOneOf) GetWhiteOk() (*BoardGameStream200ResponseOneOfWhite, bool)`

GetWhiteOk returns a tuple with the White field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhite

`func (o *BoardGameStream200ResponseOneOf) SetWhite(v BoardGameStream200ResponseOneOfWhite)`

SetWhite sets White field to given value.


### GetBlack

`func (o *BoardGameStream200ResponseOneOf) GetBlack() BoardGameStream200ResponseOneOfWhite`

GetBlack returns the Black field if non-nil, zero value otherwise.

### GetBlackOk

`func (o *BoardGameStream200ResponseOneOf) GetBlackOk() (*BoardGameStream200ResponseOneOfWhite, bool)`

GetBlackOk returns a tuple with the Black field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlack

`func (o *BoardGameStream200ResponseOneOf) SetBlack(v BoardGameStream200ResponseOneOfWhite)`

SetBlack sets Black field to given value.


### GetInitialFen

`func (o *BoardGameStream200ResponseOneOf) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *BoardGameStream200ResponseOneOf) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *BoardGameStream200ResponseOneOf) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.


### GetState

`func (o *BoardGameStream200ResponseOneOf) GetState() BoardGameStream200ResponseOneOfState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BoardGameStream200ResponseOneOf) GetStateOk() (*BoardGameStream200ResponseOneOfState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BoardGameStream200ResponseOneOf) SetState(v BoardGameStream200ResponseOneOfState)`

SetState sets State field to given value.


### GetDaysPerTurn

`func (o *BoardGameStream200ResponseOneOf) GetDaysPerTurn() int32`

GetDaysPerTurn returns the DaysPerTurn field if non-nil, zero value otherwise.

### GetDaysPerTurnOk

`func (o *BoardGameStream200ResponseOneOf) GetDaysPerTurnOk() (*int32, bool)`

GetDaysPerTurnOk returns a tuple with the DaysPerTurn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysPerTurn

`func (o *BoardGameStream200ResponseOneOf) SetDaysPerTurn(v int32)`

SetDaysPerTurn sets DaysPerTurn field to given value.

### HasDaysPerTurn

`func (o *BoardGameStream200ResponseOneOf) HasDaysPerTurn() bool`

HasDaysPerTurn returns a boolean if a field has been set.

### GetTournamentId

`func (o *BoardGameStream200ResponseOneOf) GetTournamentId() string`

GetTournamentId returns the TournamentId field if non-nil, zero value otherwise.

### GetTournamentIdOk

`func (o *BoardGameStream200ResponseOneOf) GetTournamentIdOk() (*string, bool)`

GetTournamentIdOk returns a tuple with the TournamentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTournamentId

`func (o *BoardGameStream200ResponseOneOf) SetTournamentId(v string)`

SetTournamentId sets TournamentId field to given value.

### HasTournamentId

`func (o *BoardGameStream200ResponseOneOf) HasTournamentId() bool`

HasTournamentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


