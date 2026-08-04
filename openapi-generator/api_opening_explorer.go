/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.161
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
	"reflect"
)


type OpeningExplorerAPI interface {

	/*
	OpeningExplorerLichess Lichess games

	**Endpoint: <https://explorer.lichess.org/lichess>**

Aggregated rated games from all Lichess players.

Example: `curl https://explorer.lichess.org/lichess?variant=standard&speeds=blitz,rapid,classical&ratings=2200,2500&fen=rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR%20w%20KQkq%20-%200%201`


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OpeningExplorerAPIOpeningExplorerLichessRequest
	*/
	OpeningExplorerLichess(ctx context.Context) OpeningExplorerAPIOpeningExplorerLichessRequest

	// OpeningExplorerLichessExecute executes the request
	//  @return OpeningExplorerLichess
	OpeningExplorerLichessExecute(r OpeningExplorerAPIOpeningExplorerLichessRequest) (*OpeningExplorerLichess, *http.Response, error)

	/*
	OpeningExplorerMaster Masters database

	**Endpoint: <https://explorer.lichess.org/masters>**

Example: `curl https://explorer.lichess.org/masters?play=d2d4,d7d5,c2c4,c7c6,c4d5`


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OpeningExplorerAPIOpeningExplorerMasterRequest
	*/
	OpeningExplorerMaster(ctx context.Context) OpeningExplorerAPIOpeningExplorerMasterRequest

	// OpeningExplorerMasterExecute executes the request
	//  @return OpeningExplorerMasters
	OpeningExplorerMasterExecute(r OpeningExplorerAPIOpeningExplorerMasterRequest) (*OpeningExplorerMasters, *http.Response, error)

	/*
	OpeningExplorerMasterGame OTB master game

	**Endpoint: `https://explorer.lichess.org/masters/pgn/{gameId}`**

Example: `curl https://explorer.lichess.org/masters/pgn/aAbqI4ey`


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return OpeningExplorerAPIOpeningExplorerMasterGameRequest
	*/
	OpeningExplorerMasterGame(ctx context.Context, gameId string) OpeningExplorerAPIOpeningExplorerMasterGameRequest

	// OpeningExplorerMasterGameExecute executes the request
	//  @return string
	OpeningExplorerMasterGameExecute(r OpeningExplorerAPIOpeningExplorerMasterGameRequest) (string, *http.Response, error)

	/*
	OpeningExplorerPlayer Player games

	**Endpoint: <https://explorer.lichess.org/player>**

Games of a Lichess player.

Responds with a stream of [newline delimited JSON](#description/streaming-with-nd-json). Will start indexing
on demand, immediately respond with the current results, and stream
more updates until indexing is complete. The stream is throttled
and deduplicated. Empty lines may be sent to avoid timeouts.

Will index new games at most once per minute, and revisit previously
ongoing games at most once every day.

Example: `curl https://explorer.lichess.org/player?player=revoof&color=white&play=d2d4,d7d5&recentGames=1`


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return OpeningExplorerAPIOpeningExplorerPlayerRequest
	*/
	OpeningExplorerPlayer(ctx context.Context) OpeningExplorerAPIOpeningExplorerPlayerRequest

	// OpeningExplorerPlayerExecute executes the request
	//  @return OpeningExplorerPlayer
	OpeningExplorerPlayerExecute(r OpeningExplorerAPIOpeningExplorerPlayerRequest) (*OpeningExplorerPlayer, *http.Response, error)
}

// OpeningExplorerAPIService OpeningExplorerAPI service
type OpeningExplorerAPIService service

type OpeningExplorerAPIOpeningExplorerLichessRequest struct {
	ctx context.Context
	ApiService OpeningExplorerAPI
	variant *VariantKey
	fen *string
	play *string
	speeds *[]Speed
	ratings *[]int32
	since *string
	until *string
	moves *int32
	topGames *int32
	recentGames *int32
	history *bool
}

// Variant
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Variant(variant VariantKey) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.variant = &variant
	return r
}

// X-FEN or EPD of the root position
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Fen(fen string) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.fen = &fen
	return r
}

// Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position. 
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Play(play string) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.play = &play
	return r
}

