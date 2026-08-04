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

// checks if the BroadcastRoundStudyInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastRoundStudyInfo{}

// BroadcastRoundStudyInfo struct for BroadcastRoundStudyInfo
type BroadcastRoundStudyInfo struct {
	// Whether the currently authenticated user has permission to update the study
	Writeable *bool `json:"writeable,omitempty"`
	Features *BroadcastRoundStudyInfoFeatures `json:"features,omitempty"`
}

// NewBroadcastRoundStudyInfo instantiates a new BroadcastRoundStudyInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastRoundStudyInfo() *BroadcastRoundStudyInfo {
	this := BroadcastRoundStudyInfo{}
	return &this
}

// NewBroadcastRoundStudyInfoWithDefaults instantiates a new BroadcastRoundStudyInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastRoundStudyInfoWithDefaults() *BroadcastRoundStudyInfo {
	this := BroadcastRoundStudyInfo{}
	return &this
}

// GetWriteable returns the Writeable field value if set, zero value otherwise.
func (o *BroadcastRoundStudyInfo) GetWriteable() bool {
	if o == nil || IsNil(o.Writeable) {
		var ret bool
		return ret
	}
	return *o.Writeable
}

// GetWriteableOk returns a tuple with the Writeable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundStudyInfo) GetWriteableOk() (*bool, bool) {
	if o == nil || IsNil(o.Writeable) {
		return nil, false
	}
	return o.Writeable, true
}

// HasWriteable returns a boolean if a field has been set.
func (o *BroadcastRoundStudyInfo) HasWriteable() bool {
	if o != nil && !IsNil(o.Writeable) {
		return true
	}

	return false
}

// SetWriteable gets a reference to the given bool and assigns it to the Writeable field.
func (o *BroadcastRoundStudyInfo) SetWriteable(v bool) {
	o.Writeable = &v
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *BroadcastRoundStudyInfo) GetFeatures() BroadcastRoundStudyInfoFeatures {
	if o == nil || IsNil(o.Features) {
		var ret BroadcastRoundStudyInfoFeatures
		return ret
	}
	return *o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundStudyInfo) GetFeaturesOk() (*BroadcastRoundStudyInfoFeatures, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *BroadcastRoundStudyInfo) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given BroadcastRoundStudyInfoFeatures and assigns it to the Features field.
func (o *BroadcastRoundStudyInfo) SetFeatures(v BroadcastRoundStudyInfoFeatures) {
	o.Features = &v
}

func (o BroadcastRoundStudyInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastRoundStudyInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Writeable) {
		toSerialize["writeable"] = o.Writeable
	}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	return toSerialize, nil
}

type NullableBroadcastRoundStudyInfo struct {
	value *BroadcastRoundStudyInfo
	isSet bool
}

func (v NullableBroadcastRoundStudyInfo) Get() *BroadcastRoundStudyInfo {
	return v.value
}

func (v *NullableBroadcastRoundStudyInfo) Set(val *BroadcastRoundStudyInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastRoundStudyInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastRoundStudyInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastRoundStudyInfo(val *BroadcastRoundStudyInfo) *NullableBroadcastRoundStudyInfo {
	return &NullableBroadcastRoundStudyInfo{value: val, isSet: true}
}

func (v NullableBroadcastRoundStudyInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastRoundStudyInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


