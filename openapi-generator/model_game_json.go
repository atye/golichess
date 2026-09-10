/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.171
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameJson{}

// GameJson struct for GameJson
type GameJson struct {
	Id string `json:"id"`
	Rated bool `json:"rated"`
	Variant VariantKey `json:"variant"`
	Speed Speed `json:"speed"`
	Perf string `json:"perf"`
	CreatedAt int64 `json:"createdAt"`
	LastMoveAt int64 `json:"lastMoveAt"`
	Status GameStatusName `json:"status"`
	Source *string `json:"source,omitempty"`
	Players GamePlayers `json:"players"`
	InitialFen *string `json:"initialFen,omitempty"`
	Winner *GameColor `json:"winner,omitempty"`
	Opening *GameOpening `json:"opening,omitempty"`
	Moves *string `json:"moves,omitempty"`
	Pgn *string `json:"pgn,omitempty"`
	DaysPerTurn *int32 `json:"daysPerTurn,omitempty"`
	Analysis []GameMoveAnalysis `json:"analysis,omitempty"`
	ArenaTour *GameJsonArenaTour `json:"arenaTour,omitempty"`
	SwissTour *GameJsonSwissTour `json:"swissTour,omitempty"`
	Clock *GameJsonClock `json:"clock,omitempty"`
	Clocks []int32 `json:"clocks,omitempty"`
	Division *GameJsonDivision `json:"division,omitempty"`
}

type _GameJson GameJson

// NewGameJson instantiates a new GameJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameJson(id string, rated bool, variant VariantKey, speed Speed, perf string, createdAt int64, lastMoveAt int64, status GameStatusName, players GamePlayers) *GameJson {
	this := GameJson{}
	this.Id = id
	this.Rated = rated
	this.Variant = variant
	this.Speed = speed
	this.Perf = perf
	this.CreatedAt = createdAt
	this.LastMoveAt = lastMoveAt
	this.Status = status
	this.Players = players
	return &this
}

// NewGameJsonWithDefaults instantiates a new GameJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameJsonWithDefaults() *GameJson {
	this := GameJson{}
	var variant VariantKey = VARIANTKEY_STANDARD
	this.Variant = variant
	return &this
}

// GetId returns the Id field value
func (o *GameJson) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameJson) SetId(v string) {
	o.Id = v
}

// GetRated returns the Rated field value
func (o *GameJson) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *GameJson) SetRated(v bool) {
	o.Rated = v
}

// GetVariant returns the Variant field value
func (o *GameJson) GetVariant() VariantKey {
	if o == nil {
		var ret VariantKey
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetVariantOk() (*VariantKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *GameJson) SetVariant(v VariantKey) {
	o.Variant = v
}

// GetSpeed returns the Speed field value
func (o *GameJson) GetSpeed() Speed {
	if o == nil {
		var ret Speed
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetSpeedOk() (*Speed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *GameJson) SetSpeed(v Speed) {
	o.Speed = v
}

// GetPerf returns the Perf field value
func (o *GameJson) GetPerf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetPerfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *GameJson) SetPerf(v string) {
	o.Perf = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *GameJson) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *GameJson) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetLastMoveAt returns the LastMoveAt field value
func (o *GameJson) GetLastMoveAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.LastMoveAt
}

// GetLastMoveAtOk returns a tuple with the LastMoveAt field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetLastMoveAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMoveAt, true
}

// SetLastMoveAt sets field value
func (o *GameJson) SetLastMoveAt(v int64) {
	o.LastMoveAt = v
}

// GetStatus returns the Status field value
func (o *GameJson) GetStatus() GameStatusName {
	if o == nil {
		var ret GameStatusName
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetStatusOk() (*GameStatusName, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GameJson) SetStatus(v GameStatusName) {
	o.Status = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GameJson) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GameJson) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *GameJson) SetSource(v string) {
	o.Source = &v
}

// GetPlayers returns the Players field value
func (o *GameJson) GetPlayers() GamePlayers {
	if o == nil {
		var ret GamePlayers
		return ret
	}

	return o.Players
}

