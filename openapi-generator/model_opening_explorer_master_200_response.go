/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-generator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OpeningExplorerMaster200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpeningExplorerMaster200Response{}

// OpeningExplorerMaster200Response struct for OpeningExplorerMaster200Response
type OpeningExplorerMaster200Response struct {
	Opening NullableOpeningExplorerMaster200ResponseOpening `json:"opening"`
	White int32 `json:"white"`
	Draws int32 `json:"draws"`
	Black int32 `json:"black"`
	Moves []OpeningExplorerMaster200ResponseMovesInner `json:"moves"`
	TopGames []OpeningExplorerMaster200ResponseTopGamesInner `json:"topGames"`
}

type _OpeningExplorerMaster200Response OpeningExplorerMaster200Response

// NewOpeningExplorerMaster200Response instantiates a new OpeningExplorerMaster200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpeningExplorerMaster200Response(opening NullableOpeningExplorerMaster200ResponseOpening, white int32, draws int32, black int32, moves []OpeningExplorerMaster200ResponseMovesInner, topGames []OpeningExplorerMaster200ResponseTopGamesInner) *OpeningExplorerMaster200Response {
	this := OpeningExplorerMaster200Response{}
	this.Opening = opening
	this.White = white
	this.Draws = draws
	this.Black = black
	this.Moves = moves
	this.TopGames = topGames
	return &this
}

// NewOpeningExplorerMaster200ResponseWithDefaults instantiates a new OpeningExplorerMaster200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpeningExplorerMaster200ResponseWithDefaults() *OpeningExplorerMaster200Response {
	this := OpeningExplorerMaster200Response{}
	return &this
}

// GetOpening returns the Opening field value
// If the value is explicit nil, the zero value for OpeningExplorerMaster200ResponseOpening will be returned
func (o *OpeningExplorerMaster200Response) GetOpening() OpeningExplorerMaster200ResponseOpening {
	if o == nil || o.Opening.Get() == nil {
		var ret OpeningExplorerMaster200ResponseOpening
		return ret
	}

	return *o.Opening.Get()
}

// GetOpeningOk returns a tuple with the Opening field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpeningExplorerMaster200Response) GetOpeningOk() (*OpeningExplorerMaster200ResponseOpening, bool) {
	if o == nil {
		return nil, false
	}
	return o.Opening.Get(), o.Opening.IsSet()
}

// SetOpening sets field value
func (o *OpeningExplorerMaster200Response) SetOpening(v OpeningExplorerMaster200ResponseOpening) {
	o.Opening.Set(&v)
}

// GetWhite returns the White field value
func (o *OpeningExplorerMaster200Response) GetWhite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200Response) GetWhiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *OpeningExplorerMaster200Response) SetWhite(v int32) {
	o.White = v
}

// GetDraws returns the Draws field value
func (o *OpeningExplorerMaster200Response) GetDraws() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draws
}

// GetDrawsOk returns a tuple with the Draws field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200Response) GetDrawsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draws, true
}

// SetDraws sets field value
func (o *OpeningExplorerMaster200Response) SetDraws(v int32) {
	o.Draws = v
}

// GetBlack returns the Black field value
func (o *OpeningExplorerMaster200Response) GetBlack() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200Response) GetBlackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *OpeningExplorerMaster200Response) SetBlack(v int32) {
	o.Black = v
}

// GetMoves returns the Moves field value
func (o *OpeningExplorerMaster200Response) GetMoves() []OpeningExplorerMaster200ResponseMovesInner {
	if o == nil {
		var ret []OpeningExplorerMaster200ResponseMovesInner
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200Response) GetMovesOk() ([]OpeningExplorerMaster200ResponseMovesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Moves, true
}

// SetMoves sets field value
func (o *OpeningExplorerMaster200Response) SetMoves(v []OpeningExplorerMaster200ResponseMovesInner) {
	o.Moves = v
}

// GetTopGames returns the TopGames field value
func (o *OpeningExplorerMaster200Response) GetTopGames() []OpeningExplorerMaster200ResponseTopGamesInner {
	if o == nil {
		var ret []OpeningExplorerMaster200ResponseTopGamesInner
		return ret
	}

	return o.TopGames
}

// GetTopGamesOk returns a tuple with the TopGames field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200Response) GetTopGamesOk() ([]OpeningExplorerMaster200ResponseTopGamesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopGames, true
}

// SetTopGames sets field value
func (o *OpeningExplorerMaster200Response) SetTopGames(v []OpeningExplorerMaster200ResponseTopGamesInner) {
	o.TopGames = v
}

func (o OpeningExplorerMaster200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpeningExplorerMaster200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["opening"] = o.Opening.Get()
	toSerialize["white"] = o.White
	toSerialize["draws"] = o.Draws
	toSerialize["black"] = o.Black
	toSerialize["moves"] = o.Moves
	toSerialize["topGames"] = o.TopGames
	return toSerialize, nil
}

func (o *OpeningExplorerMaster200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"opening",
		"white",
		"draws",
		"black",
		"moves",
		"topGames",
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

	varOpeningExplorerMaster200Response := _OpeningExplorerMaster200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpeningExplorerMaster200Response)

	if err != nil {
		return err
	}

	*o = OpeningExplorerMaster200Response(varOpeningExplorerMaster200Response)

	return err
}

type NullableOpeningExplorerMaster200Response struct {
	value *OpeningExplorerMaster200Response
	isSet bool
}

func (v NullableOpeningExplorerMaster200Response) Get() *OpeningExplorerMaster200Response {
	return v.value
}

func (v *NullableOpeningExplorerMaster200Response) Set(val *OpeningExplorerMaster200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableOpeningExplorerMaster200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableOpeningExplorerMaster200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpeningExplorerMaster200Response(val *OpeningExplorerMaster200Response) *NullableOpeningExplorerMaster200Response {
	return &NullableOpeningExplorerMaster200Response{value: val, isSet: true}
}

func (v NullableOpeningExplorerMaster200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpeningExplorerMaster200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


