/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.169
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GameEventInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameEventInfo{}

// GameEventInfo struct for GameEventInfo
type GameEventInfo struct {
	FullId string `json:"fullId"`
	GameId string `json:"gameId"`
	Fen *string `json:"fen,omitempty"`
	Color *GameColor `json:"color,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	Source *GameSource `json:"source,omitempty"`
	Status *GameStatus `json:"status,omitempty"`
	Variant *Variant `json:"variant,omitempty"`
	Speed *Speed `json:"speed,omitempty"`
	Perf *string `json:"perf,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	Rated *bool `json:"rated,omitempty"`
	HasMoved *bool `json:"hasMoved,omitempty"`
	Opponent *GameEventOpponent `json:"opponent,omitempty"`
	IsMyTurn *bool `json:"isMyTurn,omitempty"`
	SecondsLeft *int32 `json:"secondsLeft,omitempty"`
	Winner *GameColor `json:"winner,omitempty"`
	RatingDiff *int32 `json:"ratingDiff,omitempty"`
	Compat *GameCompat `json:"compat,omitempty"`
	Id *string `json:"id,omitempty"`
	TournamentId *string `json:"tournamentId,omitempty"`
}

type _GameEventInfo GameEventInfo

// NewGameEventInfo instantiates a new GameEventInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameEventInfo(fullId string, gameId string) *GameEventInfo {
	this := GameEventInfo{}
	this.FullId = fullId
	this.GameId = gameId
	return &this
}

// NewGameEventInfoWithDefaults instantiates a new GameEventInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameEventInfoWithDefaults() *GameEventInfo {
	this := GameEventInfo{}
	return &this
}

// GetFullId returns the FullId field value
func (o *GameEventInfo) GetFullId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullId
}

// GetFullIdOk returns a tuple with the FullId field value
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetFullIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullId, true
}

// SetFullId sets field value
func (o *GameEventInfo) SetFullId(v string) {
	o.FullId = v
}

// GetGameId returns the GameId field value
func (o *GameEventInfo) GetGameId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetGameIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameId, true
}

