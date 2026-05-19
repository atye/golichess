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
	"fmt"
)

// checks if the Timeline200ResponseEntriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Timeline200ResponseEntriesInner{}

// Timeline200ResponseEntriesInner struct for Timeline200ResponseEntriesInner
type Timeline200ResponseEntriesInner struct {
	Timeline200ResponseEntriesInnerAnyOf *Timeline200ResponseEntriesInnerAnyOf
	Timeline200ResponseEntriesInnerAnyOf1 *Timeline200ResponseEntriesInnerAnyOf1
	Timeline200ResponseEntriesInnerAnyOf10 *Timeline200ResponseEntriesInnerAnyOf10
	Timeline200ResponseEntriesInnerAnyOf11 *Timeline200ResponseEntriesInnerAnyOf11
	Timeline200ResponseEntriesInnerAnyOf12 *Timeline200ResponseEntriesInnerAnyOf12
	Timeline200ResponseEntriesInnerAnyOf13 *Timeline200ResponseEntriesInnerAnyOf13
	Timeline200ResponseEntriesInnerAnyOf2 *Timeline200ResponseEntriesInnerAnyOf2
	Timeline200ResponseEntriesInnerAnyOf3 *Timeline200ResponseEntriesInnerAnyOf3
	Timeline200ResponseEntriesInnerAnyOf4 *Timeline200ResponseEntriesInnerAnyOf4
	Timeline200ResponseEntriesInnerAnyOf5 *Timeline200ResponseEntriesInnerAnyOf5
	Timeline200ResponseEntriesInnerAnyOf6 *Timeline200ResponseEntriesInnerAnyOf6
	Timeline200ResponseEntriesInnerAnyOf7 *Timeline200ResponseEntriesInnerAnyOf7
	Timeline200ResponseEntriesInnerAnyOf8 *Timeline200ResponseEntriesInnerAnyOf8
	Timeline200ResponseEntriesInnerAnyOf9 *Timeline200ResponseEntriesInnerAnyOf9
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Timeline200ResponseEntriesInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'blog-post'
	if jsonDict["type"] == "blog-post" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'follow'
	if jsonDict["type"] == "follow" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'forum-post'
	if jsonDict["type"] == "forum-post" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'game-end'
	if jsonDict["type"] == "game-end" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'plan-renew'
	if jsonDict["type"] == "plan-renew" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'plan-start'
	if jsonDict["type"] == "plan-start" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'simul-create'
	if jsonDict["type"] == "simul-create" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'simul-join'
	if jsonDict["type"] == "simul-join" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'stream-start'
	if jsonDict["type"] == "stream-start" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'study-like'
	if jsonDict["type"] == "study-like" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'team-create'
	if jsonDict["type"] == "team-create" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'team-join'
	if jsonDict["type"] == "team-join" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'tour-join'
	if jsonDict["type"] == "tour-join" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'ublog-post'
	if jsonDict["type"] == "ublog-post" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'ublog-post-like'
	if jsonDict["type"] == "ublog-post-like" {
		// try to unmarshal JSON data into ERRORUNKNOWN
		err = json.Unmarshal(data, &dst.ERRORUNKNOWN);
		if err == nil {
			jsonERRORUNKNOWN, _ := json.Marshal(dst.ERRORUNKNOWN)
			if string(jsonERRORUNKNOWN) == "{}" { // empty struct
				dst.ERRORUNKNOWN = nil
			} else {
				return nil // data stored in dst.ERRORUNKNOWN, return on the first match
			}
		} else {
			dst.ERRORUNKNOWN = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_1'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_1" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf1
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf1);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf1, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf1)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf1) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf1 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf1, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf1 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_10'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_10" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf10
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf10);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf10, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf10)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf10) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf10 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf10, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf10 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_11'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_11" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf11
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf11);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf11, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf11)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf11) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf11 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf11, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf11 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_12'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_12" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf12
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf12);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf12, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf12)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf12) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf12 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf12, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf12 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_13'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_13" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf13
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf13);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf13, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf13)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf13) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf13 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf13, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf13 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_2'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_2" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf2
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf2);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf2, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf2)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf2) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf2 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf2, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf2 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_3'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_3" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf3
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf3);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf3, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf3)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf3) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf3 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf3, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf3 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_4'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_4" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf4
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf4);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf4, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf4)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf4) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf4 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf4, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf4 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_5'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_5" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf5
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf5);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf5, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf5)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf5) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf5 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf5, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf5 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_6'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_6" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf6
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf6);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf6, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf6)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf6) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf6 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf6, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf6 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_7'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_7" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf7
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf7);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf7, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf7)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf7) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf7 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf7, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf7 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_8'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_8" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf8
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf8);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf8, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf8)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf8) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf8 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf8, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf8 = nil
		}
	}

	// check if the discriminator value is 'timeline_200_response_entries_inner_anyOf_9'
	if jsonDict["type"] == "timeline_200_response_entries_inner_anyOf_9" {
		// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf9
		err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf9);
		if err == nil {
			jsonTimeline200ResponseEntriesInnerAnyOf9, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf9)
			if string(jsonTimeline200ResponseEntriesInnerAnyOf9) == "{}" { // empty struct
				dst.Timeline200ResponseEntriesInnerAnyOf9 = nil
			} else {
				return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf9, return on the first match
			}
		} else {
			dst.Timeline200ResponseEntriesInnerAnyOf9 = nil
		}
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf1
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf1);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf1, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf1)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf1) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf1 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf1, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf1 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf10
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf10);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf10, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf10)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf10) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf10 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf10, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf10 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf11
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf11);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf11, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf11)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf11) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf11 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf11, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf11 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf12
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf12);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf12, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf12)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf12) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf12 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf12, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf12 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf13
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf13);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf13, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf13)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf13) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf13 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf13, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf13 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf2
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf2);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf2, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf2)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf2) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf2 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf2, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf2 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf3
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf3);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf3, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf3)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf3) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf3 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf3, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf3 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf4
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf4);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf4, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf4)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf4) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf4 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf4, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf4 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf5
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf5);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf5, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf5)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf5) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf5 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf5, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf5 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf6
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf6);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf6, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf6)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf6) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf6 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf6, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf6 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf7
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf7);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf7, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf7)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf7) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf7 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf7, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf7 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf8
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf8);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf8, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf8)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf8) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf8 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf8, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf8 = nil
	}

	// try to unmarshal JSON data into Timeline200ResponseEntriesInnerAnyOf9
	err = json.Unmarshal(data, &dst.Timeline200ResponseEntriesInnerAnyOf9);
	if err == nil {
		jsonTimeline200ResponseEntriesInnerAnyOf9, _ := json.Marshal(dst.Timeline200ResponseEntriesInnerAnyOf9)
		if string(jsonTimeline200ResponseEntriesInnerAnyOf9) == "{}" { // empty struct
			dst.Timeline200ResponseEntriesInnerAnyOf9 = nil
		} else {
			return nil // data stored in dst.Timeline200ResponseEntriesInnerAnyOf9, return on the first match
		}
	} else {
		dst.Timeline200ResponseEntriesInnerAnyOf9 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(Timeline200ResponseEntriesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Timeline200ResponseEntriesInner) MarshalJSON() ([]byte, error) {
	if src.Timeline200ResponseEntriesInnerAnyOf != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf1 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf1)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf10 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf10)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf11 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf11)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf12 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf12)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf13 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf13)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf2 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf2)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf3 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf3)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf4 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf4)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf5 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf5)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf6 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf6)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf7 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf7)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf8 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf8)
	}

	if src.Timeline200ResponseEntriesInnerAnyOf9 != nil {
		return json.Marshal(&src.Timeline200ResponseEntriesInnerAnyOf9)
	}

	return nil, nil // no data in anyOf schemas
}

