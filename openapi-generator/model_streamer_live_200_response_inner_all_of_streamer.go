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
)

// checks if the StreamerLive200ResponseInnerAllOfStreamer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StreamerLive200ResponseInnerAllOfStreamer{}

// StreamerLive200ResponseInnerAllOfStreamer struct for StreamerLive200ResponseInnerAllOfStreamer
type StreamerLive200ResponseInnerAllOfStreamer struct {
	Name *string `json:"name,omitempty"`
	Headline *string `json:"headline,omitempty"`
	Description *string `json:"description,omitempty"`
	Twitch *string `json:"twitch,omitempty"`
	Youtube *string `json:"youtube,omitempty"`
	Image *string `json:"image,omitempty"`
}

// NewStreamerLive200ResponseInnerAllOfStreamer instantiates a new StreamerLive200ResponseInnerAllOfStreamer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamerLive200ResponseInnerAllOfStreamer() *StreamerLive200ResponseInnerAllOfStreamer {
	this := StreamerLive200ResponseInnerAllOfStreamer{}
	return &this
}

// NewStreamerLive200ResponseInnerAllOfStreamerWithDefaults instantiates a new StreamerLive200ResponseInnerAllOfStreamer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamerLive200ResponseInnerAllOfStreamerWithDefaults() *StreamerLive200ResponseInnerAllOfStreamer {
	this := StreamerLive200ResponseInnerAllOfStreamer{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetName(v string) {
	o.Name = &v
}

// GetHeadline returns the Headline field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetHeadline() string {
	if o == nil || IsNil(o.Headline) {
		var ret string
		return ret
	}
	return *o.Headline
}

// GetHeadlineOk returns a tuple with the Headline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetHeadlineOk() (*string, bool) {
	if o == nil || IsNil(o.Headline) {
		return nil, false
	}
	return o.Headline, true
}

// HasHeadline returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasHeadline() bool {
	if o != nil && !IsNil(o.Headline) {
		return true
	}

	return false
}

// SetHeadline gets a reference to the given string and assigns it to the Headline field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetHeadline(v string) {
	o.Headline = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetDescription(v string) {
	o.Description = &v
}

// GetTwitch returns the Twitch field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetTwitch() string {
	if o == nil || IsNil(o.Twitch) {
		var ret string
		return ret
	}
	return *o.Twitch
}

// GetTwitchOk returns a tuple with the Twitch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetTwitchOk() (*string, bool) {
	if o == nil || IsNil(o.Twitch) {
		return nil, false
	}
	return o.Twitch, true
}

// HasTwitch returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasTwitch() bool {
	if o != nil && !IsNil(o.Twitch) {
		return true
	}

	return false
}

// SetTwitch gets a reference to the given string and assigns it to the Twitch field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetTwitch(v string) {
	o.Twitch = &v
}

// GetYoutube returns the Youtube field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetYoutube() string {
	if o == nil || IsNil(o.Youtube) {
		var ret string
		return ret
	}
	return *o.Youtube
}

// GetYoutubeOk returns a tuple with the Youtube field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetYoutubeOk() (*string, bool) {
	if o == nil || IsNil(o.Youtube) {
		return nil, false
	}
	return o.Youtube, true
}

// HasYoutube returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasYoutube() bool {
	if o != nil && !IsNil(o.Youtube) {
		return true
	}

	return false
}

// SetYoutube gets a reference to the given string and assigns it to the Youtube field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetYoutube(v string) {
	o.Youtube = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *StreamerLive200ResponseInnerAllOfStreamer) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *StreamerLive200ResponseInnerAllOfStreamer) SetImage(v string) {
	o.Image = &v
}

func (o StreamerLive200ResponseInnerAllOfStreamer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StreamerLive200ResponseInnerAllOfStreamer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Headline) {
		toSerialize["headline"] = o.Headline
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Twitch) {
		toSerialize["twitch"] = o.Twitch
	}
	if !IsNil(o.Youtube) {
		toSerialize["youtube"] = o.Youtube
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	return toSerialize, nil
}

type NullableStreamerLive200ResponseInnerAllOfStreamer struct {
	value *StreamerLive200ResponseInnerAllOfStreamer
	isSet bool
}

func (v NullableStreamerLive200ResponseInnerAllOfStreamer) Get() *StreamerLive200ResponseInnerAllOfStreamer {
	return v.value
}

func (v *NullableStreamerLive200ResponseInnerAllOfStreamer) Set(val *StreamerLive200ResponseInnerAllOfStreamer) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamerLive200ResponseInnerAllOfStreamer) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamerLive200ResponseInnerAllOfStreamer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamerLive200ResponseInnerAllOfStreamer(val *StreamerLive200ResponseInnerAllOfStreamer) *NullableStreamerLive200ResponseInnerAllOfStreamer {
	return &NullableStreamerLive200ResponseInnerAllOfStreamer{value: val, isSet: true}
}

func (v NullableStreamerLive200ResponseInnerAllOfStreamer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamerLive200ResponseInnerAllOfStreamer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


