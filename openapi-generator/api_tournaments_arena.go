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


type TournamentsArenaAPI interface {

	/*
	ApiTeamArena Get team Arena tournaments

	Get all Arena tournaments relevant to a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId ID of the team
	@return TournamentsArenaAPIApiTeamArenaRequest
	*/
	ApiTeamArena(ctx context.Context, teamId string) TournamentsArenaAPIApiTeamArenaRequest

	// ApiTeamArenaExecute executes the request
	//  @return ApiTournament200ResponseCreatedInner
	ApiTeamArenaExecute(r TournamentsArenaAPIApiTeamArenaRequest) (*ApiTournament200ResponseCreatedInner, *http.Response, error)

	/*
	ApiTournament Get current tournaments

	Get recently active and finished tournaments.
This API is used to display the [Lichess tournament schedule](https://lichess.org/tournament).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TournamentsArenaAPIApiTournamentRequest
	*/
	ApiTournament(ctx context.Context) TournamentsArenaAPIApiTournamentRequest

	// ApiTournamentExecute executes the request
	//  @return ApiTournament200Response
	ApiTournamentExecute(r TournamentsArenaAPIApiTournamentRequest) (*ApiTournament200Response, *http.Response, error)

	/*
	ApiTournamentJoin Join an Arena tournament

	Join an Arena tournament, possibly with a password and/or a team.
Also unpauses if you had previously [paused](#tag/arena-tournaments/POST/api/tournament/{id}/withdraw) the tournament.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIApiTournamentJoinRequest
	*/
	ApiTournamentJoin(ctx context.Context, id string) TournamentsArenaAPIApiTournamentJoinRequest

	// ApiTournamentJoinExecute executes the request
	//  @return AccountKidPost200Response
	ApiTournamentJoinExecute(r TournamentsArenaAPIApiTournamentJoinRequest) (*AccountKidPost200Response, *http.Response, error)

	/*
	ApiTournamentPost Create a new Arena tournament

	Create a public or private Arena tournament.
This endpoint mirrors the form on <https://lichess.org/tournament/new>.
You can create up to 12 public tournaments per day, or 24 private tournaments.
A team battle can be created by specifying the `teamBattleByTeam` argument.
Additional restrictions:
  - clockTime + clockIncrement > 0
  - 15s and 0+1 variant tournaments cannot be rated
  - Clock time in comparison to tournament length must be reasonable: 3 <= (minutes * 60) / (96 * clockTime + 48 * clockIncrement + 15) <= 150


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TournamentsArenaAPIApiTournamentPostRequest
	*/
	ApiTournamentPost(ctx context.Context) TournamentsArenaAPIApiTournamentPostRequest

	// ApiTournamentPostExecute executes the request
	//  @return ApiTournamentPost200Response
	ApiTournamentPostExecute(r TournamentsArenaAPIApiTournamentPostRequest) (*ApiTournamentPost200Response, *http.Response, error)

	/*
	ApiTournamentTeamBattlePost Update a team battle

	Set the teams and number of leaders of a team battle.
To update the other attributes of a team battle, use the [tournament update endpoint](#tag/arena-tournaments/POST/api/tournament/{id}).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID
	@return TournamentsArenaAPIApiTournamentTeamBattlePostRequest
	*/
	ApiTournamentTeamBattlePost(ctx context.Context, id string) TournamentsArenaAPIApiTournamentTeamBattlePostRequest

	// ApiTournamentTeamBattlePostExecute executes the request
	//  @return Tournament200Response
	ApiTournamentTeamBattlePostExecute(r TournamentsArenaAPIApiTournamentTeamBattlePostRequest) (*Tournament200Response, *http.Response, error)

	/*
	ApiTournamentTerminate Terminate an Arena tournament

	Terminate an Arena tournament


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIApiTournamentTerminateRequest
	*/
	ApiTournamentTerminate(ctx context.Context, id string) TournamentsArenaAPIApiTournamentTerminateRequest

	// ApiTournamentTerminateExecute executes the request
	//  @return AccountKidPost200Response
	ApiTournamentTerminateExecute(r TournamentsArenaAPIApiTournamentTerminateRequest) (*AccountKidPost200Response, *http.Response, error)

	/*
	ApiTournamentUpdate Update an Arena tournament

	Update an Arena tournament.
Be mindful not to make important changes to ongoing tournaments.
Can be used to update a team battle.
Additional restrictions:
  - clockTime + clockIncrement > 0
  - 15s and 0+1 variant tournaments cannot be rated
  - Clock time in comparison to tournament length must be reasonable: 3 <= (minutes * 60) / (96 * clockTime + 48 * clockIncrement + 15) <= 150


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIApiTournamentUpdateRequest
	*/
	ApiTournamentUpdate(ctx context.Context, id string) TournamentsArenaAPIApiTournamentUpdateRequest

	// ApiTournamentUpdateExecute executes the request
	//  @return Tournament200Response
	ApiTournamentUpdateExecute(r TournamentsArenaAPIApiTournamentUpdateRequest) (*Tournament200Response, *http.Response, error)

	/*
	ApiTournamentWithdraw Pause or leave an Arena tournament

	Leave a future Arena tournament, or take a break on an ongoing Arena tournament.
It's possible to join again later. Points and streaks are preserved.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIApiTournamentWithdrawRequest
	*/
	ApiTournamentWithdraw(ctx context.Context, id string) TournamentsArenaAPIApiTournamentWithdrawRequest

	// ApiTournamentWithdrawExecute executes the request
	//  @return AccountKidPost200Response
	ApiTournamentWithdrawExecute(r TournamentsArenaAPIApiTournamentWithdrawRequest) (*AccountKidPost200Response, *http.Response, error)

	/*
	ApiUserNameTournamentCreated Get tournaments created by a user

	Get all tournaments created by a given user.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).
The stream is throttled, depending on who is making the request:
  - Anonymous request: 20 tournaments per second
  - [OAuth2 authenticated](#description/authentication) request: 30 tournaments per second
  - Authenticated, downloading your own tournaments: 50 tournaments per second


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The user whose created tournaments to fetch
	@return TournamentsArenaAPIApiUserNameTournamentCreatedRequest
	*/
	ApiUserNameTournamentCreated(ctx context.Context, username string) TournamentsArenaAPIApiUserNameTournamentCreatedRequest

	// ApiUserNameTournamentCreatedExecute executes the request
	//  @return ApiTournament200ResponseCreatedInner
	ApiUserNameTournamentCreatedExecute(r TournamentsArenaAPIApiUserNameTournamentCreatedRequest) (*ApiTournament200ResponseCreatedInner, *http.Response, error)

	/*
	ApiUserNameTournamentPlayed Get tournaments played by a user

	Get all tournaments played by a given user.
Tournaments are sorted by reverse chronological order of start date (last played first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).
The stream is throttled, depending on who is making the request:
  - Anonymous request: 20 tournaments per second
  - [OAuth2 authenticated](#description/authentication) request: 30 tournaments per second
  - Authenticated, downloading your own tournaments: 50 tournaments per second


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username The user whose played tournaments to fetch
	@return TournamentsArenaAPIApiUserNameTournamentPlayedRequest
	*/
	ApiUserNameTournamentPlayed(ctx context.Context, username string) TournamentsArenaAPIApiUserNameTournamentPlayedRequest

	// ApiUserNameTournamentPlayedExecute executes the request
	//  @return ApiUserNameTournamentPlayed200Response
	ApiUserNameTournamentPlayedExecute(r TournamentsArenaAPIApiUserNameTournamentPlayedRequest) (*ApiUserNameTournamentPlayed200Response, *http.Response, error)

	/*
	GamesByTournament Export games of an Arena tournament

	Download games of a tournament in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIGamesByTournamentRequest
	*/
	GamesByTournament(ctx context.Context, id string) TournamentsArenaAPIGamesByTournamentRequest

	// GamesByTournamentExecute executes the request
	//  @return string
	GamesByTournamentExecute(r TournamentsArenaAPIGamesByTournamentRequest) (string, *http.Response, error)

	/*
	ResultsByTournament Get results of an Arena tournament

	Players of an Arena tournament, with their score and performance, sorted by rank (best first).
**Players are streamed as [ndjson](#description/streaming-with-nd-json)**, i.e. one JSON object per line.
If called on an ongoing tournament, results can be inconsistent
due to ranking changes while the players are being streamed.
Use on finished tournaments for guaranteed consistency.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPIResultsByTournamentRequest
	*/
	ResultsByTournament(ctx context.Context, id string) TournamentsArenaAPIResultsByTournamentRequest

	// ResultsByTournamentExecute executes the request
	//  @return ResultsByTournament200Response
	ResultsByTournamentExecute(r TournamentsArenaAPIResultsByTournamentRequest) (*ResultsByTournament200Response, *http.Response, error)

	/*
	TeamsByTournament Get team standing of a team battle

	Teams of a team battle tournament, with top players, sorted by rank (best first).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPITeamsByTournamentRequest
	*/
	TeamsByTournament(ctx context.Context, id string) TournamentsArenaAPITeamsByTournamentRequest

	// TeamsByTournamentExecute executes the request
	//  @return TeamsByTournament200Response
	TeamsByTournamentExecute(r TournamentsArenaAPITeamsByTournamentRequest) (*TeamsByTournament200Response, *http.Response, error)

	/*
	Tournament Get info about an Arena tournament

	Get detailed info about recently finished, current, or upcoming tournament's duels, player standings, and other info.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsArenaAPITournamentRequest
	*/
	Tournament(ctx context.Context, id string) TournamentsArenaAPITournamentRequest

	// TournamentExecute executes the request
	//  @return Tournament200Response
	TournamentExecute(r TournamentsArenaAPITournamentRequest) (*Tournament200Response, *http.Response, error)
}

