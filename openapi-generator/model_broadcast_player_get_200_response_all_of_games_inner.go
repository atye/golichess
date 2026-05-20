/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi-generator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastPlayerGet200ResponseAllOfGamesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastPlayerGet200ResponseAllOfGamesInner{}

// BroadcastPlayerGet200ResponseAllOfGamesInner struct for BroadcastPlayerGet200ResponseAllOfGamesInner
type BroadcastPlayerGet200ResponseAllOfGamesInner struct {
	// ID of the round
	Round string `json:"round"`
	// The game ID. Analogous to chapterId.
	Id string `json:"id"`
	Opponent BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent `json:"opponent"`
	Color NullableString `json:"color"`
	Points *string `json:"points,omitempty"`
	CustomPoints *float32 `json:"customPoints,omitempty"`
	// The change in rating for the player as a result of this game
	RatingDiff *int32 `json:"ratingDiff,omitempty"`
	// FIDE rating category 
	FideTC string `json:"fideTC"`
	Ongoing *bool `json:"ongoing,omitempty"`
}

type _BroadcastPlayerGet200ResponseAllOfGamesInner BroadcastPlayerGet200ResponseAllOfGamesInner

// NewBroadcastPlayerGet200ResponseAllOfGamesInner instantiates a new BroadcastPlayerGet200ResponseAllOfGamesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastPlayerGet200ResponseAllOfGamesInner(round string, id string, opponent BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent, color NullableString, fideTC string) *BroadcastPlayerGet200ResponseAllOfGamesInner {
	this := BroadcastPlayerGet200ResponseAllOfGamesInner{}
	this.Round = round
	this.Id = id
	this.Opponent = opponent
	this.Color = color
	this.FideTC = fideTC
	return &this
}

// NewBroadcastPlayerGet200ResponseAllOfGamesInnerWithDefaults instantiates a new BroadcastPlayerGet200ResponseAllOfGamesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastPlayerGet200ResponseAllOfGamesInnerWithDefaults() *BroadcastPlayerGet200ResponseAllOfGamesInner {
	this := BroadcastPlayerGet200ResponseAllOfGamesInner{}
	return &this
}

// GetRound returns the Round field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRound() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Round
}

// GetRoundOk returns a tuple with the Round field value
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRoundOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Round, true
}

// SetRound sets field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetRound(v string) {
	o.Round = v
}

// GetId returns the Id field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetId(v string) {
	o.Id = v
}

// GetOpponent returns the Opponent field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOpponent() BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent {
	if o == nil {
		var ret BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent
		return ret
	}

	return o.Opponent
}

// GetOpponentOk returns a tuple with the Opponent field value
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOpponentOk() (*BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Opponent, true
}

// SetOpponent sets field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetOpponent(v BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent) {
	o.Opponent = v
}

// GetColor returns the Color field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetColor() string {
	if o == nil || o.Color.Get() == nil {
		var ret string
		return ret
	}

	return *o.Color.Get()
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetColorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Color.Get(), o.Color.IsSet()
}

// SetColor sets field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetColor(v string) {
	o.Color.Set(&v)
}

// GetPoints returns the Points field value if set, zero value otherwise.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetPoints() string {
	if o == nil || IsNil(o.Points) {
		var ret string
		return ret
	}
	return *o.Points
}

// GetPointsOk returns a tuple with the Points field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetPointsOk() (*string, bool) {
	if o == nil || IsNil(o.Points) {
		return nil, false
	}
	return o.Points, true
}

// HasPoints returns a boolean if a field has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasPoints() bool {
	if o != nil && !IsNil(o.Points) {
		return true
	}

	return false
}

// SetPoints gets a reference to the given string and assigns it to the Points field.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetPoints(v string) {
	o.Points = &v
}

// GetCustomPoints returns the CustomPoints field value if set, zero value otherwise.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetCustomPoints() float32 {
	if o == nil || IsNil(o.CustomPoints) {
		var ret float32
		return ret
	}
	return *o.CustomPoints
}

// GetCustomPointsOk returns a tuple with the CustomPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetCustomPointsOk() (*float32, bool) {
	if o == nil || IsNil(o.CustomPoints) {
		return nil, false
	}
	return o.CustomPoints, true
}

// HasCustomPoints returns a boolean if a field has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasCustomPoints() bool {
	if o != nil && !IsNil(o.CustomPoints) {
		return true
	}

	return false
}

