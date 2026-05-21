# Leaderboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Users** | [**[]TopUser**](TopUser.md) |  | 

## Methods

### NewLeaderboard

`func NewLeaderboard(users []TopUser, ) *Leaderboard`

NewLeaderboard instantiates a new Leaderboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardWithDefaults

`func NewLeaderboardWithDefaults() *Leaderboard`

NewLeaderboardWithDefaults instantiates a new Leaderboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsers

`func (o *Leaderboard) GetUsers() []TopUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Leaderboard) GetUsersOk() (*[]TopUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Leaderboard) SetUsers(v []TopUser)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


