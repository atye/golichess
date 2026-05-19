# ApiStormDashboard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | [**[]ApiStormDashboard200ResponseDaysInner**](ApiStormDashboard200ResponseDaysInner.md) |  | 
**High** | [**ApiStormDashboard200ResponseHigh**](ApiStormDashboard200ResponseHigh.md) |  | 

## Methods

### NewApiStormDashboard200Response

`func NewApiStormDashboard200Response(days []ApiStormDashboard200ResponseDaysInner, high ApiStormDashboard200ResponseHigh, ) *ApiStormDashboard200Response`

NewApiStormDashboard200Response instantiates a new ApiStormDashboard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStormDashboard200ResponseWithDefaults

`func NewApiStormDashboard200ResponseWithDefaults() *ApiStormDashboard200Response`

NewApiStormDashboard200ResponseWithDefaults instantiates a new ApiStormDashboard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *ApiStormDashboard200Response) GetDays() []ApiStormDashboard200ResponseDaysInner`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *ApiStormDashboard200Response) GetDaysOk() (*[]ApiStormDashboard200ResponseDaysInner, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *ApiStormDashboard200Response) SetDays(v []ApiStormDashboard200ResponseDaysInner)`

SetDays sets Days field to given value.


### GetHigh

`func (o *ApiStormDashboard200Response) GetHigh() ApiStormDashboard200ResponseHigh`

GetHigh returns the High field if non-nil, zero value otherwise.

### GetHighOk

`func (o *ApiStormDashboard200Response) GetHighOk() (*ApiStormDashboard200ResponseHigh, bool)`

GetHighOk returns a tuple with the High field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHigh

`func (o *ApiStormDashboard200Response) SetHigh(v ApiStormDashboard200ResponseHigh)`

SetHigh sets High field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