// GetPlayersOk returns a tuple with the Players field value
// and a boolean to check if the value has been set.
func (o *GameJson) GetPlayersOk() (*GamePlayers, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Players, true
}

// SetPlayers sets field value
func (o *GameJson) SetPlayers(v GamePlayers) {
	o.Players = v
}

// GetInitialFen returns the InitialFen field value if set, zero value otherwise.
func (o *GameJson) GetInitialFen() string {
	if o == nil || IsNil(o.InitialFen) {
		var ret string
		return ret
	}
	return *o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetInitialFenOk() (*string, bool) {
	if o == nil || IsNil(o.InitialFen) {
		return nil, false
	}
	return o.InitialFen, true
}

// HasInitialFen returns a boolean if a field has been set.
func (o *GameJson) HasInitialFen() bool {
	if o != nil && !IsNil(o.InitialFen) {
		return true
	}

	return false
}

// SetInitialFen gets a reference to the given string and assigns it to the InitialFen field.
func (o *GameJson) SetInitialFen(v string) {
	o.InitialFen = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *GameJson) GetWinner() GameColor {
	if o == nil || IsNil(o.Winner) {
		var ret GameColor
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetWinnerOk() (*GameColor, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *GameJson) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given GameColor and assigns it to the Winner field.
func (o *GameJson) SetWinner(v GameColor) {
	o.Winner = &v
}

// GetOpening returns the Opening field value if set, zero value otherwise.
func (o *GameJson) GetOpening() GameOpening {
	if o == nil || IsNil(o.Opening) {
		var ret GameOpening
		return ret
	}
	return *o.Opening
}

// GetOpeningOk returns a tuple with the Opening field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetOpeningOk() (*GameOpening, bool) {
	if o == nil || IsNil(o.Opening) {
		return nil, false
	}
	return o.Opening, true
}

// HasOpening returns a boolean if a field has been set.
func (o *GameJson) HasOpening() bool {
	if o != nil && !IsNil(o.Opening) {
		return true
	}

	return false
}

// SetOpening gets a reference to the given GameOpening and assigns it to the Opening field.
func (o *GameJson) SetOpening(v GameOpening) {
	o.Opening = &v
}

// GetMoves returns the Moves field value if set, zero value otherwise.
func (o *GameJson) GetMoves() string {
	if o == nil || IsNil(o.Moves) {
		var ret string
		return ret
	}
	return *o.Moves
}

// GetMovesOk returns a tuple with the Moves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetMovesOk() (*string, bool) {
	if o == nil || IsNil(o.Moves) {
		return nil, false
	}
	return o.Moves, true
}

// HasMoves returns a boolean if a field has been set.
func (o *GameJson) HasMoves() bool {
	if o != nil && !IsNil(o.Moves) {
		return true
	}

	return false
}

// SetMoves gets a reference to the given string and assigns it to the Moves field.
func (o *GameJson) SetMoves(v string) {
	o.Moves = &v
}

// GetPgn returns the Pgn field value if set, zero value otherwise.
func (o *GameJson) GetPgn() string {
	if o == nil || IsNil(o.Pgn) {
		var ret string
		return ret
	}
	return *o.Pgn
}

// GetPgnOk returns a tuple with the Pgn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetPgnOk() (*string, bool) {
	if o == nil || IsNil(o.Pgn) {
		return nil, false
	}
	return o.Pgn, true
}

// HasPgn returns a boolean if a field has been set.
func (o *GameJson) HasPgn() bool {
	if o != nil && !IsNil(o.Pgn) {
		return true
	}

	return false
}

// SetPgn gets a reference to the given string and assigns it to the Pgn field.
func (o *GameJson) SetPgn(v string) {
	o.Pgn = &v
}

// GetDaysPerTurn returns the DaysPerTurn field value if set, zero value otherwise.
func (o *GameJson) GetDaysPerTurn() int32 {
	if o == nil || IsNil(o.DaysPerTurn) {
		var ret int32
		return ret
	}
	return *o.DaysPerTurn
}

