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

// checks if the ApiUsersStatus200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiUsersStatus200ResponseInner{}

// ApiUsersStatus200ResponseInner struct for ApiUsersStatus200ResponseInner
type ApiUsersStatus200ResponseInner struct {
	Id string `json:"id"`
	Name string `json:"name"`
	// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
	Flair *string `json:"flair,omitempty"`
	Title *Title `json:"title,omitempty"`
	Online *bool `json:"online,omitempty"`
	Playing *bool `json:"playing,omitempty"`
	Streaming *bool `json:"streaming,omitempty"`
	// Use patronColor value instead to determine if player is a patron. 
	// Deprecated
	Patron *bool `json:"patron,omitempty"`
	// Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron. 
	PatronColor *int32 `json:"patronColor,omitempty"`
}

type _ApiUsersStatus200ResponseInner ApiUsersStatus200ResponseInner

// NewApiUsersStatus200ResponseInner instantiates a new ApiUsersStatus200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiUsersStatus200ResponseInner(id string, name string) *ApiUsersStatus200ResponseInner {
	this := ApiUsersStatus200ResponseInner{}
	this.Id = id
	this.Name = name
	return &this
}

// NewApiUsersStatus200ResponseInnerWithDefaults instantiates a new ApiUsersStatus200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiUsersStatus200ResponseInnerWithDefaults() *ApiUsersStatus200ResponseInner {
	this := ApiUsersStatus200ResponseInner{}
	return &this
}

// GetId returns the Id field value
func (o *ApiUsersStatus200ResponseInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ApiUsersStatus200ResponseInner) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ApiUsersStatus200ResponseInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ApiUsersStatus200ResponseInner) SetName(v string) {
	o.Name = v
}

// GetFlair returns the Flair field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetFlair() string {
	if o == nil || IsNil(o.Flair) {
		var ret string
		return ret
	}
	return *o.Flair
}

// GetFlairOk returns a tuple with the Flair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetFlairOk() (*string, bool) {
	if o == nil || IsNil(o.Flair) {
		return nil, false
	}
	return o.Flair, true
}

// HasFlair returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasFlair() bool {
	if o != nil && !IsNil(o.Flair) {
		return true
	}

	return false
}

// SetFlair gets a reference to the given string and assigns it to the Flair field.
func (o *ApiUsersStatus200ResponseInner) SetFlair(v string) {
	o.Flair = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetTitle() Title {
	if o == nil || IsNil(o.Title) {
		var ret Title
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetTitleOk() (*Title, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given Title and assigns it to the Title field.
func (o *ApiUsersStatus200ResponseInner) SetTitle(v Title) {
	o.Title = &v
}

// GetOnline returns the Online field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetOnline() bool {
	if o == nil || IsNil(o.Online) {
		var ret bool
		return ret
	}
	return *o.Online
}

// GetOnlineOk returns a tuple with the Online field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetOnlineOk() (*bool, bool) {
	if o == nil || IsNil(o.Online) {
		return nil, false
	}
	return o.Online, true
}

// HasOnline returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasOnline() bool {
	if o != nil && !IsNil(o.Online) {
		return true
	}

	return false
}

// SetOnline gets a reference to the given bool and assigns it to the Online field.
func (o *ApiUsersStatus200ResponseInner) SetOnline(v bool) {
	o.Online = &v
}

// GetPlaying returns the Playing field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetPlaying() bool {
	if o == nil || IsNil(o.Playing) {
		var ret bool
		return ret
	}
	return *o.Playing
}

// GetPlayingOk returns a tuple with the Playing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetPlayingOk() (*bool, bool) {
	if o == nil || IsNil(o.Playing) {
		return nil, false
	}
	return o.Playing, true
}

// HasPlaying returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasPlaying() bool {
	if o != nil && !IsNil(o.Playing) {
		return true
	}

	return false
}

// SetPlaying gets a reference to the given bool and assigns it to the Playing field.
func (o *ApiUsersStatus200ResponseInner) SetPlaying(v bool) {
	o.Playing = &v
}

