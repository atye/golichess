/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.161
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TvChannels200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TvChannels200Response{}

// TvChannels200Response struct for TvChannels200Response
type TvChannels200Response struct {
	Bot TvGame `json:"bot"`
	Blitz TvGame `json:"blitz"`
	RacingKings TvGame `json:"racingKings"`
	UltraBullet TvGame `json:"ultraBullet"`
	Bullet TvGame `json:"bullet"`
	Classical TvGame `json:"classical"`
	ThreeCheck TvGame `json:"threeCheck"`
	Antichess TvGame `json:"antichess"`
	Computer TvGame `json:"computer"`
	Horde TvGame `json:"horde"`
	Rapid TvGame `json:"rapid"`
	Atomic TvGame `json:"atomic"`
	Crazyhouse TvGame `json:"crazyhouse"`
	Chess960 TvGame `json:"chess960"`
	KingOfTheHill TvGame `json:"kingOfTheHill"`
	Best TvGame `json:"best"`
}

type _TvChannels200Response TvChannels200Response

// NewTvChannels200Response instantiates a new TvChannels200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTvChannels200Response(bot TvGame, blitz TvGame, racingKings TvGame, ultraBullet TvGame, bullet TvGame, classical TvGame, threeCheck TvGame, antichess TvGame, computer TvGame, horde TvGame, rapid TvGame, atomic TvGame, crazyhouse TvGame, chess960 TvGame, kingOfTheHill TvGame, best TvGame) *TvChannels200Response {
	this := TvChannels200Response{}
	this.Bot = bot
	this.Blitz = blitz
	this.RacingKings = racingKings
	this.UltraBullet = ultraBullet
	this.Bullet = bullet
	this.Classical = classical
	this.ThreeCheck = threeCheck
	this.Antichess = antichess
	this.Computer = computer
	this.Horde = horde
	this.Rapid = rapid
	this.Atomic = atomic
	this.Crazyhouse = crazyhouse
	this.Chess960 = chess960
	this.KingOfTheHill = kingOfTheHill
	this.Best = best
	return &this
}

// NewTvChannels200ResponseWithDefaults instantiates a new TvChannels200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTvChannels200ResponseWithDefaults() *TvChannels200Response {
	this := TvChannels200Response{}
	return &this
}

// GetBot returns the Bot field value
func (o *TvChannels200Response) GetBot() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Bot
}

// GetBotOk returns a tuple with the Bot field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetBotOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bot, true
}

// SetBot sets field value
func (o *TvChannels200Response) SetBot(v TvGame) {
	o.Bot = v
}

// GetBlitz returns the Blitz field value
func (o *TvChannels200Response) GetBlitz() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Blitz
}

// GetBlitzOk returns a tuple with the Blitz field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetBlitzOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Blitz, true
}

// SetBlitz sets field value
func (o *TvChannels200Response) SetBlitz(v TvGame) {
	o.Blitz = v
}

// GetRacingKings returns the RacingKings field value
func (o *TvChannels200Response) GetRacingKings() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.RacingKings
}

// GetRacingKingsOk returns a tuple with the RacingKings field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetRacingKingsOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RacingKings, true
}

// SetRacingKings sets field value
func (o *TvChannels200Response) SetRacingKings(v TvGame) {
	o.RacingKings = v
}

// GetUltraBullet returns the UltraBullet field value
func (o *TvChannels200Response) GetUltraBullet() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.UltraBullet
}

// GetUltraBulletOk returns a tuple with the UltraBullet field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetUltraBulletOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UltraBullet, true
}

// SetUltraBullet sets field value
func (o *TvChannels200Response) SetUltraBullet(v TvGame) {
	o.UltraBullet = v
}

// GetBullet returns the Bullet field value
func (o *TvChannels200Response) GetBullet() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Bullet
}

// GetBulletOk returns a tuple with the Bullet field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetBulletOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bullet, true
}

// SetBullet sets field value
func (o *TvChannels200Response) SetBullet(v TvGame) {
	o.Bullet = v
}

// GetClassical returns the Classical field value
func (o *TvChannels200Response) GetClassical() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Classical
}

// GetClassicalOk returns a tuple with the Classical field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetClassicalOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classical, true
}

// SetClassical sets field value
func (o *TvChannels200Response) SetClassical(v TvGame) {
	o.Classical = v
}

// GetThreeCheck returns the ThreeCheck field value
func (o *TvChannels200Response) GetThreeCheck() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.ThreeCheck
}

// GetThreeCheckOk returns a tuple with the ThreeCheck field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetThreeCheckOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ThreeCheck, true
}

// SetThreeCheck sets field value
func (o *TvChannels200Response) SetThreeCheck(v TvGame) {
	o.ThreeCheck = v
}

// GetAntichess returns the Antichess field value
func (o *TvChannels200Response) GetAntichess() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Antichess
}

// GetAntichessOk returns a tuple with the Antichess field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetAntichessOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Antichess, true
}

