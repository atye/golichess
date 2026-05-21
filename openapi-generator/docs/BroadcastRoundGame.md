# BroadcastRoundGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]BroadcastRoundGamePlayersInner**](BroadcastRoundGamePlayersInner.md) |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Check** | Pointer to **string** |  | [optional] 
**ThinkTime** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | The result of the game | [optional] 

## Methods

### NewBroadcastRoundGame

`func NewBroadcastRoundGame(id string, name string, ) *BroadcastRoundGame`

NewBroadcastRoundGame instantiates a new BroadcastRoundGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundGameWithDefaults

`func NewBroadcastRoundGameWithDefaults() *BroadcastRoundGame`

NewBroadcastRoundGameWithDefaults instantiates a new BroadcastRoundGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastRoundGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastRoundGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastRoundGame) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastRoundGame) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastRoundGame) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastRoundGame) SetName(v string)`

SetName sets Name field to given value.


### GetFen

`func (o *BroadcastRoundGame) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *BroadcastRoundGame) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *BroadcastRoundGame) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *BroadcastRoundGame) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetPlayers

`func (o *BroadcastRoundGame) GetPlayers() []BroadcastRoundGamePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastRoundGame) GetPlayersOk() (*[]BroadcastRoundGamePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastRoundGame) SetPlayers(v []BroadcastRoundGamePlayersInner)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *BroadcastRoundGame) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetLastMove

`func (o *BroadcastRoundGame) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *BroadcastRoundGame) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *BroadcastRoundGame) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *BroadcastRoundGame) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetCheck

`func (o *BroadcastRoundGame) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *BroadcastRoundGame) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *BroadcastRoundGame) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *BroadcastRoundGame) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetThinkTime

`func (o *BroadcastRoundGame) GetThinkTime() int32`

GetThinkTime returns the ThinkTime field if non-nil, zero value otherwise.

### GetThinkTimeOk

`func (o *BroadcastRoundGame) GetThinkTimeOk() (*int32, bool)`

GetThinkTimeOk returns a tuple with the ThinkTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkTime

`func (o *BroadcastRoundGame) SetThinkTime(v int32)`

SetThinkTime sets ThinkTime field to given value.

### HasThinkTime

`func (o *BroadcastRoundGame) HasThinkTime() bool`

HasThinkTime returns a boolean if a field has been set.

### GetStatus

`func (o *BroadcastRoundGame) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BroadcastRoundGame) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BroadcastRoundGame) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BroadcastRoundGame) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


