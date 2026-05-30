/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.145
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the GamePlayerUserAnalysis type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GamePlayerUserAnalysis{}

// GamePlayerUserAnalysis struct for GamePlayerUserAnalysis
type GamePlayerUserAnalysis struct {
	Inaccuracy int32 `json:"inaccuracy"`
	Mistake int32 `json:"mistake"`
	Blunder int32 `json:"blunder"`
	Acpl int32 `json:"acpl"`
	Accuracy *int32 `json:"accuracy,omitempty"`
}

type _GamePlayerUserAnalysis GamePlayerUserAnalysis

// NewGamePlayerUserAnalysis instantiates a new GamePlayerUserAnalysis object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGamePlayerUserAnalysis(inaccuracy int32, mistake int32, blunder int32, acpl int32) *GamePlayerUserAnalysis {
	this := GamePlayerUserAnalysis{}
	this.Inaccuracy = inaccuracy
	this.Mistake = mistake
	this.Blunder = blunder
	this.Acpl = acpl
	return &this
}

// NewGamePlayerUserAnalysisWithDefaults instantiates a new GamePlayerUserAnalysis object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGamePlayerUserAnalysisWithDefaults() *GamePlayerUserAnalysis {
	this := GamePlayerUserAnalysis{}
	return &this
}

// GetInaccuracy returns the Inaccuracy field value
func (o *GamePlayerUserAnalysis) GetInaccuracy() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Inaccuracy
}

// GetInaccuracyOk returns a tuple with the Inaccuracy field value
// and a boolean to check if the value has been set.
func (o *GamePlayerUserAnalysis) GetInaccuracyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Inaccuracy, true
}

// SetInaccuracy sets field value
func (o *GamePlayerUserAnalysis) SetInaccuracy(v int32) {
	o.Inaccuracy = v
}

// GetMistake returns the Mistake field value
func (o *GamePlayerUserAnalysis) GetMistake() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Mistake
}

// GetMistakeOk returns a tuple with the Mistake field value
// and a boolean to check if the value has been set.
func (o *GamePlayerUserAnalysis) GetMistakeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mistake, true
}

// SetMistake sets field value
func (o *GamePlayerUserAnalysis) SetMistake(v int32) {
	o.Mistake = v
}

// GetBlunder returns the Blunder field value
func (o *GamePlayerUserAnalysis) GetBlunder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Blunder
}

// GetBlunderOk returns a tuple with the Blunder field value
// and a boolean to check if the value has been set.
func (o *GamePlayerUserAnalysis) GetBlunderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blunder, true
}

// SetBlunder sets field value
func (o *GamePlayerUserAnalysis) SetBlunder(v int32) {
	o.Blunder = v
}

// GetAcpl returns the Acpl field value
func (o *GamePlayerUserAnalysis) GetAcpl() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Acpl
}

// GetAcplOk returns a tuple with the Acpl field value
// and a boolean to check if the value has been set.
func (o *GamePlayerUserAnalysis) GetAcplOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Acpl, true
}

// SetAcpl sets field value
func (o *GamePlayerUserAnalysis) SetAcpl(v int32) {
	o.Acpl = v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *GamePlayerUserAnalysis) GetAccuracy() int32 {
	if o == nil || IsNil(o.Accuracy) {
		var ret int32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GamePlayerUserAnalysis) GetAccuracyOk() (*int32, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *GamePlayerUserAnalysis) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given int32 and assigns it to the Accuracy field.
func (o *GamePlayerUserAnalysis) SetAccuracy(v int32) {
	o.Accuracy = &v
}

func (o GamePlayerUserAnalysis) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GamePlayerUserAnalysis) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inaccuracy"] = o.Inaccuracy
	toSerialize["mistake"] = o.Mistake
	toSerialize["blunder"] = o.Blunder
	toSerialize["acpl"] = o.Acpl
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	return toSerialize, nil
}

func (o *GamePlayerUserAnalysis) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inaccuracy",
		"mistake",
		"blunder",
		"acpl",
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

	varGamePlayerUserAnalysis := _GamePlayerUserAnalysis{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGamePlayerUserAnalysis)

	if err != nil {
		return err
	}

	*o = GamePlayerUserAnalysis(varGamePlayerUserAnalysis)

	return err
}

type NullableGamePlayerUserAnalysis struct {
	value *GamePlayerUserAnalysis
	isSet bool
}

func (v NullableGamePlayerUserAnalysis) Get() *GamePlayerUserAnalysis {
	return v.value
}

func (v *NullableGamePlayerUserAnalysis) Set(val *GamePlayerUserAnalysis) {
	v.value = val
	v.isSet = true
}

func (v NullableGamePlayerUserAnalysis) IsSet() bool {
	return v.isSet
}

func (v *NullableGamePlayerUserAnalysis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGamePlayerUserAnalysis(val *GamePlayerUserAnalysis) *NullableGamePlayerUserAnalysis {
	return &NullableGamePlayerUserAnalysis{value: val, isSet: true}
}

func (v NullableGamePlayerUserAnalysis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGamePlayerUserAnalysis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


