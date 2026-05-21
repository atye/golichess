# ArenaTournamentFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FullName** | **string** |  | 
**Rated** | Pointer to **bool** |  | [optional] 
**Spotlight** | Pointer to [**ArenaTournamentFullSpotlight**](ArenaTournamentFullSpotlight.md) |  | [optional] 
**Berserkable** | Pointer to **bool** |  | [optional] 
**OnlyTitled** | Pointer to **bool** |  | [optional] 
**Clock** | [**Clock**](Clock.md) |  | 
**Minutes** | Pointer to **int32** |  | [optional] 
**CreatedBy** | Pointer to **string** |  | [optional] 
**System** | Pointer to **string** |  | [optional] 
**SecondsToStart** | Pointer to **int32** |  | [optional] 
**SecondsToFinish** | Pointer to **int32** |  | [optional] 
**IsFinished** | Pointer to **bool** |  | [optional] 
**IsRecentlyFinished** | Pointer to **bool** |  | [optional] 
**PairingsClosed** | Pointer to **bool** |  | [optional] 
**StartsAt** | Pointer to **string** |  | [optional] 
**NbPlayers** | **int32** |  | 
**Verdicts** | Pointer to [**Verdicts**](Verdicts.md) |  | [optional] 
**Quote** | Pointer to [**ArenaTournamentFullQuote**](ArenaTournamentFullQuote.md) |  | [optional] 
**GreatPlayer** | Pointer to [**ArenaTournamentFullGreatPlayer**](ArenaTournamentFullGreatPlayer.md) |  | [optional] 
**AllowList** | Pointer to **[]string** | List of usernames allowed to join the tournament | [optional] 
**HasMaxRating** | Pointer to **bool** |  | [optional] 
**MaxRating** | Pointer to [**ArenaRatingObj**](ArenaRatingObj.md) |  | [optional] 
**MinRating** | Pointer to [**ArenaRatingObj**](ArenaRatingObj.md) |  | [optional] 
**MinRatedGames** | Pointer to [**ArenaTournamentMinRatedGames**](ArenaTournamentMinRatedGames.md) |  | [optional] 
**BotsAllowed** | Pointer to **bool** |  | [optional] 
**MinAccountAgeInDays** | Pointer to **int32** |  | [optional] 
**Perf** | Pointer to [**ArenaTournamentFullPerf**](ArenaTournamentFullPerf.md) |  | [optional] 
**Schedule** | Pointer to [**ArenaTournamentFullSchedule**](ArenaTournamentFullSchedule.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] 
**Duels** | Pointer to [**[]ArenaTournamentFullDuelsInner**](ArenaTournamentFullDuelsInner.md) |  | [optional] 
**Standing** | Pointer to [**ArenaTournamentFullStanding**](ArenaTournamentFullStanding.md) |  | [optional] 
**Featured** | Pointer to [**ArenaTournamentFullFeatured**](ArenaTournamentFullFeatured.md) |  | [optional] 
**Podium** | Pointer to [**[]ArenaTournamentFullPodiumInner**](ArenaTournamentFullPodiumInner.md) |  | [optional] 
**Stats** | Pointer to [**ArenaTournamentFullStats**](ArenaTournamentFullStats.md) |  | [optional] 
**MyUsername** | Pointer to **string** |  | [optional] 

## Methods

### NewArenaTournamentFull

`func NewArenaTournamentFull(id string, fullName string, clock Clock, nbPlayers int32, ) *ArenaTournamentFull`

NewArenaTournamentFull instantiates a new ArenaTournamentFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaTournamentFullWithDefaults

`func NewArenaTournamentFullWithDefaults() *ArenaTournamentFull`

NewArenaTournamentFullWithDefaults instantiates a new ArenaTournamentFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArenaTournamentFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArenaTournamentFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArenaTournamentFull) SetId(v string)`

SetId sets Id field to given value.


### GetFullName

`func (o *ArenaTournamentFull) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ArenaTournamentFull) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ArenaTournamentFull) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetRated

