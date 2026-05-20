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

// checks if the BulkPairingList200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkPairingList200ResponseInner{}

// BulkPairingList200ResponseInner struct for BulkPairingList200ResponseInner
type BulkPairingList200ResponseInner struct {
	Id string `json:"id"`
	Games []BulkPairingList200ResponseInnerGamesInner `json:"games"`
	Variant string `json:"variant"`
	Clock ApiTournament200ResponseCreatedInnerClock `json:"clock"`
	PairAt int32 `json:"pairAt"`
	PairedAt NullableInt32 `json:"pairedAt"`
	Rated bool `json:"rated"`
	StartClocksAt int32 `json:"startClocksAt"`
	ScheduledAt int32 `json:"scheduledAt"`
}

type _BulkPairingList200ResponseInner BulkPairingList200ResponseInner

// NewBulkPairingList200ResponseInner instantiates a new BulkPairingList200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkPairingList200ResponseInner(id string, games []BulkPairingList200ResponseInnerGamesInner, variant string, clock ApiTournament200ResponseCreatedInnerClock, pairAt int32, pairedAt NullableInt32, rated bool, startClocksAt int32, scheduledAt int32) *BulkPairingList200ResponseInner {
	this := BulkPairingList200ResponseInner{}
	this.Id = id
	this.Games = games
	this.Variant = variant
	this.Clock = clock
	this.PairAt = pairAt
	this.PairedAt = pairedAt
	this.Rated = rated
	this.StartClocksAt = startClocksAt
	this.ScheduledAt = scheduledAt
	return &this
}

// NewBulkPairingList200ResponseInnerWithDefaults instantiates a new BulkPairingList200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkPairingList200ResponseInnerWithDefaults() *BulkPairingList200ResponseInner {
	this := BulkPairingList200ResponseInner{}
	var variant string = "standard"
	this.Variant = variant
	return &this
}

// GetId returns the Id field value
func (o *BulkPairingList200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BulkPairingList200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetGames returns the Games field value
func (o *BulkPairingList200ResponseInner) GetGames() []BulkPairingList200ResponseInnerGamesInner {
	if o == nil {
		var ret []BulkPairingList200ResponseInnerGamesInner
		return ret
	}

	return o.Games
}

// GetGamesOk returns a tuple with the Games field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetGamesOk() ([]BulkPairingList200ResponseInnerGamesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Games, true
}

// SetGames sets field value
func (o *BulkPairingList200ResponseInner) SetGames(v []BulkPairingList200ResponseInnerGamesInner) {
	o.Games = v
}

// GetVariant returns the Variant field value
func (o *BulkPairingList200ResponseInner) GetVariant() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Variant
}

// GetVariantOk returns a tuple with the Variant field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetVariantOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Variant, true
}

// SetVariant sets field value
func (o *BulkPairingList200ResponseInner) SetVariant(v string) {
	o.Variant = v
}

// GetClock returns the Clock field value
func (o *BulkPairingList200ResponseInner) GetClock() ApiTournament200ResponseCreatedInnerClock {
	if o == nil {
		var ret ApiTournament200ResponseCreatedInnerClock
		return ret
	}

	return o.Clock
}

// GetClockOk returns a tuple with the Clock field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetClockOk() (*ApiTournament200ResponseCreatedInnerClock, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clock, true
}

// SetClock sets field value
func (o *BulkPairingList200ResponseInner) SetClock(v ApiTournament200ResponseCreatedInnerClock) {
	o.Clock = v
}

// GetPairAt returns the PairAt field value
func (o *BulkPairingList200ResponseInner) GetPairAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PairAt
}

// GetPairAtOk returns a tuple with the PairAt field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetPairAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PairAt, true
}

// SetPairAt sets field value
func (o *BulkPairingList200ResponseInner) SetPairAt(v int32) {
	o.PairAt = v
}

// GetPairedAt returns the PairedAt field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BulkPairingList200ResponseInner) GetPairedAt() int32 {
	if o == nil || o.PairedAt.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PairedAt.Get()
}

// GetPairedAtOk returns a tuple with the PairedAt field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BulkPairingList200ResponseInner) GetPairedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PairedAt.Get(), o.PairedAt.IsSet()
}

// SetPairedAt sets field value
func (o *BulkPairingList200ResponseInner) SetPairedAt(v int32) {
	o.PairedAt.Set(&v)
}

// GetRated returns the Rated field value
func (o *BulkPairingList200ResponseInner) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *BulkPairingList200ResponseInner) SetRated(v bool) {
	o.Rated = v
}

// GetStartClocksAt returns the StartClocksAt field value
func (o *BulkPairingList200ResponseInner) GetStartClocksAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StartClocksAt
}

// GetStartClocksAtOk returns a tuple with the StartClocksAt field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetStartClocksAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartClocksAt, true
}

// SetStartClocksAt sets field value
func (o *BulkPairingList200ResponseInner) SetStartClocksAt(v int32) {
	o.StartClocksAt = v
}

// GetScheduledAt returns the ScheduledAt field value
func (o *BulkPairingList200ResponseInner) GetScheduledAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ScheduledAt
}

// GetScheduledAtOk returns a tuple with the ScheduledAt field value
// and a boolean to check if the value has been set.
func (o *BulkPairingList200ResponseInner) GetScheduledAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScheduledAt, true
}

// SetScheduledAt sets field value
func (o *BulkPairingList200ResponseInner) SetScheduledAt(v int32) {
	o.ScheduledAt = v
}

func (o BulkPairingList200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkPairingList200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["games"] = o.Games
	toSerialize["variant"] = o.Variant
	toSerialize["clock"] = o.Clock
	toSerialize["pairAt"] = o.PairAt
	toSerialize["pairedAt"] = o.PairedAt.Get()
	toSerialize["rated"] = o.Rated
	toSerialize["startClocksAt"] = o.StartClocksAt
	toSerialize["scheduledAt"] = o.ScheduledAt
	return toSerialize, nil
}

func (o *BulkPairingList200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"games",
		"variant",
		"clock",
		"pairAt",
		"pairedAt",
		"rated",
		"startClocksAt",
		"scheduledAt",
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

	varBulkPairingList200ResponseInner := _BulkPairingList200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkPairingList200ResponseInner)

	if err != nil {
		return err
	}

	*o = BulkPairingList200ResponseInner(varBulkPairingList200ResponseInner)

	return err
}

type NullableBulkPairingList200ResponseInner struct {
	value *BulkPairingList200ResponseInner
	isSet bool
}

func (v NullableBulkPairingList200ResponseInner) Get() *BulkPairingList200ResponseInner {
	return v.value
}

func (v *NullableBulkPairingList200ResponseInner) Set(val *BulkPairingList200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkPairingList200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkPairingList200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkPairingList200ResponseInner(val *BulkPairingList200ResponseInner) *NullableBulkPairingList200ResponseInner {
	return &NullableBulkPairingList200ResponseInner{value: val, isSet: true}
}

func (v NullableBulkPairingList200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkPairingList200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