// Comma separated list of game speeds to filter by
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Speeds(speeds []Speed) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.speeds = &speeds
	return r
}

// Comma separated list of ratings groups to filter by. Each group ranges from its value to the next higher group in the enum (&#x60;0&#x60; from 0 to 999, &#x60;1000&#x60; from 1000 to 1199, ..., &#x60;2500&#x60; from 2500 to any rating above). 
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Ratings(ratings []int32) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.ratings = &ratings
	return r
}

// Include only games from this month or later
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Since(since string) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.since = &since
	return r
}

// Include only games from this month or earlier
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Until(until string) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.until = &until
	return r
}

// Number of most common moves to display
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Moves(moves int32) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.moves = &moves
	return r
}

// Maximum number of top games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected. 
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) TopGames(topGames int32) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.topGames = &topGames
	return r
}

// Maximum number of recent games to display.  Due to the way banned users are handled internally, the response may contain fewer games than expected. 
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) RecentGames(recentGames int32) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.recentGames = &recentGames
	return r
}

// Optionally retrieve history
func (r OpeningExplorerAPIOpeningExplorerLichessRequest) History(history bool) OpeningExplorerAPIOpeningExplorerLichessRequest {
	r.history = &history
	return r
}

func (r OpeningExplorerAPIOpeningExplorerLichessRequest) Execute() (*OpeningExplorerLichess, *http.Response, error) {
	return r.ApiService.OpeningExplorerLichessExecute(r)
}

/*
OpeningExplorerLichess Lichess games

**Endpoint: <https://explorer.lichess.org/lichess>**

Aggregated rated games from all Lichess players.

Example: `curl https://explorer.lichess.org/lichess?variant=standard&speeds=blitz,rapid,classical&ratings=2200,2500&fen=rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR%20w%20KQkq%20-%200%201`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpeningExplorerAPIOpeningExplorerLichessRequest
*/
func (a *OpeningExplorerAPIService) OpeningExplorerLichess(ctx context.Context) OpeningExplorerAPIOpeningExplorerLichessRequest {
	return OpeningExplorerAPIOpeningExplorerLichessRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpeningExplorerLichess
func (a *OpeningExplorerAPIService) OpeningExplorerLichessExecute(r OpeningExplorerAPIOpeningExplorerLichessRequest) (*OpeningExplorerLichess, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpeningExplorerLichess
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpeningExplorerAPIService.OpeningExplorerLichess")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/lichess"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variant", r.variant, "form", "")
	} else {
		var defaultValue VariantKey = "standard"
		parameterAddToHeaderOrQuery(localVarQueryParams, "variant", defaultValue, "form", "")
		r.variant = &defaultValue
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
	}
	if r.play != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", r.play, "form", "")
	} else {
		var defaultValue string = ""
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", defaultValue, "form", "")
		r.play = &defaultValue
	}
	if r.speeds != nil {
		t := *r.speeds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "speeds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "speeds", t, "form", "multi")
		}
	}
	if r.ratings != nil {
		t := *r.ratings
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "ratings", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "ratings", t, "form", "multi")
		}
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	} else {
		var defaultValue string = "1952-01"
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", defaultValue, "form", "")
		r.since = &defaultValue
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "form", "")
	} else {
		var defaultValue string = "3000-12"
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", defaultValue, "form", "")
		r.until = &defaultValue
	}
	if r.moves != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", r.moves, "form", "")
	} else {
		var defaultValue int32 = 12
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", defaultValue, "form", "")
		r.moves = &defaultValue
	}
	if r.topGames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "topGames", r.topGames, "form", "")
	} else {
		var defaultValue int32 = 4
		parameterAddToHeaderOrQuery(localVarQueryParams, "topGames", defaultValue, "form", "")
		r.topGames = &defaultValue
	}
	if r.recentGames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recentGames", r.recentGames, "form", "")
	} else {
		var defaultValue int32 = 4
		parameterAddToHeaderOrQuery(localVarQueryParams, "recentGames", defaultValue, "form", "")
		r.recentGames = &defaultValue
	}
	if r.history != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "history", r.history, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "history", defaultValue, "form", "")
		r.history = &defaultValue
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

