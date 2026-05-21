# ExternalEngine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique engine registration ID. | 
**Name** | **string** | Display name of the engine. | 
**ClientSecret** | **string** | A secret token that can be used to [*request* analysis](#tag/external-engine/POST/api/external-engine/{id}/analyse) from this external engine.  | 
**UserId** | **string** | The user this engine has been registered for. | 
**MaxThreads** | **int32** | Maximum number of available threads. | 
**MaxHash** | **int32** | Maximum available hash table size, in MiB. | 
**Variants** | [**[]UciVariant**](UciVariant.md) | List of supported chess variants. | 
**ProviderData** | Pointer to **NullableString** | Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the &#x60;providerSecret&#x60;.  | [optional] 

## Methods

### NewExternalEngine

`func NewExternalEngine(id string, name string, clientSecret string, userId string, maxThreads int32, maxHash int32, variants []UciVariant, ) *ExternalEngine`

NewExternalEngine instantiates a new ExternalEngine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalEngineWithDefaults

`func NewExternalEngineWithDefaults() *ExternalEngine`

NewExternalEngineWithDefaults instantiates a new ExternalEngine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalEngine) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalEngine) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalEngine) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExternalEngine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalEngine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalEngine) SetName(v string)`

SetName sets Name field to given value.


### GetClientSecret

`func (o *ExternalEngine) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *ExternalEngine) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *ExternalEngine) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetUserId

`func (o *ExternalEngine) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ExternalEngine) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ExternalEngine) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetMaxThreads

`func (o *ExternalEngine) GetMaxThreads() int32`

GetMaxThreads returns the MaxThreads field if non-nil, zero value otherwise.

### GetMaxThreadsOk

`func (o *ExternalEngine) GetMaxThreadsOk() (*int32, bool)`

GetMaxThreadsOk returns a tuple with the MaxThreads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxThreads

`func (o *ExternalEngine) SetMaxThreads(v int32)`

SetMaxThreads sets MaxThreads field to given value.


### GetMaxHash

`func (o *ExternalEngine) GetMaxHash() int32`

GetMaxHash returns the MaxHash field if non-nil, zero value otherwise.

### GetMaxHashOk

`func (o *ExternalEngine) GetMaxHashOk() (*int32, bool)`

GetMaxHashOk returns a tuple with the MaxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxHash

`func (o *ExternalEngine) SetMaxHash(v int32)`

SetMaxHash sets MaxHash field to given value.


### GetVariants

`func (o *ExternalEngine) GetVariants() []UciVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *ExternalEngine) GetVariantsOk() (*[]UciVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *ExternalEngine) SetVariants(v []UciVariant)`

SetVariants sets Variants field to given value.


### GetProviderData

`func (o *ExternalEngine) GetProviderData() string`

GetProviderData returns the ProviderData field if non-nil, zero value otherwise.

### GetProviderDataOk

`func (o *ExternalEngine) GetProviderDataOk() (*string, bool)`

GetProviderDataOk returns a tuple with the ProviderData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderData

`func (o *ExternalEngine) SetProviderData(v string)`

SetProviderData sets ProviderData field to given value.

### HasProviderData

`func (o *ExternalEngine) HasProviderData() bool`

HasProviderData returns a boolean if a field has been set.

### SetProviderDataNil

`func (o *ExternalEngine) SetProviderDataNil(b bool)`

 SetProviderDataNil sets the value for ProviderData to be an explicit nil

### UnsetProviderData
`func (o *ExternalEngine) UnsetProviderData()`

UnsetProviderData ensures that no value is present for ProviderData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


