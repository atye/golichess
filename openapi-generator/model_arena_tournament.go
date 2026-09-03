/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.168
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ArenaTournament type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArenaTournament{}

// ArenaTournament struct for ArenaTournament
type ArenaTournament struct {
	Id string `json:"id"`
	CreatedBy string `json:"createdBy"`
	System string `json:"system"`
	Minutes int32 `json:"minutes"`
	Clock Clock `json:"clock"`
	Rated bool `json:"rated"`
	FullName string `json:"fullName"`
	NbPlayers int32 `json:"nbPlayers"`
	Variant Variant `json:"variant"`
	StartsAt int64 `json:"startsAt"`
	FinishesAt int64 `json:"finishesAt"`
	Status ArenaStatus `json:"status"`
	Perf ArenaPerf `json:"perf"`
	SecondsToStart *int32 `json:"secondsToStart,omitempty"`
	HasMaxRating *bool `json:"hasMaxRating,omitempty"`
	MaxRating *ArenaRatingObj `json:"maxRating,omitempty"`
	MinRating *ArenaRatingObj `json:"minRating,omitempty"`
	MinRatedGames *ArenaTournamentMinRatedGames `json:"minRatedGames,omitempty"`
	BotsAllowed *bool `json:"botsAllowed,omitempty"`
	MinAccountAgeInDays *int32 `json:"minAccountAgeInDays,omitempty"`
	OnlyTitled *bool `json:"onlyTitled,omitempty"`
	TeamMember *string `json:"teamMember,omitempty"`
	Private *bool `json:"private,omitempty"`
	Position *ArenaPosition `json:"position,omitempty"`
	Schedule *ArenaTournamentSchedule `json:"schedule,omitempty"`
	TeamBattle *ArenaTournamentTeamBattle `json:"teamBattle,omitempty"`
	Winner *LightUser `json:"winner,omitempty"`
}

type _ArenaTournament ArenaTournament

// NewArenaTournament instantiates a new ArenaTournament object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArenaTournament(id string, createdBy string, system string, minutes int32, clock Clock, rated bool, fullName string, nbPlayers int32, variant Variant, startsAt int64, finishesAt int64, status ArenaStatus, perf ArenaPerf) *ArenaTournament {
	this := ArenaTournament{}
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

// NewArenaTournamentWithDefaults instantiates a new ArenaTournament object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArenaTournamentWithDefaults() *ArenaTournament {
	this := ArenaTournament{}
	return &this
}

// GetId returns the Id field value
func (o *ArenaTournament) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArenaTournament) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *ArenaTournament) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *ArenaTournament) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetSystem returns the System field value
func (o *ArenaTournament) GetSystem() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.System
}

// GetSystemOk returns a tuple with the System field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetSystemOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.System, true
}

// SetSystem sets field value
func (o *ArenaTournament) SetSystem(v string) {
	o.System = v
}

// GetMinutes returns the Minutes field value
func (o *ArenaTournament) GetMinutes() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetMinutesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Minutes, true
}

// SetMinutes sets field value
func (o *ArenaTournament) SetMinutes(v int32) {
	o.Minutes = v
}

// GetClock returns the Clock field value
func (o *ArenaTournament) GetClock() Clock {
	if o == nil {
		var ret Clock
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetClockOk() (*Clock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *ArenaTournament) SetClock(v Clock) {
	o.Clock = v
}

// GetRated returns the Rated field value
func (o *ArenaTournament) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *ArenaTournament) SetRated(v bool) {
	o.Rated = v
}

// GetFullName returns the FullName field value
func (o *ArenaTournament) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *ArenaTournament) SetFullName(v string) {
	o.FullName = v
}

// GetNbPlayers returns the NbPlayers field value
func (o *ArenaTournament) GetNbPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPlayers
}

// GetNbPlayersOk returns a tuple with the NbPlayers field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetNbPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPlayers, true
}

// SetNbPlayers sets field value
func (o *ArenaTournament) SetNbPlayers(v int32) {
	o.NbPlayers = v
}

// GetVariant returns the Variant field value
func (o *ArenaTournament) GetVariant() Variant {
	if o == nil {
		var ret Variant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetVariantOk() (*Variant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ArenaTournament) SetVariant(v Variant) {
	o.Variant = v
}

// GetStartsAt returns the StartsAt field value
func (o *ArenaTournament) GetStartsAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetStartsAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *ArenaTournament) SetStartsAt(v int64) {
	o.StartsAt = v
}

// GetFinishesAt returns the FinishesAt field value
func (o *ArenaTournament) GetFinishesAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FinishesAt
}

