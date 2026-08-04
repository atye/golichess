# TeamUpdateMsg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Date** | **int32** |  | 
**Sender** | [**LightUser**](LightUser.md) |  | 
**Team** | Pointer to [**LightTeam**](LightTeam.md) |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewTeamUpdateMsg

`func NewTeamUpdateMsg(id string, date int32, sender LightUser, text string, ) *TeamUpdateMsg`

NewTeamUpdateMsg instantiates a new TeamUpdateMsg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdateMsgWithDefaults

`func NewTeamUpdateMsgWithDefaults() *TeamUpdateMsg`

NewTeamUpdateMsgWithDefaults instantiates a new TeamUpdateMsg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamUpdateMsg) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamUpdateMsg) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamUpdateMsg) SetId(v string)`

SetId sets Id field to given value.


### GetDate

`func (o *TeamUpdateMsg) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *TeamUpdateMsg) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *TeamUpdateMsg) SetDate(v int32)`

SetDate sets Date field to given value.


### GetSender

`func (o *TeamUpdateMsg) GetSender() LightUser`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *TeamUpdateMsg) GetSenderOk() (*LightUser, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *TeamUpdateMsg) SetSender(v LightUser)`

SetSender sets Sender field to given value.


### GetTeam

`func (o *TeamUpdateMsg) GetTeam() LightTeam`

GetTeam returns the Team field if non-nil, zero value otherwise.

### GetTeamOk

`func (o *TeamUpdateMsg) GetTeamOk() (*LightTeam, bool)`

GetTeamOk returns a tuple with the Team field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeam

`func (o *TeamUpdateMsg) SetTeam(v LightTeam)`

SetTeam sets Team field to given value.

### HasTeam

`func (o *TeamUpdateMsg) HasTeam() bool`

HasTeam returns a boolean if a field has been set.

### GetText

`func (o *TeamUpdateMsg) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TeamUpdateMsg) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TeamUpdateMsg) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


