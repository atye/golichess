/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.168
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the PerfStatStat type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PerfStatStat{}

// PerfStatStat struct for PerfStatStat
type PerfStatStat struct {
	Highest *PerfStatStatHighest `json:"highest,omitempty"`
	Lowest *PerfStatStatHighest `json:"lowest,omitempty"`
	BestWins PerfStatStatBestWins `json:"bestWins"`
	WorstLosses PerfStatStatBestWins `json:"worstLosses"`
	Count PerfStatStatCount `json:"count"`
	ResultStreak PerfStatStatResultStreak `json:"resultStreak"`
	PlayStreak PerfStatStatPlayStreak `json:"playStreak"`
}

type _PerfStatStat PerfStatStat

// NewPerfStatStat instantiates a new PerfStatStat object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPerfStatStat(bestWins PerfStatStatBestWins, worstLosses PerfStatStatBestWins, count PerfStatStatCount, resultStreak PerfStatStatResultStreak, playStreak PerfStatStatPlayStreak) *PerfStatStat {
	this := PerfStatStat{}
	this.BestWins = bestWins
	this.WorstLosses = worstLosses
	this.Count = count
	this.ResultStreak = resultStreak
	this.PlayStreak = playStreak
	return &this
}

// NewPerfStatStatWithDefaults instantiates a new PerfStatStat object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPerfStatStatWithDefaults() *PerfStatStat {
	this := PerfStatStat{}
	return &this
}

// GetHighest returns the Highest field value if set, zero value otherwise.
func (o *PerfStatStat) GetHighest() PerfStatStatHighest {
	if o == nil || IsNil(o.Highest) {
		var ret PerfStatStatHighest
		return ret
	}
	return *o.Highest
}

// GetHighestOk returns a tuple with the Highest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetHighestOk() (*PerfStatStatHighest, bool) {
	if o == nil || IsNil(o.Highest) {
		return nil, false
	}
	return o.Highest, true
}

// HasHighest returns a boolean if a field has been set.
func (o *PerfStatStat) HasHighest() bool {
	if o != nil && !IsNil(o.Highest) {
		return true
	}

	return false
}

// SetHighest gets a reference to the given PerfStatStatHighest and assigns it to the Highest field.
func (o *PerfStatStat) SetHighest(v PerfStatStatHighest) {
	o.Highest = &v
}

// GetLowest returns the Lowest field value if set, zero value otherwise.
func (o *PerfStatStat) GetLowest() PerfStatStatHighest {
	if o == nil || IsNil(o.Lowest) {
		var ret PerfStatStatHighest
		return ret
	}
	return *o.Lowest
}

// GetLowestOk returns a tuple with the Lowest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetLowestOk() (*PerfStatStatHighest, bool) {
	if o == nil || IsNil(o.Lowest) {
		return nil, false
	}
	return o.Lowest, true
}

// HasLowest returns a boolean if a field has been set.
func (o *PerfStatStat) HasLowest() bool {
	if o != nil && !IsNil(o.Lowest) {
		return true
	}

	return false
}

// SetLowest gets a reference to the given PerfStatStatHighest and assigns it to the Lowest field.
func (o *PerfStatStat) SetLowest(v PerfStatStatHighest) {
	o.Lowest = &v
}

// GetBestWins returns the BestWins field value
func (o *PerfStatStat) GetBestWins() PerfStatStatBestWins {
	if o == nil {
		var ret PerfStatStatBestWins
		return ret
	}

	return o.BestWins
}

// GetBestWinsOk returns a tuple with the BestWins field value
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetBestWinsOk() (*PerfStatStatBestWins, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BestWins, true
}

// SetBestWins sets field value
func (o *PerfStatStat) SetBestWins(v PerfStatStatBestWins) {
	o.BestWins = v
}

// GetWorstLosses returns the WorstLosses field value
func (o *PerfStatStat) GetWorstLosses() PerfStatStatBestWins {
	if o == nil {
		var ret PerfStatStatBestWins
		return ret
	}

	return o.WorstLosses
}

// GetWorstLossesOk returns a tuple with the WorstLosses field value
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetWorstLossesOk() (*PerfStatStatBestWins, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorstLosses, true
}

// SetWorstLosses sets field value
func (o *PerfStatStat) SetWorstLosses(v PerfStatStatBestWins) {
	o.WorstLosses = v
}

// GetCount returns the Count field value
func (o *PerfStatStat) GetCount() PerfStatStatCount {
	if o == nil {
		var ret PerfStatStatCount
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetCountOk() (*PerfStatStatCount, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *PerfStatStat) SetCount(v PerfStatStatCount) {
	o.Count = v
}

// GetResultStreak returns the ResultStreak field value
func (o *PerfStatStat) GetResultStreak() PerfStatStatResultStreak {
	if o == nil {
		var ret PerfStatStatResultStreak
		return ret
	}

	return o.ResultStreak
}

// GetResultStreakOk returns a tuple with the ResultStreak field value
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetResultStreakOk() (*PerfStatStatResultStreak, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultStreak, true
}

// SetResultStreak sets field value
func (o *PerfStatStat) SetResultStreak(v PerfStatStatResultStreak) {
	o.ResultStreak = v
}

// GetPlayStreak returns the PlayStreak field value
func (o *PerfStatStat) GetPlayStreak() PerfStatStatPlayStreak {
	if o == nil {
		var ret PerfStatStatPlayStreak
		return ret
	}

	return o.PlayStreak
}

// GetPlayStreakOk returns a tuple with the PlayStreak field value
// and a boolean to check if the value has been set.
func (o *PerfStatStat) GetPlayStreakOk() (*PerfStatStatPlayStreak, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayStreak, true
}

// SetPlayStreak sets field value
func (o *PerfStatStat) SetPlayStreak(v PerfStatStatPlayStreak) {
	o.PlayStreak = v
}

func (o PerfStatStat) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PerfStatStat) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Highest) {
		toSerialize["highest"] = o.Highest
	}
	if !IsNil(o.Lowest) {
		toSerialize["lowest"] = o.Lowest
	}
	toSerialize["bestWins"] = o.BestWins
	toSerialize["worstLosses"] = o.WorstLosses
	toSerialize["count"] = o.Count
	toSerialize["resultStreak"] = o.ResultStreak
	toSerialize["playStreak"] = o.PlayStreak
	return toSerialize, nil
}

func (o *PerfStatStat) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"bestWins",
		"worstLosses",
		"count",
		"resultStreak",
		"playStreak",
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

	varPerfStatStat := _PerfStatStat{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPerfStatStat)

	if err != nil {
		return err
	}

	*o = PerfStatStat(varPerfStatStat)

	return err
}

type NullablePerfStatStat struct {
	value *PerfStatStat
	isSet bool
}

func (v NullablePerfStatStat) Get() *PerfStatStat {
	return v.value
}

func (v *NullablePerfStatStat) Set(val *PerfStatStat) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfStatStat) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfStatStat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfStatStat(val *PerfStatStat) *NullablePerfStatStat {
	return &NullablePerfStatStat{value: val, isSet: true}
}

func (v NullablePerfStatStat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfStatStat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


