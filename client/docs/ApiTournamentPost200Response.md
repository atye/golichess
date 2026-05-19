# ApiTournamentPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**FullName** | **string** |  | 
**Rated** | Pointer to **bool** |  | [optional] 
**Spotlight** | Pointer to [**ApiTournamentPost200ResponseSpotlight**](ApiTournamentPost200ResponseSpotlight.md) |  | [optional] 
**Berserkable** | Pointer to **bool** |  | [optional] 
**OnlyTitled** | Pointer to **bool** |  | [optional] 
**Clock** | [**ApiTournament200ResponseCreatedInnerClock**](ApiTournament200ResponseCreatedInnerClock.md) |  | 
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
**Verdicts** | Pointer to [**ApiTournamentPost200ResponseVerdicts**](ApiTournamentPost200ResponseVerdicts.md) |  | [optional] 
**Quote** | Pointer to [**ApiTournamentPost200ResponseQuote**](ApiTournamentPost200ResponseQuote.md) |  | [optional] 
**GreatPlayer** | Pointer to [**ApiTournamentPost200ResponseGreatPlayer**](ApiTournamentPost200ResponseGreatPlayer.md) |  | [optional] 
**AllowList** | Pointer to **[]string** | List of usernames allowed to join the tournament | [optional] 
**HasMaxRating** | Pointer to **bool** |  | [optional] 
**MaxRating** | Pointer to [**ApiTournament200ResponseCreatedInnerMaxRating**](ApiTournament200ResponseCreatedInnerMaxRating.md) |  | [optional] 
**MinRating** | Pointer to [**ApiTournament200ResponseCreatedInnerMaxRating**](ApiTournament200ResponseCreatedInnerMaxRating.md) |  | [optional] 
**MinRatedGames** | Pointer to [**ApiTournament200ResponseCreatedInnerMinRatedGames**](ApiTournament200ResponseCreatedInnerMinRatedGames.md) |  | [optional] 
**BotsAllowed** | Pointer to **bool** |  | [optional] 
**MinAccountAgeInDays** | Pointer to **int32** |  | [optional] 
**Perf** | Pointer to [**ApiTournamentPost200ResponsePerf**](ApiTournamentPost200ResponsePerf.md) |  | [optional] 
**Schedule** | Pointer to [**ApiTournamentPost200ResponseSchedule**](ApiTournamentPost200ResponseSchedule.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Variant** | Pointer to **string** |  | [optional] 
**Duels** | Pointer to [**[]ApiTournamentPost200ResponseDuelsInner**](ApiTournamentPost200ResponseDuelsInner.md) |  | [optional] 
**Standing** | Pointer to [**ApiTournamentPost200ResponseStanding**](ApiTournamentPost200ResponseStanding.md) |  | [optional] 
**Featured** | Pointer to [**ApiTournamentPost200ResponseFeatured**](ApiTournamentPost200ResponseFeatured.md) |  | [optional] 
**Podium** | Pointer to [**[]ApiTournamentPost200ResponsePodiumInner**](ApiTournamentPost200ResponsePodiumInner.md) |  | [optional] 
**Stats** | Pointer to [**ApiTournamentPost200ResponseStats**](ApiTournamentPost200ResponseStats.md) |  | [optional] 
**MyUsername** | Pointer to **string** |  | [optional] 

## Methods

### NewApiTournamentPost200Response

`func NewApiTournamentPost200Response(id string, fullName string, clock ApiTournament200ResponseCreatedInnerClock, nbPlayers int32, ) *ApiTournamentPost200Response`

NewApiTournamentPost200Response instantiates a new ApiTournamentPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiTournamentPost200ResponseWithDefaults

`func NewApiTournamentPost200ResponseWithDefaults() *ApiTournamentPost200Response`

NewApiTournamentPost200ResponseWithDefaults instantiates a new ApiTournamentPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiTournamentPost200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiTournamentPost200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiTournamentPost200Response) SetId(v string)`

SetId sets Id field to given value.


### GetFullName

`func (o *ApiTournamentPost200Response) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *ApiTournamentPost200Response) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *ApiTournamentPost200Response) SetFullName(v string)`

SetFullName sets FullName field to given value.


### GetRated

`func (o *ApiTournamentPost200Response) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiTournamentPost200Response) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiTournamentPost200Response) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ApiTournamentPost200Response) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetSpotlight

