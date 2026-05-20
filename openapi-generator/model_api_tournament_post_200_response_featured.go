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
)

// checks if the ApiTournamentPost200ResponseFeatured type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiTournamentPost200ResponseFeatured{}

// ApiTournamentPost200ResponseFeatured struct for ApiTournamentPost200ResponseFeatured
type ApiTournamentPost200ResponseFeatured struct {
	Id *string `json:"id,omitempty"`
	Fen *string `json:"fen,omitempty"`
	Orientation *string `json:"orientation,omitempty"`
	Color *string `json:"color,omitempty"`
	LastMove *string `json:"lastMove,omitempty"`
	White *ApiTournamentPost200ResponseFeaturedWhite `json:"white,omitempty"`
	Black *ApiTournamentPost200ResponseFeaturedWhite `json:"black,omitempty"`
	C *ApiTournamentPost200ResponseFeaturedC `json:"c,omitempty"`
}

// NewApiTournamentPost200ResponseFeatured instantiates a new ApiTournamentPost200ResponseFeatured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiTournamentPost200ResponseFeatured() *ApiTournamentPost200ResponseFeatured {
	this := ApiTournamentPost200ResponseFeatured{}
	return &this
}

// NewApiTournamentPost200ResponseFeaturedWithDefaults instantiates a new ApiTournamentPost200ResponseFeatured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiTournamentPost200ResponseFeaturedWithDefaults() *ApiTournamentPost200ResponseFeatured {
	this := ApiTournamentPost200ResponseFeatured{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ApiTournamentPost200ResponseFeatured) SetId(v string) {
	o.Id = &v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *ApiTournamentPost200ResponseFeatured) SetFen(v string) {
	o.Fen = &v
}

// GetOrientation returns the Orientation field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetOrientation() string {
	if o == nil || IsNil(o.Orientation) {
		var ret string
		return ret
	}
	return *o.Orientation
}

// GetOrientationOk returns a tuple with the Orientation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetOrientationOk() (*string, bool) {
	if o == nil || IsNil(o.Orientation) {
		return nil, false
	}
	return o.Orientation, true
}

// HasOrientation returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasOrientation() bool {
	if o != nil && !IsNil(o.Orientation) {
		return true
	}

	return false
}

// SetOrientation gets a reference to the given string and assigns it to the Orientation field.
func (o *ApiTournamentPost200ResponseFeatured) SetOrientation(v string) {
	o.Orientation = &v
}

// GetColor returns the Color field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetColor() string {
	if o == nil || IsNil(o.Color) {
		var ret string
		return ret
	}
	return *o.Color
}

// GetColorOk returns a tuple with the Color field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetColorOk() (*string, bool) {
	if o == nil || IsNil(o.Color) {
		return nil, false
	}
	return o.Color, true
}

// HasColor returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasColor() bool {
	if o != nil && !IsNil(o.Color) {
		return true
	}

	return false
}

// SetColor gets a reference to the given string and assigns it to the Color field.
func (o *ApiTournamentPost200ResponseFeatured) SetColor(v string) {
	o.Color = &v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *ApiTournamentPost200ResponseFeatured) SetLastMove(v string) {
	o.LastMove = &v
}

// GetWhite returns the White field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetWhite() ApiTournamentPost200ResponseFeaturedWhite {
	if o == nil || IsNil(o.White) {
		var ret ApiTournamentPost200ResponseFeaturedWhite
		return ret
	}
	return *o.White
}

// GetWhiteOk returns a tuple with the White field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetWhiteOk() (*ApiTournamentPost200ResponseFeaturedWhite, bool) {
	if o == nil || IsNil(o.White) {
		return nil, false
	}
	return o.White, true
}

// HasWhite returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasWhite() bool {
	if o != nil && !IsNil(o.White) {
		return true
	}

	return false
}

// SetWhite gets a reference to the given ApiTournamentPost200ResponseFeaturedWhite and assigns it to the White field.
func (o *ApiTournamentPost200ResponseFeatured) SetWhite(v ApiTournamentPost200ResponseFeaturedWhite) {
	o.White = &v
}

// GetBlack returns the Black field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetBlack() ApiTournamentPost200ResponseFeaturedWhite {
	if o == nil || IsNil(o.Black) {
		var ret ApiTournamentPost200ResponseFeaturedWhite
		return ret
	}
	return *o.Black
}

// GetBlackOk returns a tuple with the Black field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetBlackOk() (*ApiTournamentPost200ResponseFeaturedWhite, bool) {
	if o == nil || IsNil(o.Black) {
		return nil, false
	}
	return o.Black, true
}

// HasBlack returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasBlack() bool {
	if o != nil && !IsNil(o.Black) {
		return true
	}

	return false
}

// SetBlack gets a reference to the given ApiTournamentPost200ResponseFeaturedWhite and assigns it to the Black field.
func (o *ApiTournamentPost200ResponseFeatured) SetBlack(v ApiTournamentPost200ResponseFeaturedWhite) {
	o.Black = &v
}

// GetC returns the C field value if set, zero value otherwise.
func (o *ApiTournamentPost200ResponseFeatured) GetC() ApiTournamentPost200ResponseFeaturedC {
	if o == nil || IsNil(o.C) {
		var ret ApiTournamentPost200ResponseFeaturedC
		return ret
	}
	return *o.C
}

// GetCOk returns a tuple with the C field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiTournamentPost200ResponseFeatured) GetCOk() (*ApiTournamentPost200ResponseFeaturedC, bool) {
	if o == nil || IsNil(o.C) {
		return nil, false
	}
	return o.C, true
}

// HasC returns a boolean if a field has been set.
func (o *ApiTournamentPost200ResponseFeatured) HasC() bool {
	if o != nil && !IsNil(o.C) {
		return true
	}

	return false
}

// SetC gets a reference to the given ApiTournamentPost200ResponseFeaturedC and assigns it to the C field.
func (o *ApiTournamentPost200ResponseFeatured) SetC(v ApiTournamentPost200ResponseFeaturedC) {
	o.C = &v
}

func (o ApiTournamentPost200ResponseFeatured) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiTournamentPost200ResponseFeatured) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if !IsNil(o.Orientation) {
		toSerialize["orientation"] = o.Orientation
	}
	if !IsNil(o.Color) {
		toSerialize["color"] = o.Color
	}
	if !IsNil(o.LastMove) {
		toSerialize["lastMove"] = o.LastMove
	}
	if !IsNil(o.White) {
		toSerialize["white"] = o.White
	}
	if !IsNil(o.Black) {
		toSerialize["black"] = o.Black
	}
	if !IsNil(o.C) {
		toSerialize["c"] = o.C
	}
	return toSerialize, nil
}

type NullableApiTournamentPost200ResponseFeatured struct {
	value *ApiTournamentPost200ResponseFeatured
	isSet bool
}

func (v NullableApiTournamentPost200ResponseFeatured) Get() *ApiTournamentPost200ResponseFeatured {
	return v.value
}

func (v *NullableApiTournamentPost200ResponseFeatured) Set(val *ApiTournamentPost200ResponseFeatured) {
	v.value = val
	v.isSet = true
}

func (v NullableApiTournamentPost200ResponseFeatured) IsSet() bool {
	return v.isSet
}

func (v *NullableApiTournamentPost200ResponseFeatured) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiTournamentPost200ResponseFeatured(val *ApiTournamentPost200ResponseFeatured) *NullableApiTournamentPost200ResponseFeatured {
	return &NullableApiTournamentPost200ResponseFeatured{value: val, isSet: true}
}

func (v NullableApiTournamentPost200ResponseFeatured) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiTournamentPost200ResponseFeatured) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