// GetFinishesAtOk returns a tuple with the FinishesAt field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetFinishesAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FinishesAt, true
}

// SetFinishesAt sets field value
func (o *ArenaTournament) SetFinishesAt(v int64) {
	o.FinishesAt = v
}

// GetStatus returns the Status field value
func (o *ArenaTournament) GetStatus() ArenaStatus {
	if o == nil {
		var ret ArenaStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetStatusOk() (*ArenaStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ArenaTournament) SetStatus(v ArenaStatus) {
	o.Status = v
}

// GetPerf returns the Perf field value
func (o *ArenaTournament) GetPerf() ArenaPerf {
	if o == nil {
		var ret ArenaPerf
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetPerfOk() (*ArenaPerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *ArenaTournament) SetPerf(v ArenaPerf) {
	o.Perf = v
}

// GetSecondsToStart returns the SecondsToStart field value if set, zero value otherwise.
func (o *ArenaTournament) GetSecondsToStart() int32 {
	if o == nil || IsNil(o.SecondsToStart) {
		var ret int32
		return ret
	}
	return *o.SecondsToStart
}

// GetSecondsToStartOk returns a tuple with the SecondsToStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetSecondsToStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsToStart) {
		return nil, false
	}
	return o.SecondsToStart, true
}

// HasSecondsToStart returns a boolean if a field has been set.
func (o *ArenaTournament) HasSecondsToStart() bool {
	if o != nil && !IsNil(o.SecondsToStart) {
		return true
	}

	return false
}

// SetSecondsToStart gets a reference to the given int32 and assigns it to the SecondsToStart field.
func (o *ArenaTournament) SetSecondsToStart(v int32) {
	o.SecondsToStart = &v
}

// GetHasMaxRating returns the HasMaxRating field value if set, zero value otherwise.
func (o *ArenaTournament) GetHasMaxRating() bool {
	if o == nil || IsNil(o.HasMaxRating) {
		var ret bool
		return ret
	}
	return *o.HasMaxRating
}

// GetHasMaxRatingOk returns a tuple with the HasMaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetHasMaxRatingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMaxRating) {
		return nil, false
	}
	return o.HasMaxRating, true
}

// HasHasMaxRating returns a boolean if a field has been set.
func (o *ArenaTournament) HasHasMaxRating() bool {
	if o != nil && !IsNil(o.HasMaxRating) {
		return true
	}

	return false
}

// SetHasMaxRating gets a reference to the given bool and assigns it to the HasMaxRating field.
func (o *ArenaTournament) SetHasMaxRating(v bool) {
	o.HasMaxRating = &v
}

// GetMaxRating returns the MaxRating field value if set, zero value otherwise.
func (o *ArenaTournament) GetMaxRating() ArenaRatingObj {
	if o == nil || IsNil(o.MaxRating) {
		var ret ArenaRatingObj
		return ret
	}
	return *o.MaxRating
}

// GetMaxRatingOk returns a tuple with the MaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetMaxRatingOk() (*ArenaRatingObj, bool) {
	if o == nil || IsNil(o.MaxRating) {
		return nil, false
	}
	return o.MaxRating, true
}

// HasMaxRating returns a boolean if a field has been set.
func (o *ArenaTournament) HasMaxRating() bool {
	if o != nil && !IsNil(o.MaxRating) {
		return true
	}

	return false
}

// SetMaxRating gets a reference to the given ArenaRatingObj and assigns it to the MaxRating field.
func (o *ArenaTournament) SetMaxRating(v ArenaRatingObj) {
	o.MaxRating = &v
}

// GetMinRating returns the MinRating field value if set, zero value otherwise.
func (o *ArenaTournament) GetMinRating() ArenaRatingObj {
	if o == nil || IsNil(o.MinRating) {
		var ret ArenaRatingObj
		return ret
	}
	return *o.MinRating
}

// GetMinRatingOk returns a tuple with the MinRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetMinRatingOk() (*ArenaRatingObj, bool) {
	if o == nil || IsNil(o.MinRating) {
		return nil, false
	}
	return o.MinRating, true
}

// HasMinRating returns a boolean if a field has been set.
func (o *ArenaTournament) HasMinRating() bool {
	if o != nil && !IsNil(o.MinRating) {
		return true
	}

	return false
}

