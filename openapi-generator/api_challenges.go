/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.162
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


type ChallengesAPI interface {

	/*
	AdminChallengeTokens Admin challenge tokens

	**This endpoint can only be used by Lichess administrators. It will not work if you do not have the appropriate permissions.** Tournament organizers should instead use [OAuth](#tag/OAuth) to obtain `challenge:write` tokens from users in order to perform bulk pairing.*
Create and obtain `challenge:write` tokens for multiple users.
If a similar token already exists for a user, it is reused. This endpoint is idempotent.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ChallengesAPIAdminChallengeTokensRequest
	*/
	AdminChallengeTokens(ctx context.Context) ChallengesAPIAdminChallengeTokensRequest

	// AdminChallengeTokensExecute executes the request
	//  @return map[string]string
	AdminChallengeTokensExecute(r ChallengesAPIAdminChallengeTokensRequest) (map[string]string, *http.Response, error)

	/*
	ChallengeAccept Accept a challenge

	Accept an incoming challenge.
You should receive a `gameStart` event on the [incoming events stream](#tag/board/GET/api/board/game/stream/{gameId}).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param challengeId
	@return ChallengesAPIChallengeAcceptRequest
	*/
	ChallengeAccept(ctx context.Context, challengeId string) ChallengesAPIChallengeAcceptRequest

	// ChallengeAcceptExecute executes the request
	//  @return Ok
	ChallengeAcceptExecute(r ChallengesAPIChallengeAcceptRequest) (*Ok, *http.Response, error)

	/*
	ChallengeAi Challenge the AI

	Start a game with Lichess AI.
You will be notified on the [event stream](#tag/board/GET/api/board/game/stream/{gameId}) that a new game has started.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ChallengesAPIChallengeAiRequest
	*/
	ChallengeAi(ctx context.Context) ChallengesAPIChallengeAiRequest

	// ChallengeAiExecute executes the request
	//  @return ChallengeAi201Response
	ChallengeAiExecute(r ChallengesAPIChallengeAiRequest) (*ChallengeAi201Response, *http.Response, error)

	/*
	ChallengeCancel Cancel a challenge

	Cancel a challenge you sent, or aborts the game if the challenge was accepted, but the game was not yet played.
Note that the ID of a game is the same as the ID of the challenge that created it.
Works for user challenges and open challenges alike.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param challengeId
	@return ChallengesAPIChallengeCancelRequest
	*/
	ChallengeCancel(ctx context.Context, challengeId string) ChallengesAPIChallengeCancelRequest

	// ChallengeCancelExecute executes the request
	//  @return Ok
	ChallengeCancelExecute(r ChallengesAPIChallengeCancelRequest) (*Ok, *http.Response, error)

	/*
	ChallengeCreate Create a challenge

	Challenge someone to play. The targeted player can choose to accept or decline.
If the challenge is accepted, you will be notified on the [event stream](#tag/board/GET/api/board/game/stream/{gameId})
that a new game has started. The game ID will be the same as the challenge ID.
Challenges for realtime games (not correspondence) expire after 20s if not accepted.
To prevent that, use the `keepAliveStream` flag described below.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return ChallengesAPIChallengeCreateRequest
	*/
	ChallengeCreate(ctx context.Context, username string) ChallengesAPIChallengeCreateRequest

	// ChallengeCreateExecute executes the request
	//  @return ChallengeJson
	ChallengeCreateExecute(r ChallengesAPIChallengeCreateRequest) (*ChallengeJson, *http.Response, error)

	/*
	ChallengeDecline Decline a challenge

	Decline an incoming challenge.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param challengeId
	@return ChallengesAPIChallengeDeclineRequest
	*/
	ChallengeDecline(ctx context.Context, challengeId string) ChallengesAPIChallengeDeclineRequest

	// ChallengeDeclineExecute executes the request
	//  @return Ok
	ChallengeDeclineExecute(r ChallengesAPIChallengeDeclineRequest) (*Ok, *http.Response, error)

	/*
	ChallengeList List your challenges

	Get a list of challenges created by or targeted at you.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ChallengesAPIChallengeListRequest
	*/
	ChallengeList(ctx context.Context) ChallengesAPIChallengeListRequest

	// ChallengeListExecute executes the request
	//  @return ChallengeList200Response
	ChallengeListExecute(r ChallengesAPIChallengeListRequest) (*ChallengeList200Response, *http.Response, error)

	/*
	ChallengeOpen Open-ended challenge

	Create a challenge that any 2 players can join.
Share the URL of the challenge. the first 2 players to click it will be paired for a game.
The response body also contains `whiteUrl` and `blackUrl`.
You can control which color each player gets by giving them these URLs,
instead of the main challenge URL.
Open challenges expire after 24h.
If the challenge creation is [authenticated with OAuth2](#description/authentication),
then you can use the [challenge cancel endpoint](#tag/challenges/POST/api/challenge/{challengeId}/cancel) to cancel it.
To directly pair 2 known players, use [this endpoint](#tag/bulk-pairings/GET/api/bulk-pairing) instead.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ChallengesAPIChallengeOpenRequest
	*/
	ChallengeOpen(ctx context.Context) ChallengesAPIChallengeOpenRequest

	// ChallengeOpenExecute executes the request
	//  @return ChallengeOpenJson
	ChallengeOpenExecute(r ChallengesAPIChallengeOpenRequest) (*ChallengeOpenJson, *http.Response, error)

	/*
	ChallengeShow Show one challenge

	Get details about a challenge, even if it has been recently accepted, canceled or declined.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param challengeId The challenge ID
	@return ChallengesAPIChallengeShowRequest
	*/
	ChallengeShow(ctx context.Context, challengeId string) ChallengesAPIChallengeShowRequest

	// ChallengeShowExecute executes the request
	//  @return ChallengeJson
	ChallengeShowExecute(r ChallengesAPIChallengeShowRequest) (*ChallengeJson, *http.Response, error)

	/*
	ChallengeStartClocks Start clocks of a game

	Start the clocks of a game immediately, even if a player has not yet made a move.
Requires the OAuth tokens of both players with `challenge:write` scope.
If the clocks have already started, the call will have no effect.

For AI games with only one player, omit the `token2` parameter.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return ChallengesAPIChallengeStartClocksRequest
	*/
	ChallengeStartClocks(ctx context.Context, gameId string) ChallengesAPIChallengeStartClocksRequest

	// ChallengeStartClocksExecute executes the request
	//  @return Ok
	ChallengeStartClocksExecute(r ChallengesAPIChallengeStartClocksRequest) (*Ok, *http.Response, error)

	/*
	RoundAddTime Add time to the opponent clock

	Add seconds to the opponent's clock. Can be used to create games with time odds.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param seconds How many seconds to give
	@return ChallengesAPIRoundAddTimeRequest
	*/
	RoundAddTime(ctx context.Context, gameId string, seconds int32) ChallengesAPIRoundAddTimeRequest

	// RoundAddTimeExecute executes the request
	//  @return Ok
	RoundAddTimeExecute(r ChallengesAPIRoundAddTimeRequest) (*Ok, *http.Response, error)
}

