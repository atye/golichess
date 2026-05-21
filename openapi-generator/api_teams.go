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


type TeamsAPI interface {

	/*
	ApiTeamArena Get team Arena tournaments

	Get all Arena tournaments relevant to a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId ID of the team
	@return TeamsAPIApiTeamArenaRequest
	*/
	ApiTeamArena(ctx context.Context, teamId string) TeamsAPIApiTeamArenaRequest

	// ApiTeamArenaExecute executes the request
	//  @return ArenaTournament
	ApiTeamArenaExecute(r TeamsAPIApiTeamArenaRequest) (*ArenaTournament, *http.Response, error)

	/*
	ApiTeamSwiss Get team swiss tournaments

	Get all swiss tournaments of a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPIApiTeamSwissRequest
	*/
	ApiTeamSwiss(ctx context.Context, teamId string) TeamsAPIApiTeamSwissRequest

	// ApiTeamSwissExecute executes the request
	//  @return SwissTournament
	ApiTeamSwissExecute(r TeamsAPIApiTeamSwissRequest) (*SwissTournament, *http.Response, error)

	/*
	TeamAll Get popular teams

	Paginator of the most popular teams.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeamsAPITeamAllRequest
	*/
	TeamAll(ctx context.Context) TeamsAPITeamAllRequest

	// TeamAllExecute executes the request
	//  @return TeamPaginatorJson
	TeamAllExecute(r TeamsAPITeamAllRequest) (*TeamPaginatorJson, *http.Response, error)

	/*
	TeamIdJoin Join a team

	Join a team.
If the team requires a password but the `password` field is incorrect,
then the call fails with `403 Forbidden`.
Similarly, if the team join policy requires a confirmation but the
`message` parameter is not given, then the call fails with
`403 Forbidden`.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamIdJoinRequest
	*/
	TeamIdJoin(ctx context.Context, teamId string) TeamsAPITeamIdJoinRequest

	// TeamIdJoinExecute executes the request
	//  @return Ok
	TeamIdJoinExecute(r TeamsAPITeamIdJoinRequest) (*Ok, *http.Response, error)

	/*
	TeamIdKickUserId Kick a user from your team

	Kick a member out of one of your teams.
- <https://lichess.org/team>


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@param userId
	@return TeamsAPITeamIdKickUserIdRequest
	*/
	TeamIdKickUserId(ctx context.Context, teamId string, userId string) TeamsAPITeamIdKickUserIdRequest

	// TeamIdKickUserIdExecute executes the request
	//  @return Ok
	TeamIdKickUserIdExecute(r TeamsAPITeamIdKickUserIdRequest) (*Ok, *http.Response, error)

	/*
	TeamIdPmAll Message all members

	Send a private message to all members of a team.
You must be a team leader with the "Messages" permission.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamIdPmAllRequest
	*/
	TeamIdPmAll(ctx context.Context, teamId string) TeamsAPITeamIdPmAllRequest

	// TeamIdPmAllExecute executes the request
	//  @return Ok
	TeamIdPmAllExecute(r TeamsAPITeamIdPmAllRequest) (*Ok, *http.Response, error)

	/*
	TeamIdQuit Leave a team

	Leave a team.
- <https://lichess.org/team>


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamIdQuitRequest
	*/
	TeamIdQuit(ctx context.Context, teamId string) TeamsAPITeamIdQuitRequest

	// TeamIdQuitExecute executes the request
	//  @return Ok
	TeamIdQuitExecute(r TeamsAPITeamIdQuitRequest) (*Ok, *http.Response, error)

	/*
	TeamIdUsers Get members of a team

	Members are sorted by reverse chronological order of joining the team (most recent first).
OAuth is only required if the list of members is private.
Up to 5,000 users are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamIdUsersRequest
	*/
	TeamIdUsers(ctx context.Context, teamId string) TeamsAPITeamIdUsersRequest

	// TeamIdUsersExecute executes the request
	//  @return TeamIdUsers200Response
	TeamIdUsersExecute(r TeamsAPITeamIdUsersRequest) (*TeamIdUsers200Response, *http.Response, error)

	/*
	TeamOfUsername Teams of a player

	All the teams a player is a member of.
Teams that hide their player list are only included if you also belong to the team.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return TeamsAPITeamOfUsernameRequest
	*/
	TeamOfUsername(ctx context.Context, username string) TeamsAPITeamOfUsernameRequest

	// TeamOfUsernameExecute executes the request
	//  @return []Team
	TeamOfUsernameExecute(r TeamsAPITeamOfUsernameRequest) ([]Team, *http.Response, error)

	/*
	TeamRequestAccept Accept join request

	Accept someone's request to join your team

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@param userId
	@return TeamsAPITeamRequestAcceptRequest
	*/
	TeamRequestAccept(ctx context.Context, teamId string, userId string) TeamsAPITeamRequestAcceptRequest

	// TeamRequestAcceptExecute executes the request
	//  @return Ok
	TeamRequestAcceptExecute(r TeamsAPITeamRequestAcceptRequest) (*Ok, *http.Response, error)

	/*
	TeamRequestDecline Decline join request

	Decline someone's request to join your team

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@param userId
	@return TeamsAPITeamRequestDeclineRequest
	*/
	TeamRequestDecline(ctx context.Context, teamId string, userId string) TeamsAPITeamRequestDeclineRequest

	// TeamRequestDeclineExecute executes the request
	//  @return Ok
	TeamRequestDeclineExecute(r TeamsAPITeamRequestDeclineRequest) (*Ok, *http.Response, error)

	/*
	TeamRequests Get join requests

	Get pending join requests of your team

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamRequestsRequest
	*/
	TeamRequests(ctx context.Context, teamId string) TeamsAPITeamRequestsRequest

	// TeamRequestsExecute executes the request
	//  @return []TeamRequestWithUser
	TeamRequestsExecute(r TeamsAPITeamRequestsRequest) ([]TeamRequestWithUser, *http.Response, error)

	/*
	TeamSearch Search teams

	Paginator of team search results for a keyword.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return TeamsAPITeamSearchRequest
	*/
	TeamSearch(ctx context.Context) TeamsAPITeamSearchRequest

	// TeamSearchExecute executes the request
	//  @return TeamPaginatorJson
	TeamSearchExecute(r TeamsAPITeamSearchRequest) (*TeamPaginatorJson, *http.Response, error)

	/*
	TeamShow Get a single team

	Public info about a team. Includes the list of publicly visible leaders.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TeamsAPITeamShowRequest
	*/
	TeamShow(ctx context.Context, teamId string) TeamsAPITeamShowRequest

	// TeamShowExecute executes the request
	//  @return Team
	TeamShowExecute(r TeamsAPITeamShowRequest) (*Team, *http.Response, error)
}

