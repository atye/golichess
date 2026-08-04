# TeamUpdatesByTeamInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Team** | [**LightTeam**](LightTeam.md) |  | 
**Last** | **float32** |  | 
**Unread** | **int32** |  | 

## Methods

### NewTeamUpdatesByTeamInner

`func NewTeamUpdatesByTeamInner(team LightTeam, last float32, unread int32, ) *TeamUpdatesByTeamInner`

NewTeamUpdatesByTeamInner instantiates a new TeamUpdatesByTeamInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdatesByTeamInnerWithDefaults

`func NewTeamUpdatesByTeamInnerWithDefaults() *TeamUpdatesByTeamInner`

NewTeamUpdatesByTeamInnerWithDefaults instantiates a new TeamUpdatesByTeamInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTeam

`func (o *TeamUpdatesByTeamInner) GetTeam() LightTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamUpdatesByTeamInner) GetTeamOk() (*LightTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamUpdatesByTeamInner) SetTeam(v LightTeam)`

SetTeam sets Team field to given value.


### GetLast

`func (o *TeamUpdatesByTeamInner) GetLast() float32`

GetLast returns the Last field if non-nil, zero value otherwise.

### GetLastOk

`func (o *TeamUpdatesByTeamInner) GetLastOk() (*float32, bool)`

GetLastOk returns a tuple with the Last field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast

`func (o *TeamUpdatesByTeamInner) SetLast(v float32)`

SetLast sets Last field to given value.


### GetUnread

`func (o *TeamUpdatesByTeamInner) GetUnread() int32`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *TeamUpdatesByTeamInner) GetUnreadOk() (*int32, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *TeamUpdatesByTeamInner) SetUnread(v int32)`

SetUnread sets Unread field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


