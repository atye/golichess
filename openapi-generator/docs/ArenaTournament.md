# ArenaTournament

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**System** | **string** |  | 
**Minutes** | **int32** |  | 
**Clock** | [**Clock**](Clock.md) |  | 
**Rated** | **bool** |  | 
**FullName** | **string** |  | 
**NbPlayers** | **int32** |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**StartsAt** | **int64** |  | 
**FinishesAt** | **int64** |  | 
**Status** | [**ArenaStatus**](ArenaStatus.md) |  | 
**Perf** | [**ArenaPerf**](ArenaPerf.md) |  | 
**SecondsToStart** | Pointer to **int32** |  | [optional] 
**HasMaxRating** | Pointer to **bool** |  | [optional] 
**MaxRating** | Pointer to [**ArenaRatingObj**](ArenaRatingObj.md) |  | [optional] 
**MinRating** | Pointer to [**ArenaRatingObj**](ArenaRatingObj.md) |  | [optional] 
**MinRatedGames** | Pointer to [**ArenaTournamentMinRatedGames**](ArenaTournamentMinRatedGames.md) |  | [optional] 
**BotsAllowed** | Pointer to **bool** |  | [optional] 
**MinAccountAgeInDays** | Pointer to **int32** |  | [optional] 
**OnlyTitled** | Pointer to **bool** |  | [optional] 
**TeamMember** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to [**ArenaPosition**](ArenaPosition.md) |  | [optional] 
**Schedule** | Pointer to [**ArenaTournamentSchedule**](ArenaTournamentSchedule.md) |  | [optional] 
**TeamBattle** | Pointer to [**ArenaTournamentTeamBattle**](ArenaTournamentTeamBattle.md) |  | [optional] 
**Winner** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 

## Methods

### NewArenaTournament

`func NewArenaTournament(id string, createdBy string, system string, minutes int32, clock Clock, rated bool, fullName string, nbPlayers int32, variant Variant, startsAt int64, finishesAt int64, status ArenaStatus, perf ArenaPerf, ) *ArenaTournament`

NewArenaTournament instantiates a new ArenaTournament object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaTournamentWithDefaults

`func NewArenaTournamentWithDefaults() *ArenaTournament`

NewArenaTournamentWithDefaults instantiates a new ArenaTournament object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArenaTournament) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArenaTournament) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArenaTournament) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ArenaTournament) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArenaTournament) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArenaTournament) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetSystem

`func (o *ArenaTournament) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ArenaTournament) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ArenaTournament) SetSystem(v string)`

SetSystem sets System field to given value.


### GetMinutes

`func (o *ArenaTournament) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *ArenaTournament) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *ArenaTournament) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.


### GetClock

`func (o *ArenaTournament) GetClock() Clock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ArenaTournament) GetClockOk() (*Clock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ArenaTournament) SetClock(v Clock)`

SetClock sets Clock field to given value.


### GetRated

`func (o *ArenaTournament) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ArenaTournament) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ArenaTournament) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetFullName

`func (o *ArenaTournament) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ArenaTournament) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ArenaTournament) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetNbPlayers

`func (o *ArenaTournament) GetNbPlayers() int32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *ArenaTournament) GetNbPlayersOk() (*int32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *ArenaTournament) SetNbPlayers(v int32)`

SetNbPlayers sets NbPlayers field to given value.


### GetVariant

`func (o *ArenaTournament) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ArenaTournament) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ArenaTournament) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetStartsAt

