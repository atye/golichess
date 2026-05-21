# PerfStat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**PerfStatUser**](PerfStatUser.md) |  | 
**Perf** | [**PerfStatPerf**](PerfStatPerf.md) |  | 
**Rank** | **NullableInt32** |  | 
**Percentile** | **float32** |  | 
**Stat** | [**PerfStatStat**](PerfStatStat.md) |  | 

## Methods

### NewPerfStat

`func NewPerfStat(user PerfStatUser, perf PerfStatPerf, rank NullableInt32, percentile float32, stat PerfStatStat, ) *PerfStat`

NewPerfStat instantiates a new PerfStat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPerfStatWithDefaults

`func NewPerfStatWithDefaults() *PerfStat`

NewPerfStatWithDefaults instantiates a new PerfStat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *PerfStat) GetUser() PerfStatUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *PerfStat) GetUserOk() (*PerfStatUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *PerfStat) SetUser(v PerfStatUser)`

SetUser sets User field to given value.


### GetPerf

`func (o *PerfStat) GetPerf() PerfStatPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *PerfStat) GetPerfOk() (*PerfStatPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *PerfStat) SetPerf(v PerfStatPerf)`

SetPerf sets Perf field to given value.


### GetRank

`func (o *PerfStat) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *PerfStat) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *PerfStat) SetRank(v int32)`

SetRank sets Rank field to given value.


### SetRankNil

`func (o *PerfStat) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *PerfStat) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetPercentile

`func (o *PerfStat) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *PerfStat) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *PerfStat) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.


### GetStat

`func (o *PerfStat) GetStat() PerfStatStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *PerfStat) GetStatOk() (*PerfStatStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *PerfStat) SetStat(v PerfStatStat)`

SetStat sets Stat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


