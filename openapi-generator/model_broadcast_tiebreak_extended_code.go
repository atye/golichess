/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [Rust general API](https://github.com/obazin/litchee) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api-demo/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api-demo/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api-demo/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.174
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"fmt"
)

// BroadcastTiebreakExtendedCode Extended tiebreak code
type BroadcastTiebreakExtendedCode string

// List of BroadcastTiebreakExtendedCode
const (
	BROADCASTTIEBREAKEXTENDEDCODE_AOB BroadcastTiebreakExtendedCode = "AOB"
	BROADCASTTIEBREAKEXTENDEDCODE_APPO BroadcastTiebreakExtendedCode = "APPO"
	BROADCASTTIEBREAKEXTENDEDCODE_APRO BroadcastTiebreakExtendedCode = "APRO"
	BROADCASTTIEBREAKEXTENDEDCODE_ARO BroadcastTiebreakExtendedCode = "ARO"
	BROADCASTTIEBREAKEXTENDEDCODE_ARO_C1 BroadcastTiebreakExtendedCode = "ARO-C1"
	BROADCASTTIEBREAKEXTENDEDCODE_ARO_C2 BroadcastTiebreakExtendedCode = "ARO-C2"
	BROADCASTTIEBREAKEXTENDEDCODE_ARO_M1 BroadcastTiebreakExtendedCode = "ARO-M1"
	BROADCASTTIEBREAKEXTENDEDCODE_ARO_M2 BroadcastTiebreakExtendedCode = "ARO-M2"
	BROADCASTTIEBREAKEXTENDEDCODE_BH BroadcastTiebreakExtendedCode = "BH"
	BROADCASTTIEBREAKEXTENDEDCODE_BH_C1 BroadcastTiebreakExtendedCode = "BH-C1"
	BROADCASTTIEBREAKEXTENDEDCODE_BH_C2 BroadcastTiebreakExtendedCode = "BH-C2"
	BROADCASTTIEBREAKEXTENDEDCODE_BH_M1 BroadcastTiebreakExtendedCode = "BH-M1"
	BROADCASTTIEBREAKEXTENDEDCODE_BH_M2 BroadcastTiebreakExtendedCode = "BH-M2"
	BROADCASTTIEBREAKEXTENDEDCODE_BPG BroadcastTiebreakExtendedCode = "BPG"
	BROADCASTTIEBREAKEXTENDEDCODE_BWG BroadcastTiebreakExtendedCode = "BWG"
	BROADCASTTIEBREAKEXTENDEDCODE_DE BroadcastTiebreakExtendedCode = "DE"
	BROADCASTTIEBREAKEXTENDEDCODE_FB BroadcastTiebreakExtendedCode = "FB"
	BROADCASTTIEBREAKEXTENDEDCODE_FB_C1 BroadcastTiebreakExtendedCode = "FB-C1"
	BROADCASTTIEBREAKEXTENDEDCODE_FB_C2 BroadcastTiebreakExtendedCode = "FB-C2"
	BROADCASTTIEBREAKEXTENDEDCODE_FB_M1 BroadcastTiebreakExtendedCode = "FB-M1"
	BROADCASTTIEBREAKEXTENDEDCODE_FB_M2 BroadcastTiebreakExtendedCode = "FB-M2"
	BROADCASTTIEBREAKEXTENDEDCODE_KS BroadcastTiebreakExtendedCode = "KS"
	BROADCASTTIEBREAKEXTENDEDCODE_PS BroadcastTiebreakExtendedCode = "PS"
	BROADCASTTIEBREAKEXTENDEDCODE_PS_C1 BroadcastTiebreakExtendedCode = "PS-C1"
	BROADCASTTIEBREAKEXTENDEDCODE_PS_C2 BroadcastTiebreakExtendedCode = "PS-C2"
	BROADCASTTIEBREAKEXTENDEDCODE_PS_M1 BroadcastTiebreakExtendedCode = "PS-M1"
	BROADCASTTIEBREAKEXTENDEDCODE_PS_M2 BroadcastTiebreakExtendedCode = "PS-M2"
	BROADCASTTIEBREAKEXTENDEDCODE_PTP BroadcastTiebreakExtendedCode = "PTP"
	BROADCASTTIEBREAKEXTENDEDCODE_SB BroadcastTiebreakExtendedCode = "SB"
	BROADCASTTIEBREAKEXTENDEDCODE_SB_C1 BroadcastTiebreakExtendedCode = "SB-C1"
	BROADCASTTIEBREAKEXTENDEDCODE_SB_C2 BroadcastTiebreakExtendedCode = "SB-C2"
	BROADCASTTIEBREAKEXTENDEDCODE_SB_M1 BroadcastTiebreakExtendedCode = "SB-M1"
	BROADCASTTIEBREAKEXTENDEDCODE_SB_M2 BroadcastTiebreakExtendedCode = "SB-M2"
	BROADCASTTIEBREAKEXTENDEDCODE_TPR BroadcastTiebreakExtendedCode = "TPR"
	BROADCASTTIEBREAKEXTENDEDCODE_WON BroadcastTiebreakExtendedCode = "WON"
	BROADCASTTIEBREAKEXTENDEDCODE_UNKNOWN_DEFAULT_OPEN_API BroadcastTiebreakExtendedCode = "unknown_default_open_api"
)

