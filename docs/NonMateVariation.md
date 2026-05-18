# NonMateVariation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cp** | **int32** | Evaluation in centi-pawns, from White&#39;s point of view | 
**Moves** | **string** | Variation in UCI notation (King to rook for Chess960-compatible castling notation)  | 

## Methods

### NewNonMateVariation

`func NewNonMateVariation(cp int32, moves string, ) *NonMateVariation`

NewNonMateVariation instantiates a new NonMateVariation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNonMateVariationWithDefaults

`func NewNonMateVariationWithDefaults() *NonMateVariation`

NewNonMateVariationWithDefaults instantiates a new NonMateVariation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCp

`func (o *NonMateVariation) GetCp() int32`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *NonMateVariation) GetCpOk() (*int32, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *NonMateVariation) SetCp(v int32)`

SetCp sets Cp field to given value.


### GetMoves

`func (o *NonMateVariation) GetMoves() string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *NonMateVariation) GetMovesOk() (*string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *NonMateVariation) SetMoves(v string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


