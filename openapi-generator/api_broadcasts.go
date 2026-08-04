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
)


type BroadcastsAPI interface {

	/*
	BroadcastAllRoundsPgn Export all rounds as PGN

	Download all games of all rounds of a broadcast in PGN format.
If a `study:read` [OAuth token](#tag/OAuth) is provided,
the private rounds where the user is a contributor will be available.
You may want to [download only the games of a single round](#tag/broadcasts/GET/api/broadcast/round/{broadcastRoundId}.pgn) instead.

To get real-time updates about an ongoing tournament, please use the 
[round PGN stream](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn) or
[group PGN stream](#tag/broadcasts/GET/api/stream/broadcast/group/{broadcastGroupId}.pgn) endpoints instead.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@return BroadcastsAPIBroadcastAllRoundsPgnRequest
	*/
	BroadcastAllRoundsPgn(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastAllRoundsPgnRequest

	// BroadcastAllRoundsPgnExecute executes the request
	//  @return string
	BroadcastAllRoundsPgnExecute(r BroadcastsAPIBroadcastAllRoundsPgnRequest) (string, *http.Response, error)

	/*
	BroadcastMyRoundsGet Get your broadcast rounds

	Stream all broadcast rounds you are a member of.
Also includes broadcasts rounds you did not create, but were invited to.
Also includes broadcasts rounds where you're a non-writing member. See the `writeable` flag in the response.
Rounds are ordered by rank, which is roughly chronological, most recent first, slightly pondered with popularity.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BroadcastsAPIBroadcastMyRoundsGetRequest
	*/
	BroadcastMyRoundsGet(ctx context.Context) BroadcastsAPIBroadcastMyRoundsGetRequest

	// BroadcastMyRoundsGetExecute executes the request
	//  @return BroadcastMyRound
	BroadcastMyRoundsGetExecute(r BroadcastsAPIBroadcastMyRoundsGetRequest) (*BroadcastMyRound, *http.Response, error)

	/*
	BroadcastPlayerGet Get a player of a broadcast

	Get the details of a specific player and their games from a broadcast tournament.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@param playerId The unique player ID within the broadcast. This is usually their fideId.  If the player does not have a fideId, it is their name. Consult the [list of players for the broadcast](#tag/broadcasts/GET/broadcast/{broadcastTournamentId}/players) for which ID to use. 
	@return BroadcastsAPIBroadcastPlayerGetRequest
	*/
	BroadcastPlayerGet(ctx context.Context, broadcastTournamentId string, playerId string) BroadcastsAPIBroadcastPlayerGetRequest

	// BroadcastPlayerGetExecute executes the request
	//  @return BroadcastPlayerEntryWithFideAndGames
	BroadcastPlayerGetExecute(r BroadcastsAPIBroadcastPlayerGetRequest) (*BroadcastPlayerEntryWithFideAndGames, *http.Response, error)

	/*
	BroadcastPlayersGet Get players of a broadcast

	Get the list of players of a broadcast tournament, if available.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@return BroadcastsAPIBroadcastPlayersGetRequest
	*/
	BroadcastPlayersGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastPlayersGetRequest

	// BroadcastPlayersGetExecute executes the request
	//  @return []BroadcastPlayerEntry
	BroadcastPlayersGetExecute(r BroadcastsAPIBroadcastPlayersGetRequest) ([]BroadcastPlayerEntry, *http.Response, error)

	/*
	BroadcastPush Push PGN to a broadcast round

	Update a broadcast with new PGN.
Only for broadcasts without a source URL.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastRoundId The broadcast round ID
	@return BroadcastsAPIBroadcastPushRequest
	*/
	BroadcastPush(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastPushRequest

	// BroadcastPushExecute executes the request
	//  @return BroadcastPgnPush
	BroadcastPushExecute(r BroadcastsAPIBroadcastPushRequest) (*BroadcastPgnPush, *http.Response, error)

	/*
	BroadcastRoundCreate Create a broadcast round

	Create a new broadcast round to relay external games.
This endpoint accepts the same form data as the web form.

Choose one between `syncUrl`, `syncUrls`, `syncIds` and `syncUsers`, if it is missing, the broadcast needs to be fed by [pushing PGN to it](#tag/broadcasts/POST/api/broadcast/round/{broadcastRoundId}/push)


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@return BroadcastsAPIBroadcastRoundCreateRequest
	*/
	BroadcastRoundCreate(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastRoundCreateRequest

	// BroadcastRoundCreateExecute executes the request
	//  @return BroadcastRoundNew
	BroadcastRoundCreateExecute(r BroadcastsAPIBroadcastRoundCreateRequest) (*BroadcastRoundNew, *http.Response, error)

	/*
	BroadcastRoundGet Get a broadcast round

	Get information about a broadcast round.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentSlug The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
	@param broadcastRoundSlug The broadcast round slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
	@param broadcastRoundId The broadcast Round ID
	@return BroadcastsAPIBroadcastRoundGetRequest
	*/
	BroadcastRoundGet(ctx context.Context, broadcastTournamentSlug string, broadcastRoundSlug string, broadcastRoundId string) BroadcastsAPIBroadcastRoundGetRequest

	// BroadcastRoundGetExecute executes the request
	//  @return BroadcastRound
	BroadcastRoundGetExecute(r BroadcastsAPIBroadcastRoundGetRequest) (*BroadcastRound, *http.Response, error)

	/*
	BroadcastRoundPgn Export one round as PGN

	Download all games of a single round of a broadcast tournament in PGN format.
You *could* poll this endpoint to get updates about a tournament, but it would be slow,
and very inefficient.
Instead, consider [streaming the tournament](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn) to get
a new PGN every time a game is updated, in real-time.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastRoundId The round ID
	@return BroadcastsAPIBroadcastRoundPgnRequest
	*/
	BroadcastRoundPgn(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundPgnRequest

	// BroadcastRoundPgnExecute executes the request
	//  @return string
	BroadcastRoundPgnExecute(r BroadcastsAPIBroadcastRoundPgnRequest) (string, *http.Response, error)

	/*
	BroadcastRoundReset Reset a broadcast round

	Remove any games from the broadcast round and reset it to its initial state.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastRoundId The broadcast round ID
	@return BroadcastsAPIBroadcastRoundResetRequest
	*/
	BroadcastRoundReset(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundResetRequest

	// BroadcastRoundResetExecute executes the request
	//  @return Ok
	BroadcastRoundResetExecute(r BroadcastsAPIBroadcastRoundResetRequest) (*Ok, *http.Response, error)

	/*
	BroadcastRoundUpdate Update a broadcast round

	Update information about a broadcast round.
This endpoint accepts the same form data as the web form.
All fields must be populated with data. Missing fields will override the broadcast with empty data.
For instance, if you omit `startDate`, then any pre-existing start date will be removed.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastRoundId The broadcast round ID
	@return BroadcastsAPIBroadcastRoundUpdateRequest
	*/
	BroadcastRoundUpdate(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundUpdateRequest

	// BroadcastRoundUpdateExecute executes the request
	//  @return BroadcastRound
	BroadcastRoundUpdateExecute(r BroadcastsAPIBroadcastRoundUpdateRequest) (*BroadcastRound, *http.Response, error)

	/*
	BroadcastStreamGroupPgn Stream ongoing broadcast rounds of a group as PGN

	For a given broadcast group ([example](https://lichess.org/broadcast/fide-world-cadets-cup-2026/albQx5zq)),
selects all the ongoing and recently finished rounds, and sends all games of these rounds in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the rounds.
This is the best way to get updates about an ongoing broadcast with multiple concurrent rounds.
To stream a single round, use [this endpoint instead](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastGroupId The broadcast group ID. It's the last past of the group URL, which can be found on the broadcast page overview section. [Example](https://lichess.org/broadcast/fide-world-cadets-cup-2026-o12/round-9/FUQFT39K) the group URL is  https://lichess.org/broadcast/fide-world-cadets-cup-2026/albQx5zq and therefore the group ID is `albQx5zq`. 
	@return BroadcastsAPIBroadcastStreamGroupPgnRequest
	*/
	BroadcastStreamGroupPgn(ctx context.Context, broadcastGroupId string) BroadcastsAPIBroadcastStreamGroupPgnRequest

	// BroadcastStreamGroupPgnExecute executes the request
	//  @return string
	BroadcastStreamGroupPgnExecute(r BroadcastsAPIBroadcastStreamGroupPgnRequest) (string, *http.Response, error)

	/*
	BroadcastStreamRoundPgn Stream an ongoing broadcast round as PGN

	This streaming endpoint first sends all games of a broadcast round in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the round.
This is the best way to get updates about an ongoing round. Streaming means no polling,
and no pollings means no latency, and minimum impact on the server.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastRoundId The broadcast round ID
	@return BroadcastsAPIBroadcastStreamRoundPgnRequest
	*/
	BroadcastStreamRoundPgn(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastStreamRoundPgnRequest

	// BroadcastStreamRoundPgnExecute executes the request
	//  @return string
	BroadcastStreamRoundPgnExecute(r BroadcastsAPIBroadcastStreamRoundPgnRequest) (string, *http.Response, error)

	/*
	BroadcastStreamTourPgn Stream ongoing broadcast rounds of a tournament as PGN

	For a given broadcast tournament ([example](https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg)),
selects all the ongoing and recently finished rounds, and sends all games of these rounds in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the rounds.
This is the best way to get updates about an ongoing broadcast tournament across all its rounds.
To stream a single round, use [this endpoint instead](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTourId The broadcast tournament ID. It's the last past of the tournament URL, which can be found on the broadcast page overview section. [Example](https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg) the tournament URL is  https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg and therefore the tournament ID is `jfEpUuzg `. 
	@return BroadcastsAPIBroadcastStreamTourPgnRequest
	*/
	BroadcastStreamTourPgn(ctx context.Context, broadcastTourId string) BroadcastsAPIBroadcastStreamTourPgnRequest

	// BroadcastStreamTourPgnExecute executes the request
	//  @return string
	BroadcastStreamTourPgnExecute(r BroadcastsAPIBroadcastStreamTourPgnRequest) (string, *http.Response, error)

	/*
	BroadcastTeamLeaderboardGet Get the team leaderboard of a broadcast

	Get the team leaderboard of a broadcast tournament, if available.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@return BroadcastsAPIBroadcastTeamLeaderboardGetRequest
	*/
	BroadcastTeamLeaderboardGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTeamLeaderboardGetRequest

	// BroadcastTeamLeaderboardGetExecute executes the request
	//  @return []BroadcastTeamLeaderboardEntry
	BroadcastTeamLeaderboardGetExecute(r BroadcastsAPIBroadcastTeamLeaderboardGetRequest) ([]BroadcastTeamLeaderboardEntry, *http.Response, error)

	/*
	BroadcastTourCreate Create a broadcast tournament

	Create a new broadcast tournament to relay external games.
This endpoint accepts the same form data as the [web form](https://lichess.org/broadcast/new).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BroadcastsAPIBroadcastTourCreateRequest
	*/
	BroadcastTourCreate(ctx context.Context) BroadcastsAPIBroadcastTourCreateRequest

	// BroadcastTourCreateExecute executes the request
	//  @return BroadcastWithRounds
	BroadcastTourCreateExecute(r BroadcastsAPIBroadcastTourCreateRequest) (*BroadcastWithRounds, *http.Response, error)

	/*
	BroadcastTourGet Get a broadcast tournament

	Get information about a broadcast tournament.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast tournament ID
	@return BroadcastsAPIBroadcastTourGetRequest
	*/
	BroadcastTourGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTourGetRequest

	// BroadcastTourGetExecute executes the request
	//  @return BroadcastWithRoundsAndFullGroup
	BroadcastTourGetExecute(r BroadcastsAPIBroadcastTourGetRequest) (*BroadcastWithRoundsAndFullGroup, *http.Response, error)

	/*
	BroadcastTourUpdate Update your broadcast tournament

	Update information about a broadcast tournament that you created.
This endpoint accepts the same form data as the web form.
All fields must be populated with data. Missing fields will override the broadcast with empty data.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param broadcastTournamentId The broadcast ID
	@return BroadcastsAPIBroadcastTourUpdateRequest
	*/
	BroadcastTourUpdate(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTourUpdateRequest

	// BroadcastTourUpdateExecute executes the request
	//  @return Ok
	BroadcastTourUpdateExecute(r BroadcastsAPIBroadcastTourUpdateRequest) (*Ok, *http.Response, error)

	/*
	BroadcastsByUser Get broadcasts created by a user

	Get all incoming, ongoing, and finished official broadcasts.
The broadcasts are sorted by created date, most recent first.

If you are authenticated as the user whose broadcasts you are requesting, you will also see your private and unlisted broadcasts.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return BroadcastsAPIBroadcastsByUserRequest
	*/
	BroadcastsByUser(ctx context.Context, username string) BroadcastsAPIBroadcastsByUserRequest

	// BroadcastsByUserExecute executes the request
	//  @return BroadcastsByUser200Response
	BroadcastsByUserExecute(r BroadcastsAPIBroadcastsByUserRequest) (*BroadcastsByUser200Response, *http.Response, error)

	/*
	BroadcastsOfficial Get official broadcasts

	Returns active (a round is scheduled or ongoing) official broadcasts sorted by tier. 
After that, returns finished broadcasts sorted by most recent sync time.
Broadcasts are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BroadcastsAPIBroadcastsOfficialRequest
	*/
	BroadcastsOfficial(ctx context.Context) BroadcastsAPIBroadcastsOfficialRequest

	// BroadcastsOfficialExecute executes the request
	//  @return BroadcastWithRounds
	BroadcastsOfficialExecute(r BroadcastsAPIBroadcastsOfficialRequest) (*BroadcastWithRounds, *http.Response, error)

	/*
	BroadcastsSearch Search broadcasts

	Search across recent official broadcasts.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BroadcastsAPIBroadcastsSearchRequest
	*/
	BroadcastsSearch(ctx context.Context) BroadcastsAPIBroadcastsSearchRequest

	// BroadcastsSearchExecute executes the request
	//  @return BroadcastsSearch200Response
	BroadcastsSearchExecute(r BroadcastsAPIBroadcastsSearchRequest) (*BroadcastsSearch200Response, *http.Response, error)

	/*
	BroadcastsTop Get paginated top broadcast previews

	The same data, in the same order, as can be seen on [https://lichess.org/broadcast](/broadcast).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return BroadcastsAPIBroadcastsTopRequest
	*/
	BroadcastsTop(ctx context.Context) BroadcastsAPIBroadcastsTopRequest

	// BroadcastsTopExecute executes the request
	//  @return BroadcastTop
	BroadcastsTopExecute(r BroadcastsAPIBroadcastsTopRequest) (*BroadcastTop, *http.Response, error)
}

// BroadcastsAPIService BroadcastsAPI service
type BroadcastsAPIService service

type BroadcastsAPIBroadcastAllRoundsPgnRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
	clocks *bool
	comments *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r BroadcastsAPIBroadcastAllRoundsPgnRequest) Clocks(clocks bool) BroadcastsAPIBroadcastAllRoundsPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60; 
func (r BroadcastsAPIBroadcastAllRoundsPgnRequest) Comments(comments bool) BroadcastsAPIBroadcastAllRoundsPgnRequest {
	r.comments = &comments
	return r
}

func (r BroadcastsAPIBroadcastAllRoundsPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BroadcastAllRoundsPgnExecute(r)
}

/*
BroadcastAllRoundsPgn Export all rounds as PGN

Download all games of all rounds of a broadcast in PGN format.
If a `study:read` [OAuth token](#tag/OAuth) is provided,
the private rounds where the user is a contributor will be available.
You may want to [download only the games of a single round](#tag/broadcasts/GET/api/broadcast/round/{broadcastRoundId}.pgn) instead.

To get real-time updates about an ongoing tournament, please use the 
[round PGN stream](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn) or
[group PGN stream](#tag/broadcasts/GET/api/stream/broadcast/group/{broadcastGroupId}.pgn) endpoints instead.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @return BroadcastsAPIBroadcastAllRoundsPgnRequest
*/
func (a *BroadcastsAPIService) BroadcastAllRoundsPgn(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastAllRoundsPgnRequest {
	return BroadcastsAPIBroadcastAllRoundsPgnRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return string
func (a *BroadcastsAPIService) BroadcastAllRoundsPgnExecute(r BroadcastsAPIBroadcastAllRoundsPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastAllRoundsPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/{broadcastTournamentId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
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

type BroadcastsAPIBroadcastMyRoundsGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	nb *int32
}

// How many rounds to get
func (r BroadcastsAPIBroadcastMyRoundsGetRequest) Nb(nb int32) BroadcastsAPIBroadcastMyRoundsGetRequest {
	r.nb = &nb
	return r
}

func (r BroadcastsAPIBroadcastMyRoundsGetRequest) Execute() (*BroadcastMyRound, *http.Response, error) {
	return r.ApiService.BroadcastMyRoundsGetExecute(r)
}

/*
BroadcastMyRoundsGet Get your broadcast rounds

Stream all broadcast rounds you are a member of.
Also includes broadcasts rounds you did not create, but were invited to.
Also includes broadcasts rounds where you're a non-writing member. See the `writeable` flag in the response.
Rounds are ordered by rank, which is roughly chronological, most recent first, slightly pondered with popularity.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BroadcastsAPIBroadcastMyRoundsGetRequest
*/
func (a *BroadcastsAPIService) BroadcastMyRoundsGet(ctx context.Context) BroadcastsAPIBroadcastMyRoundsGetRequest {
	return BroadcastsAPIBroadcastMyRoundsGetRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BroadcastMyRound
func (a *BroadcastsAPIService) BroadcastMyRoundsGetExecute(r BroadcastsAPIBroadcastMyRoundsGetRequest) (*BroadcastMyRound, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastMyRound
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastMyRoundsGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/my-rounds"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
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

type BroadcastsAPIBroadcastPlayerGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
	playerId string
}

func (r BroadcastsAPIBroadcastPlayerGetRequest) Execute() (*BroadcastPlayerEntryWithFideAndGames, *http.Response, error) {
	return r.ApiService.BroadcastPlayerGetExecute(r)
}

/*
BroadcastPlayerGet Get a player of a broadcast

Get the details of a specific player and their games from a broadcast tournament.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @param playerId The unique player ID within the broadcast. This is usually their fideId.  If the player does not have a fideId, it is their name. Consult the [list of players for the broadcast](#tag/broadcasts/GET/broadcast/{broadcastTournamentId}/players) for which ID to use. 
 @return BroadcastsAPIBroadcastPlayerGetRequest
*/
func (a *BroadcastsAPIService) BroadcastPlayerGet(ctx context.Context, broadcastTournamentId string, playerId string) BroadcastsAPIBroadcastPlayerGetRequest {
	return BroadcastsAPIBroadcastPlayerGetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
		playerId: playerId,
	}
}

// Execute executes the request
//  @return BroadcastPlayerEntryWithFideAndGames
func (a *BroadcastsAPIService) BroadcastPlayerGetExecute(r BroadcastsAPIBroadcastPlayerGetRequest) (*BroadcastPlayerEntryWithFideAndGames, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastPlayerEntryWithFideAndGames
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastPlayerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/{broadcastTournamentId}/players/{playerId}"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playerId"+"}", url.PathEscape(parameterValueToString(r.playerId, "playerId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
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

type BroadcastsAPIBroadcastPlayersGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
}

func (r BroadcastsAPIBroadcastPlayersGetRequest) Execute() ([]BroadcastPlayerEntry, *http.Response, error) {
	return r.ApiService.BroadcastPlayersGetExecute(r)
}

/*
BroadcastPlayersGet Get players of a broadcast

Get the list of players of a broadcast tournament, if available.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @return BroadcastsAPIBroadcastPlayersGetRequest
*/
func (a *BroadcastsAPIService) BroadcastPlayersGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastPlayersGetRequest {
	return BroadcastsAPIBroadcastPlayersGetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return []BroadcastPlayerEntry
func (a *BroadcastsAPIService) BroadcastPlayersGetExecute(r BroadcastsAPIBroadcastPlayersGetRequest) ([]BroadcastPlayerEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []BroadcastPlayerEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastPlayersGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/{broadcastTournamentId}/players"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
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

type BroadcastsAPIBroadcastPushRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastRoundId string
	body *string
}

// The PGN. It can contain up to 100 games, separated by a double new line.
func (r BroadcastsAPIBroadcastPushRequest) Body(body string) BroadcastsAPIBroadcastPushRequest {
	r.body = &body
	return r
}

func (r BroadcastsAPIBroadcastPushRequest) Execute() (*BroadcastPgnPush, *http.Response, error) {
	return r.ApiService.BroadcastPushExecute(r)
}

/*
BroadcastPush Push PGN to a broadcast round

Update a broadcast with new PGN.
Only for broadcasts without a source URL.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastRoundId The broadcast round ID
 @return BroadcastsAPIBroadcastPushRequest
*/
func (a *BroadcastsAPIService) BroadcastPush(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastPushRequest {
	return BroadcastsAPIBroadcastPushRequest{
		ApiService: a,
		ctx: ctx,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return BroadcastPgnPush
func (a *BroadcastsAPIService) BroadcastPushExecute(r BroadcastsAPIBroadcastPushRequest) (*BroadcastPgnPush, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastPgnPush
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastPush")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/round/{broadcastRoundId}/push"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
	}
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v BroadcastPush400Response
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

type BroadcastsAPIBroadcastRoundCreateRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
	name *string
	syncUrl *string
	syncUrls *string
	syncIds *string
	syncUsers *string
	onlyRound *int32
	slices *string
	syncSource *string
	startsAt *int64
	startsAfterPrevious *bool
	delay *int32
	status *string
	rated *bool
	customScoring *BroadcastCustomScoring
	teamCustomScoring *BroadcastCustomPointsPerColor
	period *int32
}

// Name of the broadcast round. Example: &#x60;Round 1&#x60; 
func (r BroadcastsAPIBroadcastRoundCreateRequest) Name(name string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.name = &name
	return r
}

// URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/games.pgn &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastRoundCreateRequest) SyncUrl(syncUrl string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.syncUrl = &syncUrl
	return r
}

// URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastRoundCreateRequest) SyncUrls(syncUrls string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.syncUrls = &syncUrls
	return r
}

// Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) SyncIds(syncIds string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.syncIds = &syncIds
	return r
}

// Up to 100 Lichess usernames, separated by spaces 
func (r BroadcastsAPIBroadcastRoundCreateRequest) SyncUsers(syncUsers string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.syncUsers = &syncUsers
	return r
}

// Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: &#x60;&#x60;&#x60;txt [Round \\\&quot;3\\\&quot;] [Round \\\&quot;3.1\\\&quot;] &#x60;&#x60;&#x60; If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) OnlyRound(onlyRound int32) BroadcastsAPIBroadcastRoundCreateRequest {
	r.onlyRound = &onlyRound
	return r
}

// Select slices of the games  Optional. Select games based on their position in the source. &#x60;&#x60;&#x60;txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 &#x60;&#x60;&#x60; Slicing is done after filtering by round number.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) Slices(slices string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.slices = &slices
	return r
}

// Where the games come from. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) SyncSource(syncSource string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.syncSource = &syncSource
	return r
}

// Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: &#x60;1356998400070&#x60; 
func (r BroadcastsAPIBroadcastRoundCreateRequest) StartsAt(startsAt int64) BroadcastsAPIBroadcastRoundCreateRequest {
	r.startsAt = &startsAt
	return r
}

// The start date is unknown, and the round will start automatically when the previous round completes. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) StartsAfterPrevious(startsAfterPrevious bool) BroadcastsAPIBroadcastRoundCreateRequest {
	r.startsAfterPrevious = &startsAfterPrevious
	return r
}

// Delay in seconds for movements to appear on the broadcast. Leave it empty if you don&#39;t need it. Example: &#x60;900&#x60; (15 min) 
func (r BroadcastsAPIBroadcastRoundCreateRequest) Delay(delay int32) BroadcastsAPIBroadcastRoundCreateRequest {
	r.delay = &delay
	return r
}

// Lichess can usually detect the round status, but you can also set it manually if needed. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) Status(status string) BroadcastsAPIBroadcastRoundCreateRequest {
	r.status = &status
	return r
}

// Whether the round is used when calculating players&#39; rating changes.
func (r BroadcastsAPIBroadcastRoundCreateRequest) Rated(rated bool) BroadcastsAPIBroadcastRoundCreateRequest {
	r.rated = &rated
	return r
}

func (r BroadcastsAPIBroadcastRoundCreateRequest) CustomScoring(customScoring BroadcastCustomScoring) BroadcastsAPIBroadcastRoundCreateRequest {
	r.customScoring = &customScoring
	return r
}

// Scoring overrides for a team match win or draw. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) TeamCustomScoring(teamCustomScoring BroadcastCustomPointsPerColor) BroadcastsAPIBroadcastRoundCreateRequest {
	r.teamCustomScoring = &teamCustomScoring
	return r
}

// (Only for Admins) Waiting time for each poll. 
func (r BroadcastsAPIBroadcastRoundCreateRequest) Period(period int32) BroadcastsAPIBroadcastRoundCreateRequest {
	r.period = &period
	return r
}

func (r BroadcastsAPIBroadcastRoundCreateRequest) Execute() (*BroadcastRoundNew, *http.Response, error) {
	return r.ApiService.BroadcastRoundCreateExecute(r)
}

/*
BroadcastRoundCreate Create a broadcast round

Create a new broadcast round to relay external games.
This endpoint accepts the same form data as the web form.

Choose one between `syncUrl`, `syncUrls`, `syncIds` and `syncUsers`, if it is missing, the broadcast needs to be fed by [pushing PGN to it](#tag/broadcasts/POST/api/broadcast/round/{broadcastRoundId}/push)


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @return BroadcastsAPIBroadcastRoundCreateRequest
*/
func (a *BroadcastsAPIService) BroadcastRoundCreate(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastRoundCreateRequest {
	return BroadcastsAPIBroadcastRoundCreateRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return BroadcastRoundNew
func (a *BroadcastsAPIService) BroadcastRoundCreateExecute(r BroadcastsAPIBroadcastRoundCreateRequest) (*BroadcastRoundNew, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastRoundNew
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastRoundCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/{broadcastTournamentId}/new"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, reportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 80 {
		return localVarReturnValue, nil, reportError("name must have less than 80 elements")
	}
	if r.syncUrl == nil {
		return localVarReturnValue, nil, reportError("syncUrl is required and must be specified")
	}
	if r.syncUrls == nil {
		return localVarReturnValue, nil, reportError("syncUrls is required and must be specified")
	}
	if r.syncIds == nil {
		return localVarReturnValue, nil, reportError("syncIds is required and must be specified")
	}
	if r.syncUsers == nil {
		return localVarReturnValue, nil, reportError("syncUsers is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUrl", r.syncUrl, "", "")
	if r.onlyRound != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "onlyRound", r.onlyRound, "", "")
	}
	if r.slices != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "slices", r.slices, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUrls", r.syncUrls, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncIds", r.syncIds, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUsers", r.syncUsers, "", "")
	if r.syncSource != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "syncSource", r.syncSource, "", "")
	}
	if r.startsAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAt", r.startsAt, "", "")
	}
	if r.startsAfterPrevious != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAfterPrevious", r.startsAfterPrevious, "", "")
	}
	if r.delay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delay", r.delay, "", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "status", r.status, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.customScoring != nil {
		paramJson, err := parameterToJson(*r.customScoring)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("customScoring", paramJson)
	}
	if r.teamCustomScoring != nil {
		paramJson, err := parameterToJson(*r.teamCustomScoring)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("teamCustomScoring", paramJson)
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "period", r.period, "", "")
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

type BroadcastsAPIBroadcastRoundGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentSlug string
	broadcastRoundSlug string
	broadcastRoundId string
}

func (r BroadcastsAPIBroadcastRoundGetRequest) Execute() (*BroadcastRound, *http.Response, error) {
	return r.ApiService.BroadcastRoundGetExecute(r)
}

/*
BroadcastRoundGet Get a broadcast round

Get information about a broadcast round.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentSlug The broadcast tournament slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
 @param broadcastRoundSlug The broadcast round slug. Only used for SEO, the slug can be safely replaced by `-`. Only the `broadcastRoundId` is actually used.
 @param broadcastRoundId The broadcast Round ID
 @return BroadcastsAPIBroadcastRoundGetRequest
*/
func (a *BroadcastsAPIService) BroadcastRoundGet(ctx context.Context, broadcastTournamentSlug string, broadcastRoundSlug string, broadcastRoundId string) BroadcastsAPIBroadcastRoundGetRequest {
	return BroadcastsAPIBroadcastRoundGetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentSlug: broadcastTournamentSlug,
		broadcastRoundSlug: broadcastRoundSlug,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return BroadcastRound
func (a *BroadcastsAPIService) BroadcastRoundGetExecute(r BroadcastsAPIBroadcastRoundGetRequest) (*BroadcastRound, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastRound
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastRoundGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/{broadcastTournamentSlug}/{broadcastRoundSlug}/{broadcastRoundId}"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentSlug"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentSlug, "broadcastTournamentSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundSlug"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundSlug, "broadcastRoundSlug")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
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

type BroadcastsAPIBroadcastRoundPgnRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastRoundId string
	clocks *bool
	comments *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r BroadcastsAPIBroadcastRoundPgnRequest) Clocks(clocks bool) BroadcastsAPIBroadcastRoundPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60; 
func (r BroadcastsAPIBroadcastRoundPgnRequest) Comments(comments bool) BroadcastsAPIBroadcastRoundPgnRequest {
	r.comments = &comments
	return r
}

func (r BroadcastsAPIBroadcastRoundPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BroadcastRoundPgnExecute(r)
}

/*
BroadcastRoundPgn Export one round as PGN

Download all games of a single round of a broadcast tournament in PGN format.
You *could* poll this endpoint to get updates about a tournament, but it would be slow,
and very inefficient.
Instead, consider [streaming the tournament](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn) to get
a new PGN every time a game is updated, in real-time.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastRoundId The round ID
 @return BroadcastsAPIBroadcastRoundPgnRequest
*/
func (a *BroadcastsAPIService) BroadcastRoundPgn(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundPgnRequest {
	return BroadcastsAPIBroadcastRoundPgnRequest{
		ApiService: a,
		ctx: ctx,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return string
func (a *BroadcastsAPIService) BroadcastRoundPgnExecute(r BroadcastsAPIBroadcastRoundPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastRoundPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/round/{broadcastRoundId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
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

type BroadcastsAPIBroadcastRoundResetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastRoundId string
}

func (r BroadcastsAPIBroadcastRoundResetRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BroadcastRoundResetExecute(r)
}

/*
BroadcastRoundReset Reset a broadcast round

Remove any games from the broadcast round and reset it to its initial state.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastRoundId The broadcast round ID
 @return BroadcastsAPIBroadcastRoundResetRequest
*/
func (a *BroadcastsAPIService) BroadcastRoundReset(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundResetRequest {
	return BroadcastsAPIBroadcastRoundResetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BroadcastsAPIService) BroadcastRoundResetExecute(r BroadcastsAPIBroadcastRoundResetRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastRoundReset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/round/{broadcastRoundId}/reset"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
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

type BroadcastsAPIBroadcastRoundUpdateRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastRoundId string
	name *string
	syncUrl *string
	syncUrls *string
	syncIds *string
	syncUsers *string
	patch *bool
	onlyRound *int32
	slices *string
	syncSource *string
	startsAt *int64
	startsAfterPrevious *bool
	delay *int32
	status *string
	rated *bool
	customScoring *BroadcastCustomScoring
	teamCustomScoring *BroadcastCustomPointsPerColor
	period *int32
}

// Name of the broadcast round. Example: &#x60;Round 1&#x60; 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Name(name string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.name = &name
	return r
}

// URL that Lichess will poll to get updates about the games. It must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/games.pgn &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) SyncUrl(syncUrl string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.syncUrl = &syncUrl
	return r
}

// URLs that Lichess will poll to get updates about the games, separated by newlines. They must be publicly accessible from the Internet.  Example: &#x60;&#x60;&#x60;txt https://myserver.org/myevent/round-10/game-1.pgn https://myserver.org/myevent/round-10/game-2.pgn &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) SyncUrls(syncUrls string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.syncUrls = &syncUrls
	return r
}

// Lichess game IDs - Up to 100 Lichess game IDs, separated by spaces. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) SyncIds(syncIds string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.syncIds = &syncIds
	return r
}

// Up to 100 Lichess usernames, separated by spaces 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) SyncUsers(syncUsers string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.syncUsers = &syncUsers
	return r
}

// Only update the provided fields, leaving others unchanged
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Patch(patch bool) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.patch = &patch
	return r
}