// SetMinRating gets a reference to the given ArenaRatingObj and assigns it to the MinRating field.
func (o *ArenaTournament) SetMinRating(v ArenaRatingObj) {
	o.MinRating = &v
}

// GetMinRatedGames returns the MinRatedGames field value if set, zero value otherwise.
func (o *ArenaTournament) GetMinRatedGames() ArenaTournamentMinRatedGames {
	if o == nil || IsNil(o.MinRatedGames) {
		var ret ArenaTournamentMinRatedGames
		return ret
	}
	return *o.MinRatedGames
}

// GetMinRatedGamesOk returns a tuple with the MinRatedGames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetMinRatedGamesOk() (*ArenaTournamentMinRatedGames, bool) {
	if o == nil || IsNil(o.MinRatedGames) {
		return nil, false
	}
	return o.MinRatedGames, true
}

// HasMinRatedGames returns a boolean if a field has been set.
func (o *ArenaTournament) HasMinRatedGames() bool {
	if o != nil && !IsNil(o.MinRatedGames) {
		return true
	}

	return false
}

// SetMinRatedGames gets a reference to the given ArenaTournamentMinRatedGames and assigns it to the MinRatedGames field.
func (o *ArenaTournament) SetMinRatedGames(v ArenaTournamentMinRatedGames) {
	o.MinRatedGames = &v
}

// GetBotsAllowed returns the BotsAllowed field value if set, zero value otherwise.
func (o *ArenaTournament) GetBotsAllowed() bool {
	if o == nil || IsNil(o.BotsAllowed) {
		var ret bool
		return ret
	}
	return *o.BotsAllowed
}

// GetBotsAllowedOk returns a tuple with the BotsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetBotsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BotsAllowed) {
		return nil, false
	}
	return o.BotsAllowed, true
}

// HasBotsAllowed returns a boolean if a field has been set.
func (o *ArenaTournament) HasBotsAllowed() bool {
	if o != nil && !IsNil(o.BotsAllowed) {
		return true
	}

	return false
}

// SetBotsAllowed gets a reference to the given bool and assigns it to the BotsAllowed field.
func (o *ArenaTournament) SetBotsAllowed(v bool) {
	o.BotsAllowed = &v
}

// GetMinAccountAgeInDays returns the MinAccountAgeInDays field value if set, zero value otherwise.
func (o *ArenaTournament) GetMinAccountAgeInDays() int32 {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		var ret int32
		return ret
	}
	return *o.MinAccountAgeInDays
}

// GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetMinAccountAgeInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		return nil, false
	}
	return o.MinAccountAgeInDays, true
}

// HasMinAccountAgeInDays returns a boolean if a field has been set.
func (o *ArenaTournament) HasMinAccountAgeInDays() bool {
	if o != nil && !IsNil(o.MinAccountAgeInDays) {
		return true
	}

	return false
}

// SetMinAccountAgeInDays gets a reference to the given int32 and assigns it to the MinAccountAgeInDays field.
func (o *ArenaTournament) SetMinAccountAgeInDays(v int32) {
	o.MinAccountAgeInDays = &v
}

// GetOnlyTitled returns the OnlyTitled field value if set, zero value otherwise.
func (o *ArenaTournament) GetOnlyTitled() bool {
	if o == nil || IsNil(o.OnlyTitled) {
		var ret bool
		return ret
	}
	return *o.OnlyTitled
}

// GetOnlyTitledOk returns a tuple with the OnlyTitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetOnlyTitledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyTitled) {
		return nil, false
	}
	return o.OnlyTitled, true
}

// HasOnlyTitled returns a boolean if a field has been set.
func (o *ArenaTournament) HasOnlyTitled() bool {
	if o != nil && !IsNil(o.OnlyTitled) {
		return true
	}

	return false
}

// SetOnlyTitled gets a reference to the given bool and assigns it to the OnlyTitled field.
func (o *ArenaTournament) SetOnlyTitled(v bool) {
	o.OnlyTitled = &v
}

// GetTeamMember returns the TeamMember field value if set, zero value otherwise.
func (o *ArenaTournament) GetTeamMember() string {
	if o == nil || IsNil(o.TeamMember) {
		var ret string
		return ret
	}
	return *o.TeamMember
}