// SetAntichess sets field value
func (o *TvChannels200Response) SetAntichess(v TvGame) {
	o.Antichess = v
}

// GetComputer returns the Computer field value
func (o *TvChannels200Response) GetComputer() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Computer
}

// GetComputerOk returns a tuple with the Computer field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetComputerOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Computer, true
}

// SetComputer sets field value
func (o *TvChannels200Response) SetComputer(v TvGame) {
	o.Computer = v
}

// GetHorde returns the Horde field value
func (o *TvChannels200Response) GetHorde() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Horde
}

// GetHordeOk returns a tuple with the Horde field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetHordeOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Horde, true
}

// SetHorde sets field value
func (o *TvChannels200Response) SetHorde(v TvGame) {
	o.Horde = v
}

// GetRapid returns the Rapid field value
func (o *TvChannels200Response) GetRapid() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Rapid
}

// GetRapidOk returns a tuple with the Rapid field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetRapidOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rapid, true
}

// SetRapid sets field value
func (o *TvChannels200Response) SetRapid(v TvGame) {
	o.Rapid = v
}

// GetAtomic returns the Atomic field value
func (o *TvChannels200Response) GetAtomic() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetAtomicOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Atomic, true
}

// SetAtomic sets field value
func (o *TvChannels200Response) SetAtomic(v TvGame) {
	o.Atomic = v
}

// GetCrazyhouse returns the Crazyhouse field value
func (o *TvChannels200Response) GetCrazyhouse() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Crazyhouse
}

// GetCrazyhouseOk returns a tuple with the Crazyhouse field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetCrazyhouseOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Crazyhouse, true
}

// SetCrazyhouse sets field value
func (o *TvChannels200Response) SetCrazyhouse(v TvGame) {
	o.Crazyhouse = v
}

// GetChess960 returns the Chess960 field value
func (o *TvChannels200Response) GetChess960() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Chess960
}

// GetChess960Ok returns a tuple with the Chess960 field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetChess960Ok() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chess960, true
}

// SetChess960 sets field value
func (o *TvChannels200Response) SetChess960(v TvGame) {
	o.Chess960 = v
}

// GetKingOfTheHill returns the KingOfTheHill field value
func (o *TvChannels200Response) GetKingOfTheHill() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.KingOfTheHill
}

// GetKingOfTheHillOk returns a tuple with the KingOfTheHill field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetKingOfTheHillOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KingOfTheHill, true
}

// SetKingOfTheHill sets field value
func (o *TvChannels200Response) SetKingOfTheHill(v TvGame) {
	o.KingOfTheHill = v
}

// GetBest returns the Best field value
func (o *TvChannels200Response) GetBest() TvGame {
	if o == nil {
		var ret TvGame
		return ret
	}

	return o.Best
}

// GetBestOk returns a tuple with the Best field value
// and a boolean to check if the value has been set.
func (o *TvChannels200Response) GetBestOk() (*TvGame, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Best, true
}

// SetBest sets field value
func (o *TvChannels200Response) SetBest(v TvGame) {
	o.Best = v
}

func (o TvChannels200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TvChannels200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bot"] = o.Bot
	toSerialize["blitz"] = o.Blitz
	toSerialize["racingKings"] = o.RacingKings
	toSerialize["ultraBullet"] = o.UltraBullet
	toSerialize["bullet"] = o.Bullet
	toSerialize["classical"] = o.Classical
	toSerialize["threeCheck"] = o.ThreeCheck
	toSerialize["antichess"] = o.Antichess
	toSerialize["computer"] = o.Computer
	toSerialize["horde"] = o.Horde
	toSerialize["rapid"] = o.Rapid
	toSerialize["atomic"] = o.Atomic
	toSerialize["crazyhouse"] = o.Crazyhouse
	toSerialize["chess960"] = o.Chess960
	toSerialize["kingOfTheHill"] = o.KingOfTheHill
	toSerialize["best"] = o.Best
	return toSerialize, nil
}

func (o *TvChannels200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bot",
		"blitz",
		"racingKings",
		"ultraBullet",
		"bullet",
		"classical",
		"threeCheck",
		"antichess",
		"computer",
		"horde",
		"rapid",
		"atomic",
		"crazyhouse",
		"chess960",
		"kingOfTheHill",
		"best",
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

	varTvChannels200Response := _TvChannels200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTvChannels200Response)

	if err != nil {
		return err
	}

	*o = TvChannels200Response(varTvChannels200Response)

	return err
}

type NullableTvChannels200Response struct {
	value *TvChannels200Response
	isSet bool
}

func (v NullableTvChannels200Response) Get() *TvChannels200Response {
	return v.value
}

func (v *NullableTvChannels200Response) Set(val *TvChannels200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableTvChannels200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableTvChannels200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTvChannels200Response(val *TvChannels200Response) *NullableTvChannels200Response {
	return &NullableTvChannels200Response{value: val, isSet: true}
}

func (v NullableTvChannels200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTvChannels200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


