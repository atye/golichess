# ApiStreamEvent200ResponseOneOf2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Challenge** | [**ApiStreamEvent200ResponseOneOf2Challenge**](ApiStreamEvent200ResponseOneOf2Challenge.md) |  | 
**Compat** | Pointer to [**ApiStreamEvent200ResponseOneOfGameCompat**](ApiStreamEvent200ResponseOneOfGameCompat.md) |  | [optional] 

## Methods

### NewApiStreamEvent200ResponseOneOf2

`func NewApiStreamEvent200ResponseOneOf2(type_ string, challenge ApiStreamEvent200ResponseOneOf2Challenge, ) *ApiStreamEvent200ResponseOneOf2`

NewApiStreamEvent200ResponseOneOf2 instantiates a new ApiStreamEvent200ResponseOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStreamEvent200ResponseOneOf2WithDefaults

`func NewApiStreamEvent200ResponseOneOf2WithDefaults() *ApiStreamEvent200ResponseOneOf2`

NewApiStreamEvent200ResponseOneOf2WithDefaults instantiates a new ApiStreamEvent200ResponseOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ApiStreamEvent200ResponseOneOf2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiStreamEvent200ResponseOneOf2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiStreamEvent200ResponseOneOf2) SetType(v string)`

SetType sets Type field to given value.


### GetChallenge

`func (o *ApiStreamEvent200ResponseOneOf2) GetChallenge() ApiStreamEvent200ResponseOneOf2Challenge`

GetChallenge returns the Challenge field if non-nil, zero value otherwise.

### GetChallengeOk

`func (o *ApiStreamEvent200ResponseOneOf2) GetChallengeOk() (*ApiStreamEvent200ResponseOneOf2Challenge, bool)`

GetChallengeOk returns a tuple with the Challenge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChallenge

`func (o *ApiStreamEvent200ResponseOneOf2) SetChallenge(v ApiStreamEvent200ResponseOneOf2Challenge)`

SetChallenge sets Challenge field to given value.


### GetCompat

`func (o *ApiStreamEvent200ResponseOneOf2) GetCompat() ApiStreamEvent200ResponseOneOfGameCompat`

GetCompat returns the Compat field if non-nil, zero value otherwise.

### GetCompatOk

`func (o *ApiStreamEvent200ResponseOneOf2) GetCompatOk() (*ApiStreamEvent200ResponseOneOfGameCompat, bool)`

GetCompatOk returns a tuple with the Compat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompat

`func (o *ApiStreamEvent200ResponseOneOf2) SetCompat(v ApiStreamEvent200ResponseOneOfGameCompat)`

SetCompat sets Compat field to given value.

### HasCompat

`func (o *ApiStreamEvent200ResponseOneOf2) HasCompat() bool`

HasCompat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


