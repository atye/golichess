/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-generator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OpeningExplorerLichess200ResponseHistoryInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpeningExplorerLichess200ResponseHistoryInner{}

// OpeningExplorerLichess200ResponseHistoryInner struct for OpeningExplorerLichess200ResponseHistoryInner
type OpeningExplorerLichess200ResponseHistoryInner struct {
	Month string `json:"month"`
	White int32 `json:"white"`
	Draws int32 `json:"draws"`
	Black int32 `json:"black"`
}

type _OpeningExplorerLichess200ResponseHistoryInner OpeningExplorerLichess200ResponseHistoryInner

// NewOpeningExplorerLichess200ResponseHistoryInner instantiates a new OpeningExplorerLichess200ResponseHistoryInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpeningExplorerLichess200ResponseHistoryInner(month string, white int32, draws int32, black int32) *OpeningExplorerLichess200ResponseHistoryInner {
	this := OpeningExplorerLichess200ResponseHistoryInner{}
	this.Month = month
	this.White = white
	this.Draws = draws
	this.Black = black
	return &this
}

// NewOpeningExplorerLichess200ResponseHistoryInnerWithDefaults instantiates a new OpeningExplorerLichess200ResponseHistoryInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpeningExplorerLichess200ResponseHistoryInnerWithDefaults() *OpeningExplorerLichess200ResponseHistoryInner {
	this := OpeningExplorerLichess200ResponseHistoryInner{}
	return &this
}

// GetMonth returns the Month field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetMonth() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Month
}

// GetMonthOk returns a tuple with the Month field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetMonthOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Month, true
}

// SetMonth sets field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) SetMonth(v string) {
	o.Month = v
}

// GetWhite returns the White field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetWhite() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetWhiteOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) SetWhite(v int32) {
	o.White = v
}

// GetDraws returns the Draws field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetDraws() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draws
}

// GetDrawsOk returns a tuple with the Draws field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetDrawsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draws, true
}

// SetDraws sets field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) SetDraws(v int32) {
	o.Draws = v
}

// GetBlack returns the Black field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetBlack() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerLichess200ResponseHistoryInner) GetBlackOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *OpeningExplorerLichess200ResponseHistoryInner) SetBlack(v int32) {
	o.Black = v
}

func (o OpeningExplorerLichess200ResponseHistoryInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpeningExplorerLichess200ResponseHistoryInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["month"] = o.Month
	toSerialize["white"] = o.White
	toSerialize["draws"] = o.Draws
	toSerialize["black"] = o.Black
	return toSerialize, nil
}

func (o *OpeningExplorerLichess200ResponseHistoryInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"month",
		"white",
		"draws",
		"black",
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

	varOpeningExplorerLichess200ResponseHistoryInner := _OpeningExplorerLichess200ResponseHistoryInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpeningExplorerLichess200ResponseHistoryInner)

	if err != nil {
		return err
	}

	*o = OpeningExplorerLichess200ResponseHistoryInner(varOpeningExplorerLichess200ResponseHistoryInner)

	return err
}

type NullableOpeningExplorerLichess200ResponseHistoryInner struct {
	value *OpeningExplorerLichess200ResponseHistoryInner
	isSet bool
}

func (v NullableOpeningExplorerLichess200ResponseHistoryInner) Get() *OpeningExplorerLichess200ResponseHistoryInner {
	return v.value
}

func (v *NullableOpeningExplorerLichess200ResponseHistoryInner) Set(val *OpeningExplorerLichess200ResponseHistoryInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpeningExplorerLichess200ResponseHistoryInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpeningExplorerLichess200ResponseHistoryInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpeningExplorerLichess200ResponseHistoryInner(val *OpeningExplorerLichess200ResponseHistoryInner) *NullableOpeningExplorerLichess200ResponseHistoryInner {
	return &NullableOpeningExplorerLichess200ResponseHistoryInner{value: val, isSet: true}
}

func (v NullableOpeningExplorerLichess200ResponseHistoryInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpeningExplorerLichess200ResponseHistoryInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


