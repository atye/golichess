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


type GamesAPI interface {

	/*
	ApiAccountPlaying Get my ongoing games

	Get the ongoing games of the current user.
Real-time and correspondence games are included.
The most urgent games are listed first.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIApiAccountPlayingRequest
	*/
	ApiAccountPlaying(ctx context.Context) GamesAPIApiAccountPlayingRequest

	// ApiAccountPlayingExecute executes the request
	//  @return ApiAccountPlaying200Response
	ApiAccountPlayingExecute(r GamesAPIApiAccountPlayingRequest) (*ApiAccountPlaying200Response, *http.Response, error)

	/*
	ApiExportBookmarks Export your bookmarked games

	Download all games bookmarked by you, in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
We recommend streaming the response, for it can be very long.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIApiExportBookmarksRequest
	*/
	ApiExportBookmarks(ctx context.Context) GamesAPIApiExportBookmarksRequest

	// ApiExportBookmarksExecute executes the request
	//  @return GamePgn200Response
	ApiExportBookmarksExecute(r GamesAPIApiExportBookmarksRequest) (*GamePgn200Response, *http.Response, error)

	/*
	ApiGamesUser Export games of a user

	Download all games of any user in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
We recommend streaming the response, for it can be very long.
<https://lichess.org/@/german11> for instance has more than 500,000 games.
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second
  - Authenticated, downloading your own games: 60 games per second


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The user name.
	@return GamesAPIApiGamesUserRequest
	*/
	ApiGamesUser(ctx context.Context, username string) GamesAPIApiGamesUserRequest

	// ApiGamesUserExecute executes the request
	//  @return GamePgn200Response
	ApiGamesUserExecute(r GamesAPIApiGamesUserRequest) (*GamePgn200Response, *http.Response, error)

	/*
	ApiImportedGamesUser Export your imported games

	Download all games imported by you. Games are exported in PGN format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIApiImportedGamesUserRequest
	*/
	ApiImportedGamesUser(ctx context.Context) GamesAPIApiImportedGamesUserRequest

	// ApiImportedGamesUserExecute executes the request
	//  @return string
	ApiImportedGamesUserExecute(r GamesAPIApiImportedGamesUserRequest) (string, *http.Response, error)

	/*
	ApiUserCurrentGame Export ongoing game of a user

	Download the ongoing game, or the last game played, of a user.
Available in either PGN or JSON format.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return GamesAPIApiUserCurrentGameRequest
	*/
	ApiUserCurrentGame(ctx context.Context, username string) GamesAPIApiUserCurrentGameRequest

	// ApiUserCurrentGameExecute executes the request
	//  @return GamePgn200Response
	ApiUserCurrentGameExecute(r GamesAPIApiUserCurrentGameRequest) (*GamePgn200Response, *http.Response, error)

	/*
	BookmarkToggle Bookmark a game

	Add or remove a bookmark on a game, for the logged in user.
By default, the bookmark is toggled: added if absent, removed if present.
Use the `v` parameter to explicitly set the bookmark instead, making the request idempotent.
Bookmarked games can be downloaded with the [export your bookmarked games](#tag/games/GET/api/games/export/bookmarks) endpoint.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return GamesAPIBookmarkToggleRequest
	*/
	BookmarkToggle(ctx context.Context, gameId string) GamesAPIBookmarkToggleRequest

	// BookmarkToggleExecute executes the request
	BookmarkToggleExecute(r GamesAPIBookmarkToggleRequest) (*http.Response, error)

	/*
	GameChatGet Fetch the spectator game chat

	Get the messages posted in the public spectator chat of a game.

Games also have a private players chat, which only the 2 players can see.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId
	@return GamesAPIGameChatGetRequest
	*/
	GameChatGet(ctx context.Context, gameId string) GamesAPIGameChatGetRequest

	// GameChatGetExecute executes the request
	//  @return []SpectatorGameChatInner
	GameChatGetExecute(r GamesAPIGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error)

	/*
	GameImport Import one game

	Import a game from PGN. See <https://lichess.org/paste>.
Rate limiting: 200 games per hour for OAuth requests, 100 games per hour for anonymous requests.
To broadcast ongoing games, consider [pushing to a broadcast instead](#tag/broadcasts/POST/api/broadcast/round/{broadcastRoundId}/push).
To analyse a position or a line, just construct an analysis board URL (most standard tags supported if URL-encoded):
[https://lichess.org/analysis/pgn/e4_e5_Nf3_Nc6_Bc4_Bc5_Bxf7+](https://lichess.org/analysis/pgn/e4_e5_Nf3_Nc6_Bc4_Bc5_Bxf7+)


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIGameImportRequest
	*/
	GameImport(ctx context.Context) GamesAPIGameImportRequest

	// GameImportExecute executes the request
	//  @return GameImport200Response
	GameImportExecute(r GamesAPIGameImportRequest) (*GameImport200Response, *http.Response, error)

	/*
	GamePgn Export one game

	Download one game in either PGN or JSON format.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param gameId The game ID
	@return GamesAPIGamePgnRequest
	*/
	GamePgn(ctx context.Context, gameId string) GamesAPIGamePgnRequest

	// GamePgnExecute executes the request
	//  @return GamePgn200Response
	GamePgnExecute(r GamesAPIGamePgnRequest) (*GamePgn200Response, *http.Response, error)

	/*
	GamesByIds Stream games by IDs

	Creates a stream of games from an arbitrary streamId, and a list of game IDs.
The stream first outputs the games that already exists, then emits an event each time a game is started or finished.
Games are streamed as [ndjson](#description/streaming-with-nd-json).
Maximum number of games: 500 for anonymous requests, or 1000 for [OAuth2 authenticated](#description/authentication) requests.
While the stream is open, it is possible to [add new game IDs to watch](#tag/games/POST/api/stream/games/{streamId}/add).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return GamesAPIGamesByIdsRequest
	*/
	GamesByIds(ctx context.Context, streamId string) GamesAPIGamesByIdsRequest

	// GamesByIdsExecute executes the request
	//  @return []GameStreamGame
	GamesByIdsExecute(r GamesAPIGamesByIdsRequest) ([]GameStreamGame, *http.Response, error)

	/*
	GamesByIdsAdd Add game IDs to stream

	Add new game IDs for [an existing stream](#tag/games/POST/api/stream/games/{streamId}) to watch.
The stream will immediately outputs the games that already exists, then emit an event each time a game is started or finished.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param streamId
	@return GamesAPIGamesByIdsAddRequest
	*/
	GamesByIdsAdd(ctx context.Context, streamId string) GamesAPIGamesByIdsAddRequest

	// GamesByIdsAddExecute executes the request
	//  @return Ok
	GamesByIdsAddExecute(r GamesAPIGamesByIdsAddRequest) (*Ok, *http.Response, error)

	/*
	GamesByUsers Stream games of users

	Stream the games played between a list of users, in real time.
Only games where **both players** are part of the list are included.
The stream emits an event each time a game is started or finished.
To also get all current ongoing games at the beginning of the stream, use the `withCurrentGames` flag.
Games are streamed as [ndjson](#description/streaming-with-nd-json).
Maximum number of users: 300.
The method is `POST` so a longer list of IDs can be sent in the request body.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIGamesByUsersRequest
	*/
	GamesByUsers(ctx context.Context) GamesAPIGamesByUsersRequest

	// GamesByUsersExecute executes the request
	//  @return []GameStreamGame
	GamesByUsersExecute(r GamesAPIGamesByUsersRequest) ([]GameStreamGame, *http.Response, error)

	/*
	GamesExportIds Export games by IDs

	Download games by IDs in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.
Games are sorted by reverse chronological order (most recent first)
The method is `POST` so a longer list of IDs can be sent in the request body.
300 IDs can be submitted.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return GamesAPIGamesExportIdsRequest
	*/
	GamesExportIds(ctx context.Context) GamesAPIGamesExportIdsRequest

	// GamesExportIdsExecute executes the request
	//  @return GamePgn200Response
	GamesExportIdsExecute(r GamesAPIGamesExportIdsRequest) (*GamePgn200Response, *http.Response, error)

	/*
	StreamGame Stream moves of a game

	Stream positions and moves of any ongoing game, in [ndjson](#description/streaming-with-nd-json).
A description of the game is sent as a first message.
Then a message is sent each time a move is played.
Finally, a description of the game is sent when it finishes, and the stream is closed.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.
No more than 8 game streams can be opened at the same time from the same IP address.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id
	@return GamesAPIStreamGameRequest
	*/
	StreamGame(ctx context.Context, id string) GamesAPIStreamGameRequest

	// StreamGameExecute executes the request
	//  @return []MoveStreamEntry
	StreamGameExecute(r GamesAPIStreamGameRequest) ([]MoveStreamEntry, *http.Response, error)
}

