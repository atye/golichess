# PerfStatStatPlayStreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nb** | [**PerfStatStatResultStreakWin**](PerfStatStatResultStreakWin.md) |  | 
**Time** | [**PerfStatStatResultStreakWin**](PerfStatStatResultStreakWin.md) |  | 
**LastDate** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewPerfStatStatPlayStreak

`func NewPerfStatStatPlayStreak(nb PerfStatStatResultStreakWin, time PerfStatStatResultStreakWin, ) *PerfStatStatPlayStreak`

NewPerfStatStatPlayStreak instantiates a new PerfStatStatPlayStreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfStatStatPlayStreakWithDefaults

`func NewPerfStatStatPlayStreakWithDefaults() *PerfStatStatPlayStreak`

NewPerfStatStatPlayStreakWithDefaults instantiates a new PerfStatStatPlayStreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNb

`func (o *PerfStatStatPlayStreak) GetNb() PerfStatStatResultStreakWin`

GetNb returns the Nb field if non-nil, zero value otherwise.

### GetNbOk

`func (o *PerfStatStatPlayStreak) GetNbOk() (*PerfStatStatResultStreakWin, bool)`

GetNbOk returns a tuple with the Nb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNb

`func (o *PerfStatStatPlayStreak) SetNb(v PerfStatStatResultStreakWin)`

SetNb sets Nb field to given value.


### GetTime

`func (o *PerfStatStatPlayStreak) GetTime() PerfStatStatResultStreakWin`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *PerfStatStatPlayStreak) GetTimeOk() (*PerfStatStatResultStreakWin, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *PerfStatStatPlayStreak) SetTime(v PerfStatStatResultStreakWin)`

SetTime sets Time field to given value.


### GetLastDate

`func (o *PerfStatStatPlayStreak) GetLastDate() time.Time`

GetLastDate returns the LastDate field if non-nil, zero value otherwise.

### GetLastDateOk

`func (o *PerfStatStatPlayStreak) GetLastDateOk() (*time.Time, bool)`

GetLastDateOk returns a tuple with the LastDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDate

`func (o *PerfStatStatPlayStreak) SetLastDate(v time.Time)`

SetLastDate sets LastDate field to given value.

### HasLastDate

`func (o *PerfStatStatPlayStreak) HasLastDate() bool`

HasLastDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


