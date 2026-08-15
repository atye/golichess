/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.163
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ArenaTournamentFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArenaTournamentFull{}

// ArenaTournamentFull struct for ArenaTournamentFull
type ArenaTournamentFull struct {
	Id string `json:"id"`
	FullName string `json:"fullName"`
	Rated *bool `json:"rated,omitempty"`
	Spotlight *ArenaTournamentFullSpotlight `json:"spotlight,omitempty"`
	Berserkable *bool `json:"berserkable,omitempty"`
	OnlyTitled *bool `json:"onlyTitled,omitempty"`
	Clock Clock `json:"clock"`
	Minutes *int32 `json:"minutes,omitempty"`
	CreatedBy *string `json:"createdBy,omitempty"`
	System *string `json:"system,omitempty"`
	SecondsToStart *int32 `json:"secondsToStart,omitempty"`
	SecondsToFinish *int32 `json:"secondsToFinish,omitempty"`
	IsFinished *bool `json:"isFinished,omitempty"`
	IsRecentlyFinished *bool `json:"isRecentlyFinished,omitempty"`
	PairingsClosed *bool `json:"pairingsClosed,omitempty"`
	StartsAt *string `json:"startsAt,omitempty"`
	NbPlayers int32 `json:"nbPlayers"`
	Verdicts *Verdicts `json:"verdicts,omitempty"`
	Quote *ArenaTournamentFullQuote `json:"quote,omitempty"`
	GreatPlayer *ArenaTournamentFullGreatPlayer `json:"greatPlayer,omitempty"`
	// List of usernames allowed to join the tournament
	AllowList []string `json:"allowList,omitempty"`
	HasMaxRating *bool `json:"hasMaxRating,omitempty"`
	MaxRating *ArenaRatingObj `json:"maxRating,omitempty"`
	MinRating *ArenaRatingObj `json:"minRating,omitempty"`
	MinRatedGames *ArenaTournamentMinRatedGames `json:"minRatedGames,omitempty"`
	BotsAllowed *bool `json:"botsAllowed,omitempty"`
	MinAccountAgeInDays *int32 `json:"minAccountAgeInDays,omitempty"`
	Perf *ArenaTournamentFullPerf `json:"perf,omitempty"`
	Schedule *ArenaTournamentFullSchedule `json:"schedule,omitempty"`
	Description *string `json:"description,omitempty"`
	Variant *string `json:"variant,omitempty"`
	Duels []ArenaTournamentFullDuelsInner `json:"duels,omitempty"`
	Standing *ArenaTournamentFullStanding `json:"standing,omitempty"`
	Featured *ArenaTournamentFullFeatured `json:"featured,omitempty"`
	Podium []ArenaTournamentFullPodiumInner `json:"podium,omitempty"`
	Stats *ArenaTournamentFullStats `json:"stats,omitempty"`
	MyUsername *string `json:"myUsername,omitempty"`
}

type _ArenaTournamentFull ArenaTournamentFull

// NewArenaTournamentFull instantiates a new ArenaTournamentFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArenaTournamentFull(id string, fullName string, clock Clock, nbPlayers int32) *ArenaTournamentFull {
	this := ArenaTournamentFull{}
	this.Id = id
	this.FullName = fullName
	this.Clock = clock
	this.NbPlayers = nbPlayers
	return &this
}

// NewArenaTournamentFullWithDefaults instantiates a new ArenaTournamentFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArenaTournamentFullWithDefaults() *ArenaTournamentFull {
	this := ArenaTournamentFull{}
	return &this
}

// GetId returns the Id field value
func (o *ArenaTournamentFull) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ArenaTournamentFull) SetId(v string) {
	o.Id = v
}

// GetFullName returns the FullName field value
func (o *ArenaTournamentFull) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *ArenaTournamentFull) SetFullName(v string) {
	o.FullName = v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *ArenaTournamentFull) SetRated(v bool) {
	o.Rated = &v
}

// GetSpotlight returns the Spotlight field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetSpotlight() ArenaTournamentFullSpotlight {
	if o == nil || IsNil(o.Spotlight) {
		var ret ArenaTournamentFullSpotlight
		return ret
	}
	return *o.Spotlight
}

// GetSpotlightOk returns a tuple with the Spotlight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetSpotlightOk() (*ArenaTournamentFullSpotlight, bool) {
	if o == nil || IsNil(o.Spotlight) {
		return nil, false
	}
	return o.Spotlight, true
}

