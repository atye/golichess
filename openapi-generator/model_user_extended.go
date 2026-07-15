/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.153
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the UserExtended type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserExtended{}

// UserExtended struct for UserExtended
type UserExtended struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Perfs *Perfs `json:"perfs,omitempty"`
	Title *Title `json:"title,omitempty"`
	// See [available flair list and images](https://github.com/lichess-org/lila/tree/master/public/flair)
	Flair *string `json:"flair,omitempty"`
	CreatedAt *int64 `json:"createdAt,omitempty"`
	// only appears if a user's account is closed
	Disabled *bool `json:"disabled,omitempty"`
	// only appears if a user's account is marked for the violation of [Lichess TOS](https://lichess.org/terms-of-service)
	TosViolation *bool `json:"tosViolation,omitempty"`
	Profile *Profile `json:"profile,omitempty"`
	SeenAt *int64 `json:"seenAt,omitempty"`
	PlayTime *PlayTime `json:"playTime,omitempty"`
	// Use patronColor value instead to determine if player is a patron. 
	// Deprecated
	Patron *bool `json:"patron,omitempty"`
	// Players can choose a color for their Patron wings. See [here for the color mappings](https://github.com/lichess-org/lila/blob/master/ui/lib/css/abstract/_patron-colors.scss).  The presence of this field indicates the player is an active Patron. 
	PatronColor *int32 `json:"patronColor,omitempty"`
	Verified *bool `json:"verified,omitempty"`
	Url string `json:"url"`
	Playing *string `json:"playing,omitempty"`
	Count *Count `json:"count,omitempty"`
	Streaming *bool `json:"streaming,omitempty"`
	Streamer *UserStreamer `json:"streamer,omitempty"`
	// only appears if the request is [authenticated with OAuth2](#description/authentication)
	Followable *bool `json:"followable,omitempty"`
	// only appears if the request is [authenticated with OAuth2](#description/authentication)
	Following *bool `json:"following,omitempty"`
	// only appears if the request is [authenticated with OAuth2](#description/authentication)
	Blocking *bool `json:"blocking,omitempty"`
	FideId *float32 `json:"fideId,omitempty"`
}

type _UserExtended UserExtended

// NewUserExtended instantiates a new UserExtended object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserExtended(id string, username string, url string) *UserExtended {
	this := UserExtended{}
	this.Id = id
	this.Username = username
	this.Url = url
	return &this
}

// NewUserExtendedWithDefaults instantiates a new UserExtended object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserExtendedWithDefaults() *UserExtended {
	this := UserExtended{}
	return &this
}

// GetId returns the Id field value
func (o *UserExtended) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *UserExtended) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *UserExtended) SetId(v string) {
	o.Id = v
}

// GetUsername returns the Username field value
func (o *UserExtended) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserExtended) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserExtended) SetUsername(v string) {
	o.Username = v
}

// GetPerfs returns the Perfs field value if set, zero value otherwise.
func (o *UserExtended) GetPerfs() Perfs {
	if o == nil || IsNil(o.Perfs) {
		var ret Perfs
		return ret
	}
	return *o.Perfs
}

// GetPerfsOk returns a tuple with the Perfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetPerfsOk() (*Perfs, bool) {
	if o == nil || IsNil(o.Perfs) {
		return nil, false
	}
	return o.Perfs, true
}

// HasPerfs returns a boolean if a field has been set.
func (o *UserExtended) HasPerfs() bool {
	if o != nil && !IsNil(o.Perfs) {
		return true
	}

	return false
}

