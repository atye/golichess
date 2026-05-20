# ApiPuzzleId200ResponseGamePlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Rating** | **int32** |  | 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 

## Methods

### NewApiPuzzleId200ResponseGamePlayersInner

`func NewApiPuzzleId200ResponseGamePlayersInner(color string, id string, name string, rating int32, ) *ApiPuzzleId200ResponseGamePlayersInner`

NewApiPuzzleId200ResponseGamePlayersInner instantiates a new ApiPuzzleId200ResponseGamePlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleId200ResponseGamePlayersInnerWithDefaults

`func NewApiPuzzleId200ResponseGamePlayersInnerWithDefaults() *ApiPuzzleId200ResponseGamePlayersInner`

NewApiPuzzleId200ResponseGamePlayersInnerWithDefaults instantiates a new ApiPuzzleId200ResponseGamePlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetFlair

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiPuzzleId200ResponseGamePlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetId

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetPatron

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiPuzzleId200ResponseGamePlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiPuzzleId200ResponseGamePlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetRating

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetTitle

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiPuzzleId200ResponseGamePlayersInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiPuzzleId200ResponseGamePlayersInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiPuzzleId200ResponseGamePlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


