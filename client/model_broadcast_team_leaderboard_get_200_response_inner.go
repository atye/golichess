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

// checks if the BroadcastTeamLeaderboardGet200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTeamLeaderboardGet200ResponseInner{}

// BroadcastTeamLeaderboardGet200ResponseInner struct for BroadcastTeamLeaderboardGet200ResponseInner
type BroadcastTeamLeaderboardGet200ResponseInner struct {
	// The name of the team
	Name string `json:"name"`
	// Total match points scored
	Mp float32 `json:"mp"`
	// Total game points scored
	Gp float32 `json:"gp"`
	// The average rating of the team's players
	AverageRating *int32 `json:"averageRating,omitempty"`
	Matches []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner `json:"matches"`
	// Players who have played for the team and their overall score
	Players []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner `json:"players"`
}

type _BroadcastTeamLeaderboardGet200ResponseInner BroadcastTeamLeaderboardGet200ResponseInner

// NewBroadcastTeamLeaderboardGet200ResponseInner instantiates a new BroadcastTeamLeaderboardGet200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTeamLeaderboardGet200ResponseInner(name string, mp float32, gp float32, matches []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner, players []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) *BroadcastTeamLeaderboardGet200ResponseInner {
	this := BroadcastTeamLeaderboardGet200ResponseInner{}
	this.Name = name
	this.Mp = mp
	this.Gp = gp
	this.Matches = matches
	this.Players = players
	return &this
}

// NewBroadcastTeamLeaderboardGet200ResponseInnerWithDefaults instantiates a new BroadcastTeamLeaderboardGet200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTeamLeaderboardGet200ResponseInnerWithDefaults() *BroadcastTeamLeaderboardGet200ResponseInner {
	this := BroadcastTeamLeaderboardGet200ResponseInner{}
	return &this
}

// GetName returns the Name field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetMp returns the Mp field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Mp
}

// GetMpOk returns a tuple with the Mp field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMpOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mp, true
}

// SetMp sets field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetMp(v float32) {
	o.Mp = v
}

// GetGp returns the Gp field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetGp() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.Gp
}

// GetGpOk returns a tuple with the Gp field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetGpOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Gp, true
}

// SetGp sets field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetGp(v float32) {
	o.Gp = v
}

// GetAverageRating returns the AverageRating field value if set, zero value otherwise.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetAverageRating() int32 {
	if o == nil || IsNil(o.AverageRating) {
		var ret int32
		return ret
	}
	return *o.AverageRating
}

// GetAverageRatingOk returns a tuple with the AverageRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetAverageRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageRating) {
		return nil, false
	}
	return o.AverageRating, true
}

// HasAverageRating returns a boolean if a field has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) HasAverageRating() bool {
	if o != nil && !IsNil(o.AverageRating) {
		return true
	}

	return false
}

// SetAverageRating gets a reference to the given int32 and assigns it to the AverageRating field.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetAverageRating(v int32) {
	o.AverageRating = &v
}

// GetMatches returns the Matches field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMatches() []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner {
	if o == nil {
		var ret []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner
		return ret
	}

	return o.Matches
}

// GetMatchesOk returns a tuple with the Matches field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetMatchesOk() ([]BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Matches, true
}

// SetMatches sets field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetMatches(v []BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner) {
	o.Matches = v
}

// GetPlayers returns the Players field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetPlayers() []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner {
	if o == nil {
		var ret []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner
		return ret
	}

	return o.Players
}

// GetPlayersOk returns a tuple with the Players field value
// and a boolean to check if the value has been set.
func (o *BroadcastTeamLeaderboardGet200ResponseInner) GetPlayersOk() ([]BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Players, true
}

// SetPlayers sets field value
func (o *BroadcastTeamLeaderboardGet200ResponseInner) SetPlayers(v []BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner) {
	o.Players = v
}

func (o BroadcastTeamLeaderboardGet200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTeamLeaderboardGet200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["mp"] = o.Mp
	toSerialize["gp"] = o.Gp
	if !IsNil(o.AverageRating) {
		toSerialize["averageRating"] = o.AverageRating
	}
	toSerialize["matches"] = o.Matches
	toSerialize["players"] = o.Players
	return toSerialize, nil
}

func (o *BroadcastTeamLeaderboardGet200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"mp",
		"gp",
		"matches",
		"players",
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

	varBroadcastTeamLeaderboardGet200ResponseInner := _BroadcastTeamLeaderboardGet200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastTeamLeaderboardGet200ResponseInner)

	if err != nil {
		return err
	}

	*o = BroadcastTeamLeaderboardGet200ResponseInner(varBroadcastTeamLeaderboardGet200ResponseInner)

	return err
}

type NullableBroadcastTeamLeaderboardGet200ResponseInner struct {
	value *BroadcastTeamLeaderboardGet200ResponseInner
	isSet bool
}

func (v NullableBroadcastTeamLeaderboardGet200ResponseInner) Get() *BroadcastTeamLeaderboardGet200ResponseInner {
	return v.value
}

func (v *NullableBroadcastTeamLeaderboardGet200ResponseInner) Set(val *BroadcastTeamLeaderboardGet200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTeamLeaderboardGet200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTeamLeaderboardGet200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTeamLeaderboardGet200ResponseInner(val *BroadcastTeamLeaderboardGet200ResponseInner) *NullableBroadcastTeamLeaderboardGet200ResponseInner {
	return &NullableBroadcastTeamLeaderboardGet200ResponseInner{value: val, isSet: true}
}

func (v NullableBroadcastTeamLeaderboardGet200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTeamLeaderboardGet200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