// GetStreaming returns the Streaming field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetStreaming() bool {
	if o == nil || IsNil(o.Streaming) {
		var ret bool
		return ret
	}
	return *o.Streaming
}

// GetStreamingOk returns a tuple with the Streaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.Streaming) {
		return nil, false
	}
	return o.Streaming, true
}

// HasStreaming returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasStreaming() bool {
	if o != nil && !IsNil(o.Streaming) {
		return true
	}

	return false
}

// SetStreaming gets a reference to the given bool and assigns it to the Streaming field.
func (o *ApiUsersStatus200ResponseInner) SetStreaming(v bool) {
	o.Streaming = &v
}

// GetPatron returns the Patron field value if set, zero value otherwise.
// Deprecated
func (o *ApiUsersStatus200ResponseInner) GetPatron() bool {
	if o == nil || IsNil(o.Patron) {
		var ret bool
		return ret
	}
	return *o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *ApiUsersStatus200ResponseInner) GetPatronOk() (*bool, bool) {
	if o == nil || IsNil(o.Patron) {
		return nil, false
	}
	return o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasPatron() bool {
	if o != nil && !IsNil(o.Patron) {
		return true
	}

	return false
}

// SetPatron gets a reference to the given bool and assigns it to the Patron field.
// Deprecated
func (o *ApiUsersStatus200ResponseInner) SetPatron(v bool) {
	o.Patron = &v
}

// GetPatronColor returns the PatronColor field value if set, zero value otherwise.
func (o *ApiUsersStatus200ResponseInner) GetPatronColor() int32 {
	if o == nil || IsNil(o.PatronColor) {
		var ret int32
		return ret
	}
	return *o.PatronColor
}

// GetPatronColorOk returns a tuple with the PatronColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiUsersStatus200ResponseInner) GetPatronColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PatronColor) {
		return nil, false
	}
	return o.PatronColor, true
}

// HasPatronColor returns a boolean if a field has been set.
func (o *ApiUsersStatus200ResponseInner) HasPatronColor() bool {
	if o != nil && !IsNil(o.PatronColor) {
		return true
	}

	return false
}

// SetPatronColor gets a reference to the given int32 and assigns it to the PatronColor field.
func (o *ApiUsersStatus200ResponseInner) SetPatronColor(v int32) {
	o.PatronColor = &v
}

func (o ApiUsersStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiUsersStatus200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Flair) {
		toSerialize["flair"] = o.Flair
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Online) {
		toSerialize["online"] = o.Online
	}
	if !IsNil(o.Playing) {
		toSerialize["playing"] = o.Playing
	}
	if !IsNil(o.Streaming) {
		toSerialize["streaming"] = o.Streaming
	}
	if !IsNil(o.Patron) {
		toSerialize["patron"] = o.Patron
	}
	if !IsNil(o.PatronColor) {
		toSerialize["patronColor"] = o.PatronColor
	}
	return toSerialize, nil
}

func (o *ApiUsersStatus200ResponseInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varApiUsersStatus200ResponseInner := _ApiUsersStatus200ResponseInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApiUsersStatus200ResponseInner)

	if err != nil {
		return err
	}

	*o = ApiUsersStatus200ResponseInner(varApiUsersStatus200ResponseInner)

	return err
}

type NullableApiUsersStatus200ResponseInner struct {
	value *ApiUsersStatus200ResponseInner
	isSet bool
}

func (v NullableApiUsersStatus200ResponseInner) Get() *ApiUsersStatus200ResponseInner {
	return v.value
}

func (v *NullableApiUsersStatus200ResponseInner) Set(val *ApiUsersStatus200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableApiUsersStatus200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableApiUsersStatus200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiUsersStatus200ResponseInner(val *ApiUsersStatus200ResponseInner) *NullableApiUsersStatus200ResponseInner {
	return &NullableApiUsersStatus200ResponseInner{value: val, isSet: true}
}

func (v NullableApiUsersStatus200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiUsersStatus200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


