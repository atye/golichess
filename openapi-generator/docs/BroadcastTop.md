# BroadcastTop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**[]BroadcastWithLastRound**](BroadcastWithLastRound.md) |  | [optional] 
**Upcoming** | Pointer to [**[]BroadcastWithLastRound**](BroadcastWithLastRound.md) |  | [optional] 
**Past** | Pointer to [**BroadcastTopPast**](BroadcastTopPast.md) |  | [optional] 

## Methods

### NewBroadcastTop

`func NewBroadcastTop() *BroadcastTop`

NewBroadcastTop instantiates a new BroadcastTop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTopWithDefaults

`func NewBroadcastTopWithDefaults() *BroadcastTop`

NewBroadcastTopWithDefaults instantiates a new BroadcastTop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BroadcastTop) GetActive() []BroadcastWithLastRound`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BroadcastTop) GetActiveOk() (*[]BroadcastWithLastRound, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BroadcastTop) SetActive(v []BroadcastWithLastRound)`

SetActive sets Active field to given value.

### HasActive

`func (o *BroadcastTop) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUpcoming

`func (o *BroadcastTop) GetUpcoming() []BroadcastWithLastRound`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *BroadcastTop) GetUpcomingOk() (*[]BroadcastWithLastRound, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *BroadcastTop) SetUpcoming(v []BroadcastWithLastRound)`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *BroadcastTop) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.

### GetPast

`func (o *BroadcastTop) GetPast() BroadcastTopPast`

GetPast returns the Past field if non-nil, zero value otherwise.

### GetPastOk

`func (o *BroadcastTop) GetPastOk() (*BroadcastTopPast, bool)`

GetPastOk returns a tuple with the Past field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPast

`func (o *BroadcastTop) SetPast(v BroadcastTopPast)`

SetPast sets Past field to given value.

### HasPast

`func (o *BroadcastTop) HasPast() bool`

HasPast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