// HasSpotlight returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasSpotlight() bool {
	if o != nil && !IsNil(o.Spotlight) {
		return true
	}

	return false
}

// SetSpotlight gets a reference to the given ArenaTournamentFullSpotlight and assigns it to the Spotlight field.
func (o *ArenaTournamentFull) SetSpotlight(v ArenaTournamentFullSpotlight) {
	o.Spotlight = &v
}

// GetBerserkable returns the Berserkable field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetBerserkable() bool {
	if o == nil || IsNil(o.Berserkable) {
		var ret bool
		return ret
	}
	return *o.Berserkable
}

// GetBerserkableOk returns a tuple with the Berserkable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetBerserkableOk() (*bool, bool) {
	if o == nil || IsNil(o.Berserkable) {
		return nil, false
	}
	return o.Berserkable, true
}

// HasBerserkable returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasBerserkable() bool {
	if o != nil && !IsNil(o.Berserkable) {
		return true
	}

	return false
}

// SetBerserkable gets a reference to the given bool and assigns it to the Berserkable field.
func (o *ArenaTournamentFull) SetBerserkable(v bool) {
	o.Berserkable = &v
}

// GetOnlyTitled returns the OnlyTitled field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetOnlyTitled() bool {
	if o == nil || IsNil(o.OnlyTitled) {
		var ret bool
		return ret
	}
	return *o.OnlyTitled
}

// GetOnlyTitledOk returns a tuple with the OnlyTitled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetOnlyTitledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnlyTitled) {
		return nil, false
	}
	return o.OnlyTitled, true
}

// HasOnlyTitled returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasOnlyTitled() bool {
	if o != nil && !IsNil(o.OnlyTitled) {
		return true
	}

	return false
}

// SetOnlyTitled gets a reference to the given bool and assigns it to the OnlyTitled field.
func (o *ArenaTournamentFull) SetOnlyTitled(v bool) {
	o.OnlyTitled = &v
}

// GetClock returns the Clock field value
func (o *ArenaTournamentFull) GetClock() Clock {
	if o == nil {
		var ret Clock
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetClockOk() (*Clock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *ArenaTournamentFull) SetClock(v Clock) {
	o.Clock = v
}

// GetMinutes returns the Minutes field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMinutes() int32 {
	if o == nil || IsNil(o.Minutes) {
		var ret int32
		return ret
	}
	return *o.Minutes
}

// GetMinutesOk returns a tuple with the Minutes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMinutesOk() (*int32, bool) {
	if o == nil || IsNil(o.Minutes) {
		return nil, false
	}
	return o.Minutes, true
}

// HasMinutes returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMinutes() bool {
	if o != nil && !IsNil(o.Minutes) {
		return true
	}

	return false
}

// SetMinutes gets a reference to the given int32 and assigns it to the Minutes field.
func (o *ArenaTournamentFull) SetMinutes(v int32) {
	o.Minutes = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *ArenaTournamentFull) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetSystem() string {
	if o == nil || IsNil(o.System) {
		var ret string
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetSystemOk() (*string, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given string and assigns it to the System field.
func (o *ArenaTournamentFull) SetSystem(v string) {
	o.System = &v
}

// GetSecondsToStart returns the SecondsToStart field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetSecondsToStart() int32 {
	if o == nil || IsNil(o.SecondsToStart) {
		var ret int32
		return ret
	}
	return *o.SecondsToStart
}

// GetSecondsToStartOk returns a tuple with the SecondsToStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetSecondsToStartOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsToStart) {
		return nil, false
	}
	return o.SecondsToStart, true
}

// HasSecondsToStart returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasSecondsToStart() bool {
	if o != nil && !IsNil(o.SecondsToStart) {
		return true
	}

	return false
}

// SetSecondsToStart gets a reference to the given int32 and assigns it to the SecondsToStart field.
func (o *ArenaTournamentFull) SetSecondsToStart(v int32) {
	o.SecondsToStart = &v
}

// GetSecondsToFinish returns the SecondsToFinish field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetSecondsToFinish() int32 {
	if o == nil || IsNil(o.SecondsToFinish) {
		var ret int32
		return ret
	}
	return *o.SecondsToFinish
}

