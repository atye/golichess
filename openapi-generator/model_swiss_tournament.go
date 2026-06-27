/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.149
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SwissTournament type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwissTournament{}

// SwissTournament struct for SwissTournament
type SwissTournament struct {
	Id string `json:"id"`
	CreatedBy string `json:"createdBy"`
	StartsAt string `json:"startsAt"`
	Name string `json:"name"`
	Clock SwissTournamentClock `json:"clock"`
	Variant string `json:"variant"`
	Round float32 `json:"round"`
	NbRounds float32 `json:"nbRounds"`
	NbPlayers float32 `json:"nbPlayers"`
	NbOngoing float32 `json:"nbOngoing"`
	Status SwissStatus `json:"status"`
	Stats *SwissTournamentStats `json:"stats,omitempty"`
	Rated bool `json:"rated"`
	Verdicts Verdicts `json:"verdicts"`
	NextRound *SwissTournamentNextRound `json:"nextRound,omitempty"`
}

type _SwissTournament SwissTournament

// NewSwissTournament instantiates a new SwissTournament object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwissTournament(id string, createdBy string, startsAt string, name string, clock SwissTournamentClock, variant string, round float32, nbRounds float32, nbPlayers float32, nbOngoing float32, status SwissStatus, rated bool, verdicts Verdicts) *SwissTournament {
	this := SwissTournament{}
	this.Id = id
	this.CreatedBy = createdBy
	this.StartsAt = startsAt
	this.Name = name
	this.Clock = clock
	this.Variant = variant
	this.Round = round
	this.NbRounds = nbRounds
	this.NbPlayers = nbPlayers
	this.NbOngoing = nbOngoing
	this.Status = status
	this.Rated = rated
	this.Verdicts = verdicts
	return &this
}

// NewSwissTournamentWithDefaults instantiates a new SwissTournament object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwissTournamentWithDefaults() *SwissTournament {
	this := SwissTournament{}
	return &this
}

// GetId returns the Id field value
func (o *SwissTournament) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SwissTournament) SetId(v string) {
	o.Id = v
}

// GetCreatedBy returns the CreatedBy field value
func (o *SwissTournament) GetCreatedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetCreatedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedBy, true
}

// SetCreatedBy sets field value
func (o *SwissTournament) SetCreatedBy(v string) {
	o.CreatedBy = v
}

// GetStartsAt returns the StartsAt field value
func (o *SwissTournament) GetStartsAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetStartsAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartsAt, true
}

// SetStartsAt sets field value
func (o *SwissTournament) SetStartsAt(v string) {
	o.StartsAt = v
}

// GetName returns the Name field value
func (o *SwissTournament) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SwissTournament) SetName(v string) {
	o.Name = v
}

// GetClock returns the Clock field value
func (o *SwissTournament) GetClock() SwissTournamentClock {
	if o == nil {
		var ret SwissTournamentClock
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetClockOk() (*SwissTournamentClock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *SwissTournament) SetClock(v SwissTournamentClock) {
	o.Clock = v
}

// GetVariant returns the Variant field value
func (o *SwissTournament) GetVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *SwissTournament) SetVariant(v string) {
	o.Variant = v
}

// GetRound returns the Round field value
func (o *SwissTournament) GetRound() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetRoundOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *SwissTournament) SetRound(v float32) {
	o.Round = v
}

// GetNbRounds returns the NbRounds field value
func (o *SwissTournament) GetNbRounds() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NbRounds
}

// GetNbRoundsOk returns a tuple with the NbRounds field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetNbRoundsOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbRounds, true
}

// SetNbRounds sets field value
func (o *SwissTournament) SetNbRounds(v float32) {
	o.NbRounds = v
}

// GetNbPlayers returns the NbPlayers field value
func (o *SwissTournament) GetNbPlayers() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NbPlayers
}

// GetNbPlayersOk returns a tuple with the NbPlayers field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetNbPlayersOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPlayers, true
}

// SetNbPlayers sets field value
func (o *SwissTournament) SetNbPlayers(v float32) {
	o.NbPlayers = v
}

// GetNbOngoing returns the NbOngoing field value
func (o *SwissTournament) GetNbOngoing() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.NbOngoing
}

// GetNbOngoingOk returns a tuple with the NbOngoing field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetNbOngoingOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbOngoing, true
}

