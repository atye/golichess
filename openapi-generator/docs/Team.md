# Team

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Leader** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 
**Leaders** | Pointer to [**[]LightUser**](LightUser.md) |  | [optional] 
**NbMembers** | Pointer to **int32** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**Joined** | Pointer to **bool** |  | [optional] 
**Requested** | Pointer to **bool** |  | [optional] 

## Methods

### NewTeam

`func NewTeam(id string, name string, ) *Team`

NewTeam instantiates a new Team object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamWithDefaults

`func NewTeamWithDefaults() *Team`

NewTeamWithDefaults instantiates a new Team object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Team) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Team) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Team) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Team) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Team) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Team) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *Team) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Team) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Team) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Team) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlair

`func (o *Team) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *Team) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *Team) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *Team) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetLeader

`func (o *Team) GetLeader() LightUser`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *Team) GetLeaderOk() (*LightUser, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *Team) SetLeader(v LightUser)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *Team) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetLeaders

`func (o *Team) GetLeaders() []LightUser`

GetLeaders returns the Leaders field if non-nil, zero value otherwise.

### GetLeadersOk

`func (o *Team) GetLeadersOk() (*[]LightUser, bool)`

GetLeadersOk returns a tuple with the Leaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaders

`func (o *Team) SetLeaders(v []LightUser)`

SetLeaders sets Leaders field to given value.

### HasLeaders

`func (o *Team) HasLeaders() bool`

HasLeaders returns a boolean if a field has been set.

### GetNbMembers

`func (o *Team) GetNbMembers() int32`

GetNbMembers returns the NbMembers field if non-nil, zero value otherwise.

### GetNbMembersOk

`func (o *Team) GetNbMembersOk() (*int32, bool)`

GetNbMembersOk returns a tuple with the NbMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbMembers

`func (o *Team) SetNbMembers(v int32)`

SetNbMembers sets NbMembers field to given value.

### HasNbMembers

`func (o *Team) HasNbMembers() bool`

HasNbMembers returns a boolean if a field has been set.

### GetOpen

`func (o *Team) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *Team) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *Team) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *Team) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetJoined

`func (o *Team) GetJoined() bool`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *Team) GetJoinedOk() (*bool, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *Team) SetJoined(v bool)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *Team) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetRequested

`func (o *Team) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *Team) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *Team) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *Team) HasRequested() bool`

HasRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


