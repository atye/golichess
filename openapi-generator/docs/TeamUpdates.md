# TeamUpdates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updates** | [**TeamUpdatesPager**](TeamUpdatesPager.md) |  | 
**ByTeam** | [**[]TeamUpdatesByTeamInner**](TeamUpdatesByTeamInner.md) |  | 

## Methods

### NewTeamUpdates

`func NewTeamUpdates(updates TeamUpdatesPager, byTeam []TeamUpdatesByTeamInner, ) *TeamUpdates`

NewTeamUpdates instantiates a new TeamUpdates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdatesWithDefaults

`func NewTeamUpdatesWithDefaults() *TeamUpdates`

NewTeamUpdatesWithDefaults instantiates a new TeamUpdates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdates

`func (o *TeamUpdates) GetUpdates() TeamUpdatesPager`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *TeamUpdates) GetUpdatesOk() (*TeamUpdatesPager, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *TeamUpdates) SetUpdates(v TeamUpdatesPager)`

SetUpdates sets Updates field to given value.


### GetByTeam

`func (o *TeamUpdates) GetByTeam() []TeamUpdatesByTeamInner`

GetByTeam returns the ByTeam field if non-nil, zero value otherwise.

### GetByTeamOk

`func (o *TeamUpdates) GetByTeamOk() (*[]TeamUpdatesByTeamInner, bool)`

GetByTeamOk returns a tuple with the ByTeam field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByTeam

`func (o *TeamUpdates) SetByTeam(v []TeamUpdatesByTeamInner)`

SetByTeam sets ByTeam field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


