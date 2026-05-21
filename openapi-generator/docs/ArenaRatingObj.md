# ArenaRatingObj

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Perf** | Pointer to [**PerfType**](PerfType.md) |  | [optional] 
**Rating** | **int32** |  | 

## Methods

### NewArenaRatingObj

`func NewArenaRatingObj(rating int32, ) *ArenaRatingObj`

NewArenaRatingObj instantiates a new ArenaRatingObj object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArenaRatingObjWithDefaults

`func NewArenaRatingObjWithDefaults() *ArenaRatingObj`

NewArenaRatingObjWithDefaults instantiates a new ArenaRatingObj object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPerf

`func (o *ArenaRatingObj) GetPerf() PerfType`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ArenaRatingObj) GetPerfOk() (*PerfType, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ArenaRatingObj) SetPerf(v PerfType)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ArenaRatingObj) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRating

`func (o *ArenaRatingObj) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ArenaRatingObj) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ArenaRatingObj) SetRating(v int32)`

SetRating sets Rating field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


