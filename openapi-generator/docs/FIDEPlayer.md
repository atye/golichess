# FIDEPlayer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | Pointer to [**Title**](Title.md) |  | [optional] 
**Federation** | **string** |  | 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Inactive** | Pointer to **int32** |  | [optional] 
**Standard** | Pointer to **int32** |  | [optional] 
**Rapid** | Pointer to **int32** |  | [optional] 
**Blitz** | Pointer to **int32** |  | [optional] 
**Gender** | Pointer to **string** | FIDE uses mandatory binary gender. | [optional] 
**Photo** | Pointer to [**FIDEPlayerPhoto**](FIDEPlayerPhoto.md) |  | [optional] 

## Methods

### NewFIDEPlayer

`func NewFIDEPlayer(id int32, name string, federation string, ) *FIDEPlayer`

NewFIDEPlayer instantiates a new FIDEPlayer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFIDEPlayerWithDefaults

`func NewFIDEPlayerWithDefaults() *FIDEPlayer`

NewFIDEPlayerWithDefaults instantiates a new FIDEPlayer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FIDEPlayer) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FIDEPlayer) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FIDEPlayer) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FIDEPlayer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FIDEPlayer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FIDEPlayer) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FIDEPlayer) GetTitle() Title`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FIDEPlayer) GetTitleOk() (*Title, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FIDEPlayer) SetTitle(v Title)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FIDEPlayer) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFederation

`func (o *FIDEPlayer) GetFederation() string`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *FIDEPlayer) GetFederationOk() (*string, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *FIDEPlayer) SetFederation(v string)`

SetFederation sets Federation field to given value.


### GetYear

`func (o *FIDEPlayer) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FIDEPlayer) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FIDEPlayer) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *FIDEPlayer) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *FIDEPlayer) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *FIDEPlayer) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetInactive

`func (o *FIDEPlayer) GetInactive() int32`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *FIDEPlayer) GetInactiveOk() (*int32, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *FIDEPlayer) SetInactive(v int32)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *FIDEPlayer) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetStandard

`func (o *FIDEPlayer) GetStandard() int32`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *FIDEPlayer) GetStandardOk() (*int32, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *FIDEPlayer) SetStandard(v int32)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *FIDEPlayer) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetRapid

`func (o *FIDEPlayer) GetRapid() int32`

GetRapid returns the Rapid field if non-nil, zero value otherwise.

### GetRapidOk

`func (o *FIDEPlayer) GetRapidOk() (*int32, bool)`

GetRapidOk returns a tuple with the Rapid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapid

`func (o *FIDEPlayer) SetRapid(v int32)`

SetRapid sets Rapid field to given value.

### HasRapid

`func (o *FIDEPlayer) HasRapid() bool`

HasRapid returns a boolean if a field has been set.

### GetBlitz

`func (o *FIDEPlayer) GetBlitz() int32`

GetBlitz returns the Blitz field if non-nil, zero value otherwise.

### GetBlitzOk

`func (o *FIDEPlayer) GetBlitzOk() (*int32, bool)`

GetBlitzOk returns a tuple with the Blitz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlitz

`func (o *FIDEPlayer) SetBlitz(v int32)`

SetBlitz sets Blitz field to given value.

### HasBlitz

`func (o *FIDEPlayer) HasBlitz() bool`

HasBlitz returns a boolean if a field has been set.

### GetGender

`func (o *FIDEPlayer) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *FIDEPlayer) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *FIDEPlayer) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *FIDEPlayer) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPhoto

`func (o *FIDEPlayer) GetPhoto() FIDEPlayerPhoto`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *FIDEPlayer) GetPhotoOk() (*FIDEPlayerPhoto, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *FIDEPlayer) SetPhoto(v FIDEPlayerPhoto)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *FIDEPlayer) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


