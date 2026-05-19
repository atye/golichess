# MateVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mate** | **int32** | Evaluation in moves to mate, from White&#39;s point of view | 
**Moves** | **string** | Variation in UCI notation (King to rook for Chess960-compatible castling notation)  | 

## Methods

### NewMateVariation

`func NewMateVariation(mate int32, moves string, ) *MateVariation`

NewMateVariation instantiates a new MateVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMateVariationWithDefaults

`func NewMateVariationWithDefaults() *MateVariation`

NewMateVariationWithDefaults instantiates a new MateVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMate

`func (o *MateVariation) GetMate() int32`

GetMate returns the Mate field if non-nil, zero value otherwise.

### GetMateOk

`func (o *MateVariation) GetMateOk() (*int32, bool)`

GetMateOk returns a tuple with the Mate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMate

`func (o *MateVariation) SetMate(v int32)`

SetMate sets Mate field to given value.


### GetMoves

`func (o *MateVariation) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *MateVariation) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *MateVariation) SetMoves(v string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


