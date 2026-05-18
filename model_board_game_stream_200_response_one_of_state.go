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

// checks if the BoardGameStream200ResponseOneOfState type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BoardGameStream200ResponseOneOfState{}

// BoardGameStream200ResponseOneOfState struct for BoardGameStream200ResponseOneOfState
type BoardGameStream200ResponseOneOfState struct {
	Type string `json:"type"`
	// Current moves in UCI format (King to rook for Chess690-compatible castling notation) 
	Moves string `json:"moves"`
	// Integer of milliseconds White has left on the clock
	Wtime int32 `json:"wtime"`
	// Integer of milliseconds Black has left on the clock
	Btime int32 `json:"btime"`
	// Integer of White Fisher increment.
	Winc int32 `json:"winc"`
	// Integer of Black Fisher increment.
	Binc int32 `json:"binc"`
	Status string `json:"status"`
	// Color of the winner, if any
	Winner *string `json:"winner,omitempty"`
	// true if white is offering draw, else omitted
	Wdraw *bool `json:"wdraw,omitempty"`
	// true if black is offering draw, else omitted
	Bdraw *bool `json:"bdraw,omitempty"`
	// true if white is proposing takeback, else omitted
	Wtakeback *bool `json:"wtakeback,omitempty"`
	// true if black is proposing takeback, else omitted
	Btakeback *bool `json:"btakeback,omitempty"`
	Expiration *BoardGameStream200ResponseOneOfStateExpiration `json:"expiration,omitempty"`
}

type _BoardGameStream200ResponseOneOfState BoardGameStream200ResponseOneOfState

// NewBoardGameStream200ResponseOneOfState instantiates a new BoardGameStream200ResponseOneOfState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBoardGameStream200ResponseOneOfState(type_ string, moves string, wtime int32, btime int32, winc int32, binc int32, status string) *BoardGameStream200ResponseOneOfState {
	this := BoardGameStream200ResponseOneOfState{}
	this.Type = type_
	this.Moves = moves
	this.Wtime = wtime
	this.Btime = btime
	this.Winc = winc
	this.Binc = binc
	this.Status = status
	return &this
}

// NewBoardGameStream200ResponseOneOfStateWithDefaults instantiates a new BoardGameStream200ResponseOneOfState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBoardGameStream200ResponseOneOfStateWithDefaults() *BoardGameStream200ResponseOneOfState {
	this := BoardGameStream200ResponseOneOfState{}
	return &this
}

// GetType returns the Type field value
func (o *BoardGameStream200ResponseOneOfState) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BoardGameStream200ResponseOneOfState) SetType(v string) {
	o.Type = v
}

// GetMoves returns the Moves field value
func (o *BoardGameStream200ResponseOneOfState) GetMoves() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Moves
}

// GetMovesOk returns a tuple with the Moves field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetMovesOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Moves, true
}

// SetMoves sets field value
func (o *BoardGameStream200ResponseOneOfState) SetMoves(v string) {
	o.Moves = v
}

// GetWtime returns the Wtime field value
func (o *BoardGameStream200ResponseOneOfState) GetWtime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Wtime
}

// GetWtimeOk returns a tuple with the Wtime field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetWtimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Wtime, true
}

// SetWtime sets field value
func (o *BoardGameStream200ResponseOneOfState) SetWtime(v int32) {
	o.Wtime = v
}

// GetBtime returns the Btime field value
func (o *BoardGameStream200ResponseOneOfState) GetBtime() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Btime
}

// GetBtimeOk returns a tuple with the Btime field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetBtimeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Btime, true
}

// SetBtime sets field value
func (o *BoardGameStream200ResponseOneOfState) SetBtime(v int32) {
	o.Btime = v
}

// GetWinc returns the Winc field value
func (o *BoardGameStream200ResponseOneOfState) GetWinc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Winc
}

// GetWincOk returns a tuple with the Winc field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetWincOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Winc, true
}