// SetGameId sets field value
func (o *GameEventInfo) SetGameId(v string) {
	o.GameId = v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *GameEventInfo) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *GameEventInfo) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *GameEventInfo) SetFen(v string) {
	o.Fen = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *GameEventInfo) GetColor() GameColor {
	if o == nil || IsNil(o.Color) {
		var ret GameColor
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetColorOk() (*GameColor, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *GameEventInfo) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given GameColor and assigns it to the Color field.
func (o *GameEventInfo) SetColor(v GameColor) {
	o.Color = &v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *GameEventInfo) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *GameEventInfo) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *GameEventInfo) SetLastMove(v string) {
	o.LastMove = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *GameEventInfo) GetSource() GameSource {
	if o == nil || IsNil(o.Source) {
		var ret GameSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetSourceOk() (*GameSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *GameEventInfo) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given GameSource and assigns it to the Source field.
func (o *GameEventInfo) SetSource(v GameSource) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GameEventInfo) GetStatus() GameStatus {
	if o == nil || IsNil(o.Status) {
		var ret GameStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetStatusOk() (*GameStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GameEventInfo) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given GameStatus and assigns it to the Status field.
func (o *GameEventInfo) SetStatus(v GameStatus) {
	o.Status = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *GameEventInfo) GetVariant() Variant {
	if o == nil || IsNil(o.Variant) {
		var ret Variant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetVariantOk() (*Variant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *GameEventInfo) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given Variant and assigns it to the Variant field.
func (o *GameEventInfo) SetVariant(v Variant) {
	o.Variant = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *GameEventInfo) GetSpeed() Speed {
	if o == nil || IsNil(o.Speed) {
		var ret Speed
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetSpeedOk() (*Speed, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *GameEventInfo) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given Speed and assigns it to the Speed field.
func (o *GameEventInfo) SetSpeed(v Speed) {
	o.Speed = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *GameEventInfo) GetPerf() string {
	if o == nil || IsNil(o.Perf) {
		var ret string
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetPerfOk() (*string, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *GameEventInfo) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given string and assigns it to the Perf field.
func (o *GameEventInfo) SetPerf(v string) {
	o.Perf = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *GameEventInfo) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *GameEventInfo) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *GameEventInfo) SetRating(v int32) {
	o.Rating = &v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *GameEventInfo) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *GameEventInfo) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *GameEventInfo) SetRated(v bool) {
	o.Rated = &v
}

// GetHasMoved returns the HasMoved field value if set, zero value otherwise.
func (o *GameEventInfo) GetHasMoved() bool {
	if o == nil || IsNil(o.HasMoved) {
		var ret bool
		return ret
	}
	return *o.HasMoved
}

// GetHasMovedOk returns a tuple with the HasMoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetHasMovedOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoved) {
		return nil, false
	}
	return o.HasMoved, true
}

// HasHasMoved returns a boolean if a field has been set.
func (o *GameEventInfo) HasHasMoved() bool {
	if o != nil && !IsNil(o.HasMoved) {
		return true
	}

	return false
}

// SetHasMoved gets a reference to the given bool and assigns it to the HasMoved field.
func (o *GameEventInfo) SetHasMoved(v bool) {
	o.HasMoved = &v
}

// GetOpponent returns the Opponent field value if set, zero value otherwise.
func (o *GameEventInfo) GetOpponent() GameEventOpponent {
	if o == nil || IsNil(o.Opponent) {
		var ret GameEventOpponent
		return ret
	}
	return *o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetOpponentOk() (*GameEventOpponent, bool) {
	if o == nil || IsNil(o.Opponent) {
		return nil, false
	}
	return o.Opponent, true
}

// HasOpponent returns a boolean if a field has been set.
func (o *GameEventInfo) HasOpponent() bool {
	if o != nil && !IsNil(o.Opponent) {
		return true
	}

	return false
}

// SetOpponent gets a reference to the given GameEventOpponent and assigns it to the Opponent field.
func (o *GameEventInfo) SetOpponent(v GameEventOpponent) {
	o.Opponent = &v
}

// GetIsMyTurn returns the IsMyTurn field value if set, zero value otherwise.
func (o *GameEventInfo) GetIsMyTurn() bool {
	if o == nil || IsNil(o.IsMyTurn) {
		var ret bool
		return ret
	}
	return *o.IsMyTurn
}

// GetIsMyTurnOk returns a tuple with the IsMyTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetIsMyTurnOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMyTurn) {
		return nil, false
	}
	return o.IsMyTurn, true
}

// HasIsMyTurn returns a boolean if a field has been set.
func (o *GameEventInfo) HasIsMyTurn() bool {
	if o != nil && !IsNil(o.IsMyTurn) {
		return true
	}

	return false
}

// SetIsMyTurn gets a reference to the given bool and assigns it to the IsMyTurn field.
func (o *GameEventInfo) SetIsMyTurn(v bool) {
	o.IsMyTurn = &v
}

// GetSecondsLeft returns the SecondsLeft field value if set, zero value otherwise.
func (o *GameEventInfo) GetSecondsLeft() int32 {
	if o == nil || IsNil(o.SecondsLeft) {
		var ret int32
		return ret
	}
	return *o.SecondsLeft
}

// GetSecondsLeftOk returns a tuple with the SecondsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetSecondsLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsLeft) {
		return nil, false
	}
	return o.SecondsLeft, true
}

// HasSecondsLeft returns a boolean if a field has been set.
func (o *GameEventInfo) HasSecondsLeft() bool {
	if o != nil && !IsNil(o.SecondsLeft) {
		return true
	}

	return false
}

// SetSecondsLeft gets a reference to the given int32 and assigns it to the SecondsLeft field.
func (o *GameEventInfo) SetSecondsLeft(v int32) {
	o.SecondsLeft = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *GameEventInfo) GetWinner() GameColor {
	if o == nil || IsNil(o.Winner) {
		var ret GameColor
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetWinnerOk() (*GameColor, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *GameEventInfo) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given GameColor and assigns it to the Winner field.
func (o *GameEventInfo) SetWinner(v GameColor) {
	o.Winner = &v
}

// GetRatingDiff returns the RatingDiff field value if set, zero value otherwise.
func (o *GameEventInfo) GetRatingDiff() int32 {
	if o == nil || IsNil(o.RatingDiff) {
		var ret int32
		return ret
	}
	return *o.RatingDiff
}

// GetRatingDiffOk returns a tuple with the RatingDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetRatingDiffOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingDiff) {
		return nil, false
	}
	return o.RatingDiff, true
}

