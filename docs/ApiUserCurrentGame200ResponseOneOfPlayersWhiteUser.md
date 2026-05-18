# ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUser

`func NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUser(id string, name string, ) *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser`

NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUser instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUserWithDefaults

`func NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUserWithDefaults() *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser`

NewApiUserCurrentGame200ResponseOneOfPlayersWhiteUserWithDefaults instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFlair

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


