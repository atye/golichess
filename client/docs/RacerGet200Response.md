# RacerGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier of the puzzle race | 
**Owner** | **string** | Owner of the puzzle race | 
**Players** | [**[]RacerGet200ResponsePlayersInner**](RacerGet200ResponsePlayersInner.md) | List of players participating in the race | 
**Puzzles** | [**[]RacerGet200ResponsePuzzlesInner**](RacerGet200ResponsePuzzlesInner.md) | List of puzzles in the race | 
**FinishesAt** | **int32** | Timestamp in milliseconds when the race finishes | 
**StartsAt** | **int32** | Timestamp in milliseconds when the race started | 

## Methods

### NewRacerGet200Response

`func NewRacerGet200Response(id string, owner string, players []RacerGet200ResponsePlayersInner, puzzles []RacerGet200ResponsePuzzlesInner, finishesAt int32, startsAt int32, ) *RacerGet200Response`

NewRacerGet200Response instantiates a new RacerGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRacerGet200ResponseWithDefaults

`func NewRacerGet200ResponseWithDefaults() *RacerGet200Response`

NewRacerGet200ResponseWithDefaults instantiates a new RacerGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RacerGet200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RacerGet200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RacerGet200Response) SetId(v string)`

SetId sets Id field to given value.


### GetOwner

`func (o *RacerGet200Response) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *RacerGet200Response) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *RacerGet200Response) SetOwner(v string)`

SetOwner sets Owner field to given value.


### GetPlayers

`func (o *RacerGet200Response) GetPlayers() []RacerGet200ResponsePlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *RacerGet200Response) GetPlayersOk() (*[]RacerGet200ResponsePlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *RacerGet200Response) SetPlayers(v []RacerGet200ResponsePlayersInner)`

SetPlayers sets Players field to given value.


### GetPuzzles

`func (o *RacerGet200Response) GetPuzzles() []RacerGet200ResponsePuzzlesInner`

GetPuzzles returns the Puzzles field if non-nil, zero value otherwise.

### GetPuzzlesOk

`func (o *RacerGet200Response) GetPuzzlesOk() (*[]RacerGet200ResponsePuzzlesInner, bool)`

GetPuzzlesOk returns a tuple with the Puzzles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPuzzles

`func (o *RacerGet200Response) SetPuzzles(v []RacerGet200ResponsePuzzlesInner)`

SetPuzzles sets Puzzles field to given value.


### GetFinishesAt

`func (o *RacerGet200Response) GetFinishesAt() int32`

GetFinishesAt returns the FinishesAt field if non-nil, zero value otherwise.

### GetFinishesAtOk

`func (o *RacerGet200Response) GetFinishesAtOk() (*int32, bool)`

GetFinishesAtOk returns a tuple with the FinishesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishesAt

`func (o *RacerGet200Response) SetFinishesAt(v int32)`

SetFinishesAt sets FinishesAt field to given value.


### GetStartsAt

`func (o *RacerGet200Response) GetStartsAt() int32`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *RacerGet200Response) GetStartsAtOk() (*int32, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *RacerGet200Response) SetStartsAt(v int32)`

SetStartsAt sets StartsAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


