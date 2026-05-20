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

// checks if the ApiUserActivity200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserActivity200ResponseInner{}

// ApiUserActivity200ResponseInner struct for ApiUserActivity200ResponseInner
type ApiUserActivity200ResponseInner struct {
	Interval ApiUserActivity200ResponseInnerInterval `json:"interval"`
	Games *ApiUserActivity200ResponseInnerGames `json:"games,omitempty"`
	Puzzles *ApiUserActivity200ResponseInnerPuzzles `json:"puzzles,omitempty"`
	Storm *ApiUser200ResponseAllOfPerfsStorm `json:"storm,omitempty"`
	Racer *ApiUser200ResponseAllOfPerfsStorm `json:"racer,omitempty"`
	Streak *ApiUser200ResponseAllOfPerfsStorm `json:"streak,omitempty"`
	Tournaments *ApiUserActivity200ResponseInnerTournaments `json:"tournaments,omitempty"`
	Practice []ApiUserActivity200ResponseInnerPracticeInner `json:"practice,omitempty"`
	Simuls []string `json:"simuls,omitempty"`
	CorrespondenceMoves *ApiUserActivity200ResponseInnerCorrespondenceMoves `json:"correspondenceMoves,omitempty"`
	CorrespondenceEnds *ApiUserActivity200ResponseInnerCorrespondenceEnds `json:"correspondenceEnds,omitempty"`
	Follows *ApiUserActivity200ResponseInnerFollows `json:"follows,omitempty"`
	Studies []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament `json:"studies,omitempty"`
	Teams []ApiUserActivity200ResponseInnerTeamsInner `json:"teams,omitempty"`
	Posts []ApiUserActivity200ResponseInnerPostsInner `json:"posts,omitempty"`
	Patron *ApiUserActivity200ResponseInnerPatron `json:"patron,omitempty"`
	Stream *bool `json:"stream,omitempty"`
}

type _ApiUserActivity200ResponseInner ApiUserActivity200ResponseInner

// NewApiUserActivity200ResponseInner instantiates a new ApiUserActivity200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserActivity200ResponseInner(interval ApiUserActivity200ResponseInnerInterval) *ApiUserActivity200ResponseInner {
	this := ApiUserActivity200ResponseInner{}
	this.Interval = interval
	return &this
}

// NewApiUserActivity200ResponseInnerWithDefaults instantiates a new ApiUserActivity200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserActivity200ResponseInnerWithDefaults() *ApiUserActivity200ResponseInner {
	this := ApiUserActivity200ResponseInner{}
	return &this
}

// GetInterval returns the Interval field value
func (o *ApiUserActivity200ResponseInner) GetInterval() ApiUserActivity200ResponseInnerInterval {
	if o == nil {
		var ret ApiUserActivity200ResponseInnerInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetIntervalOk() (*ApiUserActivity200ResponseInnerInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *ApiUserActivity200ResponseInner) SetInterval(v ApiUserActivity200ResponseInnerInterval) {
	o.Interval = v
}

// GetGames returns the Games field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetGames() ApiUserActivity200ResponseInnerGames {
	if o == nil || IsNil(o.Games) {
		var ret ApiUserActivity200ResponseInnerGames
		return ret
	}
	return *o.Games
}

// GetGamesOk returns a tuple with the Games field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetGamesOk() (*ApiUserActivity200ResponseInnerGames, bool) {
	if o == nil || IsNil(o.Games) {
		return nil, false
	}
	return o.Games, true
}

// HasGames returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasGames() bool {
	if o != nil && !IsNil(o.Games) {
		return true
	}

	return false
}

// SetGames gets a reference to the given ApiUserActivity200ResponseInnerGames and assigns it to the Games field.
func (o *ApiUserActivity200ResponseInner) SetGames(v ApiUserActivity200ResponseInnerGames) {
	o.Games = &v
}

