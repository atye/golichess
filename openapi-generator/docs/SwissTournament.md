# SwissTournament

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**StartsAt** | **string** |  | 
**Name** | **string** |  | 
**Clock** | [**SwissTournamentClock**](SwissTournamentClock.md) |  | 
**Variant** | **string** |  | 
**Round** | **float32** |  | 
**NbRounds** | **float32** |  | 
**NbPlayers** | **float32** |  | 
**NbOngoing** | **float32** |  | 
**Status** | [**SwissStatus**](SwissStatus.md) |  | 
**Stats** | Pointer to [**SwissTournamentStats**](SwissTournamentStats.md) |  | [optional] 
**Rated** | **bool** |  | 
**Verdicts** | [**Verdicts**](Verdicts.md) |  | 
**NextRound** | Pointer to [**SwissTournamentNextRound**](SwissTournamentNextRound.md) |  | [optional] 

## Methods

### NewSwissTournament

`func NewSwissTournament(id string, createdBy string, startsAt string, name string, clock SwissTournamentClock, variant string, round float32, nbRounds float32, nbPlayers float32, nbOngoing float32, status SwissStatus, rated bool, verdicts Verdicts, ) *SwissTournament`

NewSwissTournament instantiates a new SwissTournament object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwissTournamentWithDefaults

`func NewSwissTournamentWithDefaults() *SwissTournament`

NewSwissTournamentWithDefaults instantiates a new SwissTournament object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SwissTournament) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SwissTournament) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SwissTournament) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *SwissTournament) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *SwissTournament) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *SwissTournament) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetStartsAt

`func (o *SwissTournament) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *SwissTournament) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *SwissTournament) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.


### GetName

`func (o *SwissTournament) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SwissTournament) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SwissTournament) SetName(v string)`

SetName sets Name field to given value.


### GetClock

`func (o *SwissTournament) GetClock() SwissTournamentClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *SwissTournament) GetClockOk() (*SwissTournamentClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *SwissTournament) SetClock(v SwissTournamentClock)`

SetClock sets Clock field to given value.


### GetVariant

`func (o *SwissTournament) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *SwissTournament) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *SwissTournament) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetRound

`func (o *SwissTournament) GetRound() float32`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *SwissTournament) GetRoundOk() (*float32, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *SwissTournament) SetRound(v float32)`

SetRound sets Round field to given value.


### GetNbRounds

`func (o *SwissTournament) GetNbRounds() float32`

GetNbRounds returns the NbRounds field if non-nil, zero value otherwise.

### GetNbRoundsOk

`func (o *SwissTournament) GetNbRoundsOk() (*float32, bool)`

GetNbRoundsOk returns a tuple with the NbRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbRounds

`func (o *SwissTournament) SetNbRounds(v float32)`

SetNbRounds sets NbRounds field to given value.


### GetNbPlayers

`func (o *SwissTournament) GetNbPlayers() float32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *SwissTournament) GetNbPlayersOk() (*float32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *SwissTournament) SetNbPlayers(v float32)`

SetNbPlayers sets NbPlayers field to given value.


### GetNbOngoing

`func (o *SwissTournament) GetNbOngoing() float32`

GetNbOngoing returns the NbOngoing field if non-nil, zero value otherwise.

### GetNbOngoingOk

`func (o *SwissTournament) GetNbOngoingOk() (*float32, bool)`

GetNbOngoingOk returns a tuple with the NbOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOngoing

`func (o *SwissTournament) SetNbOngoing(v float32)`

SetNbOngoing sets NbOngoing field to given value.


### GetStatus

`func (o *SwissTournament) GetStatus() SwissStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SwissTournament) GetStatusOk() (*SwissStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SwissTournament) SetStatus(v SwissStatus)`

SetStatus sets Status field to given value.


### GetStats

`func (o *SwissTournament) GetStats() SwissTournamentStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *SwissTournament) GetStatsOk() (*SwissTournamentStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *SwissTournament) SetStats(v SwissTournamentStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *SwissTournament) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRated

`func (o *SwissTournament) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *SwissTournament) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *SwissTournament) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetVerdicts

`func (o *SwissTournament) GetVerdicts() Verdicts`

GetVerdicts returns the Verdicts field if non-nil, zero value otherwise.

### GetVerdictsOk

`func (o *SwissTournament) GetVerdictsOk() (*Verdicts, bool)`

GetVerdictsOk returns a tuple with the Verdicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdicts

`func (o *SwissTournament) SetVerdicts(v Verdicts)`

SetVerdicts sets Verdicts field to given value.


### GetNextRound

`func (o *SwissTournament) GetNextRound() SwissTournamentNextRound`

GetNextRound returns the NextRound field if non-nil, zero value otherwise.

### GetNextRoundOk

`func (o *SwissTournament) GetNextRoundOk() (*SwissTournamentNextRound, bool)`

GetNextRoundOk returns a tuple with the NextRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRound

`func (o *SwissTournament) SetNextRound(v SwissTournamentNextRound)`

SetNextRound sets NextRound field to given value.

### HasNextRound

`func (o *SwissTournament) HasNextRound() bool`

HasNextRound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


