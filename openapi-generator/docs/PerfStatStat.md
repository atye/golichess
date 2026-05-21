# PerfStatStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highest** | Pointer to [**PerfStatStatHighest**](PerfStatStatHighest.md) |  | [optional] 
**Lowest** | Pointer to [**PerfStatStatHighest**](PerfStatStatHighest.md) |  | [optional] 
**BestWins** | [**PerfStatStatBestWins**](PerfStatStatBestWins.md) |  | 
**WorstLosses** | [**PerfStatStatBestWins**](PerfStatStatBestWins.md) |  | 
**Count** | [**PerfStatStatCount**](PerfStatStatCount.md) |  | 
**ResultStreak** | [**PerfStatStatResultStreak**](PerfStatStatResultStreak.md) |  | 
**PlayStreak** | [**PerfStatStatPlayStreak**](PerfStatStatPlayStreak.md) |  | 

## Methods

### NewPerfStatStat

`func NewPerfStatStat(bestWins PerfStatStatBestWins, worstLosses PerfStatStatBestWins, count PerfStatStatCount, resultStreak PerfStatStatResultStreak, playStreak PerfStatStatPlayStreak, ) *PerfStatStat`

NewPerfStatStat instantiates a new PerfStatStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfStatStatWithDefaults

`func NewPerfStatStatWithDefaults() *PerfStatStat`

NewPerfStatStatWithDefaults instantiates a new PerfStatStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighest

`func (o *PerfStatStat) GetHighest() PerfStatStatHighest`

GetHighest returns the Highest field if non-nil, zero value otherwise.

### GetHighestOk

`func (o *PerfStatStat) GetHighestOk() (*PerfStatStatHighest, bool)`

GetHighestOk returns a tuple with the Highest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighest

`func (o *PerfStatStat) SetHighest(v PerfStatStatHighest)`

SetHighest sets Highest field to given value.

### HasHighest

`func (o *PerfStatStat) HasHighest() bool`

HasHighest returns a boolean if a field has been set.

### GetLowest

`func (o *PerfStatStat) GetLowest() PerfStatStatHighest`

GetLowest returns the Lowest field if non-nil, zero value otherwise.

### GetLowestOk

`func (o *PerfStatStat) GetLowestOk() (*PerfStatStatHighest, bool)`

GetLowestOk returns a tuple with the Lowest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowest

`func (o *PerfStatStat) SetLowest(v PerfStatStatHighest)`

SetLowest sets Lowest field to given value.

### HasLowest

`func (o *PerfStatStat) HasLowest() bool`

HasLowest returns a boolean if a field has been set.

### GetBestWins

`func (o *PerfStatStat) GetBestWins() PerfStatStatBestWins`

GetBestWins returns the BestWins field if non-nil, zero value otherwise.

### GetBestWinsOk

`func (o *PerfStatStat) GetBestWinsOk() (*PerfStatStatBestWins, bool)`

GetBestWinsOk returns a tuple with the BestWins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestWins

`func (o *PerfStatStat) SetBestWins(v PerfStatStatBestWins)`

SetBestWins sets BestWins field to given value.


### GetWorstLosses

`func (o *PerfStatStat) GetWorstLosses() PerfStatStatBestWins`

GetWorstLosses returns the WorstLosses field if non-nil, zero value otherwise.

### GetWorstLossesOk

`func (o *PerfStatStat) GetWorstLossesOk() (*PerfStatStatBestWins, bool)`

GetWorstLossesOk returns a tuple with the WorstLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorstLosses

`func (o *PerfStatStat) SetWorstLosses(v PerfStatStatBestWins)`

SetWorstLosses sets WorstLosses field to given value.


### GetCount

`func (o *PerfStatStat) GetCount() PerfStatStatCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PerfStatStat) GetCountOk() (*PerfStatStatCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PerfStatStat) SetCount(v PerfStatStatCount)`

SetCount sets Count field to given value.


### GetResultStreak

`func (o *PerfStatStat) GetResultStreak() PerfStatStatResultStreak`

GetResultStreak returns the ResultStreak field if non-nil, zero value otherwise.

### GetResultStreakOk

`func (o *PerfStatStat) GetResultStreakOk() (*PerfStatStatResultStreak, bool)`

GetResultStreakOk returns a tuple with the ResultStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultStreak

`func (o *PerfStatStat) SetResultStreak(v PerfStatStatResultStreak)`

SetResultStreak sets ResultStreak field to given value.


### GetPlayStreak

`func (o *PerfStatStat) GetPlayStreak() PerfStatStatPlayStreak`

GetPlayStreak returns the PlayStreak field if non-nil, zero value otherwise.

### GetPlayStreakOk

`func (o *PerfStatStat) GetPlayStreakOk() (*PerfStatStatPlayStreak, bool)`

GetPlayStreakOk returns a tuple with the PlayStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayStreak

`func (o *PerfStatStat) SetPlayStreak(v PerfStatStatPlayStreak)`

SetPlayStreak sets PlayStreak field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


