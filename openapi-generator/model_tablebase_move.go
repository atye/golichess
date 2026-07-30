/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.155
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TablebaseMove type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TablebaseMove{}

// TablebaseMove struct for TablebaseMove
type TablebaseMove struct {
	Uci string `json:"uci"`
	San string `json:"san"`
	Category string `json:"category"`
	Dtz NullableInt32 `json:"dtz,omitempty"`
	PreciseDtz NullableInt32 `json:"precise_dtz,omitempty"`
	Dtc NullableInt32 `json:"dtc,omitempty"`
	Dtm NullableInt32 `json:"dtm,omitempty"`
	Dtw NullableInt32 `json:"dtw,omitempty"`
	Zeroing *bool `json:"zeroing,omitempty"`
	Conversion *bool `json:"conversion,omitempty"`
	Checkmate *bool `json:"checkmate,omitempty"`
	Stalemate *bool `json:"stalemate,omitempty"`
	VariantWin *bool `json:"variant_win,omitempty"`
	VariantLoss *bool `json:"variant_loss,omitempty"`
	InsufficientMaterial *bool `json:"insufficient_material,omitempty"`
}

type _TablebaseMove TablebaseMove

// NewTablebaseMove instantiates a new TablebaseMove object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTablebaseMove(uci string, san string, category string) *TablebaseMove {
	this := TablebaseMove{}
	this.Uci = uci
	this.San = san
	this.Category = category
	return &this
}

// NewTablebaseMoveWithDefaults instantiates a new TablebaseMove object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTablebaseMoveWithDefaults() *TablebaseMove {
	this := TablebaseMove{}
	return &this
}

// GetUci returns the Uci field value
func (o *TablebaseMove) GetUci() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uci
}

// GetUciOk returns a tuple with the Uci field value
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetUciOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uci, true
}

// SetUci sets field value
func (o *TablebaseMove) SetUci(v string) {
	o.Uci = v
}

// GetSan returns the San field value
func (o *TablebaseMove) GetSan() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.San
}

// GetSanOk returns a tuple with the San field value
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetSanOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.San, true
}

// SetSan sets field value
func (o *TablebaseMove) SetSan(v string) {
	o.San = v
}

// GetCategory returns the Category field value
func (o *TablebaseMove) GetCategory() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetCategoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *TablebaseMove) SetCategory(v string) {
	o.Category = v
}

// GetDtz returns the Dtz field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TablebaseMove) GetDtz() int32 {
	if o == nil || IsNil(o.Dtz.Get()) {
		var ret int32
		return ret
	}
	return *o.Dtz.Get()
}

// GetDtzOk returns a tuple with the Dtz field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TablebaseMove) GetDtzOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dtz.Get(), o.Dtz.IsSet()
}

// HasDtz returns a boolean if a field has been set.
func (o *TablebaseMove) HasDtz() bool {
	if o != nil && o.Dtz.IsSet() {
		return true
	}

	return false
}

// SetDtz gets a reference to the given NullableInt32 and assigns it to the Dtz field.
func (o *TablebaseMove) SetDtz(v int32) {
	o.Dtz.Set(&v)
}
// SetDtzNil sets the value for Dtz to be an explicit nil
func (o *TablebaseMove) SetDtzNil() {
	o.Dtz.Set(nil)
}

// UnsetDtz ensures that no value is present for Dtz, not even an explicit nil
func (o *TablebaseMove) UnsetDtz() {
	o.Dtz.Unset()
}

// GetPreciseDtz returns the PreciseDtz field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TablebaseMove) GetPreciseDtz() int32 {
	if o == nil || IsNil(o.PreciseDtz.Get()) {
		var ret int32
		return ret
	}
	return *o.PreciseDtz.Get()
}

// GetPreciseDtzOk returns a tuple with the PreciseDtz field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TablebaseMove) GetPreciseDtzOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PreciseDtz.Get(), o.PreciseDtz.IsSet()
}

// HasPreciseDtz returns a boolean if a field has been set.
func (o *TablebaseMove) HasPreciseDtz() bool {
	if o != nil && o.PreciseDtz.IsSet() {
		return true
	}

	return false
}

