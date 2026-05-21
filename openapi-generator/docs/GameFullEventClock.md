# GameFullEventClock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Initial** | Pointer to **int64** | Initial time in milliseconds | [optional] 
**Increment** | Pointer to **int64** | Increment time in milliseconds | [optional] 

## Methods

### NewGameFullEventClock

`func NewGameFullEventClock() *GameFullEventClock`

NewGameFullEventClock instantiates a new GameFullEventClock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameFullEventClockWithDefaults

`func NewGameFullEventClockWithDefaults() *GameFullEventClock`

NewGameFullEventClockWithDefaults instantiates a new GameFullEventClock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitial

`func (o *GameFullEventClock) GetInitial() int64`

GetInitial returns the Initial field if non-nil, zero value otherwise.

### GetInitialOk

`func (o *GameFullEventClock) GetInitialOk() (*int64, bool)`

GetInitialOk returns a tuple with the Initial field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitial

`func (o *GameFullEventClock) SetInitial(v int64)`

SetInitial sets Initial field to given value.

### HasInitial

`func (o *GameFullEventClock) HasInitial() bool`

HasInitial returns a boolean if a field has been set.

### GetIncrement

`func (o *GameFullEventClock) GetIncrement() int64`

GetIncrement returns the Increment field if non-nil, zero value otherwise.

### GetIncrementOk

`func (o *GameFullEventClock) GetIncrementOk() (*int64, bool)`

GetIncrementOk returns a tuple with the Increment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncrement

`func (o *GameFullEventClock) SetIncrement(v int64)`

SetIncrement sets Increment field to given value.

### HasIncrement

`func (o *GameFullEventClock) HasIncrement() bool`

HasIncrement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


