# ApiCloudEval200ResponsePvsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cp** | **int32** | Evaluation in centi-pawns, from White&#39;s point of view | 
**Moves** | **string** | Variation in UCI notation (King to rook for Chess960-compatible castling notation)  | 
**Mate** | **int32** | Evaluation in moves to mate, from White&#39;s point of view | 

## Methods

### NewApiCloudEval200ResponsePvsInner

`func NewApiCloudEval200ResponsePvsInner(cp int32, moves string, mate int32, ) *ApiCloudEval200ResponsePvsInner`

NewApiCloudEval200ResponsePvsInner instantiates a new ApiCloudEval200ResponsePvsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCloudEval200ResponsePvsInnerWithDefaults

`func NewApiCloudEval200ResponsePvsInnerWithDefaults() *ApiCloudEval200ResponsePvsInner`

NewApiCloudEval200ResponsePvsInnerWithDefaults instantiates a new ApiCloudEval200ResponsePvsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCp

`func (o *ApiCloudEval200ResponsePvsInner) GetCp() int32`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *ApiCloudEval200ResponsePvsInner) GetCpOk() (*int32, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *ApiCloudEval200ResponsePvsInner) SetCp(v int32)`

SetCp sets Cp field to given value.


### GetMoves

`func (o *ApiCloudEval200ResponsePvsInner) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiCloudEval200ResponsePvsInner) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiCloudEval200ResponsePvsInner) SetMoves(v string)`

SetMoves sets Moves field to given value.


### GetMate

`func (o *ApiCloudEval200ResponsePvsInner) GetMate() int32`

GetMate returns the Mate field if non-nil, zero value otherwise.

### GetMateOk

`func (o *ApiCloudEval200ResponsePvsInner) GetMateOk() (*int32, bool)`

GetMateOk returns a tuple with the Mate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMate

`func (o *ApiCloudEval200ResponsePvsInner) SetMate(v int32)`

SetMate sets Mate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


