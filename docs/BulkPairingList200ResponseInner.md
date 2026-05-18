# BulkPairingList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Games** | [**[]BulkPairingList200ResponseInnerGamesInner**](BulkPairingList200ResponseInnerGamesInner.md) |  | 
**Variant** | **string** |  | [default to "standard"]
**Clock** | [**ApiTournament200ResponseCreatedInnerClock**](ApiTournament200ResponseCreatedInnerClock.md) |  | 
**PairAt** | **int32** |  | 
**PairedAt** | **NullableInt32** |  | 
**Rated** | **bool** |  | 
**StartClocksAt** | **int32** |  | 
**ScheduledAt** | **int32** |  | 

## Methods

### NewBulkPairingList200ResponseInner

`func NewBulkPairingList200ResponseInner(id string, games []BulkPairingList200ResponseInnerGamesInner, variant string, clock ApiTournament200ResponseCreatedInnerClock, pairAt int32, pairedAt NullableInt32, rated bool, startClocksAt int32, scheduledAt int32, ) *BulkPairingList200ResponseInner`

NewBulkPairingList200ResponseInner instantiates a new BulkPairingList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPairingList200ResponseInnerWithDefaults

`func NewBulkPairingList200ResponseInnerWithDefaults() *BulkPairingList200ResponseInner`

NewBulkPairingList200ResponseInnerWithDefaults instantiates a new BulkPairingList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkPairingList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkPairingList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkPairingList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetGames

`func (o *BulkPairingList200ResponseInner) GetGames() []BulkPairingList200ResponseInnerGamesInner`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BulkPairingList200ResponseInner) GetGamesOk() (*[]BulkPairingList200ResponseInnerGamesInner, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BulkPairingList200ResponseInner) SetGames(v []BulkPairingList200ResponseInnerGamesInner)`

SetGames sets Games field to given value.


### GetVariant

`func (o *BulkPairingList200ResponseInner) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *BulkPairingList200ResponseInner) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *BulkPairingList200ResponseInner) SetVariant(v string)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *BulkPairingList200ResponseInner) GetClock() ApiTournament200ResponseCreatedInnerClock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *BulkPairingList200ResponseInner) GetClockOk() (*ApiTournament200ResponseCreatedInnerClock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *BulkPairingList200ResponseInner) SetClock(v ApiTournament200ResponseCreatedInnerClock)`

SetClock sets Clock field to given value.


### GetPairAt

`func (o *BulkPairingList200ResponseInner) GetPairAt() int32`

GetPairAt returns the PairAt field if non-nil, zero value otherwise.

### GetPairAtOk

`func (o *BulkPairingList200ResponseInner) GetPairAtOk() (*int32, bool)`

GetPairAtOk returns a tuple with the PairAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairAt

`func (o *BulkPairingList200ResponseInner) SetPairAt(v int32)`

SetPairAt sets PairAt field to given value.


### GetPairedAt

`func (o *BulkPairingList200ResponseInner) GetPairedAt() int32`

GetPairedAt returns the PairedAt field if non-nil, zero value otherwise.

### GetPairedAtOk

`func (o *BulkPairingList200ResponseInner) GetPairedAtOk() (*int32, bool)`

GetPairedAtOk returns a tuple with the PairedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAt

`func (o *BulkPairingList200ResponseInner) SetPairedAt(v int32)`

SetPairedAt sets PairedAt field to given value.


### SetPairedAtNil

`func (o *BulkPairingList200ResponseInner) SetPairedAtNil(b bool)`

 SetPairedAtNil sets the value for PairedAt to be an explicit nil

### UnsetPairedAt
`func (o *BulkPairingList200ResponseInner) UnsetPairedAt()`

UnsetPairedAt ensures that no value is present for PairedAt, not even an explicit nil
### GetRated

`func (o *BulkPairingList200ResponseInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *BulkPairingList200ResponseInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *BulkPairingList200ResponseInner) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetStartClocksAt

`func (o *BulkPairingList200ResponseInner) GetStartClocksAt() int32`

GetStartClocksAt returns the StartClocksAt field if non-nil, zero value otherwise.

### GetStartClocksAtOk

`func (o *BulkPairingList200ResponseInner) GetStartClocksAtOk() (*int32, bool)`

GetStartClocksAtOk returns a tuple with the StartClocksAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartClocksAt

`func (o *BulkPairingList200ResponseInner) SetStartClocksAt(v int32)`

SetStartClocksAt sets StartClocksAt field to given value.


### GetScheduledAt

`func (o *BulkPairingList200ResponseInner) GetScheduledAt() int32`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *BulkPairingList200ResponseInner) GetScheduledAtOk() (*int32, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *BulkPairingList200ResponseInner) SetScheduledAt(v int32)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


