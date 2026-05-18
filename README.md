# Go API client for lichess

# Introduction
Welcome to the reference for the Lichess API! Lichess is free/libre,
open-source chess server powered by volunteers and donations.
- Get help in the [Lichess Discord channel](https://discord.gg/lichess)
- API demo app with OAuth2 login and gameplay: [source](https://github.com/lichess-org/api-demo) / [demo](https://lichess-org.github.io/api-demo/)
- API UI app with OAuth2 login and endpoint forms: [source](https://github.com/lichess-org/api-ui) / [website](https://lichess.org/api/ui)
- [Contribute to this documentation on Github](https://github.com/lichess-org/api)
- Check out [Lichess widgets to embed in your website](https://lichess.org/developers)
- [Download all Lichess rated games](https://database.lichess.org/)
- [Download all Lichess puzzles with themes, ratings and votes](https://database.lichess.org/#puzzles)
- [Download all evaluated positions](https://database.lichess.org/#evals)

## Endpoint
All requests go to `https://lichess.org` (unless otherwise specified).

## Clients
- [Python general API](https://github.com/lichess-org/berserk)
- [MicroPython general API](https://github.com/mkomon/uberserk)
- [Python general API - async](https://pypi.org/project/async-lichess-sdk)
- [Python Lichess Bot](https://github.com/lichess-bot-devs/lichess-bot)
- [Python Board API for Certabo](https://github.com/haklein/certabo-lichess)
- [Java general API](https://github.com/tors42/chariot)
- [JavaScript & TypeScript general API](https://github.com/devjiwonchoi/equine)
- [LichessNET - C# API Wrapper](https://github.com/Rabergsel/LichessNET)
- [.NET general API](https://github.com/Dblike/LichessSharp)

## Rate limiting
All requests are rate limited using various strategies,
to ensure the API remains responsive for everyone.
Only make one request at a time.
If you receive an HTTP response with a [429 status](https://en.wikipedia.org/wiki/List_of_HTTP_status_codes#429),
you have exceded one of the rate limits.
In most cases, waiting one minute before retrying will be sufficient, but some limits may require longer.
Reduce your request frequency before retrying.

## Streaming with ND-JSON
Some API endpoints stream their responses as [Newline Delimited JSON a.k.a. **nd-json**](https://github.com/ndjson/ndjson-spec), with one JSON object per line.

Here's a [JavaScript utility function](https://gist.github.com/ornicar/a097406810939cf7be1df8ea30e94f3e) to help reading NDJSON streamed responses.

## Authentication
### Which authentication method is right for me?
[Read about the Lichess API authentication methods and code examples](https://github.com/lichess-org/api/blob/master/example/README.md)

### Personal Access Token
Personal API access tokens allow you to quickly interact with Lichess API without going through an OAuth flow.
- [Generate a personal access token](https://lichess.org/account/oauth/token)
- `curl https://lichess.org/api/account -H \"Authorization: Bearer {token}\"`
- [NodeJS example](https://github.com/lichess-org/api/tree/master/example/oauth-personal-token)

### Token Security
- Keep your tokens secret. Do not share them in public repositories or public forums.
- Your tokens can be used to make your account perform arbitrary actions (within the limits of the tokens' scope). You remain responsible for all activities on your account.
- Do not hardcode tokens in your application's code. Use environment variables or a secure storage and ensure they are not shipped/exposed to users. Be especially careful that they are not included in frontend bundles or apps that are shipped to users.
- If you suspect a token has been compromised, revoke it immediately.

To see your active tokens or revoke them, see [your Personal API access tokens](https://lichess.org/account/oauth/token).

### Authorization Code Flow with PKCE
The authorization code flow with PKCE allows your users to **login with Lichess**.
Lichess supports unregistered and public clients (no client authentication, choose any unique client id).
The only accepted code challenge method is `S256`.
Access tokens are long-lived (expect one year), unless they are revoked.
Refresh tokens are not supported.

See the [documentation for the OAuth endpoints](#tag/OAuth) or
the [PKCE RFC](https://datatracker.ietf.org/doc/html/rfc7636#section-4) for a precise protocol description.

- [Demo app](https://lichess-org.github.io/api-demo/)
- [Minimal client-side example](https://github.com/lichess-org/api/tree/master/example/oauth-app)
- [Flask/Python example](https://github.com/lakinwecker/lichess-oauth-flask)
- [Java example](https://github.com/tors42/lichess-oauth-pkce-app)
- [NodeJS Passport strategy to login with Lichess OAuth2](https://www.npmjs.com/package/passport-lichess)

#### Real life examples
- [PyChess](https://github.com/gbtami/pychess-variants) ([source code](https://github.com/gbtami/pychess-variants))
- [Lichess4545](https://www.lichess4545.com/) ([source code](https://github.com/cyanfish/heltour))
- [English Chess Federation](https://ecf.octoknight.com/)
- [Rotherham Online Chess](https://rotherhamonlinechess.azurewebsites.net/tournaments)

### Token format
Access tokens and authorization codes match `^[A-Za-z0-9_]+$`.
The length of tokens can be increased without notice. Make sure your application can handle at least 512 characters.
By convention tokens have a recognizable prefix, but do not rely on this.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 2.0.143
- Package version: 1.0.0
- Generator version: 7.23.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://lichess.org/api](https://lichess.org/api)

## Installation

Import the package in a go file in your project and run `go mod tidy`:

```go
import lichess "github.com/atye/golichess"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `lichess.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), lichess.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `lichess.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), lichess.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `lichess.ContextOperationServerIndices` and `lichess.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), lichess.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), lichess.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://lichess.org*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AccountAPI* | [**Account**](docs/AccountAPI.md#account) | **Get** /api/account/preferences | Get my preferences
*AccountAPI* | [**AccountEmail**](docs/AccountAPI.md#accountemail) | **Get** /api/account/email | Get my email address
*AccountAPI* | [**AccountKid**](docs/AccountAPI.md#accountkid) | **Get** /api/account/kid | Get my kid mode status
*AccountAPI* | [**AccountKidPost**](docs/AccountAPI.md#accountkidpost) | **Post** /api/account/kid | Set my kid mode status
*AccountAPI* | [**AccountMe**](docs/AccountAPI.md#accountme) | **Get** /api/account | Get my profile
*AccountAPI* | [**Timeline**](docs/AccountAPI.md#timeline) | **Get** /api/timeline | Get my timeline
*AnalysisAPI* | [**ApiCloudEval**](docs/AnalysisAPI.md#apicloudeval) | **Get** /api/cloud-eval | Get cloud evaluation of a position.
*BoardAPI* | [**ApiBoardSeek**](docs/BoardAPI.md#apiboardseek) | **Post** /api/board/seek | Create a seek
*BoardAPI* | [**ApiStreamEvent**](docs/BoardAPI.md#apistreamevent) | **Get** /api/stream/event | Stream incoming events
*BoardAPI* | [**BoardGameAbort**](docs/BoardAPI.md#boardgameabort) | **Post** /api/board/game/{gameId}/abort | Abort a game
*BoardAPI* | [**BoardGameBerserk**](docs/BoardAPI.md#boardgameberserk) | **Post** /api/board/game/{gameId}/berserk | Berserk a tournament game
*BoardAPI* | [**BoardGameChatGet**](docs/BoardAPI.md#boardgamechatget) | **Get** /api/board/game/{gameId}/chat | Fetch the player chat
*BoardAPI* | [**BoardGameChatPost**](docs/BoardAPI.md#boardgamechatpost) | **Post** /api/board/game/{gameId}/chat | Write in the chat
*BoardAPI* | [**BoardGameClaimDraw**](docs/BoardAPI.md#boardgameclaimdraw) | **Post** /api/board/game/{gameId}/claim-draw | Claim draw of a game
*BoardAPI* | [**BoardGameClaimVictory**](docs/BoardAPI.md#boardgameclaimvictory) | **Post** /api/board/game/{gameId}/claim-victory | Claim victory of a game
*BoardAPI* | [**BoardGameDraw**](docs/BoardAPI.md#boardgamedraw) | **Post** /api/board/game/{gameId}/draw/{accept} | Handle draw offers
*BoardAPI* | [**BoardGameMove**](docs/BoardAPI.md#boardgamemove) | **Post** /api/board/game/{gameId}/move/{move} | Make a Board move
*BoardAPI* | [**BoardGameResign**](docs/BoardAPI.md#boardgameresign) | **Post** /api/board/game/{gameId}/resign | Resign a game
*BoardAPI* | [**BoardGameStream**](docs/BoardAPI.md#boardgamestream) | **Get** /api/board/game/stream/{gameId} | Stream Board game state
*BoardAPI* | [**BoardGameTakeback**](docs/BoardAPI.md#boardgametakeback) | **Post** /api/board/game/{gameId}/takeback/{accept} | Handle takeback offers
*BotAPI* | [**ApiBotOnline**](docs/BotAPI.md#apibotonline) | **Get** /api/bot/online | Get online bots
*BotAPI* | [**ApiStreamEvent**](docs/BotAPI.md#apistreamevent) | **Get** /api/stream/event | Stream incoming events
*BotAPI* | [**BotAccountUpgrade**](docs/BotAPI.md#botaccountupgrade) | **Post** /api/bot/account/upgrade | Upgrade to Bot account
*BotAPI* | [**BotGameAbort**](docs/BotAPI.md#botgameabort) | **Post** /api/bot/game/{gameId}/abort | Abort a game
*BotAPI* | [**BotGameChat**](docs/BotAPI.md#botgamechat) | **Post** /api/bot/game/{gameId}/chat | Write in the chat
*BotAPI* | [**BotGameChatGet**](docs/BotAPI.md#botgamechatget) | **Get** /api/bot/game/{gameId}/chat | Fetch the game chat
*BotAPI* | [**BotGameClaimDraw**](docs/BotAPI.md#botgameclaimdraw) | **Post** /api/bot/game/{gameId}/claim-draw | Claim draw of a game
*BotAPI* | [**BotGameClaimVictory**](docs/BotAPI.md#botgameclaimvictory) | **Post** /api/bot/game/{gameId}/claim-victory | Claim victory of a game
*BotAPI* | [**BotGameDraw**](docs/BotAPI.md#botgamedraw) | **Post** /api/bot/game/{gameId}/draw/{accept} | Handle draw offers
*BotAPI* | [**BotGameMove**](docs/BotAPI.md#botgamemove) | **Post** /api/bot/game/{gameId}/move/{move} | Make a Bot move
*BotAPI* | [**BotGameResign**](docs/BotAPI.md#botgameresign) | **Post** /api/bot/game/{gameId}/resign | Resign a game
*BotAPI* | [**BotGameStream**](docs/BotAPI.md#botgamestream) | **Get** /api/bot/game/stream/{gameId} | Stream Bot game state
*BotAPI* | [**BotGameTakeback**](docs/BotAPI.md#botgametakeback) | **Post** /api/bot/game/{gameId}/takeback/{accept} | Handle takeback offers
*BroadcastsAPI* | [**BroadcastAllRoundsPgn**](docs/BroadcastsAPI.md#broadcastallroundspgn) | **Get** /api/broadcast/{broadcastTournamentId}.pgn | Export all rounds as PGN
*BroadcastsAPI* | [**BroadcastMyRoundsGet**](docs/BroadcastsAPI.md#broadcastmyroundsget) | **Get** /api/broadcast/my-rounds | Get your broadcast rounds
*BroadcastsAPI* | [**BroadcastPlayerGet**](docs/BroadcastsAPI.md#broadcastplayerget) | **Get** /broadcast/{broadcastTournamentId}/players/{playerId} | Get a player of a broadcast
*BroadcastsAPI* | [**BroadcastPlayersGet**](docs/BroadcastsAPI.md#broadcastplayersget) | **Get** /broadcast/{broadcastTournamentId}/players | Get players of a broadcast
*BroadcastsAPI* | [**BroadcastPush**](docs/BroadcastsAPI.md#broadcastpush) | **Post** /api/broadcast/round/{broadcastRoundId}/push | Push PGN to a broadcast round
*BroadcastsAPI* | [**BroadcastRoundCreate**](docs/BroadcastsAPI.md#broadcastroundcreate) | **Post** /broadcast/{broadcastTournamentId}/new | Create a broadcast round
*BroadcastsAPI* | [**BroadcastRoundGet**](docs/BroadcastsAPI.md#broadcastroundget) | **Get** /api/broadcast/{broadcastTournamentSlug}/{broadcastRoundSlug}/{broadcastRoundId} | Get a broadcast round
*BroadcastsAPI* | [**BroadcastRoundPgn**](docs/BroadcastsAPI.md#broadcastroundpgn) | **Get** /api/broadcast/round/{broadcastRoundId}.pgn | Export one round as PGN
*BroadcastsAPI* | [**BroadcastRoundReset**](docs/BroadcastsAPI.md#broadcastroundreset) | **Post** /api/broadcast/round/{broadcastRoundId}/reset | Reset a broadcast round
*BroadcastsAPI* | [**BroadcastRoundUpdate**](docs/BroadcastsAPI.md#broadcastroundupdate) | **Post** /broadcast/round/{broadcastRoundId}/edit | Update a broadcast round
*BroadcastsAPI* | [**BroadcastStreamRoundPgn**](docs/BroadcastsAPI.md#broadcaststreamroundpgn) | **Get** /api/stream/broadcast/round/{broadcastRoundId}.pgn | Stream an ongoing broadcast round as PGN
*BroadcastsAPI* | [**BroadcastTeamLeaderboardGet**](docs/BroadcastsAPI.md#broadcastteamleaderboardget) | **Get** /broadcast/{broadcastTournamentId}/teams/standings | Get the team leaderboard of a broadcast
*BroadcastsAPI* | [**BroadcastTourCreate**](docs/BroadcastsAPI.md#broadcasttourcreate) | **Post** /broadcast/new | Create a broadcast tournament
*BroadcastsAPI* | [**BroadcastTourGet**](docs/BroadcastsAPI.md#broadcasttourget) | **Get** /api/broadcast/{broadcastTournamentId} | Get a broadcast tournament
*BroadcastsAPI* | [**BroadcastTourUpdate**](docs/BroadcastsAPI.md#broadcasttourupdate) | **Post** /broadcast/{broadcastTournamentId}/edit | Update your broadcast tournament
*BroadcastsAPI* | [**BroadcastsByUser**](docs/BroadcastsAPI.md#broadcastsbyuser) | **Get** /api/broadcast/by/{username} | Get broadcasts created by a user
*BroadcastsAPI* | [**BroadcastsOfficial**](docs/BroadcastsAPI.md#broadcastsofficial) | **Get** /api/broadcast | Get official broadcasts
*BroadcastsAPI* | [**BroadcastsSearch**](docs/BroadcastsAPI.md#broadcastssearch) | **Get** /api/broadcast/search | Search broadcasts
*BroadcastsAPI* | [**BroadcastsTop**](docs/BroadcastsAPI.md#broadcaststop) | **Get** /api/broadcast/top | Get paginated top broadcast previews
*BulkPairingsAPI* | [**BulkPairingCreate**](docs/BulkPairingsAPI.md#bulkpairingcreate) | **Post** /api/bulk-pairing | Create a bulk pairing
*BulkPairingsAPI* | [**BulkPairingDelete**](docs/BulkPairingsAPI.md#bulkpairingdelete) | **Delete** /api/bulk-pairing/{id} | Cancel a bulk pairing
*BulkPairingsAPI* | [**BulkPairingGet**](docs/BulkPairingsAPI.md#bulkpairingget) | **Get** /api/bulk-pairing/{id} | Show a bulk pairing
*BulkPairingsAPI* | [**BulkPairingIdGamesGet**](docs/BulkPairingsAPI.md#bulkpairingidgamesget) | **Get** /api/bulk-pairing/{id}/games | Export games of a bulk pairing
*BulkPairingsAPI* | [**BulkPairingList**](docs/BulkPairingsAPI.md#bulkpairinglist) | **Get** /api/bulk-pairing | View your bulk pairings
*BulkPairingsAPI* | [**BulkPairingStartClocks**](docs/BulkPairingsAPI.md#bulkpairingstartclocks) | **Post** /api/bulk-pairing/{id}/start-clocks | Manually start clocks
*ChallengesAPI* | [**AdminChallengeTokens**](docs/ChallengesAPI.md#adminchallengetokens) | **Post** /api/token/admin-challenge | Admin challenge tokens
*ChallengesAPI* | [**ChallengeAccept**](docs/ChallengesAPI.md#challengeaccept) | **Post** /api/challenge/{challengeId}/accept | Accept a challenge
*ChallengesAPI* | [**ChallengeAi**](docs/ChallengesAPI.md#challengeai) | **Post** /api/challenge/ai | Challenge the AI
*ChallengesAPI* | [**ChallengeCancel**](docs/ChallengesAPI.md#challengecancel) | **Post** /api/challenge/{challengeId}/cancel | Cancel a challenge
*ChallengesAPI* | [**ChallengeCreate**](docs/ChallengesAPI.md#challengecreate) | **Post** /api/challenge/{username} | Create a challenge
*ChallengesAPI* | [**ChallengeDecline**](docs/ChallengesAPI.md#challengedecline) | **Post** /api/challenge/{challengeId}/decline | Decline a challenge
*ChallengesAPI* | [**ChallengeList**](docs/ChallengesAPI.md#challengelist) | **Get** /api/challenge | List your challenges
*ChallengesAPI* | [**ChallengeOpen**](docs/ChallengesAPI.md#challengeopen) | **Post** /api/challenge/open | Open-ended challenge
*ChallengesAPI* | [**ChallengeShow**](docs/ChallengesAPI.md#challengeshow) | **Get** /api/challenge/{challengeId}/show | Show one challenge
*ChallengesAPI* | [**ChallengeStartClocks**](docs/ChallengesAPI.md#challengestartclocks) | **Post** /api/challenge/{gameId}/start-clocks | Start clocks of a game
*ChallengesAPI* | [**RoundAddTime**](docs/ChallengesAPI.md#roundaddtime) | **Post** /api/round/{gameId}/add-time/{seconds} | Add time to the opponent clock
*ExternalEngineAPI* | [**ApiExternalEngineAcquire**](docs/ExternalEngineAPI.md#apiexternalengineacquire) | **Post** /api/external-engine/work | Acquire analysis request
*ExternalEngineAPI* | [**ApiExternalEngineAnalyse**](docs/ExternalEngineAPI.md#apiexternalengineanalyse) | **Post** /api/external-engine/{id}/analyse | Analyse with external engine
*ExternalEngineAPI* | [**ApiExternalEngineCreate**](docs/ExternalEngineAPI.md#apiexternalenginecreate) | **Post** /api/external-engine | Create external engine
*ExternalEngineAPI* | [**ApiExternalEngineDelete**](docs/ExternalEngineAPI.md#apiexternalenginedelete) | **Delete** /api/external-engine/{id} | Delete external engine
*ExternalEngineAPI* | [**ApiExternalEngineGet**](docs/ExternalEngineAPI.md#apiexternalengineget) | **Get** /api/external-engine/{id} | Get external engine
*ExternalEngineAPI* | [**ApiExternalEngineList**](docs/ExternalEngineAPI.md#apiexternalenginelist) | **Get** /api/external-engine | List external engines
*ExternalEngineAPI* | [**ApiExternalEnginePut**](docs/ExternalEngineAPI.md#apiexternalengineput) | **Put** /api/external-engine/{id} | Update external engine
*ExternalEngineAPI* | [**ApiExternalEngineSubmit**](docs/ExternalEngineAPI.md#apiexternalenginesubmit) | **Post** /api/external-engine/work/{id} | Answer analysis request
*FIDEAPI* | [**FidePlayerGet**](docs/FIDEAPI.md#fideplayerget) | **Get** /api/fide/player/{playerId} | Get a FIDE player
*FIDEAPI* | [**FidePlayerRatings**](docs/FIDEAPI.md#fideplayerratings) | **Get** /api/fide/player/{playerId}/ratings | Get ratings history of a FIDE player
*FIDEAPI* | [**FidePlayerSearch**](docs/FIDEAPI.md#fideplayersearch) | **Get** /api/fide/player | Search FIDE players
*GamesAPI* | [**ApiAccountPlaying**](docs/GamesAPI.md#apiaccountplaying) | **Get** /api/account/playing | Get my ongoing games
*GamesAPI* | [**ApiExportBookmarks**](docs/GamesAPI.md#apiexportbookmarks) | **Get** /api/games/export/bookmarks | Export your bookmarked games
*GamesAPI* | [**ApiGamesUser**](docs/GamesAPI.md#apigamesuser) | **Get** /api/games/user/{username} | Export games of a user
*GamesAPI* | [**ApiImportedGamesUser**](docs/GamesAPI.md#apiimportedgamesuser) | **Get** /api/games/export/imports | Export your imported games
*GamesAPI* | [**ApiUserCurrentGame**](docs/GamesAPI.md#apiusercurrentgame) | **Get** /api/user/{username}/current-game | Export ongoing game of a user
*GamesAPI* | [**GameChatGet**](docs/GamesAPI.md#gamechatget) | **Get** /game/{gameId}/chat | Fetch the spectator game chat
*GamesAPI* | [**GameImport**](docs/GamesAPI.md#gameimport) | **Post** /api/import | Import one game
*GamesAPI* | [**GamePgn**](docs/GamesAPI.md#gamepgn) | **Get** /game/export/{gameId} | Export one game
*GamesAPI* | [**GamesByIds**](docs/GamesAPI.md#gamesbyids) | **Post** /api/stream/games/{streamId} | Stream games by IDs
*GamesAPI* | [**GamesByIdsAdd**](docs/GamesAPI.md#gamesbyidsadd) | **Post** /api/stream/games/{streamId}/add | Add game IDs to stream
*GamesAPI* | [**GamesByUsers**](docs/GamesAPI.md#gamesbyusers) | **Post** /api/stream/games-by-users | Stream games of users
*GamesAPI* | [**GamesExportIds**](docs/GamesAPI.md#gamesexportids) | **Post** /api/games/export/_ids | Export games by IDs
*GamesAPI* | [**StreamGame**](docs/GamesAPI.md#streamgame) | **Get** /api/stream/game/{id} | Stream moves of a game
*MessagingAPI* | [**InboxUsername**](docs/MessagingAPI.md#inboxusername) | **Post** /inbox/{username} | Send a private message
*OAuthAPI* | [**ApiToken**](docs/OAuthAPI.md#apitoken) | **Post** /api/token | Obtain access token
*OAuthAPI* | [**ApiTokenDelete**](docs/OAuthAPI.md#apitokendelete) | **Delete** /api/token | Revoke access token
*OAuthAPI* | [**Oauth**](docs/OAuthAPI.md#oauth) | **Get** /oauth | Request authorization code
*OAuthAPI* | [**TokenTest**](docs/OAuthAPI.md#tokentest) | **Post** /api/token/test | Test multiple OAuth tokens
*OpeningExplorerAPI* | [**OpeningExplorerLichess**](docs/OpeningExplorerAPI.md#openingexplorerlichess) | **Get** /lichess | Lichess games
*OpeningExplorerAPI* | [**OpeningExplorerMaster**](docs/OpeningExplorerAPI.md#openingexplorermaster) | **Get** /masters | Masters database
*OpeningExplorerAPI* | [**OpeningExplorerMasterGame**](docs/OpeningExplorerAPI.md#openingexplorermastergame) | **Get** /masters/pgn/{gameId} | OTB master game
*OpeningExplorerAPI* | [**OpeningExplorerPlayer**](docs/OpeningExplorerAPI.md#openingexplorerplayer) | **Get** /player | Player games
*PuzzlesAPI* | [**ApiPuzzleActivity**](docs/PuzzlesAPI.md#apipuzzleactivity) | **Get** /api/puzzle/activity | Get your puzzle activity
*PuzzlesAPI* | [**ApiPuzzleBatchSelect**](docs/PuzzlesAPI.md#apipuzzlebatchselect) | **Get** /api/puzzle/batch/{angle} | Get multiple puzzles at once
*PuzzlesAPI* | [**ApiPuzzleBatchSolve**](docs/PuzzlesAPI.md#apipuzzlebatchsolve) | **Post** /api/puzzle/batch/{angle} | Solve multiple puzzles at once
*PuzzlesAPI* | [**ApiPuzzleDaily**](docs/PuzzlesAPI.md#apipuzzledaily) | **Get** /api/puzzle/daily | Get the daily puzzle
*PuzzlesAPI* | [**ApiPuzzleDashboard**](docs/PuzzlesAPI.md#apipuzzledashboard) | **Get** /api/puzzle/dashboard/{days} | Get your puzzle dashboard
*PuzzlesAPI* | [**ApiPuzzleId**](docs/PuzzlesAPI.md#apipuzzleid) | **Get** /api/puzzle/{id} | Get a puzzle by its ID
*PuzzlesAPI* | [**ApiPuzzleNext**](docs/PuzzlesAPI.md#apipuzzlenext) | **Get** /api/puzzle/next | Get a new puzzle
*PuzzlesAPI* | [**ApiPuzzleReplay**](docs/PuzzlesAPI.md#apipuzzlereplay) | **Get** /api/puzzle/replay/{days}/{theme} | Get puzzles to replay
*PuzzlesAPI* | [**ApiStormDashboard**](docs/PuzzlesAPI.md#apistormdashboard) | **Get** /api/storm/dashboard/{username} | Get the storm dashboard of a player
*PuzzlesAPI* | [**RacerGet**](docs/PuzzlesAPI.md#racerget) | **Get** /api/racer/{id} | Get puzzle race results
*PuzzlesAPI* | [**RacerPost**](docs/PuzzlesAPI.md#racerpost) | **Post** /api/racer | Create and join a puzzle race
*RelationsAPI* | [**ApiUserFollowing**](docs/RelationsAPI.md#apiuserfollowing) | **Get** /api/rel/following | Get users followed by the logged in user
*RelationsAPI* | [**BlockUser**](docs/RelationsAPI.md#blockuser) | **Post** /api/rel/block/{username} | Block a player
*RelationsAPI* | [**FollowUser**](docs/RelationsAPI.md#followuser) | **Post** /api/rel/follow/{username} | Follow a player
*RelationsAPI* | [**UnblockUser**](docs/RelationsAPI.md#unblockuser) | **Post** /api/rel/unblock/{username} | Unblock a player
*RelationsAPI* | [**UnfollowUser**](docs/RelationsAPI.md#unfollowuser) | **Post** /api/rel/unfollow/{username} | Unfollow a player
*SimulsAPI* | [**ApiSimul**](docs/SimulsAPI.md#apisimul) | **Get** /api/simul | Get current simuls
*StudiesAPI* | [**ApiStudyChapterMoves**](docs/StudiesAPI.md#apistudychaptermoves) | **Post** /api/study/{studyId}/{chapterId}/moves | Update the moves of a study chapter
*StudiesAPI* | [**ApiStudyChapterTags**](docs/StudiesAPI.md#apistudychaptertags) | **Post** /api/study/{studyId}/{chapterId}/tags | Update PGN tags of a study chapter
*StudiesAPI* | [**ApiStudyImportPGN**](docs/StudiesAPI.md#apistudyimportpgn) | **Post** /api/study/{studyId}/import-pgn | Import PGN into a study
*StudiesAPI* | [**ApiStudyPost**](docs/StudiesAPI.md#apistudypost) | **Post** /api/study | Create a new Study
*StudiesAPI* | [**ApiStudyStudyIdChapterIdDelete**](docs/StudiesAPI.md#apistudystudyidchapteriddelete) | **Delete** /api/study/{studyId}/{chapterId} | Delete a study chapter
*StudiesAPI* | [**StudyAllChaptersHead**](docs/StudiesAPI.md#studyallchaptershead) | **Head** /api/study/{studyId}.pgn | Study metadata
*StudiesAPI* | [**StudyAllChaptersPgn**](docs/StudiesAPI.md#studyallchapterspgn) | **Get** /api/study/{studyId}.pgn | Export all chapters
*StudiesAPI* | [**StudyChapterPgn**](docs/StudiesAPI.md#studychapterpgn) | **Get** /api/study/{studyId}/{chapterId}.pgn | Export one study chapter
*StudiesAPI* | [**StudyExportAllPgn**](docs/StudiesAPI.md#studyexportallpgn) | **Get** /api/study/by/{username}/export.pgn | Export all studies of a user
*StudiesAPI* | [**StudyListMetadata**](docs/StudiesAPI.md#studylistmetadata) | **Get** /api/study/by/{username} | List studies of a user
*TVAPI* | [**TvChannelFeed**](docs/TVAPI.md#tvchannelfeed) | **Get** /api/tv/{channel}/feed | Stream current TV game of a TV channel
*TVAPI* | [**TvChannelGames**](docs/TVAPI.md#tvchannelgames) | **Get** /api/tv/{channel} | Get best ongoing games of a TV channel
*TVAPI* | [**TvChannels**](docs/TVAPI.md#tvchannels) | **Get** /api/tv/channels | Get current TV games
*TVAPI* | [**TvFeed**](docs/TVAPI.md#tvfeed) | **Get** /api/tv/feed | Stream current TV game
*TablebaseAPI* | [**AntichessAtomic**](docs/TablebaseAPI.md#antichessatomic) | **Get** /antichess | Tablebase lookup for Antichess
*TablebaseAPI* | [**TablebaseAtomic**](docs/TablebaseAPI.md#tablebaseatomic) | **Get** /atomic | Tablebase lookup for Atomic chess
*TablebaseAPI* | [**TablebaseStandard**](docs/TablebaseAPI.md#tablebasestandard) | **Get** /standard | Tablebase lookup
*TeamsAPI* | [**ApiTeamArena**](docs/TeamsAPI.md#apiteamarena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
*TeamsAPI* | [**ApiTeamSwiss**](docs/TeamsAPI.md#apiteamswiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
*TeamsAPI* | [**TeamAll**](docs/TeamsAPI.md#teamall) | **Get** /api/team/all | Get popular teams
*TeamsAPI* | [**TeamIdJoin**](docs/TeamsAPI.md#teamidjoin) | **Post** /team/{teamId}/join | Join a team
*TeamsAPI* | [**TeamIdKickUserId**](docs/TeamsAPI.md#teamidkickuserid) | **Post** /api/team/{teamId}/kick/{userId} | Kick a user from your team
*TeamsAPI* | [**TeamIdPmAll**](docs/TeamsAPI.md#teamidpmall) | **Post** /team/{teamId}/pm-all | Message all members
*TeamsAPI* | [**TeamIdQuit**](docs/TeamsAPI.md#teamidquit) | **Post** /team/{teamId}/quit | Leave a team
*TeamsAPI* | [**TeamIdUsers**](docs/TeamsAPI.md#teamidusers) | **Get** /api/team/{teamId}/users | Get members of a team
*TeamsAPI* | [**TeamOfUsername**](docs/TeamsAPI.md#teamofusername) | **Get** /api/team/of/{username} | Teams of a player
*TeamsAPI* | [**TeamRequestAccept**](docs/TeamsAPI.md#teamrequestaccept) | **Post** /api/team/{teamId}/request/{userId}/accept | Accept join request
*TeamsAPI* | [**TeamRequestDecline**](docs/TeamsAPI.md#teamrequestdecline) | **Post** /api/team/{teamId}/request/{userId}/decline | Decline join request
*TeamsAPI* | [**TeamRequests**](docs/TeamsAPI.md#teamrequests) | **Get** /api/team/{teamId}/requests | Get join requests
*TeamsAPI* | [**TeamSearch**](docs/TeamsAPI.md#teamsearch) | **Get** /api/team/search | Search teams
*TeamsAPI* | [**TeamShow**](docs/TeamsAPI.md#teamshow) | **Get** /api/team/{teamId} | Get a single team
*TournamentsArenaAPI* | [**ApiTeamArena**](docs/TournamentsArenaAPI.md#apiteamarena) | **Get** /api/team/{teamId}/arena | Get team Arena tournaments
*TournamentsArenaAPI* | [**ApiTournament**](docs/TournamentsArenaAPI.md#apitournament) | **Get** /api/tournament | Get current tournaments
*TournamentsArenaAPI* | [**ApiTournamentJoin**](docs/TournamentsArenaAPI.md#apitournamentjoin) | **Post** /api/tournament/{id}/join | Join an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentPost**](docs/TournamentsArenaAPI.md#apitournamentpost) | **Post** /api/tournament | Create a new Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentTeamBattlePost**](docs/TournamentsArenaAPI.md#apitournamentteambattlepost) | **Post** /api/tournament/team-battle/{id} | Update a team battle
*TournamentsArenaAPI* | [**ApiTournamentTerminate**](docs/TournamentsArenaAPI.md#apitournamentterminate) | **Post** /api/tournament/{id}/terminate | Terminate an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentUpdate**](docs/TournamentsArenaAPI.md#apitournamentupdate) | **Post** /api/tournament/{id} | Update an Arena tournament
*TournamentsArenaAPI* | [**ApiTournamentWithdraw**](docs/TournamentsArenaAPI.md#apitournamentwithdraw) | **Post** /api/tournament/{id}/withdraw | Pause or leave an Arena tournament
*TournamentsArenaAPI* | [**ApiUserNameTournamentCreated**](docs/TournamentsArenaAPI.md#apiusernametournamentcreated) | **Get** /api/user/{username}/tournament/created | Get tournaments created by a user
*TournamentsArenaAPI* | [**ApiUserNameTournamentPlayed**](docs/TournamentsArenaAPI.md#apiusernametournamentplayed) | **Get** /api/user/{username}/tournament/played | Get tournaments played by a user
*TournamentsArenaAPI* | [**GamesByTournament**](docs/TournamentsArenaAPI.md#gamesbytournament) | **Get** /api/tournament/{id}/games | Export games of an Arena tournament
*TournamentsArenaAPI* | [**ResultsByTournament**](docs/TournamentsArenaAPI.md#resultsbytournament) | **Get** /api/tournament/{id}/results | Get results of an Arena tournament
*TournamentsArenaAPI* | [**TeamsByTournament**](docs/TournamentsArenaAPI.md#teamsbytournament) | **Get** /api/tournament/{id}/teams | Get team standing of a team battle
*TournamentsArenaAPI* | [**Tournament**](docs/TournamentsArenaAPI.md#tournament) | **Get** /api/tournament/{id} | Get info about an Arena tournament
*TournamentsSwissAPI* | [**ApiSwissJoin**](docs/TournamentsSwissAPI.md#apiswissjoin) | **Post** /api/swiss/{id}/join | Join a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissNew**](docs/TournamentsSwissAPI.md#apiswissnew) | **Post** /api/swiss/new/{teamId} | Create a new Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissScheduleNextRound**](docs/TournamentsSwissAPI.md#apiswissschedulenextround) | **Post** /api/swiss/{id}/schedule-next-round | Manually schedule the next round
*TournamentsSwissAPI* | [**ApiSwissTerminate**](docs/TournamentsSwissAPI.md#apiswissterminate) | **Post** /api/swiss/{id}/terminate | Terminate a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissUpdate**](docs/TournamentsSwissAPI.md#apiswissupdate) | **Post** /api/swiss/{id}/edit | Update a Swiss tournament
*TournamentsSwissAPI* | [**ApiSwissWithdraw**](docs/TournamentsSwissAPI.md#apiswisswithdraw) | **Post** /api/swiss/{id}/withdraw | Pause or leave a swiss tournament
*TournamentsSwissAPI* | [**ApiTeamSwiss**](docs/TournamentsSwissAPI.md#apiteamswiss) | **Get** /api/team/{teamId}/swiss | Get team swiss tournaments
*TournamentsSwissAPI* | [**GamesBySwiss**](docs/TournamentsSwissAPI.md#gamesbyswiss) | **Get** /api/swiss/{id}/games | Export games of a Swiss tournament
*TournamentsSwissAPI* | [**ResultsBySwiss**](docs/TournamentsSwissAPI.md#resultsbyswiss) | **Get** /api/swiss/{id}/results | Get results of a swiss tournament
*TournamentsSwissAPI* | [**Swiss**](docs/TournamentsSwissAPI.md#swiss) | **Get** /api/swiss/{id} | Get info about a Swiss tournament
*TournamentsSwissAPI* | [**SwissTrf**](docs/TournamentsSwissAPI.md#swisstrf) | **Get** /swiss/{id}.trf | Export TRF of a Swiss tournament
*UsersAPI* | [**ApiCrosstable**](docs/UsersAPI.md#apicrosstable) | **Get** /api/crosstable/{user1}/{user2} | Get crosstable
*UsersAPI* | [**ApiPlayerAutocomplete**](docs/UsersAPI.md#apiplayerautocomplete) | **Get** /api/player/autocomplete | Autocomplete usernames
*UsersAPI* | [**ApiUser**](docs/UsersAPI.md#apiuser) | **Get** /api/user/{username} | Get user public data
*UsersAPI* | [**ApiUserActivity**](docs/UsersAPI.md#apiuseractivity) | **Get** /api/user/{username}/activity | Get user activity
*UsersAPI* | [**ApiUserPerf**](docs/UsersAPI.md#apiuserperf) | **Get** /api/user/{username}/perf/{perf} | Get performance statistics of a user
*UsersAPI* | [**ApiUserRatingHistory**](docs/UsersAPI.md#apiuserratinghistory) | **Get** /api/user/{username}/rating-history | Get rating history of a user
*UsersAPI* | [**ApiUsers**](docs/UsersAPI.md#apiusers) | **Post** /api/users | Get users by ID
*UsersAPI* | [**ApiUsersStatus**](docs/UsersAPI.md#apiusersstatus) | **Get** /api/users/status | Get real-time users status
*UsersAPI* | [**Player**](docs/UsersAPI.md#player) | **Get** /api/player | Get all top 10
*UsersAPI* | [**PlayerTopNbPerfType**](docs/UsersAPI.md#playertopnbperftype) | **Get** /api/player/top/{nb}/{perfType} | Get one leaderboard
*UsersAPI* | [**ReadNote**](docs/UsersAPI.md#readnote) | **Get** /api/user/{username}/note | Get notes for a user
*UsersAPI* | [**StreamerLive**](docs/UsersAPI.md#streamerlive) | **Get** /api/streamer/live | Get live streamers
*UsersAPI* | [**WriteNote**](docs/UsersAPI.md#writenote) | **Post** /api/user/{username}/note | Add a note for a user


## Documentation For Models

 - [AIOpponent](docs/AIOpponent.md)
 - [Account200Response](docs/Account200Response.md)
 - [Account200ResponsePrefs](docs/Account200ResponsePrefs.md)
 - [AccountEmail200Response](docs/AccountEmail200Response.md)
 - [AccountKid200Response](docs/AccountKid200Response.md)
 - [AccountKidPost200Response](docs/AccountKidPost200Response.md)
 - [AccountMe200Response](docs/AccountMe200Response.md)
 - [ApiAccountPlaying200Response](docs/ApiAccountPlaying200Response.md)
 - [ApiAccountPlaying200ResponseNowPlayingInner](docs/ApiAccountPlaying200ResponseNowPlayingInner.md)
 - [ApiAccountPlaying200ResponseNowPlayingInnerOpponent](docs/ApiAccountPlaying200ResponseNowPlayingInnerOpponent.md)
 - [ApiAccountPlaying200ResponseNowPlayingInnerVariant](docs/ApiAccountPlaying200ResponseNowPlayingInnerVariant.md)
 - [ApiBoardSeek200Response](docs/ApiBoardSeek200Response.md)
 - [ApiCloudEval200Response](docs/ApiCloudEval200Response.md)
 - [ApiCloudEval200ResponsePvsInner](docs/ApiCloudEval200ResponsePvsInner.md)
 - [ApiCloudEval404Response](docs/ApiCloudEval404Response.md)
 - [ApiCrosstable200Response](docs/ApiCrosstable200Response.md)
 - [ApiExternalEngineAcquire200Response](docs/ApiExternalEngineAcquire200Response.md)
 - [ApiExternalEngineAcquireRequest](docs/ApiExternalEngineAcquireRequest.md)
 - [ApiExternalEngineAnalyse200Response](docs/ApiExternalEngineAnalyse200Response.md)
 - [ApiExternalEngineAnalyse200ResponsePvsInner](docs/ApiExternalEngineAnalyse200ResponsePvsInner.md)
 - [ApiExternalEngineAnalyseRequest](docs/ApiExternalEngineAnalyseRequest.md)
 - [ApiExternalEngineAnalyseRequestWork](docs/ApiExternalEngineAnalyseRequestWork.md)
 - [ApiExternalEngineAnalyseRequestWorkOneOf](docs/ApiExternalEngineAnalyseRequestWorkOneOf.md)
 - [ApiExternalEngineAnalyseRequestWorkOneOf1](docs/ApiExternalEngineAnalyseRequestWorkOneOf1.md)
 - [ApiExternalEngineAnalyseRequestWorkOneOf2](docs/ApiExternalEngineAnalyseRequestWorkOneOf2.md)
 - [ApiExternalEngineCreateRequest](docs/ApiExternalEngineCreateRequest.md)
 - [ApiExternalEngineList200ResponseInner](docs/ApiExternalEngineList200ResponseInner.md)
 - [ApiPlayerAutocomplete200Response](docs/ApiPlayerAutocomplete200Response.md)
 - [ApiPlayerAutocomplete200ResponseOneOf](docs/ApiPlayerAutocomplete200ResponseOneOf.md)
 - [ApiPlayerAutocomplete200ResponseOneOfResultInner](docs/ApiPlayerAutocomplete200ResponseOneOfResultInner.md)
 - [ApiPuzzleActivity200Response](docs/ApiPuzzleActivity200Response.md)
 - [ApiPuzzleActivity200ResponsePuzzle](docs/ApiPuzzleActivity200ResponsePuzzle.md)
 - [ApiPuzzleBatchSelect200Response](docs/ApiPuzzleBatchSelect200Response.md)
 - [ApiPuzzleBatchSolve200Response](docs/ApiPuzzleBatchSolve200Response.md)
 - [ApiPuzzleBatchSolve200ResponseRoundsInner](docs/ApiPuzzleBatchSolve200ResponseRoundsInner.md)
 - [ApiPuzzleBatchSolveRequest](docs/ApiPuzzleBatchSolveRequest.md)
 - [ApiPuzzleBatchSolveRequestSolutionsInner](docs/ApiPuzzleBatchSolveRequestSolutionsInner.md)
 - [ApiPuzzleDaily200Response](docs/ApiPuzzleDaily200Response.md)
 - [ApiPuzzleDaily200ResponseGame](docs/ApiPuzzleDaily200ResponseGame.md)
 - [ApiPuzzleDaily200ResponseGamePerf](docs/ApiPuzzleDaily200ResponseGamePerf.md)
 - [ApiPuzzleDaily200ResponseGamePlayersInner](docs/ApiPuzzleDaily200ResponseGamePlayersInner.md)
 - [ApiPuzzleDaily200ResponsePuzzle](docs/ApiPuzzleDaily200ResponsePuzzle.md)
 - [ApiPuzzleDashboard200Response](docs/ApiPuzzleDashboard200Response.md)
 - [ApiPuzzleDashboard200ResponseGlobal](docs/ApiPuzzleDashboard200ResponseGlobal.md)
 - [ApiPuzzleDashboard200ResponseThemesValue](docs/ApiPuzzleDashboard200ResponseThemesValue.md)
 - [ApiPuzzleId200Response](docs/ApiPuzzleId200Response.md)
 - [ApiPuzzleId200ResponseGame](docs/ApiPuzzleId200ResponseGame.md)
 - [ApiPuzzleId200ResponseGamePlayersInner](docs/ApiPuzzleId200ResponseGamePlayersInner.md)
 - [ApiPuzzleReplay200Response](docs/ApiPuzzleReplay200Response.md)
 - [ApiPuzzleReplay200ResponseAngle](docs/ApiPuzzleReplay200ResponseAngle.md)
 - [ApiPuzzleReplay200ResponseReplay](docs/ApiPuzzleReplay200ResponseReplay.md)
 - [ApiPuzzleReplay404Response](docs/ApiPuzzleReplay404Response.md)
 - [ApiSimul200Response](docs/ApiSimul200Response.md)
 - [ApiSimul200ResponsePendingInner](docs/ApiSimul200ResponsePendingInner.md)
 - [ApiSimul200ResponsePendingInnerHost](docs/ApiSimul200ResponsePendingInnerHost.md)
 - [ApiSimul200ResponsePendingInnerVariantsInner](docs/ApiSimul200ResponsePendingInnerVariantsInner.md)
 - [ApiStormDashboard200Response](docs/ApiStormDashboard200Response.md)
 - [ApiStormDashboard200ResponseDaysInner](docs/ApiStormDashboard200ResponseDaysInner.md)
 - [ApiStormDashboard200ResponseHigh](docs/ApiStormDashboard200ResponseHigh.md)
 - [ApiStreamEvent200Response](docs/ApiStreamEvent200Response.md)
 - [ApiStreamEvent200ResponseOneOf](docs/ApiStreamEvent200ResponseOneOf.md)
 - [ApiStreamEvent200ResponseOneOf1](docs/ApiStreamEvent200ResponseOneOf1.md)
 - [ApiStreamEvent200ResponseOneOf1Game](docs/ApiStreamEvent200ResponseOneOf1Game.md)
 - [ApiStreamEvent200ResponseOneOf2](docs/ApiStreamEvent200ResponseOneOf2.md)
 - [ApiStreamEvent200ResponseOneOf2Challenge](docs/ApiStreamEvent200ResponseOneOf2Challenge.md)
 - [ApiStreamEvent200ResponseOneOf2ChallengeChallenger](docs/ApiStreamEvent200ResponseOneOf2ChallengeChallenger.md)
 - [ApiStreamEvent200ResponseOneOf2ChallengeDestUser](docs/ApiStreamEvent200ResponseOneOf2ChallengeDestUser.md)
 - [ApiStreamEvent200ResponseOneOf2ChallengePerf](docs/ApiStreamEvent200ResponseOneOf2ChallengePerf.md)
 - [ApiStreamEvent200ResponseOneOf2ChallengeTimeControl](docs/ApiStreamEvent200ResponseOneOf2ChallengeTimeControl.md)
 - [ApiStreamEvent200ResponseOneOf3](docs/ApiStreamEvent200ResponseOneOf3.md)
 - [ApiStreamEvent200ResponseOneOf3Challenge](docs/ApiStreamEvent200ResponseOneOf3Challenge.md)
 - [ApiStreamEvent200ResponseOneOf3ChallengeChallenger](docs/ApiStreamEvent200ResponseOneOf3ChallengeChallenger.md)
 - [ApiStreamEvent200ResponseOneOf4](docs/ApiStreamEvent200ResponseOneOf4.md)
 - [ApiStreamEvent200ResponseOneOf4Challenge](docs/ApiStreamEvent200ResponseOneOf4Challenge.md)
 - [ApiStreamEvent200ResponseOneOfGame](docs/ApiStreamEvent200ResponseOneOfGame.md)
 - [ApiStreamEvent200ResponseOneOfGameCompat](docs/ApiStreamEvent200ResponseOneOfGameCompat.md)
 - [ApiStreamEvent200ResponseOneOfGameOpponent](docs/ApiStreamEvent200ResponseOneOfGameOpponent.md)
 - [ApiStudyImportPGN200Response](docs/ApiStudyImportPGN200Response.md)
 - [ApiStudyImportPGN200ResponseChaptersInner](docs/ApiStudyImportPGN200ResponseChaptersInner.md)
 - [ApiStudyImportPGN200ResponseChaptersInnerPlayersInner](docs/ApiStudyImportPGN200ResponseChaptersInnerPlayersInner.md)
 - [ApiStudyPost200Response](docs/ApiStudyPost200Response.md)
 - [ApiSwissNew200Response](docs/ApiSwissNew200Response.md)
 - [ApiSwissNew200ResponseClock](docs/ApiSwissNew200ResponseClock.md)
 - [ApiSwissNew200ResponseNextRound](docs/ApiSwissNew200ResponseNextRound.md)
 - [ApiSwissNew200ResponseStats](docs/ApiSwissNew200ResponseStats.md)
 - [ApiSwissUpdate401Response](docs/ApiSwissUpdate401Response.md)
 - [ApiToken200Response](docs/ApiToken200Response.md)
 - [ApiToken400Response](docs/ApiToken400Response.md)
 - [ApiTournament200Response](docs/ApiTournament200Response.md)
 - [ApiTournament200ResponseCreatedInner](docs/ApiTournament200ResponseCreatedInner.md)
 - [ApiTournament200ResponseCreatedInnerClock](docs/ApiTournament200ResponseCreatedInnerClock.md)
 - [ApiTournament200ResponseCreatedInnerMaxRating](docs/ApiTournament200ResponseCreatedInnerMaxRating.md)
 - [ApiTournament200ResponseCreatedInnerMinRatedGames](docs/ApiTournament200ResponseCreatedInnerMinRatedGames.md)
 - [ApiTournament200ResponseCreatedInnerPerf](docs/ApiTournament200ResponseCreatedInnerPerf.md)
 - [ApiTournament200ResponseCreatedInnerPosition](docs/ApiTournament200ResponseCreatedInnerPosition.md)
 - [ApiTournament200ResponseCreatedInnerSchedule](docs/ApiTournament200ResponseCreatedInnerSchedule.md)
 - [ApiTournament200ResponseCreatedInnerTeamBattle](docs/ApiTournament200ResponseCreatedInnerTeamBattle.md)
 - [ApiTournamentPost200Response](docs/ApiTournamentPost200Response.md)
 - [ApiTournamentPost200ResponseDuelsInner](docs/ApiTournamentPost200ResponseDuelsInner.md)
 - [ApiTournamentPost200ResponseDuelsInnerPInner](docs/ApiTournamentPost200ResponseDuelsInnerPInner.md)
 - [ApiTournamentPost200ResponseFeatured](docs/ApiTournamentPost200ResponseFeatured.md)
 - [ApiTournamentPost200ResponseFeaturedC](docs/ApiTournamentPost200ResponseFeaturedC.md)
 - [ApiTournamentPost200ResponseFeaturedWhite](docs/ApiTournamentPost200ResponseFeaturedWhite.md)
 - [ApiTournamentPost200ResponseGreatPlayer](docs/ApiTournamentPost200ResponseGreatPlayer.md)
 - [ApiTournamentPost200ResponsePerf](docs/ApiTournamentPost200ResponsePerf.md)
 - [ApiTournamentPost200ResponsePodiumInner](docs/ApiTournamentPost200ResponsePodiumInner.md)
 - [ApiTournamentPost200ResponsePodiumInnerNb](docs/ApiTournamentPost200ResponsePodiumInnerNb.md)
 - [ApiTournamentPost200ResponseQuote](docs/ApiTournamentPost200ResponseQuote.md)
 - [ApiTournamentPost200ResponseSchedule](docs/ApiTournamentPost200ResponseSchedule.md)
 - [ApiTournamentPost200ResponseSpotlight](docs/ApiTournamentPost200ResponseSpotlight.md)
 - [ApiTournamentPost200ResponseStanding](docs/ApiTournamentPost200ResponseStanding.md)
 - [ApiTournamentPost200ResponseStandingPlayersInner](docs/ApiTournamentPost200ResponseStandingPlayersInner.md)
 - [ApiTournamentPost200ResponseStandingPlayersInnerSheet](docs/ApiTournamentPost200ResponseStandingPlayersInnerSheet.md)
 - [ApiTournamentPost200ResponseStats](docs/ApiTournamentPost200ResponseStats.md)
 - [ApiTournamentPost200ResponseVerdicts](docs/ApiTournamentPost200ResponseVerdicts.md)
 - [ApiTournamentPost200ResponseVerdictsListInner](docs/ApiTournamentPost200ResponseVerdictsListInner.md)
 - [ApiTournamentPost400Response](docs/ApiTournamentPost400Response.md)
 - [ApiUser200Response](docs/ApiUser200Response.md)
 - [ApiUser200ResponseAllOfCount](docs/ApiUser200ResponseAllOfCount.md)
 - [ApiUser200ResponseAllOfPerfs](docs/ApiUser200ResponseAllOfPerfs.md)
 - [ApiUser200ResponseAllOfPerfsChess960](docs/ApiUser200ResponseAllOfPerfsChess960.md)
 - [ApiUser200ResponseAllOfPerfsStorm](docs/ApiUser200ResponseAllOfPerfsStorm.md)
 - [ApiUser200ResponseAllOfPlayTime](docs/ApiUser200ResponseAllOfPlayTime.md)
 - [ApiUser200ResponseAllOfProfile](docs/ApiUser200ResponseAllOfProfile.md)
 - [ApiUser200ResponseAllOfStreamer](docs/ApiUser200ResponseAllOfStreamer.md)
 - [ApiUser200ResponseAllOfStreamerTwitch](docs/ApiUser200ResponseAllOfStreamerTwitch.md)
 - [ApiUser200ResponseAllOfStreamerYoutube](docs/ApiUser200ResponseAllOfStreamerYoutube.md)
 - [ApiUserActivity200ResponseInner](docs/ApiUserActivity200ResponseInner.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceEnds](docs/ApiUserActivity200ResponseInnerCorrespondenceEnds.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceEndsCorrespondence](docs/ApiUserActivity200ResponseInnerCorrespondenceEndsCorrespondence.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceMoves](docs/ApiUserActivity200ResponseInnerCorrespondenceMoves.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner](docs/ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInner.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent](docs/ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponent.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponentOneOf](docs/ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponentOneOf.md)
 - [ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponentOneOf1](docs/ApiUserActivity200ResponseInnerCorrespondenceMovesGamesInnerOpponentOneOf1.md)
 - [ApiUserActivity200ResponseInnerFollows](docs/ApiUserActivity200ResponseInnerFollows.md)
 - [ApiUserActivity200ResponseInnerFollowsIn](docs/ApiUserActivity200ResponseInnerFollowsIn.md)
 - [ApiUserActivity200ResponseInnerGames](docs/ApiUserActivity200ResponseInnerGames.md)
 - [ApiUserActivity200ResponseInnerGamesChess960](docs/ApiUserActivity200ResponseInnerGamesChess960.md)
 - [ApiUserActivity200ResponseInnerGamesChess960Rp](docs/ApiUserActivity200ResponseInnerGamesChess960Rp.md)
 - [ApiUserActivity200ResponseInnerInterval](docs/ApiUserActivity200ResponseInnerInterval.md)
 - [ApiUserActivity200ResponseInnerPatron](docs/ApiUserActivity200ResponseInnerPatron.md)
 - [ApiUserActivity200ResponseInnerPostsInner](docs/ApiUserActivity200ResponseInnerPostsInner.md)
 - [ApiUserActivity200ResponseInnerPostsInnerPostsInner](docs/ApiUserActivity200ResponseInnerPostsInnerPostsInner.md)
 - [ApiUserActivity200ResponseInnerPracticeInner](docs/ApiUserActivity200ResponseInnerPracticeInner.md)
 - [ApiUserActivity200ResponseInnerPuzzles](docs/ApiUserActivity200ResponseInnerPuzzles.md)
 - [ApiUserActivity200ResponseInnerTeamsInner](docs/ApiUserActivity200ResponseInnerTeamsInner.md)
 - [ApiUserActivity200ResponseInnerTournaments](docs/ApiUserActivity200ResponseInnerTournaments.md)
 - [ApiUserActivity200ResponseInnerTournamentsBestInner](docs/ApiUserActivity200ResponseInnerTournamentsBestInner.md)
 - [ApiUserActivity200ResponseInnerTournamentsBestInnerTournament](docs/ApiUserActivity200ResponseInnerTournamentsBestInnerTournament.md)
 - [ApiUserCurrentGame200Response](docs/ApiUserCurrentGame200Response.md)
 - [ApiUserCurrentGame200ResponseOneOf](docs/ApiUserCurrentGame200ResponseOneOf.md)
 - [ApiUserCurrentGame200ResponseOneOfPlayers](docs/ApiUserCurrentGame200ResponseOneOfPlayers.md)
 - [ApiUserCurrentGame200ResponseOneOfPlayersWhite](docs/ApiUserCurrentGame200ResponseOneOfPlayersWhite.md)
 - [ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser](docs/ApiUserCurrentGame200ResponseOneOfPlayersWhiteUser.md)
 - [ApiUserNameTournamentPlayed200Response](docs/ApiUserNameTournamentPlayed200Response.md)
 - [ApiUserNameTournamentPlayed200ResponsePlayer](docs/ApiUserNameTournamentPlayed200ResponsePlayer.md)
 - [ApiUserPerf200Response](docs/ApiUserPerf200Response.md)
 - [ApiUserPerf200ResponsePerf](docs/ApiUserPerf200ResponsePerf.md)
 - [ApiUserPerf200ResponsePerfGlicko](docs/ApiUserPerf200ResponsePerfGlicko.md)
 - [ApiUserPerf200ResponseStat](docs/ApiUserPerf200ResponseStat.md)
 - [ApiUserPerf200ResponseStatBestWins](docs/ApiUserPerf200ResponseStatBestWins.md)
 - [ApiUserPerf200ResponseStatBestWinsResultsInner](docs/ApiUserPerf200ResponseStatBestWinsResultsInner.md)
 - [ApiUserPerf200ResponseStatBestWinsResultsInnerOpId](docs/ApiUserPerf200ResponseStatBestWinsResultsInnerOpId.md)
 - [ApiUserPerf200ResponseStatCount](docs/ApiUserPerf200ResponseStatCount.md)
 - [ApiUserPerf200ResponseStatHighest](docs/ApiUserPerf200ResponseStatHighest.md)
 - [ApiUserPerf200ResponseStatPlayStreak](docs/ApiUserPerf200ResponseStatPlayStreak.md)
 - [ApiUserPerf200ResponseStatPlayStreakNb](docs/ApiUserPerf200ResponseStatPlayStreakNb.md)
 - [ApiUserPerf200ResponseStatPlayStreakNbCur](docs/ApiUserPerf200ResponseStatPlayStreakNbCur.md)
 - [ApiUserPerf200ResponseStatResultStreak](docs/ApiUserPerf200ResponseStatResultStreak.md)
 - [ApiUserPerf200ResponseStatResultStreakLoss](docs/ApiUserPerf200ResponseStatResultStreakLoss.md)
 - [ApiUserPerf200ResponseStatResultStreakLossMax](docs/ApiUserPerf200ResponseStatResultStreakLossMax.md)
 - [ApiUserPerf200ResponseStatResultStreakLossMaxFrom](docs/ApiUserPerf200ResponseStatResultStreakLossMaxFrom.md)
 - [ApiUserPerf200ResponseStatResultStreakWin](docs/ApiUserPerf200ResponseStatResultStreakWin.md)
 - [ApiUserPerf200ResponseStatResultStreakWinCur](docs/ApiUserPerf200ResponseStatResultStreakWinCur.md)
 - [ApiUserPerf200ResponseStatResultStreakWinCurFrom](docs/ApiUserPerf200ResponseStatResultStreakWinCurFrom.md)
 - [ApiUserPerf200ResponseStatWorstLosses](docs/ApiUserPerf200ResponseStatWorstLosses.md)
 - [ApiUserPerf200ResponseStatWorstLossesResultsInner](docs/ApiUserPerf200ResponseStatWorstLossesResultsInner.md)
 - [ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId](docs/ApiUserPerf200ResponseStatWorstLossesResultsInnerOpId.md)
 - [ApiUserPerf200ResponseUser](docs/ApiUserPerf200ResponseUser.md)
 - [ApiUserRatingHistory200ResponseInner](docs/ApiUserRatingHistory200ResponseInner.md)
 - [ApiUsers200ResponseInner](docs/ApiUsers200ResponseInner.md)
 - [ApiUsersStatus200ResponseInner](docs/ApiUsersStatus200ResponseInner.md)
 - [BoardGameDrawAcceptParameter](docs/BoardGameDrawAcceptParameter.md)
 - [BoardGameStream200Response](docs/BoardGameStream200Response.md)
 - [BoardGameStream200ResponseOneOf](docs/BoardGameStream200ResponseOneOf.md)
 - [BoardGameStream200ResponseOneOf1](docs/BoardGameStream200ResponseOneOf1.md)
 - [BoardGameStream200ResponseOneOf2](docs/BoardGameStream200ResponseOneOf2.md)
 - [BoardGameStream200ResponseOneOfClock](docs/BoardGameStream200ResponseOneOfClock.md)
 - [BoardGameStream200ResponseOneOfPerf](docs/BoardGameStream200ResponseOneOfPerf.md)
 - [BoardGameStream200ResponseOneOfState](docs/BoardGameStream200ResponseOneOfState.md)
 - [BoardGameStream200ResponseOneOfStateExpiration](docs/BoardGameStream200ResponseOneOfStateExpiration.md)
 - [BoardGameStream200ResponseOneOfWhite](docs/BoardGameStream200ResponseOneOfWhite.md)
 - [BroadcastPlayerGet200Response](docs/BroadcastPlayerGet200Response.md)
 - [BroadcastPlayerGet200ResponseAllOfFide](docs/BroadcastPlayerGet200ResponseAllOfFide.md)
 - [BroadcastPlayerGet200ResponseAllOfFideRatings](docs/BroadcastPlayerGet200ResponseAllOfFideRatings.md)
 - [BroadcastPlayerGet200ResponseAllOfGamesInner](docs/BroadcastPlayerGet200ResponseAllOfGamesInner.md)
 - [BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent](docs/BroadcastPlayerGet200ResponseAllOfGamesInnerOpponent.md)
 - [BroadcastPlayersGet200ResponseInner](docs/BroadcastPlayersGet200ResponseInner.md)
 - [BroadcastPlayersGet200ResponseInnerAllOfPerformances](docs/BroadcastPlayersGet200ResponseInnerAllOfPerformances.md)
 - [BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs](docs/BroadcastPlayersGet200ResponseInnerAllOfRatingDiffs.md)
 - [BroadcastPlayersGet200ResponseInnerAllOfRatingsMap](docs/BroadcastPlayersGet200ResponseInnerAllOfRatingsMap.md)
 - [BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner](docs/BroadcastPlayersGet200ResponseInnerAllOfTiebreaksInner.md)
 - [BroadcastPush200Response](docs/BroadcastPush200Response.md)
 - [BroadcastPush200ResponseGamesInner](docs/BroadcastPush200ResponseGamesInner.md)
 - [BroadcastPush400Response](docs/BroadcastPush400Response.md)
 - [BroadcastRoundCreate200Response](docs/BroadcastRoundCreate200Response.md)
 - [BroadcastRoundCreate200ResponseStudy](docs/BroadcastRoundCreate200ResponseStudy.md)
 - [BroadcastRoundCreate200ResponseStudyFeatures](docs/BroadcastRoundCreate200ResponseStudyFeatures.md)
 - [BroadcastRoundGet200Response](docs/BroadcastRoundGet200Response.md)
 - [BroadcastRoundGet200ResponseGamesInner](docs/BroadcastRoundGet200ResponseGamesInner.md)
 - [BroadcastRoundGet200ResponseGamesInnerPlayersInner](docs/BroadcastRoundGet200ResponseGamesInnerPlayersInner.md)
 - [BroadcastRoundUpdate200Response](docs/BroadcastRoundUpdate200Response.md)
 - [BroadcastRoundUpdate200ResponseGamesInner](docs/BroadcastRoundUpdate200ResponseGamesInner.md)
 - [BroadcastRoundUpdate200ResponseGamesInnerPlayersInner](docs/BroadcastRoundUpdate200ResponseGamesInnerPlayersInner.md)
 - [BroadcastTeamLeaderboardGet200ResponseInner](docs/BroadcastTeamLeaderboardGet200ResponseInner.md)
 - [BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner](docs/BroadcastTeamLeaderboardGet200ResponseInnerMatchesInner.md)
 - [BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner](docs/BroadcastTeamLeaderboardGet200ResponseInnerPlayersInner.md)
 - [BroadcastsByUser200Response](docs/BroadcastsByUser200Response.md)
 - [BroadcastsByUser200ResponseCurrentPageResultsInner](docs/BroadcastsByUser200ResponseCurrentPageResultsInner.md)
 - [BroadcastsOfficial200Response](docs/BroadcastsOfficial200Response.md)
 - [BroadcastsOfficial200ResponseGroup](docs/BroadcastsOfficial200ResponseGroup.md)
 - [BroadcastsOfficial200ResponseGroupToursInner](docs/BroadcastsOfficial200ResponseGroupToursInner.md)
 - [BroadcastsOfficial200ResponsePhotosValue](docs/BroadcastsOfficial200ResponsePhotosValue.md)
 - [BroadcastsOfficial200ResponseRoundsInner](docs/BroadcastsOfficial200ResponseRoundsInner.md)
 - [BroadcastsOfficial200ResponseRoundsInnerCustomScoring](docs/BroadcastsOfficial200ResponseRoundsInnerCustomScoring.md)
 - [BroadcastsOfficial200ResponseRoundsInnerCustomScoringWhite](docs/BroadcastsOfficial200ResponseRoundsInnerCustomScoringWhite.md)
 - [BroadcastsOfficial200ResponseTour](docs/BroadcastsOfficial200ResponseTour.md)
 - [BroadcastsOfficial200ResponseTourInfo](docs/BroadcastsOfficial200ResponseTourInfo.md)
 - [BroadcastsSearch200Response](docs/BroadcastsSearch200Response.md)
 - [BroadcastsTop200Response](docs/BroadcastsTop200Response.md)
 - [BroadcastsTop200ResponseActiveInner](docs/BroadcastsTop200ResponseActiveInner.md)
 - [BroadcastsTop200ResponsePast](docs/BroadcastsTop200ResponsePast.md)
 - [BulkPairingList200ResponseInner](docs/BulkPairingList200ResponseInner.md)
 - [BulkPairingList200ResponseInnerGamesInner](docs/BulkPairingList200ResponseInnerGamesInner.md)
 - [ChallengeAi201Response](docs/ChallengeAi201Response.md)
 - [ChallengeList200Response](docs/ChallengeList200Response.md)
 - [ChallengeList200ResponseInInner](docs/ChallengeList200ResponseInInner.md)
 - [ChallengeOpen200Response](docs/ChallengeOpen200Response.md)
 - [ChallengeOpen200ResponseOpen](docs/ChallengeOpen200ResponseOpen.md)
 - [ChallengeOpen200ResponsePerf](docs/ChallengeOpen200ResponsePerf.md)
 - [Correspondence](docs/Correspondence.md)
 - [CustomPosition](docs/CustomPosition.md)
 - [Featured](docs/Featured.md)
 - [Featured1](docs/Featured1.md)
 - [Featured1PlayersInner](docs/Featured1PlayersInner.md)
 - [FeaturedPlayersInner](docs/FeaturedPlayersInner.md)
 - [Fen](docs/Fen.md)
 - [FidePlayerGet200Response](docs/FidePlayerGet200Response.md)
 - [FidePlayerRatings200Response](docs/FidePlayerRatings200Response.md)
 - [FidePlayerSearch200ResponseInner](docs/FidePlayerSearch200ResponseInner.md)
 - [GameChatGet200ResponseInner](docs/GameChatGet200ResponseInner.md)
 - [GameImport200Response](docs/GameImport200Response.md)
 - [GamePgn200Response](docs/GamePgn200Response.md)
 - [GamePgn200ResponseOneOf](docs/GamePgn200ResponseOneOf.md)
 - [GamePgn200ResponseOneOfAnalysisInner](docs/GamePgn200ResponseOneOfAnalysisInner.md)
 - [GamePgn200ResponseOneOfAnalysisInnerJudgment](docs/GamePgn200ResponseOneOfAnalysisInnerJudgment.md)
 - [GamePgn200ResponseOneOfClock](docs/GamePgn200ResponseOneOfClock.md)
 - [GamePgn200ResponseOneOfDivision](docs/GamePgn200ResponseOneOfDivision.md)
 - [GamePgn200ResponseOneOfOpening](docs/GamePgn200ResponseOneOfOpening.md)
 - [GamePgn200ResponseOneOfPlayers](docs/GamePgn200ResponseOneOfPlayers.md)
 - [GamePgn200ResponseOneOfPlayersWhite](docs/GamePgn200ResponseOneOfPlayersWhite.md)
 - [GamePgn200ResponseOneOfPlayersWhiteAnalysis](docs/GamePgn200ResponseOneOfPlayersWhiteAnalysis.md)
 - [GamesByIds200ResponseInner](docs/GamesByIds200ResponseInner.md)
 - [GamesByUsers200ResponseInner](docs/GamesByUsers200ResponseInner.md)
 - [GamesByUsers200ResponseInnerClock](docs/GamesByUsers200ResponseInnerClock.md)
 - [GamesByUsers200ResponseInnerPlayers](docs/GamesByUsers200ResponseInnerPlayers.md)
 - [GamesByUsers200ResponseInnerPlayersWhite](docs/GamesByUsers200ResponseInnerPlayersWhite.md)
 - [MateVariation](docs/MateVariation.md)
 - [NonMateVariation](docs/NonMateVariation.md)
 - [OpeningExplorerLichess200Response](docs/OpeningExplorerLichess200Response.md)
 - [OpeningExplorerLichess200ResponseHistoryInner](docs/OpeningExplorerLichess200ResponseHistoryInner.md)
 - [OpeningExplorerLichess200ResponseMovesInner](docs/OpeningExplorerLichess200ResponseMovesInner.md)
 - [OpeningExplorerLichess200ResponseMovesInnerGame](docs/OpeningExplorerLichess200ResponseMovesInnerGame.md)
 - [OpeningExplorerLichess200ResponseTopGamesInner](docs/OpeningExplorerLichess200ResponseTopGamesInner.md)
 - [OpeningExplorerMaster200Response](docs/OpeningExplorerMaster200Response.md)
 - [OpeningExplorerMaster200ResponseMovesInner](docs/OpeningExplorerMaster200ResponseMovesInner.md)
 - [OpeningExplorerMaster200ResponseMovesInnerGame](docs/OpeningExplorerMaster200ResponseMovesInnerGame.md)
 - [OpeningExplorerMaster200ResponseMovesInnerGameWhite](docs/OpeningExplorerMaster200ResponseMovesInnerGameWhite.md)
 - [OpeningExplorerMaster200ResponseOpening](docs/OpeningExplorerMaster200ResponseOpening.md)
 - [OpeningExplorerMaster200ResponseTopGamesInner](docs/OpeningExplorerMaster200ResponseTopGamesInner.md)
 - [OpeningExplorerPlayer200Response](docs/OpeningExplorerPlayer200Response.md)
 - [OpeningExplorerPlayer200ResponseMovesInner](docs/OpeningExplorerPlayer200ResponseMovesInner.md)
 - [OpeningExplorerPlayer200ResponseMovesInnerGame](docs/OpeningExplorerPlayer200ResponseMovesInnerGame.md)
 - [OpeningExplorerPlayer200ResponseRecentGamesInner](docs/OpeningExplorerPlayer200ResponseRecentGamesInner.md)
 - [Player](docs/Player.md)
 - [Player200Response](docs/Player200Response.md)
 - [Player200ResponseBlitzInner](docs/Player200ResponseBlitzInner.md)
 - [Player200ResponseBulletInner](docs/Player200ResponseBulletInner.md)
 - [Player200ResponseBulletInnerPerfsValue](docs/Player200ResponseBulletInnerPerfsValue.md)
 - [PlayerTopNbPerfType200Response](docs/PlayerTopNbPerfType200Response.md)
 - [RacerGet200Response](docs/RacerGet200Response.md)
 - [RacerGet200ResponsePlayersInner](docs/RacerGet200ResponsePlayersInner.md)
 - [RacerGet200ResponsePuzzlesInner](docs/RacerGet200ResponsePuzzlesInner.md)
 - [RacerGet404Response](docs/RacerGet404Response.md)
 - [RacerPost200Response](docs/RacerPost200Response.md)
 - [ReadNote200ResponseInner](docs/ReadNote200ResponseInner.md)
 - [RealTime](docs/RealTime.md)
 - [ResultsBySwiss200Response](docs/ResultsBySwiss200Response.md)
 - [ResultsByTournament200Response](docs/ResultsByTournament200Response.md)
 - [StreamGame200ResponseInner](docs/StreamGame200ResponseInner.md)
 - [StreamGame200ResponseInnerOneOf](docs/StreamGame200ResponseInnerOneOf.md)
 - [StreamGame200ResponseInnerOneOf1](docs/StreamGame200ResponseInnerOneOf1.md)
 - [StreamGame200ResponseInnerOneOfStatus](docs/StreamGame200ResponseInnerOneOfStatus.md)
 - [StreamGame429Response](docs/StreamGame429Response.md)
 - [StreamerLive200ResponseInner](docs/StreamerLive200ResponseInner.md)
 - [StreamerLive200ResponseInnerAllOfStream](docs/StreamerLive200ResponseInnerAllOfStream.md)
 - [StreamerLive200ResponseInnerAllOfStreamer](docs/StreamerLive200ResponseInnerAllOfStreamer.md)
 - [StudyListMetadata200Response](docs/StudyListMetadata200Response.md)
 - [TablebaseStandard200Response](docs/TablebaseStandard200Response.md)
 - [TablebaseStandard200ResponseMovesInner](docs/TablebaseStandard200ResponseMovesInner.md)
 - [TeamAll200Response](docs/TeamAll200Response.md)
 - [TeamIdUsers200Response](docs/TeamIdUsers200Response.md)
 - [TeamRequests200ResponseInner](docs/TeamRequests200ResponseInner.md)
 - [TeamRequests200ResponseInnerRequest](docs/TeamRequests200ResponseInnerRequest.md)
 - [TeamShow200Response](docs/TeamShow200Response.md)
 - [TeamsByTournament200Response](docs/TeamsByTournament200Response.md)
 - [TeamsByTournament200ResponseTeamsInner](docs/TeamsByTournament200ResponseTeamsInner.md)
 - [TeamsByTournament200ResponseTeamsInnerPlayersInner](docs/TeamsByTournament200ResponseTeamsInnerPlayersInner.md)
 - [Thematic](docs/Thematic.md)
 - [Timeline200Response](docs/Timeline200Response.md)
 - [Timeline200ResponseEntriesInner](docs/Timeline200ResponseEntriesInner.md)
 - [Timeline200ResponseEntriesInnerAnyOf](docs/Timeline200ResponseEntriesInnerAnyOf.md)
 - [Timeline200ResponseEntriesInnerAnyOf1](docs/Timeline200ResponseEntriesInnerAnyOf1.md)
 - [Timeline200ResponseEntriesInnerAnyOf10](docs/Timeline200ResponseEntriesInnerAnyOf10.md)
 - [Timeline200ResponseEntriesInnerAnyOf10Data](docs/Timeline200ResponseEntriesInnerAnyOf10Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf11](docs/Timeline200ResponseEntriesInnerAnyOf11.md)
 - [Timeline200ResponseEntriesInnerAnyOf11Data](docs/Timeline200ResponseEntriesInnerAnyOf11Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf12](docs/Timeline200ResponseEntriesInnerAnyOf12.md)
 - [Timeline200ResponseEntriesInnerAnyOf12Data](docs/Timeline200ResponseEntriesInnerAnyOf12Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf13](docs/Timeline200ResponseEntriesInnerAnyOf13.md)
 - [Timeline200ResponseEntriesInnerAnyOf13Data](docs/Timeline200ResponseEntriesInnerAnyOf13Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf1Data](docs/Timeline200ResponseEntriesInnerAnyOf1Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf2](docs/Timeline200ResponseEntriesInnerAnyOf2.md)
 - [Timeline200ResponseEntriesInnerAnyOf3](docs/Timeline200ResponseEntriesInnerAnyOf3.md)
 - [Timeline200ResponseEntriesInnerAnyOf3Data](docs/Timeline200ResponseEntriesInnerAnyOf3Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf4](docs/Timeline200ResponseEntriesInnerAnyOf4.md)
 - [Timeline200ResponseEntriesInnerAnyOf4Data](docs/Timeline200ResponseEntriesInnerAnyOf4Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf5](docs/Timeline200ResponseEntriesInnerAnyOf5.md)
 - [Timeline200ResponseEntriesInnerAnyOf5Data](docs/Timeline200ResponseEntriesInnerAnyOf5Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf6](docs/Timeline200ResponseEntriesInnerAnyOf6.md)
 - [Timeline200ResponseEntriesInnerAnyOf6Data](docs/Timeline200ResponseEntriesInnerAnyOf6Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf7](docs/Timeline200ResponseEntriesInnerAnyOf7.md)
 - [Timeline200ResponseEntriesInnerAnyOf7Data](docs/Timeline200ResponseEntriesInnerAnyOf7Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf8](docs/Timeline200ResponseEntriesInnerAnyOf8.md)
 - [Timeline200ResponseEntriesInnerAnyOf8Data](docs/Timeline200ResponseEntriesInnerAnyOf8Data.md)
 - [Timeline200ResponseEntriesInnerAnyOf9](docs/Timeline200ResponseEntriesInnerAnyOf9.md)
 - [Timeline200ResponseEntriesInnerAnyOf9Data](docs/Timeline200ResponseEntriesInnerAnyOf9Data.md)
 - [Timeline200ResponseEntriesInnerAnyOfData](docs/Timeline200ResponseEntriesInnerAnyOfData.md)
 - [TokenTest200ResponseValue](docs/TokenTest200ResponseValue.md)
 - [TokenTest200ResponseValueOneOf](docs/TokenTest200ResponseValueOneOf.md)
 - [Tournament200Response](docs/Tournament200Response.md)
 - [Tournament200ResponsePodiumInner](docs/Tournament200ResponsePodiumInner.md)
 - [Tournament200ResponseStanding](docs/Tournament200ResponseStanding.md)
 - [Tournament200ResponseStandingPlayersInner](docs/Tournament200ResponseStandingPlayersInner.md)
 - [TvChannelFeed200Response](docs/TvChannelFeed200Response.md)
 - [TvChannels200Response](docs/TvChannels200Response.md)
 - [TvChannels200ResponseBlitz](docs/TvChannels200ResponseBlitz.md)
 - [TvChannels200ResponseBot](docs/TvChannels200ResponseBot.md)
 - [TvFeed200Response](docs/TvFeed200Response.md)
 - [Unlimited](docs/Unlimited.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### OAuth2


- **Type**: OAuth
- **Flow**: accessCode
- **Authorization URL**: https://lichess.org/oauth
- **Scopes**: 
 - **preference:read**: Read your preferences
 - **preference:write**: Write your preferences
 - **email:read**: Read your email address
 - **engine:read**: Read your external engines
 - **engine:write**: Create, update, delete your external engines
 - **challenge:read**: Read incoming challenges
 - **challenge:write**: Create, accept, decline challenges
 - **challenge:bulk**: Create, delete, query bulk pairings
 - **study:read**: Read private studies and broadcasts
 - **study:write**: Create, update, delete studies and broadcasts
 - **tournament:write**: Create tournaments
 - **racer:write**: Create and join puzzle races
 - **puzzle:read**: Read puzzle activity
 - **puzzle:write**: Write puzzle activity
 - **team:read**: Read private team information
 - **team:write**: Join, leave teams
 - **team:lead**: Manage teams (kick members, send PMs)
 - **follow:read**: Read followed players
 - **follow:write**: Follow and unfollow other players
 - **msg:write**: Send private messages to other players
 - **board:play**: Play with the Board API
 - **bot:play**: Play with the Bot API. Only for [Bot accounts](#tag/bot/POST/api/bot/account/upgrade)
 - **web:mod**: Use moderator tools (within the bounds of your permissions)

Example

```go
auth := context.WithValue(context.Background(), lichess.ContextAccessToken, "ACCESSTOKENSTRING")
r, err := client.Service.Operation(auth, args)
```

Or via OAuth2 module to automatically refresh tokens and perform user authentication.

```go
import "golang.org/x/oauth2"

/* Perform OAuth2 round trip request and obtain a token */

tokenSource := oauth2cfg.TokenSource(createContext(httpClient), &token)
auth := context.WithValue(oauth2.NoContext, lichess.ContextOAuth2, tokenSource)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

contact@lichess.org