// SetPreciseDtz gets a reference to the given NullableInt32 and assigns it to the PreciseDtz field.
func (o *TablebaseMove) SetPreciseDtz(v int32) {
	o.PreciseDtz.Set(&v)
}
// SetPreciseDtzNil sets the value for PreciseDtz to be an explicit nil
func (o *TablebaseMove) SetPreciseDtzNil() {
	o.PreciseDtz.Set(nil)
}

// UnsetPreciseDtz ensures that no value is present for PreciseDtz, not even an explicit nil
func (o *TablebaseMove) UnsetPreciseDtz() {
	o.PreciseDtz.Unset()
}

// GetDtc returns the Dtc field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TablebaseMove) GetDtc() int32 {
	if o == nil || IsNil(o.Dtc.Get()) {
		var ret int32
		return ret
	}
	return *o.Dtc.Get()
}

// GetDtcOk returns a tuple with the Dtc field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TablebaseMove) GetDtcOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dtc.Get(), o.Dtc.IsSet()
}

// HasDtc returns a boolean if a field has been set.
func (o *TablebaseMove) HasDtc() bool {
	if o != nil && o.Dtc.IsSet() {
		return true
	}

	return false
}

// SetDtc gets a reference to the given NullableInt32 and assigns it to the Dtc field.
func (o *TablebaseMove) SetDtc(v int32) {
	o.Dtc.Set(&v)
}
// SetDtcNil sets the value for Dtc to be an explicit nil
func (o *TablebaseMove) SetDtcNil() {
	o.Dtc.Set(nil)
}

// UnsetDtc ensures that no value is present for Dtc, not even an explicit nil
func (o *TablebaseMove) UnsetDtc() {
	o.Dtc.Unset()
}

// GetDtm returns the Dtm field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TablebaseMove) GetDtm() int32 {
	if o == nil || IsNil(o.Dtm.Get()) {
		var ret int32
		return ret
	}
	return *o.Dtm.Get()
}

// GetDtmOk returns a tuple with the Dtm field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TablebaseMove) GetDtmOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dtm.Get(), o.Dtm.IsSet()
}

// HasDtm returns a boolean if a field has been set.
func (o *TablebaseMove) HasDtm() bool {
	if o != nil && o.Dtm.IsSet() {
		return true
	}

	return false
}

// SetDtm gets a reference to the given NullableInt32 and assigns it to the Dtm field.
func (o *TablebaseMove) SetDtm(v int32) {
	o.Dtm.Set(&v)
}
// SetDtmNil sets the value for Dtm to be an explicit nil
func (o *TablebaseMove) SetDtmNil() {
	o.Dtm.Set(nil)
}

// UnsetDtm ensures that no value is present for Dtm, not even an explicit nil
func (o *TablebaseMove) UnsetDtm() {
	o.Dtm.Unset()
}

// GetDtw returns the Dtw field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TablebaseMove) GetDtw() int32 {
	if o == nil || IsNil(o.Dtw.Get()) {
		var ret int32
		return ret
	}
	return *o.Dtw.Get()
}

// GetDtwOk returns a tuple with the Dtw field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TablebaseMove) GetDtwOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Dtw.Get(), o.Dtw.IsSet()
}

// HasDtw returns a boolean if a field has been set.
func (o *TablebaseMove) HasDtw() bool {
	if o != nil && o.Dtw.IsSet() {
		return true
	}

	return false
}

// SetDtw gets a reference to the given NullableInt32 and assigns it to the Dtw field.
func (o *TablebaseMove) SetDtw(v int32) {
	o.Dtw.Set(&v)
}
// SetDtwNil sets the value for Dtw to be an explicit nil
func (o *TablebaseMove) SetDtwNil() {
	o.Dtw.Set(nil)
}

// UnsetDtw ensures that no value is present for Dtw, not even an explicit nil
func (o *TablebaseMove) UnsetDtw() {
	o.Dtw.Unset()
}

// GetZeroing returns the Zeroing field value if set, zero value otherwise.
func (o *TablebaseMove) GetZeroing() bool {
	if o == nil || IsNil(o.Zeroing) {
		var ret bool
		return ret
	}
	return *o.Zeroing
}

