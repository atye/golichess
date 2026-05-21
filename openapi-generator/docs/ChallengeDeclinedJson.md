# ChallengeDeclinedJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | [**ChallengeStatus**](ChallengeStatus.md) |  | 
**Challenger** | [**ChallengeUser**](ChallengeUser.md) |  | 
**DestUser** | [**ChallengeUser**](ChallengeUser.md) |  | 
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
**DeclineReason** | **string** | Human readable, possibly translated reason why the challenge was declined. | 
**DeclineReasonKey** | **string** | Untranslated, computer-matchable reason why the challenge was declined. | 

## Methods

### NewChallengeDeclinedJson

`func NewChallengeDeclinedJson(id string, url string, status ChallengeStatus, challenger ChallengeUser, destUser ChallengeUser, variant Variant, rated bool, speed Speed, timeControl TimeControl, color ChallengeColor, perf ChallengeJsonPerf, declineReason string, declineReasonKey string, ) *ChallengeDeclinedJson`

NewChallengeDeclinedJson instantiates a new ChallengeDeclinedJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeDeclinedJsonWithDefaults

`func NewChallengeDeclinedJsonWithDefaults() *ChallengeDeclinedJson`

NewChallengeDeclinedJsonWithDefaults instantiates a new ChallengeDeclinedJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeDeclinedJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeDeclinedJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeDeclinedJson) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChallengeDeclinedJson) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeDeclinedJson) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeDeclinedJson) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ChallengeDeclinedJson) GetStatus() ChallengeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeDeclinedJson) GetStatusOk() (*ChallengeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeDeclinedJson) SetStatus(v ChallengeStatus)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ChallengeDeclinedJson) GetChallenger() ChallengeUser`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeDeclinedJson) GetChallengerOk() (*ChallengeUser, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeDeclinedJson) SetChallenger(v ChallengeUser)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ChallengeDeclinedJson) GetDestUser() ChallengeUser`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeDeclinedJson) GetDestUserOk() (*ChallengeUser, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeDeclinedJson) SetDestUser(v ChallengeUser)`

SetDestUser sets DestUser field to given value.


### GetVariant

`func (o *ChallengeDeclinedJson) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeDeclinedJson) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeDeclinedJson) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeDeclinedJson) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeDeclinedJson) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeDeclinedJson) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ChallengeDeclinedJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeDeclinedJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeDeclinedJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeDeclinedJson) GetTimeControl() TimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeDeclinedJson) GetTimeControlOk() (*TimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeDeclinedJson) SetTimeControl(v TimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ChallengeDeclinedJson) GetColor() ChallengeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeDeclinedJson) GetColorOk() (*ChallengeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeDeclinedJson) SetColor(v ChallengeColor)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ChallengeDeclinedJson) GetFinalColor() GameColor`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ChallengeDeclinedJson) GetFinalColorOk() (*GameColor, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ChallengeDeclinedJson) SetFinalColor(v GameColor)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ChallengeDeclinedJson) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeDeclinedJson) GetPerf() ChallengeJsonPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeDeclinedJson) GetPerfOk() (*ChallengeJsonPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeDeclinedJson) SetPerf(v ChallengeJsonPerf)`

SetPerf sets Perf field to given value.


### GetDirection

`func (o *ChallengeDeclinedJson) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *ChallengeDeclinedJson) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *ChallengeDeclinedJson) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *ChallengeDeclinedJson) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetInitialFen

`func (o *ChallengeDeclinedJson) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeDeclinedJson) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeDeclinedJson) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeDeclinedJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetRematchOf

`func (o *ChallengeDeclinedJson) GetRematchOf() string`

GetRematchOf returns the RematchOf field if non-nil, zero value otherwise.

### GetRematchOfOk

`func (o *ChallengeDeclinedJson) GetRematchOfOk() (*string, bool)`

GetRematchOfOk returns a tuple with the RematchOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRematchOf

`func (o *ChallengeDeclinedJson) SetRematchOf(v string)`

SetRematchOf sets RematchOf field to given value.

### HasRematchOf

`func (o *ChallengeDeclinedJson) HasRematchOf() bool`

HasRematchOf returns a boolean if a field has been set.

### GetDeclineReason

`func (o *ChallengeDeclinedJson) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *ChallengeDeclinedJson) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *ChallengeDeclinedJson) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.


### GetDeclineReasonKey

`func (o *ChallengeDeclinedJson) GetDeclineReasonKey() string`

GetDeclineReasonKey returns the DeclineReasonKey field if non-nil, zero value otherwise.

### GetDeclineReasonKeyOk

`func (o *ChallengeDeclinedJson) GetDeclineReasonKeyOk() (*string, bool)`

GetDeclineReasonKeyOk returns a tuple with the DeclineReasonKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReasonKey

`func (o *ChallengeDeclinedJson) SetDeclineReasonKey(v string)`

SetDeclineReasonKey sets DeclineReasonKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