`func (o *ArenaTournamentFull) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ArenaTournamentFull) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ArenaTournamentFull) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ArenaTournamentFull) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetSpotlight

`func (o *ArenaTournamentFull) GetSpotlight() ArenaTournamentFullSpotlight`

GetSpotlight returns the Spotlight field if non-nil, zero value otherwise.

### GetSpotlightOk

`func (o *ArenaTournamentFull) GetSpotlightOk() (*ArenaTournamentFullSpotlight, bool)`

GetSpotlightOk returns a tuple with the Spotlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotlight

`func (o *ArenaTournamentFull) SetSpotlight(v ArenaTournamentFullSpotlight)`

SetSpotlight sets Spotlight field to given value.

### HasSpotlight

`func (o *ArenaTournamentFull) HasSpotlight() bool`

HasSpotlight returns a boolean if a field has been set.

### GetBerserkable

`func (o *ArenaTournamentFull) GetBerserkable() bool`

GetBerserkable returns the Berserkable field if non-nil, zero value otherwise.

### GetBerserkableOk

`func (o *ArenaTournamentFull) GetBerserkableOk() (*bool, bool)`

GetBerserkableOk returns a tuple with the Berserkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerserkable

`func (o *ArenaTournamentFull) SetBerserkable(v bool)`

SetBerserkable sets Berserkable field to given value.

### HasBerserkable

`func (o *ArenaTournamentFull) HasBerserkable() bool`

HasBerserkable returns a boolean if a field has been set.

### GetOnlyTitled

`func (o *ArenaTournamentFull) GetOnlyTitled() bool`

GetOnlyTitled returns the OnlyTitled field if non-nil, zero value otherwise.

### GetOnlyTitledOk

`func (o *ArenaTournamentFull) GetOnlyTitledOk() (*bool, bool)`

GetOnlyTitledOk returns a tuple with the OnlyTitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyTitled

`func (o *ArenaTournamentFull) SetOnlyTitled(v bool)`

SetOnlyTitled sets OnlyTitled field to given value.

### HasOnlyTitled

`func (o *ArenaTournamentFull) HasOnlyTitled() bool`

HasOnlyTitled returns a boolean if a field has been set.

### GetClock

`func (o *ArenaTournamentFull) GetClock() Clock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ArenaTournamentFull) GetClockOk() (*Clock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ArenaTournamentFull) SetClock(v Clock)`

SetClock sets Clock field to given value.


### GetMinutes

`func (o *ArenaTournamentFull) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *ArenaTournamentFull) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *ArenaTournamentFull) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *ArenaTournamentFull) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArenaTournamentFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArenaTournamentFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArenaTournamentFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArenaTournamentFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSystem

`func (o *ArenaTournamentFull) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ArenaTournamentFull) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ArenaTournamentFull) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ArenaTournamentFull) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSecondsToStart

`func (o *ArenaTournamentFull) GetSecondsToStart() int32`

GetSecondsToStart returns the SecondsToStart field if non-nil, zero value otherwise.

### GetSecondsToStartOk

`func (o *ArenaTournamentFull) GetSecondsToStartOk() (*int32, bool)`

GetSecondsToStartOk returns a tuple with the SecondsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToStart

`func (o *ArenaTournamentFull) SetSecondsToStart(v int32)`

SetSecondsToStart sets SecondsToStart field to given value.

### HasSecondsToStart

`func (o *ArenaTournamentFull) HasSecondsToStart() bool`

HasSecondsToStart returns a boolean if a field has been set.

### GetSecondsToFinish

`func (o *ArenaTournamentFull) GetSecondsToFinish() int32`

GetSecondsToFinish returns the SecondsToFinish field if non-nil, zero value otherwise.

### GetSecondsToFinishOk

`func (o *ArenaTournamentFull) GetSecondsToFinishOk() (*int32, bool)`

GetSecondsToFinishOk returns a tuple with the SecondsToFinish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToFinish

`func (o *ArenaTournamentFull) SetSecondsToFinish(v int32)`

SetSecondsToFinish sets SecondsToFinish field to given value.

### HasSecondsToFinish

`func (o *ArenaTournamentFull) HasSecondsToFinish() bool`

HasSecondsToFinish returns a boolean if a field has been set.

### GetIsFinished

`func (o *ArenaTournamentFull) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *ArenaTournamentFull) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *ArenaTournamentFull) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.

