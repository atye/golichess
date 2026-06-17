/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.146
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserActivity{}

// UserActivity struct for UserActivity
type UserActivity struct {
	Interval UserActivityInterval `json:"interval"`
	Games *UserActivityGames `json:"games,omitempty"`
	Puzzles *UserActivityPuzzles `json:"puzzles,omitempty"`
	Storm *PuzzleModePerf `json:"storm,omitempty"`
	Racer *PuzzleModePerf `json:"racer,omitempty"`
	Streak *PuzzleModePerf `json:"streak,omitempty"`
	Tournaments *UserActivityTournaments `json:"tournaments,omitempty"`
	Practice []UserActivityPracticeInner `json:"practice,omitempty"`
	Simuls []string `json:"simuls,omitempty"`
	CorrespondenceMoves *UserActivityCorrespondenceMoves `json:"correspondenceMoves,omitempty"`
	CorrespondenceEnds *UserActivityCorrespondenceEnds `json:"correspondenceEnds,omitempty"`
	Follows *UserActivityFollows `json:"follows,omitempty"`
	Studies []UserActivityTournamentsBestInnerTournament `json:"studies,omitempty"`
	Teams []UserActivityTeamsInner `json:"teams,omitempty"`
	Posts []UserActivityPostsInner `json:"posts,omitempty"`
	Patron *UserActivityPatron `json:"patron,omitempty"`
	Stream *bool `json:"stream,omitempty"`
}

type _UserActivity UserActivity

// NewUserActivity instantiates a new UserActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserActivity(interval UserActivityInterval) *UserActivity {
	this := UserActivity{}
	this.Interval = interval
	return &this
}

// NewUserActivityWithDefaults instantiates a new UserActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserActivityWithDefaults() *UserActivity {
	this := UserActivity{}
	return &this
}

// GetInterval returns the Interval field value
func (o *UserActivity) GetInterval() UserActivityInterval {
	if o == nil {
		var ret UserActivityInterval
		return ret
	}

	return o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value
// and a boolean to check if the value has been set.
func (o *UserActivity) GetIntervalOk() (*UserActivityInterval, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Interval, true
}

// SetInterval sets field value
func (o *UserActivity) SetInterval(v UserActivityInterval) {
	o.Interval = v
}

// GetGames returns the Games field value if set, zero value otherwise.
func (o *UserActivity) GetGames() UserActivityGames {
	if o == nil || IsNil(o.Games) {
		var ret UserActivityGames
		return ret
	}
	return *o.Games
}

// GetGamesOk returns a tuple with the Games field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetGamesOk() (*UserActivityGames, bool) {
	if o == nil || IsNil(o.Games) {
		return nil, false
	}
	return o.Games, true
}

// HasGames returns a boolean if a field has been set.
func (o *UserActivity) HasGames() bool {
	if o != nil && !IsNil(o.Games) {
		return true
	}

	return false
}

// SetGames gets a reference to the given UserActivityGames and assigns it to the Games field.
func (o *UserActivity) SetGames(v UserActivityGames) {
	o.Games = &v
}

// GetPuzzles returns the Puzzles field value if set, zero value otherwise.
func (o *UserActivity) GetPuzzles() UserActivityPuzzles {
	if o == nil || IsNil(o.Puzzles) {
		var ret UserActivityPuzzles
		return ret
	}
	return *o.Puzzles
}

// GetPuzzlesOk returns a tuple with the Puzzles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetPuzzlesOk() (*UserActivityPuzzles, bool) {
	if o == nil || IsNil(o.Puzzles) {
		return nil, false
	}
	return o.Puzzles, true
}

// HasPuzzles returns a boolean if a field has been set.
func (o *UserActivity) HasPuzzles() bool {
	if o != nil && !IsNil(o.Puzzles) {
		return true
	}

	return false
}

// SetPuzzles gets a reference to the given UserActivityPuzzles and assigns it to the Puzzles field.
func (o *UserActivity) SetPuzzles(v UserActivityPuzzles) {
	o.Puzzles = &v
}

// GetStorm returns the Storm field value if set, zero value otherwise.
func (o *UserActivity) GetStorm() PuzzleModePerf {
	if o == nil || IsNil(o.Storm) {
		var ret PuzzleModePerf
		return ret
	}
	return *o.Storm
}

