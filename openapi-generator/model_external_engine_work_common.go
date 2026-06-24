/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.147
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ExternalEngineWorkCommon type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEngineWorkCommon{}

// ExternalEngineWorkCommon struct for ExternalEngineWorkCommon
type ExternalEngineWorkCommon struct {
	// Arbitary string that identifies the analysis session. Providers may wish to clear the hash table between sessions. 
	SessionId string `json:"sessionId"`
	// Number of threads to use for analysis.
	Threads int32 `json:"threads"`
	// Hash table size to use for analysis, in MiB.
	Hash int32 `json:"hash"`
	// Requested number of principal variations.
	MultiPv int32 `json:"multiPv"`
	Variant UciVariant `json:"variant"`
	// Initial position of the game.
	InitialFen string `json:"initialFen"`
	// List of moves played from the initial position, in UCI notation.
	Moves []string `json:"moves"`
}

type _ExternalEngineWorkCommon ExternalEngineWorkCommon

// NewExternalEngineWorkCommon instantiates a new ExternalEngineWorkCommon object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEngineWorkCommon(sessionId string, threads int32, hash int32, multiPv int32, variant UciVariant, initialFen string, moves []string) *ExternalEngineWorkCommon {
	this := ExternalEngineWorkCommon{}
	this.SessionId = sessionId
	this.Threads = threads
	this.Hash = hash
	this.MultiPv = multiPv
	this.Variant = variant
	this.InitialFen = initialFen
	this.Moves = moves
	return &this
}

// NewExternalEngineWorkCommonWithDefaults instantiates a new ExternalEngineWorkCommon object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEngineWorkCommonWithDefaults() *ExternalEngineWorkCommon {
	this := ExternalEngineWorkCommon{}
	var variant UciVariant = UCIVARIANT_CHESS
	this.Variant = variant
	return &this
}

// GetSessionId returns the SessionId field value
func (o *ExternalEngineWorkCommon) GetSessionId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SessionId
}

// GetSessionIdOk returns a tuple with the SessionId field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetSessionIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SessionId, true
}

// SetSessionId sets field value
func (o *ExternalEngineWorkCommon) SetSessionId(v string) {
	o.SessionId = v
}

// GetThreads returns the Threads field value
func (o *ExternalEngineWorkCommon) GetThreads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Threads
}

// GetThreadsOk returns a tuple with the Threads field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Threads, true
}

// SetThreads sets field value
func (o *ExternalEngineWorkCommon) SetThreads(v int32) {
	o.Threads = v
}

// GetHash returns the Hash field value
func (o *ExternalEngineWorkCommon) GetHash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Hash
}

// GetHashOk returns a tuple with the Hash field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetHashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hash, true
}

// SetHash sets field value
func (o *ExternalEngineWorkCommon) SetHash(v int32) {
	o.Hash = v
}

// GetMultiPv returns the MultiPv field value
func (o *ExternalEngineWorkCommon) GetMultiPv() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MultiPv
}

// GetMultiPvOk returns a tuple with the MultiPv field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetMultiPvOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MultiPv, true
}

// SetMultiPv sets field value
func (o *ExternalEngineWorkCommon) SetMultiPv(v int32) {
	o.MultiPv = v
}

// GetVariant returns the Variant field value
func (o *ExternalEngineWorkCommon) GetVariant() UciVariant {
	if o == nil {
		var ret UciVariant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetVariantOk() (*UciVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ExternalEngineWorkCommon) SetVariant(v UciVariant) {
	o.Variant = v
}

// GetInitialFen returns the InitialFen field value
func (o *ExternalEngineWorkCommon) GetInitialFen() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetInitialFenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialFen, true
}

// SetInitialFen sets field value
func (o *ExternalEngineWorkCommon) SetInitialFen(v string) {
	o.InitialFen = v
}

// GetMoves returns the Moves field value
func (o *ExternalEngineWorkCommon) GetMoves() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineWorkCommon) GetMovesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Moves, true
}

// SetMoves sets field value
func (o *ExternalEngineWorkCommon) SetMoves(v []string) {
	o.Moves = v
}

func (o ExternalEngineWorkCommon) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEngineWorkCommon) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sessionId"] = o.SessionId
	toSerialize["threads"] = o.Threads
	toSerialize["hash"] = o.Hash
	toSerialize["multiPv"] = o.MultiPv
	toSerialize["variant"] = o.Variant
	toSerialize["initialFen"] = o.InitialFen
	toSerialize["moves"] = o.Moves
	return toSerialize, nil
}

func (o *ExternalEngineWorkCommon) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sessionId",
		"threads",
		"hash",
		"multiPv",
		"variant",
		"initialFen",
		"moves",
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

	varExternalEngineWorkCommon := _ExternalEngineWorkCommon{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalEngineWorkCommon)

	if err != nil {
		return err
	}

	*o = ExternalEngineWorkCommon(varExternalEngineWorkCommon)

	return err
}

type NullableExternalEngineWorkCommon struct {
	value *ExternalEngineWorkCommon
	isSet bool
}

func (v NullableExternalEngineWorkCommon) Get() *ExternalEngineWorkCommon {
	return v.value
}

func (v *NullableExternalEngineWorkCommon) Set(val *ExternalEngineWorkCommon) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEngineWorkCommon) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEngineWorkCommon) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEngineWorkCommon(val *ExternalEngineWorkCommon) *NullableExternalEngineWorkCommon {
	return &NullableExternalEngineWorkCommon{value: val, isSet: true}
}

func (v NullableExternalEngineWorkCommon) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEngineWorkCommon) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