type OpeningExplorerAPIOpeningExplorerMasterRequest struct {
	ctx context.Context
	ApiService OpeningExplorerAPI
	fen *string
	play *string
	since *int32
	until *int32
	moves *int32
	topGames *int32
}

// X-FEN of the root position
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Fen(fen string) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.fen = &fen
	return r
}

// Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position. 
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Play(play string) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.play = &play
	return r
}

// Include only games from this year or later
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Since(since int32) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.since = &since
	return r
}

// Include only games from this year or earlier
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Until(until int32) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.until = &until
	return r
}

// Number of most common moves to display
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Moves(moves int32) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.moves = &moves
	return r
}

// Number of top games to display
func (r OpeningExplorerAPIOpeningExplorerMasterRequest) TopGames(topGames int32) OpeningExplorerAPIOpeningExplorerMasterRequest {
	r.topGames = &topGames
	return r
}

func (r OpeningExplorerAPIOpeningExplorerMasterRequest) Execute() (*OpeningExplorerMasters, *http.Response, error) {
	return r.ApiService.OpeningExplorerMasterExecute(r)
}

/*
OpeningExplorerMaster Masters database

**Endpoint: <https://explorer.lichess.org/masters>**

Example: `curl https://explorer.lichess.org/masters?play=d2d4,d7d5,c2c4,c7c6,c4d5`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpeningExplorerAPIOpeningExplorerMasterRequest
*/
func (a *OpeningExplorerAPIService) OpeningExplorerMaster(ctx context.Context) OpeningExplorerAPIOpeningExplorerMasterRequest {
	return OpeningExplorerAPIOpeningExplorerMasterRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpeningExplorerMasters
func (a *OpeningExplorerAPIService) OpeningExplorerMasterExecute(r OpeningExplorerAPIOpeningExplorerMasterRequest) (*OpeningExplorerMasters, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpeningExplorerMasters
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpeningExplorerAPIService.OpeningExplorerMaster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/masters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
	}
	if r.play != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", r.play, "form", "")
	} else {
		var defaultValue string = ""
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", defaultValue, "form", "")
		r.play = &defaultValue
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	} else {
		var defaultValue int32 = 1952
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", defaultValue, "form", "")
		r.since = &defaultValue
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "form", "")
	}
	if r.moves != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", r.moves, "form", "")
	} else {
		var defaultValue int32 = 12
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", defaultValue, "form", "")
		r.moves = &defaultValue
	}
	if r.topGames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "topGames", r.topGames, "form", "")
	} else {
		var defaultValue int32 = 15
		parameterAddToHeaderOrQuery(localVarQueryParams, "topGames", defaultValue, "form", "")
		r.topGames = &defaultValue
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

type OpeningExplorerAPIOpeningExplorerMasterGameRequest struct {
	ctx context.Context
	ApiService OpeningExplorerAPI
	gameId string
}

func (r OpeningExplorerAPIOpeningExplorerMasterGameRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.OpeningExplorerMasterGameExecute(r)
}

/*
OpeningExplorerMasterGame OTB master game

**Endpoint: `https://explorer.lichess.org/masters/pgn/{gameId}`**

Example: `curl https://explorer.lichess.org/masters/pgn/aAbqI4ey`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return OpeningExplorerAPIOpeningExplorerMasterGameRequest
*/
func (a *OpeningExplorerAPIService) OpeningExplorerMasterGame(ctx context.Context, gameId string) OpeningExplorerAPIOpeningExplorerMasterGameRequest {
	return OpeningExplorerAPIOpeningExplorerMasterGameRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return string
func (a *OpeningExplorerAPIService) OpeningExplorerMasterGameExecute(r OpeningExplorerAPIOpeningExplorerMasterGameRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpeningExplorerAPIService.OpeningExplorerMasterGame")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/masters/pgn/{gameId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-chess-pgn"}

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

type OpeningExplorerAPIOpeningExplorerPlayerRequest struct {
	ctx context.Context
	ApiService OpeningExplorerAPI
	player *string
	color *string
	variant *VariantKey
	fen *string
	play *string
	speeds *[]Speed
	modes *[]string
	since *string
	until *string
	moves *int32
	recentGames *int32
}

// Username or ID of the player
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Player(player string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.player = &player
	return r
}

// Look for games with *player* on the given side
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Color(color string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.color = &color
	return r
}

// Variant
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Variant(variant VariantKey) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.variant = &variant
	return r
}

// X-FEN of the root position
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Fen(fen string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.fen = &fen
	return r
}

// Comma separated sequence of legal moves in UCI notation. Play additional moves starting from &#x60;fen&#x60;. Required to find an opening name, if &#x60;fen&#x60; is not an exact match for a named position. 
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Play(play string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.play = &play
	return r
}

// Comma separated list of game speeds to look for
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Speeds(speeds []Speed) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.speeds = &speeds
	return r
}

