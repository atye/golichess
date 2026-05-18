# ApiTournament200ResponseCreatedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**System** | **string** |  | 
**Minutes** | **int32** |  | 
**Clock** | [**ApiTournament200ResponseCreatedInnerClock**](ApiTournament200ResponseCreatedInnerClock.md) |  | 
**Rated** | **bool** |  | 
**FullName** | **string** |  | 
**NbPlayers** | **int32** |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**StartsAt** | **int64** |  | 
**FinishesAt** | **int64** |  | 
**Status** | **int32** | 10: created, 20: started, 30: finished  | 
**Perf** | [**ApiTournament200ResponseCreatedInnerPerf**](ApiTournament200ResponseCreatedInnerPerf.md) |  | 
**SecondsToStart** | Pointer to **int32** |  | [optional] 
**HasMaxRating** | Pointer to **bool** |  | [optional] 
**MaxRating** | Pointer to [**ApiTournament200ResponseCreatedInnerMaxRating**](ApiTournament200ResponseCreatedInnerMaxRating.md) |  | [optional] 
**MinRating** | Pointer to [**ApiTournament200ResponseCreatedInnerMaxRating**](ApiTournament200ResponseCreatedInnerMaxRating.md) |  | [optional] 
**MinRatedGames** | Pointer to [**ApiTournament200ResponseCreatedInnerMinRatedGames**](ApiTournament200ResponseCreatedInnerMinRatedGames.md) |  | [optional] 
**BotsAllowed** | Pointer to **bool** |  | [optional] 
**MinAccountAgeInDays** | Pointer to **int32** |  | [optional] 
**OnlyTitled** | Pointer to **bool** |  | [optional] 
**TeamMember** | Pointer to **string** |  | [optional] 
**Private** | Pointer to **bool** |  | [optional] 
**Position** | Pointer to [**ApiTournament200ResponseCreatedInnerPosition**](ApiTournament200ResponseCreatedInnerPosition.md) |  | [optional] 
**Schedule** | Pointer to [**ApiTournament200ResponseCreatedInnerSchedule**](ApiTournament200ResponseCreatedInnerSchedule.md) |  | [optional] 
**TeamBattle** | Pointer to [**ApiTournament200ResponseCreatedInnerTeamBattle**](ApiTournament200ResponseCreatedInnerTeamBattle.md) |  | [optional] 
**Winner** | Pointer to [**ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | [optional] 

## Methods

### NewApiTournament200ResponseCreatedInner

`func NewApiTournament200ResponseCreatedInner(id string, createdBy string, system string, minutes int32, clock ApiTournament200ResponseCreatedInnerClock, rated bool, fullName string, nbPlayers int32, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, startsAt int64, finishesAt int64, status int32, perf ApiTournament200ResponseCreatedInnerPerf, ) *ApiTournament200ResponseCreatedInner`

NewApiTournament200ResponseCreatedInner instantiates a new ApiTournament200ResponseCreatedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTournament200ResponseCreatedInnerWithDefaults

`func NewApiTournament200ResponseCreatedInnerWithDefaults() *ApiTournament200ResponseCreatedInner`

NewApiTournament200ResponseCreatedInnerWithDefaults instantiates a new ApiTournament200ResponseCreatedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiTournament200ResponseCreatedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiTournament200ResponseCreatedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiTournament200ResponseCreatedInner) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ApiTournament200ResponseCreatedInner) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiTournament200ResponseCreatedInner) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiTournament200ResponseCreatedInner) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetSystem

`func (o *ApiTournament200ResponseCreatedInner) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ApiTournament200ResponseCreatedInner) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ApiTournament200ResponseCreatedInner) SetSystem(v string)`

SetSystem sets System field to given value.


### GetMinutes

`func (o *ApiTournament200ResponseCreatedInner) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *ApiTournament200ResponseCreatedInner) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *ApiTournament200ResponseCreatedInner) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.


### GetClock

`func (o *ApiTournament200ResponseCreatedInner) GetClock() ApiTournament200ResponseCreatedInnerClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiTournament200ResponseCreatedInner) GetClockOk() (*ApiTournament200ResponseCreatedInnerClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiTournament200ResponseCreatedInner) SetClock(v ApiTournament200ResponseCreatedInnerClock)`

SetClock sets Clock field to given value.


### GetRated

`func (o *ApiTournament200ResponseCreatedInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiTournament200ResponseCreatedInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiTournament200ResponseCreatedInner) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetFullName

