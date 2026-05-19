# ApiUserPerf200ResponseStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Highest** | Pointer to [**ApiUserPerf200ResponseStatHighest**](ApiUserPerf200ResponseStatHighest.md) |  | [optional] 
**Lowest** | Pointer to [**ApiUserPerf200ResponseStatHighest**](ApiUserPerf200ResponseStatHighest.md) |  | [optional] 
**BestWins** | [**ApiUserPerf200ResponseStatBestWins**](ApiUserPerf200ResponseStatBestWins.md) |  | 
**WorstLosses** | [**ApiUserPerf200ResponseStatWorstLosses**](ApiUserPerf200ResponseStatWorstLosses.md) |  | 
**Count** | [**ApiUserPerf200ResponseStatCount**](ApiUserPerf200ResponseStatCount.md) |  | 
**ResultStreak** | [**ApiUserPerf200ResponseStatResultStreak**](ApiUserPerf200ResponseStatResultStreak.md) |  | 
**PlayStreak** | [**ApiUserPerf200ResponseStatPlayStreak**](ApiUserPerf200ResponseStatPlayStreak.md) |  | 

## Methods

### NewApiUserPerf200ResponseStat

`func NewApiUserPerf200ResponseStat(bestWins ApiUserPerf200ResponseStatBestWins, worstLosses ApiUserPerf200ResponseStatWorstLosses, count ApiUserPerf200ResponseStatCount, resultStreak ApiUserPerf200ResponseStatResultStreak, playStreak ApiUserPerf200ResponseStatPlayStreak, ) *ApiUserPerf200ResponseStat`

NewApiUserPerf200ResponseStat instantiates a new ApiUserPerf200ResponseStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserPerf200ResponseStatWithDefaults

`func NewApiUserPerf200ResponseStatWithDefaults() *ApiUserPerf200ResponseStat`

NewApiUserPerf200ResponseStatWithDefaults instantiates a new ApiUserPerf200ResponseStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHighest

`func (o *ApiUserPerf200ResponseStat) GetHighest() ApiUserPerf200ResponseStatHighest`

GetHighest returns the Highest field if non-nil, zero value otherwise.

### GetHighestOk

`func (o *ApiUserPerf200ResponseStat) GetHighestOk() (*ApiUserPerf200ResponseStatHighest, bool)`

GetHighestOk returns a tuple with the Highest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighest

`func (o *ApiUserPerf200ResponseStat) SetHighest(v ApiUserPerf200ResponseStatHighest)`

SetHighest sets Highest field to given value.

### HasHighest

`func (o *ApiUserPerf200ResponseStat) HasHighest() bool`

HasHighest returns a boolean if a field has been set.

### GetLowest

`func (o *ApiUserPerf200ResponseStat) GetLowest() ApiUserPerf200ResponseStatHighest`

GetLowest returns the Lowest field if non-nil, zero value otherwise.

### GetLowestOk

`func (o *ApiUserPerf200ResponseStat) GetLowestOk() (*ApiUserPerf200ResponseStatHighest, bool)`

GetLowestOk returns a tuple with the Lowest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowest

`func (o *ApiUserPerf200ResponseStat) SetLowest(v ApiUserPerf200ResponseStatHighest)`

SetLowest sets Lowest field to given value.

### HasLowest

`func (o *ApiUserPerf200ResponseStat) HasLowest() bool`

HasLowest returns a boolean if a field has been set.

### GetBestWins

`func (o *ApiUserPerf200ResponseStat) GetBestWins() ApiUserPerf200ResponseStatBestWins`

GetBestWins returns the BestWins field if non-nil, zero value otherwise.

### GetBestWinsOk

`func (o *ApiUserPerf200ResponseStat) GetBestWinsOk() (*ApiUserPerf200ResponseStatBestWins, bool)`

GetBestWinsOk returns a tuple with the BestWins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestWins

`func (o *ApiUserPerf200ResponseStat) SetBestWins(v ApiUserPerf200ResponseStatBestWins)`

SetBestWins sets BestWins field to given value.


### GetWorstLosses

`func (o *ApiUserPerf200ResponseStat) GetWorstLosses() ApiUserPerf200ResponseStatWorstLosses`

GetWorstLosses returns the WorstLosses field if non-nil, zero value otherwise.

### GetWorstLossesOk

`func (o *ApiUserPerf200ResponseStat) GetWorstLossesOk() (*ApiUserPerf200ResponseStatWorstLosses, bool)`

GetWorstLossesOk returns a tuple with the WorstLosses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorstLosses

`func (o *ApiUserPerf200ResponseStat) SetWorstLosses(v ApiUserPerf200ResponseStatWorstLosses)`

SetWorstLosses sets WorstLosses field to given value.


### GetCount

`func (o *ApiUserPerf200ResponseStat) GetCount() ApiUserPerf200ResponseStatCount`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *ApiUserPerf200ResponseStat) GetCountOk() (*ApiUserPerf200ResponseStatCount, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *ApiUserPerf200ResponseStat) SetCount(v ApiUserPerf200ResponseStatCount)`

SetCount sets Count field to given value.


### GetResultStreak

`func (o *ApiUserPerf200ResponseStat) GetResultStreak() ApiUserPerf200ResponseStatResultStreak`

GetResultStreak returns the ResultStreak field if non-nil, zero value otherwise.

### GetResultStreakOk

`func (o *ApiUserPerf200ResponseStat) GetResultStreakOk() (*ApiUserPerf200ResponseStatResultStreak, bool)`

GetResultStreakOk returns a tuple with the ResultStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultStreak

`func (o *ApiUserPerf200ResponseStat) SetResultStreak(v ApiUserPerf200ResponseStatResultStreak)`

SetResultStreak sets ResultStreak field to given value.


### GetPlayStreak

`func (o *ApiUserPerf200ResponseStat) GetPlayStreak() ApiUserPerf200ResponseStatPlayStreak`

GetPlayStreak returns the PlayStreak field if non-nil, zero value otherwise.

### GetPlayStreakOk

`func (o *ApiUserPerf200ResponseStat) GetPlayStreakOk() (*ApiUserPerf200ResponseStatPlayStreak, bool)`

GetPlayStreakOk returns a tuple with the PlayStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayStreak

`func (o *ApiUserPerf200ResponseStat) SetPlayStreak(v ApiUserPerf200ResponseStatPlayStreak)`

SetPlayStreak sets PlayStreak field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


