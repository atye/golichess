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

// checks if the ExternalEngineRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEngineRegistration{}

// ExternalEngineRegistration struct for ExternalEngineRegistration
type ExternalEngineRegistration struct {
	// Display name of the engine.
	Name string `json:"name"`
	// Maximum number of available threads.
	MaxThreads int32 `json:"maxThreads"`
	// Maximum available hash table size, in MiB.
	MaxHash int32 `json:"maxHash"`
	// Optional list of supported chess variants.
	Variants []UciVariant `json:"variants,omitempty"`
	// A random token that can be used to [wait for analysis requests](#tag/external-engine/POST/api/external-engine/work) and provide analysis.  The engine provider should securely generate a random string.  The token will not be readable again, even by the user.  The analysis provider can register multiple engines with the same token, even for different users, and wait for analysis requests from any of them. In this case, the request must not be made via CORS, so that the token is not revealed to any of the users. 
	ProviderSecret string `json:"providerSecret"`
	// Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the `providerSecret`. 
	ProviderData *string `json:"providerData,omitempty"`
}

type _ExternalEngineRegistration ExternalEngineRegistration

// NewExternalEngineRegistration instantiates a new ExternalEngineRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEngineRegistration(name string, maxThreads int32, maxHash int32, providerSecret string) *ExternalEngineRegistration {
	this := ExternalEngineRegistration{}
	this.Name = name
	this.MaxThreads = maxThreads
	this.MaxHash = maxHash
	this.ProviderSecret = providerSecret
	return &this
}

// NewExternalEngineRegistrationWithDefaults instantiates a new ExternalEngineRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEngineRegistrationWithDefaults() *ExternalEngineRegistration {
	this := ExternalEngineRegistration{}
	return &this
}

// GetName returns the Name field value
func (o *ExternalEngineRegistration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalEngineRegistration) SetName(v string) {
	o.Name = v
}

// GetMaxThreads returns the MaxThreads field value
func (o *ExternalEngineRegistration) GetMaxThreads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxThreads
}

// GetMaxThreadsOk returns a tuple with the MaxThreads field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetMaxThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxThreads, true
}

// SetMaxThreads sets field value
func (o *ExternalEngineRegistration) SetMaxThreads(v int32) {
	o.MaxThreads = v
}

// GetMaxHash returns the MaxHash field value
func (o *ExternalEngineRegistration) GetMaxHash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxHash
}

// GetMaxHashOk returns a tuple with the MaxHash field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetMaxHashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxHash, true
}

// SetMaxHash sets field value
func (o *ExternalEngineRegistration) SetMaxHash(v int32) {
	o.MaxHash = v
}

// GetVariants returns the Variants field value if set, zero value otherwise.
func (o *ExternalEngineRegistration) GetVariants() []UciVariant {
	if o == nil || IsNil(o.Variants) {
		var ret []UciVariant
		return ret
	}
	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetVariantsOk() ([]UciVariant, bool) {
	if o == nil || IsNil(o.Variants) {
		return nil, false
	}
	return o.Variants, true
}

// HasVariants returns a boolean if a field has been set.
func (o *ExternalEngineRegistration) HasVariants() bool {
	if o != nil && !IsNil(o.Variants) {
		return true
	}

	return false
}

// SetVariants gets a reference to the given []UciVariant and assigns it to the Variants field.
func (o *ExternalEngineRegistration) SetVariants(v []UciVariant) {
	o.Variants = v
}

// GetProviderSecret returns the ProviderSecret field value
func (o *ExternalEngineRegistration) GetProviderSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderSecret
}

// GetProviderSecretOk returns a tuple with the ProviderSecret field value
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetProviderSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderSecret, true
}

// SetProviderSecret sets field value
func (o *ExternalEngineRegistration) SetProviderSecret(v string) {
	o.ProviderSecret = v
}

// GetProviderData returns the ProviderData field value if set, zero value otherwise.
func (o *ExternalEngineRegistration) GetProviderData() string {
	if o == nil || IsNil(o.ProviderData) {
		var ret string
		return ret
	}
	return *o.ProviderData
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalEngineRegistration) GetProviderDataOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderData) {
		return nil, false
	}
	return o.ProviderData, true
}

// HasProviderData returns a boolean if a field has been set.
func (o *ExternalEngineRegistration) HasProviderData() bool {
	if o != nil && !IsNil(o.ProviderData) {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given string and assigns it to the ProviderData field.
func (o *ExternalEngineRegistration) SetProviderData(v string) {
	o.ProviderData = &v
}

func (o ExternalEngineRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEngineRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["maxThreads"] = o.MaxThreads
	toSerialize["maxHash"] = o.MaxHash
	if !IsNil(o.Variants) {
		toSerialize["variants"] = o.Variants
	}
	toSerialize["providerSecret"] = o.ProviderSecret
	if !IsNil(o.ProviderData) {
		toSerialize["providerData"] = o.ProviderData
	}
	return toSerialize, nil
}

func (o *ExternalEngineRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"maxThreads",
		"maxHash",
		"providerSecret",
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

	varExternalEngineRegistration := _ExternalEngineRegistration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalEngineRegistration)

	if err != nil {
		return err
	}

	*o = ExternalEngineRegistration(varExternalEngineRegistration)

	return err
}

type NullableExternalEngineRegistration struct {
	value *ExternalEngineRegistration
	isSet bool
}

func (v NullableExternalEngineRegistration) Get() *ExternalEngineRegistration {
	return v.value
}

func (v *NullableExternalEngineRegistration) Set(val *ExternalEngineRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEngineRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEngineRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEngineRegistration(val *ExternalEngineRegistration) *NullableExternalEngineRegistration {
	return &NullableExternalEngineRegistration{value: val, isSet: true}
}

func (v NullableExternalEngineRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEngineRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


