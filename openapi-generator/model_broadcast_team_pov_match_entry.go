/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.163
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastTeamPOVMatchEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTeamPOVMatchEntry{}

// BroadcastTeamPOVMatchEntry struct for BroadcastTeamPOVMatchEntry
type BroadcastTeamPOVMatchEntry struct {
	RoundId string `json:"roundId"`
	// The name of the opposing team
	Opponent string `json:"opponent"`
	// Match points scored in this match
	Mp *float32 `json:"mp,omitempty"`
	// Game points scored in this match
	Gp *float32 `json:"gp,omitempty"`
	Points *BroadcastPointStr `json:"points,omitempty"`
}

type _BroadcastTeamPOVMatchEntry BroadcastTeamPOVMatchEntry

// NewBroadcastTeamPOVMatchEntry instantiates a new BroadcastTeamPOVMatchEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTeamPOVMatchEntry(roundId string, opponent string) *BroadcastTeamPOVMatchEntry {
	this := BroadcastTeamPOVMatchEntry{}
	this.RoundId = roundId
	this.Opponent = opponent
	return &this
}

// NewBroadcastTeamPOVMatchEntryWithDefaults instantiates a new BroadcastTeamPOVMatchEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTeamPOVMatchEntryWithDefaults() *BroadcastTeamPOVMatchEntry {
	this := BroadcastTeamPOVMatchEntry{}
	return &this
}

// GetRoundId returns the RoundId field value
func (o *BroadcastTeamPOVMatchEntry) GetRoundId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RoundId
}

// GetRoundIdOk returns a tuple with the RoundId field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamPOVMatchEntry) GetRoundIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoundId, true
}

// SetRoundId sets field value
func (o *BroadcastTeamPOVMatchEntry) SetRoundId(v string) {
	o.RoundId = v
}

// GetOpponent returns the Opponent field value
func (o *BroadcastTeamPOVMatchEntry) GetOpponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamPOVMatchEntry) GetOpponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Opponent, true
}

// SetOpponent sets field value
func (o *BroadcastTeamPOVMatchEntry) SetOpponent(v string) {
	o.Opponent = v
}

// GetMp returns the Mp field value if set, zero value otherwise.
func (o *BroadcastTeamPOVMatchEntry) GetMp() float32 {
	if o == nil || IsNil(o.Mp) {
		var ret float32
		return ret
	}
	return *o.Mp
}

// GetMpOk returns a tuple with the Mp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTeamPOVMatchEntry) GetMpOk() (*float32, bool) {
	if o == nil || IsNil(o.Mp) {
		return nil, false
	}
	return o.Mp, true
}

// HasMp returns a boolean if a field has been set.
func (o *BroadcastTeamPOVMatchEntry) HasMp() bool {
	if o != nil && !IsNil(o.Mp) {
		return true
	}

	return false
}

// SetMp gets a reference to the given float32 and assigns it to the Mp field.
func (o *BroadcastTeamPOVMatchEntry) SetMp(v float32) {
	o.Mp = &v
}

// GetGp returns the Gp field value if set, zero value otherwise.
func (o *BroadcastTeamPOVMatchEntry) GetGp() float32 {
	if o == nil || IsNil(o.Gp) {
		var ret float32
		return ret
	}
	return *o.Gp
}

// GetGpOk returns a tuple with the Gp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTeamPOVMatchEntry) GetGpOk() (*float32, bool) {
	if o == nil || IsNil(o.Gp) {
		return nil, false
	}
	return o.Gp, true
}

// HasGp returns a boolean if a field has been set.
func (o *BroadcastTeamPOVMatchEntry) HasGp() bool {
	if o != nil && !IsNil(o.Gp) {
		return true
	}

	return false
}

// SetGp gets a reference to the given float32 and assigns it to the Gp field.
func (o *BroadcastTeamPOVMatchEntry) SetGp(v float32) {
	o.Gp = &v
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BroadcastTeamPOVMatchEntry) GetPoints() BroadcastPointStr {
	if o == nil || IsNil(o.Points) {
		var ret BroadcastPointStr
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTeamPOVMatchEntry) GetPointsOk() (*BroadcastPointStr, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BroadcastTeamPOVMatchEntry) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given BroadcastPointStr and assigns it to the Points field.
func (o *BroadcastTeamPOVMatchEntry) SetPoints(v BroadcastPointStr) {
	o.Points = &v
}

func (o BroadcastTeamPOVMatchEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTeamPOVMatchEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["roundId"] = o.RoundId
	toSerialize["opponent"] = o.Opponent
	if !IsNil(o.Mp) {
		toSerialize["mp"] = o.Mp
	}
	if !IsNil(o.Gp) {
		toSerialize["gp"] = o.Gp
	}
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	return toSerialize, nil
}

func (o *BroadcastTeamPOVMatchEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"roundId",
		"opponent",
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

	varBroadcastTeamPOVMatchEntry := _BroadcastTeamPOVMatchEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastTeamPOVMatchEntry)

	if err != nil {
		return err
	}

	*o = BroadcastTeamPOVMatchEntry(varBroadcastTeamPOVMatchEntry)

	return err
}

type NullableBroadcastTeamPOVMatchEntry struct {
	value *BroadcastTeamPOVMatchEntry
	isSet bool
}

func (v NullableBroadcastTeamPOVMatchEntry) Get() *BroadcastTeamPOVMatchEntry {
	return v.value
}

func (v *NullableBroadcastTeamPOVMatchEntry) Set(val *BroadcastTeamPOVMatchEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTeamPOVMatchEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTeamPOVMatchEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTeamPOVMatchEntry(val *BroadcastTeamPOVMatchEntry) *NullableBroadcastTeamPOVMatchEntry {
	return &NullableBroadcastTeamPOVMatchEntry{value: val, isSet: true}
}

func (v NullableBroadcastTeamPOVMatchEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTeamPOVMatchEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


