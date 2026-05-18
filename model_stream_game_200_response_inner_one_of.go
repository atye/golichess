/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lichess

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the StreamGame200ResponseInnerOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamGame200ResponseInnerOneOf{}

// StreamGame200ResponseInnerOneOf struct for StreamGame200ResponseInnerOneOf
type StreamGame200ResponseInnerOneOf struct {
	Id string `json:"id"`
	Variant *ApiAccountPlaying200ResponseNowPlayingInnerVariant `json:"variant,omitempty"`
	Speed *string `json:"speed,omitempty"`
	Perf *string `json:"perf,omitempty"`
	Rated *bool `json:"rated,omitempty"`
	InitialFen *string `json:"initialFen,omitempty"`
	Fen *string `json:"fen,omitempty"`
	Player NullableString `json:"player,omitempty"`
	Turns *int32 `json:"turns,omitempty"`
	StartedAtTurn *int32 `json:"startedAtTurn,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *StreamGame200ResponseInnerOneOfStatus `json:"status,omitempty"`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	Players *GamePgn200ResponseOneOfPlayers `json:"players,omitempty"`
}

type _StreamGame200ResponseInnerOneOf StreamGame200ResponseInnerOneOf

// NewStreamGame200ResponseInnerOneOf instantiates a new StreamGame200ResponseInnerOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamGame200ResponseInnerOneOf(id string) *StreamGame200ResponseInnerOneOf {
	this := StreamGame200ResponseInnerOneOf{}
	this.Id = id
	return &this
}

// NewStreamGame200ResponseInnerOneOfWithDefaults instantiates a new StreamGame200ResponseInnerOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamGame200ResponseInnerOneOfWithDefaults() *StreamGame200ResponseInnerOneOf {
	this := StreamGame200ResponseInnerOneOf{}
	return &this
}

// GetId returns the Id field value
func (o *StreamGame200ResponseInnerOneOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *StreamGame200ResponseInnerOneOf) SetId(v string) {
	o.Id = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant {
	if o == nil || IsNil(o.Variant) {
		var ret ApiAccountPlaying200ResponseNowPlayingInnerVariant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given ApiAccountPlaying200ResponseNowPlayingInnerVariant and assigns it to the Variant field.
func (o *StreamGame200ResponseInnerOneOf) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant) {
	o.Variant = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *StreamGame200ResponseInnerOneOf) SetSpeed(v string) {
	o.Speed = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetPerf() string {
	if o == nil || IsNil(o.Perf) {
		var ret string
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetPerfOk() (*string, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given string and assigns it to the Perf field.
func (o *StreamGame200ResponseInnerOneOf) SetPerf(v string) {
	o.Perf = &v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *StreamGame200ResponseInnerOneOf) SetRated(v bool) {
	o.Rated = &v
}

// GetInitialFen returns the InitialFen field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetInitialFen() string {
	if o == nil || IsNil(o.InitialFen) {
		var ret string
		return ret
	}
	return *o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetInitialFenOk() (*string, bool) {
	if o == nil || IsNil(o.InitialFen) {
		return nil, false
	}
	return o.InitialFen, true
}

// HasInitialFen returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasInitialFen() bool {
	if o != nil && !IsNil(o.InitialFen) {
		return true
	}

	return false
}

// SetInitialFen gets a reference to the given string and assigns it to the InitialFen field.
func (o *StreamGame200ResponseInnerOneOf) SetInitialFen(v string) {
	o.InitialFen = &v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *StreamGame200ResponseInnerOneOf) SetFen(v string) {
	o.Fen = &v
}

// GetPlayer returns the Player field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StreamGame200ResponseInnerOneOf) GetPlayer() string {
	if o == nil || IsNil(o.Player.Get()) {
		var ret string
		return ret
	}
	return *o.Player.Get()
}

// GetPlayerOk returns a tuple with the Player field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StreamGame200ResponseInnerOneOf) GetPlayerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Player.Get(), o.Player.IsSet()
}

// HasPlayer returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasPlayer() bool {
	if o != nil && o.Player.IsSet() {
		return true
	}

	return false
}

// SetPlayer gets a reference to the given NullableString and assigns it to the Player field.
func (o *StreamGame200ResponseInnerOneOf) SetPlayer(v string) {
	o.Player.Set(&v)
}
// SetPlayerNil sets the value for Player to be an explicit nil
func (o *StreamGame200ResponseInnerOneOf) SetPlayerNil() {
	o.Player.Set(nil)
}

// UnsetPlayer ensures that no value is present for Player, not even an explicit nil
func (o *StreamGame200ResponseInnerOneOf) UnsetPlayer() {
	o.Player.Unset()
}

// GetTurns returns the Turns field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetTurns() int32 {
	if o == nil || IsNil(o.Turns) {
		var ret int32
		return ret
	}
	return *o.Turns
}

// GetTurnsOk returns a tuple with the Turns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetTurnsOk() (*int32, bool) {
	if o == nil || IsNil(o.Turns) {
		return nil, false
	}
	return o.Turns, true
}

// HasTurns returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasTurns() bool {
	if o != nil && !IsNil(o.Turns) {
		return true
	}

	return false
}

// SetTurns gets a reference to the given int32 and assigns it to the Turns field.
func (o *StreamGame200ResponseInnerOneOf) SetTurns(v int32) {
	o.Turns = &v
}

// GetStartedAtTurn returns the StartedAtTurn field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetStartedAtTurn() int32 {
	if o == nil || IsNil(o.StartedAtTurn) {
		var ret int32
		return ret
	}
	return *o.StartedAtTurn
}

// GetStartedAtTurnOk returns a tuple with the StartedAtTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetStartedAtTurnOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedAtTurn) {
		return nil, false
	}
	return o.StartedAtTurn, true
}

// HasStartedAtTurn returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasStartedAtTurn() bool {
	if o != nil && !IsNil(o.StartedAtTurn) {
		return true
	}

	return false
}

// SetStartedAtTurn gets a reference to the given int32 and assigns it to the StartedAtTurn field.
func (o *StreamGame200ResponseInnerOneOf) SetStartedAtTurn(v int32) {
	o.StartedAtTurn = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *StreamGame200ResponseInnerOneOf) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetStatus() StreamGame200ResponseInnerOneOfStatus {
	if o == nil || IsNil(o.Status) {
		var ret StreamGame200ResponseInnerOneOfStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StreamGame200ResponseInnerOneOfStatus and assigns it to the Status field.
func (o *StreamGame200ResponseInnerOneOf) SetStatus(v StreamGame200ResponseInnerOneOfStatus) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *StreamGame200ResponseInnerOneOf) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *StreamGame200ResponseInnerOneOf) SetLastMove(v string) {
	o.LastMove = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *StreamGame200ResponseInnerOneOf) GetPlayers() GamePgn200ResponseOneOfPlayers {
	if o == nil || IsNil(o.Players) {
		var ret GamePgn200ResponseOneOfPlayers
		return ret
	}
	return *o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamGame200ResponseInnerOneOf) GetPlayersOk() (*GamePgn200ResponseOneOfPlayers, bool) {
	if o == nil || IsNil(o.Players) {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *StreamGame200ResponseInnerOneOf) HasPlayers() bool {
	if o != nil && !IsNil(o.Players) {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given GamePgn200ResponseOneOfPlayers and assigns it to the Players field.
func (o *StreamGame200ResponseInnerOneOf) SetPlayers(v GamePgn200ResponseOneOfPlayers) {
	o.Players = &v
}

func (o StreamGame200ResponseInnerOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamGame200ResponseInnerOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Perf) {
		toSerialize["perf"] = o.Perf
	}
	if !IsNil(o.Rated) {
		toSerialize["rated"] = o.Rated
	}
	if !IsNil(o.InitialFen) {
		toSerialize["initialFen"] = o.InitialFen
	}
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if o.Player.IsSet() {
		toSerialize["player"] = o.Player.Get()
	}
	if !IsNil(o.Turns) {
		toSerialize["turns"] = o.Turns
	}
	if !IsNil(o.StartedAtTurn) {
		toSerialize["startedAtTurn"] = o.StartedAtTurn
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.LastMove) {
		toSerialize["lastMove"] = o.LastMove
	}
	if !IsNil(o.Players) {
		toSerialize["players"] = o.Players
	}
	return toSerialize, nil
}

func (o *StreamGame200ResponseInnerOneOf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varStreamGame200ResponseInnerOneOf := _StreamGame200ResponseInnerOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStreamGame200ResponseInnerOneOf)

	if err != nil {
		return err
	}

	*o = StreamGame200ResponseInnerOneOf(varStreamGame200ResponseInnerOneOf)

	return err
}

type NullableStreamGame200ResponseInnerOneOf struct {
	value *StreamGame200ResponseInnerOneOf
	isSet bool
}

func (v NullableStreamGame200ResponseInnerOneOf) Get() *StreamGame200ResponseInnerOneOf {
	return v.value
}

func (v *NullableStreamGame200ResponseInnerOneOf) Set(val *StreamGame200ResponseInnerOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamGame200ResponseInnerOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamGame200ResponseInnerOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamGame200ResponseInnerOneOf(val *StreamGame200ResponseInnerOneOf) *NullableStreamGame200ResponseInnerOneOf {
	return &NullableStreamGame200ResponseInnerOneOf{value: val, isSet: true}
}

func (v NullableStreamGame200ResponseInnerOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamGame200ResponseInnerOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


