/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.153
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserActivityCorrespondenceGame type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActivityCorrespondenceGame{}

// UserActivityCorrespondenceGame struct for UserActivityCorrespondenceGame
type UserActivityCorrespondenceGame struct {
	Id string `json:"id"`
	Color GameColor `json:"color"`
	Url string `json:"url"`
	Variant *VariantKey `json:"variant,omitempty"`
	Speed *string `json:"speed,omitempty"`
	Perf *string `json:"perf,omitempty"`
	Rated *bool `json:"rated,omitempty"`
	Opponent UserActivityCorrespondenceGameOpponent `json:"opponent"`
}

type _UserActivityCorrespondenceGame UserActivityCorrespondenceGame

// NewUserActivityCorrespondenceGame instantiates a new UserActivityCorrespondenceGame object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivityCorrespondenceGame(id string, color GameColor, url string, opponent UserActivityCorrespondenceGameOpponent) *UserActivityCorrespondenceGame {
	this := UserActivityCorrespondenceGame{}
	this.Id = id
	this.Color = color
	this.Url = url
	var variant VariantKey = VARIANTKEY_STANDARD
	this.Variant = &variant
	this.Opponent = opponent
	return &this
}

// NewUserActivityCorrespondenceGameWithDefaults instantiates a new UserActivityCorrespondenceGame object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityCorrespondenceGameWithDefaults() *UserActivityCorrespondenceGame {
	this := UserActivityCorrespondenceGame{}
	var variant VariantKey = VARIANTKEY_STANDARD
	this.Variant = &variant
	return &this
}

// GetId returns the Id field value
func (o *UserActivityCorrespondenceGame) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserActivityCorrespondenceGame) SetId(v string) {
	o.Id = v
}

// GetColor returns the Color field value
func (o *UserActivityCorrespondenceGame) GetColor() GameColor {
	if o == nil {
		var ret GameColor
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetColorOk() (*GameColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *UserActivityCorrespondenceGame) SetColor(v GameColor) {
	o.Color = v
}

// GetUrl returns the Url field value
func (o *UserActivityCorrespondenceGame) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UserActivityCorrespondenceGame) SetUrl(v string) {
	o.Url = v
}

// GetVariant returns the Variant field value if set, zero value otherwise.
func (o *UserActivityCorrespondenceGame) GetVariant() VariantKey {
	if o == nil || IsNil(o.Variant) {
		var ret VariantKey
		return ret
	}
	return *o.Variant
}

// GetVariantOk returns a tuple with the Variant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetVariantOk() (*VariantKey, bool) {
	if o == nil || IsNil(o.Variant) {
		return nil, false
	}
	return o.Variant, true
}

// HasVariant returns a boolean if a field has been set.
func (o *UserActivityCorrespondenceGame) HasVariant() bool {
	if o != nil && !IsNil(o.Variant) {
		return true
	}

	return false
}

// SetVariant gets a reference to the given VariantKey and assigns it to the Variant field.
func (o *UserActivityCorrespondenceGame) SetVariant(v VariantKey) {
	o.Variant = &v
}

// GetSpeed returns the Speed field value if set, zero value otherwise.
func (o *UserActivityCorrespondenceGame) GetSpeed() string {
	if o == nil || IsNil(o.Speed) {
		var ret string
		return ret
	}
	return *o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetSpeedOk() (*string, bool) {
	if o == nil || IsNil(o.Speed) {
		return nil, false
	}
	return o.Speed, true
}

// HasSpeed returns a boolean if a field has been set.
func (o *UserActivityCorrespondenceGame) HasSpeed() bool {
	if o != nil && !IsNil(o.Speed) {
		return true
	}

	return false
}

// SetSpeed gets a reference to the given string and assigns it to the Speed field.
func (o *UserActivityCorrespondenceGame) SetSpeed(v string) {
	o.Speed = &v
}

// GetPerf returns the Perf field value if set, zero value otherwise.
func (o *UserActivityCorrespondenceGame) GetPerf() string {
	if o == nil || IsNil(o.Perf) {
		var ret string
		return ret
	}
	return *o.Perf
}

// GetPerfOk returns a tuple with the Perf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetPerfOk() (*string, bool) {
	if o == nil || IsNil(o.Perf) {
		return nil, false
	}
	return o.Perf, true
}

