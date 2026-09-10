/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.171
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


type BoardAPI interface {

	/*
	ApiBoardSeek Create a seek

	Create a public seek, to start a game with a random player.

### Real-time seek

Specify the `time` and `increment` clock values.
The response is streamed but doesn't contain any information.

**Keep the connection open to keep the seek active**.

If the client closes the connection, the seek is canceled. This way, if the client terminates, the user won't be paired in a game they wouldn't play.
When the seek is accepted, or expires, the server closes the connection.

**Make sure to also have an [Event stream](#tag/board/GET/api/board/game/stream/{gameId}) open**, to be notified when a game starts.
We recommend opening the [Event stream](#tag/board/GET/api/board/game/stream/{gameId}) first, then the seek stream. This way,
you won't miss the game event if the seek is accepted immediately.

### Correspondence seek

Specify the `days` per turn value.
The response is not streamed, it immediately completes with the seek ID. The seek remains active on the server until it is joined by someone.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BoardAPIApiBoardSeekRequest
	*/
	ApiBoardSeek(ctx context.Context) BoardAPIApiBoardSeekRequest

	// ApiBoardSeekExecute executes the request
	//  @return ApiBoardSeek200Response
	ApiBoardSeekExecute(r BoardAPIApiBoardSeekRequest) (*ApiBoardSeek200Response, *http.Response, error)

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
	@return BoardAPIApiStreamEventRequest
	*/
	ApiStreamEvent(ctx context.Context) BoardAPIApiStreamEventRequest

	// ApiStreamEventExecute executes the request
	//  @return ApiStreamEvent200Response
	ApiStreamEventExecute(r BoardAPIApiStreamEventRequest) (*ApiStreamEvent200Response, *http.Response, error)

	/*
	BoardGameAbort Abort a game

	Abort a game being played with the Board API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameAbortRequest
	*/
	BoardGameAbort(ctx context.Context, gameId string) BoardAPIBoardGameAbortRequest

	// BoardGameAbortExecute executes the request
	//  @return Ok
	BoardGameAbortExecute(r BoardAPIBoardGameAbortRequest) (*Ok, *http.Response, error)

	/*
	BoardGameBerserk Berserk a tournament game

	Go berserk on an arena tournament game. Halves the clock time, grants an extra point upon winning.
Only available in arena tournaments that allow berserk, and before each player has made a move.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameBerserkRequest
	*/
	BoardGameBerserk(ctx context.Context, gameId string) BoardAPIBoardGameBerserkRequest

	// BoardGameBerserkExecute executes the request
	//  @return Ok
	BoardGameBerserkExecute(r BoardAPIBoardGameBerserkRequest) (*Ok, *http.Response, error)

	/*
	BoardGameChatGet Fetch the player chat

	Get the messages posted in the private game chat, i.e. the chat between the 2 players of the game.

Games can also have a public spectator chat.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameChatGetRequest
	*/
	BoardGameChatGet(ctx context.Context, gameId string) BoardAPIBoardGameChatGetRequest

	// BoardGameChatGetExecute executes the request
	//  @return []SpectatorGameChatInner
	BoardGameChatGetExecute(r BoardAPIBoardGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error)

	/*
	BoardGameChatPost Write in the chat

	Post a message to the player or spectator chat, in a game being played with the Board API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameChatPostRequest
	*/
	BoardGameChatPost(ctx context.Context, gameId string) BoardAPIBoardGameChatPostRequest

	// BoardGameChatPostExecute executes the request
	//  @return Ok
	BoardGameChatPostExecute(r BoardAPIBoardGameChatPostRequest) (*Ok, *http.Response, error)

	/*
	BoardGameClaimDraw Claim draw of a game

	Claim draw when the opponent has left the game for a while.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameClaimDrawRequest
	*/
	BoardGameClaimDraw(ctx context.Context, gameId string) BoardAPIBoardGameClaimDrawRequest

	// BoardGameClaimDrawExecute executes the request
	//  @return Ok
	BoardGameClaimDrawExecute(r BoardAPIBoardGameClaimDrawRequest) (*Ok, *http.Response, error)

	/*
	BoardGameClaimVictory Claim victory of a game

	Claim victory when the opponent has left the game for a while.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameClaimVictoryRequest
	*/
	BoardGameClaimVictory(ctx context.Context, gameId string) BoardAPIBoardGameClaimVictoryRequest

	// BoardGameClaimVictoryExecute executes the request
	//  @return Ok
	BoardGameClaimVictoryExecute(r BoardAPIBoardGameClaimVictoryRequest) (*Ok, *http.Response, error)

	/*
	BoardGameDraw Handle draw offers

	Create/accept/decline draw offers.
- `yes`: Offer a draw, or accept the opponent's draw offer.
- `no`: Decline a draw offer from the opponent.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param accept
	@return BoardAPIBoardGameDrawRequest
	*/
	BoardGameDraw(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BoardAPIBoardGameDrawRequest

	// BoardGameDrawExecute executes the request
	//  @return Ok
	BoardGameDrawExecute(r BoardAPIBoardGameDrawRequest) (*Ok, *http.Response, error)

	/*
	BoardGameMove Make a Board move

	Make a move in a game being played with the Board API.
The move can also contain a draw offer/agreement.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param move The move to play, in UCI format
	@return BoardAPIBoardGameMoveRequest
	*/
	BoardGameMove(ctx context.Context, gameId string, move string) BoardAPIBoardGameMoveRequest

	// BoardGameMoveExecute executes the request
	//  @return Ok
	BoardGameMoveExecute(r BoardAPIBoardGameMoveRequest) (*Ok, *http.Response, error)

	/*
	BoardGameResign Resign a game

	Resign a game being played with the Board API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameResignRequest
	*/
	BoardGameResign(ctx context.Context, gameId string) BoardAPIBoardGameResignRequest

	// BoardGameResignExecute executes the request
	//  @return Ok
	BoardGameResignExecute(r BoardAPIBoardGameResignRequest) (*Ok, *http.Response, error)

	/*
	BoardGameStream Stream Board game state

	Stream the state of a game being played with the Board API, as [ndjson](#description/streaming-with-nd-json).

Use this endpoint to get updates about the game in real-time, with a single request.

Each line is a JSON object containing a `type` field. Possible values are:
  - `gameFull` Full game data. All values are immutable, except for the `state` field.
  - `gameState` Current state of the game. Immutable values not included. Sent when a move is played, a draw is offered, or when the game ends.
  - `chatLine` Chat message sent by a user in the `room` "player" or "spectator".
  - `opponentGone` Whether the opponent has left the game, and how long before you can claim a win or draw.

The first line is always of type `gameFull`.

The server closes the stream when the game ends, or if the game has already ended.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return BoardAPIBoardGameStreamRequest
	*/
	BoardGameStream(ctx context.Context, gameId string) BoardAPIBoardGameStreamRequest

	// BoardGameStreamExecute executes the request
	//  @return BoardGameStream200Response
	BoardGameStreamExecute(r BoardAPIBoardGameStreamRequest) (*BoardGameStream200Response, *http.Response, error)

	/*
	BoardGameTakeback Handle takeback offers

	Create/accept/decline takebacks.
- `yes`: Propose a takeback, or accept the opponent's takeback offer.
- `no`: Decline a takeback offer from the opponent.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@param accept
	@return BoardAPIBoardGameTakebackRequest
	*/
	BoardGameTakeback(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BoardAPIBoardGameTakebackRequest

	// BoardGameTakebackExecute executes the request
	//  @return Ok
	BoardGameTakebackExecute(r BoardAPIBoardGameTakebackRequest) (*Ok, *http.Response, error)
}

// BoardAPIService BoardAPI service
type BoardAPIService service

type BoardAPIApiBoardSeekRequest struct {
	ctx context.Context
	ApiService BoardAPI
	time *float32
	increment *int32
	days *int32
	rated *bool
	variant *VariantKey
	ratingRange *string
	color *ChallengeColor
}

// Clock initial time in minutes. Required for real-time seeks.
func (r BoardAPIApiBoardSeekRequest) Time(time float32) BoardAPIApiBoardSeekRequest {
	r.time = &time
	return r
}

// Clock increment in seconds. Required for real-time seeks.
func (r BoardAPIApiBoardSeekRequest) Increment(increment int32) BoardAPIApiBoardSeekRequest {
	r.increment = &increment
	return r
}

// Days per turn. Required for correspondence seeks.
func (r BoardAPIApiBoardSeekRequest) Days(days int32) BoardAPIApiBoardSeekRequest {
	r.days = &days
	return r
}

// Whether the game is rated and impacts players ratings.
func (r BoardAPIApiBoardSeekRequest) Rated(rated bool) BoardAPIApiBoardSeekRequest {
	r.rated = &rated
	return r
}

func (r BoardAPIApiBoardSeekRequest) Variant(variant VariantKey) BoardAPIApiBoardSeekRequest {
	r.variant = &variant
	return r
}

// The rating range of potential opponents. Better left empty. Example: 1500-1800 
func (r BoardAPIApiBoardSeekRequest) RatingRange(ratingRange string) BoardAPIApiBoardSeekRequest {
	r.ratingRange = &ratingRange
	return r
}

// The color to play. Better left empty to automatically get 50% white.
func (r BoardAPIApiBoardSeekRequest) Color(color ChallengeColor) BoardAPIApiBoardSeekRequest {
	r.color = &color
	return r
}

func (r BoardAPIApiBoardSeekRequest) Execute() (*ApiBoardSeek200Response, *http.Response, error) {
	return r.ApiService.ApiBoardSeekExecute(r)
}

/*
ApiBoardSeek Create a seek

Create a public seek, to start a game with a random player.

### Real-time seek

Specify the `time` and `increment` clock values.
The response is streamed but doesn't contain any information.

**Keep the connection open to keep the seek active**.

If the client closes the connection, the seek is canceled. This way, if the client terminates, the user won't be paired in a game they wouldn't play.
When the seek is accepted, or expires, the server closes the connection.

**Make sure to also have an [Event stream](#tag/board/GET/api/board/game/stream/{gameId}) open**, to be notified when a game starts.
We recommend opening the [Event stream](#tag/board/GET/api/board/game/stream/{gameId}) first, then the seek stream. This way,
you won't miss the game event if the seek is accepted immediately.

### Correspondence seek

Specify the `days` per turn value.
The response is not streamed, it immediately completes with the seek ID. The seek remains active on the server until it is joined by someone.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BoardAPIApiBoardSeekRequest
*/
func (a *BoardAPIService) ApiBoardSeek(ctx context.Context) BoardAPIApiBoardSeekRequest {
	return BoardAPIApiBoardSeekRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiBoardSeek200Response
func (a *BoardAPIService) ApiBoardSeekExecute(r BoardAPIApiBoardSeekRequest) (*ApiBoardSeek200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiBoardSeek200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.ApiBoardSeek")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/seek"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.time == nil {
		return localVarReturnValue, nil, reportError("time is required and must be specified")
	}
	if *r.time < 0 {
		return localVarReturnValue, nil, reportError("time must be greater than 0")
	}
	if *r.time > 180 {
		return localVarReturnValue, nil, reportError("time must be less than 180")
	}
	if r.increment == nil {
		return localVarReturnValue, nil, reportError("increment is required and must be specified")
	}
	if *r.increment < 0 {
		return localVarReturnValue, nil, reportError("increment must be greater than 0")
	}
	if *r.increment > 180 {
		return localVarReturnValue, nil, reportError("increment must be less than 180")
	}
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
	localVarHTTPHeaderAccepts := []string{"application/json", "application/x-ndjson"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.ratingRange != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "ratingRange", r.ratingRange, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "time", r.time, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "increment", r.increment, "", "")
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "color", r.color, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "days", r.days, "", "")
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

type BoardAPIApiStreamEventRequest struct {
	ctx context.Context
	ApiService BoardAPI
}

func (r BoardAPIApiStreamEventRequest) Execute() (*ApiStreamEvent200Response, *http.Response, error) {
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
 @return BoardAPIApiStreamEventRequest
*/
func (a *BoardAPIService) ApiStreamEvent(ctx context.Context) BoardAPIApiStreamEventRequest {
	return BoardAPIApiStreamEventRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiStreamEvent200Response
func (a *BoardAPIService) ApiStreamEventExecute(r BoardAPIApiStreamEventRequest) (*ApiStreamEvent200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiStreamEvent200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.ApiStreamEvent")
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

type BoardAPIBoardGameAbortRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameAbortRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameAbortExecute(r)
}

/*
BoardGameAbort Abort a game

Abort a game being played with the Board API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameAbortRequest
*/
func (a *BoardAPIService) BoardGameAbort(ctx context.Context, gameId string) BoardAPIBoardGameAbortRequest {
	return BoardAPIBoardGameAbortRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameAbortExecute(r BoardAPIBoardGameAbortRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameAbort")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/abort"
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

type BoardAPIBoardGameBerserkRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameBerserkRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameBerserkExecute(r)
}

/*
BoardGameBerserk Berserk a tournament game

Go berserk on an arena tournament game. Halves the clock time, grants an extra point upon winning.
Only available in arena tournaments that allow berserk, and before each player has made a move.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameBerserkRequest
*/
func (a *BoardAPIService) BoardGameBerserk(ctx context.Context, gameId string) BoardAPIBoardGameBerserkRequest {
	return BoardAPIBoardGameBerserkRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameBerserkExecute(r BoardAPIBoardGameBerserkRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameBerserk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/berserk"
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

type BoardAPIBoardGameChatGetRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameChatGetRequest) Execute() ([]SpectatorGameChatInner, *http.Response, error) {
	return r.ApiService.BoardGameChatGetExecute(r)
}

/*
BoardGameChatGet Fetch the player chat

Get the messages posted in the private game chat, i.e. the chat between the 2 players of the game.

Games can also have a public spectator chat.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameChatGetRequest
*/
func (a *BoardAPIService) BoardGameChatGet(ctx context.Context, gameId string) BoardAPIBoardGameChatGetRequest {
	return BoardAPIBoardGameChatGetRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return []SpectatorGameChatInner
func (a *BoardAPIService) BoardGameChatGetExecute(r BoardAPIBoardGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SpectatorGameChatInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameChatGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/chat"
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

type BoardAPIBoardGameChatPostRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
	room *string
	text *string
}

func (r BoardAPIBoardGameChatPostRequest) Room(room string) BoardAPIBoardGameChatPostRequest {
	r.room = &room
	return r
}

func (r BoardAPIBoardGameChatPostRequest) Text(text string) BoardAPIBoardGameChatPostRequest {
	r.text = &text
	return r
}

func (r BoardAPIBoardGameChatPostRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameChatPostExecute(r)
}

/*
BoardGameChatPost Write in the chat

Post a message to the player or spectator chat, in a game being played with the Board API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameChatPostRequest
*/
func (a *BoardAPIService) BoardGameChatPost(ctx context.Context, gameId string) BoardAPIBoardGameChatPostRequest {
	return BoardAPIBoardGameChatPostRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameChatPostExecute(r BoardAPIBoardGameChatPostRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameChatPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/chat"
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

type BoardAPIBoardGameClaimDrawRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameClaimDrawRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameClaimDrawExecute(r)
}

/*
BoardGameClaimDraw Claim draw of a game

Claim draw when the opponent has left the game for a while.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameClaimDrawRequest
*/
func (a *BoardAPIService) BoardGameClaimDraw(ctx context.Context, gameId string) BoardAPIBoardGameClaimDrawRequest {
	return BoardAPIBoardGameClaimDrawRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameClaimDrawExecute(r BoardAPIBoardGameClaimDrawRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameClaimDraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/claim-draw"
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

type BoardAPIBoardGameClaimVictoryRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameClaimVictoryRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameClaimVictoryExecute(r)
}

/*
BoardGameClaimVictory Claim victory of a game

Claim victory when the opponent has left the game for a while.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameClaimVictoryRequest
*/
func (a *BoardAPIService) BoardGameClaimVictory(ctx context.Context, gameId string) BoardAPIBoardGameClaimVictoryRequest {
	return BoardAPIBoardGameClaimVictoryRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameClaimVictoryExecute(r BoardAPIBoardGameClaimVictoryRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameClaimVictory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/claim-victory"
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

type BoardAPIBoardGameDrawRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
	accept BoardGameDrawAcceptParameter
}

func (r BoardAPIBoardGameDrawRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameDrawExecute(r)
}

/*
BoardGameDraw Handle draw offers

Create/accept/decline draw offers.
- `yes`: Offer a draw, or accept the opponent's draw offer.
- `no`: Decline a draw offer from the opponent.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param accept
 @return BoardAPIBoardGameDrawRequest
*/
func (a *BoardAPIService) BoardGameDraw(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BoardAPIBoardGameDrawRequest {
	return BoardAPIBoardGameDrawRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		accept: accept,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameDrawExecute(r BoardAPIBoardGameDrawRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameDraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/draw/{accept}"
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

type BoardAPIBoardGameMoveRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
	move string
	offeringDraw *bool
}

// Whether to offer (or agree to) a draw
func (r BoardAPIBoardGameMoveRequest) OfferingDraw(offeringDraw bool) BoardAPIBoardGameMoveRequest {
	r.offeringDraw = &offeringDraw
	return r
}

func (r BoardAPIBoardGameMoveRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameMoveExecute(r)
}

/*
BoardGameMove Make a Board move

Make a move in a game being played with the Board API.
The move can also contain a draw offer/agreement.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param move The move to play, in UCI format
 @return BoardAPIBoardGameMoveRequest
*/
func (a *BoardAPIService) BoardGameMove(ctx context.Context, gameId string, move string) BoardAPIBoardGameMoveRequest {
	return BoardAPIBoardGameMoveRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		move: move,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameMoveExecute(r BoardAPIBoardGameMoveRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameMove")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/move/{move}"
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

type BoardAPIBoardGameResignRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameResignRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameResignExecute(r)
}

/*
BoardGameResign Resign a game

Resign a game being played with the Board API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameResignRequest
*/
func (a *BoardAPIService) BoardGameResign(ctx context.Context, gameId string) BoardAPIBoardGameResignRequest {
	return BoardAPIBoardGameResignRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameResignExecute(r BoardAPIBoardGameResignRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameResign")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/resign"
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

type BoardAPIBoardGameStreamRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
}

func (r BoardAPIBoardGameStreamRequest) Execute() (*BoardGameStream200Response, *http.Response, error) {
	return r.ApiService.BoardGameStreamExecute(r)
}

/*
BoardGameStream Stream Board game state

Stream the state of a game being played with the Board API, as [ndjson](#description/streaming-with-nd-json).

Use this endpoint to get updates about the game in real-time, with a single request.

Each line is a JSON object containing a `type` field. Possible values are:
  - `gameFull` Full game data. All values are immutable, except for the `state` field.
  - `gameState` Current state of the game. Immutable values not included. Sent when a move is played, a draw is offered, or when the game ends.
  - `chatLine` Chat message sent by a user in the `room` "player" or "spectator".
  - `opponentGone` Whether the opponent has left the game, and how long before you can claim a win or draw.

The first line is always of type `gameFull`.

The server closes the stream when the game ends, or if the game has already ended.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return BoardAPIBoardGameStreamRequest
*/
func (a *BoardAPIService) BoardGameStream(ctx context.Context, gameId string) BoardAPIBoardGameStreamRequest {
	return BoardAPIBoardGameStreamRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return BoardGameStream200Response
func (a *BoardAPIService) BoardGameStreamExecute(r BoardAPIBoardGameStreamRequest) (*BoardGameStream200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BoardGameStream200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameStream")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/stream/{gameId}"
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

type BoardAPIBoardGameTakebackRequest struct {
	ctx context.Context
	ApiService BoardAPI
	gameId string
	accept BoardGameDrawAcceptParameter
}

func (r BoardAPIBoardGameTakebackRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BoardGameTakebackExecute(r)
}

/*
BoardGameTakeback Handle takeback offers

Create/accept/decline takebacks.
- `yes`: Propose a takeback, or accept the opponent's takeback offer.
- `no`: Decline a takeback offer from the opponent.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @param accept
 @return BoardAPIBoardGameTakebackRequest
*/
func (a *BoardAPIService) BoardGameTakeback(ctx context.Context, gameId string, accept BoardGameDrawAcceptParameter) BoardAPIBoardGameTakebackRequest {
	return BoardAPIBoardGameTakebackRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
		accept: accept,
	}
}

// Execute executes the request
//  @return Ok
func (a *BoardAPIService) BoardGameTakebackExecute(r BoardAPIBoardGameTakebackRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BoardAPIService.BoardGameTakeback")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/board/game/{gameId}/takeback/{accept}"
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