// SetWinc sets field value
func (o *BoardGameStream200ResponseOneOfState) SetWinc(v int32) {
	o.Winc = v
}

// GetBinc returns the Binc field value
func (o *BoardGameStream200ResponseOneOfState) GetBinc() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Binc
}

// GetBincOk returns a tuple with the Binc field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetBincOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Binc, true
}

// SetBinc sets field value
func (o *BoardGameStream200ResponseOneOfState) SetBinc(v int32) {
	o.Binc = v
}

// GetStatus returns the Status field value
func (o *BoardGameStream200ResponseOneOfState) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BoardGameStream200ResponseOneOfState) SetStatus(v string) {
	o.Status = v
}

// GetWinner returns the Winner field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetWinner() string {
	if o == nil || IsNil(o.Winner) {
		var ret string
		return ret
	}
	return *o.Winner
}

// GetWinnerOk returns a tuple with the Winner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetWinnerOk() (*string, bool) {
	if o == nil || IsNil(o.Winner) {
		return nil, false
	}
	return o.Winner, true
}

// HasWinner returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasWinner() bool {
	if o != nil && !IsNil(o.Winner) {
		return true
	}

	return false
}

// SetWinner gets a reference to the given string and assigns it to the Winner field.
func (o *BoardGameStream200ResponseOneOfState) SetWinner(v string) {
	o.Winner = &v
}

// GetWdraw returns the Wdraw field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetWdraw() bool {
	if o == nil || IsNil(o.Wdraw) {
		var ret bool
		return ret
	}
	return *o.Wdraw
}

// GetWdrawOk returns a tuple with the Wdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetWdrawOk() (*bool, bool) {
	if o == nil || IsNil(o.Wdraw) {
		return nil, false
	}
	return o.Wdraw, true
}

// HasWdraw returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasWdraw() bool {
	if o != nil && !IsNil(o.Wdraw) {
		return true
	}

	return false
}

// SetWdraw gets a reference to the given bool and assigns it to the Wdraw field.
func (o *BoardGameStream200ResponseOneOfState) SetWdraw(v bool) {
	o.Wdraw = &v
}

// GetBdraw returns the Bdraw field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetBdraw() bool {
	if o == nil || IsNil(o.Bdraw) {
		var ret bool
		return ret
	}
	return *o.Bdraw
}

// GetBdrawOk returns a tuple with the Bdraw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetBdrawOk() (*bool, bool) {
	if o == nil || IsNil(o.Bdraw) {
		return nil, false
	}
	return o.Bdraw, true
}

// HasBdraw returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasBdraw() bool {
	if o != nil && !IsNil(o.Bdraw) {
		return true
	}

	return false
}

// SetBdraw gets a reference to the given bool and assigns it to the Bdraw field.
func (o *BoardGameStream200ResponseOneOfState) SetBdraw(v bool) {
	o.Bdraw = &v
}

// GetWtakeback returns the Wtakeback field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetWtakeback() bool {
	if o == nil || IsNil(o.Wtakeback) {
		var ret bool
		return ret
	}
	return *o.Wtakeback
}

// GetWtakebackOk returns a tuple with the Wtakeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetWtakebackOk() (*bool, bool) {
	if o == nil || IsNil(o.Wtakeback) {
		return nil, false
	}
	return o.Wtakeback, true
}

// HasWtakeback returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasWtakeback() bool {
	if o != nil && !IsNil(o.Wtakeback) {
		return true
	}

	return false
}

// SetWtakeback gets a reference to the given bool and assigns it to the Wtakeback field.
func (o *BoardGameStream200ResponseOneOfState) SetWtakeback(v bool) {
	o.Wtakeback = &v
}

// GetBtakeback returns the Btakeback field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetBtakeback() bool {
	if o == nil || IsNil(o.Btakeback) {
		var ret bool
		return ret
	}
	return *o.Btakeback
}

// GetBtakebackOk returns a tuple with the Btakeback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetBtakebackOk() (*bool, bool) {
	if o == nil || IsNil(o.Btakeback) {
		return nil, false
	}
	return o.Btakeback, true
}