// GetSecondsToFinishOk returns a tuple with the SecondsToFinish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetSecondsToFinishOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsToFinish) {
		return nil, false
	}
	return o.SecondsToFinish, true
}

// HasSecondsToFinish returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasSecondsToFinish() bool {
	if o != nil && !IsNil(o.SecondsToFinish) {
		return true
	}

	return false
}

// SetSecondsToFinish gets a reference to the given int32 and assigns it to the SecondsToFinish field.
func (o *ArenaTournamentFull) SetSecondsToFinish(v int32) {
	o.SecondsToFinish = &v
}

// GetIsFinished returns the IsFinished field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetIsFinished() bool {
	if o == nil || IsNil(o.IsFinished) {
		var ret bool
		return ret
	}
	return *o.IsFinished
}

// GetIsFinishedOk returns a tuple with the IsFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetIsFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsFinished) {
		return nil, false
	}
	return o.IsFinished, true
}

// HasIsFinished returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasIsFinished() bool {
	if o != nil && !IsNil(o.IsFinished) {
		return true
	}

	return false
}

// SetIsFinished gets a reference to the given bool and assigns it to the IsFinished field.
func (o *ArenaTournamentFull) SetIsFinished(v bool) {
	o.IsFinished = &v
}

// GetIsRecentlyFinished returns the IsRecentlyFinished field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetIsRecentlyFinished() bool {
	if o == nil || IsNil(o.IsRecentlyFinished) {
		var ret bool
		return ret
	}
	return *o.IsRecentlyFinished
}

// GetIsRecentlyFinishedOk returns a tuple with the IsRecentlyFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetIsRecentlyFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsRecentlyFinished) {
		return nil, false
	}
	return o.IsRecentlyFinished, true
}

// HasIsRecentlyFinished returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasIsRecentlyFinished() bool {
	if o != nil && !IsNil(o.IsRecentlyFinished) {
		return true
	}

	return false
}

// SetIsRecentlyFinished gets a reference to the given bool and assigns it to the IsRecentlyFinished field.
func (o *ArenaTournamentFull) SetIsRecentlyFinished(v bool) {
	o.IsRecentlyFinished = &v
}

// GetPairingsClosed returns the PairingsClosed field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetPairingsClosed() bool {
	if o == nil || IsNil(o.PairingsClosed) {
		var ret bool
		return ret
	}
	return *o.PairingsClosed
}

// GetPairingsClosedOk returns a tuple with the PairingsClosed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetPairingsClosedOk() (*bool, bool) {
	if o == nil || IsNil(o.PairingsClosed) {
		return nil, false
	}
	return o.PairingsClosed, true
}

// HasPairingsClosed returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasPairingsClosed() bool {
	if o != nil && !IsNil(o.PairingsClosed) {
		return true
	}

	return false
}

// SetPairingsClosed gets a reference to the given bool and assigns it to the PairingsClosed field.
func (o *ArenaTournamentFull) SetPairingsClosed(v bool) {
	o.PairingsClosed = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetStartsAt() string {
	if o == nil || IsNil(o.StartsAt) {
		var ret string
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetStartsAtOk() (*string, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given string and assigns it to the StartsAt field.
func (o *ArenaTournamentFull) SetStartsAt(v string) {
	o.StartsAt = &v
}

// GetNbPlayers returns the NbPlayers field value
func (o *ArenaTournamentFull) GetNbPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPlayers
}

// GetNbPlayersOk returns a tuple with the NbPlayers field value
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetNbPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPlayers, true
}

// SetNbPlayers sets field value
func (o *ArenaTournamentFull) SetNbPlayers(v int32) {
	o.NbPlayers = v
}

// GetVerdicts returns the Verdicts field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetVerdicts() Verdicts {
	if o == nil || IsNil(o.Verdicts) {
		var ret Verdicts
		return ret
	}
	return *o.Verdicts
}

// GetVerdictsOk returns a tuple with the Verdicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetVerdictsOk() (*Verdicts, bool) {
	if o == nil || IsNil(o.Verdicts) {
		return nil, false
	}
	return o.Verdicts, true
}

// HasVerdicts returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasVerdicts() bool {
	if o != nil && !IsNil(o.Verdicts) {
		return true
	}

	return false
}

// SetVerdicts gets a reference to the given Verdicts and assigns it to the Verdicts field.
func (o *ArenaTournamentFull) SetVerdicts(v Verdicts) {
	o.Verdicts = &v
}

