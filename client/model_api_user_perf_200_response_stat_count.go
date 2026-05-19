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

// checks if the ApiUserPerf200ResponseStatCount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserPerf200ResponseStatCount{}

// ApiUserPerf200ResponseStatCount struct for ApiUserPerf200ResponseStatCount
type ApiUserPerf200ResponseStatCount struct {
	All int32 `json:"all"`
	Rated int32 `json:"rated"`
	Win int32 `json:"win"`
	Loss int32 `json:"loss"`
	Draw int32 `json:"draw"`
	Tour int32 `json:"tour"`
	Berserk int32 `json:"berserk"`
	OpAvg float32 `json:"opAvg"`
	Seconds int32 `json:"seconds"`
	Disconnects int32 `json:"disconnects"`
}

type _ApiUserPerf200ResponseStatCount ApiUserPerf200ResponseStatCount

// NewApiUserPerf200ResponseStatCount instantiates a new ApiUserPerf200ResponseStatCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserPerf200ResponseStatCount(all int32, rated int32, win int32, loss int32, draw int32, tour int32, berserk int32, opAvg float32, seconds int32, disconnects int32) *ApiUserPerf200ResponseStatCount {
	this := ApiUserPerf200ResponseStatCount{}
	this.All = all
	this.Rated = rated
	this.Win = win
	this.Loss = loss
	this.Draw = draw
	this.Tour = tour
	this.Berserk = berserk
	this.OpAvg = opAvg
	this.Seconds = seconds
	this.Disconnects = disconnects
	return &this
}

// NewApiUserPerf200ResponseStatCountWithDefaults instantiates a new ApiUserPerf200ResponseStatCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserPerf200ResponseStatCountWithDefaults() *ApiUserPerf200ResponseStatCount {
	this := ApiUserPerf200ResponseStatCount{}
	return &this
}

// GetAll returns the All field value
func (o *ApiUserPerf200ResponseStatCount) GetAll() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetAllOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.All, true
}

// SetAll sets field value
func (o *ApiUserPerf200ResponseStatCount) SetAll(v int32) {
	o.All = v
}

// GetRated returns the Rated field value
func (o *ApiUserPerf200ResponseStatCount) GetRated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetRatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *ApiUserPerf200ResponseStatCount) SetRated(v int32) {
	o.Rated = v
}

// GetWin returns the Win field value
func (o *ApiUserPerf200ResponseStatCount) GetWin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Win
}

// GetWinOk returns a tuple with the Win field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetWinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Win, true
}

// SetWin sets field value
func (o *ApiUserPerf200ResponseStatCount) SetWin(v int32) {
	o.Win = v
}

// GetLoss returns the Loss field value
func (o *ApiUserPerf200ResponseStatCount) GetLoss() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Loss
}

// GetLossOk returns a tuple with the Loss field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetLossOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loss, true
}

// SetLoss sets field value
func (o *ApiUserPerf200ResponseStatCount) SetLoss(v int32) {
	o.Loss = v
}

// GetDraw returns the Draw field value
func (o *ApiUserPerf200ResponseStatCount) GetDraw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draw
}

// GetDrawOk returns a tuple with the Draw field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draw, true
}

// SetDraw sets field value
func (o *ApiUserPerf200ResponseStatCount) SetDraw(v int32) {
	o.Draw = v
}

// GetTour returns the Tour field value
func (o *ApiUserPerf200ResponseStatCount) GetTour() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Tour
}

// GetTourOk returns a tuple with the Tour field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetTourOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tour, true
}

// SetTour sets field value
func (o *ApiUserPerf200ResponseStatCount) SetTour(v int32) {
	o.Tour = v
}

// GetBerserk returns the Berserk field value
func (o *ApiUserPerf200ResponseStatCount) GetBerserk() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Berserk
}

// GetBerserkOk returns a tuple with the Berserk field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetBerserkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Berserk, true
}

// SetBerserk sets field value
func (o *ApiUserPerf200ResponseStatCount) SetBerserk(v int32) {
	o.Berserk = v
}

// GetOpAvg returns the OpAvg field value
func (o *ApiUserPerf200ResponseStatCount) GetOpAvg() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.OpAvg
}

// GetOpAvgOk returns a tuple with the OpAvg field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetOpAvgOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OpAvg, true
}

// SetOpAvg sets field value
func (o *ApiUserPerf200ResponseStatCount) SetOpAvg(v float32) {
	o.OpAvg = v
}

// GetSeconds returns the Seconds field value
func (o *ApiUserPerf200ResponseStatCount) GetSeconds() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Seconds
}

// GetSecondsOk returns a tuple with the Seconds field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetSecondsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Seconds, true
}

// SetSeconds sets field value
func (o *ApiUserPerf200ResponseStatCount) SetSeconds(v int32) {
	o.Seconds = v
}

// GetDisconnects returns the Disconnects field value
func (o *ApiUserPerf200ResponseStatCount) GetDisconnects() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Disconnects
}

// GetDisconnectsOk returns a tuple with the Disconnects field value
// and a boolean to check if the value has been set.
func (o *ApiUserPerf200ResponseStatCount) GetDisconnectsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Disconnects, true
}

// SetDisconnects sets field value
func (o *ApiUserPerf200ResponseStatCount) SetDisconnects(v int32) {
	o.Disconnects = v
}

func (o ApiUserPerf200ResponseStatCount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserPerf200ResponseStatCount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all"] = o.All
	toSerialize["rated"] = o.Rated
	toSerialize["win"] = o.Win
	toSerialize["loss"] = o.Loss
	toSerialize["draw"] = o.Draw
	toSerialize["tour"] = o.Tour
	toSerialize["berserk"] = o.Berserk
	toSerialize["opAvg"] = o.OpAvg
	toSerialize["seconds"] = o.Seconds
	toSerialize["disconnects"] = o.Disconnects
	return toSerialize, nil
}

func (o *ApiUserPerf200ResponseStatCount) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all",
		"rated",
		"win",
		"loss",
		"draw",
		"tour",
		"berserk",
		"opAvg",
		"seconds",
		"disconnects",
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

	varApiUserPerf200ResponseStatCount := _ApiUserPerf200ResponseStatCount{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUserPerf200ResponseStatCount)

	if err != nil {
		return err
	}

	*o = ApiUserPerf200ResponseStatCount(varApiUserPerf200ResponseStatCount)

	return err
}

type NullableApiUserPerf200ResponseStatCount struct {
	value *ApiUserPerf200ResponseStatCount
	isSet bool
}

func (v NullableApiUserPerf200ResponseStatCount) Get() *ApiUserPerf200ResponseStatCount {
	return v.value
}

func (v *NullableApiUserPerf200ResponseStatCount) Set(val *ApiUserPerf200ResponseStatCount) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserPerf200ResponseStatCount) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserPerf200ResponseStatCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserPerf200ResponseStatCount(val *ApiUserPerf200ResponseStatCount) *NullableApiUserPerf200ResponseStatCount {
	return &NullableApiUserPerf200ResponseStatCount{value: val, isSet: true}
}

func (v NullableApiUserPerf200ResponseStatCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserPerf200ResponseStatCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


