# GamePlayerAi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiLevel** | **int32** |  | 
**Analysis** | Pointer to [**GamePlayerAiAnalysis**](GamePlayerAiAnalysis.md) |  | [optional] 

## Methods

### NewGamePlayerAi

`func NewGamePlayerAi(aiLevel int32, ) *GamePlayerAi`

NewGamePlayerAi instantiates a new GamePlayerAi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGamePlayerAiWithDefaults

`func NewGamePlayerAiWithDefaults() *GamePlayerAi`

NewGamePlayerAiWithDefaults instantiates a new GamePlayerAi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiLevel

`func (o *GamePlayerAi) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *GamePlayerAi) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *GamePlayerAi) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.


### GetAnalysis

`func (o *GamePlayerAi) GetAnalysis() GamePlayerAiAnalysis`

GetAnalysis returns the Analysis field if non-nil, zero value otherwise.

### GetAnalysisOk

`func (o *GamePlayerAi) GetAnalysisOk() (*GamePlayerAiAnalysis, bool)`

GetAnalysisOk returns a tuple with the Analysis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalysis

`func (o *GamePlayerAi) SetAnalysis(v GamePlayerAiAnalysis)`

SetAnalysis sets Analysis field to given value.

### HasAnalysis

`func (o *GamePlayerAi) HasAnalysis() bool`

HasAnalysis returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