// Filter games by round number  Optional, only keep games from the source that match a round number. It uses the PGN **Round** tag. These would match round 3: &#x60;&#x60;&#x60;txt [Round \\\&quot;3\\\&quot;] [Round \\\&quot;3.1\\\&quot;] &#x60;&#x60;&#x60; If you set a round number, then games without a **Round** tag are dropped.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) OnlyRound(onlyRound int32) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.onlyRound = &onlyRound
	return r
}

// Select slices of the games  Optional. Select games based on their position in the source. &#x60;&#x60;&#x60;txt 1           only select the first board 1-4         only select the first 4 boards 1,2,3,4     same as above, first 4 boards 11-15,21-25 boards 11 to 15, and boards 21 to 25 2,3,7-9     boards 2, 3, 7, 8, and 9 &#x60;&#x60;&#x60; Slicing is done after filtering by round number.  It only works if you chose &#x60;syncUrl&#x60; or &#x60;syncUrls&#x60; as the source. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Slices(slices string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.slices = &slices
	return r
}

// Where the games come from. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) SyncSource(syncSource string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.syncSource = &syncSource
	return r
}

// Timestamp in milliseconds of broadcast round start. Leave empty to manually start the broadcast round. Example: &#x60;1356998400070&#x60; 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) StartsAt(startsAt int64) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.startsAt = &startsAt
	return r
}