// TournamentsArenaAPIService TournamentsArenaAPI service
type TournamentsArenaAPIService service

type TournamentsArenaAPIApiTeamArenaRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	teamId string
	max *int32
	status *string
	createdBy *string
	name *string
}

// How many tournaments to download.
func (r TournamentsArenaAPIApiTeamArenaRequest) Max(max int32) TournamentsArenaAPIApiTeamArenaRequest {
	r.max = &max
	return r
}

// [Filter] Only arena tournaments in this current state. 
func (r TournamentsArenaAPIApiTeamArenaRequest) Status(status string) TournamentsArenaAPIApiTeamArenaRequest {
	r.status = &status
	return r
}

// [Filter] Only arena tournaments created by a given user. 
func (r TournamentsArenaAPIApiTeamArenaRequest) CreatedBy(createdBy string) TournamentsArenaAPIApiTeamArenaRequest {
	r.createdBy = &createdBy
	return r
}

// [Filter] Only arena tournaments with a given name. 
func (r TournamentsArenaAPIApiTeamArenaRequest) Name(name string) TournamentsArenaAPIApiTeamArenaRequest {
	r.name = &name
	return r
}

func (r TournamentsArenaAPIApiTeamArenaRequest) Execute() (*ApiTournament200ResponseCreatedInner, *http.Response, error) {
	return r.ApiService.ApiTeamArenaExecute(r)
}

