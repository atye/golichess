/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiTournamentPost200ResponseStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTournamentPost200ResponseStats{}

// ApiTournamentPost200ResponseStats struct for ApiTournamentPost200ResponseStats
type ApiTournamentPost200ResponseStats struct {
	Games int32 `json:"games"`
	Moves int32 `json:"moves"`
	WhiteWins int32 `json:"whiteWins"`
	BlackWins int32 `json:"blackWins"`
	Draws int32 `json:"draws"`
	Berserks int32 `json:"berserks"`
	AverageRating int32 `json:"averageRating"`
}

type _ApiTournamentPost200ResponseStats ApiTournamentPost200ResponseStats

// NewApiTournamentPost200ResponseStats instantiates a new ApiTournamentPost200ResponseStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTournamentPost200ResponseStats(games int32, moves int32, whiteWins int32, blackWins int32, draws int32, berserks int32, averageRating int32) *ApiTournamentPost200ResponseStats {
	this := ApiTournamentPost200ResponseStats{}
	this.Games = games
	this.Moves = moves
	this.WhiteWins = whiteWins
	this.BlackWins = blackWins
	this.Draws = draws
	this.Berserks = berserks
	this.AverageRating = averageRating
	return &this
}

// NewApiTournamentPost200ResponseStatsWithDefaults instantiates a new ApiTournamentPost200ResponseStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTournamentPost200ResponseStatsWithDefaults() *ApiTournamentPost200ResponseStats {
	this := ApiTournamentPost200ResponseStats{}
	return &this
}

// GetGames returns the Games field value
func (o *ApiTournamentPost200ResponseStats) GetGames() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Games
}

// GetGamesOk returns a tuple with the Games field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetGamesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Games, true
}

// SetGames sets field value
func (o *ApiTournamentPost200ResponseStats) SetGames(v int32) {
	o.Games = v
}

// GetMoves returns the Moves field value
func (o *ApiTournamentPost200ResponseStats) GetMoves() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetMovesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Moves, true
}

// SetMoves sets field value
func (o *ApiTournamentPost200ResponseStats) SetMoves(v int32) {
	o.Moves = v
}

// GetWhiteWins returns the WhiteWins field value
func (o *ApiTournamentPost200ResponseStats) GetWhiteWins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.WhiteWins
}

// GetWhiteWinsOk returns a tuple with the WhiteWins field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetWhiteWinsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WhiteWins, true
}

// SetWhiteWins sets field value
func (o *ApiTournamentPost200ResponseStats) SetWhiteWins(v int32) {
	o.WhiteWins = v
}

// GetBlackWins returns the BlackWins field value
func (o *ApiTournamentPost200ResponseStats) GetBlackWins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.BlackWins
}

// GetBlackWinsOk returns a tuple with the BlackWins field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetBlackWinsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BlackWins, true
}

// SetBlackWins sets field value
func (o *ApiTournamentPost200ResponseStats) SetBlackWins(v int32) {
	o.BlackWins = v
}

// GetDraws returns the Draws field value
func (o *ApiTournamentPost200ResponseStats) GetDraws() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draws
}

// GetDrawsOk returns a tuple with the Draws field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetDrawsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draws, true
}

// SetDraws sets field value
func (o *ApiTournamentPost200ResponseStats) SetDraws(v int32) {
	o.Draws = v
}

// GetBerserks returns the Berserks field value
func (o *ApiTournamentPost200ResponseStats) GetBerserks() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Berserks
}

// GetBerserksOk returns a tuple with the Berserks field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetBerserksOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Berserks, true
}

// SetBerserks sets field value
func (o *ApiTournamentPost200ResponseStats) SetBerserks(v int32) {
	o.Berserks = v
}

// GetAverageRating returns the AverageRating field value
func (o *ApiTournamentPost200ResponseStats) GetAverageRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseStats) GetAverageRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageRating, true
}

// SetAverageRating sets field value
func (o *ApiTournamentPost200ResponseStats) SetAverageRating(v int32) {
	o.AverageRating = v
}

func (o ApiTournamentPost200ResponseStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTournamentPost200ResponseStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["games"] = o.Games
	toSerialize["moves"] = o.Moves
	toSerialize["whiteWins"] = o.WhiteWins
	toSerialize["blackWins"] = o.BlackWins
	toSerialize["draws"] = o.Draws
	toSerialize["berserks"] = o.Berserks
	toSerialize["averageRating"] = o.AverageRating
	return toSerialize, nil
}

func (o *ApiTournamentPost200ResponseStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"games",
		"moves",
		"whiteWins",
		"blackWins",
		"draws",
		"berserks",
		"averageRating",
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

	varApiTournamentPost200ResponseStats := _ApiTournamentPost200ResponseStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiTournamentPost200ResponseStats)

	if err != nil {
		return err
	}

	*o = ApiTournamentPost200ResponseStats(varApiTournamentPost200ResponseStats)

	return err
}

type NullableApiTournamentPost200ResponseStats struct {
	value *ApiTournamentPost200ResponseStats
	isSet bool
}

func (v NullableApiTournamentPost200ResponseStats) Get() *ApiTournamentPost200ResponseStats {
	return v.value
}

func (v *NullableApiTournamentPost200ResponseStats) Set(val *ApiTournamentPost200ResponseStats) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTournamentPost200ResponseStats) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTournamentPost200ResponseStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTournamentPost200ResponseStats(val *ApiTournamentPost200ResponseStats) *NullableApiTournamentPost200ResponseStats {
	return &NullableApiTournamentPost200ResponseStats{value: val, isSet: true}
}

func (v NullableApiTournamentPost200ResponseStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTournamentPost200ResponseStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


