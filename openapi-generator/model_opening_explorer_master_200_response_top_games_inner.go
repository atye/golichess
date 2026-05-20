/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the OpeningExplorerMaster200ResponseTopGamesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OpeningExplorerMaster200ResponseTopGamesInner{}

// OpeningExplorerMaster200ResponseTopGamesInner struct for OpeningExplorerMaster200ResponseTopGamesInner
type OpeningExplorerMaster200ResponseTopGamesInner struct {
	Uci string `json:"uci"`
	Id string `json:"id"`
	Winner NullableString `json:"winner"`
	White OpeningExplorerMaster200ResponseMovesInnerGameWhite `json:"white"`
	Black OpeningExplorerMaster200ResponseMovesInnerGameWhite `json:"black"`
	Year int32 `json:"year"`
	Month *string `json:"month,omitempty"`
}

type _OpeningExplorerMaster200ResponseTopGamesInner OpeningExplorerMaster200ResponseTopGamesInner

// NewOpeningExplorerMaster200ResponseTopGamesInner instantiates a new OpeningExplorerMaster200ResponseTopGamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOpeningExplorerMaster200ResponseTopGamesInner(uci string, id string, winner NullableString, white OpeningExplorerMaster200ResponseMovesInnerGameWhite, black OpeningExplorerMaster200ResponseMovesInnerGameWhite, year int32) *OpeningExplorerMaster200ResponseTopGamesInner {
	this := OpeningExplorerMaster200ResponseTopGamesInner{}
	this.Uci = uci
	this.Id = id
	this.Winner = winner
	this.White = white
	this.Black = black
	this.Year = year
	return &this
}

// NewOpeningExplorerMaster200ResponseTopGamesInnerWithDefaults instantiates a new OpeningExplorerMaster200ResponseTopGamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOpeningExplorerMaster200ResponseTopGamesInnerWithDefaults() *OpeningExplorerMaster200ResponseTopGamesInner {
	this := OpeningExplorerMaster200ResponseTopGamesInner{}
	return &this
}

// GetUci returns the Uci field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetUci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uci
}

// GetUciOk returns a tuple with the Uci field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetUciOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uci, true
}

// SetUci sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetUci(v string) {
	o.Uci = v
}

// GetId returns the Id field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetId(v string) {
	o.Id = v
}

// GetWinner returns the Winner field value
// If the value is explicit nil, the zero value for string will be returned
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetWinner() string {
	if o == nil || o.Winner.Get() == nil {
		var ret string
		return ret
	}

	return *o.Winner.Get()
}

// GetWinnerOk returns a tuple with the Winner field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetWinnerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Winner.Get(), o.Winner.IsSet()
}

// SetWinner sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetWinner(v string) {
	o.Winner.Set(&v)
}

// GetWhite returns the White field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetWhite() OpeningExplorerMaster200ResponseMovesInnerGameWhite {
	if o == nil {
		var ret OpeningExplorerMaster200ResponseMovesInnerGameWhite
		return ret
	}

	return o.White
}

// GetWhiteOk returns a tuple with the White field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetWhiteOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.White, true
}

// SetWhite sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetWhite(v OpeningExplorerMaster200ResponseMovesInnerGameWhite) {
	o.White = v
}

// GetBlack returns the Black field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetBlack() OpeningExplorerMaster200ResponseMovesInnerGameWhite {
	if o == nil {
		var ret OpeningExplorerMaster200ResponseMovesInnerGameWhite
		return ret
	}

	return o.Black
}

// GetBlackOk returns a tuple with the Black field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetBlackOk() (*OpeningExplorerMaster200ResponseMovesInnerGameWhite, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Black, true
}

// SetBlack sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetBlack(v OpeningExplorerMaster200ResponseMovesInnerGameWhite) {
	o.Black = v
}

// GetYear returns the Year field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetYear() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Year
}

// GetYearOk returns a tuple with the Year field value
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Year, true
}

// SetYear sets field value
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetYear(v int32) {
	o.Year = v
}

// GetMonth returns the Month field value if set, zero value otherwise.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetMonth() string {
	if o == nil || IsNil(o.Month) {
		var ret string
		return ret
	}
	return *o.Month
}

// GetMonthOk returns a tuple with the Month field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) GetMonthOk() (*string, bool) {
	if o == nil || IsNil(o.Month) {
		return nil, false
	}
	return o.Month, true
}

// HasMonth returns a boolean if a field has been set.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) HasMonth() bool {
	if o != nil && !IsNil(o.Month) {
		return true
	}

	return false
}

// SetMonth gets a reference to the given string and assigns it to the Month field.
func (o *OpeningExplorerMaster200ResponseTopGamesInner) SetMonth(v string) {
	o.Month = &v
}

func (o OpeningExplorerMaster200ResponseTopGamesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OpeningExplorerMaster200ResponseTopGamesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uci"] = o.Uci
	toSerialize["id"] = o.Id
	toSerialize["winner"] = o.Winner.Get()
	toSerialize["white"] = o.White
	toSerialize["black"] = o.Black
	toSerialize["year"] = o.Year
	if !IsNil(o.Month) {
		toSerialize["month"] = o.Month
	}
	return toSerialize, nil
}

func (o *OpeningExplorerMaster200ResponseTopGamesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uci",
		"id",
		"winner",
		"white",
		"black",
		"year",
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

	varOpeningExplorerMaster200ResponseTopGamesInner := _OpeningExplorerMaster200ResponseTopGamesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOpeningExplorerMaster200ResponseTopGamesInner)

	if err != nil {
		return err
	}

	*o = OpeningExplorerMaster200ResponseTopGamesInner(varOpeningExplorerMaster200ResponseTopGamesInner)

	return err
}

type NullableOpeningExplorerMaster200ResponseTopGamesInner struct {
	value *OpeningExplorerMaster200ResponseTopGamesInner
	isSet bool
}

func (v NullableOpeningExplorerMaster200ResponseTopGamesInner) Get() *OpeningExplorerMaster200ResponseTopGamesInner {
	return v.value
}

func (v *NullableOpeningExplorerMaster200ResponseTopGamesInner) Set(val *OpeningExplorerMaster200ResponseTopGamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOpeningExplorerMaster200ResponseTopGamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOpeningExplorerMaster200ResponseTopGamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOpeningExplorerMaster200ResponseTopGamesInner(val *OpeningExplorerMaster200ResponseTopGamesInner) *NullableOpeningExplorerMaster200ResponseTopGamesInner {
	return &NullableOpeningExplorerMaster200ResponseTopGamesInner{value: val, isSet: true}
}

func (v NullableOpeningExplorerMaster200ResponseTopGamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOpeningExplorerMaster200ResponseTopGamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


