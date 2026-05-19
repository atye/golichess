/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type BulkPairingsAPI interface {

	/*
	BulkPairingCreate Create a bulk pairing

	Schedule many games at once, up to 24h in advance.
OAuth tokens are required for all paired players, with the `challenge:write` scope.
You can schedule up to 500 games every 10 minutes. [Contact us](mailto:contact@lichess.org) if you need higher limits.
If games have a real-time clock, each player must have only one pairing.
For correspondence games, players can have multiple pairings within the same bulk.

**The entire bulk is rejected if:**
  - a token is missing
  - a token is present more than once (except in correspondence)
  - a token lacks the `challenge:write` scope
  - a player account is closed
  - a player is paired more than once (except in correspondence)
  - a bulk is already scheduled to start at the same time with the same player
  - you have 20 scheduled bulks
  - you have 1000 scheduled games

Partial bulks are never created. Either it all fails, or it all succeeds.
When it fails, it does so with an error message explaining the issue.
Failed bulks are not counted in the rate limiting, they are free.
Fix the issues, manually or programmatically, then retry to schedule the bulk.
A successful bulk creation returns a JSON bulk document. Its ID can be used for further operations.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BulkPairingsAPIBulkPairingCreateRequest
	*/
	BulkPairingCreate(ctx context.Context) BulkPairingsAPIBulkPairingCreateRequest

	// BulkPairingCreateExecute executes the request
	//  @return BulkPairingList200ResponseInner
	BulkPairingCreateExecute(r BulkPairingsAPIBulkPairingCreateRequest) (*BulkPairingList200ResponseInner, *http.Response, error)

	/*
	BulkPairingDelete Cancel a bulk pairing

	Cancel and delete a bulk pairing that is scheduled in the future.
If the games have already been created, then this does nothing.
Canceling a bulk pairing does not refund the rate limit cost of that bulk pairing.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return BulkPairingsAPIBulkPairingDeleteRequest
	*/
	BulkPairingDelete(ctx context.Context, id string) BulkPairingsAPIBulkPairingDeleteRequest

	// BulkPairingDeleteExecute executes the request
	//  @return AccountKidPost200Response
	BulkPairingDeleteExecute(r BulkPairingsAPIBulkPairingDeleteRequest) (*AccountKidPost200Response, *http.Response, error)

	/*
	BulkPairingGet Show a bulk pairing

	Get a single bulk pairing by its ID.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return BulkPairingsAPIBulkPairingGetRequest
	*/
	BulkPairingGet(ctx context.Context, id string) BulkPairingsAPIBulkPairingGetRequest

	// BulkPairingGetExecute executes the request
	//  @return BulkPairingList200ResponseInner
	BulkPairingGetExecute(r BulkPairingsAPIBulkPairingGetRequest) (*BulkPairingList200ResponseInner, *http.Response, error)

	/*
	BulkPairingIdGamesGet Export games of a bulk pairing

	Download games of a bulk in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return BulkPairingsAPIBulkPairingIdGamesGetRequest
	*/
	BulkPairingIdGamesGet(ctx context.Context, id string) BulkPairingsAPIBulkPairingIdGamesGetRequest

	// BulkPairingIdGamesGetExecute executes the request
	//  @return string
	BulkPairingIdGamesGetExecute(r BulkPairingsAPIBulkPairingIdGamesGetRequest) (string, *http.Response, error)

	/*
	BulkPairingList View your bulk pairings

	Get a list of bulk pairings you created.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BulkPairingsAPIBulkPairingListRequest
	*/
	BulkPairingList(ctx context.Context) BulkPairingsAPIBulkPairingListRequest

	// BulkPairingListExecute executes the request
	//  @return []BulkPairingList200ResponseInner
	BulkPairingListExecute(r BulkPairingsAPIBulkPairingListRequest) ([]BulkPairingList200ResponseInner, *http.Response, error)

	/*
	BulkPairingStartClocks Manually start clocks

	Immediately start all clocks of the games of a bulk pairing.
This overrides the `startClocksAt` value of an existing bulk pairing.
If the games have not yet been created (`bulk.pairAt` is in the future), then this does nothing.
If the clocks have already started (`bulk.startClocksAt` is in the past), then this does nothing.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return BulkPairingsAPIBulkPairingStartClocksRequest
	*/
	BulkPairingStartClocks(ctx context.Context, id string) BulkPairingsAPIBulkPairingStartClocksRequest

	// BulkPairingStartClocksExecute executes the request
	//  @return AccountKidPost200Response
	BulkPairingStartClocksExecute(r BulkPairingsAPIBulkPairingStartClocksRequest) (*AccountKidPost200Response, *http.Response, error)
}

