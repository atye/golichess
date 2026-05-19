# ApiSwissNew200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedBy** | **string** |  | 
**StartsAt** | **string** |  | 
**Name** | **string** |  | 
**Clock** | [**ApiSwissNew200ResponseClock**](ApiSwissNew200ResponseClock.md) |  | 
**Variant** | **string** |  | 
**Round** | **float32** |  | 
**NbRounds** | **float32** |  | 
**NbPlayers** | **float32** |  | 
**NbOngoing** | **float32** |  | 
**Status** | **string** | The current state of the swiss tournament | 
**Stats** | Pointer to [**ApiSwissNew200ResponseStats**](ApiSwissNew200ResponseStats.md) |  | [optional] 
**Rated** | **bool** |  | 
**Verdicts** | [**ApiTournamentPost200ResponseVerdicts**](ApiTournamentPost200ResponseVerdicts.md) |  | 
**NextRound** | Pointer to [**ApiSwissNew200ResponseNextRound**](ApiSwissNew200ResponseNextRound.md) |  | [optional] 

## Methods

### NewApiSwissNew200Response

`func NewApiSwissNew200Response(id string, createdBy string, startsAt string, name string, clock ApiSwissNew200ResponseClock, variant string, round float32, nbRounds float32, nbPlayers float32, nbOngoing float32, status string, rated bool, verdicts ApiTournamentPost200ResponseVerdicts, ) *ApiSwissNew200Response`

NewApiSwissNew200Response instantiates a new ApiSwissNew200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiSwissNew200ResponseWithDefaults

`func NewApiSwissNew200ResponseWithDefaults() *ApiSwissNew200Response`

NewApiSwissNew200ResponseWithDefaults instantiates a new ApiSwissNew200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiSwissNew200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiSwissNew200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiSwissNew200Response) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *ApiSwissNew200Response) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ApiSwissNew200Response) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ApiSwissNew200Response) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetStartsAt

`func (o *ApiSwissNew200Response) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ApiSwissNew200Response) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ApiSwissNew200Response) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.


### GetName

`func (o *ApiSwissNew200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiSwissNew200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiSwissNew200Response) SetName(v string)`

SetName sets Name field to given value.


### GetClock

`func (o *ApiSwissNew200Response) GetClock() ApiSwissNew200ResponseClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *ApiSwissNew200Response) GetClockOk() (*ApiSwissNew200ResponseClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *ApiSwissNew200Response) SetClock(v ApiSwissNew200ResponseClock)`

SetClock sets Clock field to given value.


### GetVariant

`func (o *ApiSwissNew200Response) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiSwissNew200Response) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiSwissNew200Response) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetRound

`func (o *ApiSwissNew200Response) GetRound() float32`

GetRound returns the Round field if non-nil, zero value otherwise.

### GetRoundOk

`func (o *ApiSwissNew200Response) GetRoundOk() (*float32, bool)`

GetRoundOk returns a tuple with the Round field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRound

`func (o *ApiSwissNew200Response) SetRound(v float32)`

SetRound sets Round field to given value.


### GetNbRounds

`func (o *ApiSwissNew200Response) GetNbRounds() float32`

GetNbRounds returns the NbRounds field if non-nil, zero value otherwise.

### GetNbRoundsOk

`func (o *ApiSwissNew200Response) GetNbRoundsOk() (*float32, bool)`

GetNbRoundsOk returns a tuple with the NbRounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbRounds

`func (o *ApiSwissNew200Response) SetNbRounds(v float32)`

SetNbRounds sets NbRounds field to given value.


### GetNbPlayers

`func (o *ApiSwissNew200Response) GetNbPlayers() float32`

GetNbPlayers returns the NbPlayers field if non-nil, zero value otherwise.

### GetNbPlayersOk

`func (o *ApiSwissNew200Response) GetNbPlayersOk() (*float32, bool)`

GetNbPlayersOk returns a tuple with the NbPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPlayers

`func (o *ApiSwissNew200Response) SetNbPlayers(v float32)`

SetNbPlayers sets NbPlayers field to given value.


### GetNbOngoing

`func (o *ApiSwissNew200Response) GetNbOngoing() float32`

GetNbOngoing returns the NbOngoing field if non-nil, zero value otherwise.

### GetNbOngoingOk

`func (o *ApiSwissNew200Response) GetNbOngoingOk() (*float32, bool)`

GetNbOngoingOk returns a tuple with the NbOngoing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbOngoing

`func (o *ApiSwissNew200Response) SetNbOngoing(v float32)`

SetNbOngoing sets NbOngoing field to given value.


### GetStatus

`func (o *ApiSwissNew200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiSwissNew200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiSwissNew200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStats

`func (o *ApiSwissNew200Response) GetStats() ApiSwissNew200ResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ApiSwissNew200Response) GetStatsOk() (*ApiSwissNew200ResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ApiSwissNew200Response) SetStats(v ApiSwissNew200ResponseStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ApiSwissNew200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetRated

`func (o *ApiSwissNew200Response) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiSwissNew200Response) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiSwissNew200Response) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetVerdicts

`func (o *ApiSwissNew200Response) GetVerdicts() ApiTournamentPost200ResponseVerdicts`

GetVerdicts returns the Verdicts field if non-nil, zero value otherwise.

### GetVerdictsOk

`func (o *ApiSwissNew200Response) GetVerdictsOk() (*ApiTournamentPost200ResponseVerdicts, bool)`

GetVerdictsOk returns a tuple with the Verdicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdicts

`func (o *ApiSwissNew200Response) SetVerdicts(v ApiTournamentPost200ResponseVerdicts)`

SetVerdicts sets Verdicts field to given value.


### GetNextRound

`func (o *ApiSwissNew200Response) GetNextRound() ApiSwissNew200ResponseNextRound`

GetNextRound returns the NextRound field if non-nil, zero value otherwise.

### GetNextRoundOk

`func (o *ApiSwissNew200Response) GetNextRoundOk() (*ApiSwissNew200ResponseNextRound, bool)`

GetNextRoundOk returns a tuple with the NextRound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRound

`func (o *ApiSwissNew200Response) SetNextRound(v ApiSwissNew200ResponseNextRound)`

SetNextRound sets NextRound field to given value.

### HasNextRound

`func (o *ApiSwissNew200Response) HasNextRound() bool`

HasNextRound returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


