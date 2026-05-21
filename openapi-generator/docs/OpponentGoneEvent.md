# OpponentGoneEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Gone** | **bool** |  | 
**ClaimWinInSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewOpponentGoneEvent

`func NewOpponentGoneEvent(type_ string, gone bool, ) *OpponentGoneEvent`

NewOpponentGoneEvent instantiates a new OpponentGoneEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpponentGoneEventWithDefaults

`func NewOpponentGoneEventWithDefaults() *OpponentGoneEvent`

NewOpponentGoneEventWithDefaults instantiates a new OpponentGoneEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *OpponentGoneEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OpponentGoneEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OpponentGoneEvent) SetType(v string)`

SetType sets Type field to given value.


### GetGone

`func (o *OpponentGoneEvent) GetGone() bool`

GetGone returns the Gone field if non-nil, zero value otherwise.

### GetGoneOk

`func (o *OpponentGoneEvent) GetGoneOk() (*bool, bool)`

GetGoneOk returns a tuple with the Gone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGone

`func (o *OpponentGoneEvent) SetGone(v bool)`

SetGone sets Gone field to given value.


### GetClaimWinInSeconds

`func (o *OpponentGoneEvent) GetClaimWinInSeconds() int32`

GetClaimWinInSeconds returns the ClaimWinInSeconds field if non-nil, zero value otherwise.

### GetClaimWinInSecondsOk

`func (o *OpponentGoneEvent) GetClaimWinInSecondsOk() (*int32, bool)`

GetClaimWinInSecondsOk returns a tuple with the ClaimWinInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimWinInSeconds

`func (o *OpponentGoneEvent) SetClaimWinInSeconds(v int32)`

SetClaimWinInSeconds sets ClaimWinInSeconds field to given value.

### HasClaimWinInSeconds

`func (o *OpponentGoneEvent) HasClaimWinInSeconds() bool`

HasClaimWinInSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


