/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lichess

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


type UsersAPI interface {

	/*
	ApiCrosstable Get crosstable

	Get total number of games, and current score, of any two users.
If the `matchup` flag is provided, and the users are currently playing, also gets the current match game number and scores.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param user1
	@param user2
	@return UsersAPIApiCrosstableRequest
	*/
	ApiCrosstable(ctx context.Context, user1 string, user2 string) UsersAPIApiCrosstableRequest

	// ApiCrosstableExecute executes the request
	//  @return ApiCrosstable200Response
	ApiCrosstableExecute(r UsersAPIApiCrosstableRequest) (*ApiCrosstable200Response, *http.Response, error)

	/*
	ApiPlayerAutocomplete Autocomplete usernames

	Provides autocompletion options for an incomplete username.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIApiPlayerAutocompleteRequest
	*/
	ApiPlayerAutocomplete(ctx context.Context) UsersAPIApiPlayerAutocompleteRequest

	// ApiPlayerAutocompleteExecute executes the request
	//  @return ApiPlayerAutocomplete200Response
	ApiPlayerAutocompleteExecute(r UsersAPIApiPlayerAutocompleteRequest) (*ApiPlayerAutocomplete200Response, *http.Response, error)

	/*
	ApiUser Get user public data

	Read public data of a user.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return UsersAPIApiUserRequest
	*/
	ApiUser(ctx context.Context, username string) UsersAPIApiUserRequest

	// ApiUserExecute executes the request
	//  @return ApiUser200Response
	ApiUserExecute(r UsersAPIApiUserRequest) (*ApiUser200Response, *http.Response, error)

	/*
	ApiUserActivity Get user activity

	Read data to generate the activity feed of a user.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return UsersAPIApiUserActivityRequest
	*/
	ApiUserActivity(ctx context.Context, username string) UsersAPIApiUserActivityRequest

	// ApiUserActivityExecute executes the request
	//  @return []ApiUserActivity200ResponseInner
	ApiUserActivityExecute(r UsersAPIApiUserActivityRequest) ([]ApiUserActivity200ResponseInner, *http.Response, error)

	/*
	ApiUserPerf Get performance statistics of a user

	Read performance statistics of a user, for a single performance.
Similar to the [performance pages on the website](https://lichess.org/@/thibault/perf/bullet).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@param perf
	@return UsersAPIApiUserPerfRequest
	*/
	ApiUserPerf(ctx context.Context, username string, perf string) UsersAPIApiUserPerfRequest

	// ApiUserPerfExecute executes the request
	//  @return ApiUserPerf200Response
	ApiUserPerfExecute(r UsersAPIApiUserPerfRequest) (*ApiUserPerf200Response, *http.Response, error)

	/*
	ApiUserRatingHistory Get rating history of a user

	Read rating history of a user, for all perf types.
There is at most one entry per day.
Format of an entry is `[year, month, day, rating]`.
`month` starts at zero (January).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return UsersAPIApiUserRatingHistoryRequest
	*/
	ApiUserRatingHistory(ctx context.Context, username string) UsersAPIApiUserRatingHistoryRequest

	// ApiUserRatingHistoryExecute executes the request
	//  @return []ApiUserRatingHistory200ResponseInner
	ApiUserRatingHistoryExecute(r UsersAPIApiUserRatingHistoryRequest) ([]ApiUserRatingHistory200ResponseInner, *http.Response, error)

	/*
	ApiUsers Get users by ID

	Get up to 300 users by their IDs. Users are returned in the same order as the IDs.
The method is `POST` to allow a longer list of IDs to be sent in the request body.
Please do not try to download all the Lichess users with this endpoint, or any other endpoint.
An API is not a way to fully export a website. We do not provide a full download of the Lichess users.
This endpoint is limited to 8,000 users every 10 minutes, and 120,000 every day.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIApiUsersRequest
	*/
	ApiUsers(ctx context.Context) UsersAPIApiUsersRequest

	// ApiUsersExecute executes the request
	//  @return []ApiUsers200ResponseInner
	ApiUsersExecute(r UsersAPIApiUsersRequest) ([]ApiUsers200ResponseInner, *http.Response, error)

	/*
	ApiUsersStatus Get real-time users status

	Read the `online`, `playing` and `streaming` flags of several users.
This API is very fast and cheap on lichess side.
So you can call it quite often (like once every 5 seconds).
Use it to track players and know when they're connected on lichess and playing games.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIApiUsersStatusRequest
	*/
	ApiUsersStatus(ctx context.Context) UsersAPIApiUsersStatusRequest

	// ApiUsersStatusExecute executes the request
	//  @return []ApiUsersStatus200ResponseInner
	ApiUsersStatusExecute(r UsersAPIApiUsersStatusRequest) ([]ApiUsersStatus200ResponseInner, *http.Response, error)

	/*
	Player Get all top 10

	Get the top 10 players for each speed and variant.
See <https://lichess.org/player>.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIPlayerRequest
	*/
	Player(ctx context.Context) UsersAPIPlayerRequest

	// PlayerExecute executes the request
	//  @return Player200Response
	PlayerExecute(r UsersAPIPlayerRequest) (*Player200Response, *http.Response, error)

	/*
	PlayerTopNbPerfType Get one leaderboard

	Get the leaderboard for a single speed or variant (a.k.a. `perfType`).
There is no leaderboard for correspondence or puzzles.
See <https://lichess.org/player/top/100/bullet>.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param nb How many users to fetch
	@param perfType The speed or variant
	@return UsersAPIPlayerTopNbPerfTypeRequest
	*/
	PlayerTopNbPerfType(ctx context.Context, nb int32, perfType string) UsersAPIPlayerTopNbPerfTypeRequest

	// PlayerTopNbPerfTypeExecute executes the request
	//  @return PlayerTopNbPerfType200Response
	PlayerTopNbPerfTypeExecute(r UsersAPIPlayerTopNbPerfTypeRequest) (*PlayerTopNbPerfType200Response, *http.Response, error)

	/*
	ReadNote Get notes for a user

	Get the private notes that you have added for a user.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return UsersAPIReadNoteRequest
	*/
	ReadNote(ctx context.Context, username string) UsersAPIReadNoteRequest

	// ReadNoteExecute executes the request
	//  @return []ReadNote200ResponseInner
	ReadNoteExecute(r UsersAPIReadNoteRequest) ([]ReadNote200ResponseInner, *http.Response, error)

	/*
	StreamerLive Get live streamers

	Get basic info about currently streaming users.
This API is very fast and cheap on lichess side.
So you can call it quite often (like once every 5 seconds).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return UsersAPIStreamerLiveRequest
	*/
	StreamerLive(ctx context.Context) UsersAPIStreamerLiveRequest

	// StreamerLiveExecute executes the request
	//  @return []StreamerLive200ResponseInner
	StreamerLiveExecute(r UsersAPIStreamerLiveRequest) ([]StreamerLive200ResponseInner, *http.Response, error)

	/*
	WriteNote Add a note for a user

	Add a private note available only to you about this account.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username
	@return UsersAPIWriteNoteRequest
	*/
	WriteNote(ctx context.Context, username string) UsersAPIWriteNoteRequest

	// WriteNoteExecute executes the request
	//  @return AccountKidPost200Response
	WriteNoteExecute(r UsersAPIWriteNoteRequest) (*AccountKidPost200Response, *http.Response, error)
}

