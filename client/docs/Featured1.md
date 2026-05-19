# Featured1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | **string** | The type of message. A summary of the game is sent as the first message and when the featured game changes. Subsequent messages are just the X-FEN, last move, and clocks.  | 
**D** | [**Featured1**](Featured1.md) |  | 

## Methods

### NewFeatured1

`func NewFeatured1(t string, d Featured1, ) *Featured1`

NewFeatured1 instantiates a new Featured1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeatured1WithDefaults

`func NewFeatured1WithDefaults() *Featured1`

NewFeatured1WithDefaults instantiates a new Featured1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *Featured1) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *Featured1) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *Featured1) SetT(v string)`

SetT sets T field to given value.


### GetD

`func (o *Featured1) GetD() Featured1`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *Featured1) GetDOk() (*Featured1, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *Featured1) SetD(v Featured1)`

SetD sets D field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


