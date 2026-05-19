# ApiExternalEngineAnalyse200ResponsePvsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **int32** | Current search depth of the pv | 
**Cp** | Pointer to **int32** | Evaluation in centi-pawns, from White&#39;s point of view | [optional] 
**Mate** | Pointer to **int32** | Evaluation in signed moves to mate, from White&#39;s point of view | [optional] 
**Moves** | **[]string** | Variation in UCI notation | 

## Methods

### NewApiExternalEngineAnalyse200ResponsePvsInner

`func NewApiExternalEngineAnalyse200ResponsePvsInner(depth int32, moves []string, ) *ApiExternalEngineAnalyse200ResponsePvsInner`

NewApiExternalEngineAnalyse200ResponsePvsInner instantiates a new ApiExternalEngineAnalyse200ResponsePvsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults

`func NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults() *ApiExternalEngineAnalyse200ResponsePvsInner`

NewApiExternalEngineAnalyse200ResponsePvsInnerWithDefaults instantiates a new ApiExternalEngineAnalyse200ResponsePvsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetCp() int32`

GetCp returns the Cp field if non-nil, zero value otherwise.

### GetCpOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetCpOk() (*int32, bool)`

GetCpOk returns a tuple with the Cp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetCp(v int32)`

SetCp sets Cp field to given value.

### HasCp

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) HasCp() bool`

HasCp returns a boolean if a field has been set.

### GetMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMate() int32`

GetMate returns the Mate field if non-nil, zero value otherwise.

### GetMateOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMateOk() (*int32, bool)`

GetMateOk returns a tuple with the Mate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetMate(v int32)`

SetMate sets Mate field to given value.

### HasMate

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) HasMate() bool`

HasMate returns a boolean if a field has been set.

### GetMoves

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMoves() []string`

GetMoves returns the Moves field if non-nil, zero value otherwise.

### GetMovesOk

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) GetMovesOk() (*[]string, bool)`

GetMovesOk returns a tuple with the Moves field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoves

`func (o *ApiExternalEngineAnalyse200ResponsePvsInner) SetMoves(v []string)`

SetMoves sets Moves field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


