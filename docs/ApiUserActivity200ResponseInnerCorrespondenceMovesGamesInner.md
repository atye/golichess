# ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Color** | **string** |  | 
**Url** | **string** |  | 
**Variant** | Pointer to **string** |  | [optional] [default to "standard"]
**Speed** | Pointer to **string** |  | [optional] 
**Perf** | Pointer to **string** |  | [optional] 
**Rated** | Pointer to **bool** |  | [optional] 
**Opponent** | [**ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent**](ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent.md) |  | 

## Methods

### NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner

`func NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner(id string, color string, url string, opponent ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent, ) *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner`

NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner instantiates a new ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerWithDefaults

`func NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerWithDefaults() *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner`

NewApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerWithDefaults instantiates a new ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetId(v string)`

SetId sets Id field to given value.


### GetColor

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetUrl

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVariant

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetVariant() string`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetVariantOk() (*string, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetVariant(v string)`

SetVariant sets Variant field to given value.

### HasVariant

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) HasVariant() bool`

HasVariant returns a boolean if a field has been set.

### GetSpeed

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.

### GetPerf

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetPerf() string`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetPerfOk() (*string, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetPerf(v string)`

SetPerf sets Perf field to given value.

### HasPerf

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) HasPerf() bool`

HasPerf returns a boolean if a field has been set.

### GetRated

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetRated(v bool)`

SetRated sets Rated field to given value.

### HasRated

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) HasRated() bool`

HasRated returns a boolean if a field has been set.

### GetOpponent

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetOpponent() ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent`

GetOpponent returns the Opponent field if non-nil, zero value otherwise.

### GetOpponentOk

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) GetOpponentOk() (*ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent, bool)`

GetOpponentOk returns a tuple with the Opponent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpponent

`func (o *ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner) SetOpponent(v ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent)`

SetOpponent sets Opponent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


