# ChallengeJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | [**ChallengeStatus**](ChallengeStatus.md) |  | 
**Challenger** | [**ChallengeUser**](ChallengeUser.md) |  | 
**DestUser** | [**NullableChallengeUser**](ChallengeUser.md) |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**Rated** | **bool** |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**TimeControl** | [**TimeControl**](TimeControl.md) |  | 
**Color** | [**ChallengeColor**](ChallengeColor.md) |  | 
**FinalColor** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**Perf** | [**ChallengeJsonPerf**](ChallengeJsonPerf.md) |  | 
**Direction** | Pointer to **string** |  | [optional] 
**InitialFen** | Pointer to **string** |  | [optional] 
**RematchOf** | Pointer to **string** |  | [optional] 

## Methods

### NewChallengeJson

`func NewChallengeJson(id string, url string, status ChallengeStatus, challenger ChallengeUser, destUser NullableChallengeUser, variant Variant, rated bool, speed Speed, timeControl TimeControl, color ChallengeColor, perf ChallengeJsonPerf, ) *ChallengeJson`

NewChallengeJson instantiates a new ChallengeJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeJsonWithDefaults

`func NewChallengeJsonWithDefaults() *ChallengeJson`

NewChallengeJsonWithDefaults instantiates a new ChallengeJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeJson) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChallengeJson) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeJson) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeJson) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ChallengeJson) GetStatus() ChallengeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeJson) GetStatusOk() (*ChallengeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeJson) SetStatus(v ChallengeStatus)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ChallengeJson) GetChallenger() ChallengeUser`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeJson) GetChallengerOk() (*ChallengeUser, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeJson) SetChallenger(v ChallengeUser)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ChallengeJson) GetDestUser() ChallengeUser`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeJson) GetDestUserOk() (*ChallengeUser, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeJson) SetDestUser(v ChallengeUser)`

SetDestUser sets DestUser field to given value.


### SetDestUserNil

`func (o *ChallengeJson) SetDestUserNil(b bool)`

 SetDestUserNil sets the value for DestUser to be an explicit nil

### UnsetDestUser
`func (o *ChallengeJson) UnsetDestUser()`

UnsetDestUser ensures that no value is present for DestUser, not even an explicit nil
### GetVariant

`func (o *ChallengeJson) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeJson) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeJson) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeJson) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeJson) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeJson) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ChallengeJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeJson) GetTimeControl() TimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeJson) GetTimeControlOk() (*TimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeJson) SetTimeControl(v TimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ChallengeJson) GetColor() ChallengeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeJson) GetColorOk() (*ChallengeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeJson) SetColor(v ChallengeColor)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ChallengeJson) GetFinalColor() GameColor`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ChallengeJson) GetFinalColorOk() (*GameColor, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ChallengeJson) SetFinalColor(v GameColor)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ChallengeJson) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeJson) GetPerf() ChallengeJsonPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeJson) GetPerfOk() (*ChallengeJsonPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeJson) SetPerf(v ChallengeJsonPerf)`

SetPerf sets Perf field to given value.


### GetDirection

`func (o *ChallengeJson) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ChallengeJson) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ChallengeJson) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ChallengeJson) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetInitialFen

`func (o *ChallengeJson) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeJson) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeJson) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetRematchOf

`func (o *ChallengeJson) GetRematchOf() string`

GetRematchOf returns the RematchOf field if non-nil, zero value otherwise.

### GetRematchOfOk

`func (o *ChallengeJson) GetRematchOfOk() (*string, bool)`

GetRematchOfOk returns a tuple with the RematchOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRematchOf

`func (o *ChallengeJson) SetRematchOf(v string)`

SetRematchOf sets RematchOf field to given value.

### HasRematchOf

`func (o *ChallengeJson) HasRematchOf() bool`

HasRematchOf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


