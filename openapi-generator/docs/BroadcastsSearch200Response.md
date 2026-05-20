# BroadcastsSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**MaxPerPage** | **int32** |  | 
**CurrentPageResults** | [**[]BroadcastsTop200ResponseActiveInner**](BroadcastsTop200ResponseActiveInner.md) |  | 
**PreviousPage** | **NullableInt32** |  | 
**NextPage** | **NullableInt32** |  | 

## Methods

### NewBroadcastsSearch200Response

`func NewBroadcastsSearch200Response(currentPage int32, maxPerPage int32, currentPageResults []BroadcastsTop200ResponseActiveInner, previousPage NullableInt32, nextPage NullableInt32, ) *BroadcastsSearch200Response`

NewBroadcastsSearch200Response instantiates a new BroadcastsSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastsSearch200ResponseWithDefaults

`func NewBroadcastsSearch200ResponseWithDefaults() *BroadcastsSearch200Response`

NewBroadcastsSearch200ResponseWithDefaults instantiates a new BroadcastsSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *BroadcastsSearch200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *BroadcastsSearch200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *BroadcastsSearch200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetMaxPerPage

`func (o *BroadcastsSearch200Response) GetMaxPerPage() int32`

GetMaxPerPage returns the MaxPerPage field if non-nil, zero value otherwise.

### GetMaxPerPageOk

`func (o *BroadcastsSearch200Response) GetMaxPerPageOk() (*int32, bool)`

GetMaxPerPageOk returns a tuple with the MaxPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerPage

`func (o *BroadcastsSearch200Response) SetMaxPerPage(v int32)`

SetMaxPerPage sets MaxPerPage field to given value.


### GetCurrentPageResults

`func (o *BroadcastsSearch200Response) GetCurrentPageResults() []BroadcastsTop200ResponseActiveInner`

GetCurrentPageResults returns the CurrentPageResults field if non-nil, zero value otherwise.

### GetCurrentPageResultsOk

`func (o *BroadcastsSearch200Response) GetCurrentPageResultsOk() (*[]BroadcastsTop200ResponseActiveInner, bool)`

GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageResults

`func (o *BroadcastsSearch200Response) SetCurrentPageResults(v []BroadcastsTop200ResponseActiveInner)`

SetCurrentPageResults sets CurrentPageResults field to given value.


### GetPreviousPage

`func (o *BroadcastsSearch200Response) GetPreviousPage() int32`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *BroadcastsSearch200Response) GetPreviousPageOk() (*int32, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *BroadcastsSearch200Response) SetPreviousPage(v int32)`

SetPreviousPage sets PreviousPage field to given value.


### SetPreviousPageNil

`func (o *BroadcastsSearch200Response) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *BroadcastsSearch200Response) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *BroadcastsSearch200Response) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *BroadcastsSearch200Response) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *BroadcastsSearch200Response) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.


### SetNextPageNil

`func (o *BroadcastsSearch200Response) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *BroadcastsSearch200Response) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