// The start date is unknown, and the round will start automatically when the previous round completes. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) StartsAfterPrevious(startsAfterPrevious bool) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.startsAfterPrevious = &startsAfterPrevious
	return r
}

// Delay in seconds for movements to appear on the broadcast. Leave it empty if you don&#39;t need it. Example: &#x60;900&#x60; (15 min) 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Delay(delay int32) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.delay = &delay
	return r
}

// Lichess can usually detect the round status, but you can also set it manually if needed. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Status(status string) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.status = &status
	return r
}

// Whether the round is used when calculating players&#39; rating changes.
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Rated(rated bool) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.rated = &rated
	return r
}

func (r BroadcastsAPIBroadcastRoundUpdateRequest) CustomScoring(customScoring BroadcastCustomScoring) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.customScoring = &customScoring
	return r
}

// Scoring overrides for a team match win or draw. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) TeamCustomScoring(teamCustomScoring BroadcastCustomPointsPerColor) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.teamCustomScoring = &teamCustomScoring
	return r
}

// (Only for Admins) Waiting time for each poll. 
func (r BroadcastsAPIBroadcastRoundUpdateRequest) Period(period int32) BroadcastsAPIBroadcastRoundUpdateRequest {
	r.period = &period
	return r
}

func (r BroadcastsAPIBroadcastRoundUpdateRequest) Execute() (*BroadcastRound, *http.Response, error) {
	return r.ApiService.BroadcastRoundUpdateExecute(r)
}

