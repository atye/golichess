# ApiExternalEngineList200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique engine registration ID. | 
**Name** | **string** | Display name of the engine. | 
**ClientSecret** | **string** | A secret token that can be used to [*request* analysis](#tag/external-engine/POST/api/external-engine/{id}/analyse) from this external engine.  | 
**UserId** | **string** | The user this engine has been registered for. | 
**MaxThreads** | **int32** | Maximum number of available threads. | 
**MaxHash** | **int32** | Maximum available hash table size, in MiB. | 
**Variants** | **[]string** | List of supported chess variants. | 
**ProviderData** | Pointer to **NullableString** | Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the &#x60;providerSecret&#x60;.  | [optional] 

## Methods

### NewApiExternalEngineList200ResponseInner

`func NewApiExternalEngineList200ResponseInner(id string, name string, clientSecret string, userId string, maxThreads int32, maxHash int32, variants []string, ) *ApiExternalEngineList200ResponseInner`

NewApiExternalEngineList200ResponseInner instantiates a new ApiExternalEngineList200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiExternalEngineList200ResponseInnerWithDefaults

`func NewApiExternalEngineList200ResponseInnerWithDefaults() *ApiExternalEngineList200ResponseInner`

NewApiExternalEngineList200ResponseInnerWithDefaults instantiates a new ApiExternalEngineList200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiExternalEngineList200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiExternalEngineList200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiExternalEngineList200ResponseInner) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ApiExternalEngineList200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiExternalEngineList200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiExternalEngineList200ResponseInner) SetName(v string)`

SetName sets Name field to given value.


### GetClientSecret

`func (o *ApiExternalEngineList200ResponseInner) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ApiExternalEngineList200ResponseInner) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ApiExternalEngineList200ResponseInner) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetUserId

`func (o *ApiExternalEngineList200ResponseInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiExternalEngineList200ResponseInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiExternalEngineList200ResponseInner) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMaxThreads

`func (o *ApiExternalEngineList200ResponseInner) GetMaxThreads() int32`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *ApiExternalEngineList200ResponseInner) GetMaxThreadsOk() (*int32, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *ApiExternalEngineList200ResponseInner) SetMaxThreads(v int32)`

SetMaxThreads sets MaxThreads field to given value.


### GetMaxHash

`func (o *ApiExternalEngineList200ResponseInner) GetMaxHash() int32`

GetMaxHash returns the MaxHash field if non-nil, zero value otherwise.

### GetMaxHashOk

`func (o *ApiExternalEngineList200ResponseInner) GetMaxHashOk() (*int32, bool)`

GetMaxHashOk returns a tuple with the MaxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHash

`func (o *ApiExternalEngineList200ResponseInner) SetMaxHash(v int32)`

SetMaxHash sets MaxHash field to given value.


### GetVariants

`func (o *ApiExternalEngineList200ResponseInner) GetVariants() []string`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ApiExternalEngineList200ResponseInner) GetVariantsOk() (*[]string, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ApiExternalEngineList200ResponseInner) SetVariants(v []string)`

SetVariants sets Variants field to given value.


### GetProviderData

`func (o *ApiExternalEngineList200ResponseInner) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *ApiExternalEngineList200ResponseInner) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *ApiExternalEngineList200ResponseInner) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *ApiExternalEngineList200ResponseInner) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### SetProviderDataNil

`func (o *ApiExternalEngineList200ResponseInner) SetProviderDataNil(b bool)`

 SetProviderDataNil sets the value for ProviderData to be an explicit nil

### UnsetProviderData
`func (o *ApiExternalEngineList200ResponseInner) UnsetProviderData()`

UnsetProviderData ensures that no value is present for ProviderData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


