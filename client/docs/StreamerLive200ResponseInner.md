# StreamerLive200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 
**Stream** | Pointer to [**StreamerLive200ResponseInnerAllOfStream**](StreamerLive200ResponseInnerAllOfStream.md) |  | [optional] 
**Streamer** | Pointer to [**StreamerLive200ResponseInnerAllOfStreamer**](StreamerLive200ResponseInnerAllOfStreamer.md) |  | [optional] 

## Methods

### NewStreamerLive200ResponseInner

`func NewStreamerLive200ResponseInner(id string, name string, ) *StreamerLive200ResponseInner`

NewStreamerLive200ResponseInner instantiates a new StreamerLive200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStreamerLive200ResponseInnerWithDefaults

`func NewStreamerLive200ResponseInnerWithDefaults() *StreamerLive200ResponseInner`

NewStreamerLive200ResponseInnerWithDefaults instantiates a new StreamerLive200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StreamerLive200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StreamerLive200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StreamerLive200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *StreamerLive200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StreamerLive200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StreamerLive200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *StreamerLive200ResponseInner) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *StreamerLive200ResponseInner) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *StreamerLive200ResponseInner) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *StreamerLive200ResponseInner) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *StreamerLive200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StreamerLive200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StreamerLive200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StreamerLive200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *StreamerLive200ResponseInner) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *StreamerLive200ResponseInner) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *StreamerLive200ResponseInner) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *StreamerLive200ResponseInner) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *StreamerLive200ResponseInner) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *StreamerLive200ResponseInner) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *StreamerLive200ResponseInner) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *StreamerLive200ResponseInner) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.

### GetStream

`func (o *StreamerLive200ResponseInner) GetStream() StreamerLive200ResponseInnerAllOfStream`

GetStream returns the Stream field if non-nil, zero value otherwise.

### GetStreamOk

`func (o *StreamerLive200ResponseInner) GetStreamOk() (*StreamerLive200ResponseInnerAllOfStream, bool)`

GetStreamOk returns a tuple with the Stream field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStream

`func (o *StreamerLive200ResponseInner) SetStream(v StreamerLive200ResponseInnerAllOfStream)`

SetStream sets Stream field to given value.

### HasStream

`func (o *StreamerLive200ResponseInner) HasStream() bool`

HasStream returns a boolean if a field has been set.

### GetStreamer

`func (o *StreamerLive200ResponseInner) GetStreamer() StreamerLive200ResponseInnerAllOfStreamer`

GetStreamer returns the Streamer field if non-nil, zero value otherwise.

### GetStreamerOk

`func (o *StreamerLive200ResponseInner) GetStreamerOk() (*StreamerLive200ResponseInnerAllOfStreamer, bool)`

GetStreamerOk returns a tuple with the Streamer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamer

`func (o *StreamerLive200ResponseInner) SetStreamer(v StreamerLive200ResponseInnerAllOfStreamer)`

SetStreamer sets Streamer field to given value.

### HasStreamer

`func (o *StreamerLive200ResponseInner) HasStreamer() bool`

HasStreamer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


