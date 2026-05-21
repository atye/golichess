/*
Lichess.org API reference

# Introduction Welcome to the reference for the Lichess API! Lichess is free/libre, open-source chess server powered by volunteers and donations. - Get help in the [Lichess Discord channel](https://discord.gg/lichess) - API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/) - API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui) - [Contribute to this documentation on Github](https://github.com/lichess-org/api) - Check out [Lichess widgets to embed in your website](https://lichess.org/developers) - [Download all Lichess rated games](https://database.lichess.org/) - [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles) - [Download all evaluated positions](https://database.lichess.org/#evals)  ## Endpoint All requests go to `https://lichess.org` (unless otherwise specified).  ## Clients - [Python general API](https://github.com/lichess-org/berserk) - [MicroPython general API](https://github.com/mkomon/uberserk) - [Python general API - async](https://pypi.org/project/async-lichess-sdk) - [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot) - [Python Board API for Certabo](https://github.com/haklein/certabo-lichess) - [Java general API](https://github.com/tors42/chariot) - [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine) - [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET) - [.NET general API](https://github.com/Dblike/LichessSharp)  ## Rate limiting All requests are rate limited using various strategies, to ensure the API remains responsive for everyone. Only make one request at a time. If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429), you have exceded one of the rate limits. In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer. Reduce your request frequency before retrying.  ## Streaming with ND-JSON Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.  Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.  ## Authentication ### Which authentication method is right for me? [Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)  ### Personal Access Token Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow. - [Generate a personal access token](https://lichess.org/account/oauth/token) - `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"` - [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)  ### Token Security - Keep your tokens secret. Do not share them in public repositories or public forums. - Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account. - Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users. - If you suspect a token has been compromised, revoke it immediately.  To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).  ### Authorization Code Flow with PKCE The authorization code flow with PKCE allows your users to **login with Lichess**. Lichess supports unregistered and public clients (no client authentication, choose any unique client id). The only accepted code challenge method is `S256`. Access tokens are long-lived (expect one year), unless they are revoked. Refresh tokens are not supported.  See the [documentation for the OAuth endpoints](#tag/OAuth) or the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.  - [Demo app](https://lichess-org.github.io/api-demo/) - [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app) - [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask) - [Java example](https://github.com/tors42/lichess-oauth-pkce-app) - [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)  #### Real life examples - [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants)) - [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour)) - [English Chess Federation](https://ecf.octoknight.com/) - [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)  ### Token format Access tokens and authorization codes match `^[A-Za-z0-9_]+$`. The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters. By convention tokens have a recognizable prefix, but do not rely on this. 

API version: 2.0.144
Contact: contact@lichess.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapigenerator

import (
	"encoding/json"
	"fmt"
	"gopkg.in/validator.v2"
)

// BoardGameStream200Response - struct for BoardGameStream200Response
type BoardGameStream200Response struct {
	ChatLineEvent *ChatLineEvent
	GameFullEvent *GameFullEvent
	GameStateEvent *GameStateEvent
	OpponentGoneEvent *OpponentGoneEvent
}

// ChatLineEventAsBoardGameStream200Response is a convenience function that returns ChatLineEvent wrapped in BoardGameStream200Response
func ChatLineEventAsBoardGameStream200Response(v *ChatLineEvent) BoardGameStream200Response {
	return BoardGameStream200Response{
		ChatLineEvent: v,
	}
}

// GameFullEventAsBoardGameStream200Response is a convenience function that returns GameFullEvent wrapped in BoardGameStream200Response
func GameFullEventAsBoardGameStream200Response(v *GameFullEvent) BoardGameStream200Response {
	return BoardGameStream200Response{
		GameFullEvent: v,
	}
}

// GameStateEventAsBoardGameStream200Response is a convenience function that returns GameStateEvent wrapped in BoardGameStream200Response
func GameStateEventAsBoardGameStream200Response(v *GameStateEvent) BoardGameStream200Response {
	return BoardGameStream200Response{
		GameStateEvent: v,
	}
}

// OpponentGoneEventAsBoardGameStream200Response is a convenience function that returns OpponentGoneEvent wrapped in BoardGameStream200Response
func OpponentGoneEventAsBoardGameStream200Response(v *OpponentGoneEvent) BoardGameStream200Response {
	return BoardGameStream200Response{
		OpponentGoneEvent: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *BoardGameStream200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatLineEvent
	err = newStrictDecoder(data).Decode(&dst.ChatLineEvent)
	if err == nil {
		jsonChatLineEvent, _ := json.Marshal(dst.ChatLineEvent)
		if string(jsonChatLineEvent) == "{}" { // empty struct
			dst.ChatLineEvent = nil
		} else {
			if err = validator.Validate(dst.ChatLineEvent); err != nil {
				dst.ChatLineEvent = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatLineEvent = nil
	}

	// try to unmarshal data into GameFullEvent
	err = newStrictDecoder(data).Decode(&dst.GameFullEvent)
	if err == nil {
		jsonGameFullEvent, _ := json.Marshal(dst.GameFullEvent)
		if string(jsonGameFullEvent) == "{}" { // empty struct
			dst.GameFullEvent = nil
		} else {
			if err = validator.Validate(dst.GameFullEvent); err != nil {
				dst.GameFullEvent = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameFullEvent = nil
	}

	// try to unmarshal data into GameStateEvent
	err = newStrictDecoder(data).Decode(&dst.GameStateEvent)
	if err == nil {
		jsonGameStateEvent, _ := json.Marshal(dst.GameStateEvent)
		if string(jsonGameStateEvent) == "{}" { // empty struct
			dst.GameStateEvent = nil
		} else {
			if err = validator.Validate(dst.GameStateEvent); err != nil {
				dst.GameStateEvent = nil
			} else {
				match++
			}
		}
	} else {
		dst.GameStateEvent = nil
	}

	// try to unmarshal data into OpponentGoneEvent
	err = newStrictDecoder(data).Decode(&dst.OpponentGoneEvent)
	if err == nil {
		jsonOpponentGoneEvent, _ := json.Marshal(dst.OpponentGoneEvent)
		if string(jsonOpponentGoneEvent) == "{}" { // empty struct
			dst.OpponentGoneEvent = nil
		} else {
			if err = validator.Validate(dst.OpponentGoneEvent); err != nil {
				dst.OpponentGoneEvent = nil
			} else {
				match++
			}
		}
	} else {
		dst.OpponentGoneEvent = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChatLineEvent = nil
		dst.GameFullEvent = nil
		dst.GameStateEvent = nil
		dst.OpponentGoneEvent = nil

		return fmt.Errorf("data matches more than one schema in oneOf(BoardGameStream200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(BoardGameStream200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src BoardGameStream200Response) MarshalJSON() ([]byte, error) {
	if src.ChatLineEvent != nil {
		return json.Marshal(&src.ChatLineEvent)
	}

	if src.GameFullEvent != nil {
		return json.Marshal(&src.GameFullEvent)
	}

	if src.GameStateEvent != nil {
		return json.Marshal(&src.GameStateEvent)
	}

	if src.OpponentGoneEvent != nil {
		return json.Marshal(&src.OpponentGoneEvent)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *BoardGameStream200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChatLineEvent != nil {
		return obj.ChatLineEvent
	}

	if obj.GameFullEvent != nil {
		return obj.GameFullEvent
	}

	if obj.GameStateEvent != nil {
		return obj.GameStateEvent
	}

	if obj.OpponentGoneEvent != nil {
		return obj.OpponentGoneEvent
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj BoardGameStream200Response) GetActualInstanceValue() (interface{}) {
	if obj.ChatLineEvent != nil {
		return *obj.ChatLineEvent
	}

	if obj.GameFullEvent != nil {
		return *obj.GameFullEvent
	}

	if obj.GameStateEvent != nil {
		return *obj.GameStateEvent
	}

	if obj.OpponentGoneEvent != nil {
		return *obj.OpponentGoneEvent
	}

	// all schemas are nil
	return nil
}

type NullableBoardGameStream200Response struct {
	value *BoardGameStream200Response
	isSet bool
}

func (v NullableBoardGameStream200Response) Get() *BoardGameStream200Response {
	return v.value
}

func (v *NullableBoardGameStream200Response) Set(val *BoardGameStream200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableBoardGameStream200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableBoardGameStream200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBoardGameStream200Response(val *BoardGameStream200Response) *NullableBoardGameStream200Response {
	return &NullableBoardGameStream200Response{value: val, isSet: true}
}

func (v NullableBoardGameStream200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBoardGameStream200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