### HasIsFinished

`func (o *ArenaTournamentFull) HasIsFinished() bool`

HasIsFinished returns a boolean if a field has been set.

### GetIsRecentlyFinished

`func (o *ArenaTournamentFull) GetIsRecentlyFinished() bool`

GetIsRecentlyFinished returns the IsRecentlyFinished field if non-nil, zero value otherwise.

### GetIsRecentlyFinishedOk

`func (o *ArenaTournamentFull) GetIsRecentlyFinishedOk() (*bool, bool)`

GetIsRecentlyFinishedOk returns a tuple with the IsRecentlyFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyFinished

`func (o *ArenaTournamentFull) SetIsRecentlyFinished(v bool)`

SetIsRecentlyFinished sets IsRecentlyFinished field to given value.

### HasIsRecentlyFinished

`func (o *ArenaTournamentFull) HasIsRecentlyFinished() bool`

HasIsRecentlyFinished returns a boolean if a field has been set.

### GetPairingsClosed

`func (o *ArenaTournamentFull) GetPairingsClosed() bool`

GetPairingsClosed returns the PairingsClosed field if non-nil, zero value otherwise.

### GetPairingsClosedOk

`func (o *ArenaTournamentFull) GetPairingsClosedOk() (*bool, bool)`

GetPairingsClosedOk returns a tuple with the PairingsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingsClosed

`func (o *ArenaTournamentFull) SetPairingsClosed(v bool)`

SetPairingsClosed sets PairingsClosed field to given value.

### HasPairingsClosed

`func (o *ArenaTournamentFull) HasPairingsClosed() bool`

HasPairingsClosed returns a boolean if a field has been set.

### GetStartsAt

`func (o *ArenaTournamentFull) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ArenaTournamentFull) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ArenaTournamentFull) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *ArenaTournamentFull) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetNbPlayers

`func (o *ArenaTournamentFull) GetNbPlayers() int32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *ArenaTournamentFull) GetNbPlayersOk() (*int32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *ArenaTournamentFull) SetNbPlayers(v int32)`

SetNbPlayers sets NbPlayers field to given value.


### GetVerdicts

`func (o *ArenaTournamentFull) GetVerdicts() Verdicts`

GetVerdicts returns the Verdicts field if non-nil, zero value otherwise.

### GetVerdictsOk

`func (o *ArenaTournamentFull) GetVerdictsOk() (*Verdicts, bool)`

GetVerdictsOk returns a tuple with the Verdicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdicts

`func (o *ArenaTournamentFull) SetVerdicts(v Verdicts)`

SetVerdicts sets Verdicts field to given value.

### HasVerdicts

`func (o *ArenaTournamentFull) HasVerdicts() bool`

HasVerdicts returns a boolean if a field has been set.

### GetQuote

`func (o *ArenaTournamentFull) GetQuote() ArenaTournamentFullQuote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *ArenaTournamentFull) GetQuoteOk() (*ArenaTournamentFullQuote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *ArenaTournamentFull) SetQuote(v ArenaTournamentFullQuote)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *ArenaTournamentFull) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetGreatPlayer

`func (o *ArenaTournamentFull) GetGreatPlayer() ArenaTournamentFullGreatPlayer`

GetGreatPlayer returns the GreatPlayer field if non-nil, zero value otherwise.

### GetGreatPlayerOk

`func (o *ArenaTournamentFull) GetGreatPlayerOk() (*ArenaTournamentFullGreatPlayer, bool)`

GetGreatPlayerOk returns a tuple with the GreatPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreatPlayer

`func (o *ArenaTournamentFull) SetGreatPlayer(v ArenaTournamentFullGreatPlayer)`

SetGreatPlayer sets GreatPlayer field to given value.

### HasGreatPlayer

`func (o *ArenaTournamentFull) HasGreatPlayer() bool`

HasGreatPlayer returns a boolean if a field has been set.

### GetAllowList

