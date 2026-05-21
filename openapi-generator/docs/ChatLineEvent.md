# ChatLineEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Room** | **string** |  | 
**Username** | **string** |  | 
**Text** | **string** |  | 

## Methods

### NewChatLineEvent

`func NewChatLineEvent(type_ string, room string, username string, text string, ) *ChatLineEvent`

NewChatLineEvent instantiates a new ChatLineEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatLineEventWithDefaults

`func NewChatLineEventWithDefaults() *ChatLineEvent`

NewChatLineEventWithDefaults instantiates a new ChatLineEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChatLineEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChatLineEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChatLineEvent) SetType(v string)`

SetType sets Type field to given value.


### GetRoom

`func (o *ChatLineEvent) GetRoom() string`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *ChatLineEvent) GetRoomOk() (*string, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *ChatLineEvent) SetRoom(v string)`

SetRoom sets Room field to given value.


### GetUsername

`func (o *ChatLineEvent) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ChatLineEvent) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ChatLineEvent) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetText

`func (o *ChatLineEvent) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ChatLineEvent) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ChatLineEvent) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