// GetPuzzles returns the Puzzles field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetPuzzles() ApiUserActivity200ResponseInnerPuzzles {
	if o == nil || IsNil(o.Puzzles) {
		var ret ApiUserActivity200ResponseInnerPuzzles
		return ret
	}
	return *o.Puzzles
}

// GetPuzzlesOk returns a tuple with the Puzzles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetPuzzlesOk() (*ApiUserActivity200ResponseInnerPuzzles, bool) {
	if o == nil || IsNil(o.Puzzles) {
		return nil, false
	}
	return o.Puzzles, true
}

// HasPuzzles returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasPuzzles() bool {
	if o != nil && !IsNil(o.Puzzles) {
		return true
	}

	return false
}

// SetPuzzles gets a reference to the given ApiUserActivity200ResponseInnerPuzzles and assigns it to the Puzzles field.
func (o *ApiUserActivity200ResponseInner) SetPuzzles(v ApiUserActivity200ResponseInnerPuzzles) {
	o.Puzzles = &v
}

// GetStorm returns the Storm field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetStorm() ApiUser200ResponseAllOfPerfsStorm {
	if o == nil || IsNil(o.Storm) {
		var ret ApiUser200ResponseAllOfPerfsStorm
		return ret
	}
	return *o.Storm
}

// GetStormOk returns a tuple with the Storm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetStormOk() (*ApiUser200ResponseAllOfPerfsStorm, bool) {
	if o == nil || IsNil(o.Storm) {
		return nil, false
	}
	return o.Storm, true
}

// HasStorm returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasStorm() bool {
	if o != nil && !IsNil(o.Storm) {
		return true
	}

	return false
}

// SetStorm gets a reference to the given ApiUser200ResponseAllOfPerfsStorm and assigns it to the Storm field.
func (o *ApiUserActivity200ResponseInner) SetStorm(v ApiUser200ResponseAllOfPerfsStorm) {
	o.Storm = &v
}

// GetRacer returns the Racer field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetRacer() ApiUser200ResponseAllOfPerfsStorm {
	if o == nil || IsNil(o.Racer) {
		var ret ApiUser200ResponseAllOfPerfsStorm
		return ret
	}
	return *o.Racer
}

// GetRacerOk returns a tuple with the Racer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetRacerOk() (*ApiUser200ResponseAllOfPerfsStorm, bool) {
	if o == nil || IsNil(o.Racer) {
		return nil, false
	}
	return o.Racer, true
}

// HasRacer returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasRacer() bool {
	if o != nil && !IsNil(o.Racer) {
		return true
	}

	return false
}

// SetRacer gets a reference to the given ApiUser200ResponseAllOfPerfsStorm and assigns it to the Racer field.
func (o *ApiUserActivity200ResponseInner) SetRacer(v ApiUser200ResponseAllOfPerfsStorm) {
	o.Racer = &v
}

// GetStreak returns the Streak field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetStreak() ApiUser200ResponseAllOfPerfsStorm {
	if o == nil || IsNil(o.Streak) {
		var ret ApiUser200ResponseAllOfPerfsStorm
		return ret
	}
	return *o.Streak
}

// GetStreakOk returns a tuple with the Streak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetStreakOk() (*ApiUser200ResponseAllOfPerfsStorm, bool) {
	if o == nil || IsNil(o.Streak) {
		return nil, false
	}
	return o.Streak, true
}

// HasStreak returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasStreak() bool {
	if o != nil && !IsNil(o.Streak) {
		return true
	}

	return false
}

// SetStreak gets a reference to the given ApiUser200ResponseAllOfPerfsStorm and assigns it to the Streak field.
func (o *ApiUserActivity200ResponseInner) SetStreak(v ApiUser200ResponseAllOfPerfsStorm) {
	o.Streak = &v
}

// GetTournaments returns the Tournaments field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetTournaments() ApiUserActivity200ResponseInnerTournaments {
	if o == nil || IsNil(o.Tournaments) {
		var ret ApiUserActivity200ResponseInnerTournaments
		return ret
	}
	return *o.Tournaments
}

