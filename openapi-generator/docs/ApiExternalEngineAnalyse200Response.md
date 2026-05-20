# ApiExternalEngineAnalyse200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Time** | **int32** | Number of milliseconds the search has been going on | 
**Depth** | **int32** | Current search depth | 
**Nodes** | **int32** | Number of nodes visited so far | 
**Pvs** | [**[]ApiExternalEngineAnalyse200ResponsePvsInner**](ApiExternalEngineAnalyse200ResponsePvsInner.md) | Information about up to 5 pvs, with the primary pv at index 0. | 

## Methods

### NewApiExternalEngineAnalyse200Response

`func NewApiExternalEngineAnalyse200Response(time int32, depth int32, nodes int32, pvs []ApiExternalEngineAnalyse200ResponsePvsInner, ) *ApiExternalEngineAnalyse200Response`

NewApiExternalEngineAnalyse200Response instantiates a new ApiExternalEngineAnalyse200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyse200ResponseWithDefaults

`func NewApiExternalEngineAnalyse200ResponseWithDefaults() *ApiExternalEngineAnalyse200Response`

NewApiExternalEngineAnalyse200ResponseWithDefaults instantiates a new ApiExternalEngineAnalyse200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTime

`func (o *ApiExternalEngineAnalyse200Response) GetTime() int32`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApiExternalEngineAnalyse200Response) GetTimeOk() (*int32, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApiExternalEngineAnalyse200Response) SetTime(v int32)`

SetTime sets Time field to given value.


### GetDepth

`func (o *ApiExternalEngineAnalyse200Response) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiExternalEngineAnalyse200Response) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiExternalEngineAnalyse200Response) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetNodes

`func (o *ApiExternalEngineAnalyse200Response) GetNodes() int32`

GetNodes returns the Nodes field if non-nil, zero value otherwise.

### GetNodesOk

`func (o *ApiExternalEngineAnalyse200Response) GetNodesOk() (*int32, bool)`

GetNodesOk returns a tuple with the Nodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodes

`func (o *ApiExternalEngineAnalyse200Response) SetNodes(v int32)`

SetNodes sets Nodes field to given value.


### GetPvs

`func (o *ApiExternalEngineAnalyse200Response) GetPvs() []ApiExternalEngineAnalyse200ResponsePvsInner`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *ApiExternalEngineAnalyse200Response) GetPvsOk() (*[]ApiExternalEngineAnalyse200ResponsePvsInner, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *ApiExternalEngineAnalyse200Response) SetPvs(v []ApiExternalEngineAnalyse200ResponsePvsInner)`

SetPvs sets Pvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


