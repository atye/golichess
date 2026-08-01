# BroadcastFormGrouping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**BroadcastFormGroupingInfo**](BroadcastFormGroupingInfo.md) |  | [optional] 
**ScoreGroups** | Pointer to **[]string** | This parameter is repeated with an index for each score group, like &#39;grouping.scoreGroups[0]&#x3D;wYigbpXq,M5YHvpOX&#39; | [optional] 

## Methods

### NewBroadcastFormGrouping

`func NewBroadcastFormGrouping() *BroadcastFormGrouping`

NewBroadcastFormGrouping instantiates a new BroadcastFormGrouping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastFormGroupingWithDefaults

`func NewBroadcastFormGroupingWithDefaults() *BroadcastFormGrouping`

NewBroadcastFormGroupingWithDefaults instantiates a new BroadcastFormGrouping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *BroadcastFormGrouping) GetInfo() BroadcastFormGroupingInfo`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *BroadcastFormGrouping) GetInfoOk() (*BroadcastFormGroupingInfo, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *BroadcastFormGrouping) SetInfo(v BroadcastFormGroupingInfo)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *BroadcastFormGrouping) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetScoreGroups

`func (o *BroadcastFormGrouping) GetScoreGroups() []string`

GetScoreGroups returns the ScoreGroups field if non-nil, zero value otherwise.

### GetScoreGroupsOk

`func (o *BroadcastFormGrouping) GetScoreGroupsOk() (*[]string, bool)`

GetScoreGroupsOk returns a tuple with the ScoreGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScoreGroups

`func (o *BroadcastFormGrouping) SetScoreGroups(v []string)`

SetScoreGroups sets ScoreGroups field to given value.

### HasScoreGroups

`func (o *BroadcastFormGrouping) HasScoreGroups() bool`

HasScoreGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


