# TeamUpdatesPager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**MaxPerPage** | **int32** |  | 
**CurrentPageResults** | [**[]TeamUpdate**](TeamUpdate.md) |  | 
**PreviousPage** | **NullableInt32** |  | 
**NextPage** | **NullableInt32** |  | 
**NbResults** | **int32** |  | 
**NbPages** | **int32** |  | 

## Methods

### NewTeamUpdatesPager

`func NewTeamUpdatesPager(currentPage int32, maxPerPage int32, currentPageResults []TeamUpdate, previousPage NullableInt32, nextPage NullableInt32, nbResults int32, nbPages int32, ) *TeamUpdatesPager`

NewTeamUpdatesPager instantiates a new TeamUpdatesPager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamUpdatesPagerWithDefaults

`func NewTeamUpdatesPagerWithDefaults() *TeamUpdatesPager`

NewTeamUpdatesPagerWithDefaults instantiates a new TeamUpdatesPager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *TeamUpdatesPager) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *TeamUpdatesPager) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *TeamUpdatesPager) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetMaxPerPage

`func (o *TeamUpdatesPager) GetMaxPerPage() int32`

GetMaxPerPage returns the MaxPerPage field if non-nil, zero value otherwise.

### GetMaxPerPageOk

`func (o *TeamUpdatesPager) GetMaxPerPageOk() (*int32, bool)`

GetMaxPerPageOk returns a tuple with the MaxPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerPage

`func (o *TeamUpdatesPager) SetMaxPerPage(v int32)`

SetMaxPerPage sets MaxPerPage field to given value.


### GetCurrentPageResults

`func (o *TeamUpdatesPager) GetCurrentPageResults() []TeamUpdate`

GetCurrentPageResults returns the CurrentPageResults field if non-nil, zero value otherwise.

### GetCurrentPageResultsOk

`func (o *TeamUpdatesPager) GetCurrentPageResultsOk() (*[]TeamUpdate, bool)`

GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageResults

`func (o *TeamUpdatesPager) SetCurrentPageResults(v []TeamUpdate)`

SetCurrentPageResults sets CurrentPageResults field to given value.


### GetPreviousPage

`func (o *TeamUpdatesPager) GetPreviousPage() int32`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *TeamUpdatesPager) GetPreviousPageOk() (*int32, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *TeamUpdatesPager) SetPreviousPage(v int32)`

SetPreviousPage sets PreviousPage field to given value.


### SetPreviousPageNil

`func (o *TeamUpdatesPager) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *TeamUpdatesPager) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *TeamUpdatesPager) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *TeamUpdatesPager) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *TeamUpdatesPager) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.


### SetNextPageNil

`func (o *TeamUpdatesPager) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *TeamUpdatesPager) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
### GetNbResults

`func (o *TeamUpdatesPager) GetNbResults() int32`

GetNbResults returns the NbResults field if non-nil, zero value otherwise.

### GetNbResultsOk

`func (o *TeamUpdatesPager) GetNbResultsOk() (*int32, bool)`

GetNbResultsOk returns a tuple with the NbResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbResults

`func (o *TeamUpdatesPager) SetNbResults(v int32)`

SetNbResults sets NbResults field to given value.


### GetNbPages

`func (o *TeamUpdatesPager) GetNbPages() int32`

GetNbPages returns the NbPages field if non-nil, zero value otherwise.

### GetNbPagesOk

`func (o *TeamUpdatesPager) GetNbPagesOk() (*int32, bool)`

GetNbPagesOk returns a tuple with the NbPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPages

`func (o *TeamUpdatesPager) SetNbPages(v int32)`

SetNbPages sets NbPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


