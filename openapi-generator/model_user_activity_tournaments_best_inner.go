/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.147
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserActivityTournamentsBestInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActivityTournamentsBestInner{}

// UserActivityTournamentsBestInner struct for UserActivityTournamentsBestInner
type UserActivityTournamentsBestInner struct {
	Tournament UserActivityTournamentsBestInnerTournament `json:"tournament"`
	NbGames int32 `json:"nbGames"`
	Score int32 `json:"score"`
	Rank int32 `json:"rank"`
	RankPercent int32 `json:"rankPercent"`
}

type _UserActivityTournamentsBestInner UserActivityTournamentsBestInner

// NewUserActivityTournamentsBestInner instantiates a new UserActivityTournamentsBestInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivityTournamentsBestInner(tournament UserActivityTournamentsBestInnerTournament, nbGames int32, score int32, rank int32, rankPercent int32) *UserActivityTournamentsBestInner {
	this := UserActivityTournamentsBestInner{}
	this.Tournament = tournament
	this.NbGames = nbGames
	this.Score = score
	this.Rank = rank
	this.RankPercent = rankPercent
	return &this
}

// NewUserActivityTournamentsBestInnerWithDefaults instantiates a new UserActivityTournamentsBestInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityTournamentsBestInnerWithDefaults() *UserActivityTournamentsBestInner {
	this := UserActivityTournamentsBestInner{}
	return &this
}

// GetTournament returns the Tournament field value
func (o *UserActivityTournamentsBestInner) GetTournament() UserActivityTournamentsBestInnerTournament {
	if o == nil {
		var ret UserActivityTournamentsBestInnerTournament
		return ret
	}

	return o.Tournament
}

// GetTournamentOk returns a tuple with the Tournament field value
// and a boolean to check if the value has been set.
func (o *UserActivityTournamentsBestInner) GetTournamentOk() (*UserActivityTournamentsBestInnerTournament, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tournament, true
}

// SetTournament sets field value
func (o *UserActivityTournamentsBestInner) SetTournament(v UserActivityTournamentsBestInnerTournament) {
	o.Tournament = v
}

// GetNbGames returns the NbGames field value
func (o *UserActivityTournamentsBestInner) GetNbGames() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbGames
}

// GetNbGamesOk returns a tuple with the NbGames field value
// and a boolean to check if the value has been set.
func (o *UserActivityTournamentsBestInner) GetNbGamesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbGames, true
}

// SetNbGames sets field value
func (o *UserActivityTournamentsBestInner) SetNbGames(v int32) {
	o.NbGames = v
}

// GetScore returns the Score field value
func (o *UserActivityTournamentsBestInner) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *UserActivityTournamentsBestInner) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *UserActivityTournamentsBestInner) SetScore(v int32) {
	o.Score = v
}

// GetRank returns the Rank field value
func (o *UserActivityTournamentsBestInner) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *UserActivityTournamentsBestInner) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *UserActivityTournamentsBestInner) SetRank(v int32) {
	o.Rank = v
}

// GetRankPercent returns the RankPercent field value
func (o *UserActivityTournamentsBestInner) GetRankPercent() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RankPercent
}

// GetRankPercentOk returns a tuple with the RankPercent field value
// and a boolean to check if the value has been set.
func (o *UserActivityTournamentsBestInner) GetRankPercentOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RankPercent, true
}

// SetRankPercent sets field value
func (o *UserActivityTournamentsBestInner) SetRankPercent(v int32) {
	o.RankPercent = v
}

func (o UserActivityTournamentsBestInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActivityTournamentsBestInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tournament"] = o.Tournament
	toSerialize["nbGames"] = o.NbGames
	toSerialize["score"] = o.Score
	toSerialize["rank"] = o.Rank
	toSerialize["rankPercent"] = o.RankPercent
	return toSerialize, nil
}

func (o *UserActivityTournamentsBestInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tournament",
		"nbGames",
		"score",
		"rank",
		"rankPercent",
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

	varUserActivityTournamentsBestInner := _UserActivityTournamentsBestInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserActivityTournamentsBestInner)

	if err != nil {
		return err
	}

	*o = UserActivityTournamentsBestInner(varUserActivityTournamentsBestInner)

	return err
}

type NullableUserActivityTournamentsBestInner struct {
	value *UserActivityTournamentsBestInner
	isSet bool
}

func (v NullableUserActivityTournamentsBestInner) Get() *UserActivityTournamentsBestInner {
	return v.value
}

func (v *NullableUserActivityTournamentsBestInner) Set(val *UserActivityTournamentsBestInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivityTournamentsBestInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivityTournamentsBestInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivityTournamentsBestInner(val *UserActivityTournamentsBestInner) *NullableUserActivityTournamentsBestInner {
	return &NullableUserActivityTournamentsBestInner{value: val, isSet: true}
}

func (v NullableUserActivityTournamentsBestInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivityTournamentsBestInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


