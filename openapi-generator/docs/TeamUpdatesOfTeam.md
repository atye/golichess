# TeamUpdatesOfTeam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | [**LightTeam**](LightTeam.md) |  | 
**Subscribed** | **bool** |  | 
**Updates** | [**TeamUpdatesPager**](TeamUpdatesPager.md) |  | 
**ByTeam** | [**[]TeamUpdatesByTeamInner**](TeamUpdatesByTeamInner.md) |  | 

## Methods

### NewTeamUpdatesOfTeam

`func NewTeamUpdatesOfTeam(team LightTeam, subscribed bool, updates TeamUpdatesPager, byTeam []TeamUpdatesByTeamInner, ) *TeamUpdatesOfTeam`

NewTeamUpdatesOfTeam instantiates a new TeamUpdatesOfTeam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdatesOfTeamWithDefaults

`func NewTeamUpdatesOfTeamWithDefaults() *TeamUpdatesOfTeam`

NewTeamUpdatesOfTeamWithDefaults instantiates a new TeamUpdatesOfTeam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *TeamUpdatesOfTeam) GetTeam() LightTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamUpdatesOfTeam) GetTeamOk() (*LightTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamUpdatesOfTeam) SetTeam(v LightTeam)`

SetTeam sets Team field to given value.


### GetSubscribed

`func (o *TeamUpdatesOfTeam) GetSubscribed() bool`

GetSubscribed returns the Subscribed field if non-nil, zero value otherwise.

### GetSubscribedOk

`func (o *TeamUpdatesOfTeam) GetSubscribedOk() (*bool, bool)`

GetSubscribedOk returns a tuple with the Subscribed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscribed

`func (o *TeamUpdatesOfTeam) SetSubscribed(v bool)`

SetSubscribed sets Subscribed field to given value.


### GetUpdates

`func (o *TeamUpdatesOfTeam) GetUpdates() TeamUpdatesPager`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *TeamUpdatesOfTeam) GetUpdatesOk() (*TeamUpdatesPager, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *TeamUpdatesOfTeam) SetUpdates(v TeamUpdatesPager)`

SetUpdates sets Updates field to given value.


### GetByTeam

`func (o *TeamUpdatesOfTeam) GetByTeam() []TeamUpdatesByTeamInner`

GetByTeam returns the ByTeam field if non-nil, zero value otherwise.

### GetByTeamOk

`func (o *TeamUpdatesOfTeam) GetByTeamOk() (*[]TeamUpdatesByTeamInner, bool)`

GetByTeamOk returns a tuple with the ByTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTeam

`func (o *TeamUpdatesOfTeam) SetByTeam(v []TeamUpdatesByTeamInner)`

SetByTeam sets ByTeam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


