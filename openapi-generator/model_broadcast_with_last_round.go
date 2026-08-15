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
)

// checks if the BroadcastWithLastRound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastWithLastRound{}

// BroadcastWithLastRound struct for BroadcastWithLastRound
type BroadcastWithLastRound struct {
	Group *string `json:"group,omitempty"`
	Tour *BroadcastTour `json:"tour,omitempty"`
	Round *BroadcastRoundInfo `json:"round,omitempty"`
}

// NewBroadcastWithLastRound instantiates a new BroadcastWithLastRound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastWithLastRound() *BroadcastWithLastRound {
	this := BroadcastWithLastRound{}
	return &this
}

// NewBroadcastWithLastRoundWithDefaults instantiates a new BroadcastWithLastRound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastWithLastRoundWithDefaults() *BroadcastWithLastRound {
	this := BroadcastWithLastRound{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BroadcastWithLastRound) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastWithLastRound) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BroadcastWithLastRound) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *BroadcastWithLastRound) SetGroup(v string) {
	o.Group = &v
}

// GetTour returns the Tour field value if set, zero value otherwise.
func (o *BroadcastWithLastRound) GetTour() BroadcastTour {
	if o == nil || IsNil(o.Tour) {
		var ret BroadcastTour
		return ret
	}
	return *o.Tour
}

// GetTourOk returns a tuple with the Tour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastWithLastRound) GetTourOk() (*BroadcastTour, bool) {
	if o == nil || IsNil(o.Tour) {
		return nil, false
	}
	return o.Tour, true
}

// HasTour returns a boolean if a field has been set.
func (o *BroadcastWithLastRound) HasTour() bool {
	if o != nil && !IsNil(o.Tour) {
		return true
	}

	return false
}

// SetTour gets a reference to the given BroadcastTour and assigns it to the Tour field.
func (o *BroadcastWithLastRound) SetTour(v BroadcastTour) {
	o.Tour = &v
}

// GetRound returns the Round field value if set, zero value otherwise.
func (o *BroadcastWithLastRound) GetRound() BroadcastRoundInfo {
	if o == nil || IsNil(o.Round) {
		var ret BroadcastRoundInfo
		return ret
	}
	return *o.Round
}

// GetRoundOk returns a tuple with the Round field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastWithLastRound) GetRoundOk() (*BroadcastRoundInfo, bool) {
	if o == nil || IsNil(o.Round) {
		return nil, false
	}
	return o.Round, true
}

// HasRound returns a boolean if a field has been set.
func (o *BroadcastWithLastRound) HasRound() bool {
	if o != nil && !IsNil(o.Round) {
		return true
	}

	return false
}

// SetRound gets a reference to the given BroadcastRoundInfo and assigns it to the Round field.
func (o *BroadcastWithLastRound) SetRound(v BroadcastRoundInfo) {
	o.Round = &v
}

func (o BroadcastWithLastRound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastWithLastRound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.Tour) {
		toSerialize["tour"] = o.Tour
	}
	if !IsNil(o.Round) {
		toSerialize["round"] = o.Round
	}
	return toSerialize, nil
}

type NullableBroadcastWithLastRound struct {
	value *BroadcastWithLastRound
	isSet bool
}

func (v NullableBroadcastWithLastRound) Get() *BroadcastWithLastRound {
	return v.value
}

func (v *NullableBroadcastWithLastRound) Set(val *BroadcastWithLastRound) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastWithLastRound) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastWithLastRound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastWithLastRound(val *BroadcastWithLastRound) *NullableBroadcastWithLastRound {
	return &NullableBroadcastWithLastRound{value: val, isSet: true}
}

func (v NullableBroadcastWithLastRound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastWithLastRound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


