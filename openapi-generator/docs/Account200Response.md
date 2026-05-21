# Account200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prefs** | Pointer to [**UserPreferences**](UserPreferences.md) |  | [optional] 
**Language** | Pointer to **string** |  | [optional] 

## Methods

### NewAccount200Response

`func NewAccount200Response() *Account200Response`

NewAccount200Response instantiates a new Account200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccount200ResponseWithDefaults

`func NewAccount200ResponseWithDefaults() *Account200Response`

NewAccount200ResponseWithDefaults instantiates a new Account200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrefs

`func (o *Account200Response) GetPrefs() UserPreferences`

GetPrefs returns the Prefs field if non-nil, zero value otherwise.

### GetPrefsOk

`func (o *Account200Response) GetPrefsOk() (*UserPreferences, bool)`

GetPrefsOk returns a tuple with the Prefs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefs

`func (o *Account200Response) SetPrefs(v UserPreferences)`

SetPrefs sets Prefs field to given value.

### HasPrefs

`func (o *Account200Response) HasPrefs() bool`

HasPrefs returns a boolean if a field has been set.

### GetLanguage

`func (o *Account200Response) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Account200Response) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Account200Response) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Account200Response) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