// GamesAPIService GamesAPI service
type GamesAPIService service

type GamesAPIApiAccountPlayingRequest struct {
	ctx context.Context
	ApiService GamesAPI
	nb *int32
}

// Max number of games to fetch
func (r GamesAPIApiAccountPlayingRequest) Nb(nb int32) GamesAPIApiAccountPlayingRequest {
	r.nb = &nb
	return r
}

func (r GamesAPIApiAccountPlayingRequest) Execute() (*ApiAccountPlaying200Response, *http.Response, error) {
	return r.ApiService.ApiAccountPlayingExecute(r)
}

/*
ApiAccountPlaying Get my ongoing games

Get the ongoing games of the current user.
Real-time and correspondence games are included.
The most urgent games are listed first.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIApiAccountPlayingRequest
*/
func (a *GamesAPIService) ApiAccountPlaying(ctx context.Context) GamesAPIApiAccountPlayingRequest {
	return GamesAPIApiAccountPlayingRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiAccountPlaying200Response
func (a *GamesAPIService) ApiAccountPlayingExecute(r GamesAPIApiAccountPlayingRequest) (*ApiAccountPlaying200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiAccountPlaying200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ApiAccountPlaying")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/account/playing"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 9
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

type GamesAPIApiExportBookmarksRequest struct {
	ctx context.Context
	ApiService GamesAPI
	accept *string
	since *int32
	until *int32
	max *int32
	moves *bool
	pgnInJson *bool
	tags *bool
	clocks *bool
	evals *bool
	accuracy *bool
	opening *bool
	division *bool
	literate *bool
	lastFen *bool
	sort *string
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r GamesAPIApiExportBookmarksRequest) Accept(accept string) GamesAPIApiExportBookmarksRequest {
	r.accept = &accept
	return r
}

// Download games bookmarked since this timestamp. Defaults to account creation date.
func (r GamesAPIApiExportBookmarksRequest) Since(since int32) GamesAPIApiExportBookmarksRequest {
	r.since = &since
	return r
}

// Download games bookmarked until this timestamp. Defaults to now.
func (r GamesAPIApiExportBookmarksRequest) Until(until int32) GamesAPIApiExportBookmarksRequest {
	r.until = &until
	return r
}

// How many bookmarked games to download. Leave empty to download all bookmarked games.
func (r GamesAPIApiExportBookmarksRequest) Max(max int32) GamesAPIApiExportBookmarksRequest {
	r.max = &max
	return r
}

// Include the PGN moves.
func (r GamesAPIApiExportBookmarksRequest) Moves(moves bool) GamesAPIApiExportBookmarksRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.
func (r GamesAPIApiExportBookmarksRequest) PgnInJson(pgnInJson bool) GamesAPIApiExportBookmarksRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r GamesAPIApiExportBookmarksRequest) Tags(tags bool) GamesAPIApiExportBookmarksRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r GamesAPIApiExportBookmarksRequest) Clocks(clocks bool) GamesAPIApiExportBookmarksRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r GamesAPIApiExportBookmarksRequest) Evals(evals bool) GamesAPIApiExportBookmarksRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r GamesAPIApiExportBookmarksRequest) Accuracy(accuracy bool) GamesAPIApiExportBookmarksRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r GamesAPIApiExportBookmarksRequest) Opening(opening bool) GamesAPIApiExportBookmarksRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r GamesAPIApiExportBookmarksRequest) Division(division bool) GamesAPIApiExportBookmarksRequest {
	r.division = &division
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r GamesAPIApiExportBookmarksRequest) Literate(literate bool) GamesAPIApiExportBookmarksRequest {
	r.literate = &literate
	return r
}

