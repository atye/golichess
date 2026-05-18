# BoardGameStream200ResponseOneOfStateExpiration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdleMillis** | **int32** | Milliseconds since the last move was played, or since the game started | 
**MillisToMove** | **int32** | Time each player has to make their first move, before the game is aborted | 

## Methods

### NewBoardGameStream200ResponseOneOfStateExpiration

`func NewBoardGameStream200ResponseOneOfStateExpiration(idleMillis int32, millisToMove int32, ) *BoardGameStream200ResponseOneOfStateExpiration`

NewBoardGameStream200ResponseOneOfStateExpiration instantiates a new BoardGameStream200ResponseOneOfStateExpiration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardGameStream200ResponseOneOfStateExpirationWithDefaults

`func NewBoardGameStream200ResponseOneOfStateExpirationWithDefaults() *BoardGameStream200ResponseOneOfStateExpiration`

NewBoardGameStream200ResponseOneOfStateExpirationWithDefaults instantiates a new BoardGameStream200ResponseOneOfStateExpiration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdleMillis

`func (o *BoardGameStream200ResponseOneOfStateExpiration) GetIdleMillis() int32`

GetIdleMillis returns the IdleMillis field if non-nil, zero value otherwise.

### GetIdleMillisOk

`func (o *BoardGameStream200ResponseOneOfStateExpiration) GetIdleMillisOk() (*int32, bool)`

GetIdleMillisOk returns a tuple with the IdleMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdleMillis

`func (o *BoardGameStream200ResponseOneOfStateExpiration) SetIdleMillis(v int32)`

SetIdleMillis sets IdleMillis field to given value.


### GetMillisToMove

`func (o *BoardGameStream200ResponseOneOfStateExpiration) GetMillisToMove() int32`

GetMillisToMove returns the MillisToMove field if non-nil, zero value otherwise.

### GetMillisToMoveOk

`func (o *BoardGameStream200ResponseOneOfStateExpiration) GetMillisToMoveOk() (*int32, bool)`

GetMillisToMoveOk returns a tuple with the MillisToMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMillisToMove

`func (o *BoardGameStream200ResponseOneOfStateExpiration) SetMillisToMove(v int32)`

SetMillisToMove sets MillisToMove field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


