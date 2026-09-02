/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.167
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Simul type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Simul{}

// Simul struct for Simul
type Simul struct {
	Id string `json:"id"`
	Host SimulHost `json:"host"`
	Name string `json:"name"`
	FullName string `json:"fullName"`
	Variants []SimulVariantsInner `json:"variants"`
	IsCreated bool `json:"isCreated"`
	IsFinished bool `json:"isFinished"`
	IsRunning bool `json:"isRunning"`
	Text *string `json:"text,omitempty"`
	EstimatedStartAt *int32 `json:"estimatedStartAt,omitempty"`
	StartedAt *int32 `json:"startedAt,omitempty"`
	FinishedAt *int32 `json:"finishedAt,omitempty"`
	NbApplicants int32 `json:"nbApplicants"`
	NbPairings int32 `json:"nbPairings"`
}

type _Simul Simul

// NewSimul instantiates a new Simul object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimul(id string, host SimulHost, name string, fullName string, variants []SimulVariantsInner, isCreated bool, isFinished bool, isRunning bool, nbApplicants int32, nbPairings int32) *Simul {
	this := Simul{}
	this.Id = id
	this.Host = host
	this.Name = name
	this.FullName = fullName
	this.Variants = variants
	this.IsCreated = isCreated
	this.IsFinished = isFinished
	this.IsRunning = isRunning
	this.NbApplicants = nbApplicants
	this.NbPairings = nbPairings
	return &this
}

// NewSimulWithDefaults instantiates a new Simul object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimulWithDefaults() *Simul {
	this := Simul{}
	return &this
}

// GetId returns the Id field value
func (o *Simul) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Simul) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Simul) SetId(v string) {
	o.Id = v
}

// GetHost returns the Host field value
func (o *Simul) GetHost() SimulHost {
	if o == nil {
		var ret SimulHost
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *Simul) GetHostOk() (*SimulHost, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *Simul) SetHost(v SimulHost) {
	o.Host = v
}

// GetName returns the Name field value
func (o *Simul) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Simul) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Simul) SetName(v string) {
	o.Name = v
}

// GetFullName returns the FullName field value
func (o *Simul) GetFullName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullName
}

// GetFullNameOk returns a tuple with the FullName field value
// and a boolean to check if the value has been set.
func (o *Simul) GetFullNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullName, true
}

// SetFullName sets field value
func (o *Simul) SetFullName(v string) {
	o.FullName = v
}

// GetVariants returns the Variants field value
func (o *Simul) GetVariants() []SimulVariantsInner {
	if o == nil {
		var ret []SimulVariantsInner
		return ret
	}

	return o.Variants
}

// GetVariantsOk returns a tuple with the Variants field value
// and a boolean to check if the value has been set.
func (o *Simul) GetVariantsOk() ([]SimulVariantsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Variants, true
}

// SetVariants sets field value
func (o *Simul) SetVariants(v []SimulVariantsInner) {
	o.Variants = v
}

// GetIsCreated returns the IsCreated field value
func (o *Simul) GetIsCreated() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsCreated
}

// GetIsCreatedOk returns a tuple with the IsCreated field value
// and a boolean to check if the value has been set.
func (o *Simul) GetIsCreatedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsCreated, true
}

// SetIsCreated sets field value
func (o *Simul) SetIsCreated(v bool) {
	o.IsCreated = v
}

// GetIsFinished returns the IsFinished field value
func (o *Simul) GetIsFinished() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsFinished
}

// GetIsFinishedOk returns a tuple with the IsFinished field value
// and a boolean to check if the value has been set.
func (o *Simul) GetIsFinishedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsFinished, true
}

// SetIsFinished sets field value
func (o *Simul) SetIsFinished(v bool) {
	o.IsFinished = v
}

// GetIsRunning returns the IsRunning field value
func (o *Simul) GetIsRunning() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRunning
}

// GetIsRunningOk returns a tuple with the IsRunning field value
// and a boolean to check if the value has been set.
func (o *Simul) GetIsRunningOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRunning, true
}

// SetIsRunning sets field value
func (o *Simul) SetIsRunning(v bool) {
	o.IsRunning = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Simul) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Simul) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Simul) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *Simul) SetText(v string) {
	o.Text = &v
}

// GetEstimatedStartAt returns the EstimatedStartAt field value if set, zero value otherwise.
func (o *Simul) GetEstimatedStartAt() int32 {
	if o == nil || IsNil(o.EstimatedStartAt) {
		var ret int32
		return ret
	}
	return *o.EstimatedStartAt
}

// GetEstimatedStartAtOk returns a tuple with the EstimatedStartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Simul) GetEstimatedStartAtOk() (*int32, bool) {
	if o == nil || IsNil(o.EstimatedStartAt) {
		return nil, false
	}
	return o.EstimatedStartAt, true
}

