# ChallengeList200ResponseInInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | **string** |  | 
**Challenger** | [**NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser**](ApiStreamEvent200ResponseOneOf2ChallengeDestUser.md) |  | 
**DestUser** | [**NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser**](ApiStreamEvent200ResponseOneOf2ChallengeDestUser.md) |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**Rated** | **bool** |  | 
**Speed** | **string** |  | 
**TimeControl** | [**ApiStreamEvent200ResponseOneOf2ChallengeTimeControl**](ApiStreamEvent200ResponseOneOf2ChallengeTimeControl.md) |  | 
**Color** | **string** | The color to play. Better left empty to automatically get 50% white. | 
**FinalColor** | Pointer to **string** | Color of the winner, if any | [optional] 
**Perf** | [**ApiStreamEvent200ResponseOneOf2ChallengePerf**](ApiStreamEvent200ResponseOneOf2ChallengePerf.md) |  | 
**Direction** | Pointer to **string** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**RematchOf** | Pointer to **string** |  | [optional] 

## Methods

### NewChallengeList200ResponseInInner

`func NewChallengeList200ResponseInInner(id string, url string, status string, challenger NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser, destUser NullableApiStreamEvent200ResponseOneOf2ChallengeDestUser, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, rated bool, speed string, timeControl ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, color string, perf ApiStreamEvent200ResponseOneOf2ChallengePerf, ) *ChallengeList200ResponseInInner`

NewChallengeList200ResponseInInner instantiates a new ChallengeList200ResponseInInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeList200ResponseInInnerWithDefaults

`func NewChallengeList200ResponseInInnerWithDefaults() *ChallengeList200ResponseInInner`

NewChallengeList200ResponseInInnerWithDefaults instantiates a new ChallengeList200ResponseInInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeList200ResponseInInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeList200ResponseInInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeList200ResponseInInner) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChallengeList200ResponseInInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeList200ResponseInInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeList200ResponseInInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ChallengeList200ResponseInInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeList200ResponseInInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeList200ResponseInInner) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ChallengeList200ResponseInInner) GetChallenger() ApiStreamEvent200ResponseOneOf2ChallengeDestUser`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeList200ResponseInInner) GetChallengerOk() (*ApiStreamEvent200ResponseOneOf2ChallengeDestUser, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeList200ResponseInInner) SetChallenger(v ApiStreamEvent200ResponseOneOf2ChallengeDestUser)`

SetChallenger sets Challenger field to given value.


### SetChallengerNil

`func (o *ChallengeList200ResponseInInner) SetChallengerNil(b bool)`

 SetChallengerNil sets the value for Challenger to be an explicit nil

### UnsetChallenger
`func (o *ChallengeList200ResponseInInner) UnsetChallenger()`

UnsetChallenger ensures that no value is present for Challenger, not even an explicit nil
### GetDestUser

`func (o *ChallengeList200ResponseInInner) GetDestUser() ApiStreamEvent200ResponseOneOf2ChallengeDestUser`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeList200ResponseInInner) GetDestUserOk() (*ApiStreamEvent200ResponseOneOf2ChallengeDestUser, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeList200ResponseInInner) SetDestUser(v ApiStreamEvent200ResponseOneOf2ChallengeDestUser)`

SetDestUser sets DestUser field to given value.


### SetDestUserNil

`func (o *ChallengeList200ResponseInInner) SetDestUserNil(b bool)`

 SetDestUserNil sets the value for DestUser to be an explicit nil

### UnsetDestUser
`func (o *ChallengeList200ResponseInInner) UnsetDestUser()`

UnsetDestUser ensures that no value is present for DestUser, not even an explicit nil
### GetVariant

`func (o *ChallengeList200ResponseInInner) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeList200ResponseInInner) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeList200ResponseInInner) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeList200ResponseInInner) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeList200ResponseInInner) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeList200ResponseInInner) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ChallengeList200ResponseInInner) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeList200ResponseInInner) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeList200ResponseInInner) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeList200ResponseInInner) GetTimeControl() ApiStreamEvent200ResponseOneOf2ChallengeTimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeList200ResponseInInner) GetTimeControlOk() (*ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeList200ResponseInInner) SetTimeControl(v ApiStreamEvent200ResponseOneOf2ChallengeTimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ChallengeList200ResponseInInner) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeList200ResponseInInner) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeList200ResponseInInner) SetColor(v string)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ChallengeList200ResponseInInner) GetFinalColor() string`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ChallengeList200ResponseInInner) GetFinalColorOk() (*string, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ChallengeList200ResponseInInner) SetFinalColor(v string)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ChallengeList200ResponseInInner) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeList200ResponseInInner) GetPerf() ApiStreamEvent200ResponseOneOf2ChallengePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeList200ResponseInInner) GetPerfOk() (*ApiStreamEvent200ResponseOneOf2ChallengePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeList200ResponseInInner) SetPerf(v ApiStreamEvent200ResponseOneOf2ChallengePerf)`

SetPerf sets Perf field to given value.


### GetDirection

`func (o *ChallengeList200ResponseInInner) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ChallengeList200ResponseInInner) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ChallengeList200ResponseInInner) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ChallengeList200ResponseInInner) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetInitialFen

`func (o *ChallengeList200ResponseInInner) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeList200ResponseInInner) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeList200ResponseInInner) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeList200ResponseInInner) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetRematchOf

`func (o *ChallengeList200ResponseInInner) GetRematchOf() string`

GetRematchOf returns the RematchOf field if non-nil, zero value otherwise.

### GetRematchOfOk

`func (o *ChallengeList200ResponseInInner) GetRematchOfOk() (*string, bool)`

GetRematchOfOk returns a tuple with the RematchOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRematchOf

`func (o *ChallengeList200ResponseInInner) SetRematchOf(v string)`

SetRematchOf sets RematchOf field to given value.

### HasRematchOf

`func (o *ChallengeList200ResponseInInner) HasRematchOf() bool`

HasRematchOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


