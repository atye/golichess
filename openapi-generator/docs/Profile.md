# Profile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flag** | Pointer to **string** |  | [optional] 
**Location** | Pointer to **string** |  | [optional] 
**Bio** | Pointer to **string** |  | [optional] 
**RealName** | Pointer to **string** |  | [optional] 
**FideRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**UscfRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**EcfRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**CfcRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**RcfRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**DsbRating** | Pointer to **int32** | only appears if a user has set them | [optional] 
**Links** | Pointer to **string** |  | [optional] 

## Methods

### NewProfile

`func NewProfile() *Profile`

NewProfile instantiates a new Profile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileWithDefaults

`func NewProfileWithDefaults() *Profile`

NewProfileWithDefaults instantiates a new Profile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlag

`func (o *Profile) GetFlag() string`

GetFlag returns the Flag field if non-nil, zero value otherwise.

### GetFlagOk

`func (o *Profile) GetFlagOk() (*string, bool)`

GetFlagOk returns a tuple with the Flag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlag

`func (o *Profile) SetFlag(v string)`

SetFlag sets Flag field to given value.

### HasFlag

`func (o *Profile) HasFlag() bool`

HasFlag returns a boolean if a field has been set.

### GetLocation

`func (o *Profile) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Profile) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Profile) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Profile) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetBio

`func (o *Profile) GetBio() string`

GetBio returns the Bio field if non-nil, zero value otherwise.

### GetBioOk

`func (o *Profile) GetBioOk() (*string, bool)`

GetBioOk returns a tuple with the Bio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBio

`func (o *Profile) SetBio(v string)`

SetBio sets Bio field to given value.

### HasBio

`func (o *Profile) HasBio() bool`

HasBio returns a boolean if a field has been set.

### GetRealName

`func (o *Profile) GetRealName() string`

GetRealName returns the RealName field if non-nil, zero value otherwise.

### GetRealNameOk

`func (o *Profile) GetRealNameOk() (*string, bool)`

GetRealNameOk returns a tuple with the RealName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealName

`func (o *Profile) SetRealName(v string)`

SetRealName sets RealName field to given value.

### HasRealName

`func (o *Profile) HasRealName() bool`

HasRealName returns a boolean if a field has been set.

### GetFideRating

`func (o *Profile) GetFideRating() int32`

GetFideRating returns the FideRating field if non-nil, zero value otherwise.

### GetFideRatingOk

`func (o *Profile) GetFideRatingOk() (*int32, bool)`

GetFideRatingOk returns a tuple with the FideRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFideRating

`func (o *Profile) SetFideRating(v int32)`

SetFideRating sets FideRating field to given value.

### HasFideRating

`func (o *Profile) HasFideRating() bool`

HasFideRating returns a boolean if a field has been set.

### GetUscfRating

`func (o *Profile) GetUscfRating() int32`

GetUscfRating returns the UscfRating field if non-nil, zero value otherwise.

### GetUscfRatingOk

`func (o *Profile) GetUscfRatingOk() (*int32, bool)`

GetUscfRatingOk returns a tuple with the UscfRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUscfRating

`func (o *Profile) SetUscfRating(v int32)`

SetUscfRating sets UscfRating field to given value.

### HasUscfRating

`func (o *Profile) HasUscfRating() bool`

HasUscfRating returns a boolean if a field has been set.

### GetEcfRating

`func (o *Profile) GetEcfRating() int32`

GetEcfRating returns the EcfRating field if non-nil, zero value otherwise.

### GetEcfRatingOk

`func (o *Profile) GetEcfRatingOk() (*int32, bool)`

GetEcfRatingOk returns a tuple with the EcfRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEcfRating

`func (o *Profile) SetEcfRating(v int32)`

SetEcfRating sets EcfRating field to given value.

### HasEcfRating

`func (o *Profile) HasEcfRating() bool`

HasEcfRating returns a boolean if a field has been set.

### GetCfcRating

`func (o *Profile) GetCfcRating() int32`

GetCfcRating returns the CfcRating field if non-nil, zero value otherwise.

### GetCfcRatingOk

`func (o *Profile) GetCfcRatingOk() (*int32, bool)`

GetCfcRatingOk returns a tuple with the CfcRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfcRating

`func (o *Profile) SetCfcRating(v int32)`

SetCfcRating sets CfcRating field to given value.

### HasCfcRating

`func (o *Profile) HasCfcRating() bool`

HasCfcRating returns a boolean if a field has been set.

### GetRcfRating

`func (o *Profile) GetRcfRating() int32`

GetRcfRating returns the RcfRating field if non-nil, zero value otherwise.

### GetRcfRatingOk

`func (o *Profile) GetRcfRatingOk() (*int32, bool)`

GetRcfRatingOk returns a tuple with the RcfRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRcfRating

`func (o *Profile) SetRcfRating(v int32)`

SetRcfRating sets RcfRating field to given value.

### HasRcfRating

`func (o *Profile) HasRcfRating() bool`

HasRcfRating returns a boolean if a field has been set.

### GetDsbRating

`func (o *Profile) GetDsbRating() int32`

GetDsbRating returns the DsbRating field if non-nil, zero value otherwise.

### GetDsbRatingOk

`func (o *Profile) GetDsbRatingOk() (*int32, bool)`

GetDsbRatingOk returns a tuple with the DsbRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsbRating

`func (o *Profile) SetDsbRating(v int32)`

SetDsbRating sets DsbRating field to given value.

### HasDsbRating

`func (o *Profile) HasDsbRating() bool`

HasDsbRating returns a boolean if a field has been set.

### GetLinks

`func (o *Profile) GetLinks() string`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *Profile) GetLinksOk() (*string, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *Profile) SetLinks(v string)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *Profile) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