// GetStormOk returns a tuple with the Storm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetStormOk() (*PuzzleModePerf, bool) {
	if o == nil || IsNil(o.Storm) {
		return nil, false
	}
	return o.Storm, true
}

// HasStorm returns a boolean if a field has been set.
func (o *UserActivity) HasStorm() bool {
	if o != nil && !IsNil(o.Storm) {
		return true
	}

	return false
}

// SetStorm gets a reference to the given PuzzleModePerf and assigns it to the Storm field.
func (o *UserActivity) SetStorm(v PuzzleModePerf) {
	o.Storm = &v
}

// GetRacer returns the Racer field value if set, zero value otherwise.
func (o *UserActivity) GetRacer() PuzzleModePerf {
	if o == nil || IsNil(o.Racer) {
		var ret PuzzleModePerf
		return ret
	}
	return *o.Racer
}

// GetRacerOk returns a tuple with the Racer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetRacerOk() (*PuzzleModePerf, bool) {
	if o == nil || IsNil(o.Racer) {
		return nil, false
	}
	return o.Racer, true
}

// HasRacer returns a boolean if a field has been set.
func (o *UserActivity) HasRacer() bool {
	if o != nil && !IsNil(o.Racer) {
		return true
	}

	return false
}

// SetRacer gets a reference to the given PuzzleModePerf and assigns it to the Racer field.
func (o *UserActivity) SetRacer(v PuzzleModePerf) {
	o.Racer = &v
}

// GetStreak returns the Streak field value if set, zero value otherwise.
func (o *UserActivity) GetStreak() PuzzleModePerf {
	if o == nil || IsNil(o.Streak) {
		var ret PuzzleModePerf
		return ret
	}
	return *o.Streak
}

// GetStreakOk returns a tuple with the Streak field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetStreakOk() (*PuzzleModePerf, bool) {
	if o == nil || IsNil(o.Streak) {
		return nil, false
	}
	return o.Streak, true
}

// HasStreak returns a boolean if a field has been set.
func (o *UserActivity) HasStreak() bool {
	if o != nil && !IsNil(o.Streak) {
		return true
	}

	return false
}

// SetStreak gets a reference to the given PuzzleModePerf and assigns it to the Streak field.
func (o *UserActivity) SetStreak(v PuzzleModePerf) {
	o.Streak = &v
}

// GetTournaments returns the Tournaments field value if set, zero value otherwise.
func (o *UserActivity) GetTournaments() UserActivityTournaments {
	if o == nil || IsNil(o.Tournaments) {
		var ret UserActivityTournaments
		return ret
	}
	return *o.Tournaments
}

// GetTournamentsOk returns a tuple with the Tournaments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetTournamentsOk() (*UserActivityTournaments, bool) {
	if o == nil || IsNil(o.Tournaments) {
		return nil, false
	}
	return o.Tournaments, true
}

// HasTournaments returns a boolean if a field has been set.
func (o *UserActivity) HasTournaments() bool {
	if o != nil && !IsNil(o.Tournaments) {
		return true
	}

	return false
}

// SetTournaments gets a reference to the given UserActivityTournaments and assigns it to the Tournaments field.
func (o *UserActivity) SetTournaments(v UserActivityTournaments) {
	o.Tournaments = &v
}

// GetPractice returns the Practice field value if set, zero value otherwise.
func (o *UserActivity) GetPractice() []UserActivityPracticeInner {
	if o == nil || IsNil(o.Practice) {
		var ret []UserActivityPracticeInner
		return ret
	}
	return o.Practice
}

// GetPracticeOk returns a tuple with the Practice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetPracticeOk() ([]UserActivityPracticeInner, bool) {
	if o == nil || IsNil(o.Practice) {
		return nil, false
	}
	return o.Practice, true
}

// HasPractice returns a boolean if a field has been set.
func (o *UserActivity) HasPractice() bool {
	if o != nil && !IsNil(o.Practice) {
		return true
	}

	return false
}

// SetPractice gets a reference to the given []UserActivityPracticeInner and assigns it to the Practice field.
func (o *UserActivity) SetPractice(v []UserActivityPracticeInner) {
	o.Practice = v
}

// GetSimuls returns the Simuls field value if set, zero value otherwise.
func (o *UserActivity) GetSimuls() []string {
	if o == nil || IsNil(o.Simuls) {
		var ret []string
		return ret
	}
	return o.Simuls
}