// ChallengesAPIService ChallengesAPI service
type ChallengesAPIService service

type ChallengesAPIAdminChallengeTokensRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	users *string
	description *string
}

// Usernames separated with commas
func (r ChallengesAPIAdminChallengeTokensRequest) Users(users string) ChallengesAPIAdminChallengeTokensRequest {
	r.users = &users
	return r
}

// User visible description of the token
func (r ChallengesAPIAdminChallengeTokensRequest) Description(description string) ChallengesAPIAdminChallengeTokensRequest {
	r.description = &description
	return r
}

func (r ChallengesAPIAdminChallengeTokensRequest) Execute() (map[string]string, *http.Response, error) {
	return r.ApiService.AdminChallengeTokensExecute(r)
}

/*
AdminChallengeTokens Admin challenge tokens

**This endpoint can only be used by Lichess administrators. It will not work if you do not have the appropriate permissions.** Tournament organizers should instead use [OAuth](#tag/OAuth) to obtain `challenge:write` tokens from users in order to perform bulk pairing.*
Create and obtain `challenge:write` tokens for multiple users.
If a similar token already exists for a user, it is reused. This endpoint is idempotent.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ChallengesAPIAdminChallengeTokensRequest
*/
func (a *ChallengesAPIService) AdminChallengeTokens(ctx context.Context) ChallengesAPIAdminChallengeTokensRequest {
	return ChallengesAPIAdminChallengeTokensRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return map[string]string
func (a *ChallengesAPIService) AdminChallengeTokensExecute(r ChallengesAPIAdminChallengeTokensRequest) (map[string]string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.AdminChallengeTokens")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/token/admin-challenge"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.users == nil {
		return localVarReturnValue, nil, reportError("users is required and must be specified")
	}
	if r.description == nil {
		return localVarReturnValue, nil, reportError("description is required and must be specified")
	}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "users", r.users, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
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
			var v Error
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

type ChallengesAPIChallengeAcceptRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	challengeId string
	color *string
}

// Accept challenge as this color (only valid if this is an [open challenge](#challenge/open))
func (r ChallengesAPIChallengeAcceptRequest) Color(color string) ChallengesAPIChallengeAcceptRequest {
	r.color = &color
	return r
}

func (r ChallengesAPIChallengeAcceptRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ChallengeAcceptExecute(r)
}

/*
ChallengeAccept Accept a challenge

Accept an incoming challenge.
You should receive a `gameStart` event on the [incoming events stream](#tag/board/GET/api/board/game/stream/{gameId}).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param challengeId
 @return ChallengesAPIChallengeAcceptRequest
*/
func (a *ChallengesAPIService) ChallengeAccept(ctx context.Context, challengeId string) ChallengesAPIChallengeAcceptRequest {
	return ChallengesAPIChallengeAcceptRequest{
		ApiService: a,
		ctx: ctx,
		challengeId: challengeId,
	}
}

// Execute executes the request
//  @return Ok
func (a *ChallengesAPIService) ChallengeAcceptExecute(r ChallengesAPIChallengeAcceptRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeAccept")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{challengeId}/accept"
	localVarPath = strings.Replace(localVarPath, "{"+"challengeId"+"}", url.PathEscape(parameterValueToString(r.challengeId, "challengeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "color", r.color, "form", "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFound
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

type ChallengesAPIChallengeAiRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	level *int32
	clockLimit *int32
	clockIncrement *int32
	days *int32
	color *ChallengeColor
	variant *VariantKey
	fen *string
}

// AI strength
func (r ChallengesAPIChallengeAiRequest) Level(level int32) ChallengesAPIChallengeAiRequest {
	r.level = &level
	return r
}

// Clock initial time in seconds. If empty, a correspondence game is created.
func (r ChallengesAPIChallengeAiRequest) ClockLimit(clockLimit int32) ChallengesAPIChallengeAiRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds. If empty, a correspondence game is created.
func (r ChallengesAPIChallengeAiRequest) ClockIncrement(clockIncrement int32) ChallengesAPIChallengeAiRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Days per move, for correspondence games. Clock settings must be omitted.
func (r ChallengesAPIChallengeAiRequest) Days(days int32) ChallengesAPIChallengeAiRequest {
	r.days = &days
	return r
}

// Which color you get to play
func (r ChallengesAPIChallengeAiRequest) Color(color ChallengeColor) ChallengesAPIChallengeAiRequest {
	r.color = &color
	return r
}

func (r ChallengesAPIChallengeAiRequest) Variant(variant VariantKey) ChallengesAPIChallengeAiRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r ChallengesAPIChallengeAiRequest) Fen(fen string) ChallengesAPIChallengeAiRequest {
	r.fen = &fen
	return r
}

func (r ChallengesAPIChallengeAiRequest) Execute() (*ChallengeAi201Response, *http.Response, error) {
	return r.ApiService.ChallengeAiExecute(r)
}

/*
ChallengeAi Challenge the AI

Start a game with Lichess AI.
You will be notified on the [event stream](#tag/board/GET/api/board/game/stream/{gameId}) that a new game has started.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ChallengesAPIChallengeAiRequest
*/
func (a *ChallengesAPIService) ChallengeAi(ctx context.Context) ChallengesAPIChallengeAiRequest {
	return ChallengesAPIChallengeAiRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChallengeAi201Response
func (a *ChallengesAPIService) ChallengeAiExecute(r ChallengesAPIChallengeAiRequest) (*ChallengeAi201Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChallengeAi201Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeAi")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/ai"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.level == nil {
		return localVarReturnValue, nil, reportError("level is required and must be specified")
	}
	if *r.level < 1 {
		return localVarReturnValue, nil, reportError("level must be greater than 1")
	}
	if *r.level > 8 {
		return localVarReturnValue, nil, reportError("level must be less than 8")
	}

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
	parameterAddToHeaderOrQuery(localVarFormParams, "level", r.level, "", "")
	if r.clockLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.limit", r.clockLimit, "", "")
	}
	if r.clockIncrement != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.increment", r.clockIncrement, "", "")
	}
	if r.days != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "days", r.days, "", "")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fen", r.fen, "", "")
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
			var v Error
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

type ChallengesAPIChallengeCancelRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	challengeId string
	opponentToken *string
}

// Optional &#x60;challenge:write&#x60; token of the opponent. If set, the game can be canceled even if both players have moved.
func (r ChallengesAPIChallengeCancelRequest) OpponentToken(opponentToken string) ChallengesAPIChallengeCancelRequest {
	r.opponentToken = &opponentToken
	return r
}

func (r ChallengesAPIChallengeCancelRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ChallengeCancelExecute(r)
}

/*
ChallengeCancel Cancel a challenge

Cancel a challenge you sent, or aborts the game if the challenge was accepted, but the game was not yet played.
Note that the ID of a game is the same as the ID of the challenge that created it.
Works for user challenges and open challenges alike.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param challengeId
 @return ChallengesAPIChallengeCancelRequest
*/
func (a *ChallengesAPIService) ChallengeCancel(ctx context.Context, challengeId string) ChallengesAPIChallengeCancelRequest {
	return ChallengesAPIChallengeCancelRequest{
		ApiService: a,
		ctx: ctx,
		challengeId: challengeId,
	}
}

// Execute executes the request
//  @return Ok
func (a *ChallengesAPIService) ChallengeCancelExecute(r ChallengesAPIChallengeCancelRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeCancel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{challengeId}/cancel"
	localVarPath = strings.Replace(localVarPath, "{"+"challengeId"+"}", url.PathEscape(parameterValueToString(r.challengeId, "challengeId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.opponentToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "opponentToken", r.opponentToken, "form", "")
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
		if localVarHTTPResponse.StatusCode == 404 {
			var v NotFound
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

type ChallengesAPIChallengeCreateRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	username string
	days *int32
	clockLimit *int32
	clockIncrement *int32
	rated *bool
	color *ChallengeColor
	variant *VariantKey
	fen *string
	keepAliveStream *bool
	rules *string
}

// Days per turn. Required for correspondence seeks.
func (r ChallengesAPIChallengeCreateRequest) Days(days int32) ChallengesAPIChallengeCreateRequest {
	r.days = &days
	return r
}

// Clock initial time in seconds. If empty, a correspondence game is created. Valid values are 0, 15, 30, 45, 60, 90, and any multiple of 60 up to 10800 (3 hours).
func (r ChallengesAPIChallengeCreateRequest) ClockLimit(clockLimit int32) ChallengesAPIChallengeCreateRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds. If empty, a correspondence game is created.
func (r ChallengesAPIChallengeCreateRequest) ClockIncrement(clockIncrement int32) ChallengesAPIChallengeCreateRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Game is rated and impacts players ratings
func (r ChallengesAPIChallengeCreateRequest) Rated(rated bool) ChallengesAPIChallengeCreateRequest {
	r.rated = &rated
	return r
}

// Which color you get to play
func (r ChallengesAPIChallengeCreateRequest) Color(color ChallengeColor) ChallengesAPIChallengeCreateRequest {
	r.color = &color
	return r
}

func (r ChallengesAPIChallengeCreateRequest) Variant(variant VariantKey) ChallengesAPIChallengeCreateRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r ChallengesAPIChallengeCreateRequest) Fen(fen string) ChallengesAPIChallengeCreateRequest {
	r.fen = &fen
	return r
}

// If set, the response is streamed as [ndjson](#description/streaming-with-nd-json). The challenge is kept alive until the connection is closed by the client. When the challenge is accepted, declined or canceled, a message of the form &#x60;{\\\&quot;done\\\&quot;:\\\&quot;accepted\\\&quot;}&#x60; is sent, then the connection is closed by the server. If not set, the response is not streamed, and the challenge expires after 20s if not accepted. 
func (r ChallengesAPIChallengeCreateRequest) KeepAliveStream(keepAliveStream bool) ChallengesAPIChallengeCreateRequest {
	r.keepAliveStream = &keepAliveStream
	return r
}

// Extra game rules separated by commas. Example: &#x60;noAbort,noRematch&#x60; 
func (r ChallengesAPIChallengeCreateRequest) Rules(rules string) ChallengesAPIChallengeCreateRequest {
	r.rules = &rules
	return r
}

func (r ChallengesAPIChallengeCreateRequest) Execute() (*ChallengeJson, *http.Response, error) {
	return r.ApiService.ChallengeCreateExecute(r)
}

/*
ChallengeCreate Create a challenge

Challenge someone to play. The targeted player can choose to accept or decline.
If the challenge is accepted, you will be notified on the [event stream](#tag/board/GET/api/board/game/stream/{gameId})
that a new game has started. The game ID will be the same as the challenge ID.
Challenges for realtime games (not correspondence) expire after 20s if not accepted.
To prevent that, use the `keepAliveStream` flag described below.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return ChallengesAPIChallengeCreateRequest
*/
func (a *ChallengesAPIService) ChallengeCreate(ctx context.Context, username string) ChallengesAPIChallengeCreateRequest {
	return ChallengesAPIChallengeCreateRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return ChallengeJson
func (a *ChallengesAPIService) ChallengeCreateExecute(r ChallengesAPIChallengeCreateRequest) (*ChallengeJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChallengeJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.days == nil {
		return localVarReturnValue, nil, reportError("days is required and must be specified")
	}

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
	if r.clockLimit != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.limit", r.clockLimit, "", "")
	}
	if r.clockIncrement != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "clock.increment", r.clockIncrement, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "days", r.days, "", "")
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fen", r.fen, "", "")
	}
	if r.keepAliveStream != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "keepAliveStream", r.keepAliveStream, "", "")
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
			var v Error
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

type ChallengesAPIChallengeDeclineRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	challengeId string
	reason *string
}

// Reason challenge was declined. It will be translated to the player&#39;s language. See [the full list in the translation file](https://github.com/ornicar/lila/blob/master/translation/source/challenge.xml#L14).
func (r ChallengesAPIChallengeDeclineRequest) Reason(reason string) ChallengesAPIChallengeDeclineRequest {
	r.reason = &reason
	return r
}

func (r ChallengesAPIChallengeDeclineRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ChallengeDeclineExecute(r)
}

/*
ChallengeDecline Decline a challenge

Decline an incoming challenge.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param challengeId
 @return ChallengesAPIChallengeDeclineRequest
*/
func (a *ChallengesAPIService) ChallengeDecline(ctx context.Context, challengeId string) ChallengesAPIChallengeDeclineRequest {
	return ChallengesAPIChallengeDeclineRequest{
		ApiService: a,
		ctx: ctx,
		challengeId: challengeId,
	}
}

// Execute executes the request
//  @return Ok
func (a *ChallengesAPIService) ChallengeDeclineExecute(r ChallengesAPIChallengeDeclineRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeDecline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{challengeId}/decline"
	localVarPath = strings.Replace(localVarPath, "{"+"challengeId"+"}", url.PathEscape(parameterValueToString(r.challengeId, "challengeId")), -1)

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
	if r.reason != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "reason", r.reason, "", "")
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
			var v NotFound
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

type ChallengesAPIChallengeListRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
}

