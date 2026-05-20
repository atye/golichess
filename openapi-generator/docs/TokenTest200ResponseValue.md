# TokenTest200ResponseValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** | Comma-separated list of scopes. Empty string if the token has no scopes. | [optional] 
**Expires** | Pointer to **int32** | Unix-timestamp in milliseconds or null if the token never expires. | [optional] 

## Methods

### NewTokenTest200ResponseValue

`func NewTokenTest200ResponseValue() *TokenTest200ResponseValue`

NewTokenTest200ResponseValue instantiates a new TokenTest200ResponseValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTest200ResponseValueWithDefaults

`func NewTokenTest200ResponseValueWithDefaults() *TokenTest200ResponseValue`

NewTokenTest200ResponseValueWithDefaults instantiates a new TokenTest200ResponseValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TokenTest200ResponseValue) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenTest200ResponseValue) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenTest200ResponseValue) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TokenTest200ResponseValue) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetScopes

`func (o *TokenTest200ResponseValue) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenTest200ResponseValue) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenTest200ResponseValue) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenTest200ResponseValue) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExpires

`func (o *TokenTest200ResponseValue) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenTest200ResponseValue) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenTest200ResponseValue) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenTest200ResponseValue) HasExpires() bool`

HasExpires returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


