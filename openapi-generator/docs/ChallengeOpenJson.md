# ChallengeOpenJson

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | [**ChallengeStatus**](ChallengeStatus.md) |  | 
**Challenger** | [**nil**](nil.md) |  | 
**DestUser** | [**nil**](nil.md) |  | 
**Variant** | [**Variant**](Variant.md) |  | 
**Rated** | **bool** |  | 
**Speed** | [**Speed**](Speed.md) |  | 
**TimeControl** | [**TimeControl**](TimeControl.md) |  | 
**Color** | [**ChallengeColor**](ChallengeColor.md) |  | 
**FinalColor** | Pointer to [**GameColor**](GameColor.md) |  | [optional] 
**Perf** | [**ChallengeOpenJsonPerf**](ChallengeOpenJsonPerf.md) |  | 
**InitialFen** | Pointer to **string** |  | [optional] 
**UrlWhite** | **string** |  | 
**UrlBlack** | **string** |  | 
**Open** | [**ChallengeOpenJsonOpen**](ChallengeOpenJsonOpen.md) |  | 

## Methods

### NewChallengeOpenJson

`func NewChallengeOpenJson(id string, url string, status ChallengeStatus, challenger nil, destUser nil, variant Variant, rated bool, speed Speed, timeControl TimeControl, color ChallengeColor, perf ChallengeOpenJsonPerf, urlWhite string, urlBlack string, open ChallengeOpenJsonOpen, ) *ChallengeOpenJson`

NewChallengeOpenJson instantiates a new ChallengeOpenJson object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeOpenJsonWithDefaults

`func NewChallengeOpenJsonWithDefaults() *ChallengeOpenJson`

NewChallengeOpenJsonWithDefaults instantiates a new ChallengeOpenJson object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeOpenJson) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeOpenJson) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeOpenJson) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChallengeOpenJson) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeOpenJson) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeOpenJson) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ChallengeOpenJson) GetStatus() ChallengeStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeOpenJson) GetStatusOk() (*ChallengeStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeOpenJson) SetStatus(v ChallengeStatus)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ChallengeOpenJson) GetChallenger() nil`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeOpenJson) GetChallengerOk() (*nil, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeOpenJson) SetChallenger(v nil)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ChallengeOpenJson) GetDestUser() nil`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeOpenJson) GetDestUserOk() (*nil, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeOpenJson) SetDestUser(v nil)`

SetDestUser sets DestUser field to given value.


### GetVariant

`func (o *ChallengeOpenJson) GetVariant() Variant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeOpenJson) GetVariantOk() (*Variant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeOpenJson) SetVariant(v Variant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeOpenJson) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeOpenJson) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeOpenJson) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ChallengeOpenJson) GetSpeed() Speed`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeOpenJson) GetSpeedOk() (*Speed, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeOpenJson) SetSpeed(v Speed)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeOpenJson) GetTimeControl() TimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeOpenJson) GetTimeControlOk() (*TimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeOpenJson) SetTimeControl(v TimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ChallengeOpenJson) GetColor() ChallengeColor`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeOpenJson) GetColorOk() (*ChallengeColor, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeOpenJson) SetColor(v ChallengeColor)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ChallengeOpenJson) GetFinalColor() GameColor`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ChallengeOpenJson) GetFinalColorOk() (*GameColor, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ChallengeOpenJson) SetFinalColor(v GameColor)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ChallengeOpenJson) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeOpenJson) GetPerf() ChallengeOpenJsonPerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeOpenJson) GetPerfOk() (*ChallengeOpenJsonPerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeOpenJson) SetPerf(v ChallengeOpenJsonPerf)`

SetPerf sets Perf field to given value.


### GetInitialFen

`func (o *ChallengeOpenJson) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeOpenJson) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeOpenJson) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeOpenJson) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetUrlWhite

`func (o *ChallengeOpenJson) GetUrlWhite() string`

GetUrlWhite returns the UrlWhite field if non-nil, zero value otherwise.

### GetUrlWhiteOk

`func (o *ChallengeOpenJson) GetUrlWhiteOk() (*string, bool)`

GetUrlWhiteOk returns a tuple with the UrlWhite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWhite

`func (o *ChallengeOpenJson) SetUrlWhite(v string)`

SetUrlWhite sets UrlWhite field to given value.


### GetUrlBlack

`func (o *ChallengeOpenJson) GetUrlBlack() string`

GetUrlBlack returns the UrlBlack field if non-nil, zero value otherwise.

### GetUrlBlackOk

`func (o *ChallengeOpenJson) GetUrlBlackOk() (*string, bool)`

GetUrlBlackOk returns a tuple with the UrlBlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBlack

`func (o *ChallengeOpenJson) SetUrlBlack(v string)`

SetUrlBlack sets UrlBlack field to given value.


### GetOpen

`func (o *ChallengeOpenJson) GetOpen() ChallengeOpenJsonOpen`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *ChallengeOpenJson) GetOpenOk() (*ChallengeOpenJsonOpen, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *ChallengeOpenJson) SetOpen(v ChallengeOpenJsonOpen)`

SetOpen sets Open field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


