# ApiUserPerf200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**ApiUserPerf200ResponseUser**](ApiUserPerf200ResponseUser.md) |  | 
**Perf** | [**ApiUserPerf200ResponsePerf**](ApiUserPerf200ResponsePerf.md) |  | 
**Rank** | **NullableInt32** |  | 
**Percentile** | **float32** |  | 
**Stat** | [**ApiUserPerf200ResponseStat**](ApiUserPerf200ResponseStat.md) |  | 

## Methods

### NewApiUserPerf200Response

`func NewApiUserPerf200Response(user ApiUserPerf200ResponseUser, perf ApiUserPerf200ResponsePerf, rank NullableInt32, percentile float32, stat ApiUserPerf200ResponseStat, ) *ApiUserPerf200Response`

NewApiUserPerf200Response instantiates a new ApiUserPerf200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiUserPerf200ResponseWithDefaults

`func NewApiUserPerf200ResponseWithDefaults() *ApiUserPerf200Response`

NewApiUserPerf200ResponseWithDefaults instantiates a new ApiUserPerf200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ApiUserPerf200Response) GetUser() ApiUserPerf200ResponseUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ApiUserPerf200Response) GetUserOk() (*ApiUserPerf200ResponseUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ApiUserPerf200Response) SetUser(v ApiUserPerf200ResponseUser)`

SetUser sets User field to given value.


### GetPerf

`func (o *ApiUserPerf200Response) GetPerf() ApiUserPerf200ResponsePerf`

GetPerf returns the Perf field if non-nil, zero value otherwise.

### GetPerfOk

`func (o *ApiUserPerf200Response) GetPerfOk() (*ApiUserPerf200ResponsePerf, bool)`

GetPerfOk returns a tuple with the Perf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerf

`func (o *ApiUserPerf200Response) SetPerf(v ApiUserPerf200ResponsePerf)`

SetPerf sets Perf field to given value.


### GetRank

`func (o *ApiUserPerf200Response) GetRank() int32`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *ApiUserPerf200Response) GetRankOk() (*int32, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *ApiUserPerf200Response) SetRank(v int32)`

SetRank sets Rank field to given value.


### SetRankNil

`func (o *ApiUserPerf200Response) SetRankNil(b bool)`

 SetRankNil sets the value for Rank to be an explicit nil

### UnsetRank
`func (o *ApiUserPerf200Response) UnsetRank()`

UnsetRank ensures that no value is present for Rank, not even an explicit nil
### GetPercentile

`func (o *ApiUserPerf200Response) GetPercentile() float32`

GetPercentile returns the Percentile field if non-nil, zero value otherwise.

### GetPercentileOk

`func (o *ApiUserPerf200Response) GetPercentileOk() (*float32, bool)`

GetPercentileOk returns a tuple with the Percentile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentile

`func (o *ApiUserPerf200Response) SetPercentile(v float32)`

SetPercentile sets Percentile field to given value.


### GetStat

`func (o *ApiUserPerf200Response) GetStat() ApiUserPerf200ResponseStat`

GetStat returns the Stat field if non-nil, zero value otherwise.

### GetStatOk

`func (o *ApiUserPerf200Response) GetStatOk() (*ApiUserPerf200ResponseStat, bool)`

GetStatOk returns a tuple with the Stat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStat

`func (o *ApiUserPerf200Response) SetStat(v ApiUserPerf200ResponseStat)`

SetStat sets Stat field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