// Include the X-FEN notation of the last position of the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. 
func (r GamesAPIApiExportBookmarksRequest) LastFen(lastFen bool) GamesAPIApiExportBookmarksRequest {
	r.lastFen = &lastFen
	return r
}

// Sort order of the bookmarks.
func (r GamesAPIApiExportBookmarksRequest) Sort(sort string) GamesAPIApiExportBookmarksRequest {
	r.sort = &sort
	return r
}

func (r GamesAPIApiExportBookmarksRequest) Execute() (*GamePgn200Response, *http.Response, error) {
	return r.ApiService.ApiExportBookmarksExecute(r)
}

/*
ApiExportBookmarks Export your bookmarked games

Download all games bookmarked by you, in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
We recommend streaming the response, for it can be very long.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIApiExportBookmarksRequest
*/
func (a *GamesAPIService) ApiExportBookmarks(ctx context.Context) GamesAPIApiExportBookmarksRequest {
	return GamesAPIApiExportBookmarksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GamePgn200Response
func (a *GamesAPIService) ApiExportBookmarksExecute(r GamesAPIApiExportBookmarksRequest) (*GamePgn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamePgn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ApiExportBookmarks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/games/export/bookmarks"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "form", "")
	}
	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "form", "")
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
	if r.lastFen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastFen", r.lastFen, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastFen", defaultValue, "form", "")
		r.lastFen = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	} else {
		var defaultValue string = "dateDesc"
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", defaultValue, "form", "")
		r.sort = &defaultValue
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

type GamesAPIApiGamesUserRequest struct {
	ctx context.Context
	ApiService GamesAPI
	username string
	accept *string
	since *int32
	until *int32
	max *int32
	vs *string
	rated *bool
	perfType *PerfType
	color *string
	analysed *bool
	moves *bool
	pgnInJson *bool
	tags *bool
	clocks *bool
	evals *bool
	accuracy *bool
	opening *bool
	division *bool
	ongoing *bool
	finished *bool
	literate *bool
	lastFen *bool
	withBookmarked *bool
	sort *string
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r GamesAPIApiGamesUserRequest) Accept(accept string) GamesAPIApiGamesUserRequest {
	r.accept = &accept
	return r
}

// Download games played since this timestamp. Defaults to account creation date.
func (r GamesAPIApiGamesUserRequest) Since(since int32) GamesAPIApiGamesUserRequest {
	r.since = &since
	return r
}

// Download games played until this timestamp. Defaults to now.
func (r GamesAPIApiGamesUserRequest) Until(until int32) GamesAPIApiGamesUserRequest {
	r.until = &until
	return r
}

// How many games to download. Leave empty to download all games.
func (r GamesAPIApiGamesUserRequest) Max(max int32) GamesAPIApiGamesUserRequest {
	r.max = &max
	return r
}

// [Filter] Only games played against this opponent
func (r GamesAPIApiGamesUserRequest) Vs(vs string) GamesAPIApiGamesUserRequest {
	r.vs = &vs
	return r
}

// [Filter] Only rated (&#x60;true&#x60;) or casual (&#x60;false&#x60;) games
func (r GamesAPIApiGamesUserRequest) Rated(rated bool) GamesAPIApiGamesUserRequest {
	r.rated = &rated
	return r
}

// [Filter] Only games in these speeds or variants. Multiple perf types can be specified, separated by a comma. Example: blitz,rapid,classical 
func (r GamesAPIApiGamesUserRequest) PerfType(perfType PerfType) GamesAPIApiGamesUserRequest {
	r.perfType = &perfType
	return r
}

// [Filter] Only games played as this color.
func (r GamesAPIApiGamesUserRequest) Color(color string) GamesAPIApiGamesUserRequest {
	r.color = &color
	return r
}

// [Filter] Only games with or without a computer analysis available
func (r GamesAPIApiGamesUserRequest) Analysed(analysed bool) GamesAPIApiGamesUserRequest {
	r.analysed = &analysed
	return r
}

// Include the PGN moves.
func (r GamesAPIApiGamesUserRequest) Moves(moves bool) GamesAPIApiGamesUserRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header.
func (r GamesAPIApiGamesUserRequest) PgnInJson(pgnInJson bool) GamesAPIApiGamesUserRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r GamesAPIApiGamesUserRequest) Tags(tags bool) GamesAPIApiGamesUserRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r GamesAPIApiGamesUserRequest) Clocks(clocks bool) GamesAPIApiGamesUserRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r GamesAPIApiGamesUserRequest) Evals(evals bool) GamesAPIApiGamesUserRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r GamesAPIApiGamesUserRequest) Accuracy(accuracy bool) GamesAPIApiGamesUserRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r GamesAPIApiGamesUserRequest) Opening(opening bool) GamesAPIApiGamesUserRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r GamesAPIApiGamesUserRequest) Division(division bool) GamesAPIApiGamesUserRequest {
	r.division = &division
	return r
}

