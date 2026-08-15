/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.163
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


type TVAPI interface {

	/*
	TvChannelFeed Stream current TV game of a TV channel

	Stream positions and moves of a current [TV channel's game](https://lichess.org/tv/rapid) in [ndjson](#description/streaming-with-nd-json).
Try it with `curl https://lichess.org/api/tv/rapid/feed`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channel The name of the channel in camel case.
	@return TVAPITvChannelFeedRequest
	*/
	TvChannelFeed(ctx context.Context, channel string) TVAPITvChannelFeedRequest

	// TvChannelFeedExecute executes the request
	//  @return TvFeed
	TvChannelFeedExecute(r TVAPITvChannelFeedRequest) (*TvFeed, *http.Response, error)

	/*
	TvChannelGames Get best ongoing games of a TV channel

	Get a list of ongoing games for a given TV channel. Similar to [lichess.org/games](https://lichess.org/games).
Available in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param channel The name of the channel in camel case.
	@return TVAPITvChannelGamesRequest
	*/
	TvChannelGames(ctx context.Context, channel string) TVAPITvChannelGamesRequest

	// TvChannelGamesExecute executes the request
	//  @return string
	TvChannelGamesExecute(r TVAPITvChannelGamesRequest) (string, *http.Response, error)

	/*
	TvChannels Get current TV games

	Get basic info about the best games being played for each speed and variant,
but also computer games and bot games.
See [lichess.org/tv](https://lichess.org/tv).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TVAPITvChannelsRequest
	*/
	TvChannels(ctx context.Context) TVAPITvChannelsRequest

	// TvChannelsExecute executes the request
	//  @return TvChannels200Response
	TvChannelsExecute(r TVAPITvChannelsRequest) (*TvChannels200Response, *http.Response, error)

	/*
	TvFeed Stream current TV game

	Stream positions and moves of the current [TV game](https://lichess.org/tv) in [ndjson](#description/streaming-with-nd-json).
Try it with `curl https://lichess.org/api/tv/feed`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TVAPITvFeedRequest
	*/
	TvFeed(ctx context.Context) TVAPITvFeedRequest

	// TvFeedExecute executes the request
	//  @return TvFeed
	TvFeedExecute(r TVAPITvFeedRequest) (*TvFeed, *http.Response, error)
}

// TVAPIService TVAPI service
type TVAPIService service

type TVAPITvChannelFeedRequest struct {
	ctx context.Context
	ApiService TVAPI
	channel string
}

func (r TVAPITvChannelFeedRequest) Execute() (*TvFeed, *http.Response, error) {
	return r.ApiService.TvChannelFeedExecute(r)
}

/*
TvChannelFeed Stream current TV game of a TV channel

Stream positions and moves of a current [TV channel's game](https://lichess.org/tv/rapid) in [ndjson](#description/streaming-with-nd-json).
Try it with `curl https://lichess.org/api/tv/rapid/feed`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The name of the channel in camel case.
 @return TVAPITvChannelFeedRequest
*/
func (a *TVAPIService) TvChannelFeed(ctx context.Context, channel string) TVAPITvChannelFeedRequest {
	return TVAPITvChannelFeedRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
//  @return TvFeed
func (a *TVAPIService) TvChannelFeedExecute(r TVAPITvChannelFeedRequest) (*TvFeed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TvFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TVAPIService.TvChannelFeed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tv/{channel}/feed"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson"}

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

type TVAPITvChannelGamesRequest struct {
	ctx context.Context
	ApiService TVAPI
	channel string
	accept *string
	nb *int32
	moves *bool
	pgnInJson *bool
	tags *bool
	clocks *bool
	opening *bool
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r TVAPITvChannelGamesRequest) Accept(accept string) TVAPITvChannelGamesRequest {
	r.accept = &accept
	return r
}

// Number of games to fetch.
func (r TVAPITvChannelGamesRequest) Nb(nb int32) TVAPITvChannelGamesRequest {
	r.nb = &nb
	return r
}

// Include the PGN moves.
func (r TVAPITvChannelGamesRequest) Moves(moves bool) TVAPITvChannelGamesRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r TVAPITvChannelGamesRequest) PgnInJson(pgnInJson bool) TVAPITvChannelGamesRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r TVAPITvChannelGamesRequest) Tags(tags bool) TVAPITvChannelGamesRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r TVAPITvChannelGamesRequest) Clocks(clocks bool) TVAPITvChannelGamesRequest {
	r.clocks = &clocks
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r TVAPITvChannelGamesRequest) Opening(opening bool) TVAPITvChannelGamesRequest {
	r.opening = &opening
	return r
}

func (r TVAPITvChannelGamesRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.TvChannelGamesExecute(r)
}

/*
TvChannelGames Get best ongoing games of a TV channel

Get a list of ongoing games for a given TV channel. Similar to [lichess.org/games](https://lichess.org/games).
Available in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param channel The name of the channel in camel case.
 @return TVAPITvChannelGamesRequest
*/
func (a *TVAPIService) TvChannelGames(ctx context.Context, channel string) TVAPITvChannelGamesRequest {
	return TVAPITvChannelGamesRequest{
		ApiService: a,
		ctx: ctx,
		channel: channel,
	}
}

// Execute executes the request
//  @return string
func (a *TVAPIService) TvChannelGamesExecute(r TVAPITvChannelGamesRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TVAPIService.TvChannelGames")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tv/{channel}"
	localVarPath = strings.Replace(localVarPath, "{"+"channel"+"}", url.PathEscape(parameterValueToString(r.channel, "channel")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 10
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", defaultValue, "form", "")
		r.nb = &defaultValue
	}
	if r.moves != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", r.moves, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", defaultValue, "form", "")
		r.moves = &defaultValue
	}
	if r.pgnInJson != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pgnInJson", r.pgnInJson, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "pgnInJson", defaultValue, "form", "")
		r.pgnInJson = &defaultValue
	}
	if r.tags != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", r.tags, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "tags", defaultValue, "form", "")
		r.tags = &defaultValue
	}
	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.opening != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "opening", r.opening, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "opening", defaultValue, "form", "")
		r.opening = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn", "application/x-ndjson"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
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

type TVAPITvChannelsRequest struct {
	ctx context.Context
	ApiService TVAPI
}

func (r TVAPITvChannelsRequest) Execute() (*TvChannels200Response, *http.Response, error) {
	return r.ApiService.TvChannelsExecute(r)
}

/*
TvChannels Get current TV games

Get basic info about the best games being played for each speed and variant,
but also computer games and bot games.
See [lichess.org/tv](https://lichess.org/tv).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TVAPITvChannelsRequest
*/
func (a *TVAPIService) TvChannels(ctx context.Context) TVAPITvChannelsRequest {
	return TVAPITvChannelsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TvChannels200Response
func (a *TVAPIService) TvChannelsExecute(r TVAPITvChannelsRequest) (*TvChannels200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TvChannels200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TVAPIService.TvChannels")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tv/channels"

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

type TVAPITvFeedRequest struct {
	ctx context.Context
	ApiService TVAPI
}

func (r TVAPITvFeedRequest) Execute() (*TvFeed, *http.Response, error) {
	return r.ApiService.TvFeedExecute(r)
}

/*
TvFeed Stream current TV game

Stream positions and moves of the current [TV game](https://lichess.org/tv) in [ndjson](#description/streaming-with-nd-json).
Try it with `curl https://lichess.org/api/tv/feed`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TVAPITvFeedRequest
*/
func (a *TVAPIService) TvFeed(ctx context.Context) TVAPITvFeedRequest {
	return TVAPITvFeedRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TvFeed
func (a *TVAPIService) TvFeedExecute(r TVAPITvFeedRequest) (*TvFeed, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TvFeed
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TVAPIService.TvFeed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tv/feed"

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
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson"}

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