// HasEstimatedStartAt returns a boolean if a field has been set.
func (o *Simul) HasEstimatedStartAt() bool {
	if o != nil && !IsNil(o.EstimatedStartAt) {
		return true
	}

	return false
}

// SetEstimatedStartAt gets a reference to the given int32 and assigns it to the EstimatedStartAt field.
func (o *Simul) SetEstimatedStartAt(v int32) {
	o.EstimatedStartAt = &v
}

// GetStartedAt returns the StartedAt field value if set, zero value otherwise.
func (o *Simul) GetStartedAt() int32 {
	if o == nil || IsNil(o.StartedAt) {
		var ret int32
		return ret
	}
	return *o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Simul) GetStartedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.StartedAt) {
		return nil, false
	}
	return o.StartedAt, true
}

// HasStartedAt returns a boolean if a field has been set.
func (o *Simul) HasStartedAt() bool {
	if o != nil && !IsNil(o.StartedAt) {
		return true
	}

	return false
}

// SetStartedAt gets a reference to the given int32 and assigns it to the StartedAt field.
func (o *Simul) SetStartedAt(v int32) {
	o.StartedAt = &v
}

// GetFinishedAt returns the FinishedAt field value if set, zero value otherwise.
func (o *Simul) GetFinishedAt() int32 {
	if o == nil || IsNil(o.FinishedAt) {
		var ret int32
		return ret
	}
	return *o.FinishedAt
}

// GetFinishedAtOk returns a tuple with the FinishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Simul) GetFinishedAtOk() (*int32, bool) {
	if o == nil || IsNil(o.FinishedAt) {
		return nil, false
	}
	return o.FinishedAt, true
}

// HasFinishedAt returns a boolean if a field has been set.
func (o *Simul) HasFinishedAt() bool {
	if o != nil && !IsNil(o.FinishedAt) {
		return true
	}

	return false
}

// SetFinishedAt gets a reference to the given int32 and assigns it to the FinishedAt field.
func (o *Simul) SetFinishedAt(v int32) {
	o.FinishedAt = &v
}

// GetNbApplicants returns the NbApplicants field value
func (o *Simul) GetNbApplicants() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbApplicants
}

// GetNbApplicantsOk returns a tuple with the NbApplicants field value
// and a boolean to check if the value has been set.
func (o *Simul) GetNbApplicantsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbApplicants, true
}

// SetNbApplicants sets field value
func (o *Simul) SetNbApplicants(v int32) {
	o.NbApplicants = v
}

// GetNbPairings returns the NbPairings field value
func (o *Simul) GetNbPairings() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NbPairings
}

// GetNbPairingsOk returns a tuple with the NbPairings field value
// and a boolean to check if the value has been set.
func (o *Simul) GetNbPairingsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NbPairings, true
}

// SetNbPairings sets field value
func (o *Simul) SetNbPairings(v int32) {
	o.NbPairings = v
}

func (o Simul) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Simul) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["host"] = o.Host
	toSerialize["name"] = o.Name
	toSerialize["fullName"] = o.FullName
	toSerialize["variants"] = o.Variants
	toSerialize["isCreated"] = o.IsCreated
	toSerialize["isFinished"] = o.IsFinished
	toSerialize["isRunning"] = o.IsRunning
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.EstimatedStartAt) {
		toSerialize["estimatedStartAt"] = o.EstimatedStartAt
	}
	if !IsNil(o.StartedAt) {
		toSerialize["startedAt"] = o.StartedAt
	}
	if !IsNil(o.FinishedAt) {
		toSerialize["finishedAt"] = o.FinishedAt
	}
	toSerialize["nbApplicants"] = o.NbApplicants
	toSerialize["nbPairings"] = o.NbPairings
	return toSerialize, nil
}

func (o *Simul) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"host",
		"name",
		"fullName",
		"variants",
		"isCreated",
		"isFinished",
		"isRunning",
		"nbApplicants",
		"nbPairings",
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

	varSimul := _Simul{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimul)

	if err != nil {
		return err
	}

	*o = Simul(varSimul)

	return err
}

type NullableSimul struct {
	value *Simul
	isSet bool
}

func (v NullableSimul) Get() *Simul {
	return v.value
}

func (v *NullableSimul) Set(val *Simul) {
	v.value = val
	v.isSet = true
}

func (v NullableSimul) IsSet() bool {
	return v.isSet
}

func (v *NullableSimul) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimul(val *Simul) *NullableSimul {
	return &NullableSimul{value: val, isSet: true}
}

func (v NullableSimul) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimul) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


