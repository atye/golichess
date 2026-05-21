# Perfs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chess960** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Atomic** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**RacingKings** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**UltraBullet** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Blitz** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**KingOfTheHill** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**ThreeCheck** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Antichess** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Crazyhouse** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Bullet** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Correspondence** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Horde** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Puzzle** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Classical** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Rapid** | Pointer to [**Perf**](Perf.md) |  | [optional] 
**Storm** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 
**Racer** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 
**Streak** | Pointer to [**PuzzleModePerf**](PuzzleModePerf.md) |  | [optional] 

## Methods

### NewPerfs

`func NewPerfs() *Perfs`

NewPerfs instantiates a new Perfs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfsWithDefaults

`func NewPerfsWithDefaults() *Perfs`

NewPerfsWithDefaults instantiates a new Perfs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChess960

`func (o *Perfs) GetChess960() Perf`

GetChess960 returns the Chess960 field if non-nil, zero value otherwise.

### GetChess960Ok

`func (o *Perfs) GetChess960Ok() (*Perf, bool)`

GetChess960Ok returns a tuple with the Chess960 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChess960

`func (o *Perfs) SetChess960(v Perf)`

SetChess960 sets Chess960 field to given value.

### HasChess960

`func (o *Perfs) HasChess960() bool`

HasChess960 returns a boolean if a field has been set.

### GetAtomic

`func (o *Perfs) GetAtomic() Perf`

GetAtomic returns the Atomic field if non-nil, zero value otherwise.

### GetAtomicOk

`func (o *Perfs) GetAtomicOk() (*Perf, bool)`

GetAtomicOk returns a tuple with the Atomic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAtomic

`func (o *Perfs) SetAtomic(v Perf)`

SetAtomic sets Atomic field to given value.

### HasAtomic

`func (o *Perfs) HasAtomic() bool`

HasAtomic returns a boolean if a field has been set.

### GetRacingKings

`func (o *Perfs) GetRacingKings() Perf`

GetRacingKings returns the RacingKings field if non-nil, zero value otherwise.

### GetRacingKingsOk

`func (o *Perfs) GetRacingKingsOk() (*Perf, bool)`

GetRacingKingsOk returns a tuple with the RacingKings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacingKings

`func (o *Perfs) SetRacingKings(v Perf)`

SetRacingKings sets RacingKings field to given value.

### HasRacingKings

`func (o *Perfs) HasRacingKings() bool`

HasRacingKings returns a boolean if a field has been set.

### GetUltraBullet

`func (o *Perfs) GetUltraBullet() Perf`

GetUltraBullet returns the UltraBullet field if non-nil, zero value otherwise.

### GetUltraBulletOk

`func (o *Perfs) GetUltraBulletOk() (*Perf, bool)`

GetUltraBulletOk returns a tuple with the UltraBullet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUltraBullet

`func (o *Perfs) SetUltraBullet(v Perf)`

SetUltraBullet sets UltraBullet field to given value.

### HasUltraBullet

`func (o *Perfs) HasUltraBullet() bool`

HasUltraBullet returns a boolean if a field has been set.

### GetBlitz

`func (o *Perfs) GetBlitz() Perf`

GetBlitz returns the Blitz field if non-nil, zero value otherwise.

### GetBlitzOk

`func (o *Perfs) GetBlitzOk() (*Perf, bool)`

GetBlitzOk returns a tuple with the Blitz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlitz

`func (o *Perfs) SetBlitz(v Perf)`

SetBlitz sets Blitz field to given value.

### HasBlitz

`func (o *Perfs) HasBlitz() bool`

HasBlitz returns a boolean if a field has been set.

### GetKingOfTheHill

`func (o *Perfs) GetKingOfTheHill() Perf`

GetKingOfTheHill returns the KingOfTheHill field if non-nil, zero value otherwise.

### GetKingOfTheHillOk

`func (o *Perfs) GetKingOfTheHillOk() (*Perf, bool)`

GetKingOfTheHillOk returns a tuple with the KingOfTheHill field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKingOfTheHill

`func (o *Perfs) SetKingOfTheHill(v Perf)`

SetKingOfTheHill sets KingOfTheHill field to given value.

### HasKingOfTheHill

`func (o *Perfs) HasKingOfTheHill() bool`

HasKingOfTheHill returns a boolean if a field has been set.

### GetThreeCheck

`func (o *Perfs) GetThreeCheck() Perf`

GetThreeCheck returns the ThreeCheck field if non-nil, zero value otherwise.

### GetThreeCheckOk

`func (o *Perfs) GetThreeCheckOk() (*Perf, bool)`

GetThreeCheckOk returns a tuple with the ThreeCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreeCheck

`func (o *Perfs) SetThreeCheck(v Perf)`

SetThreeCheck sets ThreeCheck field to given value.

### HasThreeCheck

`func (o *Perfs) HasThreeCheck() bool`

HasThreeCheck returns a boolean if a field has been set.

### GetAntichess

`func (o *Perfs) GetAntichess() Perf`

GetAntichess returns the Antichess field if non-nil, zero value otherwise.

### GetAntichessOk

`func (o *Perfs) GetAntichessOk() (*Perf, bool)`

GetAntichessOk returns a tuple with the Antichess field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAntichess

`func (o *Perfs) SetAntichess(v Perf)`

SetAntichess sets Antichess field to given value.

### HasAntichess

