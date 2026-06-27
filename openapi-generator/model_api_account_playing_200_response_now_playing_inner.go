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

// checks if the ApiAccountPlaying200ResponseNowPlayingInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAccountPlaying200ResponseNowPlayingInner{}

// ApiAccountPlaying200ResponseNowPlayingInner struct for ApiAccountPlaying200ResponseNowPlayingInner
type ApiAccountPlaying200ResponseNowPlayingInner struct {
	FullId string `json:"fullId"`
	GameId string `json:"gameId"`
	Fen string `json:"fen"`
	Color GameColor `json:"color"`
	LastMove string `json:"lastMove"`
	Source GameSource `json:"source"`
	Status *GameStatusName `json:"status,omitempty"`
	Variant Variant `json:"variant"`
	Speed Speed `json:"speed"`
	Perf PerfType `json:"perf"`
	Rated bool `json:"rated"`
	HasMoved bool `json:"hasMoved"`
	Opponent ApiAccountPlaying200ResponseNowPlayingInnerOpponent `json:"opponent"`
	IsMyTurn bool `json:"isMyTurn"`
	SecondsLeft int32 `json:"secondsLeft"`
	TournamentId *string `json:"tournamentId,omitempty"`
	SwissId *string `json:"swissId,omitempty"`
	Winner *GameColor `json:"winner,omitempty"`
	RatingDiff *int32 `json:"ratingDiff,omitempty"`
}

type _ApiAccountPlaying200ResponseNowPlayingInner ApiAccountPlaying200ResponseNowPlayingInner

// NewApiAccountPlaying200ResponseNowPlayingInner instantiates a new ApiAccountPlaying200ResponseNowPlayingInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAccountPlaying200ResponseNowPlayingInner(fullId string, gameId string, fen string, color GameColor, lastMove string, source GameSource, variant Variant, speed Speed, perf PerfType, rated bool, hasMoved bool, opponent ApiAccountPlaying200ResponseNowPlayingInnerOpponent, isMyTurn bool, secondsLeft int32) *ApiAccountPlaying200ResponseNowPlayingInner {
	this := ApiAccountPlaying200ResponseNowPlayingInner{}
	this.FullId = fullId
	this.GameId = gameId
	this.Fen = fen
	this.Color = color
	this.LastMove = lastMove
	this.Source = source
	this.Variant = variant
	this.Speed = speed
	this.Perf = perf
	this.Rated = rated
	this.HasMoved = hasMoved
	this.Opponent = opponent
	this.IsMyTurn = isMyTurn
	this.SecondsLeft = secondsLeft
	return &this
}

// NewApiAccountPlaying200ResponseNowPlayingInnerWithDefaults instantiates a new ApiAccountPlaying200ResponseNowPlayingInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAccountPlaying200ResponseNowPlayingInnerWithDefaults() *ApiAccountPlaying200ResponseNowPlayingInner {
	this := ApiAccountPlaying200ResponseNowPlayingInner{}
	return &this
}

// GetFullId returns the FullId field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFullId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullId
}

// GetFullIdOk returns a tuple with the FullId field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFullIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullId, true
}

// SetFullId sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetFullId(v string) {
	o.FullId = v
}

// GetGameId returns the GameId field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetGameId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetGameIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameId, true
}

// SetGameId sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetGameId(v string) {
	o.GameId = v
}

// GetFen returns the Fen field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fen
}

// GetFenOk returns a tuple with the Fen field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetFenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fen, true
}

// SetFen sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetFen(v string) {
	o.Fen = v
}

// GetColor returns the Color field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetColor() GameColor {
	if o == nil {
		var ret GameColor
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetColorOk() (*GameColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetColor(v GameColor) {
	o.Color = v
}

// GetLastMove returns the LastMove field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetLastMove() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetLastMoveOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LastMove, true
}

// SetLastMove sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetLastMove(v string) {
	o.LastMove = v
}

