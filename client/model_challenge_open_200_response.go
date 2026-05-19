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

// checks if the ChallengeOpen200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChallengeOpen200Response{}

// ChallengeOpen200Response struct for ChallengeOpen200Response
type ChallengeOpen200Response struct {
	Id string `json:"id"`
	Url string `json:"url"`
	Status string `json:"status"`
	Challenger nil `json:"challenger"`
	DestUser nil `json:"destUser"`
	Variant ApiAccountPlaying200ResponseNowPlayingInnerVariant `json:"variant"`
	Rated bool `json:"rated"`
	Speed string `json:"speed"`
	TimeControl ApiStreamEvent200ResponseOneOf2ChallengeTimeControl `json:"timeControl"`
	// Which color you get to play
	Color string `json:"color"`
	// Color of the winner, if any
	FinalColor *string `json:"finalColor,omitempty"`
	Perf ChallengeOpen200ResponsePerf `json:"perf"`
	InitialFen *string `json:"initialFen,omitempty"`
	UrlWhite string `json:"urlWhite"`
	UrlBlack string `json:"urlBlack"`
	Open ChallengeOpen200ResponseOpen `json:"open"`
}

type _ChallengeOpen200Response ChallengeOpen200Response

// NewChallengeOpen200Response instantiates a new ChallengeOpen200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChallengeOpen200Response(id string, url string, status string, challenger nil, destUser nil, variant ApiAccountPlaying200ResponseNowPlayingInnerVariant, rated bool, speed string, timeControl ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, color string, perf ChallengeOpen200ResponsePerf, urlWhite string, urlBlack string, open ChallengeOpen200ResponseOpen) *ChallengeOpen200Response {
	this := ChallengeOpen200Response{}
	this.Id = id
	this.Url = url
	this.Status = status
	this.Challenger = challenger
	this.DestUser = destUser
	this.Variant = variant
	this.Rated = rated
	this.Speed = speed
	this.TimeControl = timeControl
	this.Color = color
	this.Perf = perf
	this.UrlWhite = urlWhite
	this.UrlBlack = urlBlack
	this.Open = open
	return &this
}

// NewChallengeOpen200ResponseWithDefaults instantiates a new ChallengeOpen200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChallengeOpen200ResponseWithDefaults() *ChallengeOpen200Response {
	this := ChallengeOpen200Response{}
	var color string = "random"
	this.Color = color
	return &this
}

// GetId returns the Id field value
func (o *ChallengeOpen200Response) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ChallengeOpen200Response) SetId(v string) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *ChallengeOpen200Response) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *ChallengeOpen200Response) SetUrl(v string) {
	o.Url = v
}

// GetStatus returns the Status field value
func (o *ChallengeOpen200Response) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ChallengeOpen200Response) SetStatus(v string) {
	o.Status = v
}

// GetChallenger returns the Challenger field value
func (o *ChallengeOpen200Response) GetChallenger() nil {
	if o == nil {
		var ret nil
		return ret
	}

	return o.Challenger
}

// GetChallengerOk returns a tuple with the Challenger field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetChallengerOk() (*nil, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Challenger, true
}

// SetChallenger sets field value
func (o *ChallengeOpen200Response) SetChallenger(v nil) {
	o.Challenger = v
}

// GetDestUser returns the DestUser field value
func (o *ChallengeOpen200Response) GetDestUser() nil {
	if o == nil {
		var ret nil
		return ret
	}

	return o.DestUser
}

// GetDestUserOk returns a tuple with the DestUser field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetDestUserOk() (*nil, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestUser, true
}

// SetDestUser sets field value
func (o *ChallengeOpen200Response) SetDestUser(v nil) {
	o.DestUser = v
}

// GetVariant returns the Variant field value
func (o *ChallengeOpen200Response) GetVariant() ApiAccountPlaying200ResponseNowPlayingInnerVariant {
	if o == nil {
		var ret ApiAccountPlaying200ResponseNowPlayingInnerVariant
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetVariantOk() (*ApiAccountPlaying200ResponseNowPlayingInnerVariant, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *ChallengeOpen200Response) SetVariant(v ApiAccountPlaying200ResponseNowPlayingInnerVariant) {
	o.Variant = v
}

// GetRated returns the Rated field value
func (o *ChallengeOpen200Response) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *ChallengeOpen200Response) SetRated(v bool) {
	o.Rated = v
}

// GetSpeed returns the Speed field value
func (o *ChallengeOpen200Response) GetSpeed() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Speed
}

// GetSpeedOk returns a tuple with the Speed field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetSpeedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Speed, true
}

// SetSpeed sets field value
func (o *ChallengeOpen200Response) SetSpeed(v string) {
	o.Speed = v
}

// GetTimeControl returns the TimeControl field value
func (o *ChallengeOpen200Response) GetTimeControl() ApiStreamEvent200ResponseOneOf2ChallengeTimeControl {
	if o == nil {
		var ret ApiStreamEvent200ResponseOneOf2ChallengeTimeControl
		return ret
	}

	return o.TimeControl
}

// GetTimeControlOk returns a tuple with the TimeControl field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetTimeControlOk() (*ApiStreamEvent200ResponseOneOf2ChallengeTimeControl, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimeControl, true
}

// SetTimeControl sets field value
func (o *ChallengeOpen200Response) SetTimeControl(v ApiStreamEvent200ResponseOneOf2ChallengeTimeControl) {
	o.TimeControl = v
}

// GetColor returns the Color field value
func (o *ChallengeOpen200Response) GetColor() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *ChallengeOpen200Response) SetColor(v string) {
	o.Color = v
}

