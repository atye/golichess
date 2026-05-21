# TimelineUsersValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Flair** | Pointer to **string** | See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair) | [optional] 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Patron** | Pointer to **bool** | Use patronColor value instead to determine if player is a patron.  | [optional] 
**PatronColor** | Pointer to **int32** | Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron.  | [optional] 

## Methods

### NewTimelineUsersValue

`func NewTimelineUsersValue(id string, name string, ) *TimelineUsersValue`

NewTimelineUsersValue instantiates a new TimelineUsersValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTimelineUsersValueWithDefaults

`func NewTimelineUsersValueWithDefaults() *TimelineUsersValue`

NewTimelineUsersValueWithDefaults instantiates a new TimelineUsersValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TimelineUsersValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TimelineUsersValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TimelineUsersValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *TimelineUsersValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TimelineUsersValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TimelineUsersValue) SetName(v string)`

SetName sets Name field to given value.


### GetFlair

`func (o *TimelineUsersValue) GetFlair() string`

GetFlair returns the Flair field if non-nil, zero value otherwise.

### GetFlairOk

`func (o *TimelineUsersValue) GetFlairOk() (*string, bool)`

GetFlairOk returns a tuple with the Flair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlair

`func (o *TimelineUsersValue) SetFlair(v string)`

SetFlair sets Flair field to given value.

### HasFlair

`func (o *TimelineUsersValue) HasFlair() bool`

HasFlair returns a boolean if a field has been set.

### GetTitle

`func (o *TimelineUsersValue) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TimelineUsersValue) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TimelineUsersValue) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TimelineUsersValue) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPatron

`func (o *TimelineUsersValue) GetPatron() bool`

GetPatron returns the Patron field if non-nil, zero value otherwise.

### GetPatronOk

`func (o *TimelineUsersValue) GetPatronOk() (*bool, bool)`

GetPatronOk returns a tuple with the Patron field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatron

`func (o *TimelineUsersValue) SetPatron(v bool)`

SetPatron sets Patron field to given value.

### HasPatron

`func (o *TimelineUsersValue) HasPatron() bool`

HasPatron returns a boolean if a field has been set.

### GetPatronColor

`func (o *TimelineUsersValue) GetPatronColor() int32`

GetPatronColor returns the PatronColor field if non-nil, zero value otherwise.

### GetPatronColorOk

`func (o *TimelineUsersValue) GetPatronColorOk() (*int32, bool)`

GetPatronColorOk returns a tuple with the PatronColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatronColor

`func (o *TimelineUsersValue) SetPatronColor(v int32)`

SetPatronColor sets PatronColor field to given value.

### HasPatronColor

`func (o *TimelineUsersValue) HasPatronColor() bool`

HasPatronColor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