// GetDaysPerTurnOk returns a tuple with the DaysPerTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetDaysPerTurnOk() (*int32, bool) {
	if o == nil || IsNil(o.DaysPerTurn) {
		return nil, false
	}
	return o.DaysPerTurn, true
}

// HasDaysPerTurn returns a boolean if a field has been set.
func (o *GameJson) HasDaysPerTurn() bool {
	if o != nil && !IsNil(o.DaysPerTurn) {
		return true
	}

	return false
}

// SetDaysPerTurn gets a reference to the given int32 and assigns it to the DaysPerTurn field.
func (o *GameJson) SetDaysPerTurn(v int32) {
	o.DaysPerTurn = &v
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *GameJson) GetAnalysis() []GameMoveAnalysis {
	if o == nil || IsNil(o.Analysis) {
		var ret []GameMoveAnalysis
		return ret
	}
	return o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetAnalysisOk() ([]GameMoveAnalysis, bool) {
	if o == nil || IsNil(o.Analysis) {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *GameJson) HasAnalysis() bool {
	if o != nil && !IsNil(o.Analysis) {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given []GameMoveAnalysis and assigns it to the Analysis field.
func (o *GameJson) SetAnalysis(v []GameMoveAnalysis) {
	o.Analysis = v
}

// GetArenaTour returns the ArenaTour field value if set, zero value otherwise.
func (o *GameJson) GetArenaTour() GameJsonArenaTour {
	if o == nil || IsNil(o.ArenaTour) {
		var ret GameJsonArenaTour
		return ret
	}
	return *o.ArenaTour
}

// GetArenaTourOk returns a tuple with the ArenaTour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetArenaTourOk() (*GameJsonArenaTour, bool) {
	if o == nil || IsNil(o.ArenaTour) {
		return nil, false
	}
	return o.ArenaTour, true
}

// HasArenaTour returns a boolean if a field has been set.
func (o *GameJson) HasArenaTour() bool {
	if o != nil && !IsNil(o.ArenaTour) {
		return true
	}

	return false
}

// SetArenaTour gets a reference to the given GameJsonArenaTour and assigns it to the ArenaTour field.
func (o *GameJson) SetArenaTour(v GameJsonArenaTour) {
	o.ArenaTour = &v
}

// GetSwissTour returns the SwissTour field value if set, zero value otherwise.
func (o *GameJson) GetSwissTour() GameJsonSwissTour {
	if o == nil || IsNil(o.SwissTour) {
		var ret GameJsonSwissTour
		return ret
	}
	return *o.SwissTour
}

// GetSwissTourOk returns a tuple with the SwissTour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetSwissTourOk() (*GameJsonSwissTour, bool) {
	if o == nil || IsNil(o.SwissTour) {
		return nil, false
	}
	return o.SwissTour, true
}

// HasSwissTour returns a boolean if a field has been set.
func (o *GameJson) HasSwissTour() bool {
	if o != nil && !IsNil(o.SwissTour) {
		return true
	}

	return false
}

// SetSwissTour gets a reference to the given GameJsonSwissTour and assigns it to the SwissTour field.
func (o *GameJson) SetSwissTour(v GameJsonSwissTour) {
	o.SwissTour = &v
}

// GetClock returns the Clock field value if set, zero value otherwise.
func (o *GameJson) GetClock() GameJsonClock {
	if o == nil || IsNil(o.Clock) {
		var ret GameJsonClock
		return ret
	}
	return *o.Clock
}

// GetClockOk returns a tuple with the Clock field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetClockOk() (*GameJsonClock, bool) {
	if o == nil || IsNil(o.Clock) {
		return nil, false
	}
	return o.Clock, true
}

// HasClock returns a boolean if a field has been set.
func (o *GameJson) HasClock() bool {
	if o != nil && !IsNil(o.Clock) {
		return true
	}

	return false
}

// SetClock gets a reference to the given GameJsonClock and assigns it to the Clock field.
func (o *GameJson) SetClock(v GameJsonClock) {
	o.Clock = &v
}

// GetClocks returns the Clocks field value if set, zero value otherwise.
func (o *GameJson) GetClocks() []int32 {
	if o == nil || IsNil(o.Clocks) {
		var ret []int32
		return ret
	}
	return o.Clocks
}

// GetClocksOk returns a tuple with the Clocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetClocksOk() ([]int32, bool) {
	if o == nil || IsNil(o.Clocks) {
		return nil, false
	}
	return o.Clocks, true
}