// GetQuote returns the Quote field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetQuote() ArenaTournamentFullQuote {
	if o == nil || IsNil(o.Quote) {
		var ret ArenaTournamentFullQuote
		return ret
	}
	return *o.Quote
}

// GetQuoteOk returns a tuple with the Quote field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetQuoteOk() (*ArenaTournamentFullQuote, bool) {
	if o == nil || IsNil(o.Quote) {
		return nil, false
	}
	return o.Quote, true
}

// HasQuote returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasQuote() bool {
	if o != nil && !IsNil(o.Quote) {
		return true
	}

	return false
}

// SetQuote gets a reference to the given ArenaTournamentFullQuote and assigns it to the Quote field.
func (o *ArenaTournamentFull) SetQuote(v ArenaTournamentFullQuote) {
	o.Quote = &v
}

// GetGreatPlayer returns the GreatPlayer field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetGreatPlayer() ArenaTournamentFullGreatPlayer {
	if o == nil || IsNil(o.GreatPlayer) {
		var ret ArenaTournamentFullGreatPlayer
		return ret
	}
	return *o.GreatPlayer
}

// GetGreatPlayerOk returns a tuple with the GreatPlayer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetGreatPlayerOk() (*ArenaTournamentFullGreatPlayer, bool) {
	if o == nil || IsNil(o.GreatPlayer) {
		return nil, false
	}
	return o.GreatPlayer, true
}

// HasGreatPlayer returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasGreatPlayer() bool {
	if o != nil && !IsNil(o.GreatPlayer) {
		return true
	}

	return false
}

// SetGreatPlayer gets a reference to the given ArenaTournamentFullGreatPlayer and assigns it to the GreatPlayer field.
func (o *ArenaTournamentFull) SetGreatPlayer(v ArenaTournamentFullGreatPlayer) {
	o.GreatPlayer = &v
}

// GetAllowList returns the AllowList field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetAllowList() []string {
	if o == nil || IsNil(o.AllowList) {
		var ret []string
		return ret
	}
	return o.AllowList
}

// GetAllowListOk returns a tuple with the AllowList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetAllowListOk() ([]string, bool) {
	if o == nil || IsNil(o.AllowList) {
		return nil, false
	}
	return o.AllowList, true
}

// HasAllowList returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasAllowList() bool {
	if o != nil && !IsNil(o.AllowList) {
		return true
	}

	return false
}

// SetAllowList gets a reference to the given []string and assigns it to the AllowList field.
func (o *ArenaTournamentFull) SetAllowList(v []string) {
	o.AllowList = v
}

// GetHasMaxRating returns the HasMaxRating field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetHasMaxRating() bool {
	if o == nil || IsNil(o.HasMaxRating) {
		var ret bool
		return ret
	}
	return *o.HasMaxRating
}

// GetHasMaxRatingOk returns a tuple with the HasMaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetHasMaxRatingOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMaxRating) {
		return nil, false
	}
	return o.HasMaxRating, true
}

// HasHasMaxRating returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasHasMaxRating() bool {
	if o != nil && !IsNil(o.HasMaxRating) {
		return true
	}

	return false
}

// SetHasMaxRating gets a reference to the given bool and assigns it to the HasMaxRating field.
func (o *ArenaTournamentFull) SetHasMaxRating(v bool) {
	o.HasMaxRating = &v
}

// GetMaxRating returns the MaxRating field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMaxRating() ArenaRatingObj {
	if o == nil || IsNil(o.MaxRating) {
		var ret ArenaRatingObj
		return ret
	}
	return *o.MaxRating
}

// GetMaxRatingOk returns a tuple with the MaxRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMaxRatingOk() (*ArenaRatingObj, bool) {
	if o == nil || IsNil(o.MaxRating) {
		return nil, false
	}
	return o.MaxRating, true
}

// HasMaxRating returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMaxRating() bool {
	if o != nil && !IsNil(o.MaxRating) {
		return true
	}

	return false
}

// SetMaxRating gets a reference to the given ArenaRatingObj and assigns it to the MaxRating field.
func (o *ArenaTournamentFull) SetMaxRating(v ArenaRatingObj) {
	o.MaxRating = &v
}

// GetMinRating returns the MinRating field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMinRating() ArenaRatingObj {
	if o == nil || IsNil(o.MinRating) {
		var ret ArenaRatingObj
		return ret
	}
	return *o.MinRating
}

