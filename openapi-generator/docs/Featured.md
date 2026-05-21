# Featured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The game ID | 
**Orientation** | [**GameColor**](GameColor.md) |  | 
**Players** | [**[]FeaturedPlayersInner**](FeaturedPlayersInner.md) |  | 
**Fen** | **string** | The X-FEN of the current position | 

## Methods

### NewFeatured

`func NewFeatured(id string, orientation GameColor, players []FeaturedPlayersInner, fen string, ) *Featured`

NewFeatured instantiates a new Featured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturedWithDefaults

`func NewFeaturedWithDefaults() *Featured`

NewFeaturedWithDefaults instantiates a new Featured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Featured) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Featured) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Featured) SetId(v string)`

SetId sets Id field to given value.


### GetOrientation

`func (o *Featured) GetOrientation() GameColor`

GetOrientation returns the Orientation field if non-nil, zero value otherwise.

### GetOrientationOk

`func (o *Featured) GetOrientationOk() (*GameColor, bool)`

GetOrientationOk returns a tuple with the Orientation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrientation

`func (o *Featured) SetOrientation(v GameColor)`

SetOrientation sets Orientation field to given value.


### GetPlayers

`func (o *Featured) GetPlayers() []FeaturedPlayersInner`

GetPlayers returns the Players field if non-nil, zero value otherwise.

### GetPlayersOk

`func (o *Featured) GetPlayersOk() (*[]FeaturedPlayersInner, bool)`

GetPlayersOk returns a tuple with the Players field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayers

`func (o *Featured) SetPlayers(v []FeaturedPlayersInner)`

SetPlayers sets Players field to given value.


### GetFen

`func (o *Featured) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *Featured) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *Featured) SetFen(v string)`

SetFen sets Fen field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