/*
BroadcastRoundUpdate Update a broadcast round

Update information about a broadcast round.
This endpoint accepts the same form data as the web form.
All fields must be populated with data. Missing fields will override the broadcast with empty data.
For instance, if you omit `startDate`, then any pre-existing start date will be removed.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastRoundId The broadcast round ID
 @return BroadcastsAPIBroadcastRoundUpdateRequest
*/
func (a *BroadcastsAPIService) BroadcastRoundUpdate(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastRoundUpdateRequest {
	return BroadcastsAPIBroadcastRoundUpdateRequest{
		ApiService: a,
		ctx: ctx,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return BroadcastRound
func (a *BroadcastsAPIService) BroadcastRoundUpdateExecute(r BroadcastsAPIBroadcastRoundUpdateRequest) (*BroadcastRound, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastRound
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastRoundUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/round/{broadcastRoundId}/edit"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, reportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 80 {
		return localVarReturnValue, nil, reportError("name must have less than 80 elements")
	}
	if r.syncUrl == nil {
		return localVarReturnValue, nil, reportError("syncUrl is required and must be specified")
	}
	if r.syncUrls == nil {
		return localVarReturnValue, nil, reportError("syncUrls is required and must be specified")
	}
	if r.syncIds == nil {
		return localVarReturnValue, nil, reportError("syncIds is required and must be specified")
	}
	if r.syncUsers == nil {
		return localVarReturnValue, nil, reportError("syncUsers is required and must be specified")
	}

	if r.patch != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "patch", r.patch, "form", "")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUrl", r.syncUrl, "", "")
	if r.onlyRound != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "onlyRound", r.onlyRound, "", "")
	}
	if r.slices != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "slices", r.slices, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUrls", r.syncUrls, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncIds", r.syncIds, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "syncUsers", r.syncUsers, "", "")
	if r.syncSource != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "syncSource", r.syncSource, "", "")
	}
	if r.startsAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAt", r.startsAt, "", "")
	}
	if r.startsAfterPrevious != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAfterPrevious", r.startsAfterPrevious, "", "")
	}
	if r.delay != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "delay", r.delay, "", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "status", r.status, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.customScoring != nil {
		paramJson, err := parameterToJson(*r.customScoring)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("customScoring", paramJson)
	}
	if r.teamCustomScoring != nil {
		paramJson, err := parameterToJson(*r.teamCustomScoring)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("teamCustomScoring", paramJson)
	}
	if r.period != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "period", r.period, "", "")
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

