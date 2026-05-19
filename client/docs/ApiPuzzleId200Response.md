# ApiPuzzleId200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Game** | [**ApiPuzzleId200ResponseGame**](ApiPuzzleId200ResponseGame.md) |  | 
**Puzzle** | [**ApiPuzzleDaily200ResponsePuzzle**](ApiPuzzleDaily200ResponsePuzzle.md) |  | 

## Methods

### NewApiPuzzleId200Response

`func NewApiPuzzleId200Response(game ApiPuzzleId200ResponseGame, puzzle ApiPuzzleDaily200ResponsePuzzle, ) *ApiPuzzleId200Response`

NewApiPuzzleId200Response instantiates a new ApiPuzzleId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleId200ResponseWithDefaults

`func NewApiPuzzleId200ResponseWithDefaults() *ApiPuzzleId200Response`

NewApiPuzzleId200ResponseWithDefaults instantiates a new ApiPuzzleId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGame

`func (o *ApiPuzzleId200Response) GetGame() ApiPuzzleId200ResponseGame`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *ApiPuzzleId200Response) GetGameOk() (*ApiPuzzleId200ResponseGame, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *ApiPuzzleId200Response) SetGame(v ApiPuzzleId200ResponseGame)`

SetGame sets Game field to given value.


### GetPuzzle

`func (o *ApiPuzzleId200Response) GetPuzzle() ApiPuzzleDaily200ResponsePuzzle`

GetPuzzle returns the Puzzle field if non-nil, zero value otherwise.

### GetPuzzleOk

`func (o *ApiPuzzleId200Response) GetPuzzleOk() (*ApiPuzzleDaily200ResponsePuzzle, bool)`

GetPuzzleOk returns a tuple with the Puzzle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzle

`func (o *ApiPuzzleId200Response) SetPuzzle(v ApiPuzzleDaily200ResponsePuzzle)`

SetPuzzle sets Puzzle field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


