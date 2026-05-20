/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-generator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiTournament200ResponseCreatedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTournament200ResponseCreatedInner{}

// ApiTournament200ResponseCreatedInner struct for ApiTournament200ResponseCreatedInner
type ApiTournament200ResponseCreatedInner struct {
	Id string `json:"id"`
	CreatedBy string `json:"createdBy"`
	System string `json:"system"`
	Minutes int32 `json:"minutes"`
	Clock ApiTournament200ResponseCreatedInnerClock `json:"clock"`
	Rated bool `json:"rated"`
	FullName string `json:"fullName"`
	NbPlayers int32 `json:"nbPlayers"`
	Variant ApiAccountPlaying200ResponseNowPlayingInnerVariant `json:"variant"`
	StartsAt int64 `json:"startsAt"`
	FinishesAt int64 `json:"finishesAt"`
	// 10: created, 20: started, 30: finished 
	Status int32 `json:"status"`
	Perf ApiTournament200ResponseCreatedInnerPerf `json:"perf"`
	SecondsToStart *int32 `json:"secondsToStart,omitempty"`
	HasMaxRating *bool `json:"hasMaxRating,omitempty"`
	MaxRating *ApiTournament200ResponseCreatedInnerMaxRating `json:"maxRating,omitempty"`
	MinRating *ApiTournament200ResponseCreatedInnerMaxRating `json:"minRating,omitempty"`
	MinRatedGames *ApiTournament200ResponseCreatedInnerMinRatedGames `json:"minRatedGames,omitempty"`
	BotsAllowed *bool `json:"botsAllowed,omitempty"`
	MinAccountAgeInDays *int32 `json:"minAccountAgeInDays,omitempty"`
	OnlyTitled *bool `json:"onlyTitled,omitempty"`
	TeamMember *string `json:"teamMember,omitempty"`
	Private *bool `json:"private,omitempty"`
	Position *ApiTournament200ResponseCreatedInnerPosition `json:"position,omitempty"`
	Schedule *ApiTournament200ResponseCreatedInnerSchedule `json:"schedule,omitempty"`
	TeamBattle *ApiTournament200ResponseCreatedInnerTeamBattle `json:"teamBattle,omitempty"`
	Winner *ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId `json:"winner,omitempty"`
}

type _ApiTournament200ResponseCreatedInner ApiTournament200ResponseCreatedInner

// NewApiTournament200ResponseCreatedInner instantiates a new ApiTournament200ResponseCreatedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTournament200ResponseCreatedInner(id string, createdBy string, system string, minutes int32, clock ApiTournament200ResponseCreatedInnerClock, rated bool, fullName string, nbPlayers int32, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, startsAt int64, finishesAt int64, status int32, perf ApiTournament200ResponseCreatedInnerPerf) *ApiTournament200ResponseCreatedInner {
	this := ApiTournament200ResponseCreatedInner{}
	this.Id = id
	this.CreatedBy = createdBy
	this.System = system
	this.Minutes = minutes
	this.Clock = clock
	this.Rated = rated
	this.FullName = fullName
	this.NbPlayers = nbPlayers
	this.Variant = variant
	this.StartsAt = startsAt
	this.FinishesAt = finishesAt
	this.Status = status
	this.Perf = perf
	return &this
}

// NewApiTournament200ResponseCreatedInnerWithDefaults instantiates a new ApiTournament200ResponseCreatedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTournament200ResponseCreatedInnerWithDefaults() *ApiTournament200ResponseCreatedInner {
	this := ApiTournament200ResponseCreatedInner{}
	return &this
}

// GetId returns the Id field value
func (o *ApiTournament200ResponseCreatedInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiTournament200ResponseCreatedInner) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ApiTournament200ResponseCreatedInner) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ApiTournament200ResponseCreatedInner) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetSystem returns the System field value
func (o *ApiTournament200ResponseCreatedInner) GetSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *ApiTournament200ResponseCreatedInner) SetSystem(v string) {
	o.System = v
}

// GetMinutes returns the Minutes field value
func (o *ApiTournament200ResponseCreatedInner) GetMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *ApiTournament200ResponseCreatedInner) SetMinutes(v int32) {
	o.Minutes = v
}