`func (o *Perfs) HasAntichess() bool`

HasAntichess returns a boolean if a field has been set.

### GetCrazyhouse

`func (o *Perfs) GetCrazyhouse() Perf`

GetCrazyhouse returns the Crazyhouse field if non-nil, zero value otherwise.

### GetCrazyhouseOk

`func (o *Perfs) GetCrazyhouseOk() (*Perf, bool)`

GetCrazyhouseOk returns a tuple with the Crazyhouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrazyhouse

`func (o *Perfs) SetCrazyhouse(v Perf)`

SetCrazyhouse sets Crazyhouse field to given value.

### HasCrazyhouse

`func (o *Perfs) HasCrazyhouse() bool`

HasCrazyhouse returns a boolean if a field has been set.

### GetBullet

`func (o *Perfs) GetBullet() Perf`

GetBullet returns the Bullet field if non-nil, zero value otherwise.

### GetBulletOk

`func (o *Perfs) GetBulletOk() (*Perf, bool)`

GetBulletOk returns a tuple with the Bullet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBullet

`func (o *Perfs) SetBullet(v Perf)`

SetBullet sets Bullet field to given value.

### HasBullet

`func (o *Perfs) HasBullet() bool`

HasBullet returns a boolean if a field has been set.

### GetCorrespondence

`func (o *Perfs) GetCorrespondence() Perf`

GetCorrespondence returns the Correspondence field if non-nil, zero value otherwise.

### GetCorrespondenceOk

`func (o *Perfs) GetCorrespondenceOk() (*Perf, bool)`

GetCorrespondenceOk returns a tuple with the Correspondence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrespondence

`func (o *Perfs) SetCorrespondence(v Perf)`

SetCorrespondence sets Correspondence field to given value.

### HasCorrespondence

`func (o *Perfs) HasCorrespondence() bool`

HasCorrespondence returns a boolean if a field has been set.

### GetHorde

`func (o *Perfs) GetHorde() Perf`

GetHorde returns the Horde field if non-nil, zero value otherwise.

### GetHordeOk

`func (o *Perfs) GetHordeOk() (*Perf, bool)`

GetHordeOk returns a tuple with the Horde field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHorde

`func (o *Perfs) SetHorde(v Perf)`

SetHorde sets Horde field to given value.

### HasHorde

`func (o *Perfs) HasHorde() bool`

HasHorde returns a boolean if a field has been set.

### GetPuzzle

`func (o *Perfs) GetPuzzle() Perf`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *Perfs) GetPuzzleOk() (*Perf, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *Perfs) SetPuzzle(v Perf)`

SetPuzzle sets Puzzle field to given value.

### HasPuzzle

`func (o *Perfs) HasPuzzle() bool`

HasPuzzle returns a boolean if a field has been set.

### GetClassical

`func (o *Perfs) GetClassical() Perf`

GetClassical returns the Classical field if non-nil, zero value otherwise.

### GetClassicalOk

`func (o *Perfs) GetClassicalOk() (*Perf, bool)`

GetClassicalOk returns a tuple with the Classical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassical

`func (o *Perfs) SetClassical(v Perf)`

SetClassical sets Classical field to given value.

### HasClassical

`func (o *Perfs) HasClassical() bool`

HasClassical returns a boolean if a field has been set.

### GetRapid

`func (o *Perfs) GetRapid() Perf`

GetRapid returns the Rapid field if non-nil, zero value otherwise.

### GetRapidOk

`func (o *Perfs) GetRapidOk() (*Perf, bool)`

GetRapidOk returns a tuple with the Rapid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapid

`func (o *Perfs) SetRapid(v Perf)`

SetRapid sets Rapid field to given value.

### HasRapid

`func (o *Perfs) HasRapid() bool`

HasRapid returns a boolean if a field has been set.

### GetStorm

`func (o *Perfs) GetStorm() PuzzleModePerf`

GetStorm returns the Storm field if non-nil, zero value otherwise.

### GetStormOk

`func (o *Perfs) GetStormOk() (*PuzzleModePerf, bool)`

GetStormOk returns a tuple with the Storm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorm

`func (o *Perfs) SetStorm(v PuzzleModePerf)`

SetStorm sets Storm field to given value.

### HasStorm

`func (o *Perfs) HasStorm() bool`

HasStorm returns a boolean if a field has been set.

### GetRacer

`func (o *Perfs) GetRacer() PuzzleModePerf`

GetRacer returns the Racer field if non-nil, zero value otherwise.

### GetRacerOk

`func (o *Perfs) GetRacerOk() (*PuzzleModePerf, bool)`

GetRacerOk returns a tuple with the Racer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRacer

`func (o *Perfs) SetRacer(v PuzzleModePerf)`

SetRacer sets Racer field to given value.

### HasRacer

`func (o *Perfs) HasRacer() bool`

HasRacer returns a boolean if a field has been set.

### GetStreak

`func (o *Perfs) GetStreak() PuzzleModePerf`

GetStreak returns the Streak field if non-nil, zero value otherwise.

### GetStreakOk

`func (o *Perfs) GetStreakOk() (*PuzzleModePerf, bool)`

GetStreakOk returns a tuple with the Streak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreak

`func (o *Perfs) SetStreak(v PuzzleModePerf)`

SetStreak sets Streak field to given value.

### HasStreak

`func (o *Perfs) HasStreak() bool`

HasStreak returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


