/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.147
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


type TournamentsSwissAPI interface {

	/*
	ApiSwissJoin Join a Swiss tournament

	Join a Swiss tournament, possibly with a password.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIApiSwissJoinRequest
	*/
	ApiSwissJoin(ctx context.Context, id string) TournamentsSwissAPIApiSwissJoinRequest

	// ApiSwissJoinExecute executes the request
	//  @return Ok
	ApiSwissJoinExecute(r TournamentsSwissAPIApiSwissJoinRequest) (*Ok, *http.Response, error)

	/*
	ApiSwissNew Create a new Swiss tournament

	Create a Swiss tournament for your team.
This endpoint mirrors the Swiss tournament form from your team pagee.
You can create up to 12 tournaments per day.
Additional restrictions:
  - clock.limit + clock.increment > 0
  - 15s and 0+1 variant tournaments cannot be rated


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId ID of the team
	@return TournamentsSwissAPIApiSwissNewRequest
	*/
	ApiSwissNew(ctx context.Context, teamId string) TournamentsSwissAPIApiSwissNewRequest

	// ApiSwissNewExecute executes the request
	//  @return SwissTournament
	ApiSwissNewExecute(r TournamentsSwissAPIApiSwissNewRequest) (*SwissTournament, *http.Response, error)

	/*
	ApiSwissScheduleNextRound Manually schedule the next round

	Manually schedule the next round date and time of a Swiss tournament.
This sets the `roundInterval` field to `99999999`, i.e. manual scheduling.
All further rounds will need to be manually scheduled, unless the `roundInterval` field is changed back to automatic scheduling.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIApiSwissScheduleNextRoundRequest
	*/
	ApiSwissScheduleNextRound(ctx context.Context, id string) TournamentsSwissAPIApiSwissScheduleNextRoundRequest

	// ApiSwissScheduleNextRoundExecute executes the request
	ApiSwissScheduleNextRoundExecute(r TournamentsSwissAPIApiSwissScheduleNextRoundRequest) (*http.Response, error)

	/*
	ApiSwissTerminate Terminate a Swiss tournament

	Terminate a Swiss tournament


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Swiss tournament ID.
	@return TournamentsSwissAPIApiSwissTerminateRequest
	*/
	ApiSwissTerminate(ctx context.Context, id string) TournamentsSwissAPIApiSwissTerminateRequest

	// ApiSwissTerminateExecute executes the request
	//  @return Ok
	ApiSwissTerminateExecute(r TournamentsSwissAPIApiSwissTerminateRequest) (*Ok, *http.Response, error)

	/*
	ApiSwissUpdate Update a Swiss tournament

	Update a Swiss tournament.
Be mindful not to make important changes to ongoing tournaments.
Additional restrictions:
  - clock.limit + clock.increment > 0
  - 15s and 0+1 variant tournaments cannot be rated


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIApiSwissUpdateRequest
	*/
	ApiSwissUpdate(ctx context.Context, id string) TournamentsSwissAPIApiSwissUpdateRequest

	// ApiSwissUpdateExecute executes the request
	//  @return SwissTournament
	ApiSwissUpdateExecute(r TournamentsSwissAPIApiSwissUpdateRequest) (*SwissTournament, *http.Response, error)

	/*
	ApiSwissWithdraw Pause or leave a swiss tournament

	Leave a future Swiss tournament, or take a break on an ongoing Swiss tournament.
It's possible to join again later. Points are preserved.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIApiSwissWithdrawRequest
	*/
	ApiSwissWithdraw(ctx context.Context, id string) TournamentsSwissAPIApiSwissWithdrawRequest

	// ApiSwissWithdrawExecute executes the request
	//  @return Ok
	ApiSwissWithdrawExecute(r TournamentsSwissAPIApiSwissWithdrawRequest) (*Ok, *http.Response, error)

	/*
	ApiTeamSwiss Get team swiss tournaments

	Get all swiss tournaments of a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param teamId
	@return TournamentsSwissAPIApiTeamSwissRequest
	*/
	ApiTeamSwiss(ctx context.Context, teamId string) TournamentsSwissAPIApiTeamSwissRequest

	// ApiTeamSwissExecute executes the request
	//  @return SwissTournament
	ApiTeamSwissExecute(r TournamentsSwissAPIApiTeamSwissRequest) (*SwissTournament, *http.Response, error)

	/*
	GamesBySwiss Export games of a Swiss tournament

	Download games of a swiss tournament in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by chronological order.
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIGamesBySwissRequest
	*/
	GamesBySwiss(ctx context.Context, id string) TournamentsSwissAPIGamesBySwissRequest

	// GamesBySwissExecute executes the request
	//  @return string
	GamesBySwissExecute(r TournamentsSwissAPIGamesBySwissRequest) (string, *http.Response, error)

	/*
	ResultsBySwiss Get results of a swiss tournament

	Players of a swiss tournament, with their score and performance, sorted by rank (best first).
Players are streamed as [ndjson](#description/streaming-with-nd-json).
If called on an ongoing tournament, results can be inconsistent
due to ranking changes while the players are being streamed.
Use on finished tournaments for guaranteed consistency.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPIResultsBySwissRequest
	*/
	ResultsBySwiss(ctx context.Context, id string) TournamentsSwissAPIResultsBySwissRequest

	// ResultsBySwissExecute executes the request
	//  @return ResultsBySwiss200Response
	ResultsBySwissExecute(r TournamentsSwissAPIResultsBySwissRequest) (*ResultsBySwiss200Response, *http.Response, error)

	/*
	Swiss Get info about a Swiss tournament

	Get detailed info about a Swiss tournament.


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The Swiss tournament ID.
	@return TournamentsSwissAPISwissRequest
	*/
	Swiss(ctx context.Context, id string) TournamentsSwissAPISwissRequest

	// SwissExecute executes the request
	//  @return SwissTournament
	SwissExecute(r TournamentsSwissAPISwissRequest) (*SwissTournament, *http.Response, error)

	/*
	SwissTrf Export TRF of a Swiss tournament

	Download a tournament in the Tournament Report File format, the FIDE standard.
Documentation: <https://www.fide.com/FIDE/handbook/C04Annex2_TRF16.pdf>
Example: <https://lichess.org/swiss/j8rtJ5GL.trf>


	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param id The tournament ID.
	@return TournamentsSwissAPISwissTrfRequest
	*/
	SwissTrf(ctx context.Context, id string) TournamentsSwissAPISwissTrfRequest

	// SwissTrfExecute executes the request
	//  @return string
	SwissTrfExecute(r TournamentsSwissAPISwissTrfRequest) (string, *http.Response, error)
}