`func (o *ApiTournamentPost200Response) GetSpotlight() ApiTournamentPost200ResponseSpotlight`

GetSpotlight returns the Spotlight field if non-nil, zero value otherwise.

### GetSpotlightOk

`func (o *ApiTournamentPost200Response) GetSpotlightOk() (*ApiTournamentPost200ResponseSpotlight, bool)`

GetSpotlightOk returns a tuple with the Spotlight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpotlight

`func (o *ApiTournamentPost200Response) SetSpotlight(v ApiTournamentPost200ResponseSpotlight)`

SetSpotlight sets Spotlight field to given value.

### HasSpotlight

`func (o *ApiTournamentPost200Response) HasSpotlight() bool`

HasSpotlight returns a boolean if a field has been set.

### GetBerserkable

`func (o *ApiTournamentPost200Response) GetBerserkable() bool`

GetBerserkable returns the Berserkable field if non-nil, zero value otherwise.

### GetBerserkableOk

`func (o *ApiTournamentPost200Response) GetBerserkableOk() (*bool, bool)`

GetBerserkableOk returns a tuple with the Berserkable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBerserkable

`func (o *ApiTournamentPost200Response) SetBerserkable(v bool)`

SetBerserkable sets Berserkable field to given value.

### HasBerserkable

`func (o *ApiTournamentPost200Response) HasBerserkable() bool`

HasBerserkable returns a boolean if a field has been set.

### GetOnlyTitled

`func (o *ApiTournamentPost200Response) GetOnlyTitled() bool`

GetOnlyTitled returns the OnlyTitled field if non-nil, zero value otherwise.

### GetOnlyTitledOk

`func (o *ApiTournamentPost200Response) GetOnlyTitledOk() (*bool, bool)`

GetOnlyTitledOk returns a tuple with the OnlyTitled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyTitled

`func (o *ApiTournamentPost200Response) SetOnlyTitled(v bool)`

SetOnlyTitled sets OnlyTitled field to given value.

### HasOnlyTitled

`func (o *ApiTournamentPost200Response) HasOnlyTitled() bool`

HasOnlyTitled returns a boolean if a field has been set.

### GetClock

`func (o *ApiTournamentPost200Response) GetClock() ApiTournament200ResponseCreatedInnerClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiTournamentPost200Response) GetClockOk() (*ApiTournament200ResponseCreatedInnerClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiTournamentPost200Response) SetClock(v ApiTournament200ResponseCreatedInnerClock)`

SetClock sets Clock field to given value.


### GetMinutes

`func (o *ApiTournamentPost200Response) GetMinutes() int32`

GetMinutes returns the Minutes field if non-nil, zero value otherwise.

### GetMinutesOk

`func (o *ApiTournamentPost200Response) GetMinutesOk() (*int32, bool)`

GetMinutesOk returns a tuple with the Minutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinutes

`func (o *ApiTournamentPost200Response) SetMinutes(v int32)`

SetMinutes sets Minutes field to given value.

### HasMinutes

`func (o *ApiTournamentPost200Response) HasMinutes() bool`

HasMinutes returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ApiTournamentPost200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiTournamentPost200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiTournamentPost200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ApiTournamentPost200Response) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetSystem

`func (o *ApiTournamentPost200Response) GetSystem() string`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *ApiTournamentPost200Response) GetSystemOk() (*string, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *ApiTournamentPost200Response) SetSystem(v string)`

SetSystem sets System field to given value.

### HasSystem

`func (o *ApiTournamentPost200Response) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetSecondsToStart

`func (o *ApiTournamentPost200Response) GetSecondsToStart() int32`

GetSecondsToStart returns the SecondsToStart field if non-nil, zero value otherwise.

