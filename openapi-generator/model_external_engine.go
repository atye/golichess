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

// checks if the ExternalEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalEngine{}

// ExternalEngine struct for ExternalEngine
type ExternalEngine struct {
	// Unique engine registration ID.
	Id string `json:"id"`
	// Display name of the engine.
	Name string `json:"name"`
	// A secret token that can be used to [*request* analysis](#tag/external-engine/POST/api/external-engine/{id}/analyse) from this external engine. 
	ClientSecret string `json:"clientSecret"`
	// The user this engine has been registered for.
	UserId string `json:"userId"`
	// Maximum number of available threads.
	MaxThreads int32 `json:"maxThreads"`
	// Maximum available hash table size, in MiB.
	MaxHash int32 `json:"maxHash"`
	// List of supported chess variants.
	Variants []UciVariant `json:"variants"`
	// Arbitrary data that the engine provider can use for identification or bookkeeping.  Users can read this information, but updating it requires knowing or changing the `providerSecret`. 
	ProviderData NullableString `json:"providerData,omitempty"`
}

type _ExternalEngine ExternalEngine

// NewExternalEngine instantiates a new ExternalEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalEngine(id string, name string, clientSecret string, userId string, maxThreads int32, maxHash int32, variants []UciVariant) *ExternalEngine {
	this := ExternalEngine{}
	this.Id = id
	this.Name = name
	this.ClientSecret = clientSecret
	this.UserId = userId
	this.MaxThreads = maxThreads
	this.MaxHash = maxHash
	this.Variants = variants
	return &this
}

// NewExternalEngineWithDefaults instantiates a new ExternalEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalEngineWithDefaults() *ExternalEngine {
	this := ExternalEngine{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalEngine) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalEngine) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ExternalEngine) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalEngine) SetName(v string) {
	o.Name = v
}

// GetClientSecret returns the ClientSecret field value
func (o *ExternalEngine) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetClientSecretOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *ExternalEngine) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetUserId returns the UserId field value
func (o *ExternalEngine) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *ExternalEngine) SetUserId(v string) {
	o.UserId = v
}

// GetMaxThreads returns the MaxThreads field value
func (o *ExternalEngine) GetMaxThreads() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxThreads
}

// GetMaxThreadsOk returns a tuple with the MaxThreads field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetMaxThreadsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxThreads, true
}

// SetMaxThreads sets field value
func (o *ExternalEngine) SetMaxThreads(v int32) {
	o.MaxThreads = v
}

// GetMaxHash returns the MaxHash field value
func (o *ExternalEngine) GetMaxHash() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxHash
}

// GetMaxHashOk returns a tuple with the MaxHash field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetMaxHashOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxHash, true
}

// SetMaxHash sets field value
func (o *ExternalEngine) SetMaxHash(v int32) {
	o.MaxHash = v
}

// GetVariants returns the Variants field value
func (o *ExternalEngine) GetVariants() []UciVariant {
	if o == nil {
		var ret []UciVariant
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *ExternalEngine) GetVariantsOk() ([]UciVariant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *ExternalEngine) SetVariants(v []UciVariant) {
	o.Variants = v
}

// GetProviderData returns the ProviderData field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ExternalEngine) GetProviderData() string {
	if o == nil || IsNil(o.ProviderData.Get()) {
		var ret string
		return ret
	}
	return *o.ProviderData.Get()
}

// GetProviderDataOk returns a tuple with the ProviderData field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ExternalEngine) GetProviderDataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProviderData.Get(), o.ProviderData.IsSet()
}

// HasProviderData returns a boolean if a field has been set.
func (o *ExternalEngine) HasProviderData() bool {
	if o != nil && o.ProviderData.IsSet() {
		return true
	}

	return false
}

// SetProviderData gets a reference to the given NullableString and assigns it to the ProviderData field.
func (o *ExternalEngine) SetProviderData(v string) {
	o.ProviderData.Set(&v)
}
// SetProviderDataNil sets the value for ProviderData to be an explicit nil
func (o *ExternalEngine) SetProviderDataNil() {
	o.ProviderData.Set(nil)
}

// UnsetProviderData ensures that no value is present for ProviderData, not even an explicit nil
func (o *ExternalEngine) UnsetProviderData() {
	o.ProviderData.Unset()
}

func (o ExternalEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["clientSecret"] = o.ClientSecret
	toSerialize["userId"] = o.UserId
	toSerialize["maxThreads"] = o.MaxThreads
	toSerialize["maxHash"] = o.MaxHash
	toSerialize["variants"] = o.Variants
	if o.ProviderData.IsSet() {
		toSerialize["providerData"] = o.ProviderData.Get()
	}
	return toSerialize, nil
}

func (o *ExternalEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"clientSecret",
		"userId",
		"maxThreads",
		"maxHash",
		"variants",
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

	varExternalEngine := _ExternalEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExternalEngine)

	if err != nil {
		return err
	}

	*o = ExternalEngine(varExternalEngine)

	return err
}

type NullableExternalEngine struct {
	value *ExternalEngine
	isSet bool
}

func (v NullableExternalEngine) Get() *ExternalEngine {
	return v.value
}

func (v *NullableExternalEngine) Set(val *ExternalEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalEngine(val *ExternalEngine) *NullableExternalEngine {
	return &NullableExternalEngine{value: val, isSet: true}
}

func (v NullableExternalEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