// Ongoing games are delayed by a few seconds ranging from 3 to 60 depending on the time control, as to prevent cheat bots from using this API.
func (r GamesAPIApiGamesUserRequest) Ongoing(ongoing bool) GamesAPIApiGamesUserRequest {
	r.ongoing = &ongoing
	return r
}

// Include finished games. Set to &#x60;false&#x60; to only get ongoing games.
func (r GamesAPIApiGamesUserRequest) Finished(finished bool) GamesAPIApiGamesUserRequest {
	r.finished = &finished
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r GamesAPIApiGamesUserRequest) Literate(literate bool) GamesAPIApiGamesUserRequest {
	r.literate = &literate
	return r
}

// Include the X-FEN notation of the last position of the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. 
func (r GamesAPIApiGamesUserRequest) LastFen(lastFen bool) GamesAPIApiGamesUserRequest {
	r.lastFen = &lastFen
	return r
}

// Add a &#x60;bookmarked: true&#x60; JSON field when the logged in user has bookmarked the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. 
func (r GamesAPIApiGamesUserRequest) WithBookmarked(withBookmarked bool) GamesAPIApiGamesUserRequest {
	r.withBookmarked = &withBookmarked
	return r
}

// Sort order of the games.
func (r GamesAPIApiGamesUserRequest) Sort(sort string) GamesAPIApiGamesUserRequest {
	r.sort = &sort
	return r
}

func (r GamesAPIApiGamesUserRequest) Execute() (*GamePgn200Response, *http.Response, error) {
	return r.ApiService.ApiGamesUserExecute(r)
}