func (r ChallengesAPIChallengeListRequest) Execute() (*ChallengeList200Response, *http.Response, error) {
	return r.ApiService.ChallengeListExecute(r)
}

/*
ChallengeList List your challenges

Get a list of challenges created by or targeted at you.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ChallengesAPIChallengeListRequest
*/
func (a *ChallengesAPIService) ChallengeList(ctx context.Context) ChallengesAPIChallengeListRequest {
	return ChallengesAPIChallengeListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChallengeList200Response
func (a *ChallengesAPIService) ChallengeListExecute(r ChallengesAPIChallengeListRequest) (*ChallengeList200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChallengeList200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge"

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

type ChallengesAPIChallengeOpenRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	rated *bool
	clockLimit *int32
	clockIncrement *int32
	days *int32
	variant *VariantKey
	fen *string
	name *string
	rules *string
	users *string
	expiresAt *int64
}

// Game is rated and impacts players ratings
func (r ChallengesAPIChallengeOpenRequest) Rated(rated bool) ChallengesAPIChallengeOpenRequest {
	r.rated = &rated
	return r
}

// Clock initial time in seconds. If empty, a correspondence game is created.
func (r ChallengesAPIChallengeOpenRequest) ClockLimit(clockLimit int32) ChallengesAPIChallengeOpenRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds. If empty, a correspondence game is created.
func (r ChallengesAPIChallengeOpenRequest) ClockIncrement(clockIncrement int32) ChallengesAPIChallengeOpenRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Days per turn. For correspondence challenges.
func (r ChallengesAPIChallengeOpenRequest) Days(days int32) ChallengesAPIChallengeOpenRequest {
	r.days = &days
	return r
}

func (r ChallengesAPIChallengeOpenRequest) Variant(variant VariantKey) ChallengesAPIChallengeOpenRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r ChallengesAPIChallengeOpenRequest) Fen(fen string) ChallengesAPIChallengeOpenRequest {
	r.fen = &fen
	return r
}

