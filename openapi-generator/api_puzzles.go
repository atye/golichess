/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.169
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


type PuzzlesAPI interface {

	/*
	ApiPuzzleActivity Get your puzzle activity

	Download your puzzle activity in [ndjson](#description/streaming-with-nd-json) format.
Puzzle activity is sorted by reverse chronological order (most recent first)
We recommend streaming the response, for it can be very long.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PuzzlesAPIApiPuzzleActivityRequest
	*/
	ApiPuzzleActivity(ctx context.Context) PuzzlesAPIApiPuzzleActivityRequest

	// ApiPuzzleActivityExecute executes the request
	//  @return PuzzleActivity
	ApiPuzzleActivityExecute(r PuzzlesAPIApiPuzzleActivityRequest) (*PuzzleActivity, *http.Response, error)

	/*
	ApiPuzzleBatchSelect Get multiple puzzles at once

	Get a batch of random Lichess puzzles in JSON format.

If authenticated, only returns puzzles that the user has never seen before.

**DO NOT** use this endpoint to enumerate puzzles for mass download. Instead, download the [full public puzzle database](https://database.lichess.org/#puzzles).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param angle The theme or opening to filter puzzles with. Recommended: `mix`.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
	@return PuzzlesAPIApiPuzzleBatchSelectRequest
	*/
	ApiPuzzleBatchSelect(ctx context.Context, angle string) PuzzlesAPIApiPuzzleBatchSelectRequest

	// ApiPuzzleBatchSelectExecute executes the request
	//  @return PuzzleBatchSelect
	ApiPuzzleBatchSelectExecute(r PuzzlesAPIApiPuzzleBatchSelectRequest) (*PuzzleBatchSelect, *http.Response, error)

	/*
	ApiPuzzleBatchSolve Solve multiple puzzles at once

	Set puzzles as solved and update ratings.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param angle The theme or opening of the solved puzzles.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
	@return PuzzlesAPIApiPuzzleBatchSolveRequest
	*/
	ApiPuzzleBatchSolve(ctx context.Context, angle string) PuzzlesAPIApiPuzzleBatchSolveRequest

	// ApiPuzzleBatchSolveExecute executes the request
	//  @return PuzzleBatchSolveResponse
	ApiPuzzleBatchSolveExecute(r PuzzlesAPIApiPuzzleBatchSolveRequest) (*PuzzleBatchSolveResponse, *http.Response, error)

	/*
	ApiPuzzleDaily Get the daily puzzle

	Get the daily Lichess puzzle in JSON format.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PuzzlesAPIApiPuzzleDailyRequest
	*/
	ApiPuzzleDaily(ctx context.Context) PuzzlesAPIApiPuzzleDailyRequest

	// ApiPuzzleDailyExecute executes the request
	//  @return PuzzleAndGame
	ApiPuzzleDailyExecute(r PuzzlesAPIApiPuzzleDailyRequest) (*PuzzleAndGame, *http.Response, error)

	/*
	ApiPuzzleDashboard Get your puzzle dashboard

	Download your [puzzle dashboard](https://lichess.org/training/dashboard/30/dashboard) as JSON.
Also includes all puzzle themes played, with aggregated results.
Allows re-creating the [improvement/strengths](https://lichess.org/training/dashboard/30/improvementAreas) interfaces.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param days How many days to look back when aggregating puzzle results. 30 is sensible.
	@return PuzzlesAPIApiPuzzleDashboardRequest
	*/
	ApiPuzzleDashboard(ctx context.Context, days int32) PuzzlesAPIApiPuzzleDashboardRequest

	// ApiPuzzleDashboardExecute executes the request
	//  @return PuzzleDashboard
	ApiPuzzleDashboardExecute(r PuzzlesAPIApiPuzzleDashboardRequest) (*PuzzleDashboard, *http.Response, error)

	/*
	ApiPuzzleId Get a puzzle by its ID

	Get a single Lichess puzzle in JSON format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The puzzle ID
	@return PuzzlesAPIApiPuzzleIdRequest
	*/
	ApiPuzzleId(ctx context.Context, id string) PuzzlesAPIApiPuzzleIdRequest

	// ApiPuzzleIdExecute executes the request
	//  @return PuzzleAndGame
	ApiPuzzleIdExecute(r PuzzlesAPIApiPuzzleIdRequest) (*PuzzleAndGame, *http.Response, error)

	/*
	ApiPuzzleNext Get a new puzzle

	Get a random Lichess puzzle in JSON format.

If authenticated, only returns puzzles that the user has never seen before.

**DO NOT** use this endpoint to enumerate puzzles for mass download. Instead, download the [full public puzzle database](https://database.lichess.org/#puzzles).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PuzzlesAPIApiPuzzleNextRequest
	*/
	ApiPuzzleNext(ctx context.Context) PuzzlesAPIApiPuzzleNextRequest

	// ApiPuzzleNextExecute executes the request
	//  @return PuzzleAndGame
	ApiPuzzleNextExecute(r PuzzlesAPIApiPuzzleNextRequest) (*PuzzleAndGame, *http.Response, error)

	/*
	ApiPuzzleReplay Get puzzles to replay

	Gets the puzzle IDs of remaining puzzles to re-attempt in JSON format.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param days How many days to look back when aggregating puzzle results. 30 is sensible.
	@param theme The theme or opening to filter puzzles with.
	@return PuzzlesAPIApiPuzzleReplayRequest
	*/
	ApiPuzzleReplay(ctx context.Context, days int32, theme string) PuzzlesAPIApiPuzzleReplayRequest

	// ApiPuzzleReplayExecute executes the request
	//  @return PuzzleReplay
	ApiPuzzleReplayExecute(r PuzzlesAPIApiPuzzleReplayRequest) (*PuzzleReplay, *http.Response, error)

	/*
	ApiStormDashboard Get the storm dashboard of a player

	Download the [storm dashboard](https://lichess.org/storm/dashboard/mrbasso) of any player as JSON.
Contains the aggregated highscores, and the history of storm runs aggregated by days.
Use `?days=0` if you only care about the highscores.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param username Username of the player
	@return PuzzlesAPIApiStormDashboardRequest
	*/
	ApiStormDashboard(ctx context.Context, username string) PuzzlesAPIApiStormDashboardRequest

	// ApiStormDashboardExecute executes the request
	//  @return PuzzleStormDashboard
	ApiStormDashboardExecute(r PuzzlesAPIApiStormDashboardRequest) (*PuzzleStormDashboard, *http.Response, error)

	/*
	RacerGet Get puzzle race results

	Get the results of a [puzzle race](https://lichess.org/racer).
Returns information about players, puzzles, and the current status of the race.
- <https://lichess.org/racer>

Note that Lichess puzzle races are not persisted, and are only available
for 30 minutes. After that delay, they are permanently deleted.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The puzzle race ID
	@return PuzzlesAPIRacerGetRequest
	*/
	RacerGet(ctx context.Context, id string) PuzzlesAPIRacerGetRequest

	// RacerGetExecute executes the request
	//  @return PuzzleRaceResults
	RacerGetExecute(r PuzzlesAPIRacerGetRequest) (*PuzzleRaceResults, *http.Response, error)

	/*
	RacerPost Create and join a puzzle race

	Create a new private [puzzle race](https://lichess.org/racer).
The Lichess user who creates the race must join the race page,
and manually start the race when enough players have joined.
- <https://lichess.org/racer>


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return PuzzlesAPIRacerPostRequest
	*/
	RacerPost(ctx context.Context) PuzzlesAPIRacerPostRequest

	// RacerPostExecute executes the request
	//  @return PuzzleRacer
	RacerPostExecute(r PuzzlesAPIRacerPostRequest) (*PuzzleRacer, *http.Response, error)
}