// HasPerf returns a boolean if a field has been set.
func (o *UserActivityCorrespondenceGame) HasPerf() bool {
	if o != nil && !IsNil(o.Perf) {
		return true
	}

	return false
}

// SetPerf gets a reference to the given string and assigns it to the Perf field.
func (o *UserActivityCorrespondenceGame) SetPerf(v string) {
	o.Perf = &v
}

// GetRated returns the Rated field value if set, zero value otherwise.
func (o *UserActivityCorrespondenceGame) GetRated() bool {
	if o == nil || IsNil(o.Rated) {
		var ret bool
		return ret
	}
	return *o.Rated
}

// GetRatedOk returns a tuple with the Rated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetRatedOk() (*bool, bool) {
	if o == nil || IsNil(o.Rated) {
		return nil, false
	}
	return o.Rated, true
}

// HasRated returns a boolean if a field has been set.
func (o *UserActivityCorrespondenceGame) HasRated() bool {
	if o != nil && !IsNil(o.Rated) {
		return true
	}

	return false
}

// SetRated gets a reference to the given bool and assigns it to the Rated field.
func (o *UserActivityCorrespondenceGame) SetRated(v bool) {
	o.Rated = &v
}

// GetOpponent returns the Opponent field value
func (o *UserActivityCorrespondenceGame) GetOpponent() UserActivityCorrespondenceGameOpponent {
	if o == nil {
		var ret UserActivityCorrespondenceGameOpponent
		return ret
	}

	return o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value
// and a boolean to check if the value has been set.
func (o *UserActivityCorrespondenceGame) GetOpponentOk() (*UserActivityCorrespondenceGameOpponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Opponent, true
}

// SetOpponent sets field value
func (o *UserActivityCorrespondenceGame) SetOpponent(v UserActivityCorrespondenceGameOpponent) {
	o.Opponent = v
}

func (o UserActivityCorrespondenceGame) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActivityCorrespondenceGame) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["color"] = o.Color
	toSerialize["url"] = o.Url
	if !IsNil(o.Variant) {
		toSerialize["variant"] = o.Variant
	}
	if !IsNil(o.Speed) {
		toSerialize["speed"] = o.Speed
	}
	if !IsNil(o.Perf) {
		toSerialize["perf"] = o.Perf
	}
	if !IsNil(o.Rated) {
		toSerialize["rated"] = o.Rated
	}
	toSerialize["opponent"] = o.Opponent
	return toSerialize, nil
}

func (o *UserActivityCorrespondenceGame) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"color",
		"url",
		"opponent",
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

	varUserActivityCorrespondenceGame := _UserActivityCorrespondenceGame{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserActivityCorrespondenceGame)

	if err != nil {
		return err
	}

	*o = UserActivityCorrespondenceGame(varUserActivityCorrespondenceGame)

	return err
}

type NullableUserActivityCorrespondenceGame struct {
	value *UserActivityCorrespondenceGame
	isSet bool
}

func (v NullableUserActivityCorrespondenceGame) Get() *UserActivityCorrespondenceGame {
	return v.value
}

func (v *NullableUserActivityCorrespondenceGame) Set(val *UserActivityCorrespondenceGame) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivityCorrespondenceGame) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivityCorrespondenceGame) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivityCorrespondenceGame(val *UserActivityCorrespondenceGame) *NullableUserActivityCorrespondenceGame {
	return &NullableUserActivityCorrespondenceGame{value: val, isSet: true}
}

func (v NullableUserActivityCorrespondenceGame) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivityCorrespondenceGame) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


