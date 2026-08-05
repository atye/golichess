/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.162
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastsSearch200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastsSearch200Response{}

// BroadcastsSearch200Response struct for BroadcastsSearch200Response
type BroadcastsSearch200Response struct {
	CurrentPage int32 `json:"currentPage"`
	MaxPerPage int32 `json:"maxPerPage"`
	CurrentPageResults []BroadcastWithLastRound `json:"currentPageResults"`
	PreviousPage NullableInt32 `json:"previousPage"`
	NextPage NullableInt32 `json:"nextPage"`
}

type _BroadcastsSearch200Response BroadcastsSearch200Response

// NewBroadcastsSearch200Response instantiates a new BroadcastsSearch200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastsSearch200Response(currentPage int32, maxPerPage int32, currentPageResults []BroadcastWithLastRound, previousPage NullableInt32, nextPage NullableInt32) *BroadcastsSearch200Response {
	this := BroadcastsSearch200Response{}
	this.CurrentPage = currentPage
	this.MaxPerPage = maxPerPage
	this.CurrentPageResults = currentPageResults
	this.PreviousPage = previousPage
	this.NextPage = nextPage
	return &this
}

// NewBroadcastsSearch200ResponseWithDefaults instantiates a new BroadcastsSearch200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastsSearch200ResponseWithDefaults() *BroadcastsSearch200Response {
	this := BroadcastsSearch200Response{}
	return &this
}

// GetCurrentPage returns the CurrentPage field value
func (o *BroadcastsSearch200Response) GetCurrentPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentPage
}

// GetCurrentPageOk returns a tuple with the CurrentPage field value
// and a boolean to check if the value has been set.
func (o *BroadcastsSearch200Response) GetCurrentPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentPage, true
}

// SetCurrentPage sets field value
func (o *BroadcastsSearch200Response) SetCurrentPage(v int32) {
	o.CurrentPage = v
}

// GetMaxPerPage returns the MaxPerPage field value
func (o *BroadcastsSearch200Response) GetMaxPerPage() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPerPage
}

// GetMaxPerPageOk returns a tuple with the MaxPerPage field value
// and a boolean to check if the value has been set.
func (o *BroadcastsSearch200Response) GetMaxPerPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPerPage, true
}

// SetMaxPerPage sets field value
func (o *BroadcastsSearch200Response) SetMaxPerPage(v int32) {
	o.MaxPerPage = v
}

// GetCurrentPageResults returns the CurrentPageResults field value
func (o *BroadcastsSearch200Response) GetCurrentPageResults() []BroadcastWithLastRound {
	if o == nil {
		var ret []BroadcastWithLastRound
		return ret
	}

	return o.CurrentPageResults
}

// GetCurrentPageResultsOk returns a tuple with the CurrentPageResults field value
// and a boolean to check if the value has been set.
func (o *BroadcastsSearch200Response) GetCurrentPageResultsOk() ([]BroadcastWithLastRound, bool) {
	if o == nil {
		return nil, false
	}
	return o.CurrentPageResults, true
}

// SetCurrentPageResults sets field value
func (o *BroadcastsSearch200Response) SetCurrentPageResults(v []BroadcastWithLastRound) {
	o.CurrentPageResults = v
}

// GetPreviousPage returns the PreviousPage field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BroadcastsSearch200Response) GetPreviousPage() int32 {
	if o == nil || o.PreviousPage.Get() == nil {
		var ret int32
		return ret
	}

	return *o.PreviousPage.Get()
}

// GetPreviousPageOk returns a tuple with the PreviousPage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BroadcastsSearch200Response) GetPreviousPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreviousPage.Get(), o.PreviousPage.IsSet()
}

// SetPreviousPage sets field value
func (o *BroadcastsSearch200Response) SetPreviousPage(v int32) {
	o.PreviousPage.Set(&v)
}

// GetNextPage returns the NextPage field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *BroadcastsSearch200Response) GetNextPage() int32 {
	if o == nil || o.NextPage.Get() == nil {
		var ret int32
		return ret
	}

	return *o.NextPage.Get()
}

// GetNextPageOk returns a tuple with the NextPage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BroadcastsSearch200Response) GetNextPageOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NextPage.Get(), o.NextPage.IsSet()
}

// SetNextPage sets field value
func (o *BroadcastsSearch200Response) SetNextPage(v int32) {
	o.NextPage.Set(&v)
}

func (o BroadcastsSearch200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastsSearch200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["currentPage"] = o.CurrentPage
	toSerialize["maxPerPage"] = o.MaxPerPage
	toSerialize["currentPageResults"] = o.CurrentPageResults
	toSerialize["previousPage"] = o.PreviousPage.Get()
	toSerialize["nextPage"] = o.NextPage.Get()
	return toSerialize, nil
}

func (o *BroadcastsSearch200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"currentPage",
		"maxPerPage",
		"currentPageResults",
		"previousPage",
		"nextPage",
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

	varBroadcastsSearch200Response := _BroadcastsSearch200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastsSearch200Response)

	if err != nil {
		return err
	}

	*o = BroadcastsSearch200Response(varBroadcastsSearch200Response)

	return err
}

type NullableBroadcastsSearch200Response struct {
	value *BroadcastsSearch200Response
	isSet bool
}

func (v NullableBroadcastsSearch200Response) Get() *BroadcastsSearch200Response {
	return v.value
}

func (v *NullableBroadcastsSearch200Response) Set(val *BroadcastsSearch200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastsSearch200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastsSearch200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastsSearch200Response(val *BroadcastsSearch200Response) *NullableBroadcastsSearch200Response {
	return &NullableBroadcastsSearch200Response{value: val, isSet: true}
}

func (v NullableBroadcastsSearch200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastsSearch200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


