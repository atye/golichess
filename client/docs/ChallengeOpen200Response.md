# ChallengeOpen200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Status** | **string** |  | 
**Challenger** | [**nil**](nil.md) |  | 
**DestUser** | [**nil**](nil.md) |  | 
**Variant** | [**ApiAccountPlaying200ResponseNowPlayingInnerVariant**](ApiAccountPlaying200ResponseNowPlayingInnerVariant.md) |  | 
**Rated** | **bool** |  | 
**Speed** | **string** |  | 
**TimeControl** | [**ApiStreamEvent200ResponseOneOf2ChallengeTimeControl**](ApiStreamEvent200ResponseOneOf2ChallengeTimeControl.md) |  | 
**Color** | **string** | Which color you get to play | [default to "random"]
**FinalColor** | Pointer to **string** | Color of the winner, if any | [optional] 
**Perf** | [**ChallengeOpen200ResponsePerf**](ChallengeOpen200ResponsePerf.md) |  | 
**InitialFen** | Pointer to **string** |  | [optional] 
**UrlWhite** | **string** |  | 
**UrlBlack** | **string** |  | 
**Open** | [**ChallengeOpen200ResponseOpen**](ChallengeOpen200ResponseOpen.md) |  | 

## Methods

### NewChallengeOpen200Response

`func NewChallengeOpen200Response(id string, url string, status string, challenger nil, destUser nil, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, rated bool, speed string, timeControl ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, color string, perf ChallengeOpen200ResponsePerf, urlWhite string, urlBlack string, open ChallengeOpen200ResponseOpen, ) *ChallengeOpen200Response`

NewChallengeOpen200Response instantiates a new ChallengeOpen200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeOpen200ResponseWithDefaults

`func NewChallengeOpen200ResponseWithDefaults() *ChallengeOpen200Response`

NewChallengeOpen200ResponseWithDefaults instantiates a new ChallengeOpen200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChallengeOpen200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChallengeOpen200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChallengeOpen200Response) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ChallengeOpen200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ChallengeOpen200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ChallengeOpen200Response) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetStatus

`func (o *ChallengeOpen200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ChallengeOpen200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ChallengeOpen200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetChallenger

`func (o *ChallengeOpen200Response) GetChallenger() nil`

GetChallenger returns the Challenger field if non-nil, zero value otherwise.

### GetChallengerOk

`func (o *ChallengeOpen200Response) GetChallengerOk() (*nil, bool)`

GetChallengerOk returns a tuple with the Challenger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenger

`func (o *ChallengeOpen200Response) SetChallenger(v nil)`

SetChallenger sets Challenger field to given value.


### GetDestUser

`func (o *ChallengeOpen200Response) GetDestUser() nil`

GetDestUser returns the DestUser field if non-nil, zero value otherwise.

### GetDestUserOk

`func (o *ChallengeOpen200Response) GetDestUserOk() (*nil, bool)`

GetDestUserOk returns a tuple with the DestUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestUser

`func (o *ChallengeOpen200Response) SetDestUser(v nil)`

SetDestUser sets DestUser field to given value.


### GetVariant

`func (o *ChallengeOpen200Response) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant`

GetVariant returns the Variant field if non-nil, zero value otherwise.

### GetVariantOk

`func (o *ChallengeOpen200Response) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool)`

GetVariantOk returns a tuple with the Variant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariant

`func (o *ChallengeOpen200Response) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant)`

SetVariant sets Variant field to given value.


### GetRated

`func (o *ChallengeOpen200Response) GetRated() bool`

GetRated returns the Rated field if non-nil, zero value otherwise.

### GetRatedOk

`func (o *ChallengeOpen200Response) GetRatedOk() (*bool, bool)`

GetRatedOk returns a tuple with the Rated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRated

`func (o *ChallengeOpen200Response) SetRated(v bool)`

SetRated sets Rated field to given value.


### GetSpeed

`func (o *ChallengeOpen200Response) GetSpeed() string`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *ChallengeOpen200Response) GetSpeedOk() (*string, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *ChallengeOpen200Response) SetSpeed(v string)`

