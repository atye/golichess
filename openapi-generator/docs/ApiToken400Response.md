# ApiToken400Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **string** | The cause of the error. | 
**ErrorDescription** | Pointer to **string** | The reason why the request was rejected. | [optional] 

## Methods

### NewApiToken400Response

`func NewApiToken400Response(error_ string, ) *ApiToken400Response`

NewApiToken400Response instantiates a new ApiToken400Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiToken400ResponseWithDefaults

`func NewApiToken400ResponseWithDefaults() *ApiToken400Response`

NewApiToken400ResponseWithDefaults instantiates a new ApiToken400Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ApiToken400Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ApiToken400Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ApiToken400Response) SetError(v string)`

SetError sets Error field to given value.


### GetErrorDescription

`func (o *ApiToken400Response) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *ApiToken400Response) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *ApiToken400Response) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *ApiToken400Response) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


