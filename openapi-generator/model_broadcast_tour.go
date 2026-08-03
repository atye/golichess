/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.158
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the BroadcastTour type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BroadcastTour{}

// BroadcastTour struct for BroadcastTour
type BroadcastTour struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	CreatedAt int32 `json:"createdAt"`
	// Start and end dates of the tournament, as Unix timestamps in milliseconds
	Dates []int64 `json:"dates,omitempty"`
	Info *BroadcastTourInfo `json:"info,omitempty"`
	// Used to designate featured tournaments on Lichess
	Tier *int32 `json:"tier,omitempty"`
	Image *string `json:"image,omitempty"`
	// Full tournament description in markdown format, or in HTML if the html=1 query parameter is set.
	Description *string `json:"description,omitempty"`
	TeamTable *bool `json:"teamTable,omitempty"`
	ShowTeamScores *bool `json:"showTeamScores,omitempty"`
	Url string `json:"url"`
	CommunityOwner *LightUser `json:"communityOwner,omitempty"`
}

type _BroadcastTour BroadcastTour

// NewBroadcastTour instantiates a new BroadcastTour object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBroadcastTour(id string, name string, slug string, createdAt int32, url string) *BroadcastTour {
	this := BroadcastTour{}
	this.Id = id
	this.Name = name
	this.Slug = slug
	this.CreatedAt = createdAt
	this.Url = url
	return &this
}

// NewBroadcastTourWithDefaults instantiates a new BroadcastTour object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBroadcastTourWithDefaults() *BroadcastTour {
	this := BroadcastTour{}
	return &this
}

// GetId returns the Id field value
func (o *BroadcastTour) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BroadcastTour) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *BroadcastTour) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *BroadcastTour) SetName(v string) {
	o.Name = v
}

// GetSlug returns the Slug field value
func (o *BroadcastTour) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *BroadcastTour) SetSlug(v string) {
	o.Slug = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *BroadcastTour) GetCreatedAt() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetCreatedAtOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *BroadcastTour) SetCreatedAt(v int32) {
	o.CreatedAt = v
}

// GetDates returns the Dates field value if set, zero value otherwise.
func (o *BroadcastTour) GetDates() []int64 {
	if o == nil || IsNil(o.Dates) {
		var ret []int64
		return ret
	}
	return o.Dates
}

// GetDatesOk returns a tuple with the Dates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetDatesOk() ([]int64, bool) {
	if o == nil || IsNil(o.Dates) {
		return nil, false
	}
	return o.Dates, true
}

// HasDates returns a boolean if a field has been set.
func (o *BroadcastTour) HasDates() bool {
	if o != nil && !IsNil(o.Dates) {
		return true
	}

	return false
}

// SetDates gets a reference to the given []int64 and assigns it to the Dates field.
func (o *BroadcastTour) SetDates(v []int64) {
	o.Dates = v
}

// GetInfo returns the Info field value if set, zero value otherwise.
func (o *BroadcastTour) GetInfo() BroadcastTourInfo {
	if o == nil || IsNil(o.Info) {
		var ret BroadcastTourInfo
		return ret
	}
	return *o.Info
}

// GetInfoOk returns a tuple with the Info field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetInfoOk() (*BroadcastTourInfo, bool) {
	if o == nil || IsNil(o.Info) {
		return nil, false
	}
	return o.Info, true
}

// HasInfo returns a boolean if a field has been set.
func (o *BroadcastTour) HasInfo() bool {
	if o != nil && !IsNil(o.Info) {
		return true
	}

	return false
}

// SetInfo gets a reference to the given BroadcastTourInfo and assigns it to the Info field.
func (o *BroadcastTour) SetInfo(v BroadcastTourInfo) {
	o.Info = &v
}

// GetTier returns the Tier field value if set, zero value otherwise.
func (o *BroadcastTour) GetTier() int32 {
	if o == nil || IsNil(o.Tier) {
		var ret int32
		return ret
	}
	return *o.Tier
}

// GetTierOk returns a tuple with the Tier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetTierOk() (*int32, bool) {
	if o == nil || IsNil(o.Tier) {
		return nil, false
	}
	return o.Tier, true
}

// HasTier returns a boolean if a field has been set.
func (o *BroadcastTour) HasTier() bool {
	if o != nil && !IsNil(o.Tier) {
		return true
	}

	return false
}

// SetTier gets a reference to the given int32 and assigns it to the Tier field.
func (o *BroadcastTour) SetTier(v int32) {
	o.Tier = &v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *BroadcastTour) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *BroadcastTour) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *BroadcastTour) SetImage(v string) {
	o.Image = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *BroadcastTour) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *BroadcastTour) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *BroadcastTour) SetDescription(v string) {
	o.Description = &v
}

