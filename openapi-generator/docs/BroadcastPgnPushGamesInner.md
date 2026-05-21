# BroadcastPgnPushGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tags** | **map[string]string** |  | 
**Moves** | Pointer to **int32** |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewBroadcastPgnPushGamesInner

`func NewBroadcastPgnPushGamesInner(tags map[string]string, ) *BroadcastPgnPushGamesInner`

NewBroadcastPgnPushGamesInner instantiates a new BroadcastPgnPushGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPgnPushGamesInnerWithDefaults

`func NewBroadcastPgnPushGamesInnerWithDefaults() *BroadcastPgnPushGamesInner`

NewBroadcastPgnPushGamesInnerWithDefaults instantiates a new BroadcastPgnPushGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTags

`func (o *BroadcastPgnPushGamesInner) GetTags() map[string]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BroadcastPgnPushGamesInner) GetTagsOk() (*map[string]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BroadcastPgnPushGamesInner) SetTags(v map[string]string)`

SetTags sets Tags field to given value.


### GetMoves

`func (o *BroadcastPgnPushGamesInner) GetMoves() int32`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *BroadcastPgnPushGamesInner) GetMovesOk() (*int32, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *BroadcastPgnPushGamesInner) SetMoves(v int32)`

SetMoves sets Moves field to given value.

### HasMoves

`func (o *BroadcastPgnPushGamesInner) HasMoves() bool`

HasMoves returns a boolean if a field has been set.

### GetError

`func (o *BroadcastPgnPushGamesInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *BroadcastPgnPushGamesInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *BroadcastPgnPushGamesInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *BroadcastPgnPushGamesInner) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


