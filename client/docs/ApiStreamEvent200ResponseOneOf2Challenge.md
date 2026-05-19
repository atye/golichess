# ApiStreamEvent200ResponseOneOf2Challenge

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | **string** |  | 
**Challenger** | [**ApiStreamEvent200ResponseOneOf2ChallengeChallenger**](ApiStreamEvent200ResponseOneOf2ChallengeChallenger.md) |  | 
**DestUser** | [**NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser**](ApiStreamEvent200ResponseOneOf2ChallengeDestUser.md) |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**Rated** | **bool** |  | 
**Speed** | **string** |  | 
**TimeControl** | [**ApiStreamEvent200ResponseOneOf2ChallengeTimeControl**](ApiStreamEvent200ResponseOneOf2ChallengeTimeControl.md) |  | 
**Color** | **string** |  | 
**FinalColor** | Pointer to **NullableString** |  | [optional] 
**Perf** | [**ApiStreamEvent200ResponseOneOf2ChallengePerf**](ApiStreamEvent200ResponseOneOf2ChallengePerf.md) |  | 
**Direction** | Pointer to **string** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**RematchOf** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOf2Challenge

`func NewApiStreamEvent200ResponseOneOf2Challenge(id string, url string, status string, challenger ApiStreamEvent200ResponseOneOf2ChallengeChallenger, destUser NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, rated bool, speed string, timeControl ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, color string, perf ApiStreamEvent200ResponseOneOf2ChallengePerf, ) *ApiStreamEvent200ResponseOneOf2Challenge`

NewApiStreamEvent200ResponseOneOf2Challenge instantiates a new ApiStreamEvent200ResponseOneOf2Challenge object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOf2ChallengeWithDefaults

`func NewApiStreamEvent200ResponseOneOf2ChallengeWithDefaults() *ApiStreamEvent200ResponseOneOf2Challenge`

NewApiStreamEvent200ResponseOneOf2ChallengeWithDefaults instantiates a new ApiStreamEvent200ResponseOneOf2Challenge object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetChallenger() ApiStreamEvent200ResponseOneOf2ChallengeChallenger`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetChallengerOk() (*ApiStreamEvent200ResponseOneOf2ChallengeChallenger, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetChallenger(v ApiStreamEvent200ResponseOneOf2ChallengeChallenger)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetDestUser() ApiStreamEvent200ResponseOneOf2ChallengeDestUser`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetDestUserOk() (*ApiStreamEvent200ResponseOneOf2ChallengeDestUser, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetDestUser(v ApiStreamEvent200ResponseOneOf2ChallengeDestUser)`

SetDestUser sets DestUser field to given value.


### SetDestUserNil

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetDestUserNil(b bool)`

 SetDestUserNil sets the value for DestUser to be an explicit nil

### UnsetDestUser
`func (o *ApiStreamEvent200ResponseOneOf2Challenge) UnsetDestUser()`

UnsetDestUser ensures that no value is present for DestUser, not even an explicit nil
### GetVariant

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetTimeControl() ApiStreamEvent200ResponseOneOf2ChallengeTimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetTimeControlOk() (*ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetTimeControl(v ApiStreamEvent200ResponseOneOf2ChallengeTimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetColor(v string)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetFinalColor() string`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetFinalColorOk() (*string, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetFinalColor(v string)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### SetFinalColorNil

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetFinalColorNil(b bool)`

 SetFinalColorNil sets the value for FinalColor to be an explicit nil

### UnsetFinalColor
`func (o *ApiStreamEvent200ResponseOneOf2Challenge) UnsetFinalColor()`

UnsetFinalColor ensures that no value is present for FinalColor, not even an explicit nil
### GetPerf

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetPerf() ApiStreamEvent200ResponseOneOf2ChallengePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetPerfOk() (*ApiStreamEvent200ResponseOneOf2ChallengePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetPerf(v ApiStreamEvent200ResponseOneOf2ChallengePerf)`

SetPerf sets Perf field to given value.


### GetDirection

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetInitialFen

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetRematchOf

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetRematchOf() string`

GetRematchOf returns the RematchOf field if non-nil, zero value otherwise.

### GetRematchOfOk

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) GetRematchOfOk() (*string, bool)`

GetRematchOfOk returns a tuple with the RematchOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRematchOf

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) SetRematchOf(v string)`

SetRematchOf sets RematchOf field to given value.

### HasRematchOf

`func (o *ApiStreamEvent200ResponseOneOf2Challenge) HasRematchOf() bool`

HasRematchOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