type BroadcastsAPIBroadcastStreamGroupPgnRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastGroupId string
	clocks *bool
	comments *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r BroadcastsAPIBroadcastStreamGroupPgnRequest) Clocks(clocks bool) BroadcastsAPIBroadcastStreamGroupPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60; 
func (r BroadcastsAPIBroadcastStreamGroupPgnRequest) Comments(comments bool) BroadcastsAPIBroadcastStreamGroupPgnRequest {
	r.comments = &comments
	return r
}

func (r BroadcastsAPIBroadcastStreamGroupPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BroadcastStreamGroupPgnExecute(r)
}

/*
BroadcastStreamGroupPgn Stream ongoing broadcast rounds of a group as PGN

For a given broadcast group ([example](https://lichess.org/broadcast/fide-world-cadets-cup-2026/albQx5zq)),
selects all the ongoing and recently finished rounds, and sends all games of these rounds in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the rounds.
This is the best way to get updates about an ongoing broadcast with multiple concurrent rounds.
To stream a single round, use [this endpoint instead](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastGroupId The broadcast group ID. It's the last past of the group URL, which can be found on the broadcast page overview section. [Example](https://lichess.org/broadcast/fide-world-cadets-cup-2026-o12/round-9/FUQFT39K) the group URL is  https://lichess.org/broadcast/fide-world-cadets-cup-2026/albQx5zq and therefore the group ID is `albQx5zq`. 
 @return BroadcastsAPIBroadcastStreamGroupPgnRequest
*/
func (a *BroadcastsAPIService) BroadcastStreamGroupPgn(ctx context.Context, broadcastGroupId string) BroadcastsAPIBroadcastStreamGroupPgnRequest {
	return BroadcastsAPIBroadcastStreamGroupPgnRequest{
		ApiService: a,
		ctx: ctx,
		broadcastGroupId: broadcastGroupId,
	}
}

// Execute executes the request
//  @return string
func (a *BroadcastsAPIService) BroadcastStreamGroupPgnExecute(r BroadcastsAPIBroadcastStreamGroupPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastStreamGroupPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/broadcast/group/{broadcastGroupId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastGroupId"+"}", url.PathEscape(parameterValueToString(r.broadcastGroupId, "broadcastGroupId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastGroupId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastGroupId must have at least 8 elements")
	}
	if strlen(r.broadcastGroupId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastGroupId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
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

type BroadcastsAPIBroadcastStreamRoundPgnRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastRoundId string
	clocks *bool
	comments *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r BroadcastsAPIBroadcastStreamRoundPgnRequest) Clocks(clocks bool) BroadcastsAPIBroadcastStreamRoundPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60; 
func (r BroadcastsAPIBroadcastStreamRoundPgnRequest) Comments(comments bool) BroadcastsAPIBroadcastStreamRoundPgnRequest {
	r.comments = &comments
	return r
}

func (r BroadcastsAPIBroadcastStreamRoundPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BroadcastStreamRoundPgnExecute(r)
}

/*
BroadcastStreamRoundPgn Stream an ongoing broadcast round as PGN

This streaming endpoint first sends all games of a broadcast round in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the round.
This is the best way to get updates about an ongoing round. Streaming means no polling,
and no pollings means no latency, and minimum impact on the server.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastRoundId The broadcast round ID
 @return BroadcastsAPIBroadcastStreamRoundPgnRequest
*/
func (a *BroadcastsAPIService) BroadcastStreamRoundPgn(ctx context.Context, broadcastRoundId string) BroadcastsAPIBroadcastStreamRoundPgnRequest {
	return BroadcastsAPIBroadcastStreamRoundPgnRequest{
		ApiService: a,
		ctx: ctx,
		broadcastRoundId: broadcastRoundId,
	}
}

// Execute executes the request
//  @return string
func (a *BroadcastsAPIService) BroadcastStreamRoundPgnExecute(r BroadcastsAPIBroadcastStreamRoundPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastStreamRoundPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/broadcast/round/{broadcastRoundId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastRoundId"+"}", url.PathEscape(parameterValueToString(r.broadcastRoundId, "broadcastRoundId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastRoundId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have at least 8 elements")
	}
	if strlen(r.broadcastRoundId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastRoundId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
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

type BroadcastsAPIBroadcastStreamTourPgnRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTourId string
	clocks *bool
	comments *bool
}