/*
ApiTeamArena Get team Arena tournaments

Get all Arena tournaments relevant to a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId ID of the team
 @return TournamentsArenaAPIApiTeamArenaRequest
*/
func (a *TournamentsArenaAPIService) ApiTeamArena(ctx context.Context, teamId string) TournamentsArenaAPIApiTeamArenaRequest {
	return TournamentsArenaAPIApiTeamArenaRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return ApiTournament200ResponseCreatedInner
func (a *TournamentsArenaAPIService) ApiTeamArenaExecute(r TournamentsArenaAPIApiTeamArenaRequest) (*ApiTournament200ResponseCreatedInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiTournament200ResponseCreatedInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTeamArena")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/arena"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "form", "")
	} else {
		var defaultValue int32 = 100
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", defaultValue, "form", "")
		r.max = &defaultValue
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
	}
	if r.createdBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "createdBy", r.createdBy, "form", "")
	}
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
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

type TournamentsArenaAPIApiTournamentRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
}

func (r TournamentsArenaAPIApiTournamentRequest) Execute() (*ApiTournament200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentExecute(r)
}

/*
ApiTournament Get current tournaments

Get recently active and finished tournaments.
This API is used to display the [Lichess tournament schedule](https://lichess.org/tournament).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TournamentsArenaAPIApiTournamentRequest
*/
func (a *TournamentsArenaAPIService) ApiTournament(ctx context.Context) TournamentsArenaAPIApiTournamentRequest {
	return TournamentsArenaAPIApiTournamentRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiTournament200Response
func (a *TournamentsArenaAPIService) ApiTournamentExecute(r TournamentsArenaAPIApiTournamentRequest) (*ApiTournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiTournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournament")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament"

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

type TournamentsArenaAPIApiTournamentJoinRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	password *string
	team *string
	pairMeAsap *bool
}

// The tournament password, if one is required. Can also be a [user-specific entry code](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) generated and shared by the organizer. 
func (r TournamentsArenaAPIApiTournamentJoinRequest) Password(password string) TournamentsArenaAPIApiTournamentJoinRequest {
	r.password = &password
	return r
}

// The team to join the tournament with, for team battle tournaments
func (r TournamentsArenaAPIApiTournamentJoinRequest) Team(team string) TournamentsArenaAPIApiTournamentJoinRequest {
	r.team = &team
	return r
}

// If the tournament is started, attempt to pair the user, even if they are not connected to the tournament page. This expires after one minute, to avoid pairing a user who is long gone. You may call \\\&quot;join\\\&quot; again to extend the waiting. 
func (r TournamentsArenaAPIApiTournamentJoinRequest) PairMeAsap(pairMeAsap bool) TournamentsArenaAPIApiTournamentJoinRequest {
	r.pairMeAsap = &pairMeAsap
	return r
}

func (r TournamentsArenaAPIApiTournamentJoinRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentJoinExecute(r)
}

/*
ApiTournamentJoin Join an Arena tournament

Join an Arena tournament, possibly with a password and/or a team.
Also unpauses if you had previously [paused](#tag/arena-tournaments/POST/api/tournament/{id}/withdraw) the tournament.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIApiTournamentJoinRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentJoin(ctx context.Context, id string) TournamentsArenaAPIApiTournamentJoinRequest {
	return TournamentsArenaAPIApiTournamentJoinRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *TournamentsArenaAPIService) ApiTournamentJoinExecute(r TournamentsArenaAPIApiTournamentJoinRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentJoin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/join"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.team != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "team", r.team, "", "")
	}
	if r.pairMeAsap != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "pairMeAsap", r.pairMeAsap, "", "")
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

type TournamentsArenaAPIApiTournamentPostRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	clockTime *float32
	clockIncrement *int32
	minutes *int32
	name *string
	waitMinutes *int32
	startDate *int64
	variant *string
	rated *bool
	position *string
	berserkable *bool
	streakable *bool
	hasChat *bool
	description *string
	password *string
	teamBattleByTeam *string
	conditionsTeamMemberTeamId *string
	conditionsMinRatingRating *int32
	conditionsMaxRatingRating *int32
	conditionsNbRatedGameNb *int32
	conditionsAllowList *string
	conditionsBots *bool
	conditionsAccountAge *int32
}

// Clock initial time in minutes
func (r TournamentsArenaAPIApiTournamentPostRequest) ClockTime(clockTime float32) TournamentsArenaAPIApiTournamentPostRequest {
	r.clockTime = &clockTime
	return r
}

// Clock increment in seconds
func (r TournamentsArenaAPIApiTournamentPostRequest) ClockIncrement(clockIncrement int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// How long the tournament lasts, in minutes
func (r TournamentsArenaAPIApiTournamentPostRequest) Minutes(minutes int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.minutes = &minutes
	return r
}

// The tournament name. Leave empty to get a random Grandmaster name
func (r TournamentsArenaAPIApiTournamentPostRequest) Name(name string) TournamentsArenaAPIApiTournamentPostRequest {
	r.name = &name
	return r
}

// How long to wait before starting the tournament, from now, in minutes
func (r TournamentsArenaAPIApiTournamentPostRequest) WaitMinutes(waitMinutes int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.waitMinutes = &waitMinutes
	return r
}

// Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting
func (r TournamentsArenaAPIApiTournamentPostRequest) StartDate(startDate int64) TournamentsArenaAPIApiTournamentPostRequest {
	r.startDate = &startDate
	return r
}

func (r TournamentsArenaAPIApiTournamentPostRequest) Variant(variant string) TournamentsArenaAPIApiTournamentPostRequest {
	r.variant = &variant
	return r
}

// Games are rated and impact players ratings
func (r TournamentsArenaAPIApiTournamentPostRequest) Rated(rated bool) TournamentsArenaAPIApiTournamentPostRequest {
	r.rated = &rated
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r TournamentsArenaAPIApiTournamentPostRequest) Position(position string) TournamentsArenaAPIApiTournamentPostRequest {
	r.position = &position
	return r
}

// Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2
func (r TournamentsArenaAPIApiTournamentPostRequest) Berserkable(berserkable bool) TournamentsArenaAPIApiTournamentPostRequest {
	r.berserkable = &berserkable
	return r
}

// After 2 wins, consecutive wins grant 4 points instead of 2.
func (r TournamentsArenaAPIApiTournamentPostRequest) Streakable(streakable bool) TournamentsArenaAPIApiTournamentPostRequest {
	r.streakable = &streakable
	return r
}

// Whether the players can discuss in a chat
func (r TournamentsArenaAPIApiTournamentPostRequest) HasChat(hasChat bool) TournamentsArenaAPIApiTournamentPostRequest {
	r.hasChat = &hasChat
	return r
}

// Anything you want to tell players about the tournament
func (r TournamentsArenaAPIApiTournamentPostRequest) Description(description string) TournamentsArenaAPIApiTournamentPostRequest {
	r.description = &description
	return r
}

// Make the tournament private, and restrict access with a password. You can also [generate user-specific entry codes](https://github.com/lichess-org/api/tree/master/example/tournament-entry-code) based on this password. 
func (r TournamentsArenaAPIApiTournamentPostRequest) Password(password string) TournamentsArenaAPIApiTournamentPostRequest {
	r.password = &password
	return r
}

// Set the ID of a team you lead to create a team battle. The other teams can be added using the [team battle edit endpoint](#tag/arena-tournaments/POST/api/tournament/team-battle/{id}). 
func (r TournamentsArenaAPIApiTournamentPostRequest) TeamBattleByTeam(teamBattleByTeam string) TournamentsArenaAPIApiTournamentPostRequest {
	r.teamBattleByTeam = &teamBattleByTeam
	return r
}

// Restrict entry to members of a team. The teamId is the last part of a team URL, e.g. &#x60;https://lichess.org/team/coders&#x60; has teamId &#x3D; &#x60;coders&#x60;. Leave empty to let everyone join the tournament. Do not use this to create team battles, use &#x60;teamBattleByTeam&#x60; instead. 
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsTeamMemberTeamId(conditionsTeamMemberTeamId string) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsTeamMemberTeamId = &conditionsTeamMemberTeamId
	return r
}

// Minimum rating to join. Leave empty to let everyone join the tournament.
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsMinRatingRating(conditionsMinRatingRating int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsMinRatingRating = &conditionsMinRatingRating
	return r
}

// Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament.
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsMaxRatingRating(conditionsMaxRatingRating int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsMaxRatingRating = &conditionsMaxRatingRating
	return r
}

// Minimum number of rated games required to join.
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsNbRatedGameNb(conditionsNbRatedGameNb int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsNbRatedGameNb = &conditionsNbRatedGameNb
	return r
}

// Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60; 
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsAllowList(conditionsAllowList string) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsAllowList = &conditionsAllowList
	return r
}

// Whether bots are allowed to join the tournament.
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsBots(conditionsBots bool) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsBots = &conditionsBots
	return r
}

// Minium account age in days required to join.
func (r TournamentsArenaAPIApiTournamentPostRequest) ConditionsAccountAge(conditionsAccountAge int32) TournamentsArenaAPIApiTournamentPostRequest {
	r.conditionsAccountAge = &conditionsAccountAge
	return r
}

func (r TournamentsArenaAPIApiTournamentPostRequest) Execute() (*ApiTournamentPost200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentPostExecute(r)
}

/*
ApiTournamentPost Create a new Arena tournament

Create a public or private Arena tournament.
This endpoint mirrors the form on <https://lichess.org/tournament/new>.
You can create up to 12 public tournaments per day, or 24 private tournaments.
A team battle can be created by specifying the `teamBattleByTeam` argument.
Additional restrictions:
  - clockTime + clockIncrement > 0
  - 15s and 0+1 variant tournaments cannot be rated
  - Clock time in comparison to tournament length must be reasonable: 3 <= (minutes * 60) / (96 * clockTime + 48 * clockIncrement + 15) <= 150


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TournamentsArenaAPIApiTournamentPostRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentPost(ctx context.Context) TournamentsArenaAPIApiTournamentPostRequest {
	return TournamentsArenaAPIApiTournamentPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiTournamentPost200Response
func (a *TournamentsArenaAPIService) ApiTournamentPostExecute(r TournamentsArenaAPIApiTournamentPostRequest) (*ApiTournamentPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiTournamentPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clockTime == nil {
		return localVarReturnValue, nil, reportError("clockTime is required and must be specified")
	}
	if r.clockIncrement == nil {
		return localVarReturnValue, nil, reportError("clockIncrement is required and must be specified")
	}
	if r.minutes == nil {
		return localVarReturnValue, nil, reportError("minutes is required and must be specified")
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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "clockTime", r.clockTime, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "clockIncrement", r.clockIncrement, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "minutes", r.minutes, "", "")
	if r.waitMinutes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "waitMinutes", r.waitMinutes, "", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startDate", r.startDate, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.position != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "position", r.position, "", "")
	}
	if r.berserkable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "berserkable", r.berserkable, "", "")
	}
	if r.streakable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "streakable", r.streakable, "", "")
	}
	if r.hasChat != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hasChat", r.hasChat, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.teamBattleByTeam != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "teamBattleByTeam", r.teamBattleByTeam, "", "")
	}
	if r.conditionsTeamMemberTeamId != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.teamMember.teamId", r.conditionsTeamMemberTeamId, "", "")
	}
	if r.conditionsMinRatingRating != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.minRating.rating", r.conditionsMinRatingRating, "", "")
	}
	if r.conditionsMaxRatingRating != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.maxRating.rating", r.conditionsMaxRatingRating, "", "")
	}
	if r.conditionsNbRatedGameNb != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.nbRatedGame.nb", r.conditionsNbRatedGameNb, "", "")
	}
	if r.conditionsAllowList != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.allowList", r.conditionsAllowList, "", "")
	}
	if r.conditionsBots != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.bots", r.conditionsBots, "", "")
	}
	if r.conditionsAccountAge != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.accountAge", r.conditionsAccountAge, "", "")
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

type TournamentsArenaAPIApiTournamentTeamBattlePostRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	teams *string
	nbLeaders *int32
}

// All team IDs of the team battle, separated by commas. Make sure to always send the full list. Teams that are not in the list will be removed from the team battle. Example: &#x60;coders,zhigalko_sergei-fan-club,hhSwTKZv&#x60; 
func (r TournamentsArenaAPIApiTournamentTeamBattlePostRequest) Teams(teams string) TournamentsArenaAPIApiTournamentTeamBattlePostRequest {
	r.teams = &teams
	return r
}

// Number team leaders per team.
func (r TournamentsArenaAPIApiTournamentTeamBattlePostRequest) NbLeaders(nbLeaders int32) TournamentsArenaAPIApiTournamentTeamBattlePostRequest {
	r.nbLeaders = &nbLeaders
	return r
}

func (r TournamentsArenaAPIApiTournamentTeamBattlePostRequest) Execute() (*Tournament200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentTeamBattlePostExecute(r)
}

/*
ApiTournamentTeamBattlePost Update a team battle

Set the teams and number of leaders of a team battle.
To update the other attributes of a team battle, use the [tournament update endpoint](#tag/arena-tournaments/POST/api/tournament/{id}).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID
 @return TournamentsArenaAPIApiTournamentTeamBattlePostRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentTeamBattlePost(ctx context.Context, id string) TournamentsArenaAPIApiTournamentTeamBattlePostRequest {
	return TournamentsArenaAPIApiTournamentTeamBattlePostRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Tournament200Response
func (a *TournamentsArenaAPIService) ApiTournamentTeamBattlePostExecute(r TournamentsArenaAPIApiTournamentTeamBattlePostRequest) (*Tournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentTeamBattlePost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/team-battle/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if strlen(r.id) < 8 {
		return localVarReturnValue, nil, reportError("id must have at least 8 elements")
	}
	if strlen(r.id) > 8 {
		return localVarReturnValue, nil, reportError("id must have less than 8 elements")
	}
	if r.teams == nil {
		return localVarReturnValue, nil, reportError("teams is required and must be specified")
	}
	if r.nbLeaders == nil {
		return localVarReturnValue, nil, reportError("nbLeaders is required and must be specified")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "teams", r.teams, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "nbLeaders", r.nbLeaders, "", "")
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

type TournamentsArenaAPIApiTournamentTerminateRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
}

func (r TournamentsArenaAPIApiTournamentTerminateRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentTerminateExecute(r)
}

/*
ApiTournamentTerminate Terminate an Arena tournament

Terminate an Arena tournament


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIApiTournamentTerminateRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentTerminate(ctx context.Context, id string) TournamentsArenaAPIApiTournamentTerminateRequest {
	return TournamentsArenaAPIApiTournamentTerminateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *TournamentsArenaAPIService) ApiTournamentTerminateExecute(r TournamentsArenaAPIApiTournamentTerminateRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentTerminate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/terminate"
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

type TournamentsArenaAPIApiTournamentUpdateRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	clockTime *float32
	clockIncrement *int32
	minutes *int32
	name *string
	waitMinutes *int32
	startDate *int64
	variant *string
	rated *bool
	position *string
	berserkable *bool
	streakable *bool
	hasChat *bool
	description *string
	password *string
	conditionsMinRatingRating *int32
	conditionsMaxRatingRating *int32
	conditionsNbRatedGameNb *int32
	conditionsAllowList *string
	conditionsBots *bool
	conditionsAccountAge *int32
}

// Clock initial time in minutes
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ClockTime(clockTime float32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.clockTime = &clockTime
	return r
}

// Clock increment in seconds
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ClockIncrement(clockIncrement int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// How long the tournament lasts, in minutes
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Minutes(minutes int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.minutes = &minutes
	return r
}

// The tournament name. Leave empty to get a random Grandmaster name
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Name(name string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.name = &name
	return r
}

// How long to wait before starting the tournament, from now, in minutes
func (r TournamentsArenaAPIApiTournamentUpdateRequest) WaitMinutes(waitMinutes int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.waitMinutes = &waitMinutes
	return r
}

// Timestamp (in milliseconds) to start the tournament at a given date and time. Overrides the &#x60;waitMinutes&#x60; setting
func (r TournamentsArenaAPIApiTournamentUpdateRequest) StartDate(startDate int64) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.startDate = &startDate
	return r
}

func (r TournamentsArenaAPIApiTournamentUpdateRequest) Variant(variant string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.variant = &variant
	return r
}

// Games are rated and impact players ratings
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Rated(rated bool) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.rated = &rated
	return r
}

// Custom initial position (in X-FEN). Variant must be standard, fromPosition, or chess960 (if a valid 960 starting position), and the game cannot be rated.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Position(position string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.position = &position
	return r
}

// Whether the players can use berserk. Only allowed if clockIncrement &lt;&#x3D; clockTime * 2
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Berserkable(berserkable bool) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.berserkable = &berserkable
	return r
}

// After 2 wins, consecutive wins grant 4 points instead of 2.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Streakable(streakable bool) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.streakable = &streakable
	return r
}

// Whether the players can discuss in a chat
func (r TournamentsArenaAPIApiTournamentUpdateRequest) HasChat(hasChat bool) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.hasChat = &hasChat
	return r
}

// Anything you want to tell players about the tournament
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Description(description string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.description = &description
	return r
}

// Make the tournament private, and restrict access with a password
func (r TournamentsArenaAPIApiTournamentUpdateRequest) Password(password string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.password = &password
	return r
}

// Minimum rating to join. Leave empty to let everyone join the tournament.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsMinRatingRating(conditionsMinRatingRating int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsMinRatingRating = &conditionsMinRatingRating
	return r
}

// Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsMaxRatingRating(conditionsMaxRatingRating int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsMaxRatingRating = &conditionsMaxRatingRating
	return r
}

// Minimum number of rated games required to join.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsNbRatedGameNb(conditionsNbRatedGameNb int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsNbRatedGameNb = &conditionsNbRatedGameNb
	return r
}

// Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60; 
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsAllowList(conditionsAllowList string) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsAllowList = &conditionsAllowList
	return r
}

// Whether bots are allowed to join the tournament.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsBots(conditionsBots bool) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsBots = &conditionsBots
	return r
}

// Minium account age in days required to join.
func (r TournamentsArenaAPIApiTournamentUpdateRequest) ConditionsAccountAge(conditionsAccountAge int32) TournamentsArenaAPIApiTournamentUpdateRequest {
	r.conditionsAccountAge = &conditionsAccountAge
	return r
}

func (r TournamentsArenaAPIApiTournamentUpdateRequest) Execute() (*Tournament200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentUpdateExecute(r)
}

/*
ApiTournamentUpdate Update an Arena tournament

Update an Arena tournament.
Be mindful not to make important changes to ongoing tournaments.
Can be used to update a team battle.
Additional restrictions:
  - clockTime + clockIncrement > 0
  - 15s and 0+1 variant tournaments cannot be rated
  - Clock time in comparison to tournament length must be reasonable: 3 <= (minutes * 60) / (96 * clockTime + 48 * clockIncrement + 15) <= 150


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIApiTournamentUpdateRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentUpdate(ctx context.Context, id string) TournamentsArenaAPIApiTournamentUpdateRequest {
	return TournamentsArenaAPIApiTournamentUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Tournament200Response
func (a *TournamentsArenaAPIService) ApiTournamentUpdateExecute(r TournamentsArenaAPIApiTournamentUpdateRequest) (*Tournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clockTime == nil {
		return localVarReturnValue, nil, reportError("clockTime is required and must be specified")
	}
	if r.clockIncrement == nil {
		return localVarReturnValue, nil, reportError("clockIncrement is required and must be specified")
	}
	if r.minutes == nil {
		return localVarReturnValue, nil, reportError("minutes is required and must be specified")
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
	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "name", r.name, "", "")
	}
	parameterAddToHeaderOrQuery(localVarFormParams, "clockTime", r.clockTime, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "clockIncrement", r.clockIncrement, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "minutes", r.minutes, "", "")
	if r.waitMinutes != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "waitMinutes", r.waitMinutes, "", "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startDate", r.startDate, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.position != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "position", r.position, "", "")
	}
	if r.berserkable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "berserkable", r.berserkable, "", "")
	}
	if r.streakable != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "streakable", r.streakable, "", "")
	}
	if r.hasChat != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "hasChat", r.hasChat, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.conditionsMinRatingRating != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.minRating.rating", r.conditionsMinRatingRating, "", "")
	}
	if r.conditionsMaxRatingRating != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.maxRating.rating", r.conditionsMaxRatingRating, "", "")
	}
	if r.conditionsNbRatedGameNb != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.nbRatedGame.nb", r.conditionsNbRatedGameNb, "", "")
	}
	if r.conditionsAllowList != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.allowList", r.conditionsAllowList, "", "")
	}
	if r.conditionsBots != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.bots", r.conditionsBots, "", "")
	}
	if r.conditionsAccountAge != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.accountAge", r.conditionsAccountAge, "", "")
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

type TournamentsArenaAPIApiTournamentWithdrawRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
}

func (r TournamentsArenaAPIApiTournamentWithdrawRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.ApiTournamentWithdrawExecute(r)
}

/*
ApiTournamentWithdraw Pause or leave an Arena tournament

Leave a future Arena tournament, or take a break on an ongoing Arena tournament.
It's possible to join again later. Points and streaks are preserved.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIApiTournamentWithdrawRequest
*/
func (a *TournamentsArenaAPIService) ApiTournamentWithdraw(ctx context.Context, id string) TournamentsArenaAPIApiTournamentWithdrawRequest {
	return TournamentsArenaAPIApiTournamentWithdrawRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *TournamentsArenaAPIService) ApiTournamentWithdrawExecute(r TournamentsArenaAPIApiTournamentWithdrawRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiTournamentWithdraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/withdraw"
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

type TournamentsArenaAPIApiUserNameTournamentCreatedRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	username string
	nb *int32
	status *int32
}

// Max number of tournaments to fetch
func (r TournamentsArenaAPIApiUserNameTournamentCreatedRequest) Nb(nb int32) TournamentsArenaAPIApiUserNameTournamentCreatedRequest {
	r.nb = &nb
	return r
}

// Include tournaments in the given status: \&quot;Created\&quot; (10), \&quot;Started\&quot; (20), \&quot;Finished\&quot; (30) You can add this parameter more than once to include tournaments in different statuses. Example: &#x60;?status&#x3D;10&amp;status&#x3D;20&#x60; 
func (r TournamentsArenaAPIApiUserNameTournamentCreatedRequest) Status(status int32) TournamentsArenaAPIApiUserNameTournamentCreatedRequest {
	r.status = &status
	return r
}

func (r TournamentsArenaAPIApiUserNameTournamentCreatedRequest) Execute() (*ApiTournament200ResponseCreatedInner, *http.Response, error) {
	return r.ApiService.ApiUserNameTournamentCreatedExecute(r)
}

/*
ApiUserNameTournamentCreated Get tournaments created by a user

Get all tournaments created by a given user.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).
The stream is throttled, depending on who is making the request:
  - Anonymous request: 20 tournaments per second
  - [OAuth2 authenticated](#description/authentication) request: 30 tournaments per second
  - Authenticated, downloading your own tournaments: 50 tournaments per second


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user whose created tournaments to fetch
 @return TournamentsArenaAPIApiUserNameTournamentCreatedRequest
*/
func (a *TournamentsArenaAPIService) ApiUserNameTournamentCreated(ctx context.Context, username string) TournamentsArenaAPIApiUserNameTournamentCreatedRequest {
	return TournamentsArenaAPIApiUserNameTournamentCreatedRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return ApiTournament200ResponseCreatedInner
func (a *TournamentsArenaAPIService) ApiUserNameTournamentCreatedExecute(r TournamentsArenaAPIApiUserNameTournamentCreatedRequest) (*ApiTournament200ResponseCreatedInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiTournament200ResponseCreatedInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiUserNameTournamentCreated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/tournament/created"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	}
	if r.status != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "status", r.status, "form", "")
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

type TournamentsArenaAPIApiUserNameTournamentPlayedRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	username string
	nb *int32
	performance *bool
}

// Max number of tournaments to fetch
func (r TournamentsArenaAPIApiUserNameTournamentPlayedRequest) Nb(nb int32) TournamentsArenaAPIApiUserNameTournamentPlayedRequest {
	r.nb = &nb
	return r
}

// Include the player performance rating in the response, at some cost for the server. 
func (r TournamentsArenaAPIApiUserNameTournamentPlayedRequest) Performance(performance bool) TournamentsArenaAPIApiUserNameTournamentPlayedRequest {
	r.performance = &performance
	return r
}

func (r TournamentsArenaAPIApiUserNameTournamentPlayedRequest) Execute() (*ApiUserNameTournamentPlayed200Response, *http.Response, error) {
	return r.ApiService.ApiUserNameTournamentPlayedExecute(r)
}

/*
ApiUserNameTournamentPlayed Get tournaments played by a user

Get all tournaments played by a given user.
Tournaments are sorted by reverse chronological order of start date (last played first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).
The stream is throttled, depending on who is making the request:
  - Anonymous request: 20 tournaments per second
  - [OAuth2 authenticated](#description/authentication) request: 30 tournaments per second
  - Authenticated, downloading your own tournaments: 50 tournaments per second


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username The user whose played tournaments to fetch
 @return TournamentsArenaAPIApiUserNameTournamentPlayedRequest
*/
func (a *TournamentsArenaAPIService) ApiUserNameTournamentPlayed(ctx context.Context, username string) TournamentsArenaAPIApiUserNameTournamentPlayedRequest {
	return TournamentsArenaAPIApiUserNameTournamentPlayedRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return ApiUserNameTournamentPlayed200Response
func (a *TournamentsArenaAPIService) ApiUserNameTournamentPlayedExecute(r TournamentsArenaAPIApiUserNameTournamentPlayedRequest) (*ApiUserNameTournamentPlayed200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUserNameTournamentPlayed200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ApiUserNameTournamentPlayed")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/tournament/played"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	}
	if r.performance != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "performance", r.performance, "form", "")
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

type TournamentsArenaAPIGamesByTournamentRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	accept *string
	player *string
	moves *bool
	pgnInJson *bool
	tags *bool
	clocks *bool
	evals *bool
	accuracy *bool
	opening *bool
	division *bool
}

