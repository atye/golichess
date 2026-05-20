/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
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
	"strings"
)


type ExternalEngineAPI interface {

	/*
	ApiExternalEngineAcquire Acquire analysis request

	**Endpoint: `https://engine.lichess.ovh/api/external-engine/work`**
Wait for an analysis requests to any of the external engines that
have been registered with the given `secret`.
Uses long polling.
After acquiring a request, the provider should immediately
[start streaming the results](#tag/external-engine/POST/api/external-engine/work/{id}).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExternalEngineAPIApiExternalEngineAcquireRequest
	*/
	ApiExternalEngineAcquire(ctx context.Context) ExternalEngineAPIApiExternalEngineAcquireRequest

	// ApiExternalEngineAcquireExecute executes the request
	//  @return ApiExternalEngineAcquire200Response
	ApiExternalEngineAcquireExecute(r ExternalEngineAPIApiExternalEngineAcquireRequest) (*ApiExternalEngineAcquire200Response, *http.Response, error)

	/*
	ApiExternalEngineAnalyse Analyse with external engine

	**Endpoint: `https://engine.lichess.ovh/api/external-engine/{id}/analyse`**
Request analysis from an external engine.
Response content is streamed as [newline delimited JSON](#description/streaming-with-nd-json).
The properties are based on the [UCI specification](https://backscattering.de/chess/uci/#engine).
Analysis stops when the client goes away, the requested limit
is reached, or the provider goes away.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The external engine id.
	@return ExternalEngineAPIApiExternalEngineAnalyseRequest
	*/
	ApiExternalEngineAnalyse(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineAnalyseRequest

	// ApiExternalEngineAnalyseExecute executes the request
	//  @return ApiExternalEngineAnalyse200Response
	ApiExternalEngineAnalyseExecute(r ExternalEngineAPIApiExternalEngineAnalyseRequest) (*ApiExternalEngineAnalyse200Response, *http.Response, error)

	/*
	ApiExternalEngineCreate Create external engine

	Registers a new external engine for the user. It can then be selected
and used on the analysis board.
After registering, the provider should start waiting for analyis requests.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExternalEngineAPIApiExternalEngineCreateRequest
	*/
	ApiExternalEngineCreate(ctx context.Context) ExternalEngineAPIApiExternalEngineCreateRequest

	// ApiExternalEngineCreateExecute executes the request
	//  @return ApiExternalEngineList200ResponseInner
	ApiExternalEngineCreateExecute(r ExternalEngineAPIApiExternalEngineCreateRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error)

	/*
	ApiExternalEngineDelete Delete external engine

	Unregisters an external engine.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The external engine id.
	@return ExternalEngineAPIApiExternalEngineDeleteRequest
	*/
	ApiExternalEngineDelete(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineDeleteRequest

	// ApiExternalEngineDeleteExecute executes the request
	//  @return AccountKidPost200Response
	ApiExternalEngineDeleteExecute(r ExternalEngineAPIApiExternalEngineDeleteRequest) (*AccountKidPost200Response, *http.Response, error)

	/*
	ApiExternalEngineGet Get external engine

	Get properties and credentials of an external engine.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The external engine id.
	@return ExternalEngineAPIApiExternalEngineGetRequest
	*/
	ApiExternalEngineGet(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineGetRequest

	// ApiExternalEngineGetExecute executes the request
	//  @return ApiExternalEngineList200ResponseInner
	ApiExternalEngineGetExecute(r ExternalEngineAPIApiExternalEngineGetRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error)

	/*
	ApiExternalEngineList List external engines

	Lists all external engines that have been registered for the user,
and the credentials required to use them.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ExternalEngineAPIApiExternalEngineListRequest
	*/
	ApiExternalEngineList(ctx context.Context) ExternalEngineAPIApiExternalEngineListRequest

	// ApiExternalEngineListExecute executes the request
	//  @return []ApiExternalEngineList200ResponseInner
	ApiExternalEngineListExecute(r ExternalEngineAPIApiExternalEngineListRequest) ([]ApiExternalEngineList200ResponseInner, *http.Response, error)

	/*
	ApiExternalEnginePut Update external engine

	Updates the properties of an external engine.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The external engine id.
	@return ExternalEngineAPIApiExternalEnginePutRequest
	*/
	ApiExternalEnginePut(ctx context.Context, id string) ExternalEngineAPIApiExternalEnginePutRequest

	// ApiExternalEnginePutExecute executes the request
	//  @return ApiExternalEngineList200ResponseInner
	ApiExternalEnginePutExecute(r ExternalEngineAPIApiExternalEnginePutRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error)

	/*
	ApiExternalEngineSubmit Answer analysis request

	**Endpoint: `https://engine.lichess.ovh/api/external-engine/work/{id}`**
Submit a stream of analysis as [UCI output](https://backscattering.de/chess/uci/#engine-info).
* The engine should always be in `UCI_Chess960` mode.
* `UCI_AnalyseMode` enabled if available.
* It produces `info` with at least:
  - `depth`
  - `multipv` (between 1 and 5)
  - `score`
  - `nodes`
  - `time`
  - `pv`
The server may close the connection at any time, indicating that
the requester has gone away and analysis should be stopped.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return ExternalEngineAPIApiExternalEngineSubmitRequest
	*/
	ApiExternalEngineSubmit(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineSubmitRequest

	// ApiExternalEngineSubmitExecute executes the request
	ApiExternalEngineSubmitExecute(r ExternalEngineAPIApiExternalEngineSubmitRequest) (*http.Response, error)
}

// ExternalEngineAPIService ExternalEngineAPI service
type ExternalEngineAPIService service

type ExternalEngineAPIApiExternalEngineAcquireRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	apiExternalEngineAcquireRequest *ApiExternalEngineAcquireRequest
}

// Provider credentials.
func (r ExternalEngineAPIApiExternalEngineAcquireRequest) ApiExternalEngineAcquireRequest(apiExternalEngineAcquireRequest ApiExternalEngineAcquireRequest) ExternalEngineAPIApiExternalEngineAcquireRequest {
	r.apiExternalEngineAcquireRequest = &apiExternalEngineAcquireRequest
	return r
}

func (r ExternalEngineAPIApiExternalEngineAcquireRequest) Execute() (*ApiExternalEngineAcquire200Response, *http.Response, error) {
	return r.ApiService.ApiExternalEngineAcquireExecute(r)
}

/*
ApiExternalEngineAcquire Acquire analysis request

**Endpoint: `https://engine.lichess.ovh/api/external-engine/work`**
Wait for an analysis requests to any of the external engines that
have been registered with the given `secret`.
Uses long polling.
After acquiring a request, the provider should immediately
[start streaming the results](#tag/external-engine/POST/api/external-engine/work/{id}).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExternalEngineAPIApiExternalEngineAcquireRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineAcquire(ctx context.Context) ExternalEngineAPIApiExternalEngineAcquireRequest {
	return ExternalEngineAPIApiExternalEngineAcquireRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiExternalEngineAcquire200Response
func (a *ExternalEngineAPIService) ApiExternalEngineAcquireExecute(r ExternalEngineAPIApiExternalEngineAcquireRequest) (*ApiExternalEngineAcquire200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiExternalEngineAcquire200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineAcquire")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/work"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiExternalEngineAcquireRequest == nil {
		return localVarReturnValue, nil, reportError("apiExternalEngineAcquireRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiExternalEngineAcquireRequest
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

type ExternalEngineAPIApiExternalEngineAnalyseRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	id string
	apiExternalEngineAnalyseRequest *ApiExternalEngineAnalyseRequest
}

// Engine credentials and analysis request.
func (r ExternalEngineAPIApiExternalEngineAnalyseRequest) ApiExternalEngineAnalyseRequest(apiExternalEngineAnalyseRequest ApiExternalEngineAnalyseRequest) ExternalEngineAPIApiExternalEngineAnalyseRequest {
	r.apiExternalEngineAnalyseRequest = &apiExternalEngineAnalyseRequest
	return r
}

func (r ExternalEngineAPIApiExternalEngineAnalyseRequest) Execute() (*ApiExternalEngineAnalyse200Response, *http.Response, error) {
	return r.ApiService.ApiExternalEngineAnalyseExecute(r)
}

/*
ApiExternalEngineAnalyse Analyse with external engine

**Endpoint: `https://engine.lichess.ovh/api/external-engine/{id}/analyse`**
Request analysis from an external engine.
Response content is streamed as [newline delimited JSON](#description/streaming-with-nd-json).
The properties are based on the [UCI specification](https://backscattering.de/chess/uci/#engine).
Analysis stops when the client goes away, the requested limit
is reached, or the provider goes away.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The external engine id.
 @return ExternalEngineAPIApiExternalEngineAnalyseRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineAnalyse(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineAnalyseRequest {
	return ExternalEngineAPIApiExternalEngineAnalyseRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApiExternalEngineAnalyse200Response
func (a *ExternalEngineAPIService) ApiExternalEngineAnalyseExecute(r ExternalEngineAPIApiExternalEngineAnalyseRequest) (*ApiExternalEngineAnalyse200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiExternalEngineAnalyse200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineAnalyse")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/{id}/analyse"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiExternalEngineAnalyseRequest == nil {
		return localVarReturnValue, nil, reportError("apiExternalEngineAnalyseRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.apiExternalEngineAnalyseRequest
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

type ExternalEngineAPIApiExternalEngineCreateRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	apiExternalEngineCreateRequest *ApiExternalEngineCreateRequest
}

// A new external engine registration.
func (r ExternalEngineAPIApiExternalEngineCreateRequest) ApiExternalEngineCreateRequest(apiExternalEngineCreateRequest ApiExternalEngineCreateRequest) ExternalEngineAPIApiExternalEngineCreateRequest {
	r.apiExternalEngineCreateRequest = &apiExternalEngineCreateRequest
	return r
}

func (r ExternalEngineAPIApiExternalEngineCreateRequest) Execute() (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiExternalEngineCreateExecute(r)
}

/*
ApiExternalEngineCreate Create external engine

Registers a new external engine for the user. It can then be selected
and used on the analysis board.
After registering, the provider should start waiting for analyis requests.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExternalEngineAPIApiExternalEngineCreateRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineCreate(ctx context.Context) ExternalEngineAPIApiExternalEngineCreateRequest {
	return ExternalEngineAPIApiExternalEngineCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiExternalEngineList200ResponseInner
func (a *ExternalEngineAPIService) ApiExternalEngineCreateExecute(r ExternalEngineAPIApiExternalEngineCreateRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiExternalEngineList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiExternalEngineCreateRequest == nil {
		return localVarReturnValue, nil, reportError("apiExternalEngineCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiExternalEngineCreateRequest
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

type ExternalEngineAPIApiExternalEngineDeleteRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	id string
}

func (r ExternalEngineAPIApiExternalEngineDeleteRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.ApiExternalEngineDeleteExecute(r)
}

/*
ApiExternalEngineDelete Delete external engine

Unregisters an external engine.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The external engine id.
 @return ExternalEngineAPIApiExternalEngineDeleteRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineDelete(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineDeleteRequest {
	return ExternalEngineAPIApiExternalEngineDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *ExternalEngineAPIService) ApiExternalEngineDeleteExecute(r ExternalEngineAPIApiExternalEngineDeleteRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ExternalEngineAPIApiExternalEngineGetRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	id string
}

func (r ExternalEngineAPIApiExternalEngineGetRequest) Execute() (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiExternalEngineGetExecute(r)
}

/*
ApiExternalEngineGet Get external engine

Get properties and credentials of an external engine.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The external engine id.
 @return ExternalEngineAPIApiExternalEngineGetRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineGet(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineGetRequest {
	return ExternalEngineAPIApiExternalEngineGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApiExternalEngineList200ResponseInner
func (a *ExternalEngineAPIService) ApiExternalEngineGetExecute(r ExternalEngineAPIApiExternalEngineGetRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiExternalEngineList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ExternalEngineAPIApiExternalEngineListRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
}

func (r ExternalEngineAPIApiExternalEngineListRequest) Execute() ([]ApiExternalEngineList200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiExternalEngineListExecute(r)
}

/*
ApiExternalEngineList List external engines

Lists all external engines that have been registered for the user,
and the credentials required to use them.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ExternalEngineAPIApiExternalEngineListRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineList(ctx context.Context) ExternalEngineAPIApiExternalEngineListRequest {
	return ExternalEngineAPIApiExternalEngineListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiExternalEngineList200ResponseInner
func (a *ExternalEngineAPIService) ApiExternalEngineListExecute(r ExternalEngineAPIApiExternalEngineListRequest) ([]ApiExternalEngineList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiExternalEngineList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine"

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
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
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

type ExternalEngineAPIApiExternalEnginePutRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	id string
	apiExternalEngineCreateRequest *ApiExternalEngineCreateRequest
}

// A modified engine registration.
func (r ExternalEngineAPIApiExternalEnginePutRequest) ApiExternalEngineCreateRequest(apiExternalEngineCreateRequest ApiExternalEngineCreateRequest) ExternalEngineAPIApiExternalEnginePutRequest {
	r.apiExternalEngineCreateRequest = &apiExternalEngineCreateRequest
	return r
}

func (r ExternalEngineAPIApiExternalEnginePutRequest) Execute() (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiExternalEnginePutExecute(r)
}

/*
ApiExternalEnginePut Update external engine

Updates the properties of an external engine.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The external engine id.
 @return ExternalEngineAPIApiExternalEnginePutRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEnginePut(ctx context.Context, id string) ExternalEngineAPIApiExternalEnginePutRequest {
	return ExternalEngineAPIApiExternalEnginePutRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ApiExternalEngineList200ResponseInner
func (a *ExternalEngineAPIService) ApiExternalEnginePutExecute(r ExternalEngineAPIApiExternalEnginePutRequest) (*ApiExternalEngineList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiExternalEngineList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEnginePut")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.apiExternalEngineCreateRequest == nil {
		return localVarReturnValue, nil, reportError("apiExternalEngineCreateRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.apiExternalEngineCreateRequest
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

type ExternalEngineAPIApiExternalEngineSubmitRequest struct {
	ctx context.Context
	ApiService ExternalEngineAPI
	id string
	body *string
}

// Analysis results
func (r ExternalEngineAPIApiExternalEngineSubmitRequest) Body(body string) ExternalEngineAPIApiExternalEngineSubmitRequest {
	r.body = &body
	return r
}

func (r ExternalEngineAPIApiExternalEngineSubmitRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiExternalEngineSubmitExecute(r)
}

/*
ApiExternalEngineSubmit Answer analysis request

**Endpoint: `https://engine.lichess.ovh/api/external-engine/work/{id}`**
Submit a stream of analysis as [UCI output](https://backscattering.de/chess/uci/#engine-info).
* The engine should always be in `UCI_Chess960` mode.
* `UCI_AnalyseMode` enabled if available.
* It produces `info` with at least:
  - `depth`
  - `multipv` (between 1 and 5)
  - `score`
  - `nodes`
  - `time`
  - `pv`
The server may close the connection at any time, indicating that
the requester has gone away and analysis should be stopped.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ExternalEngineAPIApiExternalEngineSubmitRequest
*/
func (a *ExternalEngineAPIService) ApiExternalEngineSubmit(ctx context.Context, id string) ExternalEngineAPIApiExternalEngineSubmitRequest {
	return ExternalEngineAPIApiExternalEngineSubmitRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *ExternalEngineAPIService) ApiExternalEngineSubmitExecute(r ExternalEngineAPIApiExternalEngineSubmitRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ExternalEngineAPIService.ApiExternalEngineSubmit")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/external-engine/work/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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
	// body params
	localVarPostBody = r.body
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