func (src Timeline200ResponseEntriesInner) ToMap() (map[string]interface{}, error) {
	if src.Timeline200ResponseEntriesInnerAnyOf != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf1 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf1.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf10 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf10.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf11 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf11.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf12 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf12.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf13 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf13.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf2 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf2.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf3 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf3.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf4 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf4.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf5 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf5.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf6 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf6.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf7 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf7.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf8 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf8.ToMap()
	}

	if src.Timeline200ResponseEntriesInnerAnyOf9 != nil {
		return src.Timeline200ResponseEntriesInnerAnyOf9.ToMap()
	}

    return nil, nil // no data in anyOf schemas
}

type NullableTimeline200ResponseEntriesInner struct {
	value *Timeline200ResponseEntriesInner
	isSet bool
}

func (v NullableTimeline200ResponseEntriesInner) Get() *Timeline200ResponseEntriesInner {
	return v.value
}

func (v *NullableTimeline200ResponseEntriesInner) Set(val *Timeline200ResponseEntriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeline200ResponseEntriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeline200ResponseEntriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeline200ResponseEntriesInner(val *Timeline200ResponseEntriesInner) *NullableTimeline200ResponseEntriesInner {
	return &NullableTimeline200ResponseEntriesInner{value: val, isSet: true}
}

func (v NullableTimeline200ResponseEntriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeline200ResponseEntriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