`func (o *ApiTournament200ResponseCreatedInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApiTournament200ResponseCreatedInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApiTournament200ResponseCreatedInner) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetNbPlayers

`func (o *ApiTournament200ResponseCreatedInner) GetNbPlayers() int32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *ApiTournament200ResponseCreatedInner) GetNbPlayersOk() (*int32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *ApiTournament200ResponseCreatedInner) SetNbPlayers(v int32)`

SetNbPlayers sets NbPlayers field to given value.


### GetVariant

`func (o *ApiTournament200ResponseCreatedInner) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiTournament200ResponseCreatedInner) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiTournament200ResponseCreatedInner) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetStartsAt

`func (o *ApiTournament200ResponseCreatedInner) GetStartsAt() int64`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ApiTournament200ResponseCreatedInner) GetStartsAtOk() (*int64, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ApiTournament200ResponseCreatedInner) SetStartsAt(v int64)`

SetStartsAt sets StartsAt field to given value.


### GetFinishesAt

`func (o *ApiTournament200ResponseCreatedInner) GetFinishesAt() int64`

GetFinishesAt returns the FinishesAt field if non-nil, zero value otherwise.

### GetFinishesAtOk

`func (o *ApiTournament200ResponseCreatedInner) GetFinishesAtOk() (*int64, bool)`

GetFinishesAtOk returns a tuple with the FinishesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishesAt

`func (o *ApiTournament200ResponseCreatedInner) SetFinishesAt(v int64)`

SetFinishesAt sets FinishesAt field to given value.


### GetStatus

`func (o *ApiTournament200ResponseCreatedInner) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiTournament200ResponseCreatedInner) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiTournament200ResponseCreatedInner) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetPerf

`func (o *ApiTournament200ResponseCreatedInner) GetPerf() ApiTournament200ResponseCreatedInnerPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiTournament200ResponseCreatedInner) GetPerfOk() (*ApiTournament200ResponseCreatedInnerPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiTournament200ResponseCreatedInner) SetPerf(v ApiTournament200ResponseCreatedInnerPerf)`

SetPerf sets Perf field to given value.


### GetSecondsToStart

`func (o *ApiTournament200ResponseCreatedInner) GetSecondsToStart() int32`

GetSecondsToStart returns the SecondsToStart field if non-nil, zero value otherwise.

### GetSecondsToStartOk

`func (o *ApiTournament200ResponseCreatedInner) GetSecondsToStartOk() (*int32, bool)`

GetSecondsToStartOk returns a tuple with the SecondsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToStart

`func (o *ApiTournament200ResponseCreatedInner) SetSecondsToStart(v int32)`

SetSecondsToStart sets SecondsToStart field to given value.

### HasSecondsToStart

`func (o *ApiTournament200ResponseCreatedInner) HasSecondsToStart() bool`

HasSecondsToStart returns a boolean if a field has been set.

### GetHasMaxRating

`func (o *ApiTournament200ResponseCreatedInner) GetHasMaxRating() bool`

GetHasMaxRating returns the HasMaxRating field if non-nil, zero value otherwise.

### GetHasMaxRatingOk

`func (o *ApiTournament200ResponseCreatedInner) GetHasMaxRatingOk() (*bool, bool)`

GetHasMaxRatingOk returns a tuple with the HasMaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMaxRating

`func (o *ApiTournament200ResponseCreatedInner) SetHasMaxRating(v bool)`

SetHasMaxRating sets HasMaxRating field to given value.

### HasHasMaxRating

`func (o *ApiTournament200ResponseCreatedInner) HasHasMaxRating() bool`

HasHasMaxRating returns a boolean if a field has been set.

### GetMaxRating

`func (o *ApiTournament200ResponseCreatedInner) GetMaxRating() ApiTournament200ResponseCreatedInnerMaxRating`

