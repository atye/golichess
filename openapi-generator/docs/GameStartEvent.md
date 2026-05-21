# GameStartEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Game** | [**GameEventInfo**](GameEventInfo.md) |  | 

## Methods

### NewGameStartEvent

`func NewGameStartEvent(type_ string, game GameEventInfo, ) *GameStartEvent`

NewGameStartEvent instantiates a new GameStartEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameStartEventWithDefaults

`func NewGameStartEventWithDefaults() *GameStartEvent`

NewGameStartEventWithDefaults instantiates a new GameStartEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameStartEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameStartEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameStartEvent) SetType(v string)`

SetType sets Type field to given value.


### GetGame

`func (o *GameStartEvent) GetGame() GameEventInfo`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *GameStartEvent) GetGameOk() (*GameEventInfo, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *GameStartEvent) SetGame(v GameEventInfo)`

SetGame sets Game field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


