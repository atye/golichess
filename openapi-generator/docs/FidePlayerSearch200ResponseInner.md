# FidePlayerSearch200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **string** | only appears if the user is a titled player or a bot user | [optional] 
**Federation** | **string** |  | 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Inactive** | Pointer to **int32** |  | [optional] 
**Standard** | Pointer to **int32** |  | [optional] 
**Rapid** | Pointer to **int32** |  | [optional] 
**Blitz** | Pointer to **int32** |  | [optional] 
**Gender** | Pointer to **string** | FIDE uses mandatory binary gender. | [optional] 
**Photo** | Pointer to [**BroadcastsOfficial200ResponsePhotosValue**](BroadcastsOfficial200ResponsePhotosValue.md) |  | [optional] 

## Methods

### NewFidePlayerSearch200ResponseInner

`func NewFidePlayerSearch200ResponseInner(id int32, name string, federation string, ) *FidePlayerSearch200ResponseInner`

NewFidePlayerSearch200ResponseInner instantiates a new FidePlayerSearch200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFidePlayerSearch200ResponseInnerWithDefaults

`func NewFidePlayerSearch200ResponseInnerWithDefaults() *FidePlayerSearch200ResponseInner`

NewFidePlayerSearch200ResponseInnerWithDefaults instantiates a new FidePlayerSearch200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FidePlayerSearch200ResponseInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FidePlayerSearch200ResponseInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FidePlayerSearch200ResponseInner) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FidePlayerSearch200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FidePlayerSearch200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FidePlayerSearch200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FidePlayerSearch200ResponseInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FidePlayerSearch200ResponseInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FidePlayerSearch200ResponseInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FidePlayerSearch200ResponseInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFederation

`func (o *FidePlayerSearch200ResponseInner) GetFederation() string`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *FidePlayerSearch200ResponseInner) GetFederationOk() (*string, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *FidePlayerSearch200ResponseInner) SetFederation(v string)`

SetFederation sets Federation field to given value.


### GetYear

`func (o *FidePlayerSearch200ResponseInner) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FidePlayerSearch200ResponseInner) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FidePlayerSearch200ResponseInner) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *FidePlayerSearch200ResponseInner) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *FidePlayerSearch200ResponseInner) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *FidePlayerSearch200ResponseInner) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetInactive

`func (o *FidePlayerSearch200ResponseInner) GetInactive() int32`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *FidePlayerSearch200ResponseInner) GetInactiveOk() (*int32, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *FidePlayerSearch200ResponseInner) SetInactive(v int32)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *FidePlayerSearch200ResponseInner) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetStandard

`func (o *FidePlayerSearch200ResponseInner) GetStandard() int32`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *FidePlayerSearch200ResponseInner) GetStandardOk() (*int32, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *FidePlayerSearch200ResponseInner) SetStandard(v int32)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *FidePlayerSearch200ResponseInner) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetRapid

`func (o *FidePlayerSearch200ResponseInner) GetRapid() int32`

GetRapid returns the Rapid field if non-nil, zero value otherwise.

### GetRapidOk

`func (o *FidePlayerSearch200ResponseInner) GetRapidOk() (*int32, bool)`

GetRapidOk returns a tuple with the Rapid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapid

`func (o *FidePlayerSearch200ResponseInner) SetRapid(v int32)`

SetRapid sets Rapid field to given value.

### HasRapid

`func (o *FidePlayerSearch200ResponseInner) HasRapid() bool`

HasRapid returns a boolean if a field has been set.

### GetBlitz

`func (o *FidePlayerSearch200ResponseInner) GetBlitz() int32`

GetBlitz returns the Blitz field if non-nil, zero value otherwise.

### GetBlitzOk

`func (o *FidePlayerSearch200ResponseInner) GetBlitzOk() (*int32, bool)`

GetBlitzOk returns a tuple with the Blitz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlitz

`func (o *FidePlayerSearch200ResponseInner) SetBlitz(v int32)`

SetBlitz sets Blitz field to given value.

### HasBlitz

`func (o *FidePlayerSearch200ResponseInner) HasBlitz() bool`

HasBlitz returns a boolean if a field has been set.

### GetGender

`func (o *FidePlayerSearch200ResponseInner) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *FidePlayerSearch200ResponseInner) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *FidePlayerSearch200ResponseInner) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *FidePlayerSearch200ResponseInner) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPhoto

`func (o *FidePlayerSearch200ResponseInner) GetPhoto() BroadcastsOfficial200ResponsePhotosValue`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *FidePlayerSearch200ResponseInner) GetPhotoOk() (*BroadcastsOfficial200ResponsePhotosValue, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *FidePlayerSearch200ResponseInner) SetPhoto(v BroadcastsOfficial200ResponsePhotosValue)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *FidePlayerSearch200ResponseInner) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


