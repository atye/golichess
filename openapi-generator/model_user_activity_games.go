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
)

// checks if the UserActivityGames type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActivityGames{}

// UserActivityGames struct for UserActivityGames
type UserActivityGames struct {
	Chess960 *UserActivityScore `json:"chess960,omitempty"`
	Atomic *UserActivityScore `json:"atomic,omitempty"`
	RacingKings *UserActivityScore `json:"racingKings,omitempty"`
	UltraBullet *UserActivityScore `json:"ultraBullet,omitempty"`
	Blitz *UserActivityScore `json:"blitz,omitempty"`
	KingOfTheHill *UserActivityScore `json:"kingOfTheHill,omitempty"`
	Bullet *UserActivityScore `json:"bullet,omitempty"`
	Correspondence *UserActivityScore `json:"correspondence,omitempty"`
	Horde *UserActivityScore `json:"horde,omitempty"`
	Puzzle *UserActivityScore `json:"puzzle,omitempty"`
	Classical *UserActivityScore `json:"classical,omitempty"`
	Rapid *UserActivityScore `json:"rapid,omitempty"`
}

// NewUserActivityGames instantiates a new UserActivityGames object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivityGames() *UserActivityGames {
	this := UserActivityGames{}
	return &this
}

// NewUserActivityGamesWithDefaults instantiates a new UserActivityGames object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityGamesWithDefaults() *UserActivityGames {
	this := UserActivityGames{}
	return &this
}

// GetChess960 returns the Chess960 field value if set, zero value otherwise.
func (o *UserActivityGames) GetChess960() UserActivityScore {
	if o == nil || IsNil(o.Chess960) {
		var ret UserActivityScore
		return ret
	}
	return *o.Chess960
}

// GetChess960Ok returns a tuple with the Chess960 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetChess960Ok() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Chess960) {
		return nil, false
	}
	return o.Chess960, true
}

// HasChess960 returns a boolean if a field has been set.
func (o *UserActivityGames) HasChess960() bool {
	if o != nil && !IsNil(o.Chess960) {
		return true
	}

	return false
}

// SetChess960 gets a reference to the given UserActivityScore and assigns it to the Chess960 field.
func (o *UserActivityGames) SetChess960(v UserActivityScore) {
	o.Chess960 = &v
}

// GetAtomic returns the Atomic field value if set, zero value otherwise.
func (o *UserActivityGames) GetAtomic() UserActivityScore {
	if o == nil || IsNil(o.Atomic) {
		var ret UserActivityScore
		return ret
	}
	return *o.Atomic
}

// GetAtomicOk returns a tuple with the Atomic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetAtomicOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Atomic) {
		return nil, false
	}
	return o.Atomic, true
}

// HasAtomic returns a boolean if a field has been set.
func (o *UserActivityGames) HasAtomic() bool {
	if o != nil && !IsNil(o.Atomic) {
		return true
	}

	return false
}

// SetAtomic gets a reference to the given UserActivityScore and assigns it to the Atomic field.
func (o *UserActivityGames) SetAtomic(v UserActivityScore) {
	o.Atomic = &v
}

// GetRacingKings returns the RacingKings field value if set, zero value otherwise.
func (o *UserActivityGames) GetRacingKings() UserActivityScore {
	if o == nil || IsNil(o.RacingKings) {
		var ret UserActivityScore
		return ret
	}
	return *o.RacingKings
}

// GetRacingKingsOk returns a tuple with the RacingKings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetRacingKingsOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.RacingKings) {
		return nil, false
	}
	return o.RacingKings, true
}

// HasRacingKings returns a boolean if a field has been set.
func (o *UserActivityGames) HasRacingKings() bool {
	if o != nil && !IsNil(o.RacingKings) {
		return true
	}

	return false
}

// SetRacingKings gets a reference to the given UserActivityScore and assigns it to the RacingKings field.
func (o *UserActivityGames) SetRacingKings(v UserActivityScore) {
	o.RacingKings = &v
}

