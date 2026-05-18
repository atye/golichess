# FeaturedPlayersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | **NullableString** |  | 
**User** | [**ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId**](ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md) |  | 
**Rating** | **int32** |  | 
**Seconds** | **int32** | The player&#39;s remaining time in seconds | 

## Methods

### NewFeaturedPlayersInner

`func NewFeaturedPlayersInner(color NullableString, user ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, rating int32, seconds int32, ) *FeaturedPlayersInner`

NewFeaturedPlayersInner instantiates a new FeaturedPlayersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeaturedPlayersInnerWithDefaults

`func NewFeaturedPlayersInnerWithDefaults() *FeaturedPlayersInner`

NewFeaturedPlayersInnerWithDefaults instantiates a new FeaturedPlayersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *FeaturedPlayersInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *FeaturedPlayersInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *FeaturedPlayersInner) SetColor(v string)`

SetColor sets Color field to given value.


### SetColorNil

`func (o *FeaturedPlayersInner) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *FeaturedPlayersInner) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetUser

`func (o *FeaturedPlayersInner) GetUser() ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FeaturedPlayersInner) GetUserOk() (*ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FeaturedPlayersInner) SetUser(v ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId)`

SetUser sets User field to given value.


### GetRating

`func (o *FeaturedPlayersInner) GetRating() int32`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *FeaturedPlayersInner) GetRatingOk() (*int32, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *FeaturedPlayersInner) SetRating(v int32)`

SetRating sets Rating field to given value.


### GetSeconds

`func (o *FeaturedPlayersInner) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *FeaturedPlayersInner) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *FeaturedPlayersInner) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