// UsersAPIService UsersAPI service
type UsersAPIService service

type UsersAPIApiCrosstableRequest struct {
	ctx context.Context
	ApiService UsersAPI
	user1 string
	user2 string
	matchup *bool
}

// Whether to get the current match data, if any
func (r UsersAPIApiCrosstableRequest) Matchup(matchup bool) UsersAPIApiCrosstableRequest {
	r.matchup = &matchup
	return r
}

func (r UsersAPIApiCrosstableRequest) Execute() (*ApiCrosstable200Response, *http.Response, error) {
	return r.ApiService.ApiCrosstableExecute(r)
}

/*
ApiCrosstable Get crosstable

Get total number of games, and current score, of any two users.
If the `matchup` flag is provided, and the users are currently playing, also gets the current match game number and scores.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param user1
 @param user2
 @return UsersAPIApiCrosstableRequest
*/
func (a *UsersAPIService) ApiCrosstable(ctx context.Context, user1 string, user2 string) UsersAPIApiCrosstableRequest {
	return UsersAPIApiCrosstableRequest{
		ApiService: a,
		ctx: ctx,
		user1: user1,
		user2: user2,
	}
}

// Execute executes the request
//  @return ApiCrosstable200Response
func (a *UsersAPIService) ApiCrosstableExecute(r UsersAPIApiCrosstableRequest) (*ApiCrosstable200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiCrosstable200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiCrosstable")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/crosstable/{user1}/{user2}"
	localVarPath = strings.Replace(localVarPath, "{"+"user1"+"}", url.PathEscape(parameterValueToString(r.user1, "user1")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"user2"+"}", url.PathEscape(parameterValueToString(r.user2, "user2")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.matchup != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "matchup", r.matchup, "form", "")
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

type UsersAPIApiPlayerAutocompleteRequest struct {
	ctx context.Context
	ApiService UsersAPI
	term *string
	exists *bool
	object *bool
	names *bool
	friend *bool
	team *string
	tour *string
	swiss *string
	teacher *bool
}

// The beginning of a username
func (r UsersAPIApiPlayerAutocompleteRequest) Term(term string) UsersAPIApiPlayerAutocompleteRequest {
	r.term = &term
	return r
}

// If &#x60;true&#x60;, only checks if the user exists. 
func (r UsersAPIApiPlayerAutocompleteRequest) Exists(exists bool) UsersAPIApiPlayerAutocompleteRequest {
	r.exists = &exists
	return r
}

// - &#x60;false&#x60; returns an array of usernames - &#x60;true&#x60; returns an object with matching users 
func (r UsersAPIApiPlayerAutocompleteRequest) Object(object bool) UsersAPIApiPlayerAutocompleteRequest {
	r.object = &object
	return r
}

// - &#x60;false&#x60; returns an array of usernames - &#x60;true&#x60; returns an array of usernames with preferred casing 
func (r UsersAPIApiPlayerAutocompleteRequest) Names(names bool) UsersAPIApiPlayerAutocompleteRequest {
	r.names = &names
	return r
}

// Returns followed players matching &#x60;term&#x60; if any, else returns other players. Requires [OAuth](#tag/OAuth). 
func (r UsersAPIApiPlayerAutocompleteRequest) Friend(friend bool) UsersAPIApiPlayerAutocompleteRequest {
	r.friend = &friend
	return r
}

// Search within a team. Use team ID/slug. 
func (r UsersAPIApiPlayerAutocompleteRequest) Team(team string) UsersAPIApiPlayerAutocompleteRequest {
	r.team = &team
	return r
}

// Search within a arena tournament. Use tournament ID. 
func (r UsersAPIApiPlayerAutocompleteRequest) Tour(tour string) UsersAPIApiPlayerAutocompleteRequest {
	r.tour = &tour
	return r
}

// Search within a Swiss tournament. 
func (r UsersAPIApiPlayerAutocompleteRequest) Swiss(swiss string) UsersAPIApiPlayerAutocompleteRequest {
	r.swiss = &swiss
	return r
}

// Only search for players who also have a teacher role. 
func (r UsersAPIApiPlayerAutocompleteRequest) Teacher(teacher bool) UsersAPIApiPlayerAutocompleteRequest {
	r.teacher = &teacher
	return r
}

func (r UsersAPIApiPlayerAutocompleteRequest) Execute() (*ApiPlayerAutocomplete200Response, *http.Response, error) {
	return r.ApiService.ApiPlayerAutocompleteExecute(r)
}

/*
ApiPlayerAutocomplete Autocomplete usernames

Provides autocompletion options for an incomplete username.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersAPIApiPlayerAutocompleteRequest
*/
func (a *UsersAPIService) ApiPlayerAutocomplete(ctx context.Context) UsersAPIApiPlayerAutocompleteRequest {
	return UsersAPIApiPlayerAutocompleteRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ApiPlayerAutocomplete200Response
func (a *UsersAPIService) ApiPlayerAutocompleteExecute(r UsersAPIApiPlayerAutocompleteRequest) (*ApiPlayerAutocomplete200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiPlayerAutocomplete200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiPlayerAutocomplete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/player/autocomplete"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.term == nil {
		return localVarReturnValue, nil, reportError("term is required and must be specified")
	}
	if strlen(*r.term) < 3 {
		return localVarReturnValue, nil, reportError("term must have at least 3 elements")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "term", r.term, "form", "")
	if r.exists != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "exists", r.exists, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "exists", defaultValue, "form", "")
		r.exists = &defaultValue
	}
	if r.object != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", r.object, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "object", defaultValue, "form", "")
		r.object = &defaultValue
	}
	if r.names != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "names", r.names, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "names", defaultValue, "form", "")
		r.names = &defaultValue
	}
	if r.friend != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "friend", r.friend, "form", "")
	}
	if r.team != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "team", r.team, "form", "")
	}
	if r.tour != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tour", r.tour, "form", "")
	}
	if r.swiss != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "swiss", r.swiss, "form", "")
	}
	if r.teacher != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "teacher", r.teacher, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "teacher", defaultValue, "form", "")
		r.teacher = &defaultValue
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

type UsersAPIApiUserRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
	trophies *bool
	profile *bool
	rank *bool
	fideId *bool
}

// Include user trophies
func (r UsersAPIApiUserRequest) Trophies(trophies bool) UsersAPIApiUserRequest {
	r.trophies = &trophies
	return r
}

// Include user profile data
func (r UsersAPIApiUserRequest) Profile(profile bool) UsersAPIApiUserRequest {
	r.profile = &profile
	return r
}

// Include global lichess ranking for each perf
func (r UsersAPIApiUserRequest) Rank(rank bool) UsersAPIApiUserRequest {
	r.rank = &rank
	return r
}

// Include public FIDE ID if any
func (r UsersAPIApiUserRequest) FideId(fideId bool) UsersAPIApiUserRequest {
	r.fideId = &fideId
	return r
}

func (r UsersAPIApiUserRequest) Execute() (*ApiUser200Response, *http.Response, error) {
	return r.ApiService.ApiUserExecute(r)
}

/*
ApiUser Get user public data

Read public data of a user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return UsersAPIApiUserRequest
*/
func (a *UsersAPIService) ApiUser(ctx context.Context, username string) UsersAPIApiUserRequest {
	return UsersAPIApiUserRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return ApiUser200Response
func (a *UsersAPIService) ApiUserExecute(r UsersAPIApiUserRequest) (*ApiUser200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUser200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUser")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.trophies != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "trophies", r.trophies, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "trophies", defaultValue, "form", "")
		r.trophies = &defaultValue
	}
	if r.profile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", r.profile, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", defaultValue, "form", "")
		r.profile = &defaultValue
	}
	if r.rank != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rank", r.rank, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "rank", defaultValue, "form", "")
		r.rank = &defaultValue
	}
	if r.fideId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "fideId", r.fideId, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "fideId", defaultValue, "form", "")
		r.fideId = &defaultValue
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

type UsersAPIApiUserActivityRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
}

func (r UsersAPIApiUserActivityRequest) Execute() ([]ApiUserActivity200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiUserActivityExecute(r)
}

/*
ApiUserActivity Get user activity

Read data to generate the activity feed of a user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return UsersAPIApiUserActivityRequest
*/
func (a *UsersAPIService) ApiUserActivity(ctx context.Context, username string) UsersAPIApiUserActivityRequest {
	return UsersAPIApiUserActivityRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return []ApiUserActivity200ResponseInner
func (a *UsersAPIService) ApiUserActivityExecute(r UsersAPIApiUserActivityRequest) ([]ApiUserActivity200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiUserActivity200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUserActivity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/activity"
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

type UsersAPIApiUserPerfRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
	perf string
}

func (r UsersAPIApiUserPerfRequest) Execute() (*ApiUserPerf200Response, *http.Response, error) {
	return r.ApiService.ApiUserPerfExecute(r)
}

/*
ApiUserPerf Get performance statistics of a user

Read performance statistics of a user, for a single performance.
Similar to the [performance pages on the website](https://lichess.org/@/thibault/perf/bullet).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @param perf
 @return UsersAPIApiUserPerfRequest
*/
func (a *UsersAPIService) ApiUserPerf(ctx context.Context, username string, perf string) UsersAPIApiUserPerfRequest {
	return UsersAPIApiUserPerfRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
		perf: perf,
	}
}

