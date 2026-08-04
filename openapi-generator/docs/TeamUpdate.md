# TeamUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msg** | [**TeamUpdateMsg**](TeamUpdateMsg.md) |  | 
**Seen** | **bool** |  | 

## Methods

### NewTeamUpdate

`func NewTeamUpdate(msg TeamUpdateMsg, seen bool, ) *TeamUpdate`

NewTeamUpdate instantiates a new TeamUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdateWithDefaults

`func NewTeamUpdateWithDefaults() *TeamUpdate`

NewTeamUpdateWithDefaults instantiates a new TeamUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsg

`func (o *TeamUpdate) GetMsg() TeamUpdateMsg`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *TeamUpdate) GetMsgOk() (*TeamUpdateMsg, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *TeamUpdate) SetMsg(v TeamUpdateMsg)`

SetMsg sets Msg field to given value.


### GetSeen

`func (o *TeamUpdate) GetSeen() bool`

GetSeen returns the Seen field if non-nil, zero value otherwise.

### GetSeenOk

`func (o *TeamUpdate) GetSeenOk() (*bool, bool)`

GetSeenOk returns a tuple with the Seen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeen

`func (o *TeamUpdate) SetSeen(v bool)`

SetSeen sets Seen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


