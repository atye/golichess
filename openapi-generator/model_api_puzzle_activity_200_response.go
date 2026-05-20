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

// checks if the ApiPuzzleActivity200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiPuzzleActivity200Response{}

// ApiPuzzleActivity200Response struct for ApiPuzzleActivity200Response
type ApiPuzzleActivity200Response struct {
	Date int32 `json:"date"`
	Puzzle ApiPuzzleActivity200ResponsePuzzle `json:"puzzle"`
	Win bool `json:"win"`
}

type _ApiPuzzleActivity200Response ApiPuzzleActivity200Response

// NewApiPuzzleActivity200Response instantiates a new ApiPuzzleActivity200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiPuzzleActivity200Response(date int32, puzzle ApiPuzzleActivity200ResponsePuzzle, win bool) *ApiPuzzleActivity200Response {
	this := ApiPuzzleActivity200Response{}
	this.Date = date
	this.Puzzle = puzzle
	this.Win = win
	return &this
}

// NewApiPuzzleActivity200ResponseWithDefaults instantiates a new ApiPuzzleActivity200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiPuzzleActivity200ResponseWithDefaults() *ApiPuzzleActivity200Response {
	this := ApiPuzzleActivity200Response{}
	return &this
}

// GetDate returns the Date field value
func (o *ApiPuzzleActivity200Response) GetDate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *ApiPuzzleActivity200Response) GetDateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *ApiPuzzleActivity200Response) SetDate(v int32) {
	o.Date = v
}

// GetPuzzle returns the Puzzle field value
func (o *ApiPuzzleActivity200Response) GetPuzzle() ApiPuzzleActivity200ResponsePuzzle {
	if o == nil {
		var ret ApiPuzzleActivity200ResponsePuzzle
		return ret
	}

	return o.Puzzle
}

// GetPuzzleOk returns a tuple with the Puzzle field value
// and a boolean to check if the value has been set.
func (o *ApiPuzzleActivity200Response) GetPuzzleOk() (*ApiPuzzleActivity200ResponsePuzzle, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Puzzle, true
}

// SetPuzzle sets field value
func (o *ApiPuzzleActivity200Response) SetPuzzle(v ApiPuzzleActivity200ResponsePuzzle) {
	o.Puzzle = v
}

// GetWin returns the Win field value
func (o *ApiPuzzleActivity200Response) GetWin() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Win
}

// GetWinOk returns a tuple with the Win field value
// and a boolean to check if the value has been set.
func (o *ApiPuzzleActivity200Response) GetWinOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Win, true
}

// SetWin sets field value
func (o *ApiPuzzleActivity200Response) SetWin(v bool) {
	o.Win = v
}

func (o ApiPuzzleActivity200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiPuzzleActivity200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["date"] = o.Date
	toSerialize["puzzle"] = o.Puzzle
	toSerialize["win"] = o.Win
	return toSerialize, nil
}

func (o *ApiPuzzleActivity200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"date",
		"puzzle",
		"win",
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

	varApiPuzzleActivity200Response := _ApiPuzzleActivity200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiPuzzleActivity200Response)

	if err != nil {
		return err
	}

	*o = ApiPuzzleActivity200Response(varApiPuzzleActivity200Response)

	return err
}

type NullableApiPuzzleActivity200Response struct {
	value *ApiPuzzleActivity200Response
	isSet bool
}

func (v NullableApiPuzzleActivity200Response) Get() *ApiPuzzleActivity200Response {
	return v.value
}

func (v *NullableApiPuzzleActivity200Response) Set(val *ApiPuzzleActivity200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableApiPuzzleActivity200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableApiPuzzleActivity200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiPuzzleActivity200Response(val *ApiPuzzleActivity200Response) *NullableApiPuzzleActivity200Response {
	return &NullableApiPuzzleActivity200Response{value: val, isSet: true}
}

func (v NullableApiPuzzleActivity200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiPuzzleActivity200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