/*
ApiGamesUser Export games of a user

Download all games of any user in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
We recommend streaming the response, for it can be very long.
<https://lichess.org/@/german11> for instance has more than 500,000 games.
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second
  - Authenticated, downloading your own games: 60 games per second


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user name.
 @return GamesAPIApiGamesUserRequest
*/
func (a *GamesAPIService) ApiGamesUser(ctx context.Context, username string) GamesAPIApiGamesUserRequest {
	return GamesAPIApiGamesUserRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return GamePgn200Response
func (a *GamesAPIService) ApiGamesUserExecute(r GamesAPIApiGamesUserRequest) (*GamePgn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamePgn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ApiGamesUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/games/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
	}
	if r.until != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "until", r.until, "form", "")
	}
	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "form", "")
	}
	if r.vs != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "vs", r.vs, "form", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rated", r.rated, "form", "")
	}
	if r.perfType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "perfType", r.perfType, "form", "")
	}
	if r.color != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "color", r.color, "form", "")
	}
	if r.analysed != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "analysed", r.analysed, "form", "")
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
	if r.ongoing != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ongoing", r.ongoing, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "ongoing", defaultValue, "form", "")
		r.ongoing = &defaultValue
	}
	if r.finished != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "finished", r.finished, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "finished", defaultValue, "form", "")
		r.finished = &defaultValue
	}
	if r.literate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "literate", r.literate, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "literate", defaultValue, "form", "")
		r.literate = &defaultValue
	}
	if r.lastFen != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastFen", r.lastFen, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "lastFen", defaultValue, "form", "")
		r.lastFen = &defaultValue
	}
	if r.withBookmarked != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withBookmarked", r.withBookmarked, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "withBookmarked", defaultValue, "form", "")
		r.withBookmarked = &defaultValue
	}
	if r.sort != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", r.sort, "form", "")
	} else {
		var defaultValue string = "dateDesc"
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort", defaultValue, "form", "")
		r.sort = &defaultValue
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

type GamesAPIApiImportedGamesUserRequest struct {
	ctx context.Context
	ApiService GamesAPI
}

func (r GamesAPIApiImportedGamesUserRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.ApiImportedGamesUserExecute(r)
}

/*
ApiImportedGamesUser Export your imported games

Download all games imported by you. Games are exported in PGN format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIApiImportedGamesUserRequest
*/
func (a *GamesAPIService) ApiImportedGamesUser(ctx context.Context) GamesAPIApiImportedGamesUserRequest {
	return GamesAPIApiImportedGamesUserRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return string
func (a *GamesAPIService) ApiImportedGamesUserExecute(r GamesAPIApiImportedGamesUserRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ApiImportedGamesUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/games/export/imports"

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

type GamesAPIApiUserCurrentGameRequest struct {
	ctx context.Context
	ApiService GamesAPI
	username string
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

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/json&#x60; to get the games in JSON format. 
func (r GamesAPIApiUserCurrentGameRequest) Accept(accept string) GamesAPIApiUserCurrentGameRequest {
	r.accept = &accept
	return r
}

// Include the PGN moves.
func (r GamesAPIApiUserCurrentGameRequest) Moves(moves bool) GamesAPIApiUserCurrentGameRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r GamesAPIApiUserCurrentGameRequest) PgnInJson(pgnInJson bool) GamesAPIApiUserCurrentGameRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r GamesAPIApiUserCurrentGameRequest) Tags(tags bool) GamesAPIApiUserCurrentGameRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r GamesAPIApiUserCurrentGameRequest) Clocks(clocks bool) GamesAPIApiUserCurrentGameRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r GamesAPIApiUserCurrentGameRequest) Evals(evals bool) GamesAPIApiUserCurrentGameRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r GamesAPIApiUserCurrentGameRequest) Accuracy(accuracy bool) GamesAPIApiUserCurrentGameRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r GamesAPIApiUserCurrentGameRequest) Opening(opening bool) GamesAPIApiUserCurrentGameRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r GamesAPIApiUserCurrentGameRequest) Division(division bool) GamesAPIApiUserCurrentGameRequest {
	r.division = &division
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r GamesAPIApiUserCurrentGameRequest) Literate(literate bool) GamesAPIApiUserCurrentGameRequest {
	r.literate = &literate
	return r
}

func (r GamesAPIApiUserCurrentGameRequest) Execute() (*GamePgn200Response, *http.Response, error) {
	return r.ApiService.ApiUserCurrentGameExecute(r)
}