// GetUltraBullet returns the UltraBullet field value if set, zero value otherwise.
func (o *UserActivityGames) GetUltraBullet() UserActivityScore {
	if o == nil || IsNil(o.UltraBullet) {
		var ret UserActivityScore
		return ret
	}
	return *o.UltraBullet
}

// GetUltraBulletOk returns a tuple with the UltraBullet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetUltraBulletOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.UltraBullet) {
		return nil, false
	}
	return o.UltraBullet, true
}

// HasUltraBullet returns a boolean if a field has been set.
func (o *UserActivityGames) HasUltraBullet() bool {
	if o != nil && !IsNil(o.UltraBullet) {
		return true
	}

	return false
}

// SetUltraBullet gets a reference to the given UserActivityScore and assigns it to the UltraBullet field.
func (o *UserActivityGames) SetUltraBullet(v UserActivityScore) {
	o.UltraBullet = &v
}

// GetBlitz returns the Blitz field value if set, zero value otherwise.
func (o *UserActivityGames) GetBlitz() UserActivityScore {
	if o == nil || IsNil(o.Blitz) {
		var ret UserActivityScore
		return ret
	}
	return *o.Blitz
}

// GetBlitzOk returns a tuple with the Blitz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetBlitzOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Blitz) {
		return nil, false
	}
	return o.Blitz, true
}

// HasBlitz returns a boolean if a field has been set.
func (o *UserActivityGames) HasBlitz() bool {
	if o != nil && !IsNil(o.Blitz) {
		return true
	}

	return false
}

// SetBlitz gets a reference to the given UserActivityScore and assigns it to the Blitz field.
func (o *UserActivityGames) SetBlitz(v UserActivityScore) {
	o.Blitz = &v
}

// GetKingOfTheHill returns the KingOfTheHill field value if set, zero value otherwise.
func (o *UserActivityGames) GetKingOfTheHill() UserActivityScore {
	if o == nil || IsNil(o.KingOfTheHill) {
		var ret UserActivityScore
		return ret
	}
	return *o.KingOfTheHill
}

// GetKingOfTheHillOk returns a tuple with the KingOfTheHill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetKingOfTheHillOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.KingOfTheHill) {
		return nil, false
	}
	return o.KingOfTheHill, true
}

// HasKingOfTheHill returns a boolean if a field has been set.
func (o *UserActivityGames) HasKingOfTheHill() bool {
	if o != nil && !IsNil(o.KingOfTheHill) {
		return true
	}

	return false
}

// SetKingOfTheHill gets a reference to the given UserActivityScore and assigns it to the KingOfTheHill field.
func (o *UserActivityGames) SetKingOfTheHill(v UserActivityScore) {
	o.KingOfTheHill = &v
}

// GetBullet returns the Bullet field value if set, zero value otherwise.
func (o *UserActivityGames) GetBullet() UserActivityScore {
	if o == nil || IsNil(o.Bullet) {
		var ret UserActivityScore
		return ret
	}
	return *o.Bullet
}

// GetBulletOk returns a tuple with the Bullet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetBulletOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Bullet) {
		return nil, false
	}
	return o.Bullet, true
}

// HasBullet returns a boolean if a field has been set.
func (o *UserActivityGames) HasBullet() bool {
	if o != nil && !IsNil(o.Bullet) {
		return true
	}

	return false
}

// SetBullet gets a reference to the given UserActivityScore and assigns it to the Bullet field.
func (o *UserActivityGames) SetBullet(v UserActivityScore) {
	o.Bullet = &v
}

// GetCorrespondence returns the Correspondence field value if set, zero value otherwise.
func (o *UserActivityGames) GetCorrespondence() UserActivityScore {
	if o == nil || IsNil(o.Correspondence) {
		var ret UserActivityScore
		return ret
	}
	return *o.Correspondence
}

// GetCorrespondenceOk returns a tuple with the Correspondence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetCorrespondenceOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Correspondence) {
		return nil, false
	}
	return o.Correspondence, true
}

// HasCorrespondence returns a boolean if a field has been set.
func (o *UserActivityGames) HasCorrespondence() bool {
	if o != nil && !IsNil(o.Correspondence) {
		return true
	}

	return false
}