// PuzzlesAPIService PuzzlesAPI service
type PuzzlesAPIService service

type PuzzlesAPIApiPuzzleActivityRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	max *int32
	before *int32
	since *int32
}

// How many entries to download. Leave empty to download all activity.
func (r PuzzlesAPIApiPuzzleActivityRequest) Max(max int32) PuzzlesAPIApiPuzzleActivityRequest {
	r.max = &max
	return r
}

// Download entries before this timestamp. Defaults to now. Use &#x60;before&#x60; and &#x60;max&#x60; for pagination.
func (r PuzzlesAPIApiPuzzleActivityRequest) Before(before int32) PuzzlesAPIApiPuzzleActivityRequest {
	r.before = &before
	return r
}

// Download entries since this timestamp. Defaults to account creation date.
func (r PuzzlesAPIApiPuzzleActivityRequest) Since(since int32) PuzzlesAPIApiPuzzleActivityRequest {
	r.since = &since
	return r
}

func (r PuzzlesAPIApiPuzzleActivityRequest) Execute() (*PuzzleActivity, *http.Response, error) {
	return r.ApiService.ApiPuzzleActivityExecute(r)
}

/*
ApiPuzzleActivity Get your puzzle activity

Download your puzzle activity in [ndjson](#description/streaming-with-nd-json) format.
Puzzle activity is sorted by reverse chronological order (most recent first)
We recommend streaming the response, for it can be very long.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PuzzlesAPIApiPuzzleActivityRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleActivity(ctx context.Context) PuzzlesAPIApiPuzzleActivityRequest {
	return PuzzlesAPIApiPuzzleActivityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PuzzleActivity
func (a *PuzzlesAPIService) ApiPuzzleActivityExecute(r PuzzlesAPIApiPuzzleActivityRequest) (*PuzzleActivity, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleActivity
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleActivity")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/activity"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.max != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "max", r.max, "form", "")
	}
	if r.before != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "before", r.before, "form", "")
	}
	if r.since != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "since", r.since, "form", "")
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

type PuzzlesAPIApiPuzzleBatchSelectRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	angle string
	difficulty *string
	nb *int32
	color *string
}

// The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous.
func (r PuzzlesAPIApiPuzzleBatchSelectRequest) Difficulty(difficulty string) PuzzlesAPIApiPuzzleBatchSelectRequest {
	r.difficulty = &difficulty
	return r
}

// How many puzzles to fetch. Just set it to &#x60;1&#x60; if you only need one puzzle. 
func (r PuzzlesAPIApiPuzzleBatchSelectRequest) Nb(nb int32) PuzzlesAPIApiPuzzleBatchSelectRequest {
	r.nb = &nb
	return r
}

// The color to play. Better left empty to automatically get 50% white. Currently only works when &#x60;nb&#x3D;1&#x60;. 
func (r PuzzlesAPIApiPuzzleBatchSelectRequest) Color(color string) PuzzlesAPIApiPuzzleBatchSelectRequest {
	r.color = &color
	return r
}

func (r PuzzlesAPIApiPuzzleBatchSelectRequest) Execute() (*PuzzleBatchSelect, *http.Response, error) {
	return r.ApiService.ApiPuzzleBatchSelectExecute(r)
}

/*
ApiPuzzleBatchSelect Get multiple puzzles at once

Get a batch of random Lichess puzzles in JSON format.

If authenticated, only returns puzzles that the user has never seen before.

**DO NOT** use this endpoint to enumerate puzzles for mass download. Instead, download the [full public puzzle database](https://database.lichess.org/#puzzles).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param angle The theme or opening to filter puzzles with. Recommended: `mix`.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
 @return PuzzlesAPIApiPuzzleBatchSelectRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleBatchSelect(ctx context.Context, angle string) PuzzlesAPIApiPuzzleBatchSelectRequest {
	return PuzzlesAPIApiPuzzleBatchSelectRequest{
		ApiService: a,
		ctx: ctx,
		angle: angle,
	}
}

// Execute executes the request
//  @return PuzzleBatchSelect
func (a *PuzzlesAPIService) ApiPuzzleBatchSelectExecute(r PuzzlesAPIApiPuzzleBatchSelectRequest) (*PuzzleBatchSelect, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleBatchSelect
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleBatchSelect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/batch/{angle}"
	localVarPath = strings.Replace(localVarPath, "{"+"angle"+"}", url.PathEscape(parameterValueToString(r.angle, "angle")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.difficulty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "difficulty", r.difficulty, "form", "")
	}
	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 15
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", defaultValue, "form", "")
		r.nb = &defaultValue
	}
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

type PuzzlesAPIApiPuzzleBatchSolveRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	angle string
	puzzleBatchSolveRequest *PuzzleBatchSolveRequest
	nb *int32
}

// List of solved puzzles
func (r PuzzlesAPIApiPuzzleBatchSolveRequest) PuzzleBatchSolveRequest(puzzleBatchSolveRequest PuzzleBatchSolveRequest) PuzzlesAPIApiPuzzleBatchSolveRequest {
	r.puzzleBatchSolveRequest = &puzzleBatchSolveRequest
	return r
}

// When &gt; 0, the response includes a new puzzle batch with that many puzzles.  This is equivalent to calling [/api/puzzle/batch/{angle}](#tag/puzzles/GET/api/puzzle/batch/{angle}), and can sometimes save a request. 
func (r PuzzlesAPIApiPuzzleBatchSolveRequest) Nb(nb int32) PuzzlesAPIApiPuzzleBatchSolveRequest {
	r.nb = &nb
	return r
}

func (r PuzzlesAPIApiPuzzleBatchSolveRequest) Execute() (*PuzzleBatchSolveResponse, *http.Response, error) {
	return r.ApiService.ApiPuzzleBatchSolveExecute(r)
}

/*
ApiPuzzleBatchSolve Solve multiple puzzles at once

Set puzzles as solved and update ratings.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param angle The theme or opening of the solved puzzles.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
 @return PuzzlesAPIApiPuzzleBatchSolveRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleBatchSolve(ctx context.Context, angle string) PuzzlesAPIApiPuzzleBatchSolveRequest {
	return PuzzlesAPIApiPuzzleBatchSolveRequest{
		ApiService: a,
		ctx: ctx,
		angle: angle,
	}
}

// Execute executes the request
//  @return PuzzleBatchSolveResponse
func (a *PuzzlesAPIService) ApiPuzzleBatchSolveExecute(r PuzzlesAPIApiPuzzleBatchSolveRequest) (*PuzzleBatchSolveResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleBatchSolveResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleBatchSolve")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/batch/{angle}"
	localVarPath = strings.Replace(localVarPath, "{"+"angle"+"}", url.PathEscape(parameterValueToString(r.angle, "angle")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.puzzleBatchSolveRequest == nil {
		return localVarReturnValue, nil, reportError("puzzleBatchSolveRequest is required and must be specified")
	}

	if r.nb != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", r.nb, "form", "")
	} else {
		var defaultValue int32 = 0
		parameterAddToHeaderOrQuery(localVarQueryParams, "nb", defaultValue, "form", "")
		r.nb = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	localVarPostBody = r.puzzleBatchSolveRequest
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

type PuzzlesAPIApiPuzzleDailyRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
}

func (r PuzzlesAPIApiPuzzleDailyRequest) Execute() (*PuzzleAndGame, *http.Response, error) {
	return r.ApiService.ApiPuzzleDailyExecute(r)
}

/*
ApiPuzzleDaily Get the daily puzzle

Get the daily Lichess puzzle in JSON format.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PuzzlesAPIApiPuzzleDailyRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleDaily(ctx context.Context) PuzzlesAPIApiPuzzleDailyRequest {
	return PuzzlesAPIApiPuzzleDailyRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PuzzleAndGame
func (a *PuzzlesAPIService) ApiPuzzleDailyExecute(r PuzzlesAPIApiPuzzleDailyRequest) (*PuzzleAndGame, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleAndGame
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleDaily")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/daily"

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

type PuzzlesAPIApiPuzzleDashboardRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	days int32
}

func (r PuzzlesAPIApiPuzzleDashboardRequest) Execute() (*PuzzleDashboard, *http.Response, error) {
	return r.ApiService.ApiPuzzleDashboardExecute(r)
}

/*
ApiPuzzleDashboard Get your puzzle dashboard

Download your [puzzle dashboard](https://lichess.org/training/dashboard/30/dashboard) as JSON.
Also includes all puzzle themes played, with aggregated results.
Allows re-creating the [improvement/strengths](https://lichess.org/training/dashboard/30/improvementAreas) interfaces.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param days How many days to look back when aggregating puzzle results. 30 is sensible.
 @return PuzzlesAPIApiPuzzleDashboardRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleDashboard(ctx context.Context, days int32) PuzzlesAPIApiPuzzleDashboardRequest {
	return PuzzlesAPIApiPuzzleDashboardRequest{
		ApiService: a,
		ctx: ctx,
		days: days,
	}
}

// Execute executes the request
//  @return PuzzleDashboard
func (a *PuzzlesAPIService) ApiPuzzleDashboardExecute(r PuzzlesAPIApiPuzzleDashboardRequest) (*PuzzleDashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleDashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/dashboard/{days}"
	localVarPath = strings.Replace(localVarPath, "{"+"days"+"}", url.PathEscape(parameterValueToString(r.days, "days")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.days < 1 {
		return localVarReturnValue, nil, reportError("days must be greater than 1")
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

type PuzzlesAPIApiPuzzleIdRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	id string
}

func (r PuzzlesAPIApiPuzzleIdRequest) Execute() (*PuzzleAndGame, *http.Response, error) {
	return r.ApiService.ApiPuzzleIdExecute(r)
}

/*
ApiPuzzleId Get a puzzle by its ID

Get a single Lichess puzzle in JSON format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The puzzle ID
 @return PuzzlesAPIApiPuzzleIdRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleId(ctx context.Context, id string) PuzzlesAPIApiPuzzleIdRequest {
	return PuzzlesAPIApiPuzzleIdRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PuzzleAndGame
func (a *PuzzlesAPIService) ApiPuzzleIdExecute(r PuzzlesAPIApiPuzzleIdRequest) (*PuzzleAndGame, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleAndGame
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleId")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/{id}"
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

type PuzzlesAPIApiPuzzleNextRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	angle *string
	difficulty *string
	color *string
}

// The theme or opening to filter puzzles with.  Available themes are listed in [the lichess source code](https://github.com/ornicar/lila/blob/master/translation/source/puzzleTheme.xml) and [the lichess training themes hyperlinks](https://lichess.org/training/themes). 
func (r PuzzlesAPIApiPuzzleNextRequest) Angle(angle string) PuzzlesAPIApiPuzzleNextRequest {
	r.angle = &angle
	return r
}

// The desired puzzle difficulty, relative to the authenticated user puzzle rating, or 1500 if anonymous.
func (r PuzzlesAPIApiPuzzleNextRequest) Difficulty(difficulty string) PuzzlesAPIApiPuzzleNextRequest {
	r.difficulty = &difficulty
	return r
}

// The color to play. Better left empty to automatically get 50% white.
func (r PuzzlesAPIApiPuzzleNextRequest) Color(color string) PuzzlesAPIApiPuzzleNextRequest {
	r.color = &color
	return r
}

func (r PuzzlesAPIApiPuzzleNextRequest) Execute() (*PuzzleAndGame, *http.Response, error) {
	return r.ApiService.ApiPuzzleNextExecute(r)
}

/*
ApiPuzzleNext Get a new puzzle

Get a random Lichess puzzle in JSON format.

If authenticated, only returns puzzles that the user has never seen before.

**DO NOT** use this endpoint to enumerate puzzles for mass download. Instead, download the [full public puzzle database](https://database.lichess.org/#puzzles).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PuzzlesAPIApiPuzzleNextRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleNext(ctx context.Context) PuzzlesAPIApiPuzzleNextRequest {
	return PuzzlesAPIApiPuzzleNextRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PuzzleAndGame
func (a *PuzzlesAPIService) ApiPuzzleNextExecute(r PuzzlesAPIApiPuzzleNextRequest) (*PuzzleAndGame, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleAndGame
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleNext")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/next"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.angle != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "angle", r.angle, "form", "")
	}
	if r.difficulty != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "difficulty", r.difficulty, "form", "")
	}
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

type PuzzlesAPIApiPuzzleReplayRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	days int32
	theme string
}

func (r PuzzlesAPIApiPuzzleReplayRequest) Execute() (*PuzzleReplay, *http.Response, error) {
	return r.ApiService.ApiPuzzleReplayExecute(r)
}

/*
ApiPuzzleReplay Get puzzles to replay

Gets the puzzle IDs of remaining puzzles to re-attempt in JSON format.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param days How many days to look back when aggregating puzzle results. 30 is sensible.
 @param theme The theme or opening to filter puzzles with.
 @return PuzzlesAPIApiPuzzleReplayRequest
*/
func (a *PuzzlesAPIService) ApiPuzzleReplay(ctx context.Context, days int32, theme string) PuzzlesAPIApiPuzzleReplayRequest {
	return PuzzlesAPIApiPuzzleReplayRequest{
		ApiService: a,
		ctx: ctx,
		days: days,
		theme: theme,
	}
}