GetMaxRating returns the MaxRating field if non-nil, zero value otherwise.

### GetMaxRatingOk

`func (o *ApiTournament200ResponseCreatedInner) GetMaxRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool)`

GetMaxRatingOk returns a tuple with the MaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRating

`func (o *ApiTournament200ResponseCreatedInner) SetMaxRating(v ApiTournament200ResponseCreatedInnerMaxRating)`

SetMaxRating sets MaxRating field to given value.

### HasMaxRating

`func (o *ApiTournament200ResponseCreatedInner) HasMaxRating() bool`

HasMaxRating returns a boolean if a field has been set.

### GetMinRating

`func (o *ApiTournament200ResponseCreatedInner) GetMinRating() ApiTournament200ResponseCreatedInnerMaxRating`

GetMinRating returns the MinRating field if non-nil, zero value otherwise.

### GetMinRatingOk

`func (o *ApiTournament200ResponseCreatedInner) GetMinRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool)`

GetMinRatingOk returns a tuple with the MinRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRating

`func (o *ApiTournament200ResponseCreatedInner) SetMinRating(v ApiTournament200ResponseCreatedInnerMaxRating)`

SetMinRating sets MinRating field to given value.

### HasMinRating

`func (o *ApiTournament200ResponseCreatedInner) HasMinRating() bool`

HasMinRating returns a boolean if a field has been set.

### GetMinRatedGames

`func (o *ApiTournament200ResponseCreatedInner) GetMinRatedGames() ApiTournament200ResponseCreatedInnerMinRatedGames`

GetMinRatedGames returns the MinRatedGames field if non-nil, zero value otherwise.

### GetMinRatedGamesOk

`func (o *ApiTournament200ResponseCreatedInner) GetMinRatedGamesOk() (*ApiTournament200ResponseCreatedInnerMinRatedGames, bool)`

GetMinRatedGamesOk returns a tuple with the MinRatedGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRatedGames

`func (o *ApiTournament200ResponseCreatedInner) SetMinRatedGames(v ApiTournament200ResponseCreatedInnerMinRatedGames)`

SetMinRatedGames sets MinRatedGames field to given value.

### HasMinRatedGames

`func (o *ApiTournament200ResponseCreatedInner) HasMinRatedGames() bool`

HasMinRatedGames returns a boolean if a field has been set.

### GetBotsAllowed

`func (o *ApiTournament200ResponseCreatedInner) GetBotsAllowed() bool`

GetBotsAllowed returns the BotsAllowed field if non-nil, zero value otherwise.

### GetBotsAllowedOk

`func (o *ApiTournament200ResponseCreatedInner) GetBotsAllowedOk() (*bool, bool)`

GetBotsAllowedOk returns a tuple with the BotsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotsAllowed

`func (o *ApiTournament200ResponseCreatedInner) SetBotsAllowed(v bool)`

SetBotsAllowed sets BotsAllowed field to given value.

### HasBotsAllowed

`func (o *ApiTournament200ResponseCreatedInner) HasBotsAllowed() bool`

HasBotsAllowed returns a boolean if a field has been set.

### GetMinAccountAgeInDays

`func (o *ApiTournament200ResponseCreatedInner) GetMinAccountAgeInDays() int32`

GetMinAccountAgeInDays returns the MinAccountAgeInDays field if non-nil, zero value otherwise.

### GetMinAccountAgeInDaysOk

`func (o *ApiTournament200ResponseCreatedInner) GetMinAccountAgeInDaysOk() (*int32, bool)`

GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAccountAgeInDays

`func (o *ApiTournament200ResponseCreatedInner) SetMinAccountAgeInDays(v int32)`

SetMinAccountAgeInDays sets MinAccountAgeInDays field to given value.

### HasMinAccountAgeInDays

`func (o *ApiTournament200ResponseCreatedInner) HasMinAccountAgeInDays() bool`

HasMinAccountAgeInDays returns a boolean if a field has been set.

### GetOnlyTitled

`func (o *ApiTournament200ResponseCreatedInner) GetOnlyTitled() bool`

