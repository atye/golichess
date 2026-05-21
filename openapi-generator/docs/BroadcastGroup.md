# BroadcastGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Slug** | **string** |  | 
**Name** | **string** |  | 
**Tours** | [**[]BroadcastGroupTour**](BroadcastGroupTour.md) |  | 

## Methods

### NewBroadcastGroup

`func NewBroadcastGroup(id string, slug string, name string, tours []BroadcastGroupTour, ) *BroadcastGroup`

NewBroadcastGroup instantiates a new BroadcastGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastGroupWithDefaults

`func NewBroadcastGroupWithDefaults() *BroadcastGroup`

NewBroadcastGroupWithDefaults instantiates a new BroadcastGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BroadcastGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BroadcastGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BroadcastGroup) SetId(v string)`

SetId sets Id field to given value.


### GetSlug

`func (o *BroadcastGroup) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *BroadcastGroup) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *BroadcastGroup) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetName

`func (o *BroadcastGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastGroup) SetName(v string)`

SetName sets Name field to given value.


### GetTours

`func (o *BroadcastGroup) GetTours() []BroadcastGroupTour`

GetTours returns the Tours field if non-nil, zero value otherwise.

### GetToursOk

`func (o *BroadcastGroup) GetToursOk() (*[]BroadcastGroupTour, bool)`

GetToursOk returns a tuple with the Tours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTours

`func (o *BroadcastGroup) SetTours(v []BroadcastGroupTour)`

SetTours sets Tours field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