// BulkPairingsAPIService BulkPairingsAPI service
type BulkPairingsAPIService service

type BulkPairingsAPIBulkPairingCreateRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
	players *string
	clockLimit *int32
	clockIncrement *int32
	days *int32
	pairAt *int64
	startClocksAt *int64
	rated *bool
	variant *string
	fen *string
	message *string
	rules *string
}

// OAuth tokens of all the players to pair, with the syntax &#x60;tokenOfWhitePlayerInGame1:tokenOfBlackPlayerInGame1,tokenOfWhitePlayerInGame2:tokenOfBlackPlayerInGame2,...&#x60;. The 2 tokens of the players of a game are separated with &#x60;:&#x60;. The first token gets the white pieces. Games are separated with &#x60;,&#x60;. Up to 1000 tokens can be sent, for a max of 500 games. Each token must be included at most once. Example: &#x60;token1:token2,token3:token4,token5:token6&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) Players(players string) BulkPairingsAPIBulkPairingCreateRequest {
	r.players = &players
	return r
}

// Clock initial time in seconds. Example: &#x60;600&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) ClockLimit(clockLimit int32) BulkPairingsAPIBulkPairingCreateRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds. Example: &#x60;2&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) ClockIncrement(clockIncrement int32) BulkPairingsAPIBulkPairingCreateRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Days per turn. For correspondence games only.
func (r BulkPairingsAPIBulkPairingCreateRequest) Days(days int32) BulkPairingsAPIBulkPairingCreateRequest {
	r.days = &days
	return r
}

// Date at which the games will be created as a Unix timestamp in milliseconds. Up to 7 days in the future. Omit, or set to current date and time, to start the games immediately. Example: &#x60;1612289869919&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) PairAt(pairAt int64) BulkPairingsAPIBulkPairingCreateRequest {
	r.pairAt = &pairAt
	return r
}

// Date at which the clocks will be automatically started as a Unix timestamp in milliseconds. Up to 7 days in the future. Note that the clocks can start earlier than specified, if players start making moves in the game. If omitted, the clocks will not start automatically. Example: &#x60;1612289869919&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) StartClocksAt(startClocksAt int64) BulkPairingsAPIBulkPairingCreateRequest {
	r.startClocksAt = &startClocksAt
	return r
}

// Game is rated and impacts players ratings
func (r BulkPairingsAPIBulkPairingCreateRequest) Rated(rated bool) BulkPairingsAPIBulkPairingCreateRequest {
	r.rated = &rated
	return r
}

func (r BulkPairingsAPIBulkPairingCreateRequest) Variant(variant string) BulkPairingsAPIBulkPairingCreateRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r BulkPairingsAPIBulkPairingCreateRequest) Fen(fen string) BulkPairingsAPIBulkPairingCreateRequest {
	r.fen = &fen
	return r
}

// Message that will be sent to each player, when the game is created.  It is sent from your user account. &#x60;{opponent}&#x60; and &#x60;{game}&#x60; are placeholders that will be replaced with the opponent and the game URLs. You can omit this field to send the default message, but if you set your own message, it must at least contain the &#x60;{game}&#x60; placeholder. 
func (r BulkPairingsAPIBulkPairingCreateRequest) Message(message string) BulkPairingsAPIBulkPairingCreateRequest {
	r.message = &message
	return r
}

// Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60; 
func (r BulkPairingsAPIBulkPairingCreateRequest) Rules(rules string) BulkPairingsAPIBulkPairingCreateRequest {
	r.rules = &rules
	return r
}

