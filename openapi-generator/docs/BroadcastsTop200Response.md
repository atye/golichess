# BroadcastsTop200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to [**[]BroadcastsTop200ResponseActiveInner**](BroadcastsTop200ResponseActiveInner.md) |  | [optional] 
**Upcoming** | Pointer to [**[]BroadcastsTop200ResponseActiveInner**](BroadcastsTop200ResponseActiveInner.md) |  | [optional] 
**Past** | Pointer to [**BroadcastsTop200ResponsePast**](BroadcastsTop200ResponsePast.md) |  | [optional] 

## Methods

### NewBroadcastsTop200Response

`func NewBroadcastsTop200Response() *BroadcastsTop200Response`

NewBroadcastsTop200Response instantiates a new BroadcastsTop200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastsTop200ResponseWithDefaults

`func NewBroadcastsTop200ResponseWithDefaults() *BroadcastsTop200Response`

NewBroadcastsTop200ResponseWithDefaults instantiates a new BroadcastsTop200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *BroadcastsTop200Response) GetActive() []BroadcastsTop200ResponseActiveInner`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *BroadcastsTop200Response) GetActiveOk() (*[]BroadcastsTop200ResponseActiveInner, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *BroadcastsTop200Response) SetActive(v []BroadcastsTop200ResponseActiveInner)`

SetActive sets Active field to given value.

### HasActive

`func (o *BroadcastsTop200Response) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetUpcoming

`func (o *BroadcastsTop200Response) GetUpcoming() []BroadcastsTop200ResponseActiveInner`

GetUpcoming returns the Upcoming field if non-nil, zero value otherwise.

### GetUpcomingOk

`func (o *BroadcastsTop200Response) GetUpcomingOk() (*[]BroadcastsTop200ResponseActiveInner, bool)`

GetUpcomingOk returns a tuple with the Upcoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpcoming

`func (o *BroadcastsTop200Response) SetUpcoming(v []BroadcastsTop200ResponseActiveInner)`

SetUpcoming sets Upcoming field to given value.

### HasUpcoming

`func (o *BroadcastsTop200Response) HasUpcoming() bool`

HasUpcoming returns a boolean if a field has been set.

### GetPast

`func (o *BroadcastsTop200Response) GetPast() BroadcastsTop200ResponsePast`

GetPast returns the Past field if non-nil, zero value otherwise.

### GetPastOk

`func (o *BroadcastsTop200Response) GetPastOk() (*BroadcastsTop200ResponsePast, bool)`

GetPastOk returns a tuple with the Past field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPast

`func (o *BroadcastsTop200Response) SetPast(v BroadcastsTop200ResponsePast)`

SetPast sets Past field to given value.

### HasPast

`func (o *BroadcastsTop200Response) HasPast() bool`

HasPast returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