// GetTournamentsOk returns a tuple with the Tournaments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetTournamentsOk() (*ApiUserActivity200ResponseInnerTournaments, bool) {
	if o == nil || IsNil(o.Tournaments) {
		return nil, false
	}
	return o.Tournaments, true
}

// HasTournaments returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasTournaments() bool {
	if o != nil && !IsNil(o.Tournaments) {
		return true
	}

	return false
}

// SetTournaments gets a reference to the given ApiUserActivity200ResponseInnerTournaments and assigns it to the Tournaments field.
func (o *ApiUserActivity200ResponseInner) SetTournaments(v ApiUserActivity200ResponseInnerTournaments) {
	o.Tournaments = &v
}

// GetPractice returns the Practice field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetPractice() []ApiUserActivity200ResponseInnerPracticeInner {
	if o == nil || IsNil(o.Practice) {
		var ret []ApiUserActivity200ResponseInnerPracticeInner
		return ret
	}
	return o.Practice
}

// GetPracticeOk returns a tuple with the Practice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetPracticeOk() ([]ApiUserActivity200ResponseInnerPracticeInner, bool) {
	if o == nil || IsNil(o.Practice) {
		return nil, false
	}
	return o.Practice, true
}

// HasPractice returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasPractice() bool {
	if o != nil && !IsNil(o.Practice) {
		return true
	}

	return false
}

// SetPractice gets a reference to the given []ApiUserActivity200ResponseInnerPracticeInner and assigns it to the Practice field.
func (o *ApiUserActivity200ResponseInner) SetPractice(v []ApiUserActivity200ResponseInnerPracticeInner) {
	o.Practice = v
}

// GetSimuls returns the Simuls field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetSimuls() []string {
	if o == nil || IsNil(o.Simuls) {
		var ret []string
		return ret
	}
	return o.Simuls
}

// GetSimulsOk returns a tuple with the Simuls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetSimulsOk() ([]string, bool) {
	if o == nil || IsNil(o.Simuls) {
		return nil, false
	}
	return o.Simuls, true
}

// HasSimuls returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasSimuls() bool {
	if o != nil && !IsNil(o.Simuls) {
		return true
	}

	return false
}

// SetSimuls gets a reference to the given []string and assigns it to the Simuls field.
func (o *ApiUserActivity200ResponseInner) SetSimuls(v []string) {
	o.Simuls = v
}

// GetCorrespondenceMoves returns the CorrespondenceMoves field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetCorrespondenceMoves() ApiUserActivity200ResponseInnerCorrespondenceMoves {
	if o == nil || IsNil(o.CorrespondenceMoves) {
		var ret ApiUserActivity200ResponseInnerCorrespondenceMoves
		return ret
	}
	return *o.CorrespondenceMoves
}

// GetCorrespondenceMovesOk returns a tuple with the CorrespondenceMoves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetCorrespondenceMovesOk() (*ApiUserActivity200ResponseInnerCorrespondenceMoves, bool) {
	if o == nil || IsNil(o.CorrespondenceMoves) {
		return nil, false
	}
	return o.CorrespondenceMoves, true
}

// HasCorrespondenceMoves returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasCorrespondenceMoves() bool {
	if o != nil && !IsNil(o.CorrespondenceMoves) {
		return true
	}

	return false
}

// SetCorrespondenceMoves gets a reference to the given ApiUserActivity200ResponseInnerCorrespondenceMoves and assigns it to the CorrespondenceMoves field.
func (o *ApiUserActivity200ResponseInner) SetCorrespondenceMoves(v ApiUserActivity200ResponseInnerCorrespondenceMoves) {
	o.CorrespondenceMoves = &v
}

// GetCorrespondenceEnds returns the CorrespondenceEnds field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetCorrespondenceEnds() ApiUserActivity200ResponseInnerCorrespondenceEnds {
	if o == nil || IsNil(o.CorrespondenceEnds) {
		var ret ApiUserActivity200ResponseInnerCorrespondenceEnds
		return ret
	}
	return *o.CorrespondenceEnds
}

