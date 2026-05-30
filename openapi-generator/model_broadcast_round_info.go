/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.145
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastRoundInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastRoundInfo{}

// BroadcastRoundInfo struct for BroadcastRoundInfo
type BroadcastRoundInfo struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	CreatedAt int64 `json:"createdAt"`
	// Whether the round is used for rating calculations
	Rated bool `json:"rated"`
	Ongoing *bool `json:"ongoing,omitempty"`
	StartsAt *int64 `json:"startsAt,omitempty"`
	// The start date/time is unknown and the round will start automatically when the previous round completes
	StartsAfterPrevious *bool `json:"startsAfterPrevious,omitempty"`
	FinishedAt *int64 `json:"finishedAt,omitempty"`
	Finished *bool `json:"finished,omitempty"`
	Url string `json:"url"`
	Delay *int64 `json:"delay,omitempty"`
	CustomScoring *BroadcastCustomScoring `json:"customScoring,omitempty"`
}

type _BroadcastRoundInfo BroadcastRoundInfo

// NewBroadcastRoundInfo instantiates a new BroadcastRoundInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastRoundInfo(id string, name string, slug string, createdAt int64, rated bool, url string) *BroadcastRoundInfo {
	this := BroadcastRoundInfo{}
	this.Id = id
	this.Name = name
	this.Slug = slug
	this.CreatedAt = createdAt
	this.Rated = rated
	this.Url = url
	return &this
}

// NewBroadcastRoundInfoWithDefaults instantiates a new BroadcastRoundInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastRoundInfoWithDefaults() *BroadcastRoundInfo {
	this := BroadcastRoundInfo{}
	return &this
}

// GetId returns the Id field value
func (o *BroadcastRoundInfo) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BroadcastRoundInfo) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BroadcastRoundInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BroadcastRoundInfo) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BroadcastRoundInfo) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BroadcastRoundInfo) SetSlug(v string) {
	o.Slug = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BroadcastRoundInfo) GetCreatedAt() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetCreatedAtOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BroadcastRoundInfo) SetCreatedAt(v int64) {
	o.CreatedAt = v
}

// GetRated returns the Rated field value
func (o *BroadcastRoundInfo) GetRated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Rated
}

// GetRatedOk returns a tuple with the Rated field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetRatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rated, true
}

// SetRated sets field value
func (o *BroadcastRoundInfo) SetRated(v bool) {
	o.Rated = v
}

// GetOngoing returns the Ongoing field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetOngoing() bool {
	if o == nil || IsNil(o.Ongoing) {
		var ret bool
		return ret
	}
	return *o.Ongoing
}

// GetOngoingOk returns a tuple with the Ongoing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetOngoingOk() (*bool, bool) {
	if o == nil || IsNil(o.Ongoing) {
		return nil, false
	}
	return o.Ongoing, true
}

// HasOngoing returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasOngoing() bool {
	if o != nil && !IsNil(o.Ongoing) {
		return true
	}

	return false
}

// SetOngoing gets a reference to the given bool and assigns it to the Ongoing field.
func (o *BroadcastRoundInfo) SetOngoing(v bool) {
	o.Ongoing = &v
}

// GetStartsAt returns the StartsAt field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetStartsAt() int64 {
	if o == nil || IsNil(o.StartsAt) {
		var ret int64
		return ret
	}
	return *o.StartsAt
}

// GetStartsAtOk returns a tuple with the StartsAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetStartsAtOk() (*int64, bool) {
	if o == nil || IsNil(o.StartsAt) {
		return nil, false
	}
	return o.StartsAt, true
}

// HasStartsAt returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasStartsAt() bool {
	if o != nil && !IsNil(o.StartsAt) {
		return true
	}

	return false
}

// SetStartsAt gets a reference to the given int64 and assigns it to the StartsAt field.
func (o *BroadcastRoundInfo) SetStartsAt(v int64) {
	o.StartsAt = &v
}

// GetStartsAfterPrevious returns the StartsAfterPrevious field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetStartsAfterPrevious() bool {
	if o == nil || IsNil(o.StartsAfterPrevious) {
		var ret bool
		return ret
	}
	return *o.StartsAfterPrevious
}

// GetStartsAfterPreviousOk returns a tuple with the StartsAfterPrevious field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetStartsAfterPreviousOk() (*bool, bool) {
	if o == nil || IsNil(o.StartsAfterPrevious) {
		return nil, false
	}
	return o.StartsAfterPrevious, true
}

// HasStartsAfterPrevious returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasStartsAfterPrevious() bool {
	if o != nil && !IsNil(o.StartsAfterPrevious) {
		return true
	}

	return false
}

