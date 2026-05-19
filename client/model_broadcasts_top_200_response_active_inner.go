/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lichess

import (
	"encoding/json"
)

// checks if the BroadcastsTop200ResponseActiveInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastsTop200ResponseActiveInner{}

// BroadcastsTop200ResponseActiveInner struct for BroadcastsTop200ResponseActiveInner
type BroadcastsTop200ResponseActiveInner struct {
	Group *string `json:"group,omitempty"`
	Tour *BroadcastsOfficial200ResponseTour `json:"tour,omitempty"`
	Round *BroadcastsOfficial200ResponseRoundsInner `json:"round,omitempty"`
}

// NewBroadcastsTop200ResponseActiveInner instantiates a new BroadcastsTop200ResponseActiveInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastsTop200ResponseActiveInner() *BroadcastsTop200ResponseActiveInner {
	this := BroadcastsTop200ResponseActiveInner{}
	return &this
}

// NewBroadcastsTop200ResponseActiveInnerWithDefaults instantiates a new BroadcastsTop200ResponseActiveInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastsTop200ResponseActiveInnerWithDefaults() *BroadcastsTop200ResponseActiveInner {
	this := BroadcastsTop200ResponseActiveInner{}
	return &this
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BroadcastsTop200ResponseActiveInner) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsTop200ResponseActiveInner) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BroadcastsTop200ResponseActiveInner) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *BroadcastsTop200ResponseActiveInner) SetGroup(v string) {
	o.Group = &v
}

// GetTour returns the Tour field value if set, zero value otherwise.
func (o *BroadcastsTop200ResponseActiveInner) GetTour() BroadcastsOfficial200ResponseTour {
	if o == nil || IsNil(o.Tour) {
		var ret BroadcastsOfficial200ResponseTour
		return ret
	}
	return *o.Tour
}

// GetTourOk returns a tuple with the Tour field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsTop200ResponseActiveInner) GetTourOk() (*BroadcastsOfficial200ResponseTour, bool) {
	if o == nil || IsNil(o.Tour) {
		return nil, false
	}
	return o.Tour, true
}

// HasTour returns a boolean if a field has been set.
func (o *BroadcastsTop200ResponseActiveInner) HasTour() bool {
	if o != nil && !IsNil(o.Tour) {
		return true
	}

	return false
}

// SetTour gets a reference to the given BroadcastsOfficial200ResponseTour and assigns it to the Tour field.
func (o *BroadcastsTop200ResponseActiveInner) SetTour(v BroadcastsOfficial200ResponseTour) {
	o.Tour = &v
}

// GetRound returns the Round field value if set, zero value otherwise.
func (o *BroadcastsTop200ResponseActiveInner) GetRound() BroadcastsOfficial200ResponseRoundsInner {
	if o == nil || IsNil(o.Round) {
		var ret BroadcastsOfficial200ResponseRoundsInner
		return ret
	}
	return *o.Round
}

// GetRoundOk returns a tuple with the Round field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsTop200ResponseActiveInner) GetRoundOk() (*BroadcastsOfficial200ResponseRoundsInner, bool) {
	if o == nil || IsNil(o.Round) {
		return nil, false
	}
	return o.Round, true
}

// HasRound returns a boolean if a field has been set.
func (o *BroadcastsTop200ResponseActiveInner) HasRound() bool {
	if o != nil && !IsNil(o.Round) {
		return true
	}

	return false
}

// SetRound gets a reference to the given BroadcastsOfficial200ResponseRoundsInner and assigns it to the Round field.
func (o *BroadcastsTop200ResponseActiveInner) SetRound(v BroadcastsOfficial200ResponseRoundsInner) {
	o.Round = &v
}

func (o BroadcastsTop200ResponseActiveInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastsTop200ResponseActiveInner) ToMap() (map[string]interface{}, error) {
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

type NullableBroadcastsTop200ResponseActiveInner struct {
	value *BroadcastsTop200ResponseActiveInner
	isSet bool
}

func (v NullableBroadcastsTop200ResponseActiveInner) Get() *BroadcastsTop200ResponseActiveInner {
	return v.value
}

func (v *NullableBroadcastsTop200ResponseActiveInner) Set(val *BroadcastsTop200ResponseActiveInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastsTop200ResponseActiveInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastsTop200ResponseActiveInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastsTop200ResponseActiveInner(val *BroadcastsTop200ResponseActiveInner) *NullableBroadcastsTop200ResponseActiveInner {
	return &NullableBroadcastsTop200ResponseActiveInner{value: val, isSet: true}
}

func (v NullableBroadcastsTop200ResponseActiveInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastsTop200ResponseActiveInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