// GetZeroingOk returns a tuple with the Zeroing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetZeroingOk() (*bool, bool) {
	if o == nil || IsNil(o.Zeroing) {
		return nil, false
	}
	return o.Zeroing, true
}

// HasZeroing returns a boolean if a field has been set.
func (o *TablebaseMove) HasZeroing() bool {
	if o != nil && !IsNil(o.Zeroing) {
		return true
	}

	return false
}

// SetZeroing gets a reference to the given bool and assigns it to the Zeroing field.
func (o *TablebaseMove) SetZeroing(v bool) {
	o.Zeroing = &v
}

// GetConversion returns the Conversion field value if set, zero value otherwise.
func (o *TablebaseMove) GetConversion() bool {
	if o == nil || IsNil(o.Conversion) {
		var ret bool
		return ret
	}
	return *o.Conversion
}

// GetConversionOk returns a tuple with the Conversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetConversionOk() (*bool, bool) {
	if o == nil || IsNil(o.Conversion) {
		return nil, false
	}
	return o.Conversion, true
}

// HasConversion returns a boolean if a field has been set.
func (o *TablebaseMove) HasConversion() bool {
	if o != nil && !IsNil(o.Conversion) {
		return true
	}

	return false
}

// SetConversion gets a reference to the given bool and assigns it to the Conversion field.
func (o *TablebaseMove) SetConversion(v bool) {
	o.Conversion = &v
}

// GetCheckmate returns the Checkmate field value if set, zero value otherwise.
func (o *TablebaseMove) GetCheckmate() bool {
	if o == nil || IsNil(o.Checkmate) {
		var ret bool
		return ret
	}
	return *o.Checkmate
}

// GetCheckmateOk returns a tuple with the Checkmate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetCheckmateOk() (*bool, bool) {
	if o == nil || IsNil(o.Checkmate) {
		return nil, false
	}
	return o.Checkmate, true
}

// HasCheckmate returns a boolean if a field has been set.
func (o *TablebaseMove) HasCheckmate() bool {
	if o != nil && !IsNil(o.Checkmate) {
		return true
	}

	return false
}

// SetCheckmate gets a reference to the given bool and assigns it to the Checkmate field.
func (o *TablebaseMove) SetCheckmate(v bool) {
	o.Checkmate = &v
}

// GetStalemate returns the Stalemate field value if set, zero value otherwise.
func (o *TablebaseMove) GetStalemate() bool {
	if o == nil || IsNil(o.Stalemate) {
		var ret bool
		return ret
	}
	return *o.Stalemate
}

// GetStalemateOk returns a tuple with the Stalemate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetStalemateOk() (*bool, bool) {
	if o == nil || IsNil(o.Stalemate) {
		return nil, false
	}
	return o.Stalemate, true
}

// HasStalemate returns a boolean if a field has been set.
func (o *TablebaseMove) HasStalemate() bool {
	if o != nil && !IsNil(o.Stalemate) {
		return true
	}

	return false
}

// SetStalemate gets a reference to the given bool and assigns it to the Stalemate field.
func (o *TablebaseMove) SetStalemate(v bool) {
	o.Stalemate = &v
}

// GetVariantWin returns the VariantWin field value if set, zero value otherwise.
func (o *TablebaseMove) GetVariantWin() bool {
	if o == nil || IsNil(o.VariantWin) {
		var ret bool
		return ret
	}
	return *o.VariantWin
}

// GetVariantWinOk returns a tuple with the VariantWin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetVariantWinOk() (*bool, bool) {
	if o == nil || IsNil(o.VariantWin) {
		return nil, false
	}
	return o.VariantWin, true
}

// HasVariantWin returns a boolean if a field has been set.
func (o *TablebaseMove) HasVariantWin() bool {
	if o != nil && !IsNil(o.VariantWin) {
		return true
	}

	return false
}

// SetVariantWin gets a reference to the given bool and assigns it to the VariantWin field.
func (o *TablebaseMove) SetVariantWin(v bool) {
	o.VariantWin = &v
}

// GetVariantLoss returns the VariantLoss field value if set, zero value otherwise.
func (o *TablebaseMove) GetVariantLoss() bool {
	if o == nil || IsNil(o.VariantLoss) {
		var ret bool
		return ret
	}
	return *o.VariantLoss
}