// GetSimulsOk returns a tuple with the Simuls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetSimulsOk() ([]string, bool) {
	if o == nil || IsNil(o.Simuls) {
		return nil, false
	}
	return o.Simuls, true
}

// HasSimuls returns a boolean if a field has been set.
func (o *UserActivity) HasSimuls() bool {
	if o != nil && !IsNil(o.Simuls) {
		return true
	}

	return false
}

// SetSimuls gets a reference to the given []string and assigns it to the Simuls field.
func (o *UserActivity) SetSimuls(v []string) {
	o.Simuls = v
}

// GetCorrespondenceMoves returns the CorrespondenceMoves field value if set, zero value otherwise.
func (o *UserActivity) GetCorrespondenceMoves() UserActivityCorrespondenceMoves {
	if o == nil || IsNil(o.CorrespondenceMoves) {
		var ret UserActivityCorrespondenceMoves
		return ret
	}
	return *o.CorrespondenceMoves
}

// GetCorrespondenceMovesOk returns a tuple with the CorrespondenceMoves field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetCorrespondenceMovesOk() (*UserActivityCorrespondenceMoves, bool) {
	if o == nil || IsNil(o.CorrespondenceMoves) {
		return nil, false
	}
	return o.CorrespondenceMoves, true
}

// HasCorrespondenceMoves returns a boolean if a field has been set.
func (o *UserActivity) HasCorrespondenceMoves() bool {
	if o != nil && !IsNil(o.CorrespondenceMoves) {
		return true
	}

	return false
}

// SetCorrespondenceMoves gets a reference to the given UserActivityCorrespondenceMoves and assigns it to the CorrespondenceMoves field.
func (o *UserActivity) SetCorrespondenceMoves(v UserActivityCorrespondenceMoves) {
	o.CorrespondenceMoves = &v
}

// GetCorrespondenceEnds returns the CorrespondenceEnds field value if set, zero value otherwise.
func (o *UserActivity) GetCorrespondenceEnds() UserActivityCorrespondenceEnds {
	if o == nil || IsNil(o.CorrespondenceEnds) {
		var ret UserActivityCorrespondenceEnds
		return ret
	}
	return *o.CorrespondenceEnds
}

// GetCorrespondenceEndsOk returns a tuple with the CorrespondenceEnds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetCorrespondenceEndsOk() (*UserActivityCorrespondenceEnds, bool) {
	if o == nil || IsNil(o.CorrespondenceEnds) {
		return nil, false
	}
	return o.CorrespondenceEnds, true
}

// HasCorrespondenceEnds returns a boolean if a field has been set.
func (o *UserActivity) HasCorrespondenceEnds() bool {
	if o != nil && !IsNil(o.CorrespondenceEnds) {
		return true
	}

	return false
}

// SetCorrespondenceEnds gets a reference to the given UserActivityCorrespondenceEnds and assigns it to the CorrespondenceEnds field.
func (o *UserActivity) SetCorrespondenceEnds(v UserActivityCorrespondenceEnds) {
	o.CorrespondenceEnds = &v
}

// GetFollows returns the Follows field value if set, zero value otherwise.
func (o *UserActivity) GetFollows() UserActivityFollows {
	if o == nil || IsNil(o.Follows) {
		var ret UserActivityFollows
		return ret
	}
	return *o.Follows
}

// GetFollowsOk returns a tuple with the Follows field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetFollowsOk() (*UserActivityFollows, bool) {
	if o == nil || IsNil(o.Follows) {
		return nil, false
	}
	return o.Follows, true
}

// HasFollows returns a boolean if a field has been set.
func (o *UserActivity) HasFollows() bool {
	if o != nil && !IsNil(o.Follows) {
		return true
	}

	return false
}

// SetFollows gets a reference to the given UserActivityFollows and assigns it to the Follows field.
func (o *UserActivity) SetFollows(v UserActivityFollows) {
	o.Follows = &v
}

// GetStudies returns the Studies field value if set, zero value otherwise.
func (o *UserActivity) GetStudies() []UserActivityTournamentsBestInnerTournament {
	if o == nil || IsNil(o.Studies) {
		var ret []UserActivityTournamentsBestInnerTournament
		return ret
	}
	return o.Studies
}

// GetStudiesOk returns a tuple with the Studies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetStudiesOk() ([]UserActivityTournamentsBestInnerTournament, bool) {
	if o == nil || IsNil(o.Studies) {
		return nil, false
	}
	return o.Studies, true
}

