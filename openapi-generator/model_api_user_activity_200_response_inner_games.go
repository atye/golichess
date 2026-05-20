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

// checks if the ApiUserActivity200ResponseInnerGames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserActivity200ResponseInnerGames{}

// ApiUserActivity200ResponseInnerGames struct for ApiUserActivity200ResponseInnerGames
type ApiUserActivity200ResponseInnerGames struct {
	Chess960 *ApiUserActivity200ResponseInnerGamesChess960 `json:"chess960,omitempty"`
	Atomic *ApiUserActivity200ResponseInnerGamesChess960 `json:"atomic,omitempty"`
	RacingKings *ApiUserActivity200ResponseInnerGamesChess960 `json:"racingKings,omitempty"`
	UltraBullet *ApiUserActivity200ResponseInnerGamesChess960 `json:"ultraBullet,omitempty"`
	Blitz *ApiUserActivity200ResponseInnerGamesChess960 `json:"blitz,omitempty"`
	KingOfTheHill *ApiUserActivity200ResponseInnerGamesChess960 `json:"kingOfTheHill,omitempty"`
	Bullet *ApiUserActivity200ResponseInnerGamesChess960 `json:"bullet,omitempty"`
	Correspondence *ApiUserActivity200ResponseInnerGamesChess960 `json:"correspondence,omitempty"`
	Horde *ApiUserActivity200ResponseInnerGamesChess960 `json:"horde,omitempty"`
	Puzzle *ApiUserActivity200ResponseInnerGamesChess960 `json:"puzzle,omitempty"`
	Classical *ApiUserActivity200ResponseInnerGamesChess960 `json:"classical,omitempty"`
	Rapid *ApiUserActivity200ResponseInnerGamesChess960 `json:"rapid,omitempty"`
}

// NewApiUserActivity200ResponseInnerGames instantiates a new ApiUserActivity200ResponseInnerGames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserActivity200ResponseInnerGames() *ApiUserActivity200ResponseInnerGames {
	this := ApiUserActivity200ResponseInnerGames{}
	return &this
}

// NewApiUserActivity200ResponseInnerGamesWithDefaults instantiates a new ApiUserActivity200ResponseInnerGames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserActivity200ResponseInnerGamesWithDefaults() *ApiUserActivity200ResponseInnerGames {
	this := ApiUserActivity200ResponseInnerGames{}
	return &this
}

// GetChess960 returns the Chess960 field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetChess960() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Chess960) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Chess960
}

// GetChess960Ok returns a tuple with the Chess960 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetChess960Ok() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Chess960) {
		return nil, false
	}
	return o.Chess960, true
}

// HasChess960 returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasChess960() bool {
	if o != nil && !IsNil(o.Chess960) {
		return true
	}

	return false
}

// SetChess960 gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Chess960 field.
func (o *ApiUserActivity200ResponseInnerGames) SetChess960(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Chess960 = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetAtomic() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Atomic) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetAtomicOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Atomic) {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasAtomic() bool {
	if o != nil && !IsNil(o.Atomic) {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Atomic field.
func (o *ApiUserActivity200ResponseInnerGames) SetAtomic(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Atomic = &v
}

// GetRacingKings returns the RacingKings field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetRacingKings() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.RacingKings) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.RacingKings
}

// GetRacingKingsOk returns a tuple with the RacingKings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetRacingKingsOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.RacingKings) {
		return nil, false
	}
	return o.RacingKings, true
}

// HasRacingKings returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasRacingKings() bool {
	if o != nil && !IsNil(o.RacingKings) {
		return true
	}

	return false
}

// SetRacingKings gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the RacingKings field.
func (o *ApiUserActivity200ResponseInnerGames) SetRacingKings(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.RacingKings = &v
}

// GetUltraBullet returns the UltraBullet field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetUltraBullet() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.UltraBullet) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.UltraBullet
}

// GetUltraBulletOk returns a tuple with the UltraBullet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetUltraBulletOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.UltraBullet) {
		return nil, false
	}
	return o.UltraBullet, true
}

// HasUltraBullet returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasUltraBullet() bool {
	if o != nil && !IsNil(o.UltraBullet) {
		return true
	}

	return false
}

// SetUltraBullet gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the UltraBullet field.
func (o *ApiUserActivity200ResponseInnerGames) SetUltraBullet(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.UltraBullet = &v
}

// GetBlitz returns the Blitz field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetBlitz() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Blitz) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Blitz
}

// GetBlitzOk returns a tuple with the Blitz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetBlitzOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Blitz) {
		return nil, false
	}
	return o.Blitz, true
}

// HasBlitz returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasBlitz() bool {
	if o != nil && !IsNil(o.Blitz) {
		return true
	}

	return false
}

// SetBlitz gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Blitz field.
func (o *ApiUserActivity200ResponseInnerGames) SetBlitz(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Blitz = &v
}

// GetKingOfTheHill returns the KingOfTheHill field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetKingOfTheHill() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.KingOfTheHill) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.KingOfTheHill
}

// GetKingOfTheHillOk returns a tuple with the KingOfTheHill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetKingOfTheHillOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.KingOfTheHill) {
		return nil, false
	}
	return o.KingOfTheHill, true
}

// HasKingOfTheHill returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasKingOfTheHill() bool {
	if o != nil && !IsNil(o.KingOfTheHill) {
		return true
	}

	return false
}

// SetKingOfTheHill gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the KingOfTheHill field.
func (o *ApiUserActivity200ResponseInnerGames) SetKingOfTheHill(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.KingOfTheHill = &v
}

// GetBullet returns the Bullet field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetBullet() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Bullet) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Bullet
}

