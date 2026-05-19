/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ApiUserCurrentGame200ResponseOneOfPlayersWhite type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUserCurrentGame200ResponseOneOfPlayersWhite{}

// ApiUserCurrentGame200ResponseOneOfPlayersWhite struct for ApiUserCurrentGame200ResponseOneOfPlayersWhite
type ApiUserCurrentGame200ResponseOneOfPlayersWhite struct {
	User ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser `json:"user"`
	Rating int32 `json:"rating"`
	RatingDiff *int32 `json:"ratingDiff,omitempty"`
	Name *string `json:"name,omitempty"`
	Provisional *bool `json:"provisional,omitempty"`
	AiLevel *int32 `json:"aiLevel,omitempty"`
	Analysis *GamePgn200ResponseOneOfPlayersWhiteAnalysis `json:"analysis,omitempty"`
	Team *string `json:"team,omitempty"`
}

type _ApiUserCurrentGame200ResponseOneOfPlayersWhite ApiUserCurrentGame200ResponseOneOfPlayersWhite

// NewApiUserCurrentGame200ResponseOneOfPlayersWhite instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhite object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUserCurrentGame200ResponseOneOfPlayersWhite(user ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser, rating int32) *ApiUserCurrentGame200ResponseOneOfPlayersWhite {
	this := ApiUserCurrentGame200ResponseOneOfPlayersWhite{}
	this.User = user
	this.Rating = rating
	return &this
}

// NewApiUserCurrentGame200ResponseOneOfPlayersWhiteWithDefaults instantiates a new ApiUserCurrentGame200ResponseOneOfPlayersWhite object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUserCurrentGame200ResponseOneOfPlayersWhiteWithDefaults() *ApiUserCurrentGame200ResponseOneOfPlayersWhite {
	this := ApiUserCurrentGame200ResponseOneOfPlayersWhite{}
	return &this
}

// GetUser returns the User field value
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetUser() ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser {
	if o == nil {
		var ret ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetUserOk() (*ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetUser(v ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser) {
	o.User = v
}

// GetRating returns the Rating field value
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRating() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Rating
}

// GetRatingOk returns a tuple with the Rating field value
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rating, true
}

// SetRating sets field value
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetRating(v int32) {
	o.Rating = v
}

// GetRatingDiff returns the RatingDiff field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingDiff() int32 {
	if o == nil || IsNil(o.RatingDiff) {
		var ret int32
		return ret
	}
	return *o.RatingDiff
}

// GetRatingDiffOk returns a tuple with the RatingDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetRatingDiffOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingDiff) {
		return nil, false
	}
	return o.RatingDiff, true
}

// HasRatingDiff returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasRatingDiff() bool {
	if o != nil && !IsNil(o.RatingDiff) {
		return true
	}

	return false
}

// SetRatingDiff gets a reference to the given int32 and assigns it to the RatingDiff field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetRatingDiff(v int32) {
	o.RatingDiff = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetName(v string) {
	o.Name = &v
}

// GetProvisional returns the Provisional field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetProvisional() bool {
	if o == nil || IsNil(o.Provisional) {
		var ret bool
		return ret
	}
	return *o.Provisional
}

// GetProvisionalOk returns a tuple with the Provisional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetProvisionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Provisional) {
		return nil, false
	}
	return o.Provisional, true
}

// HasProvisional returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasProvisional() bool {
	if o != nil && !IsNil(o.Provisional) {
		return true
	}

	return false
}

// SetProvisional gets a reference to the given bool and assigns it to the Provisional field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetProvisional(v bool) {
	o.Provisional = &v
}

// GetAiLevel returns the AiLevel field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAiLevel() int32 {
	if o == nil || IsNil(o.AiLevel) {
		var ret int32
		return ret
	}
	return *o.AiLevel
}

// GetAiLevelOk returns a tuple with the AiLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAiLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.AiLevel) {
		return nil, false
	}
	return o.AiLevel, true
}

