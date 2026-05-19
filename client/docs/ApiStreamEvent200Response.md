# ApiStreamEvent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Game** | [**ApiStreamEvent200ResponseOneOf1Game**](ApiStreamEvent200ResponseOneOf1Game.md) |  | 
**Challenge** | [**ApiStreamEvent200ResponseOneOf4Challenge**](ApiStreamEvent200ResponseOneOf4Challenge.md) |  | 
**Compat** | Pointer to [**ApiStreamEvent200ResponseOneOfGameCompat**](ApiStreamEvent200ResponseOneOfGameCompat.md) |  | [optional] 

## Methods

### NewApiStreamEvent200Response

`func NewApiStreamEvent200Response(type_ string, game ApiStreamEvent200ResponseOneOf1Game, challenge ApiStreamEvent200ResponseOneOf4Challenge, ) *ApiStreamEvent200Response`

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

`func (o *ApiStreamEvent200Response) GetGame() ApiStreamEvent200ResponseOneOf1Game`

GetGame returns the Game field if non-nil, zero value otherwise.

### GetGameOk

`func (o *ApiStreamEvent200Response) GetGameOk() (*ApiStreamEvent200ResponseOneOf1Game, bool)`

GetGameOk returns a tuple with the Game field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGame

`func (o *ApiStreamEvent200Response) SetGame(v ApiStreamEvent200ResponseOneOf1Game)`

SetGame sets Game field to given value.


### GetChallenge

`func (o *ApiStreamEvent200Response) GetChallenge() ApiStreamEvent200ResponseOneOf4Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ApiStreamEvent200Response) GetChallengeOk() (*ApiStreamEvent200ResponseOneOf4Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ApiStreamEvent200Response) SetChallenge(v ApiStreamEvent200ResponseOneOf4Challenge)`

SetChallenge sets Challenge field to given value.


### GetCompat

`func (o *ApiStreamEvent200Response) GetCompat() ApiStreamEvent200ResponseOneOfGameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiStreamEvent200Response) GetCompatOk() (*ApiStreamEvent200ResponseOneOfGameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiStreamEvent200Response) SetCompat(v ApiStreamEvent200ResponseOneOfGameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ApiStreamEvent200Response) HasCompat() bool`

HasCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