// GetClock returns the Clock field value
func (o *ApiTournament200ResponseCreatedInner) GetClock() ApiTournament200ResponseCreatedInnerClock {
	if o == nil {
		var ret ApiTournament200ResponseCreatedInnerClock
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetClockOk() (*ApiTournament200ResponseCreatedInnerClock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *ApiTournament200ResponseCreatedInner) SetClock(v ApiTournament200ResponseCreatedInnerClock) {
	o.Clock = v
}

// GetRated returns the Rated field value
func (o *ApiTournament200ResponseCreatedInner) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *ApiTournament200ResponseCreatedInner) SetRated(v bool) {
	o.Rated = v
}

// GetFullName returns the FullName field value
func (o *ApiTournament200ResponseCreatedInner) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *ApiTournament200ResponseCreatedInner) SetFullName(v string) {
	o.FullName = v
}

// GetNbPlayers returns the NbPlayers field value
func (o *ApiTournament200ResponseCreatedInner) GetNbPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPlayers
}

// GetNbPlayersOk returns a tuple with the NbPlayers field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetNbPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPlayers, true
}

// SetNbPlayers sets field value
func (o *ApiTournament200ResponseCreatedInner) SetNbPlayers(v int32) {
	o.NbPlayers = v
}

// GetVariant returns the Variant field value
func (o *ApiTournament200ResponseCreatedInner) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant {
	if o == nil {
		var ret ApiAccountPlaying200ResponseNowPlayingInnerVariant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ApiTournament200ResponseCreatedInner) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant) {
	o.Variant = v
}

// GetStartsAt returns the StartsAt field value
func (o *ApiTournament200ResponseCreatedInner) GetStartsAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetStartsAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *ApiTournament200ResponseCreatedInner) SetStartsAt(v int64) {
	o.StartsAt = v
}

// GetFinishesAt returns the FinishesAt field value
func (o *ApiTournament200ResponseCreatedInner) GetFinishesAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinishesAt
}

// GetFinishesAtOk returns a tuple with the FinishesAt field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetFinishesAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishesAt, true
}

// SetFinishesAt sets field value
func (o *ApiTournament200ResponseCreatedInner) SetFinishesAt(v int64) {
	o.FinishesAt = v
}

// GetStatus returns the Status field value
func (o *ApiTournament200ResponseCreatedInner) GetStatus() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ApiTournament200ResponseCreatedInner) SetStatus(v int32) {
	o.Status = v
}

// GetPerf returns the Perf field value
func (o *ApiTournament200ResponseCreatedInner) GetPerf() ApiTournament200ResponseCreatedInnerPerf {
	if o == nil {
		var ret ApiTournament200ResponseCreatedInnerPerf
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetPerfOk() (*ApiTournament200ResponseCreatedInnerPerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *ApiTournament200ResponseCreatedInner) SetPerf(v ApiTournament200ResponseCreatedInnerPerf) {
	o.Perf = v
}

// GetSecondsToStart returns the SecondsToStart field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetSecondsToStart() int32 {
	if o == nil || IsNil(o.SecondsToStart) {
		var ret int32
		return ret
	}
	return *o.SecondsToStart
}

// GetSecondsToStartOk returns a tuple with the SecondsToStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetSecondsToStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsToStart) {
		return nil, false
	}
	return o.SecondsToStart, true
}

// HasSecondsToStart returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasSecondsToStart() bool {
	if o != nil && !IsNil(o.SecondsToStart) {
		return true
	}

	return false
}

// SetSecondsToStart gets a reference to the given int32 and assigns it to the SecondsToStart field.
func (o *ApiTournament200ResponseCreatedInner) SetSecondsToStart(v int32) {
	o.SecondsToStart = &v
}

// GetHasMaxRating returns the HasMaxRating field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetHasMaxRating() bool {
	if o == nil || IsNil(o.HasMaxRating) {
		var ret bool
		return ret
	}
	return *o.HasMaxRating
}

// GetHasMaxRatingOk returns a tuple with the HasMaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetHasMaxRatingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMaxRating) {
		return nil, false
	}
	return o.HasMaxRating, true
}