GetOnlyTitled returns the OnlyTitled field if non-nil, zero value otherwise.

### GetOnlyTitledOk

`func (o *ApiTournament200ResponseCreatedInner) GetOnlyTitledOk() (*bool, bool)`

GetOnlyTitledOk returns a tuple with the OnlyTitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyTitled

`func (o *ApiTournament200ResponseCreatedInner) SetOnlyTitled(v bool)`

SetOnlyTitled sets OnlyTitled field to given value.

### HasOnlyTitled

`func (o *ApiTournament200ResponseCreatedInner) HasOnlyTitled() bool`

HasOnlyTitled returns a boolean if a field has been set.

### GetTeamMember

`func (o *ApiTournament200ResponseCreatedInner) GetTeamMember() string`

GetTeamMember returns the TeamMember field if non-nil, zero value otherwise.

### GetTeamMemberOk

`func (o *ApiTournament200ResponseCreatedInner) GetTeamMemberOk() (*string, bool)`

GetTeamMemberOk returns a tuple with the TeamMember field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamMember

`func (o *ApiTournament200ResponseCreatedInner) SetTeamMember(v string)`

SetTeamMember sets TeamMember field to given value.

### HasTeamMember

`func (o *ApiTournament200ResponseCreatedInner) HasTeamMember() bool`

HasTeamMember returns a boolean if a field has been set.

### GetPrivate

`func (o *ApiTournament200ResponseCreatedInner) GetPrivate() bool`

GetPrivate returns the Private field if non-nil, zero value otherwise.

### GetPrivateOk

`func (o *ApiTournament200ResponseCreatedInner) GetPrivateOk() (*bool, bool)`

GetPrivateOk returns a tuple with the Private field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivate

`func (o *ApiTournament200ResponseCreatedInner) SetPrivate(v bool)`

SetPrivate sets Private field to given value.

### HasPrivate

`func (o *ApiTournament200ResponseCreatedInner) HasPrivate() bool`

HasPrivate returns a boolean if a field has been set.

### GetPosition

`func (o *ApiTournament200ResponseCreatedInner) GetPosition() ApiTournament200ResponseCreatedInnerPosition`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ApiTournament200ResponseCreatedInner) GetPositionOk() (*ApiTournament200ResponseCreatedInnerPosition, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ApiTournament200ResponseCreatedInner) SetPosition(v ApiTournament200ResponseCreatedInnerPosition)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *ApiTournament200ResponseCreatedInner) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetSchedule

`func (o *ApiTournament200ResponseCreatedInner) GetSchedule() ApiTournament200ResponseCreatedInnerSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApiTournament200ResponseCreatedInner) GetScheduleOk() (*ApiTournament200ResponseCreatedInnerSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApiTournament200ResponseCreatedInner) SetSchedule(v ApiTournament200ResponseCreatedInnerSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApiTournament200ResponseCreatedInner) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTeamBattle

`func (o *ApiTournament200ResponseCreatedInner) GetTeamBattle() ApiTournament200ResponseCreatedInnerTeamBattle`

GetTeamBattle returns the TeamBattle field if non-nil, zero value otherwise.

### GetTeamBattleOk

`func (o *ApiTournament200ResponseCreatedInner) GetTeamBattleOk() (*ApiTournament200ResponseCreatedInnerTeamBattle, bool)`

GetTeamBattleOk returns a tuple with the TeamBattle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamBattle

`func (o *ApiTournament200ResponseCreatedInner) SetTeamBattle(v ApiTournament200ResponseCreatedInnerTeamBattle)`

SetTeamBattle sets TeamBattle field to given value.

### HasTeamBattle

`func (o *ApiTournament200ResponseCreatedInner) HasTeamBattle() bool`

HasTeamBattle returns a boolean if a field has been set.

### GetWinner

`func (o *ApiTournament200ResponseCreatedInner) GetWinner() ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *ApiTournament200ResponseCreatedInner) GetWinnerOk() (*ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *ApiTournament200ResponseCreatedInner) SetWinner(v ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *ApiTournament200ResponseCreatedInner) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


