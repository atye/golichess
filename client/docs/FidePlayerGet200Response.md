# FidePlayerGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | **string** |  | 
**Title** | Pointer to **NullableString** | only appears if the user is a titled player or a bot user | [optional] 
**Federation** | **string** |  | 
**Year** | Pointer to **NullableInt32** |  | [optional] 
**Inactive** | Pointer to **int32** |  | [optional] 
**Standard** | Pointer to **int32** |  | [optional] 
**Rapid** | Pointer to **int32** |  | [optional] 
**Blitz** | Pointer to **int32** |  | [optional] 
**Gender** | Pointer to **string** | FIDE uses mandatory binary gender. | [optional] 
**Photo** | Pointer to [**BroadcastsOfficial200ResponsePhotosValue**](BroadcastsOfficial200ResponsePhotosValue.md) |  | [optional] 

## Methods

### NewFidePlayerGet200Response

`func NewFidePlayerGet200Response(id int32, name string, federation string, ) *FidePlayerGet200Response`

NewFidePlayerGet200Response instantiates a new FidePlayerGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFidePlayerGet200ResponseWithDefaults

`func NewFidePlayerGet200ResponseWithDefaults() *FidePlayerGet200Response`

NewFidePlayerGet200ResponseWithDefaults instantiates a new FidePlayerGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FidePlayerGet200Response) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FidePlayerGet200Response) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FidePlayerGet200Response) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *FidePlayerGet200Response) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FidePlayerGet200Response) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FidePlayerGet200Response) SetName(v string)`

SetName sets Name field to given value.


### GetTitle

`func (o *FidePlayerGet200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FidePlayerGet200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FidePlayerGet200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FidePlayerGet200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *FidePlayerGet200Response) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *FidePlayerGet200Response) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFederation

`func (o *FidePlayerGet200Response) GetFederation() string`

GetFederation returns the Federation field if non-nil, zero value otherwise.

### GetFederationOk

`func (o *FidePlayerGet200Response) GetFederationOk() (*string, bool)`

GetFederationOk returns a tuple with the Federation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederation

`func (o *FidePlayerGet200Response) SetFederation(v string)`

SetFederation sets Federation field to given value.


### GetYear

`func (o *FidePlayerGet200Response) GetYear() int32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *FidePlayerGet200Response) GetYearOk() (*int32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *FidePlayerGet200Response) SetYear(v int32)`

SetYear sets Year field to given value.

### HasYear

`func (o *FidePlayerGet200Response) HasYear() bool`

HasYear returns a boolean if a field has been set.

### SetYearNil

`func (o *FidePlayerGet200Response) SetYearNil(b bool)`

 SetYearNil sets the value for Year to be an explicit nil

### UnsetYear
`func (o *FidePlayerGet200Response) UnsetYear()`

UnsetYear ensures that no value is present for Year, not even an explicit nil
### GetInactive

`func (o *FidePlayerGet200Response) GetInactive() int32`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *FidePlayerGet200Response) GetInactiveOk() (*int32, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *FidePlayerGet200Response) SetInactive(v int32)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *FidePlayerGet200Response) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetStandard

`func (o *FidePlayerGet200Response) GetStandard() int32`

GetStandard returns the Standard field if non-nil, zero value otherwise.

### GetStandardOk

`func (o *FidePlayerGet200Response) GetStandardOk() (*int32, bool)`

GetStandardOk returns a tuple with the Standard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandard

`func (o *FidePlayerGet200Response) SetStandard(v int32)`

SetStandard sets Standard field to given value.

### HasStandard

`func (o *FidePlayerGet200Response) HasStandard() bool`

HasStandard returns a boolean if a field has been set.

### GetRapid

`func (o *FidePlayerGet200Response) GetRapid() int32`

GetRapid returns the Rapid field if non-nil, zero value otherwise.

### GetRapidOk

`func (o *FidePlayerGet200Response) GetRapidOk() (*int32, bool)`

GetRapidOk returns a tuple with the Rapid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRapid

`func (o *FidePlayerGet200Response) SetRapid(v int32)`

SetRapid sets Rapid field to given value.

### HasRapid

`func (o *FidePlayerGet200Response) HasRapid() bool`

HasRapid returns a boolean if a field has been set.

### GetBlitz

`func (o *FidePlayerGet200Response) GetBlitz() int32`

GetBlitz returns the Blitz field if non-nil, zero value otherwise.

### GetBlitzOk

`func (o *FidePlayerGet200Response) GetBlitzOk() (*int32, bool)`

GetBlitzOk returns a tuple with the Blitz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlitz

`func (o *FidePlayerGet200Response) SetBlitz(v int32)`

SetBlitz sets Blitz field to given value.

### HasBlitz

`func (o *FidePlayerGet200Response) HasBlitz() bool`

HasBlitz returns a boolean if a field has been set.

### GetGender

`func (o *FidePlayerGet200Response) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *FidePlayerGet200Response) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *FidePlayerGet200Response) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *FidePlayerGet200Response) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetPhoto

`func (o *FidePlayerGet200Response) GetPhoto() BroadcastsOfficial200ResponsePhotosValue`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *FidePlayerGet200Response) GetPhotoOk() (*BroadcastsOfficial200ResponsePhotosValue, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *FidePlayerGet200Response) SetPhoto(v BroadcastsOfficial200ResponsePhotosValue)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *FidePlayerGet200Response) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