// GetCorrespondenceEndsOk returns a tuple with the CorrespondenceEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetCorrespondenceEndsOk() (*ApiUserActivity200ResponseInnerCorrespondenceEnds, bool) {
	if o == nil || IsNil(o.CorrespondenceEnds) {
		return nil, false
	}
	return o.CorrespondenceEnds, true
}

// HasCorrespondenceEnds returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasCorrespondenceEnds() bool {
	if o != nil && !IsNil(o.CorrespondenceEnds) {
		return true
	}

	return false
}

// SetCorrespondenceEnds gets a reference to the given ApiUserActivity200ResponseInnerCorrespondenceEnds and assigns it to the CorrespondenceEnds field.
func (o *ApiUserActivity200ResponseInner) SetCorrespondenceEnds(v ApiUserActivity200ResponseInnerCorrespondenceEnds) {
	o.CorrespondenceEnds = &v
}

// GetFollows returns the Follows field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetFollows() ApiUserActivity200ResponseInnerFollows {
	if o == nil || IsNil(o.Follows) {
		var ret ApiUserActivity200ResponseInnerFollows
		return ret
	}
	return *o.Follows
}

// GetFollowsOk returns a tuple with the Follows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetFollowsOk() (*ApiUserActivity200ResponseInnerFollows, bool) {
	if o == nil || IsNil(o.Follows) {
		return nil, false
	}
	return o.Follows, true
}

// HasFollows returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasFollows() bool {
	if o != nil && !IsNil(o.Follows) {
		return true
	}

	return false
}

// SetFollows gets a reference to the given ApiUserActivity200ResponseInnerFollows and assigns it to the Follows field.
func (o *ApiUserActivity200ResponseInner) SetFollows(v ApiUserActivity200ResponseInnerFollows) {
	o.Follows = &v
}

// GetStudies returns the Studies field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetStudies() []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament {
	if o == nil || IsNil(o.Studies) {
		var ret []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament
		return ret
	}
	return o.Studies
}

// GetStudiesOk returns a tuple with the Studies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetStudiesOk() ([]ApiUserActivity200ResponseInnerTournamentsBestInnerTournament, bool) {
	if o == nil || IsNil(o.Studies) {
		return nil, false
	}
	return o.Studies, true
}

// HasStudies returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasStudies() bool {
	if o != nil && !IsNil(o.Studies) {
		return true
	}

	return false
}

