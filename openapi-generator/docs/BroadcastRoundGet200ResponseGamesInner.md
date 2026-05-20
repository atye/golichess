# BroadcastRoundGet200ResponseGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Fen** | Pointer to **string** |  | [optional] 
**Players** | Pointer to [**[]BroadcastRoundGet200ResponseGamesInnerPlayersInner**](BroadcastRoundGet200ResponseGamesInnerPlayersInner.md) |  | [optional] 
**LastMove** | Pointer to **string** |  | [optional] 
**Check** | Pointer to **string** |  | [optional] 
**ThinkTime** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | The result of the game | [optional] 

## Methods

### NewBroadcastRoundGet200ResponseGamesInner

`func NewBroadcastRoundGet200ResponseGamesInner(id string, name string, ) *BroadcastRoundGet200ResponseGamesInner`

NewBroadcastRoundGet200ResponseGamesInner instantiates a new BroadcastRoundGet200ResponseGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastRoundGet200ResponseGamesInnerWithDefaults

`func NewBroadcastRoundGet200ResponseGamesInnerWithDefaults() *BroadcastRoundGet200ResponseGamesInner`

NewBroadcastRoundGet200ResponseGamesInnerWithDefaults instantiates a new BroadcastRoundGet200ResponseGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastRoundGet200ResponseGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastRoundGet200ResponseGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *BroadcastRoundGet200ResponseGamesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastRoundGet200ResponseGamesInner) SetName(v string)`

SetName sets Name field to given value.


### GetFen

`func (o *BroadcastRoundGet200ResponseGamesInner) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *BroadcastRoundGet200ResponseGamesInner) SetFen(v string)`

SetFen sets Fen field to given value.

### HasFen

`func (o *BroadcastRoundGet200ResponseGamesInner) HasFen() bool`

HasFen returns a boolean if a field has been set.

### GetPlayers

`func (o *BroadcastRoundGet200ResponseGamesInner) GetPlayers() []BroadcastRoundGet200ResponseGamesInnerPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetPlayersOk() (*[]BroadcastRoundGet200ResponseGamesInnerPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *BroadcastRoundGet200ResponseGamesInner) SetPlayers(v []BroadcastRoundGet200ResponseGamesInnerPlayersInner)`

SetPlayers sets Players field to given value.

### HasPlayers

`func (o *BroadcastRoundGet200ResponseGamesInner) HasPlayers() bool`

HasPlayers returns a boolean if a field has been set.

### GetLastMove

`func (o *BroadcastRoundGet200ResponseGamesInner) GetLastMove() string`

GetLastMove returns the LastMove field if non-nil, zero value otherwise.

### GetLastMoveOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetLastMoveOk() (*string, bool)`

GetLastMoveOk returns a tuple with the LastMove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMove

`func (o *BroadcastRoundGet200ResponseGamesInner) SetLastMove(v string)`

SetLastMove sets LastMove field to given value.

### HasLastMove

`func (o *BroadcastRoundGet200ResponseGamesInner) HasLastMove() bool`

HasLastMove returns a boolean if a field has been set.

### GetCheck

`func (o *BroadcastRoundGet200ResponseGamesInner) GetCheck() string`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetCheckOk() (*string, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *BroadcastRoundGet200ResponseGamesInner) SetCheck(v string)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *BroadcastRoundGet200ResponseGamesInner) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetThinkTime

`func (o *BroadcastRoundGet200ResponseGamesInner) GetThinkTime() int32`

GetThinkTime returns the ThinkTime field if non-nil, zero value otherwise.

### GetThinkTimeOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetThinkTimeOk() (*int32, bool)`

GetThinkTimeOk returns a tuple with the ThinkTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkTime

`func (o *BroadcastRoundGet200ResponseGamesInner) SetThinkTime(v int32)`

SetThinkTime sets ThinkTime field to given value.

### HasThinkTime

`func (o *BroadcastRoundGet200ResponseGamesInner) HasThinkTime() bool`

HasThinkTime returns a boolean if a field has been set.

### GetStatus

`func (o *BroadcastRoundGet200ResponseGamesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BroadcastRoundGet200ResponseGamesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BroadcastRoundGet200ResponseGamesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BroadcastRoundGet200ResponseGamesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


