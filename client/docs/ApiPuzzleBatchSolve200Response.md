# ApiPuzzleBatchSolve200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Puzzles** | Pointer to [**[]ApiPuzzleId200Response**](ApiPuzzleId200Response.md) |  | [optional] 
**Glicko** | Pointer to [**ApiUserPerf200ResponsePerfGlicko**](ApiUserPerf200ResponsePerfGlicko.md) |  | [optional] 
**Rounds** | Pointer to [**[]ApiPuzzleBatchSolve200ResponseRoundsInner**](ApiPuzzleBatchSolve200ResponseRoundsInner.md) |  | [optional] 

## Methods

### NewApiPuzzleBatchSolve200Response

`func NewApiPuzzleBatchSolve200Response() *ApiPuzzleBatchSolve200Response`

NewApiPuzzleBatchSolve200Response instantiates a new ApiPuzzleBatchSolve200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleBatchSolve200ResponseWithDefaults

`func NewApiPuzzleBatchSolve200ResponseWithDefaults() *ApiPuzzleBatchSolve200Response`

NewApiPuzzleBatchSolve200ResponseWithDefaults instantiates a new ApiPuzzleBatchSolve200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPuzzles

`func (o *ApiPuzzleBatchSolve200Response) GetPuzzles() []ApiPuzzleId200Response`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *ApiPuzzleBatchSolve200Response) GetPuzzlesOk() (*[]ApiPuzzleId200Response, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *ApiPuzzleBatchSolve200Response) SetPuzzles(v []ApiPuzzleId200Response)`

SetPuzzles sets Puzzles field to given value.

### HasPuzzles

`func (o *ApiPuzzleBatchSolve200Response) HasPuzzles() bool`

HasPuzzles returns a boolean if a field has been set.

### GetGlicko

`func (o *ApiPuzzleBatchSolve200Response) GetGlicko() ApiUserPerf200ResponsePerfGlicko`

GetGlicko returns the Glicko field if non-nil, zero value otherwise.

### GetGlickoOk

`func (o *ApiPuzzleBatchSolve200Response) GetGlickoOk() (*ApiUserPerf200ResponsePerfGlicko, bool)`

GetGlickoOk returns a tuple with the Glicko field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlicko

`func (o *ApiPuzzleBatchSolve200Response) SetGlicko(v ApiUserPerf200ResponsePerfGlicko)`

SetGlicko sets Glicko field to given value.

### HasGlicko

`func (o *ApiPuzzleBatchSolve200Response) HasGlicko() bool`

HasGlicko returns a boolean if a field has been set.

### GetRounds

`func (o *ApiPuzzleBatchSolve200Response) GetRounds() []ApiPuzzleBatchSolve200ResponseRoundsInner`

GetRounds returns the Rounds field if non-nil, zero value otherwise.

### GetRoundsOk

`func (o *ApiPuzzleBatchSolve200Response) GetRoundsOk() (*[]ApiPuzzleBatchSolve200ResponseRoundsInner, bool)`

GetRoundsOk returns a tuple with the Rounds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRounds

`func (o *ApiPuzzleBatchSolve200Response) SetRounds(v []ApiPuzzleBatchSolve200ResponseRoundsInner)`

SetRounds sets Rounds field to given value.

### HasRounds

`func (o *ApiPuzzleBatchSolve200Response) HasRounds() bool`

HasRounds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