// HasRatingDiff returns a boolean if a field has been set.
func (o *GameEventInfo) HasRatingDiff() bool {
	if o != nil && !IsNil(o.RatingDiff) {
		return true
	}

	return false
}

// SetRatingDiff gets a reference to the given int32 and assigns it to the RatingDiff field.
func (o *GameEventInfo) SetRatingDiff(v int32) {
	o.RatingDiff = &v
}

// GetCompat returns the Compat field value if set, zero value otherwise.
func (o *GameEventInfo) GetCompat() GameCompat {
	if o == nil || IsNil(o.Compat) {
		var ret GameCompat
		return ret
	}
	return *o.Compat
}

// GetCompatOk returns a tuple with the Compat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetCompatOk() (*GameCompat, bool) {
	if o == nil || IsNil(o.Compat) {
		return nil, false
	}
	return o.Compat, true
}

// HasCompat returns a boolean if a field has been set.
func (o *GameEventInfo) HasCompat() bool {
	if o != nil && !IsNil(o.Compat) {
		return true
	}

	return false
}

// SetCompat gets a reference to the given GameCompat and assigns it to the Compat field.
func (o *GameEventInfo) SetCompat(v GameCompat) {
	o.Compat = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GameEventInfo) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GameEventInfo) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GameEventInfo) SetId(v string) {
	o.Id = &v
}

// GetTournamentId returns the TournamentId field value if set, zero value otherwise.
func (o *GameEventInfo) GetTournamentId() string {
	if o == nil || IsNil(o.TournamentId) {
		var ret string
		return ret
	}
	return *o.TournamentId
}

// GetTournamentIdOk returns a tuple with the TournamentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameEventInfo) GetTournamentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TournamentId) {
		return nil, false
	}
	return o.TournamentId, true
}

// HasTournamentId returns a boolean if a field has been set.
func (o *GameEventInfo) HasTournamentId() bool {
	if o != nil && !IsNil(o.TournamentId) {
		return true
	}

	return false
}

// SetTournamentId gets a reference to the given string and assigns it to the TournamentId field.
func (o *GameEventInfo) SetTournamentId(v string) {
	o.TournamentId = &v
}

func (o GameEventInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameEventInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullId"] = o.FullId
	toSerialize["gameId"] = o.GameId
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.LastMove) {
		toSerialize["lastMove"] = o.LastMove
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Perf) {
		toSerialize["perf"] = o.Perf
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Rated) {
		toSerialize["rated"] = o.Rated
	}
	if !IsNil(o.HasMoved) {
		toSerialize["hasMoved"] = o.HasMoved
	}
	if !IsNil(o.Opponent) {
		toSerialize["opponent"] = o.Opponent
	}
	if !IsNil(o.IsMyTurn) {
		toSerialize["isMyTurn"] = o.IsMyTurn
	}
	if !IsNil(o.SecondsLeft) {
		toSerialize["secondsLeft"] = o.SecondsLeft
	}
	if !IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}
	if !IsNil(o.RatingDiff) {
		toSerialize["ratingDiff"] = o.RatingDiff
	}
	if !IsNil(o.Compat) {
		toSerialize["compat"] = o.Compat
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.TournamentId) {
		toSerialize["tournamentId"] = o.TournamentId
	}
	return toSerialize, nil
}

func (o *GameEventInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"fullId",
		"gameId",
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

	varGameEventInfo := _GameEventInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGameEventInfo)

	if err != nil {
		return err
	}

	*o = GameEventInfo(varGameEventInfo)

	return err
}

type NullableGameEventInfo struct {
	value *GameEventInfo
	isSet bool
}

func (v NullableGameEventInfo) Get() *GameEventInfo {
	return v.value
}

func (v *NullableGameEventInfo) Set(val *GameEventInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableGameEventInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableGameEventInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameEventInfo(val *GameEventInfo) *NullableGameEventInfo {
	return &NullableGameEventInfo{value: val, isSet: true}
}

func (v NullableGameEventInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameEventInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


