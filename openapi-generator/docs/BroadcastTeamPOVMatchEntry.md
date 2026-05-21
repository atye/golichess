# BroadcastTeamPOVMatchEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RoundId** | **string** |  | 
**Opponent** | **string** | The name of the opposing team | 
**Mp** | Pointer to **float32** | Match points scored in this match | [optional] 
**Gp** | Pointer to **float32** | Game points scored in this match | [optional] 
**Points** | Pointer to [**BroadcastPointStr**](BroadcastPointStr.md) |  | [optional] 

## Methods

### NewBroadcastTeamPOVMatchEntry

`func NewBroadcastTeamPOVMatchEntry(roundId string, opponent string, ) *BroadcastTeamPOVMatchEntry`

NewBroadcastTeamPOVMatchEntry instantiates a new BroadcastTeamPOVMatchEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTeamPOVMatchEntryWithDefaults

`func NewBroadcastTeamPOVMatchEntryWithDefaults() *BroadcastTeamPOVMatchEntry`

NewBroadcastTeamPOVMatchEntryWithDefaults instantiates a new BroadcastTeamPOVMatchEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRoundId

`func (o *BroadcastTeamPOVMatchEntry) GetRoundId() string`

GetRoundId returns the RoundId field if non-nil, zero value otherwise.

### GetRoundIdOk

`func (o *BroadcastTeamPOVMatchEntry) GetRoundIdOk() (*string, bool)`

GetRoundIdOk returns a tuple with the RoundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundId

`func (o *BroadcastTeamPOVMatchEntry) SetRoundId(v string)`

SetRoundId sets RoundId field to given value.


### GetOpponent

`func (o *BroadcastTeamPOVMatchEntry) GetOpponent() string`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *BroadcastTeamPOVMatchEntry) GetOpponentOk() (*string, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *BroadcastTeamPOVMatchEntry) SetOpponent(v string)`

SetOpponent sets Opponent field to given value.


### GetMp

`func (o *BroadcastTeamPOVMatchEntry) GetMp() float32`

GetMp returns the Mp field if non-nil, zero value otherwise.

### GetMpOk

`func (o *BroadcastTeamPOVMatchEntry) GetMpOk() (*float32, bool)`

GetMpOk returns a tuple with the Mp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMp

`func (o *BroadcastTeamPOVMatchEntry) SetMp(v float32)`

SetMp sets Mp field to given value.

### HasMp

`func (o *BroadcastTeamPOVMatchEntry) HasMp() bool`

HasMp returns a boolean if a field has been set.

### GetGp

`func (o *BroadcastTeamPOVMatchEntry) GetGp() float32`

GetGp returns the Gp field if non-nil, zero value otherwise.

### GetGpOk

`func (o *BroadcastTeamPOVMatchEntry) GetGpOk() (*float32, bool)`

GetGpOk returns a tuple with the Gp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGp

`func (o *BroadcastTeamPOVMatchEntry) SetGp(v float32)`

SetGp sets Gp field to given value.

### HasGp

`func (o *BroadcastTeamPOVMatchEntry) HasGp() bool`

HasGp returns a boolean if a field has been set.

### GetPoints

`func (o *BroadcastTeamPOVMatchEntry) GetPoints() BroadcastPointStr`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *BroadcastTeamPOVMatchEntry) GetPointsOk() (*BroadcastPointStr, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *BroadcastTeamPOVMatchEntry) SetPoints(v BroadcastPointStr)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *BroadcastTeamPOVMatchEntry) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


