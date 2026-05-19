# ApiToken200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenType** | **string** |  | 
**AccessToken** | **string** |  | 
**ExpiresIn** | **int32** |  | 

## Methods

### NewApiToken200Response

`func NewApiToken200Response(tokenType string, accessToken string, expiresIn int32, ) *ApiToken200Response`

NewApiToken200Response instantiates a new ApiToken200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiToken200ResponseWithDefaults

`func NewApiToken200ResponseWithDefaults() *ApiToken200Response`

NewApiToken200ResponseWithDefaults instantiates a new ApiToken200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenType

`func (o *ApiToken200Response) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *ApiToken200Response) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *ApiToken200Response) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetAccessToken

`func (o *ApiToken200Response) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *ApiToken200Response) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *ApiToken200Response) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetExpiresIn

`func (o *ApiToken200Response) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *ApiToken200Response) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *ApiToken200Response) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


