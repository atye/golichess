# BroadcastFormGroupingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the group | [optional] 
**Tours** | Pointer to **string** | A linebreak separated list of tournament IDs to group together.  | [optional] 

## Methods

### NewBroadcastFormGroupingInfo

`func NewBroadcastFormGroupingInfo() *BroadcastFormGroupingInfo`

NewBroadcastFormGroupingInfo instantiates a new BroadcastFormGroupingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastFormGroupingInfoWithDefaults

`func NewBroadcastFormGroupingInfoWithDefaults() *BroadcastFormGroupingInfo`

NewBroadcastFormGroupingInfoWithDefaults instantiates a new BroadcastFormGroupingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BroadcastFormGroupingInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BroadcastFormGroupingInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BroadcastFormGroupingInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BroadcastFormGroupingInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTours

`func (o *BroadcastFormGroupingInfo) GetTours() string`

GetTours returns the Tours field if non-nil, zero value otherwise.

### GetToursOk

`func (o *BroadcastFormGroupingInfo) GetToursOk() (*string, bool)`

GetToursOk returns a tuple with the Tours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTours

`func (o *BroadcastFormGroupingInfo) SetTours(v string)`

SetTours sets Tours field to given value.

### HasTours

`func (o *BroadcastFormGroupingInfo) HasTours() bool`

HasTours returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