// GetSource returns the Source field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSource() GameSource {
	if o == nil {
		var ret GameSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSourceOk() (*GameSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSource(v GameSource) {
	o.Source = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetStatus() GameStatusName {
	if o == nil || IsNil(o.Status) {
		var ret GameStatusName
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetStatusOk() (*GameStatusName, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GameStatusName and assigns it to the Status field.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetStatus(v GameStatusName) {
	o.Status = &v
}

// GetVariant returns the Variant field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetVariant() Variant {
	if o == nil {
		var ret Variant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetVariantOk() (*Variant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetVariant(v Variant) {
	o.Variant = v
}

// GetSpeed returns the Speed field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSpeed() Speed {
	if o == nil {
		var ret Speed
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSpeedOk() (*Speed, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSpeed(v Speed) {
	o.Speed = v
}

// GetPerf returns the Perf field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetPerf() PerfType {
	if o == nil {
		var ret PerfType
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetPerfOk() (*PerfType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetPerf(v PerfType) {
	o.Perf = v
}

// GetRated returns the Rated field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetRated(v bool) {
	o.Rated = v
}

// GetHasMoved returns the HasMoved field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetHasMoved() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HasMoved
}

// GetHasMovedOk returns a tuple with the HasMoved field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetHasMovedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HasMoved, true
}

// SetHasMoved sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetHasMoved(v bool) {
	o.HasMoved = v
}

// GetOpponent returns the Opponent field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetOpponent() ApiAccountPlaying200ResponseNowPlayingInnerOpponent {
	if o == nil {
		var ret ApiAccountPlaying200ResponseNowPlayingInnerOpponent
		return ret
	}

	return o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetOpponentOk() (*ApiAccountPlaying200ResponseNowPlayingInnerOpponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Opponent, true
}

// SetOpponent sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetOpponent(v ApiAccountPlaying200ResponseNowPlayingInnerOpponent) {
	o.Opponent = v
}

// GetIsMyTurn returns the IsMyTurn field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetIsMyTurn() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsMyTurn
}

// GetIsMyTurnOk returns a tuple with the IsMyTurn field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetIsMyTurnOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsMyTurn, true
}

// SetIsMyTurn sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetIsMyTurn(v bool) {
	o.IsMyTurn = v
}

// GetSecondsLeft returns the SecondsLeft field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSecondsLeft() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SecondsLeft
}

// GetSecondsLeftOk returns a tuple with the SecondsLeft field value
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSecondsLeftOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SecondsLeft, true
}

// SetSecondsLeft sets field value
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSecondsLeft(v int32) {
	o.SecondsLeft = v
}

// GetTournamentId returns the TournamentId field value if set, zero value otherwise.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetTournamentId() string {
	if o == nil || IsNil(o.TournamentId) {
		var ret string
		return ret
	}
	return *o.TournamentId
}

// GetTournamentIdOk returns a tuple with the TournamentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetTournamentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TournamentId) {
		return nil, false
	}
	return o.TournamentId, true
}

// HasTournamentId returns a boolean if a field has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasTournamentId() bool {
	if o != nil && !IsNil(o.TournamentId) {
		return true
	}

	return false
}

// SetTournamentId gets a reference to the given string and assigns it to the TournamentId field.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetTournamentId(v string) {
	o.TournamentId = &v
}

// GetSwissId returns the SwissId field value if set, zero value otherwise.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSwissId() string {
	if o == nil || IsNil(o.SwissId) {
		var ret string
		return ret
	}
	return *o.SwissId
}

// GetSwissIdOk returns a tuple with the SwissId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetSwissIdOk() (*string, bool) {
	if o == nil || IsNil(o.SwissId) {
		return nil, false
	}
	return o.SwissId, true
}

// HasSwissId returns a boolean if a field has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasSwissId() bool {
	if o != nil && !IsNil(o.SwissId) {
		return true
	}

	return false
}