// SetPerfs gets a reference to the given Perfs and assigns it to the Perfs field.
func (o *UserExtended) SetPerfs(v Perfs) {
	o.Perfs = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *UserExtended) GetTitle() Title {
	if o == nil || IsNil(o.Title) {
		var ret Title
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetTitleOk() (*Title, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *UserExtended) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given Title and assigns it to the Title field.
func (o *UserExtended) SetTitle(v Title) {
	o.Title = &v
}

// GetFlair returns the Flair field value if set, zero value otherwise.
func (o *UserExtended) GetFlair() string {
	if o == nil || IsNil(o.Flair) {
		var ret string
		return ret
	}
	return *o.Flair
}

// GetFlairOk returns a tuple with the Flair field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetFlairOk() (*string, bool) {
	if o == nil || IsNil(o.Flair) {
		return nil, false
	}
	return o.Flair, true
}

// HasFlair returns a boolean if a field has been set.
func (o *UserExtended) HasFlair() bool {
	if o != nil && !IsNil(o.Flair) {
		return true
	}

	return false
}

// SetFlair gets a reference to the given string and assigns it to the Flair field.
func (o *UserExtended) SetFlair(v string) {
	o.Flair = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *UserExtended) GetCreatedAt() int64 {
	if o == nil || IsNil(o.CreatedAt) {
		var ret int64
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetCreatedAtOk() (*int64, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *UserExtended) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given int64 and assigns it to the CreatedAt field.
func (o *UserExtended) SetCreatedAt(v int64) {
	o.CreatedAt = &v
}

// GetDisabled returns the Disabled field value if set, zero value otherwise.
func (o *UserExtended) GetDisabled() bool {
	if o == nil || IsNil(o.Disabled) {
		var ret bool
		return ret
	}
	return *o.Disabled
}

// GetDisabledOk returns a tuple with the Disabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetDisabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Disabled) {
		return nil, false
	}
	return o.Disabled, true
}

// HasDisabled returns a boolean if a field has been set.
func (o *UserExtended) HasDisabled() bool {
	if o != nil && !IsNil(o.Disabled) {
		return true
	}

	return false
}

// SetDisabled gets a reference to the given bool and assigns it to the Disabled field.
func (o *UserExtended) SetDisabled(v bool) {
	o.Disabled = &v
}

// GetTosViolation returns the TosViolation field value if set, zero value otherwise.
func (o *UserExtended) GetTosViolation() bool {
	if o == nil || IsNil(o.TosViolation) {
		var ret bool
		return ret
	}
	return *o.TosViolation
}

// GetTosViolationOk returns a tuple with the TosViolation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetTosViolationOk() (*bool, bool) {
	if o == nil || IsNil(o.TosViolation) {
		return nil, false
	}
	return o.TosViolation, true
}

// HasTosViolation returns a boolean if a field has been set.
func (o *UserExtended) HasTosViolation() bool {
	if o != nil && !IsNil(o.TosViolation) {
		return true
	}

	return false
}

// SetTosViolation gets a reference to the given bool and assigns it to the TosViolation field.
func (o *UserExtended) SetTosViolation(v bool) {
	o.TosViolation = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *UserExtended) GetProfile() Profile {
	if o == nil || IsNil(o.Profile) {
		var ret Profile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetProfileOk() (*Profile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *UserExtended) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given Profile and assigns it to the Profile field.
func (o *UserExtended) SetProfile(v Profile) {
	o.Profile = &v
}

// GetSeenAt returns the SeenAt field value if set, zero value otherwise.
func (o *UserExtended) GetSeenAt() int64 {
	if o == nil || IsNil(o.SeenAt) {
		var ret int64
		return ret
	}
	return *o.SeenAt
}

// GetSeenAtOk returns a tuple with the SeenAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetSeenAtOk() (*int64, bool) {
	if o == nil || IsNil(o.SeenAt) {
		return nil, false
	}
	return o.SeenAt, true
}

// HasSeenAt returns a boolean if a field has been set.
func (o *UserExtended) HasSeenAt() bool {
	if o != nil && !IsNil(o.SeenAt) {
		return true
	}

	return false
}

// SetSeenAt gets a reference to the given int64 and assigns it to the SeenAt field.
func (o *UserExtended) SetSeenAt(v int64) {
	o.SeenAt = &v
}

// GetPlayTime returns the PlayTime field value if set, zero value otherwise.
func (o *UserExtended) GetPlayTime() PlayTime {
	if o == nil || IsNil(o.PlayTime) {
		var ret PlayTime
		return ret
	}
	return *o.PlayTime
}

// GetPlayTimeOk returns a tuple with the PlayTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetPlayTimeOk() (*PlayTime, bool) {
	if o == nil || IsNil(o.PlayTime) {
		return nil, false
	}
	return o.PlayTime, true
}

// HasPlayTime returns a boolean if a field has been set.
func (o *UserExtended) HasPlayTime() bool {
	if o != nil && !IsNil(o.PlayTime) {
		return true
	}

	return false
}

// SetPlayTime gets a reference to the given PlayTime and assigns it to the PlayTime field.
func (o *UserExtended) SetPlayTime(v PlayTime) {
	o.PlayTime = &v
}

// GetPatron returns the Patron field value if set, zero value otherwise.
// Deprecated
func (o *UserExtended) GetPatron() bool {
	if o == nil || IsNil(o.Patron) {
		var ret bool
		return ret
	}
	return *o.Patron
}

// GetPatronOk returns a tuple with the Patron field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *UserExtended) GetPatronOk() (*bool, bool) {
	if o == nil || IsNil(o.Patron) {
		return nil, false
	}
	return o.Patron, true
}

// HasPatron returns a boolean if a field has been set.
func (o *UserExtended) HasPatron() bool {
	if o != nil && !IsNil(o.Patron) {
		return true
	}

	return false
}

// SetPatron gets a reference to the given bool and assigns it to the Patron field.
// Deprecated
func (o *UserExtended) SetPatron(v bool) {
	o.Patron = &v
}