/*
ApiUserCurrentGame Export ongoing game of a user

Download the ongoing game, or the last game played, of a user.
Available in either PGN or JSON format.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return GamesAPIApiUserCurrentGameRequest
*/
func (a *GamesAPIService) ApiUserCurrentGame(ctx context.Context, username string) GamesAPIApiUserCurrentGameRequest {
	return GamesAPIApiUserCurrentGameRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return GamePgn200Response
func (a *GamesAPIService) ApiUserCurrentGameExecute(r GamesAPIApiUserCurrentGameRequest) (*GamePgn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamePgn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.ApiUserCurrentGame")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/current-game"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

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
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.evals != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "evals", r.evals, "form", "")
	} else {
		var defaultValue bool = true
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
		var defaultValue bool = true
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
	localVarHTTPHeaderAccepts := []string{"application/json"}

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

type GamesAPIBookmarkToggleRequest struct {
	ctx context.Context
	ApiService GamesAPI
	gameId string
	v *bool
}

// Explicitly set the bookmark instead of toggling it. &#x60;true&#x60; adds the bookmark, &#x60;false&#x60; removes it. 
func (r GamesAPIBookmarkToggleRequest) V(v bool) GamesAPIBookmarkToggleRequest {
	r.v = &v
	return r
}

func (r GamesAPIBookmarkToggleRequest) Execute() (*http.Response, error) {
	return r.ApiService.BookmarkToggleExecute(r)
}

/*
BookmarkToggle Bookmark a game

Add or remove a bookmark on a game, for the logged in user.
By default, the bookmark is toggled: added if absent, removed if present.
Use the `v` parameter to explicitly set the bookmark instead, making the request idempotent.
Bookmarked games can be downloaded with the [export your bookmarked games](#tag/games/GET/api/games/export/bookmarks) endpoint.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return GamesAPIBookmarkToggleRequest
*/
func (a *GamesAPIService) BookmarkToggle(ctx context.Context, gameId string) GamesAPIBookmarkToggleRequest {
	return GamesAPIBookmarkToggleRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
func (a *GamesAPIService) BookmarkToggleExecute(r GamesAPIBookmarkToggleRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.BookmarkToggle")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/bookmark/{gameId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.v != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "v", r.v, "form", "")
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

type GamesAPIGameChatGetRequest struct {
	ctx context.Context
	ApiService GamesAPI
	gameId string
}

func (r GamesAPIGameChatGetRequest) Execute() ([]SpectatorGameChatInner, *http.Response, error) {
	return r.ApiService.GameChatGetExecute(r)
}

/*
GameChatGet Fetch the spectator game chat

Get the messages posted in the public spectator chat of a game.

Games also have a private players chat, which only the 2 players can see.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId
 @return GamesAPIGameChatGetRequest
*/
func (a *GamesAPIService) GameChatGet(ctx context.Context, gameId string) GamesAPIGameChatGetRequest {
	return GamesAPIGameChatGetRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return []SpectatorGameChatInner
func (a *GamesAPIService) GameChatGetExecute(r GamesAPIGameChatGetRequest) ([]SpectatorGameChatInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SpectatorGameChatInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GameChatGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/game/{gameId}/chat"
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

type GamesAPIGameImportRequest struct {
	ctx context.Context
	ApiService GamesAPI
	pgn *string
}

// The PGN. It can contain only one game. Most standard tags are supported.
func (r GamesAPIGameImportRequest) Pgn(pgn string) GamesAPIGameImportRequest {
	r.pgn = &pgn
	return r
}

func (r GamesAPIGameImportRequest) Execute() (*GameImport200Response, *http.Response, error) {
	return r.ApiService.GameImportExecute(r)
}

/*
GameImport Import one game

Import a game from PGN. See <https://lichess.org/paste>.
Rate limiting: 200 games per hour for OAuth requests, 100 games per hour for anonymous requests.
To broadcast ongoing games, consider [pushing to a broadcast instead](#tag/broadcasts/POST/api/broadcast/round/{broadcastRoundId}/push).
To analyse a position or a line, just construct an analysis board URL (most standard tags supported if URL-encoded):
[https://lichess.org/analysis/pgn/e4_e5_Nf3_Nc6_Bc4_Bc5_Bxf7+](https://lichess.org/analysis/pgn/e4_e5_Nf3_Nc6_Bc4_Bc5_Bxf7+)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIGameImportRequest
*/
func (a *GamesAPIService) GameImport(ctx context.Context) GamesAPIGameImportRequest {
	return GamesAPIGameImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GameImport200Response
func (a *GamesAPIService) GameImportExecute(r GamesAPIGameImportRequest) (*GameImport200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GameImport200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GameImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/import"

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
	if r.pgn != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pgn", r.pgn, "", "")
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

type GamesAPIGamePgnRequest struct {
	ctx context.Context
	ApiService GamesAPI
	gameId string
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
	withBookmarked *bool
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/json&#x60; to get the games in JSON format. 
func (r GamesAPIGamePgnRequest) Accept(accept string) GamesAPIGamePgnRequest {
	r.accept = &accept
	return r
}

// Include the PGN moves.
func (r GamesAPIGamePgnRequest) Moves(moves bool) GamesAPIGamePgnRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r GamesAPIGamePgnRequest) PgnInJson(pgnInJson bool) GamesAPIGamePgnRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r GamesAPIGamePgnRequest) Tags(tags bool) GamesAPIGamePgnRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r GamesAPIGamePgnRequest) Clocks(clocks bool) GamesAPIGamePgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r GamesAPIGamePgnRequest) Evals(evals bool) GamesAPIGamePgnRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r GamesAPIGamePgnRequest) Accuracy(accuracy bool) GamesAPIGamePgnRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r GamesAPIGamePgnRequest) Opening(opening bool) GamesAPIGamePgnRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r GamesAPIGamePgnRequest) Division(division bool) GamesAPIGamePgnRequest {
	r.division = &division
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r GamesAPIGamePgnRequest) Literate(literate bool) GamesAPIGamePgnRequest {
	r.literate = &literate
	return r
}

// Add a &#x60;bookmarked: true&#x60; JSON field when the logged in user has bookmarked the game. The response type must be set to &#x60;application/x-ndjson&#x60; by the request &#x60;Accept&#x60; header. 
func (r GamesAPIGamePgnRequest) WithBookmarked(withBookmarked bool) GamesAPIGamePgnRequest {
	r.withBookmarked = &withBookmarked
	return r
}

func (r GamesAPIGamePgnRequest) Execute() (*GamePgn200Response, *http.Response, error) {
	return r.ApiService.GamePgnExecute(r)
}

/*
GamePgn Export one game

Download one game in either PGN or JSON format.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param gameId The game ID
 @return GamesAPIGamePgnRequest
*/
func (a *GamesAPIService) GamePgn(ctx context.Context, gameId string) GamesAPIGamePgnRequest {
	return GamesAPIGamePgnRequest{
		ApiService: a,
		ctx: ctx,
		gameId: gameId,
	}
}

// Execute executes the request
//  @return GamePgn200Response
func (a *GamesAPIService) GamePgnExecute(r GamesAPIGamePgnRequest) (*GamePgn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamePgn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GamePgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/game/export/{gameId}"
	localVarPath = strings.Replace(localVarPath, "{"+"gameId"+"}", url.PathEscape(parameterValueToString(r.gameId, "gameId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.gameId) < 8 {
		return localVarReturnValue, nil, reportError("gameId must have at least 8 elements")
	}
	if strlen(r.gameId) > 8 {
		return localVarReturnValue, nil, reportError("gameId must have less than 8 elements")
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
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.evals != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "evals", r.evals, "form", "")
	} else {
		var defaultValue bool = true
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
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "opening", defaultValue, "form", "")
		r.opening = &defaultValue
	}
	if r.division != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "division", r.division, "form", "")
	} else {
		var defaultValue bool = true
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
	if r.withBookmarked != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withBookmarked", r.withBookmarked, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "withBookmarked", defaultValue, "form", "")
		r.withBookmarked = &defaultValue
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

type GamesAPIGamesByIdsRequest struct {
	ctx context.Context
	ApiService GamesAPI
	streamId string
	body *string
}

// Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId01,gameId02,gameId03&#x60; 
func (r GamesAPIGamesByIdsRequest) Body(body string) GamesAPIGamesByIdsRequest {
	r.body = &body
	return r
}

func (r GamesAPIGamesByIdsRequest) Execute() ([]GameStreamGame, *http.Response, error) {
	return r.ApiService.GamesByIdsExecute(r)
}

/*
GamesByIds Stream games by IDs

Creates a stream of games from an arbitrary streamId, and a list of game IDs.
The stream first outputs the games that already exists, then emits an event each time a game is started or finished.
Games are streamed as [ndjson](#description/streaming-with-nd-json).
Maximum number of games: 500 for anonymous requests, or 1000 for [OAuth2 authenticated](#description/authentication) requests.
While the stream is open, it is possible to [add new game IDs to watch](#tag/games/POST/api/stream/games/{streamId}/add).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param streamId
 @return GamesAPIGamesByIdsRequest
*/
func (a *GamesAPIService) GamesByIds(ctx context.Context, streamId string) GamesAPIGamesByIdsRequest {
	return GamesAPIGamesByIdsRequest{
		ApiService: a,
		ctx: ctx,
		streamId: streamId,
	}
}

// Execute executes the request
//  @return []GameStreamGame
func (a *GamesAPIService) GamesByIdsExecute(r GamesAPIGamesByIdsRequest) ([]GameStreamGame, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GameStreamGame
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GamesByIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/games/{streamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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
	localVarHTTPHeaderAccepts := []string{"application/x-ndjson"}

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

type GamesAPIGamesByIdsAddRequest struct {
	ctx context.Context
	ApiService GamesAPI
	streamId string
	body *string
}

// Up to 500 or 1000 game IDs separated by commas. Example: &#x60;gameId04,gameId05,gameId06&#x60; 
func (r GamesAPIGamesByIdsAddRequest) Body(body string) GamesAPIGamesByIdsAddRequest {
	r.body = &body
	return r
}

func (r GamesAPIGamesByIdsAddRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.GamesByIdsAddExecute(r)
}

/*
GamesByIdsAdd Add game IDs to stream

Add new game IDs for [an existing stream](#tag/games/POST/api/stream/games/{streamId}) to watch.
The stream will immediately outputs the games that already exists, then emit an event each time a game is started or finished.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param streamId
 @return GamesAPIGamesByIdsAddRequest
*/
func (a *GamesAPIService) GamesByIdsAdd(ctx context.Context, streamId string) GamesAPIGamesByIdsAddRequest {
	return GamesAPIGamesByIdsAddRequest{
		ApiService: a,
		ctx: ctx,
		streamId: streamId,
	}
}

// Execute executes the request
//  @return Ok
func (a *GamesAPIService) GamesByIdsAddExecute(r GamesAPIGamesByIdsAddRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GamesByIdsAdd")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/games/{streamId}/add"
	localVarPath = strings.Replace(localVarPath, "{"+"streamId"+"}", url.PathEscape(parameterValueToString(r.streamId, "streamId")), -1)

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

type GamesAPIGamesByUsersRequest struct {
	ctx context.Context
	ApiService GamesAPI
	body *string
	withCurrentGames *bool
}

// Up to 300 user IDs separated by commas. Example: &#x60;thibault,maia1,maia5&#x60; 
func (r GamesAPIGamesByUsersRequest) Body(body string) GamesAPIGamesByUsersRequest {
	r.body = &body
	return r
}

// Include the already started games at the beginning of the stream.
func (r GamesAPIGamesByUsersRequest) WithCurrentGames(withCurrentGames bool) GamesAPIGamesByUsersRequest {
	r.withCurrentGames = &withCurrentGames
	return r
}

func (r GamesAPIGamesByUsersRequest) Execute() ([]GameStreamGame, *http.Response, error) {
	return r.ApiService.GamesByUsersExecute(r)
}

/*
GamesByUsers Stream games of users

Stream the games played between a list of users, in real time.
Only games where **both players** are part of the list are included.
The stream emits an event each time a game is started or finished.
To also get all current ongoing games at the beginning of the stream, use the `withCurrentGames` flag.
Games are streamed as [ndjson](#description/streaming-with-nd-json).
Maximum number of users: 300.
The method is `POST` so a longer list of IDs can be sent in the request body.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIGamesByUsersRequest
*/
func (a *GamesAPIService) GamesByUsers(ctx context.Context) GamesAPIGamesByUsersRequest {
	return GamesAPIGamesByUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GameStreamGame
func (a *GamesAPIService) GamesByUsersExecute(r GamesAPIGamesByUsersRequest) ([]GameStreamGame, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GameStreamGame
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GamesByUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/games-by-users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.withCurrentGames != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withCurrentGames", r.withCurrentGames, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "withCurrentGames", defaultValue, "form", "")
		r.withCurrentGames = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"text/plain"}

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

type GamesAPIGamesExportIdsRequest struct {
	ctx context.Context
	ApiService GamesAPI
	body *string
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

// Game IDs separated by commas. Up to 300.
func (r GamesAPIGamesExportIdsRequest) Body(body string) GamesAPIGamesExportIdsRequest {
	r.body = &body
	return r
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r GamesAPIGamesExportIdsRequest) Accept(accept string) GamesAPIGamesExportIdsRequest {
	r.accept = &accept
	return r
}

// Include the PGN moves.
func (r GamesAPIGamesExportIdsRequest) Moves(moves bool) GamesAPIGamesExportIdsRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r GamesAPIGamesExportIdsRequest) PgnInJson(pgnInJson bool) GamesAPIGamesExportIdsRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r GamesAPIGamesExportIdsRequest) Tags(tags bool) GamesAPIGamesExportIdsRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r GamesAPIGamesExportIdsRequest) Clocks(clocks bool) GamesAPIGamesExportIdsRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r GamesAPIGamesExportIdsRequest) Evals(evals bool) GamesAPIGamesExportIdsRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r GamesAPIGamesExportIdsRequest) Accuracy(accuracy bool) GamesAPIGamesExportIdsRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r GamesAPIGamesExportIdsRequest) Opening(opening bool) GamesAPIGamesExportIdsRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r GamesAPIGamesExportIdsRequest) Division(division bool) GamesAPIGamesExportIdsRequest {
	r.division = &division
	return r
}

