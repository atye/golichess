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

// checks if the GameFullEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameFullEvent{}

// GameFullEvent struct for GameFullEvent
type GameFullEvent struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Variant Variant `json:"variant"`
	Clock *GameFullEventClock `json:"clock,omitempty"`
	Speed Speed `json:"speed"`
	Perf GameFullEventPerf `json:"perf"`
	Rated bool `json:"rated"`
	CreatedAt int64 `json:"createdAt"`
	White GameEventPlayer `json:"white"`
	Black GameEventPlayer `json:"black"`
	InitialFen string `json:"initialFen"`
	State GameStateEvent `json:"state"`
	// If the game is correspondence
	DaysPerTurn *int32 `json:"daysPerTurn,omitempty"`
	TournamentId *string `json:"tournamentId,omitempty"`
}

type _GameFullEvent GameFullEvent

// NewGameFullEvent instantiates a new GameFullEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameFullEvent(type_ string, id string, variant Variant, speed Speed, perf GameFullEventPerf, rated bool, createdAt int64, white GameEventPlayer, black GameEventPlayer, initialFen string, state GameStateEvent) *GameFullEvent {
	this := GameFullEvent{}
	this.Type = type_
	this.Id = id
	this.Variant = variant
	this.Speed = speed
	this.Perf = perf
	this.Rated = rated
	this.CreatedAt = createdAt
	this.White = white
	this.Black = black
	this.InitialFen = initialFen
	this.State = state
	return &this
}

// NewGameFullEventWithDefaults instantiates a new GameFullEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameFullEventWithDefaults() *GameFullEvent {
	this := GameFullEvent{}
	var initialFen string = "startpos"
	this.InitialFen = initialFen
	return &this
}

// GetType returns the Type field value
func (o *GameFullEvent) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameFullEvent) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameFullEvent) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameFullEvent) SetId(v string) {
	o.Id = v
}

