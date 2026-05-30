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
	"fmt"
)

// checks if the TimelineEntriesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimelineEntriesInner{}

// TimelineEntriesInner struct for TimelineEntriesInner
type TimelineEntriesInner struct {
	TimelineEntryBlogPost *TimelineEntryBlogPost
	TimelineEntryFollow *TimelineEntryFollow
	TimelineEntryForumPost *TimelineEntryForumPost
	TimelineEntryGameEnd *TimelineEntryGameEnd
	TimelineEntryPlanRenew *TimelineEntryPlanRenew
	TimelineEntryPlanStart *TimelineEntryPlanStart
	TimelineEntrySimul *TimelineEntrySimul
	TimelineEntryStreamStart *TimelineEntryStreamStart
	TimelineEntryStudyLike *TimelineEntryStudyLike
	TimelineEntryTeamCreate *TimelineEntryTeamCreate
	TimelineEntryTeamJoin *TimelineEntryTeamJoin
	TimelineEntryTourJoin *TimelineEntryTourJoin
	TimelineEntryUblogPost *TimelineEntryUblogPost
	TimelineEntryUblogPostLike *TimelineEntryUblogPostLike
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TimelineEntriesInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON into map for the discriminator lookup")
	}

	// check if the discriminator value is 'blog-post'
	if jsonDict["type"] == "blog-post" {
		// try to unmarshal JSON data into TimelineEntryBlogPost
		err = json.Unmarshal(data, &dst.TimelineEntryBlogPost);
		if err == nil {
			jsonTimelineEntryBlogPost, _ := json.Marshal(dst.TimelineEntryBlogPost)
			if string(jsonTimelineEntryBlogPost) == "{}" { // empty struct
				dst.TimelineEntryBlogPost = nil
			} else {
				return nil // data stored in dst.TimelineEntryBlogPost, return on the first match
			}
		} else {
			dst.TimelineEntryBlogPost = nil
		}
	}

	// check if the discriminator value is 'follow'
	if jsonDict["type"] == "follow" {
		// try to unmarshal JSON data into TimelineEntryFollow
		err = json.Unmarshal(data, &dst.TimelineEntryFollow);
		if err == nil {
			jsonTimelineEntryFollow, _ := json.Marshal(dst.TimelineEntryFollow)
			if string(jsonTimelineEntryFollow) == "{}" { // empty struct
				dst.TimelineEntryFollow = nil
			} else {
				return nil // data stored in dst.TimelineEntryFollow, return on the first match
			}
		} else {
			dst.TimelineEntryFollow = nil
		}
	}

	// check if the discriminator value is 'forum-post'
	if jsonDict["type"] == "forum-post" {
		// try to unmarshal JSON data into TimelineEntryForumPost
		err = json.Unmarshal(data, &dst.TimelineEntryForumPost);
		if err == nil {
			jsonTimelineEntryForumPost, _ := json.Marshal(dst.TimelineEntryForumPost)
			if string(jsonTimelineEntryForumPost) == "{}" { // empty struct
				dst.TimelineEntryForumPost = nil
			} else {
				return nil // data stored in dst.TimelineEntryForumPost, return on the first match
			}
		} else {
			dst.TimelineEntryForumPost = nil
		}
	}

	// check if the discriminator value is 'game-end'
	if jsonDict["type"] == "game-end" {
		// try to unmarshal JSON data into TimelineEntryGameEnd
		err = json.Unmarshal(data, &dst.TimelineEntryGameEnd);
		if err == nil {
			jsonTimelineEntryGameEnd, _ := json.Marshal(dst.TimelineEntryGameEnd)
			if string(jsonTimelineEntryGameEnd) == "{}" { // empty struct
				dst.TimelineEntryGameEnd = nil
			} else {
				return nil // data stored in dst.TimelineEntryGameEnd, return on the first match
			}
		} else {
			dst.TimelineEntryGameEnd = nil
		}
	}

	// check if the discriminator value is 'plan-renew'
	if jsonDict["type"] == "plan-renew" {
		// try to unmarshal JSON data into TimelineEntryPlanRenew
		err = json.Unmarshal(data, &dst.TimelineEntryPlanRenew);
		if err == nil {
			jsonTimelineEntryPlanRenew, _ := json.Marshal(dst.TimelineEntryPlanRenew)
			if string(jsonTimelineEntryPlanRenew) == "{}" { // empty struct
				dst.TimelineEntryPlanRenew = nil
			} else {
				return nil // data stored in dst.TimelineEntryPlanRenew, return on the first match
			}
		} else {
			dst.TimelineEntryPlanRenew = nil
		}
	}

	// check if the discriminator value is 'plan-start'
	if jsonDict["type"] == "plan-start" {
		// try to unmarshal JSON data into TimelineEntryPlanStart
		err = json.Unmarshal(data, &dst.TimelineEntryPlanStart);
		if err == nil {
			jsonTimelineEntryPlanStart, _ := json.Marshal(dst.TimelineEntryPlanStart)
			if string(jsonTimelineEntryPlanStart) == "{}" { // empty struct
				dst.TimelineEntryPlanStart = nil
			} else {
				return nil // data stored in dst.TimelineEntryPlanStart, return on the first match
			}
		} else {
			dst.TimelineEntryPlanStart = nil
		}
	}

	// check if the discriminator value is 'simul-create'
	if jsonDict["type"] == "simul-create" {
		// try to unmarshal JSON data into TimelineEntrySimul
		err = json.Unmarshal(data, &dst.TimelineEntrySimul);
		if err == nil {
			jsonTimelineEntrySimul, _ := json.Marshal(dst.TimelineEntrySimul)
			if string(jsonTimelineEntrySimul) == "{}" { // empty struct
				dst.TimelineEntrySimul = nil
			} else {
				return nil // data stored in dst.TimelineEntrySimul, return on the first match
			}
		} else {
			dst.TimelineEntrySimul = nil
		}
	}

	// check if the discriminator value is 'simul-join'
	if jsonDict["type"] == "simul-join" {
		// try to unmarshal JSON data into TimelineEntrySimul
		err = json.Unmarshal(data, &dst.TimelineEntrySimul);
		if err == nil {
			jsonTimelineEntrySimul, _ := json.Marshal(dst.TimelineEntrySimul)
			if string(jsonTimelineEntrySimul) == "{}" { // empty struct
				dst.TimelineEntrySimul = nil
			} else {
				return nil // data stored in dst.TimelineEntrySimul, return on the first match
			}
		} else {
			dst.TimelineEntrySimul = nil
		}
	}

	// check if the discriminator value is 'stream-start'
	if jsonDict["type"] == "stream-start" {
		// try to unmarshal JSON data into TimelineEntryStreamStart
		err = json.Unmarshal(data, &dst.TimelineEntryStreamStart);
		if err == nil {
			jsonTimelineEntryStreamStart, _ := json.Marshal(dst.TimelineEntryStreamStart)
			if string(jsonTimelineEntryStreamStart) == "{}" { // empty struct
				dst.TimelineEntryStreamStart = nil
			} else {
				return nil // data stored in dst.TimelineEntryStreamStart, return on the first match
			}
		} else {
			dst.TimelineEntryStreamStart = nil
		}
	}

	// check if the discriminator value is 'study-like'
	if jsonDict["type"] == "study-like" {
		// try to unmarshal JSON data into TimelineEntryStudyLike
		err = json.Unmarshal(data, &dst.TimelineEntryStudyLike);
		if err == nil {
			jsonTimelineEntryStudyLike, _ := json.Marshal(dst.TimelineEntryStudyLike)
			if string(jsonTimelineEntryStudyLike) == "{}" { // empty struct
				dst.TimelineEntryStudyLike = nil
			} else {
				return nil // data stored in dst.TimelineEntryStudyLike, return on the first match
			}
		} else {
			dst.TimelineEntryStudyLike = nil
		}
	}

	// check if the discriminator value is 'team-create'
	if jsonDict["type"] == "team-create" {
		// try to unmarshal JSON data into TimelineEntryTeamCreate
		err = json.Unmarshal(data, &dst.TimelineEntryTeamCreate);
		if err == nil {
			jsonTimelineEntryTeamCreate, _ := json.Marshal(dst.TimelineEntryTeamCreate)
			if string(jsonTimelineEntryTeamCreate) == "{}" { // empty struct
				dst.TimelineEntryTeamCreate = nil
			} else {
				return nil // data stored in dst.TimelineEntryTeamCreate, return on the first match
			}
		} else {
			dst.TimelineEntryTeamCreate = nil
		}
	}

	// check if the discriminator value is 'team-join'
	if jsonDict["type"] == "team-join" {
		// try to unmarshal JSON data into TimelineEntryTeamJoin
		err = json.Unmarshal(data, &dst.TimelineEntryTeamJoin);
		if err == nil {
			jsonTimelineEntryTeamJoin, _ := json.Marshal(dst.TimelineEntryTeamJoin)
			if string(jsonTimelineEntryTeamJoin) == "{}" { // empty struct
				dst.TimelineEntryTeamJoin = nil
			} else {
				return nil // data stored in dst.TimelineEntryTeamJoin, return on the first match
			}
		} else {
			dst.TimelineEntryTeamJoin = nil
		}
	}

	// check if the discriminator value is 'tour-join'
	if jsonDict["type"] == "tour-join" {
		// try to unmarshal JSON data into TimelineEntryTourJoin
		err = json.Unmarshal(data, &dst.TimelineEntryTourJoin);
		if err == nil {
			jsonTimelineEntryTourJoin, _ := json.Marshal(dst.TimelineEntryTourJoin)
			if string(jsonTimelineEntryTourJoin) == "{}" { // empty struct
				dst.TimelineEntryTourJoin = nil
			} else {
				return nil // data stored in dst.TimelineEntryTourJoin, return on the first match
			}
		} else {
			dst.TimelineEntryTourJoin = nil
		}
	}

	// check if the discriminator value is 'ublog-post'
	if jsonDict["type"] == "ublog-post" {
		// try to unmarshal JSON data into TimelineEntryUblogPost
		err = json.Unmarshal(data, &dst.TimelineEntryUblogPost);
		if err == nil {
			jsonTimelineEntryUblogPost, _ := json.Marshal(dst.TimelineEntryUblogPost)
			if string(jsonTimelineEntryUblogPost) == "{}" { // empty struct
				dst.TimelineEntryUblogPost = nil
			} else {
				return nil // data stored in dst.TimelineEntryUblogPost, return on the first match
			}
		} else {
			dst.TimelineEntryUblogPost = nil
		}
	}

	// check if the discriminator value is 'ublog-post-like'
	if jsonDict["type"] == "ublog-post-like" {
		// try to unmarshal JSON data into TimelineEntryUblogPostLike
		err = json.Unmarshal(data, &dst.TimelineEntryUblogPostLike);
		if err == nil {
			jsonTimelineEntryUblogPostLike, _ := json.Marshal(dst.TimelineEntryUblogPostLike)
			if string(jsonTimelineEntryUblogPostLike) == "{}" { // empty struct
				dst.TimelineEntryUblogPostLike = nil
			} else {
				return nil // data stored in dst.TimelineEntryUblogPostLike, return on the first match
			}
		} else {
			dst.TimelineEntryUblogPostLike = nil
		}
	}

	// try to unmarshal JSON data into TimelineEntryBlogPost
	err = json.Unmarshal(data, &dst.TimelineEntryBlogPost);
	if err == nil {
		jsonTimelineEntryBlogPost, _ := json.Marshal(dst.TimelineEntryBlogPost)
		if string(jsonTimelineEntryBlogPost) == "{}" { // empty struct
			dst.TimelineEntryBlogPost = nil
		} else {
			return nil // data stored in dst.TimelineEntryBlogPost, return on the first match
		}
	} else {
		dst.TimelineEntryBlogPost = nil
	}

	// try to unmarshal JSON data into TimelineEntryFollow
	err = json.Unmarshal(data, &dst.TimelineEntryFollow);
	if err == nil {
		jsonTimelineEntryFollow, _ := json.Marshal(dst.TimelineEntryFollow)
		if string(jsonTimelineEntryFollow) == "{}" { // empty struct
			dst.TimelineEntryFollow = nil
		} else {
			return nil // data stored in dst.TimelineEntryFollow, return on the first match
		}
	} else {
		dst.TimelineEntryFollow = nil
	}

	// try to unmarshal JSON data into TimelineEntryForumPost
	err = json.Unmarshal(data, &dst.TimelineEntryForumPost);
	if err == nil {
		jsonTimelineEntryForumPost, _ := json.Marshal(dst.TimelineEntryForumPost)
		if string(jsonTimelineEntryForumPost) == "{}" { // empty struct
			dst.TimelineEntryForumPost = nil
		} else {
			return nil // data stored in dst.TimelineEntryForumPost, return on the first match
		}
	} else {
		dst.TimelineEntryForumPost = nil
	}

	// try to unmarshal JSON data into TimelineEntryGameEnd
	err = json.Unmarshal(data, &dst.TimelineEntryGameEnd);
	if err == nil {
		jsonTimelineEntryGameEnd, _ := json.Marshal(dst.TimelineEntryGameEnd)
		if string(jsonTimelineEntryGameEnd) == "{}" { // empty struct
			dst.TimelineEntryGameEnd = nil
		} else {
			return nil // data stored in dst.TimelineEntryGameEnd, return on the first match
		}
	} else {
		dst.TimelineEntryGameEnd = nil
	}

	// try to unmarshal JSON data into TimelineEntryPlanRenew
	err = json.Unmarshal(data, &dst.TimelineEntryPlanRenew);
	if err == nil {
		jsonTimelineEntryPlanRenew, _ := json.Marshal(dst.TimelineEntryPlanRenew)
		if string(jsonTimelineEntryPlanRenew) == "{}" { // empty struct
			dst.TimelineEntryPlanRenew = nil
		} else {
			return nil // data stored in dst.TimelineEntryPlanRenew, return on the first match
		}
	} else {
		dst.TimelineEntryPlanRenew = nil
	}

	// try to unmarshal JSON data into TimelineEntryPlanStart
	err = json.Unmarshal(data, &dst.TimelineEntryPlanStart);
	if err == nil {
		jsonTimelineEntryPlanStart, _ := json.Marshal(dst.TimelineEntryPlanStart)
		if string(jsonTimelineEntryPlanStart) == "{}" { // empty struct
			dst.TimelineEntryPlanStart = nil
		} else {
			return nil // data stored in dst.TimelineEntryPlanStart, return on the first match
		}
	} else {
		dst.TimelineEntryPlanStart = nil
	}

	// try to unmarshal JSON data into TimelineEntrySimul
	err = json.Unmarshal(data, &dst.TimelineEntrySimul);
	if err == nil {
		jsonTimelineEntrySimul, _ := json.Marshal(dst.TimelineEntrySimul)
		if string(jsonTimelineEntrySimul) == "{}" { // empty struct
			dst.TimelineEntrySimul = nil
		} else {
			return nil // data stored in dst.TimelineEntrySimul, return on the first match
		}
	} else {
		dst.TimelineEntrySimul = nil
	}

	// try to unmarshal JSON data into TimelineEntryStreamStart
	err = json.Unmarshal(data, &dst.TimelineEntryStreamStart);
	if err == nil {
		jsonTimelineEntryStreamStart, _ := json.Marshal(dst.TimelineEntryStreamStart)
		if string(jsonTimelineEntryStreamStart) == "{}" { // empty struct
			dst.TimelineEntryStreamStart = nil
		} else {
			return nil // data stored in dst.TimelineEntryStreamStart, return on the first match
		}
	} else {
		dst.TimelineEntryStreamStart = nil
	}

	// try to unmarshal JSON data into TimelineEntryStudyLike
	err = json.Unmarshal(data, &dst.TimelineEntryStudyLike);
	if err == nil {
		jsonTimelineEntryStudyLike, _ := json.Marshal(dst.TimelineEntryStudyLike)
		if string(jsonTimelineEntryStudyLike) == "{}" { // empty struct
			dst.TimelineEntryStudyLike = nil
		} else {
			return nil // data stored in dst.TimelineEntryStudyLike, return on the first match
		}
	} else {
		dst.TimelineEntryStudyLike = nil
	}

	// try to unmarshal JSON data into TimelineEntryTeamCreate
	err = json.Unmarshal(data, &dst.TimelineEntryTeamCreate);
	if err == nil {
		jsonTimelineEntryTeamCreate, _ := json.Marshal(dst.TimelineEntryTeamCreate)
		if string(jsonTimelineEntryTeamCreate) == "{}" { // empty struct
			dst.TimelineEntryTeamCreate = nil
		} else {
			return nil // data stored in dst.TimelineEntryTeamCreate, return on the first match
		}
	} else {
		dst.TimelineEntryTeamCreate = nil
	}

	// try to unmarshal JSON data into TimelineEntryTeamJoin
	err = json.Unmarshal(data, &dst.TimelineEntryTeamJoin);
	if err == nil {
		jsonTimelineEntryTeamJoin, _ := json.Marshal(dst.TimelineEntryTeamJoin)
		if string(jsonTimelineEntryTeamJoin) == "{}" { // empty struct
			dst.TimelineEntryTeamJoin = nil
		} else {
			return nil // data stored in dst.TimelineEntryTeamJoin, return on the first match
		}
	} else {
		dst.TimelineEntryTeamJoin = nil
	}

	// try to unmarshal JSON data into TimelineEntryTourJoin
	err = json.Unmarshal(data, &dst.TimelineEntryTourJoin);
	if err == nil {
		jsonTimelineEntryTourJoin, _ := json.Marshal(dst.TimelineEntryTourJoin)
		if string(jsonTimelineEntryTourJoin) == "{}" { // empty struct
			dst.TimelineEntryTourJoin = nil
		} else {
			return nil // data stored in dst.TimelineEntryTourJoin, return on the first match
		}
	} else {
		dst.TimelineEntryTourJoin = nil
	}

	// try to unmarshal JSON data into TimelineEntryUblogPost
	err = json.Unmarshal(data, &dst.TimelineEntryUblogPost);
	if err == nil {
		jsonTimelineEntryUblogPost, _ := json.Marshal(dst.TimelineEntryUblogPost)
		if string(jsonTimelineEntryUblogPost) == "{}" { // empty struct
			dst.TimelineEntryUblogPost = nil
		} else {
			return nil // data stored in dst.TimelineEntryUblogPost, return on the first match
		}
	} else {
		dst.TimelineEntryUblogPost = nil
	}

	// try to unmarshal JSON data into TimelineEntryUblogPostLike
	err = json.Unmarshal(data, &dst.TimelineEntryUblogPostLike);
	if err == nil {
		jsonTimelineEntryUblogPostLike, _ := json.Marshal(dst.TimelineEntryUblogPostLike)
		if string(jsonTimelineEntryUblogPostLike) == "{}" { // empty struct
			dst.TimelineEntryUblogPostLike = nil
		} else {
			return nil // data stored in dst.TimelineEntryUblogPostLike, return on the first match
		}
	} else {
		dst.TimelineEntryUblogPostLike = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TimelineEntriesInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TimelineEntriesInner) MarshalJSON() ([]byte, error) {
	if src.TimelineEntryBlogPost != nil {
		return json.Marshal(&src.TimelineEntryBlogPost)
	}

	if src.TimelineEntryFollow != nil {
		return json.Marshal(&src.TimelineEntryFollow)
	}

	if src.TimelineEntryForumPost != nil {
		return json.Marshal(&src.TimelineEntryForumPost)
	}

	if src.TimelineEntryGameEnd != nil {
		return json.Marshal(&src.TimelineEntryGameEnd)
	}

	if src.TimelineEntryPlanRenew != nil {
		return json.Marshal(&src.TimelineEntryPlanRenew)
	}

	if src.TimelineEntryPlanStart != nil {
		return json.Marshal(&src.TimelineEntryPlanStart)
	}

	if src.TimelineEntrySimul != nil {
		return json.Marshal(&src.TimelineEntrySimul)
	}

	if src.TimelineEntryStreamStart != nil {
		return json.Marshal(&src.TimelineEntryStreamStart)
	}

	if src.TimelineEntryStudyLike != nil {
		return json.Marshal(&src.TimelineEntryStudyLike)
	}

	if src.TimelineEntryTeamCreate != nil {
		return json.Marshal(&src.TimelineEntryTeamCreate)
	}

	if src.TimelineEntryTeamJoin != nil {
		return json.Marshal(&src.TimelineEntryTeamJoin)
	}

	if src.TimelineEntryTourJoin != nil {
		return json.Marshal(&src.TimelineEntryTourJoin)
	}

	if src.TimelineEntryUblogPost != nil {
		return json.Marshal(&src.TimelineEntryUblogPost)
	}

	if src.TimelineEntryUblogPostLike != nil {
		return json.Marshal(&src.TimelineEntryUblogPostLike)
	}

	return nil, nil // no data in anyOf schemas
}