// Insert textual annotations in the PGN about the opening, analysis variations, mistakes, and game termination. Example: &#x60;5... g4? { (-0.98 → 0.60) Mistake. Best move was h6. } (5... h6 6. d4 Ne7 7. g3 d5 8. exd5 fxg3 9. hxg3 c6 10. dxc6)&#x60; 
func (r GamesAPIGamesExportIdsRequest) Literate(literate bool) GamesAPIGamesExportIdsRequest {
	r.literate = &literate
	return r
}

func (r GamesAPIGamesExportIdsRequest) Execute() (*GamePgn200Response, *http.Response, error) {
	return r.ApiService.GamesExportIdsExecute(r)
}

/*
GamesExportIds Export games by IDs

Download games by IDs in PGN or [ndjson](#description/streaming-with-nd-json) format, depending on the request `Accept` header.
Games are sorted by reverse chronological order (most recent first)
The method is `POST` so a longer list of IDs can be sent in the request body.
300 IDs can be submitted.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return GamesAPIGamesExportIdsRequest
*/
func (a *GamesAPIService) GamesExportIds(ctx context.Context) GamesAPIGamesExportIdsRequest {
	return GamesAPIGamesExportIdsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GamePgn200Response
func (a *GamesAPIService) GamesExportIdsExecute(r GamesAPIGamesExportIdsRequest) (*GamePgn200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GamePgn200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.GamesExportIds")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/games/export/_ids"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	if r.accept != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Accept", r.accept, "simple", "")
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

type GamesAPIStreamGameRequest struct {
	ctx context.Context
	ApiService GamesAPI
	id string
}

func (r GamesAPIStreamGameRequest) Execute() ([]MoveStreamEntry, *http.Response, error) {
	return r.ApiService.StreamGameExecute(r)
}

/*
StreamGame Stream moves of a game

Stream positions and moves of any ongoing game, in [ndjson](#description/streaming-with-nd-json).
A description of the game is sent as a first message.
Then a message is sent each time a move is played.
Finally, a description of the game is sent when it finishes, and the stream is closed.
Ongoing games are delayed by 3 moves, as to prevent cheat bots from using this API.
No more than 8 game streams can be opened at the same time from the same IP address.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return GamesAPIStreamGameRequest
*/
func (a *GamesAPIService) StreamGame(ctx context.Context, id string) GamesAPIStreamGameRequest {
	return GamesAPIStreamGameRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return []MoveStreamEntry
func (a *GamesAPIService) StreamGameExecute(r GamesAPIStreamGameRequest) ([]MoveStreamEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []MoveStreamEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "GamesAPIService.StreamGame")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/game/{id}"
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v StreamGame429Response
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