`func (o *ArenaTournamentFull) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *ArenaTournamentFull) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *ArenaTournamentFull) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *ArenaTournamentFull) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetHasMaxRating

`func (o *ArenaTournamentFull) GetHasMaxRating() bool`

GetHasMaxRating returns the HasMaxRating field if non-nil, zero value otherwise.

### GetHasMaxRatingOk

`func (o *ArenaTournamentFull) GetHasMaxRatingOk() (*bool, bool)`

GetHasMaxRatingOk returns a tuple with the HasMaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMaxRating

`func (o *ArenaTournamentFull) SetHasMaxRating(v bool)`

SetHasMaxRating sets HasMaxRating field to given value.

### HasHasMaxRating

`func (o *ArenaTournamentFull) HasHasMaxRating() bool`

HasHasMaxRating returns a boolean if a field has been set.

### GetMaxRating

`func (o *ArenaTournamentFull) GetMaxRating() ArenaRatingObj`

GetMaxRating returns the MaxRating field if non-nil, zero value otherwise.

### GetMaxRatingOk

`func (o *ArenaTournamentFull) GetMaxRatingOk() (*ArenaRatingObj, bool)`

GetMaxRatingOk returns a tuple with the MaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRating

`func (o *ArenaTournamentFull) SetMaxRating(v ArenaRatingObj)`

SetMaxRating sets MaxRating field to given value.

### HasMaxRating

`func (o *ArenaTournamentFull) HasMaxRating() bool`

HasMaxRating returns a boolean if a field has been set.

### GetMinRating

`func (o *ArenaTournamentFull) GetMinRating() ArenaRatingObj`

GetMinRating returns the MinRating field if non-nil, zero value otherwise.

### GetMinRatingOk

`func (o *ArenaTournamentFull) GetMinRatingOk() (*ArenaRatingObj, bool)`

GetMinRatingOk returns a tuple with the MinRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRating

`func (o *ArenaTournamentFull) SetMinRating(v ArenaRatingObj)`

SetMinRating sets MinRating field to given value.

### HasMinRating

`func (o *ArenaTournamentFull) HasMinRating() bool`

HasMinRating returns a boolean if a field has been set.

### GetMinRatedGames

`func (o *ArenaTournamentFull) GetMinRatedGames() ArenaTournamentMinRatedGames`

GetMinRatedGames returns the MinRatedGames field if non-nil, zero value otherwise.

### GetMinRatedGamesOk

`func (o *ArenaTournamentFull) GetMinRatedGamesOk() (*ArenaTournamentMinRatedGames, bool)`

GetMinRatedGamesOk returns a tuple with the MinRatedGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRatedGames

`func (o *ArenaTournamentFull) SetMinRatedGames(v ArenaTournamentMinRatedGames)`

SetMinRatedGames sets MinRatedGames field to given value.

### HasMinRatedGames

`func (o *ArenaTournamentFull) HasMinRatedGames() bool`

HasMinRatedGames returns a boolean if a field has been set.

### GetBotsAllowed

`func (o *ArenaTournamentFull) GetBotsAllowed() bool`

GetBotsAllowed returns the BotsAllowed field if non-nil, zero value otherwise.

### GetBotsAllowedOk

`func (o *ArenaTournamentFull) GetBotsAllowedOk() (*bool, bool)`

GetBotsAllowedOk returns a tuple with the BotsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotsAllowed

`func (o *ArenaTournamentFull) SetBotsAllowed(v bool)`

SetBotsAllowed sets BotsAllowed field to given value.

### HasBotsAllowed

`func (o *ArenaTournamentFull) HasBotsAllowed() bool`

HasBotsAllowed returns a boolean if a field has been set.

### GetMinAccountAgeInDays

`func (o *ArenaTournamentFull) GetMinAccountAgeInDays() int32`

GetMinAccountAgeInDays returns the MinAccountAgeInDays field if non-nil, zero value otherwise.

### GetMinAccountAgeInDaysOk

`func (o *ArenaTournamentFull) GetMinAccountAgeInDaysOk() (*int32, bool)`

GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAccountAgeInDays

`func (o *ArenaTournamentFull) SetMinAccountAgeInDays(v int32)`

SetMinAccountAgeInDays sets MinAccountAgeInDays field to given value.

### HasMinAccountAgeInDays

`func (o *ArenaTournamentFull) HasMinAccountAgeInDays() bool`

HasMinAccountAgeInDays returns a boolean if a field has been set.

### GetPerf

`func (o *ArenaTournamentFull) GetPerf() ArenaTournamentFullPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ArenaTournamentFull) GetPerfOk() (*ArenaTournamentFullPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ArenaTournamentFull) SetPerf(v ArenaTournamentFullPerf)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ArenaTournamentFull) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetSchedule

