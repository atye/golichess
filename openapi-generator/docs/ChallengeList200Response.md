# ChallengeList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to [**[]ChallengeJson**](ChallengeJson.md) | Incoming challenges i.e. targeted at you | [optional] 
**Out** | Pointer to [**[]ChallengeJson**](ChallengeJson.md) | Outgoing challenges i.e. created by you | [optional] 

## Methods

### NewChallengeList200Response

`func NewChallengeList200Response() *ChallengeList200Response`

NewChallengeList200Response instantiates a new ChallengeList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChallengeList200ResponseWithDefaults

`func NewChallengeList200ResponseWithDefaults() *ChallengeList200Response`

NewChallengeList200ResponseWithDefaults instantiates a new ChallengeList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *ChallengeList200Response) GetIn() []ChallengeJson`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *ChallengeList200Response) GetInOk() (*[]ChallengeJson, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *ChallengeList200Response) SetIn(v []ChallengeJson)`

SetIn sets In field to given value.

### HasIn

`func (o *ChallengeList200Response) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetOut

`func (o *ChallengeList200Response) GetOut() []ChallengeJson`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *ChallengeList200Response) GetOutOk() (*[]ChallengeJson, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *ChallengeList200Response) SetOut(v []ChallengeJson)`

SetOut sets Out field to given value.

### HasOut

`func (o *ChallengeList200Response) HasOut() bool`

HasOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


