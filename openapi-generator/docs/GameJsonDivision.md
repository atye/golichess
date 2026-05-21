# GameJsonDivision

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Middle** | Pointer to **int32** | Ply at which the middlegame begins | [optional] 
**End** | Pointer to **int32** | Ply at which the endgame begins | [optional] 

## Methods

### NewGameJsonDivision

`func NewGameJsonDivision() *GameJsonDivision`

NewGameJsonDivision instantiates a new GameJsonDivision object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameJsonDivisionWithDefaults

`func NewGameJsonDivisionWithDefaults() *GameJsonDivision`

NewGameJsonDivisionWithDefaults instantiates a new GameJsonDivision object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMiddle

`func (o *GameJsonDivision) GetMiddle() int32`

GetMiddle returns the Middle field if non-nil, zero value otherwise.

### GetMiddleOk

`func (o *GameJsonDivision) GetMiddleOk() (*int32, bool)`

GetMiddleOk returns a tuple with the Middle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiddle

`func (o *GameJsonDivision) SetMiddle(v int32)`

SetMiddle sets Middle field to given value.

### HasMiddle

`func (o *GameJsonDivision) HasMiddle() bool`

HasMiddle returns a boolean if a field has been set.

### GetEnd

`func (o *GameJsonDivision) GetEnd() int32`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *GameJsonDivision) GetEndOk() (*int32, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *GameJsonDivision) SetEnd(v int32)`

SetEnd sets End field to given value.

### HasEnd

`func (o *GameJsonDivision) HasEnd() bool`

HasEnd returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


