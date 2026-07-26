# UserStreamer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Twitch** | Pointer to [**UserStreamerTwitch**](UserStreamerTwitch.md) |  | [optional] 
**Youtube** | Pointer to [**UserStreamerYoutube**](UserStreamerYoutube.md) |  | [optional] 

## Methods

### NewUserStreamer

`func NewUserStreamer() *UserStreamer`

NewUserStreamer instantiates a new UserStreamer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserStreamerWithDefaults

`func NewUserStreamerWithDefaults() *UserStreamer`

NewUserStreamerWithDefaults instantiates a new UserStreamer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTwitch

`func (o *UserStreamer) GetTwitch() UserStreamerTwitch`

GetTwitch returns the Twitch field if non-nil, zero value otherwise.

### GetTwitchOk

`func (o *UserStreamer) GetTwitchOk() (*UserStreamerTwitch, bool)`

GetTwitchOk returns a tuple with the Twitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitch

`func (o *UserStreamer) SetTwitch(v UserStreamerTwitch)`

SetTwitch sets Twitch field to given value.

### HasTwitch

`func (o *UserStreamer) HasTwitch() bool`

HasTwitch returns a boolean if a field has been set.

### GetYoutube

`func (o *UserStreamer) GetYoutube() UserStreamerYoutube`

GetYoutube returns the Youtube field if non-nil, zero value otherwise.

### GetYoutubeOk

`func (o *UserStreamer) GetYoutubeOk() (*UserStreamerYoutube, bool)`

GetYoutubeOk returns a tuple with the Youtube field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYoutube

`func (o *UserStreamer) SetYoutube(v UserStreamerYoutube)`

SetYoutube sets Youtube field to given value.

### HasYoutube

`func (o *UserStreamer) HasYoutube() bool`

HasYoutube returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


