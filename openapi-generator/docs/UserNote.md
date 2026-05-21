# UserNote

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 
**To** | Pointer to [**LightUser**](LightUser.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **int64** |  | [optional] 

## Methods

### NewUserNote

`func NewUserNote() *UserNote`

NewUserNote instantiates a new UserNote object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserNoteWithDefaults

`func NewUserNoteWithDefaults() *UserNote`

NewUserNoteWithDefaults instantiates a new UserNote object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *UserNote) GetFrom() LightUser`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *UserNote) GetFromOk() (*LightUser, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *UserNote) SetFrom(v LightUser)`

SetFrom sets From field to given value.

### HasFrom

`func (o *UserNote) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *UserNote) GetTo() LightUser`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *UserNote) GetToOk() (*LightUser, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *UserNote) SetTo(v LightUser)`

SetTo sets To field to given value.

### HasTo

`func (o *UserNote) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetText

`func (o *UserNote) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *UserNote) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *UserNote) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *UserNote) HasText() bool`

HasText returns a boolean if a field has been set.

### GetDate

`func (o *UserNote) GetDate() int64`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UserNote) GetDateOk() (*int64, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UserNote) SetDate(v int64)`

SetDate sets Date field to given value.

### HasDate

`func (o *UserNote) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


