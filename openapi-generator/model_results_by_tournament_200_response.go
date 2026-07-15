/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.153
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ResultsByTournament200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResultsByTournament200Response{}

// ResultsByTournament200Response struct for ResultsByTournament200Response
type ResultsByTournament200Response struct {
	Rank int32 `json:"rank"`
	Score int32 `json:"score"`
	Rating int32 `json:"rating"`
	Username string `json:"username"`
	Performance int32 `json:"performance"`
	Title *Title `json:"title,omitempty"`
	Team *string `json:"team,omitempty"`
	// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
	Flair *string `json:"flair,omitempty"`
	// Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron. 
	PatronColor *int32 `json:"patronColor,omitempty"`
	Sheet *ArenaSheet `json:"sheet,omitempty"`
}

type _ResultsByTournament200Response ResultsByTournament200Response

// NewResultsByTournament200Response instantiates a new ResultsByTournament200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResultsByTournament200Response(rank int32, score int32, rating int32, username string, performance int32) *ResultsByTournament200Response {
	this := ResultsByTournament200Response{}
	this.Rank = rank
	this.Score = score
	this.Rating = rating
	this.Username = username
	this.Performance = performance
	return &this
}

// NewResultsByTournament200ResponseWithDefaults instantiates a new ResultsByTournament200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResultsByTournament200ResponseWithDefaults() *ResultsByTournament200Response {
	this := ResultsByTournament200Response{}
	return &this
}

// GetRank returns the Rank field value
func (o *ResultsByTournament200Response) GetRank() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rank
}

// GetRankOk returns a tuple with the Rank field value
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetRankOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rank, true
}

// SetRank sets field value
func (o *ResultsByTournament200Response) SetRank(v int32) {
	o.Rank = v
}

// GetScore returns the Score field value
func (o *ResultsByTournament200Response) GetScore() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Score
}

// GetScoreOk returns a tuple with the Score field value
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetScoreOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Score, true
}

// SetScore sets field value
func (o *ResultsByTournament200Response) SetScore(v int32) {
	o.Score = v
}

// GetRating returns the Rating field value
func (o *ResultsByTournament200Response) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *ResultsByTournament200Response) SetRating(v int32) {
	o.Rating = v
}

// GetUsername returns the Username field value
func (o *ResultsByTournament200Response) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *ResultsByTournament200Response) SetUsername(v string) {
	o.Username = v
}

// GetPerformance returns the Performance field value
func (o *ResultsByTournament200Response) GetPerformance() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetPerformanceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Performance, true
}

// SetPerformance sets field value
func (o *ResultsByTournament200Response) SetPerformance(v int32) {
	o.Performance = v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ResultsByTournament200Response) GetTitle() Title {
	if o == nil || IsNil(o.Title) {
		var ret Title
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetTitleOk() (*Title, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ResultsByTournament200Response) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given Title and assigns it to the Title field.
func (o *ResultsByTournament200Response) SetTitle(v Title) {
	o.Title = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *ResultsByTournament200Response) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *ResultsByTournament200Response) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *ResultsByTournament200Response) SetTeam(v string) {
	o.Team = &v
}

// GetFlair returns the Flair field value if set, zero value otherwise.
func (o *ResultsByTournament200Response) GetFlair() string {
	if o == nil || IsNil(o.Flair) {
		var ret string
		return ret
	}
	return *o.Flair
}

// GetFlairOk returns a tuple with the Flair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetFlairOk() (*string, bool) {
	if o == nil || IsNil(o.Flair) {
		return nil, false
	}
	return o.Flair, true
}

// HasFlair returns a boolean if a field has been set.
func (o *ResultsByTournament200Response) HasFlair() bool {
	if o != nil && !IsNil(o.Flair) {
		return true
	}

	return false
}

// SetFlair gets a reference to the given string and assigns it to the Flair field.
func (o *ResultsByTournament200Response) SetFlair(v string) {
	o.Flair = &v
}

// GetPatronColor returns the PatronColor field value if set, zero value otherwise.
func (o *ResultsByTournament200Response) GetPatronColor() int32 {
	if o == nil || IsNil(o.PatronColor) {
		var ret int32
		return ret
	}
	return *o.PatronColor
}

// GetPatronColorOk returns a tuple with the PatronColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetPatronColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PatronColor) {
		return nil, false
	}
	return o.PatronColor, true
}

// HasPatronColor returns a boolean if a field has been set.
func (o *ResultsByTournament200Response) HasPatronColor() bool {
	if o != nil && !IsNil(o.PatronColor) {
		return true
	}

	return false
}

// SetPatronColor gets a reference to the given int32 and assigns it to the PatronColor field.
func (o *ResultsByTournament200Response) SetPatronColor(v int32) {
	o.PatronColor = &v
}

// GetSheet returns the Sheet field value if set, zero value otherwise.
func (o *ResultsByTournament200Response) GetSheet() ArenaSheet {
	if o == nil || IsNil(o.Sheet) {
		var ret ArenaSheet
		return ret
	}
	return *o.Sheet
}

// GetSheetOk returns a tuple with the Sheet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResultsByTournament200Response) GetSheetOk() (*ArenaSheet, bool) {
	if o == nil || IsNil(o.Sheet) {
		return nil, false
	}
	return o.Sheet, true
}

// HasSheet returns a boolean if a field has been set.
func (o *ResultsByTournament200Response) HasSheet() bool {
	if o != nil && !IsNil(o.Sheet) {
		return true
	}

	return false
}

// SetSheet gets a reference to the given ArenaSheet and assigns it to the Sheet field.
func (o *ResultsByTournament200Response) SetSheet(v ArenaSheet) {
	o.Sheet = &v
}

func (o ResultsByTournament200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResultsByTournament200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rank"] = o.Rank
	toSerialize["score"] = o.Score
	toSerialize["rating"] = o.Rating
	toSerialize["username"] = o.Username
	toSerialize["performance"] = o.Performance
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	if !IsNil(o.Flair) {
		toSerialize["flair"] = o.Flair
	}
	if !IsNil(o.PatronColor) {
		toSerialize["patronColor"] = o.PatronColor
	}
	if !IsNil(o.Sheet) {
		toSerialize["sheet"] = o.Sheet
	}
	return toSerialize, nil
}

func (o *ResultsByTournament200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rank",
		"score",
		"rating",
		"username",
		"performance",
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

	varResultsByTournament200Response := _ResultsByTournament200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varResultsByTournament200Response)

	if err != nil {
		return err
	}

	*o = ResultsByTournament200Response(varResultsByTournament200Response)

	return err
}

type NullableResultsByTournament200Response struct {
	value *ResultsByTournament200Response
	isSet bool
}

func (v NullableResultsByTournament200Response) Get() *ResultsByTournament200Response {
	return v.value
}

func (v *NullableResultsByTournament200Response) Set(val *ResultsByTournament200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableResultsByTournament200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableResultsByTournament200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultsByTournament200Response(val *ResultsByTournament200Response) *NullableResultsByTournament200Response {
	return &NullableResultsByTournament200Response{value: val, isSet: true}
}

func (v NullableResultsByTournament200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultsByTournament200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


