# BroadcastPhotosValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Small** | **string** | URL of a small (100x100) thumbnail of the photo | 
**Medium** | **string** | URL of a medium (500x500) version of the photo | 
**Credit** | Pointer to **string** | If set, then you should make it appear next to the photo | [optional] 

## Methods

### NewBroadcastPhotosValue

`func NewBroadcastPhotosValue(small string, medium string, ) *BroadcastPhotosValue`

NewBroadcastPhotosValue instantiates a new BroadcastPhotosValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBroadcastPhotosValueWithDefaults

`func NewBroadcastPhotosValueWithDefaults() *BroadcastPhotosValue`

NewBroadcastPhotosValueWithDefaults instantiates a new BroadcastPhotosValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmall

`func (o *BroadcastPhotosValue) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *BroadcastPhotosValue) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *BroadcastPhotosValue) SetSmall(v string)`

SetSmall sets Small field to given value.


### GetMedium

`func (o *BroadcastPhotosValue) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *BroadcastPhotosValue) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *BroadcastPhotosValue) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetCredit

`func (o *BroadcastPhotosValue) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *BroadcastPhotosValue) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *BroadcastPhotosValue) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *BroadcastPhotosValue) HasCredit() bool`

HasCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