// GetVariantLossOk returns a tuple with the VariantLoss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetVariantLossOk() (*bool, bool) {
	if o == nil || IsNil(o.VariantLoss) {
		return nil, false
	}
	return o.VariantLoss, true
}

// HasVariantLoss returns a boolean if a field has been set.
func (o *TablebaseMove) HasVariantLoss() bool {
	if o != nil && !IsNil(o.VariantLoss) {
		return true
	}

	return false
}

// SetVariantLoss gets a reference to the given bool and assigns it to the VariantLoss field.
func (o *TablebaseMove) SetVariantLoss(v bool) {
	o.VariantLoss = &v
}

// GetInsufficientMaterial returns the InsufficientMaterial field value if set, zero value otherwise.
func (o *TablebaseMove) GetInsufficientMaterial() bool {
	if o == nil || IsNil(o.InsufficientMaterial) {
		var ret bool
		return ret
	}
	return *o.InsufficientMaterial
}

// GetInsufficientMaterialOk returns a tuple with the InsufficientMaterial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TablebaseMove) GetInsufficientMaterialOk() (*bool, bool) {
	if o == nil || IsNil(o.InsufficientMaterial) {
		return nil, false
	}
	return o.InsufficientMaterial, true
}

// HasInsufficientMaterial returns a boolean if a field has been set.
func (o *TablebaseMove) HasInsufficientMaterial() bool {
	if o != nil && !IsNil(o.InsufficientMaterial) {
		return true
	}

	return false
}

// SetInsufficientMaterial gets a reference to the given bool and assigns it to the InsufficientMaterial field.
func (o *TablebaseMove) SetInsufficientMaterial(v bool) {
	o.InsufficientMaterial = &v
}

func (o TablebaseMove) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TablebaseMove) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uci"] = o.Uci
	toSerialize["san"] = o.San
	toSerialize["category"] = o.Category
	if o.Dtz.IsSet() {
		toSerialize["dtz"] = o.Dtz.Get()
	}
	if o.PreciseDtz.IsSet() {
		toSerialize["precise_dtz"] = o.PreciseDtz.Get()
	}
	if o.Dtc.IsSet() {
		toSerialize["dtc"] = o.Dtc.Get()
	}
	if o.Dtm.IsSet() {
		toSerialize["dtm"] = o.Dtm.Get()
	}
	if o.Dtw.IsSet() {
		toSerialize["dtw"] = o.Dtw.Get()
	}
	if !IsNil(o.Zeroing) {
		toSerialize["zeroing"] = o.Zeroing
	}
	if !IsNil(o.Conversion) {
		toSerialize["conversion"] = o.Conversion
	}
	if !IsNil(o.Checkmate) {
		toSerialize["checkmate"] = o.Checkmate
	}
	if !IsNil(o.Stalemate) {
		toSerialize["stalemate"] = o.Stalemate
	}
	if !IsNil(o.VariantWin) {
		toSerialize["variant_win"] = o.VariantWin
	}
	if !IsNil(o.VariantLoss) {
		toSerialize["variant_loss"] = o.VariantLoss
	}
	if !IsNil(o.InsufficientMaterial) {
		toSerialize["insufficient_material"] = o.InsufficientMaterial
	}
	return toSerialize, nil
}

func (o *TablebaseMove) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uci",
		"san",
		"category",
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

	varTablebaseMove := _TablebaseMove{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTablebaseMove)

	if err != nil {
		return err
	}

	*o = TablebaseMove(varTablebaseMove)

	return err
}

type NullableTablebaseMove struct {
	value *TablebaseMove
	isSet bool
}

func (v NullableTablebaseMove) Get() *TablebaseMove {
	return v.value
}

func (v *NullableTablebaseMove) Set(val *TablebaseMove) {
	v.value = val
	v.isSet = true
}

func (v NullableTablebaseMove) IsSet() bool {
	return v.isSet
}

func (v *NullableTablebaseMove) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTablebaseMove(val *TablebaseMove) *NullableTablebaseMove {
	return &NullableTablebaseMove{value: val, isSet: true}
}

func (v NullableTablebaseMove) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTablebaseMove) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