`func (o *ArenaTournament) GetStartsAt() int64`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ArenaTournament) GetStartsAtOk() (*int64, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ArenaTournament) SetStartsAt(v int64)`

SetStartsAt sets StartsAt field to given value.


### GetFinishesAt

`func (o *ArenaTournament) GetFinishesAt() int64`

GetFinishesAt returns the FinishesAt field if non-nil, zero value otherwise.

### GetFinishesAtOk

`func (o *ArenaTournament) GetFinishesAtOk() (*int64, bool)`

GetFinishesAtOk returns a tuple with the FinishesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishesAt

`func (o *ArenaTournament) SetFinishesAt(v int64)`

SetFinishesAt sets FinishesAt field to given value.


### GetStatus

`func (o *ArenaTournament) GetStatus() ArenaStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArenaTournament) GetStatusOk() (*ArenaStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArenaTournament) SetStatus(v ArenaStatus)`

SetStatus sets Status field to given value.


### GetPerf

`func (o *ArenaTournament) GetPerf() ArenaPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ArenaTournament) GetPerfOk() (*ArenaPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ArenaTournament) SetPerf(v ArenaPerf)`

SetPerf sets Perf field to given value.


### GetSecondsToStart

`func (o *ArenaTournament) GetSecondsToStart() int32`

GetSecondsToStart returns the SecondsToStart field if non-nil, zero value otherwise.

### GetSecondsToStartOk

`func (o *ArenaTournament) GetSecondsToStartOk() (*int32, bool)`

GetSecondsToStartOk returns a tuple with the SecondsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToStart

`func (o *ArenaTournament) SetSecondsToStart(v int32)`

SetSecondsToStart sets SecondsToStart field to given value.

### HasSecondsToStart

`func (o *ArenaTournament) HasSecondsToStart() bool`

HasSecondsToStart returns a boolean if a field has been set.

### GetHasMaxRating

`func (o *ArenaTournament) GetHasMaxRating() bool`

GetHasMaxRating returns the HasMaxRating field if non-nil, zero value otherwise.

### GetHasMaxRatingOk

`func (o *ArenaTournament) GetHasMaxRatingOk() (*bool, bool)`

GetHasMaxRatingOk returns a tuple with the HasMaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMaxRating

`func (o *ArenaTournament) SetHasMaxRating(v bool)`

SetHasMaxRating sets HasMaxRating field to given value.

### HasHasMaxRating

`func (o *ArenaTournament) HasHasMaxRating() bool`

HasHasMaxRating returns a boolean if a field has been set.

### GetMaxRating

`func (o *ArenaTournament) GetMaxRating() ArenaRatingObj`

GetMaxRating returns the MaxRating field if non-nil, zero value otherwise.

### GetMaxRatingOk

`func (o *ArenaTournament) GetMaxRatingOk() (*ArenaRatingObj, bool)`

GetMaxRatingOk returns a tuple with the MaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRating

`func (o *ArenaTournament) SetMaxRating(v ArenaRatingObj)`

SetMaxRating sets MaxRating field to given value.

### HasMaxRating

`func (o *ArenaTournament) HasMaxRating() bool`

HasMaxRating returns a boolean if a field has been set.

### GetMinRating

`func (o *ArenaTournament) GetMinRating() ArenaRatingObj`

GetMinRating returns the MinRating field if non-nil, zero value otherwise.

### GetMinRatingOk

`func (o *ArenaTournament) GetMinRatingOk() (*ArenaRatingObj, bool)`

GetMinRatingOk returns a tuple with the MinRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRating

`func (o *ArenaTournament) SetMinRating(v ArenaRatingObj)`

SetMinRating sets MinRating field to given value.

### HasMinRating

`func (o *ArenaTournament) HasMinRating() bool`

HasMinRating returns a boolean if a field has been set.

### GetMinRatedGames

`func (o *ArenaTournament) GetMinRatedGames() ArenaTournamentMinRatedGames`

GetMinRatedGames returns the MinRatedGames field if non-nil, zero value otherwise.

### GetMinRatedGamesOk

`func (o *ArenaTournament) GetMinRatedGamesOk() (*ArenaTournamentMinRatedGames, bool)`

GetMinRatedGamesOk returns a tuple with the MinRatedGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRatedGames

`func (o *ArenaTournament) SetMinRatedGames(v ArenaTournamentMinRatedGames)`

SetMinRatedGames sets MinRatedGames field to given value.

### HasMinRatedGames

`func (o *ArenaTournament) HasMinRatedGames() bool`

HasMinRatedGames returns a boolean if a field has been set.

### GetBotsAllowed

`func (o *ArenaTournament) GetBotsAllowed() bool`

GetBotsAllowed returns the BotsAllowed field if non-nil, zero value otherwise.

### GetBotsAllowedOk

`func (o *ArenaTournament) GetBotsAllowedOk() (*bool, bool)`

GetBotsAllowedOk returns a tuple with the BotsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotsAllowed

`func (o *ArenaTournament) SetBotsAllowed(v bool)`

SetBotsAllowed sets BotsAllowed field to given value.

### HasBotsAllowed

`func (o *ArenaTournament) HasBotsAllowed() bool`

HasBotsAllowed returns a boolean if a field has been set.

### GetMinAccountAgeInDays

`func (o *ArenaTournament) GetMinAccountAgeInDays() int32`

GetMinAccountAgeInDays returns the MinAccountAgeInDays field if non-nil, zero value otherwise.

### GetMinAccountAgeInDaysOk

`func (o *ArenaTournament) GetMinAccountAgeInDaysOk() (*int32, bool)`

GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAccountAgeInDays

`func (o *ArenaTournament) SetMinAccountAgeInDays(v int32)`

SetMinAccountAgeInDays sets MinAccountAgeInDays field to given value.

### HasMinAccountAgeInDays

`func (o *ArenaTournament) HasMinAccountAgeInDays() bool`

HasMinAccountAgeInDays returns a boolean if a field has been set.

### GetOnlyTitled

`func (o *ArenaTournament) GetOnlyTitled() bool`

GetOnlyTitled returns the OnlyTitled field if non-nil, zero value otherwise.

### GetOnlyTitledOk

`func (o *ArenaTournament) GetOnlyTitledOk() (*bool, bool)`

GetOnlyTitledOk returns a tuple with the OnlyTitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyTitled

`func (o *ArenaTournament) SetOnlyTitled(v bool)`

SetOnlyTitled sets OnlyTitled field to given value.

### HasOnlyTitled

`func (o *ArenaTournament) HasOnlyTitled() bool`

HasOnlyTitled returns a boolean if a field has been set.

### GetTeamMember

`func (o *ArenaTournament) GetTeamMember() string`

GetTeamMember returns the TeamMember field if non-nil, zero value otherwise.

### GetTeamMemberOk

`func (o *ArenaTournament) GetTeamMemberOk() (*string, bool)`

GetTeamMemberOk returns a tuple with the TeamMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMember

`func (o *ArenaTournament) SetTeamMember(v string)`

SetTeamMember sets TeamMember field to given value.

### HasTeamMember

`func (o *ArenaTournament) HasTeamMember() bool`

HasTeamMember returns a boolean if a field has been set.

### GetPrivate

`func (o *ArenaTournament) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ArenaTournament) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ArenaTournament) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ArenaTournament) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPosition

`func (o *ArenaTournament) GetPosition() ArenaPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ArenaTournament) GetPositionOk() (*ArenaPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ArenaTournament) SetPosition(v ArenaPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ArenaTournament) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSchedule

`func (o *ArenaTournament) GetSchedule() ArenaTournamentSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ArenaTournament) GetScheduleOk() (*ArenaTournamentSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ArenaTournament) SetSchedule(v ArenaTournamentSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ArenaTournament) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTeamBattle

`func (o *ArenaTournament) GetTeamBattle() ArenaTournamentTeamBattle`

GetTeamBattle returns the TeamBattle field if non-nil, zero value otherwise.

### GetTeamBattleOk

`func (o *ArenaTournament) GetTeamBattleOk() (*ArenaTournamentTeamBattle, bool)`

GetTeamBattleOk returns a tuple with the TeamBattle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamBattle

`func (o *ArenaTournament) SetTeamBattle(v ArenaTournamentTeamBattle)`

SetTeamBattle sets TeamBattle field to given value.

### HasTeamBattle

`func (o *ArenaTournament) HasTeamBattle() bool`

HasTeamBattle returns a boolean if a field has been set.

### GetWinner

`func (o *ArenaTournament) GetWinner() LightUser`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ArenaTournament) GetWinnerOk() (*LightUser, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ArenaTournament) SetWinner(v LightUser)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ArenaTournament) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


