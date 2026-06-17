/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.146
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


type OAuthAPI interface {

	/*
	ApiToken Obtain access token

	OAuth2 token endpoint. Exchanges an authorization code for an access token.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OAuthAPIApiTokenRequest
	*/
	ApiToken(ctx context.Context) OAuthAPIApiTokenRequest

	// ApiTokenExecute executes the request
	//  @return ApiToken200Response
	ApiTokenExecute(r OAuthAPIApiTokenRequest) (*ApiToken200Response, *http.Response, error)

	/*
	ApiTokenDelete Revoke access token

	Revokes the access token sent as Bearer for this request.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OAuthAPIApiTokenDeleteRequest
	*/
	ApiTokenDelete(ctx context.Context) OAuthAPIApiTokenDeleteRequest

	// ApiTokenDeleteExecute executes the request
	ApiTokenDeleteExecute(r OAuthAPIApiTokenDeleteRequest) (*http.Response, error)

	/*
	Oauth Request authorization code

	OAuth2 authorization endpoint.
Start the OAuth2 Authorization Code Flow with PKCE by securely
generating two random strings unique to each authorization
request:

* `code_verifier`
* `state`

Store these in session storage. Make sure not to reveal `code_verifier`
to eavesdroppers. Do not show it in URLs, do not abuse `state` to store
it, do not send it over insecure connections. However it is fine if
the user themselves can extract `code_verifier`, which will always be
possible for fully client-side apps.
Then send the user to this endpoint. They will be prompted to grant
authorization and then be redirected back to the given `redirect_uri`.
If the authorization failed, the following query string parameters will
be appended to the redirection:

* `error`, in particular with value `access_denied` if the user
   cancelled authorization
* `error_description` to aid debugging
* `state`, exactly as passed in the `state` parameter

If the authorization succeeded, the following query string parameters
will be appended to the redirection:

* `code`, containing a fresh short-lived authorization code
* `state`, exactly as passed in the `state` parameter

Next, to defend against cross site request forgery, check that the
returned `state` matches the `state` you originally generated.

Finally, continue by using the authorization code to
[obtain an access token](#tag/oauth/POST/api/token).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OAuthAPIOauthRequest
	*/
	Oauth(ctx context.Context) OAuthAPIOauthRequest

	// OauthExecute executes the request
	OauthExecute(r OAuthAPIOauthRequest) (*http.Response, error)

	/*
	TokenTest Test multiple OAuth tokens

	For up to 1000 OAuth tokens,
returns their associated user ID and scopes,
or `null` if the token is invalid.
The method is `POST` so a longer list of tokens can be sent in the request body.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OAuthAPITokenTestRequest
	*/
	TokenTest(ctx context.Context) OAuthAPITokenTestRequest

	// TokenTestExecute executes the request
	//  @return map[string]TokenTest200ResponseValue
	TokenTestExecute(r OAuthAPITokenTestRequest) (*map[string]TokenTest200ResponseValue, *http.Response, error)
}

// OAuthAPIService OAuthAPI service
type OAuthAPIService service

type OAuthAPIApiTokenRequest struct {
	ctx context.Context
	ApiService OAuthAPI
	grantType *string
	code *string
	codeVerifier *string
	redirectUri *string
	clientId *string
}

func (r OAuthAPIApiTokenRequest) GrantType(grantType string) OAuthAPIApiTokenRequest {
	r.grantType = &grantType
	return r
}

// The authorization code that was sent in the &#x60;code&#x60; parameter to your &#x60;redirect_uri&#x60;.
func (r OAuthAPIApiTokenRequest) Code(code string) OAuthAPIApiTokenRequest {
	r.code = &code
	return r
}

// A &#x60;code_challenge&#x60; was used to request the authorization code. This must be the &#x60;code_verifier&#x60; it was derived from.
func (r OAuthAPIApiTokenRequest) CodeVerifier(codeVerifier string) OAuthAPIApiTokenRequest {
	r.codeVerifier = &codeVerifier
	return r
}

