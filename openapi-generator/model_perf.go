/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.166
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Perf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Perf{}

// Perf struct for Perf
type Perf struct {
	Games int32 `json:"games"`
	Rating int32 `json:"rating"`
	// rating deviation
	Rd int32 `json:"rd"`
	Prog int32 `json:"prog"`
	// only appears if a user's perf rating are [provisional](https://lichess.org/faq#provisional)
	Prov *bool `json:"prov,omitempty"`
	// global lichess ranking, only appears for recently active players
	Rank *int32 `json:"rank,omitempty"`
}

type _Perf Perf

// NewPerf instantiates a new Perf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerf(games int32, rating int32, rd int32, prog int32) *Perf {
	this := Perf{}
	this.Games = games
	this.Rating = rating
	this.Rd = rd
	this.Prog = prog
	return &this
}

// NewPerfWithDefaults instantiates a new Perf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfWithDefaults() *Perf {
	this := Perf{}
	return &this
}

// GetGames returns the Games field value
func (o *Perf) GetGames() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Games
}

// GetGamesOk returns a tuple with the Games field value
// and a boolean to check if the value has been set.
func (o *Perf) GetGamesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Games, true
}

// SetGames sets field value
func (o *Perf) SetGames(v int32) {
	o.Games = v
}

// GetRating returns the Rating field value
func (o *Perf) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *Perf) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *Perf) SetRating(v int32) {
	o.Rating = v
}

// GetRd returns the Rd field value
func (o *Perf) GetRd() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rd
}

// GetRdOk returns a tuple with the Rd field value
// and a boolean to check if the value has been set.
func (o *Perf) GetRdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rd, true
}

// SetRd sets field value
func (o *Perf) SetRd(v int32) {
	o.Rd = v
}

// GetProg returns the Prog field value
func (o *Perf) GetProg() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Prog
}

// GetProgOk returns a tuple with the Prog field value
// and a boolean to check if the value has been set.
func (o *Perf) GetProgOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prog, true
}

// SetProg sets field value
func (o *Perf) SetProg(v int32) {
	o.Prog = v
}

// GetProv returns the Prov field value if set, zero value otherwise.
func (o *Perf) GetProv() bool {
	if o == nil || IsNil(o.Prov) {
		var ret bool
		return ret
	}
	return *o.Prov
}

// GetProvOk returns a tuple with the Prov field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perf) GetProvOk() (*bool, bool) {
	if o == nil || IsNil(o.Prov) {
		return nil, false
	}
	return o.Prov, true
}

// HasProv returns a boolean if a field has been set.
func (o *Perf) HasProv() bool {
	if o != nil && !IsNil(o.Prov) {
		return true
	}

	return false
}

// SetProv gets a reference to the given bool and assigns it to the Prov field.
func (o *Perf) SetProv(v bool) {
	o.Prov = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *Perf) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Perf) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *Perf) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *Perf) SetRank(v int32) {
	o.Rank = &v
}

func (o Perf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Perf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["games"] = o.Games
	toSerialize["rating"] = o.Rating
	toSerialize["rd"] = o.Rd
	toSerialize["prog"] = o.Prog
	if !IsNil(o.Prov) {
		toSerialize["prov"] = o.Prov
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

func (o *Perf) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"games",
		"rating",
		"rd",
		"prog",
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

	varPerf := _Perf{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPerf)

	if err != nil {
		return err
	}

	*o = Perf(varPerf)

	return err
}

type NullablePerf struct {
	value *Perf
	isSet bool
}

func (v NullablePerf) Get() *Perf {
	return v.value
}

func (v *NullablePerf) Set(val *Perf) {
	v.value = val
	v.isSet = true
}

func (v NullablePerf) IsSet() bool {
	return v.isSet
}

func (v *NullablePerf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerf(val *Perf) *NullablePerf {
	return &NullablePerf{value: val, isSet: true}
}

func (v NullablePerf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