// SetCorrespondence gets a reference to the given UserActivityScore and assigns it to the Correspondence field.
func (o *UserActivityGames) SetCorrespondence(v UserActivityScore) {
	o.Correspondence = &v
}

// GetHorde returns the Horde field value if set, zero value otherwise.
func (o *UserActivityGames) GetHorde() UserActivityScore {
	if o == nil || IsNil(o.Horde) {
		var ret UserActivityScore
		return ret
	}
	return *o.Horde
}

// GetHordeOk returns a tuple with the Horde field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetHordeOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Horde) {
		return nil, false
	}
	return o.Horde, true
}

// HasHorde returns a boolean if a field has been set.
func (o *UserActivityGames) HasHorde() bool {
	if o != nil && !IsNil(o.Horde) {
		return true
	}

	return false
}

// SetHorde gets a reference to the given UserActivityScore and assigns it to the Horde field.
func (o *UserActivityGames) SetHorde(v UserActivityScore) {
	o.Horde = &v
}

// GetPuzzle returns the Puzzle field value if set, zero value otherwise.
func (o *UserActivityGames) GetPuzzle() UserActivityScore {
	if o == nil || IsNil(o.Puzzle) {
		var ret UserActivityScore
		return ret
	}
	return *o.Puzzle
}

// GetPuzzleOk returns a tuple with the Puzzle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetPuzzleOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Puzzle) {
		return nil, false
	}
	return o.Puzzle, true
}

// HasPuzzle returns a boolean if a field has been set.
func (o *UserActivityGames) HasPuzzle() bool {
	if o != nil && !IsNil(o.Puzzle) {
		return true
	}

	return false
}

// SetPuzzle gets a reference to the given UserActivityScore and assigns it to the Puzzle field.
func (o *UserActivityGames) SetPuzzle(v UserActivityScore) {
	o.Puzzle = &v
}

// GetClassical returns the Classical field value if set, zero value otherwise.
func (o *UserActivityGames) GetClassical() UserActivityScore {
	if o == nil || IsNil(o.Classical) {
		var ret UserActivityScore
		return ret
	}
	return *o.Classical
}

// GetClassicalOk returns a tuple with the Classical field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetClassicalOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Classical) {
		return nil, false
	}
	return o.Classical, true
}

// HasClassical returns a boolean if a field has been set.
func (o *UserActivityGames) HasClassical() bool {
	if o != nil && !IsNil(o.Classical) {
		return true
	}

	return false
}

// SetClassical gets a reference to the given UserActivityScore and assigns it to the Classical field.
func (o *UserActivityGames) SetClassical(v UserActivityScore) {
	o.Classical = &v
}

// GetRapid returns the Rapid field value if set, zero value otherwise.
func (o *UserActivityGames) GetRapid() UserActivityScore {
	if o == nil || IsNil(o.Rapid) {
		var ret UserActivityScore
		return ret
	}
	return *o.Rapid
}

// GetRapidOk returns a tuple with the Rapid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivityGames) GetRapidOk() (*UserActivityScore, bool) {
	if o == nil || IsNil(o.Rapid) {
		return nil, false
	}
	return o.Rapid, true
}

// HasRapid returns a boolean if a field has been set.
func (o *UserActivityGames) HasRapid() bool {
	if o != nil && !IsNil(o.Rapid) {
		return true
	}

	return false
}

// SetRapid gets a reference to the given UserActivityScore and assigns it to the Rapid field.
func (o *UserActivityGames) SetRapid(v UserActivityScore) {
	o.Rapid = &v
}

func (o UserActivityGames) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActivityGames) ToMap() (map[string]interface{}, error) {
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

type NullableUserActivityGames struct {
	value *UserActivityGames
	isSet bool
}

func (v NullableUserActivityGames) Get() *UserActivityGames {
	return v.value
}

func (v *NullableUserActivityGames) Set(val *UserActivityGames) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivityGames) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivityGames) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivityGames(val *UserActivityGames) *NullableUserActivityGames {
	return &NullableUserActivityGames{value: val, isSet: true}
}

func (v NullableUserActivityGames) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivityGames) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