// HasClocks returns a boolean if a field has been set.
func (o *GameJson) HasClocks() bool {
	if o != nil && !IsNil(o.Clocks) {
		return true
	}

	return false
}

// SetClocks gets a reference to the given []int32 and assigns it to the Clocks field.
func (o *GameJson) SetClocks(v []int32) {
	o.Clocks = v
}

// GetDivision returns the Division field value if set, zero value otherwise.
func (o *GameJson) GetDivision() GameJsonDivision {
	if o == nil || IsNil(o.Division) {
		var ret GameJsonDivision
		return ret
	}
	return *o.Division
}

// GetDivisionOk returns a tuple with the Division field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameJson) GetDivisionOk() (*GameJsonDivision, bool) {
	if o == nil || IsNil(o.Division) {
		return nil, false
	}
	return o.Division, true
}

// HasDivision returns a boolean if a field has been set.
func (o *GameJson) HasDivision() bool {
	if o != nil && !IsNil(o.Division) {
		return true
	}

	return false
}

// SetDivision gets a reference to the given GameJsonDivision and assigns it to the Division field.
func (o *GameJson) SetDivision(v GameJsonDivision) {
	o.Division = &v
}

func (o GameJson) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["rated"] = o.Rated
	toSerialize["variant"] = o.Variant
	toSerialize["speed"] = o.Speed
	toSerialize["perf"] = o.Perf
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["lastMoveAt"] = o.LastMoveAt
	toSerialize["status"] = o.Status
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	toSerialize["players"] = o.Players
	if !IsNil(o.InitialFen) {
		toSerialize["initialFen"] = o.InitialFen
	}
	if !IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}
	if !IsNil(o.Opening) {
		toSerialize["opening"] = o.Opening
	}
	if !IsNil(o.Moves) {
		toSerialize["moves"] = o.Moves
	}
	if !IsNil(o.Pgn) {
		toSerialize["pgn"] = o.Pgn
	}
	if !IsNil(o.DaysPerTurn) {
		toSerialize["daysPerTurn"] = o.DaysPerTurn
	}
	if !IsNil(o.Analysis) {
		toSerialize["analysis"] = o.Analysis
	}
	if !IsNil(o.ArenaTour) {
		toSerialize["arenaTour"] = o.ArenaTour
	}
	if !IsNil(o.SwissTour) {
		toSerialize["swissTour"] = o.SwissTour
	}
	if !IsNil(o.Clock) {
		toSerialize["clock"] = o.Clock
	}
	if !IsNil(o.Clocks) {
		toSerialize["clocks"] = o.Clocks
	}
	if !IsNil(o.Division) {
		toSerialize["division"] = o.Division
	}
	return toSerialize, nil
}

func (o *GameJson) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"rated",
		"variant",
		"speed",
		"perf",
		"createdAt",
		"lastMoveAt",
		"status",
		"players",
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

	varGameJson := _GameJson{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameJson)

	if err != nil {
		return err
	}

	*o = GameJson(varGameJson)

	return err
}

type NullableGameJson struct {
	value *GameJson
	isSet bool
}

func (v NullableGameJson) Get() *GameJson {
	return v.value
}

func (v *NullableGameJson) Set(val *GameJson) {
	v.value = val
	v.isSet = true
}

func (v NullableGameJson) IsSet() bool {
	return v.isSet
}

func (v *NullableGameJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameJson(val *GameJson) *NullableGameJson {
	return &NullableGameJson{value: val, isSet: true}
}

func (v NullableGameJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