// Must match the &#x60;redirect_uri&#x60; used to request the authorization code.
func (r OAuthAPIApiTokenRequest) RedirectUri(redirectUri string) OAuthAPIApiTokenRequest {
	r.redirectUri = &redirectUri
	return r
}

// Must match the &#x60;client_id&#x60; used to request the authorization code.
func (r OAuthAPIApiTokenRequest) ClientId(clientId string) OAuthAPIApiTokenRequest {
	r.clientId = &clientId
	return r
}

func (r OAuthAPIApiTokenRequest) Execute() (*ApiToken200Response, *http.Response, error) {
	return r.ApiService.ApiTokenExecute(r)
}

/*
ApiToken Obtain access token

OAuth2 token endpoint. Exchanges an authorization code for an access token.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OAuthAPIApiTokenRequest
*/
func (a *OAuthAPIService) ApiToken(ctx context.Context) OAuthAPIApiTokenRequest {
	return OAuthAPIApiTokenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiToken200Response
func (a *OAuthAPIService) ApiTokenExecute(r OAuthAPIApiTokenRequest) (*ApiToken200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiToken200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuthAPIService.ApiToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.grantType != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "grant_type", r.grantType, "", "")
	}
	if r.code != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "code", r.code, "", "")
	}
	if r.codeVerifier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "code_verifier", r.codeVerifier, "", "")
	}
	if r.redirectUri != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "redirect_uri", r.redirectUri, "", "")
	}
	if r.clientId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "client_id", r.clientId, "", "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v OAuthError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type OAuthAPIApiTokenDeleteRequest struct {
	ctx context.Context
	ApiService OAuthAPI
}

func (r OAuthAPIApiTokenDeleteRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiTokenDeleteExecute(r)
}

/*
ApiTokenDelete Revoke access token

Revokes the access token sent as Bearer for this request.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OAuthAPIApiTokenDeleteRequest
*/
func (a *OAuthAPIService) ApiTokenDelete(ctx context.Context) OAuthAPIApiTokenDeleteRequest {
	return OAuthAPIApiTokenDeleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OAuthAPIService) ApiTokenDeleteExecute(r OAuthAPIApiTokenDeleteRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuthAPIService.ApiTokenDelete")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OAuthAPIOauthRequest struct {
	ctx context.Context
	ApiService OAuthAPI
	responseType *string
	clientId *string
	redirectUri *string
	codeChallengeMethod *string
	codeChallenge *string
	scope *string
	username *string
	state *string
}

// Must be &#x60;code&#x60;.
func (r OAuthAPIOauthRequest) ResponseType(responseType string) OAuthAPIOauthRequest {
	r.responseType = &responseType
	return r
}

// Arbitrary identifier that uniquely identifies your application.
func (r OAuthAPIOauthRequest) ClientId(clientId string) OAuthAPIOauthRequest {
	r.clientId = &clientId
	return r
}

// The absolute URL that the user should be redirected to with the authorization result.
func (r OAuthAPIOauthRequest) RedirectUri(redirectUri string) OAuthAPIOauthRequest {
	r.redirectUri = &redirectUri
	return r
}

// Must be &#x60;S256&#x60;.
func (r OAuthAPIOauthRequest) CodeChallengeMethod(codeChallengeMethod string) OAuthAPIOauthRequest {
	r.codeChallengeMethod = &codeChallengeMethod
	return r
}

// Compute &#x60;BASE64URL(SHA256(code_verifier))&#x60;.
func (r OAuthAPIOauthRequest) CodeChallenge(codeChallenge string) OAuthAPIOauthRequest {
	r.codeChallenge = &codeChallenge
	return r
}

// Space separated list of requested OAuth scopes, if any.
func (r OAuthAPIOauthRequest) Scope(scope string) OAuthAPIOauthRequest {
	r.scope = &scope
	return r
}

