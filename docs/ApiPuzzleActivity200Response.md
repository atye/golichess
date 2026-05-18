# ApiPuzzleActivity200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **int32** |  | 
**Puzzle** | [**ApiPuzzleActivity200ResponsePuzzle**](ApiPuzzleActivity200ResponsePuzzle.md) |  | 
**Win** | **bool** |  | 

## Methods

### NewApiPuzzleActivity200Response

`func NewApiPuzzleActivity200Response(date int32, puzzle ApiPuzzleActivity200ResponsePuzzle, win bool, ) *ApiPuzzleActivity200Response`

NewApiPuzzleActivity200Response instantiates a new ApiPuzzleActivity200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleActivity200ResponseWithDefaults

`func NewApiPuzzleActivity200ResponseWithDefaults() *ApiPuzzleActivity200Response`

NewApiPuzzleActivity200ResponseWithDefaults instantiates a new ApiPuzzleActivity200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ApiPuzzleActivity200Response) GetDate() int32`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ApiPuzzleActivity200Response) GetDateOk() (*int32, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ApiPuzzleActivity200Response) SetDate(v int32)`

SetDate sets Date field to given value.


### GetPuzzle

`func (o *ApiPuzzleActivity200Response) GetPuzzle() ApiPuzzleActivity200ResponsePuzzle`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *ApiPuzzleActivity200Response) GetPuzzleOk() (*ApiPuzzleActivity200ResponsePuzzle, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *ApiPuzzleActivity200Response) SetPuzzle(v ApiPuzzleActivity200ResponsePuzzle)`

SetPuzzle sets Puzzle field to given value.


### GetWin

`func (o *ApiPuzzleActivity200Response) GetWin() bool`

GetWin returns the Win field if non-nil, zero value otherwise.

### GetWinOk

`func (o *ApiPuzzleActivity200Response) GetWinOk() (*bool, bool)`

GetWinOk returns a tuple with the Win field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWin

`func (o *ApiPuzzleActivity200Response) SetWin(v bool)`

SetWin sets Win field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


