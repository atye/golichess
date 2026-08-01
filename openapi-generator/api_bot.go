/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.157
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


type BotAPI interface {

	/*
	ApiBotOnline Get online bots

	Stream the [online bot users](https://lichess.org/player/bots), as [ndjson](#description/streaming-with-nd-json).

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BotAPIApiBotOnlineRequest
	*/
	ApiBotOnline(ctx context.Context) BotAPIApiBotOnlineRequest

	// ApiBotOnlineExecute executes the request
	//  @return User
	ApiBotOnlineExecute(r BotAPIApiBotOnlineRequest) (*User, *http.Response, error)

	/*
	ApiStreamEvent Stream incoming events

	Stream the events reaching a lichess user in real time as [ndjson](#description/streaming-with-nd-json).

An empty line is sent every 7 seconds for keep alive purposes.

Each non-empty line is a JSON object containing a `type` field. Possible values are:
- `gameStart` Start of a game
- `gameFinish` Completion of a game
- `challenge` A player sends you a challenge or you challenge someone
- `challengeCanceled` A player cancels their challenge to you
- `challengeDeclined` The opponent declines your challenge

When the stream opens, all current challenges and games are sent.

Only one global event stream can be active at a time. When the stream opens, the previous one with the same access token is closed.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BotAPIApiStreamEventRequest
	*/
	ApiStreamEvent(ctx context.Context) BotAPIApiStreamEventRequest

	// ApiStreamEventExecute executes the request
	//  @return ApiStreamEvent200Response
	ApiStreamEventExecute(r BotAPIApiStreamEventRequest) (*ApiStreamEvent200Response, *http.Response, error)

	/*
	BotAccountUpgrade Upgrade to Bot account

	Upgrade a lichess player account into a Bot account. Only Bot accounts can use the Bot API.
The account **cannot have played any game** before becoming a Bot account. The upgrade is **irreversible**. The account will only be able to play as a Bot.
To upgrade an account to Bot, use the [official lichess-bot client](https://github.com/lichess-bot-devs/lichess-bot), or follow these steps:
- Create an [API access token](https://lichess.org/account/oauth/token/create?scopes[]=bot:play) with "Play bot moves" permission.
- `curl -d '' https://lichess.org/api/bot/account/upgrade -H "Authorization: Bearer <yourTokenHere>"`
To know if an account has already been upgraded, use the [Get my profile API](#tag/account/GET/api/account):
the `title` field should be set to `BOT`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BotAPIBotAccountUpgradeRequest
	*/
	BotAccountUpgrade(ctx context.Context) BotAPIBotAccountUpgradeRequest

	// BotAccountUpgradeExecute executes the request
	//  @return Ok
	BotAccountUpgradeExecute(r BotAPIBotAccountUpgradeRequest) (*Ok, *http.Response, error)

	/*
	BotGameAbort Abort a game

	Abort a game being played with the Bot API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameAbortRequest
	*/
	BotGameAbort(ctx context.Context, gameId string) BotAPIBotGameAbortRequest

	// BotGameAbortExecute executes the request
	//  @return Ok
	BotGameAbortExecute(r BotAPIBotGameAbortRequest) (*Ok, *http.Response, error)

	/*
	BotGameChat Write in the chat

	Post a message to the player or spectator chat, in a game being played with the Bot API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameChatRequest
	*/
	BotGameChat(ctx context.Context, gameId string) BotAPIBotGameChatRequest

	// BotGameChatExecute executes the request
	//  @return Ok
	BotGameChatExecute(r BotAPIBotGameChatRequest) (*Ok, *http.Response, error)

	/*
	BotGameChatGet Fetch the game chat

	Get the messages posted in the game chat


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameChatGetRequest
	*/
	BotGameChatGet(ctx context.Context, gameId string) BotAPIBotGameChatGetRequest

	// BotGameChatGetExecute executes the request
	//  @return []SpectatorGameChatInner
	BotGameChatGetExecute(r BotAPIBotGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error)

	/*
	BotGameClaimDraw Claim draw of a game

	Claim draw when the opponent has left the game for a while.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameClaimDrawRequest
	*/
	BotGameClaimDraw(ctx context.Context, gameId string) BotAPIBotGameClaimDrawRequest

	// BotGameClaimDrawExecute executes the request
	//  @return Ok
	BotGameClaimDrawExecute(r BotAPIBotGameClaimDrawRequest) (*Ok, *http.Response, error)

	/*
	BotGameClaimVictory Claim victory of a game

	Claim victory when the opponent has left the game for a while.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameClaimVictoryRequest
	*/
	BotGameClaimVictory(ctx context.Context, gameId string) BotAPIBotGameClaimVictoryRequest

	// BotGameClaimVictoryExecute executes the request
	//  @return Ok
	BotGameClaimVictoryExecute(r BotAPIBotGameClaimVictoryRequest) (*Ok, *http.Response, error)

	/*
	BotGameDraw Handle draw offers

	Create/accept/decline draw offers with the Bot API.
- `yes`: Offer a draw, or accept the opponent's draw offer.
- `no`: Decline a draw offer from the opponent.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param accept
	@return BotAPIBotGameDrawRequest
	*/
	BotGameDraw(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BotAPIBotGameDrawRequest

	// BotGameDrawExecute executes the request
	//  @return Ok
	BotGameDrawExecute(r BotAPIBotGameDrawRequest) (*Ok, *http.Response, error)

	/*
	BotGameMove Make a Bot move

	Make a move in a game being played with the Bot API.
The move can also contain a draw offer/agreement.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param move The move to play, in UCI format
	@return BotAPIBotGameMoveRequest
	*/
	BotGameMove(ctx context.Context, gameId string, move string) BotAPIBotGameMoveRequest

	// BotGameMoveExecute executes the request
	//  @return Ok
	BotGameMoveExecute(r BotAPIBotGameMoveRequest) (*Ok, *http.Response, error)

	/*
	BotGameResign Resign a game

	Resign a game being played with the Bot API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameResignRequest
	*/
	BotGameResign(ctx context.Context, gameId string) BotAPIBotGameResignRequest

	// BotGameResignExecute executes the request
	//  @return Ok
	BotGameResignExecute(r BotAPIBotGameResignRequest) (*Ok, *http.Response, error)

	/*
	BotGameStream Stream Bot game state

	Stream the state of a game being played with the Bot API, as [ndjson](#description/streaming-with-nd-json).
Use this endpoint to get updates about the game in real-time, with a single request.
Each line is a JSON object containing a `type` field. Possible values are:
- `gameFull` Full game data. All values are immutable, except for the `state` field.
- `gameState` Current state of the game. Immutable values not included.
- `chatLine` Chat message sent by a user (or the bot itself) in the `room` "player" or "spectator".
- `opponentGone` Whether the opponent has left the game, and how long before you can claim a win or draw.
The first line is always of type `gameFull`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BotAPIBotGameStreamRequest
	*/
	BotGameStream(ctx context.Context, gameId string) BotAPIBotGameStreamRequest

	// BotGameStreamExecute executes the request
	//  @return BoardGameStream200Response
	BotGameStreamExecute(r BotAPIBotGameStreamRequest) (*BoardGameStream200Response, *http.Response, error)

	/*
	BotGameTakeback Handle takeback offers

	Create/accept/decline takebacks with the Bot API.
- `yes`: Propose a takeback, or accept the opponent's takeback offer.
- `no`: Decline a takeback offer from the opponent.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param accept
	@return BotAPIBotGameTakebackRequest
	*/
	BotGameTakeback(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BotAPIBotGameTakebackRequest

	// BotGameTakebackExecute executes the request
	//  @return Ok
	BotGameTakebackExecute(r BotAPIBotGameTakebackRequest) (*Ok, *http.Response, error)
}

// BotAPIService BotAPI service
type BotAPIService service

type BotAPIApiBotOnlineRequest struct {
	ctx context.Context
	ApiService BotAPI
	nb *int32
}

// How many bot users to fetch
func (r BotAPIApiBotOnlineRequest) Nb(nb int32) BotAPIApiBotOnlineRequest {
	r.nb = &nb
	return r
}

func (r BotAPIApiBotOnlineRequest) Execute() (*User, *http.Response, error) {
	return r.ApiService.ApiBotOnlineExecute(r)
}

/*
ApiBotOnline Get online bots

Stream the [online bot users](https://lichess.org/player/bots), as [ndjson](#description/streaming-with-nd-json).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BotAPIApiBotOnlineRequest
*/
func (a *BotAPIService) ApiBotOnline(ctx context.Context) BotAPIApiBotOnlineRequest {
	return BotAPIApiBotOnlineRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return User
func (a *BotAPIService) ApiBotOnlineExecute(r BotAPIApiBotOnlineRequest) (*User, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *User
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.ApiBotOnline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/online"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 100
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", defaultValue, "form", "")
		r.nb = &defaultValue
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

type BotAPIApiStreamEventRequest struct {
	ctx context.Context
	ApiService BotAPI
}

func (r BotAPIApiStreamEventRequest) Execute() (*ApiStreamEvent200Response, *http.Response, error) {
	return r.ApiService.ApiStreamEventExecute(r)
}

/*
ApiStreamEvent Stream incoming events

Stream the events reaching a lichess user in real time as [ndjson](#description/streaming-with-nd-json).

An empty line is sent every 7 seconds for keep alive purposes.

Each non-empty line is a JSON object containing a `type` field. Possible values are:
- `gameStart` Start of a game
- `gameFinish` Completion of a game
- `challenge` A player sends you a challenge or you challenge someone
- `challengeCanceled` A player cancels their challenge to you
- `challengeDeclined` The opponent declines your challenge

When the stream opens, all current challenges and games are sent.

Only one global event stream can be active at a time. When the stream opens, the previous one with the same access token is closed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BotAPIApiStreamEventRequest
*/
func (a *BotAPIService) ApiStreamEvent(ctx context.Context) BotAPIApiStreamEventRequest {
	return BotAPIApiStreamEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStreamEvent200Response
func (a *BotAPIService) ApiStreamEventExecute(r BotAPIApiStreamEventRequest) (*ApiStreamEvent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStreamEvent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.ApiStreamEvent")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/event"

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

type BotAPIBotAccountUpgradeRequest struct {
	ctx context.Context
	ApiService BotAPI
}

func (r BotAPIBotAccountUpgradeRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotAccountUpgradeExecute(r)
}

/*
BotAccountUpgrade Upgrade to Bot account

Upgrade a lichess player account into a Bot account. Only Bot accounts can use the Bot API.
The account **cannot have played any game** before becoming a Bot account. The upgrade is **irreversible**. The account will only be able to play as a Bot.
To upgrade an account to Bot, use the [official lichess-bot client](https://github.com/lichess-bot-devs/lichess-bot), or follow these steps:
- Create an [API access token](https://lichess.org/account/oauth/token/create?scopes[]=bot:play) with "Play bot moves" permission.
- `curl -d '' https://lichess.org/api/bot/account/upgrade -H "Authorization: Bearer <yourTokenHere>"`
To know if an account has already been upgraded, use the [Get my profile API](#tag/account/GET/api/account):
the `title` field should be set to `BOT`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BotAPIBotAccountUpgradeRequest
*/
func (a *BotAPIService) BotAccountUpgrade(ctx context.Context) BotAPIBotAccountUpgradeRequest {
	return BotAPIBotAccountUpgradeRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotAccountUpgradeExecute(r BotAPIBotAccountUpgradeRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotAccountUpgrade")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/account/upgrade"

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

type BotAPIBotGameAbortRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameAbortRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameAbortExecute(r)
}

/*
BotGameAbort Abort a game

Abort a game being played with the Bot API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameAbortRequest
*/
func (a *BotAPIService) BotGameAbort(ctx context.Context, gameId string) BotAPIBotGameAbortRequest {
	return BotAPIBotGameAbortRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameAbortExecute(r BotAPIBotGameAbortRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameAbort")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/abort"
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

type BotAPIBotGameChatRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
	room *string
	text *string
}

func (r BotAPIBotGameChatRequest) Room(room string) BotAPIBotGameChatRequest {
	r.room = &room
	return r
}

func (r BotAPIBotGameChatRequest) Text(text string) BotAPIBotGameChatRequest {
	r.text = &text
	return r
}

func (r BotAPIBotGameChatRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameChatExecute(r)
}

/*
BotGameChat Write in the chat

Post a message to the player or spectator chat, in a game being played with the Bot API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameChatRequest
*/
func (a *BotAPIService) BotGameChat(ctx context.Context, gameId string) BotAPIBotGameChatRequest {
	return BotAPIBotGameChatRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameChatExecute(r BotAPIBotGameChatRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameChat")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/chat"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.room == nil {
		return localVarReturnValue, nil, reportError("room is required and must be specified")
	}
	if r.text == nil {
		return localVarReturnValue, nil, reportError("text is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "room", r.room, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "text", r.text, "", "")
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

type BotAPIBotGameChatGetRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameChatGetRequest) Execute() ([]SpectatorGameChatInner, *http.Response, error) {
	return r.ApiService.BotGameChatGetExecute(r)
}

/*
BotGameChatGet Fetch the game chat

Get the messages posted in the game chat


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameChatGetRequest
*/
func (a *BotAPIService) BotGameChatGet(ctx context.Context, gameId string) BotAPIBotGameChatGetRequest {
	return BotAPIBotGameChatGetRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return []SpectatorGameChatInner
func (a *BotAPIService) BotGameChatGetExecute(r BotAPIBotGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SpectatorGameChatInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameChatGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/chat"
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

type BotAPIBotGameClaimDrawRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameClaimDrawRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameClaimDrawExecute(r)
}

/*
BotGameClaimDraw Claim draw of a game

Claim draw when the opponent has left the game for a while.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameClaimDrawRequest
*/
func (a *BotAPIService) BotGameClaimDraw(ctx context.Context, gameId string) BotAPIBotGameClaimDrawRequest {
	return BotAPIBotGameClaimDrawRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameClaimDrawExecute(r BotAPIBotGameClaimDrawRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameClaimDraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/claim-draw"
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

type BotAPIBotGameClaimVictoryRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameClaimVictoryRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameClaimVictoryExecute(r)
}

/*
BotGameClaimVictory Claim victory of a game

Claim victory when the opponent has left the game for a while.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameClaimVictoryRequest
*/
func (a *BotAPIService) BotGameClaimVictory(ctx context.Context, gameId string) BotAPIBotGameClaimVictoryRequest {
	return BotAPIBotGameClaimVictoryRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameClaimVictoryExecute(r BotAPIBotGameClaimVictoryRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameClaimVictory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/claim-victory"
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

type BotAPIBotGameDrawRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
	accept BoardGameDrawAcceptParameter
}

func (r BotAPIBotGameDrawRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameDrawExecute(r)
}

/*
BotGameDraw Handle draw offers

Create/accept/decline draw offers with the Bot API.
- `yes`: Offer a draw, or accept the opponent's draw offer.
- `no`: Decline a draw offer from the opponent.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param accept
 @return BotAPIBotGameDrawRequest
*/
func (a *BotAPIService) BotGameDraw(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BotAPIBotGameDrawRequest {
	return BotAPIBotGameDrawRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		accept: accept,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameDrawExecute(r BotAPIBotGameDrawRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameDraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/draw/{accept}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accept"+"}", url.PathEscape(parameterValueToString(r.accept, "accept")), -1)

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

type BotAPIBotGameMoveRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
	move string
	offeringDraw *bool
}

// Whether to offer (or agree to) a draw
func (r BotAPIBotGameMoveRequest) OfferingDraw(offeringDraw bool) BotAPIBotGameMoveRequest {
	r.offeringDraw = &offeringDraw
	return r
}

func (r BotAPIBotGameMoveRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameMoveExecute(r)
}

/*
BotGameMove Make a Bot move

Make a move in a game being played with the Bot API.
The move can also contain a draw offer/agreement.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param move The move to play, in UCI format
 @return BotAPIBotGameMoveRequest
*/
func (a *BotAPIService) BotGameMove(ctx context.Context, gameId string, move string) BotAPIBotGameMoveRequest {
	return BotAPIBotGameMoveRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		move: move,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameMoveExecute(r BotAPIBotGameMoveRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameMove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/move/{move}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"move"+"}", url.PathEscape(parameterValueToString(r.move, "move")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.offeringDraw != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "offeringDraw", r.offeringDraw, "form", "")
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

type BotAPIBotGameResignRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameResignRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameResignExecute(r)
}

/*
BotGameResign Resign a game

Resign a game being played with the Bot API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameResignRequest
*/
func (a *BotAPIService) BotGameResign(ctx context.Context, gameId string) BotAPIBotGameResignRequest {
	return BotAPIBotGameResignRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameResignExecute(r BotAPIBotGameResignRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameResign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/resign"
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

type BotAPIBotGameStreamRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
}

func (r BotAPIBotGameStreamRequest) Execute() (*BoardGameStream200Response, *http.Response, error) {
	return r.ApiService.BotGameStreamExecute(r)
}

/*
BotGameStream Stream Bot game state

Stream the state of a game being played with the Bot API, as [ndjson](#description/streaming-with-nd-json).
Use this endpoint to get updates about the game in real-time, with a single request.
Each line is a JSON object containing a `type` field. Possible values are:
- `gameFull` Full game data. All values are immutable, except for the `state` field.
- `gameState` Current state of the game. Immutable values not included.
- `chatLine` Chat message sent by a user (or the bot itself) in the `room` "player" or "spectator".
- `opponentGone` Whether the opponent has left the game, and how long before you can claim a win or draw.
The first line is always of type `gameFull`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BotAPIBotGameStreamRequest
*/
func (a *BotAPIService) BotGameStream(ctx context.Context, gameId string) BotAPIBotGameStreamRequest {
	return BotAPIBotGameStreamRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return BoardGameStream200Response
func (a *BotAPIService) BotGameStreamExecute(r BotAPIBotGameStreamRequest) (*BoardGameStream200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BoardGameStream200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/stream/{gameId}"
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
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson", "application/json"}

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

type BotAPIBotGameTakebackRequest struct {
	ctx context.Context
	ApiService BotAPI
	gameId string
	accept BoardGameDrawAcceptParameter
}

func (r BotAPIBotGameTakebackRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BotGameTakebackExecute(r)
}

/*
BotGameTakeback Handle takeback offers

Create/accept/decline takebacks with the Bot API.
- `yes`: Propose a takeback, or accept the opponent's takeback offer.
- `no`: Decline a takeback offer from the opponent.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param accept
 @return BotAPIBotGameTakebackRequest
*/
func (a *BotAPIService) BotGameTakeback(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BotAPIBotGameTakebackRequest {
	return BotAPIBotGameTakebackRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		accept: accept,
	}
}

// Execute executes the request
//  @return Ok
func (a *BotAPIService) BotGameTakebackExecute(r BotAPIBotGameTakebackRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BotAPIService.BotGameTakeback")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/bot/game/{gameId}/takeback/{accept}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"accept"+"}", url.PathEscape(parameterValueToString(r.accept, "accept")), -1)

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