// GetVariant returns the Variant field value
func (o *GameFullEvent) GetVariant() Variant {
	if o == nil {
		var ret Variant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetVariantOk() (*Variant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *GameFullEvent) SetVariant(v Variant) {
	o.Variant = v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *GameFullEvent) GetClock() GameFullEventClock {
	if o == nil || IsNil(o.Clock) {
		var ret GameFullEventClock
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetClockOk() (*GameFullEventClock, bool) {
	if o == nil || IsNil(o.Clock) {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *GameFullEvent) HasClock() bool {
	if o != nil && !IsNil(o.Clock) {
		return true
	}

	return false
}

// SetClock gets a reference to the given GameFullEventClock and assigns it to the Clock field.
func (o *GameFullEvent) SetClock(v GameFullEventClock) {
	o.Clock = &v
}

// GetSpeed returns the Speed field value
func (o *GameFullEvent) GetSpeed() Speed {
	if o == nil {
		var ret Speed
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetSpeedOk() (*Speed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *GameFullEvent) SetSpeed(v Speed) {
	o.Speed = v
}

// GetPerf returns the Perf field value
func (o *GameFullEvent) GetPerf() GameFullEventPerf {
	if o == nil {
		var ret GameFullEventPerf
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetPerfOk() (*GameFullEventPerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *GameFullEvent) SetPerf(v GameFullEventPerf) {
	o.Perf = v
}

// GetRated returns the Rated field value
func (o *GameFullEvent) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *GameFullEvent) SetRated(v bool) {
	o.Rated = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GameFullEvent) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GameFullEvent) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetWhite returns the White field value
func (o *GameFullEvent) GetWhite() GameEventPlayer {
	if o == nil {
		var ret GameEventPlayer
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetWhiteOk() (*GameEventPlayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *GameFullEvent) SetWhite(v GameEventPlayer) {
	o.White = v
}

// GetBlack returns the Black field value
func (o *GameFullEvent) GetBlack() GameEventPlayer {
	if o == nil {
		var ret GameEventPlayer
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetBlackOk() (*GameEventPlayer, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *GameFullEvent) SetBlack(v GameEventPlayer) {
	o.Black = v
}

// GetInitialFen returns the InitialFen field value
func (o *GameFullEvent) GetInitialFen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetInitialFenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialFen, true
}

// SetInitialFen sets field value
func (o *GameFullEvent) SetInitialFen(v string) {
	o.InitialFen = v
}

// GetState returns the State field value
func (o *GameFullEvent) GetState() GameStateEvent {
	if o == nil {
		var ret GameStateEvent
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetStateOk() (*GameStateEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *GameFullEvent) SetState(v GameStateEvent) {
	o.State = v
}

// GetDaysPerTurn returns the DaysPerTurn field value if set, zero value otherwise.
func (o *GameFullEvent) GetDaysPerTurn() int32 {
	if o == nil || IsNil(o.DaysPerTurn) {
		var ret int32
		return ret
	}
	return *o.DaysPerTurn
}

// GetDaysPerTurnOk returns a tuple with the DaysPerTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetDaysPerTurnOk() (*int32, bool) {
	if o == nil || IsNil(o.DaysPerTurn) {
		return nil, false
	}
	return o.DaysPerTurn, true
}

// HasDaysPerTurn returns a boolean if a field has been set.
func (o *GameFullEvent) HasDaysPerTurn() bool {
	if o != nil && !IsNil(o.DaysPerTurn) {
		return true
	}

	return false
}

// SetDaysPerTurn gets a reference to the given int32 and assigns it to the DaysPerTurn field.
func (o *GameFullEvent) SetDaysPerTurn(v int32) {
	o.DaysPerTurn = &v
}

// GetTournamentId returns the TournamentId field value if set, zero value otherwise.
func (o *GameFullEvent) GetTournamentId() string {
	if o == nil || IsNil(o.TournamentId) {
		var ret string
		return ret
	}
	return *o.TournamentId
}

// GetTournamentIdOk returns a tuple with the TournamentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameFullEvent) GetTournamentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TournamentId) {
		return nil, false
	}
	return o.TournamentId, true
}

// HasTournamentId returns a boolean if a field has been set.
func (o *GameFullEvent) HasTournamentId() bool {
	if o != nil && !IsNil(o.TournamentId) {
		return true
	}

	return false
}

// SetTournamentId gets a reference to the given string and assigns it to the TournamentId field.
func (o *GameFullEvent) SetTournamentId(v string) {
	o.TournamentId = &v
}

func (o GameFullEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameFullEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	toSerialize["variant"] = o.Variant
	if !IsNil(o.Clock) {
		toSerialize["clock"] = o.Clock
	}
	toSerialize["speed"] = o.Speed
	toSerialize["perf"] = o.Perf
	toSerialize["rated"] = o.Rated
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["white"] = o.White
	toSerialize["black"] = o.Black
	toSerialize["initialFen"] = o.InitialFen
	toSerialize["state"] = o.State
	if !IsNil(o.DaysPerTurn) {
		toSerialize["daysPerTurn"] = o.DaysPerTurn
	}
	if !IsNil(o.TournamentId) {
		toSerialize["tournamentId"] = o.TournamentId
	}
	return toSerialize, nil
}

func (o *GameFullEvent) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"id",
		"variant",
		"speed",
		"perf",
		"rated",
		"createdAt",
		"white",
		"black",
		"initialFen",
		"state",
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

	varGameFullEvent := _GameFullEvent{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameFullEvent)

	if err != nil {
		return err
	}

	*o = GameFullEvent(varGameFullEvent)

	return err
}

type NullableGameFullEvent struct {
	value *GameFullEvent
	isSet bool
}

func (v NullableGameFullEvent) Get() *GameFullEvent {
	return v.value
}

func (v *NullableGameFullEvent) Set(val *GameFullEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableGameFullEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableGameFullEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameFullEvent(val *GameFullEvent) *NullableGameFullEvent {
	return &NullableGameFullEvent{value: val, isSet: true}
}

func (v NullableGameFullEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameFullEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


