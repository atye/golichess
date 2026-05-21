# Clock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | **int32** |  | 
**Increment** | **int32** |  | 

## Methods

### NewClock

`func NewClock(limit int32, increment int32, ) *Clock`

NewClock instantiates a new Clock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClockWithDefaults

`func NewClockWithDefaults() *Clock`

NewClockWithDefaults instantiates a new Clock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *Clock) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *Clock) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *Clock) SetLimit(v int32)`

SetLimit sets Limit field to given value.


### GetIncrement

`func (o *Clock) GetIncrement() int32`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *Clock) GetIncrementOk() (*int32, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *Clock) SetIncrement(v int32)`

SetIncrement sets Increment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


