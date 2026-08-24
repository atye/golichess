/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.166
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Count type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Count{}

// Count struct for Count
type Count struct {
	All int32 `json:"all"`
	Rated int32 `json:"rated"`
	Ai *int32 `json:"ai,omitempty"`
	Draw int32 `json:"draw"`
	DrawH *int32 `json:"drawH,omitempty"`
	Loss int32 `json:"loss"`
	LossH *int32 `json:"lossH,omitempty"`
	Win int32 `json:"win"`
	WinH *int32 `json:"winH,omitempty"`
	Bookmark int32 `json:"bookmark"`
	Playing int32 `json:"playing"`
	Import int32 `json:"import"`
	Me int32 `json:"me"`
}

type _Count Count

// NewCount instantiates a new Count object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCount(all int32, rated int32, draw int32, loss int32, win int32, bookmark int32, playing int32, import_ int32, me int32) *Count {
	this := Count{}
	this.All = all
	this.Rated = rated
	this.Draw = draw
	this.Loss = loss
	this.Win = win
	this.Bookmark = bookmark
	this.Playing = playing
	this.Import = import_
	this.Me = me
	return &this
}

// NewCountWithDefaults instantiates a new Count object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountWithDefaults() *Count {
	this := Count{}
	return &this
}

// GetAll returns the All field value
func (o *Count) GetAll() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.All
}

// GetAllOk returns a tuple with the All field value
// and a boolean to check if the value has been set.
func (o *Count) GetAllOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.All, true
}

// SetAll sets field value
func (o *Count) SetAll(v int32) {
	o.All = v
}

// GetRated returns the Rated field value
func (o *Count) GetRated() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *Count) GetRatedOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *Count) SetRated(v int32) {
	o.Rated = v
}

// GetAi returns the Ai field value if set, zero value otherwise.
func (o *Count) GetAi() int32 {
	if o == nil || IsNil(o.Ai) {
		var ret int32
		return ret
	}
	return *o.Ai
}

// GetAiOk returns a tuple with the Ai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Count) GetAiOk() (*int32, bool) {
	if o == nil || IsNil(o.Ai) {
		return nil, false
	}
	return o.Ai, true
}

// HasAi returns a boolean if a field has been set.
func (o *Count) HasAi() bool {
	if o != nil && !IsNil(o.Ai) {
		return true
	}

	return false
}

// SetAi gets a reference to the given int32 and assigns it to the Ai field.
func (o *Count) SetAi(v int32) {
	o.Ai = &v
}

// GetDraw returns the Draw field value
func (o *Count) GetDraw() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Draw
}

// GetDrawOk returns a tuple with the Draw field value
// and a boolean to check if the value has been set.
func (o *Count) GetDrawOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Draw, true
}

// SetDraw sets field value
func (o *Count) SetDraw(v int32) {
	o.Draw = v
}

// GetDrawH returns the DrawH field value if set, zero value otherwise.
func (o *Count) GetDrawH() int32 {
	if o == nil || IsNil(o.DrawH) {
		var ret int32
		return ret
	}
	return *o.DrawH
}

// GetDrawHOk returns a tuple with the DrawH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Count) GetDrawHOk() (*int32, bool) {
	if o == nil || IsNil(o.DrawH) {
		return nil, false
	}
	return o.DrawH, true
}

// HasDrawH returns a boolean if a field has been set.
func (o *Count) HasDrawH() bool {
	if o != nil && !IsNil(o.DrawH) {
		return true
	}

	return false
}

// SetDrawH gets a reference to the given int32 and assigns it to the DrawH field.
func (o *Count) SetDrawH(v int32) {
	o.DrawH = &v
}

// GetLoss returns the Loss field value
func (o *Count) GetLoss() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Loss
}

// GetLossOk returns a tuple with the Loss field value
// and a boolean to check if the value has been set.
func (o *Count) GetLossOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Loss, true
}

// SetLoss sets field value
func (o *Count) SetLoss(v int32) {
	o.Loss = v
}

// GetLossH returns the LossH field value if set, zero value otherwise.
func (o *Count) GetLossH() int32 {
	if o == nil || IsNil(o.LossH) {
		var ret int32
		return ret
	}
	return *o.LossH
}

// GetLossHOk returns a tuple with the LossH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Count) GetLossHOk() (*int32, bool) {
	if o == nil || IsNil(o.LossH) {
		return nil, false
	}
	return o.LossH, true
}