// Include clock comments in the PGN moves, when available. Example: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; 
func (r BroadcastsAPIBroadcastStreamTourPgnRequest) Clocks(clocks bool) BroadcastsAPIBroadcastStreamTourPgnRequest {
	r.clocks = &clocks
	return r
}

// Include analysis comments in the PGN moves, when available. Example: &#x60;12. Bxf6 { [%eval 0.23] }&#x60; 
func (r BroadcastsAPIBroadcastStreamTourPgnRequest) Comments(comments bool) BroadcastsAPIBroadcastStreamTourPgnRequest {
	r.comments = &comments
	return r
}

func (r BroadcastsAPIBroadcastStreamTourPgnRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.BroadcastStreamTourPgnExecute(r)
}

/*
BroadcastStreamTourPgn Stream ongoing broadcast rounds of a tournament as PGN

For a given broadcast tournament ([example](https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg)),
selects all the ongoing and recently finished rounds, and sends all games of these rounds in PGN format.
Then, it waits for new moves to be played. As soon as it happens, the entire PGN of the game is sent to the stream.
The stream will also send PGNs when games are added to the rounds.
This is the best way to get updates about an ongoing broadcast tournament across all its rounds.
To stream a single round, use [this endpoint instead](#tag/broadcasts/GET/api/stream/broadcast/round/{broadcastRoundId}.pgn).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTourId The broadcast tournament ID. It's the last past of the tournament URL, which can be found on the broadcast page overview section. [Example](https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg) the tournament URL is  https://lichess.org/broadcast/sparkassen-chess-trophy-2026-open-a/jfEpUuzg and therefore the tournament ID is `jfEpUuzg `. 
 @return BroadcastsAPIBroadcastStreamTourPgnRequest
*/
func (a *BroadcastsAPIService) BroadcastStreamTourPgn(ctx context.Context, broadcastTourId string) BroadcastsAPIBroadcastStreamTourPgnRequest {
	return BroadcastsAPIBroadcastStreamTourPgnRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTourId: broadcastTourId,
	}
}

// Execute executes the request
//  @return string
func (a *BroadcastsAPIService) BroadcastStreamTourPgnExecute(r BroadcastsAPIBroadcastStreamTourPgnRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastStreamTourPgn")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/stream/broadcast/tour/{broadcastTourId}.pgn"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTourId"+"}", url.PathEscape(parameterValueToString(r.broadcastTourId, "broadcastTourId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTourId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTourId must have at least 8 elements")
	}
	if strlen(r.broadcastTourId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTourId must have less than 8 elements")
	}

	if r.clocks != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", r.clocks, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "clocks", defaultValue, "form", "")
		r.clocks = &defaultValue
	}
	if r.comments != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", r.comments, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "comments", defaultValue, "form", "")
		r.comments = &defaultValue
	}
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

type BroadcastsAPIBroadcastTeamLeaderboardGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
}

func (r BroadcastsAPIBroadcastTeamLeaderboardGetRequest) Execute() ([]BroadcastTeamLeaderboardEntry, *http.Response, error) {
	return r.ApiService.BroadcastTeamLeaderboardGetExecute(r)
}

/*
BroadcastTeamLeaderboardGet Get the team leaderboard of a broadcast

Get the team leaderboard of a broadcast tournament, if available.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @return BroadcastsAPIBroadcastTeamLeaderboardGetRequest
*/
func (a *BroadcastsAPIService) BroadcastTeamLeaderboardGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTeamLeaderboardGetRequest {
	return BroadcastsAPIBroadcastTeamLeaderboardGetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return []BroadcastTeamLeaderboardEntry
func (a *BroadcastsAPIService) BroadcastTeamLeaderboardGetExecute(r BroadcastsAPIBroadcastTeamLeaderboardGetRequest) ([]BroadcastTeamLeaderboardEntry, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []BroadcastTeamLeaderboardEntry
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastTeamLeaderboardGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/{broadcastTournamentId}/teams/standings"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
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

type BroadcastsAPIBroadcastTourCreateRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	name *string
	info *BroadcastTourInfo
	markdown *string
	showScores *bool
	showRatingDiffs *bool
	teamTable *bool
	visibility *string
	players *string
	teams *string
	tier *int32
	tiebreaks *[]BroadcastTiebreakExtendedCode
	grouping *BroadcastFormGrouping
}

// Name of the broadcast tournament.  Example: &#x60;Sinquefield Cup&#x60; 
func (r BroadcastsAPIBroadcastTourCreateRequest) Name(name string) BroadcastsAPIBroadcastTourCreateRequest {
	r.name = &name
	return r
}

func (r BroadcastsAPIBroadcastTourCreateRequest) Info(info BroadcastTourInfo) BroadcastsAPIBroadcastTourCreateRequest {
	r.info = &info
	return r
}

// Optional long description of the broadcast. Markdown is supported. 
func (r BroadcastsAPIBroadcastTourCreateRequest) Markdown(markdown string) BroadcastsAPIBroadcastTourCreateRequest {
	r.markdown = &markdown
	return r
}

// Show players scores based on game results 
func (r BroadcastsAPIBroadcastTourCreateRequest) ShowScores(showScores bool) BroadcastsAPIBroadcastTourCreateRequest {
	r.showScores = &showScores
	return r
}

// Show player&#39;s rating diffs 
func (r BroadcastsAPIBroadcastTourCreateRequest) ShowRatingDiffs(showRatingDiffs bool) BroadcastsAPIBroadcastTourCreateRequest {
	r.showRatingDiffs = &showRatingDiffs
	return r
}

// Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags. 
func (r BroadcastsAPIBroadcastTourCreateRequest) TeamTable(teamTable bool) BroadcastsAPIBroadcastTourCreateRequest {
	r.teamTable = &teamTable
	return r
}

// Who can view the broadcast. * &#x60;public&#x60;: Default. Anyone can view the broadcast * &#x60;unlisted&#x60;: Only people with the link can view the broadcast * &#x60;private&#x60;: Only the broadcast owner(s) can view the broadcast 
func (r BroadcastsAPIBroadcastTourCreateRequest) Visibility(visibility string) BroadcastsAPIBroadcastTourCreateRequest {
	r.visibility = &visibility
	return r
}

// Optional replace player names, ratings and titles.  One line per player, formatted as such:  &#x60;&#x60;&#x60;txt player name / FIDE ID &#x60;&#x60;&#x60;  Example:  &#x60;&#x60;&#x60;txt Magnus Carlsen / 1503014 &#x60;&#x60;&#x60;  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\&quot;Jorge Rick Vito\\\&quot; will match \\\&quot;Jorge Rick\\\&quot;, \\\&quot;jorge vito\\\&quot;, \\\&quot;Rick, Vito\\\&quot;, etc.  If the player is NM or WNM, you can:  &#x60;&#x60;&#x60;txt player name / FIDE ID / title &#x60;&#x60;&#x60;  Alternatively, you may set tags manually, like so:  &#x60;&#x60;&#x60;txt player name / rating / title / new name &#x60;&#x60;&#x60;  All values are optional. Example: &#x60;&#x60;&#x60;txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastTourCreateRequest) Players(players string) BroadcastsAPIBroadcastTourCreateRequest {
	r.players = &players
	return r
}

// Optional: assign players to teams  One line per player, formatted as such: &#x60;&#x60;&#x60;txt Team name; Fide Id or Player name &#x60;&#x60;&#x60;  Example: &#x60;&#x60;&#x60;txt Team Cats ; 3408230 Team Dogs ; Scooby Doo &#x60;&#x60;&#x60;  By default the PGN tags WhiteTeam and BlackTeam are used. 
func (r BroadcastsAPIBroadcastTourCreateRequest) Teams(teams string) BroadcastsAPIBroadcastTourCreateRequest {
	r.teams = &teams
	return r
}

// Optional, for Lichess admins only, used to feature on /broadcast.  * &#x60;3&#x60; for Official: normal tier * &#x60;4&#x60; for Official: high tier * &#x60;5&#x60; for Official: best tier 
func (r BroadcastsAPIBroadcastTourCreateRequest) Tier(tier int32) BroadcastsAPIBroadcastTourCreateRequest {
	r.tier = &tier
	return r
}

func (r BroadcastsAPIBroadcastTourCreateRequest) Tiebreaks(tiebreaks []BroadcastTiebreakExtendedCode) BroadcastsAPIBroadcastTourCreateRequest {
	r.tiebreaks = &tiebreaks
	return r
}

