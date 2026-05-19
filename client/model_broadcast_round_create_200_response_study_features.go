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
)

// checks if the BroadcastRoundCreate200ResponseStudyFeatures type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastRoundCreate200ResponseStudyFeatures{}

// BroadcastRoundCreate200ResponseStudyFeatures struct for BroadcastRoundCreate200ResponseStudyFeatures
type BroadcastRoundCreate200ResponseStudyFeatures struct {
	// Whether chat is enabled for the currently authenticated user
	Chat *bool `json:"chat,omitempty"`
	// Whether engine analysis is enabled for the currently authenticated user
	Computer *bool `json:"computer,omitempty"`
	// Whether the opening explorer + tablebase is enabled for the currently authenticated user
	Explorer *bool `json:"explorer,omitempty"`
}

// NewBroadcastRoundCreate200ResponseStudyFeatures instantiates a new BroadcastRoundCreate200ResponseStudyFeatures object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastRoundCreate200ResponseStudyFeatures() *BroadcastRoundCreate200ResponseStudyFeatures {
	this := BroadcastRoundCreate200ResponseStudyFeatures{}
	return &this
}

// NewBroadcastRoundCreate200ResponseStudyFeaturesWithDefaults instantiates a new BroadcastRoundCreate200ResponseStudyFeatures object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastRoundCreate200ResponseStudyFeaturesWithDefaults() *BroadcastRoundCreate200ResponseStudyFeatures {
	this := BroadcastRoundCreate200ResponseStudyFeatures{}
	return &this
}

// GetChat returns the Chat field value if set, zero value otherwise.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetChat() bool {
	if o == nil || IsNil(o.Chat) {
		var ret bool
		return ret
	}
	return *o.Chat
}

// GetChatOk returns a tuple with the Chat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetChatOk() (*bool, bool) {
	if o == nil || IsNil(o.Chat) {
		return nil, false
	}
	return o.Chat, true
}

// HasChat returns a boolean if a field has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) HasChat() bool {
	if o != nil && !IsNil(o.Chat) {
		return true
	}

	return false
}

// SetChat gets a reference to the given bool and assigns it to the Chat field.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) SetChat(v bool) {
	o.Chat = &v
}

// GetComputer returns the Computer field value if set, zero value otherwise.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetComputer() bool {
	if o == nil || IsNil(o.Computer) {
		var ret bool
		return ret
	}
	return *o.Computer
}

// GetComputerOk returns a tuple with the Computer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetComputerOk() (*bool, bool) {
	if o == nil || IsNil(o.Computer) {
		return nil, false
	}
	return o.Computer, true
}

// HasComputer returns a boolean if a field has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) HasComputer() bool {
	if o != nil && !IsNil(o.Computer) {
		return true
	}

	return false
}

// SetComputer gets a reference to the given bool and assigns it to the Computer field.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) SetComputer(v bool) {
	o.Computer = &v
}

// GetExplorer returns the Explorer field value if set, zero value otherwise.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetExplorer() bool {
	if o == nil || IsNil(o.Explorer) {
		var ret bool
		return ret
	}
	return *o.Explorer
}

// GetExplorerOk returns a tuple with the Explorer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) GetExplorerOk() (*bool, bool) {
	if o == nil || IsNil(o.Explorer) {
		return nil, false
	}
	return o.Explorer, true
}

// HasExplorer returns a boolean if a field has been set.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) HasExplorer() bool {
	if o != nil && !IsNil(o.Explorer) {
		return true
	}

	return false
}

// SetExplorer gets a reference to the given bool and assigns it to the Explorer field.
func (o *BroadcastRoundCreate200ResponseStudyFeatures) SetExplorer(v bool) {
	o.Explorer = &v
}

func (o BroadcastRoundCreate200ResponseStudyFeatures) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastRoundCreate200ResponseStudyFeatures) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chat) {
		toSerialize["chat"] = o.Chat
	}
	if !IsNil(o.Computer) {
		toSerialize["computer"] = o.Computer
	}
	if !IsNil(o.Explorer) {
		toSerialize["explorer"] = o.Explorer
	}
	return toSerialize, nil
}

type NullableBroadcastRoundCreate200ResponseStudyFeatures struct {
	value *BroadcastRoundCreate200ResponseStudyFeatures
	isSet bool
}

func (v NullableBroadcastRoundCreate200ResponseStudyFeatures) Get() *BroadcastRoundCreate200ResponseStudyFeatures {
	return v.value
}

func (v *NullableBroadcastRoundCreate200ResponseStudyFeatures) Set(val *BroadcastRoundCreate200ResponseStudyFeatures) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastRoundCreate200ResponseStudyFeatures) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastRoundCreate200ResponseStudyFeatures) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastRoundCreate200ResponseStudyFeatures(val *BroadcastRoundCreate200ResponseStudyFeatures) *NullableBroadcastRoundCreate200ResponseStudyFeatures {
	return &NullableBroadcastRoundCreate200ResponseStudyFeatures{value: val, isSet: true}
}

func (v NullableBroadcastRoundCreate200ResponseStudyFeatures) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastRoundCreate200ResponseStudyFeatures) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