// Optional name for the challenge, that players will see on the challenge page.
func (r ChallengesAPIChallengeOpenRequest) Name(name string) ChallengesAPIChallengeOpenRequest {
	r.name = &name
	return r
}

// Extra game rules separated by commas. Example: &#x60;noRematch,noGiveTime&#x60; The &#x60;noAbort&#x60; rule is available for Lichess admins only 
func (r ChallengesAPIChallengeOpenRequest) Rules(rules string) ChallengesAPIChallengeOpenRequest {
	r.rules = &rules
	return r
}

// Optional pair of usernames, separated by a comma. If set, only these users will be allowed to join the game. The first username gets the white pieces. Example: &#x60;Username1,Username2&#x60; 
func (r ChallengesAPIChallengeOpenRequest) Users(users string) ChallengesAPIChallengeOpenRequest {
	r.users = &users
	return r
}

// Timestamp in milliseconds to expire the challenge. Defaults to 24h after creation. Can&#39;t be more than 2 weeks after creation.
func (r ChallengesAPIChallengeOpenRequest) ExpiresAt(expiresAt int64) ChallengesAPIChallengeOpenRequest {
	r.expiresAt = &expiresAt
	return r
}

func (r ChallengesAPIChallengeOpenRequest) Execute() (*ChallengeOpenJson, *http.Response, error) {
	return r.ApiService.ChallengeOpenExecute(r)
}

