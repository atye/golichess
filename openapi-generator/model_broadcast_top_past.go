/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.171
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
)

// checks if the BroadcastTopPast type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTopPast{}

// BroadcastTopPast struct for BroadcastTopPast
type BroadcastTopPast struct {
	CurrentPage *int32 `json:"currentPage,omitempty"`
	MaxPerPage *int32 `json:"maxPerPage,omitempty"`
	CurrentPageResults []BroadcastWithLastRound `json:"currentPageResults,omitempty"`
	PreviousPage NullableInt32 `json:"previousPage,omitempty"`
	NextPage NullableInt32 `json:"nextPage,omitempty"`
}

// NewBroadcastTopPast instantiates a new BroadcastTopPast object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTopPast() *BroadcastTopPast {
	this := BroadcastTopPast{}
	return &this
}

// NewBroadcastTopPastWithDefaults instantiates a new BroadcastTopPast object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTopPastWithDefaults() *BroadcastTopPast {
	this := BroadcastTopPast{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value if set, zero value otherwise.
func (o *BroadcastTopPast) GetCurrentPage() int32 {
	if o == nil || IsNil(o.CurrentPage) {
		var ret int32
		return ret
	}
	return *o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTopPast) GetCurrentPageOk() (*int32, bool) {
	if o == nil || IsNil(o.CurrentPage) {
		return nil, false
	}
	return o.CurrentPage, true
}

// HasCurrentPage returns a boolean if a field has been set.
func (o *BroadcastTopPast) HasCurrentPage() bool {
	if o != nil && !IsNil(o.CurrentPage) {
		return true
	}

	return false
}

// SetCurrentPage gets a reference to the given int32 and assigns it to the CurrentPage field.
func (o *BroadcastTopPast) SetCurrentPage(v int32) {
	o.CurrentPage = &v
}

// GetMaxPerPage returns the MaxPerPage field value if set, zero value otherwise.
func (o *BroadcastTopPast) GetMaxPerPage() int32 {
	if o == nil || IsNil(o.MaxPerPage) {
		var ret int32
		return ret
	}
	return *o.MaxPerPage
}

// GetMaxPerPageOk returns a tuple with the MaxPerPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTopPast) GetMaxPerPageOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPerPage) {
		return nil, false
	}
	return o.MaxPerPage, true
}

// HasMaxPerPage returns a boolean if a field has been set.
func (o *BroadcastTopPast) HasMaxPerPage() bool {
	if o != nil && !IsNil(o.MaxPerPage) {
		return true
	}

	return false
}

// SetMaxPerPage gets a reference to the given int32 and assigns it to the MaxPerPage field.
func (o *BroadcastTopPast) SetMaxPerPage(v int32) {
	o.MaxPerPage = &v
}

// GetCurrentPageResults returns the CurrentPageResults field value if set, zero value otherwise.
func (o *BroadcastTopPast) GetCurrentPageResults() []BroadcastWithLastRound {
	if o == nil || IsNil(o.CurrentPageResults) {
		var ret []BroadcastWithLastRound
		return ret
	}
	return o.CurrentPageResults
}

// GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTopPast) GetCurrentPageResultsOk() ([]BroadcastWithLastRound, bool) {
	if o == nil || IsNil(o.CurrentPageResults) {
		return nil, false
	}
	return o.CurrentPageResults, true
}

// HasCurrentPageResults returns a boolean if a field has been set.
func (o *BroadcastTopPast) HasCurrentPageResults() bool {
	if o != nil && !IsNil(o.CurrentPageResults) {
		return true
	}

	return false
}

// SetCurrentPageResults gets a reference to the given []BroadcastWithLastRound and assigns it to the CurrentPageResults field.
func (o *BroadcastTopPast) SetCurrentPageResults(v []BroadcastWithLastRound) {
	o.CurrentPageResults = v
}

// GetPreviousPage returns the PreviousPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BroadcastTopPast) GetPreviousPage() int32 {
	if o == nil || IsNil(o.PreviousPage.Get()) {
		var ret int32
		return ret
	}
	return *o.PreviousPage.Get()
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BroadcastTopPast) GetPreviousPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPage.Get(), o.PreviousPage.IsSet()
}

// HasPreviousPage returns a boolean if a field has been set.
func (o *BroadcastTopPast) HasPreviousPage() bool {
	if o != nil && o.PreviousPage.IsSet() {
		return true
	}

	return false
}

// SetPreviousPage gets a reference to the given NullableInt32 and assigns it to the PreviousPage field.
func (o *BroadcastTopPast) SetPreviousPage(v int32) {
	o.PreviousPage.Set(&v)
}
// SetPreviousPageNil sets the value for PreviousPage to be an explicit nil
func (o *BroadcastTopPast) SetPreviousPageNil() {
	o.PreviousPage.Set(nil)
}

// UnsetPreviousPage ensures that no value is present for PreviousPage, not even an explicit nil
func (o *BroadcastTopPast) UnsetPreviousPage() {
	o.PreviousPage.Unset()
}

// GetNextPage returns the NextPage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BroadcastTopPast) GetNextPage() int32 {
	if o == nil || IsNil(o.NextPage.Get()) {
		var ret int32
		return ret
	}
	return *o.NextPage.Get()
}

// GetNextPageOk returns a tuple with the NextPage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BroadcastTopPast) GetNextPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPage.Get(), o.NextPage.IsSet()
}

// HasNextPage returns a boolean if a field has been set.
func (o *BroadcastTopPast) HasNextPage() bool {
	if o != nil && o.NextPage.IsSet() {
		return true
	}

	return false
}

// SetNextPage gets a reference to the given NullableInt32 and assigns it to the NextPage field.
func (o *BroadcastTopPast) SetNextPage(v int32) {
	o.NextPage.Set(&v)
}
// SetNextPageNil sets the value for NextPage to be an explicit nil
func (o *BroadcastTopPast) SetNextPageNil() {
	o.NextPage.Set(nil)
}

// UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil
func (o *BroadcastTopPast) UnsetNextPage() {
	o.NextPage.Unset()
}

func (o BroadcastTopPast) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTopPast) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CurrentPage) {
		toSerialize["currentPage"] = o.CurrentPage
	}
	if !IsNil(o.MaxPerPage) {
		toSerialize["maxPerPage"] = o.MaxPerPage
	}
	if !IsNil(o.CurrentPageResults) {
		toSerialize["currentPageResults"] = o.CurrentPageResults
	}
	if o.PreviousPage.IsSet() {
		toSerialize["previousPage"] = o.PreviousPage.Get()
	}
	if o.NextPage.IsSet() {
		toSerialize["nextPage"] = o.NextPage.Get()
	}
	return toSerialize, nil
}

type NullableBroadcastTopPast struct {
	value *BroadcastTopPast
	isSet bool
}

func (v NullableBroadcastTopPast) Get() *BroadcastTopPast {
	return v.value
}

func (v *NullableBroadcastTopPast) Set(val *BroadcastTopPast) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTopPast) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTopPast) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTopPast(val *BroadcastTopPast) *NullableBroadcastTopPast {
	return &NullableBroadcastTopPast{value: val, isSet: true}
}

func (v NullableBroadcastTopPast) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTopPast) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


