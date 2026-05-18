# RacerGet200ResponsePlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Player username | 
**Score** | **int32** | Player&#39;s current score in the race | 
**Id** | Pointer to **string** | User ID. Missing if player is anonymous. | [optional] 
**Flair** | Pointer to **string** | User&#39;s flair icon | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewRacerGet200ResponsePlayersInner

`func NewRacerGet200ResponsePlayersInner(name string, score int32, ) *RacerGet200ResponsePlayersInner`

NewRacerGet200ResponsePlayersInner instantiates a new RacerGet200ResponsePlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacerGet200ResponsePlayersInnerWithDefaults

`func NewRacerGet200ResponsePlayersInnerWithDefaults() *RacerGet200ResponsePlayersInner`

NewRacerGet200ResponsePlayersInnerWithDefaults instantiates a new RacerGet200ResponsePlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RacerGet200ResponsePlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RacerGet200ResponsePlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RacerGet200ResponsePlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetScore

`func (o *RacerGet200ResponsePlayersInner) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RacerGet200ResponsePlayersInner) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RacerGet200ResponsePlayersInner) SetScore(v int32)`

SetScore sets Score field to given value.


### GetId

`func (o *RacerGet200ResponsePlayersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RacerGet200ResponsePlayersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RacerGet200ResponsePlayersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RacerGet200ResponsePlayersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFlair

`func (o *RacerGet200ResponsePlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *RacerGet200ResponsePlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *RacerGet200ResponsePlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *RacerGet200ResponsePlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *RacerGet200ResponsePlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *RacerGet200ResponsePlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *RacerGet200ResponsePlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *RacerGet200ResponsePlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *RacerGet200ResponsePlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *RacerGet200ResponsePlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *RacerGet200ResponsePlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *RacerGet200ResponsePlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