// GetMinRatingOk returns a tuple with the MinRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMinRatingOk() (*ArenaRatingObj, bool) {
	if o == nil || IsNil(o.MinRating) {
		return nil, false
	}
	return o.MinRating, true
}

// HasMinRating returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMinRating() bool {
	if o != nil && !IsNil(o.MinRating) {
		return true
	}

	return false
}

// SetMinRating gets a reference to the given ArenaRatingObj and assigns it to the MinRating field.
func (o *ArenaTournamentFull) SetMinRating(v ArenaRatingObj) {
	o.MinRating = &v
}

// GetMinRatedGames returns the MinRatedGames field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMinRatedGames() ArenaTournamentMinRatedGames {
	if o == nil || IsNil(o.MinRatedGames) {
		var ret ArenaTournamentMinRatedGames
		return ret
	}
	return *o.MinRatedGames
}

// GetMinRatedGamesOk returns a tuple with the MinRatedGames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMinRatedGamesOk() (*ArenaTournamentMinRatedGames, bool) {
	if o == nil || IsNil(o.MinRatedGames) {
		return nil, false
	}
	return o.MinRatedGames, true
}

// HasMinRatedGames returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMinRatedGames() bool {
	if o != nil && !IsNil(o.MinRatedGames) {
		return true
	}

	return false
}

// SetMinRatedGames gets a reference to the given ArenaTournamentMinRatedGames and assigns it to the MinRatedGames field.
func (o *ArenaTournamentFull) SetMinRatedGames(v ArenaTournamentMinRatedGames) {
	o.MinRatedGames = &v
}

// GetBotsAllowed returns the BotsAllowed field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetBotsAllowed() bool {
	if o == nil || IsNil(o.BotsAllowed) {
		var ret bool
		return ret
	}
	return *o.BotsAllowed
}

// GetBotsAllowedOk returns a tuple with the BotsAllowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetBotsAllowedOk() (*bool, bool) {
	if o == nil || IsNil(o.BotsAllowed) {
		return nil, false
	}
	return o.BotsAllowed, true
}

// HasBotsAllowed returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasBotsAllowed() bool {
	if o != nil && !IsNil(o.BotsAllowed) {
		return true
	}

	return false
}

// SetBotsAllowed gets a reference to the given bool and assigns it to the BotsAllowed field.
func (o *ArenaTournamentFull) SetBotsAllowed(v bool) {
	o.BotsAllowed = &v
}

// GetMinAccountAgeInDays returns the MinAccountAgeInDays field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMinAccountAgeInDays() int32 {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		var ret int32
		return ret
	}
	return *o.MinAccountAgeInDays
}

// GetMinAccountAgeInDaysOk returns a tuple with the MinAccountAgeInDays field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMinAccountAgeInDaysOk() (*int32, bool) {
	if o == nil || IsNil(o.MinAccountAgeInDays) {
		return nil, false
	}
	return o.MinAccountAgeInDays, true
}

// HasMinAccountAgeInDays returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMinAccountAgeInDays() bool {
	if o != nil && !IsNil(o.MinAccountAgeInDays) {
		return true
	}

	return false
}