// SetCustomPoints gets a reference to the given float32 and assigns it to the CustomPoints field.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetCustomPoints(v float32) {
	o.CustomPoints = &v
}

// GetRatingDiff returns the RatingDiff field value if set, zero value otherwise.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRatingDiff() int32 {
	if o == nil || IsNil(o.RatingDiff) {
		var ret int32
		return ret
	}
	return *o.RatingDiff
}

// GetRatingDiffOk returns a tuple with the RatingDiff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetRatingDiffOk() (*int32, bool) {
	if o == nil || IsNil(o.RatingDiff) {
		return nil, false
	}
	return o.RatingDiff, true
}

// HasRatingDiff returns a boolean if a field has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasRatingDiff() bool {
	if o != nil && !IsNil(o.RatingDiff) {
		return true
	}

	return false
}

// SetRatingDiff gets a reference to the given int32 and assigns it to the RatingDiff field.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetRatingDiff(v int32) {
	o.RatingDiff = &v
}

// GetFideTC returns the FideTC field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetFideTC() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FideTC
}

// GetFideTCOk returns a tuple with the FideTC field value
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetFideTCOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FideTC, true
}

// SetFideTC sets field value
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetFideTC(v string) {
	o.FideTC = v
}

// GetOngoing returns the Ongoing field value if set, zero value otherwise.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOngoing() bool {
	if o == nil || IsNil(o.Ongoing) {
		var ret bool
		return ret
	}
	return *o.Ongoing
}

// GetOngoingOk returns a tuple with the Ongoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) GetOngoingOk() (*bool, bool) {
	if o == nil || IsNil(o.Ongoing) {
		return nil, false
	}
	return o.Ongoing, true
}

// HasOngoing returns a boolean if a field has been set.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) HasOngoing() bool {
	if o != nil && !IsNil(o.Ongoing) {
		return true
	}

	return false
}

// SetOngoing gets a reference to the given bool and assigns it to the Ongoing field.
func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) SetOngoing(v bool) {
	o.Ongoing = &v
}

func (o BroadcastPlayerGet200ResponseAllOfGamesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastPlayerGet200ResponseAllOfGamesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["round"] = o.Round
	toSerialize["id"] = o.Id
	toSerialize["opponent"] = o.Opponent
	toSerialize["color"] = o.Color.Get()
	if !IsNil(o.Points) {
		toSerialize["points"] = o.Points
	}
	if !IsNil(o.CustomPoints) {
		toSerialize["customPoints"] = o.CustomPoints
	}
	if !IsNil(o.RatingDiff) {
		toSerialize["ratingDiff"] = o.RatingDiff
	}
	toSerialize["fideTC"] = o.FideTC
	if !IsNil(o.Ongoing) {
		toSerialize["ongoing"] = o.Ongoing
	}
	return toSerialize, nil
}

func (o *BroadcastPlayerGet200ResponseAllOfGamesInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"round",
		"id",
		"opponent",
		"color",
		"fideTC",
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

	varBroadcastPlayerGet200ResponseAllOfGamesInner := _BroadcastPlayerGet200ResponseAllOfGamesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastPlayerGet200ResponseAllOfGamesInner)

	if err != nil {
		return err
	}

	*o = BroadcastPlayerGet200ResponseAllOfGamesInner(varBroadcastPlayerGet200ResponseAllOfGamesInner)

	return err
}

type NullableBroadcastPlayerGet200ResponseAllOfGamesInner struct {
	value *BroadcastPlayerGet200ResponseAllOfGamesInner
	isSet bool
}

func (v NullableBroadcastPlayerGet200ResponseAllOfGamesInner) Get() *BroadcastPlayerGet200ResponseAllOfGamesInner {
	return v.value
}

func (v *NullableBroadcastPlayerGet200ResponseAllOfGamesInner) Set(val *BroadcastPlayerGet200ResponseAllOfGamesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastPlayerGet200ResponseAllOfGamesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastPlayerGet200ResponseAllOfGamesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastPlayerGet200ResponseAllOfGamesInner(val *BroadcastPlayerGet200ResponseAllOfGamesInner) *NullableBroadcastPlayerGet200ResponseAllOfGamesInner {
	return &NullableBroadcastPlayerGet200ResponseAllOfGamesInner{value: val, isSet: true}
}

func (v NullableBroadcastPlayerGet200ResponseAllOfGamesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastPlayerGet200ResponseAllOfGamesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