func (r BulkPairingsAPIBulkPairingCreateRequest) Execute() (*BulkPairingList200ResponseInner, *http.Response, error) {
	return r.ApiService.BulkPairingCreateExecute(r)
}

/*
BulkPairingCreate Create a bulk pairing

Schedule many games at once, up to 24h in advance.
OAuth tokens are required for all paired players, with the `challenge:write` scope.
You can schedule up to 500 games every 10 minutes. [Contact us](mailto:contact@lichess.org) if you need higher limits.
If games have a real-time clock, each player must have only one pairing.
For correspondence games, players can have multiple pairings within the same bulk.

**The entire bulk is rejected if:**
  - a token is missing
  - a token is present more than once (except in correspondence)
  - a token lacks the `challenge:write` scope
  - a player account is closed
  - a player is paired more than once (except in correspondence)
  - a bulk is already scheduled to start at the same time with the same player
  - you have 20 scheduled bulks
  - you have 1000 scheduled games

Partial bulks are never created. Either it all fails, or it all succeeds.
When it fails, it does so with an error message explaining the issue.
Failed bulks are not counted in the rate limiting, they are free.
Fix the issues, manually or programmatically, then retry to schedule the bulk.
A successful bulk creation returns a JSON bulk document. Its ID can be used for further operations.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BulkPairingsAPIBulkPairingCreateRequest
*/
func (a *BulkPairingsAPIService) BulkPairingCreate(ctx context.Context) BulkPairingsAPIBulkPairingCreateRequest {
	return BulkPairingsAPIBulkPairingCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BulkPairingList200ResponseInner
func (a *BulkPairingsAPIService) BulkPairingCreateExecute(r BulkPairingsAPIBulkPairingCreateRequest) (*BulkPairingList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkPairingList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing"

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
	if r.players != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "players", r.players, "", "")
	}
	if r.clockLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.limit", r.clockLimit, "", "")
	}
	if r.clockIncrement != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.increment", r.clockIncrement, "", "")
	}
	if r.days != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "days", r.days, "", "")
	}
	if r.pairAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pairAt", r.pairAt, "", "")
	}
	if r.startClocksAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startClocksAt", r.startClocksAt, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fen", r.fen, "", "")
	}
	if r.message != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message", r.message, "", "")
	}
	if r.rules != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rules", r.rules, "", "")
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
			var v ApiTournamentPost400Response
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

type BulkPairingsAPIBulkPairingDeleteRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
	id string
}

func (r BulkPairingsAPIBulkPairingDeleteRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.BulkPairingDeleteExecute(r)
}

/*
BulkPairingDelete Cancel a bulk pairing

Cancel and delete a bulk pairing that is scheduled in the future.
If the games have already been created, then this does nothing.
Canceling a bulk pairing does not refund the rate limit cost of that bulk pairing.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return BulkPairingsAPIBulkPairingDeleteRequest
*/
func (a *BulkPairingsAPIService) BulkPairingDelete(ctx context.Context, id string) BulkPairingsAPIBulkPairingDeleteRequest {
	return BulkPairingsAPIBulkPairingDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *BulkPairingsAPIService) BulkPairingDeleteExecute(r BulkPairingsAPIBulkPairingDeleteRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing/{id}"
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v RacerGet404Response
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

type BulkPairingsAPIBulkPairingGetRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
	id string
}

func (r BulkPairingsAPIBulkPairingGetRequest) Execute() (*BulkPairingList200ResponseInner, *http.Response, error) {
	return r.ApiService.BulkPairingGetExecute(r)
}

