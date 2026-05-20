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
)

// checks if the Tournament200ResponsePodiumInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tournament200ResponsePodiumInner{}

// Tournament200ResponsePodiumInner struct for Tournament200ResponsePodiumInner
type Tournament200ResponsePodiumInner struct {
	Name *string `json:"name,omitempty"`
	// only appears if the user is a titled player or a bot user
	Title *string `json:"title,omitempty"`
	// Use patronColor value instead to determine if player is a patron. 
	// Deprecated
	Patron *bool `json:"patron,omitempty"`
	// Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron. 
	PatronColor *int32 `json:"patronColor,omitempty"`
	// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
	Flair *string `json:"flair,omitempty"`
	Rank *int32 `json:"rank,omitempty"`
	Rating *int32 `json:"rating,omitempty"`
	Score *int32 `json:"score,omitempty"`
	Nb *ApiTournamentPost200ResponsePodiumInnerNb `json:"nb,omitempty"`
	Performance *int32 `json:"performance,omitempty"`
}

// NewTournament200ResponsePodiumInner instantiates a new Tournament200ResponsePodiumInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTournament200ResponsePodiumInner() *Tournament200ResponsePodiumInner {
	this := Tournament200ResponsePodiumInner{}
	return &this
}

// NewTournament200ResponsePodiumInnerWithDefaults instantiates a new Tournament200ResponsePodiumInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTournament200ResponsePodiumInnerWithDefaults() *Tournament200ResponsePodiumInner {
	this := Tournament200ResponsePodiumInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Tournament200ResponsePodiumInner) SetName(v string) {
	o.Name = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Tournament200ResponsePodiumInner) SetTitle(v string) {
	o.Title = &v
}

// GetPatron returns the Patron field value if set, zero value otherwise.
// Deprecated
func (o *Tournament200ResponsePodiumInner) GetPatron() bool {
	if o == nil || IsNil(o.Patron) {
		var ret bool
		return ret
	}
	return *o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *Tournament200ResponsePodiumInner) GetPatronOk() (*bool, bool) {
	if o == nil || IsNil(o.Patron) {
		return nil, false
	}
	return o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasPatron() bool {
	if o != nil && !IsNil(o.Patron) {
		return true
	}

	return false
}

// SetPatron gets a reference to the given bool and assigns it to the Patron field.
// Deprecated
func (o *Tournament200ResponsePodiumInner) SetPatron(v bool) {
	o.Patron = &v
}

// GetPatronColor returns the PatronColor field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetPatronColor() int32 {
	if o == nil || IsNil(o.PatronColor) {
		var ret int32
		return ret
	}
	return *o.PatronColor
}

// GetPatronColorOk returns a tuple with the PatronColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetPatronColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PatronColor) {
		return nil, false
	}
	return o.PatronColor, true
}

// HasPatronColor returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasPatronColor() bool {
	if o != nil && !IsNil(o.PatronColor) {
		return true
	}

	return false
}

// SetPatronColor gets a reference to the given int32 and assigns it to the PatronColor field.
func (o *Tournament200ResponsePodiumInner) SetPatronColor(v int32) {
	o.PatronColor = &v
}

// GetFlair returns the Flair field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetFlair() string {
	if o == nil || IsNil(o.Flair) {
		var ret string
		return ret
	}
	return *o.Flair
}

// GetFlairOk returns a tuple with the Flair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetFlairOk() (*string, bool) {
	if o == nil || IsNil(o.Flair) {
		return nil, false
	}
	return o.Flair, true
}

// HasFlair returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasFlair() bool {
	if o != nil && !IsNil(o.Flair) {
		return true
	}

	return false
}

// SetFlair gets a reference to the given string and assigns it to the Flair field.
func (o *Tournament200ResponsePodiumInner) SetFlair(v string) {
	o.Flair = &v
}

// GetRank returns the Rank field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetRank() int32 {
	if o == nil || IsNil(o.Rank) {
		var ret int32
		return ret
	}
	return *o.Rank
}

// GetRankOk returns a tuple with the Rank field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetRankOk() (*int32, bool) {
	if o == nil || IsNil(o.Rank) {
		return nil, false
	}
	return o.Rank, true
}

// HasRank returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasRank() bool {
	if o != nil && !IsNil(o.Rank) {
		return true
	}

	return false
}

// SetRank gets a reference to the given int32 and assigns it to the Rank field.
func (o *Tournament200ResponsePodiumInner) SetRank(v int32) {
	o.Rank = &v
}

