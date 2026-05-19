/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiStreamEvent200ResponseOneOfGame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiStreamEvent200ResponseOneOfGame{}

// ApiStreamEvent200ResponseOneOfGame struct for ApiStreamEvent200ResponseOneOfGame
type ApiStreamEvent200ResponseOneOfGame struct {
	FullId string `json:"fullId"`
	GameId string `json:"gameId"`
	Fen *string `json:"fen,omitempty"`
	Color NullableString `json:"color,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	Source *string `json:"source,omitempty"`
	Status *StreamGame200ResponseInnerOneOfStatus `json:"status,omitempty"`
	Variant *ApiAccountPlaying200ResponseNowPlayingInnerVariant `json:"variant,omitempty"`
	Speed *string `json:"speed,omitempty"`
	Perf *string `json:"perf,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	Rated *bool `json:"rated,omitempty"`
	HasMoved *bool `json:"hasMoved,omitempty"`
	Opponent *ApiStreamEvent200ResponseOneOfGameOpponent `json:"opponent,omitempty"`
	IsMyTurn *bool `json:"isMyTurn,omitempty"`
	SecondsLeft *int32 `json:"secondsLeft,omitempty"`
	Winner NullableString `json:"winner,omitempty"`
	RatingDiff *int32 `json:"ratingDiff,omitempty"`
	Compat *ApiStreamEvent200ResponseOneOfGameCompat `json:"compat,omitempty"`
	Id *string `json:"id,omitempty"`
	TournamentId *string `json:"tournamentId,omitempty"`
}

type _ApiStreamEvent200ResponseOneOfGame ApiStreamEvent200ResponseOneOfGame

// NewApiStreamEvent200ResponseOneOfGame instantiates a new ApiStreamEvent200ResponseOneOfGame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStreamEvent200ResponseOneOfGame(fullId string, gameId string) *ApiStreamEvent200ResponseOneOfGame {
	this := ApiStreamEvent200ResponseOneOfGame{}
	this.FullId = fullId
	this.GameId = gameId
	return &this
}

// NewApiStreamEvent200ResponseOneOfGameWithDefaults instantiates a new ApiStreamEvent200ResponseOneOfGame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStreamEvent200ResponseOneOfGameWithDefaults() *ApiStreamEvent200ResponseOneOfGame {
	this := ApiStreamEvent200ResponseOneOfGame{}
	return &this
}

// GetFullId returns the FullId field value
func (o *ApiStreamEvent200ResponseOneOfGame) GetFullId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullId
}

// GetFullIdOk returns a tuple with the FullId field value
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetFullIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullId, true
}

// SetFullId sets field value
func (o *ApiStreamEvent200ResponseOneOfGame) SetFullId(v string) {
	o.FullId = v
}

// GetGameId returns the GameId field value
func (o *ApiStreamEvent200ResponseOneOfGame) GetGameId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GameId
}

// GetGameIdOk returns a tuple with the GameId field value
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetGameIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GameId, true
}

// SetGameId sets field value
func (o *ApiStreamEvent200ResponseOneOfGame) SetGameId(v string) {
	o.GameId = v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetFen(v string) {
	o.Fen = &v
}

// GetColor returns the Color field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiStreamEvent200ResponseOneOfGame) GetColor() string {
	if o == nil || IsNil(o.Color.Get()) {
		var ret string
		return ret
	}
	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiStreamEvent200ResponseOneOfGame) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// HasColor returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasColor() bool {
	if o != nil && o.Color.IsSet() {
		return true
	}

	return false
}

// SetColor gets a reference to the given NullableString and assigns it to the Color field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetColor(v string) {
	o.Color.Set(&v)
}
// SetColorNil sets the value for Color to be an explicit nil
func (o *ApiStreamEvent200ResponseOneOfGame) SetColorNil() {
	o.Color.Set(nil)
}