// TeamsAPIService TeamsAPI service
type TeamsAPIService service

type TeamsAPIApiTeamArenaRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	max *int32
	status *ArenaStatusName
	createdBy *string
	name *string
}

// How many tournaments to download.
func (r TeamsAPIApiTeamArenaRequest) Max(max int32) TeamsAPIApiTeamArenaRequest {
	r.max = &max
	return r
}

// [Filter] Only arena tournaments in this current state. 
func (r TeamsAPIApiTeamArenaRequest) Status(status ArenaStatusName) TeamsAPIApiTeamArenaRequest {
	r.status = &status
	return r
}

// [Filter] Only arena tournaments created by a given user. 
func (r TeamsAPIApiTeamArenaRequest) CreatedBy(createdBy string) TeamsAPIApiTeamArenaRequest {
	r.createdBy = &createdBy
	return r
}

// [Filter] Only arena tournaments with a given name. 
func (r TeamsAPIApiTeamArenaRequest) Name(name string) TeamsAPIApiTeamArenaRequest {
	r.name = &name
	return r
}

func (r TeamsAPIApiTeamArenaRequest) Execute() (*ArenaTournament, *http.Response, error) {
	return r.ApiService.ApiTeamArenaExecute(r)
}

/*
ApiTeamArena Get team Arena tournaments

Get all Arena tournaments relevant to a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId ID of the team
 @return TeamsAPIApiTeamArenaRequest
*/
func (a *TeamsAPIService) ApiTeamArena(ctx context.Context, teamId string) TeamsAPIApiTeamArenaRequest {
	return TeamsAPIApiTeamArenaRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return ArenaTournament
func (a *TeamsAPIService) ApiTeamArenaExecute(r TeamsAPIApiTeamArenaRequest) (*ArenaTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ArenaTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.ApiTeamArena")
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

type TeamsAPIApiTeamSwissRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	max *int32
	status *SwissStatus
	createdBy *string
	name *string
}

// How many tournaments to download.
func (r TeamsAPIApiTeamSwissRequest) Max(max int32) TeamsAPIApiTeamSwissRequest {
	r.max = &max
	return r
}

// [Filter] Only swiss tournaments in this current state. 
func (r TeamsAPIApiTeamSwissRequest) Status(status SwissStatus) TeamsAPIApiTeamSwissRequest {
	r.status = &status
	return r
}

// [Filter] Only swiss tournaments created by a given user. 
func (r TeamsAPIApiTeamSwissRequest) CreatedBy(createdBy string) TeamsAPIApiTeamSwissRequest {
	r.createdBy = &createdBy
	return r
}

// [Filter] Only swiss tournaments with a given name. 
func (r TeamsAPIApiTeamSwissRequest) Name(name string) TeamsAPIApiTeamSwissRequest {
	r.name = &name
	return r
}

func (r TeamsAPIApiTeamSwissRequest) Execute() (*SwissTournament, *http.Response, error) {
	return r.ApiService.ApiTeamSwissExecute(r)
}

/*
ApiTeamSwiss Get team swiss tournaments

Get all swiss tournaments of a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPIApiTeamSwissRequest
*/
func (a *TeamsAPIService) ApiTeamSwiss(ctx context.Context, teamId string) TeamsAPIApiTeamSwissRequest {
	return TeamsAPIApiTeamSwissRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return SwissTournament
func (a *TeamsAPIService) ApiTeamSwissExecute(r TeamsAPIApiTeamSwissRequest) (*SwissTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwissTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.ApiTeamSwiss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/swiss"
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

type TeamsAPITeamAllRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	page *int32
}

func (r TeamsAPITeamAllRequest) Page(page int32) TeamsAPITeamAllRequest {
	r.page = &page
	return r
}

func (r TeamsAPITeamAllRequest) Execute() (*TeamPaginatorJson, *http.Response, error) {
	return r.ApiService.TeamAllExecute(r)
}

/*
TeamAll Get popular teams

Paginator of the most popular teams.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeamsAPITeamAllRequest
*/
func (a *TeamsAPIService) TeamAll(ctx context.Context) TeamsAPITeamAllRequest {
	return TeamsAPITeamAllRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamPaginatorJson
func (a *TeamsAPIService) TeamAllExecute(r TeamsAPITeamAllRequest) (*TeamPaginatorJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPaginatorJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/all"

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

type TeamsAPITeamIdJoinRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	message *string
	password *string
}

// Required if team manually reviews admission requests.
func (r TeamsAPITeamIdJoinRequest) Message(message string) TeamsAPITeamIdJoinRequest {
	r.message = &message
	return r
}

// Optional password, if the team requires one.
func (r TeamsAPITeamIdJoinRequest) Password(password string) TeamsAPITeamIdJoinRequest {
	r.password = &password
	return r
}

func (r TeamsAPITeamIdJoinRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamIdJoinExecute(r)
}

/*
TeamIdJoin Join a team

Join a team.
If the team requires a password but the `password` field is incorrect,
then the call fails with `403 Forbidden`.
Similarly, if the team join policy requires a confirmation but the
`message` parameter is not given, then the call fails with
`403 Forbidden`.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamIdJoinRequest
*/
func (a *TeamsAPIService) TeamIdJoin(ctx context.Context, teamId string) TeamsAPITeamIdJoinRequest {
	return TeamsAPITeamIdJoinRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamIdJoinExecute(r TeamsAPITeamIdJoinRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamIdJoin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team/{teamId}/join"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

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
	if r.message != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message", r.message, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
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

type TeamsAPITeamIdKickUserIdRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	userId string
}

func (r TeamsAPITeamIdKickUserIdRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamIdKickUserIdExecute(r)
}

/*
TeamIdKickUserId Kick a user from your team

Kick a member out of one of your teams.
- <https://lichess.org/team>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @param userId
 @return TeamsAPITeamIdKickUserIdRequest
*/
func (a *TeamsAPIService) TeamIdKickUserId(ctx context.Context, teamId string, userId string) TeamsAPITeamIdKickUserIdRequest {
	return TeamsAPITeamIdKickUserIdRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamIdKickUserIdExecute(r TeamsAPITeamIdKickUserIdRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamIdKickUserId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/kick/{userId}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type TeamsAPITeamIdPmAllRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	message *string
}

// The message to send to all your team members.
func (r TeamsAPITeamIdPmAllRequest) Message(message string) TeamsAPITeamIdPmAllRequest {
	r.message = &message
	return r
}

func (r TeamsAPITeamIdPmAllRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamIdPmAllExecute(r)
}

/*
TeamIdPmAll Message all members

Send a private message to all members of a team.
You must be a team leader with the "Messages" permission.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamIdPmAllRequest
*/
func (a *TeamsAPIService) TeamIdPmAll(ctx context.Context, teamId string) TeamsAPITeamIdPmAllRequest {
	return TeamsAPITeamIdPmAllRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamIdPmAllExecute(r TeamsAPITeamIdPmAllRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamIdPmAll")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team/{teamId}/pm-all"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

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
	if r.message != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "message", r.message, "", "")
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

type TeamsAPITeamIdQuitRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
}

func (r TeamsAPITeamIdQuitRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamIdQuitExecute(r)
}

/*
TeamIdQuit Leave a team

Leave a team.
- <https://lichess.org/team>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamIdQuitRequest
*/
func (a *TeamsAPIService) TeamIdQuit(ctx context.Context, teamId string) TeamsAPITeamIdQuitRequest {
	return TeamsAPITeamIdQuitRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamIdQuitExecute(r TeamsAPITeamIdQuitRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamIdQuit")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/team/{teamId}/quit"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

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

type TeamsAPITeamIdUsersRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	full *bool
}

// Full user documents with performance ratings. This limits the response to 1,000 users. 
func (r TeamsAPITeamIdUsersRequest) Full(full bool) TeamsAPITeamIdUsersRequest {
	r.full = &full
	return r
}

func (r TeamsAPITeamIdUsersRequest) Execute() (*TeamIdUsers200Response, *http.Response, error) {
	return r.ApiService.TeamIdUsersExecute(r)
}

/*
TeamIdUsers Get members of a team

Members are sorted by reverse chronological order of joining the team (most recent first).
OAuth is only required if the list of members is private.
Up to 5,000 users are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamIdUsersRequest
*/
func (a *TeamsAPIService) TeamIdUsers(ctx context.Context, teamId string) TeamsAPITeamIdUsersRequest {
	return TeamsAPITeamIdUsersRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return TeamIdUsers200Response
func (a *TeamsAPIService) TeamIdUsersExecute(r TeamsAPITeamIdUsersRequest) (*TeamIdUsers200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamIdUsers200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamIdUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/users"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.full != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "full", r.full, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "full", defaultValue, "form", "")
		r.full = &defaultValue
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

type TeamsAPITeamOfUsernameRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	username string
}

func (r TeamsAPITeamOfUsernameRequest) Execute() ([]Team, *http.Response, error) {
	return r.ApiService.TeamOfUsernameExecute(r)
}

/*
TeamOfUsername Teams of a player

All the teams a player is a member of.
Teams that hide their player list are only included if you also belong to the team.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return TeamsAPITeamOfUsernameRequest
*/
func (a *TeamsAPIService) TeamOfUsername(ctx context.Context, username string) TeamsAPITeamOfUsernameRequest {
	return TeamsAPITeamOfUsernameRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return []Team
func (a *TeamsAPIService) TeamOfUsernameExecute(r TeamsAPITeamOfUsernameRequest) ([]Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamOfUsername")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/of/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

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

type TeamsAPITeamRequestAcceptRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	userId string
}

func (r TeamsAPITeamRequestAcceptRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamRequestAcceptExecute(r)
}

/*
TeamRequestAccept Accept join request

Accept someone's request to join your team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @param userId
 @return TeamsAPITeamRequestAcceptRequest
*/
func (a *TeamsAPIService) TeamRequestAccept(ctx context.Context, teamId string, userId string) TeamsAPITeamRequestAcceptRequest {
	return TeamsAPITeamRequestAcceptRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamRequestAcceptExecute(r TeamsAPITeamRequestAcceptRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamRequestAccept")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/request/{userId}/accept"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type TeamsAPITeamRequestDeclineRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	userId string
}

func (r TeamsAPITeamRequestDeclineRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.TeamRequestDeclineExecute(r)
}

/*
TeamRequestDecline Decline join request

Decline someone's request to join your team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @param userId
 @return TeamsAPITeamRequestDeclineRequest
*/
func (a *TeamsAPIService) TeamRequestDecline(ctx context.Context, teamId string, userId string) TeamsAPITeamRequestDeclineRequest {
	return TeamsAPITeamRequestDeclineRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
		userId: userId,
	}
}

// Execute executes the request
//  @return Ok
func (a *TeamsAPIService) TeamRequestDeclineExecute(r TeamsAPITeamRequestDeclineRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamRequestDecline")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/request/{userId}/decline"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"userId"+"}", url.PathEscape(parameterValueToString(r.userId, "userId")), -1)

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

type TeamsAPITeamRequestsRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
	declined *bool
}

// Get the declined join requests
func (r TeamsAPITeamRequestsRequest) Declined(declined bool) TeamsAPITeamRequestsRequest {
	r.declined = &declined
	return r
}

func (r TeamsAPITeamRequestsRequest) Execute() ([]TeamRequestWithUser, *http.Response, error) {
	return r.ApiService.TeamRequestsExecute(r)
}

/*
TeamRequests Get join requests

Get pending join requests of your team

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamRequestsRequest
*/
func (a *TeamsAPIService) TeamRequests(ctx context.Context, teamId string) TeamsAPITeamRequestsRequest {
	return TeamsAPITeamRequestsRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return []TeamRequestWithUser
func (a *TeamsAPIService) TeamRequestsExecute(r TeamsAPITeamRequestsRequest) ([]TeamRequestWithUser, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TeamRequestWithUser
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamRequests")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}/requests"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.declined != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "declined", r.declined, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "declined", defaultValue, "form", "")
		r.declined = &defaultValue
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

type TeamsAPITeamSearchRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	text *string
	page *int32
}

func (r TeamsAPITeamSearchRequest) Text(text string) TeamsAPITeamSearchRequest {
	r.text = &text
	return r
}

func (r TeamsAPITeamSearchRequest) Page(page int32) TeamsAPITeamSearchRequest {
	r.page = &page
	return r
}

func (r TeamsAPITeamSearchRequest) Execute() (*TeamPaginatorJson, *http.Response, error) {
	return r.ApiService.TeamSearchExecute(r)
}

/*
TeamSearch Search teams

Paginator of team search results for a keyword.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return TeamsAPITeamSearchRequest
*/
func (a *TeamsAPIService) TeamSearch(ctx context.Context) TeamsAPITeamSearchRequest {
	return TeamsAPITeamSearchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return TeamPaginatorJson
func (a *TeamsAPIService) TeamSearchExecute(r TeamsAPITeamSearchRequest) (*TeamPaginatorJson, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TeamPaginatorJson
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamSearch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/search"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.text != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "text", r.text, "form", "")
	}
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

type TeamsAPITeamShowRequest struct {
	ctx context.Context
	ApiService TeamsAPI
	teamId string
}

func (r TeamsAPITeamShowRequest) Execute() (*Team, *http.Response, error) {
	return r.ApiService.TeamShowExecute(r)
}

/*
TeamShow Get a single team

Public info about a team. Includes the list of publicly visible leaders.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TeamsAPITeamShowRequest
*/
func (a *TeamsAPIService) TeamShow(ctx context.Context, teamId string) TeamsAPITeamShowRequest {
	return TeamsAPITeamShowRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return Team
func (a *TeamsAPIService) TeamShowExecute(r TeamsAPITeamShowRequest) (*Team, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Team
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TeamsAPIService.TeamShow")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/team/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

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
