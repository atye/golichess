# ChallengeDeclinedEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Challenge** | [**ChallengeDeclinedJson**](ChallengeDeclinedJson.md) |  | 

## Methods

### NewChallengeDeclinedEvent

`func NewChallengeDeclinedEvent(type_ string, challenge ChallengeDeclinedJson, ) *ChallengeDeclinedEvent`

NewChallengeDeclinedEvent instantiates a new ChallengeDeclinedEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeDeclinedEventWithDefaults

`func NewChallengeDeclinedEventWithDefaults() *ChallengeDeclinedEvent`

NewChallengeDeclinedEventWithDefaults instantiates a new ChallengeDeclinedEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ChallengeDeclinedEvent) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ChallengeDeclinedEvent) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ChallengeDeclinedEvent) SetType(v string)`

SetType sets Type field to given value.


### GetChallenge

`func (o *ChallengeDeclinedEvent) GetChallenge() ChallengeDeclinedJson`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ChallengeDeclinedEvent) GetChallengeOk() (*ChallengeDeclinedJson, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ChallengeDeclinedEvent) SetChallenge(v ChallengeDeclinedJson)`

SetChallenge sets Challenge field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


