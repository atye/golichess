/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.162
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
)

// checks if the StudyImportPgnChaptersChaptersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StudyImportPgnChaptersChaptersInner{}

// StudyImportPgnChaptersChaptersInner struct for StudyImportPgnChaptersChaptersInner
type StudyImportPgnChaptersChaptersInner struct {
	// The chapter ID
	Id *string `json:"id,omitempty"`
	// The chapter name
	Name *string `json:"name,omitempty"`
	Players []StudyImportPgnChaptersChaptersInnerPlayersInner `json:"players,omitempty"`
	// The chapter status
	Status *string `json:"status,omitempty"`
}

// NewStudyImportPgnChaptersChaptersInner instantiates a new StudyImportPgnChaptersChaptersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStudyImportPgnChaptersChaptersInner() *StudyImportPgnChaptersChaptersInner {
	this := StudyImportPgnChaptersChaptersInner{}
	return &this
}

// NewStudyImportPgnChaptersChaptersInnerWithDefaults instantiates a new StudyImportPgnChaptersChaptersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStudyImportPgnChaptersChaptersInnerWithDefaults() *StudyImportPgnChaptersChaptersInner {
	this := StudyImportPgnChaptersChaptersInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StudyImportPgnChaptersChaptersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyImportPgnChaptersChaptersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StudyImportPgnChaptersChaptersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StudyImportPgnChaptersChaptersInner) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StudyImportPgnChaptersChaptersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyImportPgnChaptersChaptersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StudyImportPgnChaptersChaptersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StudyImportPgnChaptersChaptersInner) SetName(v string) {
	o.Name = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *StudyImportPgnChaptersChaptersInner) GetPlayers() []StudyImportPgnChaptersChaptersInnerPlayersInner {
	if o == nil || IsNil(o.Players) {
		var ret []StudyImportPgnChaptersChaptersInnerPlayersInner
		return ret
	}
	return o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyImportPgnChaptersChaptersInner) GetPlayersOk() ([]StudyImportPgnChaptersChaptersInnerPlayersInner, bool) {
	if o == nil || IsNil(o.Players) {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *StudyImportPgnChaptersChaptersInner) HasPlayers() bool {
	if o != nil && !IsNil(o.Players) {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given []StudyImportPgnChaptersChaptersInnerPlayersInner and assigns it to the Players field.
func (o *StudyImportPgnChaptersChaptersInner) SetPlayers(v []StudyImportPgnChaptersChaptersInnerPlayersInner) {
	o.Players = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *StudyImportPgnChaptersChaptersInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StudyImportPgnChaptersChaptersInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *StudyImportPgnChaptersChaptersInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *StudyImportPgnChaptersChaptersInner) SetStatus(v string) {
	o.Status = &v
}

func (o StudyImportPgnChaptersChaptersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StudyImportPgnChaptersChaptersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Players) {
		toSerialize["players"] = o.Players
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableStudyImportPgnChaptersChaptersInner struct {
	value *StudyImportPgnChaptersChaptersInner
	isSet bool
}

func (v NullableStudyImportPgnChaptersChaptersInner) Get() *StudyImportPgnChaptersChaptersInner {
	return v.value
}

func (v *NullableStudyImportPgnChaptersChaptersInner) Set(val *StudyImportPgnChaptersChaptersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableStudyImportPgnChaptersChaptersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableStudyImportPgnChaptersChaptersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStudyImportPgnChaptersChaptersInner(val *StudyImportPgnChaptersChaptersInner) *NullableStudyImportPgnChaptersChaptersInner {
	return &NullableStudyImportPgnChaptersChaptersInner{value: val, isSet: true}
}

func (v NullableStudyImportPgnChaptersChaptersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStudyImportPgnChaptersChaptersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