// GetBulletOk returns a tuple with the Bullet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetBulletOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Bullet) {
		return nil, false
	}
	return o.Bullet, true
}

// HasBullet returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasBullet() bool {
	if o != nil && !IsNil(o.Bullet) {
		return true
	}

	return false
}

// SetBullet gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Bullet field.
func (o *ApiUserActivity200ResponseInnerGames) SetBullet(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Bullet = &v
}

// GetCorrespondence returns the Correspondence field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetCorrespondence() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Correspondence) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Correspondence
}

// GetCorrespondenceOk returns a tuple with the Correspondence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetCorrespondenceOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Correspondence) {
		return nil, false
	}
	return o.Correspondence, true
}

// HasCorrespondence returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasCorrespondence() bool {
	if o != nil && !IsNil(o.Correspondence) {
		return true
	}

	return false
}

// SetCorrespondence gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Correspondence field.
func (o *ApiUserActivity200ResponseInnerGames) SetCorrespondence(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Correspondence = &v
}

// GetHorde returns the Horde field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetHorde() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Horde) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Horde
}

// GetHordeOk returns a tuple with the Horde field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetHordeOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Horde) {
		return nil, false
	}
	return o.Horde, true
}

// HasHorde returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasHorde() bool {
	if o != nil && !IsNil(o.Horde) {
		return true
	}

	return false
}

// SetHorde gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Horde field.
func (o *ApiUserActivity200ResponseInnerGames) SetHorde(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Horde = &v
}

// GetPuzzle returns the Puzzle field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetPuzzle() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Puzzle) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Puzzle
}

// GetPuzzleOk returns a tuple with the Puzzle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetPuzzleOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Puzzle) {
		return nil, false
	}
	return o.Puzzle, true
}

// HasPuzzle returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasPuzzle() bool {
	if o != nil && !IsNil(o.Puzzle) {
		return true
	}

	return false
}

// SetPuzzle gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Puzzle field.
func (o *ApiUserActivity200ResponseInnerGames) SetPuzzle(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Puzzle = &v
}

// GetClassical returns the Classical field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetClassical() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Classical) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Classical
}

// GetClassicalOk returns a tuple with the Classical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetClassicalOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Classical) {
		return nil, false
	}
	return o.Classical, true
}

// HasClassical returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasClassical() bool {
	if o != nil && !IsNil(o.Classical) {
		return true
	}

	return false
}

// SetClassical gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Classical field.
func (o *ApiUserActivity200ResponseInnerGames) SetClassical(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Classical = &v
}

// GetRapid returns the Rapid field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInnerGames) GetRapid() ApiUserActivity200ResponseInnerGamesChess960 {
	if o == nil || IsNil(o.Rapid) {
		var ret ApiUserActivity200ResponseInnerGamesChess960
		return ret
	}
	return *o.Rapid
}

// GetRapidOk returns a tuple with the Rapid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInnerGames) GetRapidOk() (*ApiUserActivity200ResponseInnerGamesChess960, bool) {
	if o == nil || IsNil(o.Rapid) {
		return nil, false
	}
	return o.Rapid, true
}

// HasRapid returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInnerGames) HasRapid() bool {
	if o != nil && !IsNil(o.Rapid) {
		return true
	}

	return false
}

// SetRapid gets a reference to the given ApiUserActivity200ResponseInnerGamesChess960 and assigns it to the Rapid field.
func (o *ApiUserActivity200ResponseInnerGames) SetRapid(v ApiUserActivity200ResponseInnerGamesChess960) {
	o.Rapid = &v
}

func (o ApiUserActivity200ResponseInnerGames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserActivity200ResponseInnerGames) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Chess960) {
		toSerialize["chess960"] = o.Chess960
	}
	if !IsNil(o.Atomic) {
		toSerialize["atomic"] = o.Atomic
	}
	if !IsNil(o.RacingKings) {
		toSerialize["racingKings"] = o.RacingKings
	}
	if !IsNil(o.UltraBullet) {
		toSerialize["ultraBullet"] = o.UltraBullet
	}
	if !IsNil(o.Blitz) {
		toSerialize["blitz"] = o.Blitz
	}
	if !IsNil(o.KingOfTheHill) {
		toSerialize["kingOfTheHill"] = o.KingOfTheHill
	}
	if !IsNil(o.Bullet) {
		toSerialize["bullet"] = o.Bullet
	}
	if !IsNil(o.Correspondence) {
		toSerialize["correspondence"] = o.Correspondence
	}
	if !IsNil(o.Horde) {
		toSerialize["horde"] = o.Horde
	}
	if !IsNil(o.Puzzle) {
		toSerialize["puzzle"] = o.Puzzle
	}
	if !IsNil(o.Classical) {
		toSerialize["classical"] = o.Classical
	}
	if !IsNil(o.Rapid) {
		toSerialize["rapid"] = o.Rapid
	}
	return toSerialize, nil
}

type NullableApiUserActivity200ResponseInnerGames struct {
	value *ApiUserActivity200ResponseInnerGames
	isSet bool
}

func (v NullableApiUserActivity200ResponseInnerGames) Get() *ApiUserActivity200ResponseInnerGames {
	return v.value
}

func (v *NullableApiUserActivity200ResponseInnerGames) Set(val *ApiUserActivity200ResponseInnerGames) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserActivity200ResponseInnerGames) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserActivity200ResponseInnerGames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserActivity200ResponseInnerGames(val *ApiUserActivity200ResponseInnerGames) *NullableApiUserActivity200ResponseInnerGames {
	return &NullableApiUserActivity200ResponseInnerGames{value: val, isSet: true}
}

func (v NullableApiUserActivity200ResponseInnerGames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserActivity200ResponseInnerGames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


