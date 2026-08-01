/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.157
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastPlayerEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastPlayerEntry{}

// BroadcastPlayerEntry struct for BroadcastPlayerEntry
type BroadcastPlayerEntry struct {
	Name string `json:"name"`
	Title *Title `json:"title,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	FideId *int32 `json:"fideId,omitempty"`
	Team *string `json:"team,omitempty"`
	Fed *string `json:"fed,omitempty"`
	Score *float32 `json:"score,omitempty"`
	Played *int32 `json:"played,omitempty"`
	// Rating differences by FIDE time control. 
	RatingDiffs *StatByFideTC `json:"ratingDiffs,omitempty"`
	// Player's ratings at the time of the tournament. 
	RatingsMap *StatByFideTC `json:"ratingsMap,omitempty"`
	// Performance ratings by FIDE time control. 
	Performances *StatByFideTC `json:"performances,omitempty"`
	Tiebreaks []BroadcastPlayerTiebreak `json:"tiebreaks,omitempty"`
	Rank *int32 `json:"rank,omitempty"`
}

type _BroadcastPlayerEntry BroadcastPlayerEntry

// NewBroadcastPlayerEntry instantiates a new BroadcastPlayerEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastPlayerEntry(name string) *BroadcastPlayerEntry {
	this := BroadcastPlayerEntry{}
	this.Name = name
	return &this
}

// NewBroadcastPlayerEntryWithDefaults instantiates a new BroadcastPlayerEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastPlayerEntryWithDefaults() *BroadcastPlayerEntry {
	this := BroadcastPlayerEntry{}
	return &this
}

// GetName returns the Name field value
func (o *BroadcastPlayerEntry) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BroadcastPlayerEntry) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetTitle() Title {
	if o == nil || IsNil(o.Title) {
		var ret Title
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetTitleOk() (*Title, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given Title and assigns it to the Title field.
func (o *BroadcastPlayerEntry) SetTitle(v Title) {
	o.Title = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *BroadcastPlayerEntry) SetRating(v int32) {
	o.Rating = &v
}

// GetFideId returns the FideId field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetFideId() int32 {
	if o == nil || IsNil(o.FideId) {
		var ret int32
		return ret
	}
	return *o.FideId
}

// GetFideIdOk returns a tuple with the FideId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetFideIdOk() (*int32, bool) {
	if o == nil || IsNil(o.FideId) {
		return nil, false
	}
	return o.FideId, true
}

// HasFideId returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasFideId() bool {
	if o != nil && !IsNil(o.FideId) {
		return true
	}

	return false
}

// SetFideId gets a reference to the given int32 and assigns it to the FideId field.
func (o *BroadcastPlayerEntry) SetFideId(v int32) {
	o.FideId = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *BroadcastPlayerEntry) SetTeam(v string) {
	o.Team = &v
}

// GetFed returns the Fed field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetFed() string {
	if o == nil || IsNil(o.Fed) {
		var ret string
		return ret
	}
	return *o.Fed
}

// GetFedOk returns a tuple with the Fed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetFedOk() (*string, bool) {
	if o == nil || IsNil(o.Fed) {
		return nil, false
	}
	return o.Fed, true
}

// HasFed returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasFed() bool {
	if o != nil && !IsNil(o.Fed) {
		return true
	}

	return false
}

// SetFed gets a reference to the given string and assigns it to the Fed field.
func (o *BroadcastPlayerEntry) SetFed(v string) {
	o.Fed = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetScore() float32 {
	if o == nil || IsNil(o.Score) {
		var ret float32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetScoreOk() (*float32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given float32 and assigns it to the Score field.
func (o *BroadcastPlayerEntry) SetScore(v float32) {
	o.Score = &v
}

// GetPlayed returns the Played field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetPlayed() int32 {
	if o == nil || IsNil(o.Played) {
		var ret int32
		return ret
	}
	return *o.Played
}

// GetPlayedOk returns a tuple with the Played field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetPlayedOk() (*int32, bool) {
	if o == nil || IsNil(o.Played) {
		return nil, false
	}
	return o.Played, true
}

// HasPlayed returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasPlayed() bool {
	if o != nil && !IsNil(o.Played) {
		return true
	}

	return false
}

// SetPlayed gets a reference to the given int32 and assigns it to the Played field.
func (o *BroadcastPlayerEntry) SetPlayed(v int32) {
	o.Played = &v
}

// GetRatingDiffs returns the RatingDiffs field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetRatingDiffs() StatByFideTC {
	if o == nil || IsNil(o.RatingDiffs) {
		var ret StatByFideTC
		return ret
	}
	return *o.RatingDiffs
}

// GetRatingDiffsOk returns a tuple with the RatingDiffs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetRatingDiffsOk() (*StatByFideTC, bool) {
	if o == nil || IsNil(o.RatingDiffs) {
		return nil, false
	}
	return o.RatingDiffs, true
}

// HasRatingDiffs returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasRatingDiffs() bool {
	if o != nil && !IsNil(o.RatingDiffs) {
		return true
	}

	return false
}

// SetRatingDiffs gets a reference to the given StatByFideTC and assigns it to the RatingDiffs field.
func (o *BroadcastPlayerEntry) SetRatingDiffs(v StatByFideTC) {
	o.RatingDiffs = &v
}

// GetRatingsMap returns the RatingsMap field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetRatingsMap() StatByFideTC {
	if o == nil || IsNil(o.RatingsMap) {
		var ret StatByFideTC
		return ret
	}
	return *o.RatingsMap
}

// GetRatingsMapOk returns a tuple with the RatingsMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetRatingsMapOk() (*StatByFideTC, bool) {
	if o == nil || IsNil(o.RatingsMap) {
		return nil, false
	}
	return o.RatingsMap, true
}

// HasRatingsMap returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasRatingsMap() bool {
	if o != nil && !IsNil(o.RatingsMap) {
		return true
	}

	return false
}

// SetRatingsMap gets a reference to the given StatByFideTC and assigns it to the RatingsMap field.
func (o *BroadcastPlayerEntry) SetRatingsMap(v StatByFideTC) {
	o.RatingsMap = &v
}

// GetPerformances returns the Performances field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetPerformances() StatByFideTC {
	if o == nil || IsNil(o.Performances) {
		var ret StatByFideTC
		return ret
	}
	return *o.Performances
}

// GetPerformancesOk returns a tuple with the Performances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetPerformancesOk() (*StatByFideTC, bool) {
	if o == nil || IsNil(o.Performances) {
		return nil, false
	}
	return o.Performances, true
}

// HasPerformances returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasPerformances() bool {
	if o != nil && !IsNil(o.Performances) {
		return true
	}

	return false
}

// SetPerformances gets a reference to the given StatByFideTC and assigns it to the Performances field.
func (o *BroadcastPlayerEntry) SetPerformances(v StatByFideTC) {
	o.Performances = &v
}

// GetTiebreaks returns the Tiebreaks field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetTiebreaks() []BroadcastPlayerTiebreak {
	if o == nil || IsNil(o.Tiebreaks) {
		var ret []BroadcastPlayerTiebreak
		return ret
	}
	return o.Tiebreaks
}

// GetTiebreaksOk returns a tuple with the Tiebreaks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetTiebreaksOk() ([]BroadcastPlayerTiebreak, bool) {
	if o == nil || IsNil(o.Tiebreaks) {
		return nil, false
	}
	return o.Tiebreaks, true
}

// HasTiebreaks returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasTiebreaks() bool {
	if o != nil && !IsNil(o.Tiebreaks) {
		return true
	}

	return false
}

// SetTiebreaks gets a reference to the given []BroadcastPlayerTiebreak and assigns it to the Tiebreaks field.
func (o *BroadcastPlayerEntry) SetTiebreaks(v []BroadcastPlayerTiebreak) {
	o.Tiebreaks = v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *BroadcastPlayerEntry) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerEntry) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *BroadcastPlayerEntry) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *BroadcastPlayerEntry) SetRank(v int32) {
	o.Rank = &v
}

func (o BroadcastPlayerEntry) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastPlayerEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.FideId) {
		toSerialize["fideId"] = o.FideId
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.Fed) {
		toSerialize["fed"] = o.Fed
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Played) {
		toSerialize["played"] = o.Played
	}
	if !IsNil(o.RatingDiffs) {
		toSerialize["ratingDiffs"] = o.RatingDiffs
	}
	if !IsNil(o.RatingsMap) {
		toSerialize["ratingsMap"] = o.RatingsMap
	}
	if !IsNil(o.Performances) {
		toSerialize["performances"] = o.Performances
	}
	if !IsNil(o.Tiebreaks) {
		toSerialize["tiebreaks"] = o.Tiebreaks
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	return toSerialize, nil
}

func (o *BroadcastPlayerEntry) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varBroadcastPlayerEntry := _BroadcastPlayerEntry{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastPlayerEntry)

	if err != nil {
		return err
	}

	*o = BroadcastPlayerEntry(varBroadcastPlayerEntry)

	return err
}

type NullableBroadcastPlayerEntry struct {
	value *BroadcastPlayerEntry
	isSet bool
}

func (v NullableBroadcastPlayerEntry) Get() *BroadcastPlayerEntry {
	return v.value
}

func (v *NullableBroadcastPlayerEntry) Set(val *BroadcastPlayerEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastPlayerEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastPlayerEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastPlayerEntry(val *BroadcastPlayerEntry) *NullableBroadcastPlayerEntry {
	return &NullableBroadcastPlayerEntry{value: val, isSet: true}
}

func (v NullableBroadcastPlayerEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastPlayerEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