// GetPatronColor returns the PatronColor field value if set, zero value otherwise.
func (o *UserExtended) GetPatronColor() int32 {
	if o == nil || IsNil(o.PatronColor) {
		var ret int32
		return ret
	}
	return *o.PatronColor
}

// GetPatronColorOk returns a tuple with the PatronColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetPatronColorOk() (*int32, bool) {
	if o == nil || IsNil(o.PatronColor) {
		return nil, false
	}
	return o.PatronColor, true
}

// HasPatronColor returns a boolean if a field has been set.
func (o *UserExtended) HasPatronColor() bool {
	if o != nil && !IsNil(o.PatronColor) {
		return true
	}

	return false
}

// SetPatronColor gets a reference to the given int32 and assigns it to the PatronColor field.
func (o *UserExtended) SetPatronColor(v int32) {
	o.PatronColor = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *UserExtended) GetVerified() bool {
	if o == nil || IsNil(o.Verified) {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetVerifiedOk() (*bool, bool) {
	if o == nil || IsNil(o.Verified) {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *UserExtended) HasVerified() bool {
	if o != nil && !IsNil(o.Verified) {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *UserExtended) SetVerified(v bool) {
	o.Verified = &v
}

// GetUrl returns the Url field value
func (o *UserExtended) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *UserExtended) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *UserExtended) SetUrl(v string) {
	o.Url = v
}

// GetPlaying returns the Playing field value if set, zero value otherwise.
func (o *UserExtended) GetPlaying() string {
	if o == nil || IsNil(o.Playing) {
		var ret string
		return ret
	}
	return *o.Playing
}

// GetPlayingOk returns a tuple with the Playing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetPlayingOk() (*string, bool) {
	if o == nil || IsNil(o.Playing) {
		return nil, false
	}
	return o.Playing, true
}

// HasPlaying returns a boolean if a field has been set.
func (o *UserExtended) HasPlaying() bool {
	if o != nil && !IsNil(o.Playing) {
		return true
	}

	return false
}

// SetPlaying gets a reference to the given string and assigns it to the Playing field.
func (o *UserExtended) SetPlaying(v string) {
	o.Playing = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *UserExtended) GetCount() Count {
	if o == nil || IsNil(o.Count) {
		var ret Count
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetCountOk() (*Count, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *UserExtended) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given Count and assigns it to the Count field.
func (o *UserExtended) SetCount(v Count) {
	o.Count = &v
}

// GetStreaming returns the Streaming field value if set, zero value otherwise.
func (o *UserExtended) GetStreaming() bool {
	if o == nil || IsNil(o.Streaming) {
		var ret bool
		return ret
	}
	return *o.Streaming
}

// GetStreamingOk returns a tuple with the Streaming field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetStreamingOk() (*bool, bool) {
	if o == nil || IsNil(o.Streaming) {
		return nil, false
	}
	return o.Streaming, true
}

// HasStreaming returns a boolean if a field has been set.
func (o *UserExtended) HasStreaming() bool {
	if o != nil && !IsNil(o.Streaming) {
		return true
	}

	return false
}

// SetStreaming gets a reference to the given bool and assigns it to the Streaming field.
func (o *UserExtended) SetStreaming(v bool) {
	o.Streaming = &v
}

// GetStreamer returns the Streamer field value if set, zero value otherwise.
func (o *UserExtended) GetStreamer() UserStreamer {
	if o == nil || IsNil(o.Streamer) {
		var ret UserStreamer
		return ret
	}
	return *o.Streamer
}

// GetStreamerOk returns a tuple with the Streamer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetStreamerOk() (*UserStreamer, bool) {
	if o == nil || IsNil(o.Streamer) {
		return nil, false
	}
	return o.Streamer, true
}

// HasStreamer returns a boolean if a field has been set.
func (o *UserExtended) HasStreamer() bool {
	if o != nil && !IsNil(o.Streamer) {
		return true
	}

	return false
}

// SetStreamer gets a reference to the given UserStreamer and assigns it to the Streamer field.
func (o *UserExtended) SetStreamer(v UserStreamer) {
	o.Streamer = &v
}

// GetFollowable returns the Followable field value if set, zero value otherwise.
func (o *UserExtended) GetFollowable() bool {
	if o == nil || IsNil(o.Followable) {
		var ret bool
		return ret
	}
	return *o.Followable
}

// GetFollowableOk returns a tuple with the Followable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetFollowableOk() (*bool, bool) {
	if o == nil || IsNil(o.Followable) {
		return nil, false
	}
	return o.Followable, true
}

// HasFollowable returns a boolean if a field has been set.
func (o *UserExtended) HasFollowable() bool {
	if o != nil && !IsNil(o.Followable) {
		return true
	}

	return false
}

