# BroadcastsByUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentPage** | **int32** |  | 
**MaxPerPage** | **int32** |  | 
**CurrentPageResults** | [**[]BroadcastsByUser200ResponseCurrentPageResultsInner**](BroadcastsByUser200ResponseCurrentPageResultsInner.md) |  | 
**NbResults** | **int32** |  | 
**PreviousPage** | **NullableInt32** |  | 
**NextPage** | **NullableInt32** |  | 
**NbPages** | **int32** |  | 

## Methods

### NewBroadcastsByUser200Response

`func NewBroadcastsByUser200Response(currentPage int32, maxPerPage int32, currentPageResults []BroadcastsByUser200ResponseCurrentPageResultsInner, nbResults int32, previousPage NullableInt32, nextPage NullableInt32, nbPages int32, ) *BroadcastsByUser200Response`

NewBroadcastsByUser200Response instantiates a new BroadcastsByUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastsByUser200ResponseWithDefaults

`func NewBroadcastsByUser200ResponseWithDefaults() *BroadcastsByUser200Response`

NewBroadcastsByUser200ResponseWithDefaults instantiates a new BroadcastsByUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentPage

`func (o *BroadcastsByUser200Response) GetCurrentPage() int32`

GetCurrentPage returns the CurrentPage field if non-nil, zero value otherwise.

### GetCurrentPageOk

`func (o *BroadcastsByUser200Response) GetCurrentPageOk() (*int32, bool)`

GetCurrentPageOk returns a tuple with the CurrentPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPage

`func (o *BroadcastsByUser200Response) SetCurrentPage(v int32)`

SetCurrentPage sets CurrentPage field to given value.


### GetMaxPerPage

`func (o *BroadcastsByUser200Response) GetMaxPerPage() int32`

GetMaxPerPage returns the MaxPerPage field if non-nil, zero value otherwise.

### GetMaxPerPageOk

`func (o *BroadcastsByUser200Response) GetMaxPerPageOk() (*int32, bool)`

GetMaxPerPageOk returns a tuple with the MaxPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerPage

`func (o *BroadcastsByUser200Response) SetMaxPerPage(v int32)`

SetMaxPerPage sets MaxPerPage field to given value.


### GetCurrentPageResults

`func (o *BroadcastsByUser200Response) GetCurrentPageResults() []BroadcastsByUser200ResponseCurrentPageResultsInner`

GetCurrentPageResults returns the CurrentPageResults field if non-nil, zero value otherwise.

### GetCurrentPageResultsOk

`func (o *BroadcastsByUser200Response) GetCurrentPageResultsOk() (*[]BroadcastsByUser200ResponseCurrentPageResultsInner, bool)`

GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPageResults

`func (o *BroadcastsByUser200Response) SetCurrentPageResults(v []BroadcastsByUser200ResponseCurrentPageResultsInner)`

SetCurrentPageResults sets CurrentPageResults field to given value.


### GetNbResults

`func (o *BroadcastsByUser200Response) GetNbResults() int32`

GetNbResults returns the NbResults field if non-nil, zero value otherwise.

### GetNbResultsOk

`func (o *BroadcastsByUser200Response) GetNbResultsOk() (*int32, bool)`

GetNbResultsOk returns a tuple with the NbResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbResults

`func (o *BroadcastsByUser200Response) SetNbResults(v int32)`

SetNbResults sets NbResults field to given value.


### GetPreviousPage

`func (o *BroadcastsByUser200Response) GetPreviousPage() int32`

GetPreviousPage returns the PreviousPage field if non-nil, zero value otherwise.

### GetPreviousPageOk

`func (o *BroadcastsByUser200Response) GetPreviousPageOk() (*int32, bool)`

GetPreviousPageOk returns a tuple with the PreviousPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousPage

`func (o *BroadcastsByUser200Response) SetPreviousPage(v int32)`

SetPreviousPage sets PreviousPage field to given value.


### SetPreviousPageNil

`func (o *BroadcastsByUser200Response) SetPreviousPageNil(b bool)`

 SetPreviousPageNil sets the value for PreviousPage to be an explicit nil

### UnsetPreviousPage
`func (o *BroadcastsByUser200Response) UnsetPreviousPage()`

UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
### GetNextPage

`func (o *BroadcastsByUser200Response) GetNextPage() int32`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *BroadcastsByUser200Response) GetNextPageOk() (*int32, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *BroadcastsByUser200Response) SetNextPage(v int32)`

SetNextPage sets NextPage field to given value.


### SetNextPageNil

`func (o *BroadcastsByUser200Response) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *BroadcastsByUser200Response) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
### GetNbPages

`func (o *BroadcastsByUser200Response) GetNbPages() int32`

GetNbPages returns the NbPages field if non-nil, zero value otherwise.

### GetNbPagesOk

`func (o *BroadcastsByUser200Response) GetNbPagesOk() (*int32, bool)`

GetNbPagesOk returns a tuple with the NbPages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNbPages

`func (o *BroadcastsByUser200Response) SetNbPages(v int32)`

SetNbPages sets NbPages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