// GetRating returns the Rating field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetRating() int32 {
	if o == nil || IsNil(o.Rating) {
		var ret int32
		return ret
	}
	return *o.Rating
}

// GetRatingOk returns a tuple with the Rating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.Rating) {
		return nil, false
	}
	return o.Rating, true
}

// HasRating returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasRating() bool {
	if o != nil && !IsNil(o.Rating) {
		return true
	}

	return false
}

// SetRating gets a reference to the given int32 and assigns it to the Rating field.
func (o *Tournament200ResponsePodiumInner) SetRating(v int32) {
	o.Rating = &v
}

// GetScore returns the Score field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetScore() int32 {
	if o == nil || IsNil(o.Score) {
		var ret int32
		return ret
	}
	return *o.Score
}

// GetScoreOk returns a tuple with the Score field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetScoreOk() (*int32, bool) {
	if o == nil || IsNil(o.Score) {
		return nil, false
	}
	return o.Score, true
}

// HasScore returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasScore() bool {
	if o != nil && !IsNil(o.Score) {
		return true
	}

	return false
}

// SetScore gets a reference to the given int32 and assigns it to the Score field.
func (o *Tournament200ResponsePodiumInner) SetScore(v int32) {
	o.Score = &v
}

// GetNb returns the Nb field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetNb() ApiTournamentPost200ResponsePodiumInnerNb {
	if o == nil || IsNil(o.Nb) {
		var ret ApiTournamentPost200ResponsePodiumInnerNb
		return ret
	}
	return *o.Nb
}

// GetNbOk returns a tuple with the Nb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetNbOk() (*ApiTournamentPost200ResponsePodiumInnerNb, bool) {
	if o == nil || IsNil(o.Nb) {
		return nil, false
	}
	return o.Nb, true
}

// HasNb returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasNb() bool {
	if o != nil && !IsNil(o.Nb) {
		return true
	}

	return false
}

// SetNb gets a reference to the given ApiTournamentPost200ResponsePodiumInnerNb and assigns it to the Nb field.
func (o *Tournament200ResponsePodiumInner) SetNb(v ApiTournamentPost200ResponsePodiumInnerNb) {
	o.Nb = &v
}

// GetPerformance returns the Performance field value if set, zero value otherwise.
func (o *Tournament200ResponsePodiumInner) GetPerformance() int32 {
	if o == nil || IsNil(o.Performance) {
		var ret int32
		return ret
	}
	return *o.Performance
}

// GetPerformanceOk returns a tuple with the Performance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tournament200ResponsePodiumInner) GetPerformanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Performance) {
		return nil, false
	}
	return o.Performance, true
}

// HasPerformance returns a boolean if a field has been set.
func (o *Tournament200ResponsePodiumInner) HasPerformance() bool {
	if o != nil && !IsNil(o.Performance) {
		return true
	}

	return false
}

// SetPerformance gets a reference to the given int32 and assigns it to the Performance field.
func (o *Tournament200ResponsePodiumInner) SetPerformance(v int32) {
	o.Performance = &v
}

func (o Tournament200ResponsePodiumInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tournament200ResponsePodiumInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Patron) {
		toSerialize["patron"] = o.Patron
	}
	if !IsNil(o.PatronColor) {
		toSerialize["patronColor"] = o.PatronColor
	}
	if !IsNil(o.Flair) {
		toSerialize["flair"] = o.Flair
	}
	if !IsNil(o.Rank) {
		toSerialize["rank"] = o.Rank
	}
	if !IsNil(o.Rating) {
		toSerialize["rating"] = o.Rating
	}
	if !IsNil(o.Score) {
		toSerialize["score"] = o.Score
	}
	if !IsNil(o.Nb) {
		toSerialize["nb"] = o.Nb
	}
	if !IsNil(o.Performance) {
		toSerialize["performance"] = o.Performance
	}
	return toSerialize, nil
}

type NullableTournament200ResponsePodiumInner struct {
	value *Tournament200ResponsePodiumInner
	isSet bool
}

func (v NullableTournament200ResponsePodiumInner) Get() *Tournament200ResponsePodiumInner {
	return v.value
}

func (v *NullableTournament200ResponsePodiumInner) Set(val *Tournament200ResponsePodiumInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTournament200ResponsePodiumInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTournament200ResponsePodiumInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTournament200ResponsePodiumInner(val *Tournament200ResponsePodiumInner) *NullableTournament200ResponsePodiumInner {
	return &NullableTournament200ResponsePodiumInner{value: val, isSet: true}
}

func (v NullableTournament200ResponsePodiumInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTournament200ResponsePodiumInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


