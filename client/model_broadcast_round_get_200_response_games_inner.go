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
	"bytes"
	"fmt"
)

// checks if the BroadcastRoundGet200ResponseGamesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastRoundGet200ResponseGamesInner{}

// BroadcastRoundGet200ResponseGamesInner struct for BroadcastRoundGet200ResponseGamesInner
type BroadcastRoundGet200ResponseGamesInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Fen *string `json:"fen,omitempty"`
	Players []BroadcastRoundGet200ResponseGamesInnerPlayersInner `json:"players,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	Check *string `json:"check,omitempty"`
	ThinkTime *int32 `json:"thinkTime,omitempty"`
	// The result of the game
	Status *string `json:"status,omitempty"`
}

type _BroadcastRoundGet200ResponseGamesInner BroadcastRoundGet200ResponseGamesInner

// NewBroadcastRoundGet200ResponseGamesInner instantiates a new BroadcastRoundGet200ResponseGamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastRoundGet200ResponseGamesInner(id string, name string) *BroadcastRoundGet200ResponseGamesInner {
	this := BroadcastRoundGet200ResponseGamesInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewBroadcastRoundGet200ResponseGamesInnerWithDefaults instantiates a new BroadcastRoundGet200ResponseGamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastRoundGet200ResponseGamesInnerWithDefaults() *BroadcastRoundGet200ResponseGamesInner {
	this := BroadcastRoundGet200ResponseGamesInner{}
	return &this
}

// GetId returns the Id field value
func (o *BroadcastRoundGet200ResponseGamesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BroadcastRoundGet200ResponseGamesInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BroadcastRoundGet200ResponseGamesInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BroadcastRoundGet200ResponseGamesInner) SetName(v string) {
	o.Name = v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetFen(v string) {
	o.Fen = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetPlayers() []BroadcastRoundGet200ResponseGamesInnerPlayersInner {
	if o == nil || IsNil(o.Players) {
		var ret []BroadcastRoundGet200ResponseGamesInnerPlayersInner
		return ret
	}
	return o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetPlayersOk() ([]BroadcastRoundGet200ResponseGamesInnerPlayersInner, bool) {
	if o == nil || IsNil(o.Players) {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasPlayers() bool {
	if o != nil && !IsNil(o.Players) {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given []BroadcastRoundGet200ResponseGamesInnerPlayersInner and assigns it to the Players field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetPlayers(v []BroadcastRoundGet200ResponseGamesInnerPlayersInner) {
	o.Players = v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetLastMove(v string) {
	o.LastMove = &v
}

// GetCheck returns the Check field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetCheck() string {
	if o == nil || IsNil(o.Check) {
		var ret string
		return ret
	}
	return *o.Check
}

// GetCheckOk returns a tuple with the Check field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetCheckOk() (*string, bool) {
	if o == nil || IsNil(o.Check) {
		return nil, false
	}
	return o.Check, true
}

// HasCheck returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasCheck() bool {
	if o != nil && !IsNil(o.Check) {
		return true
	}

	return false
}

// SetCheck gets a reference to the given string and assigns it to the Check field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetCheck(v string) {
	o.Check = &v
}

// GetThinkTime returns the ThinkTime field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetThinkTime() int32 {
	if o == nil || IsNil(o.ThinkTime) {
		var ret int32
		return ret
	}
	return *o.ThinkTime
}

// GetThinkTimeOk returns a tuple with the ThinkTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetThinkTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ThinkTime) {
		return nil, false
	}
	return o.ThinkTime, true
}

// HasThinkTime returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasThinkTime() bool {
	if o != nil && !IsNil(o.ThinkTime) {
		return true
	}

	return false
}

// SetThinkTime gets a reference to the given int32 and assigns it to the ThinkTime field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetThinkTime(v int32) {
	o.ThinkTime = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BroadcastRoundGet200ResponseGamesInner) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BroadcastRoundGet200ResponseGamesInner) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BroadcastRoundGet200ResponseGamesInner) SetStatus(v string) {
	o.Status = &v
}

func (o BroadcastRoundGet200ResponseGamesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastRoundGet200ResponseGamesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if !IsNil(o.Players) {
		toSerialize["players"] = o.Players
	}
	if !IsNil(o.LastMove) {
		toSerialize["lastMove"] = o.LastMove
	}
	if !IsNil(o.Check) {
		toSerialize["check"] = o.Check
	}
	if !IsNil(o.ThinkTime) {
		toSerialize["thinkTime"] = o.ThinkTime
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

func (o *BroadcastRoundGet200ResponseGamesInner) UnmarshalJSON(data []byte) (err error) {
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

	varBroadcastRoundGet200ResponseGamesInner := _BroadcastRoundGet200ResponseGamesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastRoundGet200ResponseGamesInner)

	if err != nil {
		return err
	}

	*o = BroadcastRoundGet200ResponseGamesInner(varBroadcastRoundGet200ResponseGamesInner)

	return err
}

type NullableBroadcastRoundGet200ResponseGamesInner struct {
	value *BroadcastRoundGet200ResponseGamesInner
	isSet bool
}

func (v NullableBroadcastRoundGet200ResponseGamesInner) Get() *BroadcastRoundGet200ResponseGamesInner {
	return v.value
}

func (v *NullableBroadcastRoundGet200ResponseGamesInner) Set(val *BroadcastRoundGet200ResponseGamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastRoundGet200ResponseGamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastRoundGet200ResponseGamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastRoundGet200ResponseGamesInner(val *BroadcastRoundGet200ResponseGamesInner) *NullableBroadcastRoundGet200ResponseGamesInner {
	return &NullableBroadcastRoundGet200ResponseGamesInner{value: val, isSet: true}
}

func (v NullableBroadcastRoundGet200ResponseGamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastRoundGet200ResponseGamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


