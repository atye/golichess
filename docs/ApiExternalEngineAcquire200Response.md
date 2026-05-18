# ApiExternalEngineAcquire200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Work** | [**ApiExternalEngineAnalyseRequestWork**](ApiExternalEngineAnalyseRequestWork.md) |  | 
**Engine** | [**ApiExternalEngineList200ResponseInner**](ApiExternalEngineList200ResponseInner.md) |  | 

## Methods

### NewApiExternalEngineAcquire200Response

`func NewApiExternalEngineAcquire200Response(id string, work ApiExternalEngineAnalyseRequestWork, engine ApiExternalEngineList200ResponseInner, ) *ApiExternalEngineAcquire200Response`

NewApiExternalEngineAcquire200Response instantiates a new ApiExternalEngineAcquire200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineAcquire200ResponseWithDefaults

`func NewApiExternalEngineAcquire200ResponseWithDefaults() *ApiExternalEngineAcquire200Response`

NewApiExternalEngineAcquire200ResponseWithDefaults instantiates a new ApiExternalEngineAcquire200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiExternalEngineAcquire200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiExternalEngineAcquire200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiExternalEngineAcquire200Response) SetId(v string)`

SetId sets Id field to given value.


### GetWork

`func (o *ApiExternalEngineAcquire200Response) GetWork() ApiExternalEngineAnalyseRequestWork`

GetWork returns the Work field if non-nil, zero value otherwise.

### GetWorkOk

`func (o *ApiExternalEngineAcquire200Response) GetWorkOk() (*ApiExternalEngineAnalyseRequestWork, bool)`

GetWorkOk returns a tuple with the Work field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWork

`func (o *ApiExternalEngineAcquire200Response) SetWork(v ApiExternalEngineAnalyseRequestWork)`

SetWork sets Work field to given value.


### GetEngine

`func (o *ApiExternalEngineAcquire200Response) GetEngine() ApiExternalEngineList200ResponseInner`

GetEngine returns the Engine field if non-nil, zero value otherwise.

### GetEngineOk

`func (o *ApiExternalEngineAcquire200Response) GetEngineOk() (*ApiExternalEngineList200ResponseInner, bool)`

GetEngineOk returns a tuple with the Engine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngine

`func (o *ApiExternalEngineAcquire200Response) SetEngine(v ApiExternalEngineList200ResponseInner)`

SetEngine sets Engine field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


