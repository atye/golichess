# CloudEval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Depth** | **int32** |  | 
**Fen** | **string** |  | 
**Knodes** | **int32** |  | 
**Pvs** | [**[]CloudEvalPvsInner**](CloudEvalPvsInner.md) |  | 

## Methods

### NewCloudEval

`func NewCloudEval(depth int32, fen string, knodes int32, pvs []CloudEvalPvsInner, ) *CloudEval`

NewCloudEval instantiates a new CloudEval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudEvalWithDefaults

`func NewCloudEvalWithDefaults() *CloudEval`

NewCloudEvalWithDefaults instantiates a new CloudEval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDepth

`func (o *CloudEval) GetDepth() int32`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *CloudEval) GetDepthOk() (*int32, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *CloudEval) SetDepth(v int32)`

SetDepth sets Depth field to given value.


### GetFen

`func (o *CloudEval) GetFen() string`

GetFen returns the Fen field if non-nil, zero value otherwise.

### GetFenOk

`func (o *CloudEval) GetFenOk() (*string, bool)`

GetFenOk returns a tuple with the Fen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFen

`func (o *CloudEval) SetFen(v string)`

SetFen sets Fen field to given value.


### GetKnodes

`func (o *CloudEval) GetKnodes() int32`

GetKnodes returns the Knodes field if non-nil, zero value otherwise.

### GetKnodesOk

`func (o *CloudEval) GetKnodesOk() (*int32, bool)`

GetKnodesOk returns a tuple with the Knodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnodes

`func (o *CloudEval) SetKnodes(v int32)`

SetKnodes sets Knodes field to given value.


### GetPvs

`func (o *CloudEval) GetPvs() []CloudEvalPvsInner`

GetPvs returns the Pvs field if non-nil, zero value otherwise.

### GetPvsOk

`func (o *CloudEval) GetPvsOk() (*[]CloudEvalPvsInner, bool)`

GetPvsOk returns a tuple with the Pvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPvs

`func (o *CloudEval) SetPvs(v []CloudEvalPvsInner)`

SetPvs sets Pvs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


