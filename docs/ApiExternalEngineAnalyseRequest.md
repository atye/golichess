# ApiExternalEngineAnalyseRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientSecret** | **string** |  | 
**Work** | [**ApiExternalEngineAnalyseRequestWork**](ApiExternalEngineAnalyseRequestWork.md) |  | 

## Methods

### NewApiExternalEngineAnalyseRequest

`func NewApiExternalEngineAnalyseRequest(clientSecret string, work ApiExternalEngineAnalyseRequestWork, ) *ApiExternalEngineAnalyseRequest`

NewApiExternalEngineAnalyseRequest instantiates a new ApiExternalEngineAnalyseRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAnalyseRequestWithDefaults

`func NewApiExternalEngineAnalyseRequestWithDefaults() *ApiExternalEngineAnalyseRequest`

NewApiExternalEngineAnalyseRequestWithDefaults instantiates a new ApiExternalEngineAnalyseRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientSecret

`func (o *ApiExternalEngineAnalyseRequest) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiExternalEngineAnalyseRequest) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiExternalEngineAnalyseRequest) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetWork

`func (o *ApiExternalEngineAnalyseRequest) GetWork() ApiExternalEngineAnalyseRequestWork`

GetWork returns the Work field if non-nil, zero value otherwise.

### GetWorkOk

`func (o *ApiExternalEngineAnalyseRequest) GetWorkOk() (*ApiExternalEngineAnalyseRequestWork, bool)`

GetWorkOk returns a tuple with the Work field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWork

`func (o *ApiExternalEngineAnalyseRequest) SetWork(v ApiExternalEngineAnalyseRequestWork)`

SetWork sets Work field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