/*
BulkPairingGet Show a bulk pairing

Get a single bulk pairing by its ID.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return BulkPairingsAPIBulkPairingGetRequest
*/
func (a *BulkPairingsAPIService) BulkPairingGet(ctx context.Context, id string) BulkPairingsAPIBulkPairingGetRequest {
	return BulkPairingsAPIBulkPairingGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return BulkPairingList200ResponseInner
func (a *BulkPairingsAPIService) BulkPairingGetExecute(r BulkPairingsAPIBulkPairingGetRequest) (*BulkPairingList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BulkPairingList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing/{id}"
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v RacerGet404Response
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

type BulkPairingsAPIBulkPairingIdGamesGetRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
	id string
	accept *string
	moves *bool
	pgnInJson *bool
	tags *bool
	clocks *bool
	evals *bool
	accuracy *bool
	opening *bool
	division *bool
	literate *bool
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Accept(accept string) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.accept = &accept
	return r
}

// Include the PGN moves.
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Moves(moves bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) PgnInJson(pgnInJson bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Tags(tags bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Clocks(clocks bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Evals(evals bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Accuracy(accuracy bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Opening(opening bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Division(division bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.division = &division
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Literate(literate bool) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	r.literate = &literate
	return r
}

func (r BulkPairingsAPIBulkPairingIdGamesGetRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BulkPairingIdGamesGetExecute(r)
}

/*
BulkPairingIdGamesGet Export games of a bulk pairing

Download games of a bulk in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return BulkPairingsAPIBulkPairingIdGamesGetRequest
*/
func (a *BulkPairingsAPIService) BulkPairingIdGamesGet(ctx context.Context, id string) BulkPairingsAPIBulkPairingIdGamesGetRequest {
	return BulkPairingsAPIBulkPairingIdGamesGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return string
func (a *BulkPairingsAPIService) BulkPairingIdGamesGetExecute(r BulkPairingsAPIBulkPairingIdGamesGetRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingIdGamesGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing/{id}/games"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
	if r.evals != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "evals", r.evals, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "evals", defaultValue, "form", "")
		r.evals = &defaultValue
	}
	if r.accuracy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "accuracy", r.accuracy, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "accuracy", defaultValue, "form", "")
		r.accuracy = &defaultValue
	}
	if r.opening != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "opening", r.opening, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "opening", defaultValue, "form", "")
		r.opening = &defaultValue
	}
	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", defaultValue, "form", "")
		r.division = &defaultValue
	}
	if r.literate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "literate", r.literate, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "literate", defaultValue, "form", "")
		r.literate = &defaultValue
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

type BulkPairingsAPIBulkPairingListRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
}

func (r BulkPairingsAPIBulkPairingListRequest) Execute() ([]BulkPairingList200ResponseInner, *http.Response, error) {
	return r.ApiService.BulkPairingListExecute(r)
}

/*
BulkPairingList View your bulk pairings

Get a list of bulk pairings you created.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BulkPairingsAPIBulkPairingListRequest
*/
func (a *BulkPairingsAPIService) BulkPairingList(ctx context.Context) BulkPairingsAPIBulkPairingListRequest {
	return BulkPairingsAPIBulkPairingListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []BulkPairingList200ResponseInner
func (a *BulkPairingsAPIService) BulkPairingListExecute(r BulkPairingsAPIBulkPairingListRequest) ([]BulkPairingList200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []BulkPairingList200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing"

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

type BulkPairingsAPIBulkPairingStartClocksRequest struct {
	ctx context.Context
	ApiService BulkPairingsAPI
	id string
}

func (r BulkPairingsAPIBulkPairingStartClocksRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.BulkPairingStartClocksExecute(r)
}

/*
BulkPairingStartClocks Manually start clocks

Immediately start all clocks of the games of a bulk pairing.
This overrides the `startClocksAt` value of an existing bulk pairing.
If the games have not yet been created (`bulk.pairAt` is in the future), then this does nothing.
If the clocks have already started (`bulk.startClocksAt` is in the past), then this does nothing.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return BulkPairingsAPIBulkPairingStartClocksRequest
*/
func (a *BulkPairingsAPIService) BulkPairingStartClocks(ctx context.Context, id string) BulkPairingsAPIBulkPairingStartClocksRequest {
	return BulkPairingsAPIBulkPairingStartClocksRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *BulkPairingsAPIService) BulkPairingStartClocksExecute(r BulkPairingsAPIBulkPairingStartClocksRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BulkPairingsAPIService.BulkPairingStartClocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bulk-pairing/{id}/start-clocks"
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v RacerGet404Response
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
