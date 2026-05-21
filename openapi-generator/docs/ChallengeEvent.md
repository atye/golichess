# ChallengeEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Challenge** | [**ChallengeJson**](ChallengeJson.md) |  | 
**Compat** | Pointer to [**GameCompat**](GameCompat.md) |  | [optional] 

## Methods

### NewChallengeEvent

`func NewChallengeEvent(type_ string, challenge ChallengeJson, ) *ChallengeEvent`

NewChallengeEvent instantiates a new ChallengeEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeEventWithDefaults

`func NewChallengeEventWithDefaults() *ChallengeEvent`

NewChallengeEventWithDefaults instantiates a new ChallengeEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeEvent) SetType(v string)`

SetType sets Type field to given value.


### GetChallenge

`func (o *ChallengeEvent) GetChallenge() ChallengeJson`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ChallengeEvent) GetChallengeOk() (*ChallengeJson, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ChallengeEvent) SetChallenge(v ChallengeJson)`

SetChallenge sets Challenge field to given value.


### GetCompat

`func (o *ChallengeEvent) GetCompat() GameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ChallengeEvent) GetCompatOk() (*GameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ChallengeEvent) SetCompat(v GameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ChallengeEvent) HasCompat() bool`

HasCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