// GetTeamMemberOk returns a tuple with the TeamMember field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetTeamMemberOk() (*string, bool) {
	if o == nil || IsNil(o.TeamMember) {
		return nil, false
	}
	return o.TeamMember, true
}

// HasTeamMember returns a boolean if a field has been set.
func (o *ArenaTournament) HasTeamMember() bool {
	if o != nil && !IsNil(o.TeamMember) {
		return true
	}

	return false
}

// SetTeamMember gets a reference to the given string and assigns it to the TeamMember field.
func (o *ArenaTournament) SetTeamMember(v string) {
	o.TeamMember = &v
}

// GetPrivate returns the Private field value if set, zero value otherwise.
func (o *ArenaTournament) GetPrivate() bool {
	if o == nil || IsNil(o.Private) {
		var ret bool
		return ret
	}
	return *o.Private
}

// GetPrivateOk returns a tuple with the Private field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetPrivateOk() (*bool, bool) {
	if o == nil || IsNil(o.Private) {
		return nil, false
	}
	return o.Private, true
}

// HasPrivate returns a boolean if a field has been set.
func (o *ArenaTournament) HasPrivate() bool {
	if o != nil && !IsNil(o.Private) {
		return true
	}

	return false
}

// SetPrivate gets a reference to the given bool and assigns it to the Private field.
func (o *ArenaTournament) SetPrivate(v bool) {
	o.Private = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *ArenaTournament) GetPosition() ArenaPosition {
	if o == nil || IsNil(o.Position) {
		var ret ArenaPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetPositionOk() (*ArenaPosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *ArenaTournament) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given ArenaPosition and assigns it to the Position field.
func (o *ArenaTournament) SetPosition(v ArenaPosition) {
	o.Position = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ArenaTournament) GetSchedule() ArenaTournamentSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret ArenaTournamentSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetScheduleOk() (*ArenaTournamentSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ArenaTournament) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ArenaTournamentSchedule and assigns it to the Schedule field.
func (o *ArenaTournament) SetSchedule(v ArenaTournamentSchedule) {
	o.Schedule = &v
}

// GetTeamBattle returns the TeamBattle field value if set, zero value otherwise.
func (o *ArenaTournament) GetTeamBattle() ArenaTournamentTeamBattle {
	if o == nil || IsNil(o.TeamBattle) {
		var ret ArenaTournamentTeamBattle
		return ret
	}
	return *o.TeamBattle
}

// GetTeamBattleOk returns a tuple with the TeamBattle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetTeamBattleOk() (*ArenaTournamentTeamBattle, bool) {
	if o == nil || IsNil(o.TeamBattle) {
		return nil, false
	}
	return o.TeamBattle, true
}

// HasTeamBattle returns a boolean if a field has been set.
func (o *ArenaTournament) HasTeamBattle() bool {
	if o != nil && !IsNil(o.TeamBattle) {
		return true
	}

	return false
}

// SetTeamBattle gets a reference to the given ArenaTournamentTeamBattle and assigns it to the TeamBattle field.
func (o *ArenaTournament) SetTeamBattle(v ArenaTournamentTeamBattle) {
	o.TeamBattle = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *ArenaTournament) GetWinner() LightUser {
	if o == nil || IsNil(o.Winner) {
		var ret LightUser
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournament) GetWinnerOk() (*LightUser, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *ArenaTournament) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given LightUser and assigns it to the Winner field.
func (o *ArenaTournament) SetWinner(v LightUser) {
	o.Winner = &v
}

func (o ArenaTournament) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArenaTournament) ToMap() (map[string]interface{}, error) {
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

func (o *ArenaTournament) UnmarshalJSON(data []byte) (err error) {
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

	varArenaTournament := _ArenaTournament{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArenaTournament)

	if err != nil {
		return err
	}

	*o = ArenaTournament(varArenaTournament)

	return err
}

type NullableArenaTournament struct {
	value *ArenaTournament
	isSet bool
}

func (v NullableArenaTournament) Get() *ArenaTournament {
	return v.value
}

func (v *NullableArenaTournament) Set(val *ArenaTournament) {
	v.value = val
	v.isSet = true
}

func (v NullableArenaTournament) IsSet() bool {
	return v.isSet
}

func (v *NullableArenaTournament) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArenaTournament(val *ArenaTournament) *NullableArenaTournament {
	return &NullableArenaTournament{value: val, isSet: true}
}

func (v NullableArenaTournament) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArenaTournament) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


