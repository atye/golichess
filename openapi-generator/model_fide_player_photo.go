/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.164
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FIDEPlayerPhoto type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FIDEPlayerPhoto{}

// FIDEPlayerPhoto struct for FIDEPlayerPhoto
type FIDEPlayerPhoto struct {
	// URL of a small (100x100) thumbnail of the photo
	Small string `json:"small"`
	// URL of a medium (500x500) version of the photo
	Medium string `json:"medium"`
	// If set, then you should make it appear next to the photo
	Credit *string `json:"credit,omitempty"`
}

type _FIDEPlayerPhoto FIDEPlayerPhoto

// NewFIDEPlayerPhoto instantiates a new FIDEPlayerPhoto object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFIDEPlayerPhoto(small string, medium string) *FIDEPlayerPhoto {
	this := FIDEPlayerPhoto{}
	this.Small = small
	this.Medium = medium
	return &this
}

// NewFIDEPlayerPhotoWithDefaults instantiates a new FIDEPlayerPhoto object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFIDEPlayerPhotoWithDefaults() *FIDEPlayerPhoto {
	this := FIDEPlayerPhoto{}
	return &this
}

// GetSmall returns the Small field value
func (o *FIDEPlayerPhoto) GetSmall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Small
}

// GetSmallOk returns a tuple with the Small field value
// and a boolean to check if the value has been set.
func (o *FIDEPlayerPhoto) GetSmallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Small, true
}

// SetSmall sets field value
func (o *FIDEPlayerPhoto) SetSmall(v string) {
	o.Small = v
}

// GetMedium returns the Medium field value
func (o *FIDEPlayerPhoto) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *FIDEPlayerPhoto) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Medium, true
}

// SetMedium sets field value
func (o *FIDEPlayerPhoto) SetMedium(v string) {
	o.Medium = v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *FIDEPlayerPhoto) GetCredit() string {
	if o == nil || IsNil(o.Credit) {
		var ret string
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FIDEPlayerPhoto) GetCreditOk() (*string, bool) {
	if o == nil || IsNil(o.Credit) {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *FIDEPlayerPhoto) HasCredit() bool {
	if o != nil && !IsNil(o.Credit) {
		return true
	}

	return false
}

// SetCredit gets a reference to the given string and assigns it to the Credit field.
func (o *FIDEPlayerPhoto) SetCredit(v string) {
	o.Credit = &v
}

func (o FIDEPlayerPhoto) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FIDEPlayerPhoto) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["small"] = o.Small
	toSerialize["medium"] = o.Medium
	if !IsNil(o.Credit) {
		toSerialize["credit"] = o.Credit
	}
	return toSerialize, nil
}

func (o *FIDEPlayerPhoto) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"small",
		"medium",
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

	varFIDEPlayerPhoto := _FIDEPlayerPhoto{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFIDEPlayerPhoto)

	if err != nil {
		return err
	}

	*o = FIDEPlayerPhoto(varFIDEPlayerPhoto)

	return err
}

type NullableFIDEPlayerPhoto struct {
	value *FIDEPlayerPhoto
	isSet bool
}

func (v NullableFIDEPlayerPhoto) Get() *FIDEPlayerPhoto {
	return v.value
}

func (v *NullableFIDEPlayerPhoto) Set(val *FIDEPlayerPhoto) {
	v.value = val
	v.isSet = true
}

func (v NullableFIDEPlayerPhoto) IsSet() bool {
	return v.isSet
}

func (v *NullableFIDEPlayerPhoto) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFIDEPlayerPhoto(val *FIDEPlayerPhoto) *NullableFIDEPlayerPhoto {
	return &NullableFIDEPlayerPhoto{value: val, isSet: true}
}

func (v NullableFIDEPlayerPhoto) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFIDEPlayerPhoto) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


