# ApiPuzzleDaily200ResponseGamePlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **NullableString** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Rating** | **int32** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 

## Methods

### NewApiPuzzleDaily200ResponseGamePlayersInner

`func NewApiPuzzleDaily200ResponseGamePlayersInner(color NullableString, id string, name string, rating int32, ) *ApiPuzzleDaily200ResponseGamePlayersInner`

NewApiPuzzleDaily200ResponseGamePlayersInner instantiates a new ApiPuzzleDaily200ResponseGamePlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiPuzzleDaily200ResponseGamePlayersInnerWithDefaults

`func NewApiPuzzleDaily200ResponseGamePlayersInnerWithDefaults() *ApiPuzzleDaily200ResponseGamePlayersInner`

NewApiPuzzleDaily200ResponseGamePlayersInnerWithDefaults instantiates a new ApiPuzzleDaily200ResponseGamePlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetFlair

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetId

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetName(v string)`

SetName sets Name field to given value.


### GetPatron

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetRating

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetTitle

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ApiPuzzleDaily200ResponseGamePlayersInner) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


