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
)

// checks if the PuzzleBatchSolveResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PuzzleBatchSolveResponse{}

// PuzzleBatchSolveResponse struct for PuzzleBatchSolveResponse
type PuzzleBatchSolveResponse struct {
	Puzzles []PuzzleAndGame `json:"puzzles,omitempty"`
	Glicko *PuzzleGlicko `json:"glicko,omitempty"`
	Rounds []PuzzleBatchSolveResponseRoundsInner `json:"rounds,omitempty"`
}

// NewPuzzleBatchSolveResponse instantiates a new PuzzleBatchSolveResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPuzzleBatchSolveResponse() *PuzzleBatchSolveResponse {
	this := PuzzleBatchSolveResponse{}
	return &this
}

// NewPuzzleBatchSolveResponseWithDefaults instantiates a new PuzzleBatchSolveResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPuzzleBatchSolveResponseWithDefaults() *PuzzleBatchSolveResponse {
	this := PuzzleBatchSolveResponse{}
	return &this
}

// GetPuzzles returns the Puzzles field value if set, zero value otherwise.
func (o *PuzzleBatchSolveResponse) GetPuzzles() []PuzzleAndGame {
	if o == nil || IsNil(o.Puzzles) {
		var ret []PuzzleAndGame
		return ret
	}
	return o.Puzzles
}

// GetPuzzlesOk returns a tuple with the Puzzles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuzzleBatchSolveResponse) GetPuzzlesOk() ([]PuzzleAndGame, bool) {
	if o == nil || IsNil(o.Puzzles) {
		return nil, false
	}
	return o.Puzzles, true
}

// HasPuzzles returns a boolean if a field has been set.
func (o *PuzzleBatchSolveResponse) HasPuzzles() bool {
	if o != nil && !IsNil(o.Puzzles) {
		return true
	}

	return false
}

// SetPuzzles gets a reference to the given []PuzzleAndGame and assigns it to the Puzzles field.
func (o *PuzzleBatchSolveResponse) SetPuzzles(v []PuzzleAndGame) {
	o.Puzzles = v
}

// GetGlicko returns the Glicko field value if set, zero value otherwise.
func (o *PuzzleBatchSolveResponse) GetGlicko() PuzzleGlicko {
	if o == nil || IsNil(o.Glicko) {
		var ret PuzzleGlicko
		return ret
	}
	return *o.Glicko
}

// GetGlickoOk returns a tuple with the Glicko field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuzzleBatchSolveResponse) GetGlickoOk() (*PuzzleGlicko, bool) {
	if o == nil || IsNil(o.Glicko) {
		return nil, false
	}
	return o.Glicko, true
}

// HasGlicko returns a boolean if a field has been set.
func (o *PuzzleBatchSolveResponse) HasGlicko() bool {
	if o != nil && !IsNil(o.Glicko) {
		return true
	}

	return false
}

// SetGlicko gets a reference to the given PuzzleGlicko and assigns it to the Glicko field.
func (o *PuzzleBatchSolveResponse) SetGlicko(v PuzzleGlicko) {
	o.Glicko = &v
}

// GetRounds returns the Rounds field value if set, zero value otherwise.
func (o *PuzzleBatchSolveResponse) GetRounds() []PuzzleBatchSolveResponseRoundsInner {
	if o == nil || IsNil(o.Rounds) {
		var ret []PuzzleBatchSolveResponseRoundsInner
		return ret
	}
	return o.Rounds
}

// GetRoundsOk returns a tuple with the Rounds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuzzleBatchSolveResponse) GetRoundsOk() ([]PuzzleBatchSolveResponseRoundsInner, bool) {
	if o == nil || IsNil(o.Rounds) {
		return nil, false
	}
	return o.Rounds, true
}

// HasRounds returns a boolean if a field has been set.
func (o *PuzzleBatchSolveResponse) HasRounds() bool {
	if o != nil && !IsNil(o.Rounds) {
		return true
	}

	return false
}

// SetRounds gets a reference to the given []PuzzleBatchSolveResponseRoundsInner and assigns it to the Rounds field.
func (o *PuzzleBatchSolveResponse) SetRounds(v []PuzzleBatchSolveResponseRoundsInner) {
	o.Rounds = v
}

func (o PuzzleBatchSolveResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PuzzleBatchSolveResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Puzzles) {
		toSerialize["puzzles"] = o.Puzzles
	}
	if !IsNil(o.Glicko) {
		toSerialize["glicko"] = o.Glicko
	}
	if !IsNil(o.Rounds) {
		toSerialize["rounds"] = o.Rounds
	}
	return toSerialize, nil
}

type NullablePuzzleBatchSolveResponse struct {
	value *PuzzleBatchSolveResponse
	isSet bool
}

func (v NullablePuzzleBatchSolveResponse) Get() *PuzzleBatchSolveResponse {
	return v.value
}

func (v *NullablePuzzleBatchSolveResponse) Set(val *PuzzleBatchSolveResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePuzzleBatchSolveResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePuzzleBatchSolveResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePuzzleBatchSolveResponse(val *PuzzleBatchSolveResponse) *NullablePuzzleBatchSolveResponse {
	return &NullablePuzzleBatchSolveResponse{value: val, isSet: true}
}

func (v NullablePuzzleBatchSolveResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePuzzleBatchSolveResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


