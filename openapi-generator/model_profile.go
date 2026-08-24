/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.166
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
)

// checks if the Profile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Profile{}

// Profile struct for Profile
type Profile struct {
	Flag *string `json:"flag,omitempty"`
	Location *string `json:"location,omitempty"`
	Bio *string `json:"bio,omitempty"`
	RealName *string `json:"realName,omitempty"`
	// only appears if a user has set them
	FideRating *int32 `json:"fideRating,omitempty"`
	// only appears if a user has set them
	UscfRating *int32 `json:"uscfRating,omitempty"`
	// only appears if a user has set them
	EcfRating *int32 `json:"ecfRating,omitempty"`
	// only appears if a user has set them
	CfcRating *int32 `json:"cfcRating,omitempty"`
	// only appears if a user has set them
	RcfRating *int32 `json:"rcfRating,omitempty"`
	// only appears if a user has set them
	DsbRating *int32 `json:"dsbRating,omitempty"`
	Links *string `json:"links,omitempty"`
}

// NewProfile instantiates a new Profile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfile() *Profile {
	this := Profile{}
	return &this
}

// NewProfileWithDefaults instantiates a new Profile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileWithDefaults() *Profile {
	this := Profile{}
	return &this
}

// GetFlag returns the Flag field value if set, zero value otherwise.
func (o *Profile) GetFlag() string {
	if o == nil || IsNil(o.Flag) {
		var ret string
		return ret
	}
	return *o.Flag
}

// GetFlagOk returns a tuple with the Flag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetFlagOk() (*string, bool) {
	if o == nil || IsNil(o.Flag) {
		return nil, false
	}
	return o.Flag, true
}

// HasFlag returns a boolean if a field has been set.
func (o *Profile) HasFlag() bool {
	if o != nil && !IsNil(o.Flag) {
		return true
	}

	return false
}

// SetFlag gets a reference to the given string and assigns it to the Flag field.
func (o *Profile) SetFlag(v string) {
	o.Flag = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *Profile) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *Profile) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *Profile) SetLocation(v string) {
	o.Location = &v
}

// GetBio returns the Bio field value if set, zero value otherwise.
func (o *Profile) GetBio() string {
	if o == nil || IsNil(o.Bio) {
		var ret string
		return ret
	}
	return *o.Bio
}

// GetBioOk returns a tuple with the Bio field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetBioOk() (*string, bool) {
	if o == nil || IsNil(o.Bio) {
		return nil, false
	}
	return o.Bio, true
}

// HasBio returns a boolean if a field has been set.
func (o *Profile) HasBio() bool {
	if o != nil && !IsNil(o.Bio) {
		return true
	}

	return false
}

// SetBio gets a reference to the given string and assigns it to the Bio field.
func (o *Profile) SetBio(v string) {
	o.Bio = &v
}

// GetRealName returns the RealName field value if set, zero value otherwise.
func (o *Profile) GetRealName() string {
	if o == nil || IsNil(o.RealName) {
		var ret string
		return ret
	}
	return *o.RealName
}

// GetRealNameOk returns a tuple with the RealName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetRealNameOk() (*string, bool) {
	if o == nil || IsNil(o.RealName) {
		return nil, false
	}
	return o.RealName, true
}

// HasRealName returns a boolean if a field has been set.
func (o *Profile) HasRealName() bool {
	if o != nil && !IsNil(o.RealName) {
		return true
	}

	return false
}

// SetRealName gets a reference to the given string and assigns it to the RealName field.
func (o *Profile) SetRealName(v string) {
	o.RealName = &v
}

// GetFideRating returns the FideRating field value if set, zero value otherwise.
func (o *Profile) GetFideRating() int32 {
	if o == nil || IsNil(o.FideRating) {
		var ret int32
		return ret
	}
	return *o.FideRating
}

// GetFideRatingOk returns a tuple with the FideRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetFideRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.FideRating) {
		return nil, false
	}
	return o.FideRating, true
}

// HasFideRating returns a boolean if a field has been set.
func (o *Profile) HasFideRating() bool {
	if o != nil && !IsNil(o.FideRating) {
		return true
	}

	return false
}

// SetFideRating gets a reference to the given int32 and assigns it to the FideRating field.
func (o *Profile) SetFideRating(v int32) {
	o.FideRating = &v
}

// GetUscfRating returns the UscfRating field value if set, zero value otherwise.
func (o *Profile) GetUscfRating() int32 {
	if o == nil || IsNil(o.UscfRating) {
		var ret int32
		return ret
	}
	return *o.UscfRating
}

// GetUscfRatingOk returns a tuple with the UscfRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetUscfRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.UscfRating) {
		return nil, false
	}
	return o.UscfRating, true
}

// HasUscfRating returns a boolean if a field has been set.
func (o *Profile) HasUscfRating() bool {
	if o != nil && !IsNil(o.UscfRating) {
		return true
	}

	return false
}

