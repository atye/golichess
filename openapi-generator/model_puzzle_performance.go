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

// checks if the PuzzlePerformance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PuzzlePerformance{}

// PuzzlePerformance struct for PuzzlePerformance
type PuzzlePerformance struct {
	FirstWins int32 `json:"firstWins"`
	Nb int32 `json:"nb"`
	Performance int32 `json:"performance"`
	PuzzleRatingAvg int32 `json:"puzzleRatingAvg"`
	ReplayWins int32 `json:"replayWins"`
}

type _PuzzlePerformance PuzzlePerformance

// NewPuzzlePerformance instantiates a new PuzzlePerformance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPuzzlePerformance(firstWins int32, nb int32, performance int32, puzzleRatingAvg int32, replayWins int32) *PuzzlePerformance {
	this := PuzzlePerformance{}
	this.FirstWins = firstWins
	this.Nb = nb
	this.Performance = performance
	this.PuzzleRatingAvg = puzzleRatingAvg
	this.ReplayWins = replayWins
	return &this
}

// NewPuzzlePerformanceWithDefaults instantiates a new PuzzlePerformance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPuzzlePerformanceWithDefaults() *PuzzlePerformance {
	this := PuzzlePerformance{}
	return &this
}

// GetFirstWins returns the FirstWins field value
func (o *PuzzlePerformance) GetFirstWins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FirstWins
}

// GetFirstWinsOk returns a tuple with the FirstWins field value
// and a boolean to check if the value has been set.
func (o *PuzzlePerformance) GetFirstWinsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FirstWins, true
}

// SetFirstWins sets field value
func (o *PuzzlePerformance) SetFirstWins(v int32) {
	o.FirstWins = v
}

// GetNb returns the Nb field value
func (o *PuzzlePerformance) GetNb() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Nb
}

// GetNbOk returns a tuple with the Nb field value
// and a boolean to check if the value has been set.
func (o *PuzzlePerformance) GetNbOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nb, true
}

// SetNb sets field value
func (o *PuzzlePerformance) SetNb(v int32) {
	o.Nb = v
}

// GetPerformance returns the Performance field value
func (o *PuzzlePerformance) GetPerformance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value
// and a boolean to check if the value has been set.
func (o *PuzzlePerformance) GetPerformanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Performance, true
}

// SetPerformance sets field value
func (o *PuzzlePerformance) SetPerformance(v int32) {
	o.Performance = v
}

// GetPuzzleRatingAvg returns the PuzzleRatingAvg field value
func (o *PuzzlePerformance) GetPuzzleRatingAvg() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PuzzleRatingAvg
}

// GetPuzzleRatingAvgOk returns a tuple with the PuzzleRatingAvg field value
// and a boolean to check if the value has been set.
func (o *PuzzlePerformance) GetPuzzleRatingAvgOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PuzzleRatingAvg, true
}

// SetPuzzleRatingAvg sets field value
func (o *PuzzlePerformance) SetPuzzleRatingAvg(v int32) {
	o.PuzzleRatingAvg = v
}

// GetReplayWins returns the ReplayWins field value
func (o *PuzzlePerformance) GetReplayWins() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ReplayWins
}

// GetReplayWinsOk returns a tuple with the ReplayWins field value
// and a boolean to check if the value has been set.
func (o *PuzzlePerformance) GetReplayWinsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReplayWins, true
}

// SetReplayWins sets field value
func (o *PuzzlePerformance) SetReplayWins(v int32) {
	o.ReplayWins = v
}

func (o PuzzlePerformance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PuzzlePerformance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["firstWins"] = o.FirstWins
	toSerialize["nb"] = o.Nb
	toSerialize["performance"] = o.Performance
	toSerialize["puzzleRatingAvg"] = o.PuzzleRatingAvg
	toSerialize["replayWins"] = o.ReplayWins
	return toSerialize, nil
}

func (o *PuzzlePerformance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"firstWins",
		"nb",
		"performance",
		"puzzleRatingAvg",
		"replayWins",
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

	varPuzzlePerformance := _PuzzlePerformance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPuzzlePerformance)

	if err != nil {
		return err
	}

	*o = PuzzlePerformance(varPuzzlePerformance)

	return err
}

type NullablePuzzlePerformance struct {
	value *PuzzlePerformance
	isSet bool
}

func (v NullablePuzzlePerformance) Get() *PuzzlePerformance {
	return v.value
}

func (v *NullablePuzzlePerformance) Set(val *PuzzlePerformance) {
	v.value = val
	v.isSet = true
}

func (v NullablePuzzlePerformance) IsSet() bool {
	return v.isSet
}

func (v *NullablePuzzlePerformance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePuzzlePerformance(val *PuzzlePerformance) *NullablePuzzlePerformance {
	return &NullablePuzzlePerformance{value: val, isSet: true}
}

func (v NullablePuzzlePerformance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePuzzlePerformance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


