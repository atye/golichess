# FIDEPlayerPhoto

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Small** | **string** | URL of a small (100x100) thumbnail of the photo | 
**Medium** | **string** | URL of a medium (500x500) version of the photo | 
**Credit** | Pointer to **string** | If set, then you should make it appear next to the photo | [optional] 

## Methods

### NewFIDEPlayerPhoto

`func NewFIDEPlayerPhoto(small string, medium string, ) *FIDEPlayerPhoto`

NewFIDEPlayerPhoto instantiates a new FIDEPlayerPhoto object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDEPlayerPhotoWithDefaults

`func NewFIDEPlayerPhotoWithDefaults() *FIDEPlayerPhoto`

NewFIDEPlayerPhotoWithDefaults instantiates a new FIDEPlayerPhoto object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSmall

`func (o *FIDEPlayerPhoto) GetSmall() string`

GetSmall returns the Small field if non-nil, zero value otherwise.

### GetSmallOk

`func (o *FIDEPlayerPhoto) GetSmallOk() (*string, bool)`

GetSmallOk returns a tuple with the Small field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmall

`func (o *FIDEPlayerPhoto) SetSmall(v string)`

SetSmall sets Small field to given value.


### GetMedium

`func (o *FIDEPlayerPhoto) GetMedium() string`

GetMedium returns the Medium field if non-nil, zero value otherwise.

### GetMediumOk

`func (o *FIDEPlayerPhoto) GetMediumOk() (*string, bool)`

GetMediumOk returns a tuple with the Medium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedium

`func (o *FIDEPlayerPhoto) SetMedium(v string)`

SetMedium sets Medium field to given value.


### GetCredit

`func (o *FIDEPlayerPhoto) GetCredit() string`

GetCredit returns the Credit field if non-nil, zero value otherwise.

### GetCreditOk

`func (o *FIDEPlayerPhoto) GetCreditOk() (*string, bool)`

GetCreditOk returns a tuple with the Credit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredit

`func (o *FIDEPlayerPhoto) SetCredit(v string)`

SetCredit sets Credit field to given value.

### HasCredit

`func (o *FIDEPlayerPhoto) HasCredit() bool`

HasCredit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