// SetSwissId gets a reference to the given string and assigns it to the SwissId field.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetSwissId(v string) {
	o.SwissId = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetWinner() GameColor {
	if o == nil || IsNil(o.Winner) {
		var ret GameColor
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetWinnerOk() (*GameColor, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given GameColor and assigns it to the Winner field.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetWinner(v GameColor) {
	o.Winner = &v
}

// GetRatingDiff returns the RatingDiff field value if set, zero value otherwise.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatingDiff() int32 {
	if o == nil || IsNil(o.RatingDiff) {
		var ret int32
		return ret
	}
	return *o.RatingDiff
}

// GetRatingDiffOk returns a tuple with the RatingDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) GetRatingDiffOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingDiff) {
		return nil, false
	}
	return o.RatingDiff, true
}

// HasRatingDiff returns a boolean if a field has been set.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) HasRatingDiff() bool {
	if o != nil && !IsNil(o.RatingDiff) {
		return true
	}

	return false
}

// SetRatingDiff gets a reference to the given int32 and assigns it to the RatingDiff field.
func (o *ApiAccountPlaying200ResponseNowPlayingInner) SetRatingDiff(v int32) {
	o.RatingDiff = &v
}

func (o ApiAccountPlaying200ResponseNowPlayingInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAccountPlaying200ResponseNowPlayingInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullId"] = o.FullId
	toSerialize["gameId"] = o.GameId
	toSerialize["fen"] = o.Fen
	toSerialize["color"] = o.Color
	toSerialize["lastMove"] = o.LastMove
	toSerialize["source"] = o.Source
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	toSerialize["variant"] = o.Variant
	toSerialize["speed"] = o.Speed
	toSerialize["perf"] = o.Perf
	toSerialize["rated"] = o.Rated
	toSerialize["hasMoved"] = o.HasMoved
	toSerialize["opponent"] = o.Opponent
	toSerialize["isMyTurn"] = o.IsMyTurn
	toSerialize["secondsLeft"] = o.SecondsLeft
	if !IsNil(o.TournamentId) {
		toSerialize["tournamentId"] = o.TournamentId
	}
	if !IsNil(o.SwissId) {
		toSerialize["swissId"] = o.SwissId
	}
	if !IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}
	if !IsNil(o.RatingDiff) {
		toSerialize["ratingDiff"] = o.RatingDiff
	}
	return toSerialize, nil
}

func (o *ApiAccountPlaying200ResponseNowPlayingInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fullId",
		"gameId",
		"fen",
		"color",
		"lastMove",
		"source",
		"variant",
		"speed",
		"perf",
		"rated",
		"hasMoved",
		"opponent",
		"isMyTurn",
		"secondsLeft",
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

	varApiAccountPlaying200ResponseNowPlayingInner := _ApiAccountPlaying200ResponseNowPlayingInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiAccountPlaying200ResponseNowPlayingInner)

	if err != nil {
		return err
	}

	*o = ApiAccountPlaying200ResponseNowPlayingInner(varApiAccountPlaying200ResponseNowPlayingInner)

	return err
}

type NullableApiAccountPlaying200ResponseNowPlayingInner struct {
	value *ApiAccountPlaying200ResponseNowPlayingInner
	isSet bool
}

func (v NullableApiAccountPlaying200ResponseNowPlayingInner) Get() *ApiAccountPlaying200ResponseNowPlayingInner {
	return v.value
}

func (v *NullableApiAccountPlaying200ResponseNowPlayingInner) Set(val *ApiAccountPlaying200ResponseNowPlayingInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAccountPlaying200ResponseNowPlayingInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAccountPlaying200ResponseNowPlayingInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAccountPlaying200ResponseNowPlayingInner(val *ApiAccountPlaying200ResponseNowPlayingInner) *NullableApiAccountPlaying200ResponseNowPlayingInner {
	return &NullableApiAccountPlaying200ResponseNowPlayingInner{value: val, isSet: true}
}

func (v NullableApiAccountPlaying200ResponseNowPlayingInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAccountPlaying200ResponseNowPlayingInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


