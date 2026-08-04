/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.161
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
)

// checks if the BroadcastTourInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTourInfo{}

// BroadcastTourInfo Additional display information about the tournament
type BroadcastTourInfo struct {
	// Tournament format. Example: `\"8-player round-robin\" or \"5-round Swiss\"` 
	Format *string `json:"format,omitempty"`
	// Time control. Example: `\"Classical\" or \"Rapid\" or \"Rapid & Blitz\"` 
	Tc *string `json:"tc,omitempty"`
	FideTC *FideTimeControl `json:"fideTC,omitempty"`
	// Timezone of the tournament. Example: `America/New_York`. See [list of possible timezone identifiers](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones) for more. 
	TimeZone *string `json:"timeZone,omitempty"`
	// Tournament location
	Location *string `json:"location,omitempty"`
	// Mentioning up to 4 of the best players participating. 
	Players *string `json:"players,omitempty"`
	// Official website. External website URL
	Website *string `json:"website,omitempty"`
	// Official standings website. External website URL, e.g. chess-results.com, info64.org 
	Standings *string `json:"standings,omitempty"`
	// External URL to the official tournament regulations. 
	Regulations *string `json:"regulations,omitempty"`
}

// NewBroadcastTourInfo instantiates a new BroadcastTourInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTourInfo() *BroadcastTourInfo {
	this := BroadcastTourInfo{}
	return &this
}

// NewBroadcastTourInfoWithDefaults instantiates a new BroadcastTourInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTourInfoWithDefaults() *BroadcastTourInfo {
	this := BroadcastTourInfo{}
	return &this
}

// GetFormat returns the Format field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetFormat() string {
	if o == nil || IsNil(o.Format) {
		var ret string
		return ret
	}
	return *o.Format
}

// GetFormatOk returns a tuple with the Format field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetFormatOk() (*string, bool) {
	if o == nil || IsNil(o.Format) {
		return nil, false
	}
	return o.Format, true
}

// HasFormat returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasFormat() bool {
	if o != nil && !IsNil(o.Format) {
		return true
	}

	return false
}

// SetFormat gets a reference to the given string and assigns it to the Format field.
func (o *BroadcastTourInfo) SetFormat(v string) {
	o.Format = &v
}

// GetTc returns the Tc field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetTc() string {
	if o == nil || IsNil(o.Tc) {
		var ret string
		return ret
	}
	return *o.Tc
}

// GetTcOk returns a tuple with the Tc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetTcOk() (*string, bool) {
	if o == nil || IsNil(o.Tc) {
		return nil, false
	}
	return o.Tc, true
}

// HasTc returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasTc() bool {
	if o != nil && !IsNil(o.Tc) {
		return true
	}

	return false
}

// SetTc gets a reference to the given string and assigns it to the Tc field.
func (o *BroadcastTourInfo) SetTc(v string) {
	o.Tc = &v
}

// GetFideTC returns the FideTC field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetFideTC() FideTimeControl {
	if o == nil || IsNil(o.FideTC) {
		var ret FideTimeControl
		return ret
	}
	return *o.FideTC
}

// GetFideTCOk returns a tuple with the FideTC field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetFideTCOk() (*FideTimeControl, bool) {
	if o == nil || IsNil(o.FideTC) {
		return nil, false
	}
	return o.FideTC, true
}

// HasFideTC returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasFideTC() bool {
	if o != nil && !IsNil(o.FideTC) {
		return true
	}

	return false
}

// SetFideTC gets a reference to the given FideTimeControl and assigns it to the FideTC field.
func (o *BroadcastTourInfo) SetFideTC(v FideTimeControl) {
	o.FideTC = &v
}

// GetTimeZone returns the TimeZone field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetTimeZone() string {
	if o == nil || IsNil(o.TimeZone) {
		var ret string
		return ret
	}
	return *o.TimeZone
}

// GetTimeZoneOk returns a tuple with the TimeZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetTimeZoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimeZone) {
		return nil, false
	}
	return o.TimeZone, true
}

// HasTimeZone returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasTimeZone() bool {
	if o != nil && !IsNil(o.TimeZone) {
		return true
	}

	return false
}

