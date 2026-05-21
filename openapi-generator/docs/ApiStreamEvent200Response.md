# ApiStreamEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Game** | [**GameEventInfo**](GameEventInfo.md) |  | 
**Challenge** | [**ChallengeDeclinedJson**](ChallengeDeclinedJson.md) |  | 
**Compat** | Pointer to [**GameCompat**](GameCompat.md) |  | [optional] 

## Methods

### NewApiStreamEvent200Response

`func NewApiStreamEvent200Response(type_ string, game GameEventInfo, challenge ChallengeDeclinedJson, ) *ApiStreamEvent200Response`

NewApiStreamEvent200Response instantiates a new ApiStreamEvent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseWithDefaults

`func NewApiStreamEvent200ResponseWithDefaults() *ApiStreamEvent200Response`

NewApiStreamEvent200ResponseWithDefaults instantiates a new ApiStreamEvent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiStreamEvent200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStreamEvent200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStreamEvent200Response) SetType(v string)`

SetType sets Type field to given value.


### GetGame

`func (o *ApiStreamEvent200Response) GetGame() GameEventInfo`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *ApiStreamEvent200Response) GetGameOk() (*GameEventInfo, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *ApiStreamEvent200Response) SetGame(v GameEventInfo)`

SetGame sets Game field to given value.


### GetChallenge

`func (o *ApiStreamEvent200Response) GetChallenge() ChallengeDeclinedJson`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ApiStreamEvent200Response) GetChallengeOk() (*ChallengeDeclinedJson, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ApiStreamEvent200Response) SetChallenge(v ChallengeDeclinedJson)`

SetChallenge sets Challenge field to given value.


### GetCompat

`func (o *ApiStreamEvent200Response) GetCompat() GameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiStreamEvent200Response) GetCompatOk() (*GameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiStreamEvent200Response) SetCompat(v GameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ApiStreamEvent200Response) HasCompat() bool`

HasCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


