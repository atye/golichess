# BroadcastTopPast

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | Pointer to **int32** |  | [optional] 
**MaxPerPage** | Pointer to **int32** |  | [optional] 
**CurrentPageResults** | Pointer to [**[]BroadcastWithLastRound**](BroadcastWithLastRound.md) |  | [optional] 
**PreviousPage** | Pointer to **NullableInt32** |  | [optional] 
**NextPage** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewBroadcastTopPast

`func NewBroadcastTopPast() *BroadcastTopPast`

NewBroadcastTopPast instantiates a new BroadcastTopPast object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastTopPastWithDefaults

`func NewBroadcastTopPastWithDefaults() *BroadcastTopPast`

NewBroadcastTopPastWithDefaults instantiates a new BroadcastTopPast object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *BroadcastTopPast) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *BroadcastTopPast) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *BroadcastTopPast) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.

### HasCurrentPage

`func (o *BroadcastTopPast) HasCurrentPage() bool`

HasCurrentPage returns a boolean if a field has been set.

### GetMaxPerPage

`func (o *BroadcastTopPast) GetMaxPerPage() int32`

GetMaxPerPage returns the MaxPerPage field if non-nil, zero value otherwise.

### GetMaxPerPageOk

`func (o *BroadcastTopPast) GetMaxPerPageOk() (*int32, bool)`

GetMaxPerPageOk returns a tuple with the MaxPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerPage

`func (o *BroadcastTopPast) SetMaxPerPage(v int32)`

SetMaxPerPage sets MaxPerPage field to given value.

### HasMaxPerPage

`func (o *BroadcastTopPast) HasMaxPerPage() bool`

HasMaxPerPage returns a boolean if a field has been set.

### GetCurrentPageResults

`func (o *BroadcastTopPast) GetCurrentPageResults() []BroadcastWithLastRound`

GetCurrentPageResults returns the CurrentPageResults field if non-nil, zero value otherwise.

### GetCurrentPageResultsOk

`func (o *BroadcastTopPast) GetCurrentPageResultsOk() (*[]BroadcastWithLastRound, bool)`

GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageResults

`func (o *BroadcastTopPast) SetCurrentPageResults(v []BroadcastWithLastRound)`

SetCurrentPageResults sets CurrentPageResults field to given value.

### HasCurrentPageResults

`func (o *BroadcastTopPast) HasCurrentPageResults() bool`

HasCurrentPageResults returns a boolean if a field has been set.

### GetPreviousPage

`func (o *BroadcastTopPast) GetPreviousPage() int32`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *BroadcastTopPast) GetPreviousPageOk() (*int32, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *BroadcastTopPast) SetPreviousPage(v int32)`

SetPreviousPage sets PreviousPage field to given value.

### HasPreviousPage

`func (o *BroadcastTopPast) HasPreviousPage() bool`

HasPreviousPage returns a boolean if a field has been set.

### SetPreviousPageNil

`func (o *BroadcastTopPast) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *BroadcastTopPast) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *BroadcastTopPast) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *BroadcastTopPast) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *BroadcastTopPast) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *BroadcastTopPast) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### SetNextPageNil

`func (o *BroadcastTopPast) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *BroadcastTopPast) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