// SetStudies gets a reference to the given []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament and assigns it to the Studies field.
func (o *ApiUserActivity200ResponseInner) SetStudies(v []ApiUserActivity200ResponseInnerTournamentsBestInnerTournament) {
	o.Studies = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetTeams() []ApiUserActivity200ResponseInnerTeamsInner {
	if o == nil || IsNil(o.Teams) {
		var ret []ApiUserActivity200ResponseInnerTeamsInner
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetTeamsOk() ([]ApiUserActivity200ResponseInnerTeamsInner, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []ApiUserActivity200ResponseInnerTeamsInner and assigns it to the Teams field.
func (o *ApiUserActivity200ResponseInner) SetTeams(v []ApiUserActivity200ResponseInnerTeamsInner) {
	o.Teams = v
}

// GetPosts returns the Posts field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetPosts() []ApiUserActivity200ResponseInnerPostsInner {
	if o == nil || IsNil(o.Posts) {
		var ret []ApiUserActivity200ResponseInnerPostsInner
		return ret
	}
	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetPostsOk() ([]ApiUserActivity200ResponseInnerPostsInner, bool) {
	if o == nil || IsNil(o.Posts) {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasPosts() bool {
	if o != nil && !IsNil(o.Posts) {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []ApiUserActivity200ResponseInnerPostsInner and assigns it to the Posts field.
func (o *ApiUserActivity200ResponseInner) SetPosts(v []ApiUserActivity200ResponseInnerPostsInner) {
	o.Posts = v
}

// GetPatron returns the Patron field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetPatron() ApiUserActivity200ResponseInnerPatron {
	if o == nil || IsNil(o.Patron) {
		var ret ApiUserActivity200ResponseInnerPatron
		return ret
	}
	return *o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetPatronOk() (*ApiUserActivity200ResponseInnerPatron, bool) {
	if o == nil || IsNil(o.Patron) {
		return nil, false
	}
	return o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasPatron() bool {
	if o != nil && !IsNil(o.Patron) {
		return true
	}

	return false
}

// SetPatron gets a reference to the given ApiUserActivity200ResponseInnerPatron and assigns it to the Patron field.
func (o *ApiUserActivity200ResponseInner) SetPatron(v ApiUserActivity200ResponseInnerPatron) {
	o.Patron = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *ApiUserActivity200ResponseInner) GetStream() bool {
	if o == nil || IsNil(o.Stream) {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserActivity200ResponseInner) GetStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *ApiUserActivity200ResponseInner) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *ApiUserActivity200ResponseInner) SetStream(v bool) {
	o.Stream = &v
}

func (o ApiUserActivity200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserActivity200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["interval"] = o.Interval
	if !IsNil(o.Games) {
		toSerialize["games"] = o.Games
	}
	if !IsNil(o.Puzzles) {
		toSerialize["puzzles"] = o.Puzzles
	}
	if !IsNil(o.Storm) {
		toSerialize["storm"] = o.Storm
	}
	if !IsNil(o.Racer) {
		toSerialize["racer"] = o.Racer
	}
	if !IsNil(o.Streak) {
		toSerialize["streak"] = o.Streak
	}
	if !IsNil(o.Tournaments) {
		toSerialize["tournaments"] = o.Tournaments
	}
	if !IsNil(o.Practice) {
		toSerialize["practice"] = o.Practice
	}
	if !IsNil(o.Simuls) {
		toSerialize["simuls"] = o.Simuls
	}
	if !IsNil(o.CorrespondenceMoves) {
		toSerialize["correspondenceMoves"] = o.CorrespondenceMoves
	}
	if !IsNil(o.CorrespondenceEnds) {
		toSerialize["correspondenceEnds"] = o.CorrespondenceEnds
	}
	if !IsNil(o.Follows) {
		toSerialize["follows"] = o.Follows
	}
	if !IsNil(o.Studies) {
		toSerialize["studies"] = o.Studies
	}
	if !IsNil(o.Teams) {
		toSerialize["teams"] = o.Teams
	}
	if !IsNil(o.Posts) {
		toSerialize["posts"] = o.Posts
	}
	if !IsNil(o.Patron) {
		toSerialize["patron"] = o.Patron
	}
	if !IsNil(o.Stream) {
		toSerialize["stream"] = o.Stream
	}
	return toSerialize, nil
}

func (o *ApiUserActivity200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"interval",
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

	varApiUserActivity200ResponseInner := _ApiUserActivity200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUserActivity200ResponseInner)

	if err != nil {
		return err
	}

	*o = ApiUserActivity200ResponseInner(varApiUserActivity200ResponseInner)

	return err
}

type NullableApiUserActivity200ResponseInner struct {
	value *ApiUserActivity200ResponseInner
	isSet bool
}

func (v NullableApiUserActivity200ResponseInner) Get() *ApiUserActivity200ResponseInner {
	return v.value
}

func (v *NullableApiUserActivity200ResponseInner) Set(val *ApiUserActivity200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserActivity200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserActivity200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserActivity200ResponseInner(val *ApiUserActivity200ResponseInner) *NullableApiUserActivity200ResponseInner {
	return &NullableApiUserActivity200ResponseInner{value: val, isSet: true}
}

func (v NullableApiUserActivity200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserActivity200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


