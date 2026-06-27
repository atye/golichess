/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.149
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Team type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Team{}

// Team struct for Team
type Team struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
	Flair *string `json:"flair,omitempty"`
	Leader *LightUser `json:"leader,omitempty"`
	Leaders []LightUser `json:"leaders,omitempty"`
	NbMembers *int32 `json:"nbMembers,omitempty"`
	Open *bool `json:"open,omitempty"`
	Joined *bool `json:"joined,omitempty"`
	Requested *bool `json:"requested,omitempty"`
}

type _Team Team

// NewTeam instantiates a new Team object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTeam(id string, name string) *Team {
	this := Team{}
	this.Id = id
	this.Name = name
	return &this
}

// NewTeamWithDefaults instantiates a new Team object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTeamWithDefaults() *Team {
	this := Team{}
	return &this
}

// GetId returns the Id field value
func (o *Team) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Team) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Team) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Team) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Team) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Team) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Team) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Team) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Team) SetDescription(v string) {
	o.Description = &v
}

// GetFlair returns the Flair field value if set, zero value otherwise.
func (o *Team) GetFlair() string {
	if o == nil || IsNil(o.Flair) {
		var ret string
		return ret
	}
	return *o.Flair
}

// GetFlairOk returns a tuple with the Flair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetFlairOk() (*string, bool) {
	if o == nil || IsNil(o.Flair) {
		return nil, false
	}
	return o.Flair, true
}

// HasFlair returns a boolean if a field has been set.
func (o *Team) HasFlair() bool {
	if o != nil && !IsNil(o.Flair) {
		return true
	}

	return false
}

// SetFlair gets a reference to the given string and assigns it to the Flair field.
func (o *Team) SetFlair(v string) {
	o.Flair = &v
}

// GetLeader returns the Leader field value if set, zero value otherwise.
func (o *Team) GetLeader() LightUser {
	if o == nil || IsNil(o.Leader) {
		var ret LightUser
		return ret
	}
	return *o.Leader
}

// GetLeaderOk returns a tuple with the Leader field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetLeaderOk() (*LightUser, bool) {
	if o == nil || IsNil(o.Leader) {
		return nil, false
	}
	return o.Leader, true
}

// HasLeader returns a boolean if a field has been set.
func (o *Team) HasLeader() bool {
	if o != nil && !IsNil(o.Leader) {
		return true
	}

	return false
}

// SetLeader gets a reference to the given LightUser and assigns it to the Leader field.
func (o *Team) SetLeader(v LightUser) {
	o.Leader = &v
}

// GetLeaders returns the Leaders field value if set, zero value otherwise.
func (o *Team) GetLeaders() []LightUser {
	if o == nil || IsNil(o.Leaders) {
		var ret []LightUser
		return ret
	}
	return o.Leaders
}

// GetLeadersOk returns a tuple with the Leaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetLeadersOk() ([]LightUser, bool) {
	if o == nil || IsNil(o.Leaders) {
		return nil, false
	}
	return o.Leaders, true
}

// HasLeaders returns a boolean if a field has been set.
func (o *Team) HasLeaders() bool {
	if o != nil && !IsNil(o.Leaders) {
		return true
	}

	return false
}

// SetLeaders gets a reference to the given []LightUser and assigns it to the Leaders field.
func (o *Team) SetLeaders(v []LightUser) {
	o.Leaders = v
}

// GetNbMembers returns the NbMembers field value if set, zero value otherwise.
func (o *Team) GetNbMembers() int32 {
	if o == nil || IsNil(o.NbMembers) {
		var ret int32
		return ret
	}
	return *o.NbMembers
}

// GetNbMembersOk returns a tuple with the NbMembers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetNbMembersOk() (*int32, bool) {
	if o == nil || IsNil(o.NbMembers) {
		return nil, false
	}
	return o.NbMembers, true
}

// HasNbMembers returns a boolean if a field has been set.
func (o *Team) HasNbMembers() bool {
	if o != nil && !IsNil(o.NbMembers) {
		return true
	}

	return false
}

// SetNbMembers gets a reference to the given int32 and assigns it to the NbMembers field.
func (o *Team) SetNbMembers(v int32) {
	o.NbMembers = &v
}

// GetOpen returns the Open field value if set, zero value otherwise.
func (o *Team) GetOpen() bool {
	if o == nil || IsNil(o.Open) {
		var ret bool
		return ret
	}
	return *o.Open
}

// GetOpenOk returns a tuple with the Open field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetOpenOk() (*bool, bool) {
	if o == nil || IsNil(o.Open) {
		return nil, false
	}
	return o.Open, true
}

// HasOpen returns a boolean if a field has been set.
func (o *Team) HasOpen() bool {
	if o != nil && !IsNil(o.Open) {
		return true
	}

	return false
}

// SetOpen gets a reference to the given bool and assigns it to the Open field.
func (o *Team) SetOpen(v bool) {
	o.Open = &v
}

// GetJoined returns the Joined field value if set, zero value otherwise.
func (o *Team) GetJoined() bool {
	if o == nil || IsNil(o.Joined) {
		var ret bool
		return ret
	}
	return *o.Joined
}

// GetJoinedOk returns a tuple with the Joined field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetJoinedOk() (*bool, bool) {
	if o == nil || IsNil(o.Joined) {
		return nil, false
	}
	return o.Joined, true
}

// HasJoined returns a boolean if a field has been set.
func (o *Team) HasJoined() bool {
	if o != nil && !IsNil(o.Joined) {
		return true
	}

	return false
}

// SetJoined gets a reference to the given bool and assigns it to the Joined field.
func (o *Team) SetJoined(v bool) {
	o.Joined = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *Team) GetRequested() bool {
	if o == nil || IsNil(o.Requested) {
		var ret bool
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Team) GetRequestedOk() (*bool, bool) {
	if o == nil || IsNil(o.Requested) {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *Team) HasRequested() bool {
	if o != nil && !IsNil(o.Requested) {
		return true
	}

	return false
}

// SetRequested gets a reference to the given bool and assigns it to the Requested field.
func (o *Team) SetRequested(v bool) {
	o.Requested = &v
}

func (o Team) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Team) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Flair) {
		toSerialize["flair"] = o.Flair
	}
	if !IsNil(o.Leader) {
		toSerialize["leader"] = o.Leader
	}
	if !IsNil(o.Leaders) {
		toSerialize["leaders"] = o.Leaders
	}
	if !IsNil(o.NbMembers) {
		toSerialize["nbMembers"] = o.NbMembers
	}
	if !IsNil(o.Open) {
		toSerialize["open"] = o.Open
	}
	if !IsNil(o.Joined) {
		toSerialize["joined"] = o.Joined
	}
	if !IsNil(o.Requested) {
		toSerialize["requested"] = o.Requested
	}
	return toSerialize, nil
}

func (o *Team) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varTeam := _Team{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTeam)

	if err != nil {
		return err
	}

	*o = Team(varTeam)

	return err
}

type NullableTeam struct {
	value *Team
	isSet bool
}

func (v NullableTeam) Get() *Team {
	return v.value
}

func (v *NullableTeam) Set(val *Team) {
	v.value = val
	v.isSet = true
}

func (v NullableTeam) IsSet() bool {
	return v.isSet
}

func (v *NullableTeam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeam(val *Team) *NullableTeam {
	return &NullableTeam{value: val, isSet: true}
}

func (v NullableTeam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