// SetTimeZone gets a reference to the given string and assigns it to the TimeZone field.
func (o *BroadcastTourInfo) SetTimeZone(v string) {
	o.TimeZone = &v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetLocation() string {
	if o == nil || IsNil(o.Location) {
		var ret string
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetLocationOk() (*string, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given string and assigns it to the Location field.
func (o *BroadcastTourInfo) SetLocation(v string) {
	o.Location = &v
}

// GetPlayers returns the Players field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetPlayers() string {
	if o == nil || IsNil(o.Players) {
		var ret string
		return ret
	}
	return *o.Players
}

// GetPlayersOk returns a tuple with the Players field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetPlayersOk() (*string, bool) {
	if o == nil || IsNil(o.Players) {
		return nil, false
	}
	return o.Players, true
}

// HasPlayers returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasPlayers() bool {
	if o != nil && !IsNil(o.Players) {
		return true
	}

	return false
}

// SetPlayers gets a reference to the given string and assigns it to the Players field.
func (o *BroadcastTourInfo) SetPlayers(v string) {
	o.Players = &v
}

// GetWebsite returns the Website field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetWebsite() string {
	if o == nil || IsNil(o.Website) {
		var ret string
		return ret
	}
	return *o.Website
}

// GetWebsiteOk returns a tuple with the Website field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetWebsiteOk() (*string, bool) {
	if o == nil || IsNil(o.Website) {
		return nil, false
	}
	return o.Website, true
}

// HasWebsite returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasWebsite() bool {
	if o != nil && !IsNil(o.Website) {
		return true
	}

	return false
}

// SetWebsite gets a reference to the given string and assigns it to the Website field.
func (o *BroadcastTourInfo) SetWebsite(v string) {
	o.Website = &v
}

// GetStandings returns the Standings field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetStandings() string {
	if o == nil || IsNil(o.Standings) {
		var ret string
		return ret
	}
	return *o.Standings
}

// GetStandingsOk returns a tuple with the Standings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetStandingsOk() (*string, bool) {
	if o == nil || IsNil(o.Standings) {
		return nil, false
	}
	return o.Standings, true
}

// HasStandings returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasStandings() bool {
	if o != nil && !IsNil(o.Standings) {
		return true
	}

	return false
}

// SetStandings gets a reference to the given string and assigns it to the Standings field.
func (o *BroadcastTourInfo) SetStandings(v string) {
	o.Standings = &v
}

// GetRegulations returns the Regulations field value if set, zero value otherwise.
func (o *BroadcastTourInfo) GetRegulations() string {
	if o == nil || IsNil(o.Regulations) {
		var ret string
		return ret
	}
	return *o.Regulations
}

// GetRegulationsOk returns a tuple with the Regulations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTourInfo) GetRegulationsOk() (*string, bool) {
	if o == nil || IsNil(o.Regulations) {
		return nil, false
	}
	return o.Regulations, true
}

// HasRegulations returns a boolean if a field has been set.
func (o *BroadcastTourInfo) HasRegulations() bool {
	if o != nil && !IsNil(o.Regulations) {
		return true
	}

	return false
}

// SetRegulations gets a reference to the given string and assigns it to the Regulations field.
func (o *BroadcastTourInfo) SetRegulations(v string) {
	o.Regulations = &v
}

func (o BroadcastTourInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTourInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Format) {
		toSerialize["format"] = o.Format
	}
	if !IsNil(o.Tc) {
		toSerialize["tc"] = o.Tc
	}
	if !IsNil(o.FideTC) {
		toSerialize["fideTC"] = o.FideTC
	}
	if !IsNil(o.TimeZone) {
		toSerialize["timeZone"] = o.TimeZone
	}
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Players) {
		toSerialize["players"] = o.Players
	}
	if !IsNil(o.Website) {
		toSerialize["website"] = o.Website
	}
	if !IsNil(o.Standings) {
		toSerialize["standings"] = o.Standings
	}
	if !IsNil(o.Regulations) {
		toSerialize["regulations"] = o.Regulations
	}
	return toSerialize, nil
}

type NullableBroadcastTourInfo struct {
	value *BroadcastTourInfo
	isSet bool
}

func (v NullableBroadcastTourInfo) Get() *BroadcastTourInfo {
	return v.value
}

func (v *NullableBroadcastTourInfo) Set(val *BroadcastTourInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTourInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTourInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTourInfo(val *BroadcastTourInfo) *NullableBroadcastTourInfo {
	return &NullableBroadcastTourInfo{value: val, isSet: true}
}

func (v NullableBroadcastTourInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTourInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


