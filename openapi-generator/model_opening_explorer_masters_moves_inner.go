/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.161
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OpeningExplorerMastersMovesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpeningExplorerMastersMovesInner{}

// OpeningExplorerMastersMovesInner struct for OpeningExplorerMastersMovesInner
type OpeningExplorerMastersMovesInner struct {
	Uci string `json:"uci"`
	San string `json:"san"`
	AverageRating int32 `json:"averageRating"`
	White int32 `json:"white"`
	Draws int32 `json:"draws"`
	Black int32 `json:"black"`
	Game NullableOpeningExplorerMastersGame `json:"game"`
	Opening NullableOpeningExplorerOpening `json:"opening"`
}

type _OpeningExplorerMastersMovesInner OpeningExplorerMastersMovesInner

// NewOpeningExplorerMastersMovesInner instantiates a new OpeningExplorerMastersMovesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpeningExplorerMastersMovesInner(uci string, san string, averageRating int32, white int32, draws int32, black int32, game NullableOpeningExplorerMastersGame, opening NullableOpeningExplorerOpening) *OpeningExplorerMastersMovesInner {
	this := OpeningExplorerMastersMovesInner{}
	this.Uci = uci
	this.San = san
	this.AverageRating = averageRating
	this.White = white
	this.Draws = draws
	this.Black = black
	this.Game = game
	this.Opening = opening
	return &this
}

// NewOpeningExplorerMastersMovesInnerWithDefaults instantiates a new OpeningExplorerMastersMovesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpeningExplorerMastersMovesInnerWithDefaults() *OpeningExplorerMastersMovesInner {
	this := OpeningExplorerMastersMovesInner{}
	return &this
}

// GetUci returns the Uci field value
func (o *OpeningExplorerMastersMovesInner) GetUci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uci
}

// GetUciOk returns a tuple with the Uci field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetUciOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uci, true
}

// SetUci sets field value
func (o *OpeningExplorerMastersMovesInner) SetUci(v string) {
	o.Uci = v
}

// GetSan returns the San field value
func (o *OpeningExplorerMastersMovesInner) GetSan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.San
}

// GetSanOk returns a tuple with the San field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetSanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.San, true
}

// SetSan sets field value
func (o *OpeningExplorerMastersMovesInner) SetSan(v string) {
	o.San = v
}

// GetAverageRating returns the AverageRating field value
func (o *OpeningExplorerMastersMovesInner) GetAverageRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetAverageRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageRating, true
}

// SetAverageRating sets field value
func (o *OpeningExplorerMastersMovesInner) SetAverageRating(v int32) {
	o.AverageRating = v
}

// GetWhite returns the White field value
func (o *OpeningExplorerMastersMovesInner) GetWhite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetWhiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *OpeningExplorerMastersMovesInner) SetWhite(v int32) {
	o.White = v
}

// GetDraws returns the Draws field value
func (o *OpeningExplorerMastersMovesInner) GetDraws() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draws
}

// GetDrawsOk returns a tuple with the Draws field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetDrawsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draws, true
}

// SetDraws sets field value
func (o *OpeningExplorerMastersMovesInner) SetDraws(v int32) {
	o.Draws = v
}

// GetBlack returns the Black field value
func (o *OpeningExplorerMastersMovesInner) GetBlack() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMastersMovesInner) GetBlackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *OpeningExplorerMastersMovesInner) SetBlack(v int32) {
	o.Black = v
}

// GetGame returns the Game field value
// If the value is explicit nil, the zero value for OpeningExplorerMastersGame will be returned
func (o *OpeningExplorerMastersMovesInner) GetGame() OpeningExplorerMastersGame {
	if o == nil || o.Game.Get() == nil {
		var ret OpeningExplorerMastersGame
		return ret
	}

	return *o.Game.Get()
}

// GetGameOk returns a tuple with the Game field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpeningExplorerMastersMovesInner) GetGameOk() (*OpeningExplorerMastersGame, bool) {
	if o == nil {
		return nil, false
	}
	return o.Game.Get(), o.Game.IsSet()
}

// SetGame sets field value
func (o *OpeningExplorerMastersMovesInner) SetGame(v OpeningExplorerMastersGame) {
	o.Game.Set(&v)
}

// GetOpening returns the Opening field value
// If the value is explicit nil, the zero value for OpeningExplorerOpening will be returned
func (o *OpeningExplorerMastersMovesInner) GetOpening() OpeningExplorerOpening {
	if o == nil || o.Opening.Get() == nil {
		var ret OpeningExplorerOpening
		return ret
	}

	return *o.Opening.Get()
}

// GetOpeningOk returns a tuple with the Opening field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpeningExplorerMastersMovesInner) GetOpeningOk() (*OpeningExplorerOpening, bool) {
	if o == nil {
		return nil, false
	}
	return o.Opening.Get(), o.Opening.IsSet()
}

// SetOpening sets field value
func (o *OpeningExplorerMastersMovesInner) SetOpening(v OpeningExplorerOpening) {
	o.Opening.Set(&v)
}

func (o OpeningExplorerMastersMovesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpeningExplorerMastersMovesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uci"] = o.Uci
	toSerialize["san"] = o.San
	toSerialize["averageRating"] = o.AverageRating
	toSerialize["white"] = o.White
	toSerialize["draws"] = o.Draws
	toSerialize["black"] = o.Black
	toSerialize["game"] = o.Game.Get()
	toSerialize["opening"] = o.Opening.Get()
	return toSerialize, nil
}

func (o *OpeningExplorerMastersMovesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uci",
		"san",
		"averageRating",
		"white",
		"draws",
		"black",
		"game",
		"opening",
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

	varOpeningExplorerMastersMovesInner := _OpeningExplorerMastersMovesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpeningExplorerMastersMovesInner)

	if err != nil {
		return err
	}

	*o = OpeningExplorerMastersMovesInner(varOpeningExplorerMastersMovesInner)

	return err
}

type NullableOpeningExplorerMastersMovesInner struct {
	value *OpeningExplorerMastersMovesInner
	isSet bool
}

func (v NullableOpeningExplorerMastersMovesInner) Get() *OpeningExplorerMastersMovesInner {
	return v.value
}

func (v *NullableOpeningExplorerMastersMovesInner) Set(val *OpeningExplorerMastersMovesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpeningExplorerMastersMovesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpeningExplorerMastersMovesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpeningExplorerMastersMovesInner(val *OpeningExplorerMastersMovesInner) *NullableOpeningExplorerMastersMovesInner {
	return &NullableOpeningExplorerMastersMovesInner{value: val, isSet: true}
}

func (v NullableOpeningExplorerMastersMovesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpeningExplorerMastersMovesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