func (src TimelineEntriesInner) ToMap() (map[string]interface{}, error) {
	if src.TimelineEntryBlogPost != nil {
		return src.TimelineEntryBlogPost.ToMap()
	}

	if src.TimelineEntryFollow != nil {
		return src.TimelineEntryFollow.ToMap()
	}

	if src.TimelineEntryForumPost != nil {
		return src.TimelineEntryForumPost.ToMap()
	}

	if src.TimelineEntryGameEnd != nil {
		return src.TimelineEntryGameEnd.ToMap()
	}

	if src.TimelineEntryPlanRenew != nil {
		return src.TimelineEntryPlanRenew.ToMap()
	}

	if src.TimelineEntryPlanStart != nil {
		return src.TimelineEntryPlanStart.ToMap()
	}

	if src.TimelineEntrySimul != nil {
		return src.TimelineEntrySimul.ToMap()
	}

	if src.TimelineEntryStreamStart != nil {
		return src.TimelineEntryStreamStart.ToMap()
	}

	if src.TimelineEntryStudyLike != nil {
		return src.TimelineEntryStudyLike.ToMap()
	}

	if src.TimelineEntryTeamCreate != nil {
		return src.TimelineEntryTeamCreate.ToMap()
	}

	if src.TimelineEntryTeamJoin != nil {
		return src.TimelineEntryTeamJoin.ToMap()
	}

	if src.TimelineEntryTourJoin != nil {
		return src.TimelineEntryTourJoin.ToMap()
	}

	if src.TimelineEntryUblogPost != nil {
		return src.TimelineEntryUblogPost.ToMap()
	}

	if src.TimelineEntryUblogPostLike != nil {
		return src.TimelineEntryUblogPostLike.ToMap()
	}

    return nil, nil // no data in anyOf schemas
}

type NullableTimelineEntriesInner struct {
	value *TimelineEntriesInner
	isSet bool
}

func (v NullableTimelineEntriesInner) Get() *TimelineEntriesInner {
	return v.value
}

func (v *NullableTimelineEntriesInner) Set(val *TimelineEntriesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTimelineEntriesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTimelineEntriesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimelineEntriesInner(val *TimelineEntriesInner) *NullableTimelineEntriesInner {
	return &NullableTimelineEntriesInner{value: val, isSet: true}
}

func (v NullableTimelineEntriesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimelineEntriesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