// HasHasMaxRating returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasHasMaxRating() bool {
	if o != nil && !IsNil(o.HasMaxRating) {
		return true
	}

	return false
}

// SetHasMaxRating gets a reference to the given bool and assigns it to the HasMaxRating field.
func (o *ApiTournament200ResponseCreatedInner) SetHasMaxRating(v bool) {
	o.HasMaxRating = &v
}

// GetMaxRating returns the MaxRating field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetMaxRating() ApiTournament200ResponseCreatedInnerMaxRating {
	if o == nil || IsNil(o.MaxRating) {
		var ret ApiTournament200ResponseCreatedInnerMaxRating
		return ret
	}
	return *o.MaxRating
}

// GetMaxRatingOk returns a tuple with the MaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetMaxRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool) {
	if o == nil || IsNil(o.MaxRating) {
		return nil, false
	}
	return o.MaxRating, true
}

// HasMaxRating returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasMaxRating() bool {
	if o != nil && !IsNil(o.MaxRating) {
		return true
	}

	return false
}

// SetMaxRating gets a reference to the given ApiTournament200ResponseCreatedInnerMaxRating and assigns it to the MaxRating field.
func (o *ApiTournament200ResponseCreatedInner) SetMaxRating(v ApiTournament200ResponseCreatedInnerMaxRating) {
	o.MaxRating = &v
}

// GetMinRating returns the MinRating field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetMinRating() ApiTournament200ResponseCreatedInnerMaxRating {
	if o == nil || IsNil(o.MinRating) {
		var ret ApiTournament200ResponseCreatedInnerMaxRating
		return ret
	}
	return *o.MinRating
}

// GetMinRatingOk returns a tuple with the MinRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetMinRatingOk() (*ApiTournament200ResponseCreatedInnerMaxRating, bool) {
	if o == nil || IsNil(o.MinRating) {
		return nil, false
	}
	return o.MinRating, true
}

// HasMinRating returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasMinRating() bool {
	if o != nil && !IsNil(o.MinRating) {
		return true
	}

	return false
}

// SetMinRating gets a reference to the given ApiTournament200ResponseCreatedInnerMaxRating and assigns it to the MinRating field.
func (o *ApiTournament200ResponseCreatedInner) SetMinRating(v ApiTournament200ResponseCreatedInnerMaxRating) {
	o.MinRating = &v
}

// GetMinRatedGames returns the MinRatedGames field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetMinRatedGames() ApiTournament200ResponseCreatedInnerMinRatedGames {
	if o == nil || IsNil(o.MinRatedGames) {
		var ret ApiTournament200ResponseCreatedInnerMinRatedGames
		return ret
	}
	return *o.MinRatedGames
}

// GetMinRatedGamesOk returns a tuple with the MinRatedGames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetMinRatedGamesOk() (*ApiTournament200ResponseCreatedInnerMinRatedGames, bool) {
	if o == nil || IsNil(o.MinRatedGames) {
		return nil, false
	}
	return o.MinRatedGames, true
}

// HasMinRatedGames returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasMinRatedGames() bool {
	if o != nil && !IsNil(o.MinRatedGames) {
		return true
	}

	return false
}

// SetMinRatedGames gets a reference to the given ApiTournament200ResponseCreatedInnerMinRatedGames and assigns it to the MinRatedGames field.
func (o *ApiTournament200ResponseCreatedInner) SetMinRatedGames(v ApiTournament200ResponseCreatedInnerMinRatedGames) {
	o.MinRatedGames = &v
}

// GetBotsAllowed returns the BotsAllowed field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetBotsAllowed() bool {
	if o == nil || IsNil(o.BotsAllowed) {
		var ret bool
		return ret
	}
	return *o.BotsAllowed
}

// GetBotsAllowedOk returns a tuple with the BotsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetBotsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BotsAllowed) {
		return nil, false
	}
	return o.BotsAllowed, true
}

// HasBotsAllowed returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasBotsAllowed() bool {
	if o != nil && !IsNil(o.BotsAllowed) {
		return true
	}

	return false
}

