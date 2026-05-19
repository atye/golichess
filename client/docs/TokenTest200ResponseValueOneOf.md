# TokenTest200ResponseValueOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | Pointer to **string** |  | [optional] 
**Scopes** | Pointer to **string** | Comma-separated list of scopes. Empty string if the token has no scopes. | [optional] 
**Expires** | Pointer to **NullableInt32** | Unix-timestamp in milliseconds or null if the token never expires. | [optional] 

## Methods

### NewTokenTest200ResponseValueOneOf

`func NewTokenTest200ResponseValueOneOf() *TokenTest200ResponseValueOneOf`

NewTokenTest200ResponseValueOneOf instantiates a new TokenTest200ResponseValueOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokenTest200ResponseValueOneOfWithDefaults

`func NewTokenTest200ResponseValueOneOfWithDefaults() *TokenTest200ResponseValueOneOf`

NewTokenTest200ResponseValueOneOfWithDefaults instantiates a new TokenTest200ResponseValueOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *TokenTest200ResponseValueOneOf) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *TokenTest200ResponseValueOneOf) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *TokenTest200ResponseValueOneOf) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *TokenTest200ResponseValueOneOf) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetScopes

`func (o *TokenTest200ResponseValueOneOf) GetScopes() string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *TokenTest200ResponseValueOneOf) GetScopesOk() (*string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *TokenTest200ResponseValueOneOf) SetScopes(v string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *TokenTest200ResponseValueOneOf) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetExpires

`func (o *TokenTest200ResponseValueOneOf) GetExpires() int32`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *TokenTest200ResponseValueOneOf) GetExpiresOk() (*int32, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *TokenTest200ResponseValueOneOf) SetExpires(v int32)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *TokenTest200ResponseValueOneOf) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### SetExpiresNil

`func (o *TokenTest200ResponseValueOneOf) SetExpiresNil(b bool)`

 SetExpiresNil sets the value for Expires to be an explicit nil

### UnsetExpires
`func (o *TokenTest200ResponseValueOneOf) UnsetExpires()`

UnsetExpires ensures that no value is present for Expires, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


