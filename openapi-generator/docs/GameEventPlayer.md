# GameEventPlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiLevel** | Pointer to **int32** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to [**NullableTitle**](Title.md) |  | [optional] 
**Rating** | Pointer to **int32** |  | [optional] 
**Provisional** | Pointer to **bool** |  | [optional] 

## Methods

### NewGameEventPlayer

`func NewGameEventPlayer(id string, name string, ) *GameEventPlayer`

NewGameEventPlayer instantiates a new GameEventPlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameEventPlayerWithDefaults

`func NewGameEventPlayerWithDefaults() *GameEventPlayer`

NewGameEventPlayerWithDefaults instantiates a new GameEventPlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiLevel

`func (o *GameEventPlayer) GetAiLevel() int32`

GetAiLevel returns the AiLevel field if non-nil, zero value otherwise.

### GetAiLevelOk

`func (o *GameEventPlayer) GetAiLevelOk() (*int32, bool)`

GetAiLevelOk returns a tuple with the AiLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiLevel

`func (o *GameEventPlayer) SetAiLevel(v int32)`

SetAiLevel sets AiLevel field to given value.

### HasAiLevel

`func (o *GameEventPlayer) HasAiLevel() bool`

HasAiLevel returns a boolean if a field has been set.

### GetId

`func (o *GameEventPlayer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameEventPlayer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameEventPlayer) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GameEventPlayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GameEventPlayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GameEventPlayer) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *GameEventPlayer) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GameEventPlayer) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GameEventPlayer) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GameEventPlayer) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GameEventPlayer) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GameEventPlayer) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetRating

`func (o *GameEventPlayer) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *GameEventPlayer) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *GameEventPlayer) SetRating(v int32)`

SetRating sets Rating field to given value.

### HasRating

`func (o *GameEventPlayer) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetProvisional

`func (o *GameEventPlayer) GetProvisional() bool`

GetProvisional returns the Provisional field if non-nil, zero value otherwise.

### GetProvisionalOk

`func (o *GameEventPlayer) GetProvisionalOk() (*bool, bool)`

GetProvisionalOk returns a tuple with the Provisional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisional

`func (o *GameEventPlayer) SetProvisional(v bool)`

SetProvisional sets Provisional field to given value.

### HasProvisional

`func (o *GameEventPlayer) HasProvisional() bool`

HasProvisional returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