// SetBotsAllowed gets a reference to the given bool and assigns it to the BotsAllowed field.
func (o *ApiTournament200ResponseCreatedInner) SetBotsAllowed(v bool) {
	o.BotsAllowed = &v
}

// GetMinAccountAgeInDays returns the MinAccountAgeInDays field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetMinAccountAgeInDays() int32 {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		var ret int32
		return ret
	}
	return *o.MinAccountAgeInDays
}

// GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetMinAccountAgeInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		return nil, false
	}
	return o.MinAccountAgeInDays, true
}

// HasMinAccountAgeInDays returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasMinAccountAgeInDays() bool {
	if o != nil && !IsNil(o.MinAccountAgeInDays) {
		return true
	}

	return false
}

// SetMinAccountAgeInDays gets a reference to the given int32 and assigns it to the MinAccountAgeInDays field.
func (o *ApiTournament200ResponseCreatedInner) SetMinAccountAgeInDays(v int32) {
	o.MinAccountAgeInDays = &v
}

// GetOnlyTitled returns the OnlyTitled field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetOnlyTitled() bool {
	if o == nil || IsNil(o.OnlyTitled) {
		var ret bool
		return ret
	}
	return *o.OnlyTitled
}

// GetOnlyTitledOk returns a tuple with the OnlyTitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetOnlyTitledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyTitled) {
		return nil, false
	}
	return o.OnlyTitled, true
}

// HasOnlyTitled returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasOnlyTitled() bool {
	if o != nil && !IsNil(o.OnlyTitled) {
		return true
	}

	return false
}

// SetOnlyTitled gets a reference to the given bool and assigns it to the OnlyTitled field.
func (o *ApiTournament200ResponseCreatedInner) SetOnlyTitled(v bool) {
	o.OnlyTitled = &v
}

// GetTeamMember returns the TeamMember field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetTeamMember() string {
	if o == nil || IsNil(o.TeamMember) {
		var ret string
		return ret
	}
	return *o.TeamMember
}

// GetTeamMemberOk returns a tuple with the TeamMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetTeamMemberOk() (*string, bool) {
	if o == nil || IsNil(o.TeamMember) {
		return nil, false
	}
	return o.TeamMember, true
}

// HasTeamMember returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasTeamMember() bool {
	if o != nil && !IsNil(o.TeamMember) {
		return true
	}

	return false
}