// HasStudies returns a boolean if a field has been set.
func (o *UserActivity) HasStudies() bool {
	if o != nil && !IsNil(o.Studies) {
		return true
	}

	return false
}

// SetStudies gets a reference to the given []UserActivityTournamentsBestInnerTournament and assigns it to the Studies field.
func (o *UserActivity) SetStudies(v []UserActivityTournamentsBestInnerTournament) {
	o.Studies = v
}

// GetTeams returns the Teams field value if set, zero value otherwise.
func (o *UserActivity) GetTeams() []UserActivityTeamsInner {
	if o == nil || IsNil(o.Teams) {
		var ret []UserActivityTeamsInner
		return ret
	}
	return o.Teams
}

// GetTeamsOk returns a tuple with the Teams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetTeamsOk() ([]UserActivityTeamsInner, bool) {
	if o == nil || IsNil(o.Teams) {
		return nil, false
	}
	return o.Teams, true
}

// HasTeams returns a boolean if a field has been set.
func (o *UserActivity) HasTeams() bool {
	if o != nil && !IsNil(o.Teams) {
		return true
	}

	return false
}

// SetTeams gets a reference to the given []UserActivityTeamsInner and assigns it to the Teams field.
func (o *UserActivity) SetTeams(v []UserActivityTeamsInner) {
	o.Teams = v
}

// GetPosts returns the Posts field value if set, zero value otherwise.
func (o *UserActivity) GetPosts() []UserActivityPostsInner {
	if o == nil || IsNil(o.Posts) {
		var ret []UserActivityPostsInner
		return ret
	}
	return o.Posts
}

// GetPostsOk returns a tuple with the Posts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetPostsOk() ([]UserActivityPostsInner, bool) {
	if o == nil || IsNil(o.Posts) {
		return nil, false
	}
	return o.Posts, true
}

// HasPosts returns a boolean if a field has been set.
func (o *UserActivity) HasPosts() bool {
	if o != nil && !IsNil(o.Posts) {
		return true
	}

	return false
}

// SetPosts gets a reference to the given []UserActivityPostsInner and assigns it to the Posts field.
func (o *UserActivity) SetPosts(v []UserActivityPostsInner) {
	o.Posts = v
}

// GetPatron returns the Patron field value if set, zero value otherwise.
func (o *UserActivity) GetPatron() UserActivityPatron {
	if o == nil || IsNil(o.Patron) {
		var ret UserActivityPatron
		return ret
	}
	return *o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetPatronOk() (*UserActivityPatron, bool) {
	if o == nil || IsNil(o.Patron) {
		return nil, false
	}
	return o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *UserActivity) HasPatron() bool {
	if o != nil && !IsNil(o.Patron) {
		return true
	}

	return false
}

// SetPatron gets a reference to the given UserActivityPatron and assigns it to the Patron field.
func (o *UserActivity) SetPatron(v UserActivityPatron) {
	o.Patron = &v
}

// GetStream returns the Stream field value if set, zero value otherwise.
func (o *UserActivity) GetStream() bool {
	if o == nil || IsNil(o.Stream) {
		var ret bool
		return ret
	}
	return *o.Stream
}

// GetStreamOk returns a tuple with the Stream field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserActivity) GetStreamOk() (*bool, bool) {
	if o == nil || IsNil(o.Stream) {
		return nil, false
	}
	return o.Stream, true
}

// HasStream returns a boolean if a field has been set.
func (o *UserActivity) HasStream() bool {
	if o != nil && !IsNil(o.Stream) {
		return true
	}

	return false
}

// SetStream gets a reference to the given bool and assigns it to the Stream field.
func (o *UserActivity) SetStream(v bool) {
	o.Stream = &v
}

func (o UserActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserActivity) ToMap() (map[string]interface{}, error) {
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

func (o *UserActivity) UnmarshalJSON(data []byte) (err error) {
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

	varUserActivity := _UserActivity{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserActivity)

	if err != nil {
		return err
	}

	*o = UserActivity(varUserActivity)

	return err
}

type NullableUserActivity struct {
	value *UserActivity
	isSet bool
}

func (v NullableUserActivity) Get() *UserActivity {
	return v.value
}

func (v *NullableUserActivity) Set(val *UserActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableUserActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableUserActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserActivity(val *UserActivity) *NullableUserActivity {
	return &NullableUserActivity{value: val, isSet: true}
}

func (v NullableUserActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