// HasAiLevel returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasAiLevel() bool {
	if o != nil && !IsNil(o.AiLevel) {
		return true
	}

	return false
}

// SetAiLevel gets a reference to the given int32 and assigns it to the AiLevel field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetAiLevel(v int32) {
	o.AiLevel = &v
}

// GetAnalysis returns the Analysis field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAnalysis() GamePgn200ResponseOneOfPlayersWhiteAnalysis {
	if o == nil || IsNil(o.Analysis) {
		var ret GamePgn200ResponseOneOfPlayersWhiteAnalysis
		return ret
	}
	return *o.Analysis
}

// GetAnalysisOk returns a tuple with the Analysis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetAnalysisOk() (*GamePgn200ResponseOneOfPlayersWhiteAnalysis, bool) {
	if o == nil || IsNil(o.Analysis) {
		return nil, false
	}
	return o.Analysis, true
}

// HasAnalysis returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasAnalysis() bool {
	if o != nil && !IsNil(o.Analysis) {
		return true
	}

	return false
}

// SetAnalysis gets a reference to the given GamePgn200ResponseOneOfPlayersWhiteAnalysis and assigns it to the Analysis field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetAnalysis(v GamePgn200ResponseOneOfPlayersWhiteAnalysis) {
	o.Analysis = &v
}

// GetTeam returns the Team field value if set, zero value otherwise.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetTeam() string {
	if o == nil || IsNil(o.Team) {
		var ret string
		return ret
	}
	return *o.Team
}

// GetTeamOk returns a tuple with the Team field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) GetTeamOk() (*string, bool) {
	if o == nil || IsNil(o.Team) {
		return nil, false
	}
	return o.Team, true
}

// HasTeam returns a boolean if a field has been set.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) HasTeam() bool {
	if o != nil && !IsNil(o.Team) {
		return true
	}

	return false
}

// SetTeam gets a reference to the given string and assigns it to the Team field.
func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) SetTeam(v string) {
	o.Team = &v
}

func (o ApiUserCurrentGame200ResponseOneOfPlayersWhite) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUserCurrentGame200ResponseOneOfPlayersWhite) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["user"] = o.User
	toSerialize["rating"] = o.Rating
	if !IsNil(o.RatingDiff) {
		toSerialize["ratingDiff"] = o.RatingDiff
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Provisional) {
		toSerialize["provisional"] = o.Provisional
	}
	if !IsNil(o.AiLevel) {
		toSerialize["aiLevel"] = o.AiLevel
	}
	if !IsNil(o.Analysis) {
		toSerialize["analysis"] = o.Analysis
	}
	if !IsNil(o.Team) {
		toSerialize["team"] = o.Team
	}
	return toSerialize, nil
}

func (o *ApiUserCurrentGame200ResponseOneOfPlayersWhite) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user",
		"rating",
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

	varApiUserCurrentGame200ResponseOneOfPlayersWhite := _ApiUserCurrentGame200ResponseOneOfPlayersWhite{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUserCurrentGame200ResponseOneOfPlayersWhite)

	if err != nil {
		return err
	}

	*o = ApiUserCurrentGame200ResponseOneOfPlayersWhite(varApiUserCurrentGame200ResponseOneOfPlayersWhite)

	return err
}

type NullableApiUserCurrentGame200ResponseOneOfPlayersWhite struct {
	value *ApiUserCurrentGame200ResponseOneOfPlayersWhite
	isSet bool
}

func (v NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) Get() *ApiUserCurrentGame200ResponseOneOfPlayersWhite {
	return v.value
}

func (v *NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) Set(val *ApiUserCurrentGame200ResponseOneOfPlayersWhite) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUserCurrentGame200ResponseOneOfPlayersWhite(val *ApiUserCurrentGame200ResponseOneOfPlayersWhite) *NullableApiUserCurrentGame200ResponseOneOfPlayersWhite {
	return &NullableApiUserCurrentGame200ResponseOneOfPlayersWhite{value: val, isSet: true}
}

func (v NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUserCurrentGame200ResponseOneOfPlayersWhite) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


