# UserActivityCorrespondenceGame

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Color** | [**GameColor**](GameColor.md) |  | 
**Url** | **string** |  | 
**Variant** | Pointer to [**VariantKey**](VariantKey.md) |  | [optional] [default to VARIANTKEY_STANDARD]
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**Opponent** | [**UserActivityCorrespondenceGameOpponent**](UserActivityCorrespondenceGameOpponent.md) |  | 

## Methods

### NewUserActivityCorrespondenceGame

`func NewUserActivityCorrespondenceGame(id string, color GameColor, url string, opponent UserActivityCorrespondenceGameOpponent, ) *UserActivityCorrespondenceGame`

NewUserActivityCorrespondenceGame instantiates a new UserActivityCorrespondenceGame object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivityCorrespondenceGameWithDefaults

`func NewUserActivityCorrespondenceGameWithDefaults() *UserActivityCorrespondenceGame`

NewUserActivityCorrespondenceGameWithDefaults instantiates a new UserActivityCorrespondenceGame object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UserActivityCorrespondenceGame) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserActivityCorrespondenceGame) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserActivityCorrespondenceGame) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *UserActivityCorrespondenceGame) GetColor() GameColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UserActivityCorrespondenceGame) GetColorOk() (*GameColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UserActivityCorrespondenceGame) SetColor(v GameColor)`

SetColor sets Color field to given value.


### GetUrl

`func (o *UserActivityCorrespondenceGame) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserActivityCorrespondenceGame) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserActivityCorrespondenceGame) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVariant

`func (o *UserActivityCorrespondenceGame) GetVariant() VariantKey`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *UserActivityCorrespondenceGame) GetVariantOk() (*VariantKey, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *UserActivityCorrespondenceGame) SetVariant(v VariantKey)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *UserActivityCorrespondenceGame) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *UserActivityCorrespondenceGame) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *UserActivityCorrespondenceGame) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *UserActivityCorrespondenceGame) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *UserActivityCorrespondenceGame) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *UserActivityCorrespondenceGame) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *UserActivityCorrespondenceGame) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *UserActivityCorrespondenceGame) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *UserActivityCorrespondenceGame) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *UserActivityCorrespondenceGame) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *UserActivityCorrespondenceGame) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *UserActivityCorrespondenceGame) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *UserActivityCorrespondenceGame) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetOpponent

`func (o *UserActivityCorrespondenceGame) GetOpponent() UserActivityCorrespondenceGameOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *UserActivityCorrespondenceGame) GetOpponentOk() (*UserActivityCorrespondenceGameOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *UserActivityCorrespondenceGame) SetOpponent(v UserActivityCorrespondenceGameOpponent)`

SetOpponent sets Opponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