func (r BroadcastsAPIBroadcastTourCreateRequest) Grouping(grouping BroadcastFormGrouping) BroadcastsAPIBroadcastTourCreateRequest {
	r.grouping = &grouping
	return r
}

func (r BroadcastsAPIBroadcastTourCreateRequest) Execute() (*BroadcastWithRounds, *http.Response, error) {
	return r.ApiService.BroadcastTourCreateExecute(r)
}

/*
BroadcastTourCreate Create a broadcast tournament

Create a new broadcast tournament to relay external games.
This endpoint accepts the same form data as the [web form](https://lichess.org/broadcast/new).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BroadcastsAPIBroadcastTourCreateRequest
*/
func (a *BroadcastsAPIService) BroadcastTourCreate(ctx context.Context) BroadcastsAPIBroadcastTourCreateRequest {
	return BroadcastsAPIBroadcastTourCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BroadcastWithRounds
func (a *BroadcastsAPIService) BroadcastTourCreateExecute(r BroadcastsAPIBroadcastTourCreateRequest) (*BroadcastWithRounds, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastWithRounds
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastTourCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/new"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, reportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 80 {
		return localVarReturnValue, nil, reportError("name must have less than 80 elements")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	if r.info != nil {
		paramJson, err := parameterToJson(*r.info)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("info", paramJson)
	}
	if r.markdown != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "markdown", r.markdown, "", "")
	}
	if r.showScores != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "showScores", r.showScores, "", "")
	}
	if r.showRatingDiffs != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "showRatingDiffs", r.showRatingDiffs, "", "")
	}
	if r.teamTable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "teamTable", r.teamTable, "", "")
	}
	if r.visibility != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "visibility", r.visibility, "", "")
	}
	if r.players != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "players", r.players, "", "")
	}
	if r.teams != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "teams", r.teams, "", "")
	}
	if r.tier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tier", r.tier, "", "")
	}
	if r.tiebreaks != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tiebreaks", r.tiebreaks, "", "csv")
	}
	if r.grouping != nil {
		paramJson, err := parameterToJson(*r.grouping)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("grouping", paramJson)
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

type BroadcastsAPIBroadcastTourGetRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
}

func (r BroadcastsAPIBroadcastTourGetRequest) Execute() (*BroadcastWithRoundsAndFullGroup, *http.Response, error) {
	return r.ApiService.BroadcastTourGetExecute(r)
}

/*
BroadcastTourGet Get a broadcast tournament

Get information about a broadcast tournament.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast tournament ID
 @return BroadcastsAPIBroadcastTourGetRequest
*/
func (a *BroadcastsAPIService) BroadcastTourGet(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTourGetRequest {
	return BroadcastsAPIBroadcastTourGetRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return BroadcastWithRoundsAndFullGroup
func (a *BroadcastsAPIService) BroadcastTourGetExecute(r BroadcastsAPIBroadcastTourGetRequest) (*BroadcastWithRoundsAndFullGroup, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastWithRoundsAndFullGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastTourGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/{broadcastTournamentId}"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
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

type BroadcastsAPIBroadcastTourUpdateRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	broadcastTournamentId string
	name *string
	info *BroadcastTourInfo
	markdown *string
	showScores *bool
	showRatingDiffs *bool
	teamTable *bool
	visibility *string
	players *string
	teams *string
	tier *int32
	tiebreaks *[]BroadcastTiebreakExtendedCode
	grouping *BroadcastFormGrouping
}

// Name of the broadcast tournament.  Example: &#x60;Sinquefield Cup&#x60; 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Name(name string) BroadcastsAPIBroadcastTourUpdateRequest {
	r.name = &name
	return r
}

func (r BroadcastsAPIBroadcastTourUpdateRequest) Info(info BroadcastTourInfo) BroadcastsAPIBroadcastTourUpdateRequest {
	r.info = &info
	return r
}

// Optional long description of the broadcast. Markdown is supported. 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Markdown(markdown string) BroadcastsAPIBroadcastTourUpdateRequest {
	r.markdown = &markdown
	return r
}

// Show players scores based on game results 
func (r BroadcastsAPIBroadcastTourUpdateRequest) ShowScores(showScores bool) BroadcastsAPIBroadcastTourUpdateRequest {
	r.showScores = &showScores
	return r
}

// Show player&#39;s rating diffs 
func (r BroadcastsAPIBroadcastTourUpdateRequest) ShowRatingDiffs(showRatingDiffs bool) BroadcastsAPIBroadcastTourUpdateRequest {
	r.showRatingDiffs = &showRatingDiffs
	return r
}

// Show a team leaderboard. Requires WhiteTeam and BlackTeam PGN tags. 
func (r BroadcastsAPIBroadcastTourUpdateRequest) TeamTable(teamTable bool) BroadcastsAPIBroadcastTourUpdateRequest {
	r.teamTable = &teamTable
	return r
}

// Who can view the broadcast. * &#x60;public&#x60;: Default. Anyone can view the broadcast * &#x60;unlisted&#x60;: Only people with the link can view the broadcast * &#x60;private&#x60;: Only the broadcast owner(s) can view the broadcast 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Visibility(visibility string) BroadcastsAPIBroadcastTourUpdateRequest {
	r.visibility = &visibility
	return r
}

// Optional replace player names, ratings and titles.  One line per player, formatted as such:  &#x60;&#x60;&#x60;txt player name / FIDE ID &#x60;&#x60;&#x60;  Example:  &#x60;&#x60;&#x60;txt Magnus Carlsen / 1503014 &#x60;&#x60;&#x60;  Player names ignore case and punctuation, and match all possible combinations of 2 words: \\\&quot;Jorge Rick Vito\\\&quot; will match \\\&quot;Jorge Rick\\\&quot;, \\\&quot;jorge vito\\\&quot;, \\\&quot;Rick, Vito\\\&quot;, etc.  If the player is NM or WNM, you can:  &#x60;&#x60;&#x60;txt player name / FIDE ID / title &#x60;&#x60;&#x60;  Alternatively, you may set tags manually, like so:  &#x60;&#x60;&#x60;txt player name / rating / title / new name &#x60;&#x60;&#x60;  All values are optional. Example: &#x60;&#x60;&#x60;txt Magnus Carlsen / 2863 / GM YouGotLittUp / 1890 / / Louis Litt &#x60;&#x60;&#x60; 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Players(players string) BroadcastsAPIBroadcastTourUpdateRequest {
	r.players = &players
	return r
}

// Optional: assign players to teams  One line per player, formatted as such: &#x60;&#x60;&#x60;txt Team name; Fide Id or Player name &#x60;&#x60;&#x60;  Example: &#x60;&#x60;&#x60;txt Team Cats ; 3408230 Team Dogs ; Scooby Doo &#x60;&#x60;&#x60;  By default the PGN tags WhiteTeam and BlackTeam are used. 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Teams(teams string) BroadcastsAPIBroadcastTourUpdateRequest {
	r.teams = &teams
	return r
}

// Optional, for Lichess admins only, used to feature on /broadcast.  * &#x60;3&#x60; for Official: normal tier * &#x60;4&#x60; for Official: high tier * &#x60;5&#x60; for Official: best tier 
func (r BroadcastsAPIBroadcastTourUpdateRequest) Tier(tier int32) BroadcastsAPIBroadcastTourUpdateRequest {
	r.tier = &tier
	return r
}

func (r BroadcastsAPIBroadcastTourUpdateRequest) Tiebreaks(tiebreaks []BroadcastTiebreakExtendedCode) BroadcastsAPIBroadcastTourUpdateRequest {
	r.tiebreaks = &tiebreaks
	return r
}

func (r BroadcastsAPIBroadcastTourUpdateRequest) Grouping(grouping BroadcastFormGrouping) BroadcastsAPIBroadcastTourUpdateRequest {
	r.grouping = &grouping
	return r
}

func (r BroadcastsAPIBroadcastTourUpdateRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.BroadcastTourUpdateExecute(r)
}

/*
BroadcastTourUpdate Update your broadcast tournament

Update information about a broadcast tournament that you created.
This endpoint accepts the same form data as the web form.
All fields must be populated with data. Missing fields will override the broadcast with empty data.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param broadcastTournamentId The broadcast ID
 @return BroadcastsAPIBroadcastTourUpdateRequest
*/
func (a *BroadcastsAPIService) BroadcastTourUpdate(ctx context.Context, broadcastTournamentId string) BroadcastsAPIBroadcastTourUpdateRequest {
	return BroadcastsAPIBroadcastTourUpdateRequest{
		ApiService: a,
		ctx: ctx,
		broadcastTournamentId: broadcastTournamentId,
	}
}

