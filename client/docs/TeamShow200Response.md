# TeamShow200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Leader** | Pointer to [**ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | [optional] 
**Leaders** | Pointer to [**[]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | [optional] 
**NbMembers** | Pointer to **int32** |  | [optional] 
**Open** | Pointer to **bool** |  | [optional] 
**Joined** | Pointer to **bool** |  | [optional] 
**Requested** | Pointer to **bool** |  | [optional] 

## Methods

### NewTeamShow200Response

`func NewTeamShow200Response(id string, name string, ) *TeamShow200Response`

NewTeamShow200Response instantiates a new TeamShow200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamShow200ResponseWithDefaults

`func NewTeamShow200ResponseWithDefaults() *TeamShow200Response`

NewTeamShow200ResponseWithDefaults instantiates a new TeamShow200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TeamShow200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TeamShow200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TeamShow200Response) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TeamShow200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TeamShow200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TeamShow200Response) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *TeamShow200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TeamShow200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TeamShow200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TeamShow200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlair

`func (o *TeamShow200Response) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *TeamShow200Response) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *TeamShow200Response) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *TeamShow200Response) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetLeader

`func (o *TeamShow200Response) GetLeader() ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *TeamShow200Response) GetLeaderOk() (*ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *TeamShow200Response) SetLeader(v ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *TeamShow200Response) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetLeaders

`func (o *TeamShow200Response) GetLeaders() []ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetLeaders returns the Leaders field if non-nil, zero value otherwise.

### GetLeadersOk

`func (o *TeamShow200Response) GetLeadersOk() (*[]ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetLeadersOk returns a tuple with the Leaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaders

`func (o *TeamShow200Response) SetLeaders(v []ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetLeaders sets Leaders field to given value.

### HasLeaders

`func (o *TeamShow200Response) HasLeaders() bool`

HasLeaders returns a boolean if a field has been set.

### GetNbMembers

`func (o *TeamShow200Response) GetNbMembers() int32`

GetNbMembers returns the NbMembers field if non-nil, zero value otherwise.

### GetNbMembersOk

`func (o *TeamShow200Response) GetNbMembersOk() (*int32, bool)`

GetNbMembersOk returns a tuple with the NbMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbMembers

`func (o *TeamShow200Response) SetNbMembers(v int32)`

SetNbMembers sets NbMembers field to given value.

### HasNbMembers

`func (o *TeamShow200Response) HasNbMembers() bool`

HasNbMembers returns a boolean if a field has been set.

### GetOpen

`func (o *TeamShow200Response) GetOpen() bool`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *TeamShow200Response) GetOpenOk() (*bool, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *TeamShow200Response) SetOpen(v bool)`

SetOpen sets Open field to given value.

### HasOpen

`func (o *TeamShow200Response) HasOpen() bool`

HasOpen returns a boolean if a field has been set.

### GetJoined

`func (o *TeamShow200Response) GetJoined() bool`

GetJoined returns the Joined field if non-nil, zero value otherwise.

### GetJoinedOk

`func (o *TeamShow200Response) GetJoinedOk() (*bool, bool)`

GetJoinedOk returns a tuple with the Joined field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoined

`func (o *TeamShow200Response) SetJoined(v bool)`

SetJoined sets Joined field to given value.

### HasJoined

`func (o *TeamShow200Response) HasJoined() bool`

HasJoined returns a boolean if a field has been set.

### GetRequested

`func (o *TeamShow200Response) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *TeamShow200Response) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *TeamShow200Response) SetRequested(v bool)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *TeamShow200Response) HasRequested() bool`

HasRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