// Comma separated list of modes
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Modes(modes []string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.modes = &modes
	return r
}

// Include only games from this month or later
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Since(since string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.since = &since
	return r
}

// Include only games from this month or earlier
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Until(until string) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.until = &until
	return r
}

// Number of most common moves to display
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Moves(moves int32) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.moves = &moves
	return r
}

// Number of recent games to display
func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) RecentGames(recentGames int32) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	r.recentGames = &recentGames
	return r
}

func (r OpeningExplorerAPIOpeningExplorerPlayerRequest) Execute() (*OpeningExplorerPlayer, *http.Response, error) {
	return r.ApiService.OpeningExplorerPlayerExecute(r)
}

/*
OpeningExplorerPlayer Player games

**Endpoint: <https://explorer.lichess.org/player>**

Games of a Lichess player.

Responds with a stream of [newline delimited JSON](#description/streaming-with-nd-json). Will start indexing
on demand, immediately respond with the current results, and stream
more updates until indexing is complete. The stream is throttled
and deduplicated. Empty lines may be sent to avoid timeouts.

Will index new games at most once per minute, and revisit previously
ongoing games at most once every day.

Example: `curl https://explorer.lichess.org/player?player=revoof&color=white&play=d2d4,d7d5&recentGames=1`


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return OpeningExplorerAPIOpeningExplorerPlayerRequest
*/
func (a *OpeningExplorerAPIService) OpeningExplorerPlayer(ctx context.Context) OpeningExplorerAPIOpeningExplorerPlayerRequest {
	return OpeningExplorerAPIOpeningExplorerPlayerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return OpeningExplorerPlayer
func (a *OpeningExplorerAPIService) OpeningExplorerPlayerExecute(r OpeningExplorerAPIOpeningExplorerPlayerRequest) (*OpeningExplorerPlayer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *OpeningExplorerPlayer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "OpeningExplorerAPIService.OpeningExplorerPlayer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/player"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.player == nil {
		return localVarReturnValue, nil, reportError("player is required and must be specified")
	}
	if r.color == nil {
		return localVarReturnValue, nil, reportError("color is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "player", r.player, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "color", r.color, "form", "")
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "variant", r.variant, "form", "")
	} else {
		var defaultValue VariantKey = "standard"
		parameterAddToHeaderOrQuery(localVarQueryParams, "variant", defaultValue, "form", "")
		r.variant = &defaultValue
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fen", r.fen, "form", "")
	}
	if r.play != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", r.play, "form", "")
	} else {
		var defaultValue string = ""
		parameterAddToHeaderOrQuery(localVarQueryParams, "play", defaultValue, "form", "")
		r.play = &defaultValue
	}
	if r.speeds != nil {
		t := *r.speeds
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "speeds", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "speeds", t, "form", "multi")
		}
	}
	if r.modes != nil {
		t := *r.modes
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "modes", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "modes", t, "form", "multi")
		}
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	} else {
		var defaultValue string = "1952-01"
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", defaultValue, "form", "")
		r.since = &defaultValue
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "form", "")
	} else {
		var defaultValue string = "3000-12"
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", defaultValue, "form", "")
		r.until = &defaultValue
	}
	if r.moves != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "moves", r.moves, "form", "")
	}
	if r.recentGames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "recentGames", r.recentGames, "form", "")
	} else {
		var defaultValue int32 = 8
		parameterAddToHeaderOrQuery(localVarQueryParams, "recentGames", defaultValue, "form", "")
		r.recentGames = &defaultValue
	}
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
