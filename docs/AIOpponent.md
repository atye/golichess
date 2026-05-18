# AIOpponent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**nil**](nil.md) |  | 
**Username** | **string** |  | 
**Ai** | **int32** | AI level, from 1 to 8, where 1 is the weakest and 8 is the strongest. | 

## Methods

### NewAIOpponent

`func NewAIOpponent(id nil, username string, ai int32, ) *AIOpponent`

NewAIOpponent instantiates a new AIOpponent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAIOpponentWithDefaults

`func NewAIOpponentWithDefaults() *AIOpponent`

NewAIOpponentWithDefaults instantiates a new AIOpponent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AIOpponent) GetId() nil`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AIOpponent) GetIdOk() (*nil, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AIOpponent) SetId(v nil)`

SetId sets Id field to given value.


### GetUsername

`func (o *AIOpponent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AIOpponent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AIOpponent) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetAi

`func (o *AIOpponent) GetAi() int32`

GetAi returns the Ai field if non-nil, zero value otherwise.

### GetAiOk

`func (o *AIOpponent) GetAiOk() (*int32, bool)`

GetAiOk returns a tuple with the Ai field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAi

`func (o *AIOpponent) SetAi(v int32)`

SetAi sets Ai field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


