/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.158
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Top10s type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Top10s{}

// Top10s struct for Top10s
type Top10s struct {
	Bullet []TopUser `json:"bullet"`
	Blitz []TopUser `json:"blitz"`
	Rapid []TopUser `json:"rapid"`
	Classical []TopUser `json:"classical"`
	UltraBullet []TopUser `json:"ultraBullet"`
	Crazyhouse []TopUser `json:"crazyhouse"`
	Chess960 []TopUser `json:"chess960"`
	KingOfTheHill []TopUser `json:"kingOfTheHill"`
	ThreeCheck []TopUser `json:"threeCheck"`
	Antichess []TopUser `json:"antichess"`
	Atomic []TopUser `json:"atomic"`
	Horde []TopUser `json:"horde"`
	RacingKings []TopUser `json:"racingKings"`
}

type _Top10s Top10s

// NewTop10s instantiates a new Top10s object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTop10s(bullet []TopUser, blitz []TopUser, rapid []TopUser, classical []TopUser, ultraBullet []TopUser, crazyhouse []TopUser, chess960 []TopUser, kingOfTheHill []TopUser, threeCheck []TopUser, antichess []TopUser, atomic []TopUser, horde []TopUser, racingKings []TopUser) *Top10s {
	this := Top10s{}
	this.Bullet = bullet
	this.Blitz = blitz
	this.Rapid = rapid
	this.Classical = classical
	this.UltraBullet = ultraBullet
	this.Crazyhouse = crazyhouse
	this.Chess960 = chess960
	this.KingOfTheHill = kingOfTheHill
	this.ThreeCheck = threeCheck
	this.Antichess = antichess
	this.Atomic = atomic
	this.Horde = horde
	this.RacingKings = racingKings
	return &this
}

// NewTop10sWithDefaults instantiates a new Top10s object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTop10sWithDefaults() *Top10s {
	this := Top10s{}
	return &this
}

// GetBullet returns the Bullet field value
func (o *Top10s) GetBullet() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Bullet
}

// GetBulletOk returns a tuple with the Bullet field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetBulletOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Bullet, true
}

// SetBullet sets field value
func (o *Top10s) SetBullet(v []TopUser) {
	o.Bullet = v
}

// GetBlitz returns the Blitz field value
func (o *Top10s) GetBlitz() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Blitz
}

// GetBlitzOk returns a tuple with the Blitz field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetBlitzOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Blitz, true
}

// SetBlitz sets field value
func (o *Top10s) SetBlitz(v []TopUser) {
	o.Blitz = v
}

// GetRapid returns the Rapid field value
func (o *Top10s) GetRapid() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Rapid
}

// GetRapidOk returns a tuple with the Rapid field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetRapidOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rapid, true
}

// SetRapid sets field value
func (o *Top10s) SetRapid(v []TopUser) {
	o.Rapid = v
}

// GetClassical returns the Classical field value
func (o *Top10s) GetClassical() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Classical
}

// GetClassicalOk returns a tuple with the Classical field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetClassicalOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Classical, true
}

// SetClassical sets field value
func (o *Top10s) SetClassical(v []TopUser) {
	o.Classical = v
}

// GetUltraBullet returns the UltraBullet field value
func (o *Top10s) GetUltraBullet() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.UltraBullet
}

// GetUltraBulletOk returns a tuple with the UltraBullet field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetUltraBulletOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.UltraBullet, true
}

// SetUltraBullet sets field value
func (o *Top10s) SetUltraBullet(v []TopUser) {
	o.UltraBullet = v
}

// GetCrazyhouse returns the Crazyhouse field value
func (o *Top10s) GetCrazyhouse() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Crazyhouse
}

// GetCrazyhouseOk returns a tuple with the Crazyhouse field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetCrazyhouseOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Crazyhouse, true
}

// SetCrazyhouse sets field value
func (o *Top10s) SetCrazyhouse(v []TopUser) {
	o.Crazyhouse = v
}

// GetChess960 returns the Chess960 field value
func (o *Top10s) GetChess960() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Chess960
}

// GetChess960Ok returns a tuple with the Chess960 field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetChess960Ok() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Chess960, true
}