// SetNbOngoing sets field value
func (o *SwissTournament) SetNbOngoing(v float32) {
	o.NbOngoing = v
}

// GetStatus returns the Status field value
func (o *SwissTournament) GetStatus() SwissStatus {
	if o == nil {
		var ret SwissStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetStatusOk() (*SwissStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *SwissTournament) SetStatus(v SwissStatus) {
	o.Status = v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *SwissTournament) GetStats() SwissTournamentStats {
	if o == nil || IsNil(o.Stats) {
		var ret SwissTournamentStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetStatsOk() (*SwissTournamentStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// HasStats returns a boolean if a field has been set.
func (o *SwissTournament) HasStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given SwissTournamentStats and assigns it to the Stats field.
func (o *SwissTournament) SetStats(v SwissTournamentStats) {
	o.Stats = &v
}

// GetRated returns the Rated field value
func (o *SwissTournament) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *SwissTournament) SetRated(v bool) {
	o.Rated = v
}

// GetVerdicts returns the Verdicts field value
func (o *SwissTournament) GetVerdicts() Verdicts {
	if o == nil {
		var ret Verdicts
		return ret
	}

	return o.Verdicts
}

// GetVerdictsOk returns a tuple with the Verdicts field value
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetVerdictsOk() (*Verdicts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Verdicts, true
}

// SetVerdicts sets field value
func (o *SwissTournament) SetVerdicts(v Verdicts) {
	o.Verdicts = v
}

// GetNextRound returns the NextRound field value if set, zero value otherwise.
func (o *SwissTournament) GetNextRound() SwissTournamentNextRound {
	if o == nil || IsNil(o.NextRound) {
		var ret SwissTournamentNextRound
		return ret
	}
	return *o.NextRound
}

// GetNextRoundOk returns a tuple with the NextRound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwissTournament) GetNextRoundOk() (*SwissTournamentNextRound, bool) {
	if o == nil || IsNil(o.NextRound) {
		return nil, false
	}
	return o.NextRound, true
}

// HasNextRound returns a boolean if a field has been set.
func (o *SwissTournament) HasNextRound() bool {
	if o != nil && !IsNil(o.NextRound) {
		return true
	}

	return false
}

// SetNextRound gets a reference to the given SwissTournamentNextRound and assigns it to the NextRound field.
func (o *SwissTournament) SetNextRound(v SwissTournamentNextRound) {
	o.NextRound = &v
}

func (o SwissTournament) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwissTournament) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["createdBy"] = o.CreatedBy
	toSerialize["startsAt"] = o.StartsAt
	toSerialize["name"] = o.Name
	toSerialize["clock"] = o.Clock
	toSerialize["variant"] = o.Variant
	toSerialize["round"] = o.Round
	toSerialize["nbRounds"] = o.NbRounds
	toSerialize["nbPlayers"] = o.NbPlayers
	toSerialize["nbOngoing"] = o.NbOngoing
	toSerialize["status"] = o.Status
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	toSerialize["rated"] = o.Rated
	toSerialize["verdicts"] = o.Verdicts
	if !IsNil(o.NextRound) {
		toSerialize["nextRound"] = o.NextRound
	}
	return toSerialize, nil
}

func (o *SwissTournament) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"createdBy",
		"startsAt",
		"name",
		"clock",
		"variant",
		"round",
		"nbRounds",
		"nbPlayers",
		"nbOngoing",
		"status",
		"rated",
		"verdicts",
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

	varSwissTournament := _SwissTournament{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSwissTournament)

	if err != nil {
		return err
	}

	*o = SwissTournament(varSwissTournament)

	return err
}

type NullableSwissTournament struct {
	value *SwissTournament
	isSet bool
}

func (v NullableSwissTournament) Get() *SwissTournament {
	return v.value
}

func (v *NullableSwissTournament) Set(val *SwissTournament) {
	v.value = val
	v.isSet = true
}

func (v NullableSwissTournament) IsSet() bool {
	return v.isSet
}

func (v *NullableSwissTournament) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwissTournament(val *SwissTournament) *NullableSwissTournament {
	return &NullableSwissTournament{value: val, isSet: true}
}

func (v NullableSwissTournament) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwissTournament) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