`func (o *ArenaTournamentFull) GetSchedule() ArenaTournamentFullSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ArenaTournamentFull) GetScheduleOk() (*ArenaTournamentFullSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ArenaTournamentFull) SetSchedule(v ArenaTournamentFullSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ArenaTournamentFull) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetDescription

`func (o *ArenaTournamentFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArenaTournamentFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArenaTournamentFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArenaTournamentFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariant

`func (o *ArenaTournamentFull) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ArenaTournamentFull) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ArenaTournamentFull) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ArenaTournamentFull) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetDuels

`func (o *ArenaTournamentFull) GetDuels() []ArenaTournamentFullDuelsInner`

GetDuels returns the Duels field if non-nil, zero value otherwise.

### GetDuelsOk

`func (o *ArenaTournamentFull) GetDuelsOk() (*[]ArenaTournamentFullDuelsInner, bool)`

GetDuelsOk returns a tuple with the Duels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuels

`func (o *ArenaTournamentFull) SetDuels(v []ArenaTournamentFullDuelsInner)`

SetDuels sets Duels field to given value.

### HasDuels

`func (o *ArenaTournamentFull) HasDuels() bool`

HasDuels returns a boolean if a field has been set.

### GetStanding

`func (o *ArenaTournamentFull) GetStanding() ArenaTournamentFullStanding`

GetStanding returns the Standing field if non-nil, zero value otherwise.

### GetStandingOk

`func (o *ArenaTournamentFull) GetStandingOk() (*ArenaTournamentFullStanding, bool)`

GetStandingOk returns a tuple with the Standing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanding

`func (o *ArenaTournamentFull) SetStanding(v ArenaTournamentFullStanding)`

SetStanding sets Standing field to given value.

### HasStanding

`func (o *ArenaTournamentFull) HasStanding() bool`

HasStanding returns a boolean if a field has been set.

### GetFeatured

`func (o *ArenaTournamentFull) GetFeatured() ArenaTournamentFullFeatured`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ArenaTournamentFull) GetFeaturedOk() (*ArenaTournamentFullFeatured, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ArenaTournamentFull) SetFeatured(v ArenaTournamentFullFeatured)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ArenaTournamentFull) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPodium

`func (o *ArenaTournamentFull) GetPodium() []ArenaTournamentFullPodiumInner`

GetPodium returns the Podium field if non-nil, zero value otherwise.

### GetPodiumOk

`func (o *ArenaTournamentFull) GetPodiumOk() (*[]ArenaTournamentFullPodiumInner, bool)`

GetPodiumOk returns a tuple with the Podium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodium

`func (o *ArenaTournamentFull) SetPodium(v []ArenaTournamentFullPodiumInner)`

SetPodium sets Podium field to given value.

### HasPodium

`func (o *ArenaTournamentFull) HasPodium() bool`

HasPodium returns a boolean if a field has been set.

### GetStats

`func (o *ArenaTournamentFull) GetStats() ArenaTournamentFullStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ArenaTournamentFull) GetStatsOk() (*ArenaTournamentFullStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ArenaTournamentFull) SetStats(v ArenaTournamentFullStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ArenaTournamentFull) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetMyUsername

`func (o *ArenaTournamentFull) GetMyUsername() string`

GetMyUsername returns the MyUsername field if non-nil, zero value otherwise.

### GetMyUsernameOk

`func (o *ArenaTournamentFull) GetMyUsernameOk() (*string, bool)`

GetMyUsernameOk returns a tuple with the MyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyUsername

`func (o *ArenaTournamentFull) SetMyUsername(v string)`

SetMyUsername sets MyUsername field to given value.

### HasMyUsername

`func (o *ArenaTournamentFull) HasMyUsername() bool`

HasMyUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


