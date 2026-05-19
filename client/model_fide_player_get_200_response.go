/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.143
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lichess

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FidePlayerGet200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FidePlayerGet200Response{}

// FidePlayerGet200Response struct for FidePlayerGet200Response
type FidePlayerGet200Response struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
	// only appears if the user is a titled player or a bot user
	Title NullableString `json:"title,omitempty"`
	Federation string `json:"federation"`
	Year NullableInt32 `json:"year,omitempty"`
	Inactive *int32 `json:"inactive,omitempty"`
	Standard *int32 `json:"standard,omitempty"`
	Rapid *int32 `json:"rapid,omitempty"`
	Blitz *int32 `json:"blitz,omitempty"`
	// FIDE uses mandatory binary gender.
	Gender *string `json:"gender,omitempty"`
	Photo *BroadcastsOfficial200ResponsePhotosValue `json:"photo,omitempty"`
}

type _FidePlayerGet200Response FidePlayerGet200Response

// NewFidePlayerGet200Response instantiates a new FidePlayerGet200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFidePlayerGet200Response(id int32, name string, federation string) *FidePlayerGet200Response {
	this := FidePlayerGet200Response{}
	this.Id = id
	this.Name = name
	this.Federation = federation
	return &this
}

// NewFidePlayerGet200ResponseWithDefaults instantiates a new FidePlayerGet200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFidePlayerGet200ResponseWithDefaults() *FidePlayerGet200Response {
	this := FidePlayerGet200Response{}
	return &this
}

// GetId returns the Id field value
func (o *FidePlayerGet200Response) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *FidePlayerGet200Response) SetId(v int32) {
	o.Id = v
}

// GetName returns the Name field value
func (o *FidePlayerGet200Response) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FidePlayerGet200Response) SetName(v string) {
	o.Name = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FidePlayerGet200Response) GetTitle() string {
	if o == nil || IsNil(o.Title.Get()) {
		var ret string
		return ret
	}
	return *o.Title.Get()
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FidePlayerGet200Response) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Title.Get(), o.Title.IsSet()
}

// HasTitle returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasTitle() bool {
	if o != nil && o.Title.IsSet() {
		return true
	}

	return false
}

// SetTitle gets a reference to the given NullableString and assigns it to the Title field.
func (o *FidePlayerGet200Response) SetTitle(v string) {
	o.Title.Set(&v)
}
// SetTitleNil sets the value for Title to be an explicit nil
func (o *FidePlayerGet200Response) SetTitleNil() {
	o.Title.Set(nil)
}

// UnsetTitle ensures that no value is present for Title, not even an explicit nil
func (o *FidePlayerGet200Response) UnsetTitle() {
	o.Title.Unset()
}

// GetFederation returns the Federation field value
func (o *FidePlayerGet200Response) GetFederation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Federation
}

// GetFederationOk returns a tuple with the Federation field value
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetFederationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Federation, true
}

// SetFederation sets field value
func (o *FidePlayerGet200Response) SetFederation(v string) {
	o.Federation = v
}

// GetYear returns the Year field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *FidePlayerGet200Response) GetYear() int32 {
	if o == nil || IsNil(o.Year.Get()) {
		var ret int32
		return ret
	}
	return *o.Year.Get()
}

// GetYearOk returns a tuple with the Year field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *FidePlayerGet200Response) GetYearOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Year.Get(), o.Year.IsSet()
}

// HasYear returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasYear() bool {
	if o != nil && o.Year.IsSet() {
		return true
	}

	return false
}

// SetYear gets a reference to the given NullableInt32 and assigns it to the Year field.
func (o *FidePlayerGet200Response) SetYear(v int32) {
	o.Year.Set(&v)
}
// SetYearNil sets the value for Year to be an explicit nil
func (o *FidePlayerGet200Response) SetYearNil() {
	o.Year.Set(nil)
}

// UnsetYear ensures that no value is present for Year, not even an explicit nil
func (o *FidePlayerGet200Response) UnsetYear() {
	o.Year.Unset()
}

// GetInactive returns the Inactive field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetInactive() int32 {
	if o == nil || IsNil(o.Inactive) {
		var ret int32
		return ret
	}
	return *o.Inactive
}

// GetInactiveOk returns a tuple with the Inactive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetInactiveOk() (*int32, bool) {
	if o == nil || IsNil(o.Inactive) {
		return nil, false
	}
	return o.Inactive, true
}

// HasInactive returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasInactive() bool {
	if o != nil && !IsNil(o.Inactive) {
		return true
	}

	return false
}

// SetInactive gets a reference to the given int32 and assigns it to the Inactive field.
func (o *FidePlayerGet200Response) SetInactive(v int32) {
	o.Inactive = &v
}

// GetStandard returns the Standard field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetStandard() int32 {
	if o == nil || IsNil(o.Standard) {
		var ret int32
		return ret
	}
	return *o.Standard
}

// GetStandardOk returns a tuple with the Standard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetStandardOk() (*int32, bool) {
	if o == nil || IsNil(o.Standard) {
		return nil, false
	}
	return o.Standard, true
}

// HasStandard returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasStandard() bool {
	if o != nil && !IsNil(o.Standard) {
		return true
	}

	return false
}

