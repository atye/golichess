# ApiCloudEval200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **int32** |  | 
**Fen** | **string** |  | 
**Knodes** | **int32** |  | 
**Pvs** | [**[]ApiCloudEval200ResponsePvsInner**](ApiCloudEval200ResponsePvsInner.md) |  | 

## Methods

### NewApiCloudEval200Response

`func NewApiCloudEval200Response(depth int32, fen string, knodes int32, pvs []ApiCloudEval200ResponsePvsInner, ) *ApiCloudEval200Response`

NewApiCloudEval200Response instantiates a new ApiCloudEval200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiCloudEval200ResponseWithDefaults

`func NewApiCloudEval200ResponseWithDefaults() *ApiCloudEval200Response`

NewApiCloudEval200ResponseWithDefaults instantiates a new ApiCloudEval200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *ApiCloudEval200Response) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *ApiCloudEval200Response) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *ApiCloudEval200Response) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetFen

`func (o *ApiCloudEval200Response) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *ApiCloudEval200Response) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *ApiCloudEval200Response) SetFen(v string)`

SetFen sets Fen field to given value.


### GetKnodes

`func (o *ApiCloudEval200Response) GetKnodes() int32`

GetKnodes returns the Knodes field if non-nil, zero value otherwise.

### GetKnodesOk

`func (o *ApiCloudEval200Response) GetKnodesOk() (*int32, bool)`

GetKnodesOk returns a tuple with the Knodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnodes

`func (o *ApiCloudEval200Response) SetKnodes(v int32)`

SetKnodes sets Knodes field to given value.


### GetPvs

`func (o *ApiCloudEval200Response) GetPvs() []ApiCloudEval200ResponsePvsInner`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *ApiCloudEval200Response) GetPvsOk() (*[]ApiCloudEval200ResponsePvsInner, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *ApiCloudEval200Response) SetPvs(v []ApiCloudEval200ResponsePvsInner)`

SetPvs sets Pvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