// Execute executes the request
//  @return Ok
func (a *BroadcastsAPIService) BroadcastTourUpdateExecute(r BroadcastsAPIBroadcastTourUpdateRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastTourUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/broadcast/{broadcastTournamentId}/edit"
	localVarPath = strings.Replace(localVarPath, "{"+"broadcastTournamentId"+"}", url.PathEscape(parameterValueToString(r.broadcastTournamentId, "broadcastTournamentId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.broadcastTournamentId) < 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have at least 8 elements")
	}
	if strlen(r.broadcastTournamentId) > 8 {
		return localVarReturnValue, nil, reportError("broadcastTournamentId must have less than 8 elements")
	}
	if r.name == nil {
		return localVarReturnValue, nil, reportError("name is required and must be specified")
	}
	if strlen(*r.name) < 3 {
		return localVarReturnValue, nil, reportError("name must have at least 3 elements")
	}
	if strlen(*r.name) > 80 {
		return localVarReturnValue, nil, reportError("name must have less than 80 elements")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	if r.info != nil {
		paramJson, err := parameterToJson(*r.info)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("info", paramJson)
	}
	if r.markdown != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "markdown", r.markdown, "", "")
	}
	if r.showScores != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "showScores", r.showScores, "", "")
	}
	if r.showRatingDiffs != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "showRatingDiffs", r.showRatingDiffs, "", "")
	}
	if r.teamTable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "teamTable", r.teamTable, "", "")
	}
	if r.visibility != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "visibility", r.visibility, "", "")
	}
	if r.players != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "players", r.players, "", "")
	}
	if r.teams != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "teams", r.teams, "", "")
	}
	if r.tier != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tier", r.tier, "", "")
	}
	if r.tiebreaks != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "tiebreaks", r.tiebreaks, "", "csv")
	}
	if r.grouping != nil {
		paramJson, err := parameterToJson(*r.grouping)
		if err != nil {
			return localVarReturnValue, nil, err
		}
		localVarFormParams.Add("grouping", paramJson)
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

type BroadcastsAPIBroadcastsByUserRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	username string
	page *int32
	html *bool
}

func (r BroadcastsAPIBroadcastsByUserRequest) Page(page int32) BroadcastsAPIBroadcastsByUserRequest {
	r.page = &page
	return r
}

// Convert the \&quot;description\&quot; field from markdown to HTML
func (r BroadcastsAPIBroadcastsByUserRequest) Html(html bool) BroadcastsAPIBroadcastsByUserRequest {
	r.html = &html
	return r
}

func (r BroadcastsAPIBroadcastsByUserRequest) Execute() (*BroadcastsByUser200Response, *http.Response, error) {
	return r.ApiService.BroadcastsByUserExecute(r)
}

/*
BroadcastsByUser Get broadcasts created by a user

Get all incoming, ongoing, and finished official broadcasts.
The broadcasts are sorted by created date, most recent first.

If you are authenticated as the user whose broadcasts you are requesting, you will also see your private and unlisted broadcasts.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return BroadcastsAPIBroadcastsByUserRequest
*/
func (a *BroadcastsAPIService) BroadcastsByUser(ctx context.Context, username string) BroadcastsAPIBroadcastsByUserRequest {
	return BroadcastsAPIBroadcastsByUserRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return BroadcastsByUser200Response
func (a *BroadcastsAPIService) BroadcastsByUserExecute(r BroadcastsAPIBroadcastsByUserRequest) (*BroadcastsByUser200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastsByUser200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastsByUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/by/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.html != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "html", r.html, "form", "")
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

type BroadcastsAPIBroadcastsOfficialRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	nb *int32
	html *bool
	live *bool
}

// Max number of broadcasts to fetch
func (r BroadcastsAPIBroadcastsOfficialRequest) Nb(nb int32) BroadcastsAPIBroadcastsOfficialRequest {
	r.nb = &nb
	return r
}

// Convert the \&quot;description\&quot; field from markdown to HTML
func (r BroadcastsAPIBroadcastsOfficialRequest) Html(html bool) BroadcastsAPIBroadcastsOfficialRequest {
	r.html = &html
	return r
}

// [Filter] only broadcasts where a round is ongoing, i.e. started and not finished
func (r BroadcastsAPIBroadcastsOfficialRequest) Live(live bool) BroadcastsAPIBroadcastsOfficialRequest {
	r.live = &live
	return r
}

func (r BroadcastsAPIBroadcastsOfficialRequest) Execute() (*BroadcastWithRounds, *http.Response, error) {
	return r.ApiService.BroadcastsOfficialExecute(r)
}

/*
BroadcastsOfficial Get official broadcasts

Returns active (a round is scheduled or ongoing) official broadcasts sorted by tier. 
After that, returns finished broadcasts sorted by most recent sync time.
Broadcasts are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BroadcastsAPIBroadcastsOfficialRequest
*/
func (a *BroadcastsAPIService) BroadcastsOfficial(ctx context.Context) BroadcastsAPIBroadcastsOfficialRequest {
	return BroadcastsAPIBroadcastsOfficialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BroadcastWithRounds
func (a *BroadcastsAPIService) BroadcastsOfficialExecute(r BroadcastsAPIBroadcastsOfficialRequest) (*BroadcastWithRounds, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastWithRounds
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastsOfficial")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 20
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", defaultValue, "form", "")
		r.nb = &defaultValue
	}
	if r.html != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "html", r.html, "form", "")
	}
	if r.live != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "live", r.live, "form", "")
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

type BroadcastsAPIBroadcastsSearchRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	page *int32
	q *string
}

// Which page to fetch.
func (r BroadcastsAPIBroadcastsSearchRequest) Page(page int32) BroadcastsAPIBroadcastsSearchRequest {
	r.page = &page
	return r
}

// Search term
func (r BroadcastsAPIBroadcastsSearchRequest) Q(q string) BroadcastsAPIBroadcastsSearchRequest {
	r.q = &q
	return r
}

func (r BroadcastsAPIBroadcastsSearchRequest) Execute() (*BroadcastsSearch200Response, *http.Response, error) {
	return r.ApiService.BroadcastsSearchExecute(r)
}

/*
BroadcastsSearch Search broadcasts

Search across recent official broadcasts.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BroadcastsAPIBroadcastsSearchRequest
*/
func (a *BroadcastsAPIService) BroadcastsSearch(ctx context.Context) BroadcastsAPIBroadcastsSearchRequest {
	return BroadcastsAPIBroadcastsSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BroadcastsSearch200Response
func (a *BroadcastsAPIService) BroadcastsSearchExecute(r BroadcastsAPIBroadcastsSearchRequest) (*BroadcastsSearch200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastsSearch200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastsSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.q != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "q", r.q, "form", "")
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

type BroadcastsAPIBroadcastsTopRequest struct {
	ctx context.Context
	ApiService BroadcastsAPI
	page *int32
	html *bool
}

// Which page to fetch. Only page 1 has \&quot;active\&quot; broadcasts.
func (r BroadcastsAPIBroadcastsTopRequest) Page(page int32) BroadcastsAPIBroadcastsTopRequest {
	r.page = &page
	return r
}

// Convert the \&quot;description\&quot; field from markdown to HTML
func (r BroadcastsAPIBroadcastsTopRequest) Html(html bool) BroadcastsAPIBroadcastsTopRequest {
	r.html = &html
	return r
}

func (r BroadcastsAPIBroadcastsTopRequest) Execute() (*BroadcastTop, *http.Response, error) {
	return r.ApiService.BroadcastsTopExecute(r)
}

/*
BroadcastsTop Get paginated top broadcast previews

The same data, in the same order, as can be seen on [https://lichess.org/broadcast](/broadcast).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return BroadcastsAPIBroadcastsTopRequest
*/
func (a *BroadcastsAPIService) BroadcastsTop(ctx context.Context) BroadcastsAPIBroadcastsTopRequest {
	return BroadcastsAPIBroadcastsTopRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return BroadcastTop
func (a *BroadcastsAPIService) BroadcastsTopExecute(r BroadcastsAPIBroadcastsTopRequest) (*BroadcastTop, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *BroadcastTop
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "BroadcastsAPIService.BroadcastsTop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/broadcast/top"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue int32 = 1
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", defaultValue, "form", "")
		r.page = &defaultValue
	}
	if r.html != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "html", r.html, "form", "")
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
