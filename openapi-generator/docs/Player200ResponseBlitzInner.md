# Player200ResponseBlitzInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Perfs** | Pointer to [**map[string]Player200ResponseBulletInnerPerfsValue**](Player200ResponseBulletInnerPerfsValue.md) |  | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 

## Methods

### NewPlayer200ResponseBlitzInner

`func NewPlayer200ResponseBlitzInner(id string, username string, ) *Player200ResponseBlitzInner`

NewPlayer200ResponseBlitzInner instantiates a new Player200ResponseBlitzInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlayer200ResponseBlitzInnerWithDefaults

`func NewPlayer200ResponseBlitzInnerWithDefaults() *Player200ResponseBlitzInner`

NewPlayer200ResponseBlitzInnerWithDefaults instantiates a new Player200ResponseBlitzInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Player200ResponseBlitzInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Player200ResponseBlitzInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Player200ResponseBlitzInner) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *Player200ResponseBlitzInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Player200ResponseBlitzInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Player200ResponseBlitzInner) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *Player200ResponseBlitzInner) GetPerfs() map[string]Player200ResponseBulletInnerPerfsValue`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *Player200ResponseBlitzInner) GetPerfsOk() (*map[string]Player200ResponseBulletInnerPerfsValue, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *Player200ResponseBlitzInner) SetPerfs(v map[string]Player200ResponseBulletInnerPerfsValue)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *Player200ResponseBlitzInner) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *Player200ResponseBlitzInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Player200ResponseBlitzInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Player200ResponseBlitzInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Player200ResponseBlitzInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *Player200ResponseBlitzInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *Player200ResponseBlitzInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *Player200ResponseBlitzInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *Player200ResponseBlitzInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *Player200ResponseBlitzInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *Player200ResponseBlitzInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *Player200ResponseBlitzInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *Player200ResponseBlitzInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetOnline

`func (o *Player200ResponseBlitzInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *Player200ResponseBlitzInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *Player200ResponseBlitzInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *Player200ResponseBlitzInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