// Execute executes the request
//  @return PuzzleReplay
func (a *PuzzlesAPIService) ApiPuzzleReplayExecute(r PuzzlesAPIApiPuzzleReplayRequest) (*PuzzleReplay, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleReplay
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiPuzzleReplay")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/puzzle/replay/{days}/{theme}"
	localVarPath = strings.Replace(localVarPath, "{"+"days"+"}", url.PathEscape(parameterValueToString(r.days, "days")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"theme"+"}", url.PathEscape(parameterValueToString(r.theme, "theme")), -1)

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
			var v ApiPuzzleReplay404Response
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

type PuzzlesAPIApiStormDashboardRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	username string
	days *int32
}

// How many days of history to return
func (r PuzzlesAPIApiStormDashboardRequest) Days(days int32) PuzzlesAPIApiStormDashboardRequest {
	r.days = &days
	return r
}

func (r PuzzlesAPIApiStormDashboardRequest) Execute() (*PuzzleStormDashboard, *http.Response, error) {
	return r.ApiService.ApiStormDashboardExecute(r)
}

/*
ApiStormDashboard Get the storm dashboard of a player

Download the [storm dashboard](https://lichess.org/storm/dashboard/mrbasso) of any player as JSON.
Contains the aggregated highscores, and the history of storm runs aggregated by days.
Use `?days=0` if you only care about the highscores.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param username Username of the player
 @return PuzzlesAPIApiStormDashboardRequest
*/
func (a *PuzzlesAPIService) ApiStormDashboard(ctx context.Context, username string) PuzzlesAPIApiStormDashboardRequest {
	return PuzzlesAPIApiStormDashboardRequest{
		ApiService: a,
		ctx: ctx,
		username: username,
	}
}

// Execute executes the request
//  @return PuzzleStormDashboard
func (a *PuzzlesAPIService) ApiStormDashboardExecute(r PuzzlesAPIApiStormDashboardRequest) (*PuzzleStormDashboard, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleStormDashboard
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.ApiStormDashboard")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/storm/dashboard/{username}"
	localVarPath = strings.Replace(localVarPath, "{"+"username"+"}", url.PathEscape(parameterValueToString(r.username, "username")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.days != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "days", r.days, "form", "")
	} else {
		var defaultValue int32 = 30
		parameterAddToHeaderOrQuery(localVarQueryParams, "days", defaultValue, "form", "")
		r.days = &defaultValue
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

type PuzzlesAPIRacerGetRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
	id string
}

func (r PuzzlesAPIRacerGetRequest) Execute() (*PuzzleRaceResults, *http.Response, error) {
	return r.ApiService.RacerGetExecute(r)
}

/*
RacerGet Get puzzle race results

Get the results of a [puzzle race](https://lichess.org/racer).
Returns information about players, puzzles, and the current status of the race.
- <https://lichess.org/racer>

Note that Lichess puzzle races are not persisted, and are only available
for 30 minutes. After that delay, they are permanently deleted.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The puzzle race ID
 @return PuzzlesAPIRacerGetRequest
*/
func (a *PuzzlesAPIService) RacerGet(ctx context.Context, id string) PuzzlesAPIRacerGetRequest {
	return PuzzlesAPIRacerGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PuzzleRaceResults
func (a *PuzzlesAPIService) RacerGetExecute(r PuzzlesAPIRacerGetRequest) (*PuzzleRaceResults, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleRaceResults
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.RacerGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/racer/{id}"
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

type PuzzlesAPIRacerPostRequest struct {
	ctx context.Context
	ApiService PuzzlesAPI
}

func (r PuzzlesAPIRacerPostRequest) Execute() (*PuzzleRacer, *http.Response, error) {
	return r.ApiService.RacerPostExecute(r)
}

/*
RacerPost Create and join a puzzle race

Create a new private [puzzle race](https://lichess.org/racer).
The Lichess user who creates the race must join the race page,
and manually start the race when enough players have joined.
- <https://lichess.org/racer>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return PuzzlesAPIRacerPostRequest
*/
func (a *PuzzlesAPIService) RacerPost(ctx context.Context) PuzzlesAPIRacerPostRequest {
	return PuzzlesAPIRacerPostRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PuzzleRacer
func (a *PuzzlesAPIService) RacerPostExecute(r PuzzlesAPIRacerPostRequest) (*PuzzleRacer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PuzzleRacer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "PuzzlesAPIService.RacerPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/racer"

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