// SetTeamMember gets a reference to the given string and assigns it to the TeamMember field.
func (o *ApiTournament200ResponseCreatedInner) SetTeamMember(v string) {
	o.TeamMember = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ApiTournament200ResponseCreatedInner) SetPrivate(v bool) {
	o.Private = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetPosition() ApiTournament200ResponseCreatedInnerPosition {
	if o == nil || IsNil(o.Position) {
		var ret ApiTournament200ResponseCreatedInnerPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetPositionOk() (*ApiTournament200ResponseCreatedInnerPosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given ApiTournament200ResponseCreatedInnerPosition and assigns it to the Position field.
func (o *ApiTournament200ResponseCreatedInner) SetPosition(v ApiTournament200ResponseCreatedInnerPosition) {
	o.Position = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetSchedule() ApiTournament200ResponseCreatedInnerSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret ApiTournament200ResponseCreatedInnerSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetScheduleOk() (*ApiTournament200ResponseCreatedInnerSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ApiTournament200ResponseCreatedInnerSchedule and assigns it to the Schedule field.
func (o *ApiTournament200ResponseCreatedInner) SetSchedule(v ApiTournament200ResponseCreatedInnerSchedule) {
	o.Schedule = &v
}

// GetTeamBattle returns the TeamBattle field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetTeamBattle() ApiTournament200ResponseCreatedInnerTeamBattle {
	if o == nil || IsNil(o.TeamBattle) {
		var ret ApiTournament200ResponseCreatedInnerTeamBattle
		return ret
	}
	return *o.TeamBattle
}

// GetTeamBattleOk returns a tuple with the TeamBattle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetTeamBattleOk() (*ApiTournament200ResponseCreatedInnerTeamBattle, bool) {
	if o == nil || IsNil(o.TeamBattle) {
		return nil, false
	}
	return o.TeamBattle, true
}

// HasTeamBattle returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasTeamBattle() bool {
	if o != nil && !IsNil(o.TeamBattle) {
		return true
	}

	return false
}

// SetTeamBattle gets a reference to the given ApiTournament200ResponseCreatedInnerTeamBattle and assigns it to the TeamBattle field.
func (o *ApiTournament200ResponseCreatedInner) SetTeamBattle(v ApiTournament200ResponseCreatedInnerTeamBattle) {
	o.TeamBattle = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *ApiTournament200ResponseCreatedInner) GetWinner() ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId {
	if o == nil || IsNil(o.Winner) {
		var ret ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournament200ResponseCreatedInner) GetWinnerOk() (*ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *ApiTournament200ResponseCreatedInner) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId and assigns it to the Winner field.
func (o *ApiTournament200ResponseCreatedInner) SetWinner(v ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId) {
	o.Winner = &v
}

func (o ApiTournament200ResponseCreatedInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTournament200ResponseCreatedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["system"] = o.System
	toSerialize["minutes"] = o.Minutes
	toSerialize["clock"] = o.Clock
	toSerialize["rated"] = o.Rated
	toSerialize["fullName"] = o.FullName
	toSerialize["nbPlayers"] = o.NbPlayers
	toSerialize["variant"] = o.Variant
	toSerialize["startsAt"] = o.StartsAt
	toSerialize["finishesAt"] = o.FinishesAt
	toSerialize["status"] = o.Status
	toSerialize["perf"] = o.Perf
	if !IsNil(o.SecondsToStart) {
		toSerialize["secondsToStart"] = o.SecondsToStart
	}
	if !IsNil(o.HasMaxRating) {
		toSerialize["hasMaxRating"] = o.HasMaxRating
	}
	if !IsNil(o.MaxRating) {
		toSerialize["maxRating"] = o.MaxRating
	}
	if !IsNil(o.MinRating) {
		toSerialize["minRating"] = o.MinRating
	}
	if !IsNil(o.MinRatedGames) {
		toSerialize["minRatedGames"] = o.MinRatedGames
	}
	if !IsNil(o.BotsAllowed) {
		toSerialize["botsAllowed"] = o.BotsAllowed
	}
	if !IsNil(o.MinAccountAgeInDays) {
		toSerialize["minAccountAgeInDays"] = o.MinAccountAgeInDays
	}
	if !IsNil(o.OnlyTitled) {
		toSerialize["onlyTitled"] = o.OnlyTitled
	}
	if !IsNil(o.TeamMember) {
		toSerialize["teamMember"] = o.TeamMember
	}
	if !IsNil(o.Private) {
		toSerialize["private"] = o.Private
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.TeamBattle) {
		toSerialize["teamBattle"] = o.TeamBattle
	}
	if !IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}
	return toSerialize, nil
}

func (o *ApiTournament200ResponseCreatedInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"system",
		"minutes",
		"clock",
		"rated",
		"fullName",
		"nbPlayers",
		"variant",
		"startsAt",
		"finishesAt",
		"status",
		"perf",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varApiTournament200ResponseCreatedInner := _ApiTournament200ResponseCreatedInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTournament200ResponseCreatedInner)

	if err != nil {
		return err
	}

	*o = ApiTournament200ResponseCreatedInner(varApiTournament200ResponseCreatedInner)

	return err
}

type NullableApiTournament200ResponseCreatedInner struct {
	value *ApiTournament200ResponseCreatedInner
	isSet bool
}

func (v NullableApiTournament200ResponseCreatedInner) Get() *ApiTournament200ResponseCreatedInner {
	return v.value
}

func (v *NullableApiTournament200ResponseCreatedInner) Set(val *ApiTournament200ResponseCreatedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTournament200ResponseCreatedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTournament200ResponseCreatedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTournament200ResponseCreatedInner(val *ApiTournament200ResponseCreatedInner) *NullableApiTournament200ResponseCreatedInner {
	return &NullableApiTournament200ResponseCreatedInner{value: val, isSet: true}
}

func (v NullableApiTournament200ResponseCreatedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTournament200ResponseCreatedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