// SetStandard gets a reference to the given int32 and assigns it to the Standard field.
func (o *FidePlayerGet200Response) SetStandard(v int32) {
	o.Standard = &v
}

// GetRapid returns the Rapid field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetRapid() int32 {
	if o == nil || IsNil(o.Rapid) {
		var ret int32
		return ret
	}
	return *o.Rapid
}

// GetRapidOk returns a tuple with the Rapid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetRapidOk() (*int32, bool) {
	if o == nil || IsNil(o.Rapid) {
		return nil, false
	}
	return o.Rapid, true
}

// HasRapid returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasRapid() bool {
	if o != nil && !IsNil(o.Rapid) {
		return true
	}

	return false
}

// SetRapid gets a reference to the given int32 and assigns it to the Rapid field.
func (o *FidePlayerGet200Response) SetRapid(v int32) {
	o.Rapid = &v
}

// GetBlitz returns the Blitz field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetBlitz() int32 {
	if o == nil || IsNil(o.Blitz) {
		var ret int32
		return ret
	}
	return *o.Blitz
}

// GetBlitzOk returns a tuple with the Blitz field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetBlitzOk() (*int32, bool) {
	if o == nil || IsNil(o.Blitz) {
		return nil, false
	}
	return o.Blitz, true
}

// HasBlitz returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasBlitz() bool {
	if o != nil && !IsNil(o.Blitz) {
		return true
	}

	return false
}

// SetBlitz gets a reference to the given int32 and assigns it to the Blitz field.
func (o *FidePlayerGet200Response) SetBlitz(v int32) {
	o.Blitz = &v
}

// GetGender returns the Gender field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetGender() string {
	if o == nil || IsNil(o.Gender) {
		var ret string
		return ret
	}
	return *o.Gender
}

// GetGenderOk returns a tuple with the Gender field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetGenderOk() (*string, bool) {
	if o == nil || IsNil(o.Gender) {
		return nil, false
	}
	return o.Gender, true
}

// HasGender returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasGender() bool {
	if o != nil && !IsNil(o.Gender) {
		return true
	}

	return false
}

// SetGender gets a reference to the given string and assigns it to the Gender field.
func (o *FidePlayerGet200Response) SetGender(v string) {
	o.Gender = &v
}

// GetPhoto returns the Photo field value if set, zero value otherwise.
func (o *FidePlayerGet200Response) GetPhoto() BroadcastsOfficial200ResponsePhotosValue {
	if o == nil || IsNil(o.Photo) {
		var ret BroadcastsOfficial200ResponsePhotosValue
		return ret
	}
	return *o.Photo
}

// GetPhotoOk returns a tuple with the Photo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FidePlayerGet200Response) GetPhotoOk() (*BroadcastsOfficial200ResponsePhotosValue, bool) {
	if o == nil || IsNil(o.Photo) {
		return nil, false
	}
	return o.Photo, true
}

// HasPhoto returns a boolean if a field has been set.
func (o *FidePlayerGet200Response) HasPhoto() bool {
	if o != nil && !IsNil(o.Photo) {
		return true
	}

	return false
}

// SetPhoto gets a reference to the given BroadcastsOfficial200ResponsePhotosValue and assigns it to the Photo field.
func (o *FidePlayerGet200Response) SetPhoto(v BroadcastsOfficial200ResponsePhotosValue) {
	o.Photo = &v
}

func (o FidePlayerGet200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FidePlayerGet200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if o.Title.IsSet() {
		toSerialize["title"] = o.Title.Get()
	}
	toSerialize["federation"] = o.Federation
	if o.Year.IsSet() {
		toSerialize["year"] = o.Year.Get()
	}
	if !IsNil(o.Inactive) {
		toSerialize["inactive"] = o.Inactive
	}
	if !IsNil(o.Standard) {
		toSerialize["standard"] = o.Standard
	}
	if !IsNil(o.Rapid) {
		toSerialize["rapid"] = o.Rapid
	}
	if !IsNil(o.Blitz) {
		toSerialize["blitz"] = o.Blitz
	}
	if !IsNil(o.Gender) {
		toSerialize["gender"] = o.Gender
	}
	if !IsNil(o.Photo) {
		toSerialize["photo"] = o.Photo
	}
	return toSerialize, nil
}

func (o *FidePlayerGet200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"federation",
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

	varFidePlayerGet200Response := _FidePlayerGet200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFidePlayerGet200Response)

	if err != nil {
		return err
	}

	*o = FidePlayerGet200Response(varFidePlayerGet200Response)

	return err
}

type NullableFidePlayerGet200Response struct {
	value *FidePlayerGet200Response
	isSet bool
}

func (v NullableFidePlayerGet200Response) Get() *FidePlayerGet200Response {
	return v.value
}

func (v *NullableFidePlayerGet200Response) Set(val *FidePlayerGet200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFidePlayerGet200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFidePlayerGet200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFidePlayerGet200Response(val *FidePlayerGet200Response) *NullableFidePlayerGet200Response {
	return &NullableFidePlayerGet200Response{value: val, isSet: true}
}

func (v NullableFidePlayerGet200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFidePlayerGet200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