// GetFinalColor returns the FinalColor field value if set, zero value otherwise.
func (o *ChallengeOpen200Response) GetFinalColor() string {
	if o == nil || IsNil(o.FinalColor) {
		var ret string
		return ret
	}
	return *o.FinalColor
}

// GetFinalColorOk returns a tuple with the FinalColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetFinalColorOk() (*string, bool) {
	if o == nil || IsNil(o.FinalColor) {
		return nil, false
	}
	return o.FinalColor, true
}

// HasFinalColor returns a boolean if a field has been set.
func (o *ChallengeOpen200Response) HasFinalColor() bool {
	if o != nil && !IsNil(o.FinalColor) {
		return true
	}

	return false
}

// SetFinalColor gets a reference to the given string and assigns it to the FinalColor field.
func (o *ChallengeOpen200Response) SetFinalColor(v string) {
	o.FinalColor = &v
}

// GetPerf returns the Perf field value
func (o *ChallengeOpen200Response) GetPerf() ChallengeOpen200ResponsePerf {
	if o == nil {
		var ret ChallengeOpen200ResponsePerf
		return ret
	}

	return o.Perf
}

// GetPerfOk returns a tuple with the Perf field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetPerfOk() (*ChallengeOpen200ResponsePerf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Perf, true
}

// SetPerf sets field value
func (o *ChallengeOpen200Response) SetPerf(v ChallengeOpen200ResponsePerf) {
	o.Perf = v
}

// GetInitialFen returns the InitialFen field value if set, zero value otherwise.
func (o *ChallengeOpen200Response) GetInitialFen() string {
	if o == nil || IsNil(o.InitialFen) {
		var ret string
		return ret
	}
	return *o.InitialFen
}

// GetInitialFenOk returns a tuple with the InitialFen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetInitialFenOk() (*string, bool) {
	if o == nil || IsNil(o.InitialFen) {
		return nil, false
	}
	return o.InitialFen, true
}

// HasInitialFen returns a boolean if a field has been set.
func (o *ChallengeOpen200Response) HasInitialFen() bool {
	if o != nil && !IsNil(o.InitialFen) {
		return true
	}

	return false
}

// SetInitialFen gets a reference to the given string and assigns it to the InitialFen field.
func (o *ChallengeOpen200Response) SetInitialFen(v string) {
	o.InitialFen = &v
}

// GetUrlWhite returns the UrlWhite field value
func (o *ChallengeOpen200Response) GetUrlWhite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlWhite
}

// GetUrlWhiteOk returns a tuple with the UrlWhite field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetUrlWhiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlWhite, true
}

// SetUrlWhite sets field value
func (o *ChallengeOpen200Response) SetUrlWhite(v string) {
	o.UrlWhite = v
}

// GetUrlBlack returns the UrlBlack field value
func (o *ChallengeOpen200Response) GetUrlBlack() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UrlBlack
}

// GetUrlBlackOk returns a tuple with the UrlBlack field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetUrlBlackOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UrlBlack, true
}

// SetUrlBlack sets field value
func (o *ChallengeOpen200Response) SetUrlBlack(v string) {
	o.UrlBlack = v
}

// GetOpen returns the Open field value
func (o *ChallengeOpen200Response) GetOpen() ChallengeOpen200ResponseOpen {
	if o == nil {
		var ret ChallengeOpen200ResponseOpen
		return ret
	}

	return o.Open
}

// GetOpenOk returns a tuple with the Open field value
// and a boolean to check if the value has been set.
func (o *ChallengeOpen200Response) GetOpenOk() (*ChallengeOpen200ResponseOpen, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Open, true
}

// SetOpen sets field value
func (o *ChallengeOpen200Response) SetOpen(v ChallengeOpen200ResponseOpen) {
	o.Open = v
}

func (o ChallengeOpen200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChallengeOpen200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["status"] = o.Status
	toSerialize["challenger"] = o.Challenger
	toSerialize["destUser"] = o.DestUser
	toSerialize["variant"] = o.Variant
	toSerialize["rated"] = o.Rated
	toSerialize["speed"] = o.Speed
	toSerialize["timeControl"] = o.TimeControl
	toSerialize["color"] = o.Color
	if !IsNil(o.FinalColor) {
		toSerialize["finalColor"] = o.FinalColor
	}
	toSerialize["perf"] = o.Perf
	if !IsNil(o.InitialFen) {
		toSerialize["initialFen"] = o.InitialFen
	}
	toSerialize["urlWhite"] = o.UrlWhite
	toSerialize["urlBlack"] = o.UrlBlack
	toSerialize["open"] = o.Open
	return toSerialize, nil
}

func (o *ChallengeOpen200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"status",
		"challenger",
		"destUser",
		"variant",
		"rated",
		"speed",
		"timeControl",
		"color",
		"perf",
		"urlWhite",
		"urlBlack",
		"open",
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

	varChallengeOpen200Response := _ChallengeOpen200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChallengeOpen200Response)

	if err != nil {
		return err
	}

	*o = ChallengeOpen200Response(varChallengeOpen200Response)

	return err
}

type NullableChallengeOpen200Response struct {
	value *ChallengeOpen200Response
	isSet bool
}

func (v NullableChallengeOpen200Response) Get() *ChallengeOpen200Response {
	return v.value
}

func (v *NullableChallengeOpen200Response) Set(val *ChallengeOpen200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableChallengeOpen200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableChallengeOpen200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChallengeOpen200Response(val *ChallengeOpen200Response) *NullableChallengeOpen200Response {
	return &NullableChallengeOpen200Response{value: val, isSet: true}
}

func (v NullableChallengeOpen200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChallengeOpen200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