// SetMinAccountAgeInDays gets a reference to the given int32 and assigns it to the MinAccountAgeInDays field.
func (o *ArenaTournamentFull) SetMinAccountAgeInDays(v int32) {
	o.MinAccountAgeInDays = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetPerf() ArenaTournamentFullPerf {
	if o == nil || IsNil(o.Perf) {
		var ret ArenaTournamentFullPerf
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetPerfOk() (*ArenaTournamentFullPerf, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given ArenaTournamentFullPerf and assigns it to the Perf field.
func (o *ArenaTournamentFull) SetPerf(v ArenaTournamentFullPerf) {
	o.Perf = &v
}

// GetSchedule returns the Schedule field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetSchedule() ArenaTournamentFullSchedule {
	if o == nil || IsNil(o.Schedule) {
		var ret ArenaTournamentFullSchedule
		return ret
	}
	return *o.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetScheduleOk() (*ArenaTournamentFullSchedule, bool) {
	if o == nil || IsNil(o.Schedule) {
		return nil, false
	}
	return o.Schedule, true
}

// HasSchedule returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasSchedule() bool {
	if o != nil && !IsNil(o.Schedule) {
		return true
	}

	return false
}

// SetSchedule gets a reference to the given ArenaTournamentFullSchedule and assigns it to the Schedule field.
func (o *ArenaTournamentFull) SetSchedule(v ArenaTournamentFullSchedule) {
	o.Schedule = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ArenaTournamentFull) SetDescription(v string) {
	o.Description = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetVariant() string {
	if o == nil || IsNil(o.Variant) {
		var ret string
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetVariantOk() (*string, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given string and assigns it to the Variant field.
func (o *ArenaTournamentFull) SetVariant(v string) {
	o.Variant = &v
}

// GetDuels returns the Duels field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetDuels() []ArenaTournamentFullDuelsInner {
	if o == nil || IsNil(o.Duels) {
		var ret []ArenaTournamentFullDuelsInner
		return ret
	}
	return o.Duels
}

// GetDuelsOk returns a tuple with the Duels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetDuelsOk() ([]ArenaTournamentFullDuelsInner, bool) {
	if o == nil || IsNil(o.Duels) {
		return nil, false
	}
	return o.Duels, true
}

// HasDuels returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasDuels() bool {
	if o != nil && !IsNil(o.Duels) {
		return true
	}

	return false
}

// SetDuels gets a reference to the given []ArenaTournamentFullDuelsInner and assigns it to the Duels field.
func (o *ArenaTournamentFull) SetDuels(v []ArenaTournamentFullDuelsInner) {
	o.Duels = v
}

// GetStanding returns the Standing field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetStanding() ArenaTournamentFullStanding {
	if o == nil || IsNil(o.Standing) {
		var ret ArenaTournamentFullStanding
		return ret
	}
	return *o.Standing
}

// GetStandingOk returns a tuple with the Standing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetStandingOk() (*ArenaTournamentFullStanding, bool) {
	if o == nil || IsNil(o.Standing) {
		return nil, false
	}
	return o.Standing, true
}

// HasStanding returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasStanding() bool {
	if o != nil && !IsNil(o.Standing) {
		return true
	}

	return false
}

// SetStanding gets a reference to the given ArenaTournamentFullStanding and assigns it to the Standing field.
func (o *ArenaTournamentFull) SetStanding(v ArenaTournamentFullStanding) {
	o.Standing = &v
}

// GetFeatured returns the Featured field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetFeatured() ArenaTournamentFullFeatured {
	if o == nil || IsNil(o.Featured) {
		var ret ArenaTournamentFullFeatured
		return ret
	}
	return *o.Featured
}

// GetFeaturedOk returns a tuple with the Featured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetFeaturedOk() (*ArenaTournamentFullFeatured, bool) {
	if o == nil || IsNil(o.Featured) {
		return nil, false
	}
	return o.Featured, true
}

// HasFeatured returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasFeatured() bool {
	if o != nil && !IsNil(o.Featured) {
		return true
	}

	return false
}

// SetFeatured gets a reference to the given ArenaTournamentFullFeatured and assigns it to the Featured field.
func (o *ArenaTournamentFull) SetFeatured(v ArenaTournamentFullFeatured) {
	o.Featured = &v
}

// GetPodium returns the Podium field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetPodium() []ArenaTournamentFullPodiumInner {
	if o == nil || IsNil(o.Podium) {
		var ret []ArenaTournamentFullPodiumInner
		return ret
	}
	return o.Podium
}

// GetPodiumOk returns a tuple with the Podium field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetPodiumOk() ([]ArenaTournamentFullPodiumInner, bool) {
	if o == nil || IsNil(o.Podium) {
		return nil, false
	}
	return o.Podium, true
}

// HasPodium returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasPodium() bool {
	if o != nil && !IsNil(o.Podium) {
		return true
	}

	return false
}