// HasBtakeback returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasBtakeback() bool {
	if o != nil && !IsNil(o.Btakeback) {
		return true
	}

	return false
}

// SetBtakeback gets a reference to the given bool and assigns it to the Btakeback field.
func (o *BoardGameStream200ResponseOneOfState) SetBtakeback(v bool) {
	o.Btakeback = &v
}

// GetExpiration returns the Expiration field value if set, zero value otherwise.
func (o *BoardGameStream200ResponseOneOfState) GetExpiration() BoardGameStream200ResponseOneOfStateExpiration {
	if o == nil || IsNil(o.Expiration) {
		var ret BoardGameStream200ResponseOneOfStateExpiration
		return ret
	}
	return *o.Expiration
}

// GetExpirationOk returns a tuple with the Expiration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BoardGameStream200ResponseOneOfState) GetExpirationOk() (*BoardGameStream200ResponseOneOfStateExpiration, bool) {
	if o == nil || IsNil(o.Expiration) {
		return nil, false
	}
	return o.Expiration, true
}

// HasExpiration returns a boolean if a field has been set.
func (o *BoardGameStream200ResponseOneOfState) HasExpiration() bool {
	if o != nil && !IsNil(o.Expiration) {
		return true
	}

	return false
}

// SetExpiration gets a reference to the given BoardGameStream200ResponseOneOfStateExpiration and assigns it to the Expiration field.
func (o *BoardGameStream200ResponseOneOfState) SetExpiration(v BoardGameStream200ResponseOneOfStateExpiration) {
	o.Expiration = &v
}

func (o BoardGameStream200ResponseOneOfState) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BoardGameStream200ResponseOneOfState) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["moves"] = o.Moves
	toSerialize["wtime"] = o.Wtime
	toSerialize["btime"] = o.Btime
	toSerialize["winc"] = o.Winc
	toSerialize["binc"] = o.Binc
	toSerialize["status"] = o.Status
	if !IsNil(o.Winner) {
		toSerialize["winner"] = o.Winner
	}
	if !IsNil(o.Wdraw) {
		toSerialize["wdraw"] = o.Wdraw
	}
	if !IsNil(o.Bdraw) {
		toSerialize["bdraw"] = o.Bdraw
	}
	if !IsNil(o.Wtakeback) {
		toSerialize["wtakeback"] = o.Wtakeback
	}
	if !IsNil(o.Btakeback) {
		toSerialize["btakeback"] = o.Btakeback
	}
	if !IsNil(o.Expiration) {
		toSerialize["expiration"] = o.Expiration
	}
	return toSerialize, nil
}

func (o *BoardGameStream200ResponseOneOfState) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"moves",
		"wtime",
		"btime",
		"winc",
		"binc",
		"status",
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

	varBoardGameStream200ResponseOneOfState := _BoardGameStream200ResponseOneOfState{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBoardGameStream200ResponseOneOfState)

	if err != nil {
		return err
	}

	*o = BoardGameStream200ResponseOneOfState(varBoardGameStream200ResponseOneOfState)

	return err
}

type NullableBoardGameStream200ResponseOneOfState struct {
	value *BoardGameStream200ResponseOneOfState
	isSet bool
}

func (v NullableBoardGameStream200ResponseOneOfState) Get() *BoardGameStream200ResponseOneOfState {
	return v.value
}

func (v *NullableBoardGameStream200ResponseOneOfState) Set(val *BoardGameStream200ResponseOneOfState) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardGameStream200ResponseOneOfState) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardGameStream200ResponseOneOfState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardGameStream200ResponseOneOfState(val *BoardGameStream200ResponseOneOfState) *NullableBoardGameStream200ResponseOneOfState {
	return &NullableBoardGameStream200ResponseOneOfState{value: val, isSet: true}
}

func (v NullableBoardGameStream200ResponseOneOfState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardGameStream200ResponseOneOfState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