// Specify the desired response format. Use &#x60;application/x-chess-pgn&#x60; to get the games in PGN format. Use &#x60;application/x-ndjson&#x60; to get the games in ndjson format. [Read about ndjson here](#description/streaming-with-nd-json) and how you can parse it in Javascript. 
func (r TournamentsArenaAPIGamesByTournamentRequest) Accept(accept string) TournamentsArenaAPIGamesByTournamentRequest {
	r.accept = &accept
	return r
}

// Only games of a particular player. Leave empty to fetch games of all players.
func (r TournamentsArenaAPIGamesByTournamentRequest) Player(player string) TournamentsArenaAPIGamesByTournamentRequest {
	r.player = &player
	return r
}

// Include the PGN moves.
func (r TournamentsArenaAPIGamesByTournamentRequest) Moves(moves bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r TournamentsArenaAPIGamesByTournamentRequest) PgnInJson(pgnInJson bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r TournamentsArenaAPIGamesByTournamentRequest) Tags(tags bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r TournamentsArenaAPIGamesByTournamentRequest) Clocks(clocks bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r TournamentsArenaAPIGamesByTournamentRequest) Evals(evals bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r TournamentsArenaAPIGamesByTournamentRequest) Accuracy(accuracy bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r TournamentsArenaAPIGamesByTournamentRequest) Opening(opening bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r TournamentsArenaAPIGamesByTournamentRequest) Division(division bool) TournamentsArenaAPIGamesByTournamentRequest {
	r.division = &division
	return r
}