/*
ChallengeOpen Open-ended challenge

Create a challenge that any 2 players can join.
Share the URL of the challenge. the first 2 players to click it will be paired for a game.
The response body also contains `whiteUrl` and `blackUrl`.
You can control which color each player gets by giving them these URLs,
instead of the main challenge URL.
Open challenges expire after 24h.
If the challenge creation is [authenticated with OAuth2](#description/authentication),
then you can use the [challenge cancel endpoint](#tag/challenges/POST/api/challenge/{challengeId}/cancel) to cancel it.
To directly pair 2 known players, use [this endpoint](#tag/bulk-pairings/GET/api/bulk-pairing) instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ChallengesAPIChallengeOpenRequest
*/
func (a *ChallengesAPIService) ChallengeOpen(ctx context.Context) ChallengesAPIChallengeOpenRequest {
	return ChallengesAPIChallengeOpenRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ChallengeOpenJson
func (a *ChallengesAPIService) ChallengeOpenExecute(r ChallengesAPIChallengeOpenRequest) (*ChallengeOpenJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChallengeOpenJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeOpen")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/open"

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
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
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
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.fen != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "fen", r.fen, "", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	if r.rules != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rules", r.rules, "", "")
	}
	if r.users != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "users", r.users, "", "")
	}
	if r.expiresAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "expiresAt", r.expiresAt, "", "")
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
			var v Error
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

type ChallengesAPIChallengeShowRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	challengeId string
}

func (r ChallengesAPIChallengeShowRequest) Execute() (*ChallengeJson, *http.Response, error) {
	return r.ApiService.ChallengeShowExecute(r)
}

/*
ChallengeShow Show one challenge

Get details about a challenge, even if it has been recently accepted, canceled or declined.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param challengeId The challenge ID
 @return ChallengesAPIChallengeShowRequest
*/
func (a *ChallengesAPIService) ChallengeShow(ctx context.Context, challengeId string) ChallengesAPIChallengeShowRequest {
	return ChallengesAPIChallengeShowRequest{
		ApiService: a,
		ctx: ctx,
		challengeId: challengeId,
	}
}

// Execute executes the request
//  @return ChallengeJson
func (a *ChallengesAPIService) ChallengeShowExecute(r ChallengesAPIChallengeShowRequest) (*ChallengeJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ChallengeJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{challengeId}/show"
	localVarPath = strings.Replace(localVarPath, "{"+"challengeId"+"}", url.PathEscape(parameterValueToString(r.challengeId, "challengeId")), -1)

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

type ChallengesAPIChallengeStartClocksRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	gameId string
	token1 *string
	token2 *string
}

// OAuth token of a player
func (r ChallengesAPIChallengeStartClocksRequest) Token1(token1 string) ChallengesAPIChallengeStartClocksRequest {
	r.token1 = &token1
	return r
}