// HasLossH returns a boolean if a field has been set.
func (o *Count) HasLossH() bool {
	if o != nil && !IsNil(o.LossH) {
		return true
	}

	return false
}

// SetLossH gets a reference to the given int32 and assigns it to the LossH field.
func (o *Count) SetLossH(v int32) {
	o.LossH = &v
}

// GetWin returns the Win field value
func (o *Count) GetWin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Win
}

// GetWinOk returns a tuple with the Win field value
// and a boolean to check if the value has been set.
func (o *Count) GetWinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Win, true
}

// SetWin sets field value
func (o *Count) SetWin(v int32) {
	o.Win = v
}

// GetWinH returns the WinH field value if set, zero value otherwise.
func (o *Count) GetWinH() int32 {
	if o == nil || IsNil(o.WinH) {
		var ret int32
		return ret
	}
	return *o.WinH
}

// GetWinHOk returns a tuple with the WinH field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Count) GetWinHOk() (*int32, bool) {
	if o == nil || IsNil(o.WinH) {
		return nil, false
	}
	return o.WinH, true
}

// HasWinH returns a boolean if a field has been set.
func (o *Count) HasWinH() bool {
	if o != nil && !IsNil(o.WinH) {
		return true
	}

	return false
}

// SetWinH gets a reference to the given int32 and assigns it to the WinH field.
func (o *Count) SetWinH(v int32) {
	o.WinH = &v
}

// GetBookmark returns the Bookmark field value
func (o *Count) GetBookmark() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Bookmark
}

// GetBookmarkOk returns a tuple with the Bookmark field value
// and a boolean to check if the value has been set.
func (o *Count) GetBookmarkOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bookmark, true
}

// SetBookmark sets field value
func (o *Count) SetBookmark(v int32) {
	o.Bookmark = v
}

// GetPlaying returns the Playing field value
func (o *Count) GetPlaying() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Playing
}

// GetPlayingOk returns a tuple with the Playing field value
// and a boolean to check if the value has been set.
func (o *Count) GetPlayingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Playing, true
}

// SetPlaying sets field value
func (o *Count) SetPlaying(v int32) {
	o.Playing = v
}

// GetImport returns the Import field value
func (o *Count) GetImport() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Import
}

// GetImportOk returns a tuple with the Import field value
// and a boolean to check if the value has been set.
func (o *Count) GetImportOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Import, true
}

// SetImport sets field value
func (o *Count) SetImport(v int32) {
	o.Import = v
}

// GetMe returns the Me field value
func (o *Count) GetMe() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Me
}

// GetMeOk returns a tuple with the Me field value
// and a boolean to check if the value has been set.
func (o *Count) GetMeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Me, true
}

// SetMe sets field value
func (o *Count) SetMe(v int32) {
	o.Me = v
}

func (o Count) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Count) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["all"] = o.All
	toSerialize["rated"] = o.Rated
	if !IsNil(o.Ai) {
		toSerialize["ai"] = o.Ai
	}
	toSerialize["draw"] = o.Draw
	if !IsNil(o.DrawH) {
		toSerialize["drawH"] = o.DrawH
	}
	toSerialize["loss"] = o.Loss
	if !IsNil(o.LossH) {
		toSerialize["lossH"] = o.LossH
	}
	toSerialize["win"] = o.Win
	if !IsNil(o.WinH) {
		toSerialize["winH"] = o.WinH
	}
	toSerialize["bookmark"] = o.Bookmark
	toSerialize["playing"] = o.Playing
	toSerialize["import"] = o.Import
	toSerialize["me"] = o.Me
	return toSerialize, nil
}

func (o *Count) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"all",
		"rated",
		"draw",
		"loss",
		"win",
		"bookmark",
		"playing",
		"import",
		"me",
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

	varCount := _Count{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCount)

	if err != nil {
		return err
	}

	*o = Count(varCount)

	return err
}

type NullableCount struct {
	value *Count
	isSet bool
}

func (v NullableCount) Get() *Count {
	return v.value
}

func (v *NullableCount) Set(val *Count) {
	v.value = val
	v.isSet = true
}

func (v NullableCount) IsSet() bool {
	return v.isSet
}

func (v *NullableCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCount(val *Count) *NullableCount {
	return &NullableCount{value: val, isSet: true}
}

func (v NullableCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


