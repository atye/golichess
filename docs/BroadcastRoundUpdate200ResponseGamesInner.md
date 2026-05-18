# BroadcastRoundUpdate200ResponseGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]BroadcastRoundUpdate200ResponseGamesInnerPlayersInner**](BroadcastRoundUpdate200ResponseGamesInnerPlayersInner.md) |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Check** | Pointer to **string** |  | [optional] 
**ThinkTime** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | The result of the game | [optional] 

## Methods

### NewBroadcastRoundUpdate200ResponseGamesInner

`func NewBroadcastRoundUpdate200ResponseGamesInner(id string, name string, ) *BroadcastRoundUpdate200ResponseGamesInner`

NewBroadcastRoundUpdate200ResponseGamesInner instantiates a new BroadcastRoundUpdate200ResponseGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundUpdate200ResponseGamesInnerWithDefaults

`func NewBroadcastRoundUpdate200ResponseGamesInnerWithDefaults() *BroadcastRoundUpdate200ResponseGamesInner`

NewBroadcastRoundUpdate200ResponseGamesInnerWithDefaults instantiates a new BroadcastRoundUpdate200ResponseGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetName(v string)`

SetName sets Name field to given value.


### GetFen

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetPlayers

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetPlayers() []BroadcastRoundUpdate200ResponseGamesInnerPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetPlayersOk() (*[]BroadcastRoundUpdate200ResponseGamesInnerPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetPlayers(v []BroadcastRoundUpdate200ResponseGamesInnerPlayersInner)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetLastMove

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetCheck

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetThinkTime

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetThinkTime() int32`

GetThinkTime returns the ThinkTime field if non-nil, zero value otherwise.

### GetThinkTimeOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetThinkTimeOk() (*int32, bool)`

GetThinkTimeOk returns a tuple with the ThinkTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkTime

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetThinkTime(v int32)`

SetThinkTime sets ThinkTime field to given value.

### HasThinkTime

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasThinkTime() bool`

HasThinkTime returns a boolean if a field has been set.

### GetStatus

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BroadcastRoundUpdate200ResponseGamesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BroadcastRoundUpdate200ResponseGamesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BroadcastRoundUpdate200ResponseGamesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