### GetSecondsToStartOk

`func (o *ApiTournamentPost200Response) GetSecondsToStartOk() (*int32, bool)`

GetSecondsToStartOk returns a tuple with the SecondsToStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToStart

`func (o *ApiTournamentPost200Response) SetSecondsToStart(v int32)`

SetSecondsToStart sets SecondsToStart field to given value.

### HasSecondsToStart

`func (o *ApiTournamentPost200Response) HasSecondsToStart() bool`

HasSecondsToStart returns a boolean if a field has been set.

### GetSecondsToFinish

`func (o *ApiTournamentPost200Response) GetSecondsToFinish() int32`

GetSecondsToFinish returns the SecondsToFinish field if non-nil, zero value otherwise.

### GetSecondsToFinishOk

`func (o *ApiTournamentPost200Response) GetSecondsToFinishOk() (*int32, bool)`

GetSecondsToFinishOk returns a tuple with the SecondsToFinish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsToFinish

`func (o *ApiTournamentPost200Response) SetSecondsToFinish(v int32)`

SetSecondsToFinish sets SecondsToFinish field to given value.

### HasSecondsToFinish

`func (o *ApiTournamentPost200Response) HasSecondsToFinish() bool`

HasSecondsToFinish returns a boolean if a field has been set.

### GetIsFinished

`func (o *ApiTournamentPost200Response) GetIsFinished() bool`

GetIsFinished returns the IsFinished field if non-nil, zero value otherwise.

### GetIsFinishedOk

`func (o *ApiTournamentPost200Response) GetIsFinishedOk() (*bool, bool)`

GetIsFinishedOk returns a tuple with the IsFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinished

`func (o *ApiTournamentPost200Response) SetIsFinished(v bool)`

SetIsFinished sets IsFinished field to given value.

### HasIsFinished

`func (o *ApiTournamentPost200Response) HasIsFinished() bool`

HasIsFinished returns a boolean if a field has been set.

### GetIsRecentlyFinished

`func (o *ApiTournamentPost200Response) GetIsRecentlyFinished() bool`

GetIsRecentlyFinished returns the IsRecentlyFinished field if non-nil, zero value otherwise.

### GetIsRecentlyFinishedOk

`func (o *ApiTournamentPost200Response) GetIsRecentlyFinishedOk() (*bool, bool)`

GetIsRecentlyFinishedOk returns a tuple with the IsRecentlyFinished field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRecentlyFinished

`func (o *ApiTournamentPost200Response) SetIsRecentlyFinished(v bool)`

SetIsRecentlyFinished sets IsRecentlyFinished field to given value.

### HasIsRecentlyFinished

`func (o *ApiTournamentPost200Response) HasIsRecentlyFinished() bool`

HasIsRecentlyFinished returns a boolean if a field has been set.

### GetPairingsClosed

`func (o *ApiTournamentPost200Response) GetPairingsClosed() bool`

GetPairingsClosed returns the PairingsClosed field if non-nil, zero value otherwise.

### GetPairingsClosedOk

`func (o *ApiTournamentPost200Response) GetPairingsClosedOk() (*bool, bool)`

GetPairingsClosedOk returns a tuple with the PairingsClosed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingsClosed

`func (o *ApiTournamentPost200Response) SetPairingsClosed(v bool)`

SetPairingsClosed sets PairingsClosed field to given value.

### HasPairingsClosed

`func (o *ApiTournamentPost200Response) HasPairingsClosed() bool`

HasPairingsClosed returns a boolean if a field has been set.

### GetStartsAt

