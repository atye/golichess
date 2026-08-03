/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.158
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the MoveStreamEntryOneOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MoveStreamEntryOneOf{}

// MoveStreamEntryOneOf struct for MoveStreamEntryOneOf
type MoveStreamEntryOneOf struct {
	Id string `json:"id"`
	Variant *Variant `json:"variant,omitempty"`
	Speed *Speed `json:"speed,omitempty"`
	Perf *PerfType `json:"perf,omitempty"`
	Rated *bool `json:"rated,omitempty"`
	InitialFen *string `json:"initialFen,omitempty"`
	Fen *string `json:"fen,omitempty"`
	Player *GameColor `json:"player,omitempty"`
	Turns *int32 `json:"turns,omitempty"`
	StartedAtTurn *int32 `json:"startedAtTurn,omitempty"`
	Source *GameSource `json:"source,omitempty"`
	Status *GameStatus `json:"status,omitempty"`
	CreatedAt *int32 `json:"createdAt,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	Players *GamePlayers `json:"players,omitempty"`
}

type _MoveStreamEntryOneOf MoveStreamEntryOneOf

// NewMoveStreamEntryOneOf instantiates a new MoveStreamEntryOneOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoveStreamEntryOneOf(id string) *MoveStreamEntryOneOf {
	this := MoveStreamEntryOneOf{}
	this.Id = id
	return &this
}

// NewMoveStreamEntryOneOfWithDefaults instantiates a new MoveStreamEntryOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoveStreamEntryOneOfWithDefaults() *MoveStreamEntryOneOf {
	this := MoveStreamEntryOneOf{}
	return &this
}

// GetId returns the Id field value
func (o *MoveStreamEntryOneOf) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *MoveStreamEntryOneOf) SetId(v string) {
	o.Id = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetVariant() Variant {
	if o == nil || IsNil(o.Variant) {
		var ret Variant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetVariantOk() (*Variant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given Variant and assigns it to the Variant field.
func (o *MoveStreamEntryOneOf) SetVariant(v Variant) {
	o.Variant = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetSpeed() Speed {
	if o == nil || IsNil(o.Speed) {
		var ret Speed
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetSpeedOk() (*Speed, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given Speed and assigns it to the Speed field.
func (o *MoveStreamEntryOneOf) SetSpeed(v Speed) {
	o.Speed = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetPerf() PerfType {
	if o == nil || IsNil(o.Perf) {
		var ret PerfType
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetPerfOk() (*PerfType, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given PerfType and assigns it to the Perf field.
func (o *MoveStreamEntryOneOf) SetPerf(v PerfType) {
	o.Perf = &v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *MoveStreamEntryOneOf) SetRated(v bool) {
	o.Rated = &v
}

// GetInitialFen returns the InitialFen field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetInitialFen() string {
	if o == nil || IsNil(o.InitialFen) {
		var ret string
		return ret
	}
	return *o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetInitialFenOk() (*string, bool) {
	if o == nil || IsNil(o.InitialFen) {
		return nil, false
	}
	return o.InitialFen, true
}

// HasInitialFen returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasInitialFen() bool {
	if o != nil && !IsNil(o.InitialFen) {
		return true
	}

	return false
}

// SetInitialFen gets a reference to the given string and assigns it to the InitialFen field.
func (o *MoveStreamEntryOneOf) SetInitialFen(v string) {
	o.InitialFen = &v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *MoveStreamEntryOneOf) SetFen(v string) {
	o.Fen = &v
}

// GetPlayer returns the Player field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetPlayer() GameColor {
	if o == nil || IsNil(o.Player) {
		var ret GameColor
		return ret
	}
	return *o.Player
}

// GetPlayerOk returns a tuple with the Player field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetPlayerOk() (*GameColor, bool) {
	if o == nil || IsNil(o.Player) {
		return nil, false
	}
	return o.Player, true
}

// HasPlayer returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasPlayer() bool {
	if o != nil && !IsNil(o.Player) {
		return true
	}

	return false
}

// SetPlayer gets a reference to the given GameColor and assigns it to the Player field.
func (o *MoveStreamEntryOneOf) SetPlayer(v GameColor) {
	o.Player = &v
}

// GetTurns returns the Turns field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetTurns() int32 {
	if o == nil || IsNil(o.Turns) {
		var ret int32
		return ret
	}
	return *o.Turns
}

// GetTurnsOk returns a tuple with the Turns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetTurnsOk() (*int32, bool) {
	if o == nil || IsNil(o.Turns) {
		return nil, false
	}
	return o.Turns, true
}

// HasTurns returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasTurns() bool {
	if o != nil && !IsNil(o.Turns) {
		return true
	}

	return false
}

// SetTurns gets a reference to the given int32 and assigns it to the Turns field.
func (o *MoveStreamEntryOneOf) SetTurns(v int32) {
	o.Turns = &v
}

// GetStartedAtTurn returns the StartedAtTurn field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetStartedAtTurn() int32 {
	if o == nil || IsNil(o.StartedAtTurn) {
		var ret int32
		return ret
	}
	return *o.StartedAtTurn
}

// GetStartedAtTurnOk returns a tuple with the StartedAtTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetStartedAtTurnOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedAtTurn) {
		return nil, false
	}
	return o.StartedAtTurn, true
}

// HasStartedAtTurn returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasStartedAtTurn() bool {
	if o != nil && !IsNil(o.StartedAtTurn) {
		return true
	}

	return false
}

// SetStartedAtTurn gets a reference to the given int32 and assigns it to the StartedAtTurn field.
func (o *MoveStreamEntryOneOf) SetStartedAtTurn(v int32) {
	o.StartedAtTurn = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetSource() GameSource {
	if o == nil || IsNil(o.Source) {
		var ret GameSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetSourceOk() (*GameSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GameSource and assigns it to the Source field.
func (o *MoveStreamEntryOneOf) SetSource(v GameSource) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetStatus() GameStatus {
	if o == nil || IsNil(o.Status) {
		var ret GameStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetStatusOk() (*GameStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GameStatus and assigns it to the Status field.
func (o *MoveStreamEntryOneOf) SetStatus(v GameStatus) {
	o.Status = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetCreatedAt() int32 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int32
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetCreatedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int32 and assigns it to the CreatedAt field.
func (o *MoveStreamEntryOneOf) SetCreatedAt(v int32) {
	o.CreatedAt = &v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *MoveStreamEntryOneOf) SetLastMove(v string) {
	o.LastMove = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *MoveStreamEntryOneOf) GetPlayers() GamePlayers {
	if o == nil || IsNil(o.Players) {
		var ret GamePlayers
		return ret
	}
	return *o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoveStreamEntryOneOf) GetPlayersOk() (*GamePlayers, bool) {
	if o == nil || IsNil(o.Players) {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *MoveStreamEntryOneOf) HasPlayers() bool {
	if o != nil && !IsNil(o.Players) {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given GamePlayers and assigns it to the Players field.
func (o *MoveStreamEntryOneOf) SetPlayers(v GamePlayers) {
	o.Players = &v
}

func (o MoveStreamEntryOneOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MoveStreamEntryOneOf) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Player) {
		toSerialize["player"] = o.Player
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

func (o *MoveStreamEntryOneOf) UnmarshalJSON(data []byte) (err error) {
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

	varMoveStreamEntryOneOf := _MoveStreamEntryOneOf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMoveStreamEntryOneOf)

	if err != nil {
		return err
	}

	*o = MoveStreamEntryOneOf(varMoveStreamEntryOneOf)

	return err
}

type NullableMoveStreamEntryOneOf struct {
	value *MoveStreamEntryOneOf
	isSet bool
}

func (v NullableMoveStreamEntryOneOf) Get() *MoveStreamEntryOneOf {
	return v.value
}

func (v *NullableMoveStreamEntryOneOf) Set(val *MoveStreamEntryOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMoveStreamEntryOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMoveStreamEntryOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoveStreamEntryOneOf(val *MoveStreamEntryOneOf) *NullableMoveStreamEntryOneOf {
	return &NullableMoveStreamEntryOneOf{value: val, isSet: true}
}

func (v NullableMoveStreamEntryOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoveStreamEntryOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


