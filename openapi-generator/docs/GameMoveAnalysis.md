# GameMoveAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eval** | Pointer to **int32** | Evaluation in centipawns | [optional] 
**Mate** | Pointer to **int32** | Number of moves until forced mate | [optional] 
**Best** | Pointer to **string** | Best move in UCI notation (only if played move was inaccurate) | [optional] 
**Variation** | Pointer to **string** | Best variation in SAN notation (only if played move was inaccurate) | [optional] 
**Judgment** | Pointer to [**GameMoveAnalysisJudgment**](GameMoveAnalysisJudgment.md) |  | [optional] 

## Methods

### NewGameMoveAnalysis

`func NewGameMoveAnalysis() *GameMoveAnalysis`

NewGameMoveAnalysis instantiates a new GameMoveAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameMoveAnalysisWithDefaults

`func NewGameMoveAnalysisWithDefaults() *GameMoveAnalysis`

NewGameMoveAnalysisWithDefaults instantiates a new GameMoveAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEval

`func (o *GameMoveAnalysis) GetEval() int32`

GetEval returns the Eval field if non-nil, zero value otherwise.

### GetEvalOk

`func (o *GameMoveAnalysis) GetEvalOk() (*int32, bool)`

GetEvalOk returns a tuple with the Eval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEval

`func (o *GameMoveAnalysis) SetEval(v int32)`

SetEval sets Eval field to given value.

### HasEval

`func (o *GameMoveAnalysis) HasEval() bool`

HasEval returns a boolean if a field has been set.

### GetMate

`func (o *GameMoveAnalysis) GetMate() int32`

GetMate returns the Mate field if non-nil, zero value otherwise.

### GetMateOk

`func (o *GameMoveAnalysis) GetMateOk() (*int32, bool)`

GetMateOk returns a tuple with the Mate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMate

`func (o *GameMoveAnalysis) SetMate(v int32)`

SetMate sets Mate field to given value.

### HasMate

`func (o *GameMoveAnalysis) HasMate() bool`

HasMate returns a boolean if a field has been set.

### GetBest

`func (o *GameMoveAnalysis) GetBest() string`

GetBest returns the Best field if non-nil, zero value otherwise.

### GetBestOk

`func (o *GameMoveAnalysis) GetBestOk() (*string, bool)`

GetBestOk returns a tuple with the Best field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBest

`func (o *GameMoveAnalysis) SetBest(v string)`

SetBest sets Best field to given value.

### HasBest

`func (o *GameMoveAnalysis) HasBest() bool`

HasBest returns a boolean if a field has been set.

### GetVariation

`func (o *GameMoveAnalysis) GetVariation() string`

GetVariation returns the Variation field if non-nil, zero value otherwise.

### GetVariationOk

`func (o *GameMoveAnalysis) GetVariationOk() (*string, bool)`

GetVariationOk returns a tuple with the Variation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariation

`func (o *GameMoveAnalysis) SetVariation(v string)`

SetVariation sets Variation field to given value.

### HasVariation

`func (o *GameMoveAnalysis) HasVariation() bool`

HasVariation returns a boolean if a field has been set.

### GetJudgment

`func (o *GameMoveAnalysis) GetJudgment() GameMoveAnalysisJudgment`

GetJudgment returns the Judgment field if non-nil, zero value otherwise.

### GetJudgmentOk

`func (o *GameMoveAnalysis) GetJudgmentOk() (*GameMoveAnalysisJudgment, bool)`

GetJudgmentOk returns a tuple with the Judgment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJudgment

`func (o *GameMoveAnalysis) SetJudgment(v GameMoveAnalysisJudgment)`

SetJudgment sets Judgment field to given value.

### HasJudgment

`func (o *GameMoveAnalysis) HasJudgment() bool`

HasJudgment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