// SetPodium gets a reference to the given []ArenaTournamentFullPodiumInner and assigns it to the Podium field.
func (o *ArenaTournamentFull) SetPodium(v []ArenaTournamentFullPodiumInner) {
	o.Podium = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetStats() ArenaTournamentFullStats {
	if o == nil || IsNil(o.Stats) {
		var ret ArenaTournamentFullStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetStatsOk() (*ArenaTournamentFullStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ArenaTournamentFullStats and assigns it to the Stats field.
func (o *ArenaTournamentFull) SetStats(v ArenaTournamentFullStats) {
	o.Stats = &v
}

// GetMyUsername returns the MyUsername field value if set, zero value otherwise.
func (o *ArenaTournamentFull) GetMyUsername() string {
	if o == nil || IsNil(o.MyUsername) {
		var ret string
		return ret
	}
	return *o.MyUsername
}

// GetMyUsernameOk returns a tuple with the MyUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArenaTournamentFull) GetMyUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.MyUsername) {
		return nil, false
	}
	return o.MyUsername, true
}

// HasMyUsername returns a boolean if a field has been set.
func (o *ArenaTournamentFull) HasMyUsername() bool {
	if o != nil && !IsNil(o.MyUsername) {
		return true
	}

	return false
}

// SetMyUsername gets a reference to the given string and assigns it to the MyUsername field.
func (o *ArenaTournamentFull) SetMyUsername(v string) {
	o.MyUsername = &v
}

func (o ArenaTournamentFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArenaTournamentFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["fullName"] = o.FullName
	if !IsNil(o.Rated) {
		toSerialize["rated"] = o.Rated
	}
	if !IsNil(o.Spotlight) {
		toSerialize["spotlight"] = o.Spotlight
	}
	if !IsNil(o.Berserkable) {
		toSerialize["berserkable"] = o.Berserkable
	}
	if !IsNil(o.OnlyTitled) {
		toSerialize["onlyTitled"] = o.OnlyTitled
	}
	toSerialize["clock"] = o.Clock
	if !IsNil(o.Minutes) {
		toSerialize["minutes"] = o.Minutes
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.SecondsToStart) {
		toSerialize["secondsToStart"] = o.SecondsToStart
	}
	if !IsNil(o.SecondsToFinish) {
		toSerialize["secondsToFinish"] = o.SecondsToFinish
	}
	if !IsNil(o.IsFinished) {
		toSerialize["isFinished"] = o.IsFinished
	}
	if !IsNil(o.IsRecentlyFinished) {
		toSerialize["isRecentlyFinished"] = o.IsRecentlyFinished
	}
	if !IsNil(o.PairingsClosed) {
		toSerialize["pairingsClosed"] = o.PairingsClosed
	}
	if !IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	toSerialize["nbPlayers"] = o.NbPlayers
	if !IsNil(o.Verdicts) {
		toSerialize["verdicts"] = o.Verdicts
	}
	if !IsNil(o.Quote) {
		toSerialize["quote"] = o.Quote
	}
	if !IsNil(o.GreatPlayer) {
		toSerialize["greatPlayer"] = o.GreatPlayer
	}
	if !IsNil(o.AllowList) {
		toSerialize["allowList"] = o.AllowList
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
	if !IsNil(o.Perf) {
		toSerialize["perf"] = o.Perf
	}
	if !IsNil(o.Schedule) {
		toSerialize["schedule"] = o.Schedule
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Duels) {
		toSerialize["duels"] = o.Duels
	}
	if !IsNil(o.Standing) {
		toSerialize["standing"] = o.Standing
	}
	if !IsNil(o.Featured) {
		toSerialize["featured"] = o.Featured
	}
	if !IsNil(o.Podium) {
		toSerialize["podium"] = o.Podium
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.MyUsername) {
		toSerialize["myUsername"] = o.MyUsername
	}
	return toSerialize, nil
}

func (o *ArenaTournamentFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"fullName",
		"clock",
		"nbPlayers",
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

	varArenaTournamentFull := _ArenaTournamentFull{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varArenaTournamentFull)

	if err != nil {
		return err
	}

	*o = ArenaTournamentFull(varArenaTournamentFull)

	return err
}

type NullableArenaTournamentFull struct {
	value *ArenaTournamentFull
	isSet bool
}

func (v NullableArenaTournamentFull) Get() *ArenaTournamentFull {
	return v.value
}

func (v *NullableArenaTournamentFull) Set(val *ArenaTournamentFull) {
	v.value = val
	v.isSet = true
}

func (v NullableArenaTournamentFull) IsSet() bool {
	return v.isSet
}

func (v *NullableArenaTournamentFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArenaTournamentFull(val *ArenaTournamentFull) *NullableArenaTournamentFull {
	return &NullableArenaTournamentFull{value: val, isSet: true}
}

func (v NullableArenaTournamentFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArenaTournamentFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