// GetTeamTable returns the TeamTable field value if set, zero value otherwise.
func (o *BroadcastTour) GetTeamTable() bool {
	if o == nil || IsNil(o.TeamTable) {
		var ret bool
		return ret
	}
	return *o.TeamTable
}

// GetTeamTableOk returns a tuple with the TeamTable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetTeamTableOk() (*bool, bool) {
	if o == nil || IsNil(o.TeamTable) {
		return nil, false
	}
	return o.TeamTable, true
}

// HasTeamTable returns a boolean if a field has been set.
func (o *BroadcastTour) HasTeamTable() bool {
	if o != nil && !IsNil(o.TeamTable) {
		return true
	}

	return false
}

// SetTeamTable gets a reference to the given bool and assigns it to the TeamTable field.
func (o *BroadcastTour) SetTeamTable(v bool) {
	o.TeamTable = &v
}

// GetShowTeamScores returns the ShowTeamScores field value if set, zero value otherwise.
func (o *BroadcastTour) GetShowTeamScores() bool {
	if o == nil || IsNil(o.ShowTeamScores) {
		var ret bool
		return ret
	}
	return *o.ShowTeamScores
}

// GetShowTeamScoresOk returns a tuple with the ShowTeamScores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetShowTeamScoresOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowTeamScores) {
		return nil, false
	}
	return o.ShowTeamScores, true
}

// HasShowTeamScores returns a boolean if a field has been set.
func (o *BroadcastTour) HasShowTeamScores() bool {
	if o != nil && !IsNil(o.ShowTeamScores) {
		return true
	}

	return false
}

// SetShowTeamScores gets a reference to the given bool and assigns it to the ShowTeamScores field.
func (o *BroadcastTour) SetShowTeamScores(v bool) {
	o.ShowTeamScores = &v
}

// GetUrl returns the Url field value
func (o *BroadcastTour) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *BroadcastTour) SetUrl(v string) {
	o.Url = v
}

// GetCommunityOwner returns the CommunityOwner field value if set, zero value otherwise.
func (o *BroadcastTour) GetCommunityOwner() LightUser {
	if o == nil || IsNil(o.CommunityOwner) {
		var ret LightUser
		return ret
	}
	return *o.CommunityOwner
}

// GetCommunityOwnerOk returns a tuple with the CommunityOwner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BroadcastTour) GetCommunityOwnerOk() (*LightUser, bool) {
	if o == nil || IsNil(o.CommunityOwner) {
		return nil, false
	}
	return o.CommunityOwner, true
}

// HasCommunityOwner returns a boolean if a field has been set.
func (o *BroadcastTour) HasCommunityOwner() bool {
	if o != nil && !IsNil(o.CommunityOwner) {
		return true
	}

	return false
}

// SetCommunityOwner gets a reference to the given LightUser and assigns it to the CommunityOwner field.
func (o *BroadcastTour) SetCommunityOwner(v LightUser) {
	o.CommunityOwner = &v
}

func (o BroadcastTour) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BroadcastTour) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["slug"] = o.Slug
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Dates) {
		toSerialize["dates"] = o.Dates
	}
	if !IsNil(o.Info) {
		toSerialize["info"] = o.Info
	}
	if !IsNil(o.Tier) {
		toSerialize["tier"] = o.Tier
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.TeamTable) {
		toSerialize["teamTable"] = o.TeamTable
	}
	if !IsNil(o.ShowTeamScores) {
		toSerialize["showTeamScores"] = o.ShowTeamScores
	}
	toSerialize["url"] = o.Url
	if !IsNil(o.CommunityOwner) {
		toSerialize["communityOwner"] = o.CommunityOwner
	}
	return toSerialize, nil
}

func (o *BroadcastTour) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
		"slug",
		"createdAt",
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

	varBroadcastTour := _BroadcastTour{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBroadcastTour)

	if err != nil {
		return err
	}

	*o = BroadcastTour(varBroadcastTour)

	return err
}

type NullableBroadcastTour struct {
	value *BroadcastTour
	isSet bool
}

func (v NullableBroadcastTour) Get() *BroadcastTour {
	return v.value
}

func (v *NullableBroadcastTour) Set(val *BroadcastTour) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTour) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTour) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTour(val *BroadcastTour) *NullableBroadcastTour {
	return &NullableBroadcastTour{value: val, isSet: true}
}

func (v NullableBroadcastTour) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTour) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


