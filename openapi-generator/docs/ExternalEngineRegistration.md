# ExternalEngineRegistration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name of the engine. | 
**MaxThreads** | **int32** | Maximum number of available threads. | 
**MaxHash** | **int32** | Maximum available hash table size, in MiB. | 
**Variants** | Pointer to [**[]UciVariant**](UciVariant.md) | Optional list of supported chess variants. | [optional] 
**ProviderSecret** | **string** | A random token that can be used to [wait for analysis requests](#tag/external-engine/POST/api/external-engine/work) and provide analysis.  The engine provider should securely generate a random string.  The token will not be readable again, even by the user.  The analysis provider can register multiple engines with the same token, even for different users, and wait for analysis requests from any of them. In this case, the request must not be made via CORS, so that the token is not revealed to any of the users.  | 
**ProviderData** | Pointer to **string** | Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the &#x60;providerSecret&#x60;.  | [optional] 

## Methods

### NewExternalEngineRegistration

`func NewExternalEngineRegistration(name string, maxThreads int32, maxHash int32, providerSecret string, ) *ExternalEngineRegistration`

NewExternalEngineRegistration instantiates a new ExternalEngineRegistration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineRegistrationWithDefaults

`func NewExternalEngineRegistrationWithDefaults() *ExternalEngineRegistration`

NewExternalEngineRegistrationWithDefaults instantiates a new ExternalEngineRegistration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ExternalEngineRegistration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalEngineRegistration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalEngineRegistration) SetName(v string)`

SetName sets Name field to given value.


### GetMaxThreads

`func (o *ExternalEngineRegistration) GetMaxThreads() int32`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *ExternalEngineRegistration) GetMaxThreadsOk() (*int32, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *ExternalEngineRegistration) SetMaxThreads(v int32)`

SetMaxThreads sets MaxThreads field to given value.


### GetMaxHash

`func (o *ExternalEngineRegistration) GetMaxHash() int32`

GetMaxHash returns the MaxHash field if non-nil, zero value otherwise.

### GetMaxHashOk

`func (o *ExternalEngineRegistration) GetMaxHashOk() (*int32, bool)`

GetMaxHashOk returns a tuple with the MaxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHash

`func (o *ExternalEngineRegistration) SetMaxHash(v int32)`

SetMaxHash sets MaxHash field to given value.


### GetVariants

`func (o *ExternalEngineRegistration) GetVariants() []UciVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExternalEngineRegistration) GetVariantsOk() (*[]UciVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExternalEngineRegistration) SetVariants(v []UciVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *ExternalEngineRegistration) HasVariants() bool`

HasVariants returns a boolean if a field has been set.

### GetProviderSecret

`func (o *ExternalEngineRegistration) GetProviderSecret() string`

GetProviderSecret returns the ProviderSecret field if non-nil, zero value otherwise.

### GetProviderSecretOk

`func (o *ExternalEngineRegistration) GetProviderSecretOk() (*string, bool)`

GetProviderSecretOk returns a tuple with the ProviderSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderSecret

`func (o *ExternalEngineRegistration) SetProviderSecret(v string)`

SetProviderSecret sets ProviderSecret field to given value.


### GetProviderData

`func (o *ExternalEngineRegistration) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *ExternalEngineRegistration) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *ExternalEngineRegistration) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *ExternalEngineRegistration) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