// All allowed values of BroadcastTiebreakExtendedCode enum
var AllowedBroadcastTiebreakExtendedCodeEnumValues = []BroadcastTiebreakExtendedCode{
	"AOB",
	"APPO",
	"APRO",
	"ARO",
	"ARO-C1",
	"ARO-C2",
	"ARO-M1",
	"ARO-M2",
	"BH",
	"BH-C1",
	"BH-C2",
	"BH-M1",
	"BH-M2",
	"BPG",
	"BWG",
	"DE",
	"FB",
	"FB-C1",
	"FB-C2",
	"FB-M1",
	"FB-M2",
	"KS",
	"PS",
	"PS-C1",
	"PS-C2",
	"PS-M1",
	"PS-M2",
	"PTP",
	"SB",
	"SB-C1",
	"SB-C2",
	"SB-M1",
	"SB-M2",
	"TPR",
	"WON",
	"unknown_default_open_api",
}

func (v *BroadcastTiebreakExtendedCode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BroadcastTiebreakExtendedCode(value)
	for _, existing := range AllowedBroadcastTiebreakExtendedCodeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	*v = BROADCASTTIEBREAKEXTENDEDCODE_UNKNOWN_DEFAULT_OPEN_API
	return nil
}

// NewBroadcastTiebreakExtendedCodeFromValue returns a pointer to a valid BroadcastTiebreakExtendedCode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBroadcastTiebreakExtendedCodeFromValue(v string) (*BroadcastTiebreakExtendedCode, error) {
	ev := BroadcastTiebreakExtendedCode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BroadcastTiebreakExtendedCode: valid values are %v", v, AllowedBroadcastTiebreakExtendedCodeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BroadcastTiebreakExtendedCode) IsValid() bool {
	for _, existing := range AllowedBroadcastTiebreakExtendedCodeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BroadcastTiebreakExtendedCode value
func (v BroadcastTiebreakExtendedCode) Ptr() *BroadcastTiebreakExtendedCode {
	return &v
}

type NullableBroadcastTiebreakExtendedCode struct {
	value *BroadcastTiebreakExtendedCode
	isSet bool
}

func (v NullableBroadcastTiebreakExtendedCode) Get() *BroadcastTiebreakExtendedCode {
	return v.value
}

func (v *NullableBroadcastTiebreakExtendedCode) Set(val *BroadcastTiebreakExtendedCode) {
	v.value = val
	v.isSet = true
}

func (v NullableBroadcastTiebreakExtendedCode) IsSet() bool {
	return v.isSet
}

func (v *NullableBroadcastTiebreakExtendedCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBroadcastTiebreakExtendedCode(val *BroadcastTiebreakExtendedCode) *NullableBroadcastTiebreakExtendedCode {
	return &NullableBroadcastTiebreakExtendedCode{value: val, isSet: true}
}

func (v NullableBroadcastTiebreakExtendedCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBroadcastTiebreakExtendedCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