`func (o *ApiTournamentPost200Response) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ApiTournamentPost200Response) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ApiTournamentPost200Response) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.

### HasStartsAt

`func (o *ApiTournamentPost200Response) HasStartsAt() bool`

HasStartsAt returns a boolean if a field has been set.

### GetNbPlayers

`func (o *ApiTournamentPost200Response) GetNbPlayers() int32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *ApiTournamentPost200Response) GetNbPlayersOk() (*int32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *ApiTournamentPost200Response) SetNbPlayers(v int32)`

SetNbPlayers sets NbPlayers field to given value.


### GetVerdicts

`func (o *ApiTournamentPost200Response) GetVerdicts() ApiTournamentPost200ResponseVerdicts`

GetVerdicts returns the Verdicts field if non-nil, zero value otherwise.

### GetVerdictsOk

`func (o *ApiTournamentPost200Response) GetVerdictsOk() (*ApiTournamentPost200ResponseVerdicts, bool)`

GetVerdictsOk returns a tuple with the Verdicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdicts

`func (o *ApiTournamentPost200Response) SetVerdicts(v ApiTournamentPost200ResponseVerdicts)`

SetVerdicts sets Verdicts field to given value.

### HasVerdicts

`func (o *ApiTournamentPost200Response) HasVerdicts() bool`

HasVerdicts returns a boolean if a field has been set.

### GetQuote

`func (o *ApiTournamentPost200Response) GetQuote() ApiTournamentPost200ResponseQuote`

GetQuote returns the Quote field if non-nil, zero value otherwise.

### GetQuoteOk

`func (o *ApiTournamentPost200Response) GetQuoteOk() (*ApiTournamentPost200ResponseQuote, bool)`

GetQuoteOk returns a tuple with the Quote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuote

`func (o *ApiTournamentPost200Response) SetQuote(v ApiTournamentPost200ResponseQuote)`

SetQuote sets Quote field to given value.

### HasQuote

`func (o *ApiTournamentPost200Response) HasQuote() bool`

HasQuote returns a boolean if a field has been set.

### GetGreatPlayer

`func (o *ApiTournamentPost200Response) GetGreatPlayer() ApiTournamentPost200ResponseGreatPlayer`

GetGreatPlayer returns the GreatPlayer field if non-nil, zero value otherwise.

### GetGreatPlayerOk

`func (o *ApiTournamentPost200Response) GetGreatPlayerOk() (*ApiTournamentPost200ResponseGreatPlayer, bool)`

GetGreatPlayerOk returns a tuple with the GreatPlayer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreatPlayer

`func (o *ApiTournamentPost200Response) SetGreatPlayer(v ApiTournamentPost200ResponseGreatPlayer)`

SetGreatPlayer sets GreatPlayer field to given value.

### HasGreatPlayer

`func (o *ApiTournamentPost200Response) HasGreatPlayer() bool`

HasGreatPlayer returns a boolean if a field has been set.

### GetAllowList

`func (o *ApiTournamentPost200Response) GetAllowList() []string`

GetAllowList returns the AllowList field if non-nil, zero value otherwise.

### GetAllowListOk

`func (o *ApiTournamentPost200Response) GetAllowListOk() (*[]string, bool)`

GetAllowListOk returns a tuple with the AllowList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowList

`func (o *ApiTournamentPost200Response) SetAllowList(v []string)`

SetAllowList sets AllowList field to given value.

### HasAllowList

`func (o *ApiTournamentPost200Response) HasAllowList() bool`

HasAllowList returns a boolean if a field has been set.

### GetHasMaxRating

`func (o *ApiTournamentPost200Response) GetHasMaxRating() bool`

GetHasMaxRating returns the HasMaxRating field if non-nil, zero value otherwise.

### GetHasMaxRatingOk

`func (o *ApiTournamentPost200Response) GetHasMaxRatingOk() (*bool, bool)`

GetHasMaxRatingOk returns a tuple with the HasMaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMaxRating

`func (o *ApiTournamentPost200Response) SetHasMaxRating(v bool)`

SetHasMaxRating sets HasMaxRating field to given value.

### HasHasMaxRating

`func (o *ApiTournamentPost200Response) HasHasMaxRating() bool`

HasHasMaxRating returns a boolean if a field has been set.

### GetMaxRating

`func (o *ApiTournamentPost200Response) GetMaxRating() ApiTournament200ResponseCreatedInnerMaxRating`

GetMaxRating returns the MaxRating field if non-nil, zero value otherwise.

### GetMaxRatingOk

`func (o *ApiTournamentPost200Response) GetMaxRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool)`

GetMaxRatingOk returns a tuple with the MaxRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRating

`func (o *ApiTournamentPost200Response) SetMaxRating(v ApiTournament200ResponseCreatedInnerMaxRating)`

SetMaxRating sets MaxRating field to given value.

### HasMaxRating

`func (o *ApiTournamentPost200Response) HasMaxRating() bool`

HasMaxRating returns a boolean if a field has been set.

### GetMinRating

`func (o *ApiTournamentPost200Response) GetMinRating() ApiTournament200ResponseCreatedInnerMaxRating`

GetMinRating returns the MinRating field if non-nil, zero value otherwise.

### GetMinRatingOk

`func (o *ApiTournamentPost200Response) GetMinRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool)`

GetMinRatingOk returns a tuple with the MinRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRating

`func (o *ApiTournamentPost200Response) SetMinRating(v ApiTournament200ResponseCreatedInnerMaxRating)`

SetMinRating sets MinRating field to given value.

### HasMinRating

`func (o *ApiTournamentPost200Response) HasMinRating() bool`

HasMinRating returns a boolean if a field has been set.

### GetMinRatedGames

`func (o *ApiTournamentPost200Response) GetMinRatedGames() ApiTournament200ResponseCreatedInnerMinRatedGames`

GetMinRatedGames returns the MinRatedGames field if non-nil, zero value otherwise.

### GetMinRatedGamesOk

`func (o *ApiTournamentPost200Response) GetMinRatedGamesOk() (*ApiTournament200ResponseCreatedInnerMinRatedGames, bool)`

GetMinRatedGamesOk returns a tuple with the MinRatedGames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinRatedGames

`func (o *ApiTournamentPost200Response) SetMinRatedGames(v ApiTournament200ResponseCreatedInnerMinRatedGames)`

SetMinRatedGames sets MinRatedGames field to given value.

### HasMinRatedGames

`func (o *ApiTournamentPost200Response) HasMinRatedGames() bool`

HasMinRatedGames returns a boolean if a field has been set.

### GetBotsAllowed

`func (o *ApiTournamentPost200Response) GetBotsAllowed() bool`

GetBotsAllowed returns the BotsAllowed field if non-nil, zero value otherwise.

### GetBotsAllowedOk

`func (o *ApiTournamentPost200Response) GetBotsAllowedOk() (*bool, bool)`

GetBotsAllowedOk returns a tuple with the BotsAllowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBotsAllowed

`func (o *ApiTournamentPost200Response) SetBotsAllowed(v bool)`

SetBotsAllowed sets BotsAllowed field to given value.

### HasBotsAllowed

`func (o *ApiTournamentPost200Response) HasBotsAllowed() bool`

HasBotsAllowed returns a boolean if a field has been set.

### GetMinAccountAgeInDays

`func (o *ApiTournamentPost200Response) GetMinAccountAgeInDays() int32`

GetMinAccountAgeInDays returns the MinAccountAgeInDays field if non-nil, zero value otherwise.

### GetMinAccountAgeInDaysOk

`func (o *ApiTournamentPost200Response) GetMinAccountAgeInDaysOk() (*int32, bool)`

GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAccountAgeInDays

`func (o *ApiTournamentPost200Response) SetMinAccountAgeInDays(v int32)`

SetMinAccountAgeInDays sets MinAccountAgeInDays field to given value.

### HasMinAccountAgeInDays

`func (o *ApiTournamentPost200Response) HasMinAccountAgeInDays() bool`

HasMinAccountAgeInDays returns a boolean if a field has been set.

### GetPerf

`func (o *ApiTournamentPost200Response) GetPerf() ApiTournamentPost200ResponsePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiTournamentPost200Response) GetPerfOk() (*ApiTournamentPost200ResponsePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiTournamentPost200Response) SetPerf(v ApiTournamentPost200ResponsePerf)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ApiTournamentPost200Response) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetSchedule

`func (o *ApiTournamentPost200Response) GetSchedule() ApiTournamentPost200ResponseSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *ApiTournamentPost200Response) GetScheduleOk() (*ApiTournamentPost200ResponseSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *ApiTournamentPost200Response) SetSchedule(v ApiTournamentPost200ResponseSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *ApiTournamentPost200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetDescription

`func (o *ApiTournamentPost200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiTournamentPost200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiTournamentPost200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiTournamentPost200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVariant

`func (o *ApiTournamentPost200Response) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiTournamentPost200Response) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiTournamentPost200Response) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ApiTournamentPost200Response) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetDuels

`func (o *ApiTournamentPost200Response) GetDuels() []ApiTournamentPost200ResponseDuelsInner`

GetDuels returns the Duels field if non-nil, zero value otherwise.

### GetDuelsOk

`func (o *ApiTournamentPost200Response) GetDuelsOk() (*[]ApiTournamentPost200ResponseDuelsInner, bool)`

GetDuelsOk returns a tuple with the Duels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuels

`func (o *ApiTournamentPost200Response) SetDuels(v []ApiTournamentPost200ResponseDuelsInner)`

SetDuels sets Duels field to given value.

### HasDuels

`func (o *ApiTournamentPost200Response) HasDuels() bool`

HasDuels returns a boolean if a field has been set.

### GetStanding

`func (o *ApiTournamentPost200Response) GetStanding() ApiTournamentPost200ResponseStanding`

GetStanding returns the Standing field if non-nil, zero value otherwise.

### GetStandingOk

`func (o *ApiTournamentPost200Response) GetStandingOk() (*ApiTournamentPost200ResponseStanding, bool)`

GetStandingOk returns a tuple with the Standing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStanding

`func (o *ApiTournamentPost200Response) SetStanding(v ApiTournamentPost200ResponseStanding)`

SetStanding sets Standing field to given value.

### HasStanding

`func (o *ApiTournamentPost200Response) HasStanding() bool`

HasStanding returns a boolean if a field has been set.

### GetFeatured

`func (o *ApiTournamentPost200Response) GetFeatured() ApiTournamentPost200ResponseFeatured`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *ApiTournamentPost200Response) GetFeaturedOk() (*ApiTournamentPost200ResponseFeatured, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *ApiTournamentPost200Response) SetFeatured(v ApiTournamentPost200ResponseFeatured)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *ApiTournamentPost200Response) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPodium

`func (o *ApiTournamentPost200Response) GetPodium() []ApiTournamentPost200ResponsePodiumInner`

GetPodium returns the Podium field if non-nil, zero value otherwise.

### GetPodiumOk

`func (o *ApiTournamentPost200Response) GetPodiumOk() (*[]ApiTournamentPost200ResponsePodiumInner, bool)`

GetPodiumOk returns a tuple with the Podium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodium

`func (o *ApiTournamentPost200Response) SetPodium(v []ApiTournamentPost200ResponsePodiumInner)`

SetPodium sets Podium field to given value.

### HasPodium

`func (o *ApiTournamentPost200Response) HasPodium() bool`

HasPodium returns a boolean if a field has been set.

### GetStats

`func (o *ApiTournamentPost200Response) GetStats() ApiTournamentPost200ResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ApiTournamentPost200Response) GetStatsOk() (*ApiTournamentPost200ResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ApiTournamentPost200Response) SetStats(v ApiTournamentPost200ResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ApiTournamentPost200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetMyUsername

`func (o *ApiTournamentPost200Response) GetMyUsername() string`

GetMyUsername returns the MyUsername field if non-nil, zero value otherwise.

### GetMyUsernameOk

`func (o *ApiTournamentPost200Response) GetMyUsernameOk() (*string, bool)`

GetMyUsernameOk returns a tuple with the MyUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMyUsername

`func (o *ApiTournamentPost200Response) SetMyUsername(v string)`

SetMyUsername sets MyUsername field to given value.

### HasMyUsername

`func (o *ApiTournamentPost200Response) HasMyUsername() bool`

HasMyUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