// SetFollowable gets a reference to the given bool and assigns it to the Followable field.
func (o *UserExtended) SetFollowable(v bool) {
	o.Followable = &v
}

// GetFollowing returns the Following field value if set, zero value otherwise.
func (o *UserExtended) GetFollowing() bool {
	if o == nil || IsNil(o.Following) {
		var ret bool
		return ret
	}
	return *o.Following
}

// GetFollowingOk returns a tuple with the Following field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetFollowingOk() (*bool, bool) {
	if o == nil || IsNil(o.Following) {
		return nil, false
	}
	return o.Following, true
}

// HasFollowing returns a boolean if a field has been set.
func (o *UserExtended) HasFollowing() bool {
	if o != nil && !IsNil(o.Following) {
		return true
	}

	return false
}

// SetFollowing gets a reference to the given bool and assigns it to the Following field.
func (o *UserExtended) SetFollowing(v bool) {
	o.Following = &v
}

// GetBlocking returns the Blocking field value if set, zero value otherwise.
func (o *UserExtended) GetBlocking() bool {
	if o == nil || IsNil(o.Blocking) {
		var ret bool
		return ret
	}
	return *o.Blocking
}

// GetBlockingOk returns a tuple with the Blocking field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetBlockingOk() (*bool, bool) {
	if o == nil || IsNil(o.Blocking) {
		return nil, false
	}
	return o.Blocking, true
}

// HasBlocking returns a boolean if a field has been set.
func (o *UserExtended) HasBlocking() bool {
	if o != nil && !IsNil(o.Blocking) {
		return true
	}

	return false
}

// SetBlocking gets a reference to the given bool and assigns it to the Blocking field.
func (o *UserExtended) SetBlocking(v bool) {
	o.Blocking = &v
}

// GetFideId returns the FideId field value if set, zero value otherwise.
func (o *UserExtended) GetFideId() float32 {
	if o == nil || IsNil(o.FideId) {
		var ret float32
		return ret
	}
	return *o.FideId
}

// GetFideIdOk returns a tuple with the FideId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserExtended) GetFideIdOk() (*float32, bool) {
	if o == nil || IsNil(o.FideId) {
		return nil, false
	}
	return o.FideId, true
}

// HasFideId returns a boolean if a field has been set.
func (o *UserExtended) HasFideId() bool {
	if o != nil && !IsNil(o.FideId) {
		return true
	}

	return false
}

// SetFideId gets a reference to the given float32 and assigns it to the FideId field.
func (o *UserExtended) SetFideId(v float32) {
	o.FideId = &v
}

func (o UserExtended) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserExtended) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["username"] = o.Username
	if !IsNil(o.Perfs) {
		toSerialize["perfs"] = o.Perfs
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Flair) {
		toSerialize["flair"] = o.Flair
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.Disabled) {
		toSerialize["disabled"] = o.Disabled
	}
	if !IsNil(o.TosViolation) {
		toSerialize["tosViolation"] = o.TosViolation
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.SeenAt) {
		toSerialize["seenAt"] = o.SeenAt
	}
	if !IsNil(o.PlayTime) {
		toSerialize["playTime"] = o.PlayTime
	}
	if !IsNil(o.Patron) {
		toSerialize["patron"] = o.Patron
	}
	if !IsNil(o.PatronColor) {
		toSerialize["patronColor"] = o.PatronColor
	}
	if !IsNil(o.Verified) {
		toSerialize["verified"] = o.Verified
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.Playing) {
		toSerialize["playing"] = o.Playing
	}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.Streaming) {
		toSerialize["streaming"] = o.Streaming
	}
	if !IsNil(o.Streamer) {
		toSerialize["streamer"] = o.Streamer
	}
	if !IsNil(o.Followable) {
		toSerialize["followable"] = o.Followable
	}
	if !IsNil(o.Following) {
		toSerialize["following"] = o.Following
	}
	if !IsNil(o.Blocking) {
		toSerialize["blocking"] = o.Blocking
	}
	if !IsNil(o.FideId) {
		toSerialize["fideId"] = o.FideId
	}
	return toSerialize, nil
}

func (o *UserExtended) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"username",
		"url",
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

	varUserExtended := _UserExtended{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserExtended)

	if err != nil {
		return err
	}

	*o = UserExtended(varUserExtended)

	return err
}

type NullableUserExtended struct {
	value *UserExtended
	isSet bool
}

func (v NullableUserExtended) Get() *UserExtended {
	return v.value
}

func (v *NullableUserExtended) Set(val *UserExtended) {
	v.value = val
	v.isSet = true
}

func (v NullableUserExtended) IsSet() bool {
	return v.isSet
}

func (v *NullableUserExtended) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserExtended(val *UserExtended) *NullableUserExtended {
	return &NullableUserExtended{value: val, isSet: true}
}

func (v NullableUserExtended) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserExtended) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