func (r TournamentsArenaAPIGamesByTournamentRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GamesByTournamentExecute(r)
}

/*
GamesByTournament Export games of an Arena tournament

Download games of a tournament in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by reverse chronological order (most recent first).
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIGamesByTournamentRequest
*/
func (a *TournamentsArenaAPIService) GamesByTournament(ctx context.Context, id string) TournamentsArenaAPIGamesByTournamentRequest {
	return TournamentsArenaAPIGamesByTournamentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return string
func (a *TournamentsArenaAPIService) GamesByTournamentExecute(r TournamentsArenaAPIGamesByTournamentRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.GamesByTournament")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/games"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.player != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "player", r.player, "form", "")
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

type TournamentsArenaAPIResultsByTournamentRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	nb *int32
	sheet *bool
}

// Max number of players to fetch
func (r TournamentsArenaAPIResultsByTournamentRequest) Nb(nb int32) TournamentsArenaAPIResultsByTournamentRequest {
	r.nb = &nb
	return r
}

// Add a &#x60;sheet&#x60; field to the player document. It&#39;s an expensive server computation that slows down the stream. 
func (r TournamentsArenaAPIResultsByTournamentRequest) Sheet(sheet bool) TournamentsArenaAPIResultsByTournamentRequest {
	r.sheet = &sheet
	return r
}

