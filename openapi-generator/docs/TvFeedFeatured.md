# TvFeedFeatured

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**T** | **string** | The type of message. A summary of the game is sent as the first message and when the featured game changes. Subsequent messages are just the X-FEN, last move, and clocks.  | 
**D** | [**Featured**](Featured.md) |  | 

## Methods

### NewTvFeedFeatured

`func NewTvFeedFeatured(t string, d Featured, ) *TvFeedFeatured`

NewTvFeedFeatured instantiates a new TvFeedFeatured object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTvFeedFeaturedWithDefaults

`func NewTvFeedFeaturedWithDefaults() *TvFeedFeatured`

NewTvFeedFeaturedWithDefaults instantiates a new TvFeedFeatured object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetT

`func (o *TvFeedFeatured) GetT() string`

GetT returns the T field if non-nil, zero value otherwise.

### GetTOk

`func (o *TvFeedFeatured) GetTOk() (*string, bool)`

GetTOk returns a tuple with the T field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT

`func (o *TvFeedFeatured) SetT(v string)`

SetT sets T field to given value.


### GetD

`func (o *TvFeedFeatured) GetD() Featured`

GetD returns the D field if non-nil, zero value otherwise.

### GetDOk

`func (o *TvFeedFeatured) GetDOk() (*Featured, bool)`

GetDOk returns a tuple with the D field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetD

`func (o *TvFeedFeatured) SetD(v Featured)`

SetD sets D field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


