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

// checks if the BroadcastsOfficial200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastsOfficial200Response{}

// BroadcastsOfficial200Response struct for BroadcastsOfficial200Response
type BroadcastsOfficial200Response struct {
	Tour BroadcastsOfficial200ResponseTour `json:"tour"`
	Group *BroadcastsOfficial200ResponseGroup `json:"group,omitempty"`
	Rounds []BroadcastsOfficial200ResponseRoundsInner `json:"rounds"`
	DefaultRoundId *string `json:"defaultRoundId,omitempty"`
	// Photos of players, when available. The object keys are FIDE IDs
	Photos map[string]BroadcastsOfficial200ResponsePhotosValue `json:"photos,omitempty"`
}

type _BroadcastsOfficial200Response BroadcastsOfficial200Response

// NewBroadcastsOfficial200Response instantiates a new BroadcastsOfficial200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastsOfficial200Response(tour BroadcastsOfficial200ResponseTour, rounds []BroadcastsOfficial200ResponseRoundsInner) *BroadcastsOfficial200Response {
	this := BroadcastsOfficial200Response{}
	this.Tour = tour
	this.Rounds = rounds
	return &this
}

// NewBroadcastsOfficial200ResponseWithDefaults instantiates a new BroadcastsOfficial200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastsOfficial200ResponseWithDefaults() *BroadcastsOfficial200Response {
	this := BroadcastsOfficial200Response{}
	return &this
}

// GetTour returns the Tour field value
func (o *BroadcastsOfficial200Response) GetTour() BroadcastsOfficial200ResponseTour {
	if o == nil {
		var ret BroadcastsOfficial200ResponseTour
		return ret
	}

	return o.Tour
}

// GetTourOk returns a tuple with the Tour field value
// and a boolean to check if the value has been set.
func (o *BroadcastsOfficial200Response) GetTourOk() (*BroadcastsOfficial200ResponseTour, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tour, true
}

// SetTour sets field value
func (o *BroadcastsOfficial200Response) SetTour(v BroadcastsOfficial200ResponseTour) {
	o.Tour = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *BroadcastsOfficial200Response) GetGroup() BroadcastsOfficial200ResponseGroup {
	if o == nil || IsNil(o.Group) {
		var ret BroadcastsOfficial200ResponseGroup
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsOfficial200Response) GetGroupOk() (*BroadcastsOfficial200ResponseGroup, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *BroadcastsOfficial200Response) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given BroadcastsOfficial200ResponseGroup and assigns it to the Group field.
func (o *BroadcastsOfficial200Response) SetGroup(v BroadcastsOfficial200ResponseGroup) {
	o.Group = &v
}

// GetRounds returns the Rounds field value
func (o *BroadcastsOfficial200Response) GetRounds() []BroadcastsOfficial200ResponseRoundsInner {
	if o == nil {
		var ret []BroadcastsOfficial200ResponseRoundsInner
		return ret
	}

	return o.Rounds
}

// GetRoundsOk returns a tuple with the Rounds field value
// and a boolean to check if the value has been set.
func (o *BroadcastsOfficial200Response) GetRoundsOk() ([]BroadcastsOfficial200ResponseRoundsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Rounds, true
}

// SetRounds sets field value
func (o *BroadcastsOfficial200Response) SetRounds(v []BroadcastsOfficial200ResponseRoundsInner) {
	o.Rounds = v
}

// GetDefaultRoundId returns the DefaultRoundId field value if set, zero value otherwise.
func (o *BroadcastsOfficial200Response) GetDefaultRoundId() string {
	if o == nil || IsNil(o.DefaultRoundId) {
		var ret string
		return ret
	}
	return *o.DefaultRoundId
}

// GetDefaultRoundIdOk returns a tuple with the DefaultRoundId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsOfficial200Response) GetDefaultRoundIdOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRoundId) {
		return nil, false
	}
	return o.DefaultRoundId, true
}

// HasDefaultRoundId returns a boolean if a field has been set.
func (o *BroadcastsOfficial200Response) HasDefaultRoundId() bool {
	if o != nil && !IsNil(o.DefaultRoundId) {
		return true
	}

	return false
}

// SetDefaultRoundId gets a reference to the given string and assigns it to the DefaultRoundId field.
func (o *BroadcastsOfficial200Response) SetDefaultRoundId(v string) {
	o.DefaultRoundId = &v
}

// GetPhotos returns the Photos field value if set, zero value otherwise.
func (o *BroadcastsOfficial200Response) GetPhotos() map[string]BroadcastsOfficial200ResponsePhotosValue {
	if o == nil || IsNil(o.Photos) {
		var ret map[string]BroadcastsOfficial200ResponsePhotosValue
		return ret
	}
	return o.Photos
}

// GetPhotosOk returns a tuple with the Photos field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastsOfficial200Response) GetPhotosOk() (map[string]BroadcastsOfficial200ResponsePhotosValue, bool) {
	if o == nil || IsNil(o.Photos) {
		return map[string]BroadcastsOfficial200ResponsePhotosValue{}, false
	}
	return o.Photos, true
}

// HasPhotos returns a boolean if a field has been set.
func (o *BroadcastsOfficial200Response) HasPhotos() bool {
	if o != nil && !IsNil(o.Photos) {
		return true
	}

	return false
}

// SetPhotos gets a reference to the given map[string]BroadcastsOfficial200ResponsePhotosValue and assigns it to the Photos field.
func (o *BroadcastsOfficial200Response) SetPhotos(v map[string]BroadcastsOfficial200ResponsePhotosValue) {
	o.Photos = v
}

func (o BroadcastsOfficial200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastsOfficial200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tour"] = o.Tour
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	toSerialize["rounds"] = o.Rounds
	if !IsNil(o.DefaultRoundId) {
		toSerialize["defaultRoundId"] = o.DefaultRoundId
	}
	if !IsNil(o.Photos) {
		toSerialize["photos"] = o.Photos
	}
	return toSerialize, nil
}

func (o *BroadcastsOfficial200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tour",
		"rounds",
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

	varBroadcastsOfficial200Response := _BroadcastsOfficial200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastsOfficial200Response)

	if err != nil {
		return err
	}

	*o = BroadcastsOfficial200Response(varBroadcastsOfficial200Response)

	return err
}

type NullableBroadcastsOfficial200Response struct {
	value *BroadcastsOfficial200Response
	isSet bool
}

func (v NullableBroadcastsOfficial200Response) Get() *BroadcastsOfficial200Response {
	return v.value
}

func (v *NullableBroadcastsOfficial200Response) Set(val *BroadcastsOfficial200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastsOfficial200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastsOfficial200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastsOfficial200Response(val *BroadcastsOfficial200Response) *NullableBroadcastsOfficial200Response {
	return &NullableBroadcastsOfficial200Response{value: val, isSet: true}
}

func (v NullableBroadcastsOfficial200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastsOfficial200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