// Execute executes the request
//  @return ApiUserPerf200Response
func (a *UsersAPIService) ApiUserPerfExecute(r UsersAPIApiUserPerfRequest) (*ApiUserPerf200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiUserPerf200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUserPerf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/perf/{perf}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perf"+"}", url.PathEscape(parameterValueToString(r.perf, "perf")), -1)

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

type UsersAPIApiUserRatingHistoryRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
}

func (r UsersAPIApiUserRatingHistoryRequest) Execute() ([]ApiUserRatingHistory200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiUserRatingHistoryExecute(r)
}

/*
ApiUserRatingHistory Get rating history of a user

Read rating history of a user, for all perf types.
There is at most one entry per day.
Format of an entry is `[year, month, day, rating]`.
`month` starts at zero (January).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return UsersAPIApiUserRatingHistoryRequest
*/
func (a *UsersAPIService) ApiUserRatingHistory(ctx context.Context, username string) UsersAPIApiUserRatingHistoryRequest {
	return UsersAPIApiUserRatingHistoryRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return []ApiUserRatingHistory200ResponseInner
func (a *UsersAPIService) ApiUserRatingHistoryExecute(r UsersAPIApiUserRatingHistoryRequest) ([]ApiUserRatingHistory200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiUserRatingHistory200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUserRatingHistory")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/rating-history"
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

type UsersAPIApiUsersRequest struct {
	ctx context.Context
	ApiService UsersAPI
	body *string
	profile *bool
	rank *bool
}

// User IDs separated by commas.
func (r UsersAPIApiUsersRequest) Body(body string) UsersAPIApiUsersRequest {
	r.body = &body
	return r
}

// Include user profile data
func (r UsersAPIApiUsersRequest) Profile(profile bool) UsersAPIApiUsersRequest {
	r.profile = &profile
	return r
}

// Include global lichess ranking for each perf
func (r UsersAPIApiUsersRequest) Rank(rank bool) UsersAPIApiUsersRequest {
	r.rank = &rank
	return r
}

func (r UsersAPIApiUsersRequest) Execute() ([]ApiUsers200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiUsersExecute(r)
}

/*
ApiUsers Get users by ID

Get up to 300 users by their IDs. Users are returned in the same order as the IDs.
The method is `POST` to allow a longer list of IDs to be sent in the request body.
Please do not try to download all the Lichess users with this endpoint, or any other endpoint.
An API is not a way to fully export a website. We do not provide a full download of the Lichess users.
This endpoint is limited to 8,000 users every 10 minutes, and 120,000 every day.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersAPIApiUsersRequest
*/
func (a *UsersAPIService) ApiUsers(ctx context.Context) UsersAPIApiUsersRequest {
	return UsersAPIApiUsersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiUsers200ResponseInner
func (a *UsersAPIService) ApiUsersExecute(r UsersAPIApiUsersRequest) ([]ApiUsers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiUsers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUsers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	if r.profile != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", r.profile, "form", "")
	} else {
		var defaultValue bool = true
		parameterAddToHeaderOrQuery(localVarQueryParams, "profile", defaultValue, "form", "")
		r.profile = &defaultValue
	}
	if r.rank != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "rank", r.rank, "form", "")
	} else {
		var defaultValue bool = false
		parameterAddToHeaderOrQuery(localVarQueryParams, "rank", defaultValue, "form", "")
		r.rank = &defaultValue
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

type UsersAPIApiUsersStatusRequest struct {
	ctx context.Context
	ApiService UsersAPI
	ids *string
	withSignal *bool
	withGameIds *bool
	withGameMetas *bool
}

// User IDs separated by commas. Up to 100 IDs.
func (r UsersAPIApiUsersStatusRequest) Ids(ids string) UsersAPIApiUsersStatusRequest {
	r.ids = &ids
	return r
}

// Also return the network signal of the player, when available. It ranges from 1 (poor connection, lag &gt; 500ms) to 4 (great connection, lag &lt; 150ms) Defaults to &#x60;false&#x60; to preserve server resources. 
func (r UsersAPIApiUsersStatusRequest) WithSignal(withSignal bool) UsersAPIApiUsersStatusRequest {
	r.withSignal = &withSignal
	return r
}

// Also return the ID of the game being played, if any, for each player, in a &#x60;playingId&#x60; field. Defaults to &#x60;false&#x60; to preserve server resources. 
func (r UsersAPIApiUsersStatusRequest) WithGameIds(withGameIds bool) UsersAPIApiUsersStatusRequest {
	r.withGameIds = &withGameIds
	return r
}

// Also return the id, time control and variant of the game being played, if any, for each player, in a &#x60;playing&#x60; field. Defaults to &#x60;false&#x60; to preserve server resources. Disables &#x60;withGameIds&#x60;. 
func (r UsersAPIApiUsersStatusRequest) WithGameMetas(withGameMetas bool) UsersAPIApiUsersStatusRequest {
	r.withGameMetas = &withGameMetas
	return r
}

func (r UsersAPIApiUsersStatusRequest) Execute() ([]ApiUsersStatus200ResponseInner, *http.Response, error) {
	return r.ApiService.ApiUsersStatusExecute(r)
}

/*
ApiUsersStatus Get real-time users status

Read the `online`, `playing` and `streaming` flags of several users.
This API is very fast and cheap on lichess side.
So you can call it quite often (like once every 5 seconds).
Use it to track players and know when they're connected on lichess and playing games.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersAPIApiUsersStatusRequest
*/
func (a *UsersAPIService) ApiUsersStatus(ctx context.Context) UsersAPIApiUsersStatusRequest {
	return UsersAPIApiUsersStatusRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []ApiUsersStatus200ResponseInner
func (a *UsersAPIService) ApiUsersStatusExecute(r UsersAPIApiUsersStatusRequest) ([]ApiUsersStatus200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ApiUsersStatus200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ApiUsersStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/users/status"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.ids == nil {
		return localVarReturnValue, nil, reportError("ids is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "ids", r.ids, "form", "")
	if r.withSignal != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withSignal", r.withSignal, "form", "")
	}
	if r.withGameIds != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withGameIds", r.withGameIds, "form", "")
	}
	if r.withGameMetas != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "withGameMetas", r.withGameMetas, "form", "")
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

type UsersAPIPlayerRequest struct {
	ctx context.Context
	ApiService UsersAPI
}

func (r UsersAPIPlayerRequest) Execute() (*Player200Response, *http.Response, error) {
	return r.ApiService.PlayerExecute(r)
}

/*
Player Get all top 10

Get the top 10 players for each speed and variant.
See <https://lichess.org/player>.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersAPIPlayerRequest
*/
func (a *UsersAPIService) Player(ctx context.Context) UsersAPIPlayerRequest {
	return UsersAPIPlayerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return Player200Response
func (a *UsersAPIService) PlayerExecute(r UsersAPIPlayerRequest) (*Player200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Player200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.Player")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/player"

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

type UsersAPIPlayerTopNbPerfTypeRequest struct {
	ctx context.Context
	ApiService UsersAPI
	nb int32
	perfType string
}

func (r UsersAPIPlayerTopNbPerfTypeRequest) Execute() (*PlayerTopNbPerfType200Response, *http.Response, error) {
	return r.ApiService.PlayerTopNbPerfTypeExecute(r)
}

/*
PlayerTopNbPerfType Get one leaderboard

Get the leaderboard for a single speed or variant (a.k.a. `perfType`).
There is no leaderboard for correspondence or puzzles.
See <https://lichess.org/player/top/100/bullet>.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param nb How many users to fetch
 @param perfType The speed or variant
 @return UsersAPIPlayerTopNbPerfTypeRequest
*/
func (a *UsersAPIService) PlayerTopNbPerfType(ctx context.Context, nb int32, perfType string) UsersAPIPlayerTopNbPerfTypeRequest {
	return UsersAPIPlayerTopNbPerfTypeRequest{
		ApiService: a,
		ctx: ctx,
		nb: nb,
		perfType: perfType,
	}
}

// Execute executes the request
//  @return PlayerTopNbPerfType200Response
func (a *UsersAPIService) PlayerTopNbPerfTypeExecute(r UsersAPIPlayerTopNbPerfTypeRequest) (*PlayerTopNbPerfType200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PlayerTopNbPerfType200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.PlayerTopNbPerfType")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/player/top/{nb}/{perfType}"
	localVarPath = strings.Replace(localVarPath, "{"+"nb"+"}", url.PathEscape(parameterValueToString(r.nb, "nb")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"perfType"+"}", url.PathEscape(parameterValueToString(r.perfType, "perfType")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.nb < 1 {
		return localVarReturnValue, nil, reportError("nb must be greater than 1")
	}
	if r.nb > 100 {
		return localVarReturnValue, nil, reportError("nb must be less than 100")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/vnd.lichess.v3+json"}

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

type UsersAPIReadNoteRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
}

func (r UsersAPIReadNoteRequest) Execute() ([]ReadNote200ResponseInner, *http.Response, error) {
	return r.ApiService.ReadNoteExecute(r)
}

/*
ReadNote Get notes for a user

Get the private notes that you have added for a user.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return UsersAPIReadNoteRequest
*/
func (a *UsersAPIService) ReadNote(ctx context.Context, username string) UsersAPIReadNoteRequest {
	return UsersAPIReadNoteRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return []ReadNote200ResponseInner
func (a *UsersAPIService) ReadNoteExecute(r UsersAPIReadNoteRequest) ([]ReadNote200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ReadNote200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.ReadNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/note"
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

type UsersAPIStreamerLiveRequest struct {
	ctx context.Context
	ApiService UsersAPI
}

func (r UsersAPIStreamerLiveRequest) Execute() ([]StreamerLive200ResponseInner, *http.Response, error) {
	return r.ApiService.StreamerLiveExecute(r)
}

/*
StreamerLive Get live streamers

Get basic info about currently streaming users.
This API is very fast and cheap on lichess side.
So you can call it quite often (like once every 5 seconds).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return UsersAPIStreamerLiveRequest
*/
func (a *UsersAPIService) StreamerLive(ctx context.Context) UsersAPIStreamerLiveRequest {
	return UsersAPIStreamerLiveRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []StreamerLive200ResponseInner
func (a *UsersAPIService) StreamerLiveExecute(r UsersAPIStreamerLiveRequest) ([]StreamerLive200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []StreamerLive200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.StreamerLive")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/streamer/live"

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

type UsersAPIWriteNoteRequest struct {
	ctx context.Context
	ApiService UsersAPI
	username string
	text *string
}

// The contents of the note
func (r UsersAPIWriteNoteRequest) Text(text string) UsersAPIWriteNoteRequest {
	r.text = &text
	return r
}

func (r UsersAPIWriteNoteRequest) Execute() (*AccountKidPost200Response, *http.Response, error) {
	return r.ApiService.WriteNoteExecute(r)
}

/*
WriteNote Add a note for a user

Add a private note available only to you about this account.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username
 @return UsersAPIWriteNoteRequest
*/
func (a *UsersAPIService) WriteNote(ctx context.Context, username string) UsersAPIWriteNoteRequest {
	return UsersAPIWriteNoteRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return AccountKidPost200Response
func (a *UsersAPIService) WriteNoteExecute(r UsersAPIWriteNoteRequest) (*AccountKidPost200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AccountKidPost200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UsersAPIService.WriteNote")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/user/{username}/note"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
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