// OAuth token of the other player. Omit for AI games that have only one player.
func (r ChallengesAPIChallengeStartClocksRequest) Token2(token2 string) ChallengesAPIChallengeStartClocksRequest {
	r.token2 = &token2
	return r
}

func (r ChallengesAPIChallengeStartClocksRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ChallengeStartClocksExecute(r)
}

/*
ChallengeStartClocks Start clocks of a game

Start the clocks of a game immediately, even if a player has not yet made a move.
Requires the OAuth tokens of both players with `challenge:write` scope.
If the clocks have already started, the call will have no effect.

For AI games with only one player, omit the `token2` parameter.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return ChallengesAPIChallengeStartClocksRequest
*/
func (a *ChallengesAPIService) ChallengeStartClocks(ctx context.Context, gameId string) ChallengesAPIChallengeStartClocksRequest {
	return ChallengesAPIChallengeStartClocksRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *ChallengesAPIService) ChallengeStartClocksExecute(r ChallengesAPIChallengeStartClocksRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.ChallengeStartClocks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/challenge/{gameId}/start-clocks"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.token1 == nil {
		return localVarReturnValue, nil, reportError("token1 is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "token1", r.token1, "form", "")
	if r.token2 != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "token2", r.token2, "form", "")
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

type ChallengesAPIRoundAddTimeRequest struct {
	ctx context.Context
	ApiService ChallengesAPI
	gameId string
	seconds int32
}

func (r ChallengesAPIRoundAddTimeRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.RoundAddTimeExecute(r)
}

/*
RoundAddTime Add time to the opponent clock

Add seconds to the opponent's clock. Can be used to create games with time odds.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param seconds How many seconds to give
 @return ChallengesAPIRoundAddTimeRequest
*/
func (a *ChallengesAPIService) RoundAddTime(ctx context.Context, gameId string, seconds int32) ChallengesAPIRoundAddTimeRequest {
	return ChallengesAPIRoundAddTimeRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		seconds: seconds,
	}
}

// Execute executes the request
//  @return Ok
func (a *ChallengesAPIService) RoundAddTimeExecute(r ChallengesAPIRoundAddTimeRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ChallengesAPIService.RoundAddTime")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/round/{gameId}/add-time/{seconds}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"seconds"+"}", url.PathEscape(parameterValueToString(r.seconds, "seconds")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.seconds < 5 {
		return localVarReturnValue, nil, reportError("seconds must be greater than 5")
	}
	if r.seconds > 60 {
		return localVarReturnValue, nil, reportError("seconds must be less than 60")
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