// UnsetColor ensures that no value is present for Color, not even an explicit nil
func (o *ApiStreamEvent200ResponseOneOfGame) UnsetColor() {
	o.Color.Unset()
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetLastMove(v string) {
	o.LastMove = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSource() string {
	if o == nil || IsNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSourceOk() (*string, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetSource(v string) {
	o.Source = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetStatus() StreamGame200ResponseInnerOneOfStatus {
	if o == nil || IsNil(o.Status) {
		var ret StreamGame200ResponseInnerOneOfStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetStatusOk() (*StreamGame200ResponseInnerOneOfStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StreamGame200ResponseInnerOneOfStatus and assigns it to the Status field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetStatus(v StreamGame200ResponseInnerOneOfStatus) {
	o.Status = &v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant {
	if o == nil || IsNil(o.Variant) {
		var ret ApiAccountPlaying200ResponseNowPlayingInnerVariant
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given ApiAccountPlaying200ResponseNowPlayingInnerVariant and assigns it to the Variant field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant) {
	o.Variant = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetSpeed(v string) {
	o.Speed = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetPerf() string {
	if o == nil || IsNil(o.Perf) {
		var ret string
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetPerfOk() (*string, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given string and assigns it to the Perf field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetPerf(v string) {
	o.Perf = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetRating(v int32) {
	o.Rating = &v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetRated(v bool) {
	o.Rated = &v
}

// GetHasMoved returns the HasMoved field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetHasMoved() bool {
	if o == nil || IsNil(o.HasMoved) {
		var ret bool
		return ret
	}
	return *o.HasMoved
}

// GetHasMovedOk returns a tuple with the HasMoved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetHasMovedOk() (*bool, bool) {
	if o == nil || IsNil(o.HasMoved) {
		return nil, false
	}
	return o.HasMoved, true
}

// HasHasMoved returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasHasMoved() bool {
	if o != nil && !IsNil(o.HasMoved) {
		return true
	}

	return false
}

// SetHasMoved gets a reference to the given bool and assigns it to the HasMoved field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetHasMoved(v bool) {
	o.HasMoved = &v
}

// GetOpponent returns the Opponent field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetOpponent() ApiStreamEvent200ResponseOneOfGameOpponent {
	if o == nil || IsNil(o.Opponent) {
		var ret ApiStreamEvent200ResponseOneOfGameOpponent
		return ret
	}
	return *o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetOpponentOk() (*ApiStreamEvent200ResponseOneOfGameOpponent, bool) {
	if o == nil || IsNil(o.Opponent) {
		return nil, false
	}
	return o.Opponent, true
}

// HasOpponent returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasOpponent() bool {
	if o != nil && !IsNil(o.Opponent) {
		return true
	}

	return false
}

// SetOpponent gets a reference to the given ApiStreamEvent200ResponseOneOfGameOpponent and assigns it to the Opponent field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetOpponent(v ApiStreamEvent200ResponseOneOfGameOpponent) {
	o.Opponent = &v
}

// GetIsMyTurn returns the IsMyTurn field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetIsMyTurn() bool {
	if o == nil || IsNil(o.IsMyTurn) {
		var ret bool
		return ret
	}
	return *o.IsMyTurn
}

// GetIsMyTurnOk returns a tuple with the IsMyTurn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetIsMyTurnOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMyTurn) {
		return nil, false
	}
	return o.IsMyTurn, true
}

// HasIsMyTurn returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasIsMyTurn() bool {
	if o != nil && !IsNil(o.IsMyTurn) {
		return true
	}

	return false
}

// SetIsMyTurn gets a reference to the given bool and assigns it to the IsMyTurn field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetIsMyTurn(v bool) {
	o.IsMyTurn = &v
}

// GetSecondsLeft returns the SecondsLeft field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSecondsLeft() int32 {
	if o == nil || IsNil(o.SecondsLeft) {
		var ret int32
		return ret
	}
	return *o.SecondsLeft
}

// GetSecondsLeftOk returns a tuple with the SecondsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetSecondsLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.SecondsLeft) {
		return nil, false
	}
	return o.SecondsLeft, true
}

// HasSecondsLeft returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasSecondsLeft() bool {
	if o != nil && !IsNil(o.SecondsLeft) {
		return true
	}

	return false
}

// SetSecondsLeft gets a reference to the given int32 and assigns it to the SecondsLeft field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetSecondsLeft(v int32) {
	o.SecondsLeft = &v
}

// GetWinner returns the Winner field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ApiStreamEvent200ResponseOneOfGame) GetWinner() string {
	if o == nil || IsNil(o.Winner.Get()) {
		var ret string
		return ret
	}
	return *o.Winner.Get()
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ApiStreamEvent200ResponseOneOfGame) GetWinnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Winner.Get(), o.Winner.IsSet()
}

// HasWinner returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasWinner() bool {
	if o != nil && o.Winner.IsSet() {
		return true
	}

	return false
}

// SetWinner gets a reference to the given NullableString and assigns it to the Winner field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetWinner(v string) {
	o.Winner.Set(&v)
}
// SetWinnerNil sets the value for Winner to be an explicit nil
func (o *ApiStreamEvent200ResponseOneOfGame) SetWinnerNil() {
	o.Winner.Set(nil)
}

