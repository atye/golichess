# ApiUsersStatus200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Online** | Pointer to **bool** |  | [optional] 
**Playing** | Pointer to **bool** |  | [optional] 
**Streaming** | Pointer to **bool** |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewApiUsersStatus200ResponseInner

`func NewApiUsersStatus200ResponseInner(id string, name string, ) *ApiUsersStatus200ResponseInner`

NewApiUsersStatus200ResponseInner instantiates a new ApiUsersStatus200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUsersStatus200ResponseInnerWithDefaults

`func NewApiUsersStatus200ResponseInnerWithDefaults() *ApiUsersStatus200ResponseInner`

NewApiUsersStatus200ResponseInnerWithDefaults instantiates a new ApiUsersStatus200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUsersStatus200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUsersStatus200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUsersStatus200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiUsersStatus200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiUsersStatus200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiUsersStatus200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *ApiUsersStatus200ResponseInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *ApiUsersStatus200ResponseInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *ApiUsersStatus200ResponseInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *ApiUsersStatus200ResponseInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *ApiUsersStatus200ResponseInner) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ApiUsersStatus200ResponseInner) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ApiUsersStatus200ResponseInner) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ApiUsersStatus200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetOnline

`func (o *ApiUsersStatus200ResponseInner) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *ApiUsersStatus200ResponseInner) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *ApiUsersStatus200ResponseInner) SetOnline(v bool)`

SetOnline sets Online field to given value.

### HasOnline

`func (o *ApiUsersStatus200ResponseInner) HasOnline() bool`

HasOnline returns a boolean if a field has been set.

### GetPlaying

`func (o *ApiUsersStatus200ResponseInner) GetPlaying() bool`

GetPlaying returns the Playing field if non-nil, zero value otherwise.

### GetPlayingOk

`func (o *ApiUsersStatus200ResponseInner) GetPlayingOk() (*bool, bool)`

GetPlayingOk returns a tuple with the Playing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaying

`func (o *ApiUsersStatus200ResponseInner) SetPlaying(v bool)`

SetPlaying sets Playing field to given value.

### HasPlaying

`func (o *ApiUsersStatus200ResponseInner) HasPlaying() bool`

HasPlaying returns a boolean if a field has been set.

### GetStreaming

`func (o *ApiUsersStatus200ResponseInner) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *ApiUsersStatus200ResponseInner) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *ApiUsersStatus200ResponseInner) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.

### HasStreaming

`func (o *ApiUsersStatus200ResponseInner) HasStreaming() bool`

HasStreaming returns a boolean if a field has been set.

### GetPatron

`func (o *ApiUsersStatus200ResponseInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *ApiUsersStatus200ResponseInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *ApiUsersStatus200ResponseInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *ApiUsersStatus200ResponseInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *ApiUsersStatus200ResponseInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *ApiUsersStatus200ResponseInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *ApiUsersStatus200ResponseInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *ApiUsersStatus200ResponseInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


