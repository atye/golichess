# \OAuthAPI

All URIs are relative to *https://lichess.org*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApiToken**](OAuthAPI.md#ApiToken) | **Post** /api/token | Obtain access token
[**ApiTokenDelete**](OAuthAPI.md#ApiTokenDelete) | **Delete** /api/token | Revoke access token
[**Oauth**](OAuthAPI.md#Oauth) | **Get** /oauth | Request authorization code
[**TokenTest**](OAuthAPI.md#TokenTest) | **Post** /api/token/test | Test multiple OAuth tokens



## ApiToken

> ApiToken200Response ApiToken(ctx).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).ClientId(clientId).Execute()

Obtain access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	grantType := "grantType_example" // string |  (optional)
	code := "code_example" // string | The authorization code that was sent in the `code` parameter to your `redirect_uri`. (optional)
	codeVerifier := "codeVerifier_example" // string | A `code_challenge` was used to request the authorization code. This must be the `code_verifier` it was derived from. (optional)
	redirectUri := "redirectUri_example" // string | Must match the `redirect_uri` used to request the authorization code. (optional)
	clientId := "clientId_example" // string | Must match the `client_id` used to request the authorization code. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.ApiToken(context.Background()).GrantType(grantType).Code(code).CodeVerifier(codeVerifier).RedirectUri(redirectUri).ClientId(clientId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.ApiToken``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ApiToken`: ApiToken200Response
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.ApiToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApiTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantType** | **string** |  | 
 **code** | **string** | The authorization code that was sent in the &#x60;code&#x60; parameter to your &#x60;redirect_uri&#x60;. | 
 **codeVerifier** | **string** | A &#x60;code_challenge&#x60; was used to request the authorization code. This must be the &#x60;code_verifier&#x60; it was derived from. | 
 **redirectUri** | **string** | Must match the &#x60;redirect_uri&#x60; used to request the authorization code. | 
 **clientId** | **string** | Must match the &#x60;client_id&#x60; used to request the authorization code. | 

### Return type

[**ApiToken200Response**](ApiToken200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApiTokenDelete

> ApiTokenDelete(ctx).Execute()

Revoke access token



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuthAPI.ApiTokenDelete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.ApiTokenDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApiTokenDeleteRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Oauth

> Oauth(ctx).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Scope(scope).Username(username).State(state).Execute()

Request authorization code



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	responseType := "responseType_example" // string | Must be `code`.
	clientId := "example.com" // string | Arbitrary identifier that uniquely identifies your application.
	redirectUri := "redirectUri_example" // string | The absolute URL that the user should be redirected to with the authorization result.
	codeChallengeMethod := "codeChallengeMethod_example" // string | Must be `S256`.
	codeChallenge := "codeChallenge_example" // string | Compute `BASE64URL(SHA256(code_verifier))`.
	scope := "scope_example" // string | Space separated list of requested OAuth scopes, if any. (optional)
	username := "username_example" // string | Hint that you want the user to log in with a specific Lichess username. (optional)
	state := "state_example" // string | Arbitrary state that will be returned verbatim with the authorization result. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OAuthAPI.Oauth(context.Background()).ResponseType(responseType).ClientId(clientId).RedirectUri(redirectUri).CodeChallengeMethod(codeChallengeMethod).CodeChallenge(codeChallenge).Scope(scope).Username(username).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.Oauth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOauthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **responseType** | **string** | Must be &#x60;code&#x60;. | 
 **clientId** | **string** | Arbitrary identifier that uniquely identifies your application. | 
 **redirectUri** | **string** | The absolute URL that the user should be redirected to with the authorization result. | 
 **codeChallengeMethod** | **string** | Must be &#x60;S256&#x60;. | 
 **codeChallenge** | **string** | Compute &#x60;BASE64URL(SHA256(code_verifier))&#x60;. | 
 **scope** | **string** | Space separated list of requested OAuth scopes, if any. | 
 **username** | **string** | Hint that you want the user to log in with a specific Lichess username. | 
 **state** | **string** | Arbitrary state that will be returned verbatim with the authorization result. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TokenTest

> map[string]TokenTest200ResponseValue TokenTest(ctx).Body(body).Execute()

Test multiple OAuth tokens



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/atye/golichess/openapigenerator"
)

func main() {
	body := "lip_jose,lip_badToken " // string | OAuth tokens separated by commas. Up to 1000.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OAuthAPI.TokenTest(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OAuthAPI.TokenTest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TokenTest`: map[string]TokenTest200ResponseValue
	fmt.Fprintf(os.Stdout, "Response from `OAuthAPI.TokenTest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTokenTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **string** | OAuth tokens separated by commas. Up to 1000. | 

### Return type

[**map[string]TokenTest200ResponseValue**](TokenTest200ResponseValue.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: text/plain
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