SetSpeed sets Speed field to given value.


### GetTimeControl

`func (o *ChallengeOpen200Response) GetTimeControl() ApiStreamEvent200ResponseOneOf2ChallengeTimeControl`

GetTimeControl returns the TimeControl field if non-nil, zero value otherwise.

### GetTimeControlOk

`func (o *ChallengeOpen200Response) GetTimeControlOk() (*ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, bool)`

GetTimeControlOk returns a tuple with the TimeControl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeControl

`func (o *ChallengeOpen200Response) SetTimeControl(v ApiStreamEvent200ResponseOneOf2ChallengeTimeControl)`

SetTimeControl sets TimeControl field to given value.


### GetColor

`func (o *ChallengeOpen200Response) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ChallengeOpen200Response) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ChallengeOpen200Response) SetColor(v string)`

SetColor sets Color field to given value.


### GetFinalColor

`func (o *ChallengeOpen200Response) GetFinalColor() string`

GetFinalColor returns the FinalColor field if non-nil, zero value otherwise.

### GetFinalColorOk

`func (o *ChallengeOpen200Response) GetFinalColorOk() (*string, bool)`

GetFinalColorOk returns a tuple with the FinalColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalColor

`func (o *ChallengeOpen200Response) SetFinalColor(v string)`

SetFinalColor sets FinalColor field to given value.

### HasFinalColor

`func (o *ChallengeOpen200Response) HasFinalColor() bool`

HasFinalColor returns a boolean if a field has been set.

### GetPerf

`func (o *ChallengeOpen200Response) GetPerf() ChallengeOpen200ResponsePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ChallengeOpen200Response) GetPerfOk() (*ChallengeOpen200ResponsePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ChallengeOpen200Response) SetPerf(v ChallengeOpen200ResponsePerf)`

SetPerf sets Perf field to given value.


### GetInitialFen

`func (o *ChallengeOpen200Response) GetInitialFen() string`

GetInitialFen returns the InitialFen field if non-nil, zero value otherwise.

### GetInitialFenOk

`func (o *ChallengeOpen200Response) GetInitialFenOk() (*string, bool)`

GetInitialFenOk returns a tuple with the InitialFen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialFen

`func (o *ChallengeOpen200Response) SetInitialFen(v string)`

SetInitialFen sets InitialFen field to given value.

### HasInitialFen

`func (o *ChallengeOpen200Response) HasInitialFen() bool`

HasInitialFen returns a boolean if a field has been set.

### GetUrlWhite

`func (o *ChallengeOpen200Response) GetUrlWhite() string`

GetUrlWhite returns the UrlWhite field if non-nil, zero value otherwise.

### GetUrlWhiteOk

`func (o *ChallengeOpen200Response) GetUrlWhiteOk() (*string, bool)`

GetUrlWhiteOk returns a tuple with the UrlWhite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlWhite

`func (o *ChallengeOpen200Response) SetUrlWhite(v string)`

SetUrlWhite sets UrlWhite field to given value.


### GetUrlBlack

`func (o *ChallengeOpen200Response) GetUrlBlack() string`

GetUrlBlack returns the UrlBlack field if non-nil, zero value otherwise.

### GetUrlBlackOk

`func (o *ChallengeOpen200Response) GetUrlBlackOk() (*string, bool)`

GetUrlBlackOk returns a tuple with the UrlBlack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBlack

`func (o *ChallengeOpen200Response) SetUrlBlack(v string)`

SetUrlBlack sets UrlBlack field to given value.


### GetOpen

`func (o *ChallengeOpen200Response) GetOpen() ChallengeOpen200ResponseOpen`

GetOpen returns the Open field if non-nil, zero value otherwise.

### GetOpenOk

`func (o *ChallengeOpen200Response) GetOpenOk() (*ChallengeOpen200ResponseOpen, bool)`

GetOpenOk returns a tuple with the Open field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpen

`func (o *ChallengeOpen200Response) SetOpen(v ChallengeOpen200ResponseOpen)`

SetOpen sets Open field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


