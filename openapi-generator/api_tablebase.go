/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.165
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


type TablebaseAPI interface {

	/*
	AntichessAtomic Tablebase lookup for Antichess

	**Endpoint: <https://tablebase.lichess.org>**


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TablebaseAPIAntichessAtomicRequest
	*/
	AntichessAtomic(ctx context.Context) TablebaseAPIAntichessAtomicRequest

	// AntichessAtomicExecute executes the request
	//  @return TablebaseJson
	AntichessAtomicExecute(r TablebaseAPIAntichessAtomicRequest) (*TablebaseJson, *http.Response, error)

	/*
	TablebaseAtomic Tablebase lookup for Atomic chess

	**Endpoint: <https://tablebase.lichess.org>**


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TablebaseAPITablebaseAtomicRequest
	*/
	TablebaseAtomic(ctx context.Context) TablebaseAPITablebaseAtomicRequest

	// TablebaseAtomicExecute executes the request
	//  @return TablebaseJson
	TablebaseAtomicExecute(r TablebaseAPITablebaseAtomicRequest) (*TablebaseJson, *http.Response, error)

	/*
	TablebaseStandard Tablebase lookup

	**Endpoint: <https://tablebase.lichess.org>**

Example: `curl http://tablebase.lichess.org/standard?fen=4k3/6KP/8/8/8/8/7p/8_w_-_-_0_1`


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TablebaseAPITablebaseStandardRequest
	*/
	TablebaseStandard(ctx context.Context) TablebaseAPITablebaseStandardRequest

	// TablebaseStandardExecute executes the request
	//  @return TablebaseJson
	TablebaseStandardExecute(r TablebaseAPITablebaseStandardRequest) (*TablebaseJson, *http.Response, error)
}

// TablebaseAPIService TablebaseAPI service
type TablebaseAPIService service

type TablebaseAPIAntichessAtomicRequest struct {
	ctx context.Context
	ApiService TablebaseAPI
	fen *string
}

// X-FEN of the position. Underscores allowed.
func (r TablebaseAPIAntichessAtomicRequest) Fen(fen string) TablebaseAPIAntichessAtomicRequest {
	r.fen = &fen
	return r
}

func (r TablebaseAPIAntichessAtomicRequest) Execute() (*TablebaseJson, *http.Response, error) {
	return r.ApiService.AntichessAtomicExecute(r)
}

/*
AntichessAtomic Tablebase lookup for Antichess

**Endpoint: <https://tablebase.lichess.org>**


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TablebaseAPIAntichessAtomicRequest
*/
func (a *TablebaseAPIService) AntichessAtomic(ctx context.Context) TablebaseAPIAntichessAtomicRequest {
	return TablebaseAPIAntichessAtomicRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TablebaseJson
func (a *TablebaseAPIService) AntichessAtomicExecute(r TablebaseAPIAntichessAtomicRequest) (*TablebaseJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TablebaseJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TablebaseAPIService.AntichessAtomic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/antichess"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fen == nil {
		return localVarReturnValue, nil, reportError("fen is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
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

type TablebaseAPITablebaseAtomicRequest struct {
	ctx context.Context
	ApiService TablebaseAPI
	fen *string
}

// X-FEN of the position. Underscores allowed.
func (r TablebaseAPITablebaseAtomicRequest) Fen(fen string) TablebaseAPITablebaseAtomicRequest {
	r.fen = &fen
	return r
}

func (r TablebaseAPITablebaseAtomicRequest) Execute() (*TablebaseJson, *http.Response, error) {
	return r.ApiService.TablebaseAtomicExecute(r)
}

/*
TablebaseAtomic Tablebase lookup for Atomic chess

**Endpoint: <https://tablebase.lichess.org>**


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TablebaseAPITablebaseAtomicRequest
*/
func (a *TablebaseAPIService) TablebaseAtomic(ctx context.Context) TablebaseAPITablebaseAtomicRequest {
	return TablebaseAPITablebaseAtomicRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TablebaseJson
func (a *TablebaseAPIService) TablebaseAtomicExecute(r TablebaseAPITablebaseAtomicRequest) (*TablebaseJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TablebaseJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TablebaseAPIService.TablebaseAtomic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/atomic"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fen == nil {
		return localVarReturnValue, nil, reportError("fen is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
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

type TablebaseAPITablebaseStandardRequest struct {
	ctx context.Context
	ApiService TablebaseAPI
	fen *string
	dtc *string
}

// X-FEN of the position. Underscores allowed.
func (r TablebaseAPITablebaseStandardRequest) Fen(fen string) TablebaseAPITablebaseStandardRequest {
	r.fen = &fen
	return r
}

// When to query the tablebase for &#x60;dtc&#x60; values. The default is &#x60;auxiliary&#x60;, i.e., only when the position is not covered by one of the other tablebases. 
func (r TablebaseAPITablebaseStandardRequest) Dtc(dtc string) TablebaseAPITablebaseStandardRequest {
	r.dtc = &dtc
	return r
}

func (r TablebaseAPITablebaseStandardRequest) Execute() (*TablebaseJson, *http.Response, error) {
	return r.ApiService.TablebaseStandardExecute(r)
}

/*
TablebaseStandard Tablebase lookup

**Endpoint: <https://tablebase.lichess.org>**

Example: `curl http://tablebase.lichess.org/standard?fen=4k3/6KP/8/8/8/8/7p/8_w_-_-_0_1`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TablebaseAPITablebaseStandardRequest
*/
func (a *TablebaseAPIService) TablebaseStandard(ctx context.Context) TablebaseAPITablebaseStandardRequest {
	return TablebaseAPITablebaseStandardRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TablebaseJson
func (a *TablebaseAPIService) TablebaseStandardExecute(r TablebaseAPITablebaseStandardRequest) (*TablebaseJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TablebaseJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TablebaseAPIService.TablebaseStandard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/standard"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.fen == nil {
		return localVarReturnValue, nil, reportError("fen is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
	if r.dtc != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "dtc", r.dtc, "form", "")
	}
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
