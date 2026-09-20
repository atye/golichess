# BulkPairing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Games** | [**[]BulkPairingGamesInner**](BulkPairingGamesInner.md) |  | 
**Variant** | [**VariantKey**](VariantKey.md) |  | [default to VARIANTKEY_STANDARD]
**Clock** | [**Clock**](Clock.md) |  | 
**PairAt** | **int32** |  | 
**PairedAt** | **NullableInt32** |  | 
**Rated** | **bool** |  | 
**StartClocksAt** | **NullableInt32** |  | 
**ScheduledAt** | **int32** |  | 

## Methods

### NewBulkPairing

`func NewBulkPairing(id string, games []BulkPairingGamesInner, variant VariantKey, clock Clock, pairAt int32, pairedAt NullableInt32, rated bool, startClocksAt NullableInt32, scheduledAt int32, ) *BulkPairing`

NewBulkPairing instantiates a new BulkPairing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPairingWithDefaults

`func NewBulkPairingWithDefaults() *BulkPairing`

NewBulkPairingWithDefaults instantiates a new BulkPairing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BulkPairing) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BulkPairing) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BulkPairing) SetId(v string)`

SetId sets Id field to given value.


### GetGames

`func (o *BulkPairing) GetGames() []BulkPairingGamesInner`

GetGames returns the Games field if non-nil, zero value otherwise.

### GetGamesOk

`func (o *BulkPairing) GetGamesOk() (*[]BulkPairingGamesInner, bool)`

GetGamesOk returns a tuple with the Games field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGames

`func (o *BulkPairing) SetGames(v []BulkPairingGamesInner)`

SetGames sets Games field to given value.


### GetVariant

`func (o *BulkPairing) GetVariant() VariantKey`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *BulkPairing) GetVariantOk() (*VariantKey, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *BulkPairing) SetVariant(v VariantKey)`

SetVariant sets Variant field to given value.


### GetClock

`func (o *BulkPairing) GetClock() Clock`

GetClock returns the Clock field if non-nil, zero value otherwise.

### GetClockOk

`func (o *BulkPairing) GetClockOk() (*Clock, bool)`

GetClockOk returns a tuple with the Clock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClock

`func (o *BulkPairing) SetClock(v Clock)`

SetClock sets Clock field to given value.


### GetPairAt

`func (o *BulkPairing) GetPairAt() int32`

GetPairAt returns the PairAt field if non-nil, zero value otherwise.

### GetPairAtOk

`func (o *BulkPairing) GetPairAtOk() (*int32, bool)`

GetPairAtOk returns a tuple with the PairAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairAt

`func (o *BulkPairing) SetPairAt(v int32)`

SetPairAt sets PairAt field to given value.


### GetPairedAt

`func (o *BulkPairing) GetPairedAt() int32`

GetPairedAt returns the PairedAt field if non-nil, zero value otherwise.

### GetPairedAtOk

`func (o *BulkPairing) GetPairedAtOk() (*int32, bool)`

GetPairedAtOk returns a tuple with the PairedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairedAt

`func (o *BulkPairing) SetPairedAt(v int32)`

SetPairedAt sets PairedAt field to given value.


### SetPairedAtNil

`func (o *BulkPairing) SetPairedAtNil(b bool)`

 SetPairedAtNil sets the value for PairedAt to be an explicit nil

### UnsetPairedAt
`func (o *BulkPairing) UnsetPairedAt()`

UnsetPairedAt ensures that no value is present for PairedAt, not even an explicit nil
### GetRated

`func (o *BulkPairing) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *BulkPairing) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *BulkPairing) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetStartClocksAt

`func (o *BulkPairing) GetStartClocksAt() int32`

GetStartClocksAt returns the StartClocksAt field if non-nil, zero value otherwise.

### GetStartClocksAtOk

`func (o *BulkPairing) GetStartClocksAtOk() (*int32, bool)`

GetStartClocksAtOk returns a tuple with the StartClocksAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartClocksAt

`func (o *BulkPairing) SetStartClocksAt(v int32)`

SetStartClocksAt sets StartClocksAt field to given value.


### SetStartClocksAtNil

`func (o *BulkPairing) SetStartClocksAtNil(b bool)`

 SetStartClocksAtNil sets the value for StartClocksAt to be an explicit nil

### UnsetStartClocksAt
`func (o *BulkPairing) UnsetStartClocksAt()`

UnsetStartClocksAt ensures that no value is present for StartClocksAt, not even an explicit nil
### GetScheduledAt

`func (o *BulkPairing) GetScheduledAt() int32`

GetScheduledAt returns the ScheduledAt field if non-nil, zero value otherwise.

### GetScheduledAtOk

`func (o *BulkPairing) GetScheduledAtOk() (*int32, bool)`

GetScheduledAtOk returns a tuple with the ScheduledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledAt

`func (o *BulkPairing) SetScheduledAt(v int32)`

SetScheduledAt sets ScheduledAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