// Hint that you want the user to log in with a specific Lichess username.
func (r OAuthAPIOauthRequest) Username(username string) OAuthAPIOauthRequest {
	r.username = &username
	return r
}

// Arbitrary state that will be returned verbatim with the authorization result.
func (r OAuthAPIOauthRequest) State(state string) OAuthAPIOauthRequest {
	r.state = &state
	return r
}

func (r OAuthAPIOauthRequest) Execute() (*http.Response, error) {
	return r.ApiService.OauthExecute(r)
}

/*
Oauth Request authorization code

OAuth2 authorization endpoint.
Start the OAuth2 Authorization Code Flow with PKCE by securely
generating two random strings unique to each authorization
request:

* `code_verifier`
* `state`

Store these in session storage. Make sure not to reveal `code_verifier`
to eavesdroppers. Do not show it in URLs, do not abuse `state` to store
it, do not send it over insecure connections. However it is fine if
the user themselves can extract `code_verifier`, which will always be
possible for fully client-side apps.
Then send the user to this endpoint. They will be prompted to grant
authorization and then be redirected back to the given `redirect_uri`.
If the authorization failed, the following query string parameters will
be appended to the redirection:

* `error`, in particular with value `access_denied` if the user
   cancelled authorization
* `error_description` to aid debugging
* `state`, exactly as passed in the `state` parameter

If the authorization succeeded, the following query string parameters
will be appended to the redirection:

* `code`, containing a fresh short-lived authorization code
* `state`, exactly as passed in the `state` parameter

Next, to defend against cross site request forgery, check that the
returned `state` matches the `state` you originally generated.

Finally, continue by using the authorization code to
[obtain an access token](#tag/oauth/POST/api/token).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OAuthAPIOauthRequest
*/
func (a *OAuthAPIService) Oauth(ctx context.Context) OAuthAPIOauthRequest {
	return OAuthAPIOauthRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
func (a *OAuthAPIService) OauthExecute(r OAuthAPIOauthRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuthAPIService.Oauth")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.responseType == nil {
		return nil, reportError("responseType is required and must be specified")
	}
	if r.clientId == nil {
		return nil, reportError("clientId is required and must be specified")
	}
	if r.redirectUri == nil {
		return nil, reportError("redirectUri is required and must be specified")
	}
	if r.codeChallengeMethod == nil {
		return nil, reportError("codeChallengeMethod is required and must be specified")
	}
	if r.codeChallenge == nil {
		return nil, reportError("codeChallenge is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "response_type", r.responseType, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "client_id", r.clientId, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "redirect_uri", r.redirectUri, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "code_challenge_method", r.codeChallengeMethod, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "code_challenge", r.codeChallenge, "form", "")
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
	}
	if r.username != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "username", r.username, "form", "")
	}
	if r.state != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "state", r.state, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type OAuthAPITokenTestRequest struct {
	ctx context.Context
	ApiService OAuthAPI
	body *string
}

// OAuth tokens separated by commas. Up to 1000.
func (r OAuthAPITokenTestRequest) Body(body string) OAuthAPITokenTestRequest {
	r.body = &body
	return r
}

func (r OAuthAPITokenTestRequest) Execute() (*map[string]TokenTest200ResponseValue, *http.Response, error) {
	return r.ApiService.TokenTestExecute(r)
}

/*
TokenTest Test multiple OAuth tokens

For up to 1000 OAuth tokens,
returns their associated user ID and scopes,
or `null` if the token is invalid.
The method is `POST` so a longer list of tokens can be sent in the request body.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OAuthAPITokenTestRequest
*/
func (a *OAuthAPIService) TokenTest(ctx context.Context) OAuthAPITokenTestRequest {
	return OAuthAPITokenTestRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]TokenTest200ResponseValue
func (a *OAuthAPIService) TokenTestExecute(r OAuthAPITokenTestRequest) (*map[string]TokenTest200ResponseValue, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *map[string]TokenTest200ResponseValue
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OAuthAPIService.TokenTest")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/token/test"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
