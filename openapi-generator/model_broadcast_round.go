/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api-demo/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api-demo/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api-demo/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.174
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastRound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastRound{}

// BroadcastRound struct for BroadcastRound
type BroadcastRound struct {
	Round BroadcastRoundInfo `json:"round"`
	Tour BroadcastTour `json:"tour"`
	Study BroadcastRoundStudyInfo `json:"study"`
	Games []BroadcastRoundGame `json:"games"`
	Group *BroadcastGroup `json:"group,omitempty"`
	// Indicates if the user making the request is subscribed to the broadcast
	IsSubscribed *bool `json:"isSubscribed,omitempty"`
	// Photos of players, when available. The object keys are FIDE IDs
	Photos map[string]BroadcastPhotosValue `json:"photos"`
}

type _BroadcastRound BroadcastRound

// NewBroadcastRound instantiates a new BroadcastRound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastRound(round BroadcastRoundInfo, tour BroadcastTour, study BroadcastRoundStudyInfo, games []BroadcastRoundGame, photos map[string]BroadcastPhotosValue) *BroadcastRound {
	this := BroadcastRound{}
	this.Round = round
	this.Tour = tour
	this.Study = study
	this.Games = games
	this.Photos = photos
	return &this
}

// NewBroadcastRoundWithDefaults instantiates a new BroadcastRound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastRoundWithDefaults() *BroadcastRound {
	this := BroadcastRound{}
	return &this
}

// GetRound returns the Round field value
func (o *BroadcastRound) GetRound() BroadcastRoundInfo {
	if o == nil {
		var ret BroadcastRoundInfo
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetRoundOk() (*BroadcastRoundInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *BroadcastRound) SetRound(v BroadcastRoundInfo) {
	o.Round = v
}

// GetTour returns the Tour field value
func (o *BroadcastRound) GetTour() BroadcastTour {
	if o == nil {
		var ret BroadcastTour
		return ret
	}

	return o.Tour
}

// GetTourOk returns a tuple with the Tour field value
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetTourOk() (*BroadcastTour, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tour, true
}

// SetTour sets field value
func (o *BroadcastRound) SetTour(v BroadcastTour) {
	o.Tour = v
}

// GetStudy returns the Study field value
func (o *BroadcastRound) GetStudy() BroadcastRoundStudyInfo {
	if o == nil {
		var ret BroadcastRoundStudyInfo
		return ret
	}

	return o.Study
}

// GetStudyOk returns a tuple with the Study field value
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetStudyOk() (*BroadcastRoundStudyInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Study, true
}

// SetStudy sets field value
func (o *BroadcastRound) SetStudy(v BroadcastRoundStudyInfo) {
	o.Study = v
}

// GetGames returns the Games field value
func (o *BroadcastRound) GetGames() []BroadcastRoundGame {
	if o == nil {
		var ret []BroadcastRoundGame
		return ret
	}

	return o.Games
}

// GetGamesOk returns a tuple with the Games field value
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetGamesOk() ([]BroadcastRoundGame, bool) {
	if o == nil {
		return nil, false
	}
	return o.Games, true
}

// SetGames sets field value
func (o *BroadcastRound) SetGames(v []BroadcastRoundGame) {
	o.Games = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BroadcastRound) GetGroup() BroadcastGroup {
	if o == nil || IsNil(o.Group) {
		var ret BroadcastGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetGroupOk() (*BroadcastGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BroadcastRound) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BroadcastGroup and assigns it to the Group field.
func (o *BroadcastRound) SetGroup(v BroadcastGroup) {
	o.Group = &v
}

// GetIsSubscribed returns the IsSubscribed field value if set, zero value otherwise.
func (o *BroadcastRound) GetIsSubscribed() bool {
	if o == nil || IsNil(o.IsSubscribed) {
		var ret bool
		return ret
	}
	return *o.IsSubscribed
}

// GetIsSubscribedOk returns a tuple with the IsSubscribed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetIsSubscribedOk() (*bool, bool) {
	if o == nil || IsNil(o.IsSubscribed) {
		return nil, false
	}
	return o.IsSubscribed, true
}

// HasIsSubscribed returns a boolean if a field has been set.
func (o *BroadcastRound) HasIsSubscribed() bool {
	if o != nil && !IsNil(o.IsSubscribed) {
		return true
	}

	return false
}

// SetIsSubscribed gets a reference to the given bool and assigns it to the IsSubscribed field.
func (o *BroadcastRound) SetIsSubscribed(v bool) {
	o.IsSubscribed = &v
}

// GetPhotos returns the Photos field value
func (o *BroadcastRound) GetPhotos() map[string]BroadcastPhotosValue {
	if o == nil {
		var ret map[string]BroadcastPhotosValue
		return ret
	}

	return o.Photos
}

// GetPhotosOk returns a tuple with the Photos field value
// and a boolean to check if the value has been set.
func (o *BroadcastRound) GetPhotosOk() (map[string]BroadcastPhotosValue, bool) {
	if o == nil {
		return map[string]BroadcastPhotosValue{}, false
	}
	return o.Photos, true
}

// SetPhotos sets field value
func (o *BroadcastRound) SetPhotos(v map[string]BroadcastPhotosValue) {
	o.Photos = v
}

func (o BroadcastRound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastRound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["round"] = o.Round
	toSerialize["tour"] = o.Tour
	toSerialize["study"] = o.Study
	toSerialize["games"] = o.Games
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	if !IsNil(o.IsSubscribed) {
		toSerialize["isSubscribed"] = o.IsSubscribed
	}
	toSerialize["photos"] = o.Photos
	return toSerialize, nil
}

func (o *BroadcastRound) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"round",
		"tour",
		"study",
		"games",
		"photos",
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

	varBroadcastRound := _BroadcastRound{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastRound)

	if err != nil {
		return err
	}

	*o = BroadcastRound(varBroadcastRound)

	return err
}

type NullableBroadcastRound struct {
	value *BroadcastRound
	isSet bool
}

func (v NullableBroadcastRound) Get() *BroadcastRound {
	return v.value
}

func (v *NullableBroadcastRound) Set(val *BroadcastRound) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastRound) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastRound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastRound(val *BroadcastRound) *NullableBroadcastRound {
	return &NullableBroadcastRound{value: val, isSet: true}
}

func (v NullableBroadcastRound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastRound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