// TournamentsSwissAPIService TournamentsSwissAPI service
type TournamentsSwissAPIService service

type TournamentsSwissAPIApiSwissJoinRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
	password *string
}

// The tournament password, if one is required
func (r TournamentsSwissAPIApiSwissJoinRequest) Password(password string) TournamentsSwissAPIApiSwissJoinRequest {
	r.password = &password
	return r
}

func (r TournamentsSwissAPIApiSwissJoinRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ApiSwissJoinExecute(r)
}

/*
ApiSwissJoin Join a Swiss tournament

Join a Swiss tournament, possibly with a password.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIApiSwissJoinRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissJoin(ctx context.Context, id string) TournamentsSwissAPIApiSwissJoinRequest {
	return TournamentsSwissAPIApiSwissJoinRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ok
func (a *TournamentsSwissAPIService) ApiSwissJoinExecute(r TournamentsSwissAPIApiSwissJoinRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissJoin")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/join"
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

type TournamentsSwissAPIApiSwissNewRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	teamId string
	clockLimit *int32
	clockIncrement *int32
	nbRounds *int32
	name *string
	startsAt *int64
	roundInterval *int32
	variant *VariantKey
	position *string
	description *string
	rated *bool
	password *string
	forbiddenPairings *string
	manualPairings *string
	chatFor *int32
	conditionsMinRatingRating *int32
	conditionsMaxRatingRating *int32
	conditionsNbRatedGameNb *int32
	conditionsPlayYourGames *bool
	conditionsAllowList *string
}

// Clock initial time in seconds
func (r TournamentsSwissAPIApiSwissNewRequest) ClockLimit(clockLimit int32) TournamentsSwissAPIApiSwissNewRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds
func (r TournamentsSwissAPIApiSwissNewRequest) ClockIncrement(clockIncrement int32) TournamentsSwissAPIApiSwissNewRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Maximum number of rounds to play
func (r TournamentsSwissAPIApiSwissNewRequest) NbRounds(nbRounds int32) TournamentsSwissAPIApiSwissNewRequest {
	r.nbRounds = &nbRounds
	return r
}

// The tournament name. Leave empty to get a random Grandmaster name
func (r TournamentsSwissAPIApiSwissNewRequest) Name(name string) TournamentsSwissAPIApiSwissNewRequest {
	r.name = &name
	return r
}

// Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation.
func (r TournamentsSwissAPIApiSwissNewRequest) StartsAt(startsAt int64) TournamentsSwissAPIApiSwissNewRequest {
	r.startsAt = &startsAt
	return r
}

// How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI. If empty or -1, a sensible value is picked automatically. 
func (r TournamentsSwissAPIApiSwissNewRequest) RoundInterval(roundInterval int32) TournamentsSwissAPIApiSwissNewRequest {
	r.roundInterval = &roundInterval
	return r
}

func (r TournamentsSwissAPIApiSwissNewRequest) Variant(variant VariantKey) TournamentsSwissAPIApiSwissNewRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated.
func (r TournamentsSwissAPIApiSwissNewRequest) Position(position string) TournamentsSwissAPIApiSwissNewRequest {
	r.position = &position
	return r
}

// Anything you want to tell players about the tournament
func (r TournamentsSwissAPIApiSwissNewRequest) Description(description string) TournamentsSwissAPIApiSwissNewRequest {
	r.description = &description
	return r
}

// Games are rated and impact players ratings
func (r TournamentsSwissAPIApiSwissNewRequest) Rated(rated bool) TournamentsSwissAPIApiSwissNewRequest {
	r.rated = &rated
	return r
}

// Make the tournament private and restrict access with a password.
func (r TournamentsSwissAPIApiSwissNewRequest) Password(password string) TournamentsSwissAPIApiSwissNewRequest {
	r.password = &password
	return r
}

// Usernames of players that must not play together. Two usernames per line, separated by a space. 
func (r TournamentsSwissAPIApiSwissNewRequest) ForbiddenPairings(forbiddenPairings string) TournamentsSwissAPIApiSwissNewRequest {
	r.forbiddenPairings = &forbiddenPairings
	return r
}

// Manual pairings for the next round. Two usernames per line, separated by a space. Example: &#x60;&#x60;&#x60; PlayerA PlayerB PlayerC PlayerD &#x60;&#x60;&#x60; To give a bye (1 point) to a player instead of a pairing, add a line like so: &#x60;&#x60;&#x60; PlayerE 1 &#x60;&#x60;&#x60; Missing players will be considered absent and get zero points. 
func (r TournamentsSwissAPIApiSwissNewRequest) ManualPairings(manualPairings string) TournamentsSwissAPIApiSwissNewRequest {
	r.manualPairings = &manualPairings
	return r
}

// Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players 
func (r TournamentsSwissAPIApiSwissNewRequest) ChatFor(chatFor int32) TournamentsSwissAPIApiSwissNewRequest {
	r.chatFor = &chatFor
	return r
}

// Minimum rating to join. Leave empty to let everyone join the tournament.
func (r TournamentsSwissAPIApiSwissNewRequest) ConditionsMinRatingRating(conditionsMinRatingRating int32) TournamentsSwissAPIApiSwissNewRequest {
	r.conditionsMinRatingRating = &conditionsMinRatingRating
	return r
}

// Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament.
func (r TournamentsSwissAPIApiSwissNewRequest) ConditionsMaxRatingRating(conditionsMaxRatingRating int32) TournamentsSwissAPIApiSwissNewRequest {
	r.conditionsMaxRatingRating = &conditionsMaxRatingRating
	return r
}

// Minimum number of rated games required to join.
func (r TournamentsSwissAPIApiSwissNewRequest) ConditionsNbRatedGameNb(conditionsNbRatedGameNb int32) TournamentsSwissAPIApiSwissNewRequest {
	r.conditionsNbRatedGameNb = &conditionsNbRatedGameNb
	return r
}

// Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won&#39;t be able to enter yours. This results in a better swiss experience for the players who actually show up. 
func (r TournamentsSwissAPIApiSwissNewRequest) ConditionsPlayYourGames(conditionsPlayYourGames bool) TournamentsSwissAPIApiSwissNewRequest {
	r.conditionsPlayYourGames = &conditionsPlayYourGames
	return r
}

// Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60; 
func (r TournamentsSwissAPIApiSwissNewRequest) ConditionsAllowList(conditionsAllowList string) TournamentsSwissAPIApiSwissNewRequest {
	r.conditionsAllowList = &conditionsAllowList
	return r
}

func (r TournamentsSwissAPIApiSwissNewRequest) Execute() (*SwissTournament, *http.Response, error) {
	return r.ApiService.ApiSwissNewExecute(r)
}

/*
ApiSwissNew Create a new Swiss tournament

Create a Swiss tournament for your team.
This endpoint mirrors the Swiss tournament form from your team pagee.
You can create up to 12 tournaments per day.
Additional restrictions:
  - clock.limit + clock.increment > 0
  - 15s and 0+1 variant tournaments cannot be rated


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId ID of the team
 @return TournamentsSwissAPIApiSwissNewRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissNew(ctx context.Context, teamId string) TournamentsSwissAPIApiSwissNewRequest {
	return TournamentsSwissAPIApiSwissNewRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return SwissTournament
func (a *TournamentsSwissAPIService) ApiSwissNewExecute(r TournamentsSwissAPIApiSwissNewRequest) (*SwissTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwissTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissNew")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/new/{teamId}"
	localVarPath = strings.Replace(localVarPath, "{"+"teamId"+"}", url.PathEscape(parameterValueToString(r.teamId, "teamId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clockLimit == nil {
		return localVarReturnValue, nil, reportError("clockLimit is required and must be specified")
	}
	if r.clockIncrement == nil {
		return localVarReturnValue, nil, reportError("clockIncrement is required and must be specified")
	}
	if *r.clockIncrement < 0 {
		return localVarReturnValue, nil, reportError("clockIncrement must be greater than 0")
	}
	if *r.clockIncrement > 120 {
		return localVarReturnValue, nil, reportError("clockIncrement must be less than 120")
	}
	if r.nbRounds == nil {
		return localVarReturnValue, nil, reportError("nbRounds is required and must be specified")
	}
	if *r.nbRounds < 3 {
		return localVarReturnValue, nil, reportError("nbRounds must be greater than 3")
	}
	if *r.nbRounds > 100 {
		return localVarReturnValue, nil, reportError("nbRounds must be less than 100")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "clock.limit", r.clockLimit, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "clock.increment", r.clockIncrement, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "nbRounds", r.nbRounds, "", "")
	if r.startsAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAt", r.startsAt, "", "")
	}
	if r.roundInterval != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "roundInterval", r.roundInterval, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.position != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "position", r.position, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.forbiddenPairings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "forbiddenPairings", r.forbiddenPairings, "", "")
	}
	if r.manualPairings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "manualPairings", r.manualPairings, "", "")
	}
	if r.chatFor != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "chatFor", r.chatFor, "", "")
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
	if r.conditionsPlayYourGames != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.playYourGames", r.conditionsPlayYourGames, "", "")
	}
	if r.conditionsAllowList != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.allowList", r.conditionsAllowList, "", "")
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

type TournamentsSwissAPIApiSwissScheduleNextRoundRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
	date *int64
}

// Timestamp in milliseconds to start the next round at a given date and time.
func (r TournamentsSwissAPIApiSwissScheduleNextRoundRequest) Date(date int64) TournamentsSwissAPIApiSwissScheduleNextRoundRequest {
	r.date = &date
	return r
}

func (r TournamentsSwissAPIApiSwissScheduleNextRoundRequest) Execute() (*http.Response, error) {
	return r.ApiService.ApiSwissScheduleNextRoundExecute(r)
}

/*
ApiSwissScheduleNextRound Manually schedule the next round

Manually schedule the next round date and time of a Swiss tournament.
This sets the `roundInterval` field to `99999999`, i.e. manual scheduling.
All further rounds will need to be manually scheduled, unless the `roundInterval` field is changed back to automatic scheduling.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIApiSwissScheduleNextRoundRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissScheduleNextRound(ctx context.Context, id string) TournamentsSwissAPIApiSwissScheduleNextRoundRequest {
	return TournamentsSwissAPIApiSwissScheduleNextRoundRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
func (a *TournamentsSwissAPIService) ApiSwissScheduleNextRoundExecute(r TournamentsSwissAPIApiSwissScheduleNextRoundRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissScheduleNextRound")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/schedule-next-round"
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
	if r.date != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "date", r.date, "", "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v SwissUnauthorisedEdit
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TournamentsSwissAPIApiSwissTerminateRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
}

func (r TournamentsSwissAPIApiSwissTerminateRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ApiSwissTerminateExecute(r)
}

/*
ApiSwissTerminate Terminate a Swiss tournament

Terminate a Swiss tournament


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The Swiss tournament ID.
 @return TournamentsSwissAPIApiSwissTerminateRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissTerminate(ctx context.Context, id string) TournamentsSwissAPIApiSwissTerminateRequest {
	return TournamentsSwissAPIApiSwissTerminateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ok
func (a *TournamentsSwissAPIService) ApiSwissTerminateExecute(r TournamentsSwissAPIApiSwissTerminateRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissTerminate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/terminate"
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

type TournamentsSwissAPIApiSwissUpdateRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
	clockLimit *int32
	clockIncrement *int32
	nbRounds *int32
	name *string
	startsAt *int64
	roundInterval *int32
	variant *VariantKey
	position *string
	description *string
	rated *bool
	password *string
	forbiddenPairings *string
	manualPairings *string
	chatFor *int32
	conditionsMinRatingRating *int32
	conditionsMaxRatingRating *int32
	conditionsNbRatedGameNb *int32
	conditionsPlayYourGames *bool
	conditionsAllowList *string
}

// Clock initial time in seconds
func (r TournamentsSwissAPIApiSwissUpdateRequest) ClockLimit(clockLimit int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.clockLimit = &clockLimit
	return r
}

// Clock increment in seconds
func (r TournamentsSwissAPIApiSwissUpdateRequest) ClockIncrement(clockIncrement int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.clockIncrement = &clockIncrement
	return r
}

// Maximum number of rounds to play
func (r TournamentsSwissAPIApiSwissUpdateRequest) NbRounds(nbRounds int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.nbRounds = &nbRounds
	return r
}

// The tournament name. Leave empty to get a random Grandmaster name
func (r TournamentsSwissAPIApiSwissUpdateRequest) Name(name string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.name = &name
	return r
}

// Timestamp in milliseconds to start the tournament at a given date and time. By default, it starts 10 minutes after creation.
func (r TournamentsSwissAPIApiSwissUpdateRequest) StartsAt(startsAt int64) TournamentsSwissAPIApiSwissUpdateRequest {
	r.startsAt = &startsAt
	return r
}

// How long to wait between each round, in seconds. Set to 99999999 to manually schedule each round from the tournament UI. If empty or -1, a sensible value is picked automatically. 
func (r TournamentsSwissAPIApiSwissUpdateRequest) RoundInterval(roundInterval int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.roundInterval = &roundInterval
	return r
}

func (r TournamentsSwissAPIApiSwissUpdateRequest) Variant(variant VariantKey) TournamentsSwissAPIApiSwissUpdateRequest {
	r.variant = &variant
	return r
}

// Custom initial position (in X-FEN). Variant must be standard and the game cannot be rated.
func (r TournamentsSwissAPIApiSwissUpdateRequest) Position(position string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.position = &position
	return r
}

// Anything you want to tell players about the tournament
func (r TournamentsSwissAPIApiSwissUpdateRequest) Description(description string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.description = &description
	return r
}

// Games are rated and impact players ratings
func (r TournamentsSwissAPIApiSwissUpdateRequest) Rated(rated bool) TournamentsSwissAPIApiSwissUpdateRequest {
	r.rated = &rated
	return r
}

// Make the tournament private and restrict access with a password.
func (r TournamentsSwissAPIApiSwissUpdateRequest) Password(password string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.password = &password
	return r
}

// Usernames of players that must not play together. Two usernames per line, separated by a space. 
func (r TournamentsSwissAPIApiSwissUpdateRequest) ForbiddenPairings(forbiddenPairings string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.forbiddenPairings = &forbiddenPairings
	return r
}

// Manual pairings for the next round. Two usernames per line, separated by a space. Example: &#x60;&#x60;&#x60; PlayerA PlayerB PlayerC PlayerD &#x60;&#x60;&#x60; To give a bye (1 point) to a player instead of a pairing, add a line like so: &#x60;&#x60;&#x60; PlayerE 1 &#x60;&#x60;&#x60; Missing players will be considered absent and get zero points. 
func (r TournamentsSwissAPIApiSwissUpdateRequest) ManualPairings(manualPairings string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.manualPairings = &manualPairings
	return r
}

// Who can read and write in the chat. - 0  &#x3D; No-one - 10 &#x3D; Only team leaders - 20 &#x3D; Only team members - 30 &#x3D; All Lichess players 
func (r TournamentsSwissAPIApiSwissUpdateRequest) ChatFor(chatFor int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.chatFor = &chatFor
	return r
}

// Minimum rating to join. Leave empty to let everyone join the tournament.
func (r TournamentsSwissAPIApiSwissUpdateRequest) ConditionsMinRatingRating(conditionsMinRatingRating int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.conditionsMinRatingRating = &conditionsMinRatingRating
	return r
}

// Maximum rating to join. Based on best rating reached in the last 7 days. Leave empty to let everyone join the tournament.
func (r TournamentsSwissAPIApiSwissUpdateRequest) ConditionsMaxRatingRating(conditionsMaxRatingRating int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.conditionsMaxRatingRating = &conditionsMaxRatingRating
	return r
}

// Minimum number of rated games required to join.
func (r TournamentsSwissAPIApiSwissUpdateRequest) ConditionsNbRatedGameNb(conditionsNbRatedGameNb int32) TournamentsSwissAPIApiSwissUpdateRequest {
	r.conditionsNbRatedGameNb = &conditionsNbRatedGameNb
	return r
}

// Only let players join if they have played their last swiss game. If they failed to show up in a recent swiss event, they won&#39;t be able to enter yours. This results in a better swiss experience for the players who actually show up. 
func (r TournamentsSwissAPIApiSwissUpdateRequest) ConditionsPlayYourGames(conditionsPlayYourGames bool) TournamentsSwissAPIApiSwissUpdateRequest {
	r.conditionsPlayYourGames = &conditionsPlayYourGames
	return r
}

// Predefined list of usernames that are allowed to join, separated by commas. If this list is non-empty, then usernames absent from this list will be forbidden to join. Adding &#x60;%titled&#x60; to the list additionally allows any titled player to join. Example: &#x60;thibault,german11,%titled&#x60; 
func (r TournamentsSwissAPIApiSwissUpdateRequest) ConditionsAllowList(conditionsAllowList string) TournamentsSwissAPIApiSwissUpdateRequest {
	r.conditionsAllowList = &conditionsAllowList
	return r
}

func (r TournamentsSwissAPIApiSwissUpdateRequest) Execute() (*SwissTournament, *http.Response, error) {
	return r.ApiService.ApiSwissUpdateExecute(r)
}

/*
ApiSwissUpdate Update a Swiss tournament

Update a Swiss tournament.
Be mindful not to make important changes to ongoing tournaments.
Additional restrictions:
  - clock.limit + clock.increment > 0
  - 15s and 0+1 variant tournaments cannot be rated


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIApiSwissUpdateRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissUpdate(ctx context.Context, id string) TournamentsSwissAPIApiSwissUpdateRequest {
	return TournamentsSwissAPIApiSwissUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SwissTournament
func (a *TournamentsSwissAPIService) ApiSwissUpdateExecute(r TournamentsSwissAPIApiSwissUpdateRequest) (*SwissTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwissTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/edit"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clockLimit == nil {
		return localVarReturnValue, nil, reportError("clockLimit is required and must be specified")
	}
	if r.clockIncrement == nil {
		return localVarReturnValue, nil, reportError("clockIncrement is required and must be specified")
	}
	if *r.clockIncrement < 0 {
		return localVarReturnValue, nil, reportError("clockIncrement must be greater than 0")
	}
	if *r.clockIncrement > 120 {
		return localVarReturnValue, nil, reportError("clockIncrement must be less than 120")
	}
	if r.nbRounds == nil {
		return localVarReturnValue, nil, reportError("nbRounds is required and must be specified")
	}
	if *r.nbRounds < 3 {
		return localVarReturnValue, nil, reportError("nbRounds must be greater than 3")
	}
	if *r.nbRounds > 100 {
		return localVarReturnValue, nil, reportError("nbRounds must be less than 100")
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
	parameterAddToHeaderOrQuery(localVarFormParams, "clock.limit", r.clockLimit, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "clock.increment", r.clockIncrement, "", "")
	parameterAddToHeaderOrQuery(localVarFormParams, "nbRounds", r.nbRounds, "", "")
	if r.startsAt != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "startsAt", r.startsAt, "", "")
	}
	if r.roundInterval != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "roundInterval", r.roundInterval, "", "")
	}
	if r.variant != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "variant", r.variant, "", "")
	}
	if r.position != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "position", r.position, "", "")
	}
	if r.description != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "description", r.description, "", "")
	}
	if r.rated != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "rated", r.rated, "", "")
	}
	if r.password != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "password", r.password, "", "")
	}
	if r.forbiddenPairings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "forbiddenPairings", r.forbiddenPairings, "", "")
	}
	if r.manualPairings != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "manualPairings", r.manualPairings, "", "")
	}
	if r.chatFor != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "chatFor", r.chatFor, "", "")
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
	if r.conditionsPlayYourGames != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.playYourGames", r.conditionsPlayYourGames, "", "")
	}
	if r.conditionsAllowList != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "conditions.allowList", r.conditionsAllowList, "", "")
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
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v SwissUnauthorisedEdit
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

type TournamentsSwissAPIApiSwissWithdrawRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
}

func (r TournamentsSwissAPIApiSwissWithdrawRequest) Execute() (*Ok, *http.Response, error) {
	return r.ApiService.ApiSwissWithdrawExecute(r)
}

/*
ApiSwissWithdraw Pause or leave a swiss tournament

Leave a future Swiss tournament, or take a break on an ongoing Swiss tournament.
It's possible to join again later. Points are preserved.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIApiSwissWithdrawRequest
*/
func (a *TournamentsSwissAPIService) ApiSwissWithdraw(ctx context.Context, id string) TournamentsSwissAPIApiSwissWithdrawRequest {
	return TournamentsSwissAPIApiSwissWithdrawRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return Ok
func (a *TournamentsSwissAPIService) ApiSwissWithdrawExecute(r TournamentsSwissAPIApiSwissWithdrawRequest) (*Ok, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Ok
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiSwissWithdraw")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/withdraw"
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

type TournamentsSwissAPIApiTeamSwissRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	teamId string
	max *int32
	status *SwissStatus
	createdBy *string
	name *string
}

// How many tournaments to download.
func (r TournamentsSwissAPIApiTeamSwissRequest) Max(max int32) TournamentsSwissAPIApiTeamSwissRequest {
	r.max = &max
	return r
}

// [Filter] Only swiss tournaments in this current state. 
func (r TournamentsSwissAPIApiTeamSwissRequest) Status(status SwissStatus) TournamentsSwissAPIApiTeamSwissRequest {
	r.status = &status
	return r
}

// [Filter] Only swiss tournaments created by a given user. 
func (r TournamentsSwissAPIApiTeamSwissRequest) CreatedBy(createdBy string) TournamentsSwissAPIApiTeamSwissRequest {
	r.createdBy = &createdBy
	return r
}

// [Filter] Only swiss tournaments with a given name. 
func (r TournamentsSwissAPIApiTeamSwissRequest) Name(name string) TournamentsSwissAPIApiTeamSwissRequest {
	r.name = &name
	return r
}

func (r TournamentsSwissAPIApiTeamSwissRequest) Execute() (*SwissTournament, *http.Response, error) {
	return r.ApiService.ApiTeamSwissExecute(r)
}

/*
ApiTeamSwiss Get team swiss tournaments

Get all swiss tournaments of a team.
Tournaments are sorted by reverse chronological order of start date (last starting first).
Tournaments are streamed as [ndjson](#description/streaming-with-nd-json).


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param teamId
 @return TournamentsSwissAPIApiTeamSwissRequest
*/
func (a *TournamentsSwissAPIService) ApiTeamSwiss(ctx context.Context, teamId string) TournamentsSwissAPIApiTeamSwissRequest {
	return TournamentsSwissAPIApiTeamSwissRequest{
		ApiService: a,
		ctx: ctx,
		teamId: teamId,
	}
}

// Execute executes the request
//  @return SwissTournament
func (a *TournamentsSwissAPIService) ApiTeamSwissExecute(r TournamentsSwissAPIApiTeamSwissRequest) (*SwissTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwissTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ApiTeamSwiss")
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

type TournamentsSwissAPIGamesBySwissRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
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
func (r TournamentsSwissAPIGamesBySwissRequest) Accept(accept string) TournamentsSwissAPIGamesBySwissRequest {
	r.accept = &accept
	return r
}

// Only the games played by a given player
func (r TournamentsSwissAPIGamesBySwissRequest) Player(player string) TournamentsSwissAPIGamesBySwissRequest {
	r.player = &player
	return r
}

// Include the PGN moves.
func (r TournamentsSwissAPIGamesBySwissRequest) Moves(moves bool) TournamentsSwissAPIGamesBySwissRequest {
	r.moves = &moves
	return r
}

// Include the full PGN within the JSON response, in a &#x60;pgn&#x60; field.
func (r TournamentsSwissAPIGamesBySwissRequest) PgnInJson(pgnInJson bool) TournamentsSwissAPIGamesBySwissRequest {
	r.pgnInJson = &pgnInJson
	return r
}

// Include the PGN tags.
func (r TournamentsSwissAPIGamesBySwissRequest) Tags(tags bool) TournamentsSwissAPIGamesBySwissRequest {
	r.tags = &tags
	return r
}

// Include clock status when available. Either as PGN comments: &#x60;2. exd5 { [%clk 1:01:27] } e5 { [%clk 1:01:28] }&#x60; Or in a &#x60;clocks&#x60; JSON field, as centisecond integers, depending on the response type. 
func (r TournamentsSwissAPIGamesBySwissRequest) Clocks(clocks bool) TournamentsSwissAPIGamesBySwissRequest {
	r.clocks = &clocks
	return r
}

// Include analysis evaluations and comments, when available. Either as PGN comments: &#x60;12. Bxf6 { [%eval 0.23] } a3 { [%eval -1.09] }&#x60; Or in an &#x60;analysis&#x60; JSON field, depending on the response type. 
func (r TournamentsSwissAPIGamesBySwissRequest) Evals(evals bool) TournamentsSwissAPIGamesBySwissRequest {
	r.evals = &evals
	return r
}

// Include [accuracy percent](https://lichess.org/page/accuracy) of each player, when available. Only available in JSON. 
func (r TournamentsSwissAPIGamesBySwissRequest) Accuracy(accuracy bool) TournamentsSwissAPIGamesBySwissRequest {
	r.accuracy = &accuracy
	return r
}

// Include the opening name. Example: &#x60;[Opening \&quot;King&#39;s Gambit Accepted, King&#39;s Knight Gambit\&quot;]&#x60; 
func (r TournamentsSwissAPIGamesBySwissRequest) Opening(opening bool) TournamentsSwissAPIGamesBySwissRequest {
	r.opening = &opening
	return r
}

// Plies which mark the beginning of the middlegame and endgame. Only available in JSON 
func (r TournamentsSwissAPIGamesBySwissRequest) Division(division bool) TournamentsSwissAPIGamesBySwissRequest {
	r.division = &division
	return r
}

func (r TournamentsSwissAPIGamesBySwissRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.GamesBySwissExecute(r)
}

/*
GamesBySwiss Export games of a Swiss tournament

Download games of a swiss tournament in PGN or [ndjson](#description/streaming-with-nd-json) format.
Games are sorted by chronological order.
The game stream is throttled, depending on who is making the request:
  - Anonymous request: 20 games per second
  - [OAuth2 authenticated](#description/authentication) request: 30 games per second


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIGamesBySwissRequest
*/
func (a *TournamentsSwissAPIService) GamesBySwiss(ctx context.Context, id string) TournamentsSwissAPIGamesBySwissRequest {
	return TournamentsSwissAPIGamesBySwissRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return string
func (a *TournamentsSwissAPIService) GamesBySwissExecute(r TournamentsSwissAPIGamesBySwissRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.GamesBySwiss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/games"
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

type TournamentsSwissAPIResultsBySwissRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
	nb *int32
}

// Max number of players to fetch
func (r TournamentsSwissAPIResultsBySwissRequest) Nb(nb int32) TournamentsSwissAPIResultsBySwissRequest {
	r.nb = &nb
	return r
}

func (r TournamentsSwissAPIResultsBySwissRequest) Execute() (*ResultsBySwiss200Response, *http.Response, error) {
	return r.ApiService.ResultsBySwissExecute(r)
}

/*
ResultsBySwiss Get results of a swiss tournament

Players of a swiss tournament, with their score and performance, sorted by rank (best first).
Players are streamed as [ndjson](#description/streaming-with-nd-json).
If called on an ongoing tournament, results can be inconsistent
due to ranking changes while the players are being streamed.
Use on finished tournaments for guaranteed consistency.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPIResultsBySwissRequest
*/
func (a *TournamentsSwissAPIService) ResultsBySwiss(ctx context.Context, id string) TournamentsSwissAPIResultsBySwissRequest {
	return TournamentsSwissAPIResultsBySwissRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return ResultsBySwiss200Response
func (a *TournamentsSwissAPIService) ResultsBySwissExecute(r TournamentsSwissAPIResultsBySwissRequest) (*ResultsBySwiss200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ResultsBySwiss200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.ResultsBySwiss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}/results"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

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

type TournamentsSwissAPISwissRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
}

func (r TournamentsSwissAPISwissRequest) Execute() (*SwissTournament, *http.Response, error) {
	return r.ApiService.SwissExecute(r)
}

/*
Swiss Get info about a Swiss tournament

Get detailed info about a Swiss tournament.


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The Swiss tournament ID.
 @return TournamentsSwissAPISwissRequest
*/
func (a *TournamentsSwissAPIService) Swiss(ctx context.Context, id string) TournamentsSwissAPISwissRequest {
	return TournamentsSwissAPISwissRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return SwissTournament
func (a *TournamentsSwissAPIService) SwissExecute(r TournamentsSwissAPISwissRequest) (*SwissTournament, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *SwissTournament
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.Swiss")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/swiss/{id}"
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

type TournamentsSwissAPISwissTrfRequest struct {
	ctx context.Context
	ApiService TournamentsSwissAPI
	id string
}

func (r TournamentsSwissAPISwissTrfRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SwissTrfExecute(r)
}

/*
SwissTrf Export TRF of a Swiss tournament

Download a tournament in the Tournament Report File format, the FIDE standard.
Documentation: <https://www.fide.com/FIDE/handbook/C04Annex2_TRF16.pdf>
Example: <https://lichess.org/swiss/j8rtJ5GL.trf>


 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The tournament ID.
 @return TournamentsSwissAPISwissTrfRequest
*/
func (a *TournamentsSwissAPIService) SwissTrf(ctx context.Context, id string) TournamentsSwissAPISwissTrfRequest {
	return TournamentsSwissAPISwissTrfRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return string
func (a *TournamentsSwissAPIService) SwissTrfExecute(r TournamentsSwissAPISwissTrfRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "TournamentsSwissAPIService.SwissTrf")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/swiss/{id}.trf"
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
	localVarHTTPHeaderAccepts := []string{"text/plain"}

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