// SetChess960 sets field value
func (o *Top10s) SetChess960(v []TopUser) {
	o.Chess960 = v
}

// GetKingOfTheHill returns the KingOfTheHill field value
func (o *Top10s) GetKingOfTheHill() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.KingOfTheHill
}

// GetKingOfTheHillOk returns a tuple with the KingOfTheHill field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetKingOfTheHillOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.KingOfTheHill, true
}

// SetKingOfTheHill sets field value
func (o *Top10s) SetKingOfTheHill(v []TopUser) {
	o.KingOfTheHill = v
}

// GetThreeCheck returns the ThreeCheck field value
func (o *Top10s) GetThreeCheck() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.ThreeCheck
}

// GetThreeCheckOk returns a tuple with the ThreeCheck field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetThreeCheckOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.ThreeCheck, true
}

// SetThreeCheck sets field value
func (o *Top10s) SetThreeCheck(v []TopUser) {
	o.ThreeCheck = v
}

// GetAntichess returns the Antichess field value
func (o *Top10s) GetAntichess() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Antichess
}

// GetAntichessOk returns a tuple with the Antichess field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetAntichessOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Antichess, true
}

// SetAntichess sets field value
func (o *Top10s) SetAntichess(v []TopUser) {
	o.Antichess = v
}

// GetAtomic returns the Atomic field value
func (o *Top10s) GetAtomic() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetAtomicOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Atomic, true
}

// SetAtomic sets field value
func (o *Top10s) SetAtomic(v []TopUser) {
	o.Atomic = v
}

// GetHorde returns the Horde field value
func (o *Top10s) GetHorde() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.Horde
}

// GetHordeOk returns a tuple with the Horde field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetHordeOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Horde, true
}

// SetHorde sets field value
func (o *Top10s) SetHorde(v []TopUser) {
	o.Horde = v
}

// GetRacingKings returns the RacingKings field value
func (o *Top10s) GetRacingKings() []TopUser {
	if o == nil {
		var ret []TopUser
		return ret
	}

	return o.RacingKings
}

// GetRacingKingsOk returns a tuple with the RacingKings field value
// and a boolean to check if the value has been set.
func (o *Top10s) GetRacingKingsOk() ([]TopUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.RacingKings, true
}

// SetRacingKings sets field value
func (o *Top10s) SetRacingKings(v []TopUser) {
	o.RacingKings = v
}

func (o Top10s) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Top10s) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bullet"] = o.Bullet
	toSerialize["blitz"] = o.Blitz
	toSerialize["rapid"] = o.Rapid
	toSerialize["classical"] = o.Classical
	toSerialize["ultraBullet"] = o.UltraBullet
	toSerialize["crazyhouse"] = o.Crazyhouse
	toSerialize["chess960"] = o.Chess960
	toSerialize["kingOfTheHill"] = o.KingOfTheHill
	toSerialize["threeCheck"] = o.ThreeCheck
	toSerialize["antichess"] = o.Antichess
	toSerialize["atomic"] = o.Atomic
	toSerialize["horde"] = o.Horde
	toSerialize["racingKings"] = o.RacingKings
	return toSerialize, nil
}

func (o *Top10s) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bullet",
		"blitz",
		"rapid",
		"classical",
		"ultraBullet",
		"crazyhouse",
		"chess960",
		"kingOfTheHill",
		"threeCheck",
		"antichess",
		"atomic",
		"horde",
		"racingKings",
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

	varTop10s := _Top10s{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTop10s)

	if err != nil {
		return err
	}

	*o = Top10s(varTop10s)

	return err
}

type NullableTop10s struct {
	value *Top10s
	isSet bool
}

func (v NullableTop10s) Get() *Top10s {
	return v.value
}

func (v *NullableTop10s) Set(val *Top10s) {
	v.value = val
	v.isSet = true
}

func (v NullableTop10s) IsSet() bool {
	return v.isSet
}

func (v *NullableTop10s) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTop10s(val *Top10s) *NullableTop10s {
	return &NullableTop10s{value: val, isSet: true}
}

func (v NullableTop10s) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTop10s) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


