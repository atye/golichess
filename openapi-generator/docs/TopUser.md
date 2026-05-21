# TopUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Username** | **string** |  | 
**Perfs** | Pointer to [**map[string]TopUserPerfsValue**](TopUserPerfsValue.md) |  | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 

## Methods

### NewTopUser

`func NewTopUser(id string, username string, ) *TopUser`

NewTopUser instantiates a new TopUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTopUserWithDefaults

`func NewTopUserWithDefaults() *TopUser`

NewTopUserWithDefaults instantiates a new TopUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TopUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TopUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TopUser) SetId(v string)`

SetId sets Id field to given value.


### GetUsername

`func (o *TopUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *TopUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *TopUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetPerfs

`func (o *TopUser) GetPerfs() map[string]TopUserPerfsValue`

GetPerfs returns the Perfs field if non-nil, zero value otherwise.

### GetPerfsOk

`func (o *TopUser) GetPerfsOk() (*map[string]TopUserPerfsValue, bool)`

GetPerfsOk returns a tuple with the Perfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerfs

`func (o *TopUser) SetPerfs(v map[string]TopUserPerfsValue)`

SetPerfs sets Perfs field to given value.

### HasPerfs

`func (o *TopUser) HasPerfs() bool`

HasPerfs returns a boolean if a field has been set.

### GetTitle

`func (o *TopUser) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TopUser) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TopUser) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TopUser) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *TopUser) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *TopUser) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *TopUser) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *TopUser) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *TopUser) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *TopUser) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *TopUser) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *TopUser) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetOnline

`func (o *TopUser) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *TopUser) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *TopUser) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *TopUser) HasOnline() bool`

HasOnline returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