// SetUscfRating gets a reference to the given int32 and assigns it to the UscfRating field.
func (o *Profile) SetUscfRating(v int32) {
	o.UscfRating = &v
}

// GetEcfRating returns the EcfRating field value if set, zero value otherwise.
func (o *Profile) GetEcfRating() int32 {
	if o == nil || IsNil(o.EcfRating) {
		var ret int32
		return ret
	}
	return *o.EcfRating
}

// GetEcfRatingOk returns a tuple with the EcfRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetEcfRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.EcfRating) {
		return nil, false
	}
	return o.EcfRating, true
}

// HasEcfRating returns a boolean if a field has been set.
func (o *Profile) HasEcfRating() bool {
	if o != nil && !IsNil(o.EcfRating) {
		return true
	}

	return false
}

// SetEcfRating gets a reference to the given int32 and assigns it to the EcfRating field.
func (o *Profile) SetEcfRating(v int32) {
	o.EcfRating = &v
}

// GetCfcRating returns the CfcRating field value if set, zero value otherwise.
func (o *Profile) GetCfcRating() int32 {
	if o == nil || IsNil(o.CfcRating) {
		var ret int32
		return ret
	}
	return *o.CfcRating
}

// GetCfcRatingOk returns a tuple with the CfcRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetCfcRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.CfcRating) {
		return nil, false
	}
	return o.CfcRating, true
}

// HasCfcRating returns a boolean if a field has been set.
func (o *Profile) HasCfcRating() bool {
	if o != nil && !IsNil(o.CfcRating) {
		return true
	}

	return false
}

// SetCfcRating gets a reference to the given int32 and assigns it to the CfcRating field.
func (o *Profile) SetCfcRating(v int32) {
	o.CfcRating = &v
}

// GetRcfRating returns the RcfRating field value if set, zero value otherwise.
func (o *Profile) GetRcfRating() int32 {
	if o == nil || IsNil(o.RcfRating) {
		var ret int32
		return ret
	}
	return *o.RcfRating
}

// GetRcfRatingOk returns a tuple with the RcfRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetRcfRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.RcfRating) {
		return nil, false
	}
	return o.RcfRating, true
}

// HasRcfRating returns a boolean if a field has been set.
func (o *Profile) HasRcfRating() bool {
	if o != nil && !IsNil(o.RcfRating) {
		return true
	}

	return false
}

// SetRcfRating gets a reference to the given int32 and assigns it to the RcfRating field.
func (o *Profile) SetRcfRating(v int32) {
	o.RcfRating = &v
}

// GetDsbRating returns the DsbRating field value if set, zero value otherwise.
func (o *Profile) GetDsbRating() int32 {
	if o == nil || IsNil(o.DsbRating) {
		var ret int32
		return ret
	}
	return *o.DsbRating
}

// GetDsbRatingOk returns a tuple with the DsbRating field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetDsbRatingOk() (*int32, bool) {
	if o == nil || IsNil(o.DsbRating) {
		return nil, false
	}
	return o.DsbRating, true
}

// HasDsbRating returns a boolean if a field has been set.
func (o *Profile) HasDsbRating() bool {
	if o != nil && !IsNil(o.DsbRating) {
		return true
	}

	return false
}

// SetDsbRating gets a reference to the given int32 and assigns it to the DsbRating field.
func (o *Profile) SetDsbRating(v int32) {
	o.DsbRating = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *Profile) GetLinks() string {
	if o == nil || IsNil(o.Links) {
		var ret string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Profile) GetLinksOk() (*string, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *Profile) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given string and assigns it to the Links field.
func (o *Profile) SetLinks(v string) {
	o.Links = &v
}

func (o Profile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Profile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flag) {
		toSerialize["flag"] = o.Flag
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Bio) {
		toSerialize["bio"] = o.Bio
	}
	if !IsNil(o.RealName) {
		toSerialize["realName"] = o.RealName
	}
	if !IsNil(o.FideRating) {
		toSerialize["fideRating"] = o.FideRating
	}
	if !IsNil(o.UscfRating) {
		toSerialize["uscfRating"] = o.UscfRating
	}
	if !IsNil(o.EcfRating) {
		toSerialize["ecfRating"] = o.EcfRating
	}
	if !IsNil(o.CfcRating) {
		toSerialize["cfcRating"] = o.CfcRating
	}
	if !IsNil(o.RcfRating) {
		toSerialize["rcfRating"] = o.RcfRating
	}
	if !IsNil(o.DsbRating) {
		toSerialize["dsbRating"] = o.DsbRating
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableProfile struct {
	value *Profile
	isSet bool
}

func (v NullableProfile) Get() *Profile {
	return v.value
}

func (v *NullableProfile) Set(val *Profile) {
	v.value = val
	v.isSet = true
}

func (v NullableProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfile(val *Profile) *NullableProfile {
	return &NullableProfile{value: val, isSet: true}
}

func (v NullableProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