// UnsetWinner ensures that no value is present for Winner, not even an explicit nil
func (o *ApiStreamEvent200ResponseOneOfGame) UnsetWinner() {
	o.Winner.Unset()
}

// GetRatingDiff returns the RatingDiff field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingDiff() int32 {
	if o == nil || IsNil(o.RatingDiff) {
		var ret int32
		return ret
	}
	return *o.RatingDiff
}

// GetRatingDiffOk returns a tuple with the RatingDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetRatingDiffOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingDiff) {
		return nil, false
	}
	return o.RatingDiff, true
}

// HasRatingDiff returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasRatingDiff() bool {
	if o != nil && !IsNil(o.RatingDiff) {
		return true
	}

	return false
}

// SetRatingDiff gets a reference to the given int32 and assigns it to the RatingDiff field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetRatingDiff(v int32) {
	o.RatingDiff = &v
}

// GetCompat returns the Compat field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetCompat() ApiStreamEvent200ResponseOneOfGameCompat {
	if o == nil || IsNil(o.Compat) {
		var ret ApiStreamEvent200ResponseOneOfGameCompat
		return ret
	}
	return *o.Compat
}

// GetCompatOk returns a tuple with the Compat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetCompatOk() (*ApiStreamEvent200ResponseOneOfGameCompat, bool) {
	if o == nil || IsNil(o.Compat) {
		return nil, false
	}
	return o.Compat, true
}

// HasCompat returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasCompat() bool {
	if o != nil && !IsNil(o.Compat) {
		return true
	}

	return false
}

// SetCompat gets a reference to the given ApiStreamEvent200ResponseOneOfGameCompat and assigns it to the Compat field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetCompat(v ApiStreamEvent200ResponseOneOfGameCompat) {
	o.Compat = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetId(v string) {
	o.Id = &v
}

// GetTournamentId returns the TournamentId field value if set, zero value otherwise.
func (o *ApiStreamEvent200ResponseOneOfGame) GetTournamentId() string {
	if o == nil || IsNil(o.TournamentId) {
		var ret string
		return ret
	}
	return *o.TournamentId
}

// GetTournamentIdOk returns a tuple with the TournamentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) GetTournamentIdOk() (*string, bool) {
	if o == nil || IsNil(o.TournamentId) {
		return nil, false
	}
	return o.TournamentId, true
}

// HasTournamentId returns a boolean if a field has been set.
func (o *ApiStreamEvent200ResponseOneOfGame) HasTournamentId() bool {
	if o != nil && !IsNil(o.TournamentId) {
		return true
	}

	return false
}

// SetTournamentId gets a reference to the given string and assigns it to the TournamentId field.
func (o *ApiStreamEvent200ResponseOneOfGame) SetTournamentId(v string) {
	o.TournamentId = &v
}

func (o ApiStreamEvent200ResponseOneOfGame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiStreamEvent200ResponseOneOfGame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["fullId"] = o.FullId
	toSerialize["gameId"] = o.GameId
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if o.Color.IsSet() {
		toSerialize["color"] = o.Color.Get()
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
	if o.Winner.IsSet() {
		toSerialize["winner"] = o.Winner.Get()
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

func (o *ApiStreamEvent200ResponseOneOfGame) UnmarshalJSON(data []byte) (err error) {
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

	varApiStreamEvent200ResponseOneOfGame := _ApiStreamEvent200ResponseOneOfGame{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiStreamEvent200ResponseOneOfGame)

	if err != nil {
		return err
	}

	*o = ApiStreamEvent200ResponseOneOfGame(varApiStreamEvent200ResponseOneOfGame)

	return err
}

type NullableApiStreamEvent200ResponseOneOfGame struct {
	value *ApiStreamEvent200ResponseOneOfGame
	isSet bool
}

func (v NullableApiStreamEvent200ResponseOneOfGame) Get() *ApiStreamEvent200ResponseOneOfGame {
	return v.value
}

func (v *NullableApiStreamEvent200ResponseOneOfGame) Set(val *ApiStreamEvent200ResponseOneOfGame) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStreamEvent200ResponseOneOfGame) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStreamEvent200ResponseOneOfGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStreamEvent200ResponseOneOfGame(val *ApiStreamEvent200ResponseOneOfGame) *NullableApiStreamEvent200ResponseOneOfGame {
	return &NullableApiStreamEvent200ResponseOneOfGame{value: val, isSet: true}
}

func (v NullableApiStreamEvent200ResponseOneOfGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStreamEvent200ResponseOneOfGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