// SetStartsAfterPrevious gets a reference to the given bool and assigns it to the StartsAfterPrevious field.
func (o *BroadcastRoundInfo) SetStartsAfterPrevious(v bool) {
	o.StartsAfterPrevious = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetFinishedAt() int64 {
	if o == nil || IsNil(o.FinishedAt) {
		var ret int64
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetFinishedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given int64 and assigns it to the FinishedAt field.
func (o *BroadcastRoundInfo) SetFinishedAt(v int64) {
	o.FinishedAt = &v
}

// GetFinished returns the Finished field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetFinished() bool {
	if o == nil || IsNil(o.Finished) {
		var ret bool
		return ret
	}
	return *o.Finished
}

// GetFinishedOk returns a tuple with the Finished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetFinishedOk() (*bool, bool) {
	if o == nil || IsNil(o.Finished) {
		return nil, false
	}
	return o.Finished, true
}

// HasFinished returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasFinished() bool {
	if o != nil && !IsNil(o.Finished) {
		return true
	}

	return false
}

// SetFinished gets a reference to the given bool and assigns it to the Finished field.
func (o *BroadcastRoundInfo) SetFinished(v bool) {
	o.Finished = &v
}

// GetUrl returns the Url field value
func (o *BroadcastRoundInfo) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BroadcastRoundInfo) SetUrl(v string) {
	o.Url = v
}

// GetDelay returns the Delay field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetDelay() int64 {
	if o == nil || IsNil(o.Delay) {
		var ret int64
		return ret
	}
	return *o.Delay
}

// GetDelayOk returns a tuple with the Delay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetDelayOk() (*int64, bool) {
	if o == nil || IsNil(o.Delay) {
		return nil, false
	}
	return o.Delay, true
}

// HasDelay returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasDelay() bool {
	if o != nil && !IsNil(o.Delay) {
		return true
	}

	return false
}

// SetDelay gets a reference to the given int64 and assigns it to the Delay field.
func (o *BroadcastRoundInfo) SetDelay(v int64) {
	o.Delay = &v
}

// GetCustomScoring returns the CustomScoring field value if set, zero value otherwise.
func (o *BroadcastRoundInfo) GetCustomScoring() BroadcastCustomScoring {
	if o == nil || IsNil(o.CustomScoring) {
		var ret BroadcastCustomScoring
		return ret
	}
	return *o.CustomScoring
}

// GetCustomScoringOk returns a tuple with the CustomScoring field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastRoundInfo) GetCustomScoringOk() (*BroadcastCustomScoring, bool) {
	if o == nil || IsNil(o.CustomScoring) {
		return nil, false
	}
	return o.CustomScoring, true
}

// HasCustomScoring returns a boolean if a field has been set.
func (o *BroadcastRoundInfo) HasCustomScoring() bool {
	if o != nil && !IsNil(o.CustomScoring) {
		return true
	}

	return false
}

// SetCustomScoring gets a reference to the given BroadcastCustomScoring and assigns it to the CustomScoring field.
func (o *BroadcastRoundInfo) SetCustomScoring(v BroadcastCustomScoring) {
	o.CustomScoring = &v
}

func (o BroadcastRoundInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastRoundInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["createdAt"] = o.CreatedAt
	toSerialize["rated"] = o.Rated
	if !IsNil(o.Ongoing) {
		toSerialize["ongoing"] = o.Ongoing
	}
	if !IsNil(o.StartsAt) {
		toSerialize["startsAt"] = o.StartsAt
	}
	if !IsNil(o.StartsAfterPrevious) {
		toSerialize["startsAfterPrevious"] = o.StartsAfterPrevious
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	if !IsNil(o.Finished) {
		toSerialize["finished"] = o.Finished
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Delay) {
		toSerialize["delay"] = o.Delay
	}
	if !IsNil(o.CustomScoring) {
		toSerialize["customScoring"] = o.CustomScoring
	}
	return toSerialize, nil
}

func (o *BroadcastRoundInfo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"slug",
		"createdAt",
		"rated",
		"url",
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

	varBroadcastRoundInfo := _BroadcastRoundInfo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastRoundInfo)

	if err != nil {
		return err
	}

	*o = BroadcastRoundInfo(varBroadcastRoundInfo)

	return err
}

type NullableBroadcastRoundInfo struct {
	value *BroadcastRoundInfo
	isSet bool
}

func (v NullableBroadcastRoundInfo) Get() *BroadcastRoundInfo {
	return v.value
}

func (v *NullableBroadcastRoundInfo) Set(val *BroadcastRoundInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastRoundInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastRoundInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastRoundInfo(val *BroadcastRoundInfo) *NullableBroadcastRoundInfo {
	return &NullableBroadcastRoundInfo{value: val, isSet: true}
}

func (v NullableBroadcastRoundInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastRoundInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