func (r TournamentsArenaAPIResultsByTournamentRequest) Execute() (*ResultsByTournament200Response, *http.Response, error) {
	return r.ApiService.ResultsByTournamentExecute(r)
}

/*
ResultsByTournament Get results of an Arena tournament

Players of an Arena tournament, with their score and performance, sorted by rank (best first).
**Players are streamed as [ndjson](#description/streaming-with-nd-json)**, i.e. one JSON object per line.
If called on an ongoing tournament, results can be inconsistent
due to ranking changes while the players are being streamed.
Use on finished tournaments for guaranteed consistency.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPIResultsByTournamentRequest
*/
func (a *TournamentsArenaAPIService) ResultsByTournament(ctx context.Context, id string) TournamentsArenaAPIResultsByTournamentRequest {
	return TournamentsArenaAPIResultsByTournamentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ResultsByTournament200Response
func (a *TournamentsArenaAPIService) ResultsByTournamentExecute(r TournamentsArenaAPIResultsByTournamentRequest) (*ResultsByTournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultsByTournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.ResultsByTournament")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	}
	if r.sheet != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sheet", r.sheet, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "sheet", defaultValue, "form", "")
		r.sheet = &defaultValue
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

type TournamentsArenaAPITeamsByTournamentRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
}

func (r TournamentsArenaAPITeamsByTournamentRequest) Execute() (*TeamsByTournament200Response, *http.Response, error) {
	return r.ApiService.TeamsByTournamentExecute(r)
}

/*
TeamsByTournament Get team standing of a team battle

Teams of a team battle tournament, with top players, sorted by rank (best first).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPITeamsByTournamentRequest
*/
func (a *TournamentsArenaAPIService) TeamsByTournament(ctx context.Context, id string) TournamentsArenaAPITeamsByTournamentRequest {
	return TournamentsArenaAPITeamsByTournamentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return TeamsByTournament200Response
func (a *TournamentsArenaAPIService) TeamsByTournamentExecute(r TournamentsArenaAPITeamsByTournamentRequest) (*TeamsByTournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamsByTournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.TeamsByTournament")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}/teams"
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

type TournamentsArenaAPITournamentRequest struct {
	ctx context.Context
	ApiService TournamentsArenaAPI
	id string
	page *int32
}

// Specify which page of player standings to view.
func (r TournamentsArenaAPITournamentRequest) Page(page int32) TournamentsArenaAPITournamentRequest {
	r.page = &page
	return r
}

func (r TournamentsArenaAPITournamentRequest) Execute() (*Tournament200Response, *http.Response, error) {
	return r.ApiService.TournamentExecute(r)
}

/*
Tournament Get info about an Arena tournament

Get detailed info about recently finished, current, or upcoming tournament's duels, player standings, and other info.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsArenaAPITournamentRequest
*/
func (a *TournamentsArenaAPIService) Tournament(ctx context.Context, id string) TournamentsArenaAPITournamentRequest {
	return TournamentsArenaAPITournamentRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Tournament200Response
func (a *TournamentsArenaAPIService) TournamentExecute(r TournamentsArenaAPITournamentRequest) (*Tournament200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Tournament200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsArenaAPIService.Tournament")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/tournament/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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
