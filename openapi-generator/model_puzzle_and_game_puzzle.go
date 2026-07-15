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

// checks if the PuzzleAndGamePuzzle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PuzzleAndGamePuzzle{}

// PuzzleAndGamePuzzle struct for PuzzleAndGamePuzzle
type PuzzleAndGamePuzzle struct {
	Id string `json:"id"`
	InitialPly int32 `json:"initialPly"`
	Plays int32 `json:"plays"`
	Rating int32 `json:"rating"`
	Fen *string `json:"fen,omitempty"`
	// In UCI format, e.g. \"e2e4\"
	LastMove *string `json:"lastMove,omitempty"`
	Solution []string `json:"solution"`
	Themes []string `json:"themes"`
}

type _PuzzleAndGamePuzzle PuzzleAndGamePuzzle

// NewPuzzleAndGamePuzzle instantiates a new PuzzleAndGamePuzzle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPuzzleAndGamePuzzle(id string, initialPly int32, plays int32, rating int32, solution []string, themes []string) *PuzzleAndGamePuzzle {
	this := PuzzleAndGamePuzzle{}
	this.Id = id
	this.InitialPly = initialPly
	this.Plays = plays
	this.Rating = rating
	this.Solution = solution
	this.Themes = themes
	return &this
}

// NewPuzzleAndGamePuzzleWithDefaults instantiates a new PuzzleAndGamePuzzle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPuzzleAndGamePuzzleWithDefaults() *PuzzleAndGamePuzzle {
	this := PuzzleAndGamePuzzle{}
	return &this
}

// GetId returns the Id field value
func (o *PuzzleAndGamePuzzle) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PuzzleAndGamePuzzle) SetId(v string) {
	o.Id = v
}

// GetInitialPly returns the InitialPly field value
func (o *PuzzleAndGamePuzzle) GetInitialPly() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.InitialPly
}

// GetInitialPlyOk returns a tuple with the InitialPly field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetInitialPlyOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitialPly, true
}

// SetInitialPly sets field value
func (o *PuzzleAndGamePuzzle) SetInitialPly(v int32) {
	o.InitialPly = v
}

// GetPlays returns the Plays field value
func (o *PuzzleAndGamePuzzle) GetPlays() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Plays
}

// GetPlaysOk returns a tuple with the Plays field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetPlaysOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Plays, true
}

// SetPlays sets field value
func (o *PuzzleAndGamePuzzle) SetPlays(v int32) {
	o.Plays = v
}

// GetRating returns the Rating field value
func (o *PuzzleAndGamePuzzle) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *PuzzleAndGamePuzzle) SetRating(v int32) {
	o.Rating = v
}

// GetFen returns the Fen field value if set, zero value otherwise.
func (o *PuzzleAndGamePuzzle) GetFen() string {
	if o == nil || IsNil(o.Fen) {
		var ret string
		return ret
	}
	return *o.Fen
}

// GetFenOk returns a tuple with the Fen field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetFenOk() (*string, bool) {
	if o == nil || IsNil(o.Fen) {
		return nil, false
	}
	return o.Fen, true
}

// HasFen returns a boolean if a field has been set.
func (o *PuzzleAndGamePuzzle) HasFen() bool {
	if o != nil && !IsNil(o.Fen) {
		return true
	}

	return false
}

// SetFen gets a reference to the given string and assigns it to the Fen field.
func (o *PuzzleAndGamePuzzle) SetFen(v string) {
	o.Fen = &v
}

// GetLastMove returns the LastMove field value if set, zero value otherwise.
func (o *PuzzleAndGamePuzzle) GetLastMove() string {
	if o == nil || IsNil(o.LastMove) {
		var ret string
		return ret
	}
	return *o.LastMove
}

// GetLastMoveOk returns a tuple with the LastMove field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetLastMoveOk() (*string, bool) {
	if o == nil || IsNil(o.LastMove) {
		return nil, false
	}
	return o.LastMove, true
}

// HasLastMove returns a boolean if a field has been set.
func (o *PuzzleAndGamePuzzle) HasLastMove() bool {
	if o != nil && !IsNil(o.LastMove) {
		return true
	}

	return false
}

// SetLastMove gets a reference to the given string and assigns it to the LastMove field.
func (o *PuzzleAndGamePuzzle) SetLastMove(v string) {
	o.LastMove = &v
}

// GetSolution returns the Solution field value
func (o *PuzzleAndGamePuzzle) GetSolution() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Solution
}

// GetSolutionOk returns a tuple with the Solution field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetSolutionOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Solution, true
}

// SetSolution sets field value
func (o *PuzzleAndGamePuzzle) SetSolution(v []string) {
	o.Solution = v
}

// GetThemes returns the Themes field value
func (o *PuzzleAndGamePuzzle) GetThemes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Themes
}

// GetThemesOk returns a tuple with the Themes field value
// and a boolean to check if the value has been set.
func (o *PuzzleAndGamePuzzle) GetThemesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Themes, true
}

// SetThemes sets field value
func (o *PuzzleAndGamePuzzle) SetThemes(v []string) {
	o.Themes = v
}

func (o PuzzleAndGamePuzzle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PuzzleAndGamePuzzle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["initialPly"] = o.InitialPly
	toSerialize["plays"] = o.Plays
	toSerialize["rating"] = o.Rating
	if !IsNil(o.Fen) {
		toSerialize["fen"] = o.Fen
	}
	if !IsNil(o.LastMove) {
		toSerialize["lastMove"] = o.LastMove
	}
	toSerialize["solution"] = o.Solution
	toSerialize["themes"] = o.Themes
	return toSerialize, nil
}

func (o *PuzzleAndGamePuzzle) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"initialPly",
		"plays",
		"rating",
		"solution",
		"themes",
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

	varPuzzleAndGamePuzzle := _PuzzleAndGamePuzzle{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPuzzleAndGamePuzzle)

	if err != nil {
		return err
	}

	*o = PuzzleAndGamePuzzle(varPuzzleAndGamePuzzle)

	return err
}

type NullablePuzzleAndGamePuzzle struct {
	value *PuzzleAndGamePuzzle
	isSet bool
}

func (v NullablePuzzleAndGamePuzzle) Get() *PuzzleAndGamePuzzle {
	return v.value
}

func (v *NullablePuzzleAndGamePuzzle) Set(val *PuzzleAndGamePuzzle) {
	v.value = val
	v.isSet = true
}

func (v NullablePuzzleAndGamePuzzle) IsSet() bool {
	return v.isSet
}

func (v *NullablePuzzleAndGamePuzzle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePuzzleAndGamePuzzle(val *PuzzleAndGamePuzzle) *NullablePuzzleAndGamePuzzle {
	return &NullablePuzzleAndGamePuzzle{value: val, isSet: true}
}

func (v NullablePuzzleAndGamePuzzle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePuzzleAndGamePuzzle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


